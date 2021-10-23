// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_subpixel_layout_get_type()), F: marshalSubpixelLayout},
		{T: externglib.Type(C.gdk_monitor_get_type()), F: marshalMonitorrer},
	})
}

// SubpixelLayout: this enumeration describes how the red, green and blue
// components of physical pixels on an output device are laid out.
type SubpixelLayout C.gint

const (
	// SubpixelLayoutUnknown: layout is not known.
	SubpixelLayoutUnknown SubpixelLayout = iota
	// SubpixelLayoutNone: not organized in this way.
	SubpixelLayoutNone
	// SubpixelLayoutHorizontalRGB: layout is horizontal, the order is RGB.
	SubpixelLayoutHorizontalRGB
	// SubpixelLayoutHorizontalBGR: layout is horizontal, the order is BGR.
	SubpixelLayoutHorizontalBGR
	// SubpixelLayoutVerticalRGB: layout is vertical, the order is RGB.
	SubpixelLayoutVerticalRGB
	// SubpixelLayoutVerticalBGR: layout is vertical, the order is BGR.
	SubpixelLayoutVerticalBGR
)

func marshalSubpixelLayout(p uintptr) (interface{}, error) {
	return SubpixelLayout(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SubpixelLayout.
func (s SubpixelLayout) String() string {
	switch s {
	case SubpixelLayoutUnknown:
		return "Unknown"
	case SubpixelLayoutNone:
		return "None"
	case SubpixelLayoutHorizontalRGB:
		return "HorizontalRGB"
	case SubpixelLayoutHorizontalBGR:
		return "HorizontalBGR"
	case SubpixelLayoutVerticalRGB:
		return "VerticalRGB"
	case SubpixelLayoutVerticalBGR:
		return "VerticalBGR"
	default:
		return fmt.Sprintf("SubpixelLayout(%d)", s)
	}
}

// Monitor: GdkMonitor objects represent the individual outputs that are
// associated with a GdkDisplay.
//
// GdkDisplay keeps a GListModel to enumerate and monitor monitors with
// gdk.Display.GetMonitors(). You can use gdk.Display.GetMonitorAtSurface() to
// find a particular monitor.
type Monitor struct {
	*externglib.Object
}

func wrapMonitor(obj *externglib.Object) *Monitor {
	return &Monitor{
		Object: obj,
	}
}

func marshalMonitorrer(p uintptr) (interface{}, error) {
	return wrapMonitor(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Connector gets the name of the monitor's connector, if available.
func (monitor *Monitor) Connector() string {
	var _arg0 *C.GdkMonitor // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.gdk_monitor_get_connector(_arg0)
	runtime.KeepAlive(monitor)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Display gets the display that this monitor belongs to.
func (monitor *Monitor) Display() *Display {
	var _arg0 *C.GdkMonitor // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.gdk_monitor_get_display(_arg0)
	runtime.KeepAlive(monitor)

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// Geometry retrieves the size and position of the monitor within the display
// coordinate space.
//
// The returned geometry is in ”application pixels”, not in ”device pixels” (see
// gdk.Monitor.GetScaleFactor()).
func (monitor *Monitor) Geometry() *Rectangle {
	var _arg0 *C.GdkMonitor  // out
	var _arg1 C.GdkRectangle // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	C.gdk_monitor_get_geometry(_arg0, &_arg1)
	runtime.KeepAlive(monitor)

	var _geometry *Rectangle // out

	_geometry = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _geometry
}

// HeightMm gets the height in millimeters of the monitor.
func (monitor *Monitor) HeightMm() int {
	var _arg0 *C.GdkMonitor // out
	var _cret C.int         // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.gdk_monitor_get_height_mm(_arg0)
	runtime.KeepAlive(monitor)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Manufacturer gets the name or PNP ID of the monitor's manufacturer.
//
// Note that this value might also vary depending on actual display backend.
//
// The PNP ID registry is located at https://uefi.org/pnp_id_list
// (https://uefi.org/pnp_id_list).
func (monitor *Monitor) Manufacturer() string {
	var _arg0 *C.GdkMonitor // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.gdk_monitor_get_manufacturer(_arg0)
	runtime.KeepAlive(monitor)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Model gets the string identifying the monitor model, if available.
func (monitor *Monitor) Model() string {
	var _arg0 *C.GdkMonitor // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.gdk_monitor_get_model(_arg0)
	runtime.KeepAlive(monitor)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// RefreshRate gets the refresh rate of the monitor, if available.
//
// The value is in milli-Hertz, so a refresh rate of 60Hz is returned as 60000.
func (monitor *Monitor) RefreshRate() int {
	var _arg0 *C.GdkMonitor // out
	var _cret C.int         // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.gdk_monitor_get_refresh_rate(_arg0)
	runtime.KeepAlive(monitor)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ScaleFactor gets the internal scale factor that maps from monitor coordinates
// to device pixels.
//
// On traditional systems this is 1, but on very high density outputs it can be
// a higher value (often 2).
//
// This can be used if you want to create pixel based data for a particular
// monitor, but most of the time you’re drawing to a surface where it is better
// to use gdk.Surface.GetScaleFactor() instead.
func (monitor *Monitor) ScaleFactor() int {
	var _arg0 *C.GdkMonitor // out
	var _cret C.int         // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.gdk_monitor_get_scale_factor(_arg0)
	runtime.KeepAlive(monitor)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SubpixelLayout gets information about the layout of red, green and blue
// primaries for pixels.
func (monitor *Monitor) SubpixelLayout() SubpixelLayout {
	var _arg0 *C.GdkMonitor       // out
	var _cret C.GdkSubpixelLayout // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.gdk_monitor_get_subpixel_layout(_arg0)
	runtime.KeepAlive(monitor)

	var _subpixelLayout SubpixelLayout // out

	_subpixelLayout = SubpixelLayout(_cret)

	return _subpixelLayout
}

// WidthMm gets the width in millimeters of the monitor.
func (monitor *Monitor) WidthMm() int {
	var _arg0 *C.GdkMonitor // out
	var _cret C.int         // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.gdk_monitor_get_width_mm(_arg0)
	runtime.KeepAlive(monitor)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IsValid returns TRUE if the monitor object corresponds to a physical monitor.
//
// The monitor becomes invalid when the physical monitor is unplugged or
// removed.
func (monitor *Monitor) IsValid() bool {
	var _arg0 *C.GdkMonitor // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.gdk_monitor_is_valid(_arg0)
	runtime.KeepAlive(monitor)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ConnectInvalidate: emitted when the output represented by monitor gets
// disconnected.
func (monitor *Monitor) ConnectInvalidate(f func()) externglib.SignalHandle {
	return monitor.Connect("invalidate", f)
}
