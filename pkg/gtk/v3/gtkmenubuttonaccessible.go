// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_menu_button_accessible_get_type()), F: marshalMenuButtonAccessibler},
	})
}

type MenuButtonAccessible struct {
	ToggleButtonAccessible
}

func wrapMenuButtonAccessible(obj *externglib.Object) *MenuButtonAccessible {
	return &MenuButtonAccessible{
		ToggleButtonAccessible: ToggleButtonAccessible{
			ButtonAccessible: ButtonAccessible{
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
				Action: atk.Action{
					Object: obj,
				},
				Image: atk.Image{
					Object: obj,
				},
				Object: obj,
			},
		},
	}
}

func marshalMenuButtonAccessibler(p uintptr) (interface{}, error) {
	return wrapMenuButtonAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
