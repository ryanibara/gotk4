// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gio2_TlsClientConnectionInterface_copy_session_state(void*, void*);
import "C"

// glib.Type values for gtlsclientconnection.go.
var GTypeTLSClientConnection = coreglib.Type(C.g_tls_client_connection_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeTLSClientConnection, F: marshalTLSClientConnection},
	})
}

// TLSClientConnectionOverrider contains methods that are overridable.
type TLSClientConnectionOverrider interface {
	// CopySessionState: possibly copies session state from one connection to
	// another, for use in TLS session resumption. This is not normally needed,
	// but may be used when the same session needs to be used between different
	// endpoints, as is required by some protocols, such as FTP over TLS. source
	// should have already completed a handshake and, since TLS 1.3, it should
	// have been used to read data at least once. conn should not have completed
	// a handshake.
	//
	// It is not possible to know whether a call to this function will actually
	// do anything. Because session resumption is normally used only for
	// performance benefit, the TLS backend might not implement this function.
	// Even if implemented, it may not actually succeed in allowing conn to
	// resume source's TLS session, because the server may not have sent a
	// session resumption token to source, or it may refuse to accept the token
	// from conn. There is no way to know whether a call to this function is
	// actually successful.
	//
	// Using this function is not required to benefit from session resumption.
	// If the TLS backend supports session resumption, the session will be
	// resumed automatically if it is possible to do so without weakening the
	// privacy guarantees normally provided by TLS, without need to call this
	// function. For example, with TLS 1.3, a session ticket will be
	// automatically copied from any ClientConnection that has previously
	// received session tickets from the server, provided a ticket is available
	// that has not previously been used for session resumption, since session
	// ticket reuse would be a privacy weakness. Using this function causes the
	// ticket to be copied without regard for privacy considerations.
	//
	// The function takes the following parameters:
	//
	//    - source: ClientConnection.
	//
	CopySessionState(source TLSClientConnectioner)
}

// TLSClientConnection is the client-side subclass of Connection, representing a
// client-side TLS connection.
//
// TLSClientConnection wraps an interface. This means the user can get the
// underlying type by calling Cast().
type TLSClientConnection struct {
	_ [0]func() // equal guard
	TLSConnection
}

var (
	_ TLSConnectioner = (*TLSClientConnection)(nil)
)

// TLSClientConnectioner describes TLSClientConnection's interface methods.
type TLSClientConnectioner interface {
	coreglib.Objector

	// CopySessionState: possibly copies session state from one connection to
	// another, for use in TLS session resumption.
	CopySessionState(source TLSClientConnectioner)
	// ServerIdentity gets conn's expected server identity.
	ServerIdentity() *SocketConnectable
	// UseSSL3: SSL 3.0 is no longer supported.
	UseSSL3() bool
	// SetServerIdentity sets conn's expected server identity, which is used
	// both to tell servers on virtual hosts which certificate to present, and
	// also to let conn know what name to look for in the certificate when
	// performing G_TLS_CERTIFICATE_BAD_IDENTITY validation, if enabled.
	SetServerIdentity(identity SocketConnectabler)
	// SetUseSSL3: since GLib 2.42.1, SSL 3.0 is no longer supported.
	SetUseSSL3(useSsl3 bool)
}

var _ TLSClientConnectioner = (*TLSClientConnection)(nil)

func ifaceInitTLSClientConnectioner(gifacePtr, data C.gpointer) {
	iface := (*C.GTlsClientConnectionInterface)(unsafe.Pointer(gifacePtr))
	iface.copy_session_state = (*[0]byte)(C._gotk4_gio2_TlsClientConnectionInterface_copy_session_state)
}

//export _gotk4_gio2_TlsClientConnectionInterface_copy_session_state
func _gotk4_gio2_TlsClientConnectionInterface_copy_session_state(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TLSClientConnectionOverrider)

	var _source TLSClientConnectioner // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.TLSClientConnectioner is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(TLSClientConnectioner)
			return ok
		})
		rv, ok := casted.(TLSClientConnectioner)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.TLSClientConnectioner")
		}
		_source = rv
	}

	iface.CopySessionState(_source)
}

func wrapTLSClientConnection(obj *coreglib.Object) *TLSClientConnection {
	return &TLSClientConnection{
		TLSConnection: TLSConnection{
			IOStream: IOStream{
				Object: obj,
			},
		},
	}
}

