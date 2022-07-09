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
// #include <glib-object.h>
import "C"

// GTypeDTLSClientConnection returns the GType for the type DTLSClientConnection.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDTLSClientConnection() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "DtlsClientConnection").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDTLSClientConnection)
	return gtype
}

// DTLSClientConnectionOverrider contains methods that are overridable.
type DTLSClientConnectionOverrider interface {
}

// DTLSClientConnection is the client-side subclass of Connection, representing
// a client-side DTLS connection.
//
// DTLSClientConnection wraps an interface. This means the user can get the
// underlying type by calling Cast().
type DTLSClientConnection struct {
	_ [0]func() // equal guard
	DTLSConnection
}

var ()

// DTLSClientConnectioner describes DTLSClientConnection's interface methods.
type DTLSClientConnectioner interface {
	coreglib.Objector

	// ServerIdentity gets conn's expected server identity.
	ServerIdentity() *SocketConnectable
	// SetServerIdentity sets conn's expected server identity, which is used
	// both to tell servers on virtual hosts which certificate to present, and
	// also to let conn know what name to look for in the certificate when
	// performing G_TLS_CERTIFICATE_BAD_IDENTITY validation, if enabled.
	SetServerIdentity(identity SocketConnectabler)
}

var _ DTLSClientConnectioner = (*DTLSClientConnection)(nil)

func ifaceInitDTLSClientConnectioner(gifacePtr, data C.gpointer) {
}

func wrapDTLSClientConnection(obj *coreglib.Object) *DTLSClientConnection {
	return &DTLSClientConnection{
		DTLSConnection: DTLSConnection{
			DatagramBased: DatagramBased{
				Object: obj,
			},
		},
	}
}

func marshalDTLSClientConnection(p uintptr) (interface{}, error) {
	return wrapDTLSClientConnection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ServerIdentity gets conn's expected server identity.
//
// The function returns the following values:
//
//    - socketConnectable describing the expected server identity, or NULL if the
//      expected identity is not known.
//
func (conn *DTLSClientConnection) ServerIdentity() *SocketConnectable {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_info := girepository.MustFind("Gio", "DtlsClientConnection")
	_gret := _info.InvokeIfaceMethod("get_server_identity", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(conn)

	var _socketConnectable *SocketConnectable // out

	_socketConnectable = wrapSocketConnectable(coreglib.Take(unsafe.Pointer(_cret)))

	return _socketConnectable
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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(identity).Native()))

	_info := girepository.MustFind("Gio", "DtlsClientConnection")
	_info.InvokeIfaceMethod("set_server_identity", _args[:], nil)

	runtime.KeepAlive(conn)
	runtime.KeepAlive(identity)
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
func NewDTLSClientConnection(baseSocket DatagramBasedder, serverIdentity SocketConnectabler) (*DTLSClientConnection, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(baseSocket).Native()))
	if serverIdentity != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(serverIdentity).Native()))
	}

	_info := girepository.MustFind("Gio", "new")
	_gret := _info.Invoke(_args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(baseSocket)
	runtime.KeepAlive(serverIdentity)

	var _dtlsClientConnection *DTLSClientConnection // out
	var _goerr error                                // out

	_dtlsClientConnection = wrapDTLSClientConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dtlsClientConnection, _goerr
}
