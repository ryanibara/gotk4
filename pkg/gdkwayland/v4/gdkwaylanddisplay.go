// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-wayland gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/wayland/gdkwayland.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_wayland_display_get_type()), F: marshalWaylandDisplay},
	})
}

// WaylandDisplay: the Wayland implementation of `GdkDisplay`.
//
// Beyond the regular [class@Gdk.Display] API, the Wayland implementation
// provides access to Wayland objects such as the `wl_display` with
// [method@GdkWayland.WaylandDisplay.get_wl_display], the `wl_compositor` with
// [method@GdkWayland.WaylandDisplay.get_wl_compositor].
//
// You can find out what Wayland globals are supported by a display with
// [method@GdkWayland.WaylandDisplay.query_registry].
type WaylandDisplay interface {
	gdk.Display

	// StartupNotificationID gets the startup notification ID for a Wayland
	// display, or nil if no ID has been defined.
	StartupNotificationID() string
	// QueryRegistryWaylandDisplay returns true if the the interface was found
	// in the display `wl_registry.global` handler.
	QueryRegistryWaylandDisplay(global string) bool
	// SetCursorThemeWaylandDisplay sets the cursor theme for the given
	// @display.
	SetCursorThemeWaylandDisplay(name string, size int)
	// SetStartupNotificationIDWaylandDisplay sets the startup notification ID
	// for a display.
	//
	// This is usually taken from the value of the `DESKTOP_STARTUP_ID`
	// environment variable, but in some cases (such as the application not
	// being launched using exec()) it can come from other sources.
	//
	// The startup ID is also what is used to signal that the startup is
	// complete (for example, when opening a window or when calling
	// [method@Gdk.Display.notify_startup_complete]).
	SetStartupNotificationIDWaylandDisplay(startupId string)
}

// waylandDisplay implements the WaylandDisplay class.
type waylandDisplay struct {
	gdk.Display
}

// WrapWaylandDisplay wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandDisplay(obj *externglib.Object) WaylandDisplay {
	return waylandDisplay{
		Display: gdk.WrapDisplay(obj),
	}
}

func marshalWaylandDisplay(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandDisplay(obj), nil
}

func (d waylandDisplay) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_wayland_display_get_startup_notification_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d waylandDisplay) QueryRegistryWaylandDisplay(global string) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(global))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_wayland_display_query_registry(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d waylandDisplay) SetCursorThemeWaylandDisplay(name string, size int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 C.int         // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(size)

	C.gdk_wayland_display_set_cursor_theme(_arg0, _arg1, _arg2)
}

func (d waylandDisplay) SetStartupNotificationIDWaylandDisplay(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(startupId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_wayland_display_set_startup_notification_id(_arg0, _arg1)
}
