// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_store_get_type()), F: marshalTreeStorer},
	})
}

// TreeStorer describes TreeStore's methods.
type TreeStorer interface {
	// Append appends a new row to @tree_store.
	Append(parent *TreeIter) TreeIter
	// Clear removes all rows from @tree_store
	Clear()
	// Insert creates a new row at @position.
	Insert(parent *TreeIter, position int) TreeIter
	// InsertAfter inserts a new row after @sibling.
	InsertAfter(parent *TreeIter, sibling *TreeIter) TreeIter
	// InsertBefore inserts a new row before @sibling.
	InsertBefore(parent *TreeIter, sibling *TreeIter) TreeIter
	// InsertWithValuesv: variant of gtk_tree_store_insert_with_values() which
	// takes the columns and values as two arrays, instead of varargs.
	InsertWithValuesv(parent *TreeIter, position int, columns []int, values []externglib.Value) TreeIter
	// IsAncestor returns true if @iter is an ancestor of @descendant.
	IsAncestor(iter *TreeIter, descendant *TreeIter) bool
	// IterDepth returns the depth of @iter.
	IterDepth(iter *TreeIter) int
	// IterIsValid: WARNING: This function is slow.
	IterIsValid(iter *TreeIter) bool
	// MoveAfter moves @iter in @tree_store to the position after @position.
	MoveAfter(iter *TreeIter, position *TreeIter)
	// MoveBefore moves @iter in @tree_store to the position before @position.
	MoveBefore(iter *TreeIter, position *TreeIter)
	// Prepend prepends a new row to @tree_store.
	Prepend(parent *TreeIter) TreeIter
	// Remove removes @iter from @tree_store.
	Remove(iter *TreeIter) bool
	// SetColumnTypes: this function is meant primarily for #GObjects that
	// inherit from TreeStore, and should only be used when constructing a new
	// TreeStore.
	SetColumnTypes(types []externglib.Type)
	// SetValue sets the data in the cell specified by @iter and @column.
	SetValue(iter *TreeIter, column int, value *externglib.Value)
	// SetValuesv: variant of gtk_tree_store_set_valist() which takes the
	// columns and values as two arrays, instead of varargs.
	SetValuesv(iter *TreeIter, columns []int, values []externglib.Value)
	// Swap swaps @a and @b in the same level of @tree_store.
	Swap(a *TreeIter, b *TreeIter)
}

// TreeStore object is a list model for use with a TreeView widget. It
// implements the TreeModel interface, and consequentially, can use all of the
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
//    </object>
type TreeStore struct {
	*externglib.Object

	Buildable
	TreeDragDest
	TreeDragSource
	TreeSortable
}

var (
	_ TreeStorer      = (*TreeStore)(nil)
	_ gextras.Nativer = (*TreeStore)(nil)
)

func wrapTreeStore(obj *externglib.Object) TreeStorer {
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

// NewTreeStoreV: non vararg creation function. Used primarily by language
// bindings.
func NewTreeStoreV(types []externglib.Type) *TreeStore {
	var _arg2 *C.GType
	var _arg1 C.gint
	var _cret *C.GtkTreeStore // in

	_arg1 = C.gint(len(types))
	_arg2 = (*C.GType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(types))
		for i := range types {
			out[i] = (C.GType)(types[i])
		}
	}

	_cret = C.gtk_tree_store_newv(_arg1, _arg2)

	var _treeStore *TreeStore // out

	_treeStore = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*TreeStore)

	return _treeStore
}

// Append appends a new row to @tree_store. If @parent is non-nil, then it will
// append the new row after the last child of @parent, otherwise it will append
// a row to the top level. @iter will be changed to point to this new row. The
// row will be empty after this function is called. To fill in values, you need
// to call gtk_tree_store_set() or gtk_tree_store_set_value().
func (treeStore *TreeStore) Append(parent *TreeIter) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent))

	C.gtk_tree_store_append(_arg0, &_arg1, _arg2)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(unsafe.Pointer((&_arg1)))

	return _iter
}

// Clear removes all rows from @tree_store
func (treeStore *TreeStore) Clear() {
	var _arg0 *C.GtkTreeStore // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))

	C.gtk_tree_store_clear(_arg0)
}

