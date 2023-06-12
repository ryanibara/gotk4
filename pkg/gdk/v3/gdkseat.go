// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
// extern void _gotk4_gdk3_Seat_ConnectToolRemoved(gpointer, GdkDeviceTool*, guintptr);
// extern void _gotk4_gdk3_Seat_ConnectToolAdded(gpointer, GdkDeviceTool*, guintptr);
// extern void _gotk4_gdk3_Seat_ConnectDeviceRemoved(gpointer, GdkDevice*, guintptr);
// extern void _gotk4_gdk3_Seat_ConnectDeviceAdded(gpointer, GdkDevice*, guintptr);
// extern void _gotk4_gdk3_SeatGrabPrepareFunc(GdkSeat*, GdkWindow*, gpointer);
import "C"

// GType values.
var (
	GTypeSeat = coreglib.Type(C.gdk_seat_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSeat, F: marshalSeat},
	})
}

// Seat object represents a collection of input devices that belong to a user.
type Seat struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Seat)(nil)
)

// Seater describes types inherited from class Seat.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Seater interface {
	coreglib.Objector
	baseSeat() *Seat
}

var _ Seater = (*Seat)(nil)

func wrapSeat(obj *coreglib.Object) *Seat {
	return &Seat{
		Object: obj,
	}
}

func marshalSeat(p uintptr) (interface{}, error) {
	return wrapSeat(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (seat *Seat) baseSeat() *Seat {
	return seat
}

// BaseSeat returns the underlying base object.
func BaseSeat(obj Seater) *Seat {
	return obj.baseSeat()
}

// ConnectDeviceAdded signal is emitted when a new input device is related to
// this seat.
func (seat *Seat) ConnectDeviceAdded(f func(device Devicer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(seat, "device-added", false, unsafe.Pointer(C._gotk4_gdk3_Seat_ConnectDeviceAdded), f)
}

// ConnectDeviceRemoved signal is emitted when an input device is removed (e.g.
// unplugged).
func (seat *Seat) ConnectDeviceRemoved(f func(device Devicer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(seat, "device-removed", false, unsafe.Pointer(C._gotk4_gdk3_Seat_ConnectDeviceRemoved), f)
}

// ConnectToolAdded signal is emitted whenever a new tool is made known to the
// seat. The tool may later be assigned to a device (i.e. on proximity with a
// tablet). The device will emit the Device::tool-changed signal accordingly.
//
// A same tool may be used by several devices.
func (seat *Seat) ConnectToolAdded(f func(tool *DeviceTool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(seat, "tool-added", false, unsafe.Pointer(C._gotk4_gdk3_Seat_ConnectToolAdded), f)
}

// ConnectToolRemoved: this signal is emitted whenever a tool is no longer known
// to this seat.
func (seat *Seat) ConnectToolRemoved(f func(tool *DeviceTool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(seat, "tool-removed", false, unsafe.Pointer(C._gotk4_gdk3_Seat_ConnectToolRemoved), f)
}

// Capabilities returns the capabilities this Seat currently has.
//
// The function returns the following values:
//
//   - seatCapabilities: seat capabilities.
//
func (seat *Seat) Capabilities() SeatCapabilities {
	var _arg0 *C.GdkSeat            // out
	var _cret C.GdkSeatCapabilities // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(coreglib.InternObject(seat).Native()))

	_cret = C.gdk_seat_get_capabilities(_arg0)
	runtime.KeepAlive(seat)

	var _seatCapabilities SeatCapabilities // out

	_seatCapabilities = SeatCapabilities(_cret)

	return _seatCapabilities
}

// Display returns the Display this seat belongs to.
//
// The function returns the following values:
//
//   - display This object is owned by GTK+ and must not be freed.
//
func (seat *Seat) Display() *Display {
	var _arg0 *C.GdkSeat    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(coreglib.InternObject(seat).Native()))

	_cret = C.gdk_seat_get_display(_arg0)
	runtime.KeepAlive(seat)

	var _display *Display // out

	_display = wrapDisplay(coreglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// Keyboard returns the master device that routes keyboard events.
//
// The function returns the following values:
//
//   - device (optional): master Device with keyboard capabilities. This object
//     is owned by GTK+ and must not be freed.
//
func (seat *Seat) Keyboard() Devicer {
	var _arg0 *C.GdkSeat   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(coreglib.InternObject(seat).Native()))

	_cret = C.gdk_seat_get_keyboard(_arg0)
	runtime.KeepAlive(seat)

	var _device Devicer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Devicer)
				return ok
			})
			rv, ok := casted.(Devicer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
			}
			_device = rv
		}
	}

	return _device
}

// Pointer returns the master device that routes pointer events.
//
// The function returns the following values:
//
//   - device (optional): master Device with pointer capabilities. This object
//     is owned by GTK+ and must not be freed.
//
func (seat *Seat) Pointer() Devicer {
	var _arg0 *C.GdkSeat   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(coreglib.InternObject(seat).Native()))

	_cret = C.gdk_seat_get_pointer(_arg0)
	runtime.KeepAlive(seat)

	var _device Devicer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Devicer)
				return ok
			})
			rv, ok := casted.(Devicer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
			}
			_device = rv
		}
	}

	return _device
}

