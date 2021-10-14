// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_store_get_type()), F: marshalTreeStorer},
	})
}

// TreeStore: tree-like data structure that can be used with the GtkTreeView
//
// The TreeStore object is a list model for use with a TreeView widget. It
// implements the TreeModel interface, and consequently, can use all of the
// methods available there. It also implements the TreeSortable interface so it
// can be sorted by the view. Finally, it also implements the tree [drag and
// drop][gtk3-GtkTreeView-drag-and-drop] interfaces.
//
//
// GtkTreeStore as GtkBuildable
//
// The GtkTreeStore implementation of the Buildable interface allows to specify
// the model columns with a <columns> element that may contain multiple <column>
// elements, each specifying one model column. The “type” attribute specifies
// the data type for the column.
//
// An example of a UI Definition fragment for a tree store:
//
//    <object class="GtkTreeStore">
//      <columns>
//        <column type="gchararray"/>
//        <column type="gchararray"/>
//        <column type="gint"/>
//      </columns>
//    </object>.
type TreeStore struct {
	*externglib.Object

	Buildable
	TreeDragDest
	TreeDragSource
	TreeSortable
}

func wrapTreeStore(obj *externglib.Object) *TreeStore {
	return &TreeStore{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
		TreeDragDest: TreeDragDest{
			Object: obj,
		},
		TreeDragSource: TreeDragSource{
			Object: obj,
		},
		TreeSortable: TreeSortable{
			TreeModel: TreeModel{
				Object: obj,
			},
		},
	}
}

func marshalTreeStorer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTreeStore(obj), nil
}

// NewTreeStore: non vararg creation function. Used primarily by language
// bindings.
//
// The function takes the following parameters:
//
//    - types: array of #GType types for the columns, from first to last.
//
func NewTreeStore(types []externglib.Type) *TreeStore {
	var _arg2 *C.GType // out
	var _arg1 C.int
	var _cret *C.GtkTreeStore // in

	_arg1 = (C.int)(len(types))
	_arg2 = (*C.GType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GType)(_arg2), len(types))
		for i := range types {
			out[i] = C.GType(types[i])
		}
	}

	_cret = C.gtk_tree_store_newv(_arg1, _arg2)
	runtime.KeepAlive(types)

	var _treeStore *TreeStore // out

	_treeStore = wrapTreeStore(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _treeStore
}

// Append appends a new row to tree_store. If parent is non-NULL, then it will
// append the new row after the last child of parent, otherwise it will append a
// row to the top level. iter will be changed to point to this new row. The row
// will be empty after this function is called. To fill in values, you need to
// call gtk_tree_store_set() or gtk_tree_store_set_value().
//
// The function takes the following parameters:
//
//    - parent: valid TreeIter, or NULL.
//
func (treeStore *TreeStore) Append(parent *TreeIter) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	if parent != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(parent)))
	}

	C.gtk_tree_store_append(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(parent)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// Clear removes all rows from tree_store.
func (treeStore *TreeStore) Clear() {
	var _arg0 *C.GtkTreeStore // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))

	C.gtk_tree_store_clear(_arg0)
	runtime.KeepAlive(treeStore)
}

// Insert creates a new row at position. If parent is non-NULL, then the row
// will be made a child of parent. Otherwise, the row will be created at the
// toplevel. If position is -1 or is larger than the number of rows at that
// level, then the new row will be inserted to the end of the list. iter will be
// changed to point to this new row. The row will be empty after this function
// is called. To fill in values, you need to call gtk_tree_store_set() or
// gtk_tree_store_set_value().
//
// The function takes the following parameters:
//
//    - parent: valid TreeIter, or NULL.
//    - position to insert the new row, or -1 for last.
//
func (treeStore *TreeStore) Insert(parent *TreeIter, position int) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out
	var _arg3 C.int           // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	if parent != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(parent)))
	}
	_arg3 = C.int(position)

	C.gtk_tree_store_insert(_arg0, &_arg1, _arg2, _arg3)
	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(position)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// InsertAfter inserts a new row after sibling. If sibling is NULL, then the row
// will be prepended to parent ’s children. If parent and sibling are NULL, then
// the row will be prepended to the toplevel. If both sibling and parent are
// set, then parent must be the parent of sibling. When sibling is set, parent
// is optional.
//
// iter will be changed to point to this new row. The row will be empty after
// this function is called. To fill in values, you need to call
// gtk_tree_store_set() or gtk_tree_store_set_value().
//
// The function takes the following parameters:
//
//    - parent: valid TreeIter, or NULL.
//    - sibling: valid TreeIter, or NULL.
//
func (treeStore *TreeStore) InsertAfter(parent, sibling *TreeIter) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out
	var _arg3 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	if parent != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(parent)))
	}
	if sibling != nil {
		_arg3 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(sibling)))
	}

	C.gtk_tree_store_insert_after(_arg0, &_arg1, _arg2, _arg3)
	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(sibling)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// InsertBefore inserts a new row before sibling. If sibling is NULL, then the
