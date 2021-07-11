// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
// #include <glib-object.h>
//
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_file_io_stream_get_type()), F: marshalFileIOStreamer},
	})
}

// FileIOStreamOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FileIOStreamOverrider interface {
	//
	CanSeek() bool
	//
	CanTruncate() bool
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
	QueryInfo(attributes string, cancellable Cancellabler) (*FileInfo, error)
	// QueryInfoAsync: asynchronously queries the @stream for a Info. When
	// completed, @callback will be called with a Result which can be used to
	// finish the operation with g_file_io_stream_query_info_finish().
	//
	// For the synchronous version of this function, see
	// g_file_io_stream_query_info().
	QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback)
	// QueryInfoFinish finalizes the asynchronous query started by
	// g_file_io_stream_query_info_async().
	QueryInfoFinish(result AsyncResulter) (*FileInfo, error)
	//
	Tell() int64
	//
	TruncateFn(size int64, cancellable Cancellabler) error
}

// FileIOStreamer describes FileIOStream's methods.
type FileIOStreamer interface {
	// Etag gets the entity tag for the file when it has been written.
	Etag() string
	// QueryInfo queries a file io stream for the given @attributes.
	QueryInfo(attributes string, cancellable Cancellabler) (*FileInfo, error)
	// QueryInfoAsync: asynchronously queries the @stream for a Info.
	QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback)
	// QueryInfoFinish finalizes the asynchronous query started by
	// g_file_io_stream_query_info_async().
	QueryInfoFinish(result AsyncResulter) (*FileInfo, error)
}

// FileIOStream provides io streams that both read and write to the same file
// handle.
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
type FileIOStream struct {
	IOStream

	Seekable
}

var (
	_ FileIOStreamer  = (*FileIOStream)(nil)
	_ gextras.Nativer = (*FileIOStream)(nil)
)

func wrapFileIOStream(obj *externglib.Object) FileIOStreamer {
	return &FileIOStream{
		IOStream: IOStream{
			Object: obj,
		},
		Seekable: Seekable{
			Object: obj,
		},
	}
}

func marshalFileIOStreamer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFileIOStream(obj), nil
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *FileIOStream) Native() uintptr {
	return v.IOStream.Object.Native()
}

// Etag gets the entity tag for the file when it has been written. This must be
// called after the stream has been written and closed, as the etag can change
// while writing.
func (stream *FileIOStream) Etag() string {
	var _arg0 *C.GFileIOStream // out
	var _cret *C.char          // in

	_arg0 = (*C.GFileIOStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_file_io_stream_get_etag(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// QueryInfo queries a file io stream for the given @attributes. This function
// blocks while querying the stream. For the asynchronous version of this
// function, see g_file_io_stream_query_info_async(). While the stream is
// blocked, the stream will set the pending flag internally, and any other
// operations on the stream will fail with G_IO_ERROR_PENDING.
//
// Can fail if the stream was already closed (with @error being set to
// G_IO_ERROR_CLOSED), the stream has pending operations (with @error being set
// to G_IO_ERROR_PENDING), or if querying info is not supported for the stream's
// interface (with @error being set to G_IO_ERROR_NOT_SUPPORTED). I all cases of
// failure, nil will be returned.
//
// If @cancellable is not nil, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be set, and nil will be returned.
func (stream *FileIOStream) QueryInfo(attributes string, cancellable Cancellabler) (*FileInfo, error) {
	var _arg0 *C.GFileIOStream // out
	var _arg1 *C.char          // out
	var _arg2 *C.GCancellable  // out
	var _cret *C.GFileInfo     // in
	var _cerr *C.GError        // in

	_arg0 = (*C.GFileIOStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(attributes)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_file_io_stream_query_info(_arg0, _arg1, _arg2, &_cerr)

	var _fileInfo *FileInfo // out
	var _goerr error        // out

	_fileInfo = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*FileInfo)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _fileInfo, _goerr
}

// QueryInfoAsync: asynchronously queries the @stream for a Info. When
// completed, @callback will be called with a Result which can be used to finish
// the operation with g_file_io_stream_query_info_finish().
//
// For the synchronous version of this function, see
// g_file_io_stream_query_info().
func (stream *FileIOStream) QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback) {
	var _arg0 *C.GFileIOStream      // out
	var _arg1 *C.char               // out
	var _arg2 C.int                 // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GFileIOStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(attributes)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(ioPriority)
	_arg3 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(gbox.Assign(callback))

	C.g_file_io_stream_query_info_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// QueryInfoFinish finalizes the asynchronous query started by
// g_file_io_stream_query_info_async().
func (stream *FileIOStream) QueryInfoFinish(result AsyncResulter) (*FileInfo, error) {
	var _arg0 *C.GFileIOStream // out
	var _arg1 *C.GAsyncResult  // out
	var _cret *C.GFileInfo     // in
	var _cerr *C.GError        // in

	_arg0 = (*C.GFileIOStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.g_file_io_stream_query_info_finish(_arg0, _arg1, &_cerr)

	var _fileInfo *FileInfo // out
	var _goerr error        // out

	_fileInfo = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*FileInfo)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _fileInfo, _goerr
}
