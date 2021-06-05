// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_file_input_stream_get_type()), F: marshalFileInputStream},
	})
}

// FileInputStream: GFileInputStream provides input streams that take their
// content from a file.
//
// GFileInputStream implements #GSeekable, which allows the input stream to jump
// to arbitrary positions in the file, provided the filesystem of the file
// allows it. To find the position of a file input stream, use
// g_seekable_tell(). To find out if a file input stream supports seeking, use
// g_seekable_can_seek(). To position a file input stream, use
// g_seekable_seek().
type FileInputStream interface {
	InputStream
	Seekable

	// QueryInfo queries a file input stream the given @attributes. This
	// function blocks while querying the stream. For the asynchronous
	// (non-blocking) version of this function, see
	// g_file_input_stream_query_info_async(). While the stream is blocked, the
	// stream will set the pending flag internally, and any other operations on
	// the stream will fail with G_IO_ERROR_PENDING.
	QueryInfo(attributes string, cancellable Cancellable) (fileInfo FileInfo, err error)
	// QueryInfoAsync queries the stream information asynchronously. When the
	// operation is finished @callback will be called. You can then call
	// g_file_input_stream_query_info_finish() to get the result of the
	// operation.
	//
	// For the synchronous version of this function, see
	// g_file_input_stream_query_info().
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be set
	QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// QueryInfoFinish finishes an asynchronous info query operation.
	QueryInfoFinish(result AsyncResult) (fileInfo FileInfo, err error)
}

// fileInputStream implements the FileInputStream interface.
type fileInputStream struct {
	InputStream
	Seekable
}

var _ FileInputStream = (*fileInputStream)(nil)

// WrapFileInputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileInputStream(obj *externglib.Object) FileInputStream {
	return FileInputStream{
		InputStream: WrapInputStream(obj),
		Seekable:    WrapSeekable(obj),
	}
}

func marshalFileInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileInputStream(obj), nil
}

// QueryInfo queries a file input stream the given @attributes. This
// function blocks while querying the stream. For the asynchronous
// (non-blocking) version of this function, see
// g_file_input_stream_query_info_async(). While the stream is blocked, the
// stream will set the pending flag internally, and any other operations on
// the stream will fail with G_IO_ERROR_PENDING.
func (s fileInputStream) QueryInfo(attributes string, cancellable Cancellable) (fileInfo FileInfo, err error) {
	var arg0 *C.GFileInputStream
	var arg1 *C.char
	var arg2 *C.GCancellable

	arg0 = (*C.GFileInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var errout *C.GError
	var goerr error
	var cret *C.GFileInfo
	var ret2 FileInfo

	cret = C.g_file_input_stream_query_info(arg0, attributes, cancellable, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(FileInfo)

	return goerr, ret2
}

// QueryInfoAsync queries the stream information asynchronously. When the
// operation is finished @callback will be called. You can then call
// g_file_input_stream_query_info_finish() to get the result of the
// operation.
//
// For the synchronous version of this function, see
// g_file_input_stream_query_info().
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be set
func (s fileInputStream) QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GFileInputStream

	arg0 = (*C.GFileInputStream)(unsafe.Pointer(s.Native()))

	C.g_file_input_stream_query_info_async(arg0, attributes, ioPriority, cancellable, callback, userData)
}

// QueryInfoFinish finishes an asynchronous info query operation.
func (s fileInputStream) QueryInfoFinish(result AsyncResult) (fileInfo FileInfo, err error) {
	var arg0 *C.GFileInputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GFileInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var goerr error
	var cret *C.GFileInfo
	var ret2 FileInfo

	cret = C.g_file_input_stream_query_info_finish(arg0, result, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(FileInfo)

	return goerr, ret2
}
