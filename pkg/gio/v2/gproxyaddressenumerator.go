// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"reflect"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeProxyAddressEnumerator = coreglib.Type(C.g_proxy_address_enumerator_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeProxyAddressEnumerator, F: marshalProxyAddressEnumerator},
	})
}

// ProxyAddressEnumeratorOverrider contains methods that are overridable.
type ProxyAddressEnumeratorOverrider interface {
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

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeProxyAddressEnumerator,
		GoType:        reflect.TypeOf((*ProxyAddressEnumerator)(nil)),
		InitClass:     initClassProxyAddressEnumerator,
		FinalizeClass: finalizeClassProxyAddressEnumerator,
	})
}

func initClassProxyAddressEnumerator(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		InitProxyAddressEnumerator(*ProxyAddressEnumeratorClass)
	}); ok {
		klass := (*ProxyAddressEnumeratorClass)(gextras.NewStructNative(gclass))
		goval.InitProxyAddressEnumerator(klass)
	}
}

func finalizeClassProxyAddressEnumerator(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		FinalizeProxyAddressEnumerator(*ProxyAddressEnumeratorClass)
	}); ok {
		klass := (*ProxyAddressEnumeratorClass)(gextras.NewStructNative(gclass))
		goval.FinalizeProxyAddressEnumerator(klass)
	}
}

func wrapProxyAddressEnumerator(obj *coreglib.Object) *ProxyAddressEnumerator {
	return &ProxyAddressEnumerator{
		SocketAddressEnumerator: SocketAddressEnumerator{
			Object: obj,
		},
	}
}

func marshalProxyAddressEnumerator(p uintptr) (interface{}, error) {
	return wrapProxyAddressEnumerator(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ProxyAddressEnumeratorClass class structure for AddressEnumerator.
//
// An instance of this type is always passed by reference.
type ProxyAddressEnumeratorClass struct {
	*proxyAddressEnumeratorClass
}

// proxyAddressEnumeratorClass is the struct that's finalized.
type proxyAddressEnumeratorClass struct {
	native *C.GProxyAddressEnumeratorClass
}
