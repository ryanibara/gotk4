// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_CssProvider_ConnectParsingError(gpointer, void*, void*, guintptr);
import "C"

// GTypeCSSProvider returns the GType for the type CSSProvider.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCSSProvider() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "CssProvider").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalCSSProvider)
	return gtype
}

// CSSProvider: GtkCssProvider is an object implementing the GtkStyleProvider
// interface for CSS.
//
// It is able to parse CSS-like input in order to style widgets.
//
// An application can make GTK parse a specific CSS style sheet by calling
// gtk.CSSProvider.LoadFromFile() or gtk.CSSProvider.LoadFromResource() and
// adding the provider with gtk.StyleContext.AddProvider() or
// gtk.StyleContext().AddProviderForDisplay.
//
// In addition, certain files will be read when GTK is initialized. First, the
// file $XDG_CONFIG_HOME/gtk-4.0/gtk.css is loaded if it exists. Then, GTK loads
// the first existing file among
// XDG_DATA_HOME/themes/THEME/gtk-VERSION/gtk-VARIANT.css,
// $HOME/.themes/THEME/gtk-VERSION/gtk-VARIANT.css,
// $XDG_DATA_DIRS/themes/THEME/gtk-VERSION/gtk-VARIANT.css and
// DATADIR/share/themes/THEME/gtk-VERSION/gtk-VARIANT.css, where THEME is the
// name of the current theme (see the gtk.Settings:gtk-theme-name setting),
// VARIANT is the variant to load (see the
// gtk.Settings:gtk-application-prefer-dark-theme setting), DATADIR is the
// prefix configured when GTK was compiled (unless overridden by the
// GTK_DATA_PREFIX environment variable), and VERSION is the GTK version number.
// If no file is found for the current version, GTK tries older versions all the
// way back to 4.0.
//
// To track errors while loading CSS, connect to the
// gtk.CSSProvider::parsing-error signal.
type CSSProvider struct {
	_ [0]func() // equal guard
	*coreglib.Object

	StyleProvider
}

var (
	_ coreglib.Objector = (*CSSProvider)(nil)
)

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

//export _gotk4_gtk4_CssProvider_ConnectParsingError
func _gotk4_gtk4_CssProvider_ConnectParsingError(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 C.guintptr) {
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
			C.free(intern.C)
		},
	)
	_err = gerror.Take(unsafe.Pointer(arg2))

	f(_section, _err)
}

// ConnectParsingError signals that a parsing error occurred.
//
// The path, line and position describe the actual location of the error as
// accurately as possible.
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
	return coreglib.ConnectGeneratedClosure(cssProvider, "parsing-error", false, unsafe.Pointer(C._gotk4_gtk4_CssProvider_ConnectParsingError), f)
}

// NewCSSProvider returns a newly created GtkCssProvider.
//
// The function returns the following values:
//
//    - cssProvider: new GtkCssProvider.
//
func NewCSSProvider() *CSSProvider {
	_gret := girepository.MustFind("Gtk", "CssProvider").InvokeMethod("new_CssProvider", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _cssProvider *CSSProvider // out

	_cssProvider = wrapCSSProvider(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cssProvider
}

// LoadFromData loads data into css_provider.
//
// This clears any previously loaded information.
//
// The function takes the following parameters:
//
//    - data: CSS data loaded in memory.
//
func (cssProvider *CSSProvider) LoadFromData(data string) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cssProvider).Native()))
	*(*C.gssize)(unsafe.Pointer(&_args[2])) = (C.gssize)(len(data))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(C.calloc(C.size_t((len(data) + 1)), C.size_t(C.sizeof_char)))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_args[1])))), len(data)), data)
	defer C.free(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_args[1]))))

	girepository.MustFind("Gtk", "CssProvider").InvokeMethod("load_from_data", _args[:], nil)

	runtime.KeepAlive(cssProvider)
	runtime.KeepAlive(data)
}

// LoadFromFile loads the data contained in file into css_provider.
//
// This clears any previously loaded information.
//
// The function takes the following parameters:
//
//    - file: GFile pointing to a file to load.
//
func (cssProvider *CSSProvider) LoadFromFile(file gio.Filer) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cssProvider).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	girepository.MustFind("Gtk", "CssProvider").InvokeMethod("load_from_file", _args[:], nil)

	runtime.KeepAlive(cssProvider)
	runtime.KeepAlive(file)
}

// LoadFromPath loads the data contained in path into css_provider.
//
// This clears any previously loaded information.
//
// The function takes the following parameters:
//
//    - path of a filename to load, in the GLib filename encoding.
//
func (cssProvider *CSSProvider) LoadFromPath(path string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cssProvider).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "CssProvider").InvokeMethod("load_from_path", _args[:], nil)

	runtime.KeepAlive(cssProvider)
	runtime.KeepAlive(path)
}

// LoadFromResource loads the data contained in the resource at resource_path
// into the css_provider.
//
// This clears any previously loaded information.
//
// The function takes the following parameters:
//
//    - resourcePath: GResource resource path.
//
func (cssProvider *CSSProvider) LoadFromResource(resourcePath string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cssProvider).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "CssProvider").InvokeMethod("load_from_resource", _args[:], nil)

	runtime.KeepAlive(cssProvider)
	runtime.KeepAlive(resourcePath)
}

// LoadNamed loads a theme from the usual theme paths.
//
// The actual process of finding the theme might change between releases, but it
// is guaranteed that this function uses the same mechanism to load the theme
// that GTK uses for loading its own theme.
//
// The function takes the following parameters:
//
//    - name: theme name.
//    - variant (optional) to load, for example, "dark", or NULL for the default.
//
func (provider *CSSProvider) LoadNamed(name, variant string) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))
	if variant != "" {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(variant)))
		defer C.free(unsafe.Pointer(_args[2]))
	}

	girepository.MustFind("Gtk", "CssProvider").InvokeMethod("load_named", _args[:], nil)

	runtime.KeepAlive(provider)
	runtime.KeepAlive(name)
	runtime.KeepAlive(variant)
}

// String converts the provider into a string representation in CSS format.
//
// Using gtk.CSSProvider.LoadFromData() with the return value from this function
// on a new provider created with gtk.CSSProvider.New will basically create a
// duplicate of this provider.
//
// The function returns the following values:
//
//    - utf8: new string representing the provider.
//
func (provider *CSSProvider) String() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	_gret := girepository.MustFind("Gtk", "CssProvider").InvokeMethod("to_string", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(provider)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
