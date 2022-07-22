// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gio2_InputStreamClass_close_finish(GInputStream*, GAsyncResult*, GError**);
// extern gboolean _gotk4_gio2_InputStreamClass_close_fn(GInputStream*, GCancellable*, GError**);
// extern gssize _gotk4_gio2_InputStreamClass_read_finish(GInputStream*, GAsyncResult*, GError**);
// extern gssize _gotk4_gio2_InputStreamClass_skip(GInputStream*, gsize, GCancellable*, GError**);
// extern gssize _gotk4_gio2_InputStreamClass_skip_finish(GInputStream*, GAsyncResult*, GError**);
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// GType values.
var (
	GTypeInputStream = coreglib.Type(C.g_input_stream_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeInputStream, F: marshalInputStream},
	})
}

// InputStreamOverrider contains methods that are overridable.
type InputStreamOverrider interface {
	// CloseFinish finishes closing a stream asynchronously, started from
	// g_input_stream_close_async().
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	CloseFinish(result AsyncResulter) error
	// The function takes the following parameters:
	//
	CloseFn(ctx context.Context) error
	// ReadFinish finishes an asynchronous stream read operation.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	// The function returns the following values:
	//
	//    - gssize: number of bytes read in, or -1 on error, or 0 on end of file.
	//
	ReadFinish(result AsyncResulter) (int, error)
	// Skip tries to skip count bytes from the stream. Will block during the
	// operation.
	//
	// This is identical to g_input_stream_read(), from a behaviour standpoint,
	// but the bytes that are skipped are not returned to the user. Some streams
	// have an implementation that is more efficient than reading the data.
	//
	// This function is optional for inherited classes, as the default
	// implementation emulates it using read.
	//
	// If cancellable is not NULL, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable object, NULL to ignore.
	//    - count: number of bytes that will be skipped from the stream.
	//
	// The function returns the following values:
	//
	//    - gssize: number of bytes skipped, or -1 on error.
	//
	Skip(ctx context.Context, count uint) (int, error)
	// SkipFinish finishes a stream skip operation.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	// The function returns the following values:
	//
	//    - gssize: size of the bytes skipped, or -1 on error.
	//
	SkipFinish(result AsyncResulter) (int, error)
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
type InputStream struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*InputStream)(nil)
)

// InputStreamer describes types inherited from class InputStream.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type InputStreamer interface {
	coreglib.Objector
	baseInputStream() *InputStream
}

var _ InputStreamer = (*InputStream)(nil)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:     GTypeInputStream,
		GoType:    reflect.TypeOf((*InputStream)(nil)),
		InitClass: initClassInputStream,
	})
}

func initClassInputStream(gclass unsafe.Pointer, goval any) {

	pclass := (*C.GInputStreamClass)(unsafe.Pointer(gclass))

	if _, ok := goval.(interface {
		CloseFinish(result AsyncResulter) error
	}); ok {
		pclass.close_finish = (*[0]byte)(C._gotk4_gio2_InputStreamClass_close_finish)
	}

	if _, ok := goval.(interface {
		CloseFn(ctx context.Context) error
	}); ok {
		pclass.close_fn = (*[0]byte)(C._gotk4_gio2_InputStreamClass_close_fn)
	}

	if _, ok := goval.(interface {
		ReadFinish(result AsyncResulter) (int, error)
	}); ok {
		pclass.read_finish = (*[0]byte)(C._gotk4_gio2_InputStreamClass_read_finish)
	}

	if _, ok := goval.(interface {
		Skip(ctx context.Context, count uint) (int, error)
	}); ok {
		pclass.skip = (*[0]byte)(C._gotk4_gio2_InputStreamClass_skip)
	}

	if _, ok := goval.(interface {
		SkipFinish(result AsyncResulter) (int, error)
	}); ok {
		pclass.skip_finish = (*[0]byte)(C._gotk4_gio2_InputStreamClass_skip_finish)
	}
	if goval, ok := goval.(interface{ InitInputStream(*InputStreamClass) }); ok {
		klass := (*InputStreamClass)(gextras.NewStructNative(gclass))
		goval.InitInputStream(klass)
	}
}

//export _gotk4_gio2_InputStreamClass_close_finish
func _gotk4_gio2_InputStreamClass_close_finish(arg0 *C.GInputStream, arg1 *C.GAsyncResult, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface {
		CloseFinish(result AsyncResulter) error
	})

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	_goerr := iface.CloseFinish(_result)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_InputStreamClass_close_fn
func _gotk4_gio2_InputStreamClass_close_fn(arg0 *C.GInputStream, arg1 *C.GCancellable, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface {
		CloseFn(ctx context.Context) error
	})

	var _cancellable context.Context // out

	if arg1 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg1))
	}

	_goerr := iface.CloseFn(_cancellable)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_InputStreamClass_read_finish
