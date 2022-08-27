// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void callbackDelete(gpointer);
// extern void _gotk4_gtk4_TreeSortable_ConnectSortColumnChanged(gpointer, guintptr);
// extern int _gotk4_gtk4_TreeIterCompareFunc(GtkTreeModel*, GtkTreeIter*, GtkTreeIter*, gpointer);
// gboolean _gotk4_gtk4_TreeSortable_virtual_get_sort_column_id(void* fnptr, GtkTreeSortable* arg0, int* arg1, GtkSortType* arg2) {
//   return ((gboolean (*)(GtkTreeSortable*, int*, GtkSortType*))(fnptr))(arg0, arg1, arg2);
// };
// gboolean _gotk4_gtk4_TreeSortable_virtual_has_default_sort_func(void* fnptr, GtkTreeSortable* arg0) {
//   return ((gboolean (*)(GtkTreeSortable*))(fnptr))(arg0);
// };
// void _gotk4_gtk4_TreeSortable_virtual_set_default_sort_func(void* fnptr, GtkTreeSortable* arg0, GtkTreeIterCompareFunc arg1, gpointer arg2, GDestroyNotify arg3) {
//   ((void (*)(GtkTreeSortable*, GtkTreeIterCompareFunc, gpointer, GDestroyNotify))(fnptr))(arg0, arg1, arg2, arg3);
// };
// void _gotk4_gtk4_TreeSortable_virtual_set_sort_column_id(void* fnptr, GtkTreeSortable* arg0, int arg1, GtkSortType arg2) {
//   ((void (*)(GtkTreeSortable*, int, GtkSortType))(fnptr))(arg0, arg1, arg2);
// };
// void _gotk4_gtk4_TreeSortable_virtual_set_sort_func(void* fnptr, GtkTreeSortable* arg0, int arg1, GtkTreeIterCompareFunc arg2, gpointer arg3, GDestroyNotify arg4) {
//   ((void (*)(GtkTreeSortable*, int, GtkTreeIterCompareFunc, gpointer, GDestroyNotify))(fnptr))(arg0, arg1, arg2, arg3, arg4);
// };
// void _gotk4_gtk4_TreeSortable_virtual_sort_column_changed(void* fnptr, GtkTreeSortable* arg0) {
//   ((void (*)(GtkTreeSortable*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeTreeSortable = coreglib.Type(C.gtk_tree_sortable_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTreeSortable, F: marshalTreeSortable},
	})
}

// TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID uses the default sort function in a
// gtk.TreeSortable.
//
// See also: gtk.TreeSortable.SetSortColumnID().
const TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID = -1

// TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID disables sorting in a gtk.TreeSortable.
//
// See also: gtk.TreeSortable.SetSortColumnID().
const TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID = -2

// TreeIterCompareFunc should return a negative integer, zero, or a positive
// integer if a sorts before b, a sorts with b, or a sorts after b respectively.
//
// If two iters compare as equal, their order in the sorted model is undefined.
// In order to ensure that the TreeSortable behaves as expected, the
// GtkTreeIterCompareFunc must define a partial order on the model, i.e. it must
// be reflexive, antisymmetric and transitive.
//
// For example, if model is a product catalogue, then a compare function for the
// “price” column could be one which returns price_of(a) - price_of(b).
type TreeIterCompareFunc func(model TreeModeller, a, b *TreeIter) (gint int)

// TreeSortable: interface for sortable models used by GtkTreeView
//
// TreeSortable is an interface to be implemented by tree models which support
// sorting. The TreeView uses the methods provided by this interface to sort the
// model.
//
// TreeSortable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type TreeSortable struct {
	_ [0]func() // equal guard
	TreeModel
}

var ()

