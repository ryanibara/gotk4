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
		{T: externglib.Type(C.g_tls_connection_get_type()), F: marshalTLSConnectioner},
	})
}

// TLSConnectionOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TLSConnectionOverrider interface {
	AcceptCertificate(peerCert TLSCertificater, errors TLSCertificateFlags) bool
	BindingData(typ TLSChannelBindingType, data []byte) error
	// Handshake attempts a TLS handshake on conn.
	//
	// On the client side, it is never necessary to call this method; although
	// the connection needs to perform a handshake after connecting (or after
	// sending a "STARTTLS"-type command), Connection will handle this for you
	// automatically when you try to send or receive data on the connection. You
	// can call g_tls_connection_handshake() manually if you want to know
	// whether the initial handshake succeeded or failed (as opposed to just
	// immediately trying to use conn to read or write, in which case, if it
	// fails, it may not be possible to tell if it failed before or after
	// completing the handshake), but beware that servers may reject client
	// authentication after the handshake has completed, so a successful
	// handshake does not indicate the connection will be usable.
	//
	// Likewise, on the server side, although a handshake is necessary at the
	// beginning of the communication, you do not need to call this function
	// explicitly unless you want clearer error reporting.
	//
	// Previously, calling g_tls_connection_handshake() after the initial
	// handshake would trigger a rehandshake; however, this usage was deprecated
	// in GLib 2.60 because rehandshaking was removed from the TLS protocol in
	// TLS 1.3. Since GLib 2.64, calling this function after the initial
	// handshake will no longer do anything.
	//
	// When using a Connection created by Client, the Client performs the
	// initial handshake, so calling this function manually is not recommended.
	//
	// Connection::accept_certificate may be emitted during the handshake.
	Handshake(ctx context.Context) error
	// HandshakeAsync: asynchronously performs a TLS handshake on conn. See
	// g_tls_connection_handshake() for more information.
	HandshakeAsync(ctx context.Context, ioPriority int, callback AsyncReadyCallback)
	// HandshakeFinish: finish an asynchronous TLS handshake operation. See
	// g_tls_connection_handshake() for more information.
	HandshakeFinish(result AsyncResulter) error
}

// TLSConnection is the base TLS connection class type, which wraps a OStream
// and provides TLS encryption on top of it. Its subclasses, ClientConnection
// and ServerConnection, implement client-side and server-side TLS,
// respectively.
//
// For DTLS (Datagram TLS) support, see Connection.
type TLSConnection struct {
	IOStream
}

var (
	_ IOStreamer = (*TLSConnection)(nil)
)

// TLSConnectioner describes types inherited from class TLSConnection.

// To get the original type, the caller must assert this to an interface or
// another type.
type TLSConnectioner interface {
	externglib.Objector
	baseTLSConnection() *TLSConnection
}

var _ TLSConnectioner = (*TLSConnection)(nil)

func wrapTLSConnection(obj *externglib.Object) *TLSConnection {
	return &TLSConnection{
		IOStream: IOStream{
			Object: obj,
		},
	}
}

func marshalTLSConnectioner(p uintptr) (interface{}, error) {
	return wrapTLSConnection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// EmitAcceptCertificate: used by Connection implementations to emit the
// Connection::accept-certificate signal.
//
// The function takes the following parameters:
//
//    - peerCert peer's Certificate.
//    - errors problems with peer_cert.
//
func (conn *TLSConnection) EmitAcceptCertificate(peerCert TLSCertificater, errors TLSCertificateFlags) bool {
	var _arg0 *C.GTlsConnection      // out
	var _arg1 *C.GTlsCertificate     // out
	var _arg2 C.GTlsCertificateFlags // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(peerCert.Native()))
	_arg2 = C.GTlsCertificateFlags(errors)

	_cret = C.g_tls_connection_emit_accept_certificate(_arg0, _arg1, _arg2)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(peerCert)
	runtime.KeepAlive(errors)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Certificate gets conn's certificate, as set by
// g_tls_connection_set_certificate().
func (conn *TLSConnection) Certificate() TLSCertificater {
	var _arg0 *C.GTlsConnection  // out
	var _cret *C.GTlsCertificate // in

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))

	_cret = C.g_tls_connection_get_certificate(_arg0)
	runtime.KeepAlive(conn)

	var _tlsCertificate TLSCertificater // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(TLSCertificater)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.TLSCertificater")
			}
			_tlsCertificate = rv
		}
	}

	return _tlsCertificate
}

