// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_display_get_type()), F: marshalX11Display},
	})
}

// X11SetSmClientID sets the `SM_CLIENT_ID` property on the application’s leader
// window so that the window manager can save the application’s state using the
// X11R6 ICCCM session management protocol.
//
// See the X Session Management Library documentation for more information on
// session management and the Inter-Client Communication Conventions Manual
func X11SetSmClientID(smClientId string) {
	var _arg1 *C.char // out

	_arg1 = (*C.char)(C.CString(smClientId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_x11_set_sm_client_id(_arg1)
}

type X11Display interface {
	gdk.Display

	// ErrorTrapPopX11Display pops the error trap pushed by
	// gdk_x11_display_error_trap_push(). Will XSync() if necessary and will
	// always block until the error is known to have occurred or not occurred,
	// so the error code can be returned.
	//
	// If you don’t need to use the return value,
	// gdk_x11_display_error_trap_pop_ignored() would be more efficient.
	ErrorTrapPopX11Display() int
	// ErrorTrapPopIgnoredX11Display pops the error trap pushed by
	// gdk_x11_display_error_trap_push(). Does not block to see if an error
	// occurred; merely records the range of requests to ignore errors for, and
	// ignores those errors if they arrive asynchronously.
	ErrorTrapPopIgnoredX11Display()
	// ErrorTrapPushX11Display begins a range of X requests on @display for
	// which X error events will be ignored. Unignored errors (when no trap is
	// pushed) will abort the application. Use gdk_x11_display_error_trap_pop()
	// or gdk_x11_display_error_trap_pop_ignored()to lift a trap pushed with
	// this function.
	ErrorTrapPushX11Display()
	// DefaultGroup returns the default group leader surface for all toplevel
	// surfaces on @display. This surface is implicitly created by GDK. See
	// gdk_x11_surface_set_group().
	DefaultGroup() gdk.Surface
	// GlxVersion retrieves the version of the GLX implementation.
	GlxVersion() (major int, minor int, ok bool)
	// PrimaryMonitor gets the primary monitor for the display.
	//
	// The primary monitor is considered the monitor where the “main desktop”
	// lives. While normal application surfaces typically allow the window
	// manager to place the surfaces, specialized desktop applications such as
	// panels should place themselves on the primary monitor.
	//
	// If no monitor is the designated primary monitor, any monitor (usually the
	// first) may be returned.
	PrimaryMonitor() gdk.Monitor
	// Screen retrieves the X11Screen of the @display.
	Screen() X11Screen
	// StartupNotificationID gets the startup notification ID for a display.
	StartupNotificationID() string
	// UserTime returns the timestamp of the last user interaction on @display.
	// The timestamp is taken from events caused by user interaction such as key
	// presses or pointer movements. See gdk_x11_surface_set_user_time().
	UserTime() uint32
	// GrabX11Display: call XGrabServer() on @display. To ungrab the display
	// again, use gdk_x11_display_ungrab().
	//
	// gdk_x11_display_grab()/gdk_x11_display_ungrab() calls can be nested.
	GrabX11Display()
	// SetCursorThemeX11Display sets the cursor theme from which the images for
	// cursor should be taken.
	//
	// If the windowing system supports it, existing cursors created with
	// gdk_cursor_new_from_name() are updated to reflect the theme change.
	// Custom cursors constructed with gdk_cursor_new_from_texture() will have
	// to be handled by the application (GTK applications can learn about cursor
	// theme changes by listening for change notification for the corresponding
	// Setting).
	SetCursorThemeX11Display(theme string, size int)
	// SetStartupNotificationIDX11Display sets the startup notification ID for a
	// display.
	//
	// This is usually taken from the value of the DESKTOP_STARTUP_ID
	// environment variable, but in some cases (such as the application not
	// being launched using exec()) it can come from other sources.
	//
	// If the ID contains the string "_TIME" then the portion following that
	// string is taken to be the X11 timestamp of the event that triggered the
	// application to be launched and the GDK current event time is set
	// accordingly.
	//
	// The startup ID is also what is used to signal that the startup is
	// complete (for example, when opening a window or when calling
	// gdk_display_notify_startup_complete()).
	SetStartupNotificationIDX11Display(startupId string)
	// SetSurfaceScaleX11Display forces a specific window scale for all windows
	// on this display, instead of using the default or user configured scale.
	// This is can be used to disable scaling support by setting @scale to 1, or
	// to programmatically set the window scale.
	//
	// Once the scale is set by this call it will not change in response to
	// later user configuration changes.
	SetSurfaceScaleX11Display(scale int)
	// StringToCompoundTextX11Display: convert a string from the encoding of the
	// current locale into a form suitable for storing in a window property.
	StringToCompoundTextX11Display(str string) (encoding string, format int, ctext []byte, gint int)
	// TextPropertyToTextListX11Display: convert a text string from the encoding
	// as it is stored in a property into an array of strings in the encoding of
	// the current locale. (The elements of the array represent the
	// nul-separated elements of the original text string.)
	TextPropertyToTextListX11Display(encoding string, format int, text byte, length int, list *string) int
	// UngrabX11Display: ungrab @display after it has been grabbed with
	// gdk_x11_display_grab().
	UngrabX11Display()
	// UTF8ToCompoundTextX11Display converts from UTF-8 to compound text.
	UTF8ToCompoundTextX11Display(str string) (string, int, []byte, bool)
}

// x11Display implements the X11Display class.
type x11Display struct {
	gdk.Display
}

// WrapX11Display wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Display(obj *externglib.Object) X11Display {
	return x11Display{
		Display: gdk.WrapDisplay(obj),
	}
}

func marshalX11Display(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Display(obj), nil
}

func (d x11Display) ErrorTrapPopX11Display() int {
	var _arg0 *C.GdkDisplay // out
	var _cret C.int         // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_error_trap_pop(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (d x11Display) ErrorTrapPopIgnoredX11Display() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_error_trap_pop_ignored(_arg0)
}

func (d x11Display) ErrorTrapPushX11Display() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_error_trap_push(_arg0)
}

