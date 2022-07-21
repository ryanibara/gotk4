// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GTypeButtonAccessible returns the GType for the type ButtonAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeButtonAccessible() coreglib.Type {
	gtype := coreglib.Type(C.gtk_button_accessible_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalButtonAccessible)
	return gtype
}

// ButtonAccessibleOverrider contains methods that are overridable.
type ButtonAccessibleOverrider interface {
}

type ButtonAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	*coreglib.Object
	atk.Action
	atk.Image
}

var (
	_ coreglib.Objector = (*ButtonAccessible)(nil)
)

func classInitButtonAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapButtonAccessible(obj *coreglib.Object) *ButtonAccessible {
	return &ButtonAccessible{
		ContainerAccessible: ContainerAccessible{
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
		},
		Object: obj,
		Action: atk.Action{
			Object: obj,
		},
		Image: atk.Image{
			Object: obj,
		},
	}
}

func marshalButtonAccessible(p uintptr) (interface{}, error) {
	return wrapButtonAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