func _gotk4_gio2_InputStreamClass_read_finish(arg0 *C.GInputStream, arg1 *C.GAsyncResult, _cerr **C.GError) (cret C.gssize) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface {
		ReadFinish(result AsyncResulter) (int, error)
	})

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	gssize, _goerr := iface.ReadFinish(_result)

	cret = C.gssize(gssize)
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_InputStreamClass_skip
func _gotk4_gio2_InputStreamClass_skip(arg0 *C.GInputStream, arg1 C.gsize, arg2 *C.GCancellable, _cerr **C.GError) (cret C.gssize) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Skip(ctx context.Context, count uint) (int, error)
	})

	var _cancellable context.Context // out
	var _count uint                  // out

	if arg2 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg2))
	}
	_count = uint(arg1)

	gssize, _goerr := iface.Skip(_cancellable, _count)

	cret = C.gssize(gssize)
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_InputStreamClass_skip_finish
func _gotk4_gio2_InputStreamClass_skip_finish(arg0 *C.GInputStream, arg1 *C.GAsyncResult, _cerr **C.GError) (cret C.gssize) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface {
		SkipFinish(result AsyncResulter) (int, error)
	})

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	gssize, _goerr := iface.SkipFinish(_result)

	cret = C.gssize(gssize)
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

func wrapInputStream(obj *coreglib.Object) *InputStream {
	return &InputStream{
		Object: obj,
	}
}

