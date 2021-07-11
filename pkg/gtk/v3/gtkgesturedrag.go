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
		{T: externglib.Type(C.gtk_gesture_drag_get_type()), F: marshalGestureDragger},
	})
}

// GestureDragger describes GestureDrag's methods.
type GestureDragger interface {
	// Offset: if the @gesture is active, this function returns true and fills
	// in @x and @y with the coordinates of the current point, as an offset to
	// the starting drag point.
	Offset() (x float64, y float64, ok bool)
	// StartPoint: if the @gesture is active, this function returns true and
	// fills in @x and @y with the drag start coordinates, in window-relative
	// coordinates.
	StartPoint() (x float64, y float64, ok bool)
}

// GestureDrag is a Gesture implementation that recognizes drag operations. The
// drag operation itself can be tracked throught the GestureDrag::drag-begin,
// GestureDrag::drag-update and GestureDrag::drag-end signals, or the relevant
// coordinates be extracted through gtk_gesture_drag_get_offset() and
// gtk_gesture_drag_get_start_point().
type GestureDrag struct {
	GestureSingle
}

var (
	_ GestureDragger  = (*GestureDrag)(nil)
	_ gextras.Nativer = (*GestureDrag)(nil)
)

func wrapGestureDrag(obj *externglib.Object) GestureDragger {
	return &GestureDrag{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureDragger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGestureDrag(obj), nil
}

// NewGestureDrag returns a newly created Gesture that recognizes drags.
func NewGestureDrag(widget Widgetter) *GestureDrag {
	var _arg1 *C.GtkWidget  // out
	var _cret *C.GtkGesture // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))

	_cret = C.gtk_gesture_drag_new(_arg1)

	var _gestureDrag *GestureDrag // out

	_gestureDrag = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*GestureDrag)

	return _gestureDrag
}

// Offset: if the @gesture is active, this function returns true and fills in @x
// and @y with the coordinates of the current point, as an offset to the
// starting drag point.
func (gesture *GestureDrag) Offset() (x float64, y float64, ok bool) {
	var _arg0 *C.GtkGestureDrag // out
	var _arg1 C.gdouble         // in
	var _arg2 C.gdouble         // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkGestureDrag)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_drag_get_offset(_arg0, &_arg1, &_arg2)

	var _x float64 // out
	var _y float64 // out
	var _ok bool   // out

	_x = float64(_arg1)
	_y = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _x, _y, _ok
}

// StartPoint: if the @gesture is active, this function returns true and fills
// in @x and @y with the drag start coordinates, in window-relative coordinates.
func (gesture *GestureDrag) StartPoint() (x float64, y float64, ok bool) {
	var _arg0 *C.GtkGestureDrag // out
	var _arg1 C.gdouble         // in
	var _arg2 C.gdouble         // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkGestureDrag)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_drag_get_start_point(_arg0, &_arg1, &_arg2)

	var _x float64 // out
	var _y float64 // out
	var _ok bool   // out

	_x = float64(_arg1)
	_y = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _x, _y, _ok
}
