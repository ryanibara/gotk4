// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern GTlsInteractionResult _gotk4_gio2_TlsInteractionClass_ask_password(GTlsInteraction*, GTlsPassword*, GCancellable*, GError**);
// extern GTlsInteractionResult _gotk4_gio2_TlsInteractionClass_ask_password_finish(GTlsInteraction*, GAsyncResult*, GError**);
// extern GTlsInteractionResult _gotk4_gio2_TlsInteractionClass_request_certificate(GTlsInteraction*, GTlsConnection*, GTlsCertificateRequestFlags, GCancellable*, GError**);
// extern GTlsInteractionResult _gotk4_gio2_TlsInteractionClass_request_certificate_finish(GTlsInteraction*, GAsyncResult*, GError**);
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// glib.Type values for gtlsinteraction.go.
var GTypeTLSInteraction = externglib.Type(C.g_tls_interaction_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeTLSInteraction, F: marshalTLSInteraction},
	})
}

// TLSInteractionOverrider contains methods that are overridable.
type TLSInteractionOverrider interface {
	externglib.Objector
	// AskPassword: run synchronous interaction to ask the user for a password.
	// In general, g_tls_interaction_invoke_ask_password() should be used
	// instead of this function.
	//
	// Derived subclasses usually implement a password prompt, although they may
	// also choose to provide a password from elsewhere. The password value will
	// be filled in and then callback will be called. Alternatively the user may
	// abort this password request, which will usually abort the TLS connection.
	//
	// If the interaction is cancelled by the cancellation object, or by the
	// user then G_TLS_INTERACTION_FAILED will be returned with an error that
	// contains a G_IO_ERROR_CANCELLED error code. Certain implementations may
	// not support immediate cancellation.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable cancellation object.
	//    - password: Password object.
	//
	// The function returns the following values:
	//
	//    - tlsInteractionResult status of the ask password interaction.
	//
	AskPassword(ctx context.Context, password *TLSPassword) (TLSInteractionResult, error)
	// AskPasswordFinish: complete an ask password user interaction request.
	// This should be once the g_tls_interaction_ask_password_async() completion
	// callback is called.
	//
	// If G_TLS_INTERACTION_HANDLED is returned, then the Password passed to
	// g_tls_interaction_ask_password() will have its password filled in.
	//
	// If the interaction is cancelled by the cancellation object, or by the
	// user then G_TLS_INTERACTION_FAILED will be returned with an error that
	// contains a G_IO_ERROR_CANCELLED error code.
	//
	// The function takes the following parameters:
	//
	//    - result passed to the callback.
	//
	// The function returns the following values:
	//
	//    - tlsInteractionResult status of the ask password interaction.
	//
	AskPasswordFinish(result AsyncResultOverrider) (TLSInteractionResult, error)
	// RequestCertificate: run synchronous interaction to ask the user to choose
	// a certificate to use with the connection. In general,
	// g_tls_interaction_invoke_request_certificate() should be used instead of
	// this function.
	//
	// Derived subclasses usually implement a certificate selector, although
	// they may also choose to provide a certificate from elsewhere.
	// Alternatively the user may abort this certificate request, which will
	// usually abort the TLS connection.
	//
	// If G_TLS_INTERACTION_HANDLED is returned, then the Connection passed to
	// g_tls_interaction_request_certificate() will have had its
	// Connection:certificate filled in.
	//
	// If the interaction is cancelled by the cancellation object, or by the
	// user then G_TLS_INTERACTION_FAILED will be returned with an error that
	// contains a G_IO_ERROR_CANCELLED error code. Certain implementations may
	// not support immediate cancellation.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable cancellation object.
	//    - connection: Connection object.
	//    - flags providing more information about the request.
	//
	// The function returns the following values:
	//
	//    - tlsInteractionResult status of the request certificate interaction.
	//
	RequestCertificate(ctx context.Context, connection TLSConnectioner, flags TLSCertificateRequestFlags) (TLSInteractionResult, error)
	// RequestCertificateFinish: complete a request certificate user interaction
	// request. This should be once the
	// g_tls_interaction_request_certificate_async() completion callback is
	// called.
	//
	// If G_TLS_INTERACTION_HANDLED is returned, then the Connection passed to
	// g_tls_interaction_request_certificate_async() will have had its
	// Connection:certificate filled in.
	//
	// If the interaction is cancelled by the cancellation object, or by the
	// user then G_TLS_INTERACTION_FAILED will be returned with an error that
	// contains a G_IO_ERROR_CANCELLED error code.
	//
	// The function takes the following parameters:
	//
	//    - result passed to the callback.
	//
	// The function returns the following values:
	//
	//    - tlsInteractionResult status of the request certificate interaction.
	//
	RequestCertificateFinish(result AsyncResultOverrider) (TLSInteractionResult, error)
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
	*externglib.Object
}

