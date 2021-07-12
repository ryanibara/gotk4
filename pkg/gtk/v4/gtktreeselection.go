// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
// void gotk4_TreeSelectionForeachFunc(GtkTreeModel*, GtkTreePath*, GtkTreeIter*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_selection_get_type()), F: marshalTreeSelectioner},
	})
}

// TreeSelectionForeachFunc: function used by
// gtk_tree_selection_selected_foreach() to map all selected rows. It will be
// called on every selected row in the view.
type TreeSelectionForeachFunc func(model *TreeModel, path *TreePath, iter *TreeIter, data cgo.Handle)

//export gotk4_TreeSelectionForeachFunc
func gotk4_TreeSelectionForeachFunc(arg0 *C.GtkTreeModel, arg1 *C.GtkTreePath, arg2 *C.GtkTreeIter, arg3 C.gpointer) {
	v := gbox.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var model *TreeModel // out
	var path *TreePath   // out
	var iter *TreeIter   // out
	var data cgo.Handle  // out

	model = wrapTreeModel(externglib.Take(unsafe.Pointer(arg0)))
	path = (*TreePath)(unsafe.Pointer(arg1))
	runtime.SetFinalizer(path, func(v *TreePath) {
		C.gtk_tree_path_free((*C.GtkTreePath)(unsafe.Pointer(v)))
	})
	iter = (*TreeIter)(unsafe.Pointer(arg2))
	runtime.SetFinalizer(iter, func(v *TreeIter) {
		C.gtk_tree_iter_free((*C.GtkTreeIter)(unsafe.Pointer(v)))
	})
	data = (cgo.Handle)(unsafe.Pointer(arg3))

	fn := v.(TreeSelectionForeachFunc)
	fn(model, path, iter, data)
}

// TreeSelectionFunc: function used by gtk_tree_selection_set_select_function()
// to filter whether or not a row may be selected. It is called whenever a row's
// state might change.
//
// A return value of true indicates to @selection that it is okay to change the
// selection.
type TreeSelectionFunc func(selection *TreeSelection, model *TreeModel, path *TreePath, pathCurrentlySelected bool, data cgo.Handle) (ok bool)

