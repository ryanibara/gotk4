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
		{T: externglib.Type(C.gtk_scale_button_accessible_get_type()), F: marshalScaleButtonAccessibler},
	})
}

type ScaleButtonAccessible struct {
	ButtonAccessible

	*externglib.Object
	atk.Value
}

var (
	_ externglib.Objector = (*ScaleButtonAccessible)(nil)
)

func wrapScaleButtonAccessible(obj *externglib.Object) *ScaleButtonAccessible {
	return &ScaleButtonAccessible{
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
		Object: obj,
		Value: atk.Value{
			Object: obj,
		},
	}
}

func marshalScaleButtonAccessibler(p uintptr) (interface{}, error) {
	return wrapScaleButtonAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
