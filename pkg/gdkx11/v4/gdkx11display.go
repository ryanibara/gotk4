// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"reflect"
	"runtime"
	"runtime/cgo"
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_display_get_type()), F: marshalX11Displayer},
	})
}

// X11SetSmClientID sets the SM_CLIENT_ID property on the application’s leader
// window so that the window manager can save the application’s state using the
// X11R6 ICCCM session management protocol.
//
// See the X Session Management Library documentation for more information on
// session management and the Inter-Client Communication Conventions Manual.
//
// The function takes the following parameters:
//
//    - smClientId (optional): client id assigned by the session manager when the
//      connection was opened, or NULL to remove the property.
//
func X11SetSmClientID(smClientId string) {
	var _arg1 *C.char // out

	if smClientId != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(smClientId)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gdk_x11_set_sm_client_id(_arg1)
	runtime.KeepAlive(smClientId)
}

type X11Display struct {
	_ [0]func() // equal guard
	gdk.Display
}

var (
	_ externglib.Objector = (*X11Display)(nil)
)

func wrapX11Display(obj *externglib.Object) *X11Display {
	return &X11Display{
		Display: gdk.Display{
			Object: obj,
		},
	}
}

func marshalX11Displayer(p uintptr) (interface{}, error) {
	return wrapX11Display(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectXevent signal is a low level signal that is emitted whenever an XEvent
// has been received.
//
// When handlers to this signal return TRUE, no other handlers will be invoked.
// In particular, the default handler for this function is GDK's own event
// handling mechanism, so by returning TRUE for an event that GDK expects to
// translate, you may break GDK and/or GTK+ in interesting ways. You have been
// warned.
//
// If you want this signal handler to queue a Event, you can use
// gdk_display_put_event().
//
// If you are interested in X GenericEvents, bear in mind that XGetEventData()
// has been already called on the event, and XFreeEventData() will be called
// afterwards.
func (display *X11Display) ConnectXevent(f func(xevent cgo.Handle) bool) externglib.SignalHandle {
	return display.Connect("xevent", f)
}

// ErrorTrapPop pops the error trap pushed by gdk_x11_display_error_trap_push().
// Will XSync() if necessary and will always block until the error is known to
// have occurred or not occurred, so the error code can be returned.
//
// If you don’t need to use the return value,
// gdk_x11_display_error_trap_pop_ignored() would be more efficient.
//
// The function returns the following values:
//
//    - gint: x error code or 0 on success.
//
func (display *X11Display) ErrorTrapPop() int {
	var _arg0 *C.GdkDisplay // out
	var _cret C.int         // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_x11_display_error_trap_pop(_arg0)
	runtime.KeepAlive(display)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ErrorTrapPopIgnored pops the error trap pushed by
// gdk_x11_display_error_trap_push(). Does not block to see if an error
// occurred; merely records the range of requests to ignore errors for, and
// ignores those errors if they arrive asynchronously.
func (display *X11Display) ErrorTrapPopIgnored() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_x11_display_error_trap_pop_ignored(_arg0)
	runtime.KeepAlive(display)
}

// ErrorTrapPush begins a range of X requests on display for which X error
// events will be ignored. Unignored errors (when no trap is pushed) will abort
// the application. Use gdk_x11_display_error_trap_pop() or
// gdk_x11_display_error_trap_pop_ignored()to lift a trap pushed with this
// function.
func (display *X11Display) ErrorTrapPush() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_x11_display_error_trap_push(_arg0)
	runtime.KeepAlive(display)
}

// DefaultGroup returns the default group leader surface for all toplevel
// surfaces on display. This surface is implicitly created by GDK. See
// gdk_x11_surface_set_group().
//
// The function returns the following values:
//
//    - surface: default group leader surface for display.
//
func (display *X11Display) DefaultGroup() gdk.Surfacer {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_x11_display_get_default_group(_arg0)
	runtime.KeepAlive(display)

	var _surface gdk.Surfacer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Surfacer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(gdk.Surfacer)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gdk.Surfacer")
		}
		_surface = rv
	}

	return _surface
}

// PrimaryMonitor gets the primary monitor for the display.
//
// The primary monitor is considered the monitor where the “main desktop” lives.
// While normal application surfaces typically allow the window manager to place
// the surfaces, specialized desktop applications such as panels should place
// themselves on the primary monitor.
//
// If no monitor is the designated primary monitor, any monitor (usually the
// first) may be returned.
//
// The function returns the following values:
//
//    - monitor: primary monitor, or any monitor if no primary monitor is
//      configured by the user.
//
func (display *X11Display) PrimaryMonitor() *gdk.Monitor {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_x11_display_get_primary_monitor(_arg0)
	runtime.KeepAlive(display)

	var _monitor *gdk.Monitor // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_monitor = &gdk.Monitor{
			Object: obj,
		}
	}

	return _monitor
}

