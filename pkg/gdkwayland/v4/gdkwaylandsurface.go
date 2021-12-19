// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/wayland/gdkwayland.h>
// extern void _gotk4_gdkwayland4_WaylandToplevelExported(GdkToplevel*, char*, gpointer);
// extern void callbackDelete(gpointer);
import "C"

// WaylandToplevelExported: callback that gets called when the handle for a
// surface has been obtained from the Wayland compositor.
//
// This callback is used in gdkwayland.WaylandToplevel.ExportHandle().
//
// The handle can be passed to other processes, for the purpose of marking
// surfaces as transient for out-of-process surfaces.
type WaylandToplevelExported func(toplevel *WaylandToplevel, handle string)

//export _gotk4_gdkwayland4_WaylandToplevelExported
func _gotk4_gdkwayland4_WaylandToplevelExported(arg1 *C.GdkToplevel, arg2 *C.char, arg3 C.gpointer) {
	var fn WaylandToplevelExported
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(WaylandToplevelExported)
	}

	var _toplevel *WaylandToplevel // out
	var _handle string             // out

	_toplevel = wrapWaylandToplevel(externglib.Take(unsafe.Pointer(arg1)))
	_handle = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	fn(_toplevel, _handle)
}

// ExportHandle: asynchronously obtains a handle for a surface that can be
// passed to other processes.
//
// When the handle has been obtained, callback will be called.
//
// It is an error to call this function on a surface that is already exported.
//
// When the handle is no longer needed,
// gdkwayland.WaylandToplevel.UnexportHandle() should be called to clean up
// resources.
//
// The main purpose for obtaining a handle is to mark a surface from another
// surface as transient for this one, see
// gdkwayland.WaylandToplevel.SetTransientForExported().
//
// Note that this API depends on an unstable Wayland protocol, and thus may
// require changes in the future.
//
// The function takes the following parameters:
//
//    - callback to call with the handle.
//
// The function returns the following values:
//
//    - ok: TRUE if the handle has been requested, FALSE if an error occurred.
//
func (toplevel *WaylandToplevel) ExportHandle(callback WaylandToplevelExported) bool {
	var _arg0 *C.GdkToplevel               // out
	var _arg1 C.GdkWaylandToplevelExported // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify
	var _cret C.gboolean // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gdkwayland4_WaylandToplevelExported)
	_arg2 = C.gpointer(gbox.Assign(callback))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	_cret = C.gdk_wayland_toplevel_export_handle(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(toplevel)
	runtime.KeepAlive(callback)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetApplicationID sets the application id on a GdkToplevel.
//
// The function takes the following parameters:
//
//    - applicationId: application id for the toplevel.
//
func (toplevel *WaylandToplevel) SetApplicationID(applicationId string) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(applicationId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_wayland_toplevel_set_application_id(_arg0, _arg1)
	runtime.KeepAlive(toplevel)
	runtime.KeepAlive(applicationId)
}

// SetTransientForExported marks toplevel as transient for the surface to which
// the given parent_handle_str refers.
//
// Typically, the handle will originate from a
// gdkwayland.WaylandToplevel.ExportHandle() call in another process.
//
// Note that this API depends on an unstable Wayland protocol, and thus may
// require changes in the future.
//
// The function takes the following parameters:
//
//    - parentHandleStr: exported handle for a surface.
//
// The function returns the following values:
//
//    - ok: TRUE if the surface has been marked as transient, FALSE if an error
//      occurred.
//
func (toplevel *WaylandToplevel) SetTransientForExported(parentHandleStr string) bool {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.char        // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(parentHandleStr)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_wayland_toplevel_set_transient_for_exported(_arg0, _arg1)
	runtime.KeepAlive(toplevel)
	runtime.KeepAlive(parentHandleStr)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnexportHandle destroys the handle that was obtained with
// gdk_wayland_toplevel_export_handle().
//
// It is an error to call this function on a surface that does not have a
// handle.
//
// Note that this API depends on an unstable Wayland protocol, and thus may
// require changes in the future.
func (toplevel *WaylandToplevel) UnexportHandle() {
	var _arg0 *C.GdkToplevel // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))

	C.gdk_wayland_toplevel_unexport_handle(_arg0)
	runtime.KeepAlive(toplevel)
}
