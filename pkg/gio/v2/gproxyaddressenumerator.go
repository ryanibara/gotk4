// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gproxyaddressenumerator.go.
var GTypeProxyAddressEnumerator = externglib.Type(C.g_proxy_address_enumerator_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeProxyAddressEnumerator, F: marshalProxyAddressEnumerator},
	})
}

// ProxyAddressEnumeratorOverrider contains methods that are overridable.
type ProxyAddressEnumeratorOverrider interface {
	externglib.Objector
}

// WrapProxyAddressEnumeratorOverrider wraps the ProxyAddressEnumeratorOverrider
// interface implementation to access the instance methods.
func WrapProxyAddressEnumeratorOverrider(obj ProxyAddressEnumeratorOverrider) *ProxyAddressEnumerator {
	return wrapProxyAddressEnumerator(externglib.BaseObject(obj))
}

// ProxyAddressEnumerator is a wrapper around AddressEnumerator which takes the
// Address instances returned by the AddressEnumerator and wraps them in Address
// instances, using the given AddressEnumerator:proxy-resolver.
//
// This enumerator will be returned (for example, by
// g_socket_connectable_enumerate()) as appropriate when a proxy is configured;
// there should be no need to manually wrap a AddressEnumerator instance with
// one.
type ProxyAddressEnumerator struct {
	_ [0]func() // equal guard
	SocketAddressEnumerator
}

var (
	_ SocketAddressEnumeratorrer = (*ProxyAddressEnumerator)(nil)
)

func classInitProxyAddressEnumeratorrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapProxyAddressEnumerator(obj *externglib.Object) *ProxyAddressEnumerator {
	return &ProxyAddressEnumerator{
		SocketAddressEnumerator: SocketAddressEnumerator{
			Object: obj,
		},
	}
}

func marshalProxyAddressEnumerator(p uintptr) (interface{}, error) {
	return wrapProxyAddressEnumerator(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
