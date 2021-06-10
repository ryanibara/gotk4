// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_focus_get_type()), F: marshalEventControllerFocus},
	})
}

// EventControllerFocus: `GtkEventControllerFocus` is an event controller to
// keep track of keyboard focus.
//
// The event controller offers [signal@Gtk.EventControllerFocus::enter] and
// [signal@Gtk.EventControllerFocus::leave] signals, as well as
// [property@Gtk.EventControllerFocus:is-focus] and
// [property@Gtk.EventControllerFocus:contains-focus] properties which are
// updated to reflect focus changes inside the widget hierarchy that is rooted
// at the controllers widget.
type EventControllerFocus interface {
	EventController

	// ContainsFocus returns true if focus is within @self or one of its
	// children.
	ContainsFocus() bool
	// IsFocus returns true if focus is within @self, but not one of its
	// children.
	IsFocus() bool
}

// eventControllerFocus implements the EventControllerFocus interface.
type eventControllerFocus struct {
	EventController
}

var _ EventControllerFocus = (*eventControllerFocus)(nil)

// WrapEventControllerFocus wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventControllerFocus(obj *externglib.Object) EventControllerFocus {
	return EventControllerFocus{
		EventController: WrapEventController(obj),
	}
}

func marshalEventControllerFocus(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventControllerFocus(obj), nil
}

// ContainsFocus returns true if focus is within @self or one of its
// children.
func (s eventControllerFocus) ContainsFocus() bool {
	var _arg0 *C.GtkEventControllerFocus

	_arg0 = (*C.GtkEventControllerFocus)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	_cret = C.gtk_event_controller_focus_contains_focus(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsFocus returns true if focus is within @self, but not one of its
// children.
func (s eventControllerFocus) IsFocus() bool {
	var _arg0 *C.GtkEventControllerFocus

	_arg0 = (*C.GtkEventControllerFocus)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	_cret = C.gtk_event_controller_focus_is_focus(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}
