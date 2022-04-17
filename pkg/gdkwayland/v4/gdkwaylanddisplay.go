// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <gdk/wayland/gdkwayland.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gdkwaylanddisplay.go.
var GTypeWaylandDisplay = externglib.Type(C.gdk_wayland_display_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeWaylandDisplay, F: marshalWaylandDisplay},
	})
}

// WaylandDisplayOverrider contains methods that are overridable.
type WaylandDisplayOverrider interface {
	externglib.Objector
}

// WrapWaylandDisplayOverrider wraps the WaylandDisplayOverrider
// interface implementation to access the instance methods.
func WrapWaylandDisplayOverrider(obj WaylandDisplayOverrider) *WaylandDisplay {
	return wrapWaylandDisplay(externglib.BaseObject(obj))
}

// WaylandDisplay: wayland implementation of GdkDisplay.
//
// Beyond the regular gdk.Display API, the Wayland implementation provides
// access to Wayland objects such as the wl_display with
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
	_ externglib.Objector = (*WaylandDisplay)(nil)
)

func classInitWaylandDisplayer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapWaylandDisplay(obj *externglib.Object) *WaylandDisplay {
	return &WaylandDisplay{
		Display: gdk.Display{
			Object: obj,
		},
	}
}

func marshalWaylandDisplay(p uintptr) (interface{}, error) {
	return wrapWaylandDisplay(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// StartupNotificationID gets the startup notification ID for a Wayland display,
// or NULL if no ID has been defined.
//
// The function returns the following values:
//
//    - utf8 (optional): startup notification ID for display, or NULL.
//
func (display *WaylandDisplay) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

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
//    - global interface to query in the registry.
//
// The function returns the following values:
//
//    - ok: TRUE if the global is offered by the compositor.
//
func (display *WaylandDisplay) QueryRegistry(global string) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
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
//    - name: new cursor theme.
//    - size to use for cursors.
//
func (display *WaylandDisplay) SetCursorTheme(name string, size int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 C.int         // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
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
// The startup ID is also what is used to signal that the startup is complete
// (for example, when opening a window or when calling
// gdk.Display.NotifyStartupComplete()).
//
// The function takes the following parameters:
//
//    - startupId: startup notification ID (must be valid utf8).
//
func (display *WaylandDisplay) SetStartupNotificationID(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(startupId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_wayland_display_set_startup_notification_id(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(startupId)
}
