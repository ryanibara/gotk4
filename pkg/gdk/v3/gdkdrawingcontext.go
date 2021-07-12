// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
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

// DrawingContexter describes DrawingContext's methods.
type DrawingContexter interface {
	// CairoContext retrieves a Cairo context to be used to draw on the Window
	// that created the DrawingContext.
	CairoContext() *cairo.Context
	// Clip retrieves a copy of the clip region used when creating the @context.
	Clip() *cairo.Region
	// Window retrieves the window that created the drawing @context.
	Window() *Window
	// IsValid checks whether the given DrawingContext is valid.
	IsValid() bool
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

var (
	_ DrawingContexter = (*DrawingContext)(nil)
	_ gextras.Nativer  = (*DrawingContext)(nil)
)

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

	_ret = (*cairo.Context)(unsafe.Pointer(_cret))

	return _ret
}

// Clip retrieves a copy of the clip region used when creating the @context.
func (context *DrawingContext) Clip() *cairo.Region {
	var _arg0 *C.GdkDrawingContext // out
	var _cret *C.cairo_region_t    // in

	_arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drawing_context_get_clip(_arg0)

	var _region *cairo.Region // out

	_region = (*cairo.Region)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_region, func(v *cairo.Region) {
		C.free(unsafe.Pointer(v))
	})

	return _region
}

// Window retrieves the window that created the drawing @context.
func (context *DrawingContext) Window() *Window {
	var _arg0 *C.GdkDrawingContext // out
	var _cret *C.GdkWindow         // in

	_arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drawing_context_get_window(_arg0)

	var _window *Window // out

	_window = wrapWindow(externglib.Take(unsafe.Pointer(_cret)))

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
