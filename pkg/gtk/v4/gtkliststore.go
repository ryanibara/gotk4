// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GTypeListStore returns the GType for the type ListStore.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeListStore() coreglib.Type {
	gtype := coreglib.Type(C.gtk_list_store_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalListStore)
	return gtype
}

// ListStoreOverrider contains methods that are overridable.
type ListStoreOverrider interface {
}

// ListStore: list-like data structure that can be used with the GtkTreeView
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
//    </object>.
type ListStore struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Buildable
	TreeDragDest
	TreeDragSource
	TreeSortable
}

var (
	_ coreglib.Objector = (*ListStore)(nil)
)

func classInitListStorer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapListStore(obj *coreglib.Object) *ListStore {
	return &ListStore{
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

func marshalListStore(p uintptr) (interface{}, error) {
	return wrapListStore(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewListStore: non-vararg creation function. Used primarily by language
// bindings.
//
// The function takes the following parameters:
//
//    - types: array of #GType types for the columns, from first to last.
//
// The function returns the following values:
//
//    - listStore: new ListStore.
//
func NewListStore(types []coreglib.Type) *ListStore {
	var _arg2 *C.GType // out
	var _arg1 C.int
	var _cret *C.GtkListStore // in

	_arg1 = (C.int)(len(types))
	_arg2 = (*C.GType)(C.calloc(C.size_t(len(types)), C.size_t(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GType)(_arg2), len(types))
		for i := range types {
			out[i] = C.GType(types[i])
		}
	}

	_cret = C.gtk_list_store_newv(_arg1, _arg2)
	runtime.KeepAlive(types)

	var _listStore *ListStore // out

	_listStore = wrapListStore(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _listStore
}

// Append appends a new row to list_store. iter will be changed to point to this
// new row. The row will be empty after this function is called. To fill in
// values, you need to call gtk_list_store_set() or gtk_list_store_set_value().
//
// The function returns the following values:
//
//    - iter: unset TreeIter to set to the appended row.
//
func (listStore *ListStore) Append() *TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))

	C.gtk_list_store_append(_arg0, &_arg1)
	runtime.KeepAlive(listStore)

	var _iter *TreeIter // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// Clear removes all rows from the list store.
func (listStore *ListStore) Clear() {
	var _arg0 *C.GtkListStore // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))

	C.gtk_list_store_clear(_arg0)
	runtime.KeepAlive(listStore)
}

// Insert creates a new row at position. iter will be changed to point to this
// new row. If position is -1 or is larger than the number of rows on the list,
// then the new row will be appended to the list. The row will be empty after
// this function is called. To fill in values, you need to call
// gtk_list_store_set() or gtk_list_store_set_value().
//
// The function takes the following parameters:
//
//    - position to insert the new row, or -1 for last.
//
// The function returns the following values:
//
//    - iter: unset TreeIter to set to the new row.
//
func (listStore *ListStore) Insert(position int) *TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 C.int           // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))
	_arg2 = C.int(position)

	C.gtk_list_store_insert(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(position)

	var _iter *TreeIter // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// InsertAfter inserts a new row after sibling. If sibling is NULL, then the row
// will be prepended to the beginning of the list. iter will be changed to point
// to this new row. The row will be empty after this function is called. To fill
// in values, you need to call gtk_list_store_set() or
// gtk_list_store_set_value().
//
// The function takes the following parameters:
//
//    - sibling (optional): valid TreeIter, or NULL.
//
// The function returns the following values:
//
//    - iter: unset TreeIter to set to the new row.
//
func (listStore *ListStore) InsertAfter(sibling *TreeIter) *TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))
	if sibling != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(sibling)))
	}

	C.gtk_list_store_insert_after(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(sibling)

	var _iter *TreeIter // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// InsertBefore inserts a new row before sibling. If sibling is NULL, then the
// row will be appended to the end of the list. iter will be changed to point to
// this new row. The row will be empty after this function is called. To fill in
// values, you need to call gtk_list_store_set() or gtk_list_store_set_value().
//
// The function takes the following parameters:
//
//    - sibling (optional): valid TreeIter, or NULL.
//
// The function returns the following values:
//
//    - iter: unset TreeIter to set to the new row.
//
func (listStore *ListStore) InsertBefore(sibling *TreeIter) *TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))
	if sibling != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(sibling)))
	}

	C.gtk_list_store_insert_before(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(sibling)

	var _iter *TreeIter // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// InsertWithValues: variant of gtk_list_store_insert_with_values() which takes
// the columns and values as two arrays, instead of varargs.
//
// This function is mainly intended for language-bindings.
//
// The function takes the following parameters:
//
//    - position to insert the new row, or -1 for last.
//    - columns: array of column numbers.
//    - values: array of GValues.
//
// The function returns the following values:
//
//    - iter (optional): unset TreeIter to set to the new row.
//
func (listStore *ListStore) InsertWithValues(position int, columns []int, values []coreglib.Value) *TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 C.int           // out
	var _arg3 *C.int          // out
	var _arg5 C.int
	var _arg4 *C.GValue // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))
	_arg2 = C.int(position)
	_arg5 = (C.int)(len(columns))
	_arg3 = (*C.int)(C.calloc(C.size_t(len(columns)), C.size_t(C.sizeof_int)))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice((*C.int)(_arg3), len(columns))
		for i := range columns {
			out[i] = C.int(columns[i])
		}
	}
	_arg5 = (C.int)(len(values))
	_arg4 = (*C.GValue)(C.calloc(C.size_t(len(values)), C.size_t(C.sizeof_GValue)))
	defer C.free(unsafe.Pointer(_arg4))
	{
		out := unsafe.Slice((*C.GValue)(_arg4), len(values))
		for i := range values {
			out[i] = *(*C.GValue)(unsafe.Pointer((&values[i]).Native()))
		}
	}

	C.gtk_list_store_insert_with_valuesv(_arg0, &_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(position)
	runtime.KeepAlive(columns)
	runtime.KeepAlive(values)

	var _iter *TreeIter // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// IterIsValid: > This function is slow. Only use it for debugging and/or
// testing > purposes.
//
// Checks if the given iter is a valid iter for this ListStore.
//
// The function takes the following parameters:
//
//    - iter: TreeIter.
//
// The function returns the following values:
//
//    - ok: TRUE if the iter is valid, FALSE if the iter is invalid.
//
func (listStore *ListStore) IterIsValid(iter *TreeIter) bool {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_list_store_iter_is_valid(_arg0, _arg1)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MoveAfter moves iter in store to the position after position. Note that this
// function only works with unsorted stores. If position is NULL, iter will be
// moved to the start of the list.
//
// The function takes the following parameters:
//
//    - iter: TreeIter.
//    - position (optional) or NULL.
//
func (store *ListStore) MoveAfter(iter, position *TreeIter) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	if position != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(position)))
	}

	C.gtk_list_store_move_after(_arg0, _arg1, _arg2)
	runtime.KeepAlive(store)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(position)
}

