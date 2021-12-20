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
		{T: externglib.Type(C.gtk_radio_button_accessible_get_type()), F: marshalRadioButtonAccessibler},
	})
}

type RadioButtonAccessible struct {
	_ [0]func() // equal guard
	ToggleButtonAccessible
}

var (
	_ externglib.Objector = (*RadioButtonAccessible)(nil)
)

func wrapRadioButtonAccessible(obj *externglib.Object) *RadioButtonAccessible {
	return &RadioButtonAccessible{
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
				Object: obj,
				Action: atk.Action{
					Object: obj,
				},
				Image: atk.Image{
					Object: obj,
				},
			},
		},
	}
}

func marshalRadioButtonAccessibler(p uintptr) (interface{}, error) {
	return wrapRadioButtonAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
