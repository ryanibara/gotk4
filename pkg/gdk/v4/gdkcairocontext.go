// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_cairo_context_get_type()), F: marshalCairoContexter},
	})
}

// CairoContexter describes CairoContext's methods.
type CairoContexter interface {
	// CairoCreate retrieves a Cairo context to be used to draw on the
	// `GdkSurface` of @context.
	CairoCreate() *cairo.Context
}

// CairoContext: `GdkCairoContext` is an object representing the
// platform-specific draw context.
//
// `GdkCairoContext`s are created for a surface using
// [method@Gdk.Surface.create_cairo_context], and the context can then be used
// to draw on that surface.
type CairoContext struct {
	DrawContext
}

var (
	_ CairoContexter  = (*CairoContext)(nil)
	_ gextras.Nativer = (*CairoContext)(nil)
)

func wrapCairoContext(obj *externglib.Object) CairoContexter {
	return &CairoContext{
		DrawContext: DrawContext{
			Object: obj,
		},
	}
}

func marshalCairoContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCairoContext(obj), nil
}

// CairoCreate retrieves a Cairo context to be used to draw on the `GdkSurface`
// of @context.
//
// A call to [method@Gdk.DrawContext.begin_frame] with this @context must have
// been done or this function will return nil.
//
// The returned context is guaranteed to be valid until
// [method@Gdk.DrawContext.end_frame] is called.
func (self *CairoContext) CairoCreate() *cairo.Context {
	var _arg0 *C.GdkCairoContext // out
	var _cret *C.cairo_t         // in

	_arg0 = (*C.GdkCairoContext)(unsafe.Pointer(self.Native()))

	_cret = C.gdk_cairo_context_cairo_create(_arg0)

	var _context *cairo.Context // out

	_context = (*cairo.Context)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_context, func(v *cairo.Context) {
		C.free(unsafe.Pointer(v))
	})

	return _context
}
