// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_texture_get_type()), F: marshalTexturer},
	})
}

// Texture: GdkTexture is the basic element used to refer to pixel data.
//
// It is primarily meant for pixel data that will not change over multiple
// frames, and will be used for a long time.
//
// There are various ways to create GdkTexture objects from a GdkPixbuf, or a
// Cairo surface, or other pixel data.
//
// The ownership of the pixel data is transferred to the GdkTexture instance;
// you can only make a copy of it, via gdk.Texture.Download().
//
// GdkTexture is an immutable object: That means you cannot change anything
// about it other than increasing the reference count via g_object_ref().
type Texture struct {
	*externglib.Object

	Paintable
}

var (
	_ externglib.Objector = (*Texture)(nil)
)

// Texturer describes types inherited from class Texture.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Texturer interface {
	externglib.Objector
	baseTexture() *Texture
}

var _ Texturer = (*Texture)(nil)

func wrapTexture(obj *externglib.Object) *Texture {
	return &Texture{
		Object: obj,
		Paintable: Paintable{
			Object: obj,
		},
	}
}

func marshalTexturer(p uintptr) (interface{}, error) {
	return wrapTexture(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (texture *Texture) baseTexture() *Texture {
	return texture
}

// BaseTexture returns the underlying base object.
func BaseTexture(obj Texturer) *Texture {
	return obj.baseTexture()
}

// NewTextureForPixbuf creates a new texture object representing the GdkPixbuf.
//
// The function takes the following parameters:
//
//    - pixbuf: GdkPixbuf.
//
// The function returns the following values:
//
//    - texture: new GdkTexture.
//
func NewTextureForPixbuf(pixbuf *gdkpixbuf.Pixbuf) *Texture {
	var _arg1 *C.GdkPixbuf  // out
	var _cret *C.GdkTexture // in

	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	_cret = C.gdk_texture_new_for_pixbuf(_arg1)
	runtime.KeepAlive(pixbuf)

	var _texture *Texture // out

	_texture = wrapTexture(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _texture
}

// NewTextureFromFile creates a new texture by loading an image from a file.
//
// The file format is detected automatically. The supported formats are PNG and
// JPEG, though more formats might be available.
//
// If NULL is returned, then error will be set.
//
// The function takes the following parameters:
//
//    - file: GFile to load.
//
// The function returns the following values:
//
//    - texture: newly-created GdkTexture or NULL if an error occurred.
//
func NewTextureFromFile(file gio.Filer) (*Texture, error) {
	var _arg1 *C.GFile      // out
	var _cret *C.GdkTexture // in
	var _cerr *C.GError     // in

	_arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))

	_cret = C.gdk_texture_new_from_file(_arg1, &_cerr)
	runtime.KeepAlive(file)

	var _texture *Texture // out
	var _goerr error      // out

	_texture = wrapTexture(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _texture, _goerr
}

// NewTextureFromResource creates a new texture by loading an image from a
// resource.
//
// The file format is detected automatically. The supported formats are PNG and
// JPEG, though more formats might be available.
//
// It is a fatal error if resource_path does not specify a valid image resource
// and the program will abort if that happens. If you are unsure about the
// validity of a resource, use gdk.Texture.NewFromFile to load it.
//
// The function takes the following parameters:
//
//    - resourcePath: path of the resource file.
//
// The function returns the following values:
//
//    - texture: newly-created GdkTexture.
//
func NewTextureFromResource(resourcePath string) *Texture {
	var _arg1 *C.char       // out
	var _cret *C.GdkTexture // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_texture_new_from_resource(_arg1)
	runtime.KeepAlive(resourcePath)

	var _texture *Texture // out

	_texture = wrapTexture(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _texture
}

// Height returns the height of the texture, in pixels.
//
// The function returns the following values:
//
//    - gint: height of the GdkTexture.
//
func (texture *Texture) Height() int {
	var _arg0 *C.GdkTexture // out
	var _cret C.int         // in

	_arg0 = (*C.GdkTexture)(unsafe.Pointer(texture.Native()))

	_cret = C.gdk_texture_get_height(_arg0)
	runtime.KeepAlive(texture)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Width returns the width of texture, in pixels.
//
// The function returns the following values:
//
//    - gint: width of the GdkTexture.
//
func (texture *Texture) Width() int {
	var _arg0 *C.GdkTexture // out
	var _cret C.int         // in

	_arg0 = (*C.GdkTexture)(unsafe.Pointer(texture.Native()))

	_cret = C.gdk_texture_get_width(_arg0)
	runtime.KeepAlive(texture)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SaveToPng: store the given texture to the filename as a PNG file.
//
// This is a utility function intended for debugging and testing. If you want
// more control over formats, proper error handling or want to store to a GFile
// or other location, you might want to look into using the gdk-pixbuf library.
//
// The function takes the following parameters:
//
//    - filename to store to.
//
// The function returns the following values:
//
//    - ok: TRUE if saving succeeded, FALSE on failure.
//
func (texture *Texture) SaveToPng(filename string) bool {
	var _arg0 *C.GdkTexture // out
	var _arg1 *C.char       // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkTexture)(unsafe.Pointer(texture.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_texture_save_to_png(_arg0, _arg1)
	runtime.KeepAlive(texture)
	runtime.KeepAlive(filename)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
