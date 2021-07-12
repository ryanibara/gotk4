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
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
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
	// Fill tries to read @count bytes from the stream into the buffer. Will
	// block during this read.
	//
	// If @count is zero, returns zero and does nothing. A value of @count
	// larger than G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT error.
	//
	// On success, the number of bytes read into the buffer is returned. It is
	// not an error if this is not the same as the requested size, as it can
	// happen e.g. near the end of a file. Zero is returned on end of file (or
	// if @count is zero), but never otherwise.
	//
	// If @count is -1 then the attempted read size is equal to the number of
	// bytes that are required to fill the buffer.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	//
	// On error -1 is returned and @error is set accordingly.
	//
	// For the asynchronous, non-blocking, version of this function, see
	// g_buffered_input_stream_fill_async().
	Fill(count int, cancellable Cancellabler) (int, error)
	// FillAsync reads data into @stream's buffer asynchronously, up to @count
	// size. @io_priority can be used to prioritize reads. For the synchronous
	// version of this function, see g_buffered_input_stream_fill().
	//
	// If @count is -1 then the attempted read size is equal to the number of
	// bytes that are required to fill the buffer.
	FillAsync(count int, ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback)
	// FillFinish finishes an asynchronous read.
	FillFinish(result AsyncResulter) (int, error)
}

// BufferedInputStreamer describes BufferedInputStream's methods.
type BufferedInputStreamer interface {
	// Fill tries to read @count bytes from the stream into the buffer.
	Fill(count int, cancellable Cancellabler) (int, error)
	// FillAsync reads data into @stream's buffer asynchronously, up to @count
	// size.
	FillAsync(count int, ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback)
	// FillFinish finishes an asynchronous read.
	FillFinish(result AsyncResulter) (int, error)
	// Available gets the size of the available data within the stream.
	Available() uint
	// BufferSize gets the size of the input buffer.
	BufferSize() uint
	// Peek peeks in the buffer, copying data of size @count into @buffer,
	// offset @offset bytes.
	Peek(buffer []byte, offset uint) uint
	// ReadByte tries to read a single byte from the stream or the buffer.
	ReadByte(cancellable Cancellabler) (int, error)
	// SetBufferSize sets the size of the internal buffer of @stream to @size,
	// or to the size of the contents of the buffer.
	SetBufferSize(size uint)
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
	FilterInputStream

	Seekable
}

var (
	_ BufferedInputStreamer = (*BufferedInputStream)(nil)
	_ gextras.Nativer       = (*BufferedInputStream)(nil)
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
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapBufferedInputStream(obj), nil
}

