// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gdk4_Device_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gdk4_Device_ConnectToolChanged(gpointer, GdkDeviceTool*, guintptr);
import "C"

// glib.Type values for gdkdevice.go.
var (
	GTypeInputSource = coreglib.Type(C.gdk_input_source_get_type())
	GTypeDevice      = coreglib.Type(C.gdk_device_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeInputSource, F: marshalInputSource},
		{T: GTypeDevice, F: marshalDevice},
	})
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
	// SourceKeyboard: device is a keyboard.
	SourceKeyboard
	// SourceTouchscreen: device is a direct-input touch device, such as a
	// touchscreen or tablet.
	SourceTouchscreen
	// SourceTouchpad: device is an indirect touch device, such as a touchpad.
	SourceTouchpad
	// SourceTrackpoint: device is a trackpoint.
	SourceTrackpoint
	// SourceTabletPad: device is a "pad", a collection of buttons, rings and
	// strips found in drawing tablets.
	SourceTabletPad
)

func marshalInputSource(p uintptr) (interface{}, error) {
	return InputSource(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for InputSource.
func (i InputSource) String() string {
	switch i {
	case SourceMouse:
		return "Mouse"
	case SourcePen:
		return "Pen"
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

// Device: GdkDevice object represents an input device, such as a keyboard, a
// mouse, or a touchpad.
//
// See the gdk.Seat documentation for more information about the various kinds
// of devices, and their relationships.
type Device struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Device)(nil)
)

// Devicer describes types inherited from class Device.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Devicer interface {
	coreglib.Objector
	baseDevice() *Device
}

var _ Devicer = (*Device)(nil)

func wrapDevice(obj *coreglib.Object) *Device {
	return &Device{
		Object: obj,
	}
}

func marshalDevice(p uintptr) (interface{}, error) {
	return wrapDevice(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (device *Device) baseDevice() *Device {
	return device
}

// BaseDevice returns the underlying base object.
func BaseDevice(obj Devicer) *Device {
	return obj.baseDevice()
}

//export _gotk4_gdk4_Device_ConnectChanged
func _gotk4_gdk4_Device_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectChanged is emitted either when the the number of either axes or keys
// changes.
//
// On X11 this will normally happen when the physical device routing events
// through the logical device changes (for example, user switches from the USB
// mouse to a tablet); in that case the logical device will change to reflect
// the axes and keys on the new physical device.
func (device *Device) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(device, "changed", false, unsafe.Pointer(C._gotk4_gdk4_Device_ConnectChanged), f)
}

//export _gotk4_gdk4_Device_ConnectToolChanged
func _gotk4_gdk4_Device_ConnectToolChanged(arg0 C.gpointer, arg1 *C.GdkDeviceTool, arg2 C.guintptr) {
	var f func(tool *DeviceTool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(tool *DeviceTool))
	}

	var _tool *DeviceTool // out

	_tool = wrapDeviceTool(coreglib.Take(unsafe.Pointer(arg1)))

	f(_tool)
}

// ConnectToolChanged is emitted on pen/eraser devices whenever tools enter or
// leave proximity.
func (device *Device) ConnectToolChanged(f func(tool *DeviceTool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(device, "tool-changed", false, unsafe.Pointer(C._gotk4_gdk4_Device_ConnectToolChanged), f)
}

// CapsLockState retrieves whether the Caps Lock modifier of the keyboard is
// locked.
//
// This is only relevant for keyboard devices.
//
// The function returns the following values:
//
//    - ok: TRUE if Caps Lock is on for device.
//
func (device *Device) CapsLockState() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**Device)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_caps_lock_state", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DeviceTool retrieves the current tool for device.
//
// The function returns the following values:
//
//    - deviceTool: GdkDeviceTool, or NULL.
//
func (device *Device) DeviceTool() *DeviceTool {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**Device)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_device_tool", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _deviceTool *DeviceTool // out

	_deviceTool = wrapDeviceTool(coreglib.Take(unsafe.Pointer(_cret)))

	return _deviceTool
}

// Display returns the GdkDisplay to which device pertains.
//
// The function returns the following values:
//
//    - display: GdkDisplay.
//
func (device *Device) Display() *Display {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**Device)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_display", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _display *Display // out

	_display = wrapDisplay(coreglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// HasCursor determines whether the pointer follows device motion.
//
// This is not meaningful for keyboard devices, which don't have a pointer.
//
// The function returns the following values:
//
//    - ok: TRUE if the pointer follows device motion.
//
func (device *Device) HasCursor() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**Device)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_has_cursor", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Name: name of the device, suitable for showing in a user interface.
//
// The function returns the following values:
//
//    - utf8: name.
//
func (device *Device) Name() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**Device)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_name", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// NumLockState retrieves whether the Num Lock modifier of the keyboard is
// locked.
//
// This is only relevant for keyboard devices.
//
// The function returns the following values:
//
//    - ok: TRUE if Num Lock is on for device.
//
func (device *Device) NumLockState() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**Device)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_num_lock_state", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NumTouches retrieves the number of touch points associated to device.
//
// The function returns the following values:
//
//    - guint: number of touch points.
//
func (device *Device) NumTouches() uint32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**Device)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_num_touches", args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// ProductID returns the product ID of this device.
//
// This ID is retrieved from the device, and does not change. See
// gdk.Device.GetVendorID() for more information.
//
// The function returns the following values:
//
//    - utf8 (optional): product ID, or NULL.
//
func (device *Device) ProductID() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**Device)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_product_id", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ScrollLockState retrieves whether the Scroll Lock modifier of the keyboard is
// locked.
//
// This is only relevant for keyboard devices.
//
// The function returns the following values:
//
//    - ok: TRUE if Scroll Lock is on for device.
//
func (device *Device) ScrollLockState() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**Device)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_scroll_lock_state", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Seat returns the GdkSeat the device belongs to.
//
// The function returns the following values:
//
//    - seat: GdkSeat.
//
func (device *Device) Seat() Seater {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**Device)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_seat", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _seat Seater // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Seater is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Seater)
			return ok
		})
		rv, ok := casted.(Seater)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Seater")
		}
		_seat = rv
	}

	return _seat
}

// Timestamp returns the timestamp of the last activity for this device.
//
// In practice, this means the timestamp of the last event that was received
// from the OS for this device. (GTK may occasionally produce events for a
// device that are not received from the OS, and will not update the timestamp).
//
// The function returns the following values:
//
//    - guint32: timestamp of the last activity for this device.
//
func (device *Device) Timestamp() uint32 {
	var args [1]girepository.Argument
	var _arg0 *C.void   // out
	var _cret C.guint32 // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**Device)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_timestamp", args[:], nil)
	_cret = *(*C.guint32)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// VendorID returns the vendor ID of this device.
//
// This ID is retrieved from the device, and does not change.
//
// This function, together with gdk.Device.GetProductID(), can be used to eg.
// compose GSettings paths to store settings for this device.
//
//     static GSettings *
//     get_device_settings (GdkDevice *device)
//     {
//       const char *vendor, *product;
//       GSettings *settings;
//       GdkDevice *device;
//       char *path;
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
//
// The function returns the following values:
//
//    - utf8 (optional): vendor ID, or NULL.
//
func (device *Device) VendorID() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**Device)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_vendor_id", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// HasBidiLayouts determines if layouts for both right-to-left and left-to-right
// languages are in use on the keyboard.
//
// This is only relevant for keyboard devices.
//
// The function returns the following values:
//
//    - ok: TRUE if there are layouts with both directions, FALSE otherwise.
//
func (device *Device) HasBidiLayouts() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**Device)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("has_bidi_layouts", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
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
