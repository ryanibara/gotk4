// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gdkx11window.go.
var GTypeX11Window = coreglib.Type(C.gdk_x11_window_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeX11Window, F: marshalX11Window},
	})
}

// X11GetServerTime: routine to get the current X server time stamp.
//
// The function takes the following parameters:
//
//    - window used for communication with the server. The window must have
//      GDK_PROPERTY_CHANGE_MASK in its events mask or a hang will result.
//
// The function returns the following values:
//
//    - guint32: time stamp.
//
func X11GetServerTime(window *X11Window) uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_gret := girepository.MustFind("GdkX11", "x11_get_server_time").Invoke(_args[:], nil)
	_cret = *(*C.guint32)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(window)

	var _guint32 uint32 // out

	_guint32 = uint32(*(*C.guint32)(unsafe.Pointer(&_cret)))

	return _guint32
}

// X11WindowOverrider contains methods that are overridable.
type X11WindowOverrider interface {
}

type X11Window struct {
	_ [0]func() // equal guard
	gdk.Window
}

var (
	_ gdk.Windower = (*X11Window)(nil)
)

func classInitX11Windower(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapX11Window(obj *coreglib.Object) *X11Window {
	return &X11Window{
		Window: gdk.Window{
			Object: obj,
		},
	}
}

func marshalX11Window(p uintptr) (interface{}, error) {
	return wrapX11Window(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Desktop gets the number of the workspace window is on.
//
// The function returns the following values:
//
//    - guint32: current workspace of window.
//
func (window *X11Window) Desktop() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_gret := girepository.MustFind("GdkX11", "X11Window").InvokeMethod("get_desktop", _args[:], nil)
	_cret = *(*C.guint32)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(window)

	var _guint32 uint32 // out

	_guint32 = uint32(*(*C.guint32)(unsafe.Pointer(&_cret)))

	return _guint32
}

// MoveToCurrentDesktop moves the window to the correct workspace when running
// under a window manager that supports multiple workspaces, as described in the
// Extended Window Manager Hints (http://www.freedesktop.org/Standards/wm-spec)
// specification. Will not do anything if the window is already on all
// workspaces.
func (window *X11Window) MoveToCurrentDesktop() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	girepository.MustFind("GdkX11", "X11Window").InvokeMethod("move_to_current_desktop", _args[:], nil)

	runtime.KeepAlive(window)
}

// MoveToDesktop moves the window to the given workspace when running unde a
// window manager that supports multiple workspaces, as described in the
// Extended Window Manager Hints (http://www.freedesktop.org/Standards/wm-spec)
// specification.
//
// The function takes the following parameters:
//
//    - desktop: number of the workspace to move the window to.
//
func (window *X11Window) MoveToDesktop(desktop uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	*(*C.guint32)(unsafe.Pointer(&_args[1])) = C.guint32(desktop)

	girepository.MustFind("GdkX11", "X11Window").InvokeMethod("move_to_desktop", _args[:], nil)

	runtime.KeepAlive(window)
	runtime.KeepAlive(desktop)
}

// SetFrameExtents: this is the same as gdk_window_set_shadow_width() but it
// only works on GdkX11Window.
//
// Deprecated: Use gdk_window_set_shadow_width() instead.
//
// The function takes the following parameters:
//
//    - left extent.
//    - right extent.
//    - top extent.
//    - bottom extent.
//
func (window *X11Window) SetFrameExtents(left, right, top, bottom int32) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(left)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(right)
	*(*C.int)(unsafe.Pointer(&_args[3])) = C.int(top)
	*(*C.int)(unsafe.Pointer(&_args[4])) = C.int(bottom)

	girepository.MustFind("GdkX11", "X11Window").InvokeMethod("set_frame_extents", _args[:], nil)

	runtime.KeepAlive(window)
	runtime.KeepAlive(left)
	runtime.KeepAlive(right)
	runtime.KeepAlive(top)
	runtime.KeepAlive(bottom)
}