// ChannelBindingData: query the TLS backend for TLS channel binding data of
// type for conn.
//
// This call retrieves TLS channel binding data as specified in RFC 5056
// (https://tools.ietf.org/html/rfc5056), RFC 5929
// (https://tools.ietf.org/html/rfc5929), and related RFCs. The binding data is
// returned in data. The data is resized by the callee using Array buffer
// management and will be freed when the data is destroyed by
// g_byte_array_unref(). If data is NULL, it will only check whether TLS backend
// is able to fetch the data (e.g. whether type is supported by the TLS
// backend). It does not guarantee that the data will be available though. That
// could happen if TLS connection does not support type or the binding data is
// not available yet due to additional negotiation or input required.
//
// The function takes the following parameters:
//
//    - typ type of data to fetch.
//
func (conn *TLSConnection) ChannelBindingData(typ TLSChannelBindingType) ([]byte, error) {
	var _arg0 *C.GTlsConnection        // out
	var _arg1 C.GTlsChannelBindingType // out
	var _arg2 C.GByteArray             // in
	var _cerr *C.GError                // in

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))
	_arg1 = C.GTlsChannelBindingType(typ)

	C.g_tls_connection_get_channel_binding_data(_arg0, _arg1, &_arg2, &_cerr)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(typ)

	var _data []byte // out
	var _goerr error // out

	_data = make([]byte, _arg2.len)
	copy(_data, unsafe.Slice((*byte)(_arg2.data), _arg2.len))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _data, _goerr
}

// Database gets the certificate database that conn uses to verify peer
// certificates. See g_tls_connection_set_database().
func (conn *TLSConnection) Database() TLSDatabaser {
	var _arg0 *C.GTlsConnection // out
	var _cret *C.GTlsDatabase   // in

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))

	_cret = C.g_tls_connection_get_database(_arg0)
	runtime.KeepAlive(conn)

	var _tlsDatabase TLSDatabaser // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(TLSDatabaser)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.TLSDatabaser")
			}
			_tlsDatabase = rv
		}
	}

	return _tlsDatabase
}

// Interaction: get the object that will be used to interact with the user. It
// will be used for things like prompting the user for passwords. If NULL is
// returned, then no user interaction will occur for this connection.
func (conn *TLSConnection) Interaction() *TLSInteraction {
	var _arg0 *C.GTlsConnection  // out
	var _cret *C.GTlsInteraction // in

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))

	_cret = C.g_tls_connection_get_interaction(_arg0)
	runtime.KeepAlive(conn)

	var _tlsInteraction *TLSInteraction // out

	if _cret != nil {
		_tlsInteraction = wrapTLSInteraction(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _tlsInteraction
}

// NegotiatedProtocol gets the name of the application-layer protocol negotiated
// during the handshake.
//
// If the peer did not use the ALPN extension, or did not advertise a protocol
// that matched one of conn's protocols, or the TLS backend does not support
// ALPN, then this will be NULL. See
// g_tls_connection_set_advertised_protocols().
func (conn *TLSConnection) NegotiatedProtocol() string {
	var _arg0 *C.GTlsConnection // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))

	_cret = C.g_tls_connection_get_negotiated_protocol(_arg0)
	runtime.KeepAlive(conn)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// PeerCertificate gets conn's peer's certificate after the handshake has
// completed or failed. (It is not set during the emission of
// Connection::accept-certificate.).
func (conn *TLSConnection) PeerCertificate() TLSCertificater {
	var _arg0 *C.GTlsConnection  // out
	var _cret *C.GTlsCertificate // in

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))

	_cret = C.g_tls_connection_get_peer_certificate(_arg0)
	runtime.KeepAlive(conn)

	var _tlsCertificate TLSCertificater // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(TLSCertificater)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.TLSCertificater")
			}
			_tlsCertificate = rv
		}
	}

	return _tlsCertificate
}

