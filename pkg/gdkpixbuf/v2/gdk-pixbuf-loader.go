// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk-pixbuf/gdk-pixbuf.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_pixbuf_loader_get_type()), F: marshalPixbufLoader},
	})
}

// PixbufLoader: incremental image loader.
//
// `GdkPixbufLoader` provides a way for applications to drive the process of
// loading an image, by letting them send the image data directly to the loader
// instead of having the loader read the data from a file. Applications can use
// this functionality instead of `gdk_pixbuf_new_from_file()` or
// `gdk_pixbuf_animation_new_from_file()` when they need to parse image data in
// small chunks. For example, it should be used when reading an image from a
// (potentially) slow network connection, or when loading an extremely large
// file.
//
// To use `GdkPixbufLoader` to load an image, create a new instance, and call
// [method@GdkPixbuf.PixbufLoader.write] to send the data to it. When done,
// [method@GdkPixbuf.PixbufLoader.close] should be called to end the stream and
// finalize everything.
//
// The loader will emit three important signals throughout the process:
//
//    - [signal@GdkPixbuf.PixbufLoader::size-prepared] will be emitted as
//      soon as the image has enough information to determine the size of
//      the image to be used. If you want to scale the image while loading
//      it, you can call [method@GdkPixbuf.PixbufLoader.set_size] in
//      response to this signal.
//    - [signal@GdkPixbuf.PixbufLoader::area-prepared] will be emitted as
//      soon as the pixbuf of the desired has been allocated. You can obtain
//      the `GdkPixbuf` instance by calling [method@GdkPixbuf.PixbufLoader.get_pixbuf].
//      If you want to use it, simply acquire a reference to it. You can
//      also call `gdk_pixbuf_loader_get_pixbuf()` later to get the same
//      pixbuf.
//    - [signal@GdkPixbuf.PixbufLoader::area-updated] will be emitted every
//      time a region is updated. This way you can update a partially
//      completed image. Note that you do not know anything about the
//      completeness of an image from the updated area. For example, in an
//      interlaced image you will need to make several passes before the
//      image is done loading.
//
//
// Loading an animation
//
// Loading an animation is almost as easy as loading an image. Once the first
// [signal@GdkPixbuf.PixbufLoader::area-prepared] signal has been emitted, you
// can call [method@GdkPixbuf.PixbufLoader.get_animation] to get the
// [class@GdkPixbuf.PixbufAnimation] instance, and then call and
// [method@GdkPixbuf.PixbufAnimation.get_iter] to get a
// [class@GdkPixbuf.PixbufAnimationIter] to retrieve the pixbuf for the desired
// time stamp.
type PixbufLoader interface {
	gextras.Objector

	// Close informs a pixbuf loader that no further writes with
	// gdk_pixbuf_loader_write() will occur, so that it can free its internal
	// loading structures.
	//
	// This function also tries to parse any data that hasn't yet been parsed;
	// if the remaining data is partial or corrupt, an error will be returned.
	//
	// If `FALSE` is returned, `error` will be set to an error from the
	// `GDK_PIXBUF_ERROR` or `G_FILE_ERROR` domains.
	//
	// If you're just cancelling a load rather than expecting it to be finished,
	// passing `NULL` for `error` to ignore it is reasonable.
	//
	// Remember that this function does not release a reference on the loader,
	// so you will need to explicitly release any reference you hold.
	Close() error
	// SetSize causes the image to be scaled while it is loaded.
	//
	// The desired image size can be determined relative to the original size of
	// the image by calling gdk_pixbuf_loader_set_size() from a signal handler
	// for the ::size-prepared signal.
	//
	// Attempts to set the desired image size are ignored after the emission of
	// the ::size-prepared signal.
	SetSize(width int, height int)
	// Write parses the next `count` bytes in the given image buffer.
	Write(buf []byte) error
}

// pixbufLoader implements the PixbufLoader interface.
type pixbufLoader struct {
	gextras.Objector
}

var _ PixbufLoader = (*pixbufLoader)(nil)

// WrapPixbufLoader wraps a GObject to the right type. It is
// primarily used internally.
func WrapPixbufLoader(obj *externglib.Object) PixbufLoader {
	return PixbufLoader{
		Objector: obj,
	}
}

func marshalPixbufLoader(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPixbufLoader(obj), nil
}

// Close informs a pixbuf loader that no further writes with
// gdk_pixbuf_loader_write() will occur, so that it can free its internal
// loading structures.
//
// This function also tries to parse any data that hasn't yet been parsed;
// if the remaining data is partial or corrupt, an error will be returned.
//
// If `FALSE` is returned, `error` will be set to an error from the
// `GDK_PIXBUF_ERROR` or `G_FILE_ERROR` domains.
//
// If you're just cancelling a load rather than expecting it to be finished,
// passing `NULL` for `error` to ignore it is reasonable.
//
// Remember that this function does not release a reference on the loader,
// so you will need to explicitly release any reference you hold.
func (l pixbufLoader) Close() error {
	var _arg0 *C.GdkPixbufLoader

	_arg0 = (*C.GdkPixbufLoader)(unsafe.Pointer(l.Native()))

	var _cerr *C.GError

	C.gdk_pixbuf_loader_close(_arg0, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetSize causes the image to be scaled while it is loaded.
//
// The desired image size can be determined relative to the original size of
// the image by calling gdk_pixbuf_loader_set_size() from a signal handler
// for the ::size-prepared signal.
//
// Attempts to set the desired image size are ignored after the emission of
// the ::size-prepared signal.
func (l pixbufLoader) SetSize(width int, height int) {
	var _arg0 *C.GdkPixbufLoader
	var _arg1 C.int
	var _arg2 C.int

	_arg0 = (*C.GdkPixbufLoader)(unsafe.Pointer(l.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	C.gdk_pixbuf_loader_set_size(_arg0, _arg1, _arg2)
}

// Write parses the next `count` bytes in the given image buffer.
func (l pixbufLoader) Write(buf []byte) error {
	var _arg0 *C.GdkPixbufLoader
	var _arg1 *C.guchar
	var _arg2 C.gsize

	_arg0 = (*C.GdkPixbufLoader)(unsafe.Pointer(l.Native()))
	_arg2 = C.gsize(len(buf))
	_arg1 = (*C.guchar)(unsafe.Pointer(&buf[0]))

	var _cerr *C.GError

	C.gdk_pixbuf_loader_write(_arg0, _arg1, _arg2, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
