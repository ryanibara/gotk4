// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern GSource* _gotk4_gio2_PollableOutputStreamInterface_create_source(GPollableOutputStream*, GCancellable*);
// extern gboolean _gotk4_gio2_PollableOutputStreamInterface_can_poll(GPollableOutputStream*);
// extern gboolean _gotk4_gio2_PollableOutputStreamInterface_is_writable(GPollableOutputStream*);
import "C"

// glib.Type values for gpollableoutputstream.go.
var GTypePollableOutputStream = coreglib.Type(C.g_pollable_output_stream_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypePollableOutputStream, F: marshalPollableOutputStream},
	})
}

// PollableOutputStreamOverrider contains methods that are overridable.
type PollableOutputStreamOverrider interface {
	// CanPoll checks if stream is actually pollable. Some classes may implement
	// OutputStream but have only certain instances of that class be pollable.
	// If this method returns FALSE, then the behavior of other OutputStream
	// methods is undefined.
	//
	// For any given stream, the value returned by this method is constant; a
	// stream cannot switch from pollable to non-pollable or vice versa.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if stream is pollable, FALSE if not.
	//
	CanPoll() bool
	// CreateSource creates a #GSource that triggers when stream can be written,
	// or cancellable is triggered or an error occurs. The callback on the
	// source is of the SourceFunc type.
	//
	// As with g_pollable_output_stream_is_writable(), it is possible that the
	// stream may not actually be writable even after the source triggers, so
	// you should use g_pollable_output_stream_write_nonblocking() rather than
	// g_output_stream_write() from the callback.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional) or NULL.
	//
	// The function returns the following values:
	//
	//    - source: new #GSource.
	//
	CreateSource(ctx context.Context) *glib.Source
	// IsWritable checks if stream can be written.
	//
	// Note that some stream types may not be able to implement this 100%
	// reliably, and it is possible that a call to g_output_stream_write() after
	// this returns TRUE would still block. To guarantee non-blocking behavior,
	// you should always use g_pollable_output_stream_write_nonblocking(), which
	// will return a G_IO_ERROR_WOULD_BLOCK error rather than blocking.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if stream is writable, FALSE if not. If an error has
	//      occurred on stream, this will result in
	//      g_pollable_output_stream_is_writable() returning TRUE, and the next
	//      attempt to write will return the error.
	//
	IsWritable() bool
}

// PollableOutputStream is implemented by Streams that can be polled for
// readiness to write. This can be used when interfacing with a non-GIO API that
// expects UNIX-file-descriptor-style asynchronous I/O rather than GIO-style.
//
// PollableOutputStream wraps an interface. This means the user can get the
// underlying type by calling Cast().
type PollableOutputStream struct {
	_ [0]func() // equal guard
	OutputStream
}

var (
	_ OutputStreamer = (*PollableOutputStream)(nil)
)

// PollableOutputStreamer describes PollableOutputStream's interface methods.
type PollableOutputStreamer interface {
	coreglib.Objector

	// CanPoll checks if stream is actually pollable.
	CanPoll() bool
	// CreateSource creates a #GSource that triggers when stream can be written,
	// or cancellable is triggered or an error occurs.
	CreateSource(ctx context.Context) *glib.Source
	// IsWritable checks if stream can be written.
	IsWritable() bool
}

var _ PollableOutputStreamer = (*PollableOutputStream)(nil)

func ifaceInitPollableOutputStreamer(gifacePtr, data C.gpointer) {
	iface := (*C.GPollableOutputStreamInterface)(unsafe.Pointer(gifacePtr))
	iface.can_poll = (*[0]byte)(C._gotk4_gio2_PollableOutputStreamInterface_can_poll)
	iface.create_source = (*[0]byte)(C._gotk4_gio2_PollableOutputStreamInterface_create_source)
	iface.is_writable = (*[0]byte)(C._gotk4_gio2_PollableOutputStreamInterface_is_writable)
}

//export _gotk4_gio2_PollableOutputStreamInterface_can_poll
func _gotk4_gio2_PollableOutputStreamInterface_can_poll(arg0 *C.GPollableOutputStream) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PollableOutputStreamOverrider)

	ok := iface.CanPoll()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_PollableOutputStreamInterface_create_source
func _gotk4_gio2_PollableOutputStreamInterface_create_source(arg0 *C.GPollableOutputStream, arg1 *C.GCancellable) (cret *C.GSource) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PollableOutputStreamOverrider)

	var _cancellable context.Context // out

	if arg1 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg1))
	}

	source := iface.CreateSource(_cancellable)

	cret = (*C.void)(gextras.StructNative(unsafe.Pointer(source)))

	return cret
}

//export _gotk4_gio2_PollableOutputStreamInterface_is_writable
func _gotk4_gio2_PollableOutputStreamInterface_is_writable(arg0 *C.GPollableOutputStream) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PollableOutputStreamOverrider)

	ok := iface.IsWritable()

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapPollableOutputStream(obj *coreglib.Object) *PollableOutputStream {
	return &PollableOutputStream{
		OutputStream: OutputStream{
			Object: obj,
		},
	}
}

func marshalPollableOutputStream(p uintptr) (interface{}, error) {
	return wrapPollableOutputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CanPoll checks if stream is actually pollable. Some classes may implement
// OutputStream but have only certain instances of that class be pollable. If
// this method returns FALSE, then the behavior of other OutputStream methods is
// undefined.
//
// For any given stream, the value returned by this method is constant; a stream
// cannot switch from pollable to non-pollable or vice versa.
//
// The function returns the following values:
//
//    - ok: TRUE if stream is pollable, FALSE if not.
//
func (stream *PollableOutputStream) CanPoll() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**PollableOutputStream)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CreateSource creates a #GSource that triggers when stream can be written, or
// cancellable is triggered or an error occurs. The callback on the source is of
// the SourceFunc type.
//
// As with g_pollable_output_stream_is_writable(), it is possible that the
// stream may not actually be writable even after the source triggers, so you
// should use g_pollable_output_stream_write_nonblocking() rather than
// g_output_stream_write() from the callback.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//
// The function returns the following values:
//
//    - source: new #GSource.
//
func (stream *PollableOutputStream) CreateSource(ctx context.Context) *glib.Source {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**PollableOutputStream)(unsafe.Pointer(&args[1])) = _arg1

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)

	var _source *glib.Source // out

	_source = (*glib.Source)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_source)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_source_unref((*C.GSource)(intern.C))
		},
	)

	return _source
}

// IsWritable checks if stream can be written.
//
// Note that some stream types may not be able to implement this 100% reliably,
// and it is possible that a call to g_output_stream_write() after this returns
// TRUE would still block. To guarantee non-blocking behavior, you should always
// use g_pollable_output_stream_write_nonblocking(), which will return a
// G_IO_ERROR_WOULD_BLOCK error rather than blocking.
//
// The function returns the following values:
//
//    - ok: TRUE if stream is writable, FALSE if not. If an error has occurred on
//      stream, this will result in g_pollable_output_stream_is_writable()
//      returning TRUE, and the next attempt to write will return the error.
//
func (stream *PollableOutputStream) IsWritable() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**PollableOutputStream)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
