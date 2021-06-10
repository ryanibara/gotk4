// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
		{T: externglib.Type(C.g_inet_address_get_type()), F: marshalInetAddress},
	})
}

// InetAddress represents an IPv4 or IPv6 internet address. Use
// g_resolver_lookup_by_name() or g_resolver_lookup_by_name_async() to look up
// the Address for a hostname. Use g_resolver_lookup_by_address() or
// g_resolver_lookup_by_address_async() to look up the hostname for a Address.
//
// To actually connect to a remote host, you will need a SocketAddress (which
// includes a Address as well as a port number).
type InetAddress interface {
	gextras.Objector

	// Equal checks if two Address instances are equal, e.g. the same address.
	Equal(otherAddress InetAddress) bool
	// IsAny tests whether @address is the "any" address for its family.
	IsAny() bool
	// IsLinkLocal tests whether @address is a link-local address (that is, if
	// it identifies a host on a local network that is not connected to the
	// Internet).
	IsLinkLocal() bool
	// IsLoopback tests whether @address is the loopback address for its family.
	IsLoopback() bool
	// IsMcGlobal tests whether @address is a global multicast address.
	IsMcGlobal() bool
	// IsMcLinkLocal tests whether @address is a link-local multicast address.
	IsMcLinkLocal() bool
	// IsMcNodeLocal tests whether @address is a node-local multicast address.
	IsMcNodeLocal() bool
	// IsMcOrgLocal tests whether @address is an organization-local multicast
	// address.
	IsMcOrgLocal() bool
	// IsMcSiteLocal tests whether @address is a site-local multicast address.
	IsMcSiteLocal() bool
	// IsMulticast tests whether @address is a multicast address.
	IsMulticast() bool
	// IsSiteLocal tests whether @address is a site-local address such as
	// 10.0.0.1 (that is, the address identifies a host on a local network that
	// can not be reached directly from the Internet, but which may have
	// outgoing Internet connectivity via a NAT or firewall).
	IsSiteLocal() bool
	// NativeSize gets the size of the native raw binary address for @address.
	// This is the size of the data that you get from g_inet_address_to_bytes().
	NativeSize() uint
	// ToBytes gets the raw binary address data from @address.
	ToBytes() *byte
	// String converts @address to string form.
	String() string
}

// inetAddress implements the InetAddress interface.
type inetAddress struct {
	gextras.Objector
}

var _ InetAddress = (*inetAddress)(nil)

// WrapInetAddress wraps a GObject to the right type. It is
// primarily used internally.
func WrapInetAddress(obj *externglib.Object) InetAddress {
	return InetAddress{
		Objector: obj,
	}
}

func marshalInetAddress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapInetAddress(obj), nil
}

// Equal checks if two Address instances are equal, e.g. the same address.
func (a inetAddress) Equal(otherAddress InetAddress) bool {
	var _arg0 *C.GInetAddress
	var _arg1 *C.GInetAddress

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GInetAddress)(unsafe.Pointer(otherAddress.Native()))

	var _cret C.gboolean

	_cret = C.g_inet_address_equal(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsAny tests whether @address is the "any" address for its family.
func (a inetAddress) IsAny() bool {
	var _arg0 *C.GInetAddress

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	_cret = C.g_inet_address_get_is_any(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsLinkLocal tests whether @address is a link-local address (that is, if
// it identifies a host on a local network that is not connected to the
// Internet).
func (a inetAddress) IsLinkLocal() bool {
	var _arg0 *C.GInetAddress

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	_cret = C.g_inet_address_get_is_link_local(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsLoopback tests whether @address is the loopback address for its family.
func (a inetAddress) IsLoopback() bool {
	var _arg0 *C.GInetAddress

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	_cret = C.g_inet_address_get_is_loopback(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsMcGlobal tests whether @address is a global multicast address.
func (a inetAddress) IsMcGlobal() bool {
	var _arg0 *C.GInetAddress

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	_cret = C.g_inet_address_get_is_mc_global(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsMcLinkLocal tests whether @address is a link-local multicast address.
func (a inetAddress) IsMcLinkLocal() bool {
	var _arg0 *C.GInetAddress

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	_cret = C.g_inet_address_get_is_mc_link_local(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsMcNodeLocal tests whether @address is a node-local multicast address.
func (a inetAddress) IsMcNodeLocal() bool {
	var _arg0 *C.GInetAddress

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	_cret = C.g_inet_address_get_is_mc_node_local(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsMcOrgLocal tests whether @address is an organization-local multicast
// address.
func (a inetAddress) IsMcOrgLocal() bool {
	var _arg0 *C.GInetAddress

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	_cret = C.g_inet_address_get_is_mc_org_local(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsMcSiteLocal tests whether @address is a site-local multicast address.
func (a inetAddress) IsMcSiteLocal() bool {
	var _arg0 *C.GInetAddress

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	_cret = C.g_inet_address_get_is_mc_site_local(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsMulticast tests whether @address is a multicast address.
func (a inetAddress) IsMulticast() bool {
	var _arg0 *C.GInetAddress

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	_cret = C.g_inet_address_get_is_multicast(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsSiteLocal tests whether @address is a site-local address such as
// 10.0.0.1 (that is, the address identifies a host on a local network that
// can not be reached directly from the Internet, but which may have
// outgoing Internet connectivity via a NAT or firewall).
func (a inetAddress) IsSiteLocal() bool {
	var _arg0 *C.GInetAddress

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	_cret = C.g_inet_address_get_is_site_local(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// NativeSize gets the size of the native raw binary address for @address.
// This is the size of the data that you get from g_inet_address_to_bytes().
func (a inetAddress) NativeSize() uint {
	var _arg0 *C.GInetAddress

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	var _cret C.gsize

	_cret = C.g_inet_address_get_native_size(_arg0)

	var _gsize uint

	_gsize = (uint)(_cret)

	return _gsize
}

// ToBytes gets the raw binary address data from @address.
func (a inetAddress) ToBytes() *byte {
	var _arg0 *C.GInetAddress

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	var _cret *C.guint8

	_cret = C.g_inet_address_to_bytes(_arg0)

	var _guint8 *byte

	_guint8 = (*byte)(_cret)

	return _guint8
}

// String converts @address to string form.
func (a inetAddress) String() string {
	var _arg0 *C.GInetAddress

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	var _cret *C.gchar

	_cret = C.g_inet_address_to_string(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
