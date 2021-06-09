package girgen

import (
	"bytes"
	"fmt"
	"go/format"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
	"github.com/pkg/errors"
)

// FileGenerator contains values to fully construct a generated file.
type FileGenerator struct {
	parent *NamespaceGenerator
	pen    *pen.PaperString // body

	name       string
	includes   []string // C include
	marshalers []string

	imports  map[string]string // optional alias value
	imported []string          // keep track of side effects

	// Keep track of local callback declarations so we can reference them in the
	// file.
	callbacks      map[string]struct{}
	callbackDelete bool

	// inserted keeps track of what was inserted once.
	inserted struct {
		IncludeGObject bool
		IncludeStdbool bool
	}
}

// NewFileGenerator creates a new FileGenerator.
func NewFileGenerator(parent *NamespaceGenerator, goFile string) *FileGenerator {
	fg := FileGenerator{
		parent: parent,
		pen:    pen.NewPaperStringSize(2 * 1024 * 1024), // 2MB

		name:      goFile,
		imports:   map[string]string{},
		callbacks: map[string]struct{}{},

		// overallocate
		marshalers: make([]string, 0, 0+
			len(parent.current.Namespace.Enums)+
			len(parent.current.Namespace.Bitfields)+
			len(parent.current.Namespace.Records)+
			len(parent.current.Namespace.Classes)+
			len(parent.current.Namespace.Interfaces),
		),
	}

	return &fg
}

func replaceExt(name, ext string) string {
	parts := strings.Split(name, ".")

	switch len(parts) {
	case 0:
		return ""
	case 1:
		return parts[0] + ext
	default:
		return strings.Join(parts[:len(parts)-1], ".") + ext
	}
}

