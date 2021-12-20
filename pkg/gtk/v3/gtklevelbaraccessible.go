// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_level_bar_accessible_get_type()), F: marshalLevelBarAccessibler},
	})
}

type LevelBarAccessible struct {
	_ [0]func() // equal guard
	WidgetAccessible

	atk.Value
}

var (
	_ externglib.Objector = (*LevelBarAccessible)(nil)
)

func wrapLevelBarAccessible(obj *externglib.Object) *LevelBarAccessible {
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

func marshalLevelBarAccessibler(p uintptr) (interface{}, error) {
	return wrapLevelBarAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
