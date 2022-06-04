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
// extern gboolean _gotk4_gio2_TlsConnectionClass_handshake(void*, void*, GError**);
// extern gboolean _gotk4_gio2_TlsConnectionClass_handshake_finish(void*, void*, GError**);
// extern void _gotk4_gio2_AsyncReadyCallback(void*, void*, gpointer);
import "C"

// glib.Type values for gtlsconnection.go.
var GTypeTLSConnection = coreglib.Type(C.g_tls_connection_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeTLSConnection, F: marshalTLSConnection},
	})
}

// TLSConnectionOverrider contains methods that are overridable.
type TLSConnectionOverrider interface {
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
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional) or NULL.
	//
	Handshake(ctx context.Context) error
	// HandshakeFinish: finish an asynchronous TLS handshake operation. See
	// g_tls_connection_handshake() for more information.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	HandshakeFinish(result AsyncResulter) error
}

// TLSConnection is the base TLS connection class type, which wraps a OStream
// and provides TLS encryption on top of it. Its subclasses, ClientConnection
// and ServerConnection, implement client-side and server-side TLS,
// respectively.
//
// For DTLS (Datagram TLS) support, see Connection.
type TLSConnection struct {
	_ [0]func() // equal guard
	IOStream
}

var (
	_ IOStreamer = (*TLSConnection)(nil)
)

// TLSConnectioner describes types inherited from class TLSConnection.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type TLSConnectioner interface {
	coreglib.Objector
	baseTLSConnection() *TLSConnection
}

var _ TLSConnectioner = (*TLSConnection)(nil)

func classInitTLSConnectioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GTlsConnectionClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GTlsConnectionClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface {
		Handshake(ctx context.Context) error
	}); ok {
		pclass.handshake = (*[0]byte)(C._gotk4_gio2_TlsConnectionClass_handshake)
	}

	if _, ok := goval.(interface {
		HandshakeFinish(result AsyncResulter) error
	}); ok {
		pclass.handshake_finish = (*[0]byte)(C._gotk4_gio2_TlsConnectionClass_handshake_finish)
	}
}

//export _gotk4_gio2_TlsConnectionClass_handshake
func _gotk4_gio2_TlsConnectionClass_handshake(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Handshake(ctx context.Context) error
	})

	var _cancellable context.Context // out

	if arg1 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg1))
	}

	_goerr := iface.Handshake(_cancellable)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_TlsConnectionClass_handshake_finish
func _gotk4_gio2_TlsConnectionClass_handshake_finish(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		HandshakeFinish(result AsyncResulter) error
	})

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	_goerr := iface.HandshakeFinish(_result)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

func wrapTLSConnection(obj *coreglib.Object) *TLSConnection {
	return &TLSConnection{
		IOStream: IOStream{
			Object: obj,
		},
	}
}