// SetFrameSyncEnabled: this function can be used to disable frame
// synchronization for a window. Normally frame synchronziation will be enabled
// or disabled based on whether the system has a compositor that supports frame
// synchronization, but if the window is not directly managed by the window
// manager, then frame synchronziation may need to be disabled. This is the case
// for a window embedded via the XEMBED protocol.
//
// The function takes the following parameters:
//
//    - frameSyncEnabled: whether frame-synchronization should be enabled.
//
func (window *X11Window) SetFrameSyncEnabled(frameSyncEnabled bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	if frameSyncEnabled {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("GdkX11", "X11Window").InvokeMethod("set_frame_sync_enabled", _args[:], nil)

	runtime.KeepAlive(window)
	runtime.KeepAlive(frameSyncEnabled)
}

// SetHideTitlebarWhenMaximized: set a hint for the window manager, requesting
// that the titlebar should be hidden when the window is maximized.
//
// Note that this property is automatically updated by GTK+, so this function
// should only be used by applications which do not use GTK+ to create toplevel
// windows.
//
// The function takes the following parameters:
//
//    - hideTitlebarWhenMaximized: whether to hide the titlebar when maximized.
//
func (window *X11Window) SetHideTitlebarWhenMaximized(hideTitlebarWhenMaximized bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	if hideTitlebarWhenMaximized {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("GdkX11", "X11Window").InvokeMethod("set_hide_titlebar_when_maximized", _args[:], nil)

	runtime.KeepAlive(window)
	runtime.KeepAlive(hideTitlebarWhenMaximized)
}

// SetThemeVariant: GTK+ applications can request a dark theme variant. In order
// to make other applications - namely window managers using GTK+ for themeing -
// aware of this choice, GTK+ uses this function to export the requested theme
// variant as _GTK_THEME_VARIANT property on toplevel windows.
//
// Note that this property is automatically updated by GTK+, so this function
// should only be used by applications which do not use GTK+ to create toplevel
// windows.
//
// The function takes the following parameters:
//
//    - variant: theme variant to export.
//
func (window *X11Window) SetThemeVariant(variant string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(variant)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("GdkX11", "X11Window").InvokeMethod("set_theme_variant", _args[:], nil)

	runtime.KeepAlive(window)
	runtime.KeepAlive(variant)
}

// SetUserTime: application can use this call to update the _NET_WM_USER_TIME
// property on a toplevel window. This property stores an Xserver time which
// represents the time of the last user input event received for this window.
// This property may be used by the window manager to alter the focus, stacking,
// and/or placement behavior of windows when they are mapped depending on
// whether the new window was created by a user action or is a "pop-up" window
// activated by a timer or some other event.
//
// Note that this property is automatically updated by GDK, so this function
// should only be used by applications which handle input events bypassing GDK.
//
// The function takes the following parameters:
//
//    - timestamp: XServer timestamp to which the property should be set.
//
func (window *X11Window) SetUserTime(timestamp uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	*(*C.guint32)(unsafe.Pointer(&_args[1])) = C.guint32(timestamp)

	girepository.MustFind("GdkX11", "X11Window").InvokeMethod("set_user_time", _args[:], nil)

	runtime.KeepAlive(window)
	runtime.KeepAlive(timestamp)
}

// SetUTF8Property: this function modifies or removes an arbitrary X11 window
// property of type UTF8_STRING. If the given window is not a toplevel window,
// it is ignored.
//
// The function takes the following parameters:
//
//    - name: property name, will be interned as an X atom.
//    - value (optional): property value, or NULL to delete.
//
func (window *X11Window) SetUTF8Property(name, value string) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))
	if value != "" {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_args[2]))
	}

	girepository.MustFind("GdkX11", "X11Window").InvokeMethod("set_utf8_property", _args[:], nil)

	runtime.KeepAlive(window)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
}
