// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_store_get_type()), F: marshalTreeStore},
	})
}

// TreeStore: the TreeStore object is a list model for use with a TreeView
// widget. It implements the TreeModel interface, and consequentially, can use
// all of the methods available there. It also implements the TreeSortable
// interface so it can be sorted by the view. Finally, it also implements the
// tree [drag and drop][gtk3-GtkTreeView-drag-and-drop] interfaces.
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
type TreeStore interface {
	gextras.Objector
	Buildable
	TreeDragDest
	TreeDragSource
	TreeModel
	TreeSortable

	// Append appends a new row to @tree_store. If @parent is non-nil, then it
	// will append the new row after the last child of @parent, otherwise it
	// will append a row to the top level. @iter will be changed to point to
	// this new row. The row will be empty after this function is called. To
	// fill in values, you need to call gtk_tree_store_set() or
	// gtk_tree_store_set_value().
	Append(parent *TreeIter) TreeIter
	// Clear removes all rows from @tree_store
	Clear()
	// Insert creates a new row at @position. If parent is non-nil, then the row
	// will be made a child of @parent. Otherwise, the row will be created at
	// the toplevel. If @position is -1 or is larger than the number of rows at
	// that level, then the new row will be inserted to the end of the list.
	// @iter will be changed to point to this new row. The row will be empty
	// after this function is called. To fill in values, you need to call
	// gtk_tree_store_set() or gtk_tree_store_set_value().
	Insert(parent *TreeIter, position int) TreeIter
	// InsertAfter inserts a new row after @sibling. If @sibling is nil, then
	// the row will be prepended to @parent ’s children. If @parent and @sibling
	// are nil, then the row will be prepended to the toplevel. If both @sibling
	// and @parent are set, then @parent must be the parent of @sibling. When
	// @sibling is set, @parent is optional.
	//
	// @iter will be changed to point to this new row. The row will be empty
	// after this function is called. To fill in values, you need to call
	// gtk_tree_store_set() or gtk_tree_store_set_value().
	InsertAfter(parent *TreeIter, sibling *TreeIter) TreeIter
	// InsertBefore inserts a new row before @sibling. If @sibling is nil, then
	// the row will be appended to @parent ’s children. If @parent and @sibling
	// are nil, then the row will be appended to the toplevel. If both @sibling
	// and @parent are set, then @parent must be the parent of @sibling. When
	// @sibling is set, @parent is optional.
	//
	// @iter will be changed to point to this new row. The row will be empty
	// after this function is called. To fill in values, you need to call
	// gtk_tree_store_set() or gtk_tree_store_set_value().
	InsertBefore(parent *TreeIter, sibling *TreeIter) TreeIter
	// InsertWithValuesv: a variant of gtk_tree_store_insert_with_values() which
	// takes the columns and values as two arrays, instead of varargs. This
	// function is mainly intended for language bindings.
	InsertWithValuesv(parent *TreeIter, position int, columns []int, values []*externglib.Value) TreeIter
	// IsAncestor returns true if @iter is an ancestor of @descendant. That is,
	// @iter is the parent (or grandparent or great-grandparent) of @descendant.
	IsAncestor(iter *TreeIter, descendant *TreeIter) bool
	// IterDepth returns the depth of @iter. This will be 0 for anything on the
	// root level, 1 for anything down a level, etc.
	IterDepth(iter *TreeIter) int
	// IterIsValid: WARNING: This function is slow. Only use it for debugging
	// and/or testing purposes.
	//
	// Checks if the given iter is a valid iter for this TreeStore.
	IterIsValid(iter *TreeIter) bool
	// MoveAfter moves @iter in @tree_store to the position after @position.
	// @iter and @position should be in the same level. Note that this function
	// only works with unsorted stores. If @position is nil, @iter will be moved
	// to the start of the level.
	MoveAfter(iter *TreeIter, position *TreeIter)
	// MoveBefore moves @iter in @tree_store to the position before @position.
	// @iter and @position should be in the same level. Note that this function
	// only works with unsorted stores. If @position is nil, @iter will be moved
	// to the end of the level.
	MoveBefore(iter *TreeIter, position *TreeIter)
	// Prepend prepends a new row to @tree_store. If @parent is non-nil, then it
	// will prepend the new row before the first child of @parent, otherwise it
	// will prepend a row to the top level. @iter will be changed to point to
	// this new row. The row will be empty after this function is called. To
	// fill in values, you need to call gtk_tree_store_set() or
	// gtk_tree_store_set_value().
	Prepend(parent *TreeIter) TreeIter
	// Remove removes @iter from @tree_store. After being removed, @iter is set
	// to the next valid row at that level, or invalidated if it previously
	// pointed to the last one.
	Remove(iter *TreeIter) bool
	// SetValue sets the data in the cell specified by @iter and @column. The
	// type of @value must be convertible to the type of the column.
	SetValue(iter *TreeIter, column int, value *externglib.Value)
	// SetValuesv: a variant of gtk_tree_store_set_valist() which takes the
	// columns and values as two arrays, instead of varargs. This function is
	// mainly intended for language bindings or in case the number of columns to
	// change is not known until run-time.
	SetValuesv(iter *TreeIter, columns []int, values []*externglib.Value)
	// Swap swaps @a and @b in the same level of @tree_store. Note that this
	// function only works with unsorted stores.
	Swap(a *TreeIter, b *TreeIter)
}

