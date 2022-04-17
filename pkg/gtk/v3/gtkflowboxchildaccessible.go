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

// glib.Type values for gtkflowboxchildaccessible.go.
var GTypeFlowBoxChildAccessible = externglib.Type(C.gtk_flow_box_child_accessible_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeFlowBoxChildAccessible, F: marshalFlowBoxChildAccessible},
	})
}

// FlowBoxChildAccessibleOverrider contains methods that are overridable.
type FlowBoxChildAccessibleOverrider interface {
	externglib.Objector
}

type FlowBoxChildAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible
}

var (
	_ externglib.Objector = (*FlowBoxChildAccessible)(nil)
)

func classInitFlowBoxChildAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFlowBoxChildAccessible(obj *externglib.Object) *FlowBoxChildAccessible {
	return &FlowBoxChildAccessible{
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
	}
}

func marshalFlowBoxChildAccessible(p uintptr) (interface{}, error) {
	return wrapFlowBoxChildAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
