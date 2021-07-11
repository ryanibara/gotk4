// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_zoom_get_type()), F: marshalGestureZoomer},
	})
}

// GestureZoomer describes GestureZoom's methods.
type GestureZoomer interface {
	// ScaleDelta: if @gesture is active, this function returns the zooming
	// difference since the gesture was recognized (hence the starting point is
	// considered 1:1).
	ScaleDelta() float64
}

// GestureZoom is a Gesture implementation able to recognize pinch/zoom
// gestures, whenever the distance between both tracked sequences changes, the
// GestureZoom::scale-changed signal is emitted to report the scale factor.
type GestureZoom struct {
	Gesture
}

var (
	_ GestureZoomer   = (*GestureZoom)(nil)
	_ gextras.Nativer = (*GestureZoom)(nil)
)

func wrapGestureZoom(obj *externglib.Object) GestureZoomer {
	return &GestureZoom{
		Gesture: Gesture{
			EventController: EventController{
				Object: obj,
			},
		},
	}
}

func marshalGestureZoomer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGestureZoom(obj), nil
}

// NewGestureZoom returns a newly created Gesture that recognizes zoom in/out
// gestures (usually known as pinch/zoom).
func NewGestureZoom(widget Widgetter) *GestureZoom {
	var _arg1 *C.GtkWidget  // out
	var _cret *C.GtkGesture // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))

	_cret = C.gtk_gesture_zoom_new(_arg1)

	var _gestureZoom *GestureZoom // out

	_gestureZoom = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*GestureZoom)

	return _gestureZoom
}

// ScaleDelta: if @gesture is active, this function returns the zooming
// difference since the gesture was recognized (hence the starting point is
// considered 1:1). If @gesture is not active, 1 is returned.
func (gesture *GestureZoom) ScaleDelta() float64 {
	var _arg0 *C.GtkGestureZoom // out
	var _cret C.gdouble         // in

	_arg0 = (*C.GtkGestureZoom)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_zoom_get_scale_delta(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}
