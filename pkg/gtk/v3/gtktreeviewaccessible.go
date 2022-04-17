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

// glib.Type values for gtktreeviewaccessible.go.
var GTypeTreeViewAccessible = externglib.Type(C.gtk_tree_view_accessible_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeTreeViewAccessible, F: marshalTreeViewAccessible},
	})
}

// TreeViewAccessibleOverrider contains methods that are overridable.
type TreeViewAccessibleOverrider interface {
	externglib.Objector
}

type TreeViewAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	*externglib.Object
	atk.Selection
	atk.Table
	CellAccessibleParent
}

var (
	_ externglib.Objector = (*TreeViewAccessible)(nil)
)

func classInitTreeViewAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTreeViewAccessible(obj *externglib.Object) *TreeViewAccessible {
	return &TreeViewAccessible{
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
		Selection: atk.Selection{
			Object: obj,
		},
		Table: atk.Table{
			Object: obj,
		},
		CellAccessibleParent: CellAccessibleParent{
			Object: obj,
		},
	}
}

func marshalTreeViewAccessible(p uintptr) (interface{}, error) {
	return wrapTreeViewAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
