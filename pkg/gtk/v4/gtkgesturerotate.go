// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_rotate_get_type()), F: marshalGestureRotate},
	})
}

// GestureRotate: `GtkGestureRotate` is a `GtkGesture` for 2-finger rotations.
//
// Whenever the angle between both handled sequences changes, the
// [signal@Gtk.GestureRotate::angle-changed] signal is emitted.
type GestureRotate interface {
	Gesture

	// AngleDelta gets the angle delta in radians.
	//
	// If @gesture is active, this function returns the angle difference in
	// radians since the gesture was first recognized. If @gesture is not
	// active, 0 is returned.
	AngleDelta() float64
}

// gestureRotate implements the GestureRotate interface.
type gestureRotate struct {
	Gesture
}

var _ GestureRotate = (*gestureRotate)(nil)

// WrapGestureRotate wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureRotate(obj *externglib.Object) GestureRotate {
	return GestureRotate{
		Gesture: WrapGesture(obj),
	}
}

func marshalGestureRotate(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureRotate(obj), nil
}

// NewGestureRotate constructs a class GestureRotate.
func NewGestureRotate() GestureRotate {
	var _cret C.GtkGestureRotate

	cret = C.gtk_gesture_rotate_new()

	var _gestureRotate GestureRotate

	_gestureRotate = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(GestureRotate)

	return _gestureRotate
}

// AngleDelta gets the angle delta in radians.
//
// If @gesture is active, this function returns the angle difference in
// radians since the gesture was first recognized. If @gesture is not
// active, 0 is returned.
func (g gestureRotate) AngleDelta() float64 {
	var _arg0 *C.GtkGestureRotate

	_arg0 = (*C.GtkGestureRotate)(unsafe.Pointer(g.Native()))

	var _cret C.double

	cret = C.gtk_gesture_rotate_get_angle_delta(_arg0)

	var _gdouble float64

	_gdouble = (float64)(_cret)

	return _gdouble
}