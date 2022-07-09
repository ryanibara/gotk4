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
// #include <glib-object.h>
// extern gboolean _gotk4_gio2_PollableInputStreamInterface_can_poll(void*);
// extern gboolean _gotk4_gio2_PollableInputStreamInterface_is_readable(void*);
// extern void* _gotk4_gio2_PollableInputStreamInterface_create_source(void*, void*);
import "C"

// GTypePollableInputStream returns the GType for the type PollableInputStream.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypePollableInputStream() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "PollableInputStream").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalPollableInputStream)
	return gtype
}

// PollableInputStreamOverrider contains methods that are overridable.
type PollableInputStreamOverrider interface {
	// CanPoll checks if stream is actually pollable. Some classes may implement
	// InputStream but have only certain instances of that class be pollable. If
	// this method returns FALSE, then the behavior of other InputStream methods
	// is undefined.
	//
	// For any given stream, the value returned by this method is constant; a
	// stream cannot switch from pollable to non-pollable or vice versa.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if stream is pollable, FALSE if not.
	//
	CanPoll() bool
	// CreateSource creates a #GSource that triggers when stream can be read, or
	// cancellable is triggered or an error occurs. The callback on the source
	// is of the SourceFunc type.
	//
	// As with g_pollable_input_stream_is_readable(), it is possible that the
	// stream may not actually be readable even after the source triggers, so
	// you should use g_pollable_input_stream_read_nonblocking() rather than
	// g_input_stream_read() from the callback.
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
	// IsReadable checks if stream can be read.
	//
	// Note that some stream types may not be able to implement this 100%
	// reliably, and it is possible that a call to g_input_stream_read() after
	// this returns TRUE would still block. To guarantee non-blocking behavior,
	// you should always use g_pollable_input_stream_read_nonblocking(), which
	// will return a G_IO_ERROR_WOULD_BLOCK error rather than blocking.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if stream is readable, FALSE if not. If an error has
	//      occurred on stream, this will result in
	//      g_pollable_input_stream_is_readable() returning TRUE, and the next
	//      attempt to read will return the error.
	//
	IsReadable() bool
}

// PollableInputStream is implemented by Streams that can be polled for
// readiness to read. This can be used when interfacing with a non-GIO API that
// expects UNIX-file-descriptor-style asynchronous I/O rather than GIO-style.
//
// PollableInputStream wraps an interface. This means the user can get the
// underlying type by calling Cast().
type PollableInputStream struct {
	_ [0]func() // equal guard
	InputStream
}

var (
	_ InputStreamer = (*PollableInputStream)(nil)
)

// PollableInputStreamer describes PollableInputStream's interface methods.
type PollableInputStreamer interface {
	coreglib.Objector

	// CanPoll checks if stream is actually pollable.
	CanPoll() bool
	// CreateSource creates a #GSource that triggers when stream can be read, or
	// cancellable is triggered or an error occurs.
	CreateSource(ctx context.Context) *glib.Source
	// IsReadable checks if stream can be read.
	IsReadable() bool
}

var _ PollableInputStreamer = (*PollableInputStream)(nil)

func ifaceInitPollableInputStreamer(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Gio", "PollableInputStreamInterface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("can_poll"))) = unsafe.Pointer(C._gotk4_gio2_PollableInputStreamInterface_can_poll)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("create_source"))) = unsafe.Pointer(C._gotk4_gio2_PollableInputStreamInterface_create_source)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("is_readable"))) = unsafe.Pointer(C._gotk4_gio2_PollableInputStreamInterface_is_readable)
}

//export _gotk4_gio2_PollableInputStreamInterface_can_poll
func _gotk4_gio2_PollableInputStreamInterface_can_poll(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PollableInputStreamOverrider)

	ok := iface.CanPoll()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_PollableInputStreamInterface_create_source
func _gotk4_gio2_PollableInputStreamInterface_create_source(arg0 *C.void, arg1 *C.void) (cret *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PollableInputStreamOverrider)

	var _cancellable context.Context // out

	if arg1 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg1))
	}

	source := iface.CreateSource(_cancellable)

	cret = (*C.void)(gextras.StructNative(unsafe.Pointer(source)))

	return cret
}

//export _gotk4_gio2_PollableInputStreamInterface_is_readable
func _gotk4_gio2_PollableInputStreamInterface_is_readable(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PollableInputStreamOverrider)

	ok := iface.IsReadable()

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapPollableInputStream(obj *coreglib.Object) *PollableInputStream {
	return &PollableInputStream{
		InputStream: InputStream{
			Object: obj,
		},
	}
}

func marshalPollableInputStream(p uintptr) (interface{}, error) {
	return wrapPollableInputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CanPoll checks if stream is actually pollable. Some classes may implement
// InputStream but have only certain instances of that class be pollable. If
// this method returns FALSE, then the behavior of other InputStream methods is
// undefined.
//
// For any given stream, the value returned by this method is constant; a stream
// cannot switch from pollable to non-pollable or vice versa.
//
// The function returns the following values:
//
//    - ok: TRUE if stream is pollable, FALSE if not.
//
func (stream *PollableInputStream) CanPoll() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_info := girepository.MustFind("Gio", "PollableInputStream")
	_gret := _info.InvokeIfaceMethod("can_poll", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// CreateSource creates a #GSource that triggers when stream can be read, or
// cancellable is triggered or an error occurs. The callback on the source is of
// the SourceFunc type.
//
// As with g_pollable_input_stream_is_readable(), it is possible that the stream
// may not actually be readable even after the source triggers, so you should
// use g_pollable_input_stream_read_nonblocking() rather than
// g_input_stream_read() from the callback.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//
// The function returns the following values:
//
//    - source: new #GSource.
//
func (stream *PollableInputStream) CreateSource(ctx context.Context) *glib.Source {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[1] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}

	_info := girepository.MustFind("Gio", "PollableInputStream")
	_gret := _info.InvokeIfaceMethod("create_source", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)

	var _source *glib.Source // out

	_source = (*glib.Source)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_source)),
		func(intern *struct{ C unsafe.Pointer }) {
			{
				args := [1]girepository.Argument{(*C.void)(intern.C)}
				girepository.MustFind("GLib", "Source").InvokeRecordMethod("free", args[:], nil)
			}
		},
	)

	return _source
}

// IsReadable checks if stream can be read.
//
// Note that some stream types may not be able to implement this 100% reliably,
// and it is possible that a call to g_input_stream_read() after this returns
// TRUE would still block. To guarantee non-blocking behavior, you should always
// use g_pollable_input_stream_read_nonblocking(), which will return a
// G_IO_ERROR_WOULD_BLOCK error rather than blocking.
//
// The function returns the following values:
//
//    - ok: TRUE if stream is readable, FALSE if not. If an error has occurred on
//      stream, this will result in g_pollable_input_stream_is_readable()
//      returning TRUE, and the next attempt to read will return the error.
//
func (stream *PollableInputStream) IsReadable() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_info := girepository.MustFind("Gio", "PollableInputStream")
	_gret := _info.InvokeIfaceMethod("is_readable", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}
