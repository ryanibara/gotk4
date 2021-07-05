// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_gl_texture_get_type()), F: marshalGLTexture},
	})
}

// GLTexture: gdkTexture representing a GL texture object.
type GLTexture interface {
	Texture

	// AsPaintable casts the class to the Paintable interface.
	AsPaintable() Paintable

	// ReleaseGLTexture releases the GL resources held by a `GdkGLTexture`.
	//
	// The texture contents are still available via the
	// [method@Gdk.Texture.download] function, after this function has been
	// called.
	ReleaseGLTexture()
}

// glTexture implements the GLTexture class.
type glTexture struct {
	Texture
}

// WrapGLTexture wraps a GObject to the right type. It is
// primarily used internally.
func WrapGLTexture(obj *externglib.Object) GLTexture {
	return glTexture{
		Texture: WrapTexture(obj),
	}
}

func marshalGLTexture(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGLTexture(obj), nil
}

func (s glTexture) ReleaseGLTexture() {
	var _arg0 *C.GdkGLTexture // out

	_arg0 = (*C.GdkGLTexture)(unsafe.Pointer(s.Native()))

	C.gdk_gl_texture_release(_arg0)
}

func (g glTexture) AsPaintable() Paintable {
	return WrapPaintable(gextras.InternObject(g))
}
