// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gtcpconnection.go.
var GTypeTCPConnection = externglib.Type(C.g_tcp_connection_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeTCPConnection, F: marshalTCPConnection},
	})
}

// TCPConnectionOverrider contains methods that are overridable.
type TCPConnectionOverrider interface {
}

// TCPConnection: this is the subclass of Connection that is created for TCP/IP
// sockets.
type TCPConnection struct {
	_ [0]func() // equal guard
	SocketConnection
}

var (
	_ IOStreamer = (*TCPConnection)(nil)
)

func classInitTCPConnectioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTCPConnection(obj *externglib.Object) *TCPConnection {
	return &TCPConnection{
		SocketConnection: SocketConnection{
			IOStream: IOStream{
				Object: obj,
			},
		},
	}
}

func marshalTCPConnection(p uintptr) (interface{}, error) {
	return wrapTCPConnection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// GracefulDisconnect checks if graceful disconnects are used. See
// g_tcp_connection_set_graceful_disconnect().
//
// The function returns the following values:
//
//    - ok: TRUE if graceful disconnect is used on close, FALSE otherwise.
//
func (connection *TCPConnection) GracefulDisconnect() bool {
	var _arg0 *C.GTcpConnection // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GTcpConnection)(unsafe.Pointer(connection.Native()))

	_cret = C.g_tcp_connection_get_graceful_disconnect(_arg0)
	runtime.KeepAlive(connection)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetGracefulDisconnect: this enables graceful disconnects on close. A graceful
// disconnect means that we signal the receiving end that the connection is
// terminated and wait for it to close the connection before closing the
// connection.
//
// A graceful disconnect means that we can be sure that we successfully sent all
// the outstanding data to the other end, or get an error reported. However, it
// also means we have to wait for all the data to reach the other side and for
// it to acknowledge this by closing the socket, which may take a while. For
// this reason it is disabled by default.
//
// The function takes the following parameters:
//
//    - gracefulDisconnect: whether to do graceful disconnects or not.
//
func (connection *TCPConnection) SetGracefulDisconnect(gracefulDisconnect bool) {
	var _arg0 *C.GTcpConnection // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GTcpConnection)(unsafe.Pointer(connection.Native()))
	if gracefulDisconnect {
		_arg1 = C.TRUE
	}

	C.g_tcp_connection_set_graceful_disconnect(_arg0, _arg1)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(gracefulDisconnect)
}
