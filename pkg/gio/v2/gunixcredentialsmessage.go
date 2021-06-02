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
		{T: externglib.Type(C.g_unix_credentials_message_get_type()), F: marshalUnixCredentialsMessage},
	})
}

type UnixCredentialsMessagePrivate struct {
	native C.GUnixCredentialsMessagePrivate
}

// WrapUnixCredentialsMessagePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapUnixCredentialsMessagePrivate(ptr unsafe.Pointer) *UnixCredentialsMessagePrivate {
	if ptr == nil {
		return nil
	}

	return (*UnixCredentialsMessagePrivate)(ptr)
}

func marshalUnixCredentialsMessagePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapUnixCredentialsMessagePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (u *UnixCredentialsMessagePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&u.native)
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

	// Credentials gets the credentials stored in @message.
	Credentials() Credentials
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

// NewUnixCredentialsMessage constructs a class UnixCredentialsMessage.
func NewUnixCredentialsMessage() UnixCredentialsMessage {
	ret := C.g_unix_credentials_message_new()

	var ret0 UnixCredentialsMessage

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(UnixCredentialsMessage)

	return ret0
}

// NewUnixCredentialsMessageWithCredentials constructs a class UnixCredentialsMessage.
func NewUnixCredentialsMessageWithCredentials(credentials Credentials) UnixCredentialsMessage {
	var arg1 *C.GCredentials

	arg1 = (*C.GCredentials)(credentials.Native())

	ret := C.g_unix_credentials_message_new_with_credentials(arg1)

	var ret0 UnixCredentialsMessage

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(UnixCredentialsMessage)

	return ret0
}

// Credentials gets the credentials stored in @message.
func (m unixCredentialsMessage) Credentials() Credentials {
	var arg0 *C.GUnixCredentialsMessage

	arg0 = (*C.GUnixCredentialsMessage)(m.Native())

	ret := C.g_unix_credentials_message_get_credentials(arg0)

	var ret0 Credentials

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Credentials)

	return ret0
}
