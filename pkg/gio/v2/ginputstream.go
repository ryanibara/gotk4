// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdbool.h>
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
		{T: externglib.Type(C.g_input_stream_get_type()), F: marshalInputStream},
	})
}

type InputStreamPrivate struct {
	native C.GInputStreamPrivate
}

// WrapInputStreamPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapInputStreamPrivate(ptr unsafe.Pointer) *InputStreamPrivate {
	if ptr == nil {
		return nil
	}

	return (*InputStreamPrivate)(ptr)
}

func marshalInputStreamPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapInputStreamPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (i *InputStreamPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}

// InputStream has functions to read from a stream (g_input_stream_read()), to
// close a stream (g_input_stream_close()) and to skip some content
// (g_input_stream_skip()).
//
// To copy the content of an input stream to an output stream without manually
// handling the reads and writes, use g_output_stream_splice().
//
// See the documentation for OStream for details of thread safety of streaming
// APIs.
//
// All of these functions have async variants too.
type InputStream interface {
	gextras.Objector

	// ClearPending clears the pending flag on @stream.
	ClearPending()
	// Close closes the stream, releasing resources related to it.
	//
	// Once the stream is closed, all other operations will return
	// G_IO_ERROR_CLOSED. Closing a stream multiple times will not return an
	// error.
	//
	// Streams will be automatically closed when the last reference is dropped,
	// but you might want to call this function to make sure resources are
	// released as early as possible.
	//
	// Some streams might keep the backing store of the stream (e.g. a file
	// descriptor) open after the stream is closed. See the documentation for
	// the individual stream for details.
	//
	// On failure the first error that happened will be reported, but the close
	// operation will finish as much as possible. A stream that failed to close
	// will still return G_IO_ERROR_CLOSED for all operations. Still, it is
	// important to check and report the error to the user.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	// Cancelling a close will still leave the stream closed, but some streams
	// can use a faster close that doesn't block to e.g. check errors.
	Close(cancellable Cancellable) error
	// CloseAsync requests an asynchronous closes of the stream, releasing
	// resources related to it. When the operation is finished @callback will be
	// called. You can then call g_input_stream_close_finish() to get the result
	// of the operation.
	//
	// For behaviour details see g_input_stream_close().
	//
	// The asynchronous methods have a default fallback that uses threads to
	// implement asynchronicity, so they are optional for inheriting classes.
	// However, if you override one you must override all.
	CloseAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// CloseFinish finishes closing a stream asynchronously, started from
	// g_input_stream_close_async().
	CloseFinish(result AsyncResult) error
	// HasPending checks if an input stream has pending actions.
	HasPending() bool
	// IsClosed checks if an input stream is closed.
	IsClosed() bool
	// Read tries to read @count bytes from the stream into the buffer starting
	// at @buffer. Will block during this read.
	//
	// If count is zero returns zero and does nothing. A value of @count larger
	// than G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT error.
	//
	// On success, the number of bytes read into the buffer is returned. It is
	// not an error if this is not the same as the requested size, as it can
	// happen e.g. near the end of a file. Zero is returned on end of file (or
	// if @count is zero), but never otherwise.
	//
	// The returned @buffer is not a nul-terminated string, it can contain nul
	// bytes at any position, and this function doesn't nul-terminate the
	// @buffer.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	//
	// On error -1 is returned and @error is set accordingly.
	Read(cancellable Cancellable) (buffer []byte, count uint, gssize int, err error)
	// ReadAll tries to read @count bytes from the stream into the buffer
	// starting at @buffer. Will block during this read.
	//
	// This function is similar to g_input_stream_read(), except it tries to
	// read as many bytes as requested, only stopping on an error or end of
	// stream.
	//
	// On a successful read of @count bytes, or if we reached the end of the
	// stream, true is returned, and @bytes_read is set to the number of bytes
	// read into @buffer.
	//
	// If there is an error during the operation false is returned and @error is
	// set to indicate the error status.
	//
	// As a special exception to the normal conventions for functions that use
	// #GError, if this function returns false (and sets @error) then
	// @bytes_read will be set to the number of bytes that were successfully
	// read before the error was encountered. This functionality is only
	// available from C. If you need it from another language then you must
	// write your own loop around g_input_stream_read().
	ReadAll(cancellable Cancellable) (buffer []byte, count uint, bytesRead uint, err error)
	// ReadAllAsync: request an asynchronous read of @count bytes from the
	// stream into the buffer starting at @buffer.
	//
	// This is the asynchronous equivalent of g_input_stream_read_all().
	//
	// Call g_input_stream_read_all_finish() to collect the result.
	//
	// Any outstanding I/O request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	ReadAllAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) (buffer []byte, count uint)
	// ReadAllFinish finishes an asynchronous stream read operation started with
	// g_input_stream_read_all_async().
	//
	// As a special exception to the normal conventions for functions that use
	// #GError, if this function returns false (and sets @error) then
	// @bytes_read will be set to the number of bytes that were successfully
	// read before the error was encountered. This functionality is only
	// available from C. If you need it from another language then you must
	// write your own loop around g_input_stream_read_async().
	ReadAllFinish(result AsyncResult) (bytesRead uint, err error)
	// ReadAsync: request an asynchronous read of @count bytes from the stream
	// into the buffer starting at @buffer. When the operation is finished
	// @callback will be called. You can then call g_input_stream_read_finish()
	// to get the result of the operation.
	//
	// During an async request no other sync and async calls are allowed on
	// @stream, and will result in G_IO_ERROR_PENDING errors.
	//
	// A value of @count larger than G_MAXSSIZE will cause a
	// G_IO_ERROR_INVALID_ARGUMENT error.
	//
	// On success, the number of bytes read into the buffer will be passed to
	// the callback. It is not an error if this is not the same as the requested
	// size, as it can happen e.g. near the end of a file, but generally we try
	// to read as many bytes as requested. Zero is returned on end of file (or
	// if @count is zero), but never otherwise.
	//
	// Any outstanding i/o request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	//
	// The asynchronous methods have a default fallback that uses threads to
	// implement asynchronicity, so they are optional for inheriting classes.
	// However, if you override one you must override all.
	ReadAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) (buffer []byte, count uint)
	// ReadBytes: like g_input_stream_read(), this tries to read @count bytes
	// from the stream in a blocking fashion. However, rather than reading into
	// a user-supplied buffer, this will create a new #GBytes containing the
	// data that was read. This may be easier to use from language bindings.
	//
	// If count is zero, returns a zero-length #GBytes and does nothing. A value
	// of @count larger than G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT
	// error.
	//
	// On success, a new #GBytes is returned. It is not an error if the size of
	// this object is not the same as the requested size, as it can happen e.g.
	// near the end of a file. A zero-length #GBytes is returned on end of file
	// (or if @count is zero), but never otherwise.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	//
	// On error nil is returned and @error is set accordingly.
	ReadBytes(count uint, cancellable Cancellable) (bytes *glib.Bytes, err error)
	// ReadBytesAsync: request an asynchronous read of @count bytes from the
	// stream into a new #GBytes. When the operation is finished @callback will
	// be called. You can then call g_input_stream_read_bytes_finish() to get
	// the result of the operation.
	//
	// During an async request no other sync and async calls are allowed on
	// @stream, and will result in G_IO_ERROR_PENDING errors.
	//
	// A value of @count larger than G_MAXSSIZE will cause a
	// G_IO_ERROR_INVALID_ARGUMENT error.
	//
	// On success, the new #GBytes will be passed to the callback. It is not an
	// error if this is smaller than the requested size, as it can happen e.g.
	// near the end of a file, but generally we try to read as many bytes as
	// requested. Zero is returned on end of file (or if @count is zero), but
	// never otherwise.
	//
	// Any outstanding I/O request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	ReadBytesAsync(count uint, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// ReadBytesFinish finishes an asynchronous stream read-into-#GBytes
	// operation.
	ReadBytesFinish(result AsyncResult) (bytes *glib.Bytes, err error)
	// ReadFinish finishes an asynchronous stream read operation.
	ReadFinish(result AsyncResult) (gssize int, err error)
	// SetPending sets @stream to have actions pending. If the pending flag is
	// already set or @stream is closed, it will return false and set @error.
	SetPending() error
	// Skip tries to skip @count bytes from the stream. Will block during the
	// operation.
	//
	// This is identical to g_input_stream_read(), from a behaviour standpoint,
	// but the bytes that are skipped are not returned to the user. Some streams
	// have an implementation that is more efficient than reading the data.
	//
	// This function is optional for inherited classes, as the default
	// implementation emulates it using read.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	Skip(count uint, cancellable Cancellable) (gssize int, err error)
	// SkipAsync: request an asynchronous skip of @count bytes from the stream.
	// When the operation is finished @callback will be called. You can then
	// call g_input_stream_skip_finish() to get the result of the operation.
	//
	// During an async request no other sync and async calls are allowed, and
	// will result in G_IO_ERROR_PENDING errors.
	//
	// A value of @count larger than G_MAXSSIZE will cause a
	// G_IO_ERROR_INVALID_ARGUMENT error.
	//
	// On success, the number of bytes skipped will be passed to the callback.
	// It is not an error if this is not the same as the requested size, as it
	// can happen e.g. near the end of a file, but generally we try to skip as
	// many bytes as requested. Zero is returned on end of file (or if @count is
	// zero), but never otherwise.
	//
	// Any outstanding i/o request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	//
	// The asynchronous methods have a default fallback that uses threads to
	// implement asynchronicity, so they are optional for inheriting classes.
	// However, if you override one, you must override all.
	SkipAsync(count uint, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// SkipFinish finishes a stream skip operation.
	SkipFinish(result AsyncResult) (gssize int, err error)
}

