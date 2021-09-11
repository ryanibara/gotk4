// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// CairoCreate creates a Cairo context for drawing to window.
//
// Note that calling cairo_reset_clip() on the resulting #cairo_t will produce
// undefined results, so avoid it at all costs.
//
// Typically, this function is used to draw on a Window out of the paint cycle
// of the toolkit; this should be avoided, as it breaks various assumptions and
// optimizations.
//
// If you are drawing on a native Window in response to a GDK_EXPOSE event you
// should use gdk_window_begin_draw_frame() and
// gdk_drawing_context_get_cairo_context() instead. GTK will automatically do
// this for you when drawing a widget.
//
// Deprecated: Use gdk_window_begin_draw_frame() and
// gdk_drawing_context_get_cairo_context() instead.
func CairoCreate(window Windower) *cairo.Context {
	var _arg1 *C.GdkWindow // out
	var _cret *C.cairo_t   // in

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gdk_cairo_create(_arg1)
	runtime.KeepAlive(window)

	var _context *cairo.Context // out

	_context = cairo.WrapContext(uintptr(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_context, func(v *cairo.Context) {
		C.cairo_destroy((*C.cairo_t)(unsafe.Pointer(v.Native())))
	})

	return _context
}

// CairoDrawFromGL: this is the main way to draw GL content in GTK+. It takes a
// render buffer ID (source_type == RENDERBUFFER) or a texture id (source_type
// == TEXTURE) and draws it onto cr with an OVER operation, respecting the
// current clip. The top left corner of the rectangle specified by x, y, width
// and height will be drawn at the current (0,0) position of the cairo_t.
//
// This will work for *all* cairo_t, as long as window is realized, but the
// fallback implementation that reads back the pixels from the buffer may be
// used in the general case. In the case of direct drawing to a window with no
// special effects applied to cr it will however use a more efficient approach.
//
// For RENDERBUFFER the code will always fall back to software for buffers with
// alpha components, so make sure you use TEXTURE if using alpha.
//
// Calling this may change the current GL context.
func CairoDrawFromGL(cr *cairo.Context, window Windower, source int, sourceType int, bufferScale int, x int, y int, width int, height int) {
	var _arg1 *C.cairo_t   // out
	var _arg2 *C.GdkWindow // out
	var _arg3 C.int        // out
	var _arg4 C.int        // out
	var _arg5 C.int        // out
	var _arg6 C.int        // out
	var _arg7 C.int        // out
	var _arg8 C.int        // out
	var _arg9 C.int        // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	_arg3 = C.int(source)
	_arg4 = C.int(sourceType)
	_arg5 = C.int(bufferScale)
	_arg6 = C.int(x)
	_arg7 = C.int(y)
	_arg8 = C.int(width)
	_arg9 = C.int(height)

	C.gdk_cairo_draw_from_gl(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(window)
	runtime.KeepAlive(source)
	runtime.KeepAlive(sourceType)
	runtime.KeepAlive(bufferScale)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// CairoGetClipRectangle: this is a convenience function around
// cairo_clip_extents(). It rounds the clip extents to integer coordinates and
// returns a boolean indicating if a clip area exists.
func CairoGetClipRectangle(cr *cairo.Context) (Rectangle, bool) {
	var _arg1 *C.cairo_t     // out
	var _arg2 C.GdkRectangle // in
	var _cret C.gboolean     // in

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))

	_cret = C.gdk_cairo_get_clip_rectangle(_arg1, &_arg2)
	runtime.KeepAlive(cr)

	var _rect Rectangle // out
	var _ok bool        // out

	_rect = *(*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret != 0 {
		_ok = true
	}

	return _rect, _ok
}

// CairoGetDrawingContext retrieves the DrawingContext that created the Cairo
// context cr.
func CairoGetDrawingContext(cr *cairo.Context) *DrawingContext {
	var _arg1 *C.cairo_t           // out
	var _cret *C.GdkDrawingContext // in

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))

	_cret = C.gdk_cairo_get_drawing_context(_arg1)
	runtime.KeepAlive(cr)

	var _drawingContext *DrawingContext // out

	if _cret != nil {
		_drawingContext = wrapDrawingContext(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _drawingContext
}

// CairoRectangle adds the given rectangle to the current path of cr.
func CairoRectangle(cr *cairo.Context, rectangle *Rectangle) {
	var _arg1 *C.cairo_t      // out
	var _arg2 *C.GdkRectangle // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))

	C.gdk_cairo_rectangle(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(rectangle)
}

// CairoRegion adds the given region to the current path of cr.
func CairoRegion(cr *cairo.Context, region *cairo.Region) {
	var _arg1 *C.cairo_t        // out
	var _arg2 *C.cairo_region_t // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))

	C.gdk_cairo_region(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(region)
}

// CairoRegionCreateFromSurface creates region that describes covers the area
// where the given surface is more than 50% opaque.
//
// This function takes into account device offsets that might be set with
// cairo_surface_set_device_offset().
func CairoRegionCreateFromSurface(surface *cairo.Surface) *cairo.Region {
	var _arg1 *C.cairo_surface_t // out
	var _cret *C.cairo_region_t  // in

	_arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_cairo_region_create_from_surface(_arg1)
	runtime.KeepAlive(surface)

	var _region *cairo.Region // out

	{
		_pp := &struct{ p unsafe.Pointer }{unsafe.Pointer(_cret)}
		_region = (*cairo.Region)(unsafe.Pointer(_pp))
	}
	runtime.SetFinalizer(_region, func(v *cairo.Region) {
		C.cairo_region_destroy((*C.cairo_region_t)(unsafe.Pointer(v.Native())))
	})

	return _region
}

// CairoSetSourceColor sets the specified Color as the source color of cr.
//
// Deprecated: Use gdk_cairo_set_source_rgba() instead.
func CairoSetSourceColor(cr *cairo.Context, color *Color) {
	var _arg1 *C.cairo_t  // out
	var _arg2 *C.GdkColor // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.GdkColor)(gextras.StructNative(unsafe.Pointer(color)))

	C.gdk_cairo_set_source_color(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(color)
}

// CairoSetSourcePixbuf sets the given pixbuf as the source pattern for cr.
//
// The pattern has an extend mode of CAIRO_EXTEND_NONE and is aligned so that
// the origin of pixbuf is pixbuf_x, pixbuf_y.
func CairoSetSourcePixbuf(cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, pixbufX float64, pixbufY float64) {
	var _arg1 *C.cairo_t   // out
	var _arg2 *C.GdkPixbuf // out
	var _arg3 C.gdouble    // out
	var _arg4 C.gdouble    // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	_arg3 = C.gdouble(pixbufX)
	_arg4 = C.gdouble(pixbufY)

	C.gdk_cairo_set_source_pixbuf(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(pixbuf)
	runtime.KeepAlive(pixbufX)
	runtime.KeepAlive(pixbufY)
}

// CairoSetSourceRGBA sets the specified RGBA as the source color of cr.
func CairoSetSourceRGBA(cr *cairo.Context, rgba *RGBA) {
	var _arg1 *C.cairo_t // out
	var _arg2 *C.GdkRGBA // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(rgba)))

	C.gdk_cairo_set_source_rgba(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(rgba)
}

// CairoSetSourceWindow sets the given window as the source pattern for cr.
//
// The pattern has an extend mode of CAIRO_EXTEND_NONE and is aligned so that
// the origin of window is x, y. The window contains all its subwindows when
// rendering.
//
// Note that the contents of window are undefined outside of the visible part of
// window, so use this function with care.
func CairoSetSourceWindow(cr *cairo.Context, window Windower, x float64, y float64) {
	var _arg1 *C.cairo_t   // out
	var _arg2 *C.GdkWindow // out
	var _arg3 C.gdouble    // out
	var _arg4 C.gdouble    // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)

	C.gdk_cairo_set_source_window(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(window)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// CairoSurfaceCreateFromPixbuf creates an image surface with the same contents
// as the pixbuf.
func CairoSurfaceCreateFromPixbuf(pixbuf *gdkpixbuf.Pixbuf, scale int, forWindow Windower) *cairo.Surface {
	var _arg1 *C.GdkPixbuf       // out
	var _arg2 C.int              // out
	var _arg3 *C.GdkWindow       // out
	var _cret *C.cairo_surface_t // in

	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	_arg2 = C.int(scale)
	if forWindow != nil {
		_arg3 = (*C.GdkWindow)(unsafe.Pointer(forWindow.Native()))
	}

	_cret = C.gdk_cairo_surface_create_from_pixbuf(_arg1, _arg2, _arg3)
	runtime.KeepAlive(pixbuf)
	runtime.KeepAlive(scale)
	runtime.KeepAlive(forWindow)

	var _surface *cairo.Surface // out

	_surface = cairo.WrapSurface(uintptr(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_surface, func(v *cairo.Surface) {
		C.cairo_surface_destroy((*C.cairo_surface_t)(unsafe.Pointer(v.Native())))
	})

	return _surface
}
