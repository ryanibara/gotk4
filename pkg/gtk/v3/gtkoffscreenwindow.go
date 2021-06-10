// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_offscreen_window_get_type()), F: marshalOffscreenWindow},
	})
}

// OffscreenWindow: gtkOffscreenWindow is strictly intended to be used for
// obtaining snapshots of widgets that are not part of a normal widget
// hierarchy. Since OffscreenWindow is a toplevel widget you cannot obtain
// snapshots of a full window with it since you cannot pack a toplevel widget in
// another toplevel.
//
// The idea is to take a widget and manually set the state of it, add it to a
// GtkOffscreenWindow and then retrieve the snapshot as a #cairo_surface_t or
// Pixbuf.
//
// GtkOffscreenWindow derives from Window only as an implementation detail.
// Applications should not use any API specific to Window to operate on this
// object. It should be treated as a Bin that has no parent widget.
//
// When contained offscreen widgets are redrawn, GtkOffscreenWindow will emit a
// Widget::damage-event signal.
type OffscreenWindow interface {
	Window
	Buildable
}

// offscreenWindow implements the OffscreenWindow interface.
type offscreenWindow struct {
	Window
	Buildable
}

var _ OffscreenWindow = (*offscreenWindow)(nil)

// WrapOffscreenWindow wraps a GObject to the right type. It is
// primarily used internally.
func WrapOffscreenWindow(obj *externglib.Object) OffscreenWindow {
	return OffscreenWindow{
		Window:    WrapWindow(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalOffscreenWindow(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapOffscreenWindow(obj), nil
}
