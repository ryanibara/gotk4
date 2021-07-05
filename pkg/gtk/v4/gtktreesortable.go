// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
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
		{T: externglib.Type(C.gtk_tree_sortable_get_type()), F: marshalTreeSortable},
	})
}

// TreeIterCompareFunc should return a negative integer, zero, or a positive
// integer if @a sorts before @b, @a sorts with @b, or @a sorts after @b
// respectively.
//
// If two iters compare as equal, their order in the sorted model is undefined.
// In order to ensure that the TreeSortable behaves as expected, the
// GtkTreeIterCompareFunc must define a partial order on the model, i.e. it must
// be reflexive, antisymmetric and transitive.
//
// For example, if @model is a product catalogue, then a compare function for
// the “price” column could be one which returns `price_of(@a) - price_of(@b)`.
type TreeIterCompareFunc func(model TreeModel, a TreeIter, b TreeIter) (gint int)

//export gotk4_TreeIterCompareFunc
func gotk4_TreeIterCompareFunc(arg0 *C.GtkTreeModel, arg1 *C.GtkTreeIter, arg2 *C.GtkTreeIter, arg3 C.gpointer) C.int {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var model TreeModel // out
	var a TreeIter      // out
	var b TreeIter      // out

	model = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(TreeModel)
	a = (TreeIter)(unsafe.Pointer(arg1))
	b = (TreeIter)(unsafe.Pointer(arg2))

	fn := v.(TreeIterCompareFunc)
	gint := fn(model, a, b)

	var cret C.int // out

	cret = C.int(gint)

	return cret
}

// TreeSortable: the interface for sortable models used by GtkTreeView
//
// TreeSortable is an interface to be implemented by tree models which support
// sorting. The TreeView uses the methods provided by this interface to sort the
// model.
type TreeSortable interface {
	TreeModel

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

// treeSortable implements the TreeSortable interface.
type treeSortable struct {
	TreeModel
}

var _ TreeSortable = (*treeSortable)(nil)

// WrapTreeSortable wraps a GObject to a type that implements
// interface TreeSortable. It is primarily used internally.
func WrapTreeSortable(obj *externglib.Object) TreeSortable {
	return treeSortable{
		TreeModel: WrapTreeModel(obj),
	}
}

func marshalTreeSortable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeSortable(obj), nil
}

func (s treeSortable) SortColumnID() (int, SortType, bool) {
	var _arg0 *C.GtkTreeSortable // out
	var _arg1 *C.int             // in
	var _arg2 *C.GtkSortType     // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_tree_sortable_get_sort_column_id(_arg0, &_arg1, &_arg2)

	var _sortColumnId int // out
	var _order SortType   // out
	var _ok bool          // out

	_sortColumnId = int(_arg1)
	{
		var refTmpIn C.GtkSortType
		var refTmpOut SortType

		refTmpIn = *_arg2

		refTmpOut = SortType(refTmpIn)

		_order = refTmpOut
	}
	if _cret != 0 {
		_ok = true
	}

	return _sortColumnId, _order, _ok
}

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

func (s treeSortable) SetSortColumnID(sortColumnId int, order SortType) {
	var _arg0 *C.GtkTreeSortable // out
	var _arg1 C.int              // out
	var _arg2 C.GtkSortType      // out

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(sortColumnId)
	_arg2 = C.GtkSortType(order)

	C.gtk_tree_sortable_set_sort_column_id(_arg0, _arg1, _arg2)
}

func (s treeSortable) SortColumnChanged() {
	var _arg0 *C.GtkTreeSortable // out

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(s.Native()))

	C.gtk_tree_sortable_sort_column_changed(_arg0)
}
