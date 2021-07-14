// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
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
	// SeatCapabilitiesNone: no input capabilities
	SeatCapabilitiesNone SeatCapabilities = 0b0
	// SeatCapabilitiesPointer: seat has a pointer (e.g. mouse)
	SeatCapabilitiesPointer SeatCapabilities = 0b1
	// SeatCapabilitiesTouch: seat has touchscreen(s) attached
	SeatCapabilitiesTouch SeatCapabilities = 0b10
	// SeatCapabilitiesTabletStylus: seat has drawing tablet(s) attached
	SeatCapabilitiesTabletStylus SeatCapabilities = 0b100
	// SeatCapabilitiesKeyboard: seat has keyboard(s) attached
	SeatCapabilitiesKeyboard SeatCapabilities = 0b1000
	// SeatCapabilitiesAllPointing: union of all pointing capabilities
	SeatCapabilitiesAllPointing SeatCapabilities = 0b111
	// SeatCapabilitiesAll: union of all capabilities
	SeatCapabilitiesAll SeatCapabilities = 0b1111
)

func marshalSeatCapabilities(p uintptr) (interface{}, error) {
	return SeatCapabilities(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// SeatGrabPrepareFunc: type of the callback used to set up window so it can be
// grabbed. A typical action would be ensuring the window is visible, although
// there's room for other initialization actions.
type SeatGrabPrepareFunc func(seat *Seat, window *Window, userData cgo.Handle)

//export _gotk4_gdk3_SeatGrabPrepareFunc
func _gotk4_gdk3_SeatGrabPrepareFunc(arg0 *C.GdkSeat, arg1 *C.GdkWindow, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var seat *Seat          // out
	var window *Window      // out
	var userData cgo.Handle // out

	seat = wrapSeat(externglib.Take(unsafe.Pointer(arg0)))
	window = wrapWindow(externglib.Take(unsafe.Pointer(arg1)))
	userData = (cgo.Handle)(unsafe.Pointer(arg2))

	fn := v.(SeatGrabPrepareFunc)
	fn(seat, window, userData)
}

// Seater describes Seat's methods.
type Seater interface {
	// Capabilities returns the capabilities this Seat currently has.
	Capabilities() SeatCapabilities
	// Display returns the Display this seat belongs to.
	Display() *Display
	// Keyboard returns the master device that routes keyboard events.
	Keyboard() *Device
	// Pointer returns the master device that routes pointer events.
	Pointer() *Device
	// Ungrab releases a grab added through gdk_seat_grab().
	Ungrab()
}

// Seat object represents a collection of input devices that belong to a user.
type Seat struct {
	*externglib.Object
}

var (
	_ Seater          = (*Seat)(nil)
	_ gextras.Nativer = (*Seat)(nil)
)

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

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// Keyboard returns the master device that routes keyboard events.
func (seat *Seat) Keyboard() *Device {
	var _arg0 *C.GdkSeat   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	_cret = C.gdk_seat_get_keyboard(_arg0)

	var _device *Device // out

	_device = wrapDevice(externglib.Take(unsafe.Pointer(_cret)))

	return _device
}

// Pointer returns the master device that routes pointer events.
func (seat *Seat) Pointer() *Device {
	var _arg0 *C.GdkSeat   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	_cret = C.gdk_seat_get_pointer(_arg0)

	var _device *Device // out

	_device = wrapDevice(externglib.Take(unsafe.Pointer(_cret)))

	return _device
}

// Ungrab releases a grab added through gdk_seat_grab().
func (seat *Seat) Ungrab() {
	var _arg0 *C.GdkSeat // out

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	C.gdk_seat_ungrab(_arg0)
}
