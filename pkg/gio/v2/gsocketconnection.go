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

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_socket_connection_get_type()), F: marshalSocketConnectioner},
	})
}

// SocketConnection is a OStream for a connected socket. They can be created
// either by Client when connecting to a host, or by Listener when accepting a
// new client.
//
// The type of the Connection object returned from these calls depends on the
// type of the underlying socket that is in use. For instance, for a TCP/IP
// connection it will be a Connection.
//
// Choosing what type of object to construct is done with the socket connection
// factory, and it is possible for 3rd parties to register custom socket
// connection types for specific combination of socket family/type/protocol
// using g_socket_connection_factory_register_type().
//
// To close a Connection, use g_io_stream_close(). Closing both substreams of
// the OStream separately will not close the underlying #GSocket.
type SocketConnection struct {
	IOStream
}

var (
	_ IOStreamer = (*SocketConnection)(nil)
)

func wrapSocketConnection(obj *externglib.Object) *SocketConnection {
	return &SocketConnection{
		IOStream: IOStream{
			Object: obj,
		},
	}
}

func marshalSocketConnectioner(p uintptr) (interface{}, error) {
	return wrapSocketConnection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectSocketConnection: connect connection to the specified remote address.
//
// The function takes the following parameters:
//
//    - ctx: GCancellable or NULL.
//    - address specifying the remote address.
//
func (connection *SocketConnection) ConnectSocketConnection(ctx context.Context, address SocketAddresser) error {
	var _arg0 *C.GSocketConnection // out
	var _arg2 *C.GCancellable      // out
	var _arg1 *C.GSocketAddress    // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketConnection)(unsafe.Pointer(connection.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GSocketAddress)(unsafe.Pointer(address.Native()))

	C.g_socket_connection_connect(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(address)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// ConnectAsync: asynchronously connect connection to the specified remote
// address.
//
// This clears the #GSocket:blocking flag on connection's underlying socket if
// it is currently set.
//
// Use g_socket_connection_connect_finish() to retrieve the result.
//
// The function takes the following parameters:
//
//    - ctx: GCancellable or NULL.
//    - address specifying the remote address.
//    - callback: ReadyCallback.
//
func (connection *SocketConnection) ConnectAsync(ctx context.Context, address SocketAddresser, callback AsyncReadyCallback) {
	var _arg0 *C.GSocketConnection  // out
	var _arg2 *C.GCancellable       // out
	var _arg1 *C.GSocketAddress     // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GSocketConnection)(unsafe.Pointer(connection.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GSocketAddress)(unsafe.Pointer(address.Native()))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_socket_connection_connect_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(address)
	runtime.KeepAlive(callback)
}

// ConnectFinish gets the result of a g_socket_connection_connect_async() call.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (connection *SocketConnection) ConnectFinish(result AsyncResulter) error {
	var _arg0 *C.GSocketConnection // out
	var _arg1 *C.GAsyncResult      // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketConnection)(unsafe.Pointer(connection.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_socket_connection_connect_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// LocalAddress: try to get the local address of a socket connection.
func (connection *SocketConnection) LocalAddress() (SocketAddresser, error) {
	var _arg0 *C.GSocketConnection // out
	var _cret *C.GSocketAddress    // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketConnection)(unsafe.Pointer(connection.Native()))

	_cret = C.g_socket_connection_get_local_address(_arg0, &_cerr)
	runtime.KeepAlive(connection)

	var _socketAddress SocketAddresser // out
	var _goerr error                   // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.SocketAddresser is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(SocketAddresser)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.SocketAddresser")
		}
		_socketAddress = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketAddress, _goerr
}

// RemoteAddress: try to get the remote address of a socket connection.
//
// Since GLib 2.40, when used with g_socket_client_connect() or
// g_socket_client_connect_async(), during emission of
// G_SOCKET_CLIENT_CONNECTING, this function will return the remote address that
// will be used for the connection. This allows applications to print e.g.
// "Connecting to example.com (10.42.77.3)...".
func (connection *SocketConnection) RemoteAddress() (SocketAddresser, error) {
	var _arg0 *C.GSocketConnection // out
	var _cret *C.GSocketAddress    // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketConnection)(unsafe.Pointer(connection.Native()))

	_cret = C.g_socket_connection_get_remote_address(_arg0, &_cerr)
	runtime.KeepAlive(connection)

	var _socketAddress SocketAddresser // out
	var _goerr error                   // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.SocketAddresser is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(SocketAddresser)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.SocketAddresser")
		}
		_socketAddress = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketAddress, _goerr
}

// Socket gets the underlying #GSocket object of the connection. This can be
// useful if you want to do something unusual on it not supported by the
// Connection APIs.
func (connection *SocketConnection) Socket() *Socket {
	var _arg0 *C.GSocketConnection // out
	var _cret *C.GSocket           // in

	_arg0 = (*C.GSocketConnection)(unsafe.Pointer(connection.Native()))

	_cret = C.g_socket_connection_get_socket(_arg0)
	runtime.KeepAlive(connection)

	var _socket *Socket // out

	_socket = wrapSocket(externglib.Take(unsafe.Pointer(_cret)))

	return _socket
}

// IsConnected checks if connection is connected. This is equivalent to calling
// g_socket_is_connected() on connection's underlying #GSocket.
func (connection *SocketConnection) IsConnected() bool {
	var _arg0 *C.GSocketConnection // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GSocketConnection)(unsafe.Pointer(connection.Native()))

	_cret = C.g_socket_connection_is_connected(_arg0)
	runtime.KeepAlive(connection)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SocketConnectionFactoryLookupType looks up the #GType to be used when
// creating socket connections on sockets with the specified family, type and
// protocol_id.
//
// If no type is registered, the Connection base type is returned.
//
// The function takes the following parameters:
//
//    - family: Family.
//    - typ: Type.
//    - protocolId: protocol id.
//
func SocketConnectionFactoryLookupType(family SocketFamily, typ SocketType, protocolId int) externglib.Type {
	var _arg1 C.GSocketFamily // out
	var _arg2 C.GSocketType   // out
	var _arg3 C.gint          // out
	var _cret C.GType         // in

	_arg1 = C.GSocketFamily(family)
	_arg2 = C.GSocketType(typ)
	_arg3 = C.gint(protocolId)

	_cret = C.g_socket_connection_factory_lookup_type(_arg1, _arg2, _arg3)
	runtime.KeepAlive(family)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(protocolId)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// SocketConnectionFactoryRegisterType looks up the #GType to be used when
// creating socket connections on sockets with the specified family, type and
// protocol.
//
// If no type is registered, the Connection base type is returned.
//
// The function takes the following parameters:
//
//    - gType inheriting from G_TYPE_SOCKET_CONNECTION.
//    - family: Family.
//    - typ: Type.
//    - protocol id.
//
func SocketConnectionFactoryRegisterType(gType externglib.Type, family SocketFamily, typ SocketType, protocol int) {
	var _arg1 C.GType         // out
	var _arg2 C.GSocketFamily // out
	var _arg3 C.GSocketType   // out
	var _arg4 C.gint          // out

	_arg1 = C.GType(gType)
	_arg2 = C.GSocketFamily(family)
	_arg3 = C.GSocketType(typ)
	_arg4 = C.gint(protocol)

	C.g_socket_connection_factory_register_type(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(gType)
	runtime.KeepAlive(family)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(protocol)
}
