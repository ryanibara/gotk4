// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
//
// gboolean gotk4_TreeModelForeachFunc(GtkTreeModel*, GtkTreePath*, GtkTreeIter*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_list_store_get_type()), F: marshalListStore},
	})
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
	Buildable
	TreeDragDest
	TreeDragSource
	TreeSortable

	AppendListStore() TreeIter

	ClearListStore()

	InsertListStore(position int) TreeIter

	InsertAfterListStore(sibling *TreeIter) TreeIter

	InsertBeforeListStore(sibling *TreeIter) TreeIter

	InsertWithValuesvListStore(position int, columns []int, values []externglib.Value) TreeIter

	IterIsValidListStore(iter *TreeIter) bool

	MoveAfterListStore(iter *TreeIter, position *TreeIter)

	MoveBeforeListStore(iter *TreeIter, position *TreeIter)

	PrependListStore() TreeIter

	RemoveListStore(iter *TreeIter) bool

	ReorderListStore(newOrder []int)

	SetColumnTypesListStore(types []externglib.Type)

	SetValueListStore(iter *TreeIter, column int, value externglib.Value)

	SetValuesvListStore(iter *TreeIter, columns []int, values []externglib.Value)

	SwapListStore(a *TreeIter, b *TreeIter)
}

// listStore implements the ListStore class.
type listStore struct {
	gextras.Objector
}

// WrapListStore wraps a GObject to the right type. It is
// primarily used internally.
func WrapListStore(obj *externglib.Object) ListStore {
	return listStore{
		Objector: obj,
	}
}

func marshalListStore(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapListStore(obj), nil
}

func NewListStoreV(types []externglib.Type) ListStore {
	var _arg2 *C.GType
	var _arg1 C.int
	var _cret *C.GtkListStore // in

	_arg1 = C.int(len(types))
	_arg2 = (*C.GType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(types))
		for i := range types {
			out[i] = (C.GType)(types[i])
		}
	}

	_cret = C.gtk_list_store_newv(_arg1, _arg2)

	var _listStore ListStore // out

	_listStore = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(ListStore)

	return _listStore
}

func (l listStore) AppendListStore() TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))

	C.gtk_list_store_append(_arg0, &_arg1)

	var _iter TreeIter // out

	{
		var refTmpIn *C.GtkTreeIter
		var refTmpOut *TreeIter

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*TreeIter)(unsafe.Pointer(refTmpIn))

		_iter = *refTmpOut
	}

	return _iter
}

func (l listStore) ClearListStore() {
	var _arg0 *C.GtkListStore // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))

	C.gtk_list_store_clear(_arg0)
}

func (l listStore) InsertListStore(position int) TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 C.int           // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))
	_arg2 = C.int(position)

	C.gtk_list_store_insert(_arg0, &_arg1, _arg2)

	var _iter TreeIter // out

	{
		var refTmpIn *C.GtkTreeIter
		var refTmpOut *TreeIter

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*TreeIter)(unsafe.Pointer(refTmpIn))

		_iter = *refTmpOut
	}

	return _iter
}

func (l listStore) InsertAfterListStore(sibling *TreeIter) TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(sibling.Native()))

	C.gtk_list_store_insert_after(_arg0, &_arg1, _arg2)

	var _iter TreeIter // out

	{
		var refTmpIn *C.GtkTreeIter
		var refTmpOut *TreeIter

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*TreeIter)(unsafe.Pointer(refTmpIn))

		_iter = *refTmpOut
	}

	return _iter
}

func (l listStore) InsertBeforeListStore(sibling *TreeIter) TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(sibling.Native()))

	C.gtk_list_store_insert_before(_arg0, &_arg1, _arg2)

	var _iter TreeIter // out

	{
		var refTmpIn *C.GtkTreeIter
		var refTmpOut *TreeIter

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*TreeIter)(unsafe.Pointer(refTmpIn))

		_iter = *refTmpOut
	}

	return _iter
}

func (l listStore) InsertWithValuesvListStore(position int, columns []int, values []externglib.Value) TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 C.int           // out
	var _arg3 *C.int
	var _arg5 C.int
	var _arg4 *C.GValue
	var _arg5 C.int

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))
	_arg2 = C.int(position)
	_arg5 = C.int(len(columns))
	_arg3 = (*C.int)(unsafe.Pointer(&columns[0]))
	_arg5 = C.int(len(values))
	_arg4 = (*C.GValue)(C.malloc(C.ulong(len(values)) * C.ulong(C.sizeof_GValue)))
	defer C.free(unsafe.Pointer(_arg4))
	{
		out := unsafe.Slice(_arg4, len(values))
		for i := range values {
			{
				var refTmpIn *externglib.Value
				var refTmpOut *C.GValue

				in0 := &values[i]
				refTmpIn = in0

				refTmpOut = (*C.GValue)(unsafe.Pointer(&refTmpIn.GValue))

				out[i] = *refTmpOut
			}
		}
	}

	C.gtk_list_store_insert_with_valuesv(_arg0, &_arg1, _arg2, _arg3, _arg4, _arg5)

	var _iter TreeIter // out

	{
		var refTmpIn *C.GtkTreeIter
		var refTmpOut *TreeIter

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*TreeIter)(unsafe.Pointer(refTmpIn))

		_iter = *refTmpOut
	}

	return _iter
}

