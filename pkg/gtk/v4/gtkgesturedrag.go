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
// extern void _gotk4_gtk4_GestureDrag_ConnectDragBegin(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk4_GestureDrag_ConnectDragEnd(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk4_GestureDrag_ConnectDragUpdate(gpointer, gdouble, gdouble, guintptr);
import "C"

// glib.Type values for gtkgesturedrag.go.
var GTypeGestureDrag = externglib.Type(C.gtk_gesture_drag_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeGestureDrag, F: marshalGestureDrag},
	})
}

// GestureDragOverrider contains methods that are overridable.
type GestureDragOverrider interface {
	externglib.Objector
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

func classInitGestureDragger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

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

func marshalGestureDrag(p uintptr) (interface{}, error) {
	return wrapGestureDrag(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_GestureDrag_ConnectDragBegin
func _gotk4_gtk4_GestureDrag_ConnectDragBegin(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(startX, startY float64)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(startX, startY float64))
	}

	var _startX float64 // out
	var _startY float64 // out

	_startX = float64(arg1)
	_startY = float64(arg2)

	f(_startX, _startY)
}

// ConnectDragBegin is emitted whenever dragging starts.
func (gesture *GestureDrag) ConnectDragBegin(f func(startX, startY float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(gesture, "drag-begin", false, unsafe.Pointer(C._gotk4_gtk4_GestureDrag_ConnectDragBegin), f)
}

//export _gotk4_gtk4_GestureDrag_ConnectDragEnd
func _gotk4_gtk4_GestureDrag_ConnectDragEnd(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(offsetX, offsetY float64)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(offsetX, offsetY float64))
	}

	var _offsetX float64 // out
	var _offsetY float64 // out

	_offsetX = float64(arg1)
	_offsetY = float64(arg2)

	f(_offsetX, _offsetY)
}

// ConnectDragEnd is emitted whenever the dragging is finished.
func (gesture *GestureDrag) ConnectDragEnd(f func(offsetX, offsetY float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(gesture, "drag-end", false, unsafe.Pointer(C._gotk4_gtk4_GestureDrag_ConnectDragEnd), f)
}

//export _gotk4_gtk4_GestureDrag_ConnectDragUpdate
func _gotk4_gtk4_GestureDrag_ConnectDragUpdate(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(offsetX, offsetY float64)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(offsetX, offsetY float64))
	}

	var _offsetX float64 // out
	var _offsetY float64 // out

	_offsetX = float64(arg1)
	_offsetY = float64(arg2)

	f(_offsetX, _offsetY)
}

// ConnectDragUpdate is emitted whenever the dragging point moves.
func (gesture *GestureDrag) ConnectDragUpdate(f func(offsetX, offsetY float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(gesture, "drag-update", false, unsafe.Pointer(C._gotk4_gtk4_GestureDrag_ConnectDragUpdate), f)
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

	_arg0 = (*C.GtkGestureDrag)(unsafe.Pointer(externglib.InternObject(gesture).Native()))

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

	_arg0 = (*C.GtkGestureDrag)(unsafe.Pointer(externglib.InternObject(gesture).Native()))

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
