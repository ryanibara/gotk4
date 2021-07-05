// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_gesture_zoom_get_type()), F: marshalGestureZoom},
	})
}

// GestureZoom is a Gesture implementation able to recognize pinch/zoom
// gestures, whenever the distance between both tracked sequences changes, the
// GestureZoom::scale-changed signal is emitted to report the scale factor.
type GestureZoom interface {
	Gesture

	// ScaleDelta: if @gesture is active, this function returns the zooming
	// difference since the gesture was recognized (hence the starting point is
	// considered 1:1). If @gesture is not active, 1 is returned.
	ScaleDelta() float64
}

// gestureZoom implements the GestureZoom class.
type gestureZoom struct {
	Gesture
}

// WrapGestureZoom wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureZoom(obj *externglib.Object) GestureZoom {
	return gestureZoom{
		Gesture: WrapGesture(obj),
	}
}

func marshalGestureZoom(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureZoom(obj), nil
}

// NewGestureZoom returns a newly created Gesture that recognizes zoom in/out
// gestures (usually known as pinch/zoom).
func NewGestureZoom(widget Widget) GestureZoom {
	var _arg1 *C.GtkWidget  // out
	var _cret *C.GtkGesture // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_gesture_zoom_new(_arg1)

	var _gestureZoom GestureZoom // out

	_gestureZoom = WrapGestureZoom(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureZoom
}

func (g gestureZoom) ScaleDelta() float64 {
	var _arg0 *C.GtkGestureZoom // out
	var _cret C.gdouble         // in

	_arg0 = (*C.GtkGestureZoom)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_zoom_get_scale_delta(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}
