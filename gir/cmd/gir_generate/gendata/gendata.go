// Package gendata contains data used to generate GTK4 bindings for Go. It
// exists primarily to be used externally.
package gendata

import (
	"errors"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	. "github.com/diamondburned/gotk4/gir/girgen/types"
	. "github.com/diamondburned/gotk4/gir/girgen/types/typeconv"
)

type Package struct {
	PkgName    string   // pkg-config name
	Namespaces []string // refer to ./cmd/gir_namespaces
}

// HasNamespace returns true if the package allows all namespaces or has the
// given namespace in the list.
func (pkg *Package) HasNamespace(n *gir.Namespace) bool {
	if pkg.Namespaces == nil {
		return true
	}

	namespace := gir.VersionedNamespace(n)
	for _, name := range pkg.Namespaces {
		if name == namespace {
			return true
		}
	}

	return false
}

// PkgExceptions contains a list of file names that won't be deleted off of
// pkg/.
var PkgExceptions = []string{
	"core",
	"cairo",
	"go.mod",
	"go.sum",
	"LICENSE",
}

// PkgGenerated contains a list of file names that are packages generated using
// the given Packages list. It is manually updated.
var PkgGenerated = []string{
	"atk",
	"gdk",
	"gdkpixbuf",
	"gdkpixdata",
	"gdkwayland",
	"gdkx11",
	"gio",
	"glib",
	"gobject",
	"graphene",
	"gsk",
	"gtk",
	"pango",
	"pangocairo",
}

// GenerateExceptions contains the keys of the underneath ImportOverrides map.
var GenerateExceptions = []string{
	"cairo-1",
}

// ImportOverrides is the list of imports to defer to another library, usually
// because it's tedious or impossible to generate.
//
// Not included: externglib (gotk3/gotk3/glib).
var ImportOverrides = map[string]string{}

// Packages lists pkg-config packages and optionally the namespaces to be
// generated. If the list of namespaces is nil, then everything is generated.
var Packages = []Package{
	{"gobject-introspection-1.0", []string{
		"GLib-2",
		"GObject-2",
		"Gio-2",
		"cairo-1",
	}},
	{"gdk-pixbuf-2.0", nil},
	{"graphene-1.0", nil},
	{"atk", nil},
	{"pango", []string{
		"Pango-1",
		"PangoCairo-1",
	}},
	{"gtk4", nil},     // includes Gdk
	{"gtk+-3.0", nil}, // includes Gdk
}

// Preprocessors defines a list of preprocessors that the main generator will
// use. It's mostly used for renaming colliding types/identifiers.
var Preprocessors = []Preprocessor{
	// Collision due to case conversions.
	TypeRenamer("GLib-2.file_test", "test_file"),
	// This collides with Native().
	TypeRenamer("Gtk-4.Native", "NativeSurface"),
	// These collide with structs of the same names.
	RenameEnumMembers("Pango-1.AttrType", "ATTR_(.*)", "ATTR_TYPE_$1"),
	RenameEnumMembers("Gsk-4.RenderNodeType", ".*", "${0}_TYPE"),
	RenameEnumMembers("Gdk-3.EventType", ".*", "${0}_TYPE"),
	// See #28.
	RemoveCIncludes("Gio-2.0.gir", "gio/gdesktopappinfo.h"),

	// Length and Value are invalid in Go. We manually handle them in GLibLogs.
	RemoveRecordFields("GLib-2.LogField", "length", "value"),

	ModifyParamDirections("Gio-2.InputStream.read", map[string]string{
		"buffer": "in",
		"count":  "in",
	}),
	ModifyParamDirections("Gio-2.InputStream.read_async", map[string]string{
		"buffer": "in",
		"count":  "in",
	}),
	ModifyParamDirections("Gio-2.InputStream.read_all", map[string]string{
		"buffer": "in",
		"count":  "in",
	}),
	ModifyParamDirections("Gio-2.InputStream.read_all_async", map[string]string{
		"buffer": "in",
		"count":  "in",
	}),
	ModifyParamDirections("Gio-2.Socket.receive", map[string]string{
		"buffer": "in",
		"size":   "in",
	}),
	ModifyParamDirections("Gio-2.Socket.receive_from", map[string]string{
		"buffer": "in",
		"size":   "in",
	}),
	ModifyParamDirections("Gio-2.Socket.receive_with_blocking", map[string]string{
		"buffer": "in",
		"size":   "in",
	}),
	ModifyParamDirections("Gio-2.DBusInterfaceGetPropertyFunc", map[string]string{
		"error": "out",
	}),
	ModifyCallable("Gdk-4.Clipboard.read_async", func(c *gir.CallableAttrs) {
		// Fix this parameter's type not being a proper array.
		p := FindParameter(c, "mime_types")
		p.Array = &gir.Array{
			CType: "const char**",
			Type:  &gir.Type{Name: "utf8"},
		}
	}),
}

