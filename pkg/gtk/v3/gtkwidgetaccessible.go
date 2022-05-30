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

// glib.Type values for gtkwidgetaccessible.go.
var GTypeWidgetAccessible = coreglib.Type(C.gtk_widget_accessible_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeWidgetAccessible, F: marshalWidgetAccessible},
	})
}

// WidgetAccessibleOverrider contains methods that are overridable.
type WidgetAccessibleOverrider interface {
}

type WidgetAccessible struct {
	_ [0]func() // equal guard
	Accessible

	atk.Component
}

var (
	_ coreglib.Objector = (*WidgetAccessible)(nil)
)

func classInitWidgetAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapWidgetAccessible(obj *coreglib.Object) *WidgetAccessible {
	return &WidgetAccessible{
		Accessible: Accessible{
			ObjectClass: atk.ObjectClass{
				Object: obj,
			},
		},
		Component: atk.Component{
			Object: obj,
		},
	}
}

func marshalWidgetAccessible(p uintptr) (interface{}, error) {
	return wrapWidgetAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
