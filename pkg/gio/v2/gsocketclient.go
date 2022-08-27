// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// void _gotk4_gio2_SocketClient_virtual_event(void* fnptr, GSocketClient* arg0, GSocketClientEvent arg1, GSocketConnectable* arg2, GIOStream* arg3) {
//   ((void (*)(GSocketClient*, GSocketClientEvent, GSocketConnectable*, GIOStream*))(fnptr))(arg0, arg1, arg2, arg3);
// };
import "C"

// AddApplicationProxy: enable proxy protocols to be handled by the application.
// When the indicated proxy protocol is returned by the Resolver, Client will
// consider this protocol as supported but will not try to find a #GProxy
// instance to handle handshaking. The application must check for this case by
// calling g_socket_connection_get_remote_address() on the returned Connection,
// and seeing if it's a Address of the appropriate type, to determine whether or
// not it needs to handle the proxy handshaking itself.
//
// This should be used for proxy protocols that are dialects of another protocol
// such as HTTP proxy. It also allows cohabitation of proxy protocols that are
// reused between protocols. A good example is HTTP. It can be used to proxy
// HTTP, FTP and Gopher and can also be use as generic socket proxy through the
// HTTP CONNECT method.
//
// When the proxy is detected as being an application proxy, TLS handshake will
// be skipped. This is required to let the application do the proxy specific
// handshake.
//
// The function takes the following parameters:
//
//    - protocol: proxy protocol.
//
func (client *SocketClient) AddApplicationProxy(protocol string) {
	var _arg0 *C.GSocketClient // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(protocol)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_socket_client_add_application_proxy(_arg0, _arg1)
	runtime.KeepAlive(client)
	runtime.KeepAlive(protocol)
}

// ConnectToService attempts to create a TCP connection to a service.
//
// This call looks up the SRV record for service at domain for the "tcp"
// protocol. It then attempts to connect, in turn, to each of the hosts
// providing the service until either a connection succeeds or there are no
// hosts remaining.
//
// Upon a successful connection, a new Connection is constructed and returned.
// The caller owns this new object and must drop their reference to it when
// finished with it.
//
// In the event of any failure (DNS error, service not found, no hosts
// connectable) NULL is returned and error (if non-NULL) is set accordingly.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - domain name.
//    - service: name of the service to connect to.
//
// The function returns the following values:
//
//    - socketConnection if successful, or NULL on error.
//
func (client *SocketClient) ConnectToService(ctx context.Context, domain, service string) (*SocketConnection, error) {
	var _arg0 *C.GSocketClient     // out
	var _arg3 *C.GCancellable      // out
	var _arg1 *C.gchar             // out
	var _arg2 *C.gchar             // out
	var _cret *C.GSocketConnection // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(service)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_socket_client_connect_to_service(_arg0, _arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(client)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(service)

	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_socketConnection = wrapSocketConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
}

// The function takes the following parameters:
//
//    - event
//    - connectable
//    - connection
//
func (client *SocketClient) event(event SocketClientEvent, connectable SocketConnectabler, connection IOStreamer) {
	gclass := (*C.GSocketClientClass)(coreglib.PeekParentClass(client))
	fnarg := gclass.event

	var _arg0 *C.GSocketClient      // out
	var _arg1 C.GSocketClientEvent  // out
	var _arg2 *C.GSocketConnectable // out
	var _arg3 *C.GIOStream          // out

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	_arg1 = C.GSocketClientEvent(event)
	_arg2 = (*C.GSocketConnectable)(unsafe.Pointer(coreglib.InternObject(connectable).Native()))
	_arg3 = (*C.GIOStream)(unsafe.Pointer(coreglib.InternObject(connection).Native()))

	C._gotk4_gio2_SocketClient_virtual_event(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(client)
	runtime.KeepAlive(event)
	runtime.KeepAlive(connectable)
	runtime.KeepAlive(connection)
}

// SocketClientClass: instance of this type is always passed by reference.
type SocketClientClass struct {
	*socketClientClass
}

// socketClientClass is the struct that's finalized.
type socketClientClass struct {
	native *C.GSocketClientClass
}
