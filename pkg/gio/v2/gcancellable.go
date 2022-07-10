// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_CancellableClass_cancelled(void*);
// extern void _gotk4_gio2_Cancellable_ConnectCancelled(gpointer, guintptr);
import "C"

// GTypeCancellable returns the GType for the type Cancellable.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCancellable() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "Cancellable").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalCancellable)
	return gtype
}

// CancellableOverrider contains methods that are overridable.
type CancellableOverrider interface {
	Cancelled()
}

// Cancellable is a thread-safe operation cancellation stack used throughout GIO
// to allow for cancellation of synchronous and asynchronous operations.
type Cancellable struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Cancellable)(nil)
)

func classInitCancellabler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gio", "CancellableClass")

	if _, ok := goval.(interface{ Cancelled() }); ok {
		o := pclass.StructFieldOffset("cancelled")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_CancellableClass_cancelled)
	}
}

//export _gotk4_gio2_CancellableClass_cancelled
func _gotk4_gio2_CancellableClass_cancelled(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Cancelled() })

	iface.Cancelled()
}

func wrapCancellable(obj *coreglib.Object) *Cancellable {
	return &Cancellable{
		Object: obj,
	}
}

func marshalCancellable(p uintptr) (interface{}, error) {
	return wrapCancellable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_Cancellable_ConnectCancelled
func _gotk4_gio2_Cancellable_ConnectCancelled(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectCancelled is emitted when the operation has been cancelled.
//
// Can be used by implementations of cancellable operations. If the operation is
// cancelled from another thread, the signal will be emitted in the thread that
// cancelled the operation, not the thread that is running the operation.
//
// Note that disconnecting from this signal (or any signal) in a multi-threaded
// program is prone to race conditions. For instance it is possible that a
// signal handler may be invoked even after a call to
// g_signal_handler_disconnect() for that handler has already returned.
//
// There is also a problem when cancellation happens right before connecting to
// the signal. If this happens the signal will unexpectedly not be emitted, and
// checking before connecting to the signal leaves a race condition where this
// is still happening.
//
// In order to make it safe and easy to connect handlers there are two helper
// functions: g_cancellable_connect() and g_cancellable_disconnect() which
// protect against problems like this.
//
// An example of how to us this:
//
//        // Make sure we don't do unnecessary work if already cancelled
//        if (g_cancellable_set_error_if_cancelled (cancellable, error))
//          return;
//
//        // Set up all the data needed to be able to handle cancellation
//        // of the operation
//        my_data = my_data_new (...);
//
//        id = 0;
//        if (cancellable)
//          id = g_cancellable_connect (cancellable,
//        			      G_CALLBACK (cancelled_handler)
//        			      data, NULL);
//
//        // cancellable operation here...
//
//        g_cancellable_disconnect (cancellable, id);
//
//        // cancelled_handler is never called after this, it is now safe
//        // to free the data
//        my_data_free (my_data);
//
// Note that the cancelled signal is emitted in the thread that the user
// cancelled from, which may be the main thread. So, the cancellable signal
// should not do something that can block.
func (cancellable *Cancellable) ConnectCancelled(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(cancellable, "cancelled", false, unsafe.Pointer(C._gotk4_gio2_Cancellable_ConnectCancelled), f)
}

// NewCancellable creates a new #GCancellable object.
//
// Applications that want to start one or more operations that should be
// cancellable should create a #GCancellable and pass it to the operations.
//
// One #GCancellable can be used in multiple consecutive operations or in
// multiple concurrent operations.
//
// The function returns the following values:
//
//    - cancellable: #GCancellable.
//
func NewCancellable() *Cancellable {
	_info := girepository.MustFind("Gio", "Cancellable")
	_gret := _info.InvokeClassMethod("new_Cancellable", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _cancellable *Cancellable // out

	_cancellable = wrapCancellable(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _cancellable
}

// Cancel will set cancellable to cancelled, and will emit the
// #GCancellable::cancelled signal. (However, see the warning about race
// conditions in the documentation for that signal if you are planning to
// connect to it.)
//
// This function is thread-safe. In other words, you can safely call it from a
// thread other than the one running the operation that was passed the
// cancellable.
//
// If cancellable is NULL, this function returns immediately for convenience.
//
// The convention within GIO is that cancelling an asynchronous operation causes
// it to complete asynchronously. That is, if you cancel the operation from the
// same thread in which it is running, then the operation's ReadyCallback will
// not be invoked until the application returns to the main loop.
func (cancellable *Cancellable) Cancel() {
	var _args [1]girepository.Argument

	if cancellable != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	_info := girepository.MustFind("Gio", "Cancellable")
	_info.InvokeClassMethod("cancel", _args[:], nil)

	runtime.KeepAlive(cancellable)
}

// Disconnect disconnects a handler from a cancellable instance similar to
// g_signal_handler_disconnect(). Additionally, in the event that a signal
// handler is currently running, this call will block until the handler has
// finished. Calling this function from a #GCancellable::cancelled signal
// handler will therefore result in a deadlock.
//
// This avoids a race condition where a thread cancels at the same time as the
// cancellable operation is finished and the signal handler is removed. See
// #GCancellable::cancelled for details on how to use this.
//
// If cancellable is NULL or handler_id is 0 this function does nothing.
//
// The function takes the following parameters:
//
//    - handlerId: handler id of the handler to be disconnected, or 0.
//
func (cancellable *Cancellable) Disconnect(handlerId uint32) {
	var _args [2]girepository.Argument

	if cancellable != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}
	*(*C.gulong)(unsafe.Pointer(&_args[1])) = C.gulong(handlerId)

	_info := girepository.MustFind("Gio", "Cancellable")
	_info.InvokeClassMethod("disconnect", _args[:], nil)

	runtime.KeepAlive(cancellable)
	runtime.KeepAlive(handlerId)
}

// Fd gets the file descriptor for a cancellable job. This can be used to
// implement cancellable operations on Unix systems. The returned fd will turn
// readable when cancellable is cancelled.
//
// You are not supposed to read from the fd yourself, just check for readable
// status. Reading to unset the readable status is done with
// g_cancellable_reset().
//
// After a successful return from this function, you should use
// g_cancellable_release_fd() to free up resources allocated for the returned
// file descriptor.
//
// See also g_cancellable_make_pollfd().
//
// The function returns the following values:
//
//    - gint: valid file descriptor. -1 if the file descriptor is not supported,
//      or on errors.
//
func (cancellable *Cancellable) Fd() int32 {
	var _args [1]girepository.Argument

	if cancellable != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	_info := girepository.MustFind("Gio", "Cancellable")
	_gret := _info.InvokeClassMethod("get_fd", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cancellable)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// IsCancelled checks if a cancellable job has been cancelled.
//
// The function returns the following values:
//
//    - ok: TRUE if cancellable is cancelled, FALSE if called with NULL or if
//      item is not cancelled.
//
func (cancellable *Cancellable) IsCancelled() bool {
	var _args [1]girepository.Argument

	if cancellable != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	_info := girepository.MustFind("Gio", "Cancellable")
	_gret := _info.InvokeClassMethod("is_cancelled", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cancellable)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// PopCurrent pops cancellable off the cancellable stack (verifying that
// cancellable is on the top of the stack).
func (cancellable *Cancellable) PopCurrent() {
	var _args [1]girepository.Argument

	if cancellable != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	_info := girepository.MustFind("Gio", "Cancellable")
	_info.InvokeClassMethod("pop_current", _args[:], nil)

	runtime.KeepAlive(cancellable)
}

// PushCurrent pushes cancellable onto the cancellable stack. The current
// cancellable can then be received using g_cancellable_get_current().
//
// This is useful when implementing cancellable operations in code that does not
// allow you to pass down the cancellable object.
//
// This is typically called automatically by e.g. #GFile operations, so you
// rarely have to call this yourself.
func (cancellable *Cancellable) PushCurrent() {
	var _args [1]girepository.Argument

	if cancellable != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	_info := girepository.MustFind("Gio", "Cancellable")
	_info.InvokeClassMethod("push_current", _args[:], nil)

	runtime.KeepAlive(cancellable)
}

// ReleaseFd releases a resources previously allocated by g_cancellable_get_fd()
// or g_cancellable_make_pollfd().
//
// For compatibility reasons with older releases, calling this function is not
// strictly required, the resources will be automatically freed when the
// cancellable is finalized. However, the cancellable will block scarce file
// descriptors until it is finalized if this function is not called. This can
// cause the application to run out of file descriptors when many #GCancellables
// are used at the same time.
func (cancellable *Cancellable) ReleaseFd() {
	var _args [1]girepository.Argument

	if cancellable != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	_info := girepository.MustFind("Gio", "Cancellable")
	_info.InvokeClassMethod("release_fd", _args[:], nil)

	runtime.KeepAlive(cancellable)
}

// Reset resets cancellable to its uncancelled state.
//
// If cancellable is currently in use by any cancellable operation then the
// behavior of this function is undefined.
//
// Note that it is generally not a good idea to reuse an existing cancellable
// for more operations after it has been cancelled once, as this function might
// tempt you to do. The recommended practice is to drop the reference to a
// cancellable after cancelling it, and let it die with the outstanding async
// operations. You should create a fresh cancellable for further async
// operations.
func (cancellable *Cancellable) Reset() {
	var _args [1]girepository.Argument

	if cancellable != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	_info := girepository.MustFind("Gio", "Cancellable")
	_info.InvokeClassMethod("reset", _args[:], nil)

	runtime.KeepAlive(cancellable)
}

// SetErrorIfCancelled: if the cancellable is cancelled, sets the error to
// notify that the operation was cancelled.
func (cancellable *Cancellable) SetErrorIfCancelled() error {
	var _args [1]girepository.Argument

	if cancellable != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	_info := girepository.MustFind("Gio", "Cancellable")
	_info.InvokeClassMethod("set_error_if_cancelled", _args[:], nil)

	runtime.KeepAlive(cancellable)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cerr))))
	}

	return _goerr
}

