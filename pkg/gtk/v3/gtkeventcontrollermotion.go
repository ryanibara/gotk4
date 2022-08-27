// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_EventControllerMotion_ConnectMotion(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk3_EventControllerMotion_ConnectLeave(gpointer, guintptr);
// extern void _gotk4_gtk3_EventControllerMotion_ConnectEnter(gpointer, gdouble, gdouble, guintptr);
import "C"

// GType values.
var (
	GTypeEventControllerMotion = coreglib.Type(C.gtk_event_controller_motion_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEventControllerMotion, F: marshalEventControllerMotion},
	})
}

// EventControllerMotion is an event controller meant for situations where you
// need to track the position of the pointer.
//
// This object was added in 3.24.
type EventControllerMotion struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*EventControllerMotion)(nil)
)

func wrapEventControllerMotion(obj *coreglib.Object) *EventControllerMotion {
	return &EventControllerMotion{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerMotion(p uintptr) (interface{}, error) {
	return wrapEventControllerMotion(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectEnter signals that the pointer has entered the widget.
func (v *EventControllerMotion) ConnectEnter(f func(x, y float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "enter", false, unsafe.Pointer(C._gotk4_gtk3_EventControllerMotion_ConnectEnter), f)
}

// ConnectLeave signals that pointer has left the widget.
func (v *EventControllerMotion) ConnectLeave(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "leave", false, unsafe.Pointer(C._gotk4_gtk3_EventControllerMotion_ConnectLeave), f)
}

// ConnectMotion is emitted when the pointer moves inside the widget.
func (v *EventControllerMotion) ConnectMotion(f func(x, y float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "motion", false, unsafe.Pointer(C._gotk4_gtk3_EventControllerMotion_ConnectMotion), f)
}
