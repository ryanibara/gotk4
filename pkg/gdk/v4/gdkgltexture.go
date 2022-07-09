// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeGLTexture returns the GType for the type GLTexture.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeGLTexture() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "GLTexture").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalGLTexture)
	return gtype
}

// GLTexture: gdkTexture representing a GL texture object.
type GLTexture struct {
	_ [0]func() // equal guard
	Texture
}

var (
	_ Texturer = (*GLTexture)(nil)
)

func wrapGLTexture(obj *coreglib.Object) *GLTexture {
	return &GLTexture{
		Texture: Texture{
			Object: obj,
			Paintable: Paintable{
				Object: obj,
			},
		},
	}
}

func marshalGLTexture(p uintptr) (interface{}, error) {
	return wrapGLTexture(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Release releases the GL resources held by a GdkGLTexture.
//
// The texture contents are still available via the gdk.Texture.Download()
// function, after this function has been called.
func (self *GLTexture) Release() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	girepository.MustFind("Gdk", "GLTexture").InvokeMethod("release", _args[:], nil)

	runtime.KeepAlive(self)
}
