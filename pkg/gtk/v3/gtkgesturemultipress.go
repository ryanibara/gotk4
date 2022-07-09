// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_GestureMultiPress_ConnectPressed(gpointer, gint, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk3_GestureMultiPress_ConnectReleased(gpointer, gint, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk3_GestureMultiPress_ConnectStopped(gpointer, guintptr);
import "C"

// GTypeGestureMultiPress returns the GType for the type GestureMultiPress.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeGestureMultiPress() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "GestureMultiPress").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalGestureMultiPress)
	return gtype
}

// GestureMultiPress is a Gesture implementation able to recognize multiple
// clicks on a nearby zone, which can be listened for through the
// GestureMultiPress::pressed signal. Whenever time or distance between clicks
// exceed the GTK+ defaults, GestureMultiPress::stopped is emitted, and the
// click counter is reset.
//
// Callers may also restrict the area that is considered valid for a >1
// touch/button press through gtk_gesture_multi_press_set_area(), so any click
// happening outside that area is considered to be a first click of its own.
type GestureMultiPress struct {
	_ [0]func() // equal guard
	GestureSingle
}

var (
	_ Gesturer = (*GestureMultiPress)(nil)
)

func wrapGestureMultiPress(obj *coreglib.Object) *GestureMultiPress {
	return &GestureMultiPress{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureMultiPress(p uintptr) (interface{}, error) {
	return wrapGestureMultiPress(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_GestureMultiPress_ConnectPressed
func _gotk4_gtk3_GestureMultiPress_ConnectPressed(arg0 C.gpointer, arg1 C.gint, arg2 C.gdouble, arg3 C.gdouble, arg4 C.guintptr) {
	var f func(nPress int32, x, y float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(nPress int32, x, y float64))
	}

	var _nPress int32 // out
	var _x float64    // out
	var _y float64    // out

	_nPress = int32(arg1)
	_x = float64(arg2)
	_y = float64(arg3)

	f(_nPress, _x, _y)
}

// ConnectPressed: this signal is emitted whenever a button or touch press
// happens.
func (gesture *GestureMultiPress) ConnectPressed(f func(nPress int32, x, y float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "pressed", false, unsafe.Pointer(C._gotk4_gtk3_GestureMultiPress_ConnectPressed), f)
}

//export _gotk4_gtk3_GestureMultiPress_ConnectReleased
func _gotk4_gtk3_GestureMultiPress_ConnectReleased(arg0 C.gpointer, arg1 C.gint, arg2 C.gdouble, arg3 C.gdouble, arg4 C.guintptr) {
	var f func(nPress int32, x, y float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(nPress int32, x, y float64))
	}

	var _nPress int32 // out
	var _x float64    // out
	var _y float64    // out

	_nPress = int32(arg1)
	_x = float64(arg2)
	_y = float64(arg3)

	f(_nPress, _x, _y)
}

// ConnectReleased: this signal is emitted when a button or touch is released.
// n_press will report the number of press that is paired to this event, note
// that GestureMultiPress::stopped may have been emitted between the press and
// its release, n_press will only start over at the next press.
func (gesture *GestureMultiPress) ConnectReleased(f func(nPress int32, x, y float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "released", false, unsafe.Pointer(C._gotk4_gtk3_GestureMultiPress_ConnectReleased), f)
}

//export _gotk4_gtk3_GestureMultiPress_ConnectStopped
func _gotk4_gtk3_GestureMultiPress_ConnectStopped(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectStopped: this signal is emitted whenever any time/distance threshold
// has been exceeded.
func (gesture *GestureMultiPress) ConnectStopped(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "stopped", false, unsafe.Pointer(C._gotk4_gtk3_GestureMultiPress_ConnectStopped), f)
}

// NewGestureMultiPress returns a newly created Gesture that recognizes single
// and multiple presses.
//
// The function takes the following parameters:
//
//    - widget: Widget.
//
// The function returns the following values:
//
//    - gestureMultiPress: newly created GestureMultiPress.
//
func NewGestureMultiPress(widget Widgetter) *GestureMultiPress {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_gret := girepository.MustFind("Gtk", "GestureMultiPress").InvokeMethod("new_GestureMultiPress", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(widget)

	var _gestureMultiPress *GestureMultiPress // out

	_gestureMultiPress = wrapGestureMultiPress(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureMultiPress
}

// Area: if an area was set through gtk_gesture_multi_press_set_area(), this
// function will return TRUE and fill in rect with the press area. See
// gtk_gesture_multi_press_set_area() for more details on what the press area
// represents.
//
// The function returns the following values:
//
//    - rect: return location for the press area.
//    - ok: TRUE if rect was filled with the press area.
//
func (gesture *GestureMultiPress) Area() (*gdk.Rectangle, bool) {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_gret := girepository.MustFind("Gtk", "GestureMultiPress").InvokeMethod("get_area", _args[:], _outs[:])
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _rect *gdk.Rectangle // out
	var _ok bool             // out

	_rect = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _rect, _ok
}

// SetArea: if rect is non-NULL, the press area will be checked to be confined
// within the rectangle, otherwise the button count will be reset so the press
// is seen as being the first one. If rect is NULL, the area will be reset to an
// unrestricted state.
//
// Note: The rectangle is only used to determine whether any non-first click
// falls within the expected area. This is not akin to an input shape.
//
// The function takes the following parameters:
//
//    - rect (optional): rectangle to receive coordinates on.
//
func (gesture *GestureMultiPress) SetArea(rect *gdk.Rectangle) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))
	if rect != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(rect)))
	}

	girepository.MustFind("Gtk", "GestureMultiPress").InvokeMethod("set_area", _args[:], nil)

	runtime.KeepAlive(gesture)
	runtime.KeepAlive(rect)
}
