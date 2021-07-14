// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
// #include <glib-object.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_socket_listener_get_type()), F: marshalSocketListenerer},
	})
}

// SocketListenerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SocketListenerOverrider interface {
	Changed()
	Event(event SocketListenerEvent, socket *Socket)
}

// SocketListenerer describes SocketListener's methods.
type SocketListenerer interface {
	// Accept blocks waiting for a client to connect to any of the sockets added
	// to the listener.
	Accept(cancellable *Cancellable) (*externglib.Object, *SocketConnection, error)
	// AcceptAsync: this is the asynchronous version of
	// g_socket_listener_accept().
	AcceptAsync(cancellable *Cancellable, callback AsyncReadyCallback)
	// AcceptFinish finishes an async accept operation.
	AcceptFinish(result AsyncResulter) (*externglib.Object, *SocketConnection, error)
	// AcceptSocket blocks waiting for a client to connect to any of the sockets
	// added to the listener.
	AcceptSocket(cancellable *Cancellable) (*externglib.Object, *Socket, error)
	// AcceptSocketAsync: this is the asynchronous version of
	// g_socket_listener_accept_socket().
	AcceptSocketAsync(cancellable *Cancellable, callback AsyncReadyCallback)
	// AcceptSocketFinish finishes an async accept operation.
	AcceptSocketFinish(result AsyncResulter) (*externglib.Object, *Socket, error)
	// AddAddress creates a socket of type type and protocol protocol, binds it
	// to address and adds it to the set of sockets we're accepting sockets
	// from.
	AddAddress(address SocketAddresser, typ SocketType, protocol SocketProtocol, sourceObject *externglib.Object) (*SocketAddress, error)
	// AddAnyInetPort listens for TCP connections on any available port number
	// for both IPv6 and IPv4 (if each is available).
	AddAnyInetPort(sourceObject *externglib.Object) (uint16, error)
	// AddInetPort: helper function for g_socket_listener_add_address() that
	// creates a TCP/IP socket listening on IPv4 and IPv6 (if supported) on the
	// specified port on all interfaces.
	AddInetPort(port uint16, sourceObject *externglib.Object) error
	// AddSocket adds socket to the set of sockets that we try to accept new
	// clients from.
	AddSocket(socket *Socket, sourceObject *externglib.Object) error
	// Close closes all the sockets in the listener.
	Close()
	// SetBacklog sets the listen backlog on the sockets in the listener.
	SetBacklog(listenBacklog int)
}

// SocketListener is an object that keeps track of a set of server sockets and
// helps you accept sockets from any of the socket, either sync or async.
//
// Add addresses and ports to listen on using g_socket_listener_add_address()
// and g_socket_listener_add_inet_port(). These will be listened on until
// g_socket_listener_close() is called. Dropping your final reference to the
// Listener will not cause g_socket_listener_close() to be called implicitly, as
// some references to the Listener may be held internally.
//
// If you want to implement a network server, also look at Service and
// SocketService which are subclasses of Listener that make this even easier.
type SocketListener struct {
	*externglib.Object
}

var (
	_ SocketListenerer = (*SocketListener)(nil)
	_ gextras.Nativer  = (*SocketListener)(nil)
)

func wrapSocketListener(obj *externglib.Object) *SocketListener {
	return &SocketListener{
		Object: obj,
	}
}

func marshalSocketListenerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSocketListener(obj), nil
}

