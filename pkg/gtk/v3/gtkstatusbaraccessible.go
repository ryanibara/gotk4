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

// glib.Type values for gtkstatusbaraccessible.go.
var GTypeStatusbarAccessible = externglib.Type(C.gtk_statusbar_accessible_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeStatusbarAccessible, F: marshalStatusbarAccessible},
	})
}

// StatusbarAccessibleOverrider contains methods that are overridable.
type StatusbarAccessibleOverrider interface {
	externglib.Objector
}

// WrapStatusbarAccessibleOverrider wraps the StatusbarAccessibleOverrider
// interface implementation to access the instance methods.
func WrapStatusbarAccessibleOverrider(obj StatusbarAccessibleOverrider) *StatusbarAccessible {
	return wrapStatusbarAccessible(externglib.BaseObject(obj))
}

type StatusbarAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible
}

var (
	_ externglib.Objector = (*StatusbarAccessible)(nil)
)

func classInitStatusbarAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapStatusbarAccessible(obj *externglib.Object) *StatusbarAccessible {
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
	return wrapStatusbarAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
