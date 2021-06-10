// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_renderer_cell_accessible_get_type()), F: marshalRendererCellAccessible},
	})
}

type RendererCellAccessible interface {
	CellAccessible
}

// rendererCellAccessible implements the RendererCellAccessible interface.
type rendererCellAccessible struct {
	CellAccessible
}

var _ RendererCellAccessible = (*rendererCellAccessible)(nil)

// WrapRendererCellAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapRendererCellAccessible(obj *externglib.Object) RendererCellAccessible {
	return RendererCellAccessible{
		CellAccessible: WrapCellAccessible(obj),
	}
}

func marshalRendererCellAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRendererCellAccessible(obj), nil
}