// PeerCertificateErrors gets the errors associated with validating conn's
// peer's certificate, after the handshake has completed or failed. (It is not
// set during the emission of Connection::accept-certificate.).
func (conn *TLSConnection) PeerCertificateErrors() TLSCertificateFlags {
	var _arg0 *C.GTlsConnection      // out
	var _cret C.GTlsCertificateFlags // in

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))

	_cret = C.g_tls_connection_get_peer_certificate_errors(_arg0)
	runtime.KeepAlive(conn)

	var _tlsCertificateFlags TLSCertificateFlags // out

	_tlsCertificateFlags = TLSCertificateFlags(_cret)

	return _tlsCertificateFlags
}

// RehandshakeMode gets conn rehandshaking mode. See
// g_tls_connection_set_rehandshake_mode() for details.
//
// Deprecated: Changing the rehandshake mode is no longer required for
// compatibility. Also, rehandshaking has been removed from the TLS protocol in
// TLS 1.3.
func (conn *TLSConnection) RehandshakeMode() TLSRehandshakeMode {
	var _arg0 *C.GTlsConnection     // out
	var _cret C.GTlsRehandshakeMode // in

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))

	_cret = C.g_tls_connection_get_rehandshake_mode(_arg0)
	runtime.KeepAlive(conn)

	var _tlsRehandshakeMode TLSRehandshakeMode // out

	_tlsRehandshakeMode = TLSRehandshakeMode(_cret)

	return _tlsRehandshakeMode
}

