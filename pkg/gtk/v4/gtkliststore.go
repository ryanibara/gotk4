// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdbool.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_list_store_get_type()), F: marshalListStore},
	})
}

type ListStorePrivate struct {
	native C.GtkListStorePrivate
}

// WrapListStorePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapListStorePrivate(ptr unsafe.Pointer) *ListStorePrivate {
	if ptr == nil {
		return nil
	}

	return (*ListStorePrivate)(ptr)
}

func marshalListStorePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapListStorePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (l *ListStorePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&l.native)
}

// ListStore: a list-like data structure that can be used with the GtkTreeView
//
// The ListStore object is a list model for use with a TreeView widget. It
// implements the TreeModel interface, and consequentialy, can use all of the
// methods available there. It also implements the TreeSortable interface so it
// can be sorted by the view. Finally, it also implements the tree [drag and
// drop][gtk4-GtkTreeView-drag-and-drop] interfaces.
//
// The ListStore can accept most GObject types as a column type, though it can’t
// accept all custom types. Internally, it will keep a copy of data passed in
// (such as a string or a boxed pointer). Columns that accept #GObjects are
// handled a little differently. The ListStore will keep a reference to the
// object instead of copying the value. As a result, if the object is modified,
// it is up to the application writer to call gtk_tree_model_row_changed() to
// emit the TreeModel::row_changed signal. This most commonly affects lists with
// Textures stored.
//
// An example for creating a simple list store:
//
//    <object class="GtkListStore">
//      <columns>
//        <column type="gchararray"/>
//        <column type="gchararray"/>
//        <column type="gint"/>
//      </columns>
//      <data>
//        <row>
//          <col id="0">John</col>
//          <col id="1">Doe</col>
//          <col id="2">25</col>
//        </row>
//        <row>
//          <col id="0">Johan</col>
//          <col id="1">Dahlin</col>
//          <col id="2">50</col>
//        </row>
//      </data>
//    </object>
type ListStore interface {
	gextras.Objector
	Buildable
	TreeDragDest
	TreeDragSource
	TreeModel
	TreeSortable

	// Append appends a new row to @list_store. @iter will be changed to point
	// to this new row. The row will be empty after this function is called. To
	// fill in values, you need to call gtk_list_store_set() or
	// gtk_list_store_set_value().
	Append() TreeIter
	// Clear removes all rows from the list store.
	Clear()
	// Insert creates a new row at @position. @iter will be changed to point to
	// this new row. If @position is -1 or is larger than the number of rows on
	// the list, then the new row will be appended to the list. The row will be
	// empty after this function is called. To fill in values, you need to call
	// gtk_list_store_set() or gtk_list_store_set_value().
	Insert(position int) TreeIter
	// InsertAfter inserts a new row after @sibling. If @sibling is nil, then
	// the row will be prepended to the beginning of the list. @iter will be
	// changed to point to this new row. The row will be empty after this
	// function is called. To fill in values, you need to call
	// gtk_list_store_set() or gtk_list_store_set_value().
	InsertAfter(sibling *TreeIter) TreeIter
	// InsertBefore inserts a new row before @sibling. If @sibling is nil, then
	// the row will be appended to the end of the list. @iter will be changed to
	// point to this new row. The row will be empty after this function is
	// called. To fill in values, you need to call gtk_list_store_set() or
	// gtk_list_store_set_value().
	InsertBefore(sibling *TreeIter) TreeIter
	// InsertWithValuesv: a variant of gtk_list_store_insert_with_values() which
	// takes the columns and values as two arrays, instead of varargs.
	//
	// This function is mainly intended for language-bindings.
	InsertWithValuesv(position int, columns []int, values []*externglib.Value) TreeIter
	// IterIsValid: > This function is slow. Only use it for debugging and/or
	// testing > purposes.
	//
	// Checks if the given iter is a valid iter for this ListStore.
	IterIsValid(iter *TreeIter) bool
	// MoveAfter moves @iter in @store to the position after @position. Note
	// that this function only works with unsorted stores. If @position is nil,
	// @iter will be moved to the start of the list.
	MoveAfter(iter *TreeIter, position *TreeIter)
	// MoveBefore moves @iter in @store to the position before @position. Note
	// that this function only works with unsorted stores. If @position is nil,
	// @iter will be moved to the end of the list.
	MoveBefore(iter *TreeIter, position *TreeIter)
	// Prepend prepends a new row to @list_store. @iter will be changed to point
	// to this new row. The row will be empty after this function is called. To
	// fill in values, you need to call gtk_list_store_set() or
	// gtk_list_store_set_value().
	Prepend() TreeIter
	// Remove removes the given row from the list store. After being removed,
	// @iter is set to be the next valid row, or invalidated if it pointed to
	// the last row in @list_store.
	Remove(iter *TreeIter) bool
	// Reorder reorders @store to follow the order indicated by @new_order. Note
	// that this function only works with unsorted stores.
	Reorder(newOrder []int)
	// SetColumnTypes: this function is meant primarily for #GObjects that
	// inherit from ListStore, and should only be used when constructing a new
	// ListStore. It will not function after a row has been added, or a method
	// on the TreeModel interface is called.
	SetColumnTypes(nColumns int, types []externglib.Type)
	// SetValue sets the data in the cell specified by @iter and @column. The
	// type of @value must be convertible to the type of the column.
	SetValue(iter *TreeIter, column int, value *externglib.Value)
	// SetValuesv: a variant of gtk_list_store_set_valist() which takes the
	// columns and values as two arrays, instead of varargs. This function is
	// mainly intended for language-bindings and in case the number of columns
	// to change is not known until run-time.
	SetValuesv(iter *TreeIter, columns []int, values []*externglib.Value)
	// Swap swaps @a and @b in @store. Note that this function only works with
	// unsorted stores.
	Swap(a *TreeIter, b *TreeIter)
}

