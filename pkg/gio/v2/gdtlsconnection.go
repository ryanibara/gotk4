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
// #include <glib-object.h>
// extern gboolean _gotk4_gio2_DtlsConnectionInterface_handshake(void*, void*, GError**);
// extern gboolean _gotk4_gio2_DtlsConnectionInterface_handshake_finish(void*, void*, GError**);
// extern gboolean _gotk4_gio2_DtlsConnectionInterface_shutdown(void*, gboolean, gboolean, void*, GError**);
// extern gboolean _gotk4_gio2_DtlsConnectionInterface_shutdown_finish(void*, void*, GError**);
// extern gchar* _gotk4_gio2_DtlsConnectionInterface_get_negotiated_protocol(void*);
// extern void _gotk4_gio2_AsyncReadyCallback(void*, void*, gpointer);
// extern void _gotk4_gio2_DtlsConnectionInterface_set_advertised_protocols(void*, gchar**);
import "C"

// GTypeDTLSConnection returns the GType for the type DTLSConnection.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDTLSConnection() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "DtlsConnection").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDTLSConnection)
	return gtype
}

// DTLSConnection is the base DTLS connection class type, which wraps a Based
// and provides DTLS encryption on top of it. Its subclasses, ClientConnection
// and ServerConnection, implement client-side and server-side DTLS,
// respectively.
//
// For TLS support, see Connection.
//
// As DTLS is datagram based, Connection implements Based, presenting a
// datagram-socket-like API for the encrypted connection. This operates over a
// base datagram connection, which is also a Based (Connection:base-socket).
//
// To close a DTLS connection, use g_dtls_connection_close().
//
// Neither ServerConnection or ClientConnection set the peer address on their
// base Based if it is a #GSocket — it is up to the caller to do that if they
// wish. If they do not, and g_socket_close() is called on the base socket, the
// Connection will not raise a G_IO_ERROR_NOT_CONNECTED error on further I/O.
//
// DTLSConnection wraps an interface. This means the user can get the
// underlying type by calling Cast().
type DTLSConnection struct {
	_ [0]func() // equal guard
	DatagramBased
}

var ()

// DTLSConnectioner describes DTLSConnection's interface methods.
type DTLSConnectioner interface {
	coreglib.Objector

	// Close the DTLS connection.
	Close(ctx context.Context) error
	// CloseAsync: asynchronously close the DTLS connection.
	CloseAsync(ctx context.Context, ioPriority int32, callback AsyncReadyCallback)
	// CloseFinish: finish an asynchronous TLS close operation.
	CloseFinish(result AsyncResulter) error
	// Certificate gets conn's certificate, as set by
	// g_dtls_connection_set_certificate().
	Certificate() TLSCertificater
	// Database gets the certificate database that conn uses to verify peer
	// certificates.
	Database() TLSDatabaser
	// Interaction: get the object that will be used to interact with the user.
	Interaction() *TLSInteraction
	// NegotiatedProtocol gets the name of the application-layer protocol
	// negotiated during the handshake.
	NegotiatedProtocol() string
	// PeerCertificate gets conn's peer's certificate after the handshake has
	// completed or failed.
	PeerCertificate() TLSCertificater
	// RequireCloseNotify tests whether or not conn expects a proper TLS close
	// notification when the connection is closed.
	RequireCloseNotify() bool
	// Handshake attempts a TLS handshake on conn.
	Handshake(ctx context.Context) error
	// HandshakeAsync: asynchronously performs a TLS handshake on conn.
	HandshakeAsync(ctx context.Context, ioPriority int32, callback AsyncReadyCallback)
	// HandshakeFinish: finish an asynchronous TLS handshake operation.
	HandshakeFinish(result AsyncResulter) error
	// SetAdvertisedProtocols sets the list of application-layer protocols to
	// advertise that the caller is willing to speak on this connection.
	SetAdvertisedProtocols(protocols []string)
	// SetCertificate: this sets the certificate that conn will present to its
	// peer during the TLS handshake.
	SetCertificate(certificate TLSCertificater)
	// SetDatabase sets the certificate database that is used to verify peer
	// certificates.
	SetDatabase(database TLSDatabaser)
	// SetInteraction: set the object that will be used to interact with the
	// user.
	SetInteraction(interaction *TLSInteraction)
	// SetRequireCloseNotify sets whether or not conn expects a proper TLS close
	// notification before the connection is closed.
	SetRequireCloseNotify(requireCloseNotify bool)
	// Shutdown: shut down part or all of a DTLS connection.
	Shutdown(ctx context.Context, shutdownRead, shutdownWrite bool) error
	// ShutdownAsync: asynchronously shut down part or all of the DTLS
	// connection.
	ShutdownAsync(ctx context.Context, shutdownRead, shutdownWrite bool, ioPriority int32, callback AsyncReadyCallback)
	// ShutdownFinish: finish an asynchronous TLS shutdown operation.
	ShutdownFinish(result AsyncResulter) error
}