// NewBufferedInputStream creates a new Stream from the given @base_stream, with
// a buffer set to the default size (4 kilobytes).
func NewBufferedInputStream(baseStream InputStreamer) *BufferedInputStream {
	var _arg1 *C.GInputStream // out
	var _cret *C.GInputStream // in

	_arg1 = (*C.GInputStream)(unsafe.Pointer((baseStream).(gextras.Nativer).Native()))

	_cret = C.g_buffered_input_stream_new(_arg1)

	var _bufferedInputStream *BufferedInputStream // out

	_bufferedInputStream = wrapBufferedInputStream(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _bufferedInputStream
}

// NewBufferedInputStreamSized creates a new InputStream from the given
// @base_stream, with a buffer set to @size.
func NewBufferedInputStreamSized(baseStream InputStreamer, size uint) *BufferedInputStream {
	var _arg1 *C.GInputStream // out
	var _arg2 C.gsize         // out
	var _cret *C.GInputStream // in

	_arg1 = (*C.GInputStream)(unsafe.Pointer((baseStream).(gextras.Nativer).Native()))
	_arg2 = C.gsize(size)

	_cret = C.g_buffered_input_stream_new_sized(_arg1, _arg2)

	var _bufferedInputStream *BufferedInputStream // out

	_bufferedInputStream = wrapBufferedInputStream(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _bufferedInputStream
}

// Fill tries to read @count bytes from the stream into the buffer. Will block
// during this read.
//
// If @count is zero, returns zero and does nothing. A value of @count larger
// than G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes read into the buffer is returned. It is not
// an error if this is not the same as the requested size, as it can happen e.g.
// near the end of a file. Zero is returned on end of file (or if @count is
// zero), but never otherwise.
//
// If @count is -1 then the attempted read size is equal to the number of bytes
// that are required to fill the buffer.
//
// If @cancellable is not nil, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned. If an operation was
// partially finished when the operation was cancelled the partial result will
// be returned, without an error.
//
// On error -1 is returned and @error is set accordingly.
//
// For the asynchronous, non-blocking, version of this function, see
// g_buffered_input_stream_fill_async().
func (stream *BufferedInputStream) Fill(count int, cancellable Cancellabler) (int, error) {
	var _arg0 *C.GBufferedInputStream // out
	var _arg1 C.gssize                // out
	var _arg2 *C.GCancellable         // out
	var _cret C.gssize                // in
	var _cerr *C.GError               // in

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = C.gssize(count)
	_arg2 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_buffered_input_stream_fill(_arg0, _arg1, _arg2, &_cerr)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gssize, _goerr
}

// FillAsync reads data into @stream's buffer asynchronously, up to @count size.
// @io_priority can be used to prioritize reads. For the synchronous version of
// this function, see g_buffered_input_stream_fill().
//
// If @count is -1 then the attempted read size is equal to the number of bytes
// that are required to fill the buffer.
func (stream *BufferedInputStream) FillAsync(count int, ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback) {
	var _arg0 *C.GBufferedInputStream // out
	var _arg1 C.gssize                // out
	var _arg2 C.int                   // out
	var _arg3 *C.GCancellable         // out
	var _arg4 C.GAsyncReadyCallback   // out
	var _arg5 C.gpointer

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = C.gssize(count)
	_arg2 = C.int(ioPriority)
	_arg3 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(gbox.Assign(callback))

	C.g_buffered_input_stream_fill_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// FillFinish finishes an asynchronous read.
func (stream *BufferedInputStream) FillFinish(result AsyncResulter) (int, error) {
	var _arg0 *C.GBufferedInputStream // out
	var _arg1 *C.GAsyncResult         // out
	var _cret C.gssize                // in
	var _cerr *C.GError               // in

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.g_buffered_input_stream_fill_finish(_arg0, _arg1, &_cerr)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gssize, _goerr
}

// Available gets the size of the available data within the stream.
func (stream *BufferedInputStream) Available() uint {
	var _arg0 *C.GBufferedInputStream // out
	var _cret C.gsize                 // in

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_buffered_input_stream_get_available(_arg0)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// BufferSize gets the size of the input buffer.
func (stream *BufferedInputStream) BufferSize() uint {
	var _arg0 *C.GBufferedInputStream // out
	var _cret C.gsize                 // in

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_buffered_input_stream_get_buffer_size(_arg0)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// Peek peeks in the buffer, copying data of size @count into @buffer, offset
// @offset bytes.
func (stream *BufferedInputStream) Peek(buffer []byte, offset uint) uint {
	var _arg0 *C.GBufferedInputStream // out
	var _arg1 *C.void
	var _arg3 C.gsize
	var _arg2 C.gsize // out
	var _cret C.gsize // in

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))
	_arg3 = C.gsize(len(buffer))
	if len(buffer) > 0 {
		_arg1 = (*C.void)(unsafe.Pointer(&buffer[0]))
	}
	_arg2 = C.gsize(offset)

	_cret = C.g_buffered_input_stream_peek(_arg0, unsafe.Pointer(_arg1), _arg2, _arg3)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// ReadByte tries to read a single byte from the stream or the buffer. Will
// block during this read.
//
// On success, the byte read from the stream is returned. On end of stream -1 is
// returned but it's not an exceptional error and @error is not set.
//
// If @cancellable is not nil, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned. If an operation was
// partially finished when the operation was cancelled the partial result will
// be returned, without an error.
//
// On error -1 is returned and @error is set accordingly.
func (stream *BufferedInputStream) ReadByte(cancellable Cancellabler) (int, error) {
	var _arg0 *C.GBufferedInputStream // out
	var _arg1 *C.GCancellable         // out
	var _cret C.int                   // in
	var _cerr *C.GError               // in

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_buffered_input_stream_read_byte(_arg0, _arg1, &_cerr)

	var _gint int    // out
	var _goerr error // out

	_gint = int(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gint, _goerr
}

// SetBufferSize sets the size of the internal buffer of @stream to @size, or to
// the size of the contents of the buffer. The buffer can never be resized
// smaller than its current contents.
func (stream *BufferedInputStream) SetBufferSize(size uint) {
	var _arg0 *C.GBufferedInputStream // out
	var _arg1 C.gsize                 // out

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = C.gsize(size)

	C.g_buffered_input_stream_set_buffer_size(_arg0, _arg1)
}