// RequireCloseNotify tests whether or not conn expects a proper TLS close
// notification when the connection is closed. See
// g_tls_connection_set_require_close_notify() for details.
func (conn *TLSConnection) RequireCloseNotify() bool {
	var _arg0 *C.GTlsConnection // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))

	_cret = C.g_tls_connection_get_require_close_notify(_arg0)
	runtime.KeepAlive(conn)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UseSystemCertDB gets whether conn uses the system certificate database to
// verify peer certificates. See g_tls_connection_set_use_system_certdb().
//
// Deprecated: Use g_tls_connection_get_database() instead.
func (conn *TLSConnection) UseSystemCertDB() bool {
	var _arg0 *C.GTlsConnection // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))

	_cret = C.g_tls_connection_get_use_system_certdb(_arg0)
	runtime.KeepAlive(conn)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Handshake attempts a TLS handshake on conn.
//
// On the client side, it is never necessary to call this method; although the
// connection needs to perform a handshake after connecting (or after sending a
// "STARTTLS"-type command), Connection will handle this for you automatically
// when you try to send or receive data on the connection. You can call
// g_tls_connection_handshake() manually if you want to know whether the initial
// handshake succeeded or failed (as opposed to just immediately trying to use
// conn to read or write, in which case, if it fails, it may not be possible to
// tell if it failed before or after completing the handshake), but beware that
// servers may reject client authentication after the handshake has completed,
// so a successful handshake does not indicate the connection will be usable.
//
// Likewise, on the server side, although a handshake is necessary at the
// beginning of the communication, you do not need to call this function
// explicitly unless you want clearer error reporting.
//
// Previously, calling g_tls_connection_handshake() after the initial handshake
// would trigger a rehandshake; however, this usage was deprecated in GLib 2.60
// because rehandshaking was removed from the TLS protocol in TLS 1.3. Since
// GLib 2.64, calling this function after the initial handshake will no longer
// do anything.
//
// When using a Connection created by Client, the Client performs the initial
// handshake, so calling this function manually is not recommended.
//
// Connection::accept_certificate may be emitted during the handshake.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//
func (conn *TLSConnection) Handshake(ctx context.Context) error {
	var _arg0 *C.GTlsConnection // out
	var _arg1 *C.GCancellable   // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	C.g_tls_connection_handshake(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(ctx)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// HandshakeAsync: asynchronously performs a TLS handshake on conn. See
// g_tls_connection_handshake() for more information.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback to call when the handshake is complete.
//
func (conn *TLSConnection) HandshakeAsync(ctx context.Context, ioPriority int, callback AsyncReadyCallback) {
	var _arg0 *C.GTlsConnection     // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.int                 // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.int(ioPriority)
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_tls_connection_handshake_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// HandshakeFinish: finish an asynchronous TLS handshake operation. See
// g_tls_connection_handshake() for more information.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (conn *TLSConnection) HandshakeFinish(result AsyncResulter) error {
	var _arg0 *C.GTlsConnection // out
	var _arg1 *C.GAsyncResult   // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_tls_connection_handshake_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetAdvertisedProtocols sets the list of application-layer protocols to
// advertise that the caller is willing to speak on this connection. The
// Application-Layer Protocol Negotiation (ALPN) extension will be used to
// negotiate a compatible protocol with the peer; use
// g_tls_connection_get_negotiated_protocol() to find the negotiated protocol
// after the handshake. Specifying NULL for the the value of protocols will
// disable ALPN negotiation.
//
// See IANA TLS ALPN Protocol IDs
// (https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids)
// for a list of registered protocol IDs.
//
// The function takes the following parameters:
//
//    - protocols: NULL-terminated array of ALPN protocol names (eg,
//    "http/1.1", "h2"), or NULL.
//
func (conn *TLSConnection) SetAdvertisedProtocols(protocols []string) {
	var _arg0 *C.GTlsConnection // out
	var _arg1 **C.gchar         // out

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))
	{
		_arg1 = (**C.gchar)(C.calloc(C.size_t((len(protocols) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(protocols)+1)
			var zero *C.gchar
			out[len(protocols)] = zero
			for i := range protocols {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(protocols[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.g_tls_connection_set_advertised_protocols(_arg0, _arg1)
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
// first. You can call g_tls_client_connection_get_accepted_cas() on the failed
// connection to get a list of Certificate Authorities that the server will
// accept certificates from.
//
// (It is also possible that a server will allow the connection with or without
// a certificate; in that case, if you don't provide a certificate, you can tell
// that the server requested one by the fact that
// g_tls_client_connection_get_accepted_cas() will return non-NULL.).
//
// The function takes the following parameters:
//
//    - certificate to use for conn.
//
func (conn *TLSConnection) SetCertificate(certificate TLSCertificater) {
	var _arg0 *C.GTlsConnection  // out
	var _arg1 *C.GTlsCertificate // out

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(certificate.Native()))

	C.g_tls_connection_set_certificate(_arg0, _arg1)
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
//    - database: Database.
//
func (conn *TLSConnection) SetDatabase(database TLSDatabaser) {
	var _arg0 *C.GTlsConnection // out
	var _arg1 *C.GTlsDatabase   // out

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))
	if database != nil {
		_arg1 = (*C.GTlsDatabase)(unsafe.Pointer(database.Native()))
	}

	C.g_tls_connection_set_database(_arg0, _arg1)
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
//    - interaction object, or NULL.
//
func (conn *TLSConnection) SetInteraction(interaction *TLSInteraction) {
	var _arg0 *C.GTlsConnection  // out
	var _arg1 *C.GTlsInteraction // out

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))
	if interaction != nil {
		_arg1 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	}

	C.g_tls_connection_set_interaction(_arg0, _arg1)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(interaction)
}

// SetRehandshakeMode: since GLib 2.64, changing the rehandshake mode is no
// longer supported and will have no effect. With TLS 1.3, rehandshaking has
// been removed from the TLS protocol, replaced by separate post-handshake
// authentication and rekey operations.
//
// Deprecated: Changing the rehandshake mode is no longer required for
// compatibility. Also, rehandshaking has been removed from the TLS protocol in
// TLS 1.3.
//
// The function takes the following parameters:
//
//    - mode: rehandshaking mode.
//
func (conn *TLSConnection) SetRehandshakeMode(mode TLSRehandshakeMode) {
	var _arg0 *C.GTlsConnection     // out
	var _arg1 C.GTlsRehandshakeMode // out

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))
	_arg1 = C.GTlsRehandshakeMode(mode)

	C.g_tls_connection_set_rehandshake_mode(_arg0, _arg1)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(mode)
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
// self-delimiting); in this case, the close notify is redundant and sometimes
// omitted. (TLS 1.1 explicitly allows this; in TLS 1.0 it is technically an
// error, but often done anyway.) You can use
// g_tls_connection_set_require_close_notify() to tell conn to allow an
// "unannounced" connection close, in which case the close will show up as a
// 0-length read, as in a non-TLS Connection, and it is up to the application to
// check that the data has been fully received.
//
// Note that this only affects the behavior when the peer closes the connection;
// when the application calls g_io_stream_close() itself on conn, this will send
// a close notification regardless of the setting of this property. If you
// explicitly want to do an unclean close, you can close conn's
// Connection:base-io-stream rather than closing conn itself, but note that this
// may only be done when no other operations are pending on conn or the base I/O
// stream.
//
// The function takes the following parameters:
//
//    - requireCloseNotify: whether or not to require close notification.
//
func (conn *TLSConnection) SetRequireCloseNotify(requireCloseNotify bool) {
	var _arg0 *C.GTlsConnection // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))
	if requireCloseNotify {
		_arg1 = C.TRUE
	}

	C.g_tls_connection_set_require_close_notify(_arg0, _arg1)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(requireCloseNotify)
}

// SetUseSystemCertDB sets whether conn uses the system certificate database to
// verify peer certificates. This is TRUE by default. If set to FALSE, then peer
// certificate validation will always set the G_TLS_CERTIFICATE_UNKNOWN_CA error
// (meaning Connection::accept-certificate will always be emitted on client-side
// connections, unless that bit is not set in
// ClientConnection:validation-flags).
//
// Deprecated: Use g_tls_connection_set_database() instead.
//
// The function takes the following parameters:
//
//    - useSystemCertdb: whether to use the system certificate database.
//
func (conn *TLSConnection) SetUseSystemCertDB(useSystemCertdb bool) {
	var _arg0 *C.GTlsConnection // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GTlsConnection)(unsafe.Pointer(conn.Native()))
	if useSystemCertdb {
		_arg1 = C.TRUE
	}

	C.g_tls_connection_set_use_system_certdb(_arg0, _arg1)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(useSystemCertdb)
}

func (conn *TLSConnection) baseTLSConnection() *TLSConnection {
	return conn
}

// BaseTLSConnection returns the underlying base object.
func BaseTLSConnection(obj TLSConnectioner) *TLSConnection {
	return obj.baseTLSConnection()
}

// ConnectAcceptCertificate: emitted during the TLS handshake after the peer
// certificate has been received. You can examine peer_cert's certification path
// by calling g_tls_certificate_get_issuer() on it.
//
// For a client-side connection, peer_cert is the server's certificate, and the
// signal will only be emitted if the certificate was not acceptable according
// to conn's ClientConnection:validation_flags. If you would like the
// certificate to be accepted despite errors, return TRUE from the signal
// handler. Otherwise, if no handler accepts the certificate, the handshake will
// fail with G_TLS_ERROR_BAD_CERTIFICATE.
//
// For a server-side connection, peer_cert is the certificate presented by the
// client, if this was requested via the server's
// ServerConnection:authentication_mode. On the server side, the signal is
// always emitted when the client presents a certificate, and the certificate
// will only be accepted if a handler returns TRUE.
//
// Note that if this signal is emitted as part of asynchronous I/O in the main
// thread, then you should not attempt to interact with the user before
// returning from the signal handler. If you want to let the user decide whether
// or not to accept the certificate, you would have to return FALSE from the
// signal handler on the first attempt, and then after the connection attempt
// returns a G_TLS_ERROR_BAD_CERTIFICATE, you can interact with the user, and if
// the user decides to accept the certificate, remember that fact, create a new
// connection, and return TRUE from the signal handler the next time.
//
// If you are doing I/O in another thread, you do not need to worry about this,
// and can simply block in the signal handler until the UI thread returns an
// answer.
func (conn *TLSConnection) ConnectAcceptCertificate(f func(peerCert TLSCertificater, errors TLSCertificateFlags) bool) externglib.SignalHandle {
	return conn.Connect("accept-certificate", f)
}
