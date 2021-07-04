// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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

	ErrorTrapPopX11Display() int

	ErrorTrapPopIgnoredX11Display()

	ErrorTrapPushX11Display()

	DefaultGroup() gdk.Surface

	GlxVersion() (major int, minor int, ok bool)

	PrimaryMonitor() gdk.Monitor

	Screen() X11Screen

	StartupNotificationID() string

	UserTime() uint32

	GrabX11Display()

	SetCursorThemeX11Display(theme string, size int)

	SetStartupNotificationIDX11Display(startupId string)

	SetSurfaceScaleX11Display(scale int)

	StringToCompoundTextX11Display(str string) (encoding string, format int, ctext []byte, gint int)

	TextPropertyToTextListX11Display(encoding string, format int, text *byte, length int, list **string) int

	UngrabX11Display()

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
	var _arg1 C.int         // in
	var _arg2 C.int         // in
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
	var _arg2 *C.char       // in
	var _arg3 C.int         // in
	var _arg4 *C.guchar
	var _arg5 C.int // in
	var _cret C.int // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_x11_display_string_to_compound_text(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)

	var _encoding string // out
	var _format int      // out
	var _ctext []byte
	var _gint int // out

	_encoding = C.GoString(_arg2)
	_format = int(_arg3)
	_ctext = unsafe.Slice((*byte)(unsafe.Pointer(_arg4)), _arg5)
	runtime.SetFinalizer(&_ctext, func(v *[]byte) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	_gint = int(_cret)

	return _encoding, _format, _ctext, _gint
}

func (d x11Display) TextPropertyToTextListX11Display(encoding string, format int, text *byte, length int, list **string) int {
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
	_arg3 = (*C.guchar)(unsafe.Pointer(text))
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
	var _arg2 *C.char       // in
	var _arg3 C.int         // in
	var _arg4 *C.guchar
	var _arg5 C.int      // in
	var _cret C.gboolean // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_x11_display_utf8_to_compound_text(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)

	var _encoding string // out
	var _format int      // out
	var _ctext []byte
	var _ok bool // out

	_encoding = C.GoString(_arg2)
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
