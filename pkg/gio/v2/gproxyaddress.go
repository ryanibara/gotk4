// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.g_proxy_address_get_type()), F: marshalProXYAddress},
	})
}

// ProXYAddress: support for proxied SocketAddress.
type ProXYAddress interface {
	InetSocketAddress
	SocketConnectable

	// DestinationHostname gets @proxy's destination hostname; that is, the name
	// of the host that will be connected to via the proxy, not the name of the
	// proxy itself.
	DestinationHostname() string
	// DestinationPort gets @proxy's destination port; that is, the port on the
	// destination host that will be connected to via the proxy, not the port
	// number of the proxy itself.
	DestinationPort() uint16
	// DestinationProtocol gets the protocol that is being spoken to the
	// destination server; eg, "http" or "ftp".
	DestinationProtocol() string
	// Password gets @proxy's password.
	Password() string
	// Protocol gets @proxy's protocol. eg, "socks" or "http"
	Protocol() string
	// URI gets the proxy URI that @proxy was constructed from.
	URI() string
	// Username gets @proxy's username.
	Username() string
}

// proXYAddress implements the ProXYAddress interface.
type proXYAddress struct {
	InetSocketAddress
	SocketConnectable
}

var _ ProXYAddress = (*proXYAddress)(nil)

// WrapProXYAddress wraps a GObject to the right type. It is
// primarily used internally.
func WrapProXYAddress(obj *externglib.Object) ProXYAddress {
	return ProXYAddress{
		InetSocketAddress: WrapInetSocketAddress(obj),
		SocketConnectable: WrapSocketConnectable(obj),
	}
}

func marshalProXYAddress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapProXYAddress(obj), nil
}

// NewProXYAddress constructs a class ProXYAddress.
func NewProXYAddress(inetaddr InetAddress, port uint16, protocol string, destHostname string, destPort uint16, username string, password string) ProXYAddress {
	var _arg1 *C.GInetAddress
	var _arg2 C.guint16
	var _arg3 *C.gchar
	var _arg4 *C.gchar
	var _arg5 C.guint16
	var _arg6 *C.gchar
	var _arg7 *C.gchar

	_arg1 = (*C.GInetAddress)(unsafe.Pointer(inetaddr.Native()))
	_arg2 = C.guint16(port)
	_arg3 = (*C.gchar)(C.CString(protocol))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(C.CString(destHostname))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = C.guint16(destPort)
	_arg6 = (*C.gchar)(C.CString(username))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = (*C.gchar)(C.CString(password))
	defer C.free(unsafe.Pointer(_arg7))

	var _cret C.GProxyAddress

	cret = C.g_proxy_address_new(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)

	var _proxyAddress ProXYAddress

	_proxyAddress = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ProXYAddress)

	return _proxyAddress
}

// DestinationHostname gets @proxy's destination hostname; that is, the name
// of the host that will be connected to via the proxy, not the name of the
// proxy itself.
func (p proXYAddress) DestinationHostname() string {
	var _arg0 *C.GProxyAddress

	_arg0 = (*C.GProxyAddress)(unsafe.Pointer(p.Native()))

	var _cret *C.gchar

	cret = C.g_proxy_address_get_destination_hostname(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// DestinationPort gets @proxy's destination port; that is, the port on the
// destination host that will be connected to via the proxy, not the port
// number of the proxy itself.
func (p proXYAddress) DestinationPort() uint16 {
	var _arg0 *C.GProxyAddress

	_arg0 = (*C.GProxyAddress)(unsafe.Pointer(p.Native()))

	var _cret C.guint16

	cret = C.g_proxy_address_get_destination_port(_arg0)

	var _guint16 uint16

	_guint16 = (uint16)(_cret)

	return _guint16
}

// DestinationProtocol gets the protocol that is being spoken to the
// destination server; eg, "http" or "ftp".
func (p proXYAddress) DestinationProtocol() string {
	var _arg0 *C.GProxyAddress

	_arg0 = (*C.GProxyAddress)(unsafe.Pointer(p.Native()))

	var _cret *C.gchar

	cret = C.g_proxy_address_get_destination_protocol(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Password gets @proxy's password.
func (p proXYAddress) Password() string {
	var _arg0 *C.GProxyAddress

	_arg0 = (*C.GProxyAddress)(unsafe.Pointer(p.Native()))

	var _cret *C.gchar

	cret = C.g_proxy_address_get_password(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Protocol gets @proxy's protocol. eg, "socks" or "http"
func (p proXYAddress) Protocol() string {
	var _arg0 *C.GProxyAddress

	_arg0 = (*C.GProxyAddress)(unsafe.Pointer(p.Native()))

	var _cret *C.gchar

	cret = C.g_proxy_address_get_protocol(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// URI gets the proxy URI that @proxy was constructed from.
func (p proXYAddress) URI() string {
	var _arg0 *C.GProxyAddress

	_arg0 = (*C.GProxyAddress)(unsafe.Pointer(p.Native()))

	var _cret *C.gchar

	cret = C.g_proxy_address_get_uri(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Username gets @proxy's username.
func (p proXYAddress) Username() string {
	var _arg0 *C.GProxyAddress

	_arg0 = (*C.GProxyAddress)(unsafe.Pointer(p.Native()))

	var _cret *C.gchar

	cret = C.g_proxy_address_get_username(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}