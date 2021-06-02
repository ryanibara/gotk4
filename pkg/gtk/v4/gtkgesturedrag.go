// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_drag_get_type()), F: marshalGestureDrag},
	})
}

// GestureDrag: `GtkGestureDrag` is a `GtkGesture` implementation for drags.
//
// The drag operation itself can be tracked throughout the
// [signal@Gtk.GestureDrag::drag-begin], [signal@Gtk.GestureDrag::drag-update]
// and [signal@Gtk.GestureDrag::drag-end] signals, and the relevant coordinates
// can be extracted through [method@Gtk.GestureDrag.get_offset] and
// [method@Gtk.GestureDrag.get_start_point].
type GestureDrag interface {
	GestureSingle

	// Offset gets the offset from the start point.
	//
	// If the @gesture is active, this function returns true and fills in @x and
	// @y with the coordinates of the current point, as an offset to the
	// starting drag point.
	Offset() (x float64, y float64, ok bool)
	// StartPoint gets the point where the drag started.
	//
	// If the @gesture is active, this function returns true and fills in @x and
	// @y with the drag start coordinates, in surface-relative coordinates.
	StartPoint() (x float64, y float64, ok bool)
}

// gestureDrag implements the GestureDrag interface.
type gestureDrag struct {
	GestureSingle
}

var _ GestureDrag = (*gestureDrag)(nil)

// WrapGestureDrag wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureDrag(obj *externglib.Object) GestureDrag {
	return GestureDrag{
		GestureSingle: WrapGestureSingle(obj),
	}
}

func marshalGestureDrag(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureDrag(obj), nil
}

// NewGestureDrag constructs a class GestureDrag.
func NewGestureDrag() GestureDrag {
	ret := C.gtk_gesture_drag_new()

	var ret0 GestureDrag

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(GestureDrag)

	return ret0
}

// Offset gets the offset from the start point.
//
// If the @gesture is active, this function returns true and fills in @x and
// @y with the coordinates of the current point, as an offset to the
// starting drag point.
func (g gestureDrag) Offset() (x float64, y float64, ok bool) {
	var arg0 *C.GtkGestureDrag
	var arg1 *C.double // out
	var arg2 *C.double // out

	arg0 = (*C.GtkGestureDrag)(g.Native())

	ret := C.gtk_gesture_drag_get_offset(arg0, &arg1, &arg2)

	var ret0 float64
	var ret1 float64
	var ret2 bool

	ret0 = float64(arg1)

	ret1 = float64(arg2)

	ret2 = C.bool(ret) != 0

	return ret0, ret1, ret2
}

// StartPoint gets the point where the drag started.
//
// If the @gesture is active, this function returns true and fills in @x and
// @y with the drag start coordinates, in surface-relative coordinates.
func (g gestureDrag) StartPoint() (x float64, y float64, ok bool) {
	var arg0 *C.GtkGestureDrag
	var arg1 *C.double // out
	var arg2 *C.double // out

	arg0 = (*C.GtkGestureDrag)(g.Native())

	ret := C.gtk_gesture_drag_get_start_point(arg0, &arg1, &arg2)

	var ret0 float64
	var ret1 float64
	var ret2 bool

	ret0 = float64(arg1)

	ret1 = float64(arg2)

	ret2 = C.bool(ret) != 0

	return ret0, ret1, ret2
}
