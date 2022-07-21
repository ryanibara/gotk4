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

// GTypeStackAccessible returns the GType for the type StackAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeStackAccessible() coreglib.Type {
	gtype := coreglib.Type(C.gtk_stack_accessible_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalStackAccessible)
	return gtype
}

// StackAccessibleOverrider contains methods that are overridable.
type StackAccessibleOverrider interface {
}

type StackAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible
}

var (
	_ coreglib.Objector = (*StackAccessible)(nil)
)

func classInitStackAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapStackAccessible(obj *coreglib.Object) *StackAccessible {
	return &StackAccessible{
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

func marshalStackAccessible(p uintptr) (interface{}, error) {
	return wrapStackAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