func (l listStore) IterIsValidListStore(iter *TreeIter) bool {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	_cret = C.gtk_list_store_iter_is_valid(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s listStore) MoveAfterListStore(iter *TreeIter, position *TreeIter) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(position.Native()))

	C.gtk_list_store_move_after(_arg0, _arg1, _arg2)
}

func (s listStore) MoveBeforeListStore(iter *TreeIter, position *TreeIter) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(position.Native()))

	C.gtk_list_store_move_before(_arg0, _arg1, _arg2)
}

func (l listStore) PrependListStore() TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))

	C.gtk_list_store_prepend(_arg0, &_arg1)

	var _iter TreeIter // out

	{
		var refTmpIn *C.GtkTreeIter
		var refTmpOut *TreeIter

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*TreeIter)(unsafe.Pointer(refTmpIn))

		_iter = *refTmpOut
	}

	return _iter
}

func (l listStore) RemoveListStore(iter *TreeIter) bool {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	_cret = C.gtk_list_store_remove(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s listStore) ReorderListStore(newOrder []int) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.int

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(s.Native()))
	{
		var zero int
		newOrder = append(newOrder, zero)
	}
	_arg1 = (*C.int)(unsafe.Pointer(&newOrder[0]))

	C.gtk_list_store_reorder(_arg0, _arg1)
}

func (l listStore) SetColumnTypesListStore(types []externglib.Type) {
	var _arg0 *C.GtkListStore // out
	var _arg2 *C.GType
	var _arg1 C.int

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))
	_arg1 = C.int(len(types))
	_arg2 = (*C.GType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(types))
		for i := range types {
			out[i] = (C.GType)(types[i])
		}
	}

	C.gtk_list_store_set_column_types(_arg0, _arg1, _arg2)
}

func (l listStore) SetValueListStore(iter *TreeIter, column int, value externglib.Value) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 C.int           // out
	var _arg3 *C.GValue       // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))
	_arg2 = C.int(column)
	_arg3 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	C.gtk_list_store_set_value(_arg0, _arg1, _arg2, _arg3)
}

func (l listStore) SetValuesvListStore(iter *TreeIter, columns []int, values []externglib.Value) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.int
	var _arg4 C.int
	var _arg3 *C.GValue
	var _arg4 C.int

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))
	_arg4 = C.int(len(columns))
	_arg2 = (*C.int)(unsafe.Pointer(&columns[0]))
	_arg4 = C.int(len(values))
	_arg3 = (*C.GValue)(C.malloc(C.ulong(len(values)) * C.ulong(C.sizeof_GValue)))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice(_arg3, len(values))
		for i := range values {
			{
				var refTmpIn *externglib.Value
				var refTmpOut *C.GValue

				in0 := &values[i]
				refTmpIn = in0

				refTmpOut = (*C.GValue)(unsafe.Pointer(&refTmpIn.GValue))

				out[i] = *refTmpOut
			}
		}
	}

	C.gtk_list_store_set_valuesv(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (s listStore) SwapListStore(a *TreeIter, b *TreeIter) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(a.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(b.Native()))

	C.gtk_list_store_swap(_arg0, _arg1, _arg2)
}

func (b listStore) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}

func (d listStore) DragDataReceived(dest *TreePath, value externglib.Value) bool {
	return WrapTreeDragDest(gextras.InternObject(d)).DragDataReceived(dest, value)
}

func (d listStore) RowDropPossible(destPath *TreePath, value externglib.Value) bool {
	return WrapTreeDragDest(gextras.InternObject(d)).RowDropPossible(destPath, value)
}

func (d listStore) DragDataDelete(path *TreePath) bool {
	return WrapTreeDragSource(gextras.InternObject(d)).DragDataDelete(path)
}

func (d listStore) DragDataGet(path *TreePath) gdk.ContentProvider {
	return WrapTreeDragSource(gextras.InternObject(d)).DragDataGet(path)
}

func (d listStore) RowDraggable(path *TreePath) bool {
	return WrapTreeDragSource(gextras.InternObject(d)).RowDraggable(path)
}