// MoveBefore moves iter in store to the position before position. Note that
// this function only works with unsorted stores. If position is NULL, iter will
// be moved to the end of the list.
//
// The function takes the following parameters:
//
//    - iter: TreeIter.
//    - position (optional) or NULL.
//
func (store *ListStore) MoveBefore(iter, position *TreeIter) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	if position != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(position)))
	}

	C.gtk_list_store_move_before(_arg0, _arg1, _arg2)
	runtime.KeepAlive(store)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(position)
}

// Prepend prepends a new row to list_store. iter will be changed to point to
// this new row. The row will be empty after this function is called. To fill in
// values, you need to call gtk_list_store_set() or gtk_list_store_set_value().
//
// The function returns the following values:
//
//    - iter: unset TreeIter to set to the prepend row.
//
func (listStore *ListStore) Prepend() *TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))

	C.gtk_list_store_prepend(_arg0, &_arg1)
	runtime.KeepAlive(listStore)

	var _iter *TreeIter // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// Remove removes the given row from the list store. After being removed, iter
// is set to be the next valid row, or invalidated if it pointed to the last row
// in list_store.
//
// The function takes the following parameters:
//
//    - iter: valid TreeIter.
//
// The function returns the following values:
//
//    - ok: TRUE if iter is valid, FALSE if not.
//
func (listStore *ListStore) Remove(iter *TreeIter) bool {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_list_store_remove(_arg0, _arg1)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Reorder reorders store to follow the order indicated by new_order. Note that
// this function only works with unsorted stores.
//
// The function takes the following parameters:
//
//    - newOrder: array of integers mapping the new position of each child to its
//      old position before the re-ordering, i.e. new_order[newpos] = oldpos. It
//      must have exactly as many items as the list store’s length.
//
func (store *ListStore) Reorder(newOrder []int) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.int          // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	{
		_arg1 = (*C.int)(C.calloc(C.size_t((len(newOrder) + 1)), C.size_t(C.sizeof_int)))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(newOrder)+1)
			var zero C.int
			out[len(newOrder)] = zero
			for i := range newOrder {
				out[i] = C.int(newOrder[i])
			}
		}
	}

	C.gtk_list_store_reorder(_arg0, _arg1)
	runtime.KeepAlive(store)
	runtime.KeepAlive(newOrder)
}

