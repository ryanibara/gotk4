// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"reflect"
	"runtime"
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
		{T: externglib.Type(C.gdk_x11_surface_get_type()), F: marshalX11Surfacer},
	})
}

// X11GetServerTime: routine to get the current X server time stamp.
//
// The function takes the following parameters:
//
//    - surface used for communication with the server. The surface must have
//      GDK_PROPERTY_CHANGE_MASK in its events mask or a hang will result.
//
// The function returns the following values:
//
//    - guint32: time stamp.
//
func X11GetServerTime(surface *X11Surface) uint32 {
	var _arg1 *C.GdkSurface // out
	var _cret C.guint32     // in

	_arg1 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_x11_get_server_time(_arg1)
	runtime.KeepAlive(surface)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

type X11Surface struct {
	_ [0]func() // equal guard
	gdk.Surface
}

var (
	_ gdk.Surfacer = (*X11Surface)(nil)
)

func wrapX11Surface(obj *externglib.Object) *X11Surface {
	return &X11Surface{
		Surface: gdk.Surface{
			Object: obj,
		},
	}
}

func marshalX11Surfacer(p uintptr) (interface{}, error) {
	return wrapX11Surface(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Desktop gets the number of the workspace surface is on.
//
// The function returns the following values:
//
//    - guint32: current workspace of surface.
//
func (surface *X11Surface) Desktop() uint32 {
	var _arg0 *C.GdkSurface // out
	var _cret C.guint32     // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_x11_surface_get_desktop(_arg0)
	runtime.KeepAlive(surface)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// Group returns the group this surface belongs to.
//
// The function returns the following values:
//
//    - ret: group of this surface;.
//
func (surface *X11Surface) Group() gdk.Surfacer {
	var _arg0 *C.GdkSurface // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_x11_surface_get_group(_arg0)
	runtime.KeepAlive(surface)

	var _ret gdk.Surfacer // out

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
		_ret = rv
	}

	return _ret
}

// MoveToCurrentDesktop moves the surface to the correct workspace when running
// under a window manager that supports multiple workspaces, as described in the
// Extended Window Manager Hints (http://www.freedesktop.org/Standards/wm-spec)
// specification. Will not do anything if the surface is already on all
// workspaces.
func (surface *X11Surface) MoveToCurrentDesktop() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	C.gdk_x11_surface_move_to_current_desktop(_arg0)
	runtime.KeepAlive(surface)
}

// MoveToDesktop moves the surface to the given workspace when running unde a
// window manager that supports multiple workspaces, as described in the
// Extended Window Manager Hints (http://www.freedesktop.org/Standards/wm-spec)
// specification.
//
// The function takes the following parameters:
//
//    - desktop: number of the workspace to move the surface to.
//
func (surface *X11Surface) MoveToDesktop(desktop uint32) {
	var _arg0 *C.GdkSurface // out
	var _arg1 C.guint32     // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	_arg1 = C.guint32(desktop)

	C.gdk_x11_surface_move_to_desktop(_arg0, _arg1)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(desktop)
}

// SetFrameSyncEnabled: this function can be used to disable frame
// synchronization for a surface. Normally frame synchronziation will be enabled
// or disabled based on whether the system has a compositor that supports frame
// synchronization, but if the surface is not directly managed by the window
// manager, then frame synchronziation may need to be disabled. This is the case
// for a surface embedded via the XEMBED protocol.
//
// The function takes the following parameters:
//
//    - frameSyncEnabled: whether frame-synchronization should be enabled.
//
func (surface *X11Surface) SetFrameSyncEnabled(frameSyncEnabled bool) {
	var _arg0 *C.GdkSurface // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	if frameSyncEnabled {
		_arg1 = C.TRUE
	}

	C.gdk_x11_surface_set_frame_sync_enabled(_arg0, _arg1)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(frameSyncEnabled)
}

// SetGroup sets the group leader of surface to be leader. See the ICCCM for
// details.
//
// The function takes the following parameters:
//
//    - leader: Surface.
//
func (surface *X11Surface) SetGroup(leader gdk.Surfacer) {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	_arg1 = (*C.GdkSurface)(unsafe.Pointer(leader.Native()))

	C.gdk_x11_surface_set_group(_arg0, _arg1)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(leader)
}

// SetSkipPagerHint sets a hint on surface that pagers should not display it.
// See the EWMH for details.
//
// The function takes the following parameters:
//
//    - skipsPager: TRUE to skip pagers.
//
func (surface *X11Surface) SetSkipPagerHint(skipsPager bool) {
	var _arg0 *C.GdkSurface // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	if skipsPager {
		_arg1 = C.TRUE
	}

	C.gdk_x11_surface_set_skip_pager_hint(_arg0, _arg1)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(skipsPager)
}

