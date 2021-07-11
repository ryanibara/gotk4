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
		{T: externglib.Type(C.gtk_label_accessible_get_type()), F: marshalLabelAccessibler},
	})
}

// LabelAccessibler describes LabelAccessible's methods.
type LabelAccessibler interface {
	privateLabelAccessible()
}

type LabelAccessible struct {
	WidgetAccessible

	atk.Hypertext
	atk.Text
}

var (
	_ LabelAccessibler = (*LabelAccessible)(nil)
	_ gextras.Nativer  = (*LabelAccessible)(nil)
)

func wrapLabelAccessible(obj *externglib.Object) LabelAccessibler {
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
	}
}

func marshalLabelAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLabelAccessible(obj), nil
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *LabelAccessible) Native() uintptr {
	return v.WidgetAccessible.Accessible.ObjectClass.Object.Native()
}

func (*LabelAccessible) privateLabelAccessible() {}
