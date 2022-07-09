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

// GTypeDTLSServerConnection returns the GType for the type DTLSServerConnection.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDTLSServerConnection() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "DtlsServerConnection").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDTLSServerConnection)
	return gtype
}

// DTLSServerConnectionOverrider contains methods that are overridable.
type DTLSServerConnectionOverrider interface {
}

// DTLSServerConnection is the server-side subclass of Connection, representing
// a server-side DTLS connection.
//
// DTLSServerConnection wraps an interface. This means the user can get the
// underlying type by calling Cast().
type DTLSServerConnection struct {
	_ [0]func() // equal guard
	DTLSConnection
}

var ()

// DTLSServerConnectioner describes DTLSServerConnection's interface methods.
type DTLSServerConnectioner interface {
	coreglib.Objector

	baseDTLSServerConnection() *DTLSServerConnection
}

var _ DTLSServerConnectioner = (*DTLSServerConnection)(nil)

func ifaceInitDTLSServerConnectioner(gifacePtr, data C.gpointer) {
}

func wrapDTLSServerConnection(obj *coreglib.Object) *DTLSServerConnection {
	return &DTLSServerConnection{
		DTLSConnection: DTLSConnection{
			DatagramBased: DatagramBased{
				Object: obj,
			},
		},
	}
}

func marshalDTLSServerConnection(p uintptr) (interface{}, error) {
	return wrapDTLSServerConnection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
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
//    - certificate (optional): default server certificate, or NULL.
//
// The function returns the following values:
//
//    - dtlsServerConnection: new ServerConnection, or NULL on error.
//
func NewDTLSServerConnection(baseSocket DatagramBasedder, certificate TLSCertificater) (*DTLSServerConnection, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(baseSocket).Native()))
	if certificate != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(certificate).Native()))
	}

	_gret := girepository.MustFind("Gio", "new").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(baseSocket)
	runtime.KeepAlive(certificate)

	var _dtlsServerConnection *DTLSServerConnection // out
	var _goerr error                                // out

	_dtlsServerConnection = wrapDTLSServerConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dtlsServerConnection, _goerr
}
