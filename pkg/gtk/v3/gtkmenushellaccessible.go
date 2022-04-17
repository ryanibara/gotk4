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

// glib.Type values for gtkmenushellaccessible.go.
var GTypeMenuShellAccessible = externglib.Type(C.gtk_menu_shell_accessible_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeMenuShellAccessible, F: marshalMenuShellAccessible},
	})
}

// MenuShellAccessibleOverrider contains methods that are overridable.
type MenuShellAccessibleOverrider interface {
	externglib.Objector
}

// WrapMenuShellAccessibleOverrider wraps the MenuShellAccessibleOverrider
// interface implementation to access the instance methods.
func WrapMenuShellAccessibleOverrider(obj MenuShellAccessibleOverrider) *MenuShellAccessible {
	return wrapMenuShellAccessible(externglib.BaseObject(obj))
}

type MenuShellAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	atk.Selection
}

var (
	_ externglib.Objector = (*MenuShellAccessible)(nil)
)

func classInitMenuShellAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapMenuShellAccessible(obj *externglib.Object) *MenuShellAccessible {
	return &MenuShellAccessible{
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

func marshalMenuShellAccessible(p uintptr) (interface{}, error) {
	return wrapMenuShellAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