//export gotk4_TreeSelectionFunc
func gotk4_TreeSelectionFunc(arg0 *C.GtkTreeSelection, arg1 *C.GtkTreeModel, arg2 *C.GtkTreePath, arg3 C.gboolean, arg4 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	var selection *TreeSelection   // out
	var model *TreeModel           // out
	var path *TreePath             // out
	var pathCurrentlySelected bool // out
	var data cgo.Handle            // out

	selection = wrapTreeSelection(externglib.Take(unsafe.Pointer(arg0)))
	model = wrapTreeModel(externglib.Take(unsafe.Pointer(arg1)))
	path = (*TreePath)(unsafe.Pointer(arg2))
	runtime.SetFinalizer(path, func(v *TreePath) {
		C.gtk_tree_path_free((*C.GtkTreePath)(unsafe.Pointer(v)))
	})
	if arg3 != 0 {
		pathCurrentlySelected = true
	}
	data = (cgo.Handle)(unsafe.Pointer(arg4))

	fn := v.(TreeSelectionFunc)
	ok := fn(selection, model, path, pathCurrentlySelected, data)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// TreeSelectioner describes TreeSelection's methods.
type TreeSelectioner interface {
	// CountSelectedRows returns the number of rows that have been selected in
	// @tree.
	CountSelectedRows() int
	// Mode gets the selection mode for @selection.
	Mode() SelectionMode
	// Selected sets @iter to the currently selected node if @selection is set
	// to K_SELECTION_SINGLE or K_SELECTION_BROWSE.
	Selected() (*TreeModel, TreeIter, bool)
	// TreeView returns the tree view associated with @selection.
	TreeView() *TreeView
	// IterIsSelected returns true if the row at @iter is currently selected.
	IterIsSelected(iter *TreeIter) bool
	// PathIsSelected returns true if the row pointed to by @path is currently
	// selected.
	PathIsSelected(path *TreePath) bool
	// SelectAll selects all the nodes.
	SelectAll()
	// SelectIter selects the specified iterator.
	SelectIter(iter *TreeIter)
	// SelectPath: select the row at @path.
	SelectPath(path *TreePath)
	// SelectRange selects a range of nodes, determined by @start_path and
	// @end_path inclusive.
	SelectRange(startPath *TreePath, endPath *TreePath)
	// SelectedForeach calls a function for each selected node.
	SelectedForeach(fn TreeSelectionForeachFunc)
	// SetMode sets the selection mode of the @selection.
	SetMode(typ SelectionMode)
	// UnselectAll unselects all the nodes.
	UnselectAll()
	// UnselectIter unselects the specified iterator.
	UnselectIter(iter *TreeIter)
	// UnselectPath unselects the row at @path.
	UnselectPath(path *TreePath)
	// UnselectRange unselects a range of nodes, determined by @start_path and
	// @end_path inclusive.
	UnselectRange(startPath *TreePath, endPath *TreePath)
}

// TreeSelection: selection object for GtkTreeView
//
// The TreeSelection object is a helper object to manage the selection for a
// TreeView widget. The TreeSelection object is automatically created when a new
// TreeView widget is created, and cannot exist independently of this widget.
// The primary reason the TreeSelection objects exists is for cleanliness of
// code and API. That is, there is no conceptual reason all these functions
// could not be methods on the TreeView widget instead of a separate function.
//
// The TreeSelection object is gotten from a TreeView by calling
// gtk_tree_view_get_selection(). It can be manipulated to check the selection
// status of the tree, as well as select and deselect individual rows. Selection
// is done completely view side. As a result, multiple views of the same model
// can have completely different selections. Additionally, you cannot change the
// selection of a row on the model that is not currently displayed by the view
// without expanding its parents first.
//
// One of the important things to remember when monitoring the selection of a
// view is that the TreeSelection::changed signal is mostly a hint. That is, it
// may only emit one signal when a range of rows is selected. Additionally, it
// may on occasion emit a TreeSelection::changed signal when nothing has
// happened (mostly as a result of programmers calling select_row on an already
// selected row).
type TreeSelection struct {
	*externglib.Object
}

var (
	_ TreeSelectioner = (*TreeSelection)(nil)
	_ gextras.Nativer = (*TreeSelection)(nil)
)

func wrapTreeSelection(obj *externglib.Object) *TreeSelection {
	return &TreeSelection{
		Object: obj,
	}
}

func marshalTreeSelectioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTreeSelection(obj), nil
}

// CountSelectedRows returns the number of rows that have been selected in
// @tree.
func (selection *TreeSelection) CountSelectedRows() int {
	var _arg0 *C.GtkTreeSelection // out
	var _cret C.int               // in

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(selection.Native()))

	_cret = C.gtk_tree_selection_count_selected_rows(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Mode gets the selection mode for @selection. See
// gtk_tree_selection_set_mode().
func (selection *TreeSelection) Mode() SelectionMode {
	var _arg0 *C.GtkTreeSelection // out
	var _cret C.GtkSelectionMode  // in

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(selection.Native()))

	_cret = C.gtk_tree_selection_get_mode(_arg0)

	var _selectionMode SelectionMode // out

	_selectionMode = SelectionMode(_cret)

	return _selectionMode
}

// Selected sets @iter to the currently selected node if @selection is set to
// K_SELECTION_SINGLE or K_SELECTION_BROWSE. @iter may be NULL if you just want
// to test if @selection has any selected nodes. @model is filled with the
// current model as a convenience. This function will not work if you use
// @selection is K_SELECTION_MULTIPLE.
func (selection *TreeSelection) Selected() (*TreeModel, TreeIter, bool) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreeModel     // in
	var _iter TreeIter
	var _cret C.gboolean // in

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(selection.Native()))

	_cret = C.gtk_tree_selection_get_selected(_arg0, &_arg1, (*C.GtkTreeIter)(unsafe.Pointer(&_iter)))

	var _model *TreeModel // out

	var _ok bool // out

	_model = wrapTreeModel(externglib.Take(unsafe.Pointer(_arg1)))

	if _cret != 0 {
		_ok = true
	}

	return _model, _iter, _ok
}

