// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_drawing_context_get_type()), F: marshalDrawingContexter},
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
// DrawingContext is available since GDK 3.22
type DrawingContext struct {
	*externglib.Object
}

func wrapDrawingContext(obj *externglib.Object) *DrawingContext {
	return &DrawingContext{
		Object: obj,
	}
}

func marshalDrawingContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDrawingContext(obj), nil
}

// CairoContext retrieves a Cairo context to be used to draw on the Window that
// created the DrawingContext.
//
// The returned context is guaranteed to be valid as long as the DrawingContext
// is valid, that is between a call to gdk_window_begin_draw_frame() and
// gdk_window_end_draw_frame().
func (context *DrawingContext) CairoContext() *cairo.Context {
	var _arg0 *C.GdkDrawingContext // out
	var _cret *C.cairo_t           // in

	_arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drawing_context_get_cairo_context(_arg0)

	var _ret *cairo.Context // out

	_ret = cairo.WrapContext(uintptr(unsafe.Pointer(_cret)))
	C.cairo_reference(_cret)
	runtime.SetFinalizer(_ret, func(v *cairo.Context) {
		C.cairo_destroy((*C.cairo_t)(unsafe.Pointer(v.Native())))
	})

	return _ret
}

// Clip retrieves a copy of the clip region used when creating the context.
func (context *DrawingContext) Clip() *cairo.Region {
	var _arg0 *C.GdkDrawingContext // out
	var _cret *C.cairo_region_t    // in

	_arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drawing_context_get_clip(_arg0)

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
func (context *DrawingContext) Window() Windower {
	var _arg0 *C.GdkDrawingContext // out
	var _cret *C.GdkWindow         // in

	_arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drawing_context_get_window(_arg0)

	var _window Windower // out

	_window = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Windower)

	return _window
}

// IsValid checks whether the given DrawingContext is valid.
func (context *DrawingContext) IsValid() bool {
	var _arg0 *C.GdkDrawingContext // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drawing_context_is_valid(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
