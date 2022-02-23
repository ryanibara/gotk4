// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
// extern void _gotk4_gdk4_Seat_ConnectDeviceAdded(gpointer, GdkDevice*, guintptr);
// extern void _gotk4_gdk4_Seat_ConnectDeviceRemoved(gpointer, GdkDevice*, guintptr);
// extern void _gotk4_gdk4_Seat_ConnectToolAdded(gpointer, GdkDeviceTool*, guintptr);
// extern void _gotk4_gdk4_Seat_ConnectToolRemoved(gpointer, GdkDeviceTool*, guintptr);
import "C"

// glib.Type values for gdkseat.go.
var (
	GTypeSeatCapabilities = externglib.Type(C.gdk_seat_capabilities_get_type())
	GTypeSeat             = externglib.Type(C.gdk_seat_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeSeatCapabilities, F: marshalSeatCapabilities},
		{T: GTypeSeat, F: marshalSeat},
	})
}

// SeatCapabilities flags describing the seat capabilities.
type SeatCapabilities C.guint

const (
	// SeatCapabilityNone: no input capabilities.
	SeatCapabilityNone SeatCapabilities = 0b0
	// SeatCapabilityPointer: seat has a pointer (e.g. mouse).
	SeatCapabilityPointer SeatCapabilities = 0b1
	// SeatCapabilityTouch: seat has touchscreen(s) attached.
	SeatCapabilityTouch SeatCapabilities = 0b10
	// SeatCapabilityTabletStylus: seat has drawing tablet(s) attached.
	SeatCapabilityTabletStylus SeatCapabilities = 0b100
	// SeatCapabilityKeyboard: seat has keyboard(s) attached.
	SeatCapabilityKeyboard SeatCapabilities = 0b1000
	// SeatCapabilityTabletPad: seat has drawing tablet pad(s) attached.
	SeatCapabilityTabletPad SeatCapabilities = 0b10000
	// SeatCapabilityAllPointing: union of all pointing capabilities.
	SeatCapabilityAllPointing SeatCapabilities = 0b111
	// SeatCapabilityAll: union of all capabilities.
	SeatCapabilityAll SeatCapabilities = 0b1111
)

func marshalSeatCapabilities(p uintptr) (interface{}, error) {
	return SeatCapabilities(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for SeatCapabilities.
func (s SeatCapabilities) String() string {
	if s == 0 {
		return "SeatCapabilities(0)"
	}

	var builder strings.Builder
	builder.Grow(178)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case SeatCapabilityNone:
			builder.WriteString("None|")
		case SeatCapabilityPointer:
			builder.WriteString("Pointer|")
		case SeatCapabilityTouch:
			builder.WriteString("Touch|")
		case SeatCapabilityTabletStylus:
			builder.WriteString("TabletStylus|")
		case SeatCapabilityKeyboard:
			builder.WriteString("Keyboard|")
		case SeatCapabilityTabletPad:
			builder.WriteString("TabletPad|")
		case SeatCapabilityAllPointing:
			builder.WriteString("AllPointing|")
		case SeatCapabilityAll:
			builder.WriteString("All|")
		default:
			builder.WriteString(fmt.Sprintf("SeatCapabilities(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s SeatCapabilities) Has(other SeatCapabilities) bool {
	return (s & other) == other
}

// Seat: GdkSeat object represents a collection of input devices that belong to
// a user.
type Seat struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Seat)(nil)
)

// Seater describes types inherited from class Seat.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Seater interface {
	externglib.Objector
	baseSeat() *Seat
}

var _ Seater = (*Seat)(nil)

func wrapSeat(obj *externglib.Object) *Seat {
	return &Seat{
		Object: obj,
	}
}

func marshalSeat(p uintptr) (interface{}, error) {
	return wrapSeat(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (seat *Seat) baseSeat() *Seat {
	return seat
}

// BaseSeat returns the underlying base object.
func BaseSeat(obj Seater) *Seat {
	return obj.baseSeat()
}

//export _gotk4_gdk4_Seat_ConnectDeviceAdded
func _gotk4_gdk4_Seat_ConnectDeviceAdded(arg0 C.gpointer, arg1 *C.GdkDevice, arg2 C.guintptr) {
	var f func(device Devicer)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(device Devicer))
	}

	var _device Devicer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Devicer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Devicer)
			return ok
		})
		rv, ok := casted.(Devicer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
		}
		_device = rv
	}

	f(_device)
}

// ConnectDeviceAdded: emitted when a new input device is related to this seat.
func (seat *Seat) ConnectDeviceAdded(f func(device Devicer)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(seat, "device-added", false, unsafe.Pointer(C._gotk4_gdk4_Seat_ConnectDeviceAdded), f)
}

//export _gotk4_gdk4_Seat_ConnectDeviceRemoved
func _gotk4_gdk4_Seat_ConnectDeviceRemoved(arg0 C.gpointer, arg1 *C.GdkDevice, arg2 C.guintptr) {
	var f func(device Devicer)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(device Devicer))
	}

	var _device Devicer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Devicer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Devicer)
			return ok
		})
		rv, ok := casted.(Devicer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
		}
		_device = rv
	}

	f(_device)
}