var _ DTLSConnectioner = (*DTLSConnection)(nil)

func wrapDTLSConnection(obj *coreglib.Object) *DTLSConnection {
	return &DTLSConnection{
		DatagramBased: DatagramBased{
			Object: obj,
		},
	}
}

func marshalDTLSConnection(p uintptr) (interface{}, error) {
	return wrapDTLSConnection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Close the DTLS connection. This is equivalent to calling
// g_dtls_connection_shutdown() to shut down both sides of the connection.
//
// Closing a Connection waits for all buffered but untransmitted data to be sent
// before it completes. It then sends a close_notify DTLS alert to the peer and
// may wait for a close_notify to be received from the peer. It does not close
// the underlying Connection:base-socket; that must be closed separately.
//
// Once conn is closed, all other operations will return G_IO_ERROR_CLOSED.
// Closing a Connection multiple times will not return an error.
//
// Connections will be automatically closed when the last reference is dropped,
// but you might want to call this function to make sure resources are released
// as early as possible.
//
// If cancellable is cancelled, the Connection may be left partially-closed and
// any pending untransmitted data may be lost. Call g_dtls_connection_close()
// again to complete closing the Connection.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//
func (conn *DTLSConnection) Close(ctx context.Context) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[1] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_info.InvokeIfaceMethod("close", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(ctx)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cerr))))
	}

	return _goerr
}

// CloseAsync: asynchronously close the DTLS connection. See
// g_dtls_connection_close() for more information.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback (optional) to call when the close operation is complete.
//
func (conn *DTLSConnection) CloseAsync(ctx context.Context, ioPriority int32, callback AsyncReadyCallback) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(ioPriority)
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[3])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[4] = C.gpointer(gbox.AssignOnce(callback))
	}

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_info.InvokeIfaceMethod("close_async", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// CloseFinish: finish an asynchronous TLS close operation. See
// g_dtls_connection_close() for more information.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (conn *DTLSConnection) CloseFinish(result AsyncResulter) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_info.InvokeIfaceMethod("close_finish", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(result)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cerr))))
	}

	return _goerr
}

// Certificate gets conn's certificate, as set by
// g_dtls_connection_set_certificate().
//
// The function returns the following values:
//
//    - tlsCertificate (optional) conn's certificate, or NULL.
//
func (conn *DTLSConnection) Certificate() TLSCertificater {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_gret := _info.InvokeIfaceMethod("get_certificate", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(conn)

	var _tlsCertificate TLSCertificater // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(TLSCertificater)
				return ok
			})
			rv, ok := casted.(TLSCertificater)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.TLSCertificater")
			}
			_tlsCertificate = rv
		}
	}

	return _tlsCertificate
}

// Database gets the certificate database that conn uses to verify peer
// certificates. See g_dtls_connection_set_database().
//
// The function returns the following values:
//
//    - tlsDatabase (optional): certificate database that conn uses or NULL.
//
func (conn *DTLSConnection) Database() TLSDatabaser {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_gret := _info.InvokeIfaceMethod("get_database", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(conn)

	var _tlsDatabase TLSDatabaser // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(TLSDatabaser)
				return ok
			})
			rv, ok := casted.(TLSDatabaser)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.TLSDatabaser")
			}
			_tlsDatabase = rv
		}
	}

	return _tlsDatabase
}

