// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
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
		{T: externglib.Type(C.g_dtls_server_connection_get_type()), F: marshalDTLSServerConnectioner},
	})
}

// DTLSServerConnection is the server-side subclass of Connection, representing
// a server-side DTLS connection.
type DTLSServerConnection struct {
	DTLSConnection
}

var ()

// DTLSServerConnectioner describes DTLSServerConnection's interface methods.
type DTLSServerConnectioner interface {
	externglib.Objector

	baseDTLSServerConnection() *DTLSServerConnection
}

var _ DTLSServerConnectioner = (*DTLSServerConnection)(nil)

func wrapDTLSServerConnection(obj *externglib.Object) *DTLSServerConnection {
	return &DTLSServerConnection{
		DTLSConnection: DTLSConnection{
			DatagramBased: DatagramBased{
				Object: obj,
			},
		},
	}
}

func marshalDTLSServerConnectioner(p uintptr) (interface{}, error) {
	return wrapDTLSServerConnection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *DTLSServerConnection) baseDTLSServerConnection() *DTLSServerConnection {
	return v
}

// BaseDTLSServerConnection returns the underlying base object.
func BaseDTLSServerConnection(obj DTLSServerConnectioner) *DTLSServerConnection {
	return obj.baseDTLSServerConnection()
}

// NewDTLSServerConnection creates a new ServerConnection wrapping base_socket.
//
// The function takes the following parameters:
//
//    - baseSocket to wrap.
//    - certificate: default server certificate, or NULL.
//
func NewDTLSServerConnection(baseSocket DatagramBasedder, certificate TLSCertificater) (DTLSServerConnectioner, error) {
	var _arg1 *C.GDatagramBased  // out
	var _arg2 *C.GTlsCertificate // out
	var _cret *C.GDatagramBased  // in
	var _cerr *C.GError          // in

	_arg1 = (*C.GDatagramBased)(unsafe.Pointer(baseSocket.Native()))
	if certificate != nil {
		_arg2 = (*C.GTlsCertificate)(unsafe.Pointer(certificate.Native()))
	}

	_cret = C.g_dtls_server_connection_new(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(baseSocket)
	runtime.KeepAlive(certificate)

	var _dtlsServerConnection DTLSServerConnectioner // out
	var _goerr error                                 // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.DTLSServerConnectioner is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(DTLSServerConnectioner)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.DTLSServerConnectioner")
		}
		_dtlsServerConnection = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dtlsServerConnection, _goerr
}
