// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_seat_capabilities_get_type()), F: marshalSeatCapabilities},
		{T: externglib.Type(C.gdk_seat_get_type()), F: marshalSeater},
	})
}

// SeatCapabilities flags describing the seat capabilities.
type SeatCapabilities int

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
	// SeatCapabilityAllPointing: union of all pointing capabilities.
	SeatCapabilityAllPointing SeatCapabilities = 0b111
	// SeatCapabilityAll: union of all capabilities.
	SeatCapabilityAll SeatCapabilities = 0b1111
)

func marshalSeatCapabilities(p uintptr) (interface{}, error) {
	return SeatCapabilities(C.g_value_get_flags((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the names in string for SeatCapabilities.
func (s SeatCapabilities) String() string {
	if s == 0 {
		return "SeatCapabilities(0)"
	}

	var builder strings.Builder
	builder.Grow(154)

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

// SeatGrabPrepareFunc: type of the callback used to set up window so it can be
// grabbed. A typical action would be ensuring the window is visible, although
// there's room for other initialization actions.
type SeatGrabPrepareFunc func(seat Seater, window Windower)

//export _gotk4_gdk3_SeatGrabPrepareFunc
func _gotk4_gdk3_SeatGrabPrepareFunc(arg0 *C.GdkSeat, arg1 *C.GdkWindow, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var seat Seater     // out
	var window Windower // out

	{
		objptr := unsafe.Pointer(arg0)
		if objptr == nil {
			panic("object of type gdk.Seater is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Seater)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Seater")
		}
		seat = rv
	}
	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Windower)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Windower")
		}
		window = rv
	}

	fn := v.(SeatGrabPrepareFunc)
	fn(seat, window)
}

// Seat object represents a collection of input devices that belong to a user.
type Seat struct {
	*externglib.Object
}

// Seater describes Seat's abstract methods.
type Seater interface {
	externglib.Objector

	// Capabilities returns the capabilities this Seat currently has.
	Capabilities() SeatCapabilities
	// Display returns the Display this seat belongs to.
	Display() *Display
	// Keyboard returns the master device that routes keyboard events.
	Keyboard() Devicer
	// Pointer returns the master device that routes pointer events.
	Pointer() Devicer
	// Slaves returns the slave devices that match the given capabilities.
	Slaves(capabilities SeatCapabilities) []Devicer
	// Ungrab releases a grab added through gdk_seat_grab().
	Ungrab()
}

var _ Seater = (*Seat)(nil)

func wrapSeat(obj *externglib.Object) *Seat {
	return &Seat{
		Object: obj,
	}
}

func marshalSeater(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSeat(obj), nil
}

// Capabilities returns the capabilities this Seat currently has.
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

// Display returns the Display this seat belongs to.
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

// Keyboard returns the master device that routes keyboard events.
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
			rv, ok := (externglib.CastObject(object)).(Devicer)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Devicer")
			}
			_device = rv
		}
	}

	return _device
}

// Pointer returns the master device that routes pointer events.
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
			rv, ok := (externglib.CastObject(object)).(Devicer)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Devicer")
			}
			_device = rv
		}
	}

	return _device
}

// Slaves returns the slave devices that match the given capabilities.
func (seat *Seat) Slaves(capabilities SeatCapabilities) []Devicer {
	var _arg0 *C.GdkSeat            // out
	var _arg1 C.GdkSeatCapabilities // out
	var _cret *C.GList              // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))
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

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Devicer)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Devicer")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}

// Ungrab releases a grab added through gdk_seat_grab().
func (seat *Seat) Ungrab() {
	var _arg0 *C.GdkSeat // out

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	C.gdk_seat_ungrab(_arg0)
	runtime.KeepAlive(seat)
}
