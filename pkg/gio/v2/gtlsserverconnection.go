// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tls_server_connection_get_type()), F: marshalTLSServerConnectioner},
	})
}

// TLSServerConnection is the server-side subclass of Connection, representing a
// server-side TLS connection.
type TLSServerConnection struct {
	TLSConnection
}

// TLSServerConnectioner describes TLSServerConnection's abstract methods.
type TLSServerConnectioner interface {
	externglib.Objector

	privateTLSServerConnection()
}

var _ TLSServerConnectioner = (*TLSServerConnection)(nil)

func wrapTLSServerConnection(obj *externglib.Object) *TLSServerConnection {
	return &TLSServerConnection{
		TLSConnection: TLSConnection{
			IOStream: IOStream{
				Object: obj,
			},
		},
	}
}

func marshalTLSServerConnectioner(p uintptr) (interface{}, error) {
	return wrapTLSServerConnection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (*TLSServerConnection) privateTLSServerConnection() {}

// NewTLSServerConnection creates a new ServerConnection wrapping base_io_stream
// (which must have pollable input and output streams).
//
// See the documentation for Connection:base-io-stream for restrictions on when
// application code can run operations on the base_io_stream after this function
// has returned.
//
// The function takes the following parameters:
//
//    - baseIoStream to wrap.
//    - certificate: default server certificate, or NULL.
//
func NewTLSServerConnection(baseIoStream IOStreamer, certificate TLSCertificater) (TLSServerConnectioner, error) {
	var _arg1 *C.GIOStream       // out
	var _arg2 *C.GTlsCertificate // out
	var _cret *C.GIOStream       // in
	var _cerr *C.GError          // in

	_arg1 = (*C.GIOStream)(unsafe.Pointer(baseIoStream.Native()))
	if certificate != nil {
		_arg2 = (*C.GTlsCertificate)(unsafe.Pointer(certificate.Native()))
	}

	_cret = C.g_tls_server_connection_new(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(baseIoStream)
	runtime.KeepAlive(certificate)

	var _tlsServerConnection TLSServerConnectioner // out
	var _goerr error                               // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.TLSServerConnectioner is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(TLSServerConnectioner)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.TLSServerConnectioner")
		}
		_tlsServerConnection = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsServerConnection, _goerr
}
