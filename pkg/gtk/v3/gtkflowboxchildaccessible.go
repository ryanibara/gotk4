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

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_flow_box_child_accessible_get_type()), F: marshalFlowBoxChildAccessibler},
	})
}

type FlowBoxChildAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible
}

var (
	_ externglib.Objector = (*FlowBoxChildAccessible)(nil)
)

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

func marshalFlowBoxChildAccessibler(p uintptr) (interface{}, error) {
	return wrapFlowBoxChildAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