// Insert creates a new row at @position. If parent is non-nil, then the row
// will be made a child of @parent. Otherwise, the row will be created at the
// toplevel. If @position is -1 or is larger than the number of rows at that
// level, then the new row will be inserted to the end of the list. @iter will
// be changed to point to this new row. The row will be empty after this
// function is called. To fill in values, you need to call gtk_tree_store_set()
// or gtk_tree_store_set_value().
func (treeStore *TreeStore) Insert(parent *TreeIter, position int) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out
	var _arg3 C.gint          // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent))
	_arg3 = C.gint(position)

	C.gtk_tree_store_insert(_arg0, &_arg1, _arg2, _arg3)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(unsafe.Pointer((&_arg1)))

	return _iter
}

// InsertAfter inserts a new row after @sibling. If @sibling is nil, then the
// row will be prepended to @parent ’s children. If @parent and @sibling are
// nil, then the row will be prepended to the toplevel. If both @sibling and
// @parent are set, then @parent must be the parent of @sibling. When @sibling
// is set, @parent is optional.
//
// @iter will be changed to point to this new row. The row will be empty after
// this function is called. To fill in values, you need to call
// gtk_tree_store_set() or gtk_tree_store_set_value().
func (treeStore *TreeStore) InsertAfter(parent *TreeIter, sibling *TreeIter) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out
	var _arg3 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent))
	_arg3 = (*C.GtkTreeIter)(unsafe.Pointer(sibling))

	C.gtk_tree_store_insert_after(_arg0, &_arg1, _arg2, _arg3)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(unsafe.Pointer((&_arg1)))

	return _iter
}

// InsertBefore inserts a new row before @sibling. If @sibling is nil, then the
// row will be appended to @parent ’s children. If @parent and @sibling are nil,
// then the row will be appended to the toplevel. If both @sibling and @parent
// are set, then @parent must be the parent of @sibling. When @sibling is set,
// @parent is optional.
//
// @iter will be changed to point to this new row. The row will be empty after
// this function is called. To fill in values, you need to call
// gtk_tree_store_set() or gtk_tree_store_set_value().
func (treeStore *TreeStore) InsertBefore(parent *TreeIter, sibling *TreeIter) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out
	var _arg3 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent))
	_arg3 = (*C.GtkTreeIter)(unsafe.Pointer(sibling))

	C.gtk_tree_store_insert_before(_arg0, &_arg1, _arg2, _arg3)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(unsafe.Pointer((&_arg1)))

	return _iter
}

// InsertWithValuesv: variant of gtk_tree_store_insert_with_values() which takes
// the columns and values as two arrays, instead of varargs. This function is
// mainly intended for language bindings.
func (treeStore *TreeStore) InsertWithValuesv(parent *TreeIter, position int, columns []int, values []externglib.Value) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out
	var _arg3 C.gint          // out
	var _arg4 *C.gint
	var _arg6 C.gint
	var _arg5 *C.GValue

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent))
	_arg3 = C.gint(position)
	_arg6 = C.gint(len(columns))
	_arg4 = (*C.gint)(unsafe.Pointer(&columns[0]))
	_arg6 = C.gint(len(values))
	_arg5 = (*C.GValue)(C.malloc(C.ulong(len(values)) * C.ulong(C.sizeof_GValue)))
	defer C.free(unsafe.Pointer(_arg5))
	{
		out := unsafe.Slice(_arg5, len(values))
		for i := range values {
			out[i] = *(*C.GValue)(unsafe.Pointer(&(&values[i]).GValue))
		}
	}

	C.gtk_tree_store_insert_with_valuesv(_arg0, &_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(unsafe.Pointer((&_arg1)))

	return _iter
}