// listStore implements the ListStore interface.
type listStore struct {
	gextras.Objector
	Buildable
	TreeDragDest
	TreeDragSource
	TreeModel
	TreeSortable
}

var _ ListStore = (*listStore)(nil)

// WrapListStore wraps a GObject to the right type. It is
// primarily used internally.
func WrapListStore(obj *externglib.Object) ListStore {
	return ListStore{
		Objector:       obj,
		Buildable:      WrapBuildable(obj),
		TreeDragDest:   WrapTreeDragDest(obj),
		TreeDragSource: WrapTreeDragSource(obj),
		TreeModel:      WrapTreeModel(obj),
		TreeSortable:   WrapTreeSortable(obj),
	}
}

func marshalListStore(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapListStore(obj), nil
}

// NewListStoreV constructs a class ListStore.
func NewListStoreV(nColumns int, types []externglib.Type) ListStore {
	var arg1 C.int
	var arg2 *C.GType

	{
		var dst []C.GType
		ptr := C.malloc(C.sizeof_GType * len(types))
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&dst))
		sliceHeader.Data = uintptr(unsafe.Pointer(ptr))
		sliceHeader.Len = len(types)
		sliceHeader.Cap = len(types)
		defer C.free(unsafe.Pointer(ptr))

		for i := 0; i < len(types); i++ {
			src := types[i]
			dst[i] = C.GType(src)
		}

		arg2 = (*C.GType)(unsafe.Pointer(ptr))
		arg1 = len(types)
	}

	ret := C.gtk_list_store_newv(arg1, arg2)

	var ret0 ListStore

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(ListStore)

	return ret0
}

// Append appends a new row to @list_store. @iter will be changed to point
// to this new row. The row will be empty after this function is called. To
// fill in values, you need to call gtk_list_store_set() or
// gtk_list_store_set_value().
func (l listStore) Append() TreeIter {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter // out

	arg0 = (*C.GtkListStore)(l.Native())

	C.gtk_list_store_append(arg0, &arg1)

	var ret0 *TreeIter

	{
		ret0 = WrapTreeIter(unsafe.Pointer(arg1))
	}

	return ret0
}

// Clear removes all rows from the list store.
func (l listStore) Clear() {
	var arg0 *C.GtkListStore

	arg0 = (*C.GtkListStore)(l.Native())

	C.gtk_list_store_clear(arg0)
}

// Insert creates a new row at @position. @iter will be changed to point to
// this new row. If @position is -1 or is larger than the number of rows on
// the list, then the new row will be appended to the list. The row will be
// empty after this function is called. To fill in values, you need to call
// gtk_list_store_set() or gtk_list_store_set_value().
func (l listStore) Insert(position int) TreeIter {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter // out
	var arg2 C.int

	arg0 = (*C.GtkListStore)(l.Native())
	arg2 = C.int(position)

	C.gtk_list_store_insert(arg0, &arg1, arg2)

	var ret0 *TreeIter

	{
		ret0 = WrapTreeIter(unsafe.Pointer(arg1))
	}

	return ret0
}

