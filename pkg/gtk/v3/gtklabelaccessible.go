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
		{T: externglib.Type(C.gtk_label_accessible_get_type()), F: marshalLabelAccessibler},
	})
}

type LabelAccessible struct {
	WidgetAccessible

	atk.Hypertext
	atk.Text
	*externglib.Object
}

func wrapLabelAccessible(obj *externglib.Object) *LabelAccessible {
	return &LabelAccessible{
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
		Hypertext: atk.Hypertext{
			Object: obj,
		},
		Text: atk.Text{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalLabelAccessibler(p uintptr) (interface{}, error) {
	return wrapLabelAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (*LabelAccessible) privateLabelAccessible() {}
