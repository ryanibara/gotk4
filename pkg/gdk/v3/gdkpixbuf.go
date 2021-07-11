// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
import "C"

// PixbufGetFromSurface transfers image data from a #cairo_surface_t and
// converts it to an RGB(A) representation inside a Pixbuf. This allows you to
// efficiently read individual pixels from cairo surfaces. For Windows, use
// gdk_pixbuf_get_from_window() instead.
//
// This function will create an RGB pixbuf with 8 bits per channel. The pixbuf
// will contain an alpha channel if the @surface contains one.
func PixbufGetFromSurface(surface *cairo.Surface, srcX int, srcY int, width int, height int) *gdkpixbuf.Pixbuf {
	var _arg1 *C.cairo_surface_t // out
	var _arg2 C.gint             // out
	var _arg3 C.gint             // out
	var _arg4 C.gint             // out
	var _arg5 C.gint             // out
	var _cret *C.GdkPixbuf       // in

	_arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface))
	_arg2 = C.gint(srcX)
	_arg3 = C.gint(srcY)
	_arg4 = C.gint(width)
	_arg5 = C.gint(height)

	_cret = C.gdk_pixbuf_get_from_surface(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	_pixbuf = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*gdkpixbuf.Pixbuf)

	return _pixbuf
}

// PixbufGetFromWindow transfers image data from a Window and converts it to an
// RGB(A) representation inside a Pixbuf. In other words, copies image data from
// a server-side drawable to a client-side RGB(A) buffer. This allows you to
// efficiently read individual pixels on the client side.
//
// This function will create an RGB pixbuf with 8 bits per channel with the size
// specified by the @width and @height arguments scaled by the scale factor of
// @window. The pixbuf will contain an alpha channel if the @window contains
// one.
//
// If the window is off the screen, then there is no image data in the
// obscured/offscreen regions to be placed in the pixbuf. The contents of
// portions of the pixbuf corresponding to the offscreen region are undefined.
//
// If the window you’re obtaining data from is partially obscured by other
// windows, then the contents of the pixbuf areas corresponding to the obscured
// regions are undefined.
//
// If the window is not mapped (typically because it’s iconified/minimized or
// not on the current workspace), then nil will be returned.
//
// If memory can’t be allocated for the return value, nil will be returned
// instead.
//
// (In short, there are several ways this function can fail, and if it fails it
// returns nil; so check the return value.)
func PixbufGetFromWindow(window Windowwer, srcX int, srcY int, width int, height int) *gdkpixbuf.Pixbuf {
	var _arg1 *C.GdkWindow // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out
	var _arg4 C.gint       // out
	var _arg5 C.gint       // out
	var _cret *C.GdkPixbuf // in

	_arg1 = (*C.GdkWindow)(unsafe.Pointer((window).(gextras.Nativer).Native()))
	_arg2 = C.gint(srcX)
	_arg3 = C.gint(srcY)
	_arg4 = C.gint(width)
	_arg5 = C.gint(height)

	_cret = C.gdk_pixbuf_get_from_window(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	_pixbuf = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*gdkpixbuf.Pixbuf)

	return _pixbuf
}