// inputStream implements the InputStream interface.
type inputStream struct {
	gextras.Objector
}

var _ InputStream = (*inputStream)(nil)

// WrapInputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapInputStream(obj *externglib.Object) InputStream {
	return InputStream{
		Objector: obj,
	}
}

func marshalInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapInputStream(obj), nil
}

// ClearPending clears the pending flag on @stream.
func (s inputStream) ClearPending() {
	var arg0 *C.GInputStream

	arg0 = (*C.GInputStream)(s.Native())

	C.g_input_stream_clear_pending(arg0)
}

// Close closes the stream, releasing resources related to it.
//
// Once the stream is closed, all other operations will return
// G_IO_ERROR_CLOSED. Closing a stream multiple times will not return an
// error.
//
// Streams will be automatically closed when the last reference is dropped,
// but you might want to call this function to make sure resources are
// released as early as possible.
//
// Some streams might keep the backing store of the stream (e.g. a file
// descriptor) open after the stream is closed. See the documentation for
// the individual stream for details.
//
// On failure the first error that happened will be reported, but the close
// operation will finish as much as possible. A stream that failed to close
// will still return G_IO_ERROR_CLOSED for all operations. Still, it is
// important to check and report the error to the user.
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
// Cancelling a close will still leave the stream closed, but some streams
// can use a faster close that doesn't block to e.g. check errors.
func (s inputStream) Close(cancellable Cancellable) error {
	var arg0 *C.GInputStream
	var arg1 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GInputStream)(s.Native())
	arg1 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_input_stream_close(arg0, arg1, &gError)

	var goError error

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return goError
}

