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
		{T: externglib.Type(C.gtk_gesture_swipe_get_type()), F: marshalGestureSwipe},
	})
}

// GestureSwipe: `GtkGestureSwipe` is a `GtkGesture` for swipe gestures.
//
// After a press/move/.../move/release sequence happens, the
// [signal@Gtk.GestureSwipe::swipe] signal will be emitted, providing the
// velocity and directionality of the sequence at the time it was lifted.
//
// If the velocity is desired in intermediate points,
// [method@Gtk.GestureSwipe.get_velocity] can be called in a
// [signal@Gtk.Gesture::update] handler.
//
// All velocities are reported in pixels/sec units.
type GestureSwipe interface {
	GestureSingle

	// Velocity gets the current velocity.
	//
	// If the gesture is recognized, this function returns true and fills in
	// @velocity_x and @velocity_y with the recorded velocity, as per the last
	// events processed.
	Velocity() (velocityX float64, velocityY float64, ok bool)
}

// gestureSwipe implements the GestureSwipe interface.
type gestureSwipe struct {
	GestureSingle
}

var _ GestureSwipe = (*gestureSwipe)(nil)

// WrapGestureSwipe wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureSwipe(obj *externglib.Object) GestureSwipe {
	return GestureSwipe{
		GestureSingle: WrapGestureSingle(obj),
	}
}

func marshalGestureSwipe(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureSwipe(obj), nil
}

// NewGestureSwipe constructs a class GestureSwipe.
func NewGestureSwipe() GestureSwipe {
	ret := C.gtk_gesture_swipe_new()

	var ret0 GestureSwipe

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(GestureSwipe)

	return ret0
}

// Velocity gets the current velocity.
//
// If the gesture is recognized, this function returns true and fills in
// @velocity_x and @velocity_y with the recorded velocity, as per the last
// events processed.
func (g gestureSwipe) Velocity() (velocityX float64, velocityY float64, ok bool) {
	var arg0 *C.GtkGestureSwipe
	var arg1 *C.double // out
	var arg2 *C.double // out

	arg0 = (*C.GtkGestureSwipe)(g.Native())

	ret := C.gtk_gesture_swipe_get_velocity(arg0, &arg1, &arg2)

	var ret0 float64
	var ret1 float64
	var ret2 bool

	ret0 = float64(arg1)

	ret1 = float64(arg2)

	ret2 = C.bool(ret) != C.false

	return ret0, ret1, ret2
}
