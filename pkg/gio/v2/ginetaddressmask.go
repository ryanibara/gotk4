// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GTypeInetAddressMask returns the GType for the type InetAddressMask.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeInetAddressMask() coreglib.Type {
	gtype := coreglib.Type(C.g_inet_address_mask_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalInetAddressMask)
	return gtype
}

// InetAddressMaskOverrider contains methods that are overridable.
type InetAddressMaskOverrider interface {
}

// InetAddressMask represents a range of IPv4 or IPv6 addresses described by a
// base address and a length indicating how many bits of the base address are
// relevant for matching purposes. These are often given in string form. Eg,
// "10.0.0.0/8", or "fe80::/10".
type InetAddressMask struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Initable
}

var (
	_ coreglib.Objector = (*InetAddressMask)(nil)
)

func classInitInetAddressMasker(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapInetAddressMask(obj *coreglib.Object) *InetAddressMask {
	return &InetAddressMask{
		Object: obj,
		Initable: Initable{
			Object: obj,
		},
	}
}

func marshalInetAddressMask(p uintptr) (interface{}, error) {
	return wrapInetAddressMask(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewInetAddressMask creates a new AddressMask representing all addresses whose
// first length bits match addr.
//
// The function takes the following parameters:
//
//    - addr: Address.
//    - length: number of bits of addr to use.
//
// The function returns the following values:
//
//    - inetAddressMask: new AddressMask, or NULL on error.
//
func NewInetAddressMask(addr *InetAddress, length uint32) (*InetAddressMask, error) {
	var _arg1 *C.GInetAddress     // out
	var _arg2 C.guint             // out
	var _cret *C.GInetAddressMask // in
	var _cerr *C.GError           // in

	_arg1 = (*C.GInetAddress)(unsafe.Pointer(coreglib.InternObject(addr).Native()))
	_arg2 = C.guint(length)

	_cret = C.g_inet_address_mask_new(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(addr)
	runtime.KeepAlive(length)

	var _inetAddressMask *InetAddressMask // out
	var _goerr error                      // out

	_inetAddressMask = wrapInetAddressMask(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _inetAddressMask, _goerr
}

// NewInetAddressMaskFromString parses mask_string as an IP address and
// (optional) length, and creates a new AddressMask. The length, if present, is
// delimited by a "/". If it is not present, then the length is assumed to be
// the full length of the address.
//
// The function takes the following parameters:
//
//    - maskString: IP address or address/length string.
//
// The function returns the following values:
//
//    - inetAddressMask: new AddressMask corresponding to string, or NULL on
//      error.
//
func NewInetAddressMaskFromString(maskString string) (*InetAddressMask, error) {
	var _arg1 *C.gchar            // out
	var _cret *C.GInetAddressMask // in
	var _cerr *C.GError           // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(maskString)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_inet_address_mask_new_from_string(_arg1, &_cerr)
	runtime.KeepAlive(maskString)

	var _inetAddressMask *InetAddressMask // out
	var _goerr error                      // out

	_inetAddressMask = wrapInetAddressMask(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _inetAddressMask, _goerr
}

// Equal tests if mask and mask2 are the same mask.
//
// The function takes the following parameters:
//
//    - mask2: another AddressMask.
//
// The function returns the following values:
//
//    - ok: whether mask and mask2 are the same mask.
//
func (mask *InetAddressMask) Equal(mask2 *InetAddressMask) bool {
	var _arg0 *C.GInetAddressMask // out
	var _arg1 *C.GInetAddressMask // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer(coreglib.InternObject(mask).Native()))
	_arg1 = (*C.GInetAddressMask)(unsafe.Pointer(coreglib.InternObject(mask2).Native()))

	_cret = C.g_inet_address_mask_equal(_arg0, _arg1)
	runtime.KeepAlive(mask)
	runtime.KeepAlive(mask2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Address gets mask's base address.
//
// The function returns the following values:
//
//    - inetAddress mask's base address.
//
func (mask *InetAddressMask) Address() *InetAddress {
	var _arg0 *C.GInetAddressMask // out
	var _cret *C.GInetAddress     // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer(coreglib.InternObject(mask).Native()))

	_cret = C.g_inet_address_mask_get_address(_arg0)
	runtime.KeepAlive(mask)

	var _inetAddress *InetAddress // out

	_inetAddress = wrapInetAddress(coreglib.Take(unsafe.Pointer(_cret)))

	return _inetAddress
}

// Family gets the Family of mask's address.
//
// The function returns the following values:
//
//    - socketFamily of mask's address.
//
func (mask *InetAddressMask) Family() SocketFamily {
	var _arg0 *C.GInetAddressMask // out
	var _cret C.GSocketFamily     // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer(coreglib.InternObject(mask).Native()))

	_cret = C.g_inet_address_mask_get_family(_arg0)
	runtime.KeepAlive(mask)

	var _socketFamily SocketFamily // out

	_socketFamily = SocketFamily(_cret)

	return _socketFamily
}

// Length gets mask's length.
//
// The function returns the following values:
//
//    - guint mask's length.
//
func (mask *InetAddressMask) Length() uint32 {
	var _arg0 *C.GInetAddressMask // out
	var _cret C.guint             // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer(coreglib.InternObject(mask).Native()))

	_cret = C.g_inet_address_mask_get_length(_arg0)
	runtime.KeepAlive(mask)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// Matches tests if address falls within the range described by mask.
//
// The function takes the following parameters:
//
//    - address: Address.
//
// The function returns the following values:
//
//    - ok: whether address falls within the range described by mask.
//
func (mask *InetAddressMask) Matches(address *InetAddress) bool {
	var _arg0 *C.GInetAddressMask // out
	var _arg1 *C.GInetAddress     // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer(coreglib.InternObject(mask).Native()))
	_arg1 = (*C.GInetAddress)(unsafe.Pointer(coreglib.InternObject(address).Native()))

	_cret = C.g_inet_address_mask_matches(_arg0, _arg1)
	runtime.KeepAlive(mask)
	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// String converts mask back to its corresponding string form.
//
// The function returns the following values:
//
//    - utf8: string corresponding to mask.
//
func (mask *InetAddressMask) String() string {
	var _arg0 *C.GInetAddressMask // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer(coreglib.InternObject(mask).Native()))

	_cret = C.g_inet_address_mask_to_string(_arg0)
	runtime.KeepAlive(mask)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
