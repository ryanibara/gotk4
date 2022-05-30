// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk4_StyleContextClass_changed(GtkStyleContext*);
import "C"

// glib.Type values for gtkstylecontext.go.
var (
	GTypeStyleContextPrintFlags = coreglib.Type(C.gtk_style_context_print_flags_get_type())
	GTypeStyleContext           = coreglib.Type(C.gtk_style_context_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeStyleContextPrintFlags, F: marshalStyleContextPrintFlags},
		{T: GTypeStyleContext, F: marshalStyleContext},
	})
}

// StyleContextPrintFlags flags that modify the behavior of
// gtk_style_context_to_string().
//
// New values may be added to this enumeration.
type StyleContextPrintFlags C.guint

const (
	// StyleContextPrintNone: default value.
	StyleContextPrintNone StyleContextPrintFlags = 0b0
	// StyleContextPrintRecurse: print the entire tree of CSS nodes starting at
	// the style context's node.
	StyleContextPrintRecurse StyleContextPrintFlags = 0b1
	// StyleContextPrintShowStyle: show the values of the CSS properties for
	// each node.
	StyleContextPrintShowStyle StyleContextPrintFlags = 0b10
	// StyleContextPrintShowChange: show information about what changes affect
	// the styles.
	StyleContextPrintShowChange StyleContextPrintFlags = 0b100
)

func marshalStyleContextPrintFlags(p uintptr) (interface{}, error) {
	return StyleContextPrintFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for StyleContextPrintFlags.
func (s StyleContextPrintFlags) String() string {
	if s == 0 {
		return "StyleContextPrintFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(101)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case StyleContextPrintNone:
			builder.WriteString("None|")
		case StyleContextPrintRecurse:
			builder.WriteString("Recurse|")
		case StyleContextPrintShowStyle:
			builder.WriteString("ShowStyle|")
		case StyleContextPrintShowChange:
			builder.WriteString("ShowChange|")
		default:
			builder.WriteString(fmt.Sprintf("StyleContextPrintFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s StyleContextPrintFlags) Has(other StyleContextPrintFlags) bool {
	return (s & other) == other
}

// StyleContextOverrider contains methods that are overridable.
type StyleContextOverrider interface {
	Changed()
}

// StyleContext: GtkStyleContext stores styling information affecting a widget.
//
// In order to construct the final style information, GtkStyleContext queries
// information from all attached GtkStyleProviders. Style providers can be
// either attached explicitly to the context through
// gtk.StyleContext.AddProvider(), or to the display through
// gtk.StyleContext().AddProviderForDisplay. The resulting style is a
// combination of all providers’ information in priority order.
//
// For GTK widgets, any GtkStyleContext returned by gtk.Widget.GetStyleContext()
// will already have a GdkDisplay and RTL/LTR information set. The style context
// will also be updated automatically if any of these settings change on the
// widget.
//
//
// Style Classes
//
// Widgets can add style classes to their context, which can be used to
// associate different styles by class. The documentation for individual widgets
// lists which style classes it uses itself, and which style classes may be
// added by applications to affect their appearance.
//
//
// Custom styling in UI libraries and applications
//
// If you are developing a library with custom widgets that render differently
// than standard components, you may need to add a GtkStyleProvider yourself
// with the GTK_STYLE_PROVIDER_PRIORITY_FALLBACK priority, either a
// GtkCssProvider or a custom object implementing the GtkStyleProvider
// interface. This way themes may still attempt to style your UI elements in a
// different way if needed so.
//
// If you are using custom styling on an applications, you probably want then to
// make your style information prevail to the theme’s, so you must use a
// GtkStyleProvider with the GTK_STYLE_PROVIDER_PRIORITY_APPLICATION priority,
// keep in mind that the user settings in XDG_CONFIG_HOME/gtk-4.0/gtk.css will
// still take precedence over your changes, as it uses the
// GTK_STYLE_PROVIDER_PRIORITY_USER priority.
type StyleContext struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*StyleContext)(nil)
)

func classInitStyleContexter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkStyleContextClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkStyleContextClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Changed() }); ok {
		pclass.changed = (*[0]byte)(C._gotk4_gtk4_StyleContextClass_changed)
	}
}

//export _gotk4_gtk4_StyleContextClass_changed
func _gotk4_gtk4_StyleContextClass_changed(arg0 *C.GtkStyleContext) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Changed() })

	iface.Changed()
}

func wrapStyleContext(obj *coreglib.Object) *StyleContext {
	return &StyleContext{
		Object: obj,
	}
}

func marshalStyleContext(p uintptr) (interface{}, error) {
	return wrapStyleContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AddClass adds a style class to context, so later uses of the style context
// will make use of this new class for styling.
//
// In the CSS file format, a GtkEntry defining a “search” class, would be
// matched by:
//
//    entry.search { ... }
//
//
// While any widget defining a “search” class would be matched by:
//
//    .search { ... }.
//
// The function takes the following parameters:
//
//    - className class name to use in styling.
//
func (context *StyleContext) AddClass(className string) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(className)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**StyleContext)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "StyleContext").InvokeMethod("add_class", args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(className)
}