// SetSkipTaskbarHint sets a hint on surface that taskbars should not display
// it. See the EWMH for details.
//
// The function takes the following parameters:
//
//    - skipsTaskbar: TRUE to skip taskbars.
//
func (surface *X11Surface) SetSkipTaskbarHint(skipsTaskbar bool) {
	var _arg0 *C.GdkSurface // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	if skipsTaskbar {
		_arg1 = C.TRUE
	}

	C.gdk_x11_surface_set_skip_taskbar_hint(_arg0, _arg1)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(skipsTaskbar)
}

// SetThemeVariant: GTK applications can request a dark theme variant. In order
// to make other applications - namely window managers using GTK for themeing -
// aware of this choice, GTK uses this function to export the requested theme
// variant as _GTK_THEME_VARIANT property on toplevel surfaces.
//
// Note that this property is automatically updated by GTK, so this function
// should only be used by applications which do not use GTK to create toplevel
// surfaces.
//
// The function takes the following parameters:
//
//    - variant: theme variant to export.
//
func (surface *X11Surface) SetThemeVariant(variant string) {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(variant)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_x11_surface_set_theme_variant(_arg0, _arg1)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(variant)
}

// SetUrgencyHint sets a hint on surface that it needs user attention. See the
// ICCCM for details.
//
// The function takes the following parameters:
//
//    - urgent: TRUE to indicate urgenct attention needed.
//
func (surface *X11Surface) SetUrgencyHint(urgent bool) {
	var _arg0 *C.GdkSurface // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	if urgent {
		_arg1 = C.TRUE
	}

	C.gdk_x11_surface_set_urgency_hint(_arg0, _arg1)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(urgent)
}

// SetUserTime: application can use this call to update the _NET_WM_USER_TIME
// property on a toplevel surface. This property stores an Xserver time which
// represents the time of the last user input event received for this surface.
// This property may be used by the window manager to alter the focus, stacking,
// and/or placement behavior of surfaces when they are mapped depending on
// whether the new surface was created by a user action or is a "pop-up" surface
// activated by a timer or some other event.
//
// Note that this property is automatically updated by GDK, so this function
// should only be used by applications which handle input events bypassing GDK.
//
// The function takes the following parameters:
//
//    - timestamp: XServer timestamp to which the property should be set.
//
func (surface *X11Surface) SetUserTime(timestamp uint32) {
	var _arg0 *C.GdkSurface // out
	var _arg1 C.guint32     // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	_arg1 = C.guint32(timestamp)

	C.gdk_x11_surface_set_user_time(_arg0, _arg1)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(timestamp)
}

// SetUTF8Property: this function modifies or removes an arbitrary X11 window
// property of type UTF8_STRING. If the given surface is not a toplevel surface,
// it is ignored.
//
// The function takes the following parameters:
//
//    - name: property name, will be interned as an X atom.
//    - value (optional): property value, or NULL to delete.
//
func (surface *X11Surface) SetUTF8Property(name, value string) {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.char       // out
	var _arg2 *C.char       // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	if value != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	C.gdk_x11_surface_set_utf8_property(_arg0, _arg1, _arg2)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
}
