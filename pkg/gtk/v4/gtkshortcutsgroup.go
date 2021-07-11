// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcuts_group_get_type()), F: marshalShortcutsGrouper},
	})
}

// ShortcutsGrouper describes ShortcutsGroup's methods.
type ShortcutsGrouper interface {
	privateShortcutsGroup()
}

// ShortcutsGroup: `GtkShortcutsGroup` represents a group of related keyboard
// shortcuts or gestures.
//
// The group has a title. It may optionally be associated with a view of the
// application, which can be used to show only relevant shortcuts depending on
// the application context.
//
// This widget is only meant to be used with [class@Gtk.ShortcutsWindow].
type ShortcutsGroup struct {
	Box
}

var (
	_ ShortcutsGrouper = (*ShortcutsGroup)(nil)
	_ gextras.Nativer  = (*ShortcutsGroup)(nil)
)

func wrapShortcutsGroup(obj *externglib.Object) ShortcutsGrouper {
	return &ShortcutsGroup{
		Box: Box{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalShortcutsGrouper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapShortcutsGroup(obj), nil
}

func (*ShortcutsGroup) privateShortcutsGroup() {}