// TreeView returns the tree view associated with @selection.
func (selection *TreeSelection) TreeView() *TreeView {
	var _arg0 *C.GtkTreeSelection // out
	var _cret *C.GtkTreeView      // in

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(selection.Native()))

	_cret = C.gtk_tree_selection_get_tree_view(_arg0)

	var _treeView *TreeView // out

	_treeView = wrapTreeView(externglib.Take(unsafe.Pointer(_cret)))

	return _treeView
}

// IterIsSelected returns true if the row at @iter is currently selected.
func (selection *TreeSelection) IterIsSelected(iter *TreeIter) bool {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreeIter      // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(selection.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))

	_cret = C.gtk_tree_selection_iter_is_selected(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PathIsSelected returns true if the row pointed to by @path is currently
// selected. If @path does not point to a valid location, false is returned
func (selection *TreeSelection) PathIsSelected(path *TreePath) bool {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreePath      // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(selection.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path))

	_cret = C.gtk_tree_selection_path_is_selected(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectAll selects all the nodes. @selection must be set to
// K_SELECTION_MULTIPLE mode.
func (selection *TreeSelection) SelectAll() {
	var _arg0 *C.GtkTreeSelection // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(selection.Native()))

	C.gtk_tree_selection_select_all(_arg0)
}

// SelectIter selects the specified iterator.
func (selection *TreeSelection) SelectIter(iter *TreeIter) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreeIter      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(selection.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.gtk_tree_selection_select_iter(_arg0, _arg1)
}

// SelectPath: select the row at @path.
func (selection *TreeSelection) SelectPath(path *TreePath) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreePath      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(selection.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_tree_selection_select_path(_arg0, _arg1)
}

// SelectRange selects a range of nodes, determined by @start_path and @end_path
// inclusive. @selection must be set to K_SELECTION_MULTIPLE mode.
func (selection *TreeSelection) SelectRange(startPath *TreePath, endPath *TreePath) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreePath      // out
	var _arg2 *C.GtkTreePath      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(selection.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(startPath))
	_arg2 = (*C.GtkTreePath)(unsafe.Pointer(endPath))

	C.gtk_tree_selection_select_range(_arg0, _arg1, _arg2)
}

// SelectedForeach calls a function for each selected node. Note that you cannot
// modify the tree or selection from within this function. As a result,
// gtk_tree_selection_get_selected_rows() might be more useful.
func (selection *TreeSelection) SelectedForeach(fn TreeSelectionForeachFunc) {
	var _arg0 *C.GtkTreeSelection           // out
	var _arg1 C.GtkTreeSelectionForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(selection.Native()))
	_arg1 = (*[0]byte)(C.gotk4_TreeSelectionForeachFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))

	C.gtk_tree_selection_selected_foreach(_arg0, _arg1, _arg2)
}

// SetMode sets the selection mode of the @selection. If the previous type was
// K_SELECTION_MULTIPLE, then the anchor is kept selected, if it was previously
// selected.
func (selection *TreeSelection) SetMode(typ SelectionMode) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 C.GtkSelectionMode  // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(selection.Native()))
	_arg1 = C.GtkSelectionMode(typ)

	C.gtk_tree_selection_set_mode(_arg0, _arg1)
}

// UnselectAll unselects all the nodes.
func (selection *TreeSelection) UnselectAll() {
	var _arg0 *C.GtkTreeSelection // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(selection.Native()))

	C.gtk_tree_selection_unselect_all(_arg0)
}

// UnselectIter unselects the specified iterator.
func (selection *TreeSelection) UnselectIter(iter *TreeIter) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreeIter      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(selection.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.gtk_tree_selection_unselect_iter(_arg0, _arg1)
}

// UnselectPath unselects the row at @path.
func (selection *TreeSelection) UnselectPath(path *TreePath) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreePath      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(selection.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_tree_selection_unselect_path(_arg0, _arg1)
}

// UnselectRange unselects a range of nodes, determined by @start_path and
// @end_path inclusive.
func (selection *TreeSelection) UnselectRange(startPath *TreePath, endPath *TreePath) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreePath      // out
	var _arg2 *C.GtkTreePath      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(selection.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(startPath))
	_arg2 = (*C.GtkTreePath)(unsafe.Pointer(endPath))

	C.gtk_tree_selection_unselect_range(_arg0, _arg1, _arg2)
}
