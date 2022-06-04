// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gdkapplaunchcontext.go.
var GTypeAppLaunchContext = coreglib.Type(C.gdk_app_launch_context_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeAppLaunchContext, F: marshalAppLaunchContext},
	})
}

// AppLaunchContext: GdkAppLaunchContext handles launching an application in a
// graphical context.
//
// It is an implementation of GAppLaunchContext that provides startup
// notification and allows to launch applications on a specific screen or
// workspace.
//
// Launching an application
//
//    GdkAppLaunchContext *context;
//
//    context = gdk_display_get_app_launch_context (display);
//
//    gdk_app_launch_context_set_display (display);
//    gdk_app_launch_context_set_timestamp (gdk_event_get_time (event));
//
//    if (!g_app_info_launch_default_for_uri ("http://www.gtk.org", context, &error))
//      g_warning ("Launching failed: s\n", error->message);
//
//    g_object_unref (context);.
type AppLaunchContext struct {
	_ [0]func() // equal guard
	gio.AppLaunchContext
}

var (
	_ coreglib.Objector = (*AppLaunchContext)(nil)
)

func wrapAppLaunchContext(obj *coreglib.Object) *AppLaunchContext {
	return &AppLaunchContext{
		AppLaunchContext: gio.AppLaunchContext{
			Object: obj,
		},
	}
}

func marshalAppLaunchContext(p uintptr) (interface{}, error) {
	return wrapAppLaunchContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Display gets the GdkDisplay that context is for.
//
// The function returns the following values:
//
//    - display of context.
//
func (context *AppLaunchContext) Display() *Display {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gdk", "AppLaunchContext").InvokeMethod("get_display", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _display *Display // out

	_display = wrapDisplay(coreglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// SetDesktop sets the workspace on which applications will be launched.
//
// This only works when running under a window manager that supports multiple
// workspaces, as described in the Extended Window Manager Hints
// (http://www.freedesktop.org/Standards/wm-spec).
//
// When the workspace is not specified or desktop is set to -1, it is up to the
// window manager to pick one, typically it will be the current workspace.
//
// The function takes the following parameters:
//
//    - desktop: number of a workspace, or -1.
//
func (context *AppLaunchContext) SetDesktop(desktop int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(desktop)

	girepository.MustFind("Gdk", "AppLaunchContext").InvokeMethod("set_desktop", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(desktop)
}

// SetIcon sets the icon for applications that are launched with this context.
//
// Window Managers can use this information when displaying startup
// notification.
//
// See also gdk.AppLaunchContext.SetIconName().
//
// The function takes the following parameters:
//
//    - icon (optional) or NULL.
//
func (context *AppLaunchContext) SetIcon(icon gio.Iconner) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	if icon != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(icon).Native()))
	}

	girepository.MustFind("Gdk", "AppLaunchContext").InvokeMethod("set_icon", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(icon)
}

// SetIconName sets the icon for applications that are launched with this
// context.
//
// The icon_name will be interpreted in the same way as the Icon field in
// desktop files. See also gdk.AppLaunchContext.SetIcon()().
//
// If both icon and icon_name are set, the icon_name takes priority. If neither
// icon or icon_name is set, the icon is taken from either the file that is
// passed to launched application or from the GAppInfo for the launched
// application itself.
//
// The function takes the following parameters:
//
//    - iconName (optional): icon name, or NULL.
//
func (context *AppLaunchContext) SetIconName(iconName string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	if iconName != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gdk", "AppLaunchContext").InvokeMethod("set_icon_name", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(iconName)
}

// SetTimestamp sets the timestamp of context.
//
// The timestamp should ideally be taken from the event that triggered the
// launch.
//
// Window managers can use this information to avoid moving the focus to the
// newly launched application when the user is busy typing in another window.
// This is also known as 'focus stealing prevention'.
//
// The function takes the following parameters:
//
//    - timestamp: timestamp.
//
func (context *AppLaunchContext) SetTimestamp(timestamp uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(*C.guint32)(unsafe.Pointer(&_args[1])) = C.guint32(timestamp)

	girepository.MustFind("Gdk", "AppLaunchContext").InvokeMethod("set_timestamp", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(timestamp)
}