// CloseAsync requests an asynchronous closes of the stream, releasing
// resources related to it. When the operation is finished @callback will be
// called. You can then call g_input_stream_close_finish() to get the result
// of the operation.
//
// For behaviour details see g_input_stream_close().
//
// The asynchronous methods have a default fallback that uses threads to
// implement asynchronicity, so they are optional for inheriting classes.
// However, if you override one you must override all.
func (s inputStream) CloseAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GInputStream
	var arg1 C.int
	var arg2 *C.GCancellable
	var arg3 C.GAsyncReadyCallback
	var arg4 C.gpointer

	arg0 = (*C.GInputStream)(s.Native())
	arg1 = C.int(ioPriority)
	arg2 = (*C.GCancellable)(cancellable.Native())
	arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	arg4 = C.gpointer(box.Assign(callback))

	C.g_input_stream_close_async(arg0, arg1, arg2, arg3, arg4)
}

// CloseFinish finishes closing a stream asynchronously, started from
// g_input_stream_close_async().
func (s inputStream) CloseFinish(result AsyncResult) error {
	var arg0 *C.GInputStream
	var arg1 *C.GAsyncResult
	var gError *C.GError

	arg0 = (*C.GInputStream)(s.Native())
	arg1 = (*C.GAsyncResult)(result.Native())

	ret := C.g_input_stream_close_finish(arg0, arg1, &gError)

	var goError error

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return goError
}

