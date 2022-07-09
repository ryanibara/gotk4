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
// #include <glib-object.h>
import "C"

// GTypeStatusbarAccessible returns the GType for the type StatusbarAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeStatusbarAccessible() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "StatusbarAccessible").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalStatusbarAccessible)
	return gtype
}

// StatusbarAccessibleOverrider contains methods that are overridable.
type StatusbarAccessibleOverrider interface {
}

type StatusbarAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible
}

var (
	_ coreglib.Objector = (*StatusbarAccessible)(nil)
)

func classInitStatusbarAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapStatusbarAccessible(obj *coreglib.Object) *StatusbarAccessible {
	return &StatusbarAccessible{
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

func marshalStatusbarAccessible(p uintptr) (interface{}, error) {
	return wrapStatusbarAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
