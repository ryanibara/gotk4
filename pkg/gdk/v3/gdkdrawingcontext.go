// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/cairo"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_drawing_context_get_type()), F: marshalDrawingContext},
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
type DrawingContext interface {
	gextras.Objector

	// CairoContext retrieves a Cairo context to be used to draw on the Window
	// that created the DrawingContext.
	//
	// The returned context is guaranteed to be valid as long as the
	// DrawingContext is valid, that is between a call to
	// gdk_window_begin_draw_frame() and gdk_window_end_draw_frame().
	CairoContext() *cairo.Context
	// Clip retrieves a copy of the clip region used when creating the @context.
	Clip() *cairo.Region
	// Window retrieves the window that created the drawing @context.
	Window() Window
	// IsValid checks whether the given DrawingContext is valid.
	IsValid() bool
}

// drawingContext implements the DrawingContext interface.
type drawingContext struct {
	gextras.Objector
}

var _ DrawingContext = (*drawingContext)(nil)

// WrapDrawingContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapDrawingContext(obj *externglib.Object) DrawingContext {
	return DrawingContext{
		Objector: obj,
	}
}

func marshalDrawingContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDrawingContext(obj), nil
}

// CairoContext retrieves a Cairo context to be used to draw on the Window
// that created the DrawingContext.
//
// The returned context is guaranteed to be valid as long as the
// DrawingContext is valid, that is between a call to
// gdk_window_begin_draw_frame() and gdk_window_end_draw_frame().
func (c drawingContext) CairoContext() *cairo.Context {
	var arg0 *C.GdkDrawingContext

	arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(c.Native()))

	var cret *C.cairo_t
	var ret1 *cairo.Context

	cret = C.gdk_drawing_context_get_cairo_context(arg0)

	ret1 = cairo.WrapContext(unsafe.Pointer(cret))

	return ret1
}

// Clip retrieves a copy of the clip region used when creating the @context.
func (c drawingContext) Clip() *cairo.Region {
	var arg0 *C.GdkDrawingContext

	arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(c.Native()))

	var cret *C.cairo_region_t
	var ret1 *cairo.Region

	cret = C.gdk_drawing_context_get_clip(arg0)

	ret1 = cairo.WrapRegion(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *cairo.Region) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Window retrieves the window that created the drawing @context.
func (c drawingContext) Window() Window {
	var arg0 *C.GdkDrawingContext

	arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(c.Native()))

	var cret *C.GdkWindow
	var ret1 Window

	cret = C.gdk_drawing_context_get_window(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Window)

	return ret1
}

// IsValid checks whether the given DrawingContext is valid.
func (c drawingContext) IsValid() bool {
	var arg0 *C.GdkDrawingContext

	arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gdk_drawing_context_is_valid(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}
