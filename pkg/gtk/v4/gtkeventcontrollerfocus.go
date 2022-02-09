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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_focus_get_type()), F: marshalEventControllerFocusser},
	})
}

// EventControllerFocus: GtkEventControllerFocus is an event controller to keep
// track of keyboard focus.
//
// The event controller offers gtk.EventControllerFocus::enter and
// gtk.EventControllerFocus::leave signals, as well as
// gtk.EventControllerFocus:is-focus and gtk.EventControllerFocus:contains-focus
// properties which are updated to reflect focus changes inside the widget
// hierarchy that is rooted at the controllers widget.
type EventControllerFocus struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*EventControllerFocus)(nil)
)

func wrapEventControllerFocus(obj *externglib.Object) *EventControllerFocus {
	return &EventControllerFocus{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerFocusser(p uintptr) (interface{}, error) {
	return wrapEventControllerFocus(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectEnter: emitted whenever the focus enters into the widget or one of its
// descendents.
//
// Note that this means you may not get an ::enter signal even though the widget
// becomes the focus location, in certain cases (such as when the focus moves
// from a descendent of the widget to the widget itself). If you are interested
// in these cases, you can monitor the gtk.EventControllerFocus:is-focus
// property for changes.
func (self *EventControllerFocus) ConnectEnter(f func()) externglib.SignalHandle {
	return self.Connect("enter", externglib.GeneratedClosure{Func: f})
}

// ConnectLeave: emitted whenever the focus leaves the widget hierarchy that is
// rooted at the widget that the controller is attached to.
//
// Note that this means you may not get a ::leave signal even though the focus
// moves away from the widget, in certain cases (such as when the focus moves
// from the widget to a descendent). If you are interested in these cases, you
// can monitor the gtk.EventControllerFocus:is-focus property for changes.
func (self *EventControllerFocus) ConnectLeave(f func()) externglib.SignalHandle {
	return self.Connect("leave", externglib.GeneratedClosure{Func: f})
}

// NewEventControllerFocus creates a new event controller that will handle focus
// events.
//
// The function returns the following values:
//
//    - eventControllerFocus: new GtkEventControllerFocus.
//
func NewEventControllerFocus() *EventControllerFocus {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_event_controller_focus_new()

	var _eventControllerFocus *EventControllerFocus // out

	_eventControllerFocus = wrapEventControllerFocus(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerFocus
}

// ContainsFocus returns TRUE if focus is within self or one of its children.
//
// The function returns the following values:
//
//    - ok: TRUE if focus is within self or one of its children.
//
func (self *EventControllerFocus) ContainsFocus() bool {
	var _arg0 *C.GtkEventControllerFocus // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkEventControllerFocus)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_event_controller_focus_contains_focus(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsFocus returns TRUE if focus is within self, but not one of its children.
//
// The function returns the following values:
//
//    - ok: TRUE if focus is within self, but not one of its children.
//
func (self *EventControllerFocus) IsFocus() bool {
	var _arg0 *C.GtkEventControllerFocus // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkEventControllerFocus)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_event_controller_focus_is_focus(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
