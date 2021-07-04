// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gdk_texture_get_type()), F: marshalTexture},
	})
}

// Texture: `GdkTexture` is the basic element used to refer to pixel data.
//
// It is primarily meant for pixel data that will not change over multiple
// frames, and will be used for a long time.
//
// There are various ways to create `GdkTexture` objects from a `GdkPixbuf`, or
// a Cairo surface, or other pixel data.
//
// The ownership of the pixel data is transferred to the `GdkTexture` instance;
// you can only make a copy of it, via [method@Gdk.Texture.download].
//
// `GdkTexture` is an immutable object: That means you cannot change anything
// about it other than increasing the reference count via g_object_ref().
type Texture interface {
	Paintable

	Height() int

	Width() int

	SaveToPngTexture(filename string) bool
}

// texture implements the Texture class.
type texture struct {
	gextras.Objector
}

// WrapTexture wraps a GObject to the right type. It is
// primarily used internally.
func WrapTexture(obj *externglib.Object) Texture {
	return texture{
		Objector: obj,
	}
}

func marshalTexture(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTexture(obj), nil
}

func NewTextureForPixbuf(pixbuf gdkpixbuf.Pixbuf) Texture {
	var _arg1 *C.GdkPixbuf  // out
	var _cret *C.GdkTexture // in

	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	_cret = C.gdk_texture_new_for_pixbuf(_arg1)

	var _texture Texture // out

	_texture = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Texture)

	return _texture
}

func NewTextureFromResource(resourcePath string) Texture {
	var _arg1 *C.char       // out
	var _cret *C.GdkTexture // in

	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_texture_new_from_resource(_arg1)

	var _texture Texture // out

	_texture = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Texture)

	return _texture
}

func (t texture) Height() int {
	var _arg0 *C.GdkTexture // out
	var _cret C.int         // in

	_arg0 = (*C.GdkTexture)(unsafe.Pointer(t.Native()))

	_cret = C.gdk_texture_get_height(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (t texture) Width() int {
	var _arg0 *C.GdkTexture // out
	var _cret C.int         // in

	_arg0 = (*C.GdkTexture)(unsafe.Pointer(t.Native()))

	_cret = C.gdk_texture_get_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (t texture) SaveToPngTexture(filename string) bool {
	var _arg0 *C.GdkTexture // out
	var _arg1 *C.char       // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkTexture)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_texture_save_to_png(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p texture) ComputeConcreteSize(specifiedWidth float64, specifiedHeight float64, defaultWidth float64, defaultHeight float64) (concreteWidth float64, concreteHeight float64) {
	return WrapPaintable(gextras.InternObject(p)).ComputeConcreteSize(specifiedWidth, specifiedHeight, defaultWidth, defaultHeight)
}

func (p texture) CurrentImage() Paintable {
	return WrapPaintable(gextras.InternObject(p)).CurrentImage()
}

func (p texture) Flags() PaintableFlags {
	return WrapPaintable(gextras.InternObject(p)).Flags()
}

func (p texture) IntrinsicAspectRatio() float64 {
	return WrapPaintable(gextras.InternObject(p)).IntrinsicAspectRatio()
}

func (p texture) IntrinsicHeight() int {
	return WrapPaintable(gextras.InternObject(p)).IntrinsicHeight()
}

func (p texture) IntrinsicWidth() int {
	return WrapPaintable(gextras.InternObject(p)).IntrinsicWidth()
}

func (p texture) InvalidateContents() {
	WrapPaintable(gextras.InternObject(p)).InvalidateContents()
}

func (p texture) InvalidateSize() {
	WrapPaintable(gextras.InternObject(p)).InvalidateSize()
}

func (p texture) Snapshot(snapshot Snapshot, width float64, height float64) {
	WrapPaintable(gextras.InternObject(p)).Snapshot(snapshot, width, height)
}
