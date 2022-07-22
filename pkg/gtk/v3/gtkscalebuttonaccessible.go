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

// GType values.
var (
	GTypeScaleButtonAccessible = coreglib.Type(C.gtk_scale_button_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeScaleButtonAccessible, F: marshalScaleButtonAccessible},
	})
}

// ScaleButtonAccessibleOverrider contains methods that are overridable.
type ScaleButtonAccessibleOverrider interface {
}

type ScaleButtonAccessible struct {
	_ [0]func() // equal guard
	ButtonAccessible

	*coreglib.Object
	atk.Value
}

var (
	_ coreglib.Objector = (*ScaleButtonAccessible)(nil)
)

func classInitScaleButtonAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapScaleButtonAccessible(obj *coreglib.Object) *ScaleButtonAccessible {
	return &ScaleButtonAccessible{
		ButtonAccessible: ButtonAccessible{
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
		},
		Object: obj,
		Value: atk.Value{
			Object: obj,
		},
	}
}

func marshalScaleButtonAccessible(p uintptr) (interface{}, error) {
	return wrapScaleButtonAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
