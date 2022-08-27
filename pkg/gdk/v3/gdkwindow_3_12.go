// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

// EventCompression: get the current event compression setting for this window.
//
// The function returns the following values:
//
//    - ok: TRUE if motion events will be compressed.
//
func (window *Window) EventCompression() bool {
	var _arg0 *C.GdkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_cret = C.gdk_window_get_event_compression(_arg0)
	runtime.KeepAlive(window)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetEventCompression determines whether or not extra unprocessed motion events
// in the event queue can be discarded. If TRUE only the most recent event will
// be delivered.
//
// Some types of applications, e.g. paint programs, need to see all motion
// events and will benefit from turning off event compression.
//
// By default, event compression is enabled.
//
// The function takes the following parameters:
//
//    - eventCompression: TRUE if motion events should be compressed.
//
func (window *Window) SetEventCompression(eventCompression bool) {
	var _arg0 *C.GdkWindow // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GdkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	if eventCompression {
		_arg1 = C.TRUE
	}

	C.gdk_window_set_event_compression(_arg0, _arg1)
	runtime.KeepAlive(window)
	runtime.KeepAlive(eventCompression)
}

// SetShadowWidth: newer GTK+ windows using client-side decorations use extra
// geometry around their frames for effects like shadows and invisible borders.
// Window managers that want to maximize windows or snap to edges need to know
// where the extents of the actual frame lie, so that users don’t feel like
// windows are snapping against random invisible edges.
//
// Note that this property is automatically updated by GTK+, so this function
// should only be used by applications which do not use GTK+ to create toplevel
// windows.
//
// The function takes the following parameters:
//
//    - left extent.
//    - right extent.
//    - top extent.
//    - bottom extent.
//
func (window *Window) SetShadowWidth(left, right, top, bottom int) {
	var _arg0 *C.GdkWindow // out
	var _arg1 C.gint       // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out
	var _arg4 C.gint       // out

	_arg0 = (*C.GdkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	_arg1 = C.gint(left)
	_arg2 = C.gint(right)
	_arg3 = C.gint(top)
	_arg4 = C.gint(bottom)

	C.gdk_window_set_shadow_width(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(window)
	runtime.KeepAlive(left)
	runtime.KeepAlive(right)
	runtime.KeepAlive(top)
	runtime.KeepAlive(bottom)
}
