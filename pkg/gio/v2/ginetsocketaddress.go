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
		{T: externglib.Type(C.g_inet_socket_address_get_type()), F: marshalInetSocketAddresser},
	})
}

// InetSocketAddresser describes InetSocketAddress's methods.
type InetSocketAddresser interface {
	// Address gets @address's Address.
	Address() *InetAddress
	// Flowinfo gets the `sin6_flowinfo` field from @address, which must be an
	// IPv6 address.
	Flowinfo() uint32
	// Port gets @address's port.
	Port() uint16
	// ScopeID gets the `sin6_scope_id` field from @address, which must be an
	// IPv6 address.
	ScopeID() uint32
}

// InetSocketAddress: IPv4 or IPv6 socket address; that is, the combination of a
// Address and a port number.
type InetSocketAddress struct {
	SocketAddress
}

var (
	_ InetSocketAddresser = (*InetSocketAddress)(nil)
	_ gextras.Nativer     = (*InetSocketAddress)(nil)
)

func wrapInetSocketAddress(obj *externglib.Object) *InetSocketAddress {
	return &InetSocketAddress{
		SocketAddress: SocketAddress{
			Object: obj,
			SocketConnectable: SocketConnectable{
				Object: obj,
			},
		},
	}
}

func marshalInetSocketAddresser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapInetSocketAddress(obj), nil
}

// NewInetSocketAddress creates a new SocketAddress for @address and @port.
func NewInetSocketAddress(address InetAddresser, port uint16) *InetSocketAddress {
	var _arg1 *C.GInetAddress   // out
	var _arg2 C.guint16         // out
	var _cret *C.GSocketAddress // in

	_arg1 = (*C.GInetAddress)(unsafe.Pointer((address).(gextras.Nativer).Native()))
	_arg2 = C.guint16(port)

	_cret = C.g_inet_socket_address_new(_arg1, _arg2)

	var _inetSocketAddress *InetSocketAddress // out

	_inetSocketAddress = wrapInetSocketAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _inetSocketAddress
}

// NewInetSocketAddressFromString creates a new SocketAddress for @address and
// @port.
//
// If @address is an IPv6 address, it can also contain a scope ID (separated
// from the address by a `%`).
func NewInetSocketAddressFromString(address string, port uint) *InetSocketAddress {
	var _arg1 *C.char           // out
	var _arg2 C.guint           // out
	var _cret *C.GSocketAddress // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(address)))
	_arg2 = C.guint(port)

	_cret = C.g_inet_socket_address_new_from_string(_arg1, _arg2)

	var _inetSocketAddress *InetSocketAddress // out

	_inetSocketAddress = wrapInetSocketAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _inetSocketAddress
}

// Address gets @address's Address.
func (address *InetSocketAddress) Address() *InetAddress {
	var _arg0 *C.GInetSocketAddress // out
	var _cret *C.GInetAddress       // in

	_arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_socket_address_get_address(_arg0)

	var _inetAddress *InetAddress // out

	_inetAddress = wrapInetAddress(externglib.Take(unsafe.Pointer(_cret)))

	return _inetAddress
}

// Flowinfo gets the `sin6_flowinfo` field from @address, which must be an IPv6
// address.
func (address *InetSocketAddress) Flowinfo() uint32 {
	var _arg0 *C.GInetSocketAddress // out
	var _cret C.guint32             // in

	_arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_socket_address_get_flowinfo(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// Port gets @address's port.
func (address *InetSocketAddress) Port() uint16 {
	var _arg0 *C.GInetSocketAddress // out
	var _cret C.guint16             // in

	_arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_socket_address_get_port(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// ScopeID gets the `sin6_scope_id` field from @address, which must be an IPv6
// address.
func (address *InetSocketAddress) ScopeID() uint32 {
	var _arg0 *C.GInetSocketAddress // out
	var _cret C.guint32             // in

	_arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_socket_address_get_scope_id(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}
