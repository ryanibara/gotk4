// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.gtk_event_controller_motion_get_type()), F: marshalEventControllerMotion},
	})
}

// EventControllerMotion is an event controller meant for situations where you
// need to track the position of the pointer.
//
// This object was added in 3.24.
type EventControllerMotion interface {
	EventController
}

// eventControllerMotion implements the EventControllerMotion interface.
type eventControllerMotion struct {
	EventController
}

var _ EventControllerMotion = (*eventControllerMotion)(nil)

// WrapEventControllerMotion wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventControllerMotion(obj *externglib.Object) EventControllerMotion {
	return EventControllerMotion{
		EventController: WrapEventController(obj),
	}
}

func marshalEventControllerMotion(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventControllerMotion(obj), nil
}

// NewEventControllerMotion constructs a class EventControllerMotion.
func NewEventControllerMotion(widget Widget) EventControllerMotion {
	var arg1 *C.GtkWidget

	arg1 = (*C.GtkWidget)(widget.Native())

	ret := C.gtk_event_controller_motion_new(arg1)

	var ret0 EventControllerMotion

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(EventControllerMotion)

	return ret0
}