// ConnectDeviceRemoved: emitted when an input device is removed (e.g.
// unplugged).
func (seat *Seat) ConnectDeviceRemoved(f func(device Devicer)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(seat, "device-removed", false, unsafe.Pointer(C._gotk4_gdk4_Seat_ConnectDeviceRemoved), f)
}

//export _gotk4_gdk4_Seat_ConnectToolAdded
func _gotk4_gdk4_Seat_ConnectToolAdded(arg0 C.gpointer, arg1 *C.GdkDeviceTool, arg2 C.guintptr) {
	var f func(tool *DeviceTool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(tool *DeviceTool))
	}

	var _tool *DeviceTool // out

	_tool = wrapDeviceTool(externglib.Take(unsafe.Pointer(arg1)))

	f(_tool)
}

// ConnectToolAdded: emitted whenever a new tool is made known to the seat.
//
// The tool may later be assigned to a device (i.e. on proximity with a tablet).
// The device will emit the [signalGdkDevice::tool-changed] signal accordingly.
//
// A same tool may be used by several devices.
func (seat *Seat) ConnectToolAdded(f func(tool *DeviceTool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(seat, "tool-added", false, unsafe.Pointer(C._gotk4_gdk4_Seat_ConnectToolAdded), f)
}

//export _gotk4_gdk4_Seat_ConnectToolRemoved
func _gotk4_gdk4_Seat_ConnectToolRemoved(arg0 C.gpointer, arg1 *C.GdkDeviceTool, arg2 C.guintptr) {
	var f func(tool *DeviceTool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(tool *DeviceTool))
	}

	var _tool *DeviceTool // out

	_tool = wrapDeviceTool(externglib.Take(unsafe.Pointer(arg1)))

	f(_tool)
}

// ConnectToolRemoved: emitted whenever a tool is no longer known to this seat.
func (seat *Seat) ConnectToolRemoved(f func(tool *DeviceTool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(seat, "tool-removed", false, unsafe.Pointer(C._gotk4_gdk4_Seat_ConnectToolRemoved), f)
}

// Capabilities returns the capabilities this GdkSeat currently has.
//
// The function returns the following values:
//
//    - seatCapabilities: seat capabilities.
//
func (seat *Seat) Capabilities() SeatCapabilities {
	var _arg0 *C.GdkSeat            // out
	var _cret C.GdkSeatCapabilities // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	_cret = C.gdk_seat_get_capabilities(_arg0)
	runtime.KeepAlive(seat)

	var _seatCapabilities SeatCapabilities // out

	_seatCapabilities = SeatCapabilities(_cret)

	return _seatCapabilities
}

// Devices returns the devices that match the given capabilities.
//
// The function takes the following parameters:
//
//    - capabilities to get devices for.
//
// The function returns the following values:
//
//    - list: list of GdkDevices. The list must be freed with g_list_free(), the
//      elements are owned by GTK and must not be freed.
//
func (seat *Seat) Devices(capabilities SeatCapabilities) []Devicer {
	var _arg0 *C.GdkSeat            // out
	var _arg1 C.GdkSeatCapabilities // out
	var _cret *C.GList              // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))
	_arg1 = C.GdkSeatCapabilities(capabilities)

	_cret = C.gdk_seat_get_devices(_arg0, _arg1)
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

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
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

// Display returns the GdkDisplay this seat belongs to.
//
// The function returns the following values:
//
//    - display: GdkDisplay. This object is owned by GTK and must not be freed.
//
func (seat *Seat) Display() *Display {
	var _arg0 *C.GdkSeat    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	_cret = C.gdk_seat_get_display(_arg0)
	runtime.KeepAlive(seat)

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// Keyboard returns the device that routes keyboard events.
//
// The function returns the following values:
//
//    - device (optional): GdkDevice with keyboard capabilities. This object is
//      owned by GTK and must not be freed.
//
func (seat *Seat) Keyboard() Devicer {
	var _arg0 *C.GdkSeat   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	_cret = C.gdk_seat_get_keyboard(_arg0)
	runtime.KeepAlive(seat)

	var _device Devicer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
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

// Pointer returns the device that routes pointer events.
//
// The function returns the following values:
//
//    - device (optional): GdkDevice with pointer capabilities. This object is
//      owned by GTK and must not be freed.
//
func (seat *Seat) Pointer() Devicer {
	var _arg0 *C.GdkSeat   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	_cret = C.gdk_seat_get_pointer(_arg0)
	runtime.KeepAlive(seat)

	var _device Devicer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
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

// Tools returns all GdkDeviceTools that are known to the application.
//
// The function returns the following values:
//
//    - list: A list of tools. Free with g_list_free().
//
func (seat *Seat) Tools() []DeviceTool {
	var _arg0 *C.GdkSeat // out
	var _cret *C.GList   // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	_cret = C.gdk_seat_get_tools(_arg0)
	runtime.KeepAlive(seat)

	var _list []DeviceTool // out

	_list = make([]DeviceTool, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GdkDeviceTool)(v)
		var dst DeviceTool // out
		dst = *wrapDeviceTool(externglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}
