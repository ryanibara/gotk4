// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <gdk/gdk.h>
import "C"

// CairoCreate creates a Cairo context for drawing to @window.
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
func CairoCreate(window Window) *cairo.Context {
	var arg1 *C.GdkWindow

	arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))

	var cret *C.cairo_t
	var ret1 *cairo.Context

	cret = C.gdk_cairo_create(window)

	ret1 = cairo.WrapContext(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *cairo.Context) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// CairoDrawFromGL: this is the main way to draw GL content in GTK+. It takes a
// render buffer ID (@source_type == RENDERBUFFER) or a texture id (@source_type
// == TEXTURE) and draws it onto @cr with an OVER operation, respecting the
// current clip. The top left corner of the rectangle specified by @x, @y,
// @width and @height will be drawn at the current (0,0) position of the
// cairo_t.
//
// This will work for *all* cairo_t, as long as @window is realized, but the
// fallback implementation that reads back the pixels from the buffer may be
// used in the general case. In the case of direct drawing to a window with no
// special effects applied to @cr it will however use a more efficient approach.
//
// For RENDERBUFFER the code will always fall back to software for buffers with
// alpha components, so make sure you use TEXTURE if using alpha.
//
// Calling this may change the current GL context.
func CairoDrawFromGL(cr *cairo.Context, window Window, source int, sourceType int, bufferScale int, x int, y int, width int, height int) {
	var arg1 *C.cairo_t
	var arg2 *C.GdkWindow
	var arg3 C.int
	var arg4 C.int
	var arg5 C.int
	var arg6 C.int
	var arg7 C.int
	var arg8 C.int
	var arg9 C.int

	arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	arg2 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	arg3 = C.int(source)
	arg4 = C.int(sourceType)
	arg5 = C.int(bufferScale)
	arg6 = C.int(x)
	arg7 = C.int(y)
	arg8 = C.int(width)
	arg9 = C.int(height)

	C.gdk_cairo_draw_from_gl(cr, window, source, sourceType, bufferScale, x, y, width, height)
}

// CairoGetClipRectangle: this is a convenience function around
// cairo_clip_extents(). It rounds the clip extents to integer coordinates and
// returns a boolean indicating if a clip area exists.
func CairoGetClipRectangle(cr *cairo.Context) (rect Rectangle, ok bool) {
	var arg1 *C.cairo_t

	arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))

	var arg2 C.GdkRectangle
	var ret2 *Rectangle
	var cret C.gboolean
	var ret2 bool

	cret = C.gdk_cairo_get_clip_rectangle(cr, &arg2)

	*ret2 = WrapRectangle(unsafe.Pointer(arg2))
	ret2 = C.bool(cret) != C.false

	return ret2, ret2
}

// CairoGetDrawingContext retrieves the DrawingContext that created the Cairo
// context @cr.
func CairoGetDrawingContext(cr *cairo.Context) DrawingContext {
	var arg1 *C.cairo_t

	arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))

	var cret *C.GdkDrawingContext
	var ret1 DrawingContext

	cret = C.gdk_cairo_get_drawing_context(cr)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(DrawingContext)

	return ret1
}

// CairoRectangle adds the given rectangle to the current path of @cr.
func CairoRectangle(cr *cairo.Context, rectangle *Rectangle) {
	var arg1 *C.cairo_t
	var arg2 *C.GdkRectangle

	arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	arg2 = (*C.GdkRectangle)(unsafe.Pointer(rectangle.Native()))

	C.gdk_cairo_rectangle(cr, rectangle)
}

// CairoRegion adds the given region to the current path of @cr.
func CairoRegion(cr *cairo.Context, region *cairo.Region) {
	var arg1 *C.cairo_t
	var arg2 *C.cairo_region_t

	arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	arg2 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))

	C.gdk_cairo_region(cr, region)
}

// CairoRegionCreateFromSurface creates region that describes covers the area
// where the given @surface is more than 50% opaque.
//
// This function takes into account device offsets that might be set with
// cairo_surface_set_device_offset().
func CairoRegionCreateFromSurface(surface *cairo.Surface) *cairo.Region {
	var arg1 *C.cairo_surface_t

	arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))

	var cret *C.cairo_region_t
	var ret1 *cairo.Region

	cret = C.gdk_cairo_region_create_from_surface(surface)

	ret1 = cairo.WrapRegion(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *cairo.Region) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// CairoSetSourceColor sets the specified Color as the source color of @cr.
func CairoSetSourceColor(cr *cairo.Context, color *Color) {
	var arg1 *C.cairo_t
	var arg2 *C.GdkColor

	arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	arg2 = (*C.GdkColor)(unsafe.Pointer(color.Native()))

	C.gdk_cairo_set_source_color(cr, color)
}

// CairoSetSourcePixbuf sets the given pixbuf as the source pattern for @cr.
//
// The pattern has an extend mode of CAIRO_EXTEND_NONE and is aligned so that
// the origin of @pixbuf is @pixbuf_x, @pixbuf_y.
func CairoSetSourcePixbuf(cr *cairo.Context, pixbuf gdkpixbuf.Pixbuf, pixbufX float64, pixbufY float64) {
	var arg1 *C.cairo_t
	var arg2 *C.GdkPixbuf
	var arg3 C.gdouble
	var arg4 C.gdouble

	arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	arg3 = C.gdouble(pixbufX)
	arg4 = C.gdouble(pixbufY)

	C.gdk_cairo_set_source_pixbuf(cr, pixbuf, pixbufX, pixbufY)
}

// CairoSetSourceRGBA sets the specified RGBA as the source color of @cr.
func CairoSetSourceRGBA(cr *cairo.Context, rgba *RGBA) {
	var arg1 *C.cairo_t
	var arg2 *C.GdkRGBA

	arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	arg2 = (*C.GdkRGBA)(unsafe.Pointer(rgba.Native()))

	C.gdk_cairo_set_source_rgba(cr, rgba)
}

// CairoSetSourceWindow sets the given window as the source pattern for @cr.
//
// The pattern has an extend mode of CAIRO_EXTEND_NONE and is aligned so that
// the origin of @window is @x, @y. The window contains all its subwindows when
// rendering.
//
// Note that the contents of @window are undefined outside of the visible part
// of @window, so use this function with care.
func CairoSetSourceWindow(cr *cairo.Context, window Window, x float64, y float64) {
	var arg1 *C.cairo_t
	var arg2 *C.GdkWindow
	var arg3 C.gdouble
	var arg4 C.gdouble

	arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	arg2 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	arg3 = C.gdouble(x)
	arg4 = C.gdouble(y)

	C.gdk_cairo_set_source_window(cr, window, x, y)
}

// CairoSurfaceCreateFromPixbuf creates an image surface with the same contents
// as the pixbuf.
func CairoSurfaceCreateFromPixbuf(pixbuf gdkpixbuf.Pixbuf, scale int, forWindow Window) *cairo.Surface {
	var arg1 *C.GdkPixbuf
	var arg2 C.int
	var arg3 *C.GdkWindow

	arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	arg2 = C.int(scale)
	arg3 = (*C.GdkWindow)(unsafe.Pointer(forWindow.Native()))

	var cret *C.cairo_surface_t
	var ret1 *cairo.Surface

	cret = C.gdk_cairo_surface_create_from_pixbuf(pixbuf, scale, forWindow)

	ret1 = cairo.WrapSurface(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *cairo.Surface) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}
