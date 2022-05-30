// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkimageaccessible.go.
var GTypeImageAccessible = coreglib.Type(C.gtk_image_accessible_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeImageAccessible, F: marshalImageAccessible},
	})
}

// ImageAccessibleOverrider contains methods that are overridable.
type ImageAccessibleOverrider interface {
}

type ImageAccessible struct {
	_ [0]func() // equal guard
	WidgetAccessible

	atk.Image
}

var (
	_ coreglib.Objector = (*ImageAccessible)(nil)
)

func classInitImageAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapImageAccessible(obj *coreglib.Object) *ImageAccessible {
	return &ImageAccessible{
		WidgetAccessible: WidgetAccessible{
			Accessible: Accessible{
				ObjectClass: atk.ObjectClass{
					Object: obj,
				},
			},
			Component: atk.Component{
				Object: obj,
			},
		},
		Image: atk.Image{
			Object: obj,
		},
	}
}

func marshalImageAccessible(p uintptr) (interface{}, error) {
	return wrapImageAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
