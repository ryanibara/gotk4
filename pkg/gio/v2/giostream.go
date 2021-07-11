// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
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
		{T: externglib.Type(C.g_io_stream_get_type()), F: marshalIOStreamer},
	})
}

// IOStreamOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type IOStreamOverrider interface {
	// CloseAsync requests an asynchronous close of the stream, releasing
	// resources related to it. When the operation is finished @callback will be
	// called. You can then call g_io_stream_close_finish() to get the result of
	// the operation.
	//
	// For behaviour details see g_io_stream_close().
	//
	// The asynchronous methods have a default fallback that uses threads to
	// implement asynchronicity, so they are optional for inheriting classes.
	// However, if you override one you must override all.
	CloseAsync(ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback)
	// CloseFinish closes a stream.
	CloseFinish(result AsyncResulter) error

	CloseFn(cancellable Cancellabler) error
	// InputStream gets the input stream for this object. This is used for
	// reading.
	InputStream() *InputStream
	// OutputStream gets the output stream for this object. This is used for
	// writing.
	OutputStream() *OutputStream
}

// IOStreamer describes IOStream's methods.
type IOStreamer interface {
	// ClearPending clears the pending flag on @stream.
	ClearPending()
	// Close closes the stream, releasing resources related to it.
	Close(cancellable Cancellabler) error
	// CloseAsync requests an asynchronous close of the stream, releasing
	// resources related to it.
	CloseAsync(ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback)
	// CloseFinish closes a stream.
	CloseFinish(result AsyncResulter) error
	// InputStream gets the input stream for this object.
	InputStream() *InputStream
	// OutputStream gets the output stream for this object.
	OutputStream() *OutputStream
	// HasPending checks if a stream has pending actions.
	HasPending() bool
	// IsClosed checks if a stream is closed.
	IsClosed() bool
	// SetPending sets @stream to have actions pending.
	SetPending() error
}

// IOStream represents an object that has both read and write streams. Generally
// the two streams act as separate input and output streams, but they share some
// common resources and state. For instance, for seekable streams, both streams
// may use the same position.
//
// Examples of OStream objects are Connection, which represents a two-way
// network connection; and IOStream, which represents a file handle opened in
// read-write mode.
//
// To do the actual reading and writing you need to get the substreams with
// g_io_stream_get_input_stream() and g_io_stream_get_output_stream().
//
// The OStream object owns the input and the output streams, not the other way
// around, so keeping the substreams alive will not keep the OStream object
// alive. If the OStream object is freed it will be closed, thus closing the
// substreams, so even if the substreams stay alive they will always return
// G_IO_ERROR_CLOSED for all operations.
//
// To close a stream use g_io_stream_close() which will close the common stream
// object and also the individual substreams. You can also close the substreams
// themselves. In most cases this only marks the substream as closed, so further
// I/O on it fails but common state in the OStream may still be open. However,
// some streams may support "half-closed" states where one direction of the
// stream is actually shut down.
//
// Operations on OStreams cannot be started while another operation on the
// OStream or its substreams is in progress. Specifically, an application can
// read from the Stream and write to the Stream simultaneously (either in
// separate threads, or as asynchronous operations in the same thread), but an
// application cannot start any OStream operation while there is a OStream,
// Stream or Stream operation in progress, and an application can’t start any
// Stream or Stream operation while there is a OStream operation in progress.
//
// This is a product of individual stream operations being associated with a
// given Context (the thread-default context at the time the operation was
// started), rather than entire streams being associated with a single Context.
//
// GIO may run operations on OStreams from other (worker) threads, and this may
// be exposed to application code in the behaviour of wrapper streams, such as
// InputStream or Connection. With such wrapper APIs, application code may only
// run operations on the base (wrapped) stream when the wrapper stream is idle.
// Note that the semantics of such operations may not be well-defined due to the
// state the wrapper stream leaves the base stream in (though they are
// guaranteed not to crash).
type IOStream struct {
	*externglib.Object
}

