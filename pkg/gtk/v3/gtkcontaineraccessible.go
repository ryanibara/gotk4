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
		{T: externglib.Type(C.gtk_container_accessible_get_type()), F: marshalContainerAccessibler},
	})
}

type ContainerAccessible struct {
	WidgetAccessible
}

func wrapContainerAccessible(obj *externglib.Object) *ContainerAccessible {
	return &ContainerAccessible{
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
	}
}

func marshalContainerAccessibler(p uintptr) (interface{}, error) {
	return wrapContainerAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
