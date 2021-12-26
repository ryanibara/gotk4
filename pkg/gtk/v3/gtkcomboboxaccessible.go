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

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_combo_box_accessible_get_type()), F: marshalComboBoxAccessibler},
	})
}

type ComboBoxAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	*externglib.Object
	atk.Action
	atk.Selection
}

var (
	_ externglib.Objector = (*ComboBoxAccessible)(nil)
)

func wrapComboBoxAccessible(obj *externglib.Object) *ComboBoxAccessible {
	return &ComboBoxAccessible{
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
	}
}

func marshalComboBoxAccessibler(p uintptr) (interface{}, error) {
	return wrapComboBoxAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
