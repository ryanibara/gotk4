// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tls_server_connection_get_type()), F: marshalTLSServerConnection},
	})
}

// NewTLSServerConnection creates a new ServerConnection wrapping
// @base_io_stream (which must have pollable input and output streams).
//
// See the documentation for Connection:base-io-stream for restrictions on when
// application code can run operations on the @base_io_stream after this
// function has returned.
func NewTLSServerConnection(baseIOStream IOStream, certificate TLSCertificate) (tlsServerConnection TLSServerConnection, err error) {
	var arg1 *C.GIOStream
	var arg2 *C.GTlsCertificate
	var gError *C.GError

	arg1 = (*C.GIOStream)(baseIOStream.Native())
	arg2 = (*C.GTlsCertificate)(certificate.Native())

	ret := C.g_tls_server_connection_new(arg1, arg2, &gError)

	var ret0 TLSServerConnection
	var goError error

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(TLSServerConnection)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// TLSServerConnection is the server-side subclass of Connection, representing a
// server-side TLS connection.
type TLSServerConnection interface {
	TLSConnection
}

// tlsServerConnection implements the TLSServerConnection interface.
type tlsServerConnection struct {
	TLSConnection
}

var _ TLSServerConnection = (*tlsServerConnection)(nil)

// WrapTLSServerConnection wraps a GObject to a type that implements interface
// TLSServerConnection. It is primarily used internally.
func WrapTLSServerConnection(obj *externglib.Object) TLSServerConnection {
	return TLSServerConnection{
		TLSConnection: WrapTLSConnection(obj),
	}
}

func marshalTLSServerConnection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTLSServerConnection(obj), nil
}
