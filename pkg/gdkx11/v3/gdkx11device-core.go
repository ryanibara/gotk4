// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_device_core_get_type()), F: marshalX11DeviceCore},
	})
}

type X11DeviceCore interface {
	gdk.Device
}

// x11DeviceCore implements the X11DeviceCore interface.
type x11DeviceCore struct {
	gdk.Device
}

var _ X11DeviceCore = (*x11DeviceCore)(nil)

// WrapX11DeviceCore wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11DeviceCore(obj *externglib.Object) X11DeviceCore {
	return X11DeviceCore{
		gdk.Device: gdk.WrapDevice(obj),
	}
}

func marshalX11DeviceCore(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11DeviceCore(obj), nil
}
