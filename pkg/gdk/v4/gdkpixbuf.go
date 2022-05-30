// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// PixbufGetFromTexture creates a new pixbuf from texture.
//
// This should generally not be used in newly written code as later stages will
// almost certainly convert the pixbuf back into a texture to draw it on screen.
//
// The function takes the following parameters:
//
//    - texture: GdkTexture.
//
// The function returns the following values:
//
//    - pixbuf (optional): new Pixbuf or NULL in case of an error.
//
func PixbufGetFromTexture(texture Texturer) *gdkpixbuf.Pixbuf {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(texture).Native()))
	*(*Texturer)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "pixbuf_get_from_texture").Invoke(args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(texture)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if _cret != nil {
		{
			obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
			_pixbuf = &gdkpixbuf.Pixbuf{
				Object: obj,
				LoadableIcon: gio.LoadableIcon{
					Icon: gio.Icon{
						Object: obj,
					},
				},
			}
		}
	}

	return _pixbuf
}
