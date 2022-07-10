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
// #include <glib-object.h>
// extern void _gotk4_gdk3_Monitor_ConnectInvalidate(gpointer, guintptr);
import "C"

// GTypeSubpixelLayout returns the GType for the type SubpixelLayout.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeSubpixelLayout() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "SubpixelLayout").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalSubpixelLayout)
	return gtype
}

// GTypeMonitor returns the GType for the type Monitor.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeMonitor() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "Monitor").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalMonitor)
	return gtype
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
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Monitor)(nil)
)

func wrapMonitor(obj *coreglib.Object) *Monitor {
	return &Monitor{
		Object: obj,
	}
}

func marshalMonitor(p uintptr) (interface{}, error) {
	return wrapMonitor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gdk3_Monitor_ConnectInvalidate
func _gotk4_gdk3_Monitor_ConnectInvalidate(arg0 C.gpointer, arg1 C.guintptr) {
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

func (monitor *Monitor) ConnectInvalidate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(monitor, "invalidate", false, unsafe.Pointer(C._gotk4_gdk3_Monitor_ConnectInvalidate), f)
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

	_info := girepository.MustFind("Gdk", "Monitor")
	_gret := _info.InvokeClassMethod("get_display", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _display *Display // out

	_display = wrapDisplay(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

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
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_info := girepository.MustFind("Gdk", "Monitor")
	_info.InvokeClassMethod("get_geometry", _args[:], _outs[:])

	runtime.KeepAlive(monitor)

	var _geometry *Rectangle // out

	_geometry = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[0])))))

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

	_info := girepository.MustFind("Gdk", "Monitor")
	_gret := _info.InvokeClassMethod("get_height_mm", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_info := girepository.MustFind("Gdk", "Monitor")
	_gret := _info.InvokeClassMethod("get_manufacturer", _args[:], nil)
	_cret := *(**C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _utf8 string // out

	if *(**C.char)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_cret)))))
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_info := girepository.MustFind("Gdk", "Monitor")
	_gret := _info.InvokeClassMethod("get_model", _args[:], nil)
	_cret := *(**C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _utf8 string // out

	if *(**C.char)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_cret)))))
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

	_info := girepository.MustFind("Gdk", "Monitor")
	_gret := _info.InvokeClassMethod("get_refresh_rate", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

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
func (monitor *Monitor) ScaleFactor() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_info := girepository.MustFind("Gdk", "Monitor")
	_gret := _info.InvokeClassMethod("get_scale_factor", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

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

	_info := girepository.MustFind("Gdk", "Monitor")
	_gret := _info.InvokeClassMethod("get_width_mm", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

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
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_info := girepository.MustFind("Gdk", "Monitor")
	_info.InvokeClassMethod("get_workarea", _args[:], _outs[:])

	runtime.KeepAlive(monitor)

	var _workarea *Rectangle // out

	_workarea = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[0])))))

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_info := girepository.MustFind("Gdk", "Monitor")
	_gret := _info.InvokeClassMethod("is_primary", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(monitor)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}
