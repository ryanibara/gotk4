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

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_proxy_address_enumerator_get_type()), F: marshalProxyAddressEnumeratorrer},
	})
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

func wrapProxyAddressEnumerator(obj *externglib.Object) *ProxyAddressEnumerator {
	return &ProxyAddressEnumerator{
		SocketAddressEnumerator: SocketAddressEnumerator{
			Object: obj,
		},
	}
}

func marshalProxyAddressEnumeratorrer(p uintptr) (interface{}, error) {
	return wrapProxyAddressEnumerator(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
