// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_invisible_get_type()), F: marshalInvisible},
	})
}

// Invisible: the Invisible widget is used internally in GTK+, and is probably
// not very useful for application developers.
//
// It is used for reliable pointer grabs and selection handling in the code for
// drag-and-drop.
type Invisible interface {
	Widget

	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable

	// Screen returns the Screen object associated with @invisible
	Screen() gdk.Screen
	// SetScreenInvisible sets the Screen where the Invisible object will be
	// displayed.
	SetScreenInvisible(screen gdk.Screen)
}

// invisible implements the Invisible class.
type invisible struct {
	Widget
}

// WrapInvisible wraps a GObject to the right type. It is
// primarily used internally.
func WrapInvisible(obj *externglib.Object) Invisible {
	return invisible{
		Widget: WrapWidget(obj),
	}
}

func marshalInvisible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapInvisible(obj), nil
}

// NewInvisible creates a new Invisible.
func NewInvisible() Invisible {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_invisible_new()

	var _invisible Invisible // out

	_invisible = WrapInvisible(externglib.Take(unsafe.Pointer(_cret)))

	return _invisible
}

// NewInvisibleForScreen creates a new Invisible object for a specified screen
func NewInvisibleForScreen(screen gdk.Screen) Invisible {
	var _arg1 *C.GdkScreen // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gtk_invisible_new_for_screen(_arg1)

	var _invisible Invisible // out

	_invisible = WrapInvisible(externglib.Take(unsafe.Pointer(_cret)))

	return _invisible
}

func (i invisible) Screen() gdk.Screen {
	var _arg0 *C.GtkInvisible // out
	var _cret *C.GdkScreen    // in

	_arg0 = (*C.GtkInvisible)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_invisible_get_screen(_arg0)

	var _screen gdk.Screen // out

	_screen = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Screen)

	return _screen
}

func (i invisible) SetScreenInvisible(screen gdk.Screen) {
	var _arg0 *C.GtkInvisible // out
	var _arg1 *C.GdkScreen    // out

	_arg0 = (*C.GtkInvisible)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	C.gtk_invisible_set_screen(_arg0, _arg1)
}

func (i invisible) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(i))
}