// treeStore implements the TreeStore interface.
type treeStore struct {
	gextras.Objector
	Buildable
	TreeDragDest
	TreeDragSource
	TreeModel
	TreeSortable
}

var _ TreeStore = (*treeStore)(nil)

// WrapTreeStore wraps a GObject to the right type. It is
// primarily used internally.
func WrapTreeStore(obj *externglib.Object) TreeStore {
	return TreeStore{
		Objector:       obj,
		Buildable:      WrapBuildable(obj),
		TreeDragDest:   WrapTreeDragDest(obj),
		TreeDragSource: WrapTreeDragSource(obj),
		TreeModel:      WrapTreeModel(obj),
		TreeSortable:   WrapTreeSortable(obj),
	}
}

func marshalTreeStore(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeStore(obj), nil
}

// Append appends a new row to @tree_store. If @parent is non-nil, then it
// will append the new row after the last child of @parent, otherwise it
// will append a row to the top level. @iter will be changed to point to
// this new row. The row will be empty after this function is called. To
// fill in values, you need to call gtk_tree_store_set() or
// gtk_tree_store_set_value().
func (t treeStore) Append(parent *TreeIter) TreeIter {
	var arg0 *C.GtkTreeStore
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent.Native()))

	var arg1 C.GtkTreeIter
	var ret1 *TreeIter

	C.gtk_tree_store_append(arg0, &arg1, parent)

	ret1 = WrapTreeIter(unsafe.Pointer(arg1))

	return ret1
}

// Clear removes all rows from @tree_store
func (t treeStore) Clear() {
	var arg0 *C.GtkTreeStore

	arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))

	C.gtk_tree_store_clear(arg0)
}

// Insert creates a new row at @position. If parent is non-nil, then the row
// will be made a child of @parent. Otherwise, the row will be created at
// the toplevel. If @position is -1 or is larger than the number of rows at
// that level, then the new row will be inserted to the end of the list.
// @iter will be changed to point to this new row. The row will be empty
// after this function is called. To fill in values, you need to call
// gtk_tree_store_set() or gtk_tree_store_set_value().
func (t treeStore) Insert(parent *TreeIter, position int) TreeIter {
	var arg0 *C.GtkTreeStore
	var arg2 *C.GtkTreeIter
	var arg3 C.gint

	arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent.Native()))
	arg3 = C.gint(position)

	var arg1 C.GtkTreeIter
	var ret1 *TreeIter

	C.gtk_tree_store_insert(arg0, &arg1, parent, position)

	ret1 = WrapTreeIter(unsafe.Pointer(arg1))

	return ret1
}

