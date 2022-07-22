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
	GTypeContainerAccessible = coreglib.Type(C.gtk_container_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeContainerAccessible, F: marshalContainerAccessible},
	})
}

// ContainerAccessibleOverrider contains methods that are overridable.
type ContainerAccessibleOverrider interface {
}

type ContainerAccessible struct {
	_ [0]func() // equal guard
	WidgetAccessible
}

var (
	_ coreglib.Objector = (*ContainerAccessible)(nil)
)

func classInitContainerAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapContainerAccessible(obj *coreglib.Object) *ContainerAccessible {
	return &ContainerAccessible{
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
	}
}

func marshalContainerAccessible(p uintptr) (interface{}, error) {
	return wrapContainerAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