var ConversionProcessors = []ConversionProcessor{
	ProcessCallback("Gio-2.AsyncReadyCallback", func(conv *Converter) {
		// Don't include the first parameter in Go.
		conv.Results[0].Skip = true
	}),
}

// Filters defines a list of GIR types to be filtered. The map key is the
// namespace, and the values are list of names.
var Filters = []FilterMatcher{
	AbsoluteFilter("C.cairo_image_surface_create"),

	// These are not in gotk3/cairo.
	AbsoluteFilter("cairo.ScaledFont"),
	AbsoluteFilter("cairo.FontType"),

	// Broadway is not included, so we don't generate code for it.
	FileFilter("gsk/broadway/gskbroadwayrenderer.h"),
	// Output buffer parameter is not actually array.
	AbsoluteFilter("GLib.unichar_to_utf8"),
	// This is useless.
	AbsoluteFilter("GLib.nullify_pointer"),
	// Requires special header, is optional function.
	AbsoluteFilter("GLib.unix_error_quark"),
	AbsoluteFilter("Gio.networking_init"),
	// Not an array type but expects an array.
	AbsoluteFilter("Gio.SimpleProxyResolver.set_ignore_hosts"),
	// These are not found.
	AbsoluteFilter("C.GdkPixbufModule"),
	AbsoluteFilter("GdkPixbuf.PixbufNonAnim"),
	AbsoluteFilter("GdkPixbuf.PixbufModulePattern"),
	AbsoluteFilter("GdkPixbuf.PixbufFormat.domain"),
	AbsoluteFilter("GdkPixbuf.PixbufFormat.flags"),
	AbsoluteFilter("GdkPixbuf.PixbufFormat.disabled"),
	// Dangerous.
	AbsoluteFilter("GLib.IOChannel.read"),
	AbsoluteFilter("GLib.Bytes.new_take"),
	AbsoluteFilter("GLib.Bytes.new_static"),
	AbsoluteFilter("GLib.Bytes.unref_to_data"),
	AbsoluteFilter("GLib.Bytes.unref_to_array"),

	FileFilter("gasyncqueue."),
	FileFilter("gatomic."),
	FileFilter("gbacktrace."),
	FileFilter("gbase64."),
	FileFilter("gbitlock."),
	FileFilter("gdataset."),
	FileFilter("gdate."),
	FileFilter("gdatetime."),
	FileFilter("gerror."), // already handled internally
	FileFilter("ghook."),
	FileFilter("glib-unix."),
	FileFilter("glist."),
	FileFilter("gmacros."),
	FileFilter("gmem."),
	FileFilter("gnetworking."), // needs header
	FileFilter("gprintf."),
	FileFilter("grcbox."),
	FileFilter("grefcount."),
	FileFilter("grefstring."),
	FileFilter("gsettingsbackend."),
	FileFilter("gslice."),
	FileFilter("gslist."),
	FileFilter("gstdio."),
	FileFilter("gstrfuncs."),
	FileFilter("gstringchunk."),
	FileFilter("gstring."),
	FileFilter("gstrvbuilder."),
	FileFilter("gtestutils."),
	FileFilter("gthread."),
	FileFilter("gthreadpool."),
	FileFilter("gtrashstack."),

	// Header-specific.
	FileFilter("gskglrenderer."),
	FileFilter("gsknglrenderer."),
	FileFilter("gskvulkanrenderer."),
	FileFilter("gdesktopappinfo."), // See #28.
	// These are not found in GTK4 for some reason, but we're ignoring it for
	// GTK3 as well.
	FileFilter("gtkpagesetupunixdialog"),
	FileFilter("gtkprintunixdialog"),
	FileFilter("gtkprinter"),
	FileFilter("gtkprintjob"),
	FileFilter("gdkprivate"),

	// These are missing on build for some reason.
	AbsoluteFilter("C.g_array_get_type"),
	AbsoluteFilter("C.g_byte_array_get_type"),
	AbsoluteFilter("C.g_bytes_get_type"),
	AbsoluteFilter("C.g_ptr_array_get_type"),
	AbsoluteFilter("C.gtk_header_bar_accessible_get_type"),
	AbsoluteFilter("C.gdk_pixbuf_non_anim_get_type"),
	AbsoluteFilter("C.gdk_window_destroy_notify"),
	AbsoluteFilter("C.gtk_print_capabilities_get_type"),
}

