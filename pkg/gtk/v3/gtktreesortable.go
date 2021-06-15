// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_sortable_get_type()), F: marshalTreeSortable},
	})
}

// TreeSortableOverrider contains methods that are overridable. This
// interface is a subset of the interface TreeSortable.
type TreeSortableOverrider interface {
	// SortColumnID fills in @sort_column_id and @order with the current sort
	// column and the order. It returns true unless the @sort_column_id is
	// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID or
	// GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID.
	SortColumnID() (int, SortType, bool)
	// HasDefaultSortFunc returns true if the model has a default sort function.
	// This is used primarily by GtkTreeViewColumns in order to determine if a
	// model can go back to the default state, or not.
	HasDefaultSortFunc() bool
	// SetSortColumnID sets the current sort column to be @sort_column_id. The
	// @sortable will resort itself to reflect this change, after emitting a
	// TreeSortable::sort-column-changed signal. @sort_column_id may either be a
	// regular column id, or one of the following special values:
	//
	// - GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID: the default sort function
	// will be used, if it is set
	//
	// - GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID: no sorting will occur
	SetSortColumnID(sortColumnId int, order SortType)
	// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
	// @sortable.
	SortColumnChanged()
}

// TreeSortable is an interface to be implemented by tree models which support
// sorting. The TreeView uses the methods provided by this interface to sort the
// model.
type TreeSortable interface {
	TreeModel
	TreeSortableOverrider
}

// treeSortable implements the TreeSortable interface.
type treeSortable struct {
	TreeModel
}

var _ TreeSortable = (*treeSortable)(nil)

// WrapTreeSortable wraps a GObject to a type that implements interface
// TreeSortable. It is primarily used internally.
func WrapTreeSortable(obj *externglib.Object) TreeSortable {
	return TreeSortable{
		TreeModel: WrapTreeModel(obj),
	}
}

func marshalTreeSortable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeSortable(obj), nil
}

// SortColumnID fills in @sort_column_id and @order with the current sort
// column and the order. It returns true unless the @sort_column_id is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID or
// GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID.
func (s treeSortable) SortColumnID() (int, SortType, bool) {
	var _arg0 *C.GtkTreeSortable // out
	var _arg1 C.gint             // in
	var _arg2 C.GtkSortType      // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_tree_sortable_get_sort_column_id(_arg0, &_arg1, &_arg2)

	var _sortColumnId int // out
	var _order SortType   // out
	var _ok bool          // out

	_sortColumnId = (int)(_arg1)
	_order = SortType(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _sortColumnId, _order, _ok
}

// HasDefaultSortFunc returns true if the model has a default sort function.
// This is used primarily by GtkTreeViewColumns in order to determine if a
// model can go back to the default state, or not.
func (s treeSortable) HasDefaultSortFunc() bool {
	var _arg0 *C.GtkTreeSortable // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_tree_sortable_has_default_sort_func(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetSortColumnID sets the current sort column to be @sort_column_id. The
// @sortable will resort itself to reflect this change, after emitting a
// TreeSortable::sort-column-changed signal. @sort_column_id may either be a
// regular column id, or one of the following special values:
//
// - GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID: the default sort function
// will be used, if it is set
//
// - GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID: no sorting will occur
func (s treeSortable) SetSortColumnID(sortColumnId int, order SortType) {
	var _arg0 *C.GtkTreeSortable // out
	var _arg1 C.gint             // out
	var _arg2 C.GtkSortType      // out

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(s.Native()))
	_arg1 = (C.gint)(sortColumnId)
	_arg2 = (C.GtkSortType)(order)

	C.gtk_tree_sortable_set_sort_column_id(_arg0, _arg1, _arg2)
}

// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
// @sortable.
func (s treeSortable) SortColumnChanged() {
	var _arg0 *C.GtkTreeSortable // out

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(s.Native()))

	C.gtk_tree_sortable_sort_column_changed(_arg0)
}
