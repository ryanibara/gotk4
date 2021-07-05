// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
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
	gextras.Objector

	// AsPaintable casts the class to the Paintable interface.
	AsPaintable() Paintable

	// Height returns the height of the @texture, in pixels.
	Height() int
	// Width returns the width of @texture, in pixels.
	Width() int
	// SaveToPngTexture: store the given @texture to the @filename as a PNG
	// file.
	//
	// This is a utility function intended for debugging and testing. If you
	// want more control over formats, proper error handling or want to store to
	// a `GFile` or other location, you might want to look into using the
	// gdk-pixbuf library.
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

// NewTextureForPixbuf creates a new texture object representing the
// `GdkPixbuf`.
func NewTextureForPixbuf(pixbuf gdkpixbuf.Pixbuf) Texture {
	var _arg1 *C.GdkPixbuf  // out
	var _cret *C.GdkTexture // in

	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	_cret = C.gdk_texture_new_for_pixbuf(_arg1)

	var _texture Texture // out

	_texture = WrapTexture(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _texture
}

// NewTextureFromResource creates a new texture by loading an image from a
// resource.
//
// The file format is detected automatically. The supported formats are PNG and
// JPEG, though more formats might be available.
//
// It is a fatal error if @resource_path does not specify a valid image resource
// and the program will abort if that happens. If you are unsure about the
// validity of a resource, use [ctor@Gdk.Texture.new_from_file] to load it.
func NewTextureFromResource(resourcePath string) Texture {
	var _arg1 *C.char       // out
	var _cret *C.GdkTexture // in

	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_texture_new_from_resource(_arg1)

	var _texture Texture // out

	_texture = WrapTexture(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

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

func (t texture) AsPaintable() Paintable {
	return WrapPaintable(gextras.InternObject(t))
}