// ImportGError ensures that gerror is imported.
func ImportGError(nsgen *girgen.NamespaceGenerator) error {
	core := file.ImportCore("gerror")

	for _, f := range nsgen.Files {
		if f.Header().HasImport(core) {
			return nil
		}
	}

	f := nsgen.MakeFile("")
	f.Header().DashImport(core)

	return nil
}

// GioArrayUseBytes is the postprocessor that adds gio/v2.UseBytes.
func GioArrayUseBytes(nsgen *girgen.NamespaceGenerator) error {
	fg, ok := nsgen.Files["garray.go"]
	if !ok {
		return nil
	}

	h := fg.Header()
	h.Import("runtime")
	h.ImportCore("gbox")
	h.ImportCore("gextras")
	h.CallbackDelete = true

	// We can use the gbox.Assign API for this. The type doesn't matter much,
	// since we're not actually going to access the data through it.

	p := fg.Pen()
	p.Line(`
		// UseBytes is similar to NewBytes, except the given Go byte slice is
		// not copied, but will be kept alive for the lifetime of the GBytes.
		// Note that the user must NOT modify data.
		//
		// Refer to g_bytes_new_with_free_func() for more information.
		func UseBytes(data []byte) *Bytes {
			byteID := gbox.Assign(data)

			v := C.g_bytes_new_with_free_func(
				C.gconstpointer(unsafe.Pointer(&data[0])),
				C.gsize(len(data)),
				C.GDestroyNotify((*[0]byte)(C.callbackDelete)),
				C.gpointer(byteID),
			)

			_bytes := (*Bytes)(gextras.NewStructNative(unsafe.Pointer(v)))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(_bytes)),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_bytes_unref((*C.GBytes)(intern.C))
				},
			)

			return _bytes
		}
	`)

	return nil
}