var (
	_ externglib.Objector = (*TLSInteraction)(nil)
)

func classInitTLSInteractioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GTlsInteractionClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GTlsInteractionClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface {
		AskPassword(ctx context.Context, password *TLSPassword) (TLSInteractionResult, error)
	}); ok {
		pclass.ask_password = (*[0]byte)(C._gotk4_gio2_TlsInteractionClass_ask_password)
	}

	if _, ok := goval.(interface {
		AskPasswordFinish(result AsyncResultOverrider) (TLSInteractionResult, error)
	}); ok {
		pclass.ask_password_finish = (*[0]byte)(C._gotk4_gio2_TlsInteractionClass_ask_password_finish)
	}

	if _, ok := goval.(interface {
		RequestCertificate(ctx context.Context, connection TLSConnectioner, flags TLSCertificateRequestFlags) (TLSInteractionResult, error)
	}); ok {
		pclass.request_certificate = (*[0]byte)(C._gotk4_gio2_TlsInteractionClass_request_certificate)
	}

	if _, ok := goval.(interface {
		RequestCertificateFinish(result AsyncResultOverrider) (TLSInteractionResult, error)
	}); ok {
		pclass.request_certificate_finish = (*[0]byte)(C._gotk4_gio2_TlsInteractionClass_request_certificate_finish)
	}
}

//export _gotk4_gio2_TlsInteractionClass_ask_password
func _gotk4_gio2_TlsInteractionClass_ask_password(arg0 *C.GTlsInteraction, arg1 *C.GTlsPassword, arg2 *C.GCancellable, _cerr **C.GError) (cret C.GTlsInteractionResult) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		AskPassword(ctx context.Context, password *TLSPassword) (TLSInteractionResult, error)
	})

	var _cancellable context.Context // out
	var _password *TLSPassword       // out

	if arg2 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg2))
	}
	_password = wrapTLSPassword(externglib.Take(unsafe.Pointer(arg1)))

	tlsInteractionResult, _goerr := iface.AskPassword(_cancellable, _password)

	cret = C.GTlsInteractionResult(tlsInteractionResult)
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_TlsInteractionClass_ask_password_finish
func _gotk4_gio2_TlsInteractionClass_ask_password_finish(arg0 *C.GTlsInteraction, arg1 *C.GAsyncResult, _cerr **C.GError) (cret C.GTlsInteractionResult) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		AskPasswordFinish(result AsyncResultOverrider) (TLSInteractionResult, error)
	})

	var _result AsyncResultOverrider // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(AsyncResultOverrider)
			return ok
		})
		rv, ok := casted.(AsyncResultOverrider)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	tlsInteractionResult, _goerr := iface.AskPasswordFinish(_result)

	cret = C.GTlsInteractionResult(tlsInteractionResult)
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_TlsInteractionClass_request_certificate
func _gotk4_gio2_TlsInteractionClass_request_certificate(arg0 *C.GTlsInteraction, arg1 *C.GTlsConnection, arg2 C.GTlsCertificateRequestFlags, arg3 *C.GCancellable, _cerr **C.GError) (cret C.GTlsInteractionResult) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		RequestCertificate(ctx context.Context, connection TLSConnectioner, flags TLSCertificateRequestFlags) (TLSInteractionResult, error)
	})

	var _cancellable context.Context      // out
	var _connection TLSConnectioner       // out
	var _flags TLSCertificateRequestFlags // out

	if arg3 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg3))
	}
	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.TLSConnectioner is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(TLSConnectioner)
			return ok
		})
		rv, ok := casted.(TLSConnectioner)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.TLSConnectioner")
		}
		_connection = rv
	}
	_flags = TLSCertificateRequestFlags(arg2)

	tlsInteractionResult, _goerr := iface.RequestCertificate(_cancellable, _connection, _flags)

	cret = C.GTlsInteractionResult(tlsInteractionResult)
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_TlsInteractionClass_request_certificate_finish
func _gotk4_gio2_TlsInteractionClass_request_certificate_finish(arg0 *C.GTlsInteraction, arg1 *C.GAsyncResult, _cerr **C.GError) (cret C.GTlsInteractionResult) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		RequestCertificateFinish(result AsyncResultOverrider) (TLSInteractionResult, error)
	})

	var _result AsyncResultOverrider // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(AsyncResultOverrider)
			return ok
		})
		rv, ok := casted.(AsyncResultOverrider)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	tlsInteractionResult, _goerr := iface.RequestCertificateFinish(_result)

	cret = C.GTlsInteractionResult(tlsInteractionResult)
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