// Interaction: get the object that will be used to interact with the user. It
// will be used for things like prompting the user for passwords. If NULL is
// returned, then no user interaction will occur for this connection.
//
// The function returns the following values:
//
//    - tlsInteraction (optional): interaction object.
//
func (conn *DTLSConnection) Interaction() *TLSInteraction {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_gret := _info.InvokeIfaceMethod("get_interaction", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(conn)

	var _tlsInteraction *TLSInteraction // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_tlsInteraction = wrapTLSInteraction(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	}

	return _tlsInteraction
}

// NegotiatedProtocol gets the name of the application-layer protocol negotiated
// during the handshake.
//
// If the peer did not use the ALPN extension, or did not advertise a protocol
// that matched one of conn's protocols, or the TLS backend does not support
// ALPN, then this will be NULL. See
// g_dtls_connection_set_advertised_protocols().
//
// The function returns the following values:
//
//    - utf8 (optional): negotiated protocol, or NULL.
//
func (conn *DTLSConnection) NegotiatedProtocol() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_gret := _info.InvokeIfaceMethod("get_negotiated_protocol", _args[:], nil)
	_cret := *(**C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(conn)

	var _utf8 string // out

	if *(**C.gchar)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_cret)))))
	}

	return _utf8
}

// PeerCertificate gets conn's peer's certificate after the handshake has
// completed or failed. (It is not set during the emission of
// Connection::accept-certificate.).
//
// The function returns the following values:
//
//    - tlsCertificate (optional) conn's peer's certificate, or NULL.
//
func (conn *DTLSConnection) PeerCertificate() TLSCertificater {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_gret := _info.InvokeIfaceMethod("get_peer_certificate", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(conn)

	var _tlsCertificate TLSCertificater // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(TLSCertificater)
				return ok
			})
			rv, ok := casted.(TLSCertificater)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.TLSCertificater")
			}
			_tlsCertificate = rv
		}
	}

	return _tlsCertificate
}

// RequireCloseNotify tests whether or not conn expects a proper TLS close
// notification when the connection is closed. See
// g_dtls_connection_set_require_close_notify() for details.
//
// The function returns the following values:
//
//    - ok: TRUE if conn requires a proper TLS close notification.
//
func (conn *DTLSConnection) RequireCloseNotify() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_gret := _info.InvokeIfaceMethod("get_require_close_notify", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(conn)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Handshake attempts a TLS handshake on conn.
//
// On the client side, it is never necessary to call this method; although the
// connection needs to perform a handshake after connecting, Connection will
// handle this for you automatically when you try to send or receive data on the
// connection. You can call g_dtls_connection_handshake() manually if you want
// to know whether the initial handshake succeeded or failed (as opposed to just
// immediately trying to use conn to read or write, in which case, if it fails,
// it may not be possible to tell if it failed before or after completing the
// handshake), but beware that servers may reject client authentication after
// the handshake has completed, so a successful handshake does not indicate the
// connection will be usable.
//
// Likewise, on the server side, although a handshake is necessary at the
// beginning of the communication, you do not need to call this function
// explicitly unless you want clearer error reporting.
//
// Previously, calling g_dtls_connection_handshake() after the initial handshake
// would trigger a rehandshake; however, this usage was deprecated in GLib 2.60
// because rehandshaking was removed from the TLS protocol in TLS 1.3. Since
// GLib 2.64, calling this function after the initial handshake will no longer
// do anything.
//
// Connection::accept_certificate may be emitted during the handshake.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//
func (conn *DTLSConnection) Handshake(ctx context.Context) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[1] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_info.InvokeIfaceMethod("handshake", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(ctx)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cerr))))
	}

	return _goerr
}

// HandshakeAsync: asynchronously performs a TLS handshake on conn. See
// g_dtls_connection_handshake() for more information.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback (optional) to call when the handshake is complete.
//
func (conn *DTLSConnection) HandshakeAsync(ctx context.Context, ioPriority int32, callback AsyncReadyCallback) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(ioPriority)
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[3])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[4] = C.gpointer(gbox.AssignOnce(callback))
	}

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_info.InvokeIfaceMethod("handshake_async", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// HandshakeFinish: finish an asynchronous TLS handshake operation. See
// g_dtls_connection_handshake() for more information.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (conn *DTLSConnection) HandshakeFinish(result AsyncResulter) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_info.InvokeIfaceMethod("handshake_finish", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(result)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cerr))))
	}

	return _goerr
}

