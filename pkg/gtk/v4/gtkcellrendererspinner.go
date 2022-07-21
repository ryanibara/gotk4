// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GTypeCellRendererSpinner returns the GType for the type CellRendererSpinner.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCellRendererSpinner() coreglib.Type {
	gtype := coreglib.Type(C.gtk_cell_renderer_spinner_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalCellRendererSpinner)
	return gtype
}

// CellRendererSpinner renders a spinning animation in a cell
//
// GtkCellRendererSpinner renders a spinning animation in a cell, very similar
// to Spinner. It can often be used as an alternative to a CellRendererProgress
// for displaying indefinite activity, instead of actual progress.
//
// To start the animation in a cell, set the CellRendererSpinner:active property
// to TRUE and increment the CellRendererSpinner:pulse property at regular
// intervals. The usual way to set the cell renderer properties for each cell is
// to bind them to columns in your tree model using e.g.
// gtk_tree_view_column_add_attribute().
type CellRendererSpinner struct {
	_ [0]func() // equal guard
	CellRenderer
}

var (
	_ CellRendererer = (*CellRendererSpinner)(nil)
)

func wrapCellRendererSpinner(obj *coreglib.Object) *CellRendererSpinner {
	return &CellRendererSpinner{
		CellRenderer: CellRenderer{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
		},
	}
}

func marshalCellRendererSpinner(p uintptr) (interface{}, error) {
	return wrapCellRendererSpinner(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCellRendererSpinner returns a new cell renderer which will show a spinner
// to indicate activity.
//
// The function returns the following values:
//
//    - cellRendererSpinner: new CellRenderer.
//
func NewCellRendererSpinner() *CellRendererSpinner {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_spinner_new()

	var _cellRendererSpinner *CellRendererSpinner // out

	_cellRendererSpinner = wrapCellRendererSpinner(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererSpinner
}
