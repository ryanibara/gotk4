// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
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

const MAX_TIMECOORD_AXES = C.MAX_TIMECOORD_AXES

// DeviceType indicates the device type. See
// [above][GdkDeviceManager.description] for more information about the meaning
// of these device types.
type DeviceType C.gint

const (
	// DeviceTypeMaster: device is a master (or virtual) device. There will be
	// an associated focus indicator on the screen.
	DeviceTypeMaster DeviceType = iota
	// DeviceTypeSlave: device is a slave (or physical) device.
	DeviceTypeSlave
	// DeviceTypeFloating: device is a physical device, currently not attached
	// to any virtual device.
	DeviceTypeFloating
)

func marshalDeviceType(p uintptr) (interface{}, error) {
	return DeviceType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for DeviceType.
func (d DeviceType) String() string {
	switch d {
	case DeviceTypeMaster:
		return "Master"
	case DeviceTypeSlave:
		return "Slave"
	case DeviceTypeFloating:
		return "Floating"
	default:
		return fmt.Sprintf("DeviceType(%d)", d)
	}
}

// InputMode: enumeration that describes the mode of an input device.
type InputMode C.gint

const (
	// ModeDisabled: device is disabled and will not report any events.
	ModeDisabled InputMode = iota
	// ModeScreen: device is enabled. The device’s coordinate space maps to the
	// entire screen.
	ModeScreen
	// ModeWindow: device is enabled. The device’s coordinate space is mapped to
	// a single window. The manner in which this window is chosen is undefined,
	// but it will typically be the same way in which the focus window for key
	// events is determined.
	ModeWindow
)

func marshalInputMode(p uintptr) (interface{}, error) {
	return InputMode(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for InputMode.
func (i InputMode) String() string {
	switch i {
	case ModeDisabled:
		return "Disabled"
	case ModeScreen:
		return "Screen"
	case ModeWindow:
		return "Window"
	default:
		return fmt.Sprintf("InputMode(%d)", i)
	}
}

// InputSource: enumeration describing the type of an input device in general
// terms.
type InputSource C.gint

const (
	// SourceMouse: device is a mouse. (This will be reported for the core
	// pointer, even if it is something else, such as a trackball.).
	SourceMouse InputSource = iota
	// SourcePen: device is a stylus of a graphics tablet or similar device.
	SourcePen
	// SourceEraser: device is an eraser. Typically, this would be the other end
	// of a stylus on a graphics tablet.
	SourceEraser
	// SourceCursor: device is a graphics tablet “puck” or similar device.
	SourceCursor
	// SourceKeyboard: device is a keyboard.
	SourceKeyboard
	// SourceTouchscreen: device is a direct-input touch device, such as a
	// touchscreen or tablet. This device type has been added in 3.4.
	SourceTouchscreen
	// SourceTouchpad: device is an indirect touch device, such as a touchpad.
	// This device type has been added in 3.4.
	SourceTouchpad
	// SourceTrackpoint: device is a trackpoint. This device type has been added
	// in 3.22.
	SourceTrackpoint
	// SourceTabletPad: device is a "pad", a collection of buttons, rings and
	// strips found in drawing tablets. This device type has been added in 3.22.
	SourceTabletPad
)

func marshalInputSource(p uintptr) (interface{}, error) {
	return InputSource(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for InputSource.
func (i InputSource) String() string {
	switch i {
	case SourceMouse:
		return "Mouse"
	case SourcePen:
		return "Pen"
	case SourceEraser:
		return "Eraser"
	case SourceCursor:
		return "Cursor"
	case SourceKeyboard:
		return "Keyboard"
	case SourceTouchscreen:
		return "Touchscreen"
	case SourceTouchpad:
		return "Touchpad"
	case SourceTrackpoint:
		return "Trackpoint"
	case SourceTabletPad:
		return "TabletPad"
	default:
		return fmt.Sprintf("InputSource(%d)", i)
	}
}

// Device object represents a single input device, such as a keyboard, a mouse,
// a touchpad, etc.
//
// See the DeviceManager documentation for more information about the various
// kinds of master and slave devices, and their relationships.
type Device struct {
	*externglib.Object
}

// Devicer describes types inherited from class Device.
// To get the original type, the caller must assert this to an interface or
// another type.
type Devicer interface {
	externglib.Objector

	// BaseDevice returns the underlying base class.
	BaseDevice() *Device
}

var _ Devicer = (*Device)(nil)

func wrapDevice(obj *externglib.Object) *Device {
	return &Device{
		Object: obj,
	}
}

func marshalDevicer(p uintptr) (interface{}, error) {
	return wrapDevice(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AssociatedDevice returns the associated device to device, if device is of
// type GDK_DEVICE_TYPE_MASTER, it will return the paired pointer or keyboard.
//
// If device is of type GDK_DEVICE_TYPE_SLAVE, it will return the master device
// to which device is attached to.
//
// If device is of type GDK_DEVICE_TYPE_FLOATING, NULL will be returned, as
// there is no associated device.
func (device *Device) AssociatedDevice() Devicer {
	var _arg0 *C.GdkDevice // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_associated_device(_arg0)
	runtime.KeepAlive(device)

	var _ret Devicer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Devicer)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Devicer")
			}
			_ret = rv
		}
	}

	return _ret
}

// Axes returns the axes currently available on the device.
func (device *Device) Axes() AxisFlags {
	var _arg0 *C.GdkDevice   // out
	var _cret C.GdkAxisFlags // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_axes(_arg0)
	runtime.KeepAlive(device)

	var _axisFlags AxisFlags // out

	_axisFlags = AxisFlags(_cret)

	return _axisFlags
}

// AxisUse returns the axis use for index_.
//
// The function takes the following parameters:
//
//    - index_: index of the axis.
//
func (device *Device) AxisUse(index_ uint) AxisUse {
	var _arg0 *C.GdkDevice // out
	var _arg1 C.guint      // out
	var _cret C.GdkAxisUse // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg1 = C.guint(index_)

	_cret = C.gdk_device_get_axis_use(_arg0, _arg1)
	runtime.KeepAlive(device)
	runtime.KeepAlive(index_)

	var _axisUse AxisUse // out

	_axisUse = AxisUse(_cret)

	return _axisUse
}

// DeviceType returns the device type for device.
func (device *Device) DeviceType() DeviceType {
	var _arg0 *C.GdkDevice    // out
	var _cret C.GdkDeviceType // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_device_type(_arg0)
	runtime.KeepAlive(device)

	var _deviceType DeviceType // out

	_deviceType = DeviceType(_cret)

	return _deviceType
}

// Display returns the Display to which device pertains.
func (device *Device) Display() *Display {
	var _arg0 *C.GdkDevice  // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_display(_arg0)
	runtime.KeepAlive(device)

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// HasCursor determines whether the pointer follows device motion. This is not
// meaningful for keyboard devices, which don't have a pointer.
func (device *Device) HasCursor() bool {
	var _arg0 *C.GdkDevice // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_has_cursor(_arg0)
	runtime.KeepAlive(device)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Key: if index_ has a valid keyval, this function will return TRUE and fill in
// keyval and modifiers with the keyval settings.
//
// The function takes the following parameters:
//
//    - index_: index of the macro button to get.
//
func (device *Device) Key(index_ uint) (uint, ModifierType, bool) {
	var _arg0 *C.GdkDevice      // out
	var _arg1 C.guint           // out
	var _arg2 C.guint           // in
	var _arg3 C.GdkModifierType // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg1 = C.guint(index_)

	_cret = C.gdk_device_get_key(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(device)
	runtime.KeepAlive(index_)

	var _keyval uint            // out
	var _modifiers ModifierType // out
	var _ok bool                // out

	_keyval = uint(_arg2)
	_modifiers = ModifierType(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _keyval, _modifiers, _ok
}

// LastEventWindow gets information about which window the given pointer device
// is in, based on events that have been received so far from the display
// server. If another application has a pointer grab, or this application has a
// grab with owner_events = FALSE, NULL may be returned even if the pointer is
// physically over one of this application's windows.
func (device *Device) LastEventWindow() Windower {
	var _arg0 *C.GdkDevice // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_last_event_window(_arg0)
	runtime.KeepAlive(device)

	var _window Windower // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Windower)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Windower")
			}
			_window = rv
		}
	}

	return _window
}

// Mode determines the mode of the device.
func (device *Device) Mode() InputMode {
	var _arg0 *C.GdkDevice   // out
	var _cret C.GdkInputMode // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_mode(_arg0)
	runtime.KeepAlive(device)

	var _inputMode InputMode // out

	_inputMode = InputMode(_cret)

	return _inputMode
}

// NAxes returns the number of axes the device currently has.
func (device *Device) NAxes() int {
	var _arg0 *C.GdkDevice // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_n_axes(_arg0)
	runtime.KeepAlive(device)

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
	runtime.KeepAlive(device)

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
	runtime.KeepAlive(device)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Position gets the current location of device. As a slave device coordinates
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
	runtime.KeepAlive(device)

	var _screen *Screen // out
	var _x int          // out
	var _y int          // out

	if _arg1 != nil {
		_screen = wrapScreen(externglib.Take(unsafe.Pointer(_arg1)))
	}
	_x = int(_arg2)
	_y = int(_arg3)

	return _screen, _x, _y
}

// PositionDouble gets the current location of device in double precision. As a
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
	runtime.KeepAlive(device)

	var _screen *Screen // out
	var _x float64      // out
	var _y float64      // out

	if _arg1 != nil {
		_screen = wrapScreen(externglib.Take(unsafe.Pointer(_arg1)))
	}
	_x = float64(_arg2)
	_y = float64(_arg3)

	return _screen, _x, _y
}

// ProductID returns the product ID of this device, or NULL if this information
// couldn't be obtained. This ID is retrieved from the device, and is thus
// constant for it. See gdk_device_get_vendor_id() for more information.
func (device *Device) ProductID() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_product_id(_arg0)
	runtime.KeepAlive(device)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Seat returns the Seat the device belongs to.
func (device *Device) Seat() Seater {
	var _arg0 *C.GdkDevice // out
	var _cret *C.GdkSeat   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_seat(_arg0)
	runtime.KeepAlive(device)

	var _seat Seater // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Seater is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Seater)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Seater")
		}
		_seat = rv
	}

	return _seat
}

// Source determines the type of the device.
func (device *Device) Source() InputSource {
	var _arg0 *C.GdkDevice     // out
	var _cret C.GdkInputSource // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_source(_arg0)
	runtime.KeepAlive(device)

	var _inputSource InputSource // out

	_inputSource = InputSource(_cret)

	return _inputSource
}

// VendorID returns the vendor ID of this device, or NULL if this information
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
//     }.
func (device *Device) VendorID() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_vendor_id(_arg0)
	runtime.KeepAlive(device)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// WindowAtPosition obtains the window underneath device, returning the location
// of the device in win_x and win_y. Returns NULL if the window tree under
// device is not known to GDK (for example, belongs to another application).
//
// As a slave device coordinates are those of its master pointer, This function
// may not be called on devices of type GDK_DEVICE_TYPE_SLAVE, unless there is
// an ongoing grab on them, see gdk_device_grab().
func (device *Device) WindowAtPosition() (winX int, winY int, window Windower) {
	var _arg0 *C.GdkDevice // out
	var _arg1 C.gint       // in
	var _arg2 C.gint       // in
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_window_at_position(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(device)

	var _winX int        // out
	var _winY int        // out
	var _window Windower // out

	_winX = int(_arg1)
	_winY = int(_arg2)
	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Windower)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Windower")
			}
			_window = rv
		}
	}

	return _winX, _winY, _window
}

