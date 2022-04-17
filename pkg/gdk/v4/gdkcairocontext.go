// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gdkcairocontext.go.
var GTypeCairoContext = externglib.Type(C.gdk_cairo_context_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeCairoContext, F: marshalCairoContext},
	})
}

// CairoContext: GdkCairoContext is an object representing the platform-specific
// draw context.
//
// GdkCairoContexts are created for a surface using
// gdk.Surface.CreateCairoContext(), and the context can then be used to draw on
// that surface.
type CairoContext struct {
	_ [0]func() // equal guard
	DrawContext
}

var (
	_ DrawContexter = (*CairoContext)(nil)
)

// CairoContexter describes types inherited from class CairoContext.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type CairoContexter interface {
	externglib.Objector
	baseCairoContext() *CairoContext
}

var _ CairoContexter = (*CairoContext)(nil)

func wrapCairoContext(obj *externglib.Object) *CairoContext {
	return &CairoContext{
		DrawContext: DrawContext{
			Object: obj,
		},
	}
}

func marshalCairoContext(p uintptr) (interface{}, error) {
	return wrapCairoContext(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *CairoContext) baseCairoContext() *CairoContext {
	return self
}

// BaseCairoContext returns the underlying base object.
func BaseCairoContext(obj CairoContexter) *CairoContext {
	return obj.baseCairoContext()
}

// CairoCreate retrieves a Cairo context to be used to draw on the GdkSurface of
// context.
//
// A call to gdk.DrawContext.BeginFrame() with this context must have been done
// or this function will return NULL.
//
// The returned context is guaranteed to be valid until
// gdk.DrawContext.EndFrame() is called.
//
// The function returns the following values:
//
//    - context (optional): cairo context to be used to draw the contents of the
//      GdkSurface. NULL is returned when context is not drawing.
//
func (self *CairoContext) CairoCreate() *cairo.Context {
	var _arg0 *C.GdkCairoContext // out
	var _cret *C.cairo_t         // in

	_arg0 = (*C.GdkCairoContext)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gdk_cairo_context_cairo_create(_arg0)
	runtime.KeepAlive(self)

	var _context *cairo.Context // out

	if _cret != nil {
		_context = cairo.WrapContext(uintptr(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(_context, func(v *cairo.Context) {
			C.cairo_destroy((*C.cairo_t)(unsafe.Pointer(v.Native())))
		})
	}

	return _context
}
