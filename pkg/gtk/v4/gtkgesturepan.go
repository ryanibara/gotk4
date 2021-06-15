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
		{T: externglib.Type(C.gtk_gesture_pan_get_type()), F: marshalGesturePan},
	})
}

// GesturePan: `GtkGesturePan` is a `GtkGesture` for pan gestures.
//
// These are drags that are locked to happen along one axis. The axis that a
// `GtkGesturePan` handles is defined at construct time, and can be changed
// through [method@Gtk.GesturePan.set_orientation].
//
// When the gesture starts to be recognized, `GtkGesturePan` will attempt to
// determine as early as possible whether the sequence is moving in the expected
// direction, and denying the sequence if this does not happen.
//
// Once a panning gesture along the expected axis is recognized, the
// [signal@Gtk.GesturePan::pan] signal will be emitted as input events are
// received, containing the offset in the given axis.
type GesturePan interface {
	GestureDrag

	// Orientation returns the orientation of the pan gestures that this
	// @gesture expects.
	Orientation() Orientation
	// SetOrientation sets the orientation to be expected on pan gestures.
	SetOrientation(orientation Orientation)
}

// gesturePan implements the GesturePan class.
type gesturePan struct {
	GestureDrag
}

var _ GesturePan = (*gesturePan)(nil)

// WrapGesturePan wraps a GObject to the right type. It is
// primarily used internally.
func WrapGesturePan(obj *externglib.Object) GesturePan {
	return gesturePan{
		GestureDrag: WrapGestureDrag(obj),
	}
}

func marshalGesturePan(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGesturePan(obj), nil
}

// NewGesturePan constructs a class GesturePan.
func NewGesturePan(orientation Orientation) GesturePan {
	var _arg1 C.GtkOrientation // out
	var _cret C.GtkGesturePan  // in

	_arg1 = (C.GtkOrientation)(orientation)

	_cret = C.gtk_gesture_pan_new(_arg1)

	var _gesturePan GesturePan // out

	_gesturePan = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(GesturePan)

	return _gesturePan
}

// Orientation returns the orientation of the pan gestures that this
// @gesture expects.
func (g gesturePan) Orientation() Orientation {
	var _arg0 *C.GtkGesturePan // out
	var _cret C.GtkOrientation // in

	_arg0 = (*C.GtkGesturePan)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_pan_get_orientation(_arg0)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

// SetOrientation sets the orientation to be expected on pan gestures.
func (g gesturePan) SetOrientation(orientation Orientation) {
	var _arg0 *C.GtkGesturePan // out
	var _arg1 C.GtkOrientation // out

	_arg0 = (*C.GtkGesturePan)(unsafe.Pointer(g.Native()))
	_arg1 = (C.GtkOrientation)(orientation)

	C.gtk_gesture_pan_set_orientation(_arg0, _arg1)
}
