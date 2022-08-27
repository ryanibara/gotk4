// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// Serial returns the current serial number of layout.
//
// The serial number is initialized to an small number larger than zero when a
// new layout is created and is increased whenever the layout is changed using
// any of the setter functions, or the PangoContext it uses has changed. The
// serial may wrap, but will never have the value 0. Since it can wrap, never
// compare it with "less than", always use "not equals".
//
// This can be used to automatically detect changes to a PangoLayout, and is
// useful for example to decide whether a layout needs redrawing. To force the
// serial to be increased, use pango.Layout.ContextChanged().
//
// The function returns the following values:
//
//    - guint: current serial number of layout.
//
func (layout *Layout) Serial() uint {
	var _arg0 *C.PangoLayout // out
	var _cret C.guint        // in

	_arg0 = (*C.PangoLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))

	_cret = C.pango_layout_get_serial(_arg0)
	runtime.KeepAlive(layout)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}
