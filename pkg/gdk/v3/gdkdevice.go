// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_device_type_get_type()), F: marshalDeviceType},
		{T: externglib.Type(C.gdk_input_mode_get_type()), F: marshalInputMode},
		{T: externglib.Type(C.gdk_input_source_get_type()), F: marshalInputSource},
		{T: externglib.Type(C.gdk_device_get_type()), F: marshalDevicer},
	})
}

// DeviceType indicates the device type. See
// [above][GdkDeviceManager.description] for more information about the meaning
// of these device types.
type DeviceType int

const (
	// Master: device is a master (or virtual) device. There will be an
	// associated focus indicator on the screen.
	DeviceTypeMaster DeviceType = iota
	// Slave: device is a slave (or physical) device.
	DeviceTypeSlave
	// Floating: device is a physical device, currently not attached to any
	// virtual device.
	DeviceTypeFloating
)

func marshalDeviceType(p uintptr) (interface{}, error) {
	return DeviceType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// InputMode: enumeration that describes the mode of an input device.
type InputMode int

const (
	// Disabled: device is disabled and will not report any events.
	InputModeDisabled InputMode = iota
	// Screen: device is enabled. The device’s coordinate space maps to the
	// entire screen.
	InputModeScreen
	// Window: device is enabled. The device’s coordinate space is mapped to a
	// single window. The manner in which this window is chosen is undefined,
	// but it will typically be the same way in which the focus window for key
	// events is determined.
	InputModeWindow
)

func marshalInputMode(p uintptr) (interface{}, error) {
	return InputMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// InputSource: enumeration describing the type of an input device in general
// terms.
type InputSource int

const (
	// Mouse: device is a mouse. (This will be reported for the core pointer,
	// even if it is something else, such as a trackball.)
	InputSourceMouse InputSource = iota
	// Pen: device is a stylus of a graphics tablet or similar device.
	InputSourcePen
	// Eraser: device is an eraser. Typically, this would be the other end of a
	// stylus on a graphics tablet.
	InputSourceEraser
	// Cursor: device is a graphics tablet “puck” or similar device.
	InputSourceCursor
	// Keyboard: device is a keyboard.
	InputSourceKeyboard
	// Touchscreen: device is a direct-input touch device, such as a touchscreen
	// or tablet. This device type has been added in 3.4.
	InputSourceTouchscreen
	// Touchpad: device is an indirect touch device, such as a touchpad. This
	// device type has been added in 3.4.
	InputSourceTouchpad
	// Trackpoint: device is a trackpoint. This device type has been added in
	// 3.22
	InputSourceTrackpoint
	// TabletPad: device is a "pad", a collection of buttons, rings and strips
	// found in drawing tablets. This device type has been added in 3.22.
	InputSourceTabletPad
)

func marshalInputSource(p uintptr) (interface{}, error) {
	return InputSource(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Devicer describes Device's methods.
type Devicer interface {
	// AssociatedDevice returns the associated device to @device, if @device is
	// of type GDK_DEVICE_TYPE_MASTER, it will return the paired pointer or
	// keyboard.
	AssociatedDevice() *Device
	// Axes returns the axes currently available on the device.
	Axes() AxisFlags
	// AxisUse returns the axis use for @index_.
	AxisUse(index_ uint) AxisUse
	// DeviceType returns the device type for @device.
	DeviceType() DeviceType
	// Display returns the Display to which @device pertains.
	Display() *Display
	// HasCursor determines whether the pointer follows device motion.
	HasCursor() bool
	// Key: if @index_ has a valid keyval, this function will return true and
	// fill in @keyval and @modifiers with the keyval settings.
	Key(index_ uint) (uint, ModifierType, bool)
	// LastEventWindow gets information about which window the given pointer
	// device is in, based on events that have been received so far from the
	// display server.
	LastEventWindow() *Window
	// Mode determines the mode of the device.
	Mode() InputMode
	// NAxes returns the number of axes the device currently has.
	NAxes() int
	// NKeys returns the number of keys the device currently has.
	NKeys() int
	// Name determines the name of the device.
	Name() string
	// Position gets the current location of @device.
	Position() (screen *Screen, x int, y int)
	// PositionDouble gets the current location of @device in double precision.
	PositionDouble() (screen *Screen, x float64, y float64)
	// ProductID returns the product ID of this device, or nil if this
	// information couldn't be obtained.
	ProductID() string
	// Seat returns the Seat the device belongs to.
	Seat() *Seat
	// Source determines the type of the device.
	Source() InputSource
	// VendorID returns the vendor ID of this device, or nil if this information
	// couldn't be obtained.
	VendorID() string
	// WindowAtPosition obtains the window underneath @device, returning the
	// location of the device in @win_x and @win_y.
	WindowAtPosition() (winX int, winY int, window *Window)
	// WindowAtPositionDouble obtains the window underneath @device, returning
	// the location of the device in @win_x and @win_y in double precision.
	WindowAtPositionDouble() (winX float64, winY float64, window *Window)
	// Ungrab: release any grab on @device.
	Ungrab(time_ uint32)
	// Warp warps @device in @display to the point @x,@y on the screen @screen,
	// unless the device is confined to a window by a grab, in which case it
	// will be moved as far as allowed by the grab.
	Warp(screen Screener, x int, y int)
}

// Device object represents a single input device, such as a keyboard, a mouse,
// a touchpad, etc.
//
// See the DeviceManager documentation for more information about the various
// kinds of master and slave devices, and their relationships.
type Device struct {
	*externglib.Object
}

var (
	_ Devicer         = (*Device)(nil)
	_ gextras.Nativer = (*Device)(nil)
)

func wrapDevice(obj *externglib.Object) Devicer {
	return &Device{
		Object: obj,
	}
}

func marshalDevicer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDevice(obj), nil
}

// AssociatedDevice returns the associated device to @device, if @device is of
// type GDK_DEVICE_TYPE_MASTER, it will return the paired pointer or keyboard.
//
// If @device is of type GDK_DEVICE_TYPE_SLAVE, it will return the master device
// to which @device is attached to.
//
// If @device is of type GDK_DEVICE_TYPE_FLOATING, nil will be returned, as
// there is no associated device.
func (device *Device) AssociatedDevice() *Device {
	var _arg0 *C.GdkDevice // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_associated_device(_arg0)

	var _ret *Device // out

	_ret = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Device)

	return _ret
}

// Axes returns the axes currently available on the device.
func (device *Device) Axes() AxisFlags {
	var _arg0 *C.GdkDevice   // out
	var _cret C.GdkAxisFlags // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_axes(_arg0)

	var _axisFlags AxisFlags // out

	_axisFlags = (AxisFlags)(_cret)

	return _axisFlags
}

// AxisUse returns the axis use for @index_.
func (device *Device) AxisUse(index_ uint) AxisUse {
	var _arg0 *C.GdkDevice // out
	var _arg1 C.guint      // out
	var _cret C.GdkAxisUse // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg1 = C.guint(index_)

	_cret = C.gdk_device_get_axis_use(_arg0, _arg1)

	var _axisUse AxisUse // out

	_axisUse = (AxisUse)(_cret)

	return _axisUse
}

// DeviceType returns the device type for @device.
func (device *Device) DeviceType() DeviceType {
	var _arg0 *C.GdkDevice    // out
	var _cret C.GdkDeviceType // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_device_type(_arg0)

	var _deviceType DeviceType // out

	_deviceType = (DeviceType)(_cret)

	return _deviceType
}

// Display returns the Display to which @device pertains.
func (device *Device) Display() *Display {
	var _arg0 *C.GdkDevice  // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_display(_arg0)

	var _display *Display // out

	_display = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Display)

	return _display
}

// HasCursor determines whether the pointer follows device motion. This is not
// meaningful for keyboard devices, which don't have a pointer.
func (device *Device) HasCursor() bool {
	var _arg0 *C.GdkDevice // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_has_cursor(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Key: if @index_ has a valid keyval, this function will return true and fill
// in @keyval and @modifiers with the keyval settings.
func (device *Device) Key(index_ uint) (uint, ModifierType, bool) {
	var _arg0 *C.GdkDevice      // out
	var _arg1 C.guint           // out
	var _arg2 C.guint           // in
	var _arg3 C.GdkModifierType // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg1 = C.guint(index_)

	_cret = C.gdk_device_get_key(_arg0, _arg1, &_arg2, &_arg3)

	var _keyval uint            // out
	var _modifiers ModifierType // out
	var _ok bool                // out

	_keyval = uint(_arg2)
	_modifiers = (ModifierType)(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _keyval, _modifiers, _ok
}

// LastEventWindow gets information about which window the given pointer device
// is in, based on events that have been received so far from the display
// server. If another application has a pointer grab, or this application has a
// grab with owner_events = false, nil may be returned even if the pointer is
// physically over one of this application's windows.
func (device *Device) LastEventWindow() *Window {
	var _arg0 *C.GdkDevice // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_last_event_window(_arg0)

	var _window *Window // out

	_window = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Window)

	return _window
}

// Mode determines the mode of the device.
func (device *Device) Mode() InputMode {
	var _arg0 *C.GdkDevice   // out
	var _cret C.GdkInputMode // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_mode(_arg0)

	var _inputMode InputMode // out

	_inputMode = (InputMode)(_cret)

	return _inputMode
}

// NAxes returns the number of axes the device currently has.
func (device *Device) NAxes() int {
	var _arg0 *C.GdkDevice // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_n_axes(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NKeys returns the number of keys the device currently has.
func (device *Device) NKeys() int {
	var _arg0 *C.GdkDevice // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_n_keys(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Name determines the name of the device.
func (device *Device) Name() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Position gets the current location of @device. As a slave device coordinates
// are those of its master pointer, This function may not be called on devices
// of type GDK_DEVICE_TYPE_SLAVE, unless there is an ongoing grab on them, see
// gdk_device_grab().
func (device *Device) Position() (screen *Screen, x int, y int) {
	var _arg0 *C.GdkDevice // out
	var _arg1 *C.GdkScreen // in
	var _arg2 C.gint       // in
	var _arg3 C.gint       // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	C.gdk_device_get_position(_arg0, &_arg1, &_arg2, &_arg3)

	var _screen *Screen // out
	var _x int          // out
	var _y int          // out

	_screen = (gextras.CastObject(externglib.Take(unsafe.Pointer(_arg1)))).(*Screen)
	_x = int(_arg2)
	_y = int(_arg3)

	return _screen, _x, _y
}

// PositionDouble gets the current location of @device in double precision. As a
// slave device's coordinates are those of its master pointer, this function may
// not be called on devices of type GDK_DEVICE_TYPE_SLAVE, unless there is an
// ongoing grab on them. See gdk_device_grab().
func (device *Device) PositionDouble() (screen *Screen, x float64, y float64) {
	var _arg0 *C.GdkDevice // out
	var _arg1 *C.GdkScreen // in
	var _arg2 C.gdouble    // in
	var _arg3 C.gdouble    // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	C.gdk_device_get_position_double(_arg0, &_arg1, &_arg2, &_arg3)

	var _screen *Screen // out
	var _x float64      // out
	var _y float64      // out

	_screen = (gextras.CastObject(externglib.Take(unsafe.Pointer(_arg1)))).(*Screen)
	_x = float64(_arg2)
	_y = float64(_arg3)

	return _screen, _x, _y
}

// ProductID returns the product ID of this device, or nil if this information
// couldn't be obtained. This ID is retrieved from the device, and is thus
// constant for it. See gdk_device_get_vendor_id() for more information.
func (device *Device) ProductID() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_product_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Seat returns the Seat the device belongs to.
func (device *Device) Seat() *Seat {
	var _arg0 *C.GdkDevice // out
	var _cret *C.GdkSeat   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_seat(_arg0)

	var _seat *Seat // out

	_seat = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Seat)

	return _seat
}

// Source determines the type of the device.
func (device *Device) Source() InputSource {
	var _arg0 *C.GdkDevice     // out
	var _cret C.GdkInputSource // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_source(_arg0)

	var _inputSource InputSource // out

	_inputSource = (InputSource)(_cret)

	return _inputSource
}

// VendorID returns the vendor ID of this device, or nil if this information
// couldn't be obtained. This ID is retrieved from the device, and is thus
// constant for it.
//
// This function, together with gdk_device_get_product_id(), can be used to eg.
// compose #GSettings paths to store settings for this device.
//
//     static GSettings *
//     get_device_settings (GdkDevice *device)
//     {
//       const gchar *vendor, *product;
//       GSettings *settings;
//       GdkDevice *device;
//       gchar *path;
//
//       vendor = gdk_device_get_vendor_id (device);
//       product = gdk_device_get_product_id (device);
//
//       path = g_strdup_printf ("/org/example/app/devices/s:s/", vendor, product);
//       settings = g_settings_new_with_path (DEVICE_SCHEMA, path);
//       g_free (path);
//
//       return settings;
//     }
func (device *Device) VendorID() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_vendor_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// WindowAtPosition obtains the window underneath @device, returning the
// location of the device in @win_x and @win_y. Returns nil if the window tree
// under @device is not known to GDK (for example, belongs to another
// application).
//
// As a slave device coordinates are those of its master pointer, This function
// may not be called on devices of type GDK_DEVICE_TYPE_SLAVE, unless there is
// an ongoing grab on them, see gdk_device_grab().
func (device *Device) WindowAtPosition() (winX int, winY int, window *Window) {
	var _arg0 *C.GdkDevice // out
	var _arg1 C.gint       // in
	var _arg2 C.gint       // in
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_window_at_position(_arg0, &_arg1, &_arg2)

	var _winX int       // out
	var _winY int       // out
	var _window *Window // out

	_winX = int(_arg1)
	_winY = int(_arg2)
	_window = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Window)

	return _winX, _winY, _window
}

// WindowAtPositionDouble obtains the window underneath @device, returning the
// location of the device in @win_x and @win_y in double precision. Returns nil
// if the window tree under @device is not known to GDK (for example, belongs to
// another application).
//
// As a slave device coordinates are those of its master pointer, This function
// may not be called on devices of type GDK_DEVICE_TYPE_SLAVE, unless there is
// an ongoing grab on them, see gdk_device_grab().
func (device *Device) WindowAtPositionDouble() (winX float64, winY float64, window *Window) {
	var _arg0 *C.GdkDevice // out
	var _arg1 C.gdouble    // in
	var _arg2 C.gdouble    // in
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_window_at_position_double(_arg0, &_arg1, &_arg2)

	var _winX float64   // out
	var _winY float64   // out
	var _window *Window // out

	_winX = float64(_arg1)
	_winY = float64(_arg2)
	_window = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Window)

	return _winX, _winY, _window
}

// Ungrab: release any grab on @device.
//
// Deprecated: Use gdk_seat_ungrab() instead.
func (device *Device) Ungrab(time_ uint32) {
	var _arg0 *C.GdkDevice // out
	var _arg1 C.guint32    // out

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg1 = C.guint32(time_)

	C.gdk_device_ungrab(_arg0, _arg1)
}

// Warp warps @device in @display to the point @x,@y on the screen @screen,
// unless the device is confined to a window by a grab, in which case it will be
// moved as far as allowed by the grab. Warping the pointer creates events as if
// the user had moved the mouse instantaneously to the destination.
//
// Note that the pointer should normally be under the control of the user. This
// function was added to cover some rare use cases like keyboard navigation
// support for the color picker in the ColorSelectionDialog.
func (device *Device) Warp(screen Screener, x int, y int) {
	var _arg0 *C.GdkDevice // out
	var _arg1 *C.GdkScreen // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer((screen).(gextras.Nativer).Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	C.gdk_device_warp(_arg0, _arg1, _arg2, _arg3)
}

// TimeCoord stores a single event in a motion history.
type TimeCoord struct {
	native C.GdkTimeCoord
}

// Native returns the underlying C source pointer.
func (t *TimeCoord) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}
