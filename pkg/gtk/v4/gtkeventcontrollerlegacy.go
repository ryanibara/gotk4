// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_legacy_get_type()), F: marshalEventControllerLegacier},
	})
}

// EventControllerLegacy: GtkEventControllerLegacy is an event controller that
// provides raw access to the event stream.
//
// It should only be used as a last resort if none of the other event
// controllers or gestures do the job.
type EventControllerLegacy struct {
	EventController
}

func wrapEventControllerLegacy(obj *externglib.Object) *EventControllerLegacy {
	return &EventControllerLegacy{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerLegacier(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEventControllerLegacy(obj), nil
}

// NewEventControllerLegacy creates a new legacy event controller.
func NewEventControllerLegacy() *EventControllerLegacy {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_event_controller_legacy_new()

	var _eventControllerLegacy *EventControllerLegacy // out

	_eventControllerLegacy = wrapEventControllerLegacy(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerLegacy
}

func (*EventControllerLegacy) privateEventControllerLegacy() {}
