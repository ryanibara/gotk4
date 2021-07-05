// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_cell_renderer_spinner_get_type()), F: marshalCellRendererSpinner},
	})
}

// CellRendererSpinner renders a spinning animation in a cell, very similar to
// Spinner. It can often be used as an alternative to a CellRendererProgress for
// displaying indefinite activity, instead of actual progress.
//
// To start the animation in a cell, set the CellRendererSpinner:active property
// to true and increment the CellRendererSpinner:pulse property at regular
// intervals. The usual way to set the cell renderer properties for each cell is
// to bind them to columns in your tree model using e.g.
// gtk_tree_view_column_add_attribute().
type CellRendererSpinner interface {
	CellRenderer
}

// cellRendererSpinner implements the CellRendererSpinner class.
type cellRendererSpinner struct {
	CellRenderer
}

// WrapCellRendererSpinner wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellRendererSpinner(obj *externglib.Object) CellRendererSpinner {
	return cellRendererSpinner{
		CellRenderer: WrapCellRenderer(obj),
	}
}

func marshalCellRendererSpinner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellRendererSpinner(obj), nil
}

// NewCellRendererSpinner returns a new cell renderer which will show a spinner
// to indicate activity.
func NewCellRendererSpinner() CellRendererSpinner {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_spinner_new()

	var _cellRendererSpinner CellRendererSpinner // out

	_cellRendererSpinner = WrapCellRendererSpinner(externglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererSpinner
}