// SetColumnTypes: this function is meant primarily for #GObjects that inherit
// from ListStore, and should only be used when constructing a new ListStore. It
// will not function after a row has been added, or a method on the TreeModel
// interface is called.
//
// The function takes the following parameters:
//
//    - types: array length n of #GTypes.
//
func (listStore *ListStore) SetColumnTypes(types []coreglib.Type) {
	var _arg0 *C.GtkListStore // out
	var _arg2 *C.GType        // out
	var _arg1 C.int

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))
	_arg1 = (C.int)(len(types))
	_arg2 = (*C.GType)(C.calloc(C.size_t(len(types)), C.size_t(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GType)(_arg2), len(types))
		for i := range types {
			out[i] = C.GType(types[i])
		}
	}

	C.gtk_list_store_set_column_types(_arg0, _arg1, _arg2)
	runtime.KeepAlive(listStore)
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
func (listStore *ListStore) SetValue(iter *TreeIter, column int, value *coreglib.Value) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 C.int           // out
	var _arg3 *C.GValue       // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg2 = C.int(column)
	_arg3 = (*C.GValue)(unsafe.Pointer(value.Native()))

	C.gtk_list_store_set_value(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(column)
	runtime.KeepAlive(value)
}

// Set: variant of gtk_list_store_set_valist() which takes the columns and
// values as two arrays, instead of varargs. This function is mainly intended
// for language-bindings and in case the number of columns to change is not
// known until run-time.
//
// The function takes the following parameters:
//
//    - iter: valid TreeIter for the row being modified.
//    - columns: array of column numbers.
//    - values: array of GValues.
//
func (listStore *ListStore) Set(iter *TreeIter, columns []int, values []coreglib.Value) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.int          // out
	var _arg4 C.int
	var _arg3 *C.GValue // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg4 = (C.int)(len(columns))
	_arg2 = (*C.int)(C.calloc(C.size_t(len(columns)), C.size_t(C.sizeof_int)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.int)(_arg2), len(columns))
		for i := range columns {
			out[i] = C.int(columns[i])
		}
	}
	_arg4 = (C.int)(len(values))
	_arg3 = (*C.GValue)(C.calloc(C.size_t(len(values)), C.size_t(C.sizeof_GValue)))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice((*C.GValue)(_arg3), len(values))
		for i := range values {
			out[i] = *(*C.GValue)(unsafe.Pointer((&values[i]).Native()))
		}
	}

	C.gtk_list_store_set_valuesv(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(columns)
	runtime.KeepAlive(values)
}

// Swap swaps a and b in store. Note that this function only works with unsorted
// stores.
//
// The function takes the following parameters:
//
//    - a: TreeIter.
//    - b: another TreeIter.
//
func (store *ListStore) Swap(a, b *TreeIter) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(a)))
	_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(b)))

	C.gtk_list_store_swap(_arg0, _arg1, _arg2)
	runtime.KeepAlive(store)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
