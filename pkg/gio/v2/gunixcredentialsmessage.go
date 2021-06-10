// Code generated by girgen. DO NOT EDIT.

package gio

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
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
		{T: externglib.Type(C.g_unix_credentials_message_get_type()), F: marshalUnixCredentialsMessage},
	})
}

// UnixCredentialsMessage: this ControlMessage contains a #GCredentials
// instance. It may be sent using g_socket_send_message() and received using
// g_socket_receive_message() over UNIX sockets (ie: sockets in the
// G_SOCKET_FAMILY_UNIX family).
//
// For an easier way to send and receive credentials over stream-oriented UNIX
// sockets, see g_unix_connection_send_credentials() and
// g_unix_connection_receive_credentials(). To receive credentials of a foreign
// process connected to a socket, use g_socket_get_credentials().
type UnixCredentialsMessage interface {
	SocketControlMessage
}

// unixCredentialsMessage implements the UnixCredentialsMessage interface.
type unixCredentialsMessage struct {
	SocketControlMessage
}

var _ UnixCredentialsMessage = (*unixCredentialsMessage)(nil)

// WrapUnixCredentialsMessage wraps a GObject to the right type. It is
// primarily used internally.
func WrapUnixCredentialsMessage(obj *externglib.Object) UnixCredentialsMessage {
	return UnixCredentialsMessage{
		SocketControlMessage: WrapSocketControlMessage(obj),
	}
}

func marshalUnixCredentialsMessage(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapUnixCredentialsMessage(obj), nil
}
