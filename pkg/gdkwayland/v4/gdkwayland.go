// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"runtime"
	_ "runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4-wayland gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/wayland/gdkwayland.h>
// #include <glib-object.h>
// extern void callbackDelete(gpointer);
// extern void _gotk4_gdkwayland4_WaylandToplevelExported(GdkToplevel*, char*, gpointer);
import "C"

// GType values.
var (
	GTypeWaylandDevice    = coreglib.Type(C.gdk_wayland_device_get_type())
	GTypeWaylandDisplay   = coreglib.Type(C.gdk_wayland_display_get_type())
	GTypeWaylandGLContext = coreglib.Type(C.gdk_wayland_gl_context_get_type())
	GTypeWaylandMonitor   = coreglib.Type(C.gdk_wayland_monitor_get_type())
	GTypeWaylandPopup     = coreglib.Type(C.gdk_wayland_popup_get_type())
	GTypeWaylandSeat      = coreglib.Type(C.gdk_wayland_seat_get_type())
	GTypeWaylandSurface   = coreglib.Type(C.gdk_wayland_surface_get_type())
	GTypeWaylandToplevel  = coreglib.Type(C.gdk_wayland_toplevel_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWaylandDevice, F: marshalWaylandDevice},
		coreglib.TypeMarshaler{T: GTypeWaylandDisplay, F: marshalWaylandDisplay},
		coreglib.TypeMarshaler{T: GTypeWaylandGLContext, F: marshalWaylandGLContext},
		coreglib.TypeMarshaler{T: GTypeWaylandMonitor, F: marshalWaylandMonitor},
		coreglib.TypeMarshaler{T: GTypeWaylandPopup, F: marshalWaylandPopup},
		coreglib.TypeMarshaler{T: GTypeWaylandSeat, F: marshalWaylandSeat},
		coreglib.TypeMarshaler{T: GTypeWaylandSurface, F: marshalWaylandSurface},
		coreglib.TypeMarshaler{T: GTypeWaylandToplevel, F: marshalWaylandToplevel},
	})
}

// WaylandToplevelExported: callback that gets called when the handle for a
// surface has been obtained from the Wayland compositor.
//
// This callback is used in gdkwayland.WaylandToplevel.ExportHandle().
//
// The handle can be passed to other processes, for the purpose of marking
// surfaces as transient for out-of-process surfaces.
type WaylandToplevelExported func(toplevel *WaylandToplevel, handle string)

// WaylandDevice: wayland implementation of GdkDevice.
//
// Beyond the regular gdk.Device API, the Wayland implementation
// provides access to Wayland objects such as the wl_seat with
// gdkwayland.WaylandDevice.GetWlSeat(), the wl_keyboard with
// gdkwayland.WaylandDevice.GetWlKeyboard() and the wl_pointer with
// gdkwayland.WaylandDevice.GetWlPointer().
type WaylandDevice struct {
	_ [0]func() // equal guard
	gdk.Device
}

var (
	_ gdk.Devicer = (*WaylandDevice)(nil)
)

func wrapWaylandDevice(obj *coreglib.Object) *WaylandDevice {
	return &WaylandDevice{
		Device: gdk.Device{
			Object: obj,
		},
	}
}