// Screen retrieves the X11Screen of the display.
//
// The function returns the following values:
//
//    - x11Screen: X11Screen.
//
func (display *X11Display) Screen() *X11Screen {
	var _arg0 *C.GdkDisplay   // out
	var _cret *C.GdkX11Screen // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_x11_display_get_screen(_arg0)
	runtime.KeepAlive(display)

	var _x11Screen *X11Screen // out

	_x11Screen = wrapX11Screen(externglib.Take(unsafe.Pointer(_cret)))

	return _x11Screen
}

// StartupNotificationID gets the startup notification ID for a display.
//
// The function returns the following values:
//
//    - utf8: startup notification ID for display.
//
func (display *X11Display) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_x11_display_get_startup_notification_id(_arg0)
	runtime.KeepAlive(display)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UserTime returns the timestamp of the last user interaction on display. The
// timestamp is taken from events caused by user interaction such as key presses
// or pointer movements. See gdk_x11_surface_set_user_time().
//
// The function returns the following values:
//
//    - guint32: timestamp of the last user interaction.
//
func (display *X11Display) UserTime() uint32 {
	var _arg0 *C.GdkDisplay // out
	var _cret C.guint32     // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_x11_display_get_user_time(_arg0)
	runtime.KeepAlive(display)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// Grab: call XGrabServer() on display. To ungrab the display again, use
// gdk_x11_display_ungrab().
//
// gdk_x11_display_grab()/gdk_x11_display_ungrab() calls can be nested.
func (display *X11Display) Grab() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_x11_display_grab(_arg0)
	runtime.KeepAlive(display)
}

// SetCursorTheme sets the cursor theme from which the images for cursor should
// be taken.
//
// If the windowing system supports it, existing cursors created with
// gdk_cursor_new_from_name() are updated to reflect the theme change. Custom
// cursors constructed with gdk_cursor_new_from_texture() will have to be
// handled by the application (GTK applications can learn about cursor theme
// changes by listening for change notification for the corresponding Setting).
//
// The function takes the following parameters:
//
//    - theme (optional): name of the cursor theme to use, or NULL to unset a
//      previously set value.
//    - size: cursor size to use, or 0 to keep the previous size.
//
func (display *X11Display) SetCursorTheme(theme string, size int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 C.int         // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	if theme != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(theme)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = C.int(size)

	C.gdk_x11_display_set_cursor_theme(_arg0, _arg1, _arg2)
	runtime.KeepAlive(display)
	runtime.KeepAlive(theme)
	runtime.KeepAlive(size)
}

// SetStartupNotificationID sets the startup notification ID for a display.
//
// This is usually taken from the value of the DESKTOP_STARTUP_ID environment
// variable, but in some cases (such as the application not being launched using
// exec()) it can come from other sources.
//
// If the ID contains the string "_TIME" then the portion following that string
// is taken to be the X11 timestamp of the event that triggered the application
// to be launched and the GDK current event time is set accordingly.
//
// The startup ID is also what is used to signal that the startup is complete
// (for example, when opening a window or when calling
// gdk_display_notify_startup_complete()).
//
// The function takes the following parameters:
//
//    - startupId: startup notification ID (must be valid utf8).
//
func (display *X11Display) SetStartupNotificationID(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(startupId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_x11_display_set_startup_notification_id(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(startupId)
}

// SetSurfaceScale forces a specific window scale for all windows on this
// display, instead of using the default or user configured scale. This is can
// be used to disable scaling support by setting scale to 1, or to
// programmatically set the window scale.
//
// Once the scale is set by this call it will not change in response to later
// user configuration changes.
//
// The function takes the following parameters:
//
//    - scale: new scale value.
//
func (display *X11Display) SetSurfaceScale(scale int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.int         // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = C.int(scale)

	C.gdk_x11_display_set_surface_scale(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(scale)
}

// Ungrab display after it has been grabbed with gdk_x11_display_grab().
func (display *X11Display) Ungrab() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_x11_display_ungrab(_arg0)
	runtime.KeepAlive(display)
}

// X11DisplayOpen tries to open a new display to the X server given by
// display_name. If opening the display fails, NULL is returned.
//
// The function takes the following parameters:
//
//    - displayName (optional): name of the X display. See the XOpenDisplay() for
//      details.
//
// The function returns the following values:
//
//    - display (optional): new display or NULL on error.
//
func X11DisplayOpen(displayName string) *gdk.Display {
	var _arg1 *C.char       // out
	var _cret *C.GdkDisplay // in

	if displayName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(displayName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gdk_x11_display_open(_arg1)
	runtime.KeepAlive(displayName)

	var _display *gdk.Display // out

	if _cret != nil {
		{
			obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
			_display = &gdk.Display{
				Object: obj,
			}
		}
	}

	return _display
}

// X11DisplaySetProgramClass sets the program class.
//
// The X11 backend uses the program class to set the class name part of the
// WM_CLASS property on toplevel windows; see the ICCCM.
//
// The function takes the following parameters:
//
//    - display: Display.
//    - programClass: string.
//
func X11DisplaySetProgramClass(display *gdk.Display, programClass string) {
	var _arg1 *C.GdkDisplay // out
	var _arg2 *C.char       // out

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(programClass)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gdk_x11_display_set_program_class(_arg1, _arg2)
	runtime.KeepAlive(display)
	runtime.KeepAlive(programClass)
}
