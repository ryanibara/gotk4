// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_buffered_input_stream_get_type()), F: marshalBufferedInputStreamer},
	})
}

// BufferedInputStreamOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type BufferedInputStreamOverrider interface {
	// Fill tries to read count bytes from the stream into the buffer. Will
	// block during this read.
	//
	// If count is zero, returns zero and does nothing. A value of count larger
	// than G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT error.
	//
	// On success, the number of bytes read into the buffer is returned. It is
	// not an error if this is not the same as the requested size, as it can
	// happen e.g. near the end of a file. Zero is returned on end of file (or
	// if count is zero), but never otherwise.
	//
	// If count is -1 then the attempted read size is equal to the number of
	// bytes that are required to fill the buffer.
	//
	// If cancellable is not NULL, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	//
	// On error -1 is returned and error is set accordingly.
	//
	// For the asynchronous, non-blocking, version of this function, see
	// g_buffered_input_stream_fill_async().
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable object, NULL to ignore.
	//    - count: number of bytes that will be read from the stream.
	//
	// The function returns the following values:
	//
	//    - gssize: number of bytes read into stream's buffer, up to count, or -1
	//      on error.
	//
	Fill(ctx context.Context, count int) (int, error)
	// FillAsync reads data into stream's buffer asynchronously, up to count
	// size. io_priority can be used to prioritize reads. For the synchronous
	// version of this function, see g_buffered_input_stream_fill().
	//
	// If count is -1 then the attempted read size is equal to the number of
	// bytes that are required to fill the buffer.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable object.
	//    - count: number of bytes that will be read from the stream.
	//    - ioPriority: [I/O priority][io-priority] of the request.
	//    - callback (optional): ReadyCallback.
	//
	FillAsync(ctx context.Context, count, ioPriority int, callback AsyncReadyCallback)
	// FillFinish finishes an asynchronous read.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	// The function returns the following values:
	//
	//    - gssize of the read stream, or -1 on an error.
	//
	FillFinish(result AsyncResulter) (int, error)
}

// BufferedInputStream: buffered input stream implements InputStream and
// provides for buffered reads.
//
// By default, InputStream's buffer size is set at 4 kilobytes.
//
// To create a buffered input stream, use g_buffered_input_stream_new(), or
// g_buffered_input_stream_new_sized() to specify the buffer's size at
// construction.
//
// To get the size of a buffer within a buffered input stream, use
// g_buffered_input_stream_get_buffer_size(). To change the size of a buffered
// input stream's buffer, use g_buffered_input_stream_set_buffer_size(). Note
// that the buffer's size cannot be reduced below the size of the data within
// the buffer.
type BufferedInputStream struct {
	_ [0]func() // equal guard
	FilterInputStream

	Seekable
}

var (
	_ FilterInputStreamer = (*BufferedInputStream)(nil)
)

func wrapBufferedInputStream(obj *externglib.Object) *BufferedInputStream {
	return &BufferedInputStream{
		FilterInputStream: FilterInputStream{
			InputStream: InputStream{
				Object: obj,
			},
		},
		Seekable: Seekable{
			Object: obj,
		},
	}
}

