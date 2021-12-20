// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"sync"
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
		{T: externglib.Type(C.gtk_check_menu_item_accessible_get_type()), F: marshalCheckMenuItemAccessibler},
	})
}

type CheckMenuItemAccessible struct {
	MenuItemAccessible

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ externglib.Objector = (*CheckMenuItemAccessible)(nil)
)

func wrapCheckMenuItemAccessible(obj *externglib.Object) *CheckMenuItemAccessible {
	return &CheckMenuItemAccessible{
		MenuItemAccessible: MenuItemAccessible{
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
			Action: atk.Action{
				Object: obj,
			},
			Selection: atk.Selection{
				Object: obj,
			},
		},
	}
}

func marshalCheckMenuItemAccessibler(p uintptr) (interface{}, error) {
	return wrapCheckMenuItemAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