func (s listStore) SortColumnID() (int, SortType, bool) {
	return WrapTreeSortable(gextras.InternObject(s)).SortColumnID()
}

func (s listStore) HasDefaultSortFunc() bool {
	return WrapTreeSortable(gextras.InternObject(s)).HasDefaultSortFunc()
}

func (s listStore) SetSortColumnID(sortColumnId int, order SortType) {
	WrapTreeSortable(gextras.InternObject(s)).SetSortColumnID(sortColumnId, order)
}

func (s listStore) SortColumnChanged() {
	WrapTreeSortable(gextras.InternObject(s)).SortColumnChanged()
}

func (t listStore) NewFilter(root *TreePath) TreeModel {
	return WrapTreeModel(gextras.InternObject(t)).NewFilter(root)
}

func (t listStore) Foreach(fn TreeModelForeachFunc) {
	WrapTreeModel(gextras.InternObject(t)).Foreach(fn)
}

func (t listStore) ColumnType(index_ int) externglib.Type {
	return WrapTreeModel(gextras.InternObject(t)).ColumnType(index_)
}

func (t listStore) Flags() TreeModelFlags {
	return WrapTreeModel(gextras.InternObject(t)).Flags()
}

func (t listStore) Iter(path *TreePath) (TreeIter, bool) {
	return WrapTreeModel(gextras.InternObject(t)).Iter(path)
}

func (t listStore) IterFirst() (TreeIter, bool) {
	return WrapTreeModel(gextras.InternObject(t)).IterFirst()
}

func (t listStore) IterFromString(pathString string) (TreeIter, bool) {
	return WrapTreeModel(gextras.InternObject(t)).IterFromString(pathString)
}

func (t listStore) NColumns() int {
	return WrapTreeModel(gextras.InternObject(t)).NColumns()
}

func (t listStore) Path(iter *TreeIter) *TreePath {
	return WrapTreeModel(gextras.InternObject(t)).Path(iter)
}

func (t listStore) StringFromIter(iter *TreeIter) string {
	return WrapTreeModel(gextras.InternObject(t)).StringFromIter(iter)
}

func (t listStore) Value(iter *TreeIter, column int) externglib.Value {
	return WrapTreeModel(gextras.InternObject(t)).Value(iter, column)
}

func (t listStore) IterChildren(parent *TreeIter) (TreeIter, bool) {
	return WrapTreeModel(gextras.InternObject(t)).IterChildren(parent)
}

func (t listStore) IterHasChild(iter *TreeIter) bool {
	return WrapTreeModel(gextras.InternObject(t)).IterHasChild(iter)
}

func (t listStore) IterNChildren(iter *TreeIter) int {
	return WrapTreeModel(gextras.InternObject(t)).IterNChildren(iter)
}

func (t listStore) IterNext(iter *TreeIter) bool {
	return WrapTreeModel(gextras.InternObject(t)).IterNext(iter)
}

func (t listStore) IterNthChild(parent *TreeIter, n int) (TreeIter, bool) {
	return WrapTreeModel(gextras.InternObject(t)).IterNthChild(parent, n)
}

func (t listStore) IterParent(child *TreeIter) (TreeIter, bool) {
	return WrapTreeModel(gextras.InternObject(t)).IterParent(child)
}

func (t listStore) IterPrevious(iter *TreeIter) bool {
	return WrapTreeModel(gextras.InternObject(t)).IterPrevious(iter)
}

func (t listStore) RefNode(iter *TreeIter) {
	WrapTreeModel(gextras.InternObject(t)).RefNode(iter)
}

func (t listStore) RowChanged(path *TreePath, iter *TreeIter) {
	WrapTreeModel(gextras.InternObject(t)).RowChanged(path, iter)
}

func (t listStore) RowDeleted(path *TreePath) {
	WrapTreeModel(gextras.InternObject(t)).RowDeleted(path)
}

func (t listStore) RowHasChildToggled(path *TreePath, iter *TreeIter) {
	WrapTreeModel(gextras.InternObject(t)).RowHasChildToggled(path, iter)
}

func (t listStore) RowInserted(path *TreePath, iter *TreeIter) {
	WrapTreeModel(gextras.InternObject(t)).RowInserted(path, iter)
}

func (t listStore) RowsReorderedWithLength(path *TreePath, iter *TreeIter, newOrder []int) {
	WrapTreeModel(gextras.InternObject(t)).RowsReorderedWithLength(path, iter, newOrder)
}

func (t listStore) UnrefNode(iter *TreeIter) {
	WrapTreeModel(gextras.InternObject(t)).UnrefNode(iter)
}
