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

// glib.Type values for gtktreeviewaccessible.go.
var GTypeTreeViewAccessible = coreglib.Type(C.gtk_tree_view_accessible_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeTreeViewAccessible, F: marshalTreeViewAccessible},
	})
}

// TreeViewAccessibleOverrider contains methods that are overridable.
type TreeViewAccessibleOverrider interface {
}

type TreeViewAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	*coreglib.Object
	atk.Selection
	atk.Table
	CellAccessibleParent
}

var (
	_ coreglib.Objector = (*TreeViewAccessible)(nil)
)

func classInitTreeViewAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTreeViewAccessible(obj *coreglib.Object) *TreeViewAccessible {
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
	return wrapTreeViewAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