// WindowAtPositionDouble obtains the window underneath device, returning the
// location of the device in win_x and win_y in double precision. Returns NULL
// if the window tree under device is not known to GDK (for example, belongs to
// another application).
//
// As a slave device coordinates are those of its master pointer, This function
// may not be called on devices of type GDK_DEVICE_TYPE_SLAVE, unless there is
// an ongoing grab on them, see gdk_device_grab().
func (device *Device) WindowAtPositionDouble() (winX float64, winY float64, window Windower) {
	var _arg0 *C.GdkDevice // out
	var _arg1 C.gdouble    // in
	var _arg2 C.gdouble    // in
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_window_at_position_double(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(device)

	var _winX float64    // out
	var _winY float64    // out
	var _window Windower // out

	_winX = float64(_arg1)
	_winY = float64(_arg2)
	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Windower)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Windower")
			}
			_window = rv
		}
	}

	return _winX, _winY, _window
}

// Grab grabs the device so that all events coming from this device are passed
// to this application until the device is ungrabbed with gdk_device_ungrab(),
// or the window becomes unviewable. This overrides any previous grab on the
// device by this client.
//
// Note that device and window need to be on the same display.
//
// Device grabs are used for operations which need complete control over the
// given device events (either pointer or keyboard). For example in GTK+ this is
// used for Drag and Drop operations, popup menus and such.
//
// Note that if the event mask of an X window has selected both button press and
// button release events, then a button press event will cause an automatic
// pointer grab until the button is released. X does this automatically since
// most applications expect to receive button press and release events in pairs.
// It is equivalent to a pointer grab on the window with owner_events set to
// TRUE.
//
// If you set up anything at the time you take the grab that needs to be cleaned
// up when the grab ends, you should handle the EventGrabBroken events that are
// emitted when the grab ends unvoluntarily.
//
// Deprecated: Use gdk_seat_grab() instead.
//
// The function takes the following parameters:
//
//    - window which will own the grab (the grab window).
//    - grabOwnership specifies the grab ownership.
//    - ownerEvents: if FALSE then all device events are reported with respect
//    to window and are only reported if selected by event_mask. If TRUE then
//    pointer events for this application are reported as normal, but pointer
//    events outside this application are reported with respect to window and
//    only if selected by event_mask. In either mode, unreported events are
//    discarded.
//    - eventMask specifies the event mask, which is used in accordance with
//    owner_events.
//    - cursor to display while the grab is active if the device is a pointer.
//    If this is NULL then the normal cursors are used for window and its
//    descendants, and the cursor for window is used elsewhere.
//    - time_: timestamp of the event which led to this pointer grab. This
//    usually comes from the Event struct, though GDK_CURRENT_TIME can be used
//    if the time isn’t known.
//
func (device *Device) Grab(window Windower, grabOwnership GrabOwnership, ownerEvents bool, eventMask EventMask, cursor Cursorrer, time_ uint32) GrabStatus {
	var _arg0 *C.GdkDevice       // out
	var _arg1 *C.GdkWindow       // out
	var _arg2 C.GdkGrabOwnership // out
	var _arg3 C.gboolean         // out
	var _arg4 C.GdkEventMask     // out
	var _arg5 *C.GdkCursor       // out
	var _arg6 C.guint32          // out
	var _cret C.GdkGrabStatus    // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = C.GdkGrabOwnership(grabOwnership)
	if ownerEvents {
		_arg3 = C.TRUE
	}
	_arg4 = C.GdkEventMask(eventMask)
	if cursor != nil {
		_arg5 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))
	}
	_arg6 = C.guint32(time_)

	_cret = C.gdk_device_grab(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(device)
	runtime.KeepAlive(window)
	runtime.KeepAlive(grabOwnership)
	runtime.KeepAlive(ownerEvents)
	runtime.KeepAlive(eventMask)
	runtime.KeepAlive(cursor)
	runtime.KeepAlive(time_)

	var _grabStatus GrabStatus // out

	_grabStatus = GrabStatus(_cret)

	return _grabStatus
}