func (d x11Display) DefaultGroup() gdk.Surface {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_get_default_group(_arg0)

	var _surface gdk.Surface // out

	_surface = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Surface)

	return _surface
}

func (d x11Display) GlxVersion() (major int, minor int, ok bool) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.int        // in
	var _arg2 *C.int        // in
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_get_glx_version(_arg0, &_arg1, &_arg2)

	var _major int // out
	var _minor int // out
	var _ok bool   // out

	_major = int(_arg1)
	_minor = int(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _major, _minor, _ok
}

func (d x11Display) PrimaryMonitor() gdk.Monitor {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_get_primary_monitor(_arg0)

	var _monitor gdk.Monitor // out

	_monitor = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Monitor)

	return _monitor
}

func (d x11Display) Screen() X11Screen {
	var _arg0 *C.GdkDisplay   // out
	var _cret *C.GdkX11Screen // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_get_screen(_arg0)

	var _x11Screen X11Screen // out

	_x11Screen = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(X11Screen)

	return _x11Screen
}

func (d x11Display) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_get_startup_notification_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d x11Display) UserTime() uint32 {
	var _arg0 *C.GdkDisplay // out
	var _cret C.guint32     // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_get_user_time(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

func (d x11Display) GrabX11Display() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_grab(_arg0)
}

func (d x11Display) SetCursorThemeX11Display(theme string, size int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 C.int         // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(theme))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(size)

	C.gdk_x11_display_set_cursor_theme(_arg0, _arg1, _arg2)
}

func (d x11Display) SetStartupNotificationIDX11Display(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(startupId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_x11_display_set_startup_notification_id(_arg0, _arg1)
}

func (d x11Display) SetSurfaceScaleX11Display(scale int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.int         // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = C.int(scale)

	C.gdk_x11_display_set_surface_scale(_arg0, _arg1)
}

func (d x11Display) StringToCompoundTextX11Display(str string) (encoding string, format int, ctext []byte, gint int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 **C.char      // in
	var _arg3 *C.int        // in
	var _arg4 *C.guchar
	var _arg5 *C.int // in
	var _cret C.int  // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_x11_display_string_to_compound_text(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)

	var _encoding string // out
	var _format int      // out
	var _ctext []byte
	var _gint int // out

	{
		var refTmpIn *C.char
		var refTmpOut string

		refTmpIn = *_arg2

		refTmpOut = C.GoString(refTmpIn)

		_encoding = refTmpOut
	}
	_format = int(_arg3)
	_ctext = unsafe.Slice((*byte)(unsafe.Pointer(_arg4)), _arg5)
	runtime.SetFinalizer(&_ctext, func(v *[]byte) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	_gint = int(_cret)

	return _encoding, _format, _ctext, _gint
}

func (d x11Display) TextPropertyToTextListX11Display(encoding string, format int, text byte, length int, list *string) int {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 C.int         // out
	var _arg3 *C.guchar     // out
	var _arg4 C.int         // out
	var _arg5 ***C.char     // out
	var _cret C.int         // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(encoding))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(format)
	_arg3 = *C.guchar(text)
	_arg4 = C.int(length)
	{
		var refTmpIn string
		var refTmpOut *C.char

		refTmpIn = list

		refTmpOut = (*C.char)(C.CString(refTmpIn))
		defer C.free(unsafe.Pointer(refTmpOut))

		out0 := &refTmpOut
		out1 := &out0
		_arg5 = out1
	}

	_cret = C.gdk_x11_display_text_property_to_text_list(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (d x11Display) UngrabX11Display() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_ungrab(_arg0)
}

func (d x11Display) UTF8ToCompoundTextX11Display(str string) (string, int, []byte, bool) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 **C.char      // in
	var _arg3 *C.int        // in
	var _arg4 *C.guchar
	var _arg5 *C.int     // in
	var _cret C.gboolean // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_x11_display_utf8_to_compound_text(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)

	var _encoding string // out
	var _format int      // out
	var _ctext []byte
	var _ok bool // out

	{
		var refTmpIn *C.char
		var refTmpOut string

		refTmpIn = *_arg2

		refTmpOut = C.GoString(refTmpIn)

		_encoding = refTmpOut
	}
	_format = int(_arg3)
	_ctext = unsafe.Slice((*byte)(unsafe.Pointer(_arg4)), _arg5)
	runtime.SetFinalizer(&_ctext, func(v *[]byte) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	if _cret != 0 {
		_ok = true
	}

	return _encoding, _format, _ctext, _ok
}
