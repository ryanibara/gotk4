// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_spin_get_type()), F: marshalCellRendererSpinner},
	})
}

// CellRendererSpinner describes CellRendererSpin's methods.
type CellRendererSpinner interface {
	privateCellRendererSpin()
}

// CellRendererSpin renders a spin button in a cell
//
// CellRendererSpin renders text in a cell like CellRendererText from which it
// is derived. But while CellRendererText offers a simple entry to edit the
// text, CellRendererSpin offers a SpinButton widget. Of course, that means that
// the text has to be parseable as a floating point number.
//
// The range of the spinbutton is taken from the adjustment property of the cell
// renderer, which can be set explicitly or mapped to a column in the tree
// model, like all properties of cell renders. CellRendererSpin also has
// properties for the CellRendererSpin:climb-rate and the number of
// CellRendererSpin:digits to display. Other SpinButton properties can be set in
// a handler for the CellRenderer::editing-started signal.
//
// The CellRendererSpin cell renderer was added in GTK 2.10.
type CellRendererSpin struct {
	CellRendererText
}

var (
	_ CellRendererSpinner = (*CellRendererSpin)(nil)
	_ gextras.Nativer     = (*CellRendererSpin)(nil)
)

func wrapCellRendererSpin(obj *externglib.Object) CellRendererSpinner {
	return &CellRendererSpin{
		CellRendererText: CellRendererText{
			CellRenderer: CellRenderer{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalCellRendererSpinner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCellRendererSpin(obj), nil
}

// NewCellRendererSpin creates a new CellRendererSpin.
func NewCellRendererSpin() *CellRendererSpin {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_spin_new()

	var _cellRendererSpin *CellRendererSpin // out

	_cellRendererSpin = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*CellRendererSpin)

	return _cellRendererSpin
}

func (*CellRendererSpin) privateCellRendererSpin() {}
