// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-wayland gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/wayland/gdkwayland.h>
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
	// WlCompositor returns the Wayland `wl_compositor` of a `GdkDisplay`.
	WlCompositor() *interface{}
	// WlDisplay returns the Wayland `wl_display` of a `GdkDisplay`.
	WlDisplay() *interface{}
	// QueryRegistry returns true if the the interface was found in the display
	// `wl_registry.global` handler.
	QueryRegistry(global string) bool
	// SetCursorTheme sets the cursor theme for the given @display.
	SetCursorTheme(name string, size int)
	// SetStartupNotificationID sets the startup notification ID for a display.
	//
	// This is usually taken from the value of the `DESKTOP_STARTUP_ID`
	// environment variable, but in some cases (such as the application not
	// being launched using exec()) it can come from other sources.
	//
	// The startup ID is also what is used to signal that the startup is
	// complete (for example, when opening a window or when calling
	// [method@Gdk.Display.notify_startup_complete]).
	SetStartupNotificationID(startupId string)
}

// waylandDisplay implements the WaylandDisplay interface.
type waylandDisplay struct {
	gdk.Display
}

var _ WaylandDisplay = (*waylandDisplay)(nil)

// WrapWaylandDisplay wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandDisplay(obj *externglib.Object) WaylandDisplay {
	return WaylandDisplay{
		gdk.Display: gdk.WrapDisplay(obj),
	}
}

func marshalWaylandDisplay(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandDisplay(obj), nil
}

// StartupNotificationID gets the startup notification ID for a Wayland
// display, or nil if no ID has been defined.
func (d waylandDisplay) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	var _cret *C.char

	_cret = C.gdk_wayland_display_get_startup_notification_id(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// WlCompositor returns the Wayland `wl_compositor` of a `GdkDisplay`.
func (d waylandDisplay) WlCompositor() *interface{} {
	var _arg0 *C.GdkDisplay

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	var _cret *C.wl_compositor

	_cret = C.gdk_wayland_display_get_wl_compositor(_arg0)

	var _gpointer *interface{}

	_gpointer = (*interface{})(_cret)

	return _gpointer
}

// WlDisplay returns the Wayland `wl_display` of a `GdkDisplay`.
func (d waylandDisplay) WlDisplay() *interface{} {
	var _arg0 *C.GdkDisplay

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	var _cret *C.wl_display

	_cret = C.gdk_wayland_display_get_wl_display(_arg0)

	var _gpointer *interface{}

	_gpointer = (*interface{})(_cret)

	return _gpointer
}

// QueryRegistry returns true if the the interface was found in the display
// `wl_registry.global` handler.
func (d waylandDisplay) QueryRegistry(global string) bool {
	var _arg0 *C.GdkDisplay
	var _arg1 *C.char

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(global))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean

	_cret = C.gdk_wayland_display_query_registry(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetCursorTheme sets the cursor theme for the given @display.
func (d waylandDisplay) SetCursorTheme(name string, size int) {
	var _arg0 *C.GdkDisplay
	var _arg1 *C.char
	var _arg2 C.int

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(size)

	C.gdk_wayland_display_set_cursor_theme(_arg0, _arg1, _arg2)
}

// SetStartupNotificationID sets the startup notification ID for a display.
//
// This is usually taken from the value of the `DESKTOP_STARTUP_ID`
// environment variable, but in some cases (such as the application not
// being launched using exec()) it can come from other sources.
//
// The startup ID is also what is used to signal that the startup is
// complete (for example, when opening a window or when calling
// [method@Gdk.Display.notify_startup_complete]).
func (d waylandDisplay) SetStartupNotificationID(startupId string) {
	var _arg0 *C.GdkDisplay
	var _arg1 *C.char

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(startupId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_wayland_display_set_startup_notification_id(_arg0, _arg1)
}
