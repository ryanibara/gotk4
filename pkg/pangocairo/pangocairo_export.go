// Code generated by girgen. DO NOT EDIT.

package pangocairo

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <pango/pangocairo.h>
import "C"

//export _gotk4_pangocairo1_ShapeRendererFunc
func _gotk4_pangocairo1_ShapeRendererFunc(arg1 *C.cairo_t, arg2 *C.PangoAttrShape, arg3 C.gboolean, arg4 C.gpointer) {
	var fn ShapeRendererFunc
	{
		v := gbox.Get(uintptr(arg4))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ShapeRendererFunc)
	}

	var _cr *cairo.Context     // out
	var _attr *pango.AttrShape // out
	var _doPath bool           // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.cairo_t)(unsafe.Pointer(v.Native())))
	})
	_attr = (*pango.AttrShape)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	if arg3 != 0 {
		_doPath = true
	}

	fn(_cr, _attr, _doPath)
}
