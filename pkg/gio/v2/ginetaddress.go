// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gchar* _gotk4_gio2_InetAddressClass_to_string(GInetAddress*);
import "C"

// glib.Type values for ginetaddress.go.
var GTypeInetAddress = coreglib.Type(C.g_inet_address_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeInetAddress, F: marshalInetAddress},
	})
}

// InetAddressOverrider contains methods that are overridable.
type InetAddressOverrider interface {
	// String converts address to string form.
	//
	// The function returns the following values:
	//
	//    - utf8: representation of address as a string, which should be freed
	//      after use.
	//
	String() string
}

// InetAddress represents an IPv4 or IPv6 internet address. Use
// g_resolver_lookup_by_name() or g_resolver_lookup_by_name_async() to look up
// the Address for a hostname. Use g_resolver_lookup_by_address() or
// g_resolver_lookup_by_address_async() to look up the hostname for a Address.
//
// To actually connect to a remote host, you will need a SocketAddress (which
// includes a Address as well as a port number).
type InetAddress struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*InetAddress)(nil)
)

func classInitInetAddresser(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GInetAddressClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GInetAddressClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ String() string }); ok {
		pclass.to_string = (*[0]byte)(C._gotk4_gio2_InetAddressClass_to_string)
	}
}

//export _gotk4_gio2_InetAddressClass_to_string
func _gotk4_gio2_InetAddressClass_to_string(arg0 *C.GInetAddress) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ String() string })

	utf8 := iface.String()

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

func wrapInetAddress(obj *coreglib.Object) *InetAddress {
	return &InetAddress{
		Object: obj,
	}
}

func marshalInetAddress(p uintptr) (interface{}, error) {
	return wrapInetAddress(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewInetAddressFromString parses string as an IP address and creates a new
// Address.
//
// The function takes the following parameters:
//
//    - str: string representation of an IP address.
//
// The function returns the following values:
//
//    - inetAddress (optional): new Address corresponding to string, or NULL if
//      string could not be parsed. Free the returned object with
//      g_object_unref().
//
func NewInetAddressFromString(str string) *InetAddress {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg0))
	*(*string)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "InetAddress").InvokeMethod("new_InetAddress_from_string", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(str)

	var _inetAddress *InetAddress // out

	if _cret != nil {
		_inetAddress = wrapInetAddress(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _inetAddress
}

// Equal checks if two Address instances are equal, e.g. the same address.
//
// The function takes the following parameters:
//
//    - otherAddress: another Address.
//
// The function returns the following values:
//
//    - ok: TRUE if address and other_address are equal, FALSE otherwise.
//
func (address *InetAddress) Equal(otherAddress *InetAddress) bool {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(otherAddress).Native()))
	*(**InetAddress)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gio", "InetAddress").InvokeMethod("equal", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(address)
	runtime.KeepAlive(otherAddress)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsAny tests whether address is the "any" address for its family.
//
// The function returns the following values:
//
//    - ok: TRUE if address is the "any" address for its family.
//
func (address *InetAddress) IsAny() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	*(**InetAddress)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "InetAddress").InvokeMethod("get_is_any", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsLinkLocal tests whether address is a link-local address (that is, if it
// identifies a host on a local network that is not connected to the Internet).
//
// The function returns the following values:
//
//    - ok: TRUE if address is a link-local address.
//
func (address *InetAddress) IsLinkLocal() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	*(**InetAddress)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "InetAddress").InvokeMethod("get_is_link_local", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsLoopback tests whether address is the loopback address for its family.
//
// The function returns the following values:
//
//    - ok: TRUE if address is the loopback address for its family.
//
func (address *InetAddress) IsLoopback() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	*(**InetAddress)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "InetAddress").InvokeMethod("get_is_loopback", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMcGlobal tests whether address is a global multicast address.
//
// The function returns the following values:
//
//    - ok: TRUE if address is a global multicast address.
//
func (address *InetAddress) IsMcGlobal() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	*(**InetAddress)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "InetAddress").InvokeMethod("get_is_mc_global", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMcLinkLocal tests whether address is a link-local multicast address.
//
// The function returns the following values:
//
//    - ok: TRUE if address is a link-local multicast address.
//
func (address *InetAddress) IsMcLinkLocal() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	*(**InetAddress)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "InetAddress").InvokeMethod("get_is_mc_link_local", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMcNodeLocal tests whether address is a node-local multicast address.
//
// The function returns the following values:
//
//    - ok: TRUE if address is a node-local multicast address.
//
func (address *InetAddress) IsMcNodeLocal() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	*(**InetAddress)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "InetAddress").InvokeMethod("get_is_mc_node_local", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMcOrgLocal tests whether address is an organization-local multicast
// address.
//
// The function returns the following values:
//
//    - ok: TRUE if address is an organization-local multicast address.
//
func (address *InetAddress) IsMcOrgLocal() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	*(**InetAddress)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "InetAddress").InvokeMethod("get_is_mc_org_local", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMcSiteLocal tests whether address is a site-local multicast address.
//
// The function returns the following values:
//
//    - ok: TRUE if address is a site-local multicast address.
//
func (address *InetAddress) IsMcSiteLocal() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	*(**InetAddress)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "InetAddress").InvokeMethod("get_is_mc_site_local", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMulticast tests whether address is a multicast address.
//
// The function returns the following values:
//
//    - ok: TRUE if address is a multicast address.
//
func (address *InetAddress) IsMulticast() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	*(**InetAddress)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "InetAddress").InvokeMethod("get_is_multicast", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSiteLocal tests whether address is a site-local address such as 10.0.0.1
// (that is, the address identifies a host on a local network that can not be
// reached directly from the Internet, but which may have outgoing Internet
// connectivity via a NAT or firewall).
//
// The function returns the following values:
//
//    - ok: TRUE if address is a site-local address.
//
func (address *InetAddress) IsSiteLocal() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	*(**InetAddress)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "InetAddress").InvokeMethod("get_is_site_local", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NativeSize gets the size of the native raw binary address for address. This
// is the size of the data that you get from g_inet_address_to_bytes().
//
// The function returns the following values:
//
//    - gsize: number of bytes used for the native version of address.
//
func (address *InetAddress) NativeSize() uint {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gsize // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	*(**InetAddress)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "InetAddress").InvokeMethod("get_native_size", args[:], nil)
	_cret = *(*C.gsize)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(address)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// String converts address to string form.
//
// The function returns the following values:
//
//    - utf8: representation of address as a string, which should be freed after
//      use.
//
func (address *InetAddress) String() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	*(**InetAddress)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "InetAddress").InvokeMethod("to_string", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(address)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