// SetAdvertisedProtocols sets the list of application-layer protocols to
// advertise that the caller is willing to speak on this connection. The
// Application-Layer Protocol Negotiation (ALPN) extension will be used to
// negotiate a compatible protocol with the peer; use
// g_dtls_connection_get_negotiated_protocol() to find the negotiated protocol
// after the handshake. Specifying NULL for the the value of protocols will
// disable ALPN negotiation.
//
// See IANA TLS ALPN Protocol IDs
// (https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids)
// for a list of registered protocol IDs.
//
// The function takes the following parameters:
//
//    - protocols (optional): NULL-terminated array of ALPN protocol names (eg,
//      "http/1.1", "h2"), or NULL.
//
func (conn *DTLSConnection) SetAdvertisedProtocols(protocols []string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	{
		*(***C.gchar)(unsafe.Pointer(&_args[1])) = (**C.gchar)(C.calloc(C.size_t((len(protocols) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(*(***C.gchar)(unsafe.Pointer(&_args[1]))))
		{
			out := unsafe.Slice(_args[1], len(protocols)+1)
			var zero *C.gchar
			out[len(protocols)] = zero
			for i := range protocols {
				*(**C.gchar)(unsafe.Pointer(&out[i])) = (*C.gchar)(unsafe.Pointer(C.CString(protocols[i])))
				defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&out[i]))))
			}
		}
	}

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_info.InvokeIfaceMethod("set_advertised_protocols", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(protocols)
}

// SetCertificate: this sets the certificate that conn will present to its peer
// during the TLS handshake. For a ServerConnection, it is mandatory to set
// this, and that will normally be done at construct time.
//
// For a ClientConnection, this is optional. If a handshake fails with
// G_TLS_ERROR_CERTIFICATE_REQUIRED, that means that the server requires a
// certificate, and if you try connecting again, you should call this method
// first. You can call g_dtls_client_connection_get_accepted_cas() on the failed
// connection to get a list of Certificate Authorities that the server will
// accept certificates from.
//
// (It is also possible that a server will allow the connection with or without
// a certificate; in that case, if you don't provide a certificate, you can tell
// that the server requested one by the fact that
// g_dtls_client_connection_get_accepted_cas() will return non-NULL.).
//
// The function takes the following parameters:
//
//    - certificate to use for conn.
//
func (conn *DTLSConnection) SetCertificate(certificate TLSCertificater) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(certificate).Native()))

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_info.InvokeIfaceMethod("set_certificate", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(certificate)
}

// SetDatabase sets the certificate database that is used to verify peer
// certificates. This is set to the default database by default. See
// g_tls_backend_get_default_database(). If set to NULL, then peer certificate
// validation will always set the G_TLS_CERTIFICATE_UNKNOWN_CA error (meaning
// Connection::accept-certificate will always be emitted on client-side
// connections, unless that bit is not set in
// ClientConnection:validation-flags).
//
// The function takes the following parameters:
//
//    - database (optional): Database.
//
func (conn *DTLSConnection) SetDatabase(database TLSDatabaser) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	if database != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(database).Native()))
	}

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_info.InvokeIfaceMethod("set_database", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(database)
}

// SetInteraction: set the object that will be used to interact with the user.
// It will be used for things like prompting the user for passwords.
//
// The interaction argument will normally be a derived subclass of Interaction.
// NULL can also be provided if no user interaction should occur for this
// connection.
//
// The function takes the following parameters:
//
//    - interaction (optional) object, or NULL.
//
func (conn *DTLSConnection) SetInteraction(interaction *TLSInteraction) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	if interaction != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(interaction).Native()))
	}

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_info.InvokeIfaceMethod("set_interaction", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(interaction)
}

