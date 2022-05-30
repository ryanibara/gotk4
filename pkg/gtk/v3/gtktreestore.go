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

// glib.Type values for gtktreestore.go.
var GTypeTreeStore = coreglib.Type(C.gtk_tree_store_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeTreeStore, F: marshalTreeStore},
	})
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

// Clear removes all rows from tree_store.
func (treeStore *TreeStore) Clear() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	*(**TreeStore)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "TreeStore").InvokeMethod("clear", args[:], nil)

	runtime.KeepAlive(treeStore)
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
	var args [3]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _arg2 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg2 = (*C.void)(gextras.StructNative(unsafe.Pointer(descendant)))
	*(**TreeStore)(unsafe.Pointer(&args[1])) = _arg1
	*(**TreeIter)(unsafe.Pointer(&args[2])) = _arg2

	_gret := girepository.MustFind("Gtk", "TreeStore").InvokeMethod("is_ancestor", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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
// The function returns the following values:
//
//    - gint: depth of iter.
//
func (treeStore *TreeStore) IterDepth(iter *TreeIter) int {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	*(**TreeStore)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "TreeStore").InvokeMethod("iter_depth", args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

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
// The function returns the following values:
//
//    - ok: TRUE if the iter is valid, FALSE if the iter is invalid.
//
func (treeStore *TreeStore) IterIsValid(iter *TreeIter) bool {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	*(**TreeStore)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "TreeStore").InvokeMethod("iter_is_valid", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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
//    - position (optional): TreeIter.
//
func (treeStore *TreeStore) MoveAfter(iter, position *TreeIter) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	if position != nil {
		_arg2 = (*C.void)(gextras.StructNative(unsafe.Pointer(position)))
	}
	*(**TreeStore)(unsafe.Pointer(&args[1])) = _arg1
	*(**TreeIter)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "TreeStore").InvokeMethod("move_after", args[:], nil)

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
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	if position != nil {
		_arg2 = (*C.void)(gextras.StructNative(unsafe.Pointer(position)))
	}
	*(**TreeStore)(unsafe.Pointer(&args[1])) = _arg1
	*(**TreeIter)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "TreeStore").InvokeMethod("move_before", args[:], nil)

	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(position)
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
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	*(**TreeStore)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "TreeStore").InvokeMethod("remove", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
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
func (treeStore *TreeStore) SetValue(iter *TreeIter, column int, value *coreglib.Value) {
	var args [4]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 C.gint  // out
	var _arg3 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg2 = C.gint(column)
	_arg3 = (*C.void)(unsafe.Pointer(value.Native()))
	*(**TreeStore)(unsafe.Pointer(&args[1])) = _arg1
	*(**TreeIter)(unsafe.Pointer(&args[2])) = _arg2
	*(*int)(unsafe.Pointer(&args[3])) = _arg3

	girepository.MustFind("Gtk", "TreeStore").InvokeMethod("set_value", args[:], nil)

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
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeStore).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(a)))
	_arg2 = (*C.void)(gextras.StructNative(unsafe.Pointer(b)))
	*(**TreeStore)(unsafe.Pointer(&args[1])) = _arg1
	*(**TreeIter)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "TreeStore").InvokeMethod("swap", args[:], nil)

	runtime.KeepAlive(treeStore)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