func marshalTLSClientConnection(p uintptr) (interface{}, error) {
	return wrapTLSClientConnection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CopySessionState: possibly copies session state from one connection to
// another, for use in TLS session resumption. This is not normally needed, but
// may be used when the same session needs to be used between different
// endpoints, as is required by some protocols, such as FTP over TLS. source
// should have already completed a handshake and, since TLS 1.3, it should have
// been used to read data at least once. conn should not have completed a
// handshake.
//
// It is not possible to know whether a call to this function will actually do
// anything. Because session resumption is normally used only for performance
// benefit, the TLS backend might not implement this function. Even if
// implemented, it may not actually succeed in allowing conn to resume source's
// TLS session, because the server may not have sent a session resumption token
// to source, or it may refuse to accept the token from conn. There is no way to
// know whether a call to this function is actually successful.
//
// Using this function is not required to benefit from session resumption. If
// the TLS backend supports session resumption, the session will be resumed
// automatically if it is possible to do so without weakening the privacy
// guarantees normally provided by TLS, without need to call this function. For
// example, with TLS 1.3, a session ticket will be automatically copied from any
// ClientConnection that has previously received session tickets from the
// server, provided a ticket is available that has not previously been used for
// session resumption, since session ticket reuse would be a privacy weakness.
// Using this function causes the ticket to be copied without regard for privacy
// considerations.
//
// The function takes the following parameters:
//
//    - source: ClientConnection.
//
func (conn *TLSClientConnection) CopySessionState(source TLSClientConnectioner) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(source).Native()))

	runtime.KeepAlive(conn)
	runtime.KeepAlive(source)
}

// ServerIdentity gets conn's expected server identity.
//
// The function returns the following values:
//
//    - socketConnectable (optional) describing the expected server identity, or
//      NULL if the expected identity is not known.
//
func (conn *TLSClientConnection) ServerIdentity() *SocketConnectable {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(conn)

	var _socketConnectable *SocketConnectable // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_socketConnectable = wrapSocketConnectable(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _socketConnectable
}

// UseSSL3: SSL 3.0 is no longer supported. See
// g_tls_client_connection_set_use_ssl3() for details.
//
// Deprecated: SSL 3.0 is insecure.
//
// The function returns the following values:
//
//    - ok: FALSE.
//
func (conn *TLSClientConnection) UseSSL3() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(conn)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetServerIdentity sets conn's expected server identity, which is used both to
// tell servers on virtual hosts which certificate to present, and also to let
// conn know what name to look for in the certificate when performing
// G_TLS_CERTIFICATE_BAD_IDENTITY validation, if enabled.
//
// The function takes the following parameters:
//
//    - identity describing the expected server identity.
//
func (conn *TLSClientConnection) SetServerIdentity(identity SocketConnectabler) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(identity).Native()))

	runtime.KeepAlive(conn)
	runtime.KeepAlive(identity)
}

// SetUseSSL3: since GLib 2.42.1, SSL 3.0 is no longer supported.
//
// From GLib 2.42.1 through GLib 2.62, this function could be used to force use
// of TLS 1.0, the lowest-supported TLS protocol version at the time. In the
// past, this was needed to connect to broken TLS servers that exhibited
// protocol version intolerance. Such servers are no longer common, and using
// TLS 1.0 is no longer considered acceptable.
//
// Since GLib 2.64, this function does nothing.
//
// Deprecated: SSL 3.0 is insecure.
//
// The function takes the following parameters:
//
//    - useSsl3: #gboolean, ignored.
//
func (conn *TLSClientConnection) SetUseSSL3(useSsl3 bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	if useSsl3 {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	runtime.KeepAlive(conn)
	runtime.KeepAlive(useSsl3)
}

// NewTLSClientConnection creates a new ClientConnection wrapping base_io_stream
// (which must have pollable input and output streams) which is assumed to
// communicate with the server identified by server_identity.
//
// See the documentation for Connection:base-io-stream for restrictions on when
// application code can run operations on the base_io_stream after this function
// has returned.
//
// The function takes the following parameters:
//
//    - baseIoStream to wrap.
//    - serverIdentity (optional): expected identity of the server.
//
// The function returns the following values:
//
//    - tlsClientConnection: new ClientConnection, or NULL on error.
//
func NewTLSClientConnection(baseIoStream IOStreamer, serverIdentity SocketConnectabler) (*TLSClientConnection, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(baseIoStream).Native()))
	if serverIdentity != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(serverIdentity).Native()))
	}

	_gret := girepository.MustFind("Gio", "new").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(baseIoStream)
	runtime.KeepAlive(serverIdentity)

	var _tlsClientConnection *TLSClientConnection // out
	var _goerr error                              // out

	_tlsClientConnection = wrapTLSClientConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsClientConnection, _goerr
}
