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

// glib.Type values for gtkshortcutsgroup.go.
var GTypeShortcutsGroup = externglib.Type(C.gtk_shortcuts_group_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeShortcutsGroup, F: marshalShortcutsGroup},
	})
}

// ShortcutsGroupOverrider contains methods that are overridable.
type ShortcutsGroupOverrider interface {
	externglib.Objector
}

// WrapShortcutsGroupOverrider wraps the ShortcutsGroupOverrider
// interface implementation to access the instance methods.
func WrapShortcutsGroupOverrider(obj ShortcutsGroupOverrider) *ShortcutsGroup {
	return wrapShortcutsGroup(externglib.BaseObject(obj))
}

// ShortcutsGroup represents a group of related keyboard shortcuts or gestures.
// The group has a title. It may optionally be associated with a view of the
// application, which can be used to show only relevant shortcuts depending on
// the application context.
//
// This widget is only meant to be used with ShortcutsWindow.
type ShortcutsGroup struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer         = (*ShortcutsGroup)(nil)
	_ externglib.Objector = (*ShortcutsGroup)(nil)
)

func classInitShortcutsGrouper(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapShortcutsGroup(obj *externglib.Object) *ShortcutsGroup {
	return &ShortcutsGroup{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalShortcutsGroup(p uintptr) (interface{}, error) {
	return wrapShortcutsGroup(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
