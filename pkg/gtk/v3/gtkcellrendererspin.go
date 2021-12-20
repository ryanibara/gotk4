// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_spin_get_type()), F: marshalCellRendererSpinner},
	})
}

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
// The CellRendererSpin cell renderer was added in GTK+ 2.10.
type CellRendererSpin struct {
	CellRendererText

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ CellRendererer = (*CellRendererSpin)(nil)
)

func wrapCellRendererSpin(obj *externglib.Object) *CellRendererSpin {
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
	return wrapCellRendererSpin(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCellRendererSpin creates a new CellRendererSpin.
//
// The function returns the following values:
//
//    - cellRendererSpin: new CellRendererSpin.
//
func NewCellRendererSpin() *CellRendererSpin {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_spin_new()

	var _cellRendererSpin *CellRendererSpin // out

	_cellRendererSpin = wrapCellRendererSpin(externglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererSpin
}