// ListSlaveDevices: if the device if of type GDK_DEVICE_TYPE_MASTER, it will
// return the list of slave devices attached to it, otherwise it will return
// NULL.
func (device *Device) ListSlaveDevices() []Devicer {
	var _arg0 *C.GdkDevice // out
	var _cret *C.GList     // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_list_slave_devices(_arg0)
	runtime.KeepAlive(device)

	var _list []Devicer // out

	if _cret != nil {
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
	}

	return _list
}

// SetAxisUse specifies how an axis of a device is used.
//
// The function takes the following parameters:
//
//    - index_: index of the axis.
//    - use specifies how the axis is used.
//
func (device *Device) SetAxisUse(index_ uint, use AxisUse) {
	var _arg0 *C.GdkDevice // out
	var _arg1 C.guint      // out
	var _arg2 C.GdkAxisUse // out

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg1 = C.guint(index_)
	_arg2 = C.GdkAxisUse(use)

	C.gdk_device_set_axis_use(_arg0, _arg1, _arg2)
	runtime.KeepAlive(device)
	runtime.KeepAlive(index_)
	runtime.KeepAlive(use)
}

// SetKey specifies the X key event to generate when a macro button of a device
// is pressed.
//
// The function takes the following parameters:
//
//    - index_: index of the macro button to set.
//    - keyval to generate.
//    - modifiers to set.
//
func (device *Device) SetKey(index_, keyval uint, modifiers ModifierType) {
	var _arg0 *C.GdkDevice      // out
	var _arg1 C.guint           // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg1 = C.guint(index_)
	_arg2 = C.guint(keyval)
	_arg3 = C.GdkModifierType(modifiers)

	C.gdk_device_set_key(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(device)
	runtime.KeepAlive(index_)
	runtime.KeepAlive(keyval)
	runtime.KeepAlive(modifiers)
}

// SetMode sets a the mode of an input device. The mode controls if the device
// is active and whether the device’s range is mapped to the entire screen or to
// a single window.
//
// Note: This is only meaningful for floating devices, master devices (and
// slaves connected to these) drive the pointer cursor, which is not limited by
// the input mode.
//
// The function takes the following parameters:
//
//    - mode: input mode.
//
func (device *Device) SetMode(mode InputMode) bool {
	var _arg0 *C.GdkDevice   // out
	var _arg1 C.GdkInputMode // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg1 = C.GdkInputMode(mode)

	_cret = C.gdk_device_set_mode(_arg0, _arg1)
	runtime.KeepAlive(device)
	runtime.KeepAlive(mode)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Ungrab: release any grab on device.
//
// Deprecated: Use gdk_seat_ungrab() instead.
//
// The function takes the following parameters:
//
//    - time_: timestap (e.g. GDK_CURRENT_TIME).
//
func (device *Device) Ungrab(time_ uint32) {
	var _arg0 *C.GdkDevice // out
	var _arg1 C.guint32    // out

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg1 = C.guint32(time_)

	C.gdk_device_ungrab(_arg0, _arg1)
	runtime.KeepAlive(device)
	runtime.KeepAlive(time_)
}

// Warp warps device in display to the point x,y on the screen screen, unless
// the device is confined to a window by a grab, in which case it will be moved
// as far as allowed by the grab. Warping the pointer creates events as if the
// user had moved the mouse instantaneously to the destination.
//
// Note that the pointer should normally be under the control of the user. This
// function was added to cover some rare use cases like keyboard navigation
// support for the color picker in the ColorSelectionDialog.
//
// The function takes the following parameters:
//
//    - screen to warp device to.
//    - x: x coordinate of the destination.
//    - y: y coordinate of the destination.
//
func (device *Device) Warp(screen *Screen, x, y int) {
	var _arg0 *C.GdkDevice // out
	var _arg1 *C.GdkScreen // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	C.gdk_device_warp(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(device)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// BaseDevice returns device.
func (device *Device) BaseDevice() *Device {
	return device
}

// ConnectChanged signal is emitted either when the Device has changed the
// number of either axes or keys. For example In X this will normally happen
// when the slave device routing events through the master device changes (for
// example, user switches from the USB mouse to a tablet), in that case the
// master device will change to reflect the new slave device axes and keys.
func (device *Device) ConnectChanged(f func()) externglib.SignalHandle {
	return device.Connect("changed", f)
}

// ConnectToolChanged signal is emitted on pen/eraser Devices whenever tools
// enter or leave proximity.
func (device *Device) ConnectToolChanged(f func(tool DeviceTool)) externglib.SignalHandle {
	return device.Connect("tool-changed", f)
}

// DeviceGrabInfoLibgtkOnly determines information about the current keyboard
// grab. This is not public API and must not be used by applications.
//
// Deprecated: The symbol was never meant to be used outside of GTK+.
//
// The function takes the following parameters:
//
//    - display for which to get the grab information.
//    - device to get the grab information from.
//
func DeviceGrabInfoLibgtkOnly(display *Display, device Devicer) (grabWindow Windower, ownerEvents bool, ok bool) {
	var _arg1 *C.GdkDisplay // out
	var _arg2 *C.GdkDevice  // out
	var _arg3 *C.GdkWindow  // in
	var _arg4 C.gboolean    // in
	var _cret C.gboolean    // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_grab_info_libgtk_only(_arg1, _arg2, &_arg3, &_arg4)
	runtime.KeepAlive(display)
	runtime.KeepAlive(device)

	var _grabWindow Windower // out
	var _ownerEvents bool    // out
	var _ok bool             // out

	{
		objptr := unsafe.Pointer(_arg3)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Windower)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Windower")
		}
		_grabWindow = rv
	}
	if _arg4 != 0 {
		_ownerEvents = true
	}
	if _cret != 0 {
		_ok = true
	}

	return _grabWindow, _ownerEvents, _ok
}

// TimeCoord stores a single event in a motion history.
//
// An instance of this type is always passed by reference.
type TimeCoord struct {
	*timeCoord
}

// timeCoord is the struct that's finalized.
type timeCoord struct {
	native *C.GdkTimeCoord
}

// Time: timestamp for this event.
func (t *TimeCoord) Time() uint32 {
	var v uint32 // out
	v = uint32(t.native.time)
	return v
}

// Axes values of the device’s axes.
func (t *TimeCoord) Axes() [128]float64 {
	var v [128]float64 // out
	v = *(*[128]float64)(unsafe.Pointer(&t.native.axes))
	return v
}