// TreeSortabler describes TreeSortable's interface methods.
type TreeSortabler interface {
	coreglib.Objector

	// SortColumnID fills in sort_column_id and order with the current sort
	// column and the order.
	SortColumnID() (int, SortType, bool)
	// HasDefaultSortFunc returns TRUE if the model has a default sort function.
	HasDefaultSortFunc() bool
	// SetDefaultSortFunc sets the default comparison function used when sorting
	// to be sort_func.
	SetDefaultSortFunc(sortFunc TreeIterCompareFunc)
	// SetSortColumnID sets the current sort column to be sort_column_id.
	SetSortColumnID(sortColumnId int, order SortType)
	// SetSortFunc sets the comparison function used when sorting to be
	// sort_func.
	SetSortFunc(sortColumnId int, sortFunc TreeIterCompareFunc)
	// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
	// sortable.
	SortColumnChanged()

	// Sort-column-changed signal is emitted when the sort column or sort order
	// of sortable is changed.
	ConnectSortColumnChanged(func()) coreglib.SignalHandle
}

var _ TreeSortabler = (*TreeSortable)(nil)

func wrapTreeSortable(obj *coreglib.Object) *TreeSortable {
	return &TreeSortable{
		TreeModel: TreeModel{
			Object: obj,
		},
	}
}

func marshalTreeSortable(p uintptr) (interface{}, error) {
	return wrapTreeSortable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectSortColumnChanged signal is emitted when the sort column or sort order
// of sortable is changed. The signal is emitted before the contents of sortable
// are resorted.
func (sortable *TreeSortable) ConnectSortColumnChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(sortable, "sort-column-changed", false, unsafe.Pointer(C._gotk4_gtk4_TreeSortable_ConnectSortColumnChanged), f)
}

// SortColumnID fills in sort_column_id and order with the current sort column
// and the order. It returns TRUE unless the sort_column_id is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID or
// GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID.
//
// The function returns the following values:
//
//    - sortColumnId: sort column id to be filled in.
//    - order to be filled in.
//    - ok: TRUE if the sort column is not one of the special sort column ids.
//
func (sortable *TreeSortable) SortColumnID() (int, SortType, bool) {
	var _arg0 *C.GtkTreeSortable // out
	var _arg1 C.int              // in
	var _arg2 C.GtkSortType      // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))

	_cret = C.gtk_tree_sortable_get_sort_column_id(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(sortable)

	var _sortColumnId int // out
	var _order SortType   // out
	var _ok bool          // out

	_sortColumnId = int(_arg1)
	_order = SortType(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _sortColumnId, _order, _ok
}

// HasDefaultSortFunc returns TRUE if the model has a default sort function.
// This is used primarily by GtkTreeViewColumns in order to determine if a model
// can go back to the default state, or not.
//
// The function returns the following values:
//
//    - ok: TRUE, if the model has a default sort function.
//
func (sortable *TreeSortable) HasDefaultSortFunc() bool {
	var _arg0 *C.GtkTreeSortable // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))

	_cret = C.gtk_tree_sortable_has_default_sort_func(_arg0)
	runtime.KeepAlive(sortable)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDefaultSortFunc sets the default comparison function used when sorting to
// be sort_func. If the current sort column id of sortable is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, then the model will sort using this
// function.
//
// If sort_func is NULL, then there will be no default comparison function. This
// means that once the model has been sorted, it can’t go back to the default
// state. In this case, when the current sort column id of sortable is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, the model will be unsorted.
//
// The function takes the following parameters:
//
//    - sortFunc: comparison function.
//
func (sortable *TreeSortable) SetDefaultSortFunc(sortFunc TreeIterCompareFunc) {
	var _arg0 *C.GtkTreeSortable       // out
	var _arg1 C.GtkTreeIterCompareFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk4_TreeIterCompareFunc)
	_arg2 = C.gpointer(gbox.Assign(sortFunc))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_tree_sortable_set_default_sort_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(sortable)
	runtime.KeepAlive(sortFunc)
}

// SetSortColumnID sets the current sort column to be sort_column_id. The
// sortable will resort itself to reflect this change, after emitting a
// TreeSortable::sort-column-changed signal. sort_column_id may either be a
// regular column id, or one of the following special values:
//
// - GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID: the default sort function will be
// used, if it is set
//
// - GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID: no sorting will occur.
//
// The function takes the following parameters:
//
//    - sortColumnId: sort column id to set.
//    - order: sort order of the column.
//
func (sortable *TreeSortable) SetSortColumnID(sortColumnId int, order SortType) {
	var _arg0 *C.GtkTreeSortable // out
	var _arg1 C.int              // out
	var _arg2 C.GtkSortType      // out

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))
	_arg1 = C.int(sortColumnId)
	_arg2 = C.GtkSortType(order)

	C.gtk_tree_sortable_set_sort_column_id(_arg0, _arg1, _arg2)
	runtime.KeepAlive(sortable)
	runtime.KeepAlive(sortColumnId)
	runtime.KeepAlive(order)
}