// row will be appended to parent ’s children. If parent and sibling are NULL,
// then the row will be appended to the toplevel. If both sibling and parent are
// set, then parent must be the parent of sibling. When sibling is set, parent
// is optional.
//
// iter will be changed to point to this new row. The row will be empty after
// this function is called. To fill in values, you need to call
// gtk_tree_store_set() or gtk_tree_store_set_value().
//
// The function takes the following parameters:
//
//    - parent: valid TreeIter, or NULL.
//    - sibling: valid TreeIter, or NULL.
//
func (treeStore *TreeStore) InsertBefore(parent, sibling *TreeIter) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out
	var _arg3 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	if parent != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(parent)))
	}
	if sibling != nil {
		_arg3 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(sibling)))
	}

	C.gtk_tree_store_insert_before(_arg0, &_arg1, _arg2, _arg3)
	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(sibling)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// InsertWithValues: variant of gtk_tree_store_insert_with_values() which takes
// the columns and values as two arrays, instead of varargs. This function is
// mainly intended for language bindings.
//
// The function takes the following parameters:
//
//    - parent: valid TreeIter, or NULL.
//    - position to insert the new row, or -1 for last.
//    - columns: array of column numbers.
//    - values: array of GValues.
//
func (treeStore *TreeStore) InsertWithValues(parent *TreeIter, position int, columns []int, values []externglib.Value) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out
	var _arg3 C.int           // out
	var _arg4 *C.int          // out
	var _arg6 C.int
	var _arg5 *C.GValue // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	if parent != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(parent)))
	}
	_arg3 = C.int(position)
	_arg6 = (C.int)(len(columns))
	_arg4 = (*C.int)(C.malloc(C.ulong(len(columns)) * C.ulong(C.sizeof_int)))
	defer C.free(unsafe.Pointer(_arg4))
	{
		out := unsafe.Slice((*C.int)(_arg4), len(columns))
		for i := range columns {
			out[i] = C.int(columns[i])
		}
	}
	_arg6 = (C.int)(len(values))
	_arg5 = (*C.GValue)(C.malloc(C.ulong(len(values)) * C.ulong(C.sizeof_GValue)))
	defer C.free(unsafe.Pointer(_arg5))
	{
		out := unsafe.Slice((*C.GValue)(_arg5), len(values))
		for i := range values {
			out[i] = *(*C.GValue)(unsafe.Pointer((&values[i]).Native()))
		}
	}

	C.gtk_tree_store_insert_with_valuesv(_arg0, &_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(position)
	runtime.KeepAlive(columns)
	runtime.KeepAlive(values)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// IsAncestor returns TRUE if iter is an ancestor of descendant. That is, iter
// is the parent (or grandparent or great-grandparent) of descendant.
//
// The function takes the following parameters:
//
//    - iter: valid TreeIter.
//    - descendant: valid TreeIter.
//
func (treeStore *TreeStore) IsAncestor(iter, descendant *TreeIter) bool {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(descendant)))

	_cret = C.gtk_tree_store_is_ancestor(_arg0, _arg1, _arg2)
	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(descendant)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IterDepth returns the depth of iter. This will be 0 for anything on the root
// level, 1 for anything down a level, etc.
//
// The function takes the following parameters:
//
//    - iter: valid TreeIter.
//
func (treeStore *TreeStore) IterDepth(iter *TreeIter) int {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.int           // in

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_tree_store_iter_depth(_arg0, _arg1)
	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(iter)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IterIsValid: WARNING: This function is slow. Only use it for debugging and/or
// testing purposes.
//
// Checks if the given iter is a valid iter for this TreeStore.
//
// The function takes the following parameters:
//
//    - iter: TreeIter.
//
func (treeStore *TreeStore) IterIsValid(iter *TreeIter) bool {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_tree_store_iter_is_valid(_arg0, _arg1)
	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MoveAfter moves iter in tree_store to the position after position. iter and
// position should be in the same level. Note that this function only works with
// unsorted stores. If position is NULL, iter will be moved to the start of the
// level.
//
// The function takes the following parameters:
//
//    - iter: TreeIter.
//    - position: TreeIter.
//
func (treeStore *TreeStore) MoveAfter(iter, position *TreeIter) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	if position != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(position)))
	}

	C.gtk_tree_store_move_after(_arg0, _arg1, _arg2)
	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(position)
}

