// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_icon_view_accessible_get_type()), F: marshalIconViewAccessibler},
	})
}

// IconViewAccessibler describes IconViewAccessible's methods.
type IconViewAccessibler interface {
	privateIconViewAccessible()
}

type IconViewAccessible struct {
	ContainerAccessible

	atk.Selection
}

var (
	_ IconViewAccessibler = (*IconViewAccessible)(nil)
	_ gextras.Nativer     = (*IconViewAccessible)(nil)
)

func wrapIconViewAccessible(obj *externglib.Object) IconViewAccessibler {
	return &IconViewAccessible{
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
		Selection: atk.Selection{
			Object: obj,
		},
	}
}

func marshalIconViewAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapIconViewAccessible(obj), nil
}

func (*IconViewAccessible) privateIconViewAccessible() {}
