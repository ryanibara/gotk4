// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
import "C"

// X11FreeCompoundText frees the data returned from
// gdk_x11_display_string_to_compound_text().
//
// The function takes the following parameters:
//
//    - ctext: pointer stored in ctext from a call to
//    gdk_x11_display_string_to_compound_text().
//
func X11FreeCompoundText(ctext *byte) {
	var _arg1 *C.guchar // out

	_arg1 = (*C.guchar)(unsafe.Pointer(ctext))

	C.gdk_x11_free_compound_text(_arg1)
	runtime.KeepAlive(ctext)
}
