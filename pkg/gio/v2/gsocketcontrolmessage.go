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
		{T: externglib.Type(C.g_socket_control_message_get_type()), F: marshalSocketControlMessage},
	})
}

// SocketControlMessage: a ControlMessage is a special-purpose utility message
// that can be sent to or received from a #GSocket. These types of messages are
// often called "ancillary data".
//
// The message can represent some sort of special instruction to or information
// from the socket or can represent a special kind of transfer to the peer (for
// example, sending a file descriptor over a UNIX socket).
//
// These messages are sent with g_socket_send_message() and received with
// g_socket_receive_message().
//
// To extend the set of control message that can be sent, subclass this class
// and override the get_size, get_level, get_type and serialize methods.
//
// To extend the set of control messages that can be received, subclass this
// class and implement the deserialize method. Also, make sure your class is
// registered with the GType typesystem before calling
// g_socket_receive_message() to read such a message.
type SocketControlMessage interface {
	gextras.Objector

	// Level returns the "level" (i.e. the originating protocol) of the control
	// message. This is often SOL_SOCKET.
	Level() int
	// MsgType returns the protocol specific type of the control message. For
	// instance, for UNIX fd passing this would be SCM_RIGHTS.
	MsgType() int
	// Size returns the space required for the control message, not including
	// headers or alignment.
	Size() uint
	// Serialize converts the data in the message to bytes placed in the
	// message.
	//
	// @data is guaranteed to have enough space to fit the size returned by
	// g_socket_control_message_get_size() on this object.
	Serialize(data interface{})
}

// socketControlMessage implements the SocketControlMessage interface.
type socketControlMessage struct {
	gextras.Objector
}

var _ SocketControlMessage = (*socketControlMessage)(nil)

// WrapSocketControlMessage wraps a GObject to the right type. It is
// primarily used internally.
func WrapSocketControlMessage(obj *externglib.Object) SocketControlMessage {
	return SocketControlMessage{
		Objector: obj,
	}
}

func marshalSocketControlMessage(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSocketControlMessage(obj), nil
}

// Level returns the "level" (i.e. the originating protocol) of the control
// message. This is often SOL_SOCKET.
func (m socketControlMessage) Level() int {
	var _arg0 *C.GSocketControlMessage

	_arg0 = (*C.GSocketControlMessage)(unsafe.Pointer(m.Native()))

	var _cret C.int

	_cret = C.g_socket_control_message_get_level(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// MsgType returns the protocol specific type of the control message. For
// instance, for UNIX fd passing this would be SCM_RIGHTS.
func (m socketControlMessage) MsgType() int {
	var _arg0 *C.GSocketControlMessage

	_arg0 = (*C.GSocketControlMessage)(unsafe.Pointer(m.Native()))

	var _cret C.int

	_cret = C.g_socket_control_message_get_msg_type(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Size returns the space required for the control message, not including
// headers or alignment.
func (m socketControlMessage) Size() uint {
	var _arg0 *C.GSocketControlMessage

	_arg0 = (*C.GSocketControlMessage)(unsafe.Pointer(m.Native()))

	var _cret C.gsize

	_cret = C.g_socket_control_message_get_size(_arg0)

	var _gsize uint

	_gsize = (uint)(_cret)

	return _gsize
}

// Serialize converts the data in the message to bytes placed in the
// message.
//
// @data is guaranteed to have enough space to fit the size returned by
// g_socket_control_message_get_size() on this object.
func (m socketControlMessage) Serialize(data interface{}) {
	var _arg0 *C.GSocketControlMessage
	var _arg1 C.gpointer

	_arg0 = (*C.GSocketControlMessage)(unsafe.Pointer(m.Native()))
	_arg1 = C.gpointer(data)

	C.g_socket_control_message_serialize(_arg0, _arg1)
}
