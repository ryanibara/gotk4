// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gio2_AsyncReadyCallback(void*, void*, gpointer);
import "C"

// DBusAddressEscapeValue: escape string so it can appear in a D-Bus address as
// the value part of a key-value pair.
//
// For instance, if string is /run/bus-for-:0, this function would return
// /run/bus-for-3A0, which could be used in a D-Bus address like
// unix:nonce-tcp:host=127.0.0.1,port=42,noncefile=/run/bus-for-3A0.
//
// The function takes the following parameters:
//
//    - str: unescaped string to be included in a D-Bus address as the value in a
//      key-value pair.
//
// The function returns the following values:
//
//    - utf8: copy of string with all non-optionally-escaped bytes escaped.
//
func DBusAddressEscapeValue(str string) string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_args[0]))

	_gret := girepository.MustFind("Gio", "dbus_address_escape_value").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(str)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// DBusAddressGetStream: asynchronously connects to an endpoint specified by
// address and sets up the connection so it is in a state to run the client-side
// of the D-Bus authentication conversation. address must be in the D-Bus
// address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// When the operation is finished, callback will be invoked. You can then call
// g_dbus_address_get_stream_finish() to get the result of the operation.
//
// This is an asynchronous failable function. See
// g_dbus_address_get_stream_sync() for the synchronous version.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - address: valid D-Bus address.
//    - callback (optional) to call when the request is satisfied.
//
func DBusAddressGetStream(ctx context.Context, address string, callback AsyncReadyCallback) {
	var _args [4]girepository.Argument

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[1] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(address)))
	defer C.free(unsafe.Pointer(_args[0]))
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[2])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[3] = C.gpointer(gbox.AssignOnce(callback))
	}

	girepository.MustFind("Gio", "dbus_address_get_stream").Invoke(_args[:], nil)

	runtime.KeepAlive(ctx)
	runtime.KeepAlive(address)
	runtime.KeepAlive(callback)
}

// DBusAddressGetStreamFinish finishes an operation started with
// g_dbus_address_get_stream().
//
// A server is not required to set a GUID, so out_guid may be set to NULL even
// on success.
//
// The function takes the following parameters:
//
//    - res obtained from the GAsyncReadyCallback passed to
//      g_dbus_address_get_stream().
//
// The function returns the following values:
//
//    - outGuid (optional): NULL or return location to store the GUID extracted
//      from address, if any.
//    - ioStream or NULL if error is set.
//
func DBusAddressGetStreamFinish(res AsyncResulter) (string, IOStreamer, error) {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(res).Native()))

	_gret := girepository.MustFind("Gio", "dbus_address_get_stream_finish").Invoke(_args[:], _outs[:])
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(res)

	var _outGuid string      // out
	var _ioStream IOStreamer // out
	var _goerr error         // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_outGuid = C.GoString((*C.gchar)(unsafe.Pointer(_outs[0])))
		defer C.free(unsafe.Pointer(_outs[0]))
	}
	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.IOStreamer is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(IOStreamer)
			return ok
		})
		rv, ok := casted.(IOStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.IOStreamer")
		}
		_ioStream = rv
	}
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _outGuid, _ioStream, _goerr
}

// DBusAddressGetStreamSync: synchronously connects to an endpoint specified by
// address and sets up the connection so it is in a state to run the client-side
// of the D-Bus authentication conversation. address must be in the D-Bus
// address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// A server is not required to set a GUID, so out_guid may be set to NULL even
// on success.
//
// This is a synchronous failable function. See g_dbus_address_get_stream() for
// the asynchronous version.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - address: valid D-Bus address.
//
// The function returns the following values:
//
//    - outGuid (optional): NULL or return location to store the GUID extracted
//      from address, if any.
//    - ioStream or NULL if error is set.
//
func DBusAddressGetStreamSync(ctx context.Context, address string) (string, IOStreamer, error) {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[1] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(address)))
	defer C.free(unsafe.Pointer(_args[0]))

	_gret := girepository.MustFind("Gio", "dbus_address_get_stream_sync").Invoke(_args[:], _outs[:])
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(ctx)
	runtime.KeepAlive(address)

	var _outGuid string      // out
	var _ioStream IOStreamer // out
	var _goerr error         // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_outGuid = C.GoString((*C.gchar)(unsafe.Pointer(_outs[0])))
		defer C.free(unsafe.Pointer(_outs[0]))
	}
	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.IOStreamer is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(IOStreamer)
			return ok
		})
		rv, ok := casted.(IOStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.IOStreamer")
		}
		_ioStream = rv
	}
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _outGuid, _ioStream, _goerr
}

// DBusIsAddress checks if string is a D-Bus address
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// This doesn't check if string is actually supported by BusServer or
// BusConnection - use g_dbus_is_supported_address() to do more checks.
//
// The function takes the following parameters:
//
//    - str: string.
//
// The function returns the following values:
//
//    - ok: TRUE if string is a valid D-Bus address, FALSE otherwise.
//
func DBusIsAddress(str string) bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_args[0]))

	_gret := girepository.MustFind("Gio", "dbus_is_address").Invoke(_args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(str)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// DBusIsSupportedAddress: like g_dbus_is_address() but also checks if the
// library supports the transports in string and that key/value pairs for each
// transport are valid. See the specification of the D-Bus address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// The function takes the following parameters:
//
//    - str: string.
//
func DBusIsSupportedAddress(str string) error {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_args[0]))

	girepository.MustFind("Gio", "dbus_is_supported_address").Invoke(_args[:], nil)

	runtime.KeepAlive(str)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
