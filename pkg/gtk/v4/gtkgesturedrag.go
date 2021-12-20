// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_drag_get_type()), F: marshalGestureDragger},
	})
}

// GestureDrag: GtkGestureDrag is a GtkGesture implementation for drags.
//
// The drag operation itself can be tracked throughout the
// gtk.GestureDrag::drag-begin, gtk.GestureDrag::drag-update and
// gtk.GestureDrag::drag-end signals, and the relevant coordinates can be
// extracted through gtk.GestureDrag.GetOffset() and
// gtk.GestureDrag.GetStartPoint().
type GestureDrag struct {
	_ [0]func() // equal guard
	GestureSingle
}

var (
	_ Gesturer = (*GestureDrag)(nil)
)

func wrapGestureDrag(obj *externglib.Object) *GestureDrag {
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
	return wrapGestureDrag(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectDragBegin: emitted whenever dragging starts.
func (gesture *GestureDrag) ConnectDragBegin(f func(startX, startY float64)) externglib.SignalHandle {
	return gesture.Connect("drag-begin", f)
}

// ConnectDragEnd: emitted whenever the dragging is finished.
func (gesture *GestureDrag) ConnectDragEnd(f func(offsetX, offsetY float64)) externglib.SignalHandle {
	return gesture.Connect("drag-end", f)
}

// ConnectDragUpdate: emitted whenever the dragging point moves.
func (gesture *GestureDrag) ConnectDragUpdate(f func(offsetX, offsetY float64)) externglib.SignalHandle {
	return gesture.Connect("drag-update", f)
}

// NewGestureDrag returns a newly created GtkGesture that recognizes drags.
//
// The function returns the following values:
//
//    - gestureDrag: newly created GtkGestureDrag.
//
func NewGestureDrag() *GestureDrag {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_drag_new()

	var _gestureDrag *GestureDrag // out

	_gestureDrag = wrapGestureDrag(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureDrag
}

// Offset gets the offset from the start point.
//
// If the gesture is active, this function returns TRUE and fills in x and y
// with the coordinates of the current point, as an offset to the starting drag
// point.
//
// The function returns the following values:
//
//    - x (optional): x offset for the current point.
//    - y (optional): y offset for the current point.
//    - ok: TRUE if the gesture is active.
//
func (gesture *GestureDrag) Offset() (x float64, y float64, ok bool) {
	var _arg0 *C.GtkGestureDrag // out
	var _arg1 C.double          // in
	var _arg2 C.double          // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkGestureDrag)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_drag_get_offset(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(gesture)

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

// StartPoint gets the point where the drag started.
//
// If the gesture is active, this function returns TRUE and fills in x and y
// with the drag start coordinates, in surface-relative coordinates.
//
// The function returns the following values:
//
//    - x (optional): x coordinate for the drag start point.
//    - y (optional): y coordinate for the drag start point.
//    - ok: TRUE if the gesture is active.
//
func (gesture *GestureDrag) StartPoint() (x float64, y float64, ok bool) {
	var _arg0 *C.GtkGestureDrag // out
	var _arg1 C.double          // in
	var _arg2 C.double          // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkGestureDrag)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_drag_get_start_point(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(gesture)

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
