// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeDrawingContext = coreglib.Type(C.gdk_drawing_context_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDrawingContext, F: marshalDrawingContext},
	})
}

// DrawingContext is an object that represents the current drawing state of a
// Window.
//
// It's possible to use a DrawingContext to draw on a Window via rendering API
// like Cairo or OpenGL.
//
// A DrawingContext can only be created by calling gdk_window_begin_draw_frame()
// and will be valid until a call to gdk_window_end_draw_frame().
//
// DrawingContext is available since GDK 3.22.
type DrawingContext struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DrawingContext)(nil)
)

func wrapDrawingContext(obj *coreglib.Object) *DrawingContext {
	return &DrawingContext{
		Object: obj,
	}
}

func marshalDrawingContext(p uintptr) (interface{}, error) {
	return wrapDrawingContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CairoContext retrieves a Cairo context to be used to draw on the Window that
// created the DrawingContext.
//
// The returned context is guaranteed to be valid as long as the DrawingContext
// is valid, that is between a call to gdk_window_begin_draw_frame() and
// gdk_window_end_draw_frame().
//
// The function returns the following values:
//
//   - ret: cairo context to be used to draw the contents of the Window.
//     The context is owned by the DrawingContext and should not be destroyed.
//
func (context *DrawingContext) CairoContext() *cairo.Context {
	var _arg0 *C.GdkDrawingContext // out
	var _cret *C.cairo_t           // in

	_arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gdk_drawing_context_get_cairo_context(_arg0)
	runtime.KeepAlive(context)

	var _ret *cairo.Context // out

	_ret = cairo.WrapContext(uintptr(unsafe.Pointer(_cret)))
	C.cairo_reference(_cret)
	runtime.SetFinalizer(_ret, func(v *cairo.Context) {
		C.cairo_destroy((*C.cairo_t)(unsafe.Pointer(v.Native())))
	})

	return _ret
}

// Clip retrieves a copy of the clip region used when creating the context.
//
// The function returns the following values:
//
//   - region (optional): cairo region.
//
func (context *DrawingContext) Clip() *cairo.Region {
	var _arg0 *C.GdkDrawingContext // out
	var _cret *C.cairo_region_t    // in

	_arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gdk_drawing_context_get_clip(_arg0)
	runtime.KeepAlive(context)

	var _region *cairo.Region // out

	if _cret != nil {
		{
			_pp := &struct{ p unsafe.Pointer }{unsafe.Pointer(_cret)}
			_region = (*cairo.Region)(unsafe.Pointer(_pp))
		}
		runtime.SetFinalizer(_region, func(v *cairo.Region) {
			C.cairo_region_destroy((*C.cairo_region_t)(unsafe.Pointer(v.Native())))
		})
	}

	return _region
}

// Window retrieves the window that created the drawing context.
//
// The function returns the following values:
//
//   - window: Window.
//
func (context *DrawingContext) Window() Windower {
	var _arg0 *C.GdkDrawingContext // out
	var _cret *C.GdkWindow         // in

	_arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gdk_drawing_context_get_window(_arg0)
	runtime.KeepAlive(context)

	var _window Windower // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Windower)
			return ok
		})
		rv, ok := casted.(Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_window = rv
	}

	return _window
}

// IsValid checks whether the given DrawingContext is valid.
//
// The function returns the following values:
//
//   - ok: TRUE if the context is valid.
//
func (context *DrawingContext) IsValid() bool {
	var _arg0 *C.GdkDrawingContext // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gdk_drawing_context_is_valid(_arg0)
	runtime.KeepAlive(context)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
