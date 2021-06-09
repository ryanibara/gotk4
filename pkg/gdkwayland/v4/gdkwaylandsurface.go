// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
)

// #cgo pkg-config: gtk4-wayland gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/wayland/gdkwayland.h>
import "C"

// WaylandToplevelExported: callback that gets called when the handle for a
// surface has been obtained from the Wayland compositor.
//
// This callback is used in [method@GdkWayland.WaylandToplevel.export_handle].
//
// The @handle can be passed to other processes, for the purpose of marking
// surfaces as transient for out-of-process surfaces.
type WaylandToplevelExported func()

//export gotk4_WaylandToplevelExported
func gotk4_WaylandToplevelExported(arg0 *C.GdkToplevel, arg1 *C.char, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(WaylandToplevelExported)
	fn()
}