// HasPending checks if an input stream has pending actions.
func (s inputStream) HasPending() bool {
	var arg0 *C.GInputStream

	arg0 = (*C.GInputStream)(s.Native())

	ret := C.g_input_stream_has_pending(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// IsClosed checks if an input stream is closed.
func (s inputStream) IsClosed() bool {
	var arg0 *C.GInputStream

	arg0 = (*C.GInputStream)(s.Native())

	ret := C.g_input_stream_is_closed(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Read tries to read @count bytes from the stream into the buffer starting
// at @buffer. Will block during this read.
//
// If count is zero returns zero and does nothing. A value of @count larger
// than G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes read into the buffer is returned. It is
// not an error if this is not the same as the requested size, as it can
// happen e.g. near the end of a file. Zero is returned on end of file (or
// if @count is zero), but never otherwise.
//
// The returned @buffer is not a nul-terminated string, it can contain nul
// bytes at any position, and this function doesn't nul-terminate the
// @buffer.
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
//
// On error -1 is returned and @error is set accordingly.
func (s inputStream) Read(cancellable Cancellable) (buffer []byte, count uint, gssize int, err error) {
	var arg0 *C.GInputStream
	var arg1 *C.void // out
	var arg2 C.gsize // out
	var arg3 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GInputStream)(s.Native())
	arg3 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_input_stream_read(arg0, &arg1, &arg2, arg3, &gError)

	var ret0 []byte
	var ret1 uint
	var ret2 int
	var goError error

	{
		ret0 = make([]byte, arg2)
		for i := 0; i < uintptr(arg2); i++ {
			src := (C.guint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i))
			ret0[i] = byte(src)
		}
	}

	ret1 = uint(arg2)

	ret2 = int(ret)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, ret1, ret2, goError
}

// ReadAll tries to read @count bytes from the stream into the buffer
// starting at @buffer. Will block during this read.
//
// This function is similar to g_input_stream_read(), except it tries to
// read as many bytes as requested, only stopping on an error or end of
// stream.
//
// On a successful read of @count bytes, or if we reached the end of the
// stream, true is returned, and @bytes_read is set to the number of bytes
// read into @buffer.
//
// If there is an error during the operation false is returned and @error is
// set to indicate the error status.
//
// As a special exception to the normal conventions for functions that use
// #GError, if this function returns false (and sets @error) then
// @bytes_read will be set to the number of bytes that were successfully
// read before the error was encountered. This functionality is only
// available from C. If you need it from another language then you must
// write your own loop around g_input_stream_read().
func (s inputStream) ReadAll(cancellable Cancellable) (buffer []byte, count uint, bytesRead uint, err error) {
	var arg0 *C.GInputStream
	var arg1 *C.void  // out
	var arg2 C.gsize  // out
	var arg3 *C.gsize // out
	var arg4 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GInputStream)(s.Native())
	arg4 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_input_stream_read_all(arg0, &arg1, &arg2, &arg3, arg4, &gError)

	var ret0 []byte
	var ret1 uint
	var ret2 uint
	var goError error

	{
		ret0 = make([]byte, arg2)
		for i := 0; i < uintptr(arg2); i++ {
			src := (C.guint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i))
			ret0[i] = byte(src)
		}
	}

	ret1 = uint(arg2)

	ret2 = uint(arg3)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, ret1, ret2, goError
}

