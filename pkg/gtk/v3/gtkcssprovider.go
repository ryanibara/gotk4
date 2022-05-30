// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_CssProviderClass_parsing_error(GtkCssProvider*, GtkCssSection*, GError*);
// extern void _gotk4_gtk3_CssProvider_ConnectParsingError(gpointer, GtkCssSection*, GError*, guintptr);
import "C"

// glib.Type values for gtkcssprovider.go.
var (
	GTypeCSSProviderError = coreglib.Type(C.gtk_css_provider_error_get_type())
	GTypeCSSProvider      = coreglib.Type(C.gtk_css_provider_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeCSSProviderError, F: marshalCSSProviderError},
		{T: GTypeCSSProvider, F: marshalCSSProvider},
	})
}

// CSSProviderError: error codes for GTK_CSS_PROVIDER_ERROR.
type CSSProviderError C.gint

const (
	// CSSProviderErrorFailed: failed.
	CSSProviderErrorFailed CSSProviderError = iota
	// CSSProviderErrorSyntax: syntax error.
	CSSProviderErrorSyntax
	// CSSProviderErrorImport: import error.
	CSSProviderErrorImport
	// CSSProviderErrorName: name error.
	CSSProviderErrorName
	// CSSProviderErrorDeprecated: deprecation error.
	CSSProviderErrorDeprecated
	// CSSProviderErrorUnknownValue: unknown value.
	CSSProviderErrorUnknownValue
)

func marshalCSSProviderError(p uintptr) (interface{}, error) {
	return CSSProviderError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CSSProviderError.
func (c CSSProviderError) String() string {
	switch c {
	case CSSProviderErrorFailed:
		return "Failed"
	case CSSProviderErrorSyntax:
		return "Syntax"
	case CSSProviderErrorImport:
		return "Import"
	case CSSProviderErrorName:
		return "Name"
	case CSSProviderErrorDeprecated:
		return "Deprecated"
	case CSSProviderErrorUnknownValue:
		return "UnknownValue"
	default:
		return fmt.Sprintf("CSSProviderError(%d)", c)
	}
}

// CSSProviderOverrider contains methods that are overridable.
type CSSProviderOverrider interface {
	// The function takes the following parameters:
	//
	//    - section
	//    - err
	//
	ParsingError(section *CSSSection, err error)
}

// CSSProvider is an object implementing the StyleProvider interface. It is able
// to parse [CSS-like][css-overview] input in order to style widgets.
//
// An application can make GTK+ parse a specific CSS style sheet by calling
// gtk_css_provider_load_from_file() or gtk_css_provider_load_from_resource()
// and adding the provider with gtk_style_context_add_provider() or
// gtk_style_context_add_provider_for_screen().
//
// In addition, certain files will be read when GTK+ is initialized. First, the
// file $XDG_CONFIG_HOME/gtk-3.0/gtk.css is loaded if it exists. Then, GTK+
// loads the first existing file among
// XDG_DATA_HOME/themes/THEME/gtk-VERSION/gtk.css,
// $HOME/.themes/THEME/gtk-VERSION/gtk.css,
// $XDG_DATA_DIRS/themes/THEME/gtk-VERSION/gtk.css and
// DATADIR/share/themes/THEME/gtk-VERSION/gtk.css, where THEME is the name of
// the current theme (see the Settings:gtk-theme-name setting), DATADIR is the
// prefix configured when GTK+ was compiled (unless overridden by the
// GTK_DATA_PREFIX environment variable), and VERSION is the GTK+ version
// number. If no file is found for the current version, GTK+ tries older
// versions all the way back to 3.0.
//
// In the same way, GTK+ tries to load a gtk-keys.css file for the current key
// theme, as defined by Settings:gtk-key-theme-name.
type CSSProvider struct {
	_ [0]func() // equal guard
	*coreglib.Object

	StyleProvider
}

var (
	_ coreglib.Objector = (*CSSProvider)(nil)
)

func classInitCSSProviderer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkCssProviderClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkCssProviderClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface {
		ParsingError(section *CSSSection, err error)
	}); ok {
		pclass.parsing_error = (*[0]byte)(C._gotk4_gtk3_CssProviderClass_parsing_error)
	}
}

//export _gotk4_gtk3_CssProviderClass_parsing_error
func _gotk4_gtk3_CssProviderClass_parsing_error(arg0 *C.GtkCssProvider, arg1 *C.GtkCssSection, arg2 *C.GError) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		ParsingError(section *CSSSection, err error)
	})

	var _section *CSSSection // out
	var _err error           // out

	_section = (*CSSSection)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	C.gtk_css_section_ref(arg1)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_section)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_css_section_unref((*C.GtkCssSection)(intern.C))
		},
	)
	_err = gerror.Take(unsafe.Pointer(arg2))

	iface.ParsingError(_section, _err)
}

func wrapCSSProvider(obj *coreglib.Object) *CSSProvider {
	return &CSSProvider{
		Object: obj,
		StyleProvider: StyleProvider{
			Object: obj,
		},
	}
}

func marshalCSSProvider(p uintptr) (interface{}, error) {
	return wrapCSSProvider(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_CssProvider_ConnectParsingError
func _gotk4_gtk3_CssProvider_ConnectParsingError(arg0 C.gpointer, arg1 *C.GtkCssSection, arg2 *C.GError, arg3 C.guintptr) {
	var f func(section *CSSSection, err error)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(section *CSSSection, err error))
	}

	var _section *CSSSection // out
	var _err error           // out

	_section = (*CSSSection)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	C.gtk_css_section_ref(arg1)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_section)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_css_section_unref((*C.GtkCssSection)(intern.C))
		},
	)
	_err = gerror.Take(unsafe.Pointer(arg2))

	f(_section, _err)
}

