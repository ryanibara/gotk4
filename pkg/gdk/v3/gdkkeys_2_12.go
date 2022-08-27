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

// HaveBidiLayouts determines if keyboard layouts for both right-to-left and
// left-to-right languages are in use.
//
// The function returns the following values:
//
//    - ok: TRUE if there are layouts in both directions, FALSE otherwise.
//
func (keymap *Keymap) HaveBidiLayouts() bool {
	var _arg0 *C.GdkKeymap // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(coreglib.InternObject(keymap).Native()))

	_cret = C.gdk_keymap_have_bidi_layouts(_arg0)
	runtime.KeepAlive(keymap)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
