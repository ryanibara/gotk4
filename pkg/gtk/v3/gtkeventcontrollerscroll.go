// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
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
		{T: externglib.Type(C.gtk_event_controller_scroll_flags_get_type()), F: marshalEventControllerScrollFlags},
		{T: externglib.Type(C.gtk_event_controller_scroll_get_type()), F: marshalEventControllerScroller},
	})
}

// EventControllerScrollFlags describes the behavior of a EventControllerScroll.
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
	// EventControllerScrollKinetic: emit EventControllerScroll::decelerate
	// after continuous scroll finishes.
	EventControllerScrollKinetic EventControllerScrollFlags = 0b1000
	// EventControllerScrollBothAxes: emit scroll on both axes.
	EventControllerScrollBothAxes EventControllerScrollFlags = 0b11
)

func marshalEventControllerScrollFlags(p uintptr) (interface{}, error) {
	return EventControllerScrollFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
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

// Has returns true if e contains other.
func (e EventControllerScrollFlags) Has(other EventControllerScrollFlags) bool {
	return (e & other) == other
}

// EventControllerScroll is an event controller meant to handle scroll events
// from mice and touchpads. It is capable of handling both discrete and
// continuous scroll events, abstracting them both on the
// EventControllerScroll::scroll signal (deltas in the discrete case are
// multiples of 1).
//
// In the case of continuous scroll events, EventControllerScroll encloses all
// EventControllerScroll::scroll events between two
// EventControllerScroll::scroll-begin and EventControllerScroll::scroll-end
// signals.
//
// The behavior of the event controller can be modified by the flags given at
// creation time, or modified at a later point through
// gtk_event_controller_scroll_set_flags() (e.g. because the scrolling
// conditions of the widget changed).
//
// The controller can be set up to emit motion for either/both vertical and
// horizontal scroll events through K_EVENT_CONTROLLER_SCROLL_VERTICAL,
// K_EVENT_CONTROLLER_SCROLL_HORIZONTAL and K_EVENT_CONTROLLER_SCROLL_BOTH. If
// any axis is disabled, the respective EventControllerScroll::scroll delta will
// be 0. Vertical scroll events will be translated to horizontal motion for the
// devices incapable of horizontal scrolling.
//
// The event controller can also be forced to emit discrete events on all
// devices through K_EVENT_CONTROLLER_SCROLL_DISCRETE. This can be used to
// implement discrete actions triggered through scroll events (e.g. switching
// across combobox options).
//
// The K_EVENT_CONTROLLER_SCROLL_KINETIC flag toggles the emission of the
// EventControllerScroll::decelerate signal, emitted at the end of scrolling
// with two X/Y velocity arguments that are consistent with the motion that was
// received.
//
// This object was added in 3.24.
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
	return wrapEventControllerScroll(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewEventControllerScroll creates a new event controller that will handle
// scroll events for the given widget.
//
// The function takes the following parameters:
//
//    - widget: Widget.
//    - flags: behavior flags.
//
func NewEventControllerScroll(widget Widgetter, flags EventControllerScrollFlags) *EventControllerScroll {
	var _arg1 *C.GtkWidget                    // out
	var _arg2 C.GtkEventControllerScrollFlags // out
	var _cret *C.GtkEventController           // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.GtkEventControllerScrollFlags(flags)

	_cret = C.gtk_event_controller_scroll_new(_arg1, _arg2)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(flags)

	var _eventControllerScroll *EventControllerScroll // out

	_eventControllerScroll = wrapEventControllerScroll(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerScroll
}

// Flags gets the flags conditioning the scroll controller behavior.
func (controller *EventControllerScroll) Flags() EventControllerScrollFlags {
	var _arg0 *C.GtkEventControllerScroll     // out
	var _cret C.GtkEventControllerScrollFlags // in

	_arg0 = (*C.GtkEventControllerScroll)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_scroll_get_flags(_arg0)
	runtime.KeepAlive(controller)

	var _eventControllerScrollFlags EventControllerScrollFlags // out

	_eventControllerScrollFlags = EventControllerScrollFlags(_cret)

	return _eventControllerScrollFlags
}

// SetFlags sets the flags conditioning scroll controller behavior.
//
// The function takes the following parameters:
//
//    - flags: behavior flags.
//
func (controller *EventControllerScroll) SetFlags(flags EventControllerScrollFlags) {
	var _arg0 *C.GtkEventControllerScroll     // out
	var _arg1 C.GtkEventControllerScrollFlags // out

	_arg0 = (*C.GtkEventControllerScroll)(unsafe.Pointer(controller.Native()))
	_arg1 = C.GtkEventControllerScrollFlags(flags)

	C.gtk_event_controller_scroll_set_flags(_arg0, _arg1)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(flags)
}

// ConnectDecelerate: emitted after scroll is finished if the
// K_EVENT_CONTROLLER_SCROLL_KINETIC flag is set. vel_x and vel_y express the
// initial velocity that was imprinted by the scroll events. vel_x and vel_y are
// expressed in pixels/ms.
func (controller *EventControllerScroll) ConnectDecelerate(f func(velX, velY float64)) externglib.SignalHandle {
	return controller.Connect("decelerate", f)
}

// ConnectScroll signals that the widget should scroll by the amount specified
// by dx and dy.
func (controller *EventControllerScroll) ConnectScroll(f func(dx, dy float64)) externglib.SignalHandle {
	return controller.Connect("scroll", f)
}

// ConnectScrollBegin signals that a new scrolling operation has begun. It will
// only be emitted on devices capable of it.
func (controller *EventControllerScroll) ConnectScrollBegin(f func()) externglib.SignalHandle {
	return controller.Connect("scroll-begin", f)
}

// ConnectScrollEnd signals that a new scrolling operation has finished. It will
// only be emitted on devices capable of it.
func (controller *EventControllerScroll) ConnectScrollEnd(f func()) externglib.SignalHandle {
	return controller.Connect("scroll-end", f)
}