// InsertAfter inserts a new row after @sibling. If @sibling is nil, then
// the row will be prepended to the beginning of the list. @iter will be
// changed to point to this new row. The row will be empty after this
// function is called. To fill in values, you need to call
// gtk_list_store_set() or gtk_list_store_set_value().
func (l listStore) InsertAfter(sibling *TreeIter) TreeIter {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter // out
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkListStore)(l.Native())
	arg2 = (*C.GtkTreeIter)(sibling.Native())

	C.gtk_list_store_insert_after(arg0, &arg1, arg2)

	var ret0 *TreeIter

	{
		ret0 = WrapTreeIter(unsafe.Pointer(arg1))
	}

	return ret0
}

// InsertBefore inserts a new row before @sibling. If @sibling is nil, then
// the row will be appended to the end of the list. @iter will be changed to
// point to this new row. The row will be empty after this function is
// called. To fill in values, you need to call gtk_list_store_set() or
// gtk_list_store_set_value().
func (l listStore) InsertBefore(sibling *TreeIter) TreeIter {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter // out
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkListStore)(l.Native())
	arg2 = (*C.GtkTreeIter)(sibling.Native())

	C.gtk_list_store_insert_before(arg0, &arg1, arg2)

	var ret0 *TreeIter

	{
		ret0 = WrapTreeIter(unsafe.Pointer(arg1))
	}

	return ret0
}

// InsertWithValuesv: a variant of gtk_list_store_insert_with_values() which
// takes the columns and values as two arrays, instead of varargs.
//
// This function is mainly intended for language-bindings.
func (l listStore) InsertWithValuesv(position int, columns []int, values []*externglib.Value) TreeIter {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter // out
	var arg2 C.int
	var arg3 *C.int
	var arg4 *C.GValue
	var arg5 C.int

	arg0 = (*C.GtkListStore)(l.Native())
	arg2 = C.int(position)
	arg3 = (*C.int)(unsafe.Pointer(&columns[0]))
	arg5 = len(columns)
	defer runtime.KeepAlive(columns)
	{
		var dst []C.GValue
		ptr := C.malloc(C.sizeof_GValue * len(values))
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&dst))
		sliceHeader.Data = uintptr(unsafe.Pointer(ptr))
		sliceHeader.Len = len(values)
		sliceHeader.Cap = len(values)
		defer C.free(unsafe.Pointer(ptr))

		for i := 0; i < len(values); i++ {
			src := values[i]
			dst[i] = (*C.GValue)(src.GValue)
		}

		arg4 = (*C.GValue)(unsafe.Pointer(ptr))
		arg5 = len(values)
	}

	C.gtk_list_store_insert_with_valuesv(arg0, &arg1, arg2, arg3, arg4, arg5)

	var ret0 *TreeIter

	{
		ret0 = WrapTreeIter(unsafe.Pointer(arg1))
	}

	return ret0
}

// IterIsValid: > This function is slow. Only use it for debugging and/or
// testing > purposes.
//
// Checks if the given iter is a valid iter for this ListStore.
func (l listStore) IterIsValid(iter *TreeIter) bool {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter

	arg0 = (*C.GtkListStore)(l.Native())
	arg1 = (*C.GtkTreeIter)(iter.Native())

	ret := C.gtk_list_store_iter_is_valid(arg0, arg1)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// MoveAfter moves @iter in @store to the position after @position. Note
// that this function only works with unsorted stores. If @position is nil,
// @iter will be moved to the start of the list.
func (s listStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkListStore)(s.Native())
	arg1 = (*C.GtkTreeIter)(iter.Native())
	arg2 = (*C.GtkTreeIter)(position.Native())

	C.gtk_list_store_move_after(arg0, arg1, arg2)
}

// MoveBefore moves @iter in @store to the position before @position. Note
// that this function only works with unsorted stores. If @position is nil,
// @iter will be moved to the end of the list.
func (s listStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkListStore)(s.Native())
	arg1 = (*C.GtkTreeIter)(iter.Native())
	arg2 = (*C.GtkTreeIter)(position.Native())

	C.gtk_list_store_move_before(arg0, arg1, arg2)
}

// Prepend prepends a new row to @list_store. @iter will be changed to point
// to this new row. The row will be empty after this function is called. To
// fill in values, you need to call gtk_list_store_set() or
// gtk_list_store_set_value().
func (l listStore) Prepend() TreeIter {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter // out

	arg0 = (*C.GtkListStore)(l.Native())

	C.gtk_list_store_prepend(arg0, &arg1)

	var ret0 *TreeIter

	{
		ret0 = WrapTreeIter(unsafe.Pointer(arg1))
	}

	return ret0
}

