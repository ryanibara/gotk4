// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/x11/gdkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_device_manager_xi2_get_type()), F: marshalX11DeviceManagerXI2},
	})
}

type X11DeviceManagerXI2 interface {
	gextras.Objector
}

// x11DeviceManagerXI2 implements the X11DeviceManagerXI2 interface.
type x11DeviceManagerXI2 struct {
	gextras.Objector
}

var _ X11DeviceManagerXI2 = (*x11DeviceManagerXI2)(nil)

// WrapX11DeviceManagerXI2 wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11DeviceManagerXI2(obj *externglib.Object) X11DeviceManagerXI2 {
	return X11DeviceManagerXI2{
		Objector: obj,
	}
}

func marshalX11DeviceManagerXI2(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11DeviceManagerXI2(obj), nil
}
