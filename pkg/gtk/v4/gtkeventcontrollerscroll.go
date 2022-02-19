// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern gboolean _gotk4_gtk4_EventControllerScroll_ConnectScroll(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk4_EventControllerScroll_ConnectDecelerate(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk4_EventControllerScroll_ConnectScrollBegin(gpointer, guintptr);
// extern void _gotk4_gtk4_EventControllerScroll_ConnectScrollEnd(gpointer, guintptr);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_scroll_flags_get_type()), F: marshalEventControllerScrollFlags},
		{T: externglib.Type(C.gtk_event_controller_scroll_get_type()), F: marshalEventControllerScroller},
	})
}

// EventControllerScrollFlags describes the behavior of a
// GtkEventControllerScroll.
type EventControllerScrollFlags C.guint

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

// EventControllerScrollOverrider contains methods that are overridable.
type EventControllerScrollOverrider interface {
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
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*EventControllerScroll)(nil)
)

func classInitEventControllerScroller(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

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

//export _gotk4_gtk4_EventControllerScroll_ConnectDecelerate
func _gotk4_gtk4_EventControllerScroll_ConnectDecelerate(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(velX, velY float64)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(velX, velY float64))
	}

	var _velX float64 // out
	var _velY float64 // out

	_velX = float64(arg1)
	_velY = float64(arg2)

	f(_velX, _velY)
}

// ConnectDecelerate: emitted after scroll is finished if the
// GTK_EVENT_CONTROLLER_SCROLL_KINETIC flag is set.
//
// vel_x and vel_y express the initial velocity that was imprinted by the scroll
// events. vel_x and vel_y are expressed in pixels/ms.
func (scroll *EventControllerScroll) ConnectDecelerate(f func(velX, velY float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(scroll, "decelerate", false, unsafe.Pointer(C._gotk4_gtk4_EventControllerScroll_ConnectDecelerate), f)
}

//export _gotk4_gtk4_EventControllerScroll_ConnectScroll
func _gotk4_gtk4_EventControllerScroll_ConnectScroll(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) (cret C.gboolean) {
	var f func(dx, dy float64) (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(dx, dy float64) (ok bool))
	}

	var _dx float64 // out
	var _dy float64 // out

	_dx = float64(arg1)
	_dy = float64(arg2)

	ok := f(_dx, _dy)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectScroll signals that the widget should scroll by the amount specified
// by dx and dy.
func (scroll *EventControllerScroll) ConnectScroll(f func(dx, dy float64) (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(scroll, "scroll", false, unsafe.Pointer(C._gotk4_gtk4_EventControllerScroll_ConnectScroll), f)
}

//export _gotk4_gtk4_EventControllerScroll_ConnectScrollBegin
func _gotk4_gtk4_EventControllerScroll_ConnectScrollBegin(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectScrollBegin signals that a new scrolling operation has begun.
//
// It will only be emitted on devices capable of it.
func (scroll *EventControllerScroll) ConnectScrollBegin(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(scroll, "scroll-begin", false, unsafe.Pointer(C._gotk4_gtk4_EventControllerScroll_ConnectScrollBegin), f)
}

//export _gotk4_gtk4_EventControllerScroll_ConnectScrollEnd
func _gotk4_gtk4_EventControllerScroll_ConnectScrollEnd(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectScrollEnd signals that a scrolling operation has finished.
//
// It will only be emitted on devices capable of it.
func (scroll *EventControllerScroll) ConnectScrollEnd(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(scroll, "scroll-end", false, unsafe.Pointer(C._gotk4_gtk4_EventControllerScroll_ConnectScrollEnd), f)
}

// NewEventControllerScroll creates a new event controller that will handle
// scroll events.
//
// The function takes the following parameters:
//
//    - flags affecting the controller behavior.
//
// The function returns the following values:
//
//    - eventControllerScroll: new GtkEventControllerScroll.
//
func NewEventControllerScroll(flags EventControllerScrollFlags) *EventControllerScroll {
	var _arg1 C.GtkEventControllerScrollFlags // out
	var _cret *C.GtkEventController           // in

	_arg1 = C.GtkEventControllerScrollFlags(flags)

	_cret = C.gtk_event_controller_scroll_new(_arg1)
	runtime.KeepAlive(flags)

	var _eventControllerScroll *EventControllerScroll // out

	_eventControllerScroll = wrapEventControllerScroll(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerScroll
}

// Flags gets the flags conditioning the scroll controller behavior.
//
// The function returns the following values:
//
//    - eventControllerScrollFlags: controller flags.
//
func (scroll *EventControllerScroll) Flags() EventControllerScrollFlags {
	var _arg0 *C.GtkEventControllerScroll     // out
	var _cret C.GtkEventControllerScrollFlags // in

	_arg0 = (*C.GtkEventControllerScroll)(unsafe.Pointer(scroll.Native()))

	_cret = C.gtk_event_controller_scroll_get_flags(_arg0)
	runtime.KeepAlive(scroll)

	var _eventControllerScrollFlags EventControllerScrollFlags // out

	_eventControllerScrollFlags = EventControllerScrollFlags(_cret)

	return _eventControllerScrollFlags
}

// SetFlags sets the flags conditioning scroll controller behavior.
//
// The function takes the following parameters:
//
//    - flags affecting the controller behavior.
//
func (scroll *EventControllerScroll) SetFlags(flags EventControllerScrollFlags) {
	var _arg0 *C.GtkEventControllerScroll     // out
	var _arg1 C.GtkEventControllerScrollFlags // out

	_arg0 = (*C.GtkEventControllerScroll)(unsafe.Pointer(scroll.Native()))
	_arg1 = C.GtkEventControllerScrollFlags(flags)

	C.gtk_event_controller_scroll_set_flags(_arg0, _arg1)
	runtime.KeepAlive(scroll)
	runtime.KeepAlive(flags)
}
