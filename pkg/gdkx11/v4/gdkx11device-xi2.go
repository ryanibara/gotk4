// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"fmt"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_device_xi2_get_type()), F: marshalX11DeviceXI2er},
	})
}

type X11DeviceType C.gint

const (
	X11DeviceTypeLogical X11DeviceType = iota
	X11DeviceTypePhysical
	X11DeviceTypeFloating
)

// String returns the name in string for X11DeviceType.
func (x X11DeviceType) String() string {
	switch x {
	case X11DeviceTypeLogical:
		return "Logical"
	case X11DeviceTypePhysical:
		return "Physical"
	case X11DeviceTypeFloating:
		return "Floating"
	default:
		return fmt.Sprintf("X11DeviceType(%d)", x)
	}
}

type X11DeviceXI2 struct {
	gdk.Device
}

func wrapX11DeviceXI2(obj *externglib.Object) *X11DeviceXI2 {
	return &X11DeviceXI2{
		Device: gdk.Device{
			Object: obj,
		},
	}
}

func marshalX11DeviceXI2er(p uintptr) (interface{}, error) {
	return wrapX11DeviceXI2(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
