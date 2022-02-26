// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_GesturePan_ConnectPan(gpointer, GtkPanDirection, gdouble, guintptr);
import "C"

// glib.Type values for gtkgesturepan.go.
var GTypeGesturePan = externglib.Type(C.gtk_gesture_pan_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeGesturePan, F: marshalGesturePan},
	})
}

// GesturePanOverrider contains methods that are overridable.
type GesturePanOverrider interface {
}

// GesturePan: GtkGesturePan is a GtkGesture for pan gestures.
//
// These are drags that are locked to happen along one axis. The axis that a
// GtkGesturePan handles is defined at construct time, and can be changed
// through gtk.GesturePan.SetOrientation().
//
// When the gesture starts to be recognized, GtkGesturePan will attempt to
// determine as early as possible whether the sequence is moving in the expected
// direction, and denying the sequence if this does not happen.
//
// Once a panning gesture along the expected axis is recognized, the
// gtk.GesturePan::pan signal will be emitted as input events are received,
// containing the offset in the given axis.
type GesturePan struct {
	_ [0]func() // equal guard
	GestureDrag
}

var (
	_ Gesturer = (*GesturePan)(nil)
)

func classInitGesturePanner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGesturePan(obj *externglib.Object) *GesturePan {
	return &GesturePan{
		GestureDrag: GestureDrag{
			GestureSingle: GestureSingle{
				Gesture: Gesture{
					EventController: EventController{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalGesturePan(p uintptr) (interface{}, error) {
	return wrapGesturePan(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_GesturePan_ConnectPan
func _gotk4_gtk4_GesturePan_ConnectPan(arg0 C.gpointer, arg1 C.GtkPanDirection, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(direction PanDirection, offset float64)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(direction PanDirection, offset float64))
	}

	var _direction PanDirection // out
	var _offset float64         // out

	_direction = PanDirection(arg1)
	_offset = float64(arg2)

	f(_direction, _offset)
}

// ConnectPan is emitted once a panning gesture along the expected axis is
// detected.
func (gesture *GesturePan) ConnectPan(f func(direction PanDirection, offset float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(gesture, "pan", false, unsafe.Pointer(C._gotk4_gtk4_GesturePan_ConnectPan), f)
}

// NewGesturePan returns a newly created GtkGesture that recognizes pan
// gestures.
//
// The function takes the following parameters:
//
//    - orientation: expected orientation.
//
// The function returns the following values:
//
//    - gesturePan: newly created GtkGesturePan.
//
func NewGesturePan(orientation Orientation) *GesturePan {
	var _arg1 C.GtkOrientation // out
	var _cret *C.GtkGesture    // in

	_arg1 = C.GtkOrientation(orientation)

	_cret = C.gtk_gesture_pan_new(_arg1)
	runtime.KeepAlive(orientation)

	var _gesturePan *GesturePan // out

	_gesturePan = wrapGesturePan(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gesturePan
}

// Orientation returns the orientation of the pan gestures that this gesture
// expects.
//
// The function returns the following values:
//
//    - orientation: expected orientation for pan gestures.
//
func (gesture *GesturePan) Orientation() Orientation {
	var _arg0 *C.GtkGesturePan // out
	var _cret C.GtkOrientation // in

	_arg0 = (*C.GtkGesturePan)(unsafe.Pointer(externglib.InternObject(gesture).Native()))

	_cret = C.gtk_gesture_pan_get_orientation(_arg0)
	runtime.KeepAlive(gesture)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

// SetOrientation sets the orientation to be expected on pan gestures.
//
// The function takes the following parameters:
//
//    - orientation: expected orientation.
//
func (gesture *GesturePan) SetOrientation(orientation Orientation) {
	var _arg0 *C.GtkGesturePan // out
	var _arg1 C.GtkOrientation // out

	_arg0 = (*C.GtkGesturePan)(unsafe.Pointer(externglib.InternObject(gesture).Native()))
	_arg1 = C.GtkOrientation(orientation)

	C.gtk_gesture_pan_set_orientation(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(orientation)
}
