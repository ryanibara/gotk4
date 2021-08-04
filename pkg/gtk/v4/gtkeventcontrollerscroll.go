// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"strings"
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
		{T: externglib.Type(C.gtk_event_controller_scroll_flags_get_type()), F: marshalEventControllerScrollFlags},
		{T: externglib.Type(C.gtk_event_controller_scroll_get_type()), F: marshalEventControllerScroller},
	})
}

// EventControllerScrollFlags describes the behavior of a
// GtkEventControllerScroll.
type EventControllerScrollFlags int

const (
	// EventControllerScrollNone: don't emit scroll.
	EventControllerScrollNone EventControllerScrollFlags = 0b0
	// EventControllerScrollVertical: emit scroll with vertical deltas.
	EventControllerScrollVertical EventControllerScrollFlags = 0b1
	// EventControllerScrollHorizontal: emit scroll with horizontal deltas.
	EventControllerScrollHorizontal EventControllerScrollFlags = 0b10
	// EventControllerScrollDiscrete: only emit deltas that are multiples of 1.
	EventControllerScrollDiscrete EventControllerScrollFlags = 0b100
	// EventControllerScrollKinetic: emit ::decelerate after continuous scroll
	// finishes.
	EventControllerScrollKinetic EventControllerScrollFlags = 0b1000
	// EventControllerScrollBothAxes: emit scroll on both axes.
	EventControllerScrollBothAxes EventControllerScrollFlags = 0b11
)

func marshalEventControllerScrollFlags(p uintptr) (interface{}, error) {
	return EventControllerScrollFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the names in string for EventControllerScrollFlags.
func (e EventControllerScrollFlags) String() string {
	if e == 0 {
		return "EventControllerScrollFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(176)

	for e != 0 {
		next := e & (e - 1)
		bit := e - next

		switch bit {
		case EventControllerScrollNone:
			builder.WriteString("None|")
		case EventControllerScrollVertical:
			builder.WriteString("Vertical|")
		case EventControllerScrollHorizontal:
			builder.WriteString("Horizontal|")
		case EventControllerScrollDiscrete:
			builder.WriteString("Discrete|")
		case EventControllerScrollKinetic:
			builder.WriteString("Kinetic|")
		case EventControllerScrollBothAxes:
			builder.WriteString("BothAxes|")
		default:
			builder.WriteString(fmt.Sprintf("EventControllerScrollFlags(0b%b)|", bit))
		}

		e = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// EventControllerScroll: GtkEventControllerScroll is an event controller that
// handles scroll events.
//
// It is capable of handling both discrete and continuous scroll events from
// mice or touchpads, abstracting them both with the
// gtk.EventControllerScroll::scroll signal. Deltas in the discrete case are
// multiples of 1.
//
// In the case of continuous scroll events, GtkEventControllerScroll encloses
// all gtk.EventControllerScroll::scroll emissions between two
// gtk.EventControllerScroll::scroll-begin and
// gtk.EventControllerScroll::scroll-end signals.
//
// The behavior of the event controller can be modified by the flags given at
// creation time, or modified at a later point through
// gtk.EventControllerScroll.SetFlags() (e.g. because the scrolling conditions
// of the widget changed).
//
// The controller can be set up to emit motion for either/both vertical and
// horizontal scroll events through GTK_EVENT_CONTROLLER_SCROLL_VERTICAL,
// GTK_EVENT_CONTROLLER_SCROLL_HORIZONTAL and
// GTK_EVENT_CONTROLLER_SCROLL_BOTH_AXES. If any axis is disabled, the
// respective gtk.EventControllerScroll::scroll delta will be 0. Vertical scroll
// events will be translated to horizontal motion for the devices incapable of
// horizontal scrolling.
//
// The event controller can also be forced to emit discrete events on all
// devices through GTK_EVENT_CONTROLLER_SCROLL_DISCRETE. This can be used to
// implement discrete actions triggered through scroll events (e.g. switching
// across combobox options).
//
// The GTK_EVENT_CONTROLLER_SCROLL_KINETIC flag toggles the emission of the
// gtk.EventControllerScroll::decelerate signal, emitted at the end of scrolling
// with two X/Y velocity arguments that are consistent with the motion that was
// received.
type EventControllerScroll struct {
	EventController
}

func wrapEventControllerScroll(obj *externglib.Object) *EventControllerScroll {
	return &EventControllerScroll{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerScroller(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEventControllerScroll(obj), nil
}

// NewEventControllerScroll creates a new event controller that will handle
// scroll events.
func NewEventControllerScroll(flags EventControllerScrollFlags) *EventControllerScroll {
	var _arg1 C.GtkEventControllerScrollFlags // out
	var _cret *C.GtkEventController           // in

	_arg1 = C.GtkEventControllerScrollFlags(flags)

	_cret = C.gtk_event_controller_scroll_new(_arg1)

	var _eventControllerScroll *EventControllerScroll // out

	_eventControllerScroll = wrapEventControllerScroll(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerScroll
}

// Flags gets the flags conditioning the scroll controller behavior.
func (scroll *EventControllerScroll) Flags() EventControllerScrollFlags {
	var _arg0 *C.GtkEventControllerScroll     // out
	var _cret C.GtkEventControllerScrollFlags // in

	_arg0 = (*C.GtkEventControllerScroll)(unsafe.Pointer(scroll.Native()))

	_cret = C.gtk_event_controller_scroll_get_flags(_arg0)

	var _eventControllerScrollFlags EventControllerScrollFlags // out

	_eventControllerScrollFlags = EventControllerScrollFlags(_cret)

	return _eventControllerScrollFlags
}

// SetFlags sets the flags conditioning scroll controller behavior.
func (scroll *EventControllerScroll) SetFlags(flags EventControllerScrollFlags) {
	var _arg0 *C.GtkEventControllerScroll     // out
	var _arg1 C.GtkEventControllerScrollFlags // out

	_arg0 = (*C.GtkEventControllerScroll)(unsafe.Pointer(scroll.Native()))
	_arg1 = C.GtkEventControllerScrollFlags(flags)

	C.gtk_event_controller_scroll_set_flags(_arg0, _arg1)
}
