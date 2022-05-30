// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gsocketconnection.go.
var GTypeSocketConnection = coreglib.Type(C.g_socket_connection_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeSocketConnection, F: marshalSocketConnection},
	})
}

// ConnectionFactoryCreateConnection creates a Connection subclass of the right
// type for socket.
//
// The function returns the following values:
//
//    - socketConnection: Connection.
//
func (socket *Socket) ConnectionFactoryCreateConnection() *SocketConnection {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(socket).Native()))
	*(**Socket)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "Socket").InvokeMethod("connection_factory_create_connection", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(socket)

	var _socketConnection *SocketConnection // out

	_socketConnection = wrapSocketConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _socketConnection
}

// SocketConnectionOverrider contains methods that are overridable.
type SocketConnectionOverrider interface {
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
	_ [0]func() // equal guard
	IOStream
}

var (
	_ IOStreamer = (*SocketConnection)(nil)
)

func classInitSocketConnectioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSocketConnection(obj *coreglib.Object) *SocketConnection {
	return &SocketConnection{
		IOStream: IOStream{
			Object: obj,
		},
	}
}

func marshalSocketConnection(p uintptr) (interface{}, error) {
	return wrapSocketConnection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectSocketConnection: connect connection to the specified remote address.
//
// The function takes the following parameters:
//
//    - ctx (optional): GCancellable or NULL.
//    - address specifying the remote address.
//
func (connection *SocketConnection) ConnectSocketConnection(ctx context.Context, address SocketAddresser) error {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg2 *C.void // out
	var _arg1 *C.void // out
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	*(**SocketConnection)(unsafe.Pointer(&args[1])) = _arg1
	*(*context.Context)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gio", "SocketConnection").InvokeMethod("connect", args[:], nil)

	runtime.KeepAlive(connection)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(address)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// ConnectFinish gets the result of a g_socket_connection_connect_async() call.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (connection *SocketConnection) ConnectFinish(result AsyncResulter) error {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))
	*(**SocketConnection)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gio", "SocketConnection").InvokeMethod("connect_finish", args[:], nil)

	runtime.KeepAlive(connection)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// LocalAddress: try to get the local address of a socket connection.
//
// The function returns the following values:
//
//    - socketAddress or NULL on error. Free the returned object with
//      g_object_unref().
//
func (connection *SocketConnection) LocalAddress() (SocketAddresser, error) {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
	*(**SocketConnection)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "SocketConnection").InvokeMethod("get_local_address", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(connection)

	var _socketAddress SocketAddresser // out
	var _goerr error                   // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.SocketAddresser is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(SocketAddresser)
			return ok
		})
		rv, ok := casted.(SocketAddresser)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.SocketAddresser")
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
//
// The function returns the following values:
//
//    - socketAddress or NULL on error. Free the returned object with
//      g_object_unref().
//
func (connection *SocketConnection) RemoteAddress() (SocketAddresser, error) {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
	*(**SocketConnection)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "SocketConnection").InvokeMethod("get_remote_address", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(connection)

	var _socketAddress SocketAddresser // out
	var _goerr error                   // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.SocketAddresser is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(SocketAddresser)
			return ok
		})
		rv, ok := casted.(SocketAddresser)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.SocketAddresser")
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
//
// The function returns the following values:
//
//    - socket or NULL on error.
//
func (connection *SocketConnection) Socket() *Socket {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
	*(**SocketConnection)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "SocketConnection").InvokeMethod("get_socket", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(connection)

	var _socket *Socket // out

	_socket = wrapSocket(coreglib.Take(unsafe.Pointer(_cret)))

	return _socket
}

// IsConnected checks if connection is connected. This is equivalent to calling
// g_socket_is_connected() on connection's underlying #GSocket.
//
// The function returns the following values:
//
//    - ok: whether connection is connected.
//
func (connection *SocketConnection) IsConnected() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
	*(**SocketConnection)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "SocketConnection").InvokeMethod("is_connected", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(connection)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
