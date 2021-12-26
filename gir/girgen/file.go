package girgen

import (
	"fmt"
	"go/format"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
	"github.com/diamondburned/gotk4/gir/girgen/types"
	"github.com/pkg/errors"
)

// FileGenerator is a file generator.
type FileGenerator struct {
	*NamespaceGenerator
	BuildTags []string // go:build lines, joined by AND (&&)
	Packages  []string // extra

	pen    *pen.PaperBuffer
	header file.Header

	name   string
	isRoot bool
}

var (
	_ types.FileGenerator      = (*FileGenerator)(nil)
	_ generators.FileGenerator = (*FileGenerator)(nil)
	_ generators.FileWriter    = (*FileGenerator)(nil)
)

// NewFileGenerator creates a new empty FileGenerator instance.
func NewFileGenerator(n *NamespaceGenerator, name string, isRoot bool) *FileGenerator {
	return &FileGenerator{
		NamespaceGenerator: n,

		pen:    pen.NewPaperBufferSize(10 * 1024), // 10KB
		name:   name,
		isRoot: isRoot,
	}
}

// Header returns the current file's header.
func (f *FileGenerator) Header() *file.Header {
	return &f.header
}

// Pen returns the current file's writing pen.
func (f *FileGenerator) Pen() *pen.Pen {
	return &f.pen.Pen
}

func (f *FileGenerator) Logln(lvl logger.Level, v ...interface{}) {
	p := fmt.Sprintf("file %s", f.name)
	f.NamespaceGenerator.Logln(lvl, logger.Prefix(v, p)...)
}

// IsEmpty returns true if the file is empty.
func (f *FileGenerator) IsEmpty() bool {
	return !f.isRoot && f.pen.Len() == 0
}

// Generate generates the final file content, completed with gofmt.
func (f *FileGenerator) Generate() ([]byte, error) {
	if len(f.header.Marshalers) > 0 {
		// Import externglib for the RegisterMarshal function.
		f.header.NeedsExternGLib()
	}

	if f.header.CallbackDelete {
		// C headers are per-file.
		f.header.AddCallbackHeader("extern void callbackDelete(gpointer);")
		// Ensure that gbox is imported, so we have the exported callbackDelete
		// function.
		f.header.DashImport(file.ImportCore("gbox"))
	}

	fpen := pen.NewPaperBufferSize(4096 + f.pen.Len()) // 4KB + pen
	fpen.Words("// Code generated by girgen. DO NOT EDIT.")
	fpen.EmptyLine()

	if len(f.BuildTags) > 0 {
		fpen.Words("//go:build", strings.Join(f.BuildTags, " && "))
		fpen.EmptyLine()
	}

	fpen.Words("package", f.PkgName)
	fpen.EmptyLine()

	if len(f.header.Imports) > 0 {
		builtins := make([]string, 0, len(f.header.Imports))
		externs := make([]string, 0, len(f.header.Imports))

		for path, alias := range f.header.Imports {
			// If path matches the import path, then this is a bug elsewhere.
			if path == f.PkgPath {
				// f.Logln(logger.Error, "importing self")
				// return nil, errors.New("cyclical import on self")
				continue
			}

			// Path containing a dot probably indicates it has a domain name.
			if !strings.Contains(path, ".") {
				builtins = append(builtins, makeImport(path, alias))
			} else {
				externs = append(externs, makeImport(path, alias))
			}
		}

		sort.Strings(builtins)
		sort.Strings(externs)

		fpen.Words("import (")
		fpen.Lines(builtins)
		if len(builtins) > 0 && len(externs) > 0 {
			fpen.EmptyLine()
		}
		fpen.Lines(externs)
		fpen.Words(")")
		fpen.EmptyLine()
	}

	if f.isRoot {
		fpen.Words("// #cgo pkg-config:", f.Pkgconfig())
		fpen.Words("// #cgo CFLAGS: -Wno-deprecated-declarations")
	}

	fpen.Words("// #include <stdlib.h>")
	if incls := f.CIncludes(); len(incls) > 0 {
		for _, incl := range incls {
			fpen.Linef("// #include <%s>", incl)
		}
	}

	if len(f.header.Callbacks) > 0 {
		for _, callback := range f.header.SortedCallbackHeaders() {
			fpen.Words(makeComment(callback))
		}
	}

	fpen.Words(`import "C"`)
	fpen.EmptyLine()

	if len(f.header.Marshalers) > 0 {
		fpen.Words("func init() {")
		fpen.Words("  externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{")
		for _, marshaler := range f.header.Marshalers {
			fpen.Words(marshaler)
		}
		fpen.Words("  })")
		fpen.Words("}")
	}

	fpen.Write(f.pen.Bytes())

	b, err := format.Source(fpen.Bytes())
	if err != nil {
		f.Logln(logger.Error, "fmt error:", err)
		return fpen.Bytes(), errors.Wrap(err, "fmt error")
	}

	return b, nil
}

func makeComment(block string) string {
	lines := strings.Split(block, "\n")
	for i := range lines {
		lines[i] = "// " + lines[i]
	}
	return strings.Join(lines, "\n")
}

func makeImport(importPath, alias string) string {
	pathBase := path.Base(importPath)

	// Check if the base is a version part.
	if strings.HasPrefix(pathBase, "v") {
		_, err := strconv.Atoi(strings.TrimPrefix(pathBase, "v"))
		if err == nil {
			// Valid version part. Trim it.
			pathBase = path.Base(path.Dir(importPath))
		}
	}

	if alias == "" || alias == pathBase {
		return strconv.Quote(importPath)
	}

	// Only use the import alias if it's provided and does not match the base
	// name of the import path for idiomaticity.
	return alias + " " + strconv.Quote(importPath)
}

// Pkgconfig returns the current repository's pkg-config names.
func (f *FileGenerator) Pkgconfig() []string {
	return append(f.Packages, f.NamespaceGenerator.Pkgconfig()...)
}

// CIncludes returns this file's sorted C includes, including the repository's C
// includes.
func (f *FileGenerator) CIncludes() []string {
	extraIncludes := f.header.SortedCIncludes()

	includes := make([]string, 0, len(extraIncludes)+len(f.current.Repository.CIncludes))
	for _, incl := range f.current.Repository.CIncludes {
		includes = append(includes, incl.Name)
	}
	includes = append(includes, extraIncludes...)

	sort.Strings(includes)
	return includes
}
