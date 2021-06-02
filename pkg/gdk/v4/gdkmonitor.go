// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_monitor_get_type()), F: marshalMonitor},
	})
}

// Monitor: `GdkMonitor` objects represent the individual outputs that are
// associated with a `GdkDisplay`.
//
// `GdkDisplay` keeps a `GListModel` to enumerate and monitor monitors with
// [method@Gdk.Display.get_monitors]. You can use
// [method@Gdk.Display.get_monitor_at_surface] to find a particular monitor.
type Monitor interface {
	gextras.Objector

	// Connector gets the name of the monitor's connector, if available.
	Connector() string
	// Display gets the display that this monitor belongs to.
	Display() Display
	// Geometry retrieves the size and position of the monitor within the
	// display coordinate space.
	//
	// The returned geometry is in ”application pixels”, not in ”device pixels”
	// (see [method@Gdk.Monitor.get_scale_factor]).
	Geometry() Rectangle
	// HeightMm gets the height in millimeters of the monitor.
	HeightMm() int
	// Manufacturer gets the name or PNP ID of the monitor's manufacturer.
	//
	// Note that this value might also vary depending on actual display backend.
	//
	// The PNP ID registry is located at https://uefi.org/pnp_id_list
	// (https://uefi.org/pnp_id_list).
	Manufacturer() string
	// Model gets the string identifying the monitor model, if available.
	Model() string
	// RefreshRate gets the refresh rate of the monitor, if available.
	//
	// The value is in milli-Hertz, so a refresh rate of 60Hz is returned as
	// 60000.
	RefreshRate() int
	// ScaleFactor gets the internal scale factor that maps from monitor
	// coordinates to device pixels.
	//
	// On traditional systems this is 1, but on very high density outputs it can
	// be a higher value (often 2).
	//
	// This can be used if you want to create pixel based data for a particular
	// monitor, but most of the time you’re drawing to a surface where it is
	// better to use [method@Gdk.Surface.get_scale_factor] instead.
	ScaleFactor() int
	// SubpixelLayout gets information about the layout of red, green and blue
	// primaries for pixels.
	SubpixelLayout() SubpixelLayout
	// WidthMm gets the width in millimeters of the monitor.
	WidthMm() int
	// IsValid returns true if the @monitor object corresponds to a physical
	// monitor.
	//
	// The @monitor becomes invalid when the physical monitor is unplugged or
	// removed.
	IsValid() bool
}

// monitor implements the Monitor interface.
type monitor struct {
	gextras.Objector
}

var _ Monitor = (*monitor)(nil)

// WrapMonitor wraps a GObject to the right type. It is
// primarily used internally.
func WrapMonitor(obj *externglib.Object) Monitor {
	return Monitor{
		Objector: obj,
	}
}

func marshalMonitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMonitor(obj), nil
}

// Connector gets the name of the monitor's connector, if available.
func (m monitor) Connector() string {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(m.Native())

	ret := C.gdk_monitor_get_connector(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// Display gets the display that this monitor belongs to.
func (m monitor) Display() Display {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(m.Native())

	ret := C.gdk_monitor_get_display(arg0)

	var ret0 Display

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Display)

	return ret0
}

// Geometry retrieves the size and position of the monitor within the
// display coordinate space.
//
// The returned geometry is in ”application pixels”, not in ”device pixels”
// (see [method@Gdk.Monitor.get_scale_factor]).
func (m monitor) Geometry() Rectangle {
	var arg0 *C.GdkMonitor
	var arg1 *C.GdkRectangle // out

	arg0 = (*C.GdkMonitor)(m.Native())

	C.gdk_monitor_get_geometry(arg0, &arg1)

	var ret0 *Rectangle

	{
		ret0 = WrapRectangle(unsafe.Pointer(arg1))
	}

	return ret0
}

// HeightMm gets the height in millimeters of the monitor.
func (m monitor) HeightMm() int {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(m.Native())

	ret := C.gdk_monitor_get_height_mm(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// Manufacturer gets the name or PNP ID of the monitor's manufacturer.
//
// Note that this value might also vary depending on actual display backend.
//
// The PNP ID registry is located at https://uefi.org/pnp_id_list
// (https://uefi.org/pnp_id_list).
func (m monitor) Manufacturer() string {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(m.Native())

	ret := C.gdk_monitor_get_manufacturer(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// Model gets the string identifying the monitor model, if available.
func (m monitor) Model() string {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(m.Native())

	ret := C.gdk_monitor_get_model(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// RefreshRate gets the refresh rate of the monitor, if available.
//
// The value is in milli-Hertz, so a refresh rate of 60Hz is returned as
// 60000.
func (m monitor) RefreshRate() int {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(m.Native())

	ret := C.gdk_monitor_get_refresh_rate(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// ScaleFactor gets the internal scale factor that maps from monitor
// coordinates to device pixels.
//
// On traditional systems this is 1, but on very high density outputs it can
// be a higher value (often 2).
//
// This can be used if you want to create pixel based data for a particular
// monitor, but most of the time you’re drawing to a surface where it is
// better to use [method@Gdk.Surface.get_scale_factor] instead.
func (m monitor) ScaleFactor() int {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(m.Native())

	ret := C.gdk_monitor_get_scale_factor(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// SubpixelLayout gets information about the layout of red, green and blue
// primaries for pixels.
func (m monitor) SubpixelLayout() SubpixelLayout {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(m.Native())

	ret := C.gdk_monitor_get_subpixel_layout(arg0)

	var ret0 SubpixelLayout

	ret0 = SubpixelLayout(ret)

	return ret0
}

// WidthMm gets the width in millimeters of the monitor.
func (m monitor) WidthMm() int {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(m.Native())

	ret := C.gdk_monitor_get_width_mm(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// IsValid returns true if the @monitor object corresponds to a physical
// monitor.
//
// The @monitor becomes invalid when the physical monitor is unplugged or
// removed.
func (m monitor) IsValid() bool {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(m.Native())

	ret := C.gdk_monitor_is_valid(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}
