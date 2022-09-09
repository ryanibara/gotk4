// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_GestureZoom_ConnectScaleChanged(gpointer, gdouble, guintptr);
import "C"

// GType values.
var (
	GTypeGestureZoom = coreglib.Type(C.gtk_gesture_zoom_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGestureZoom, F: marshalGestureZoom},
	})
}

// GestureZoom is a Gesture implementation able to recognize pinch/zoom
// gestures, whenever the distance between both tracked sequences changes, the
// GestureZoom::scale-changed signal is emitted to report the scale factor.
type GestureZoom struct {
	_ [0]func() // equal guard
	Gesture
}

var (
	_ Gesturer = (*GestureZoom)(nil)
)

func wrapGestureZoom(obj *coreglib.Object) *GestureZoom {
	return &GestureZoom{
		Gesture: Gesture{
			EventController: EventController{
				Object: obj,
			},
		},
	}
}

func marshalGestureZoom(p uintptr) (interface{}, error) {
	return wrapGestureZoom(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectScaleChanged: this signal is emitted whenever the distance between
// both tracked sequences changes.
func (gesture *GestureZoom) ConnectScaleChanged(f func(scale float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "scale-changed", false, unsafe.Pointer(C._gotk4_gtk3_GestureZoom_ConnectScaleChanged), f)
}

// NewGestureZoom returns a newly created Gesture that recognizes zoom in/out
// gestures (usually known as pinch/zoom).
//
// The function takes the following parameters:
//
//    - widget: Widget.
//
// The function returns the following values:
//
//    - gestureZoom: newly created GestureZoom.
//
func NewGestureZoom(widget Widgetter) *GestureZoom {
	var _arg1 *C.GtkWidget  // out
	var _cret *C.GtkGesture // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_cret = C.gtk_gesture_zoom_new(_arg1)
	runtime.KeepAlive(widget)

	var _gestureZoom *GestureZoom // out

	_gestureZoom = wrapGestureZoom(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureZoom
}

// ScaleDelta: if gesture is active, this function returns the zooming
// difference since the gesture was recognized (hence the starting point is
// considered 1:1). If gesture is not active, 1 is returned.
//
// The function returns the following values:
//
//    - gdouble: scale delta.
//
func (gesture *GestureZoom) ScaleDelta() float64 {
	var _arg0 *C.GtkGestureZoom // out
	var _cret C.gdouble         // in

	_arg0 = (*C.GtkGestureZoom)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_cret = C.gtk_gesture_zoom_get_scale_delta(_arg0)
	runtime.KeepAlive(gesture)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}
