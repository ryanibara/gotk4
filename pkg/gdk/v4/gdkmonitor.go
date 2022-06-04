// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gdk4_Monitor_ConnectInvalidate(gpointer, guintptr);
import "C"

// glib.Type values for gdkmonitor.go.
var (
	GTypeSubpixelLayout = coreglib.Type(C.gdk_subpixel_layout_get_type())
	GTypeMonitor        = coreglib.Type(C.gdk_monitor_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
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
	return SubpixelLayout(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
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

// Monitor: GdkMonitor objects represent the individual outputs that are
// associated with a GdkDisplay.
//
// GdkDisplay keeps a GListModel to enumerate and monitor monitors with
// gdk.Display.GetMonitors(). You can use gdk.Display.GetMonitorAtSurface() to
// find a particular monitor.
type Monitor struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Monitor)(nil)
)

func classInitMonitorrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapMonitor(obj *coreglib.Object) *Monitor {
	return &Monitor{
		Object: obj,
	}
}

func marshalMonitor(p uintptr) (interface{}, error) {
	return wrapMonitor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gdk4_Monitor_ConnectInvalidate
func _gotk4_gdk4_Monitor_ConnectInvalidate(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectInvalidate is emitted when the output represented by monitor gets
// disconnected.
func (monitor *Monitor) ConnectInvalidate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(monitor, "invalidate", false, unsafe.Pointer(C._gotk4_gdk4_Monitor_ConnectInvalidate), f)
}

// Connector gets the name of the monitor's connector, if available.
//
// The function returns the following values:
//
//    - utf8 (optional): name of the connector.
//
func (monitor *Monitor) Connector() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_gret := girepository.MustFind("Gdk", "Monitor").InvokeMethod("get_connector", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Display gets the display that this monitor belongs to.
//
// The function returns the following values:
//
//    - display: display.
//
func (monitor *Monitor) Display() *Display {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_gret := girepository.MustFind("Gdk", "Monitor").InvokeMethod("get_display", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _display *Display // out

	_display = wrapDisplay(coreglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// Geometry retrieves the size and position of the monitor within the display
// coordinate space.
//
// The returned geometry is in ”application pixels”, not in ”device pixels” (see
// gdk.Monitor.GetScaleFactor()).
//
// The function returns the following values:
//
//    - geometry: GdkRectangle to be filled with the monitor geometry.
//
func (monitor *Monitor) Geometry() *Rectangle {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	girepository.MustFind("Gdk", "Monitor").InvokeMethod("get_geometry", _args[:], _outs[:])

	runtime.KeepAlive(monitor)

	var _geometry *Rectangle // out

	_geometry = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))

	return _geometry
}

// HeightMm gets the height in millimeters of the monitor.
//
// The function returns the following values:
//
//    - gint: physical height of the monitor.
//
func (monitor *Monitor) HeightMm() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_gret := girepository.MustFind("Gdk", "Monitor").InvokeMethod("get_height_mm", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// Manufacturer gets the name or PNP ID of the monitor's manufacturer.
//
// Note that this value might also vary depending on actual display backend.
//
// The PNP ID registry is located at https://uefi.org/pnp_id_list
// (https://uefi.org/pnp_id_list).
//
// The function returns the following values:
//
//    - utf8 (optional): name of the manufacturer, or NULL.
//
func (monitor *Monitor) Manufacturer() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_gret := girepository.MustFind("Gdk", "Monitor").InvokeMethod("get_manufacturer", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Model gets the string identifying the monitor model, if available.
//
// The function returns the following values:
//
//    - utf8 (optional): monitor model, or NULL.
//
func (monitor *Monitor) Model() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_gret := girepository.MustFind("Gdk", "Monitor").InvokeMethod("get_model", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
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
func (monitor *Monitor) RefreshRate() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_gret := girepository.MustFind("Gdk", "Monitor").InvokeMethod("get_refresh_rate", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

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
//
// The function returns the following values:
//
//    - gint: scale factor.
//
func (monitor *Monitor) ScaleFactor() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_gret := girepository.MustFind("Gdk", "Monitor").InvokeMethod("get_scale_factor", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// WidthMm gets the width in millimeters of the monitor.
//
// The function returns the following values:
//
//    - gint: physical width of the monitor.
//
func (monitor *Monitor) WidthMm() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_gret := girepository.MustFind("Gdk", "Monitor").InvokeMethod("get_width_mm", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// IsValid returns TRUE if the monitor object corresponds to a physical monitor.
//
// The monitor becomes invalid when the physical monitor is unplugged or
// removed.
//
// The function returns the following values:
//
//    - ok: TRUE if the object corresponds to a physical monitor.
//
func (monitor *Monitor) IsValid() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_gret := girepository.MustFind("Gdk", "Monitor").InvokeMethod("is_valid", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}
