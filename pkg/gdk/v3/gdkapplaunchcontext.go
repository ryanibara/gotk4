// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_app_launch_context_get_type()), F: marshalAppLaunchContexter},
	})
}

// AppLaunchContexter describes AppLaunchContext's methods.
type AppLaunchContexter interface {
	// SetDesktop sets the workspace on which applications will be launched when
	// using this context when running under a window manager that supports
	// multiple workspaces, as described in the Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec).
	SetDesktop(desktop int)
	// SetDisplay sets the display on which applications will be launched when
	// using this context.
	SetDisplay(display Displayer)
	// SetIcon sets the icon for applications that are launched with this
	// context.
	SetIcon(icon gio.Iconer)
	// SetIconName sets the icon for applications that are launched with this
	// context.
	SetIconName(iconName string)
	// SetScreen sets the screen on which applications will be launched when
	// using this context.
	SetScreen(screen Screener)
	// SetTimestamp sets the timestamp of @context.
	SetTimestamp(timestamp uint32)
}

// AppLaunchContext is an implementation of LaunchContext that handles launching
// an application in a graphical context. It provides startup notification and
// allows to launch applications on a specific screen or workspace.
//
// Launching an application
//
//    GdkAppLaunchContext *context;
//
//    context = gdk_display_get_app_launch_context (display);
//
//    gdk_app_launch_context_set_screen (screen);
//    gdk_app_launch_context_set_timestamp (event->time);
//
//    if (!g_app_info_launch_default_for_uri ("http://www.gtk.org", context, &error))
//      g_warning ("Launching failed: s\n", error->message);
//
//    g_object_unref (context);
type AppLaunchContext struct {
	gio.AppLaunchContext
}

var (
	_ AppLaunchContexter = (*AppLaunchContext)(nil)
	_ gextras.Nativer    = (*AppLaunchContext)(nil)
)

func wrapAppLaunchContext(obj *externglib.Object) *AppLaunchContext {
	return &AppLaunchContext{
		AppLaunchContext: gio.AppLaunchContext{
			Object: obj,
		},
	}
}

func marshalAppLaunchContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAppLaunchContext(obj), nil
}

// NewAppLaunchContext creates a new AppLaunchContext.
//
// Deprecated: Use gdk_display_get_app_launch_context() instead.
func NewAppLaunchContext() *AppLaunchContext {
	var _cret *C.GdkAppLaunchContext // in

	_cret = C.gdk_app_launch_context_new()

	var _appLaunchContext *AppLaunchContext // out

	_appLaunchContext = wrapAppLaunchContext(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _appLaunchContext
}

// SetDesktop sets the workspace on which applications will be launched when
// using this context when running under a window manager that supports multiple
// workspaces, as described in the Extended Window Manager Hints
// (http://www.freedesktop.org/Standards/wm-spec).
//
// When the workspace is not specified or @desktop is set to -1, it is up to the
// window manager to pick one, typically it will be the current workspace.
func (context *AppLaunchContext) SetDesktop(desktop int) {
	var _arg0 *C.GdkAppLaunchContext // out
	var _arg1 C.gint                 // out

	_arg0 = (*C.GdkAppLaunchContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.gint(desktop)

	C.gdk_app_launch_context_set_desktop(_arg0, _arg1)
}

// SetDisplay sets the display on which applications will be launched when using
// this context. See also gdk_app_launch_context_set_screen().
//
// Deprecated: Use gdk_display_get_app_launch_context() instead.
func (context *AppLaunchContext) SetDisplay(display Displayer) {
	var _arg0 *C.GdkAppLaunchContext // out
	var _arg1 *C.GdkDisplay          // out

	_arg0 = (*C.GdkAppLaunchContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.GdkDisplay)(unsafe.Pointer((display).(gextras.Nativer).Native()))

	C.gdk_app_launch_context_set_display(_arg0, _arg1)
}

// SetIcon sets the icon for applications that are launched with this context.
//
// Window Managers can use this information when displaying startup
// notification.
//
// See also gdk_app_launch_context_set_icon_name().
func (context *AppLaunchContext) SetIcon(icon gio.Iconer) {
	var _arg0 *C.GdkAppLaunchContext // out
	var _arg1 *C.GIcon               // out

	_arg0 = (*C.GdkAppLaunchContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer((icon).(gextras.Nativer).Native()))

	C.gdk_app_launch_context_set_icon(_arg0, _arg1)
}

// SetIconName sets the icon for applications that are launched with this
// context. The @icon_name will be interpreted in the same way as the Icon field
// in desktop files. See also gdk_app_launch_context_set_icon().
//
// If both @icon and @icon_name are set, the @icon_name takes priority. If
// neither @icon or @icon_name is set, the icon is taken from either the file
// that is passed to launched application or from the Info for the launched
// application itself.
func (context *AppLaunchContext) SetIconName(iconName string) {
	var _arg0 *C.GdkAppLaunchContext // out
	var _arg1 *C.char                // out

	_arg0 = (*C.GdkAppLaunchContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))

	C.gdk_app_launch_context_set_icon_name(_arg0, _arg1)
}

// SetScreen sets the screen on which applications will be launched when using
// this context. See also gdk_app_launch_context_set_display().
//
// If both @screen and @display are set, the @screen takes priority. If neither
// @screen or @display are set, the default screen and display are used.
func (context *AppLaunchContext) SetScreen(screen Screener) {
	var _arg0 *C.GdkAppLaunchContext // out
	var _arg1 *C.GdkScreen           // out

	_arg0 = (*C.GdkAppLaunchContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer((screen).(gextras.Nativer).Native()))

	C.gdk_app_launch_context_set_screen(_arg0, _arg1)
}

// SetTimestamp sets the timestamp of @context. The timestamp should ideally be
// taken from the event that triggered the launch.
//
// Window managers can use this information to avoid moving the focus to the
// newly launched application when the user is busy typing in another window.
// This is also known as 'focus stealing prevention'.
func (context *AppLaunchContext) SetTimestamp(timestamp uint32) {
	var _arg0 *C.GdkAppLaunchContext // out
	var _arg1 C.guint32              // out

	_arg0 = (*C.GdkAppLaunchContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.guint32(timestamp)

	C.gdk_app_launch_context_set_timestamp(_arg0, _arg1)
}
