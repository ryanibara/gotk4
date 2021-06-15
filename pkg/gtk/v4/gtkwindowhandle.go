// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_window_handle_get_type()), F: marshalWindowHandle},
	})
}

// WindowHandle: `GtkWindowHandle` is a titlebar area widget.
//
// When added into a window, it can be dragged to move the window, and handles
// right click, double click and middle click as expected of a titlebar.
//
//
// CSS nodes
//
// `GtkWindowHandle` has a single CSS node with the name `windowhandle`.
//
//
// Accessibility
//
// `GtkWindowHandle` uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type WindowHandle interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget

	// Child gets the child widget of @self.
	Child() Widget
	// SetChild sets the child widget of @self.
	SetChild(child Widget)
}

// windowHandle implements the WindowHandle class.
type windowHandle struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ WindowHandle = (*windowHandle)(nil)

// WrapWindowHandle wraps a GObject to the right type. It is
// primarily used internally.
func WrapWindowHandle(obj *externglib.Object) WindowHandle {
	return windowHandle{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalWindowHandle(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWindowHandle(obj), nil
}

// NewWindowHandle constructs a class WindowHandle.
func NewWindowHandle() WindowHandle {
	var _cret C.GtkWindowHandle // in

	_cret = C.gtk_window_handle_new()

	var _windowHandle WindowHandle // out

	_windowHandle = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(WindowHandle)

	return _windowHandle
}

// Child gets the child widget of @self.
func (s windowHandle) Child() Widget {
	var _arg0 *C.GtkWindowHandle // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkWindowHandle)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_window_handle_get_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// SetChild sets the child widget of @self.
func (s windowHandle) SetChild(child Widget) {
	var _arg0 *C.GtkWindowHandle // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkWindowHandle)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_window_handle_set_child(_arg0, _arg1)
}
