// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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

// PixbufLoader: the GdkPixbufLoader struct contains only private fields.
type PixbufLoader interface {
	gextras.Objector

	// Close informs a pixbuf loader that no further writes with
	// gdk_pixbuf_loader_write() will occur, so that it can free its internal
	// loading structures. Also, tries to parse any data that hasn't yet been
	// parsed; if the remaining data is partial or corrupt, an error will be
	// returned. If false is returned, @error will be set to an error from the
	// K_PIXBUF_ERROR or FILE_ERROR domains. If you're just cancelling a load
	// rather than expecting it to be finished, passing nil for @error to ignore
	// it is reasonable.
	//
	// Remember that this does not unref the loader, so if you plan not to use
	// it anymore, please g_object_unref() it.
	Close() error
	// Animation queries the PixbufAnimation that a pixbuf loader is currently
	// creating. In general it only makes sense to call this function after the
	// "area-prepared" signal has been emitted by the loader. If the loader
	// doesn't have enough bytes yet (hasn't emitted the "area-prepared" signal)
	// this function will return nil.
	Animation() PixbufAnimation
	// Format obtains the available information about the format of the
	// currently loading image file.
	Format() *PixbufFormat
	// Pixbuf queries the Pixbuf that a pixbuf loader is currently creating. In
	// general it only makes sense to call this function after the
	// "area-prepared" signal has been emitted by the loader; this means that
	// enough data has been read to know the size of the image that will be
	// allocated. If the loader has not received enough data via
	// gdk_pixbuf_loader_write(), then this function returns nil. The returned
	// pixbuf will be the same in all future calls to the loader, so simply
	// calling g_object_ref() should be sufficient to continue using it.
	// Additionally, if the loader is an animation, it will return the "static
	// image" of the animation (see gdk_pixbuf_animation_get_static_image()).
	Pixbuf() Pixbuf
	// SetSize causes the image to be scaled while it is loaded. The desired
	// image size can be determined relative to the original size of the image
	// by calling gdk_pixbuf_loader_set_size() from a signal handler for the
	// ::size-prepared signal.
	//
	// Attempts to set the desired image size are ignored after the emission of
	// the ::size-prepared signal.
	SetSize(width int, height int)
	// Write: this will cause a pixbuf loader to parse the next @count bytes of
	// an image. It will return true if the data was loaded successfully, and
	// false if an error occurred. In the latter case, the loader will be
	// closed, and will not accept further writes. If false is returned, @error
	// will be set to an error from the K_PIXBUF_ERROR or FILE_ERROR domains.
	Write(buf []byte) error
	// WriteBytes: this will cause a pixbuf loader to parse a buffer inside a
	// #GBytes for an image. It will return true if the data was loaded
	// successfully, and false if an error occurred. In the latter case, the
	// loader will be closed, and will not accept further writes. If false is
	// returned, @error will be set to an error from the K_PIXBUF_ERROR or
	// FILE_ERROR domains.
	//
	// See also: gdk_pixbuf_loader_write()
	WriteBytes(buffer *glib.Bytes) error
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

// NewPixbufLoader constructs a class PixbufLoader.
func NewPixbufLoader() PixbufLoader {
	var cret C.GdkPixbufLoader
	var ret1 PixbufLoader

	cret = C.gdk_pixbuf_loader_new()

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(PixbufLoader)

	return ret1
}

// NewPixbufLoaderWithMIMEType constructs a class PixbufLoader.
func NewPixbufLoaderWithMIMEType(mimeType string) (pixbufLoader PixbufLoader, err error) {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(arg1))

	var errout *C.GError
	var goerr error
	var cret C.GdkPixbufLoader
	var ret2 PixbufLoader

	cret = C.gdk_pixbuf_loader_new_with_mime_type(mimeType, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(PixbufLoader)

	return goerr, ret2
}

// NewPixbufLoaderWithType constructs a class PixbufLoader.
func NewPixbufLoaderWithType(imageType string) (pixbufLoader PixbufLoader, err error) {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(imageType))
	defer C.free(unsafe.Pointer(arg1))

	var errout *C.GError
	var goerr error
	var cret C.GdkPixbufLoader
	var ret2 PixbufLoader

	cret = C.gdk_pixbuf_loader_new_with_type(imageType, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(PixbufLoader)

	return goerr, ret2
}

// Close informs a pixbuf loader that no further writes with
// gdk_pixbuf_loader_write() will occur, so that it can free its internal
// loading structures. Also, tries to parse any data that hasn't yet been
// parsed; if the remaining data is partial or corrupt, an error will be
// returned. If false is returned, @error will be set to an error from the
// K_PIXBUF_ERROR or FILE_ERROR domains. If you're just cancelling a load
// rather than expecting it to be finished, passing nil for @error to ignore
// it is reasonable.
//
// Remember that this does not unref the loader, so if you plan not to use
// it anymore, please g_object_unref() it.
func (l pixbufLoader) Close() error {
	var arg0 *C.GdkPixbufLoader

	arg0 = (*C.GdkPixbufLoader)(unsafe.Pointer(l.Native()))

	var errout *C.GError
	var goerr error

	C.gdk_pixbuf_loader_close(arg0, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))

	return goerr
}

