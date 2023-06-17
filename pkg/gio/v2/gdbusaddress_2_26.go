// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// DBusAddressGetForBusSync: synchronously looks up the D-Bus address for the
// well-known message bus instance specified by bus_type. This may involve using
// various platform specific mechanisms.
//
// The returned address will be in the D-Bus address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// The function takes the following parameters:
//
//   - ctx (optional) or NULL.
//   - busType: Type.
//
// The function returns the following values:
//
//   - utf8: valid D-Bus address string for bus_type or NULL if error is set.
//
func DBusAddressGetForBusSync(ctx context.Context, busType BusType) (string, error) {
	var _arg2 *C.GCancellable // out
	var _arg1 C.GBusType      // out
	var _cret *C.gchar        // in
	var _cerr *C.GError       // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GBusType(busType)

	_cret = C.g_dbus_address_get_for_bus_sync(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(busType)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _utf8, _goerr
}

// DBusAddressGetStream: asynchronously connects to an endpoint
// specified by address and sets up the connection so it is in
// a state to run the client-side of the D-Bus authentication
// conversation. address must be in the D-Bus address format
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
//   - ctx (optional) or NULL.
//   - address: valid D-Bus address.
//   - callback (optional) to call when the request is satisfied.
//
func DBusAddressGetStream(ctx context.Context, address string, callback AsyncReadyCallback) {
	var _arg2 *C.GCancellable       // out
	var _arg1 *C.gchar              // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(address)))
	defer C.free(unsafe.Pointer(_arg1))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_dbus_address_get_stream(_arg1, _arg2, _arg3, _arg4)
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
//   - res obtained from the GAsyncReadyCallback passed to
//     g_dbus_address_get_stream().
//
// The function returns the following values:
//
//   - outGuid (optional): NULL or return location to store the GUID extracted
//     from address, if any.
//   - ioStream or NULL if error is set.
//
func DBusAddressGetStreamFinish(res AsyncResulter) (string, IOStreamer, error) {
	var _arg1 *C.GAsyncResult // out
	var _arg2 *C.gchar        // in
	var _cret *C.GIOStream    // in
	var _cerr *C.GError       // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(res).Native()))

	_cret = C.g_dbus_address_get_stream_finish(_arg1, &_arg2, &_cerr)
	runtime.KeepAlive(res)

	var _outGuid string      // out
	var _ioStream IOStreamer // out
	var _goerr error         // out

	if _arg2 != nil {
		_outGuid = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
		defer C.free(unsafe.Pointer(_arg2))
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
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _outGuid, _ioStream, _goerr
}

// DBusAddressGetStreamSync: synchronously connects to an endpoint
// specified by address and sets up the connection so it is in
// a state to run the client-side of the D-Bus authentication
// conversation. address must be in the D-Bus address format
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
//   - ctx (optional) or NULL.
//   - address: valid D-Bus address.
//
// The function returns the following values:
//
//   - outGuid (optional): NULL or return location to store the GUID extracted
//     from address, if any.
//   - ioStream or NULL if error is set.
//
func DBusAddressGetStreamSync(ctx context.Context, address string) (string, IOStreamer, error) {
	var _arg3 *C.GCancellable // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.gchar        // in
	var _cret *C.GIOStream    // in
	var _cerr *C.GError       // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(address)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_address_get_stream_sync(_arg1, &_arg2, _arg3, &_cerr)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(address)

	var _outGuid string      // out
	var _ioStream IOStreamer // out
	var _goerr error         // out

	if _arg2 != nil {
		_outGuid = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
		defer C.free(unsafe.Pointer(_arg2))
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
	if _cerr != nil {
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
//   - str: string.
//
// The function returns the following values:
//
//   - ok: TRUE if string is a valid D-Bus address, FALSE otherwise.
//
func DBusIsAddress(str string) bool {
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_is_address(_arg1)
	runtime.KeepAlive(str)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DBusIsSupportedAddress: like g_dbus_is_address() but also checks if the
// library supports the transports in string and that key/value pairs for
// each transport are valid. See the specification of the D-Bus address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// The function takes the following parameters:
//
//   - str: string.
//
func DBusIsSupportedAddress(str string) error {
	var _arg1 *C.gchar  // out
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_dbus_is_supported_address(_arg1, &_cerr)
	runtime.KeepAlive(str)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
