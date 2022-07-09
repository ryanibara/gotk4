// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_GestureRotate_ConnectAngleChanged(gpointer, gdouble, gdouble, guintptr);
import "C"

// GTypeGestureRotate returns the GType for the type GestureRotate.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeGestureRotate() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "GestureRotate").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalGestureRotate)
	return gtype
}

// GestureRotate is a Gesture implementation able to recognize 2-finger
// rotations, whenever the angle between both handled sequences changes, the
// GestureRotate::angle-changed signal is emitted.
type GestureRotate struct {
	_ [0]func() // equal guard
	Gesture
}

var (
	_ Gesturer = (*GestureRotate)(nil)
)

func wrapGestureRotate(obj *coreglib.Object) *GestureRotate {
	return &GestureRotate{
		Gesture: Gesture{
			EventController: EventController{
				Object: obj,
			},
		},
	}
}

func marshalGestureRotate(p uintptr) (interface{}, error) {
	return wrapGestureRotate(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_GestureRotate_ConnectAngleChanged
func _gotk4_gtk3_GestureRotate_ConnectAngleChanged(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(angle, angleDelta float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(angle, angleDelta float64))
	}

	var _angle float64      // out
	var _angleDelta float64 // out

	_angle = float64(arg1)
	_angleDelta = float64(arg2)

	f(_angle, _angleDelta)
}

// ConnectAngleChanged: this signal is emitted when the angle between both
// tracked points changes.
func (gesture *GestureRotate) ConnectAngleChanged(f func(angle, angleDelta float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "angle-changed", false, unsafe.Pointer(C._gotk4_gtk3_GestureRotate_ConnectAngleChanged), f)
}

// NewGestureRotate returns a newly created Gesture that recognizes 2-touch
// rotation gestures.
//
// The function takes the following parameters:
//
//    - widget: Widget.
//
// The function returns the following values:
//
//    - gestureRotate: newly created GestureRotate.
//
func NewGestureRotate(widget Widgetter) *GestureRotate {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_gret := girepository.MustFind("Gtk", "GestureRotate").InvokeMethod("new_GestureRotate", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(widget)

	var _gestureRotate *GestureRotate // out

	_gestureRotate = wrapGestureRotate(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureRotate
}

// AngleDelta: if gesture is active, this function returns the angle difference
// in radians since the gesture was first recognized. If gesture is not active,
// 0 is returned.
//
// The function returns the following values:
//
//    - gdouble: angle delta in radians.
//
func (gesture *GestureRotate) AngleDelta() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_gret := girepository.MustFind("Gtk", "GestureRotate").InvokeMethod("get_angle_delta", _args[:], nil)
	_cret = *(*C.gdouble)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.gdouble)(unsafe.Pointer(&_cret)))

	return _gdouble
}
