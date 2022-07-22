// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_CancellableClass_cancelled(GCancellable*);
// extern void _gotk4_gio2_Cancellable_ConnectCancelled(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeCancellable = coreglib.Type(C.g_cancellable_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCancellable, F: marshalCancellable},
	})
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

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeCancellable,
		GoType:        reflect.TypeOf((*Cancellable)(nil)),
		InitClass:     initClassCancellable,
		FinalizeClass: finalizeClassCancellable,
	})
}

func initClassCancellable(gclass unsafe.Pointer, goval any) {

	pclass := (*C.GCancellableClass)(unsafe.Pointer(gclass))

	if _, ok := goval.(interface{ Cancelled() }); ok {
		pclass.cancelled = (*[0]byte)(C._gotk4_gio2_CancellableClass_cancelled)
	}
	if goval, ok := goval.(interface{ InitCancellable(*CancellableClass) }); ok {
		klass := (*CancellableClass)(gextras.NewStructNative(gclass))
		goval.InitCancellable(klass)
	}
}

func finalizeClassCancellable(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ FinalizeCancellable(*CancellableClass) }); ok {
		klass := (*CancellableClass)(gextras.NewStructNative(gclass))
		goval.FinalizeCancellable(klass)
	}
}

//export _gotk4_gio2_CancellableClass_cancelled
func _gotk4_gio2_CancellableClass_cancelled(arg0 *C.GCancellable) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
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
	var _cret *C.GCancellable // in

	_cret = C.g_cancellable_new()

	var _cancellable *Cancellable // out

	_cancellable = wrapCancellable(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

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
	var _arg0 *C.GCancellable // out

	if cancellable != nil {
		_arg0 = (*C.GCancellable)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	C.g_cancellable_cancel(_arg0)
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
	var _arg0 *C.GCancellable // out
	var _arg1 C.gulong        // out

	if cancellable != nil {
		_arg0 = (*C.GCancellable)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}
	_arg1 = C.gulong(handlerId)

	C.g_cancellable_disconnect(_arg0, _arg1)
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
func (cancellable *Cancellable) Fd() int {
	var _arg0 *C.GCancellable // out
	var _cret C.int           // in

	if cancellable != nil {
		_arg0 = (*C.GCancellable)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	_cret = C.g_cancellable_get_fd(_arg0)
	runtime.KeepAlive(cancellable)

	var _gint int // out

	_gint = int(_cret)

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
	var _arg0 *C.GCancellable // out
	var _cret C.gboolean      // in

	if cancellable != nil {
		_arg0 = (*C.GCancellable)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	_cret = C.g_cancellable_is_cancelled(_arg0)
	runtime.KeepAlive(cancellable)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PopCurrent pops cancellable off the cancellable stack (verifying that
// cancellable is on the top of the stack).
func (cancellable *Cancellable) PopCurrent() {
	var _arg0 *C.GCancellable // out

	if cancellable != nil {
		_arg0 = (*C.GCancellable)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	C.g_cancellable_pop_current(_arg0)
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
	var _arg0 *C.GCancellable // out

	if cancellable != nil {
		_arg0 = (*C.GCancellable)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	C.g_cancellable_push_current(_arg0)
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
	var _arg0 *C.GCancellable // out

	if cancellable != nil {
		_arg0 = (*C.GCancellable)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	C.g_cancellable_release_fd(_arg0)
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
	var _arg0 *C.GCancellable // out

	if cancellable != nil {
		_arg0 = (*C.GCancellable)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	C.g_cancellable_reset(_arg0)
	runtime.KeepAlive(cancellable)
}

// SetErrorIfCancelled: if the cancellable is cancelled, sets the error to
// notify that the operation was cancelled.
func (cancellable *Cancellable) SetErrorIfCancelled() error {
	var _arg0 *C.GCancellable // out
	var _cerr *C.GError       // in

	if cancellable != nil {
		_arg0 = (*C.GCancellable)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	C.g_cancellable_set_error_if_cancelled(_arg0, &_cerr)
	runtime.KeepAlive(cancellable)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
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
	var _arg0 *C.GCancellable // out
	var _cret *C.GSource      // in

	if cancellable != nil {
		_arg0 = (*C.GCancellable)(unsafe.Pointer(coreglib.InternObject(cancellable).Native()))
	}

	_cret = C.g_cancellable_source_new(_arg0)
	runtime.KeepAlive(cancellable)

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

// CancellableGetCurrent gets the top cancellable from the stack.
//
// The function returns the following values:
//
//    - cancellable (optional) from the top of the stack, or NULL if the stack is
//      empty.
//
func CancellableGetCurrent() *Cancellable {
	var _cret *C.GCancellable // in

	_cret = C.g_cancellable_get_current()

	var _cancellable *Cancellable // out

	if _cret != nil {
		_cancellable = wrapCancellable(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _cancellable
}

// CancellableClass: instance of this type is always passed by reference.
type CancellableClass struct {
	*cancellableClass
}

// cancellableClass is the struct that's finalized.
type cancellableClass struct {
	native *C.GCancellableClass
}
