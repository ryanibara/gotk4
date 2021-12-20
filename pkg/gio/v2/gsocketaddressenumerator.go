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
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_socket_address_enumerator_get_type()), F: marshalSocketAddressEnumeratorrer},
	})
}

// SocketAddressEnumeratorOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SocketAddressEnumeratorOverrider interface {
	// Next retrieves the next Address from enumerator. Note that this may block
	// for some amount of time. (Eg, a Address may need to do a DNS lookup
	// before it can return an address.) Use
	// g_socket_address_enumerator_next_async() if you need to avoid blocking.
	//
	// If enumerator is expected to yield addresses, but for some reason is
	// unable to (eg, because of a DNS error), then the first call to
	// g_socket_address_enumerator_next() will return an appropriate error in
	// *error. However, if the first call to g_socket_address_enumerator_next()
	// succeeds, then any further internal errors (other than cancellable being
	// triggered) will be ignored.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable object, NULL to ignore.
	//
	// The function returns the following values:
	//
	//    - socketAddress (owned by the caller), or NULL on error (in which case
	//      *error will be set) or if there are no more addresses.
	//
	Next(ctx context.Context) (SocketAddresser, error)
	// NextAsync: asynchronously retrieves the next Address from enumerator and
	// then calls callback, which must call
	// g_socket_address_enumerator_next_finish() to get the result.
	//
	// It is an error to call this multiple times before the previous callback
	// has finished.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable object, NULL to ignore.
	//    - callback (optional) to call when the request is satisfied.
	//
	NextAsync(ctx context.Context, callback AsyncReadyCallback)
	// NextFinish retrieves the result of a completed call to
	// g_socket_address_enumerator_next_async(). See
	// g_socket_address_enumerator_next() for more information about error
	// handling.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	// The function returns the following values:
	//
	//    - socketAddress (owned by the caller), or NULL on error (in which case
	//      *error will be set) or if there are no more addresses.
	//
	NextFinish(result AsyncResulter) (SocketAddresser, error)
}

// SocketAddressEnumerator is an enumerator type for Address instances. It is
// returned by enumeration functions such as g_socket_connectable_enumerate(),
// which returns a AddressEnumerator to list each Address which could be used to
// connect to that Connectable.
//
// Enumeration is typically a blocking operation, so the asynchronous methods
// g_socket_address_enumerator_next_async() and
// g_socket_address_enumerator_next_finish() should be used where possible.
//
// Each AddressEnumerator can only be enumerated once. Once
// g_socket_address_enumerator_next() has returned NULL, further enumeration
// with that AddressEnumerator is not possible, and it can be unreffed.
type SocketAddressEnumerator struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*SocketAddressEnumerator)(nil)
)

// SocketAddressEnumeratorrer describes types inherited from class SocketAddressEnumerator.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type SocketAddressEnumeratorrer interface {
	externglib.Objector
	baseSocketAddressEnumerator() *SocketAddressEnumerator
}

var _ SocketAddressEnumeratorrer = (*SocketAddressEnumerator)(nil)

func wrapSocketAddressEnumerator(obj *externglib.Object) *SocketAddressEnumerator {
	return &SocketAddressEnumerator{
		Object: obj,
	}
}

func marshalSocketAddressEnumeratorrer(p uintptr) (interface{}, error) {
	return wrapSocketAddressEnumerator(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (enumerator *SocketAddressEnumerator) baseSocketAddressEnumerator() *SocketAddressEnumerator {
	return enumerator
}

// BaseSocketAddressEnumerator returns the underlying base object.
func BaseSocketAddressEnumerator(obj SocketAddressEnumeratorrer) *SocketAddressEnumerator {
	return obj.baseSocketAddressEnumerator()
}

// Next retrieves the next Address from enumerator. Note that this may block for
// some amount of time. (Eg, a Address may need to do a DNS lookup before it can
// return an address.) Use g_socket_address_enumerator_next_async() if you need
// to avoid blocking.
//
// If enumerator is expected to yield addresses, but for some reason is unable
// to (eg, because of a DNS error), then the first call to
// g_socket_address_enumerator_next() will return an appropriate error in
// *error. However, if the first call to g_socket_address_enumerator_next()
// succeeds, then any further internal errors (other than cancellable being
// triggered) will be ignored.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//
// The function returns the following values:
//
//    - socketAddress (owned by the caller), or NULL on error (in which case
//      *error will be set) or if there are no more addresses.
//
func (enumerator *SocketAddressEnumerator) Next(ctx context.Context) (SocketAddresser, error) {
	var _arg0 *C.GSocketAddressEnumerator // out
	var _arg1 *C.GCancellable             // out
	var _cret *C.GSocketAddress           // in
	var _cerr *C.GError                   // in

	_arg0 = (*C.GSocketAddressEnumerator)(unsafe.Pointer(enumerator.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C.g_socket_address_enumerator_next(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(enumerator)
	runtime.KeepAlive(ctx)

	var _socketAddress SocketAddresser // out
	var _goerr error                   // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.SocketAddresser is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.Cast()
		rv, ok := casted.(SocketAddresser)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.SocketAddresser")
		}
		_socketAddress = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketAddress, _goerr
}

// NextAsync: asynchronously retrieves the next Address from enumerator and then
// calls callback, which must call g_socket_address_enumerator_next_finish() to
// get the result.
//
// It is an error to call this multiple times before the previous callback has
// finished.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - callback (optional) to call when the request is satisfied.
//
func (enumerator *SocketAddressEnumerator) NextAsync(ctx context.Context, callback AsyncReadyCallback) {
	var _arg0 *C.GSocketAddressEnumerator // out
	var _arg1 *C.GCancellable             // out
	var _arg2 C.GAsyncReadyCallback       // out
	var _arg3 C.gpointer

	_arg0 = (*C.GSocketAddressEnumerator)(unsafe.Pointer(enumerator.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_socket_address_enumerator_next_async(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(enumerator)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// NextFinish retrieves the result of a completed call to
// g_socket_address_enumerator_next_async(). See
// g_socket_address_enumerator_next() for more information about error handling.
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - socketAddress (owned by the caller), or NULL on error (in which case
//      *error will be set) or if there are no more addresses.
//
func (enumerator *SocketAddressEnumerator) NextFinish(result AsyncResulter) (SocketAddresser, error) {
	var _arg0 *C.GSocketAddressEnumerator // out
	var _arg1 *C.GAsyncResult             // out
	var _cret *C.GSocketAddress           // in
	var _cerr *C.GError                   // in

	_arg0 = (*C.GSocketAddressEnumerator)(unsafe.Pointer(enumerator.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_socket_address_enumerator_next_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(enumerator)
	runtime.KeepAlive(result)

	var _socketAddress SocketAddresser // out
	var _goerr error                   // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.SocketAddresser is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.Cast()
		rv, ok := casted.(SocketAddresser)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.SocketAddresser")
		}
		_socketAddress = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketAddress, _goerr
}