// AddProvider adds a style provider to context, to be used in style
// construction.
//
// Note that a style provider added by this function only affects the style of
// the widget to which context belongs. If you want to affect the style of all
// widgets, use gtk.StyleContext().AddProviderForDisplay.
//
// Note: If both priorities are the same, a GtkStyleProvider added through this
// function takes precedence over another added through
// gtk.StyleContext().AddProviderForDisplay.
//
// The function takes the following parameters:
//
//    - provider: GtkStyleProvider.
//    - priority of the style provider. The lower it is, the earlier it will be
//      used in the style construction. Typically this will be in the range
//      between GTK_STYLE_PROVIDER_PRIORITY_FALLBACK and
//      GTK_STYLE_PROVIDER_PRIORITY_USER.
//
func (context *StyleContext) AddProvider(provider StyleProviderer, priority uint) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 C.guint // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg2 = C.guint(priority)
	*(**StyleContext)(unsafe.Pointer(&args[1])) = _arg1
	*(*StyleProviderer)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "StyleContext").InvokeMethod("add_provider", args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(priority)
}

// Display returns the GdkDisplay to which context is attached.
//
// The function returns the following values:
//
//    - display: GdkDisplay.
//
func (context *StyleContext) Display() *gdk.Display {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(**StyleContext)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "StyleContext").InvokeMethod("get_display", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _display *gdk.Display // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_display = &gdk.Display{
			Object: obj,
		}
	}

	return _display
}

// HasClass returns TRUE if context currently has defined the given class name.
//
// The function takes the following parameters:
//
//    - className class name.
//
// The function returns the following values:
//
//    - ok: TRUE if context has class_name defined.
//
func (context *StyleContext) HasClass(className string) bool {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(className)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**StyleContext)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "StyleContext").InvokeMethod("has_class", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)
	runtime.KeepAlive(className)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RemoveClass removes class_name from context.
//
// The function takes the following parameters:
//
//    - className class name to remove.
//
func (context *StyleContext) RemoveClass(className string) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(className)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**StyleContext)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "StyleContext").InvokeMethod("remove_class", args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(className)
}

// RemoveProvider removes provider from the style providers list in context.
//
// The function takes the following parameters:
//
//    - provider: GtkStyleProvider.
//
func (context *StyleContext) RemoveProvider(provider StyleProviderer) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	*(**StyleContext)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "StyleContext").InvokeMethod("remove_provider", args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(provider)
}

// Restore restores context state to a previous stage.
//
// See gtk.StyleContext.Save().
func (context *StyleContext) Restore() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(**StyleContext)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "StyleContext").InvokeMethod("restore", args[:], nil)

	runtime.KeepAlive(context)
}

// Save saves the context state.
//
// This allows temporary modifications done through gtk.StyleContext.AddClass(),
// gtk.StyleContext.RemoveClass(), gtk.StyleContext.SetState() to be quickly
// reverted in one go through gtk.StyleContext.Restore().
//
// The matching call to gtk.StyleContext.Restore() must be done before GTK
// returns to the main loop.
func (context *StyleContext) Save() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(**StyleContext)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "StyleContext").InvokeMethod("save", args[:], nil)

	runtime.KeepAlive(context)
}

// SetDisplay attaches context to the given display.
//
// The display is used to add style information from “global” style providers,
// such as the display's GtkSettings instance.
//
// If you are using a GtkStyleContext returned from
// gtk.Widget.GetStyleContext(), you do not need to call this yourself.
//
// The function takes the following parameters:
//
//    - display: GdkDisplay.
//
func (context *StyleContext) SetDisplay(display *gdk.Display) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(display).Native()))
	*(**StyleContext)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "StyleContext").InvokeMethod("set_display", args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(display)
}

// StyleContextAddProviderForDisplay adds a global style provider to display,
// which will be used in style construction for all GtkStyleContexts under
// display.
//
// GTK uses this to make styling information from GtkSettings available.
//
// Note: If both priorities are the same, A GtkStyleProvider added through
// gtk.StyleContext.AddProvider() takes precedence over another added through
// this function.
//
// The function takes the following parameters:
//
//    - display: GdkDisplay.
//    - provider: GtkStyleProvider.
//    - priority of the style provider. The lower it is, the earlier it will be
//      used in the style construction. Typically this will be in the range
//      between GTK_STYLE_PROVIDER_PRIORITY_FALLBACK and
//      GTK_STYLE_PROVIDER_PRIORITY_USER.
//
func StyleContextAddProviderForDisplay(display *gdk.Display, provider StyleProviderer, priority uint) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 C.guint // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(display).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg2 = C.guint(priority)
	*(**gdk.Display)(unsafe.Pointer(&args[0])) = _arg0
	*(*StyleProviderer)(unsafe.Pointer(&args[1])) = _arg1
	*(*uint)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "add_provider_for_display").Invoke(args[:], nil)

	runtime.KeepAlive(display)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(priority)
}

// StyleContextRemoveProviderForDisplay removes provider from the global style
// providers list in display.
//
// The function takes the following parameters:
//
//    - display: GdkDisplay.
//    - provider: GtkStyleProvider.
//
func StyleContextRemoveProviderForDisplay(display *gdk.Display, provider StyleProviderer) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(display).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	*(**gdk.Display)(unsafe.Pointer(&args[0])) = _arg0
	*(*StyleProviderer)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "remove_provider_for_display").Invoke(args[:], nil)

	runtime.KeepAlive(display)
	runtime.KeepAlive(provider)
}
