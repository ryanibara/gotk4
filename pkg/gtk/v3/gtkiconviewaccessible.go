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

// glib.Type values for gtkiconviewaccessible.go.
var GTypeIconViewAccessible = externglib.Type(C.gtk_icon_view_accessible_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeIconViewAccessible, F: marshalIconViewAccessible},
	})
}

// IconViewAccessibleOverrider contains methods that are overridable.
type IconViewAccessibleOverrider interface {
}

type IconViewAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	atk.Selection
}

var (
	_ externglib.Objector = (*IconViewAccessible)(nil)
)

func classInitIconViewAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapIconViewAccessible(obj *externglib.Object) *IconViewAccessible {
	return &IconViewAccessible{
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
		Selection: atk.Selection{
			Object: obj,
		},
	}
}

func marshalIconViewAccessible(p uintptr) (interface{}, error) {
	return wrapIconViewAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
