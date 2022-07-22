// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeTextViewAccessible = coreglib.Type(C.gtk_text_view_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTextViewAccessible, F: marshalTextViewAccessible},
	})
}

// TextViewAccessibleOverrider contains methods that are overridable.
type TextViewAccessibleOverrider interface {
}

type TextViewAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	*coreglib.Object
	atk.EditableText
	atk.StreamableContent
	atk.Text
}

var (
	_ coreglib.Objector = (*TextViewAccessible)(nil)
)

func classInitTextViewAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTextViewAccessible(obj *coreglib.Object) *TextViewAccessible {
	return &TextViewAccessible{
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
		EditableText: atk.EditableText{
			Object: obj,
		},
		StreamableContent: atk.StreamableContent{
			Object: obj,
		},
		Text: atk.Text{
			Object: obj,
		},
	}
}

func marshalTextViewAccessible(p uintptr) (interface{}, error) {
	return wrapTextViewAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
