// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_motion_get_type()), F: marshalEventControllerMotioner},
	})
}

// EventControllerMotioner describes EventControllerMotion's methods.
type EventControllerMotioner interface {
	privateEventControllerMotion()
}

// EventControllerMotion is an event controller meant for situations where you
// need to track the position of the pointer.
//
// This object was added in 3.24.
type EventControllerMotion struct {
	EventController
}

var (
	_ EventControllerMotioner = (*EventControllerMotion)(nil)
	_ gextras.Nativer         = (*EventControllerMotion)(nil)
)

func wrapEventControllerMotion(obj *externglib.Object) *EventControllerMotion {
	return &EventControllerMotion{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerMotioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEventControllerMotion(obj), nil
}

// NewEventControllerMotion creates a new event controller that will handle
// motion events for the given widget.
func NewEventControllerMotion(widget Widgeter) *EventControllerMotion {
	var _arg1 *C.GtkWidget          // out
	var _cret *C.GtkEventController // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))

	_cret = C.gtk_event_controller_motion_new(_arg1)

	var _eventControllerMotion *EventControllerMotion // out

	_eventControllerMotion = wrapEventControllerMotion(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerMotion
}

func (*EventControllerMotion) privateEventControllerMotion() {}