func (fg *FileGenerator) generate() ([]byte, error) {
	if len(fg.marshalers) > 0 {
		fg.addImportAlias("github.com/gotk3/gotk3/glib", "externglib")
	}

	if fg.callbackDelete {
		fg.addImport("github.com/diamondburned/gotk4/internal/box")
	}

	var out bytes.Buffer
	// Preallocate 10KB + existing buffers.
	out.Grow(10*1024 + fg.pen.Len())

	pen := pen.NewPen(&out)
	pen.Words("// Code generated by girgen. DO NOT EDIT.")
	pen.EmptyLine()

	pen.Words("package", fg.parent.PackageName())
	pen.EmptyLine()

	if len(fg.imports) > 0 {
		builtin := make([]string, len(fg.imports))
		externs := make([]string, len(fg.imports))

		for path, alias := range fg.imports {
			if !strings.Contains(path, "/") {
				builtin = append(builtin, makeImport(path, alias))
			} else {
				externs = append(externs, makeImport(path, alias))
			}
		}

		sort.Strings(builtin)
		sort.Strings(externs)

		pen.Words("import (")

		for _, str := range builtin {
			pen.Words(str)
		}
		if len(builtin) > 0 && len(externs) > 0 {
			pen.EmptyLine()
		}
		for _, str := range externs {
			pen.Words(str)
		}

		pen.Line(")")
		pen.EmptyLine()
		pen.EmptyLine()
	}

	pen.Words(append([]string{"// #cgo pkg-config:"}, fg.parent.pkgconfig()...)...)
	pen.Words("// #cgo CFLAGS: -Wno-deprecated-declarations")

	for _, cIncl := range fg.includes {
		pen.Linef("// #include <%s>", cIncl)
	}
	for _, cIncl := range fg.parent.current.Repository.CIncludes {
		pen.Linef("// #include <%s>", cIncl.Name)
	}

	if fg.callbackDelete {
		pen.Words("//")
		pen.Words("// extern void callbackDelete(gpointer);")
	}

	if len(fg.callbacks) > 0 {
		pen.Words("//")
		for _, callback := range fg.callbackHeaders() {
			pen.Words("//", callback)
		}
	}

	pen.Words(`import "C"`)
	pen.EmptyLine()

	if len(fg.marshalers) > 0 {
		pen.Words("func init() {")
		pen.Words("  externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{")

		for _, marshaler := range fg.marshalers {
			pen.Words("      " + marshaler)
		}

		pen.Words("  })")
		pen.Words("}")
		pen.EmptyLine()
	}

	if fg.callbackDelete {
		pen.Words("//export callbackDelete")
		pen.Words("func callbackDelete(ptr C.gpointer) {")
		pen.Words("  box.Delete(box.Callback, uintptr(ptr))")
		pen.Words("}")
		pen.EmptyLine()
	}

	pen.WriteString(fg.pen.String())

	formatted, err := format.Source(out.Bytes())
	if err != nil {
		fg.Logln(LogError, "failed to fmt file", fg.name)
		return out.Bytes(), errors.Wrap(err, "fmt file "+fg.name)
	}

	return formatted, nil
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

func (fg *FileGenerator) addMarshaler(glibGetType, goName string) {
	fg.marshalers = append(fg.marshalers, fmt.Sprintf(
		`{T: externglib.Type(C.%s()), F: marshal%s},`, glibGetType, goName,
	))
}

func (fg *FileGenerator) needsStdbool() {
	if !fg.inserted.IncludeStdbool {
		fg.inserted.IncludeStdbool = true

		fg.includes = append(fg.includes, "stdbool.h")
	}
}

func (fg *FileGenerator) needsGLibObject() {
	if !fg.inserted.IncludeGObject {
		fg.inserted.IncludeGObject = true

		// Need this for g_value_get_boxed.
		fg.includes = append(fg.includes, "glib-object.h")
	}
}

func (fg *FileGenerator) needsCallbackDelete() {
	fg.callbackDelete = true
}

func (fg *FileGenerator) callbackHeaders() []string {
	headers := make([]string, 0, len(fg.callbacks))
	for callback := range fg.callbacks {
		headers = append(headers, callback)
	}

	sort.Strings(headers)
	return headers
}

// addCallbackHeader adds the given callback's C function header into the
// registry.
func (fg *FileGenerator) addCallbackHeader(header string) {
	fg.callbacks[header] = struct{}{}
}

func (fg *FileGenerator) addImport(pkgPath string) {
	fg.addImportAlias(pkgPath, "")
}

func (fg *FileGenerator) addImportAlias(pkgPath, alias string) {
	_, ok := fg.imports[pkgPath]
	if ok {
		return
	}

	fg.imports[pkgPath] = alias
	fg.imported = append(fg.imported, pkgPath)
}

// wipeImportStack wipes the internal import stack. This is useful if the caller
// wants to keep track of added imports to undo them.
func (fg *FileGenerator) wipeImportStack() {
	fg.imported = fg.imported[:0]
}

// undoImports undoes the added imports by reading back from the stack. The
// caller should call wipeImportStack before this. The import stack is wiped
// once done.
func (fg *FileGenerator) undoImports() {
	for _, pkgPath := range fg.imported {
		delete(fg.imports, pkgPath)
	}
	fg.wipeImportStack()
}

func (fg *FileGenerator) addResolvedImport(resolved *ResolvedType) {
	if resolved != nil && resolved.Import != "" && resolved.Import != fg.parent.pkgPath {
		fg.addImportAlias(resolved.Import, resolved.Package)
	}
}

// addGLibImport adds the gotk3/glib import.
func (fg *FileGenerator) addGLibImport() {
	fg.addResolvedImport(externGLibType("", gir.Type{}, ""))
}

// Namespace returns the generator's namespace that includes the repository it's
// in.
func (fg *FileGenerator) Namespace() *gir.NamespaceFindResult {
	return fg.parent.current
}

func (fg *FileGenerator) Logln(level LogLevel, v ...interface{}) {
	prefix := []interface{}{fg.name + ":"}
	prefix = append(prefix, v...)

	fg.parent.Logln(level, prefix...)
}

func (fg *FileGenerator) warnUnknownType(typ string) {
	fg.Logln(LogUnknown, strconv.Quote(typ))
}