// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeCellView returns the GType for the type CellView.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCellView() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "CellView").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalCellView)
	return gtype
}

// CellViewOverrider contains methods that are overridable.
type CellViewOverrider interface {
}

// CellView displays a single row of a TreeModel using a CellArea and
// CellAreaContext. A CellAreaContext can be provided to the CellView at
// construction time in order to keep the cellview in context of a group of cell
// views, this ensures that the renderers displayed will be properly aligned
// with eachother (like the aligned cells in the menus of ComboBox).
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
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	CellLayout
	Orientable
}

var (
	_ Widgetter         = (*CellView)(nil)
	_ coreglib.Objector = (*CellView)(nil)
)

func classInitCellViewer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapCellView(obj *coreglib.Object) *CellView {
	return &CellView{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
		},
		Object: obj,
		CellLayout: CellLayout{
			Object: obj,
		},
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalCellView(p uintptr) (interface{}, error) {
	return wrapCellView(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCellView creates a new CellView widget.
//
// The function returns the following values:
//
//    - cellView: newly created CellView widget.
//
func NewCellView() *CellView {
	_info := girepository.MustFind("Gtk", "CellView")
	_gret := _info.InvokeClassMethod("new_CellView", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _cellView *CellView // out

	_cellView = wrapCellView(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// NewCellViewWithContext creates a new CellView widget with a specific CellArea
// to layout cells and a specific CellAreaContext.
//
// Specifying the same context for a handfull of cells lets the underlying area
// synchronize the geometry for those cells, in this way alignments with
// cellviews for other rows are possible.
//
// The function takes the following parameters:
//
//    - area to layout cells.
//    - context in which to calculate cell geometry.
//
// The function returns the following values:
//
//    - cellView: newly created CellView widget.
//
func NewCellViewWithContext(area CellAreaer, context *CellAreaContext) *CellView {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(area).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_info := girepository.MustFind("Gtk", "CellView")
	_gret := _info.InvokeClassMethod("new_CellView_with_context", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(area)
	runtime.KeepAlive(context)

	var _cellView *CellView // out

	_cellView = wrapCellView(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// NewCellViewWithMarkup creates a new CellView widget, adds a CellRendererText
// to it, and makes it show markup. The text can be marked up with the [Pango
// text markup language][PangoMarkupFormat].
//
// The function takes the following parameters:
//
//    - markup: text to display in the cell view.
//
// The function returns the following values:
//
//    - cellView: newly created CellView widget.
//
func NewCellViewWithMarkup(markup string) *CellView {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(markup)))
	defer C.free(unsafe.Pointer(_args[0]))

	_info := girepository.MustFind("Gtk", "CellView")
	_gret := _info.InvokeClassMethod("new_CellView_with_markup", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(markup)

	var _cellView *CellView // out

	_cellView = wrapCellView(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// NewCellViewWithPixbuf creates a new CellView widget, adds a
// CellRendererPixbuf to it, and makes it show pixbuf.
//
// The function takes the following parameters:
//
//    - pixbuf: image to display in the cell view.
//
// The function returns the following values:
//
//    - cellView: newly created CellView widget.
//
func NewCellViewWithPixbuf(pixbuf *gdkpixbuf.Pixbuf) *CellView {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))

	_info := girepository.MustFind("Gtk", "CellView")
	_gret := _info.InvokeClassMethod("new_CellView_with_pixbuf", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(pixbuf)

	var _cellView *CellView // out

	_cellView = wrapCellView(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// NewCellViewWithText creates a new CellView widget, adds a CellRendererText to
// it, and makes it show text.
//
// The function takes the following parameters:
//
//    - text to display in the cell view.
//
// The function returns the following values:
//
//    - cellView: newly created CellView widget.
//
func NewCellViewWithText(text string) *CellView {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_args[0]))

	_info := girepository.MustFind("Gtk", "CellView")
	_gret := _info.InvokeClassMethod("new_CellView_with_text", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(text)

	var _cellView *CellView // out

	_cellView = wrapCellView(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// DisplayedRow returns a TreePath referring to the currently displayed row. If
// no row is currently displayed, NULL is returned.
//
// The function returns the following values:
//
//    - treePath (optional): currently displayed row or NULL.
//
func (cellView *CellView) DisplayedRow() *TreePath {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))

	_info := girepository.MustFind("Gtk", "CellView")
	_gret := _info.InvokeClassMethod("get_displayed_row", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cellView)

	var _treePath *TreePath // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_treePath = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_treePath)),
			func(intern *struct{ C unsafe.Pointer }) {
				{
					args := [1]girepository.Argument{(*C.void)(intern.C)}
					girepository.MustFind("Gtk", "TreePath").InvokeRecordMethod("free", args[:], nil)
				}
			},
		)
	}

	return _treePath
}

// DrawSensitive gets whether cell_view is configured to draw all of its cells
// in a sensitive state.
//
// The function returns the following values:
//
//    - ok: whether cell_view draws all of its cells in a sensitive state.
//
func (cellView *CellView) DrawSensitive() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))

	_info := girepository.MustFind("Gtk", "CellView")
	_gret := _info.InvokeClassMethod("get_draw_sensitive", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cellView)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))

	_info := girepository.MustFind("Gtk", "CellView")
	_gret := _info.InvokeClassMethod("get_fit_model", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cellView)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Model returns the model for cell_view. If no model is used NULL is returned.
//
// The function returns the following values:
//
//    - treeModel (optional) used or NULL.
//
func (cellView *CellView) Model() *TreeModel {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))

	_info := girepository.MustFind("Gtk", "CellView")
	_gret := _info.InvokeClassMethod("get_model", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cellView)

	var _treeModel *TreeModel // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_treeModel = wrapTreeModel(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _treeModel
}

// SizeOfRow sets requisition to the size needed by cell_view to display the
// model row pointed to by path.
//
// Deprecated: Combo box formerly used this to calculate the sizes for
// cellviews, now you can achieve this by either using the CellView:fit-model
// property or by setting the currently displayed row of the CellView and using
// gtk_widget_get_preferred_size().
//
// The function takes the following parameters:
//
//    - path: TreePath.
//
// The function returns the following values:
//
//    - requisition: return location for the size.
//    - ok: TRUE.
//
func (cellView *CellView) SizeOfRow(path *TreePath) (*Requisition, bool) {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(path)))

	_info := girepository.MustFind("Gtk", "CellView")
	_gret := _info.InvokeClassMethod("get_size_of_row", _args[:], _outs[:])
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cellView)
	runtime.KeepAlive(path)

	var _requisition *Requisition // out
	var _ok bool                  // out

	_requisition = (*Requisition)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _requisition, _ok
}

