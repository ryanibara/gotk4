// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gdkx11devicemanager-xi2.go.
var GTypeX11DeviceManagerXI2 = coreglib.Type(C.gdk_x11_device_manager_xi2_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeX11DeviceManagerXI2, F: marshalX11DeviceManagerXI2},
	})
}

type X11DeviceManagerXI2 struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*X11DeviceManagerXI2)(nil)
)

func wrapX11DeviceManagerXI2(obj *coreglib.Object) *X11DeviceManagerXI2 {
	return &X11DeviceManagerXI2{
		Object: obj,
	}
}

func marshalX11DeviceManagerXI2(p uintptr) (interface{}, error) {
	return wrapX11DeviceManagerXI2(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
