// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gtk4_EventControllerLegacy_ConnectEvent(gpointer, GdkEvent*, guintptr);
import "C"

// glib.Type values for gtkeventcontrollerlegacy.go.
var GTypeEventControllerLegacy = coreglib.Type(C.gtk_event_controller_legacy_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeEventControllerLegacy, F: marshalEventControllerLegacy},
	})
}

// EventControllerLegacyOverrider contains methods that are overridable.
type EventControllerLegacyOverrider interface {
}

// EventControllerLegacy: GtkEventControllerLegacy is an event controller that
// provides raw access to the event stream.
//
// It should only be used as a last resort if none of the other event
// controllers or gestures do the job.
type EventControllerLegacy struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*EventControllerLegacy)(nil)
)

func classInitEventControllerLegacier(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapEventControllerLegacy(obj *coreglib.Object) *EventControllerLegacy {
	return &EventControllerLegacy{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerLegacy(p uintptr) (interface{}, error) {
	return wrapEventControllerLegacy(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_EventControllerLegacy_ConnectEvent
func _gotk4_gtk4_EventControllerLegacy_ConnectEvent(arg0 C.gpointer, arg1 *C.GdkEvent, arg2 C.guintptr) (cret C.gboolean) {
	var f func(event gdk.Eventer) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(event gdk.Eventer) (ok bool))
	}

	var _event gdk.Eventer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Eventer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.Eventer)
			return ok
		})
		rv, ok := casted.(gdk.Eventer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Eventer")
		}
		_event = rv
	}

	ok := f(_event)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectEvent is emitted for each GDK event delivered to controller.
func (v *EventControllerLegacy) ConnectEvent(f func(event gdk.Eventer) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "event", false, unsafe.Pointer(C._gotk4_gtk4_EventControllerLegacy_ConnectEvent), f)
}

// NewEventControllerLegacy creates a new legacy event controller.
//
// The function returns the following values:
//
//    - eventControllerLegacy: newly created event controller.
//
func NewEventControllerLegacy() *EventControllerLegacy {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "EventControllerLegacy").InvokeMethod("new_EventControllerLegacy", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _eventControllerLegacy *EventControllerLegacy // out

	_eventControllerLegacy = wrapEventControllerLegacy(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerLegacy
}
