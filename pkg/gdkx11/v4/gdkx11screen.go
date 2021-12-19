// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_screen_get_type()), F: marshalX11Screener},
	})
}

// X11ScreenOverrider contains methods that are overridable.
type X11ScreenOverrider interface {
}

type X11Screen struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*X11Screen)(nil)
)

func classInitX11Screener(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapX11Screen(obj *externglib.Object) *X11Screen {
	return &X11Screen{
		Object: obj,
	}
}

func marshalX11Screener(p uintptr) (interface{}, error) {
	return wrapX11Screen(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (screen *X11Screen) ConnectWindowManagerChanged(f func()) externglib.SignalHandle {
	return screen.Connect("window-manager-changed", externglib.GeneratedClosure{Func: f})
}

// CurrentDesktop returns the current workspace for screen when running under a
// window manager that supports multiple workspaces, as described in the
// Extended Window Manager Hints (http://www.freedesktop.org/Standards/wm-spec)
// specification.
//
// The function returns the following values:
//
//    - guint32: current workspace, or 0 if workspaces are not supported.
//
func (screen *X11Screen) CurrentDesktop() uint32 {
	var _arg0 *C.GdkX11Screen // out
	var _cret C.guint32       // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_x11_screen_get_current_desktop(_arg0)
	runtime.KeepAlive(screen)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// NumberOfDesktops returns the number of workspaces for screen when running
// under a window manager that supports multiple workspaces, as described in the
// Extended Window Manager Hints (http://www.freedesktop.org/Standards/wm-spec)
// specification.
//
// The function returns the following values:
//
//    - guint32: number of workspaces, or 0 if workspaces are not supported.
//
func (screen *X11Screen) NumberOfDesktops() uint32 {
	var _arg0 *C.GdkX11Screen // out
	var _cret C.guint32       // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_x11_screen_get_number_of_desktops(_arg0)
	runtime.KeepAlive(screen)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// ScreenNumber returns the index of a X11Screen.
//
// The function returns the following values:
//
//    - gint: position of screen among the screens of its display.
//
func (screen *X11Screen) ScreenNumber() int {
	var _arg0 *C.GdkX11Screen // out
	var _cret C.int           // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_x11_screen_get_screen_number(_arg0)
	runtime.KeepAlive(screen)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// WindowManagerName returns the name of the window manager for screen.
//
// The function returns the following values:
//
//    - utf8: name of the window manager screen screen, or "unknown" if the
//      window manager is unknown. The string is owned by GDK and should not be
//      freed.
//
func (screen *X11Screen) WindowManagerName() string {
	var _arg0 *C.GdkX11Screen // out
	var _cret *C.char         // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_x11_screen_get_window_manager_name(_arg0)
	runtime.KeepAlive(screen)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SupportsNetWmHint: this function is specific to the X11 backend of GDK, and
// indicates whether the window manager supports a certain hint from the
// Extended Window Manager Hints (http://www.freedesktop.org/Standards/wm-spec)
// specification.
//
// When using this function, keep in mind that the window manager can change
// over time; so you shouldn’t use this function in a way that impacts
// persistent application state. A common bug is that your application can start
// up before the window manager does when the user logs in, and before the
// window manager starts gdk_x11_screen_supports_net_wm_hint() will return FALSE
// for every property. You can monitor the window_manager_changed signal on
// X11Screen to detect a window manager change.
//
// The function takes the following parameters:
//
//    - propertyName: name of the WM property.
//
// The function returns the following values:
//
//    - ok: TRUE if the window manager supports property.
//
func (screen *X11Screen) SupportsNetWmHint(propertyName string) bool {
	var _arg0 *C.GdkX11Screen // out
	var _arg1 *C.char         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(screen.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(propertyName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_x11_screen_supports_net_wm_hint(_arg0, _arg1)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(propertyName)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