// GLibLogs adds the following g_log_* functions:
//
//  - g_log_set_handler
//  - g_log_set_handler_full
//
func GLibLogs(nsgen *girgen.NamespaceGenerator) error {
	fg, ok := nsgen.Files["gmessages.go"]
	if !ok {
		return errors.New("missing file gmessages.go")
	}

	h := fg.Header()
	h.Import("log")
	h.ImportCore("gbox")
	h.CallbackDelete = true

	n := nsgen.Namespace()
	r := nsgen.Repositories()

	h.AddCallback(n, r.FindFullType("GLib-2.LogFunc").Type.(*gir.Callback))
	h.AddCallback(n, r.FindFullType("GLib-2.LogWriterFunc").Type.(*gir.Callback))

	p := fg.Pen()
	p.Line(`
		// LogSetHandler sets the handler used for GLib logging and returns the
		// new handler ID. It is a wrapper around g_log_set_handler and
		// g_log_set_handler_full.
		//
		// To detach a log handler, use LogRemoveHandler.
		func LogSetHandler(domain string, level LogLevelFlags, f LogFunc) uint {
			var log_domain *C.gchar
			if domain != "" {
				log_domain = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
				defer C.free(unsafe.Pointer(log_domain))
			}

			data := gbox.Assign(f)

			h := C.g_log_set_handler_full(
				log_domain,
				C.GLogLevelFlags(level), 
				C.GLogFunc((*[0]byte)(C._gotk4_glib2_LogFunc)),
				C.gpointer(data),
				C.GDestroyNotify((*[0]byte)(C.callbackDelete)),
			)

			return uint(h)
		}

		// Value returns the field's value.
		func (l *LogField) Value() string {
			if l.native.length == -1 {
				return C.GoString((*C.gchar)(unsafe.Pointer(l.native.value)))
			}
			return C.GoStringN((*C.gchar)(unsafe.Pointer(l.native.value)), C.int(l.native.length))
		}

		// LogSetWriter sets the log writer to the given callback, which should
		// take in a list of pair of key-value strings and return true if the
		// log has been successfully written. It is a wrapper around
		// g_log_set_writer_func.
		//
		// Note that this function must ONLY be called ONCE. The GLib
		// documentation states that it is an error to call it more than once.
		func LogSetWriter(f LogWriterFunc) {
			data := gbox.Assign(f)
			C.g_log_set_writer_func(
				C.GLogWriterFunc((*[0]byte)(C._gotk4_glib2_LogWriterFunc)),
				C.gpointer(data),
				C.GDestroyNotify((*[0]byte)(C.callbackDelete)),
			)
		}

		// LogUseDefaultLogger calls LogUseLogger with Go's default standard
		// logger. It is a convenient function for log.Default().
		func LogUseDefaultLogger() {
			LogUseLogger(log.Default())	
		}

		// LogUseLogger calls LogSetWriter with the given Go's standard logger.
		// Note that either this or LogSetWriter must only be called once.
		// The method will ignore all fields excet for "MESSAGE"; for more
		// sophisticated, structured log writing, use LogSetWriter.
		// The output format of the logs printed using this function is not
		// guaranteed to not change. Users who rely on the format are better off
		// using LogSetWriter.
		func LogUseLogger(l *log.Logger) {
			// Treat Lshortfile and Llongfile the same, because we don't have
			// the full path in codeFile anyway.
			Lfile := l.Flags() & (log.Lshortfile | log.Llongfile) != 0

			LogSetWriter(func(lvl LogLevelFlags, fields []LogField) LogWriterOutput {
				var message, codeFile, codeLine, codeFunc string

				for _, field := range fields {
					if !Lfile {
						if field.Key() == "MESSAGE" {
							message = field.Value()
						}
						// Skip setting code* if we don't have to.
						continue
					}

					switch field.Key() {
					case "MESSAGE":   message  = field.Value()
					case "CODE_FILE": codeFile = field.Value()
					case "CODE_LINE": codeLine = field.Value()
					case "CODE_FUNC": codeFunc = field.Value()
					}
				}

				if !Lfile || (codeFile == "" && codeLine == "") {
					l.Print(message)
					return LogWriterHandled
				}

				if codeFunc == "" {
					l.Printf("%s:%s: %s", codeFile, codeLine, message)
					return LogWriterHandled
				}

				l.Printf("%s:%s (%s): %s", codeFile, codeLine, codeFunc, message)
				return LogWriterHandled
			})
		}
	`)

	return nil
}

// Postprocessors is similar to Append, except the caller can mutate the package
// in a more flexible manner.
var Postprocessors = map[string][]girgen.Postprocessor{
	"GLib-2": {GioArrayUseBytes, GLibLogs},
	"Gio-2":  {ImportGError},
	"Gtk-3":  {ImportGError},
	"Gtk-4":  {ImportGError}, // for the marshaler
}

// Appends contains the contents of files that are appended into generated
// outputs. It is used to add custom implementations of missing functions.
var Appends = map[string]string{
	"gtk/v3/gtk.go": `
		// Init binds to the gtk_init() function. Argument parsing is not
		// supported.
		func Init() {
			C.gtk_init(nil, nil)
		}
	`,
}
