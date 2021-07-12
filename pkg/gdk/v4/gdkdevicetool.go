// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_device_tool_type_get_type()), F: marshalDeviceToolType},
		{T: externglib.Type(C.gdk_device_tool_get_type()), F: marshalDeviceTooler},
	})
}

// DeviceToolType indicates the specific type of tool being used being a tablet.
// Such as an airbrush, pencil, etc.
type DeviceToolType int

const (
	// Unknown: tool is of an unknown type.
	DeviceToolTypeUnknown DeviceToolType = iota
	// Pen: tool is a standard tablet stylus.
	DeviceToolTypePen
	// Eraser: tool is standard tablet eraser.
	DeviceToolTypeEraser
	// Brush: tool is a brush stylus.
	DeviceToolTypeBrush
	// Pencil: tool is a pencil stylus.
	DeviceToolTypePencil
	// Airbrush: tool is an airbrush stylus.
	DeviceToolTypeAirbrush
	// Mouse: tool is a mouse.
	DeviceToolTypeMouse
	// Lens: tool is a lens cursor.
	DeviceToolTypeLens
)

func marshalDeviceToolType(p uintptr) (interface{}, error) {
	return DeviceToolType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// DeviceTooler describes DeviceTool's methods.
type DeviceTooler interface {
	// Axes gets the axes of the tool.
	Axes() AxisFlags
	// HardwareID gets the hardware ID of this tool, or 0 if it's not known.
	HardwareID() uint64
	// Serial gets the serial number of this tool.
	Serial() uint64
	// ToolType gets the `GdkDeviceToolType` of the tool.
	ToolType() DeviceToolType
}

// DeviceTool: physical tool associated to a `GdkDevice`.
type DeviceTool struct {
	*externglib.Object
}

var (
	_ DeviceTooler    = (*DeviceTool)(nil)
	_ gextras.Nativer = (*DeviceTool)(nil)
)

func wrapDeviceTool(obj *externglib.Object) *DeviceTool {
	return &DeviceTool{
		Object: obj,
	}
}

func marshalDeviceTooler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDeviceTool(obj), nil
}

// Axes gets the axes of the tool.
func (tool *DeviceTool) Axes() AxisFlags {
	var _arg0 *C.GdkDeviceTool // out
	var _cret C.GdkAxisFlags   // in

	_arg0 = (*C.GdkDeviceTool)(unsafe.Pointer(tool.Native()))

	_cret = C.gdk_device_tool_get_axes(_arg0)

	var _axisFlags AxisFlags // out

	_axisFlags = AxisFlags(_cret)

	return _axisFlags
}

// HardwareID gets the hardware ID of this tool, or 0 if it's not known.
//
// When non-zero, the identificator is unique for the given tool model, meaning
// that two identical tools will share the same @hardware_id, but will have
// different serial numbers (see [method@Gdk.DeviceTool.get_serial]).
//
// This is a more concrete (and device specific) method to identify a
// `GdkDeviceTool` than [method@Gdk.DeviceTool.get_tool_type], as a tablet may
// support multiple devices with the same `GdkDeviceToolType`, but different
// hardware identificators.
func (tool *DeviceTool) HardwareID() uint64 {
	var _arg0 *C.GdkDeviceTool // out
	var _cret C.guint64        // in

	_arg0 = (*C.GdkDeviceTool)(unsafe.Pointer(tool.Native()))

	_cret = C.gdk_device_tool_get_hardware_id(_arg0)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// Serial gets the serial number of this tool.
//
// This value can be used to identify a physical tool (eg. a tablet pen) across
// program executions.
func (tool *DeviceTool) Serial() uint64 {
	var _arg0 *C.GdkDeviceTool // out
	var _cret C.guint64        // in

	_arg0 = (*C.GdkDeviceTool)(unsafe.Pointer(tool.Native()))

	_cret = C.gdk_device_tool_get_serial(_arg0)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// ToolType gets the `GdkDeviceToolType` of the tool.
func (tool *DeviceTool) ToolType() DeviceToolType {
	var _arg0 *C.GdkDeviceTool    // out
	var _cret C.GdkDeviceToolType // in

	_arg0 = (*C.GdkDeviceTool)(unsafe.Pointer(tool.Native()))

	_cret = C.gdk_device_tool_get_tool_type(_arg0)

	var _deviceToolType DeviceToolType // out

	_deviceToolType = DeviceToolType(_cret)

	return _deviceToolType
}
