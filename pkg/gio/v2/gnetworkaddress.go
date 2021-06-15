// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 glib-2.0 gobject-introspection-1.0
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
		{T: externglib.Type(C.g_network_address_get_type()), F: marshalNetworkAddress},
	})
}

// NetworkAddress provides an easy way to resolve a hostname and then attempt to
// connect to that host, handling the possibility of multiple IP addresses and
// multiple address families.
//
// The enumeration results of resolved addresses *may* be cached as long as this
// object is kept alive which may have unexpected results if alive for too long.
//
// See Connectable for an example of using the connectable interface.
type NetworkAddress interface {
	gextras.Objector
	SocketConnectable

	// Hostname gets @addr's hostname. This might be either UTF-8 or
	// ASCII-encoded, depending on what @addr was created with.
	Hostname() string
	// Port gets @addr's port number
	Port() uint16
	// Scheme gets @addr's scheme
	Scheme() string
}

// networkAddress implements the NetworkAddress class.
type networkAddress struct {
	gextras.Objector
	SocketConnectable
}

var _ NetworkAddress = (*networkAddress)(nil)

// WrapNetworkAddress wraps a GObject to the right type. It is
// primarily used internally.
func WrapNetworkAddress(obj *externglib.Object) NetworkAddress {
	return networkAddress{
		Objector:          obj,
		SocketConnectable: WrapSocketConnectable(obj),
	}
}

func marshalNetworkAddress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNetworkAddress(obj), nil
}

// NewNetworkAddress constructs a class NetworkAddress.
func NewNetworkAddress(hostname string, port uint16) NetworkAddress {
	var _arg1 *C.gchar          // out
	var _arg2 C.guint16         // out
	var _cret C.GNetworkAddress // in

	_arg1 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.guint16)(port)

	_cret = C.g_network_address_new(_arg1, _arg2)

	var _networkAddress NetworkAddress // out

	_networkAddress = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(NetworkAddress)

	return _networkAddress
}

// NewNetworkAddressLoopback constructs a class NetworkAddress.
func NewNetworkAddressLoopback(port uint16) NetworkAddress {
	var _arg1 C.guint16         // out
	var _cret C.GNetworkAddress // in

	_arg1 = (C.guint16)(port)

	_cret = C.g_network_address_new_loopback(_arg1)

	var _networkAddress NetworkAddress // out

	_networkAddress = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(NetworkAddress)

	return _networkAddress
}

// Hostname gets @addr's hostname. This might be either UTF-8 or
// ASCII-encoded, depending on what @addr was created with.
func (a networkAddress) Hostname() string {
	var _arg0 *C.GNetworkAddress // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_network_address_get_hostname(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Port gets @addr's port number
func (a networkAddress) Port() uint16 {
	var _arg0 *C.GNetworkAddress // out
	var _cret C.guint16          // in

	_arg0 = (*C.GNetworkAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_network_address_get_port(_arg0)

	var _guint16 uint16 // out

	_guint16 = (uint16)(_cret)

	return _guint16
}

// Scheme gets @addr's scheme
func (a networkAddress) Scheme() string {
	var _arg0 *C.GNetworkAddress // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_network_address_get_scheme(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}