// InsertAfter inserts a new row after @sibling. If @sibling is nil, then
// the row will be prepended to @parent ’s children. If @parent and @sibling
// are nil, then the row will be prepended to the toplevel. If both @sibling
// and @parent are set, then @parent must be the parent of @sibling. When
// @sibling is set, @parent is optional.
//
// @iter will be changed to point to this new row. The row will be empty
// after this function is called. To fill in values, you need to call
// gtk_tree_store_set() or gtk_tree_store_set_value().
func (t treeStore) InsertAfter(parent *TreeIter, sibling *TreeIter) TreeIter {
	var arg0 *C.GtkTreeStore
	var arg2 *C.GtkTreeIter
	var arg3 *C.GtkTreeIter

	arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent.Native()))
	arg3 = (*C.GtkTreeIter)(unsafe.Pointer(sibling.Native()))

	var arg1 C.GtkTreeIter
	var ret1 *TreeIter

	C.gtk_tree_store_insert_after(arg0, &arg1, parent, sibling)

	ret1 = WrapTreeIter(unsafe.Pointer(arg1))

	return ret1
}

// InsertBefore inserts a new row before @sibling. If @sibling is nil, then
// the row will be appended to @parent ’s children. If @parent and @sibling
// are nil, then the row will be appended to the toplevel. If both @sibling
// and @parent are set, then @parent must be the parent of @sibling. When
// @sibling is set, @parent is optional.
//
// @iter will be changed to point to this new row. The row will be empty
// after this function is called. To fill in values, you need to call
// gtk_tree_store_set() or gtk_tree_store_set_value().
func (t treeStore) InsertBefore(parent *TreeIter, sibling *TreeIter) TreeIter {
	var arg0 *C.GtkTreeStore
	var arg2 *C.GtkTreeIter
	var arg3 *C.GtkTreeIter

	arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent.Native()))
	arg3 = (*C.GtkTreeIter)(unsafe.Pointer(sibling.Native()))

	var arg1 C.GtkTreeIter
	var ret1 *TreeIter

	C.gtk_tree_store_insert_before(arg0, &arg1, parent, sibling)

	ret1 = WrapTreeIter(unsafe.Pointer(arg1))

	return ret1
}

// InsertWithValuesv: a variant of gtk_tree_store_insert_with_values() which
// takes the columns and values as two arrays, instead of varargs. This
// function is mainly intended for language bindings.
func (t treeStore) InsertWithValuesv(parent *TreeIter, position int, columns []int, values []*externglib.Value) TreeIter {
	var arg0 *C.GtkTreeStore

	arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))

	var arg1 C.GtkTreeIter
	var ret1 *TreeIter

	C.gtk_tree_store_insert_with_valuesv(arg0, &arg1, parent, position, columns, values, nValues)

	ret1 = WrapTreeIter(unsafe.Pointer(arg1))

	return ret1
}

