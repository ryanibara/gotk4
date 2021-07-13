// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_view_get_type()), F: marshalCellViewer},
	})
}

// CellViewer describes CellView's methods.
type CellViewer interface {
	// DisplayedRow returns a TreePath referring to the currently displayed row.
	DisplayedRow() *TreePath
	// DrawSensitive gets whether cell_view is configured to draw all of its
	// cells in a sensitive state.
	DrawSensitive() bool
	// FitModel gets whether cell_view is configured to request space to fit the
	// entire TreeModel.
	FitModel() bool
	// Model returns the model for cell_view.
	Model() *TreeModel
	// SetDisplayedRow sets the row of the model that is currently displayed by
	// the CellView.
	SetDisplayedRow(path *TreePath)
	// SetDrawSensitive sets whether cell_view should draw all of its cells in a
	// sensitive state, this is used by ComboBox menus to ensure that rows with
	// insensitive cells that contain children appear sensitive in the parent
	// menu item.
	SetDrawSensitive(drawSensitive bool)
	// SetFitModel sets whether cell_view should request space to fit the entire
	// TreeModel.
	SetFitModel(fitModel bool)
	// SetModel sets the model for cell_view.
	SetModel(model TreeModeler)
}

// CellView: widget displaying a single row of a GtkTreeModel
//
// A CellView displays a single row of a TreeModel using a CellArea and
// CellAreaContext. A CellAreaContext can be provided to the CellView at
// construction time in order to keep the cellview in context of a group of cell
// views, this ensures that the renderers displayed will be properly aligned
// with each other (like the aligned cells in the menus of ComboBox).
//
// CellView is Orientable in order to decide in which orientation the underlying
// CellAreaContext should be allocated. Taking the ComboBox menu as an example,
// cellviews should be oriented horizontally if the menus are listed
// top-to-bottom and thus all share the same width but may have separate
// individual heights (left-to-right menus should be allocated vertically since
// they all share the same height but may have variable widths).
//
//
// CSS nodes
//
// GtkCellView has a single CSS node with name cellview.
type CellView struct {
	Widget

	CellLayout
	Orientable
}

var (
	_ CellViewer      = (*CellView)(nil)
	_ gextras.Nativer = (*CellView)(nil)
)

