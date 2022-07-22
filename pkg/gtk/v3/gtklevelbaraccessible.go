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

// GType values.
var (
	GTypeLevelBarAccessible = coreglib.Type(C.gtk_level_bar_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeLevelBarAccessible, F: marshalLevelBarAccessible},
	})
}

// LevelBarAccessibleOverrider contains methods that are overridable.
type LevelBarAccessibleOverrider interface {
}

type LevelBarAccessible struct {
	_ [0]func() // equal guard
	WidgetAccessible

	atk.Value
}

var (
	_ coreglib.Objector = (*LevelBarAccessible)(nil)
)

func classInitLevelBarAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapLevelBarAccessible(obj *coreglib.Object) *LevelBarAccessible {
	return &LevelBarAccessible{
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
		Value: atk.Value{
			Object: obj,
		},
	}
}

func marshalLevelBarAccessible(p uintptr) (interface{}, error) {
	return wrapLevelBarAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