var (
	_ IOStreamer      = (*IOStream)(nil)
	_ gextras.Nativer = (*IOStream)(nil)
)

func wrapIOStream(obj *externglib.Object) IOStreamer {
	return &IOStream{
		Object: obj,
	}
}

func marshalIOStreamer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapIOStream(obj), nil
}

// ClearPending clears the pending flag on @stream.
func (stream *IOStream) ClearPending() {
	var _arg0 *C.GIOStream // out

	_arg0 = (*C.GIOStream)(unsafe.Pointer(stream.Native()))

	C.g_io_stream_clear_pending(_arg0)
}

// Close closes the stream, releasing resources related to it. This will also
// close the individual input and output streams, if they are not already
// closed.
//
// Once the stream is closed, all other operations will return
// G_IO_ERROR_CLOSED. Closing a stream multiple times will not return an error.
//
// Closing a stream will automatically flush any outstanding buffers in the
// stream.
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
// check and report the error to the user, otherwise there might be a loss of
// data as all data might not be written.
//
// If @cancellable is not NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation was
// cancelled, the error G_IO_ERROR_CANCELLED will be returned. Cancelling a
// close will still leave the stream closed, but some streams can use a faster
// close that doesn't block to e.g. check errors.
//
// The default implementation of this method just calls close on the individual
// input/output streams.
func (stream *IOStream) Close(cancellable Cancellabler) error {
	var _arg0 *C.GIOStream    // out
	var _arg1 *C.GCancellable // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GIOStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	C.g_io_stream_close(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// CloseAsync requests an asynchronous close of the stream, releasing resources
// related to it. When the operation is finished @callback will be called. You
// can then call g_io_stream_close_finish() to get the result of the operation.
//
// For behaviour details see g_io_stream_close().
//
// The asynchronous methods have a default fallback that uses threads to
// implement asynchronicity, so they are optional for inheriting classes.
// However, if you override one you must override all.
func (stream *IOStream) CloseAsync(ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback) {
	var _arg0 *C.GIOStream          // out
	var _arg1 C.int                 // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GIOStream)(unsafe.Pointer(stream.Native()))
	_arg1 = C.int(ioPriority)
	_arg2 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_io_stream_close_async(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// CloseFinish closes a stream.
func (stream *IOStream) CloseFinish(result AsyncResulter) error {
	var _arg0 *C.GIOStream    // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GIOStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	C.g_io_stream_close_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// InputStream gets the input stream for this object. This is used for reading.
func (stream *IOStream) InputStream() *InputStream {
	var _arg0 *C.GIOStream    // out
	var _cret *C.GInputStream // in

	_arg0 = (*C.GIOStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_io_stream_get_input_stream(_arg0)

	var _inputStream *InputStream // out

	_inputStream = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*InputStream)

	return _inputStream
}

// OutputStream gets the output stream for this object. This is used for
// writing.
func (stream *IOStream) OutputStream() *OutputStream {
	var _arg0 *C.GIOStream     // out
	var _cret *C.GOutputStream // in

	_arg0 = (*C.GIOStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_io_stream_get_output_stream(_arg0)

	var _outputStream *OutputStream // out

	_outputStream = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*OutputStream)

	return _outputStream
}

// HasPending checks if a stream has pending actions.
func (stream *IOStream) HasPending() bool {
	var _arg0 *C.GIOStream // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GIOStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_io_stream_has_pending(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsClosed checks if a stream is closed.
func (stream *IOStream) IsClosed() bool {
	var _arg0 *C.GIOStream // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GIOStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_io_stream_is_closed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetPending sets @stream to have actions pending. If the pending flag is
// already set or @stream is closed, it will return false and set @error.
func (stream *IOStream) SetPending() error {
	var _arg0 *C.GIOStream // out
	var _cerr *C.GError    // in

	_arg0 = (*C.GIOStream)(unsafe.Pointer(stream.Native()))

	C.g_io_stream_set_pending(_arg0, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
