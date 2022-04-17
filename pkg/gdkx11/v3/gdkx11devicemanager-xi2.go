// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gdkx11devicemanager-xi2.go.
var GTypeX11DeviceManagerXI2 = externglib.Type(C.gdk_x11_device_manager_xi2_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeX11DeviceManagerXI2, F: marshalX11DeviceManagerXI2},
	})
}

// X11DeviceManagerXI2Overrider contains methods that are overridable.
type X11DeviceManagerXI2Overrider interface {
	externglib.Objector
}

type X11DeviceManagerXI2 struct {
	_ [0]func() // equal guard
	X11DeviceManagerCore
}

var (
	_ gdk.DeviceManagerer = (*X11DeviceManagerXI2)(nil)
)

func classInitX11DeviceManagerXI2er(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapX11DeviceManagerXI2(obj *externglib.Object) *X11DeviceManagerXI2 {
	return &X11DeviceManagerXI2{
		X11DeviceManagerCore: X11DeviceManagerCore{
			DeviceManager: gdk.DeviceManager{
				Object: obj,
			},
		},
	}
}

func marshalX11DeviceManagerXI2(p uintptr) (interface{}, error) {
	return wrapX11DeviceManagerXI2(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
