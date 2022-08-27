// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdkx.h>
import "C"

// SetFrameSyncEnabled: this function can be used to disable frame
// synchronization for a window. Normally frame synchronziation will be enabled
// or disabled based on whether the system has a compositor that supports frame
// synchronization, but if the window is not directly managed by the window
// manager, then frame synchronziation may need to be disabled. This is the case
// for a window embedded via the XEMBED protocol.
//
// The function takes the following parameters:
//
//    - frameSyncEnabled: whether frame-synchronization should be enabled.
//
func (window *X11Window) SetFrameSyncEnabled(frameSyncEnabled bool) {
	var _arg0 *C.GdkWindow // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GdkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	if frameSyncEnabled {
		_arg1 = C.TRUE
	}

	C.gdk_x11_window_set_frame_sync_enabled(_arg0, _arg1)
	runtime.KeepAlive(window)
	runtime.KeepAlive(frameSyncEnabled)
}
