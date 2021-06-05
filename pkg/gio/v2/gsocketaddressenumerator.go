// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
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
		{T: externglib.Type(C.g_socket_address_enumerator_get_type()), F: marshalSocketAddressEnumerator},
	})
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
// g_socket_address_enumerator_next() has returned nil, further enumeration with
// that AddressEnumerator is not possible, and it can be unreffed.
type SocketAddressEnumerator interface {
	gextras.Objector

	// Next retrieves the next Address from @enumerator. Note that this may
	// block for some amount of time. (Eg, a Address may need to do a DNS lookup
	// before it can return an address.) Use
	// g_socket_address_enumerator_next_async() if you need to avoid blocking.
	//
	// If @enumerator is expected to yield addresses, but for some reason is
	// unable to (eg, because of a DNS error), then the first call to
	// g_socket_address_enumerator_next() will return an appropriate error in
	// *@error. However, if the first call to g_socket_address_enumerator_next()
	// succeeds, then any further internal errors (other than @cancellable being
	// triggered) will be ignored.
	Next(cancellable Cancellable) (socketAddress SocketAddress, err error)
	// NextAsync: asynchronously retrieves the next Address from @enumerator and
	// then calls @callback, which must call
	// g_socket_address_enumerator_next_finish() to get the result.
	//
	// It is an error to call this multiple times before the previous callback
	// has finished.
	NextAsync(cancellable Cancellable, callback AsyncReadyCallback)
	// NextFinish retrieves the result of a completed call to
	// g_socket_address_enumerator_next_async(). See
	// g_socket_address_enumerator_next() for more information about error
	// handling.
	NextFinish(result AsyncResult) (socketAddress SocketAddress, err error)
}

// socketAddressEnumerator implements the SocketAddressEnumerator interface.
type socketAddressEnumerator struct {
	gextras.Objector
}

var _ SocketAddressEnumerator = (*socketAddressEnumerator)(nil)

// WrapSocketAddressEnumerator wraps a GObject to the right type. It is
// primarily used internally.
func WrapSocketAddressEnumerator(obj *externglib.Object) SocketAddressEnumerator {
	return SocketAddressEnumerator{
		Objector: obj,
	}
}

func marshalSocketAddressEnumerator(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSocketAddressEnumerator(obj), nil
}

// Next retrieves the next Address from @enumerator. Note that this may
// block for some amount of time. (Eg, a Address may need to do a DNS lookup
// before it can return an address.) Use
// g_socket_address_enumerator_next_async() if you need to avoid blocking.
//
// If @enumerator is expected to yield addresses, but for some reason is
// unable to (eg, because of a DNS error), then the first call to
// g_socket_address_enumerator_next() will return an appropriate error in
// *@error. However, if the first call to g_socket_address_enumerator_next()
// succeeds, then any further internal errors (other than @cancellable being
// triggered) will be ignored.
func (e socketAddressEnumerator) Next(cancellable Cancellable) (socketAddress SocketAddress, err error) {
	var arg0 *C.GSocketAddressEnumerator
	var arg1 *C.GCancellable
	var errout *C.GError

	arg0 = (*C.GSocketAddressEnumerator)(unsafe.Pointer(e.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret *C.GSocketAddress
	var ret1 SocketAddress
	var goerr error

	cret = C.g_socket_address_enumerator_next(arg0, cancellable, &errout)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(SocketAddress)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return ret1, goerr
}

// NextAsync: asynchronously retrieves the next Address from @enumerator and
// then calls @callback, which must call
// g_socket_address_enumerator_next_finish() to get the result.
//
// It is an error to call this multiple times before the previous callback
// has finished.
func (e socketAddressEnumerator) NextAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GSocketAddressEnumerator

	arg0 = (*C.GSocketAddressEnumerator)(unsafe.Pointer(e.Native()))

	C.g_socket_address_enumerator_next_async(arg0, cancellable, callback, userData)
}

// NextFinish retrieves the result of a completed call to
// g_socket_address_enumerator_next_async(). See
// g_socket_address_enumerator_next() for more information about error
// handling.
func (e socketAddressEnumerator) NextFinish(result AsyncResult) (socketAddress SocketAddress, err error) {
	var arg0 *C.GSocketAddressEnumerator
	var arg1 *C.GAsyncResult
	var errout *C.GError

	arg0 = (*C.GSocketAddressEnumerator)(unsafe.Pointer(e.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cret *C.GSocketAddress
	var ret1 SocketAddress
	var goerr error

	cret = C.g_socket_address_enumerator_next_finish(arg0, result, &errout)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(SocketAddress)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return ret1, goerr
}
