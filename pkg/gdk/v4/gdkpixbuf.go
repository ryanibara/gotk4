// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// PixbufGetFromSurface transfers image data from a `cairo_surface_t` and
// converts it to a `GdkPixbuf`.
//
// This allows you to efficiently read individual pixels from cairo surfaces.
//
// This function will create an RGB pixbuf with 8 bits per channel. The pixbuf
// will contain an alpha channel if the @surface contains one.
func PixbufGetFromSurface(surface *cairo.Surface, srcX int, srcY int, width int, height int) gdkpixbuf.Pixbuf {
	var arg1 *C.cairo_surface_t
	var arg2 C.int
	var arg3 C.int
	var arg4 C.int
	var arg5 C.int

	arg1 = (*C.cairo_surface_t)(surface.Native())
	arg2 = C.int(srcX)
	arg3 = C.int(srcY)
	arg4 = C.int(width)
	arg5 = C.int(height)

	ret := C.gdk_pixbuf_get_from_surface(arg1, arg2, arg3, arg4, arg5)

	var ret0 gdkpixbuf.Pixbuf

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(gdkpixbuf.Pixbuf)

	return ret0
}

// PixbufGetFromTexture creates a new pixbuf from @texture.
//
// This should generally not be used in newly written code as later stages will
// almost certainly convert the pixbuf back into a texture to draw it on screen.
func PixbufGetFromTexture(texture Texture) gdkpixbuf.Pixbuf {
	var arg1 *C.GdkTexture

	arg1 = (*C.GdkTexture)(texture.Native())

	ret := C.gdk_pixbuf_get_from_texture(arg1)

	var ret0 gdkpixbuf.Pixbuf

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(gdkpixbuf.Pixbuf)

	return ret0
}