// ReadAllAsync: request an asynchronous read of @count bytes from the
// stream into the buffer starting at @buffer.
//
// This is the asynchronous equivalent of g_input_stream_read_all().
//
// Call g_input_stream_read_all_finish() to collect the result.
//
// Any outstanding I/O request with higher priority (lower numerical value)
// will be executed before an outstanding request with lower priority.
// Default priority is G_PRIORITY_DEFAULT.
func (s inputStream) ReadAllAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) (buffer []byte, count uint) {
	var arg0 *C.GInputStream
	var arg1 *C.void // out
	var arg2 C.gsize // out
	var arg3 C.int
	var arg4 *C.GCancellable
	var arg5 C.GAsyncReadyCallback
	var arg6 C.gpointer

	arg0 = (*C.GInputStream)(s.Native())
	arg3 = C.int(ioPriority)
	arg4 = (*C.GCancellable)(cancellable.Native())
	arg5 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	arg6 = C.gpointer(box.Assign(callback))

	C.g_input_stream_read_all_async(arg0, &arg1, &arg2, arg3, arg4, arg5, arg6)

	var ret0 []byte
	var ret1 uint

	{
		ret0 = make([]byte, arg2)
		for i := 0; i < uintptr(arg2); i++ {
			src := (C.guint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i))
			ret0[i] = byte(src)
		}
	}

	ret1 = uint(arg2)

	return ret0, ret1
}

// ReadAllFinish finishes an asynchronous stream read operation started with
// g_input_stream_read_all_async().
//
// As a special exception to the normal conventions for functions that use
// #GError, if this function returns false (and sets @error) then
// @bytes_read will be set to the number of bytes that were successfully
// read before the error was encountered. This functionality is only
// available from C. If you need it from another language then you must
// write your own loop around g_input_stream_read_async().
func (s inputStream) ReadAllFinish(result AsyncResult) (bytesRead uint, err error) {
	var arg0 *C.GInputStream
	var arg1 *C.GAsyncResult
	var arg2 *C.gsize // out
	var gError *C.GError

	arg0 = (*C.GInputStream)(s.Native())
	arg1 = (*C.GAsyncResult)(result.Native())

	ret := C.g_input_stream_read_all_finish(arg0, arg1, &arg2, &gError)

	var ret0 uint
	var goError error

	ret0 = uint(arg2)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// ReadAsync: request an asynchronous read of @count bytes from the stream
// into the buffer starting at @buffer. When the operation is finished
// @callback will be called. You can then call g_input_stream_read_finish()
// to get the result of the operation.
//
// During an async request no other sync and async calls are allowed on
// @stream, and will result in G_IO_ERROR_PENDING errors.
//
// A value of @count larger than G_MAXSSIZE will cause a
// G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes read into the buffer will be passed to
// the callback. It is not an error if this is not the same as the requested
// size, as it can happen e.g. near the end of a file, but generally we try
// to read as many bytes as requested. Zero is returned on end of file (or
// if @count is zero), but never otherwise.
//
// Any outstanding i/o request with higher priority (lower numerical value)
// will be executed before an outstanding request with lower priority.
// Default priority is G_PRIORITY_DEFAULT.
//
// The asynchronous methods have a default fallback that uses threads to
// implement asynchronicity, so they are optional for inheriting classes.
// However, if you override one you must override all.
func (s inputStream) ReadAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) (buffer []byte, count uint) {
	var arg0 *C.GInputStream
	var arg1 *C.void // out
	var arg2 C.gsize // out
	var arg3 C.int
	var arg4 *C.GCancellable
	var arg5 C.GAsyncReadyCallback
	var arg6 C.gpointer

	arg0 = (*C.GInputStream)(s.Native())
	arg3 = C.int(ioPriority)
	arg4 = (*C.GCancellable)(cancellable.Native())
	arg5 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	arg6 = C.gpointer(box.Assign(callback))

	C.g_input_stream_read_async(arg0, &arg1, &arg2, arg3, arg4, arg5, arg6)

	var ret0 []byte
	var ret1 uint

	{
		ret0 = make([]byte, arg2)
		for i := 0; i < uintptr(arg2); i++ {
			src := (C.guint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i))
			ret0[i] = byte(src)
		}
	}

	ret1 = uint(arg2)

	return ret0, ret1
}