// IsAncestor returns true if @iter is an ancestor of @descendant. That is,
// @iter is the parent (or grandparent or great-grandparent) of @descendant.
func (treeStore *TreeStore) IsAncestor(iter *TreeIter, descendant *TreeIter) bool {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(descendant))

	_cret = C.gtk_tree_store_is_ancestor(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IterDepth returns the depth of @iter. This will be 0 for anything on the root
// level, 1 for anything down a level, etc.
func (treeStore *TreeStore) IterDepth(iter *TreeIter) int {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))

	_cret = C.gtk_tree_store_iter_depth(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IterIsValid: WARNING: This function is slow. Only use it for debugging and/or
// testing purposes.
//
// Checks if the given iter is a valid iter for this TreeStore.
func (treeStore *TreeStore) IterIsValid(iter *TreeIter) bool {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))

	_cret = C.gtk_tree_store_iter_is_valid(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MoveAfter moves @iter in @tree_store to the position after @position. @iter
// and @position should be in the same level. Note that this function only works
// with unsorted stores. If @position is nil, @iter will be moved to the start
// of the level.
func (treeStore *TreeStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(position))

	C.gtk_tree_store_move_after(_arg0, _arg1, _arg2)
}

// MoveBefore moves @iter in @tree_store to the position before @position. @iter
// and @position should be in the same level. Note that this function only works
// with unsorted stores. If @position is nil, @iter will be moved to the end of
// the level.
func (treeStore *TreeStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(position))

	C.gtk_tree_store_move_before(_arg0, _arg1, _arg2)
}

// Prepend prepends a new row to @tree_store. If @parent is non-nil, then it
// will prepend the new row before the first child of @parent, otherwise it will
// prepend a row to the top level. @iter will be changed to point to this new
// row. The row will be empty after this function is called. To fill in values,
// you need to call gtk_tree_store_set() or gtk_tree_store_set_value().
func (treeStore *TreeStore) Prepend(parent *TreeIter) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent))

	C.gtk_tree_store_prepend(_arg0, &_arg1, _arg2)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(unsafe.Pointer((&_arg1)))

	return _iter
}

// Remove removes @iter from @tree_store. After being removed, @iter is set to
// the next valid row at that level, or invalidated if it previously pointed to
// the last one.
func (treeStore *TreeStore) Remove(iter *TreeIter) bool {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))

	_cret = C.gtk_tree_store_remove(_arg0, _arg1)

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
func (treeStore *TreeStore) SetColumnTypes(types []externglib.Type) {
	var _arg0 *C.GtkTreeStore // out
	var _arg2 *C.GType
	var _arg1 C.gint

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = C.gint(len(types))
	_arg2 = (*C.GType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(types))
		for i := range types {
			out[i] = (C.GType)(types[i])
		}
	}

	C.gtk_tree_store_set_column_types(_arg0, _arg1, _arg2)
}

// SetValue sets the data in the cell specified by @iter and @column. The type
// of @value must be convertible to the type of the column.
func (treeStore *TreeStore) SetValue(iter *TreeIter, column int, value *externglib.Value) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 C.gint          // out
	var _arg3 *C.GValue       // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))
	_arg2 = C.gint(column)
	_arg3 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	C.gtk_tree_store_set_value(_arg0, _arg1, _arg2, _arg3)
}

// SetValuesv: variant of gtk_tree_store_set_valist() which takes the columns
// and values as two arrays, instead of varargs. This function is mainly
// intended for language bindings or in case the number of columns to change is
// not known until run-time.
func (treeStore *TreeStore) SetValuesv(iter *TreeIter, columns []int, values []externglib.Value) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.gint
	var _arg4 C.gint
	var _arg3 *C.GValue

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))
	_arg4 = C.gint(len(columns))
	_arg2 = (*C.gint)(unsafe.Pointer(&columns[0]))
	_arg4 = C.gint(len(values))
	_arg3 = (*C.GValue)(C.malloc(C.ulong(len(values)) * C.ulong(C.sizeof_GValue)))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice(_arg3, len(values))
		for i := range values {
			out[i] = *(*C.GValue)(unsafe.Pointer(&(&values[i]).GValue))
		}
	}

	C.gtk_tree_store_set_valuesv(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// Swap swaps @a and @b in the same level of @tree_store. Note that this
// function only works with unsorted stores.
func (treeStore *TreeStore) Swap(a *TreeIter, b *TreeIter) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(a))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(b))

	C.gtk_tree_store_swap(_arg0, _arg1, _arg2)
}