func marshalWaylandDevice(p uintptr) (interface{}, error) {
	return wrapWaylandDevice(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NodePath returns the /dev/input/event* path of this device.
//
// For GdkDevices that possibly coalesce multiple hardware devices (eg. mouse,
// keyboard, touch,...), this function will return NULL.
//
// This is most notably implemented for devices of type GDK_SOURCE_PEN,
// GDK_SOURCE_TABLET_PAD.
//
// The function returns the following values:
//
//   - utf8 (optional): /dev/input/event* path of this device.
//
func (device *WaylandDevice) NodePath() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.char      // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_cret = C.gdk_wayland_device_get_node_path(_arg0)
	runtime.KeepAlive(device)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// WaylandDisplay: wayland implementation of GdkDisplay.
//
// Beyond the regular gdk.Display API, the Wayland implementation
// provides access to Wayland objects such as the wl_display with
// gdkwayland.WaylandDisplay.GetWlDisplay(), the wl_compositor with
// gdkwayland.WaylandDisplay.GetWlCompositor().
//
// You can find out what Wayland globals are supported by a display with
// gdkwayland.WaylandDisplay.QueryRegistry().
type WaylandDisplay struct {
	_ [0]func() // equal guard
	gdk.Display
}

var (
	_ coreglib.Objector = (*WaylandDisplay)(nil)
)

func wrapWaylandDisplay(obj *coreglib.Object) *WaylandDisplay {
	return &WaylandDisplay{
		Display: gdk.Display{
			Object: obj,
		},
	}
}

func marshalWaylandDisplay(p uintptr) (interface{}, error) {
	return wrapWaylandDisplay(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// StartupNotificationID gets the startup notification ID for a Wayland display,
// or NULL if no ID has been defined.
//
// The function returns the following values:
//
//   - utf8 (optional): startup notification ID for display, or NULL.
//
func (display *WaylandDisplay) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_cret = C.gdk_wayland_display_get_startup_notification_id(_arg0)
	runtime.KeepAlive(display)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// QueryRegistry returns TRUE if the the interface was found in the display
// wl_registry.global handler.
//
// The function takes the following parameters:
//
//   - global interface to query in the registry.
//
// The function returns the following values:
//
//   - ok: TRUE if the global is offered by the compositor.
//
func (display *WaylandDisplay) QueryRegistry(global string) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(global)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_wayland_display_query_registry(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(global)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetCursorTheme sets the cursor theme for the given display.
//
// The function takes the following parameters:
//
//   - name: new cursor theme.
//   - size to use for cursors.
//
func (display *WaylandDisplay) SetCursorTheme(name string, size int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 C.int         // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(size)

	C.gdk_wayland_display_set_cursor_theme(_arg0, _arg1, _arg2)
	runtime.KeepAlive(display)
	runtime.KeepAlive(name)
	runtime.KeepAlive(size)
}

// SetStartupNotificationID sets the startup notification ID for a display.
//
// This is usually taken from the value of the DESKTOP_STARTUP_ID environment
// variable, but in some cases (such as the application not being launched using
// exec()) it can come from other sources.
//
// The startup ID is also what is used to signal that the startup
// is complete (for example, when opening a window or when calling
// gdk.Display.NotifyStartupComplete()).
//
// The function takes the following parameters:
//
//   - startupId: startup notification ID (must be valid utf8).
//
func (display *WaylandDisplay) SetStartupNotificationID(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(startupId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_wayland_display_set_startup_notification_id(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(startupId)
}

// WaylandGLContext: wayland implementation of GdkGLContext.
type WaylandGLContext struct {
	_ [0]func() // equal guard
	gdk.GLContext
}

var (
	_ gdk.GLContexter = (*WaylandGLContext)(nil)
)

func wrapWaylandGLContext(obj *coreglib.Object) *WaylandGLContext {
	return &WaylandGLContext{
		GLContext: gdk.GLContext{
			DrawContext: gdk.DrawContext{
				Object: obj,
			},
		},
	}
}

func marshalWaylandGLContext(p uintptr) (interface{}, error) {
	return wrapWaylandGLContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// WaylandMonitor: wayland implementation of GdkMonitor.
//
// Beyond the gdk.Monitor API, the Wayland implementation offers access to the
// Wayland wl_output object with gdkwayland.WaylandMonitor.GetWlOutput().
type WaylandMonitor struct {
	_ [0]func() // equal guard
	gdk.Monitor
}

var (
	_ coreglib.Objector = (*WaylandMonitor)(nil)
)

func wrapWaylandMonitor(obj *coreglib.Object) *WaylandMonitor {
	return &WaylandMonitor{
		Monitor: gdk.Monitor{
			Object: obj,
		},
	}
}

func marshalWaylandMonitor(p uintptr) (interface{}, error) {
	return wrapWaylandMonitor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// WaylandPopup: wayland implementation of GdkPopup.
type WaylandPopup struct {
	_ [0]func() // equal guard
	WaylandSurface

	*coreglib.Object
	gdk.Popup
	gdk.Surface
}

var (
	_ coreglib.Objector = (*WaylandPopup)(nil)
	_ gdk.Surfacer      = (*WaylandPopup)(nil)
)

func wrapWaylandPopup(obj *coreglib.Object) *WaylandPopup {
	return &WaylandPopup{
		WaylandSurface: WaylandSurface{
			Surface: gdk.Surface{
				Object: obj,
			},
		},
		Object: obj,
		Popup: gdk.Popup{
			Surface: gdk.Surface{
				Object: obj,
			},
		},
		Surface: gdk.Surface{
			Object: obj,
		},
	}
}

func marshalWaylandPopup(p uintptr) (interface{}, error) {
	return wrapWaylandPopup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// WaylandSeat: wayland implementation of GdkSeat.
//
// Beyond the regular gdk.Seat API, the Wayland implementation provides access
// to the Wayland wl_seat object with gdkwayland.WaylandSeat.GetWlSeat().
type WaylandSeat struct {
	_ [0]func() // equal guard
	gdk.Seat
}

var (
	_ gdk.Seater = (*WaylandSeat)(nil)
)

func wrapWaylandSeat(obj *coreglib.Object) *WaylandSeat {
	return &WaylandSeat{
		Seat: gdk.Seat{
			Object: obj,
		},
	}
}

func marshalWaylandSeat(p uintptr) (interface{}, error) {
	return wrapWaylandSeat(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// WaylandSurface: wayland implementation of GdkSurface.
//
// Beyond the gdk.Surface API, the Wayland implementation offers access to the
// Wayland wl_surface object with gdkwayland.WaylandSurface.GetWlSurface().
type WaylandSurface struct {
	_ [0]func() // equal guard
	gdk.Surface
}

var (
	_ gdk.Surfacer = (*WaylandSurface)(nil)
)

func wrapWaylandSurface(obj *coreglib.Object) *WaylandSurface {
	return &WaylandSurface{
		Surface: gdk.Surface{
			Object: obj,
		},
	}
}

func marshalWaylandSurface(p uintptr) (interface{}, error) {
	return wrapWaylandSurface(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// WaylandToplevel: wayland implementation of GdkToplevel.
//
// Beyond the gdk.Toplevel API, the Wayland implementation has
// API to set up cross-process parent-child relationships between
// surfaces with gdkwayland.WaylandToplevel.ExportHandle() and
// gdkwayland.WaylandToplevel.SetTransientForExported().
type WaylandToplevel struct {
	_ [0]func() // equal guard
	WaylandSurface

	*coreglib.Object
	gdk.Surface
	gdk.Toplevel
}

var (
	_ coreglib.Objector = (*WaylandToplevel)(nil)
	_ gdk.Surfacer      = (*WaylandToplevel)(nil)
)

func wrapWaylandToplevel(obj *coreglib.Object) *WaylandToplevel {
	return &WaylandToplevel{
		WaylandSurface: WaylandSurface{
			Surface: gdk.Surface{
				Object: obj,
			},
		},
		Object: obj,
		Surface: gdk.Surface{
			Object: obj,
		},
		Toplevel: gdk.Toplevel{
			Surface: gdk.Surface{
				Object: obj,
			},
		},
	}
}

func marshalWaylandToplevel(p uintptr) (interface{}, error) {
	return wrapWaylandToplevel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ExportHandle: asynchronously obtains a handle for a surface that can be
// passed to other processes.
//
// When the handle has been obtained, callback will be called.
//
// It is an error to call this function on a surface that is already exported.
//
// When the handle is no longer needed,
// gdkwayland.WaylandToplevel.UnexportHandle() should be called to clean up
// resources.
//
// The main purpose for obtaining a handle is to mark a
// surface from another surface as transient for this one, see
// gdkwayland.WaylandToplevel.SetTransientForExported().
//
// Note that this API depends on an unstable Wayland protocol, and thus may
// require changes in the future.
//
// The function takes the following parameters:
//
//   - callback to call with the handle.
//
// The function returns the following values:
//
//   - ok: TRUE if the handle has been requested, FALSE if an error occurred.
//
func (toplevel *WaylandToplevel) ExportHandle(callback WaylandToplevelExported) bool {
	var _arg0 *C.GdkToplevel               // out
	var _arg1 C.GdkWaylandToplevelExported // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify
	var _cret C.gboolean // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(coreglib.InternObject(toplevel).Native()))
	_arg1 = (*[0]byte)(C._gotk4_gdkwayland4_WaylandToplevelExported)
	_arg2 = C.gpointer(gbox.Assign(callback))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	_cret = C.gdk_wayland_toplevel_export_handle(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(toplevel)
	runtime.KeepAlive(callback)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetApplicationID sets the application id on a GdkToplevel.
//
// The function takes the following parameters:
//
//   - applicationId: application id for the toplevel.
//
func (toplevel *WaylandToplevel) SetApplicationID(applicationId string) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(coreglib.InternObject(toplevel).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(applicationId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_wayland_toplevel_set_application_id(_arg0, _arg1)
	runtime.KeepAlive(toplevel)
	runtime.KeepAlive(applicationId)
}

// SetTransientForExported marks toplevel as transient for the surface to which
// the given parent_handle_str refers.
//
// Typically, the handle will originate from a
// gdkwayland.WaylandToplevel.ExportHandle() call in another process.
//
// Note that this API depends on an unstable Wayland protocol, and thus may
// require changes in the future.
//
// The function takes the following parameters:
//
//   - parentHandleStr: exported handle for a surface.
//
// The function returns the following values:
//
//   - ok: TRUE if the surface has been marked as transient, FALSE if an error
//     occurred.
//
func (toplevel *WaylandToplevel) SetTransientForExported(parentHandleStr string) bool {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.char        // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(coreglib.InternObject(toplevel).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(parentHandleStr)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_wayland_toplevel_set_transient_for_exported(_arg0, _arg1)
	runtime.KeepAlive(toplevel)
	runtime.KeepAlive(parentHandleStr)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnexportHandle destroys the handle that was obtained with
// gdk_wayland_toplevel_export_handle().
//
// It is an error to call this function on a surface that does not have a
// handle.
//
// Note that this API depends on an unstable Wayland protocol, and thus may
// require changes in the future.
func (toplevel *WaylandToplevel) UnexportHandle() {
	var _arg0 *C.GdkToplevel // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(coreglib.InternObject(toplevel).Native()))

	C.gdk_wayland_toplevel_unexport_handle(_arg0)
	runtime.KeepAlive(toplevel)
}