// ReadBytes: like g_input_stream_read(), this tries to read @count bytes
// from the stream in a blocking fashion. However, rather than reading into
// a user-supplied buffer, this will create a new #GBytes containing the
// data that was read. This may be easier to use from language bindings.
//
// If count is zero, returns a zero-length #GBytes and does nothing. A value
// of @count larger than G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT
// error.
//
// On success, a new #GBytes is returned. It is not an error if the size of
// this object is not the same as the requested size, as it can happen e.g.
// near the end of a file. A zero-length #GBytes is returned on end of file
// (or if @count is zero), but never otherwise.
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
//
// On error nil is returned and @error is set accordingly.
func (s inputStream) ReadBytes(count uint, cancellable Cancellable) (bytes *glib.Bytes, err error) {
	var arg0 *C.GInputStream
	var arg1 C.gsize
	var arg2 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GInputStream)(s.Native())
	arg1 = C.gsize(count)
	arg2 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_input_stream_read_bytes(arg0, arg1, arg2, &gError)

	var ret0 *glib.Bytes
	var goError error

	{
		ret0 = glib.WrapBytes(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *glib.Bytes) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// ReadBytesAsync: request an asynchronous read of @count bytes from the
// stream into a new #GBytes. When the operation is finished @callback will
// be called. You can then call g_input_stream_read_bytes_finish() to get
// the result of the operation.
//
// During an async request no other sync and async calls are allowed on
// @stream, and will result in G_IO_ERROR_PENDING errors.
//
// A value of @count larger than G_MAXSSIZE will cause a
// G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the new #GBytes will be passed to the callback. It is not an
// error if this is smaller than the requested size, as it can happen e.g.
// near the end of a file, but generally we try to read as many bytes as
// requested. Zero is returned on end of file (or if @count is zero), but
// never otherwise.
//
// Any outstanding I/O request with higher priority (lower numerical value)
// will be executed before an outstanding request with lower priority.
// Default priority is G_PRIORITY_DEFAULT.
func (s inputStream) ReadBytesAsync(count uint, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GInputStream
	var arg1 C.gsize
	var arg2 C.int
	var arg3 *C.GCancellable
	var arg4 C.GAsyncReadyCallback
	var arg5 C.gpointer

	arg0 = (*C.GInputStream)(s.Native())
	arg1 = C.gsize(count)
	arg2 = C.int(ioPriority)
	arg3 = (*C.GCancellable)(cancellable.Native())
	arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	arg5 = C.gpointer(box.Assign(callback))

	C.g_input_stream_read_bytes_async(arg0, arg1, arg2, arg3, arg4, arg5)
}

// ReadBytesFinish finishes an asynchronous stream read-into-#GBytes
// operation.
func (s inputStream) ReadBytesFinish(result AsyncResult) (bytes *glib.Bytes, err error) {
	var arg0 *C.GInputStream
	var arg1 *C.GAsyncResult
	var gError *C.GError

	arg0 = (*C.GInputStream)(s.Native())
	arg1 = (*C.GAsyncResult)(result.Native())

	ret := C.g_input_stream_read_bytes_finish(arg0, arg1, &gError)

	var ret0 *glib.Bytes
	var goError error

	{
		ret0 = glib.WrapBytes(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *glib.Bytes) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// ReadFinish finishes an asynchronous stream read operation.
func (s inputStream) ReadFinish(result AsyncResult) (gssize int, err error) {
	var arg0 *C.GInputStream
	var arg1 *C.GAsyncResult
	var gError *C.GError

	arg0 = (*C.GInputStream)(s.Native())
	arg1 = (*C.GAsyncResult)(result.Native())

	ret := C.g_input_stream_read_finish(arg0, arg1, &gError)

	var ret0 int
	var goError error

	ret0 = int(ret)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// SetPending sets @stream to have actions pending. If the pending flag is
// already set or @stream is closed, it will return false and set @error.
func (s inputStream) SetPending() error {
	var arg0 *C.GInputStream
	var gError *C.GError

	arg0 = (*C.GInputStream)(s.Native())

	ret := C.g_input_stream_set_pending(arg0, &gError)

	var goError error

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return goError
}

// Skip tries to skip @count bytes from the stream. Will block during the
// operation.
//
// This is identical to g_input_stream_read(), from a behaviour standpoint,
// but the bytes that are skipped are not returned to the user. Some streams
// have an implementation that is more efficient than reading the data.
//
// This function is optional for inherited classes, as the default
// implementation emulates it using read.
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
func (s inputStream) Skip(count uint, cancellable Cancellable) (gssize int, err error) {
	var arg0 *C.GInputStream
	var arg1 C.gsize
	var arg2 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GInputStream)(s.Native())
	arg1 = C.gsize(count)
	arg2 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_input_stream_skip(arg0, arg1, arg2, &gError)

	var ret0 int
	var goError error

	ret0 = int(ret)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// SkipAsync: request an asynchronous skip of @count bytes from the stream.
// When the operation is finished @callback will be called. You can then
// call g_input_stream_skip_finish() to get the result of the operation.
//
// During an async request no other sync and async calls are allowed, and
// will result in G_IO_ERROR_PENDING errors.
//
// A value of @count larger than G_MAXSSIZE will cause a
// G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes skipped will be passed to the callback.
// It is not an error if this is not the same as the requested size, as it
// can happen e.g. near the end of a file, but generally we try to skip as
// many bytes as requested. Zero is returned on end of file (or if @count is
// zero), but never otherwise.
//
// Any outstanding i/o request with higher priority (lower numerical value)
// will be executed before an outstanding request with lower priority.
// Default priority is G_PRIORITY_DEFAULT.
//
// The asynchronous methods have a default fallback that uses threads to
// implement asynchronicity, so they are optional for inheriting classes.
// However, if you override one, you must override all.
func (s inputStream) SkipAsync(count uint, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GInputStream
	var arg1 C.gsize
	var arg2 C.int
	var arg3 *C.GCancellable
	var arg4 C.GAsyncReadyCallback
	var arg5 C.gpointer

	arg0 = (*C.GInputStream)(s.Native())
	arg1 = C.gsize(count)
	arg2 = C.int(ioPriority)
	arg3 = (*C.GCancellable)(cancellable.Native())
	arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	arg5 = C.gpointer(box.Assign(callback))

	C.g_input_stream_skip_async(arg0, arg1, arg2, arg3, arg4, arg5)
}

// SkipFinish finishes a stream skip operation.
func (s inputStream) SkipFinish(result AsyncResult) (gssize int, err error) {
	var arg0 *C.GInputStream
	var arg1 *C.GAsyncResult
	var gError *C.GError

	arg0 = (*C.GInputStream)(s.Native())
	arg1 = (*C.GAsyncResult)(result.Native())

	ret := C.g_input_stream_skip_finish(arg0, arg1, &gError)

	var ret0 int
	var goError error

	ret0 = int(ret)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}