func wrapCellView(obj *externglib.Object) *CellView {
	return &CellView{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		CellLayout: CellLayout{
			Object: obj,
		},
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalCellViewer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCellView(obj), nil
}

// NewCellView creates a new CellView widget.
func NewCellView() *CellView {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_cell_view_new()

	var _cellView *CellView // out

	_cellView = wrapCellView(externglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// NewCellViewWithContext creates a new CellView widget with a specific CellArea
// to layout cells and a specific CellAreaContext.
//
// Specifying the same context for a handful of cells lets the underlying area
// synchronize the geometry for those cells, in this way alignments with
// cellviews for other rows are possible.
func NewCellViewWithContext(area CellAreaer, context *CellAreaContext) *CellView {
	var _arg1 *C.GtkCellArea        // out
	var _arg2 *C.GtkCellAreaContext // out
	var _cret *C.GtkWidget          // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer((area).(gextras.Nativer).Native()))
	_arg2 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_cell_view_new_with_context(_arg1, _arg2)

	var _cellView *CellView // out

	_cellView = wrapCellView(externglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// NewCellViewWithMarkup creates a new CellView widget, adds a CellRendererText
// to it, and makes it show markup. The text can be marked up with the [Pango
// text markup language][PangoMarkupFormat].
func NewCellViewWithMarkup(markup string) *CellView {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(markup)))

	_cret = C.gtk_cell_view_new_with_markup(_arg1)

	var _cellView *CellView // out

	_cellView = wrapCellView(externglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// NewCellViewWithText creates a new CellView widget, adds a CellRendererText to
// it, and makes it show text.
func NewCellViewWithText(text string) *CellView {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))

	_cret = C.gtk_cell_view_new_with_text(_arg1)

	var _cellView *CellView // out

	_cellView = wrapCellView(externglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// NewCellViewWithTexture creates a new CellView widget, adds a
// CellRendererPixbuf to it, and makes it show texture.
func NewCellViewWithTexture(texture gdk.Texturer) *CellView {
	var _arg1 *C.GdkTexture // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GdkTexture)(unsafe.Pointer((texture).(gextras.Nativer).Native()))

	_cret = C.gtk_cell_view_new_with_texture(_arg1)

	var _cellView *CellView // out

	_cellView = wrapCellView(externglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *CellView) Native() uintptr {
	return v.Widget.InitiallyUnowned.Object.Native()
}

// DisplayedRow returns a TreePath referring to the currently displayed row. If
// no row is currently displayed, NULL is returned.
func (cellView *CellView) DisplayedRow() *TreePath {
	var _arg0 *C.GtkCellView // out
	var _cret *C.GtkTreePath // in

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(cellView.Native()))

	_cret = C.gtk_cell_view_get_displayed_row(_arg0)

	var _treePath *TreePath // out

	_treePath = (*TreePath)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_treePath, func(v *TreePath) {
		C.gtk_tree_path_free((*C.GtkTreePath)(unsafe.Pointer(v)))
	})

	return _treePath
}

// DrawSensitive gets whether cell_view is configured to draw all of its cells
// in a sensitive state.
func (cellView *CellView) DrawSensitive() bool {
	var _arg0 *C.GtkCellView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(cellView.Native()))

	_cret = C.gtk_cell_view_get_draw_sensitive(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FitModel gets whether cell_view is configured to request space to fit the
// entire TreeModel.
func (cellView *CellView) FitModel() bool {
	var _arg0 *C.GtkCellView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(cellView.Native()))

	_cret = C.gtk_cell_view_get_fit_model(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Model returns the model for cell_view. If no model is used NULL is returned.
func (cellView *CellView) Model() *TreeModel {
	var _arg0 *C.GtkCellView  // out
	var _cret *C.GtkTreeModel // in

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(cellView.Native()))

	_cret = C.gtk_cell_view_get_model(_arg0)

	var _treeModel *TreeModel // out

	_treeModel = wrapTreeModel(externglib.Take(unsafe.Pointer(_cret)))

	return _treeModel
}

// SetDisplayedRow sets the row of the model that is currently displayed by the
// CellView. If the path is unset, then the contents of the cellview “stick” at
// their last value; this is not normally a desired result, but may be a needed
// intermediate state if say, the model for the CellView becomes temporarily
// empty.
func (cellView *CellView) SetDisplayedRow(path *TreePath) {
	var _arg0 *C.GtkCellView // out
	var _arg1 *C.GtkTreePath // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(cellView.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_cell_view_set_displayed_row(_arg0, _arg1)
}

// SetDrawSensitive sets whether cell_view should draw all of its cells in a
// sensitive state, this is used by ComboBox menus to ensure that rows with
// insensitive cells that contain children appear sensitive in the parent menu
// item.
func (cellView *CellView) SetDrawSensitive(drawSensitive bool) {
	var _arg0 *C.GtkCellView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(cellView.Native()))
	if drawSensitive {
		_arg1 = C.TRUE
	}

	C.gtk_cell_view_set_draw_sensitive(_arg0, _arg1)
}

// SetFitModel sets whether cell_view should request space to fit the entire
// TreeModel.
//
// This is used by ComboBox to ensure that the cell view displayed on the combo
// box’s button always gets enough space and does not resize when selection
// changes.
func (cellView *CellView) SetFitModel(fitModel bool) {
	var _arg0 *C.GtkCellView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(cellView.Native()))
	if fitModel {
		_arg1 = C.TRUE
	}

	C.gtk_cell_view_set_fit_model(_arg0, _arg1)
}

// SetModel sets the model for cell_view. If cell_view already has a model set,
// it will remove it before setting the new model. If model is NULL, then it
// will unset the old model.
func (cellView *CellView) SetModel(model TreeModeler) {
	var _arg0 *C.GtkCellView  // out
	var _arg1 *C.GtkTreeModel // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(cellView.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer((model).(gextras.Nativer).Native()))

	C.gtk_cell_view_set_model(_arg0, _arg1)
}
