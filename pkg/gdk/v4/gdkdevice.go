// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_input_source_get_type()), F: marshalInputSource},
		{T: externglib.Type(C.gdk_device_get_type()), F: marshalDevice},
	})
}

// InputSource: an enumeration describing the type of an input device in general
// terms.
type InputSource int

const (
	// mouse: the device is a mouse. (This will be reported for the core
	// pointer, even if it is something else, such as a trackball.)
	InputSourceMouse InputSource = 0
	// pen: the device is a stylus of a graphics tablet or similar device.
	InputSourcePen InputSource = 1
	// keyboard: the device is a keyboard.
	InputSourceKeyboard InputSource = 2
	// touchscreen: the device is a direct-input touch device, such as a
	// touchscreen or tablet
	InputSourceTouchscreen InputSource = 3
	// touchpad: the device is an indirect touch device, such as a touchpad
	InputSourceTouchpad InputSource = 4
	// trackpoint: the device is a trackpoint
	InputSourceTrackpoint InputSource = 5
	// TabletPad: the device is a "pad", a collection of buttons, rings and
	// strips found in drawing tablets
	InputSourceTabletPad InputSource = 6
)

func marshalInputSource(p uintptr) (interface{}, error) {
	return InputSource(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Device: the `GdkDevice` object represents an input device, such as a
// keyboard, a mouse, or a touchpad.
//
// See the [class@Gdk.Seat] documentation for more information about the various
// kinds of devices, and their relationships.
type Device interface {
	gextras.Objector

	CapsLockState() bool

	DeviceTool() DeviceTool

	Direction() pango.Direction

	Display() Display

	HasCursor() bool

	ModifierState() ModifierType

	Name() string

	NumLockState() bool

	NumTouches() uint

	ProductID() string

	ScrollLockState() bool

	Seat() Seat

	Source() InputSource

	SurfaceAtPosition() (winX float64, winY float64, surface Surface)

	Timestamp() uint32

	VendorID() string

	HasBidiLayoutsDevice() bool
}

// device implements the Device class.
type device struct {
	gextras.Objector
}

// WrapDevice wraps a GObject to the right type. It is
// primarily used internally.
func WrapDevice(obj *externglib.Object) Device {
	return device{
		Objector: obj,
	}
}

func marshalDevice(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDevice(obj), nil
}

func (d device) CapsLockState() bool {
	var _arg0 *C.GdkDevice // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_caps_lock_state(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d device) DeviceTool() DeviceTool {
	var _arg0 *C.GdkDevice     // out
	var _cret *C.GdkDeviceTool // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_device_tool(_arg0)

	var _deviceTool DeviceTool // out

	_deviceTool = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(DeviceTool)

	return _deviceTool
}

func (d device) Direction() pango.Direction {
	var _arg0 *C.GdkDevice     // out
	var _cret C.PangoDirection // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_direction(_arg0)

	var _direction pango.Direction // out

	_direction = pango.Direction(_cret)

	return _direction
}

func (d device) Display() Display {
	var _arg0 *C.GdkDevice  // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Display)

	return _display
}

func (d device) HasCursor() bool {
	var _arg0 *C.GdkDevice // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_has_cursor(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d device) ModifierState() ModifierType {
	var _arg0 *C.GdkDevice      // out
	var _cret C.GdkModifierType // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_modifier_state(_arg0)

	var _modifierType ModifierType // out

	_modifierType = ModifierType(_cret)

	return _modifierType
}

func (d device) Name() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.char      // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d device) NumLockState() bool {
	var _arg0 *C.GdkDevice // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_num_lock_state(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d device) NumTouches() uint {
	var _arg0 *C.GdkDevice // out
	var _cret C.guint      // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_num_touches(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (d device) ProductID() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.char      // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_product_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d device) ScrollLockState() bool {
	var _arg0 *C.GdkDevice // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_scroll_lock_state(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d device) Seat() Seat {
	var _arg0 *C.GdkDevice // out
	var _cret *C.GdkSeat   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_seat(_arg0)

	var _seat Seat // out

	_seat = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Seat)

	return _seat
}

func (d device) Source() InputSource {
	var _arg0 *C.GdkDevice     // out
	var _cret C.GdkInputSource // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_source(_arg0)

	var _inputSource InputSource // out

	_inputSource = InputSource(_cret)

	return _inputSource
}

func (d device) SurfaceAtPosition() (winX float64, winY float64, surface Surface) {
	var _arg0 *C.GdkDevice  // out
	var _arg1 C.double      // in
	var _arg2 C.double      // in
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_surface_at_position(_arg0, &_arg1, &_arg2)

	var _winX float64    // out
	var _winY float64    // out
	var _surface Surface // out

	_winX = float64(_arg1)
	_winY = float64(_arg2)
	_surface = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Surface)

	return _winX, _winY, _surface
}

func (d device) Timestamp() uint32 {
	var _arg0 *C.GdkDevice // out
	var _cret C.guint32    // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_timestamp(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

func (d device) VendorID() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.char      // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_vendor_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d device) HasBidiLayoutsDevice() bool {
	var _arg0 *C.GdkDevice // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_has_bidi_layouts(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TimeCoord: a TimeCoord stores a single event in a motion history.
type TimeCoord C.GdkTimeCoord

// WrapTimeCoord wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTimeCoord(ptr unsafe.Pointer) *TimeCoord {
	return (*TimeCoord)(ptr)
}

// Native returns the underlying C source pointer.
func (t *TimeCoord) Native() unsafe.Pointer {
	return unsafe.Pointer(t)
}
