// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tcp_wrapper_connection_get_type()), F: marshalTCPWrapperConnectioner},
	})
}

// TCPWrapperConnectioner describes TCPWrapperConnection's methods.
type TCPWrapperConnectioner interface {
	// BaseIOStream gets @conn's base OStream
	BaseIOStream() *IOStream
}

// TCPWrapperConnection can be used to wrap a OStream that is based on a
// #GSocket, but which is not actually a Connection. This is used by Client so
// that it can always return a Connection, even when the connection it has
// actually created is not directly a Connection.
type TCPWrapperConnection struct {
	TCPConnection
}

var (
	_ TCPWrapperConnectioner = (*TCPWrapperConnection)(nil)
	_ gextras.Nativer        = (*TCPWrapperConnection)(nil)
)

func wrapTCPWrapperConnection(obj *externglib.Object) *TCPWrapperConnection {
	return &TCPWrapperConnection{
		TCPConnection: TCPConnection{
			SocketConnection: SocketConnection{
				IOStream: IOStream{
					Object: obj,
				},
			},
		},
	}
}

func marshalTCPWrapperConnectioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTCPWrapperConnection(obj), nil
}

// NewTCPWrapperConnection wraps @base_io_stream and @socket together as a
// Connection.
func NewTCPWrapperConnection(baseIoStream IOStreamer, socket Socketer) *TCPWrapperConnection {
	var _arg1 *C.GIOStream         // out
	var _arg2 *C.GSocket           // out
	var _cret *C.GSocketConnection // in

	_arg1 = (*C.GIOStream)(unsafe.Pointer((baseIoStream).(gextras.Nativer).Native()))
	_arg2 = (*C.GSocket)(unsafe.Pointer((socket).(gextras.Nativer).Native()))

	_cret = C.g_tcp_wrapper_connection_new(_arg1, _arg2)

	var _tcpWrapperConnection *TCPWrapperConnection // out

	_tcpWrapperConnection = wrapTCPWrapperConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _tcpWrapperConnection
}

// BaseIOStream gets @conn's base OStream
func (conn *TCPWrapperConnection) BaseIOStream() *IOStream {
	var _arg0 *C.GTcpWrapperConnection // out
	var _cret *C.GIOStream             // in

	_arg0 = (*C.GTcpWrapperConnection)(unsafe.Pointer(conn.Native()))

	_cret = C.g_tcp_wrapper_connection_get_base_io_stream(_arg0)

	var _ioStream *IOStream // out

	_ioStream = wrapIOStream(externglib.Take(unsafe.Pointer(_cret)))

	return _ioStream
}
