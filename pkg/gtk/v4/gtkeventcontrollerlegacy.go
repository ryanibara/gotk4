// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_legacy_get_type()), F: marshalEventControllerLegacy},
	})
}

// EventControllerLegacy: `GtkEventControllerLegacy` is an event controller that
// provides raw access to the event stream.
//
// It should only be used as a last resort if none of the other event
// controllers or gestures do the job.
type EventControllerLegacy interface {
	EventController
}

// eventControllerLegacy implements the EventControllerLegacy class.
type eventControllerLegacy struct {
	EventController
}

// WrapEventControllerLegacy wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventControllerLegacy(obj *externglib.Object) EventControllerLegacy {
	return eventControllerLegacy{
		EventController: WrapEventController(obj),
	}
}

func marshalEventControllerLegacy(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventControllerLegacy(obj), nil
}

func NewEventControllerLegacy() EventControllerLegacy {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_event_controller_legacy_new()

	var _eventControllerLegacy EventControllerLegacy // out

	_eventControllerLegacy = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(EventControllerLegacy)

	return _eventControllerLegacy
}
