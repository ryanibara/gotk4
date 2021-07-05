// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
import "C"

// CairoDrawFromGL: the main way to draw GL content in GTK.
//
// It takes a render buffer ID (@source_type == RENDERBUFFER) or a texture id
// (@source_type == TEXTURE) and draws it onto @cr with an OVER operation,
// respecting the current clip. The top left corner of the rectangle specified
// by @x, @y, @width and @height will be drawn at the current (0,0) position of
// the `cairo_t`.
//
// This will work for *all* `cairo_t`, as long as @surface is realized, but the
// fallback implementation that reads back the pixels from the buffer may be
// used in the general case. In the case of direct drawing to a surface with no
// special effects applied to @cr it will however use a more efficient approach.
//
// For RENDERBUFFER the code will always fall back to software for buffers with
// alpha components, so make sure you use TEXTURE if using alpha.
//
// Calling this may change the current GL context.
func CairoDrawFromGL(cr cairo.Context, surface Surface, source int, sourceType int, bufferScale int, x int, y int, width int, height int) {
	var _arg1 *C.cairo_t    // out
	var _arg2 *C.GdkSurface // out
	var _arg3 C.int         // out
	var _arg4 C.int         // out
	var _arg5 C.int         // out
	var _arg6 C.int         // out
	var _arg7 C.int         // out
	var _arg8 C.int         // out
	var _arg9 C.int         // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg2 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	_arg3 = C.int(source)
	_arg4 = C.int(sourceType)
	_arg5 = C.int(bufferScale)
	_arg6 = C.int(x)
	_arg7 = C.int(y)
	_arg8 = C.int(width)
	_arg9 = C.int(height)

	C.gdk_cairo_draw_from_gl(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9)
}

// CairoRectangle adds the given rectangle to the current path of @cr.
func CairoRectangle(cr cairo.Context, rectangle Rectangle) {
	var _arg1 *C.cairo_t      // out
	var _arg2 *C.GdkRectangle // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg2 = (*C.GdkRectangle)(unsafe.Pointer(rectangle))

	C.gdk_cairo_rectangle(_arg1, _arg2)
}

// CairoRegion adds the given region to the current path of @cr.
func CairoRegion(cr cairo.Context, region cairo.Region) {
	var _arg1 *C.cairo_t        // out
	var _arg2 *C.cairo_region_t // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg2 = (*C.cairo_region_t)(unsafe.Pointer(region))

	C.gdk_cairo_region(_arg1, _arg2)
}

// CairoRegionCreateFromSurface creates region that covers the area where the
// given @surface is more than 50% opaque.
//
// This function takes into account device offsets that might be set with
// cairo_surface_set_device_offset().
func CairoRegionCreateFromSurface(surface cairo.Surface) cairo.Region {
	var _arg1 *C.cairo_surface_t // out
	var _cret *C.cairo_region_t  // in

	_arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface))

	_cret = C.gdk_cairo_region_create_from_surface(_arg1)

	var _region cairo.Region // out

	_region = (cairo.Region)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_region, func(v cairo.Region) {
		C.free(unsafe.Pointer(v))
	})

	return _region
}

// CairoSetSourcePixbuf sets the given pixbuf as the source pattern for @cr.
//
// The pattern has an extend mode of CAIRO_EXTEND_NONE and is aligned so that
// the origin of @pixbuf is @pixbuf_x, @pixbuf_y.
func CairoSetSourcePixbuf(cr cairo.Context, pixbuf gdkpixbuf.Pixbuf, pixbufX float64, pixbufY float64) {
	var _arg1 *C.cairo_t   // out
	var _arg2 *C.GdkPixbuf // out
	var _arg3 C.double     // out
	var _arg4 C.double     // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	_arg3 = C.double(pixbufX)
	_arg4 = C.double(pixbufY)

	C.gdk_cairo_set_source_pixbuf(_arg1, _arg2, _arg3, _arg4)
}

// CairoSetSourceRGBA sets the specified RGBA as the source color of @cr.
func CairoSetSourceRGBA(cr cairo.Context, rgba RGBA) {
	var _arg1 *C.cairo_t // out
	var _arg2 *C.GdkRGBA // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg2 = (*C.GdkRGBA)(unsafe.Pointer(rgba))

	C.gdk_cairo_set_source_rgba(_arg1, _arg2)
}