// Animation queries the PixbufAnimation that a pixbuf loader is currently
// creating. In general it only makes sense to call this function after the
// "area-prepared" signal has been emitted by the loader. If the loader
// doesn't have enough bytes yet (hasn't emitted the "area-prepared" signal)
// this function will return nil.
func (l pixbufLoader) Animation() PixbufAnimation {
	var arg0 *C.GdkPixbufLoader

	arg0 = (*C.GdkPixbufLoader)(unsafe.Pointer(l.Native()))

	var cret *C.GdkPixbufAnimation
	var ret1 PixbufAnimation

	cret = C.gdk_pixbuf_loader_get_animation(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(PixbufAnimation)

	return ret1
}

// Format obtains the available information about the format of the
// currently loading image file.
func (l pixbufLoader) Format() *PixbufFormat {
	var arg0 *C.GdkPixbufLoader

	arg0 = (*C.GdkPixbufLoader)(unsafe.Pointer(l.Native()))

	var cret *C.GdkPixbufFormat
	var ret1 *PixbufFormat

	cret = C.gdk_pixbuf_loader_get_format(arg0)

	ret1 = WrapPixbufFormat(unsafe.Pointer(cret))

	return ret1
}

// Pixbuf queries the Pixbuf that a pixbuf loader is currently creating. In
// general it only makes sense to call this function after the
// "area-prepared" signal has been emitted by the loader; this means that
// enough data has been read to know the size of the image that will be
// allocated. If the loader has not received enough data via
// gdk_pixbuf_loader_write(), then this function returns nil. The returned
// pixbuf will be the same in all future calls to the loader, so simply
// calling g_object_ref() should be sufficient to continue using it.
// Additionally, if the loader is an animation, it will return the "static
// image" of the animation (see gdk_pixbuf_animation_get_static_image()).
func (l pixbufLoader) Pixbuf() Pixbuf {
	var arg0 *C.GdkPixbufLoader

	arg0 = (*C.GdkPixbufLoader)(unsafe.Pointer(l.Native()))

	var cret *C.GdkPixbuf
	var ret1 Pixbuf

	cret = C.gdk_pixbuf_loader_get_pixbuf(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Pixbuf)

	return ret1
}

// SetSize causes the image to be scaled while it is loaded. The desired
// image size can be determined relative to the original size of the image
// by calling gdk_pixbuf_loader_set_size() from a signal handler for the
// ::size-prepared signal.
//
// Attempts to set the desired image size are ignored after the emission of
// the ::size-prepared signal.
func (l pixbufLoader) SetSize(width int, height int) {
	var arg0 *C.GdkPixbufLoader
	var arg1 C.int
	var arg2 C.int

	arg0 = (*C.GdkPixbufLoader)(unsafe.Pointer(l.Native()))
	arg1 = C.int(width)
	arg2 = C.int(height)

	C.gdk_pixbuf_loader_set_size(arg0, width, height)
}

// Write: this will cause a pixbuf loader to parse the next @count bytes of
// an image. It will return true if the data was loaded successfully, and
// false if an error occurred. In the latter case, the loader will be
// closed, and will not accept further writes. If false is returned, @error
// will be set to an error from the K_PIXBUF_ERROR or FILE_ERROR domains.
func (l pixbufLoader) Write(buf []byte) error {
	var arg0 *C.GdkPixbufLoader

	arg0 = (*C.GdkPixbufLoader)(unsafe.Pointer(l.Native()))

	var errout *C.GError
	var goerr error

	C.gdk_pixbuf_loader_write(arg0, buf, count, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))

	return goerr
}

// WriteBytes: this will cause a pixbuf loader to parse a buffer inside a
// #GBytes for an image. It will return true if the data was loaded
// successfully, and false if an error occurred. In the latter case, the
// loader will be closed, and will not accept further writes. If false is
// returned, @error will be set to an error from the K_PIXBUF_ERROR or
// FILE_ERROR domains.
//
// See also: gdk_pixbuf_loader_write()
func (l pixbufLoader) WriteBytes(buffer *glib.Bytes) error {
	var arg0 *C.GdkPixbufLoader
	var arg1 *C.GBytes

	arg0 = (*C.GdkPixbufLoader)(unsafe.Pointer(l.Native()))
	arg1 = (*C.GBytes)(unsafe.Pointer(buffer.Native()))

	var errout *C.GError
	var goerr error

	C.gdk_pixbuf_loader_write_bytes(arg0, buffer, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))

	return goerr
}