// SetSortFunc sets the comparison function used when sorting to be sort_func.
// If the current sort column id of sortable is the same as sort_column_id, then
// the model will sort using this function.
//
// The function takes the following parameters:
//
//    - sortColumnId: sort column id to set the function for.
//    - sortFunc: comparison function.
//
func (sortable *TreeSortable) SetSortFunc(sortColumnId int, sortFunc TreeIterCompareFunc) {
	var _arg0 *C.GtkTreeSortable       // out
	var _arg1 C.int                    // out
	var _arg2 C.GtkTreeIterCompareFunc // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))
	_arg1 = C.int(sortColumnId)
	_arg2 = (*[0]byte)(C._gotk4_gtk4_TreeIterCompareFunc)
	_arg3 = C.gpointer(gbox.Assign(sortFunc))
	_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_tree_sortable_set_sort_func(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(sortable)
	runtime.KeepAlive(sortColumnId)
	runtime.KeepAlive(sortFunc)
}

// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
// sortable.
func (sortable *TreeSortable) SortColumnChanged() {
	var _arg0 *C.GtkTreeSortable // out

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))

	C.gtk_tree_sortable_sort_column_changed(_arg0)
	runtime.KeepAlive(sortable)
}

// sortColumnID fills in sort_column_id and order with the current sort column
// and the order. It returns TRUE unless the sort_column_id is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID or
// GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID.
//
// The function returns the following values:
//
//    - sortColumnId: sort column id to be filled in.
//    - order to be filled in.
//    - ok: TRUE if the sort column is not one of the special sort column ids.
//
func (sortable *TreeSortable) sortColumnID() (int, SortType, bool) {
	gclass := (*C.GtkTreeSortableIface)(coreglib.PeekParentClass(sortable))
	fnarg := gclass.get_sort_column_id

	var _arg0 *C.GtkTreeSortable // out
	var _arg1 C.int              // in
	var _arg2 C.GtkSortType      // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))

	_cret = C._gotk4_gtk4_TreeSortable_virtual_get_sort_column_id(unsafe.Pointer(fnarg), _arg0, &_arg1, &_arg2)
	runtime.KeepAlive(sortable)

	var _sortColumnId int // out
	var _order SortType   // out
	var _ok bool          // out

	_sortColumnId = int(_arg1)
	_order = SortType(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _sortColumnId, _order, _ok
}

// hasDefaultSortFunc returns TRUE if the model has a default sort function.
// This is used primarily by GtkTreeViewColumns in order to determine if a model
// can go back to the default state, or not.
//
// The function returns the following values:
//
//    - ok: TRUE, if the model has a default sort function.
//
func (sortable *TreeSortable) hasDefaultSortFunc() bool {
	gclass := (*C.GtkTreeSortableIface)(coreglib.PeekParentClass(sortable))
	fnarg := gclass.has_default_sort_func

	var _arg0 *C.GtkTreeSortable // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))

	_cret = C._gotk4_gtk4_TreeSortable_virtual_has_default_sort_func(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(sortable)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// setDefaultSortFunc sets the default comparison function used when sorting to
// be sort_func. If the current sort column id of sortable is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, then the model will sort using this
// function.
//
// If sort_func is NULL, then there will be no default comparison function. This
// means that once the model has been sorted, it can’t go back to the default
// state. In this case, when the current sort column id of sortable is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, the model will be unsorted.
//
// The function takes the following parameters:
//
//    - sortFunc: comparison function.
//
func (sortable *TreeSortable) setDefaultSortFunc(sortFunc TreeIterCompareFunc) {
	gclass := (*C.GtkTreeSortableIface)(coreglib.PeekParentClass(sortable))
	fnarg := gclass.set_default_sort_func

	var _arg0 *C.GtkTreeSortable       // out
	var _arg1 C.GtkTreeIterCompareFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk4_TreeIterCompareFunc)
	_arg2 = C.gpointer(gbox.Assign(sortFunc))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C._gotk4_gtk4_TreeSortable_virtual_set_default_sort_func(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(sortable)
	runtime.KeepAlive(sortFunc)
}

// setSortColumnID sets the current sort column to be sort_column_id. The
// sortable will resort itself to reflect this change, after emitting a
// TreeSortable::sort-column-changed signal. sort_column_id may either be a
// regular column id, or one of the following special values:
//
// - GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID: the default sort function will be
// used, if it is set
//
// - GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID: no sorting will occur.
//
// The function takes the following parameters:
//
//    - sortColumnId: sort column id to set.
//    - order: sort order of the column.
//
func (sortable *TreeSortable) setSortColumnID(sortColumnId int, order SortType) {
	gclass := (*C.GtkTreeSortableIface)(coreglib.PeekParentClass(sortable))
	fnarg := gclass.set_sort_column_id

	var _arg0 *C.GtkTreeSortable // out
	var _arg1 C.int              // out
	var _arg2 C.GtkSortType      // out

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))
	_arg1 = C.int(sortColumnId)
	_arg2 = C.GtkSortType(order)

	C._gotk4_gtk4_TreeSortable_virtual_set_sort_column_id(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(sortable)
	runtime.KeepAlive(sortColumnId)
	runtime.KeepAlive(order)
}

// setSortFunc sets the comparison function used when sorting to be sort_func.
// If the current sort column id of sortable is the same as sort_column_id, then
// the model will sort using this function.
//
// The function takes the following parameters:
//
//    - sortColumnId: sort column id to set the function for.
//    - sortFunc: comparison function.
//
func (sortable *TreeSortable) setSortFunc(sortColumnId int, sortFunc TreeIterCompareFunc) {
	gclass := (*C.GtkTreeSortableIface)(coreglib.PeekParentClass(sortable))
	fnarg := gclass.set_sort_func

	var _arg0 *C.GtkTreeSortable       // out
	var _arg1 C.int                    // out
	var _arg2 C.GtkTreeIterCompareFunc // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))
	_arg1 = C.int(sortColumnId)
	_arg2 = (*[0]byte)(C._gotk4_gtk4_TreeIterCompareFunc)
	_arg3 = C.gpointer(gbox.Assign(sortFunc))
	_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C._gotk4_gtk4_TreeSortable_virtual_set_sort_func(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(sortable)
	runtime.KeepAlive(sortColumnId)
	runtime.KeepAlive(sortFunc)
}

// sortColumnChanged emits a TreeSortable::sort-column-changed signal on
// sortable.
func (sortable *TreeSortable) sortColumnChanged() {
	gclass := (*C.GtkTreeSortableIface)(coreglib.PeekParentClass(sortable))
	fnarg := gclass.sort_column_changed

	var _arg0 *C.GtkTreeSortable // out

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))

	C._gotk4_gtk4_TreeSortable_virtual_sort_column_changed(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(sortable)
}

// TreeSortableIface: instance of this type is always passed by reference.
type TreeSortableIface struct {
	*treeSortableIface
}

// treeSortableIface is the struct that's finalized.
type treeSortableIface struct {
	native *C.GtkTreeSortableIface
}