// SetBackgroundColor sets the background color of view.
//
// Deprecated: Use gtk_cell_view_set_background_rgba() instead.
//
// The function takes the following parameters:
//
//    - color: new background color.
//
func (cellView *CellView) SetBackgroundColor(color *gdk.Color) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(color)))

	_info := girepository.MustFind("Gtk", "CellView")
	_info.InvokeClassMethod("set_background_color", _args[:], nil)

	runtime.KeepAlive(cellView)
	runtime.KeepAlive(color)
}

// SetBackgroundRGBA sets the background color of cell_view.
//
// The function takes the following parameters:
//
//    - rgba: new background color.
//
func (cellView *CellView) SetBackgroundRGBA(rgba *gdk.RGBA) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(rgba)))

	_info := girepository.MustFind("Gtk", "CellView")
	_info.InvokeClassMethod("set_background_rgba", _args[:], nil)

	runtime.KeepAlive(cellView)
	runtime.KeepAlive(rgba)
}

// SetDisplayedRow sets the row of the model that is currently displayed by the
// CellView. If the path is unset, then the contents of the cellview “stick” at
// their last value; this is not normally a desired result, but may be a needed
// intermediate state if say, the model for the CellView becomes temporarily
// empty.
//
// The function takes the following parameters:
//
//    - path (optional) or NULL to unset.
//
func (cellView *CellView) SetDisplayedRow(path *TreePath) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))
	if path != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(path)))
	}

	_info := girepository.MustFind("Gtk", "CellView")
	_info.InvokeClassMethod("set_displayed_row", _args[:], nil)

	runtime.KeepAlive(cellView)
	runtime.KeepAlive(path)
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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))
	if drawSensitive {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "CellView")
	_info.InvokeClassMethod("set_draw_sensitive", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))
	if fitModel {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "CellView")
	_info.InvokeClassMethod("set_fit_model", _args[:], nil)

	runtime.KeepAlive(cellView)
	runtime.KeepAlive(fitModel)
}

// SetModel sets the model for cell_view. If cell_view already has a model set,
// it will remove it before setting the new model. If model is NULL, then it
// will unset the old model.
//
// The function takes the following parameters:
//
//    - model (optional): TreeModel.
//
func (cellView *CellView) SetModel(model TreeModeller) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))
	if model != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	_info := girepository.MustFind("Gtk", "CellView")
	_info.InvokeClassMethod("set_model", _args[:], nil)

	runtime.KeepAlive(cellView)
	runtime.KeepAlive(model)
}
