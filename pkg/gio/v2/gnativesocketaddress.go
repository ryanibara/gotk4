// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_native_socket_address_get_type()), F: marshalNativeSocketAddresser},
	})
}

// NativeSocketAddress: socket address of some unknown native type.
type NativeSocketAddress struct {
	SocketAddress
}

func wrapNativeSocketAddress(obj *externglib.Object) *NativeSocketAddress {
	return &NativeSocketAddress{
		SocketAddress: SocketAddress{
			Object: obj,
			SocketConnectable: SocketConnectable{
				Object: obj,
			},
		},
	}
}

func marshalNativeSocketAddresser(p uintptr) (interface{}, error) {
	return wrapNativeSocketAddress(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewNativeSocketAddress creates a new SocketAddress for native and len.
//
// The function takes the following parameters:
//
//    - native address object.
//    - len: length of native, in bytes.
//
func NewNativeSocketAddress(native cgo.Handle, len uint) *NativeSocketAddress {
	var _arg1 C.gpointer        // out
	var _arg2 C.gsize           // out
	var _cret *C.GSocketAddress // in

	_arg1 = (C.gpointer)(unsafe.Pointer(native))
	_arg2 = C.gsize(len)

	_cret = C.g_native_socket_address_new(_arg1, _arg2)
	runtime.KeepAlive(native)
	runtime.KeepAlive(len)

	var _nativeSocketAddress *NativeSocketAddress // out

	_nativeSocketAddress = wrapNativeSocketAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _nativeSocketAddress
}