// IsAncestor returns true if @iter is an ancestor of @descendant. That is,
// @iter is the parent (or grandparent or great-grandparent) of @descendant.
func (t treeStore) IsAncestor(iter *TreeIter, descendant *TreeIter) bool {
	var arg0 *C.GtkTreeStore
	var arg1 *C.GtkTreeIter
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(descendant.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_tree_store_is_ancestor(arg0, iter, descendant)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// IterDepth returns the depth of @iter. This will be 0 for anything on the
// root level, 1 for anything down a level, etc.
func (t treeStore) IterDepth(iter *TreeIter) int {
	var arg0 *C.GtkTreeStore
	var arg1 *C.GtkTreeIter

	arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	var cret C.gint
	var ret1 int

	cret = C.gtk_tree_store_iter_depth(arg0, iter)

	ret1 = C.gint(cret)

	return ret1
}

// IterIsValid: WARNING: This function is slow. Only use it for debugging
// and/or testing purposes.
//
// Checks if the given iter is a valid iter for this TreeStore.
func (t treeStore) IterIsValid(iter *TreeIter) bool {
	var arg0 *C.GtkTreeStore
	var arg1 *C.GtkTreeIter

	arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_tree_store_iter_is_valid(arg0, iter)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// MoveAfter moves @iter in @tree_store to the position after @position.
// @iter and @position should be in the same level. Note that this function
// only works with unsorted stores. If @position is nil, @iter will be moved
// to the start of the level.
func (t treeStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	var arg0 *C.GtkTreeStore
	var arg1 *C.GtkTreeIter
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(position.Native()))

	C.gtk_tree_store_move_after(arg0, iter, position)
}

// MoveBefore moves @iter in @tree_store to the position before @position.
// @iter and @position should be in the same level. Note that this function
// only works with unsorted stores. If @position is nil, @iter will be moved
// to the end of the level.
func (t treeStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	var arg0 *C.GtkTreeStore
	var arg1 *C.GtkTreeIter
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(position.Native()))

	C.gtk_tree_store_move_before(arg0, iter, position)
}

// Prepend prepends a new row to @tree_store. If @parent is non-nil, then it
// will prepend the new row before the first child of @parent, otherwise it
// will prepend a row to the top level. @iter will be changed to point to
// this new row. The row will be empty after this function is called. To
// fill in values, you need to call gtk_tree_store_set() or
// gtk_tree_store_set_value().
func (t treeStore) Prepend(parent *TreeIter) TreeIter {
	var arg0 *C.GtkTreeStore
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent.Native()))

	var arg1 C.GtkTreeIter
	var ret1 *TreeIter

	C.gtk_tree_store_prepend(arg0, &arg1, parent)

	ret1 = WrapTreeIter(unsafe.Pointer(arg1))

	return ret1
}

// Remove removes @iter from @tree_store. After being removed, @iter is set
// to the next valid row at that level, or invalidated if it previously
// pointed to the last one.
func (t treeStore) Remove(iter *TreeIter) bool {
	var arg0 *C.GtkTreeStore
	var arg1 *C.GtkTreeIter

	arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_tree_store_remove(arg0, iter)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// SetValue sets the data in the cell specified by @iter and @column. The
// type of @value must be convertible to the type of the column.
func (t treeStore) SetValue(iter *TreeIter, column int, value *externglib.Value) {
	var arg0 *C.GtkTreeStore
	var arg1 *C.GtkTreeIter
	var arg2 C.gint
	var arg3 *C.GValue

	arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))
	arg2 = C.gint(column)
	arg3 = (*C.GValue)(value.GValue)

	C.gtk_tree_store_set_value(arg0, iter, column, value)
}

// SetValuesv: a variant of gtk_tree_store_set_valist() which takes the
// columns and values as two arrays, instead of varargs. This function is
// mainly intended for language bindings or in case the number of columns to
// change is not known until run-time.
func (t treeStore) SetValuesv(iter *TreeIter, columns []int, values []*externglib.Value) {
	var arg0 *C.GtkTreeStore

	arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))

	C.gtk_tree_store_set_valuesv(arg0, iter, columns, values, nValues)
}

// Swap swaps @a and @b in the same level of @tree_store. Note that this
// function only works with unsorted stores.
func (t treeStore) Swap(a *TreeIter, b *TreeIter) {
	var arg0 *C.GtkTreeStore
	var arg1 *C.GtkTreeIter
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(a.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(b.Native()))

	C.gtk_tree_store_swap(arg0, a, b)
}

type TreeStorePrivate struct {
	native C.GtkTreeStorePrivate
}

// WrapTreeStorePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTreeStorePrivate(ptr unsafe.Pointer) *TreeStorePrivate {
	if ptr == nil {
		return nil
	}

	return (*TreeStorePrivate)(ptr)
}

func marshalTreeStorePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTreeStorePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TreeStorePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}
