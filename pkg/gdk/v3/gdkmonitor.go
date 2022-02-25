// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
// extern void _gotk4_gdk3_Monitor_ConnectInvalidate(gpointer, guintptr);
import "C"

// glib.Type values for gdkmonitor.go.
var (
	GTypeSubpixelLayout = externglib.Type(C.gdk_subpixel_layout_get_type())
	GTypeMonitor        = externglib.Type(C.gdk_monitor_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeSubpixelLayout, F: marshalSubpixelLayout},
		{T: GTypeMonitor, F: marshalMonitor},
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

// MonitorOverrider contains methods that are overridable.
type MonitorOverrider interface {
}

// Monitor objects represent the individual outputs that are associated with a
// Display. GdkDisplay has APIs to enumerate monitors with
// gdk_display_get_n_monitors() and gdk_display_get_monitor(), and to find
// particular monitors with gdk_display_get_primary_monitor() or
// gdk_display_get_monitor_at_window().
//
// GdkMonitor was introduced in GTK+ 3.22 and supersedes earlier APIs in
// GdkScreen to obtain monitor-related information.
type Monitor struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Monitor)(nil)
)

func classInitMonitorrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapMonitor(obj *externglib.Object) *Monitor {
	return &Monitor{
		Object: obj,
	}
}

func marshalMonitor(p uintptr) (interface{}, error) {
	return wrapMonitor(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gdk3_Monitor_ConnectInvalidate
func _gotk4_gdk3_Monitor_ConnectInvalidate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

func (monitor *Monitor) ConnectInvalidate(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(monitor, "invalidate", false, unsafe.Pointer(C._gotk4_gdk3_Monitor_ConnectInvalidate), f)
}

// Display gets the display that this monitor belongs to.
//
// The function returns the following values:
//
//    - display: display.
//
func (monitor *Monitor) Display() *Display {
	var _arg0 *C.GdkMonitor // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(externglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_get_display(_arg0)
	runtime.KeepAlive(monitor)

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// Geometry retrieves the size and position of an individual monitor within the
// display coordinate space. The returned geometry is in ”application pixels”,
// not in ”device pixels” (see gdk_monitor_get_scale_factor()).
//
// The function returns the following values:
//
//    - geometry to be filled with the monitor geometry.
//
func (monitor *Monitor) Geometry() *Rectangle {
	var _arg0 *C.GdkMonitor  // out
	var _arg1 C.GdkRectangle // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(externglib.InternObject(monitor).Native()))

	C.gdk_monitor_get_geometry(_arg0, &_arg1)
	runtime.KeepAlive(monitor)

	var _geometry *Rectangle // out

	_geometry = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _geometry
}

// HeightMm gets the height in millimeters of the monitor.
//
// The function returns the following values:
//
//    - gint: physical height of the monitor.
//
func (monitor *Monitor) HeightMm() int {
	var _arg0 *C.GdkMonitor // out
	var _cret C.int         // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(externglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_get_height_mm(_arg0)
	runtime.KeepAlive(monitor)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Manufacturer gets the name or PNP ID of the monitor's manufacturer, if
// available.
//
// Note that this value might also vary depending on actual display backend.
//
// PNP ID registry is located at https://uefi.org/pnp_id_list.
//
// The function returns the following values:
//
//    - utf8 (optional): name of the manufacturer, or NULL.
//
func (monitor *Monitor) Manufacturer() string {
	var _arg0 *C.GdkMonitor // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(externglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_get_manufacturer(_arg0)
	runtime.KeepAlive(monitor)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Model gets the a string identifying the monitor model, if available.
//
// The function returns the following values:
//
//    - utf8 (optional): monitor model, or NULL.
//
func (monitor *Monitor) Model() string {
	var _arg0 *C.GdkMonitor // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(externglib.InternObject(monitor).Native()))

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
//
// The function returns the following values:
//
//    - gint: refresh rate in milli-Hertz, or 0.
//
func (monitor *Monitor) RefreshRate() int {
	var _arg0 *C.GdkMonitor // out
	var _cret C.int         // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(externglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_get_refresh_rate(_arg0)
	runtime.KeepAlive(monitor)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ScaleFactor gets the internal scale factor that maps from monitor coordinates
// to the actual device pixels. On traditional systems this is 1, but on very
// high density outputs this can be a higher value (often 2).
//
// This can be used if you want to create pixel based data for a particular
// monitor, but most of the time you’re drawing to a window where it is better
// to use gdk_window_get_scale_factor() instead.
//
// The function returns the following values:
//
//    - gint: scale factor.
//
func (monitor *Monitor) ScaleFactor() int {
	var _arg0 *C.GdkMonitor // out
	var _cret C.int         // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(externglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_get_scale_factor(_arg0)
	runtime.KeepAlive(monitor)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SubpixelLayout gets information about the layout of red, green and blue
// primaries for each pixel in this monitor, if available.
//
// The function returns the following values:
//
//    - subpixelLayout: subpixel layout.
//
func (monitor *Monitor) SubpixelLayout() SubpixelLayout {
	var _arg0 *C.GdkMonitor       // out
	var _cret C.GdkSubpixelLayout // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(externglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_get_subpixel_layout(_arg0)
	runtime.KeepAlive(monitor)

	var _subpixelLayout SubpixelLayout // out

	_subpixelLayout = SubpixelLayout(_cret)

	return _subpixelLayout
}

// WidthMm gets the width in millimeters of the monitor.
//
// The function returns the following values:
//
//    - gint: physical width of the monitor.
//
func (monitor *Monitor) WidthMm() int {
	var _arg0 *C.GdkMonitor // out
	var _cret C.int         // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(externglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_get_width_mm(_arg0)
	runtime.KeepAlive(monitor)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Workarea retrieves the size and position of the “work area” on a monitor
// within the display coordinate space. The returned geometry is in ”application
// pixels”, not in ”device pixels” (see gdk_monitor_get_scale_factor()).
//
// The work area should be considered when positioning menus and similar popups,
// to avoid placing them below panels, docks or other desktop components.
//
// Note that not all backends may have a concept of workarea. This function will
// return the monitor geometry if a workarea is not available, or does not
// apply.
//
// The function returns the following values:
//
//    - workarea to be filled with the monitor workarea.
//
func (monitor *Monitor) Workarea() *Rectangle {
	var _arg0 *C.GdkMonitor  // out
	var _arg1 C.GdkRectangle // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(externglib.InternObject(monitor).Native()))

	C.gdk_monitor_get_workarea(_arg0, &_arg1)
	runtime.KeepAlive(monitor)

	var _workarea *Rectangle // out

	_workarea = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _workarea
}

// IsPrimary gets whether this monitor should be considered primary (see
// gdk_display_get_primary_monitor()).
//
// The function returns the following values:
//
//    - ok: TRUE if monitor is primary.
//
func (monitor *Monitor) IsPrimary() bool {
	var _arg0 *C.GdkMonitor // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(externglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_is_primary(_arg0)
	runtime.KeepAlive(monitor)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