func wrapTLSInteraction(obj *externglib.Object) *TLSInteraction {
	return &TLSInteraction{
		Object: obj,
	}
}

func marshalTLSInteraction(p uintptr) (interface{}, error) {
	return wrapTLSInteraction(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AskPassword: run synchronous interaction to ask the user for a password. In
// general, g_tls_interaction_invoke_ask_password() should be used instead of
// this function.
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
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable cancellation object.
//    - password: Password object.
//
// The function returns the following values:
//
//    - tlsInteractionResult status of the ask password interaction.
//
func (interaction *TLSInteraction) AskPassword(ctx context.Context, password *TLSPassword) (TLSInteractionResult, error) {
	var _arg0 *C.GTlsInteraction      // out
	var _arg2 *C.GCancellable         // out
	var _arg1 *C.GTlsPassword         // out
	var _cret C.GTlsInteractionResult // in
	var _cerr *C.GError               // in

	_arg0 = (*C.GTlsInteraction)(unsafe.Pointer(externglib.InternObject(interaction).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GTlsPassword)(unsafe.Pointer(externglib.InternObject(password).Native()))

	_cret = C.g_tls_interaction_ask_password(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(interaction)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(password)

	var _tlsInteractionResult TLSInteractionResult // out
	var _goerr error                               // out

	_tlsInteractionResult = TLSInteractionResult(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsInteractionResult, _goerr
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
	var _arg0 *C.GTlsInteraction    // out
	var _arg2 *C.GCancellable       // out
	var _arg1 *C.GTlsPassword       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GTlsInteraction)(unsafe.Pointer(externglib.InternObject(interaction).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GTlsPassword)(unsafe.Pointer(externglib.InternObject(password).Native()))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_tls_interaction_ask_password_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(interaction)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(password)
	runtime.KeepAlive(callback)
}

// AskPasswordFinish: complete an ask password user interaction request. This
// should be once the g_tls_interaction_ask_password_async() completion callback
// is called.
//
// If G_TLS_INTERACTION_HANDLED is returned, then the Password passed to
// g_tls_interaction_ask_password() will have its password filled in.
//
// If the interaction is cancelled by the cancellation object, or by the user
// then G_TLS_INTERACTION_FAILED will be returned with an error that contains a
// G_IO_ERROR_CANCELLED error code.
//
// The function takes the following parameters:
//
//    - result passed to the callback.
//
// The function returns the following values:
//
//    - tlsInteractionResult status of the ask password interaction.
//
func (interaction *TLSInteraction) AskPasswordFinish(result AsyncResultOverrider) (TLSInteractionResult, error) {
	var _arg0 *C.GTlsInteraction      // out
	var _arg1 *C.GAsyncResult         // out
	var _cret C.GTlsInteractionResult // in
	var _cerr *C.GError               // in

	_arg0 = (*C.GTlsInteraction)(unsafe.Pointer(externglib.InternObject(interaction).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(result).Native()))

	_cret = C.g_tls_interaction_ask_password_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(interaction)
	runtime.KeepAlive(result)

	var _tlsInteractionResult TLSInteractionResult // out
	var _goerr error                               // out

	_tlsInteractionResult = TLSInteractionResult(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsInteractionResult, _goerr
}

// InvokeAskPassword: invoke the interaction to ask the user for a password. It
// invokes this interaction in the main loop, specifically the Context returned
// by g_main_context_get_thread_default() when the interaction is created. This
// is called by called by Connection or Database to ask the user for a password.
//
// Derived subclasses usually implement a password prompt, although they may
// also choose to provide a password from elsewhere. The password value will be
// filled in and then callback will be called. Alternatively the user may abort
// this password request, which will usually abort the TLS connection.
//
// The implementation can either be a synchronous (eg: modal dialog) or an
// asynchronous one (eg: modeless dialog). This function will take care of
// calling which ever one correctly.
//
// If the interaction is cancelled by the cancellation object, or by the user
// then G_TLS_INTERACTION_FAILED will be returned with an error that contains a
// G_IO_ERROR_CANCELLED error code. Certain implementations may not support
// immediate cancellation.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable cancellation object.
//    - password: Password object.
//
// The function returns the following values:
//
//    - tlsInteractionResult status of the ask password interaction.
//
func (interaction *TLSInteraction) InvokeAskPassword(ctx context.Context, password *TLSPassword) (TLSInteractionResult, error) {
	var _arg0 *C.GTlsInteraction      // out
	var _arg2 *C.GCancellable         // out
	var _arg1 *C.GTlsPassword         // out
	var _cret C.GTlsInteractionResult // in
	var _cerr *C.GError               // in

	_arg0 = (*C.GTlsInteraction)(unsafe.Pointer(externglib.InternObject(interaction).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GTlsPassword)(unsafe.Pointer(externglib.InternObject(password).Native()))

	_cret = C.g_tls_interaction_invoke_ask_password(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(interaction)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(password)

	var _tlsInteractionResult TLSInteractionResult // out
	var _goerr error                               // out

	_tlsInteractionResult = TLSInteractionResult(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsInteractionResult, _goerr
}

// InvokeRequestCertificate: invoke the interaction to ask the user to choose a
// certificate to use with the connection. It invokes this interaction in the
// main loop, specifically the Context returned by
// g_main_context_get_thread_default() when the interaction is created. This is
// called by called by Connection when the peer requests a certificate during
// the handshake.
//
// Derived subclasses usually implement a certificate selector, although they
// may also choose to provide a certificate from elsewhere. Alternatively the
// user may abort this certificate request, which may or may not abort the TLS
// connection.
//
// The implementation can either be a synchronous (eg: modal dialog) or an
// asynchronous one (eg: modeless dialog). This function will take care of
// calling which ever one correctly.
//
// If the interaction is cancelled by the cancellation object, or by the user
// then G_TLS_INTERACTION_FAILED will be returned with an error that contains a
// G_IO_ERROR_CANCELLED error code. Certain implementations may not support
// immediate cancellation.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable cancellation object.
//    - connection: Connection object.
//    - flags providing more information about the request.
//
// The function returns the following values:
//
//    - tlsInteractionResult status of the certificate request interaction.
//
func (interaction *TLSInteraction) InvokeRequestCertificate(ctx context.Context, connection TLSConnectioner, flags TLSCertificateRequestFlags) (TLSInteractionResult, error) {
	var _arg0 *C.GTlsInteraction            // out
	var _arg3 *C.GCancellable               // out
	var _arg1 *C.GTlsConnection             // out
	var _arg2 C.GTlsCertificateRequestFlags // out
	var _cret C.GTlsInteractionResult       // in
	var _cerr *C.GError                     // in

	_arg0 = (*C.GTlsInteraction)(unsafe.Pointer(externglib.InternObject(interaction).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GTlsConnection)(unsafe.Pointer(externglib.InternObject(connection).Native()))
	_arg2 = C.GTlsCertificateRequestFlags(flags)

	_cret = C.g_tls_interaction_invoke_request_certificate(_arg0, _arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(interaction)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(flags)

	var _tlsInteractionResult TLSInteractionResult // out
	var _goerr error                               // out

	_tlsInteractionResult = TLSInteractionResult(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsInteractionResult, _goerr
}

// RequestCertificate: run synchronous interaction to ask the user to choose a
// certificate to use with the connection. In general,
// g_tls_interaction_invoke_request_certificate() should be used instead of this
// function.
//
// Derived subclasses usually implement a certificate selector, although they
// may also choose to provide a certificate from elsewhere. Alternatively the
// user may abort this certificate request, which will usually abort the TLS
// connection.
//
// If G_TLS_INTERACTION_HANDLED is returned, then the Connection passed to
// g_tls_interaction_request_certificate() will have had its
// Connection:certificate filled in.
//
// If the interaction is cancelled by the cancellation object, or by the user
// then G_TLS_INTERACTION_FAILED will be returned with an error that contains a
// G_IO_ERROR_CANCELLED error code. Certain implementations may not support
// immediate cancellation.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable cancellation object.
//    - connection: Connection object.
//    - flags providing more information about the request.
//
// The function returns the following values:
//
//    - tlsInteractionResult status of the request certificate interaction.
//
func (interaction *TLSInteraction) RequestCertificate(ctx context.Context, connection TLSConnectioner, flags TLSCertificateRequestFlags) (TLSInteractionResult, error) {
	var _arg0 *C.GTlsInteraction            // out
	var _arg3 *C.GCancellable               // out
	var _arg1 *C.GTlsConnection             // out
	var _arg2 C.GTlsCertificateRequestFlags // out
	var _cret C.GTlsInteractionResult       // in
	var _cerr *C.GError                     // in

	_arg0 = (*C.GTlsInteraction)(unsafe.Pointer(externglib.InternObject(interaction).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GTlsConnection)(unsafe.Pointer(externglib.InternObject(connection).Native()))
	_arg2 = C.GTlsCertificateRequestFlags(flags)

	_cret = C.g_tls_interaction_request_certificate(_arg0, _arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(interaction)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(flags)

	var _tlsInteractionResult TLSInteractionResult // out
	var _goerr error                               // out

	_tlsInteractionResult = TLSInteractionResult(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsInteractionResult, _goerr
}

// RequestCertificateAsync: run asynchronous interaction to ask the user for a
// certificate to use with the connection. In general,
// g_tls_interaction_invoke_request_certificate() should be used instead of this
// function.
//
// Derived subclasses usually implement a certificate selector, although they
// may also choose to provide a certificate from elsewhere. callback will be
// called when the operation completes. Alternatively the user may abort this
// certificate request, which will usually abort the TLS connection.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable cancellation object.
//    - connection: Connection object.
//    - flags providing more information about the request.
//    - callback (optional) will be called when the interaction completes.
//
func (interaction *TLSInteraction) RequestCertificateAsync(ctx context.Context, connection TLSConnectioner, flags TLSCertificateRequestFlags, callback AsyncReadyCallback) {
	var _arg0 *C.GTlsInteraction            // out
	var _arg3 *C.GCancellable               // out
	var _arg1 *C.GTlsConnection             // out
	var _arg2 C.GTlsCertificateRequestFlags // out
	var _arg4 C.GAsyncReadyCallback         // out
	var _arg5 C.gpointer

	_arg0 = (*C.GTlsInteraction)(unsafe.Pointer(externglib.InternObject(interaction).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GTlsConnection)(unsafe.Pointer(externglib.InternObject(connection).Native()))
	_arg2 = C.GTlsCertificateRequestFlags(flags)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_tls_interaction_request_certificate_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(interaction)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(callback)
}

// RequestCertificateFinish: complete a request certificate user interaction
// request. This should be once the
// g_tls_interaction_request_certificate_async() completion callback is called.
//
// If G_TLS_INTERACTION_HANDLED is returned, then the Connection passed to
// g_tls_interaction_request_certificate_async() will have had its
// Connection:certificate filled in.
//
// If the interaction is cancelled by the cancellation object, or by the user
// then G_TLS_INTERACTION_FAILED will be returned with an error that contains a
// G_IO_ERROR_CANCELLED error code.
//
// The function takes the following parameters:
//
//    - result passed to the callback.
//
// The function returns the following values:
//
//    - tlsInteractionResult status of the request certificate interaction.
//
func (interaction *TLSInteraction) RequestCertificateFinish(result AsyncResultOverrider) (TLSInteractionResult, error) {
	var _arg0 *C.GTlsInteraction      // out
	var _arg1 *C.GAsyncResult         // out
	var _cret C.GTlsInteractionResult // in
	var _cerr *C.GError               // in

	_arg0 = (*C.GTlsInteraction)(unsafe.Pointer(externglib.InternObject(interaction).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(result).Native()))

	_cret = C.g_tls_interaction_request_certificate_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(interaction)
	runtime.KeepAlive(result)

	var _tlsInteractionResult TLSInteractionResult // out
	var _goerr error                               // out

	_tlsInteractionResult = TLSInteractionResult(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsInteractionResult, _goerr
}
