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
// extern void _gotk4_gtk4_EventControllerMotion_ConnectEnter(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk4_EventControllerMotion_ConnectLeave(gpointer, guintptr);
// extern void _gotk4_gtk4_EventControllerMotion_ConnectMotion(gpointer, gdouble, gdouble, guintptr);
import "C"

// glib.Type values for gtkeventcontrollermotion.go.
var GTypeEventControllerMotion = externglib.Type(C.gtk_event_controller_motion_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeEventControllerMotion, F: marshalEventControllerMotion},
	})
}

// EventControllerMotionOverrider contains methods that are overridable.
type EventControllerMotionOverrider interface {
}

// EventControllerMotion: GtkEventControllerMotion is an event controller
// tracking the pointer position.
//
// The event controller offers gtk.EventControllerMotion::enter and
// gtk.EventControllerMotion::leave signals, as well as
// gtk.EventControllerMotion:is-pointer and
// gtk.EventControllerMotion:contains-pointer properties which are updated to
// reflect changes in the pointer position as it moves over the widget.
type EventControllerMotion struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*EventControllerMotion)(nil)
)

func classInitEventControllerMotioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapEventControllerMotion(obj *externglib.Object) *EventControllerMotion {
	return &EventControllerMotion{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerMotion(p uintptr) (interface{}, error) {
	return wrapEventControllerMotion(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_EventControllerMotion_ConnectEnter
func _gotk4_gtk4_EventControllerMotion_ConnectEnter(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(x, y float64)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(x, y float64))
	}

	var _x float64 // out
	var _y float64 // out

	_x = float64(arg1)
	_y = float64(arg2)

	f(_x, _y)
}

// ConnectEnter signals that the pointer has entered the widget.
func (self *EventControllerMotion) ConnectEnter(f func(x, y float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "enter", false, unsafe.Pointer(C._gotk4_gtk4_EventControllerMotion_ConnectEnter), f)
}

//export _gotk4_gtk4_EventControllerMotion_ConnectLeave
func _gotk4_gtk4_EventControllerMotion_ConnectLeave(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectLeave signals that the pointer has left the widget.
func (self *EventControllerMotion) ConnectLeave(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "leave", false, unsafe.Pointer(C._gotk4_gtk4_EventControllerMotion_ConnectLeave), f)
}

//export _gotk4_gtk4_EventControllerMotion_ConnectMotion
func _gotk4_gtk4_EventControllerMotion_ConnectMotion(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(x, y float64)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(x, y float64))
	}

	var _x float64 // out
	var _y float64 // out

	_x = float64(arg1)
	_y = float64(arg2)

	f(_x, _y)
}

// ConnectMotion: emitted when the pointer moves inside the widget.
func (self *EventControllerMotion) ConnectMotion(f func(x, y float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "motion", false, unsafe.Pointer(C._gotk4_gtk4_EventControllerMotion_ConnectMotion), f)
}

// NewEventControllerMotion creates a new event controller that will handle
// motion events.
//
// The function returns the following values:
//
//    - eventControllerMotion: new GtkEventControllerMotion.
//
func NewEventControllerMotion() *EventControllerMotion {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_event_controller_motion_new()

	var _eventControllerMotion *EventControllerMotion // out

	_eventControllerMotion = wrapEventControllerMotion(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerMotion
}

// ContainsPointer returns if a pointer is within self or one of its children.
//
// The function returns the following values:
//
//    - ok: TRUE if a pointer is within self or one of its children.
//
func (self *EventControllerMotion) ContainsPointer() bool {
	var _arg0 *C.GtkEventControllerMotion // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GtkEventControllerMotion)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_event_controller_motion_contains_pointer(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsPointer returns if a pointer is within self, but not one of its children.
//
// The function returns the following values:
//
//    - ok: TRUE if a pointer is within self but not one of its children.
//
func (self *EventControllerMotion) IsPointer() bool {
	var _arg0 *C.GtkEventControllerMotion // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GtkEventControllerMotion)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_event_controller_motion_is_pointer(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