// ConnectParsingError signals that a parsing error occurred. the path, line and
// position describe the actual location of the error as accurately as possible.
//
// Parsing errors are never fatal, so the parsing will resume after the error.
// Errors may however cause parts of the given data or even all of it to not be
// parsed at all. So it is a useful idea to check that the parsing succeeds by
// connecting to this signal.
//
// Note that this signal may be emitted at any time as the css provider may opt
// to defer parsing parts or all of the input to a later time than when a
// loading function was called.
func (cssProvider *CSSProvider) ConnectParsingError(f func(section *CSSSection, err error)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(cssProvider, "parsing-error", false, unsafe.Pointer(C._gotk4_gtk3_CssProvider_ConnectParsingError), f)
}

// NewCSSProvider returns a newly created CssProvider.
//
// The function returns the following values:
//
//    - cssProvider: new CssProvider.
//
func NewCSSProvider() *CSSProvider {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "CssProvider").InvokeMethod("new_CssProvider", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _cssProvider *CSSProvider // out

	_cssProvider = wrapCSSProvider(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cssProvider
}

// LoadFromData loads data into css_provider, and by doing so clears any
// previously loaded information.
//
// The function takes the following parameters:
//
//    - data: CSS data loaded in memory.
//
func (cssProvider *CSSProvider) LoadFromData(data string) error {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 C.gssize
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cssProvider).Native()))
	_arg2 = (C.gssize)(len(data))
	_arg1 = (*C.void)(C.calloc(C.size_t((len(data) + 1)), C.size_t(C.sizeof_gchar)))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(_arg1)), len(data)), data)
	defer C.free(unsafe.Pointer(_arg1))
	*(**CSSProvider)(unsafe.Pointer(&args[1])) = _arg1
	*(*string)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "CssProvider").InvokeMethod("load_from_data", args[:], nil)

	runtime.KeepAlive(cssProvider)
	runtime.KeepAlive(data)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// LoadFromFile loads the data contained in file into css_provider, making it
// clear any previously loaded information.
//
// The function takes the following parameters:
//
//    - file pointing to a file to load.
//
func (cssProvider *CSSProvider) LoadFromFile(file gio.Filer) error {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cssProvider).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(file).Native()))
	*(**CSSProvider)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "CssProvider").InvokeMethod("load_from_file", args[:], nil)

	runtime.KeepAlive(cssProvider)
	runtime.KeepAlive(file)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// LoadFromPath loads the data contained in path into css_provider, making it
// clear any previously loaded information.
//
// The function takes the following parameters:
//
//    - path of a filename to load, in the GLib filename encoding.
//
func (cssProvider *CSSProvider) LoadFromPath(path string) error {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cssProvider).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**CSSProvider)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "CssProvider").InvokeMethod("load_from_path", args[:], nil)

	runtime.KeepAlive(cssProvider)
	runtime.KeepAlive(path)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// LoadFromResource loads the data contained in the resource at resource_path
// into the CssProvider, clearing any previously loaded information.
//
// To track errors while loading CSS, connect to the CssProvider::parsing-error
// signal.
//
// The function takes the following parameters:
//
//    - resourcePath resource path.
//
func (cssProvider *CSSProvider) LoadFromResource(resourcePath string) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cssProvider).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**CSSProvider)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "CssProvider").InvokeMethod("load_from_resource", args[:], nil)

	runtime.KeepAlive(cssProvider)
	runtime.KeepAlive(resourcePath)
}

// String converts the provider into a string representation in CSS format.
//
// Using gtk_css_provider_load_from_data() with the return value from this
// function on a new provider created with gtk_css_provider_new() will basically
// create a duplicate of this provider.
//
// The function returns the following values:
//
//    - utf8: new string representing the provider.
//
func (provider *CSSProvider) String() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	*(**CSSProvider)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "CssProvider").InvokeMethod("to_string", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(provider)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// CSSProviderGetDefault returns the provider containing the style settings used
// as a fallback for all widgets.
//
// Deprecated: Use gtk_css_provider_new() instead.
//
// The function returns the following values:
//
//    - cssProvider: provider used for fallback styling. This memory is owned by
//      GTK+, and you must not free it.
//
func CSSProviderGetDefault() *CSSProvider {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "get_default").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _cssProvider *CSSProvider // out

	_cssProvider = wrapCSSProvider(coreglib.Take(unsafe.Pointer(_cret)))

	return _cssProvider
}

// CSSProviderGetNamed loads a theme from the usual theme paths.
//
// The function takes the following parameters:
//
//    - name: theme name.
//    - variant (optional) to load, for example, "dark", or NULL for the default.
//
// The function returns the following values:
//
//    - cssProvider with the theme loaded. This memory is owned by GTK+, and you
//      must not free it.
//
func CSSProviderGetNamed(name, variant string) *CSSProvider {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg0))
	if variant != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(variant)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	*(*string)(unsafe.Pointer(&args[0])) = _arg0
	*(*string)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "get_named").Invoke(args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(name)
	runtime.KeepAlive(variant)

	var _cssProvider *CSSProvider // out

	_cssProvider = wrapCSSProvider(coreglib.Take(unsafe.Pointer(_cret)))

	return _cssProvider
}