// SetRequireCloseNotify sets whether or not conn expects a proper TLS close
// notification before the connection is closed. If this is TRUE (the default),
// then conn will expect to receive a TLS close notification from its peer
// before the connection is closed, and will return a G_TLS_ERROR_EOF error if
// the connection is closed without proper notification (since this may indicate
// a network error, or man-in-the-middle attack).
//
// In some protocols, the application will know whether or not the connection
// was closed cleanly based on application-level data (because the
// application-level data includes a length field, or is somehow
// self-delimiting); in this case, the close notify is redundant and may be
// omitted. You can use g_dtls_connection_set_require_close_notify() to tell
// conn to allow an "unannounced" connection close, in which case the close will
// show up as a 0-length read, as in a non-TLS Based, and it is up to the
// application to check that the data has been fully received.
//
// Note that this only affects the behavior when the peer closes the connection;
// when the application calls g_dtls_connection_close_async() on conn itself,
// this will send a close notification regardless of the setting of this
// property. If you explicitly want to do an unclean close, you can close conn's
// Connection:base-socket rather than closing conn itself.
//
// The function takes the following parameters:
//
//    - requireCloseNotify: whether or not to require close notification.
//
func (conn *DTLSConnection) SetRequireCloseNotify(requireCloseNotify bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	if requireCloseNotify {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_info.InvokeIfaceMethod("set_require_close_notify", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(requireCloseNotify)
}

// Shutdown: shut down part or all of a DTLS connection.
//
// If shutdown_read is TRUE then the receiving side of the connection is shut
// down, and further reading is disallowed. Subsequent calls to
// g_datagram_based_receive_messages() will return G_IO_ERROR_CLOSED.
//
// If shutdown_write is TRUE then the sending side of the connection is shut
// down, and further writing is disallowed. Subsequent calls to
// g_datagram_based_send_messages() will return G_IO_ERROR_CLOSED.
//
// It is allowed for both shutdown_read and shutdown_write to be TRUE — this is
// equivalent to calling g_dtls_connection_close().
//
// If cancellable is cancelled, the Connection may be left partially-closed and
// any pending untransmitted data may be lost. Call g_dtls_connection_shutdown()
// again to complete closing the Connection.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - shutdownRead: TRUE to stop reception of incoming datagrams.
//    - shutdownWrite: TRUE to stop sending outgoing datagrams.
//
func (conn *DTLSConnection) Shutdown(ctx context.Context, shutdownRead, shutdownWrite bool) error {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[3] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	if shutdownRead {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}
	if shutdownWrite {
		*(*C.gboolean)(unsafe.Pointer(&_args[2])) = C.TRUE
	}

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_info.InvokeIfaceMethod("shutdown", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(shutdownRead)
	runtime.KeepAlive(shutdownWrite)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cerr))))
	}

	return _goerr
}

// ShutdownAsync: asynchronously shut down part or all of the DTLS connection.
// See g_dtls_connection_shutdown() for more information.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - shutdownRead: TRUE to stop reception of incoming datagrams.
//    - shutdownWrite: TRUE to stop sending outgoing datagrams.
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback (optional) to call when the shutdown operation is complete.
//
func (conn *DTLSConnection) ShutdownAsync(ctx context.Context, shutdownRead, shutdownWrite bool, ioPriority int32, callback AsyncReadyCallback) {
	var _args [7]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[4] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	if shutdownRead {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}
	if shutdownWrite {
		*(*C.gboolean)(unsafe.Pointer(&_args[2])) = C.TRUE
	}
	*(*C.int)(unsafe.Pointer(&_args[3])) = C.int(ioPriority)
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[5])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[6] = C.gpointer(gbox.AssignOnce(callback))
	}

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_info.InvokeIfaceMethod("shutdown_async", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(shutdownRead)
	runtime.KeepAlive(shutdownWrite)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// ShutdownFinish: finish an asynchronous TLS shutdown operation. See
// g_dtls_connection_shutdown() for more information.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (conn *DTLSConnection) ShutdownFinish(result AsyncResulter) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_info := girepository.MustFind("Gio", "DtlsConnection")
	_info.InvokeIfaceMethod("shutdown_finish", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(result)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cerr))))
	}

	return _goerr
}
