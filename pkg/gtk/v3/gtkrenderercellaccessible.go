// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
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
		{T: externglib.Type(C.gtk_renderer_cell_accessible_get_type()), F: marshalRendererCellAccessibler},
	})
}

// RendererCellAccessibleOverrider contains methods that are overridable.
type RendererCellAccessibleOverrider interface {
}

type RendererCellAccessible struct {
	_ [0]func() // equal guard
	CellAccessible
}

var (
	_ externglib.Objector = (*RendererCellAccessible)(nil)
)

func classInitRendererCellAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRendererCellAccessible(obj *externglib.Object) *RendererCellAccessible {
	return &RendererCellAccessible{
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
	}
}

func marshalRendererCellAccessibler(p uintptr) (interface{}, error) {
	return wrapRendererCellAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func NewRendererCellAccessible(renderer CellRendererer) *RendererCellAccessible {
	var _arg1 *C.GtkCellRenderer // out
	var _cret *C.AtkObject       // in

	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))

	_cret = C.gtk_renderer_cell_accessible_new(_arg1)
	runtime.KeepAlive(renderer)

	var _rendererCellAccessible *RendererCellAccessible // out

	_rendererCellAccessible = wrapRendererCellAccessible(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _rendererCellAccessible
}
