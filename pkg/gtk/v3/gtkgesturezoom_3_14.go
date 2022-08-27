// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

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