// Remove removes the given row from the list store. After being removed,
// @iter is set to be the next valid row, or invalidated if it pointed to
// the last row in @list_store.
func (l listStore) Remove(iter *TreeIter) bool {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter

	arg0 = (*C.GtkListStore)(l.Native())
	arg1 = (*C.GtkTreeIter)(iter.Native())

	ret := C.gtk_list_store_remove(arg0, arg1)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Reorder reorders @store to follow the order indicated by @new_order. Note
// that this function only works with unsorted stores.
func (s listStore) Reorder(newOrder []int) {
	var arg0 *C.GtkListStore
	var arg1 *C.int

	arg0 = (*C.GtkListStore)(s.Native())
	{
		var dst []C.int
		ptr := C.malloc(C.sizeof_int * (len(newOrder) + 1))
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&dst))
		sliceHeader.Data = uintptr(unsafe.Pointer(ptr))
		sliceHeader.Len = len(newOrder)
		sliceHeader.Cap = len(newOrder)
		defer C.free(unsafe.Pointer(ptr))

		for i := 0; i < len(newOrder); i++ {
			src := newOrder[i]
			dst[i] = C.int(src)
		}

		arg1 = (*C.int)(unsafe.Pointer(ptr))
	}

	C.gtk_list_store_reorder(arg0, arg1)
}

// SetColumnTypes: this function is meant primarily for #GObjects that
// inherit from ListStore, and should only be used when constructing a new
// ListStore. It will not function after a row has been added, or a method
// on the TreeModel interface is called.
func (l listStore) SetColumnTypes(nColumns int, types []externglib.Type) {
	var arg0 *C.GtkListStore
	var arg1 C.int
	var arg2 *C.GType

	arg0 = (*C.GtkListStore)(l.Native())
	{
		var dst []C.GType
		ptr := C.malloc(C.sizeof_GType * len(types))
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&dst))
		sliceHeader.Data = uintptr(unsafe.Pointer(ptr))
		sliceHeader.Len = len(types)
		sliceHeader.Cap = len(types)
		defer C.free(unsafe.Pointer(ptr))

		for i := 0; i < len(types); i++ {
			src := types[i]
			dst[i] = C.GType(src)
		}

		arg2 = (*C.GType)(unsafe.Pointer(ptr))
		arg1 = len(types)
	}

	C.gtk_list_store_set_column_types(arg0, arg1, arg2)
}

// SetValue sets the data in the cell specified by @iter and @column. The
// type of @value must be convertible to the type of the column.
func (l listStore) SetValue(iter *TreeIter, column int, value *externglib.Value) {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter
	var arg2 C.int
	var arg3 *C.GValue

	arg0 = (*C.GtkListStore)(l.Native())
	arg1 = (*C.GtkTreeIter)(iter.Native())
	arg2 = C.int(column)
	arg3 = (*C.GValue)(value.GValue)

	C.gtk_list_store_set_value(arg0, arg1, arg2, arg3)
}

// SetValuesv: a variant of gtk_list_store_set_valist() which takes the
// columns and values as two arrays, instead of varargs. This function is
// mainly intended for language-bindings and in case the number of columns
// to change is not known until run-time.
func (l listStore) SetValuesv(iter *TreeIter, columns []int, values []*externglib.Value) {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter
	var arg2 *C.int
	var arg3 *C.GValue
	var arg4 C.int

	arg0 = (*C.GtkListStore)(l.Native())
	arg1 = (*C.GtkTreeIter)(iter.Native())
	arg2 = (*C.int)(unsafe.Pointer(&columns[0]))
	arg4 = len(columns)
	defer runtime.KeepAlive(columns)
	{
		var dst []C.GValue
		ptr := C.malloc(C.sizeof_GValue * len(values))
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&dst))
		sliceHeader.Data = uintptr(unsafe.Pointer(ptr))
		sliceHeader.Len = len(values)
		sliceHeader.Cap = len(values)
		defer C.free(unsafe.Pointer(ptr))

		for i := 0; i < len(values); i++ {
			src := values[i]
			dst[i] = (*C.GValue)(src.GValue)
		}

		arg3 = (*C.GValue)(unsafe.Pointer(ptr))
		arg4 = len(values)
	}

	C.gtk_list_store_set_valuesv(arg0, arg1, arg2, arg3, arg4)
}

// Swap swaps @a and @b in @store. Note that this function only works with
// unsorted stores.
func (s listStore) Swap(a *TreeIter, b *TreeIter) {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkListStore)(s.Native())
	arg1 = (*C.GtkTreeIter)(a.Native())
	arg2 = (*C.GtkTreeIter)(b.Native())

	C.gtk_list_store_swap(arg0, arg1, arg2)
}
