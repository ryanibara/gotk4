// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_event_controller_motion_get_type()), F: marshalEventControllerMotioner},
	})
}

// EventControllerMotioner describes EventControllerMotion's methods.
type EventControllerMotioner interface {
	// ContainsPointer returns if a pointer is within @self or one of its
	// children.
	ContainsPointer() bool
	// IsPointer returns if a pointer is within @self, but not one of its
	// children.
	IsPointer() bool
}

// EventControllerMotion: `GtkEventControllerMotion` is an event controller
// tracking the pointer position.
//
// The event controller offers [signal@Gtk.EventControllerMotion::enter] and
// [signal@Gtk.EventControllerMotion::leave] signals, as well as
// [property@Gtk.EventControllerMotion:is-pointer] and
// [property@Gtk.EventControllerMotion:contains-pointer] properties which are
// updated to reflect changes in the pointer position as it moves over the
// widget.
type EventControllerMotion struct {
	EventController
}

var (
	_ EventControllerMotioner = (*EventControllerMotion)(nil)
	_ gextras.Nativer         = (*EventControllerMotion)(nil)
)

func wrapEventControllerMotion(obj *externglib.Object) EventControllerMotioner {
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
// motion events.
func NewEventControllerMotion() *EventControllerMotion {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_event_controller_motion_new()

	var _eventControllerMotion *EventControllerMotion // out

	_eventControllerMotion = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*EventControllerMotion)

	return _eventControllerMotion
}

// ContainsPointer returns if a pointer is within @self or one of its children.
func (self *EventControllerMotion) ContainsPointer() bool {
	var _arg0 *C.GtkEventControllerMotion // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GtkEventControllerMotion)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_event_controller_motion_contains_pointer(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsPointer returns if a pointer is within @self, but not one of its children.
func (self *EventControllerMotion) IsPointer() bool {
	var _arg0 *C.GtkEventControllerMotion // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GtkEventControllerMotion)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_event_controller_motion_is_pointer(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
