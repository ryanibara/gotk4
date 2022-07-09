// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeTreeStore returns the GType for the type TreeStore.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTreeStore() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "TreeStore").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalTreeStore)
	return gtype
}

// TreeStoreOverrider contains methods that are overridable.
type TreeStoreOverrider interface {
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
//    </object>.
type TreeStore struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Buildable
	TreeDragDest
	TreeDragSource
	TreeSortable
}

var (
	_ coreglib.Objector = (*TreeStore)(nil)
)

func classInitTreeStorer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTreeStore(obj *coreglib.Object) *TreeStore {
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

func marshalTreeStore(p uintptr) (interface{}, error) {
	return wrapTreeStore(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Append appends a new row to tree_store. If parent is non-NULL, then it will
// append the new row after the last child of parent, otherwise it will append a
// row to the top level. iter will be changed to point to this new row. The row
// will be empty after this function is called. To fill in values, you need to
// call gtk_tree_store_set() or gtk_tree_store_set_value().
//
// The function takes the following parameters:
//
//    - parent (optional): valid TreeIter, or NULL.
//
// The function returns the following values:
//
//    - iter: unset TreeIter to set to the appended row.
//
func (treeStore *TreeStore) Append(parent *TreeIter) *TreeIter {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	if parent != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(parent)))
	}

	girepository.MustFind("Gtk", "TreeStore").InvokeMethod("append", _args[:], _outs[:])

	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(parent)

	var _iter *TreeIter // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))

	return _iter
}

// Clear removes all rows from tree_store.
func (treeStore *TreeStore) Clear() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))

	girepository.MustFind("Gtk", "TreeStore").InvokeMethod("clear", _args[:], nil)

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
//    - parent (optional): valid TreeIter, or NULL.
//    - position to insert the new row, or -1 for last.
//
// The function returns the following values:
//
//    - iter: unset TreeIter to set to the new row.
//
func (treeStore *TreeStore) Insert(parent *TreeIter, position int32) *TreeIter {
	var _args [3]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	if parent != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(parent)))
	}
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(position)

	girepository.MustFind("Gtk", "TreeStore").InvokeMethod("insert", _args[:], _outs[:])

	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(position)

	var _iter *TreeIter // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))

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
//    - parent (optional): valid TreeIter, or NULL.
//    - sibling (optional): valid TreeIter, or NULL.
//
// The function returns the following values:
//
//    - iter: unset TreeIter to set to the new row.
//
func (treeStore *TreeStore) InsertAfter(parent, sibling *TreeIter) *TreeIter {
	var _args [3]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	if parent != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(parent)))
	}
	if sibling != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(sibling)))
	}

	girepository.MustFind("Gtk", "TreeStore").InvokeMethod("insert_after", _args[:], _outs[:])

	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(sibling)

	var _iter *TreeIter // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))

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
//    - parent (optional): valid TreeIter, or NULL.
//    - sibling (optional): valid TreeIter, or NULL.
//
// The function returns the following values:
//
//    - iter: unset TreeIter to set to the new row.
//
func (treeStore *TreeStore) InsertBefore(parent, sibling *TreeIter) *TreeIter {
	var _args [3]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	if parent != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(parent)))
	}
	if sibling != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(sibling)))
	}

	girepository.MustFind("Gtk", "TreeStore").InvokeMethod("insert_before", _args[:], _outs[:])

	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(sibling)

	var _iter *TreeIter // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))

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
// The function returns the following values:
//
//    - ok: TRUE, if iter is an ancestor of descendant.
//
func (treeStore *TreeStore) IsAncestor(iter, descendant *TreeIter) bool {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(descendant)))

	_gret := girepository.MustFind("Gtk", "TreeStore").InvokeMethod("is_ancestor", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(descendant)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
// The function returns the following values:
//
//    - gint: depth of iter.
//
func (treeStore *TreeStore) IterDepth(iter *TreeIter) int32 {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))

	_gret := girepository.MustFind("Gtk", "TreeStore").InvokeMethod("iter_depth", _args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(iter)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

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
// The function returns the following values:
//
//    - ok: TRUE if the iter is valid, FALSE if the iter is invalid.
//
func (treeStore *TreeStore) IterIsValid(iter *TreeIter) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))

	_gret := girepository.MustFind("Gtk", "TreeStore").InvokeMethod("iter_is_valid", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
//    - position (optional): TreeIter.
//
func (treeStore *TreeStore) MoveAfter(iter, position *TreeIter) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	if position != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(position)))
	}

	girepository.MustFind("Gtk", "TreeStore").InvokeMethod("move_after", _args[:], nil)

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
//    - position (optional) or NULL.
//
func (treeStore *TreeStore) MoveBefore(iter, position *TreeIter) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	if position != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(position)))
	}

	girepository.MustFind("Gtk", "TreeStore").InvokeMethod("move_before", _args[:], nil)

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
//    - parent (optional): valid TreeIter, or NULL.
//
// The function returns the following values:
//
//    - iter: unset TreeIter to set to the prepended row.
//
func (treeStore *TreeStore) Prepend(parent *TreeIter) *TreeIter {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	if parent != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(parent)))
	}

	girepository.MustFind("Gtk", "TreeStore").InvokeMethod("prepend", _args[:], _outs[:])

	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(parent)

	var _iter *TreeIter // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))

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
// The function returns the following values:
//
//    - ok: TRUE if iter is still valid, FALSE if not.
//
func (treeStore *TreeStore) Remove(iter *TreeIter) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))

	_gret := girepository.MustFind("Gtk", "TreeStore").InvokeMethod("remove", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
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
func (treeStore *TreeStore) SetValue(iter *TreeIter, column int32, value *coreglib.Value) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(column)
	*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(value.Native()))

	girepository.MustFind("Gtk", "TreeStore").InvokeMethod("set_value", _args[:], nil)

	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(column)
	runtime.KeepAlive(value)
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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(a)))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(b)))

	girepository.MustFind("Gtk", "TreeStore").InvokeMethod("swap", _args[:], nil)

	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