func marshalTLSConnection(p uintptr) (interface{}, error) {
	return wrapTLSConnection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (conn *TLSConnection) baseTLSConnection() *TLSConnection {
	return conn
}

// BaseTLSConnection returns the underlying base object.
func BaseTLSConnection(obj TLSConnectioner) *TLSConnection {
	return obj.baseTLSConnection()
}

// Certificate gets conn's certificate, as set by
// g_tls_connection_set_certificate().
//
// The function returns the following values:
//
//    - tlsCertificate (optional) conn's certificate, or NULL.
//
func (conn *TLSConnection) Certificate() TLSCertificater {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_gret := girepository.MustFind("Gio", "TlsConnection").InvokeMethod("get_certificate", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(conn)

	var _tlsCertificate TLSCertificater // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

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
// certificates. See g_tls_connection_set_database().
//
// The function returns the following values:
//
//    - tlsDatabase (optional): certificate database that conn uses or NULL.
//
func (conn *TLSConnection) Database() TLSDatabaser {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_gret := girepository.MustFind("Gio", "TlsConnection").InvokeMethod("get_database", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(conn)

	var _tlsDatabase TLSDatabaser // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

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
func (conn *TLSConnection) Interaction() *TLSInteraction {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_gret := girepository.MustFind("Gio", "TlsConnection").InvokeMethod("get_interaction", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(conn)

	var _tlsInteraction *TLSInteraction // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_tlsInteraction = wrapTLSInteraction(coreglib.Take(unsafe.Pointer(_cret)))
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
//
// The function returns the following values:
//
//    - utf8 (optional): negotiated protocol, or NULL.
//
func (conn *TLSConnection) NegotiatedProtocol() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_gret := girepository.MustFind("Gio", "TlsConnection").InvokeMethod("get_negotiated_protocol", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(conn)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
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
func (conn *TLSConnection) PeerCertificate() TLSCertificater {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_gret := girepository.MustFind("Gio", "TlsConnection").InvokeMethod("get_peer_certificate", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(conn)

	var _tlsCertificate TLSCertificater // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

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
// g_tls_connection_set_require_close_notify() for details.
//
// The function returns the following values:
//
//    - ok: TRUE if conn requires a proper TLS close notification.
//
func (conn *TLSConnection) RequireCloseNotify() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_gret := girepository.MustFind("Gio", "TlsConnection").InvokeMethod("get_require_close_notify", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(conn)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// UseSystemCertDB gets whether conn uses the system certificate database to
// verify peer certificates. See g_tls_connection_set_use_system_certdb().
//
// Deprecated: Use g_tls_connection_get_database() instead.
//
// The function returns the following values:
//
//    - ok: whether conn uses the system certificate database.
//
func (conn *TLSConnection) UseSystemCertDB() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_gret := girepository.MustFind("Gio", "TlsConnection").InvokeMethod("get_use_system_certdb", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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
//    - ctx (optional) or NULL.
//
func (conn *TLSConnection) Handshake(ctx context.Context) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[1] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}

	girepository.MustFind("Gio", "TlsConnection").InvokeMethod("handshake", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(ctx)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// HandshakeAsync: asynchronously performs a TLS handshake on conn. See
// g_tls_connection_handshake() for more information.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback (optional) to call when the handshake is complete.
//
func (conn *TLSConnection) HandshakeAsync(ctx context.Context, ioPriority int32, callback AsyncReadyCallback) {
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

	girepository.MustFind("Gio", "TlsConnection").InvokeMethod("handshake_async", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	girepository.MustFind("Gio", "TlsConnection").InvokeMethod("handshake_finish", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(result)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
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
//    - protocols (optional): NULL-terminated array of ALPN protocol names (eg,
//      "http/1.1", "h2"), or NULL.
//
func (conn *TLSConnection) SetAdvertisedProtocols(protocols []string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	{
		*(***C.void)(unsafe.Pointer(&_args[1])) = (**C.void)(C.calloc(C.size_t((len(protocols) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_args[1]))
		{
			out := unsafe.Slice(_args[1], len(protocols)+1)
			var zero *C.void
			out[len(protocols)] = zero
			for i := range protocols {
				*(**C.void)(unsafe.Pointer(&out[i])) = (*C.void)(unsafe.Pointer(C.CString(protocols[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	girepository.MustFind("Gio", "TlsConnection").InvokeMethod("set_advertised_protocols", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(certificate).Native()))

	girepository.MustFind("Gio", "TlsConnection").InvokeMethod("set_certificate", _args[:], nil)

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
func (conn *TLSConnection) SetDatabase(database TLSDatabaser) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	if database != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(database).Native()))
	}

	girepository.MustFind("Gio", "TlsConnection").InvokeMethod("set_database", _args[:], nil)

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
func (conn *TLSConnection) SetInteraction(interaction *TLSInteraction) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	if interaction != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(interaction).Native()))
	}

	girepository.MustFind("Gio", "TlsConnection").InvokeMethod("set_interaction", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	if requireCloseNotify {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gio", "TlsConnection").InvokeMethod("set_require_close_notify", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	if useSystemCertdb {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gio", "TlsConnection").InvokeMethod("set_use_system_certdb", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(useSystemCertdb)
}
