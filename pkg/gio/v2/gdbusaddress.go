// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
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

// DBusAddressEscapeValue: escape @string so it can appear in a D-Bus address as
// the value part of a key-value pair.
//
// For instance, if @string is `/run/bus-for-:0`, this function would return
// `/run/bus-for-3A0`, which could be used in a D-Bus address like
// `unix:nonce-tcp:host=127.0.0.1,port=42,noncefile=/run/bus-for-3A0`.
func DBusAddressEscapeValue(string string) string {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var ret1 string

	cret = C.g_dbus_address_escape_value(string)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// DBusAddressGetForBusSync: synchronously looks up the D-Bus address for the
// well-known message bus instance specified by @bus_type. This may involve
// using various platform specific mechanisms.
//
// The returned address will be in the D-Bus address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
func DBusAddressGetForBusSync(busType BusType, cancellable Cancellable) (utf8 string, err error) {
	var arg1 C.GBusType
	var arg2 *C.GCancellable

	arg1 = (C.GBusType)(busType)
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var errout *C.GError
	var goerr error
	var cret *C.gchar
	var ret2 string

	cret = C.g_dbus_address_get_for_bus_sync(busType, cancellable, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goerr, ret2
}

// DBusAddressGetStream: asynchronously connects to an endpoint specified by
// @address and sets up the connection so it is in a state to run the
// client-side of the D-Bus authentication conversation. @address must be in the
// D-Bus address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// When the operation is finished, @callback will be invoked. You can then call
// g_dbus_address_get_stream_finish() to get the result of the operation.
//
// This is an asynchronous failable function. See
// g_dbus_address_get_stream_sync() for the synchronous version.
func DBusAddressGetStream(address string, cancellable Cancellable, callback AsyncReadyCallback) {

	C.g_dbus_address_get_stream(address, cancellable, callback, userData)
}

// DBusAddressGetStreamFinish finishes an operation started with
// g_dbus_address_get_stream().
func DBusAddressGetStreamFinish(res AsyncResult) (outGuid string, ioStream IOStream, err error) {
	var arg1 *C.GAsyncResult

	arg1 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	var arg2 *C.gchar
	var ret2 string
	var errout *C.GError
	var goerr error
	var cret *C.GIOStream
	var ret3 IOStream

	cret = C.g_dbus_address_get_stream_finish(res, &arg2, &errout)

	*ret2 = C.GoString(arg2)
	defer C.free(unsafe.Pointer(arg2))
	goerr = gerror.Take(unsafe.Pointer(errout))
	ret3 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(IOStream)

	return ret2, goerr, ret3
}

// DBusAddressGetStreamSync: synchronously connects to an endpoint specified by
// @address and sets up the connection so it is in a state to run the
// client-side of the D-Bus authentication conversation. @address must be in the
// D-Bus address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// This is a synchronous failable function. See g_dbus_address_get_stream() for
// the asynchronous version.
func DBusAddressGetStreamSync(address string, cancellable Cancellable) (outGuid string, ioStream IOStream, err error) {
	var arg1 *C.gchar
	var arg3 *C.GCancellable

	arg1 = (*C.gchar)(C.CString(address))
	defer C.free(unsafe.Pointer(arg1))
	arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var arg2 *C.gchar
	var ret2 string
	var errout *C.GError
	var goerr error
	var cret *C.GIOStream
	var ret3 IOStream

	cret = C.g_dbus_address_get_stream_sync(address, &arg2, cancellable, &errout)

	*ret2 = C.GoString(arg2)
	defer C.free(unsafe.Pointer(arg2))
	goerr = gerror.Take(unsafe.Pointer(errout))
	ret3 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(IOStream)

	return ret2, goerr, ret3
}

// DBusIsAddress checks if @string is a D-Bus address
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// This doesn't check if @string is actually supported by BusServer or
// BusConnection - use g_dbus_is_supported_address() to do more checks.
func DBusIsAddress(string string) bool {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ret1 bool

	cret = C.g_dbus_is_address(string)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// DBusIsSupportedAddress: like g_dbus_is_address() but also checks if the
// library supports the transports in @string and that key/value pairs for each
// transport are valid. See the specification of the D-Bus address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
func DBusIsSupportedAddress(string string) error {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	var errout *C.GError
	var goerr error

	C.g_dbus_is_supported_address(string, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))

	return goerr
}
