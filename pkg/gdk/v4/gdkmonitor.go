// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
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

// Monitor: gdkMonitor objects represent the individual outputs that are
// associated with a Display. GdkDisplay keeps a Model to enumerate and monitor
// monitors with gdk_display_get_monitors(). You can use
// gdk_display_get_monitor_at_surface() to find a particular monitor.
type Monitor interface {
	gextras.Objector

	// Connector gets the name of the monitor's connector, if available.
	Connector() string
	// Display gets the display that this monitor belongs to.
	Display() Display
	// Geometry retrieves the size and position of an individual monitor within
	// the display coordinate space. The returned geometry is in ”application
	// pixels”, not in ”device pixels” (see gdk_monitor_get_scale_factor()).
	Geometry() Rectangle
	// HeightMm gets the height in millimeters of the monitor.
	HeightMm() int
	// Manufacturer gets the name or PNP ID of the monitor's manufacturer, if
	// available.
	//
	// Note that this value might also vary depending on actual display backend.
	//
	// PNP ID registry is located at https://uefi.org/pnp_id_list
	Manufacturer() string
	// Model gets the string identifying the monitor model, if available.
	Model() string
	// RefreshRate gets the refresh rate of the monitor, if available.
	//
	// The value is in milli-Hertz, so a refresh rate of 60Hz is returned as
	// 60000.
	RefreshRate() int
	// ScaleFactor gets the internal scale factor that maps from monitor
	// coordinates to the actual device pixels. On traditional systems this is
	// 1, but on very high density outputs this can be a higher value (often 2).
	//
	// This can be used if you want to create pixel based data for a particular
	// monitor, but most of the time you’re drawing to a surface where it is
	// better to use gdk_surface_get_scale_factor() instead.
	ScaleFactor() int
	// SubpixelLayout gets information about the layout of red, green and blue
	// primaries for each pixel in this monitor, if available.
	SubpixelLayout() SubpixelLayout
	// WidthMm gets the width in millimeters of the monitor.
	WidthMm() int
	// IsValid returns true if the @monitor object corresponds to a physical
	// monitor. The @monitor becomes invalid when the physical monitor is
	// unplugged or removed.
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

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret *C.char
	var ret1 string

	cret = C.gdk_monitor_get_connector(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Display gets the display that this monitor belongs to.
func (m monitor) Display() Display {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret *C.GdkDisplay
	var ret1 Display

	cret = C.gdk_monitor_get_display(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Display)

	return ret1
}

// Geometry retrieves the size and position of an individual monitor within
// the display coordinate space. The returned geometry is in ”application
// pixels”, not in ”device pixels” (see gdk_monitor_get_scale_factor()).
func (m monitor) Geometry() Rectangle {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var arg1 C.GdkRectangle
	var ret1 *Rectangle

	C.gdk_monitor_get_geometry(arg0, &arg1)

	*ret1 = WrapRectangle(unsafe.Pointer(arg1))

	return ret1
}

// HeightMm gets the height in millimeters of the monitor.
func (m monitor) HeightMm() int {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret C.int
	var ret1 int

	cret = C.gdk_monitor_get_height_mm(arg0)

	ret1 = C.int(cret)

	return ret1
}

// Manufacturer gets the name or PNP ID of the monitor's manufacturer, if
// available.
//
// Note that this value might also vary depending on actual display backend.
//
// PNP ID registry is located at https://uefi.org/pnp_id_list
func (m monitor) Manufacturer() string {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret *C.char
	var ret1 string

	cret = C.gdk_monitor_get_manufacturer(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Model gets the string identifying the monitor model, if available.
func (m monitor) Model() string {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret *C.char
	var ret1 string

	cret = C.gdk_monitor_get_model(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// RefreshRate gets the refresh rate of the monitor, if available.
//
// The value is in milli-Hertz, so a refresh rate of 60Hz is returned as
// 60000.
func (m monitor) RefreshRate() int {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret C.int
	var ret1 int

	cret = C.gdk_monitor_get_refresh_rate(arg0)

	ret1 = C.int(cret)

	return ret1
}

// ScaleFactor gets the internal scale factor that maps from monitor
// coordinates to the actual device pixels. On traditional systems this is
// 1, but on very high density outputs this can be a higher value (often 2).
//
// This can be used if you want to create pixel based data for a particular
// monitor, but most of the time you’re drawing to a surface where it is
// better to use gdk_surface_get_scale_factor() instead.
func (m monitor) ScaleFactor() int {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret C.int
	var ret1 int

	cret = C.gdk_monitor_get_scale_factor(arg0)

	ret1 = C.int(cret)

	return ret1
}

// SubpixelLayout gets information about the layout of red, green and blue
// primaries for each pixel in this monitor, if available.
func (m monitor) SubpixelLayout() SubpixelLayout {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret C.GdkSubpixelLayout
	var ret1 SubpixelLayout

	cret = C.gdk_monitor_get_subpixel_layout(arg0)

	ret1 = SubpixelLayout(cret)

	return ret1
}

// WidthMm gets the width in millimeters of the monitor.
func (m monitor) WidthMm() int {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret C.int
	var ret1 int

	cret = C.gdk_monitor_get_width_mm(arg0)

	ret1 = C.int(cret)

	return ret1
}

// IsValid returns true if the @monitor object corresponds to a physical
// monitor. The @monitor becomes invalid when the physical monitor is
// unplugged or removed.
func (m monitor) IsValid() bool {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gdk_monitor_is_valid(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}
