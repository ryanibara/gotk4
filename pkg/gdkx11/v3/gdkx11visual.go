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

// glib.Type values for gdkx11visual.go.
var GTypeX11Visual = externglib.Type(C.gdk_x11_visual_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeX11Visual, F: marshalX11Visual},
	})
}

// X11VisualOverrider contains methods that are overridable.
type X11VisualOverrider interface {
	externglib.Objector
}

type X11Visual struct {
	_ [0]func() // equal guard
	gdk.Visual
}

var (
	_ externglib.Objector = (*X11Visual)(nil)
)

func classInitX11Visualer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapX11Visual(obj *externglib.Object) *X11Visual {
	return &X11Visual{
		Visual: gdk.Visual{
			Object: obj,
		},
	}
}

func marshalX11Visual(p uintptr) (interface{}, error) {
	return wrapX11Visual(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
