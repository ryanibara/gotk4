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
import "C"

// glib.Type values for gdkdevicetool.go.
var (
	GTypeDeviceToolType = coreglib.Type(C.gdk_device_tool_type_get_type())
	GTypeDeviceTool     = coreglib.Type(C.gdk_device_tool_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeDeviceToolType, F: marshalDeviceToolType},
		{T: GTypeDeviceTool, F: marshalDeviceTool},
	})
}

// DeviceToolType indicates the specific type of tool being used being a tablet.
// Such as an airbrush, pencil, etc.
type DeviceToolType C.gint

const (
	// DeviceToolTypeUnknown: tool is of an unknown type.
	DeviceToolTypeUnknown DeviceToolType = iota
	// DeviceToolTypePen: tool is a standard tablet stylus.
	DeviceToolTypePen
	// DeviceToolTypeEraser: tool is standard tablet eraser.
	DeviceToolTypeEraser
	// DeviceToolTypeBrush: tool is a brush stylus.
	DeviceToolTypeBrush
	// DeviceToolTypePencil: tool is a pencil stylus.
	DeviceToolTypePencil
	// DeviceToolTypeAirbrush: tool is an airbrush stylus.
	DeviceToolTypeAirbrush
	// DeviceToolTypeMouse: tool is a mouse.
	DeviceToolTypeMouse
	// DeviceToolTypeLens: tool is a lens cursor.
	DeviceToolTypeLens
)

func marshalDeviceToolType(p uintptr) (interface{}, error) {
	return DeviceToolType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for DeviceToolType.
func (d DeviceToolType) String() string {
	switch d {
	case DeviceToolTypeUnknown:
		return "Unknown"
	case DeviceToolTypePen:
		return "Pen"
	case DeviceToolTypeEraser:
		return "Eraser"
	case DeviceToolTypeBrush:
		return "Brush"
	case DeviceToolTypePencil:
		return "Pencil"
	case DeviceToolTypeAirbrush:
		return "Airbrush"
	case DeviceToolTypeMouse:
		return "Mouse"
	case DeviceToolTypeLens:
		return "Lens"
	default:
		return fmt.Sprintf("DeviceToolType(%d)", d)
	}
}

// DeviceTool: physical tool associated to a GdkDevice.
type DeviceTool struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DeviceTool)(nil)
)

func wrapDeviceTool(obj *coreglib.Object) *DeviceTool {
	return &DeviceTool{
		Object: obj,
	}
}

func marshalDeviceTool(p uintptr) (interface{}, error) {
	return wrapDeviceTool(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// HardwareID gets the hardware ID of this tool, or 0 if it's not known.
//
// When non-zero, the identificator is unique for the given tool model, meaning
// that two identical tools will share the same hardware_id, but will have
// different serial numbers (see gdk.DeviceTool.GetSerial()).
//
// This is a more concrete (and device specific) method to identify a
// GdkDeviceTool than gdk.DeviceTool.GetToolType(), as a tablet may support
// multiple devices with the same GdkDeviceToolType, but different hardware
// identificators.
//
// The function returns the following values:
//
//    - guint64: hardware identificator of this tool.
//
func (tool *DeviceTool) HardwareID() uint64 {
	var args [1]girepository.Argument
	var _arg0 *C.void   // out
	var _cret C.guint64 // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(tool).Native()))
	*(**DeviceTool)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "DeviceTool").InvokeMethod("get_hardware_id", args[:], nil)
	_cret = *(*C.guint64)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(tool)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// Serial gets the serial number of this tool.
//
// This value can be used to identify a physical tool (eg. a tablet pen) across
// program executions.
//
// The function returns the following values:
//
//    - guint64: serial ID for this tool.
//
func (tool *DeviceTool) Serial() uint64 {
	var args [1]girepository.Argument
	var _arg0 *C.void   // out
	var _cret C.guint64 // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(tool).Native()))
	*(**DeviceTool)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "DeviceTool").InvokeMethod("get_serial", args[:], nil)
	_cret = *(*C.guint64)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(tool)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}