// MoveBefore moves iter in tree_store to the position before position. iter and
// position should be in the same level. Note that this function only works with
// unsorted stores. If position is NULL, iter will be moved to the end of the
// level.
//
// The function takes the following parameters:
//
//    - iter: TreeIter.
//    - position or NULL.
//
func (treeStore *TreeStore) MoveBefore(iter, position *TreeIter) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	if position != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(position)))
	}

	C.gtk_tree_store_move_before(_arg0, _arg1, _arg2)
	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(position)
}

// Prepend prepends a new row to tree_store. If parent is non-NULL, then it will
// prepend the new row before the first child of parent, otherwise it will
// prepend a row to the top level. iter will be changed to point to this new
// row. The row will be empty after this function is called. To fill in values,
// you need to call gtk_tree_store_set() or gtk_tree_store_set_value().
//
// The function takes the following parameters:
//
//    - parent: valid TreeIter, or NULL.
//
func (treeStore *TreeStore) Prepend(parent *TreeIter) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	if parent != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(parent)))
	}

	C.gtk_tree_store_prepend(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(parent)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// Remove removes iter from tree_store. After being removed, iter is set to the
// next valid row at that level, or invalidated if it previously pointed to the
// last one.
//
// The function takes the following parameters:
//
//    - iter: valid TreeIter.
//
func (treeStore *TreeStore) Remove(iter *TreeIter) bool {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_tree_store_remove(_arg0, _arg1)
	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetColumnTypes: this function is meant primarily for #GObjects that inherit
// from TreeStore, and should only be used when constructing a new TreeStore. It
// will not function after a row has been added, or a method on the TreeModel
// interface is called.
//
// The function takes the following parameters:
//
//    - types: array of #GType types, one for each column.
//
func (treeStore *TreeStore) SetColumnTypes(types []externglib.Type) {
	var _arg0 *C.GtkTreeStore // out
	var _arg2 *C.GType        // out
	var _arg1 C.int

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (C.int)(len(types))
	_arg2 = (*C.GType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GType)(_arg2), len(types))
		for i := range types {
			out[i] = C.GType(types[i])
		}
	}

	C.gtk_tree_store_set_column_types(_arg0, _arg1, _arg2)
	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(types)
}

// SetValue sets the data in the cell specified by iter and column. The type of
// value must be convertible to the type of the column.
//
// The function takes the following parameters:
//
//    - iter: valid TreeIter for the row being modified.
//    - column number to modify.
//    - value: new value for the cell.
//
func (treeStore *TreeStore) SetValue(iter *TreeIter, column int, value *externglib.Value) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 C.int           // out
	var _arg3 *C.GValue       // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg2 = C.int(column)
	_arg3 = (*C.GValue)(unsafe.Pointer(value.Native()))

	C.gtk_tree_store_set_value(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(column)
	runtime.KeepAlive(value)
}

// Set: variant of gtk_tree_store_set_valist() which takes the columns and
// values as two arrays, instead of varargs. This function is mainly intended
// for language bindings or in case the number of columns to change is not known
// until run-time.
//
// The function takes the following parameters:
//
//    - iter: valid TreeIter for the row being modified.
//    - columns: array of column numbers.
//    - values: array of GValues.
//
func (treeStore *TreeStore) Set(iter *TreeIter, columns []int, values []externglib.Value) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.int          // out
	var _arg4 C.int
	var _arg3 *C.GValue // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg4 = (C.int)(len(columns))
	_arg2 = (*C.int)(C.malloc(C.ulong(len(columns)) * C.ulong(C.sizeof_int)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.int)(_arg2), len(columns))
		for i := range columns {
			out[i] = C.int(columns[i])
		}
	}
	_arg4 = (C.int)(len(values))
	_arg3 = (*C.GValue)(C.malloc(C.ulong(len(values)) * C.ulong(C.sizeof_GValue)))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice((*C.GValue)(_arg3), len(values))
		for i := range values {
			out[i] = *(*C.GValue)(unsafe.Pointer((&values[i]).Native()))
		}
	}

	C.gtk_tree_store_set_valuesv(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(columns)
	runtime.KeepAlive(values)
}

// Swap swaps a and b in the same level of tree_store. Note that this function
// only works with unsorted stores.
//
// The function takes the following parameters:
//
//    - a: TreeIter.
//    - b: another TreeIter.
//
func (treeStore *TreeStore) Swap(a, b *TreeIter) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(a)))
	_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(b)))

	C.gtk_tree_store_swap(_arg0, _arg1, _arg2)
	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