// NewSocketListener creates a new Listener with no sockets to listen for. New
// listeners can be added with e.g. g_socket_listener_add_address() or
// g_socket_listener_add_inet_port().
func NewSocketListener() *SocketListener {
	var _cret *C.GSocketListener // in

	_cret = C.g_socket_listener_new()

	var _socketListener *SocketListener // out

	_socketListener = wrapSocketListener(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _socketListener
}

// Accept blocks waiting for a client to connect to any of the sockets added to
// the listener. Returns a Connection for the socket that was accepted.
//
// If source_object is not NULL it will be filled out with the source object
// specified when the corresponding socket or address was added to the listener.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned.
func (listener *SocketListener) Accept(cancellable *Cancellable) (*externglib.Object, *SocketConnection, error) {
	var _arg0 *C.GSocketListener   // out
	var _arg1 *C.GObject           // in
	var _arg2 *C.GCancellable      // out
	var _cret *C.GSocketConnection // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(listener.Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_socket_listener_accept(_arg0, &_arg1, _arg2, &_cerr)

	var _sourceObject *externglib.Object    // out
	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_sourceObject = externglib.Take(unsafe.Pointer(_arg1))
	_socketConnection = wrapSocketConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _sourceObject, _socketConnection, _goerr
}

// AcceptAsync: this is the asynchronous version of g_socket_listener_accept().
//
// When the operation is finished callback will be called. You can then call
// g_socket_listener_accept_socket() to get the result of the operation.
func (listener *SocketListener) AcceptAsync(cancellable *Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GSocketListener    // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(listener.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg3 = C.gpointer(gbox.AssignOnce(callback))

	C.g_socket_listener_accept_async(_arg0, _arg1, _arg2, _arg3)
}

// AcceptFinish finishes an async accept operation. See
// g_socket_listener_accept_async()
func (listener *SocketListener) AcceptFinish(result AsyncResulter) (*externglib.Object, *SocketConnection, error) {
	var _arg0 *C.GSocketListener   // out
	var _arg1 *C.GAsyncResult      // out
	var _arg2 *C.GObject           // in
	var _cret *C.GSocketConnection // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(listener.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.g_socket_listener_accept_finish(_arg0, _arg1, &_arg2, &_cerr)

	var _sourceObject *externglib.Object    // out
	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_sourceObject = externglib.Take(unsafe.Pointer(_arg2))
	_socketConnection = wrapSocketConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _sourceObject, _socketConnection, _goerr
}

// AcceptSocket blocks waiting for a client to connect to any of the sockets
// added to the listener. Returns the #GSocket that was accepted.
//
// If you want to accept the high-level Connection, not a #GSocket, which is
// often the case, then you should use g_socket_listener_accept() instead.
//
// If source_object is not NULL it will be filled out with the source object
// specified when the corresponding socket or address was added to the listener.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned.
func (listener *SocketListener) AcceptSocket(cancellable *Cancellable) (*externglib.Object, *Socket, error) {
	var _arg0 *C.GSocketListener // out
	var _arg1 *C.GObject         // in
	var _arg2 *C.GCancellable    // out
	var _cret *C.GSocket         // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(listener.Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_socket_listener_accept_socket(_arg0, &_arg1, _arg2, &_cerr)

	var _sourceObject *externglib.Object // out
	var _socket *Socket                  // out
	var _goerr error                     // out

	_sourceObject = externglib.Take(unsafe.Pointer(_arg1))
	_socket = wrapSocket(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _sourceObject, _socket, _goerr
}

// AcceptSocketAsync: this is the asynchronous version of
// g_socket_listener_accept_socket().
//
// When the operation is finished callback will be called. You can then call
// g_socket_listener_accept_socket_finish() to get the result of the operation.
func (listener *SocketListener) AcceptSocketAsync(cancellable *Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GSocketListener    // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(listener.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg3 = C.gpointer(gbox.AssignOnce(callback))

	C.g_socket_listener_accept_socket_async(_arg0, _arg1, _arg2, _arg3)
}

// AcceptSocketFinish finishes an async accept operation. See
// g_socket_listener_accept_socket_async()
func (listener *SocketListener) AcceptSocketFinish(result AsyncResulter) (*externglib.Object, *Socket, error) {
	var _arg0 *C.GSocketListener // out
	var _arg1 *C.GAsyncResult    // out
	var _arg2 *C.GObject         // in
	var _cret *C.GSocket         // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(listener.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.g_socket_listener_accept_socket_finish(_arg0, _arg1, &_arg2, &_cerr)

	var _sourceObject *externglib.Object // out
	var _socket *Socket                  // out
	var _goerr error                     // out

	_sourceObject = externglib.Take(unsafe.Pointer(_arg2))
	_socket = wrapSocket(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _sourceObject, _socket, _goerr
}

// AddAddress creates a socket of type type and protocol protocol, binds it to
// address and adds it to the set of sockets we're accepting sockets from.
//
// Note that adding an IPv6 address, depending on the platform, may or may not
// result in a listener that also accepts IPv4 connections. For more
// deterministic behavior, see g_socket_listener_add_inet_port().
//
// source_object will be passed out in the various calls to accept to identify
// this particular source, which is useful if you're listening on multiple
// addresses and do different things depending on what address is connected to.
//
// If successful and effective_address is non-NULL then it will be set to the
// address that the binding actually occurred at. This is helpful for
// determining the port number that was used for when requesting a binding to
// port 0 (ie: "any port"). This address, if requested, belongs to the caller
// and must be freed.
//
// Call g_socket_listener_close() to stop listening on address; this will not be
// done automatically when you drop your final reference to listener, as
// references may be held internally.
func (listener *SocketListener) AddAddress(address SocketAddresser, typ SocketType, protocol SocketProtocol, sourceObject *externglib.Object) (*SocketAddress, error) {
	var _arg0 *C.GSocketListener // out
	var _arg1 *C.GSocketAddress  // out
	var _arg2 C.GSocketType      // out
	var _arg3 C.GSocketProtocol  // out
	var _arg4 *C.GObject         // out
	var _arg5 *C.GSocketAddress  // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(listener.Native()))
	_arg1 = (*C.GSocketAddress)(unsafe.Pointer((address).(gextras.Nativer).Native()))
	_arg2 = C.GSocketType(typ)
	_arg3 = C.GSocketProtocol(protocol)
	_arg4 = (*C.GObject)(unsafe.Pointer(sourceObject.Native()))

	C.g_socket_listener_add_address(_arg0, _arg1, _arg2, _arg3, _arg4, &_arg5, &_cerr)

	var _effectiveAddress *SocketAddress // out
	var _goerr error                     // out

	_effectiveAddress = wrapSocketAddress(externglib.AssumeOwnership(unsafe.Pointer(_arg5)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _effectiveAddress, _goerr
}

// AddAnyInetPort listens for TCP connections on any available port number for
// both IPv6 and IPv4 (if each is available).
//
// This is useful if you need to have a socket for incoming connections but
// don't care about the specific port number.
//
// source_object will be passed out in the various calls to accept to identify
// this particular source, which is useful if you're listening on multiple
// addresses and do different things depending on what address is connected to.
func (listener *SocketListener) AddAnyInetPort(sourceObject *externglib.Object) (uint16, error) {
	var _arg0 *C.GSocketListener // out
	var _arg1 *C.GObject         // out
	var _cret C.guint16          // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(listener.Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer(sourceObject.Native()))

	_cret = C.g_socket_listener_add_any_inet_port(_arg0, _arg1, &_cerr)

	var _guint16 uint16 // out
	var _goerr error    // out

	_guint16 = uint16(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _guint16, _goerr
}

// AddInetPort: helper function for g_socket_listener_add_address() that creates
// a TCP/IP socket listening on IPv4 and IPv6 (if supported) on the specified
// port on all interfaces.
//
// source_object will be passed out in the various calls to accept to identify
// this particular source, which is useful if you're listening on multiple
// addresses and do different things depending on what address is connected to.
//
// Call g_socket_listener_close() to stop listening on port; this will not be
// done automatically when you drop your final reference to listener, as
// references may be held internally.
func (listener *SocketListener) AddInetPort(port uint16, sourceObject *externglib.Object) error {
	var _arg0 *C.GSocketListener // out
	var _arg1 C.guint16          // out
	var _arg2 *C.GObject         // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(listener.Native()))
	_arg1 = C.guint16(port)
	_arg2 = (*C.GObject)(unsafe.Pointer(sourceObject.Native()))

	C.g_socket_listener_add_inet_port(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// AddSocket adds socket to the set of sockets that we try to accept new clients
// from. The socket must be bound to a local address and listened to.
//
// source_object will be passed out in the various calls to accept to identify
// this particular source, which is useful if you're listening on multiple
// addresses and do different things depending on what address is connected to.
//
// The socket will not be automatically closed when the listener is finalized
// unless the listener held the final reference to the socket. Before GLib 2.42,
// the socket was automatically closed on finalization of the listener, even if
// references to it were held elsewhere.
func (listener *SocketListener) AddSocket(socket *Socket, sourceObject *externglib.Object) error {
	var _arg0 *C.GSocketListener // out
	var _arg1 *C.GSocket         // out
	var _arg2 *C.GObject         // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(listener.Native()))
	_arg1 = (*C.GSocket)(unsafe.Pointer(socket.Native()))
	_arg2 = (*C.GObject)(unsafe.Pointer(sourceObject.Native()))

	C.g_socket_listener_add_socket(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// Close closes all the sockets in the listener.
func (listener *SocketListener) Close() {
	var _arg0 *C.GSocketListener // out

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(listener.Native()))

	C.g_socket_listener_close(_arg0)
}

// SetBacklog sets the listen backlog on the sockets in the listener. This must
// be called before adding any sockets, addresses or ports to the Listener (for
// example, by calling g_socket_listener_add_inet_port()) to be effective.
//
// See g_socket_set_listen_backlog() for details
func (listener *SocketListener) SetBacklog(listenBacklog int) {
	var _arg0 *C.GSocketListener // out
	var _arg1 C.int              // out

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(listener.Native()))
	_arg1 = C.int(listenBacklog)

	C.g_socket_listener_set_backlog(_arg0, _arg1)
}
