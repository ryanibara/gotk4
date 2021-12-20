// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_dtls_client_connection_get_type()), F: marshalDTLSClientConnectioner},
	})
}

// DTLSClientConnection is the client-side subclass of Connection, representing
// a client-side DTLS connection.
type DTLSClientConnection struct {
	_ [0]func() // equal guard
	DTLSConnection
}

var ()

// DTLSClientConnectioner describes DTLSClientConnection's interface methods.
type DTLSClientConnectioner interface {
	externglib.Objector

	// ServerIdentity gets conn's expected server identity.
	ServerIdentity() SocketConnectabler
	// ValidationFlags gets conn's validation flags.
	ValidationFlags() TLSCertificateFlags
	// SetServerIdentity sets conn's expected server identity, which is used
	// both to tell servers on virtual hosts which certificate to present, and
	// also to let conn know what name to look for in the certificate when
	// performing G_TLS_CERTIFICATE_BAD_IDENTITY validation, if enabled.
	SetServerIdentity(identity SocketConnectabler)
	// SetValidationFlags sets conn's validation flags, to override the default
	// set of checks performed when validating a server certificate.
	SetValidationFlags(flags TLSCertificateFlags)
}

var _ DTLSClientConnectioner = (*DTLSClientConnection)(nil)

func wrapDTLSClientConnection(obj *externglib.Object) *DTLSClientConnection {
	return &DTLSClientConnection{
		DTLSConnection: DTLSConnection{
			DatagramBased: DatagramBased{
				Object: obj,
			},
		},
	}
}

func marshalDTLSClientConnectioner(p uintptr) (interface{}, error) {
	return wrapDTLSClientConnection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ServerIdentity gets conn's expected server identity.
//
// The function returns the following values:
//
//    - socketConnectable describing the expected server identity, or NULL if the
//      expected identity is not known.
//
func (conn *DTLSClientConnection) ServerIdentity() SocketConnectabler {
	var _arg0 *C.GDtlsClientConnection // out
	var _cret *C.GSocketConnectable    // in

	_arg0 = (*C.GDtlsClientConnection)(unsafe.Pointer(conn.Native()))

	_cret = C.g_dtls_client_connection_get_server_identity(_arg0)
	runtime.KeepAlive(conn)

	var _socketConnectable SocketConnectabler // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.SocketConnectabler is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(SocketConnectabler)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.SocketConnectabler")
		}
		_socketConnectable = rv
	}

	return _socketConnectable
}

// ValidationFlags gets conn's validation flags.
//
// The function returns the following values:
//
//    - tlsCertificateFlags: validation flags.
//
func (conn *DTLSClientConnection) ValidationFlags() TLSCertificateFlags {
	var _arg0 *C.GDtlsClientConnection // out
	var _cret C.GTlsCertificateFlags   // in

	_arg0 = (*C.GDtlsClientConnection)(unsafe.Pointer(conn.Native()))

	_cret = C.g_dtls_client_connection_get_validation_flags(_arg0)
	runtime.KeepAlive(conn)

	var _tlsCertificateFlags TLSCertificateFlags // out

	_tlsCertificateFlags = TLSCertificateFlags(_cret)

	return _tlsCertificateFlags
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
func (conn *DTLSClientConnection) SetServerIdentity(identity SocketConnectabler) {
	var _arg0 *C.GDtlsClientConnection // out
	var _arg1 *C.GSocketConnectable    // out

	_arg0 = (*C.GDtlsClientConnection)(unsafe.Pointer(conn.Native()))
	_arg1 = (*C.GSocketConnectable)(unsafe.Pointer(identity.Native()))

	C.g_dtls_client_connection_set_server_identity(_arg0, _arg1)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(identity)
}

// SetValidationFlags sets conn's validation flags, to override the default set
// of checks performed when validating a server certificate. By default,
// G_TLS_CERTIFICATE_VALIDATE_ALL is used.
//
// The function takes the following parameters:
//
//    - flags to use.
//
func (conn *DTLSClientConnection) SetValidationFlags(flags TLSCertificateFlags) {
	var _arg0 *C.GDtlsClientConnection // out
	var _arg1 C.GTlsCertificateFlags   // out

	_arg0 = (*C.GDtlsClientConnection)(unsafe.Pointer(conn.Native()))
	_arg1 = C.GTlsCertificateFlags(flags)

	C.g_dtls_client_connection_set_validation_flags(_arg0, _arg1)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(flags)
}

// NewDTLSClientConnection creates a new ClientConnection wrapping base_socket
// which is assumed to communicate with the server identified by
// server_identity.
//
// The function takes the following parameters:
//
//    - baseSocket to wrap.
//    - serverIdentity (optional): expected identity of the server.
//
// The function returns the following values:
//
//    - dtlsClientConnection: new ClientConnection, or NULL on error.
//
func NewDTLSClientConnection(baseSocket DatagramBasedder, serverIdentity SocketConnectabler) (DTLSClientConnectioner, error) {
	var _arg1 *C.GDatagramBased     // out
	var _arg2 *C.GSocketConnectable // out
	var _cret *C.GDatagramBased     // in
	var _cerr *C.GError             // in

	_arg1 = (*C.GDatagramBased)(unsafe.Pointer(baseSocket.Native()))
	if serverIdentity != nil {
		_arg2 = (*C.GSocketConnectable)(unsafe.Pointer(serverIdentity.Native()))
	}

	_cret = C.g_dtls_client_connection_new(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(baseSocket)
	runtime.KeepAlive(serverIdentity)

	var _dtlsClientConnection DTLSClientConnectioner // out
	var _goerr error                                 // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.DTLSClientConnectioner is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.Cast()
		rv, ok := casted.(DTLSClientConnectioner)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.DTLSClientConnectioner")
		}
		_dtlsClientConnection = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dtlsClientConnection, _goerr
}
