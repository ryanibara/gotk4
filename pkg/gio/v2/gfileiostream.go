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
		{T: externglib.Type(C.g_file_io_stream_get_type()), F: marshalFileIOStream},
	})
}

// FileIOStream: GFileIOStream provides io streams that both read and write to
// the same file handle.
//
// GFileIOStream implements #GSeekable, which allows the io stream to jump to
// arbitrary positions in the file and to truncate the file, provided the
// filesystem of the file supports these operations.
//
// To find the position of a file io stream, use g_seekable_tell().
//
// To find out if a file io stream supports seeking, use g_seekable_can_seek().
// To position a file io stream, use g_seekable_seek(). To find out if a file io
// stream supports truncating, use g_seekable_can_truncate(). To truncate a file
// io stream, use g_seekable_truncate().
//
// The default implementation of all the IOStream operations and the
// implementation of #GSeekable just call into the same operations on the output
// stream.
type FileIOStream interface {
	IOStream
	Seekable

	// Etag gets the entity tag for the file when it has been written. This must
	// be called after the stream has been written and closed, as the etag can
	// change while writing.
	Etag() string
	// QueryInfo queries a file io stream for the given @attributes. This
	// function blocks while querying the stream. For the asynchronous version
	// of this function, see g_file_io_stream_query_info_async(). While the
	// stream is blocked, the stream will set the pending flag internally, and
	// any other operations on the stream will fail with G_IO_ERROR_PENDING.
	//
	// Can fail if the stream was already closed (with @error being set to
	// G_IO_ERROR_CLOSED), the stream has pending operations (with @error being
	// set to G_IO_ERROR_PENDING), or if querying info is not supported for the
	// stream's interface (with @error being set to G_IO_ERROR_NOT_SUPPORTED). I
	// all cases of failure, nil will be returned.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be set, and nil will
	// be returned.
	QueryInfo(attributes string, cancellable Cancellable) (fileInfo FileInfo, err error)
	// QueryInfoAsync: asynchronously queries the @stream for a Info. When
	// completed, @callback will be called with a Result which can be used to
	// finish the operation with g_file_io_stream_query_info_finish().
	//
	// For the synchronous version of this function, see
	// g_file_io_stream_query_info().
	QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// QueryInfoFinish finalizes the asynchronous query started by
	// g_file_io_stream_query_info_async().
	QueryInfoFinish(result AsyncResult) (fileInfo FileInfo, err error)
}

// fileIOStream implements the FileIOStream interface.
type fileIOStream struct {
	IOStream
	Seekable
}

var _ FileIOStream = (*fileIOStream)(nil)

// WrapFileIOStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileIOStream(obj *externglib.Object) FileIOStream {
	return FileIOStream{
		IOStream: WrapIOStream(obj),
		Seekable: WrapSeekable(obj),
	}
}

func marshalFileIOStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileIOStream(obj), nil
}

// Etag gets the entity tag for the file when it has been written. This must
// be called after the stream has been written and closed, as the etag can
// change while writing.
func (s fileIOStream) Etag() string {
	var arg0 *C.GFileIOStream

	arg0 = (*C.GFileIOStream)(unsafe.Pointer(s.Native()))

	var cret *C.char
	var ret1 string

	cret = C.g_file_io_stream_get_etag(arg0)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// QueryInfo queries a file io stream for the given @attributes. This
// function blocks while querying the stream. For the asynchronous version
// of this function, see g_file_io_stream_query_info_async(). While the
// stream is blocked, the stream will set the pending flag internally, and
// any other operations on the stream will fail with G_IO_ERROR_PENDING.
//
// Can fail if the stream was already closed (with @error being set to
// G_IO_ERROR_CLOSED), the stream has pending operations (with @error being
// set to G_IO_ERROR_PENDING), or if querying info is not supported for the
// stream's interface (with @error being set to G_IO_ERROR_NOT_SUPPORTED). I
// all cases of failure, nil will be returned.
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be set, and nil will
// be returned.
func (s fileIOStream) QueryInfo(attributes string, cancellable Cancellable) (fileInfo FileInfo, err error) {
	var arg0 *C.GFileIOStream
	var arg1 *C.char
	var arg2 *C.GCancellable

	arg0 = (*C.GFileIOStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var errout *C.GError
	var goerr error
	var cret *C.GFileInfo
	var ret2 FileInfo

	cret = C.g_file_io_stream_query_info(arg0, attributes, cancellable, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(FileInfo)

	return goerr, ret2
}

// QueryInfoAsync: asynchronously queries the @stream for a Info. When
// completed, @callback will be called with a Result which can be used to
// finish the operation with g_file_io_stream_query_info_finish().
//
// For the synchronous version of this function, see
// g_file_io_stream_query_info().
func (s fileIOStream) QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GFileIOStream

	arg0 = (*C.GFileIOStream)(unsafe.Pointer(s.Native()))

	C.g_file_io_stream_query_info_async(arg0, attributes, ioPriority, cancellable, callback, userData)
}

// QueryInfoFinish finalizes the asynchronous query started by
// g_file_io_stream_query_info_async().
func (s fileIOStream) QueryInfoFinish(result AsyncResult) (fileInfo FileInfo, err error) {
	var arg0 *C.GFileIOStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GFileIOStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var goerr error
	var cret *C.GFileInfo
	var ret2 FileInfo

	cret = C.g_file_io_stream_query_info_finish(arg0, result, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(FileInfo)

	return goerr, ret2
}
