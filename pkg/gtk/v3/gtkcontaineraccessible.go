// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtkcontaineraccessible.go.
var GTypeContainerAccessible = externglib.Type(C.gtk_container_accessible_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeContainerAccessible, F: marshalContainerAccessible},
	})
}

// ContainerAccessibleOverrider contains methods that are overridable.
type ContainerAccessibleOverrider interface {
	externglib.Objector
}

type ContainerAccessible struct {
	_ [0]func() // equal guard
	WidgetAccessible
}

var (
	_ externglib.Objector = (*ContainerAccessible)(nil)
)

func classInitContainerAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapContainerAccessible(obj *externglib.Object) *ContainerAccessible {
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
	return wrapContainerAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
