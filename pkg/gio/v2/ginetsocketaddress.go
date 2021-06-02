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
		{T: externglib.Type(C.g_inet_socket_address_get_type()), F: marshalInetSocketAddress},
	})
}

// InetSocketAddress: an IPv4 or IPv6 socket address; that is, the combination
// of a Address and a port number.
type InetSocketAddress interface {
	SocketAddress
	SocketConnectable

	// Address gets @address's Address.
	Address() InetAddress
	// Flowinfo gets the `sin6_flowinfo` field from @address, which must be an
	// IPv6 address.
	Flowinfo() uint32
	// Port gets @address's port.
	Port() uint16
	// ScopeID gets the `sin6_scope_id` field from @address, which must be an
	// IPv6 address.
	ScopeID() uint32
}

// inetSocketAddress implements the InetSocketAddress interface.
type inetSocketAddress struct {
	SocketAddress
	SocketConnectable
}

var _ InetSocketAddress = (*inetSocketAddress)(nil)

// WrapInetSocketAddress wraps a GObject to the right type. It is
// primarily used internally.
func WrapInetSocketAddress(obj *externglib.Object) InetSocketAddress {
	return InetSocketAddress{
		SocketAddress:     WrapSocketAddress(obj),
		SocketConnectable: WrapSocketConnectable(obj),
	}
}

func marshalInetSocketAddress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapInetSocketAddress(obj), nil
}

// NewInetSocketAddress constructs a class InetSocketAddress.
func NewInetSocketAddress(address InetAddress, port uint16) InetSocketAddress {
	var arg1 *C.GInetAddress
	var arg2 C.guint16

	arg1 = (*C.GInetAddress)(address.Native())
	arg2 = C.guint16(port)

	ret := C.g_inet_socket_address_new(arg1, arg2)

	var ret0 InetSocketAddress

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(InetSocketAddress)

	return ret0
}

// NewInetSocketAddressFromString constructs a class InetSocketAddress.
func NewInetSocketAddressFromString(address string, port uint) InetSocketAddress {
	var arg1 *C.char
	var arg2 C.guint

	arg1 = (*C.gchar)(C.CString(address))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.guint(port)

	ret := C.g_inet_socket_address_new_from_string(arg1, arg2)

	var ret0 InetSocketAddress

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(InetSocketAddress)

	return ret0
}

// Address gets @address's Address.
func (a inetSocketAddress) Address() InetAddress {
	var arg0 *C.GInetSocketAddress

	arg0 = (*C.GInetSocketAddress)(a.Native())

	ret := C.g_inet_socket_address_get_address(arg0)

	var ret0 InetAddress

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(InetAddress)

	return ret0
}

// Flowinfo gets the `sin6_flowinfo` field from @address, which must be an
// IPv6 address.
func (a inetSocketAddress) Flowinfo() uint32 {
	var arg0 *C.GInetSocketAddress

	arg0 = (*C.GInetSocketAddress)(a.Native())

	ret := C.g_inet_socket_address_get_flowinfo(arg0)

	var ret0 uint32

	ret0 = uint32(ret)

	return ret0
}

// Port gets @address's port.
func (a inetSocketAddress) Port() uint16 {
	var arg0 *C.GInetSocketAddress

	arg0 = (*C.GInetSocketAddress)(a.Native())

	ret := C.g_inet_socket_address_get_port(arg0)

	var ret0 uint16

	ret0 = uint16(ret)

	return ret0
}

// ScopeID gets the `sin6_scope_id` field from @address, which must be an
// IPv6 address.
func (a inetSocketAddress) ScopeID() uint32 {
	var arg0 *C.GInetSocketAddress

	arg0 = (*C.GInetSocketAddress)(a.Native())

	ret := C.g_inet_socket_address_get_scope_id(arg0)

	var ret0 uint32

	ret0 = uint32(ret)

	return ret0
}

type InetSocketAddressPrivate struct {
	native C.GInetSocketAddressPrivate
}

// WrapInetSocketAddressPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapInetSocketAddressPrivate(ptr unsafe.Pointer) *InetSocketAddressPrivate {
	if ptr == nil {
		return nil
	}

	return (*InetSocketAddressPrivate)(ptr)
}

func marshalInetSocketAddressPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapInetSocketAddressPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (i *InetSocketAddressPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}
