// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtktextcellaccessible.go.
var GTypeTextCellAccessible = coreglib.Type(C.gtk_text_cell_accessible_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeTextCellAccessible, F: marshalTextCellAccessible},
	})
}

// TextCellAccessibleOverrider contains methods that are overridable.
type TextCellAccessibleOverrider interface {
}

type TextCellAccessible struct {
	_ [0]func() // equal guard
	RendererCellAccessible

	*coreglib.Object
	atk.ObjectClass
	atk.Text
}

var (
	_ coreglib.Objector = (*TextCellAccessible)(nil)
)

func classInitTextCellAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTextCellAccessible(obj *coreglib.Object) *TextCellAccessible {
	return &TextCellAccessible{
		RendererCellAccessible: RendererCellAccessible{
			CellAccessible: CellAccessible{
				Accessible: Accessible{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
				Object: obj,
				Action: atk.Action{
					Object: obj,
				},
				Component: atk.Component{
					Object: obj,
				},
				ObjectClass: atk.ObjectClass{
					Object: obj,
				},
				TableCell: atk.TableCell{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
			},
		},
		Object: obj,
		ObjectClass: atk.ObjectClass{
			Object: obj,
		},
		Text: atk.Text{
			Object: obj,
		},
	}
}

func marshalTextCellAccessible(p uintptr) (interface{}, error) {
	return wrapTextCellAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