func marshalBufferedInputStreamer(p uintptr) (interface{}, error) {
	return wrapBufferedInputStream(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBufferedInputStream creates a new Stream from the given base_stream, with
// a buffer set to the default size (4 kilobytes).
//
// The function takes the following parameters:
//
//    - baseStream: Stream.
//
// The function returns the following values:
//
//    - bufferedInputStream for the given base_stream.
//
func NewBufferedInputStream(baseStream InputStreamer) *BufferedInputStream {
	var _arg1 *C.GInputStream // out
	var _cret *C.GInputStream // in

	_arg1 = (*C.GInputStream)(unsafe.Pointer(baseStream.Native()))

	_cret = C.g_buffered_input_stream_new(_arg1)
	runtime.KeepAlive(baseStream)

	var _bufferedInputStream *BufferedInputStream // out

	_bufferedInputStream = wrapBufferedInputStream(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _bufferedInputStream
}

// NewBufferedInputStreamSized creates a new InputStream from the given
// base_stream, with a buffer set to size.
//
// The function takes the following parameters:
//
//    - baseStream: Stream.
//    - size: #gsize.
//
// The function returns the following values:
//
//    - bufferedInputStream: Stream.
//
func NewBufferedInputStreamSized(baseStream InputStreamer, size uint) *BufferedInputStream {
	var _arg1 *C.GInputStream // out
	var _arg2 C.gsize         // out
	var _cret *C.GInputStream // in

	_arg1 = (*C.GInputStream)(unsafe.Pointer(baseStream.Native()))
	_arg2 = C.gsize(size)

	_cret = C.g_buffered_input_stream_new_sized(_arg1, _arg2)
	runtime.KeepAlive(baseStream)
	runtime.KeepAlive(size)

	var _bufferedInputStream *BufferedInputStream // out

	_bufferedInputStream = wrapBufferedInputStream(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _bufferedInputStream
}

// Fill tries to read count bytes from the stream into the buffer. Will block
// during this read.
//
// If count is zero, returns zero and does nothing. A value of count larger than
// G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes read into the buffer is returned. It is not
// an error if this is not the same as the requested size, as it can happen e.g.
// near the end of a file. Zero is returned on end of file (or if count is
// zero), but never otherwise.
//
// If count is -1 then the attempted read size is equal to the number of bytes
// that are required to fill the buffer.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned. If an operation was
// partially finished when the operation was cancelled the partial result will
// be returned, without an error.
//
// On error -1 is returned and error is set accordingly.
//
// For the asynchronous, non-blocking, version of this function, see
// g_buffered_input_stream_fill_async().
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - count: number of bytes that will be read from the stream.
//
// The function returns the following values:
//
//    - gssize: number of bytes read into stream's buffer, up to count, or -1 on
//      error.
//
func (stream *BufferedInputStream) Fill(ctx context.Context, count int) (int, error) {
	var _arg0 *C.GBufferedInputStream // out
	var _arg2 *C.GCancellable         // out
	var _arg1 C.gssize                // out
	var _cret C.gssize                // in
	var _cerr *C.GError               // in

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.gssize(count)

	_cret = C.g_buffered_input_stream_fill(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(count)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gssize, _goerr
}

// FillAsync reads data into stream's buffer asynchronously, up to count size.
// io_priority can be used to prioritize reads. For the synchronous version of
// this function, see g_buffered_input_stream_fill().
//
// If count is -1 then the attempted read size is equal to the number of bytes
// that are required to fill the buffer.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object.
//    - count: number of bytes that will be read from the stream.
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback (optional): ReadyCallback.
//
func (stream *BufferedInputStream) FillAsync(ctx context.Context, count, ioPriority int, callback AsyncReadyCallback) {
	var _arg0 *C.GBufferedInputStream // out
	var _arg3 *C.GCancellable         // out
	var _arg1 C.gssize                // out
	var _arg2 C.int                   // out
	var _arg4 C.GAsyncReadyCallback   // out
	var _arg5 C.gpointer

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.gssize(count)
	_arg2 = C.int(ioPriority)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_buffered_input_stream_fill_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(count)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// FillFinish finishes an asynchronous read.
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - gssize of the read stream, or -1 on an error.
//
func (stream *BufferedInputStream) FillFinish(result AsyncResulter) (int, error) {
	var _arg0 *C.GBufferedInputStream // out
	var _arg1 *C.GAsyncResult         // out
	var _cret C.gssize                // in
	var _cerr *C.GError               // in

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_buffered_input_stream_fill_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gssize, _goerr
}

// Available gets the size of the available data within the stream.
//
// The function returns the following values:
//
//    - gsize: size of the available stream.
//
func (stream *BufferedInputStream) Available() uint {
	var _arg0 *C.GBufferedInputStream // out
	var _cret C.gsize                 // in

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_buffered_input_stream_get_available(_arg0)
	runtime.KeepAlive(stream)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// BufferSize gets the size of the input buffer.
//
// The function returns the following values:
//
//    - gsize: current buffer size.
//
func (stream *BufferedInputStream) BufferSize() uint {
	var _arg0 *C.GBufferedInputStream // out
	var _cret C.gsize                 // in

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_buffered_input_stream_get_buffer_size(_arg0)
	runtime.KeepAlive(stream)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// Peek peeks in the buffer, copying data of size count into buffer, offset
// offset bytes.
//
// The function takes the following parameters:
//
//    - buffer: pointer to an allocated chunk of memory.
//    - offset: #gsize.
//
// The function returns the following values:
//
//    - gsize of the number of bytes peeked, or -1 on error.
//
func (stream *BufferedInputStream) Peek(buffer []byte, offset uint) uint {
	var _arg0 *C.GBufferedInputStream // out
	var _arg1 *C.void                 // out
	var _arg3 C.gsize
	var _arg2 C.gsize // out
	var _cret C.gsize // in

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))
	_arg3 = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		_arg1 = (*C.void)(unsafe.Pointer(&buffer[0]))
	}
	_arg2 = C.gsize(offset)

	_cret = C.g_buffered_input_stream_peek(_arg0, unsafe.Pointer(_arg1), _arg2, _arg3)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(offset)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// PeekBuffer returns the buffer with the currently available bytes. The
// returned buffer must not be modified and will become invalid when reading
// from the stream or filling the buffer.
//
// The function returns the following values:
//
//    - guint8s: read-only buffer.
//
func (stream *BufferedInputStream) PeekBuffer() []byte {
	var _arg0 *C.GBufferedInputStream // out
	var _cret unsafe.Pointer          // in
	var _arg1 C.gsize                 // in

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_buffered_input_stream_peek_buffer(_arg0, &_arg1)
	runtime.KeepAlive(stream)

	var _guint8s []byte // out

	_guint8s = make([]byte, _arg1)
	copy(_guint8s, unsafe.Slice((*byte)(unsafe.Pointer(_cret)), _arg1))

	return _guint8s
}

// ReadByte tries to read a single byte from the stream or the buffer. Will
// block during this read.
//
// On success, the byte read from the stream is returned. On end of stream -1 is
// returned but it's not an exceptional error and error is not set.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned. If an operation was
// partially finished when the operation was cancelled the partial result will
// be returned, without an error.
//
// On error -1 is returned and error is set accordingly.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//
// The function returns the following values:
//
//    - gint: byte read from the stream, or -1 on end of stream or error.
//
func (stream *BufferedInputStream) ReadByte(ctx context.Context) (int, error) {
	var _arg0 *C.GBufferedInputStream // out
	var _arg1 *C.GCancellable         // out
	var _cret C.int                   // in
	var _cerr *C.GError               // in

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C.g_buffered_input_stream_read_byte(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)

	var _gint int    // out
	var _goerr error // out

	_gint = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gint, _goerr
}

// SetBufferSize sets the size of the internal buffer of stream to size, or to
// the size of the contents of the buffer. The buffer can never be resized
// smaller than its current contents.
//
// The function takes the following parameters:
//
//    - size: #gsize.
//
func (stream *BufferedInputStream) SetBufferSize(size uint) {
	var _arg0 *C.GBufferedInputStream // out
	var _arg1 C.gsize                 // out

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = C.gsize(size)

	C.g_buffered_input_stream_set_buffer_size(_arg0, _arg1)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(size)
}
