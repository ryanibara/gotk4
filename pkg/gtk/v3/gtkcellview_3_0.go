// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// DrawSensitive gets whether cell_view is configured to draw all of its cells
// in a sensitive state.
//
// The function returns the following values:
//
//    - ok: whether cell_view draws all of its cells in a sensitive state.
//
func (cellView *CellView) DrawSensitive() bool {
	var _arg0 *C.GtkCellView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))

	_cret = C.gtk_cell_view_get_draw_sensitive(_arg0)
	runtime.KeepAlive(cellView)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FitModel gets whether cell_view is configured to request space to fit the
// entire TreeModel.
//
// The function returns the following values:
//
//    - ok: whether cell_view requests space to fit the entire TreeModel.
//
func (cellView *CellView) FitModel() bool {
	var _arg0 *C.GtkCellView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))

	_cret = C.gtk_cell_view_get_fit_model(_arg0)
	runtime.KeepAlive(cellView)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetBackgroundRGBA sets the background color of cell_view.
//
// The function takes the following parameters:
//
//    - rgba: new background color.
//
func (cellView *CellView) SetBackgroundRGBA(rgba *gdk.RGBA) {
	var _arg0 *C.GtkCellView // out
	var _arg1 *C.GdkRGBA     // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))
	_arg1 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(rgba)))

	C.gtk_cell_view_set_background_rgba(_arg0, _arg1)
	runtime.KeepAlive(cellView)
	runtime.KeepAlive(rgba)
}

// SetDrawSensitive sets whether cell_view should draw all of its cells in a
// sensitive state, this is used by ComboBox menus to ensure that rows with
// insensitive cells that contain children appear sensitive in the parent menu
// item.
//
// The function takes the following parameters:
//
//    - drawSensitive: whether to draw all cells in a sensitive state.
//
func (cellView *CellView) SetDrawSensitive(drawSensitive bool) {
	var _arg0 *C.GtkCellView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))
	if drawSensitive {
		_arg1 = C.TRUE
	}

	C.gtk_cell_view_set_draw_sensitive(_arg0, _arg1)
	runtime.KeepAlive(cellView)
	runtime.KeepAlive(drawSensitive)
}

// SetFitModel sets whether cell_view should request space to fit the entire
// TreeModel.
//
// This is used by ComboBox to ensure that the cell view displayed on the combo
// box’s button always gets enough space and does not resize when selection
// changes.
//
// The function takes the following parameters:
//
//    - fitModel: whether cell_view should request space for the whole model.
//
func (cellView *CellView) SetFitModel(fitModel bool) {
	var _arg0 *C.GtkCellView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))
	if fitModel {
		_arg1 = C.TRUE
	}

	C.gtk_cell_view_set_fit_model(_arg0, _arg1)
	runtime.KeepAlive(cellView)
	runtime.KeepAlive(fitModel)
}
