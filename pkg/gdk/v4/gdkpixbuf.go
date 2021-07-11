// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
import "C"

// PixbufGetFromSurface transfers image data from a `cairo_surface_t` and
// converts it to a `GdkPixbuf`.
//
// This allows you to efficiently read individual pixels from cairo surfaces.
//
// This function will create an RGB pixbuf with 8 bits per channel. The pixbuf
// will contain an alpha channel if the @surface contains one.
func PixbufGetFromSurface(surface *cairo.Surface, srcX int, srcY int, width int, height int) *gdkpixbuf.Pixbuf {
	var _arg1 *C.cairo_surface_t // out
	var _arg2 C.int              // out
	var _arg3 C.int              // out
	var _arg4 C.int              // out
	var _arg5 C.int              // out
	var _cret *C.GdkPixbuf       // in

	_arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface))
	_arg2 = C.int(srcX)
	_arg3 = C.int(srcY)
	_arg4 = C.int(width)
	_arg5 = C.int(height)

	_cret = C.gdk_pixbuf_get_from_surface(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	_pixbuf = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*gdkpixbuf.Pixbuf)

	return _pixbuf
}

// PixbufGetFromTexture creates a new pixbuf from @texture.
//
// This should generally not be used in newly written code as later stages will
// almost certainly convert the pixbuf back into a texture to draw it on screen.
func PixbufGetFromTexture(texture Texturer) *gdkpixbuf.Pixbuf {
	var _arg1 *C.GdkTexture // out
	var _cret *C.GdkPixbuf  // in

	_arg1 = (*C.GdkTexture)(unsafe.Pointer((texture).(gextras.Nativer).Native()))

	_cret = C.gdk_pixbuf_get_from_texture(_arg1)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	_pixbuf = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*gdkpixbuf.Pixbuf)

	return _pixbuf
}
