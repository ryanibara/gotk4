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

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
// #include <glib-object.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_socket_client_get_type()), F: marshalSocketClienter},
	})
}

// SocketClientOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SocketClientOverrider interface {
	Event(event SocketClientEvent, connectable SocketConnectabler, connection IOStreamer)
}

// SocketClient is a lightweight high-level utility class for connecting to a
// network host using a connection oriented socket type.
//
// You create a Client object, set any options you want, and then call a sync or
// async connect operation, which returns a Connection subclass on success.
//
// The type of the Connection object returned depends on the type of the
// underlying socket that is in use. For instance, for a TCP/IP connection it
// will be a Connection.
//
// As Client is a lightweight object, you don't need to cache it. You can just
// create a new one any time you need one.
type SocketClient struct {
	*externglib.Object
}

func wrapSocketClient(obj *externglib.Object) *SocketClient {
	return &SocketClient{
		Object: obj,
	}
}

func marshalSocketClienter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSocketClient(obj), nil
}

// NewSocketClient creates a new Client with the default options.
func NewSocketClient() *SocketClient {
	var _cret *C.GSocketClient // in

	_cret = C.g_socket_client_new()

	var _socketClient *SocketClient // out

	_socketClient = wrapSocketClient(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _socketClient
}

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
func (client *SocketClient) AddApplicationProxy(protocol string) {
	var _arg0 *C.GSocketClient // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(protocol)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_socket_client_add_application_proxy(_arg0, _arg1)
	runtime.KeepAlive(client)
	runtime.KeepAlive(protocol)
}

// ConnectSocketClient tries to resolve the connectable and make a network
// connection to it.
//
// Upon a successful connection, a new Connection is constructed and returned.
// The caller owns this new object and must drop their reference to it when
// finished with it.
//
// The type of the Connection object returned depends on the type of the
// underlying socket that is used. For instance, for a TCP/IP connection it will
// be a Connection.
//
// The socket created will be the same family as the address that the
// connectable resolves to, unless family is set with
// g_socket_client_set_family() or indirectly via
// g_socket_client_set_local_address(). The socket type defaults to
// G_SOCKET_TYPE_STREAM but can be set with g_socket_client_set_socket_type().
//
// If a local address is specified with g_socket_client_set_local_address() the
// socket will be bound to this address before connecting.
func (client *SocketClient) ConnectSocketClient(ctx context.Context, connectable SocketConnectabler) (*SocketConnection, error) {
	var _arg0 *C.GSocketClient      // out
	var _arg2 *C.GCancellable       // out
	var _arg1 *C.GSocketConnectable // out
	var _cret *C.GSocketConnection  // in
	var _cerr *C.GError             // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GSocketConnectable)(unsafe.Pointer(connectable.Native()))

	_cret = C.g_socket_client_connect(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(client)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connectable)

	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_socketConnection = wrapSocketConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
}

// ConnectAsync: this is the asynchronous version of g_socket_client_connect().
//
// You may wish to prefer the asynchronous version even in synchronous command
// line programs because, since 2.60, it implements RFC 8305
// (https://tools.ietf.org/html/rfc8305) "Happy Eyeballs" recommendations to
// work around long connection timeouts in networks where IPv6 is broken by
// performing an IPv4 connection simultaneously without waiting for IPv6 to time
// out, which is not supported by the synchronous call. (This is not an API
// guarantee, and may change in the future.)
//
// When the operation is finished callback will be called. You can then call
// g_socket_client_connect_finish() to get the result of the operation.
func (client *SocketClient) ConnectAsync(ctx context.Context, connectable SocketConnectabler, callback AsyncReadyCallback) {
	var _arg0 *C.GSocketClient      // out
	var _arg2 *C.GCancellable       // out
	var _arg1 *C.GSocketConnectable // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GSocketConnectable)(unsafe.Pointer(connectable.Native()))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_socket_client_connect_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(client)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connectable)
	runtime.KeepAlive(callback)
}

// ConnectFinish finishes an async connect operation. See
// g_socket_client_connect_async().
func (client *SocketClient) ConnectFinish(result AsyncResulter) (*SocketConnection, error) {
	var _arg0 *C.GSocketClient     // out
	var _arg1 *C.GAsyncResult      // out
	var _cret *C.GSocketConnection // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_socket_client_connect_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(client)
	runtime.KeepAlive(result)

	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_socketConnection = wrapSocketConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
}

// ConnectToHost: this is a helper function for g_socket_client_connect().
//
// Attempts to create a TCP connection to the named host.
//
// host_and_port may be in any of a number of recognized formats; an IPv6
// address, an IPv4 address, or a domain name (in which case a DNS lookup is
// performed). Quoting with [] is supported for all address types. A port
// override may be specified in the usual way with a colon. Ports may be given
// as decimal numbers or symbolic names (in which case an /etc/services lookup
// is performed).
//
// If no port override is given in host_and_port then default_port will be used
// as the port number to connect to.
//
// In general, host_and_port is expected to be provided by the user (allowing
// them to give the hostname, and a port override if necessary) and default_port
// is expected to be provided by the application.
//
// In the case that an IP address is given, a single connection attempt is made.
// In the case that a name is given, multiple connection attempts may be made,
// in turn and according to the number of address records in DNS, until a
// connection succeeds.
//
// Upon a successful connection, a new Connection is constructed and returned.
// The caller owns this new object and must drop their reference to it when
// finished with it.
//
// In the event of any failure (DNS error, service not found, no hosts
// connectable) NULL is returned and error (if non-NULL) is set accordingly.
func (client *SocketClient) ConnectToHost(ctx context.Context, hostAndPort string, defaultPort uint16) (*SocketConnection, error) {
	var _arg0 *C.GSocketClient     // out
	var _arg3 *C.GCancellable      // out
	var _arg1 *C.gchar             // out
	var _arg2 C.guint16            // out
	var _cret *C.GSocketConnection // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(hostAndPort)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint16(defaultPort)

	_cret = C.g_socket_client_connect_to_host(_arg0, _arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(client)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(hostAndPort)
	runtime.KeepAlive(defaultPort)

	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_socketConnection = wrapSocketConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
}

// ConnectToHostAsync: this is the asynchronous version of
// g_socket_client_connect_to_host().
//
// When the operation is finished callback will be called. You can then call
// g_socket_client_connect_to_host_finish() to get the result of the operation.
func (client *SocketClient) ConnectToHostAsync(ctx context.Context, hostAndPort string, defaultPort uint16, callback AsyncReadyCallback) {
	var _arg0 *C.GSocketClient      // out
	var _arg3 *C.GCancellable       // out
	var _arg1 *C.gchar              // out
	var _arg2 C.guint16             // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(hostAndPort)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint16(defaultPort)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_socket_client_connect_to_host_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(client)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(hostAndPort)
	runtime.KeepAlive(defaultPort)
	runtime.KeepAlive(callback)
}

// ConnectToHostFinish finishes an async connect operation. See
// g_socket_client_connect_to_host_async().
func (client *SocketClient) ConnectToHostFinish(result AsyncResulter) (*SocketConnection, error) {
	var _arg0 *C.GSocketClient     // out
	var _arg1 *C.GAsyncResult      // out
	var _cret *C.GSocketConnection // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_socket_client_connect_to_host_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(client)
	runtime.KeepAlive(result)

	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_socketConnection = wrapSocketConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
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
func (client *SocketClient) ConnectToService(ctx context.Context, domain string, service string) (*SocketConnection, error) {
	var _arg0 *C.GSocketClient     // out
	var _arg3 *C.GCancellable      // out
	var _arg1 *C.gchar             // out
	var _arg2 *C.gchar             // out
	var _cret *C.GSocketConnection // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
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

	_socketConnection = wrapSocketConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
}

// ConnectToServiceAsync: this is the asynchronous version of
// g_socket_client_connect_to_service().
func (client *SocketClient) ConnectToServiceAsync(ctx context.Context, domain string, service string, callback AsyncReadyCallback) {
	var _arg0 *C.GSocketClient      // out
	var _arg3 *C.GCancellable       // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.gchar              // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(service)))
	defer C.free(unsafe.Pointer(_arg2))
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_socket_client_connect_to_service_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(client)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(service)
	runtime.KeepAlive(callback)
}

// ConnectToServiceFinish finishes an async connect operation. See
// g_socket_client_connect_to_service_async().
func (client *SocketClient) ConnectToServiceFinish(result AsyncResulter) (*SocketConnection, error) {
	var _arg0 *C.GSocketClient     // out
	var _arg1 *C.GAsyncResult      // out
	var _cret *C.GSocketConnection // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_socket_client_connect_to_service_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(client)
	runtime.KeepAlive(result)

	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_socketConnection = wrapSocketConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
}

// ConnectToURI: this is a helper function for g_socket_client_connect().
//
// Attempts to create a TCP connection with a network URI.
//
// uri may be any valid URI containing an "authority" (hostname/port) component.
// If a port is not specified in the URI, default_port will be used. TLS will be
// negotiated if Client:tls is TRUE. (Client does not know to automatically
// assume TLS for certain URI schemes.)
//
// Using this rather than g_socket_client_connect() or
// g_socket_client_connect_to_host() allows Client to determine when to use
// application-specific proxy protocols.
//
// Upon a successful connection, a new Connection is constructed and returned.
// The caller owns this new object and must drop their reference to it when
// finished with it.
//
// In the event of any failure (DNS error, service not found, no hosts
// connectable) NULL is returned and error (if non-NULL) is set accordingly.
func (client *SocketClient) ConnectToURI(ctx context.Context, uri string, defaultPort uint16) (*SocketConnection, error) {
	var _arg0 *C.GSocketClient     // out
	var _arg3 *C.GCancellable      // out
	var _arg1 *C.gchar             // out
	var _arg2 C.guint16            // out
	var _cret *C.GSocketConnection // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint16(defaultPort)

	_cret = C.g_socket_client_connect_to_uri(_arg0, _arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(client)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(defaultPort)

	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_socketConnection = wrapSocketConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
}

// ConnectToURIAsync: this is the asynchronous version of
// g_socket_client_connect_to_uri().
//
// When the operation is finished callback will be called. You can then call
// g_socket_client_connect_to_uri_finish() to get the result of the operation.
func (client *SocketClient) ConnectToURIAsync(ctx context.Context, uri string, defaultPort uint16, callback AsyncReadyCallback) {
	var _arg0 *C.GSocketClient      // out
	var _arg3 *C.GCancellable       // out
	var _arg1 *C.gchar              // out
	var _arg2 C.guint16             // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint16(defaultPort)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_socket_client_connect_to_uri_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(client)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(defaultPort)
	runtime.KeepAlive(callback)
}

// ConnectToURIFinish finishes an async connect operation. See
// g_socket_client_connect_to_uri_async().
func (client *SocketClient) ConnectToURIFinish(result AsyncResulter) (*SocketConnection, error) {
	var _arg0 *C.GSocketClient     // out
	var _arg1 *C.GAsyncResult      // out
	var _cret *C.GSocketConnection // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_socket_client_connect_to_uri_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(client)
	runtime.KeepAlive(result)

	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_socketConnection = wrapSocketConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
}

// EnableProxy gets the proxy enable state; see
// g_socket_client_set_enable_proxy().
func (client *SocketClient) EnableProxy() bool {
	var _arg0 *C.GSocketClient // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))

	_cret = C.g_socket_client_get_enable_proxy(_arg0)
	runtime.KeepAlive(client)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Family gets the socket family of the socket client.
//
// See g_socket_client_set_family() for details.
func (client *SocketClient) Family() SocketFamily {
	var _arg0 *C.GSocketClient // out
	var _cret C.GSocketFamily  // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))

	_cret = C.g_socket_client_get_family(_arg0)
	runtime.KeepAlive(client)

	var _socketFamily SocketFamily // out

	_socketFamily = SocketFamily(_cret)

	return _socketFamily
}

// LocalAddress gets the local address of the socket client.
//
// See g_socket_client_set_local_address() for details.
func (client *SocketClient) LocalAddress() SocketAddresser {
	var _arg0 *C.GSocketClient  // out
	var _cret *C.GSocketAddress // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))

	_cret = C.g_socket_client_get_local_address(_arg0)
	runtime.KeepAlive(client)

	var _socketAddress SocketAddresser // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(SocketAddresser)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.SocketAddresser")
			}
			_socketAddress = rv
		}
	}

	return _socketAddress
}

// Protocol gets the protocol name type of the socket client.
//
// See g_socket_client_set_protocol() for details.
func (client *SocketClient) Protocol() SocketProtocol {
	var _arg0 *C.GSocketClient  // out
	var _cret C.GSocketProtocol // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))

	_cret = C.g_socket_client_get_protocol(_arg0)
	runtime.KeepAlive(client)

	var _socketProtocol SocketProtocol // out

	_socketProtocol = SocketProtocol(_cret)

	return _socketProtocol
}

// ProxyResolver gets the Resolver being used by client. Normally, this will be
// the resolver returned by g_proxy_resolver_get_default(), but you can override
// it with g_socket_client_set_proxy_resolver().
func (client *SocketClient) ProxyResolver() ProxyResolverer {
	var _arg0 *C.GSocketClient  // out
	var _cret *C.GProxyResolver // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))

	_cret = C.g_socket_client_get_proxy_resolver(_arg0)
	runtime.KeepAlive(client)

	var _proxyResolver ProxyResolverer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.ProxyResolverer is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(ProxyResolverer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.ProxyResolverer")
		}
		_proxyResolver = rv
	}

	return _proxyResolver
}

// SocketType gets the socket type of the socket client.
//
// See g_socket_client_set_socket_type() for details.
func (client *SocketClient) SocketType() SocketType {
	var _arg0 *C.GSocketClient // out
	var _cret C.GSocketType    // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))

	_cret = C.g_socket_client_get_socket_type(_arg0)
	runtime.KeepAlive(client)

	var _socketType SocketType // out

	_socketType = SocketType(_cret)

	return _socketType
}

// Timeout gets the I/O timeout time for sockets created by client.
//
// See g_socket_client_set_timeout() for details.
func (client *SocketClient) Timeout() uint {
	var _arg0 *C.GSocketClient // out
	var _cret C.guint          // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))

	_cret = C.g_socket_client_get_timeout(_arg0)
	runtime.KeepAlive(client)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// TLS gets whether client creates TLS connections. See
// g_socket_client_set_tls() for details.
func (client *SocketClient) TLS() bool {
	var _arg0 *C.GSocketClient // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))

	_cret = C.g_socket_client_get_tls(_arg0)
	runtime.KeepAlive(client)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TLSValidationFlags gets the TLS validation flags used creating TLS
// connections via client.
func (client *SocketClient) TLSValidationFlags() TLSCertificateFlags {
	var _arg0 *C.GSocketClient       // out
	var _cret C.GTlsCertificateFlags // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))

	_cret = C.g_socket_client_get_tls_validation_flags(_arg0)
	runtime.KeepAlive(client)

	var _tlsCertificateFlags TLSCertificateFlags // out

	_tlsCertificateFlags = TLSCertificateFlags(_cret)

	return _tlsCertificateFlags
}

// SetEnableProxy sets whether or not client attempts to make connections via a
// proxy server. When enabled (the default), Client will use a Resolver to
// determine if a proxy protocol such as SOCKS is needed, and automatically do
// the necessary proxy negotiation.
//
// See also g_socket_client_set_proxy_resolver().
func (client *SocketClient) SetEnableProxy(enable bool) {
	var _arg0 *C.GSocketClient // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	if enable {
		_arg1 = C.TRUE
	}

	C.g_socket_client_set_enable_proxy(_arg0, _arg1)
	runtime.KeepAlive(client)
	runtime.KeepAlive(enable)
}

// SetFamily sets the socket family of the socket client. If this is set to
// something other than G_SOCKET_FAMILY_INVALID then the sockets created by this
// object will be of the specified family.
//
// This might be useful for instance if you want to force the local connection
// to be an ipv4 socket, even though the address might be an ipv6 mapped to ipv4
// address.
func (client *SocketClient) SetFamily(family SocketFamily) {
	var _arg0 *C.GSocketClient // out
	var _arg1 C.GSocketFamily  // out

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	_arg1 = C.GSocketFamily(family)

	C.g_socket_client_set_family(_arg0, _arg1)
	runtime.KeepAlive(client)
	runtime.KeepAlive(family)
}

// SetLocalAddress sets the local address of the socket client. The sockets
// created by this object will bound to the specified address (if not NULL)
// before connecting.
//
// This is useful if you want to ensure that the local side of the connection is
// on a specific port, or on a specific interface.
func (client *SocketClient) SetLocalAddress(address SocketAddresser) {
	var _arg0 *C.GSocketClient  // out
	var _arg1 *C.GSocketAddress // out

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	if address != nil {
		_arg1 = (*C.GSocketAddress)(unsafe.Pointer(address.Native()))
	}

	C.g_socket_client_set_local_address(_arg0, _arg1)
	runtime.KeepAlive(client)
	runtime.KeepAlive(address)
}

// SetProtocol sets the protocol of the socket client. The sockets created by
// this object will use of the specified protocol.
//
// If protocol is G_SOCKET_PROTOCOL_DEFAULT that means to use the default
// protocol for the socket family and type.
func (client *SocketClient) SetProtocol(protocol SocketProtocol) {
	var _arg0 *C.GSocketClient  // out
	var _arg1 C.GSocketProtocol // out

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	_arg1 = C.GSocketProtocol(protocol)

	C.g_socket_client_set_protocol(_arg0, _arg1)
	runtime.KeepAlive(client)
	runtime.KeepAlive(protocol)
}

// SetProxyResolver overrides the Resolver used by client. You can call this if
// you want to use specific proxies, rather than using the system default proxy
// settings.
//
// Note that whether or not the proxy resolver is actually used depends on the
// setting of Client:enable-proxy, which is not changed by this function (but
// which is TRUE by default).
func (client *SocketClient) SetProxyResolver(proxyResolver ProxyResolverer) {
	var _arg0 *C.GSocketClient  // out
	var _arg1 *C.GProxyResolver // out

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	if proxyResolver != nil {
		_arg1 = (*C.GProxyResolver)(unsafe.Pointer(proxyResolver.Native()))
	}

	C.g_socket_client_set_proxy_resolver(_arg0, _arg1)
	runtime.KeepAlive(client)
	runtime.KeepAlive(proxyResolver)
}

// SetSocketType sets the socket type of the socket client. The sockets created
// by this object will be of the specified type.
//
// It doesn't make sense to specify a type of G_SOCKET_TYPE_DATAGRAM, as
// GSocketClient is used for connection oriented services.
func (client *SocketClient) SetSocketType(typ SocketType) {
	var _arg0 *C.GSocketClient // out
	var _arg1 C.GSocketType    // out

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	_arg1 = C.GSocketType(typ)

	C.g_socket_client_set_socket_type(_arg0, _arg1)
	runtime.KeepAlive(client)
	runtime.KeepAlive(typ)
}

// SetTimeout sets the I/O timeout for sockets created by client. timeout is a
// time in seconds, or 0 for no timeout (the default).
//
// The timeout value affects the initial connection attempt as well, so setting
// this may cause calls to g_socket_client_connect(), etc, to fail with
// G_IO_ERROR_TIMED_OUT.
func (client *SocketClient) SetTimeout(timeout uint) {
	var _arg0 *C.GSocketClient // out
	var _arg1 C.guint          // out

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	_arg1 = C.guint(timeout)

	C.g_socket_client_set_timeout(_arg0, _arg1)
	runtime.KeepAlive(client)
	runtime.KeepAlive(timeout)
}

// SetTLS sets whether client creates TLS (aka SSL) connections. If tls is TRUE,
// client will wrap its connections in a ClientConnection and perform a TLS
// handshake when connecting.
//
// Note that since Client must return a Connection, but ClientConnection is not
// a Connection, this actually wraps the resulting ClientConnection in a
// WrapperConnection when returning it. You can use
// g_tcp_wrapper_connection_get_base_io_stream() on the return value to extract
// the ClientConnection.
//
// If you need to modify the behavior of the TLS handshake (eg, by setting a
// client-side certificate to use, or connecting to the
// Connection::accept-certificate signal), you can connect to client's
// Client::event signal and wait for it to be emitted with
// G_SOCKET_CLIENT_TLS_HANDSHAKING, which will give you a chance to see the
// ClientConnection before the handshake starts.
func (client *SocketClient) SetTLS(tls bool) {
	var _arg0 *C.GSocketClient // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	if tls {
		_arg1 = C.TRUE
	}

	C.g_socket_client_set_tls(_arg0, _arg1)
	runtime.KeepAlive(client)
	runtime.KeepAlive(tls)
}

// SetTLSValidationFlags sets the TLS validation flags used when creating TLS
// connections via client. The default value is G_TLS_CERTIFICATE_VALIDATE_ALL.
func (client *SocketClient) SetTLSValidationFlags(flags TLSCertificateFlags) {
	var _arg0 *C.GSocketClient       // out
	var _arg1 C.GTlsCertificateFlags // out

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(client.Native()))
	_arg1 = C.GTlsCertificateFlags(flags)

	C.g_socket_client_set_tls_validation_flags(_arg0, _arg1)
	runtime.KeepAlive(client)
	runtime.KeepAlive(flags)
}