// NewSource creates a source that triggers if cancellable is cancelled and
// calls its callback of type SourceFunc. This is primarily useful for attaching
// to another (non-cancellable) source with g_source_add_child_source() to add
// cancellability to it.
//
// For convenience, you can call this with a NULL #GCancellable, in which case
// the source will never trigger.
//
// The new #GSource will hold a reference to the #GCancellable.
//
// The function returns the following values:
//
//    - source: new #GSource.
//
func (cancellable *Cancellable) NewSource() *glib.Source {
	var _args [1]girepository.Argument

	if cancellable != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	_info := girepository.MustFind("Gio", "Cancellable")
	_gret := _info.InvokeClassMethod("source_new", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cancellable)

	var _source *glib.Source // out

	_source = (*glib.Source)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_source)),
		func(intern *struct{ C unsafe.Pointer }) {
			{
				var args [1]girepository.Argument
				*(*unsafe.Pointer)(unsafe.Pointer(&args[0])) = unsafe.Pointer(intern.C)
				girepository.MustFind("GLib", "Source").InvokeRecordMethod("free", args[:], nil)
			}
		},
	)

	return _source
}

// CancellableGetCurrent gets the top cancellable from the stack.
//
// The function returns the following values:
//
//    - cancellable (optional) from the top of the stack, or NULL if the stack is
//      empty.
//
func CancellableGetCurrent() *Cancellable {
	_info := girepository.MustFind("Gio", "get_current")
	_gret := _info.InvokeFunction(nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _cancellable *Cancellable // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_cancellable = wrapCancellable(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	}

	return _cancellable
}
