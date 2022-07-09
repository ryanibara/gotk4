package girgen

import (
	"fmt"
	"go/format"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
	"github.com/diamondburned/gotk4/gir/girgen/types"
	"github.com/pkg/errors"
)

// FileGenerator describes any file that can be generated.
type FileGenerator interface {
	generators.FileGenerator
	generators.FileWriter
	Generate() ([]byte, error)
	IsEmpty() bool
}

// CFileGenerator creates a new C file generator.
type CFileGenerator struct {
	*NamespaceGenerator
	pen    *pen.PaperBuffer
	header file.Header
}

// NewCFileGenerator creates as new C file generator.
func NewCFileGenerator(n *NamespaceGenerator) *CFileGenerator {
	return &CFileGenerator{
		NamespaceGenerator: n,
		pen:                pen.NewPaperBufferSize(10 * 1024), // 10KB
	}
}

// Header returns the C file generator's headers.
func (f *CFileGenerator) Header() *file.Header { return &f.header }

// Pen returns the C file generator's writer.
func (f *CFileGenerator) Pen() *pen.Pen { return &f.pen.Pen }

// IsEmpty returns true if the file is empty.
func (f *CFileGenerator) IsEmpty() bool {
	return f == nil || (f.pen.Len() == 0 && len(f.header.Callbacks) == 0)
}

// Generate implements FileGenerator.
func (f *CFileGenerator) Generate() ([]byte, error) {
	switch f.LinkMode() {
	case types.DynamicLinkMode:
		f.pen.Words("#include <stdlib.h>")
		if incls := f.CIncludes(); len(incls) > 0 {
			for _, incl := range incls {
				f.pen.Linef("#include <%s>", incl)
			}
		}
	case types.RuntimeLinkMode:
		f.pen.Words("#include <stdlib.h>")
		f.pen.Words("#include <glib.h>")
	}

	if len(f.header.Callbacks) > 0 {
		f.pen.EmptyLine()
		for _, callback := range f.header.SortedCallbackHeaders() {
			f.pen.Words(callback)
		}
	}

	return f.pen.Bytes(), nil
}

// CIncludes returns the list of C includes at the top of the file.
func (f *CFileGenerator) CIncludes() []string {
	return namespaceCIncludes(f.current, &f.header)
}

// GoFileGenerator is a file generator.
type GoFileGenerator struct {
	*NamespaceGenerator
	BuildTags []string // go:build lines, joined by AND (&&)
	Packages  []string // extra

	pen    *pen.PaperBuffer
	header file.Header

	name   string
	isRoot bool
}

var (
	_ FileGenerator            = (*GoFileGenerator)(nil)
	_ types.FileGenerator      = (*GoFileGenerator)(nil)
	_ generators.FileGenerator = (*GoFileGenerator)(nil)
	_ generators.FileWriter    = (*GoFileGenerator)(nil)
)

// NewGoFileGenerator creates a new empty GoFileGenerator instance.
func NewGoFileGenerator(n *NamespaceGenerator, name string, isRoot bool) *GoFileGenerator {
	return &GoFileGenerator{
		NamespaceGenerator: n,

		pen:    pen.NewPaperBufferSize(10 * 1024), // 10KB
		name:   name,
		isRoot: isRoot,
	}
}

// Name returns the current file's name.
func (f *GoFileGenerator) Name() string {
	return f.name
}

// Header returns the current file's header.
func (f *GoFileGenerator) Header() *file.Header {
	return &f.header
}

// Pen returns the current file's writing pen.
func (f *GoFileGenerator) Pen() *pen.Pen {
	return &f.pen.Pen
}

func (f *GoFileGenerator) Logln(lvl logger.Level, v ...interface{}) {
	p := fmt.Sprintf("file %s", f.name)
	f.NamespaceGenerator.Logln(lvl, logger.Prefix(v, p)...)
}

// IsEmpty returns true if the file is empty.
func (f *GoFileGenerator) IsEmpty() bool {
	return !f.isRoot && f.pen.Len() == 0
}

// Generate generates the final file content, completed with gofmt.
func (f *GoFileGenerator) Generate() ([]byte, error) {
	if len(f.header.Marshalers) > 0 {
		// Import coreglib for the RegisterMarshal function.
		f.header.NeedsExternGLib()
	}

	if f.LinkMode() == types.RuntimeLinkMode {
		f.header.ImportCore("girepository")
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

	switch f.LinkMode() {
	case types.DynamicLinkMode:
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
	case types.RuntimeLinkMode:
		fpen.Words("// #cgo pkg-config: gobject-2.0")
		fpen.Words("// #include <stdlib.h>")
		fpen.Words("// #include <glib.h>")
	}

	if len(f.header.Callbacks) > 0 {
		for _, callback := range f.header.SortedCallbackHeaders() {
			fpen.Words(makeComment(callback))
		}
	}

	fpen.Words(`import "C"`)
	fpen.EmptyLine()

	for _, m := range f.header.Marshalers {
		// Contrary to popular beliefs, RegisterGValueMarshalers is actually
		// thread-safe.
		fpen.Linef("// GType%s returns the GType for the type %[1]s.", m.GoTypeName)
		fpen.Linef("//")
		fpen.Linef("// This function has the side effect of registering a GValue marshaler")
		fpen.Linef("// globally. Use this if you need that for any reason. The function is")
		fpen.Linef("// concurrently safe to use.")
		fpen.Linef("func GType%s() coreglib.Type {", m.GoTypeName)
		fpen.Linef("  gtype := %s", m.GLibType())
		fpen.Linef("  coreglib.RegisterGValueMarshaler(gtype, marshal%s)", m.GoTypeName)
		fpen.Linef("  return gtype")
		fpen.Linef("}")
		fpen.EmptyLine()
	}

	if f.isRoot && f.LinkMode() == types.RuntimeLinkMode {
		namespace := f.Namespace().Namespace

		fpen.Words("func init() {")
		fpen.Linef("  girepository.Require(%q, %q)", namespace.Name, namespace.Version)
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
func (f *GoFileGenerator) Pkgconfig() []string {
	return append(f.Packages, f.NamespaceGenerator.Pkgconfig()...)
}

// CIncludes returns this file's sorted C includes, including the repository's C
// includes.
func (f *GoFileGenerator) CIncludes() []string {
	return namespaceCIncludes(f.current, &f.header)
}

func namespaceCIncludes(n *gir.NamespaceFindResult, h *file.Header) []string {
	extraIncludes := h.SortedCIncludes()

	includes := make([]string, 0, len(extraIncludes)+len(n.Repository.CIncludes))
	for _, incl := range n.Repository.CIncludes {
		includes = append(includes, incl.Name)
	}
	includes = append(includes, extraIncludes...)

	sort.Strings(includes)
	return includes
}