// ConnectEvent: emitted when client's activity on connectable changes state.
// Among other things, this can be used to provide progress information about a
// network connection in the UI. The meanings of the different event values are
// as follows:
//
// - G_SOCKET_CLIENT_RESOLVING: client is about to look up connectable in DNS.
// connection will be NULL.
//
// - G_SOCKET_CLIENT_RESOLVED: client has successfully resolved connectable in
// DNS. connection will be NULL.
//
// - G_SOCKET_CLIENT_CONNECTING: client is about to make a connection to a
// remote host; either a proxy server or the destination server itself.
// connection is the Connection, which is not yet connected. Since GLib 2.40,
// you can access the remote address via
// g_socket_connection_get_remote_address().
//
// - G_SOCKET_CLIENT_CONNECTED: client has successfully connected to a remote
// host. connection is the connected Connection.
//
// - G_SOCKET_CLIENT_PROXY_NEGOTIATING: client is about to negotiate with a
// proxy to get it to connect to connectable. connection is the Connection to
// the proxy server.
//
// - G_SOCKET_CLIENT_PROXY_NEGOTIATED: client has negotiated a connection to
// connectable through a proxy server. connection is the stream returned from
// g_proxy_connect(), which may or may not be a Connection.
//
// - G_SOCKET_CLIENT_TLS_HANDSHAKING: client is about to begin a TLS handshake.
// connection is a ClientConnection.
//
// - G_SOCKET_CLIENT_TLS_HANDSHAKED: client has successfully completed the TLS
// handshake. connection is a ClientConnection.
//
// - G_SOCKET_CLIENT_COMPLETE: client has either successfully connected to
// connectable (in which case connection is the Connection that it will be
// returning to the caller) or has failed (in which case connection is NULL and
// the client is about to return an error).
//
// Each event except G_SOCKET_CLIENT_COMPLETE may be emitted multiple times (or
// not at all) for a given connectable (in particular, if client ends up
// attempting to connect to more than one address). However, if client emits the
// Client::event signal at all for a given connectable, then it will always emit
// it with G_SOCKET_CLIENT_COMPLETE when it is done.
//
// Note that there may be additional ClientEvent values in the future;
// unrecognized event values should be ignored.
func (s *SocketClient) ConnectEvent(f func(event SocketClientEvent, connectable SocketConnectabler, connection IOStreamer)) glib.SignalHandle {
	return s.Connect("event", f)
}
