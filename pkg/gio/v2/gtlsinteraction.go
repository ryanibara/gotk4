// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gio2_AsyncReadyCallback(void*, void*, gpointer);
import "C"

// GTypeTLSInteraction returns the GType for the type TLSInteraction.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTLSInteraction() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "TlsInteraction").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalTLSInteraction)
	return gtype
}

// TLSInteractionOverrider contains methods that are overridable.
type TLSInteractionOverrider interface {
}

// TLSInteraction provides a mechanism for the TLS connection and database code
// to interact with the user. It can be used to ask the user for passwords.
//
// To use a Interaction with a TLS connection use
// g_tls_connection_set_interaction().
//
// Callers should instantiate a derived class that implements the various
// interaction methods to show the required dialogs.
//
// Callers should use the 'invoke' functions like
// g_tls_interaction_invoke_ask_password() to run interaction methods. These
// functions make sure that the interaction is invoked in the main loop and not
// in the current thread, if the current thread is not running the main loop.
//
// Derived classes can choose to implement whichever interactions methods they'd
// like to support by overriding those virtual methods in their class
// initialization function. Any interactions not implemented will return
// G_TLS_INTERACTION_UNHANDLED. If a derived class implements an async method,
// it must also implement the corresponding finish method.
type TLSInteraction struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TLSInteraction)(nil)
)

func classInitTLSInteractioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTLSInteraction(obj *coreglib.Object) *TLSInteraction {
	return &TLSInteraction{
		Object: obj,
	}
}

func marshalTLSInteraction(p uintptr) (interface{}, error) {
	return wrapTLSInteraction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AskPasswordAsync: run asynchronous interaction to ask the user for a
// password. In general, g_tls_interaction_invoke_ask_password() should be used
// instead of this function.
//
// Derived subclasses usually implement a password prompt, although they may
// also choose to provide a password from elsewhere. The password value will be
// filled in and then callback will be called. Alternatively the user may abort
// this password request, which will usually abort the TLS connection.
//
// If the interaction is cancelled by the cancellation object, or by the user
// then G_TLS_INTERACTION_FAILED will be returned with an error that contains a
// G_IO_ERROR_CANCELLED error code. Certain implementations may not support
// immediate cancellation.
//
// Certain implementations may not support immediate cancellation.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable cancellation object.
//    - password: Password object.
//    - callback (optional) will be called when the interaction completes.
//
func (interaction *TLSInteraction) AskPasswordAsync(ctx context.Context, password *TLSPassword, callback AsyncReadyCallback) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(interaction).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(password).Native()))
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[3])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[4] = C.gpointer(gbox.AssignOnce(callback))
	}

	girepository.MustFind("Gio", "TlsInteraction").InvokeMethod("ask_password_async", _args[:], nil)

	runtime.KeepAlive(interaction)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(password)
	runtime.KeepAlive(callback)
}