func marshalInputStream(p uintptr) (interface{}, error) {
	return wrapInputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (stream *InputStream) baseInputStream() *InputStream {
	return stream
}

// BaseInputStream returns the underlying base object.
func BaseInputStream(obj InputStreamer) *InputStream {
	return obj.baseInputStream()
}

// ClearPending clears the pending flag on stream.
func (stream *InputStream) ClearPending() {
	var _arg0 *C.GInputStream // out

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	C.g_input_stream_clear_pending(_arg0)
	runtime.KeepAlive(stream)
}

// Close closes the stream, releasing resources related to it.
//
// Once the stream is closed, all other operations will return
// G_IO_ERROR_CLOSED. Closing a stream multiple times will not return an error.
//
// Streams will be automatically closed when the last reference is dropped, but
// you might want to call this function to make sure resources are released as
// early as possible.
//
// Some streams might keep the backing store of the stream (e.g. a file
// descriptor) open after the stream is closed. See the documentation for the
// individual stream for details.
//
// On failure the first error that happened will be reported, but the close
// operation will finish as much as possible. A stream that failed to close will
// still return G_IO_ERROR_CLOSED for all operations. Still, it is important to
// check and report the error to the user.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned. Cancelling a close will
// still leave the stream closed, but some streams can use a faster close that
// doesn't block to e.g. check errors.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//
func (stream *InputStream) Close(ctx context.Context) error {
	var _arg0 *C.GInputStream // out
	var _arg1 *C.GCancellable // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	C.g_input_stream_close(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// CloseAsync requests an asynchronous closes of the stream, releasing resources
// related to it. When the operation is finished callback will be called. You
// can then call g_input_stream_close_finish() to get the result of the
// operation.
//
// For behaviour details see g_input_stream_close().
//
// The asynchronous methods have a default fallback that uses threads to
// implement asynchronicity, so they are optional for inheriting classes.
// However, if you override one you must override all.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellable object.
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (stream *InputStream) CloseAsync(ctx context.Context, ioPriority int, callback AsyncReadyCallback) {
	var _arg0 *C.GInputStream       // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.int                 // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.int(ioPriority)
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_input_stream_close_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// CloseFinish finishes closing a stream asynchronously, started from
// g_input_stream_close_async().
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (stream *InputStream) CloseFinish(result AsyncResulter) error {
	var _arg0 *C.GInputStream // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C.g_input_stream_close_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// HasPending checks if an input stream has pending actions.
//
// The function returns the following values:
//
//    - ok: TRUE if stream has pending actions.
//
func (stream *InputStream) HasPending() bool {
	var _arg0 *C.GInputStream // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_cret = C.g_input_stream_has_pending(_arg0)
	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsClosed checks if an input stream is closed.
//
// The function returns the following values:
//
//    - ok: TRUE if the stream is closed.
//
func (stream *InputStream) IsClosed() bool {
	var _arg0 *C.GInputStream // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_cret = C.g_input_stream_is_closed(_arg0)
	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Read tries to read count bytes from the stream into the buffer starting at
// buffer. Will block during this read.
//
// If count is zero returns zero and does nothing. A value of count larger than
// G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes read into the buffer is returned. It is not
// an error if this is not the same as the requested size, as it can happen e.g.
// near the end of a file. Zero is returned on end of file (or if count is
// zero), but never otherwise.
//
// The returned buffer is not a nul-terminated string, it can contain nul bytes
// at any position, and this function doesn't nul-terminate the buffer.
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
//    - buffer: a buffer to read data into (which should be at least count bytes
//      long).
//
// The function returns the following values:
//
//    - gssize: number of bytes read, or -1 on error, or 0 on end of file.
//
func (stream *InputStream) Read(ctx context.Context, buffer []byte) (int, error) {
	var _arg0 *C.GInputStream // out
	var _arg3 *C.GCancellable // out
	var _arg1 *C.void         // out
	var _arg2 C.gsize
	var _cret C.gssize  // in
	var _cerr *C.GError // in

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		_arg1 = (*C.void)(unsafe.Pointer(&buffer[0]))
	}

	_cret = C.g_input_stream_read(_arg0, unsafe.Pointer(_arg1), _arg2, _arg3, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(buffer)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gssize, _goerr
}

// ReadAll tries to read count bytes from the stream into the buffer starting at
// buffer. Will block during this read.
//
// This function is similar to g_input_stream_read(), except it tries to read as
// many bytes as requested, only stopping on an error or end of stream.
//
// On a successful read of count bytes, or if we reached the end of the stream,
// TRUE is returned, and bytes_read is set to the number of bytes read into
// buffer.
//
// If there is an error during the operation FALSE is returned and error is set
// to indicate the error status.
//
// As a special exception to the normal conventions for functions that use
// #GError, if this function returns FALSE (and sets error) then bytes_read will
// be set to the number of bytes that were successfully read before the error
// was encountered. This functionality is only available from C. If you need it
// from another language then you must write your own loop around
// g_input_stream_read().
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - buffer: a buffer to read data into (which should be at least count bytes
//      long).
//
// The function returns the following values:
//
//    - bytesRead: location to store the number of bytes that was read from the
//      stream.
//
func (stream *InputStream) ReadAll(ctx context.Context, buffer []byte) (uint, error) {
	var _arg0 *C.GInputStream // out
	var _arg4 *C.GCancellable // out
	var _arg1 *C.void         // out
	var _arg2 C.gsize
	var _arg3 C.gsize   // in
	var _cerr *C.GError // in

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		_arg1 = (*C.void)(unsafe.Pointer(&buffer[0]))
	}

	C.g_input_stream_read_all(_arg0, unsafe.Pointer(_arg1), _arg2, &_arg3, _arg4, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(buffer)

	var _bytesRead uint // out
	var _goerr error    // out

	_bytesRead = uint(_arg3)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _bytesRead, _goerr
}

// ReadAllAsync: request an asynchronous read of count bytes from the stream
// into the buffer starting at buffer.
//
// This is the asynchronous equivalent of g_input_stream_read_all().
//
// Call g_input_stream_read_all_finish() to collect the result.
//
// Any outstanding I/O request with higher priority (lower numerical value) will
// be executed before an outstanding request with lower priority. Default
// priority is G_PRIORITY_DEFAULT.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - buffer: a buffer to read data into (which should be at least count bytes
//      long).
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (stream *InputStream) ReadAllAsync(ctx context.Context, buffer []byte, ioPriority int, callback AsyncReadyCallback) {
	var _arg0 *C.GInputStream // out
	var _arg4 *C.GCancellable // out
	var _arg1 *C.void         // out
	var _arg2 C.gsize
	var _arg3 C.int                 // out
	var _arg5 C.GAsyncReadyCallback // out
	var _arg6 C.gpointer

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		_arg1 = (*C.void)(unsafe.Pointer(&buffer[0]))
	}
	_arg3 = C.int(ioPriority)
	if callback != nil {
		_arg5 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg6 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_input_stream_read_all_async(_arg0, unsafe.Pointer(_arg1), _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// ReadAllFinish finishes an asynchronous stream read operation started with
// g_input_stream_read_all_async().
//
// As a special exception to the normal conventions for functions that use
// #GError, if this function returns FALSE (and sets error) then bytes_read will
// be set to the number of bytes that were successfully read before the error
// was encountered. This functionality is only available from C. If you need it
// from another language then you must write your own loop around
// g_input_stream_read_async().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - bytesRead: location to store the number of bytes that was read from the
//      stream.
//
func (stream *InputStream) ReadAllFinish(result AsyncResulter) (uint, error) {
	var _arg0 *C.GInputStream // out
	var _arg1 *C.GAsyncResult // out
	var _arg2 C.gsize         // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C.g_input_stream_read_all_finish(_arg0, _arg1, &_arg2, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _bytesRead uint // out
	var _goerr error    // out

	_bytesRead = uint(_arg2)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _bytesRead, _goerr
}

// ReadAsync: request an asynchronous read of count bytes from the stream into
// the buffer starting at buffer. When the operation is finished callback will
// be called. You can then call g_input_stream_read_finish() to get the result
// of the operation.
//
// During an async request no other sync and async calls are allowed on stream,
// and will result in G_IO_ERROR_PENDING errors.
//
// A value of count larger than G_MAXSSIZE will cause a
// G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes read into the buffer will be passed to the
// callback. It is not an error if this is not the same as the requested size,
// as it can happen e.g. near the end of a file, but generally we try to read as
// many bytes as requested. Zero is returned on end of file (or if count is
// zero), but never otherwise.
//
// Any outstanding i/o request with higher priority (lower numerical value) will
// be executed before an outstanding request with lower priority. Default
// priority is G_PRIORITY_DEFAULT.
//
// The asynchronous methods have a default fallback that uses threads to
// implement asynchronicity, so they are optional for inheriting classes.
// However, if you override one you must override all.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - buffer: a buffer to read data into (which should be at least count bytes
//      long).
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (stream *InputStream) ReadAsync(ctx context.Context, buffer []byte, ioPriority int, callback AsyncReadyCallback) {
	var _arg0 *C.GInputStream // out
	var _arg4 *C.GCancellable // out
	var _arg1 *C.void         // out
	var _arg2 C.gsize
	var _arg3 C.int                 // out
	var _arg5 C.GAsyncReadyCallback // out
	var _arg6 C.gpointer

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		_arg1 = (*C.void)(unsafe.Pointer(&buffer[0]))
	}
	_arg3 = C.int(ioPriority)
	if callback != nil {
		_arg5 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg6 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_input_stream_read_async(_arg0, unsafe.Pointer(_arg1), _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// ReadBytes: like g_input_stream_read(), this tries to read count bytes from
// the stream in a blocking fashion. However, rather than reading into a
// user-supplied buffer, this will create a new #GBytes containing the data that
// was read. This may be easier to use from language bindings.
//
// If count is zero, returns a zero-length #GBytes and does nothing. A value of
// count larger than G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, a new #GBytes is returned. It is not an error if the size of this
// object is not the same as the requested size, as it can happen e.g. near the
// end of a file. A zero-length #GBytes is returned on end of file (or if count
// is zero), but never otherwise.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned. If an operation was
// partially finished when the operation was cancelled the partial result will
// be returned, without an error.
//
// On error NULL is returned and error is set accordingly.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - count: maximum number of bytes that will be read from the stream. Common
//      values include 4096 and 8192.
//
// The function returns the following values:
//
//    - bytes: new #GBytes, or NULL on error.
//
func (stream *InputStream) ReadBytes(ctx context.Context, count uint) (*glib.Bytes, error) {
	var _arg0 *C.GInputStream // out
	var _arg2 *C.GCancellable // out
	var _arg1 C.gsize         // out
	var _cret *C.GBytes       // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.gsize(count)

	_cret = C.g_input_stream_read_bytes(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(count)

	var _bytes *glib.Bytes // out
	var _goerr error       // out

	_bytes = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _bytes, _goerr
}

// ReadBytesAsync: request an asynchronous read of count bytes from the stream
// into a new #GBytes. When the operation is finished callback will be called.
// You can then call g_input_stream_read_bytes_finish() to get the result of the
// operation.
//
// During an async request no other sync and async calls are allowed on stream,
// and will result in G_IO_ERROR_PENDING errors.
//
// A value of count larger than G_MAXSSIZE will cause a
// G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the new #GBytes will be passed to the callback. It is not an
// error if this is smaller than the requested size, as it can happen e.g. near
// the end of a file, but generally we try to read as many bytes as requested.
// Zero is returned on end of file (or if count is zero), but never otherwise.
//
// Any outstanding I/O request with higher priority (lower numerical value) will
// be executed before an outstanding request with lower priority. Default
// priority is G_PRIORITY_DEFAULT.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - count: number of bytes that will be read from the stream.
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (stream *InputStream) ReadBytesAsync(ctx context.Context, count uint, ioPriority int, callback AsyncReadyCallback) {
	var _arg0 *C.GInputStream       // out
	var _arg3 *C.GCancellable       // out
	var _arg1 C.gsize               // out
	var _arg2 C.int                 // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.gsize(count)
	_arg2 = C.int(ioPriority)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_input_stream_read_bytes_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(count)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// ReadBytesFinish finishes an asynchronous stream read-into-#GBytes operation.
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - bytes: newly-allocated #GBytes, or NULL on error.
//
func (stream *InputStream) ReadBytesFinish(result AsyncResulter) (*glib.Bytes, error) {
	var _arg0 *C.GInputStream // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GBytes       // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_cret = C.g_input_stream_read_bytes_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _bytes *glib.Bytes // out
	var _goerr error       // out

	_bytes = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _bytes, _goerr
}

// ReadFinish finishes an asynchronous stream read operation.
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - gssize: number of bytes read in, or -1 on error, or 0 on end of file.
//
func (stream *InputStream) ReadFinish(result AsyncResulter) (int, error) {
	var _arg0 *C.GInputStream // out
	var _arg1 *C.GAsyncResult // out
	var _cret C.gssize        // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_cret = C.g_input_stream_read_finish(_arg0, _arg1, &_cerr)
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

// SetPending sets stream to have actions pending. If the pending flag is
// already set or stream is closed, it will return FALSE and set error.
func (stream *InputStream) SetPending() error {
	var _arg0 *C.GInputStream // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	C.g_input_stream_set_pending(_arg0, &_cerr)
	runtime.KeepAlive(stream)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Skip tries to skip count bytes from the stream. Will block during the
// operation.
//
// This is identical to g_input_stream_read(), from a behaviour standpoint, but
// the bytes that are skipped are not returned to the user. Some streams have an
// implementation that is more efficient than reading the data.
//
// This function is optional for inherited classes, as the default
// implementation emulates it using read.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned. If an operation was
// partially finished when the operation was cancelled the partial result will
// be returned, without an error.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - count: number of bytes that will be skipped from the stream.
//
// The function returns the following values:
//
//    - gssize: number of bytes skipped, or -1 on error.
//
func (stream *InputStream) Skip(ctx context.Context, count uint) (int, error) {
	var _arg0 *C.GInputStream // out
	var _arg2 *C.GCancellable // out
	var _arg1 C.gsize         // out
	var _cret C.gssize        // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.gsize(count)

	_cret = C.g_input_stream_skip(_arg0, _arg1, _arg2, &_cerr)
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

// SkipAsync: request an asynchronous skip of count bytes from the stream. When
// the operation is finished callback will be called. You can then call
// g_input_stream_skip_finish() to get the result of the operation.
//
// During an async request no other sync and async calls are allowed, and will
// result in G_IO_ERROR_PENDING errors.
//
// A value of count larger than G_MAXSSIZE will cause a
// G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes skipped will be passed to the callback. It is
// not an error if this is not the same as the requested size, as it can happen
// e.g. near the end of a file, but generally we try to skip as many bytes as
// requested. Zero is returned on end of file (or if count is zero), but never
// otherwise.
//
// Any outstanding i/o request with higher priority (lower numerical value) will
// be executed before an outstanding request with lower priority. Default
// priority is G_PRIORITY_DEFAULT.
//
// The asynchronous methods have a default fallback that uses threads to
// implement asynchronicity, so they are optional for inheriting classes.
// However, if you override one, you must override all.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - count: number of bytes that will be skipped from the stream.
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (stream *InputStream) SkipAsync(ctx context.Context, count uint, ioPriority int, callback AsyncReadyCallback) {
	var _arg0 *C.GInputStream       // out
	var _arg3 *C.GCancellable       // out
	var _arg1 C.gsize               // out
	var _arg2 C.int                 // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.gsize(count)
	_arg2 = C.int(ioPriority)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_input_stream_skip_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(count)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// SkipFinish finishes a stream skip operation.
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - gssize: size of the bytes skipped, or -1 on error.
//
func (stream *InputStream) SkipFinish(result AsyncResulter) (int, error) {
	var _arg0 *C.GInputStream // out
	var _arg1 *C.GAsyncResult // out
	var _cret C.gssize        // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_cret = C.g_input_stream_skip_finish(_arg0, _arg1, &_cerr)
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

// InputStreamClass: instance of this type is always passed by reference.
type InputStreamClass struct {
	*inputStreamClass
}

// inputStreamClass is the struct that's finalized.
type inputStreamClass struct {
	native *C.GInputStreamClass
}
