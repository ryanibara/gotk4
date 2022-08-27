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
	GTypeNativeSocketAddress = coreglib.Type(C.g_native_socket_address_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeNativeSocketAddress, F: marshalNativeSocketAddress},
	})
}

// NativeSocketAddressOverrides contains methods that are overridable.
type NativeSocketAddressOverrides struct {
}

func defaultNativeSocketAddressOverrides(v *NativeSocketAddress) NativeSocketAddressOverrides {
	return NativeSocketAddressOverrides{}
}

// NativeSocketAddress: socket address of some unknown native type.
type NativeSocketAddress struct {
	_ [0]func() // equal guard
	SocketAddress
}

var (
	_ SocketAddresser = (*NativeSocketAddress)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*NativeSocketAddress, *NativeSocketAddressClass, NativeSocketAddressOverrides](
		GTypeNativeSocketAddress,
		initNativeSocketAddressClass,
		wrapNativeSocketAddress,
		defaultNativeSocketAddressOverrides,
	)
}

func initNativeSocketAddressClass(gclass unsafe.Pointer, overrides NativeSocketAddressOverrides, classInitFunc func(*NativeSocketAddressClass)) {
	if classInitFunc != nil {
		class := (*NativeSocketAddressClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapNativeSocketAddress(obj *coreglib.Object) *NativeSocketAddress {
	return &NativeSocketAddress{
		SocketAddress: SocketAddress{
			Object: obj,
			SocketConnectable: SocketConnectable{
				Object: obj,
			},
		},
	}
}

func marshalNativeSocketAddress(p uintptr) (interface{}, error) {
	return wrapNativeSocketAddress(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NativeSocketAddressClass: instance of this type is always passed by
// reference.
type NativeSocketAddressClass struct {
	*nativeSocketAddressClass
}

// nativeSocketAddressClass is the struct that's finalized.
type nativeSocketAddressClass struct {
	native *C.GNativeSocketAddressClass
}

func (n *NativeSocketAddressClass) ParentClass() *SocketAddressClass {
	valptr := &n.native.parent_class
	var _v *SocketAddressClass // out
	_v = (*SocketAddressClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
