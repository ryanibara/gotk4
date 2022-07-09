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
import "C"

// GTypeTLSServerConnection returns the GType for the type TLSServerConnection.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTLSServerConnection() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "TlsServerConnection").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalTLSServerConnection)
	return gtype
}

// TLSServerConnectionOverrider contains methods that are overridable.
type TLSServerConnectionOverrider interface {
}

// TLSServerConnection is the server-side subclass of Connection, representing a
// server-side TLS connection.
//
// TLSServerConnection wraps an interface. This means the user can get the
// underlying type by calling Cast().
type TLSServerConnection struct {
	_ [0]func() // equal guard
	TLSConnection
}

var (
	_ TLSConnectioner = (*TLSServerConnection)(nil)
)

// TLSServerConnectioner describes TLSServerConnection's interface methods.
type TLSServerConnectioner interface {
	coreglib.Objector

	baseTLSServerConnection() *TLSServerConnection
}

var _ TLSServerConnectioner = (*TLSServerConnection)(nil)

func ifaceInitTLSServerConnectioner(gifacePtr, data C.gpointer) {
}

func wrapTLSServerConnection(obj *coreglib.Object) *TLSServerConnection {
	return &TLSServerConnection{
		TLSConnection: TLSConnection{
			IOStream: IOStream{
				Object: obj,
			},
		},
	}
}

func marshalTLSServerConnection(p uintptr) (interface{}, error) {
	return wrapTLSServerConnection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *TLSServerConnection) baseTLSServerConnection() *TLSServerConnection {
	return v
}

// BaseTLSServerConnection returns the underlying base object.
func BaseTLSServerConnection(obj TLSServerConnectioner) *TLSServerConnection {
	return obj.baseTLSServerConnection()
}

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
//    - certificate (optional): default server certificate, or NULL.
//
// The function returns the following values:
//
//    - tlsServerConnection: new ServerConnection, or NULL on error.
//
func NewTLSServerConnection(baseIoStream IOStreamer, certificate TLSCertificater) (*TLSServerConnection, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(baseIoStream).Native()))
	if certificate != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(certificate).Native()))
	}

	_gret := girepository.MustFind("Gio", "new").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(baseIoStream)
	runtime.KeepAlive(certificate)

	var _tlsServerConnection *TLSServerConnection // out
	var _goerr error                              // out

	_tlsServerConnection = wrapTLSServerConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsServerConnection, _goerr
}