// Slaves returns the slave devices that match the given capabilities.
//
// The function takes the following parameters:
//
//   - capabilities to get devices for.
//
// The function returns the following values:
//
//   - list of Devices. The list must be freed with g_list_free(), the elements
//     are owned by GDK and must not be freed.
//
func (seat *Seat) Slaves(capabilities SeatCapabilities) []Devicer {
	var _arg0 *C.GdkSeat            // out
	var _arg1 C.GdkSeatCapabilities // out
	var _cret *C.GList              // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(coreglib.InternObject(seat).Native()))
	_arg1 = C.GdkSeatCapabilities(capabilities)

	_cret = C.gdk_seat_get_slaves(_arg0, _arg1)
	runtime.KeepAlive(seat)
	runtime.KeepAlive(capabilities)

	var _list []Devicer // out

	_list = make([]Devicer, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GdkDevice)(v)
		var dst Devicer // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gdk.Devicer is nil")
			}

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Devicer)
				return ok
			})
			rv, ok := casted.(Devicer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}

// Grab grabs the seat so that all events corresponding to the given
// capabilities are passed to this application until the seat is ungrabbed with
// gdk_seat_ungrab(), or the window becomes hidden. This overrides any previous
// grab on the seat by this client.
//
// As a rule of thumb, if a grab is desired over GDK_SEAT_CAPABILITY_POINTER,
// all other "pointing" capabilities (eg. GDK_SEAT_CAPABILITY_TOUCH) should be
// grabbed too, so the user is able to interact with all of those while the grab
// holds, you should thus use GDK_SEAT_CAPABILITY_ALL_POINTING most commonly.
//
// Grabs are used for operations which need complete control over the events
// corresponding to the given capabilities. For example in GTK+ this is used for
// Drag and Drop operations, popup menus and such.
//
// Note that if the event mask of a Window has selected both button press and
// button release events, or touch begin and touch end, then a press event
// will cause an automatic grab until the button is released, equivalent to a
// grab on the window with owner_events set to TRUE. This is done because most
// applications expect to receive paired press and release events.
//
// If you set up anything at the time you take the grab that needs to be cleaned
// up when the grab ends, you should handle the EventGrabBroken events that are
// emitted when the grab ends unvoluntarily.
//
// The function takes the following parameters:
//
//   - window which will own the grab.
//   - capabilities that will be grabbed.
//   - ownerEvents: if FALSE then all device events are reported with respect
//     to window and are only reported if selected by event_mask. If TRUE then
//     pointer events for this application are reported as normal, but pointer
//     events outside this application are reported with respect to window and
//     only if selected by event_mask. In either mode, unreported events are
//     discarded.
//   - cursor (optional) to display while the grab is active. If this is NULL
//     then the normal cursors are used for window and its descendants, and the
//     cursor for window is used elsewhere.
//   - event (optional) that is triggering the grab, or NULL if none is
//     available.
//   - prepareFunc (optional): function to prepare the window to be grabbed,
//     it can be NULL if window is visible before this call.
//
// The function returns the following values:
//
//   - grabStatus: GDK_GRAB_SUCCESS if the grab was successful.
//
func (seat *Seat) Grab(window Windower, capabilities SeatCapabilities, ownerEvents bool, cursor Cursorrer, event *Event, prepareFunc SeatGrabPrepareFunc) GrabStatus {
	var _arg0 *C.GdkSeat               // out
	var _arg1 *C.GdkWindow             // out
	var _arg2 C.GdkSeatCapabilities    // out
	var _arg3 C.gboolean               // out
	var _arg4 *C.GdkCursor             // out
	var _arg5 *C.GdkEvent              // out
	var _arg6 C.GdkSeatGrabPrepareFunc // out
	var _arg7 C.gpointer
	var _cret C.GdkGrabStatus // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(coreglib.InternObject(seat).Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	_arg2 = C.GdkSeatCapabilities(capabilities)
	if ownerEvents {
		_arg3 = C.TRUE
	}
	if cursor != nil {
		_arg4 = (*C.GdkCursor)(unsafe.Pointer(coreglib.InternObject(cursor).Native()))
	}
	if event != nil {
		_arg5 = (*C.GdkEvent)(gextras.StructNative(unsafe.Pointer(event)))
	}
	if prepareFunc != nil {
		_arg6 = (*[0]byte)(C._gotk4_gdk3_SeatGrabPrepareFunc)
		_arg7 = C.gpointer(gbox.Assign(prepareFunc))
		defer gbox.Delete(uintptr(_arg7))
	}

	_cret = C.gdk_seat_grab(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(seat)
	runtime.KeepAlive(window)
	runtime.KeepAlive(capabilities)
	runtime.KeepAlive(ownerEvents)
	runtime.KeepAlive(cursor)
	runtime.KeepAlive(event)
	runtime.KeepAlive(prepareFunc)

	var _grabStatus GrabStatus // out

	_grabStatus = GrabStatus(_cret)

	return _grabStatus
}

// Ungrab releases a grab added through gdk_seat_grab().
func (seat *Seat) Ungrab() {
	var _arg0 *C.GdkSeat // out

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(coreglib.InternObject(seat).Native()))

	C.gdk_seat_ungrab(_arg0)
	runtime.KeepAlive(seat)
}
