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

// glib.Type values for gtktreemodelsort.go.
var GTypeTreeModelSort = coreglib.Type(C.gtk_tree_model_sort_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeTreeModelSort, F: marshalTreeModelSort},
	})
}

// TreeModelSortOverrider contains methods that are overridable.
type TreeModelSortOverrider interface {
}

// TreeModelSort is a model which implements the TreeSortable interface. It does
// not hold any data itself, but rather is created with a child model and
// proxies its data. It has identical column types to this child model, and the
// changes in the child are propagated. The primary purpose of this model is to
// provide a way to sort a different model without modifying it. Note that the
// sort function used by TreeModelSort is not guaranteed to be stable.
//
// The use of this is best demonstrated through an example. In the following
// sample code we create two TreeView widgets each with a view of the same data.
// As the model is wrapped here by a TreeModelSort, the two TreeViews can each
// sort their view of the data without affecting the other. By contrast, if we
// simply put the same model in each widget, then sorting the first would sort
// the second.
//
// Using a TreeModelSort
//
//    void
//    selection_changed (GtkTreeSelection *selection, gpointer data)
//    {
//      GtkTreeModel *sort_model = NULL;
//      GtkTreeModel *child_model;
//      GtkTreeIter sort_iter;
//      GtkTreeIter child_iter;
//      char *some_data = NULL;
//      char *modified_data;
//
//      // Get the current selected row and the model.
//      if (! gtk_tree_selection_get_selected (selection,
//                                             &sort_model,
//                                             &sort_iter))
//        return;
//
//      // Look up the current value on the selected row and get
//      // a new value to change it to.
//      gtk_tree_model_get (GTK_TREE_MODEL (sort_model), &sort_iter,
//                          COLUMN_1, &some_data,
//                          -1);
//
//      modified_data = change_the_data (some_data);
//      g_free (some_data);
//
//      // Get an iterator on the child model, instead of the sort model.
//      gtk_tree_model_sort_convert_iter_to_child_iter (GTK_TREE_MODEL_SORT (sort_model),
//                                                      &child_iter,
//                                                      &sort_iter);
//
//      // Get the child model and change the value of the row. In this
//      // example, the child model is a GtkListStore. It could be any other
//      // type of model, though.
//      child_model = gtk_tree_model_sort_get_model (GTK_TREE_MODEL_SORT (sort_model));
//      gtk_list_store_set (GTK_LIST_STORE (child_model), &child_iter,
//                          COLUMN_1, &modified_data,
//                          -1);
//      g_free (modified_data);
//    }.
type TreeModelSort struct {
	_ [0]func() // equal guard
	*coreglib.Object

	TreeDragSource
	TreeSortable
}

var (
	_ coreglib.Objector = (*TreeModelSort)(nil)
)

func classInitTreeModelSorter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTreeModelSort(obj *coreglib.Object) *TreeModelSort {
	return &TreeModelSort{
		Object: obj,
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

func marshalTreeModelSort(p uintptr) (interface{}, error) {
	return wrapTreeModelSort(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTreeModelSortWithModel creates a new TreeModelSort, with child_model as
// the child model.
//
// The function takes the following parameters:
//
//    - childModel: TreeModel.
//
// The function returns the following values:
//
//    - treeModelSort: new TreeModelSort.
//
func NewTreeModelSortWithModel(childModel TreeModeller) *TreeModelSort {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(childModel).Native()))

	_gret := girepository.MustFind("Gtk", "TreeModelSort").InvokeMethod("new_TreeModelSort_with_model", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(childModel)

	var _treeModelSort *TreeModelSort // out

	_treeModelSort = wrapTreeModelSort(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _treeModelSort
}

// ClearCache: this function should almost never be called. It clears the
// tree_model_sort of any cached iterators that haven’t been reffed with
// gtk_tree_model_ref_node(). This might be useful if the child model being
// sorted is static (and doesn’t change often) and there has been a lot of
// unreffed access to nodes. As a side effect of this function, all unreffed
// iters will be invalid.
func (treeModelSort *TreeModelSort) ClearCache() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeModelSort).Native()))

	girepository.MustFind("Gtk", "TreeModelSort").InvokeMethod("clear_cache", _args[:], nil)

	runtime.KeepAlive(treeModelSort)
}

// ConvertChildIterToIter sets sort_iter to point to the row in tree_model_sort
// that corresponds to the row pointed at by child_iter. If sort_iter was not
// set, FALSE is returned. Note: a boolean is only returned since 2.14.
//
// The function takes the following parameters:
//
//    - childIter: valid TreeIter pointing to a row on the child model.
//
// The function returns the following values:
//
//    - sortIter: uninitialized TreeIter.
//    - ok: TRUE, if sort_iter was set, i.e. if sort_iter is a valid iterator
//      pointer to a visible row in the child model.
//
func (treeModelSort *TreeModelSort) ConvertChildIterToIter(childIter *TreeIter) (*TreeIter, bool) {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeModelSort).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(childIter)))

	_gret := girepository.MustFind("Gtk", "TreeModelSort").InvokeMethod("convert_child_iter_to_iter", _args[:], _outs[:])
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(treeModelSort)
	runtime.KeepAlive(childIter)

	var _sortIter *TreeIter // out
	var _ok bool            // out

	_sortIter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _sortIter, _ok
}

// ConvertChildPathToPath converts child_path to a path relative to
// tree_model_sort. That is, child_path points to a path in the child model. The
// returned path will point to the same row in the sorted model. If child_path
// isn’t a valid path on the child model, then NULL is returned.
//
// The function takes the following parameters:
//
//    - childPath to convert.
//
// The function returns the following values:
//
//    - treePath (optional): newly allocated TreePath, or NULL.
//
func (treeModelSort *TreeModelSort) ConvertChildPathToPath(childPath *TreePath) *TreePath {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeModelSort).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(childPath)))

	_gret := girepository.MustFind("Gtk", "TreeModelSort").InvokeMethod("convert_child_path_to_path", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(treeModelSort)
	runtime.KeepAlive(childPath)

	var _treePath *TreePath // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_treePath = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_treePath)),
			func(intern *struct{ C unsafe.Pointer }) {
				{
					args := [1]girepository.Argument{(*C.void)(intern.C)}
					girepository.MustFind("Gtk", "TreePath").InvokeMethod("free", args[:], nil)
				}
			},
		)
	}

	return _treePath
}

// ConvertIterToChildIter sets child_iter to point to the row pointed to by
// sorted_iter.
//
// The function takes the following parameters:
//
//    - sortedIter: valid TreeIter pointing to a row on tree_model_sort.
//
// The function returns the following values:
//
//    - childIter: uninitialized TreeIter.
//
func (treeModelSort *TreeModelSort) ConvertIterToChildIter(sortedIter *TreeIter) *TreeIter {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeModelSort).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(sortedIter)))

	girepository.MustFind("Gtk", "TreeModelSort").InvokeMethod("convert_iter_to_child_iter", _args[:], _outs[:])

	runtime.KeepAlive(treeModelSort)
	runtime.KeepAlive(sortedIter)

	var _childIter *TreeIter // out

	_childIter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))

	return _childIter
}

// ConvertPathToChildPath converts sorted_path to a path on the child model of
// tree_model_sort. That is, sorted_path points to a location in
// tree_model_sort. The returned path will point to the same location in the
// model not being sorted. If sorted_path does not point to a location in the
// child model, NULL is returned.
//
// The function takes the following parameters:
//
//    - sortedPath to convert.
//
// The function returns the following values:
//
//    - treePath (optional): newly allocated TreePath, or NULL.
//
func (treeModelSort *TreeModelSort) ConvertPathToChildPath(sortedPath *TreePath) *TreePath {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeModelSort).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(sortedPath)))

	_gret := girepository.MustFind("Gtk", "TreeModelSort").InvokeMethod("convert_path_to_child_path", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(treeModelSort)
	runtime.KeepAlive(sortedPath)

	var _treePath *TreePath // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_treePath = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_treePath)),
			func(intern *struct{ C unsafe.Pointer }) {
				{
					args := [1]girepository.Argument{(*C.void)(intern.C)}
					girepository.MustFind("Gtk", "TreePath").InvokeMethod("free", args[:], nil)
				}
			},
		)
	}

	return _treePath
}

// Model returns the model the TreeModelSort is sorting.
//
// The function returns the following values:
//
//    - treeModel: "child model" being sorted.
//
func (treeModel *TreeModelSort) Model() *TreeModel {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeModel).Native()))

	_gret := girepository.MustFind("Gtk", "TreeModelSort").InvokeMethod("get_model", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(treeModel)

	var _treeModel *TreeModel // out

	_treeModel = wrapTreeModel(coreglib.Take(unsafe.Pointer(_cret)))

	return _treeModel
}

// IterIsValid: > This function is slow. Only use it for debugging and/or
// testing > purposes.
//
// Checks if the given iter is a valid iter for this TreeModelSort.
//
// The function takes the following parameters:
//
//    - iter: TreeIter.
//
// The function returns the following values:
//
//    - ok: TRUE if the iter is valid, FALSE if the iter is invalid.
//
func (treeModelSort *TreeModelSort) IterIsValid(iter *TreeIter) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeModelSort).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))

	_gret := girepository.MustFind("Gtk", "TreeModelSort").InvokeMethod("iter_is_valid", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(treeModelSort)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ResetDefaultSortFunc: this resets the default sort function to be in the
// “unsorted” state. That is, it is in the same order as the child model. It
// will re-sort the model to be in the same order as the child model only if the
// TreeModelSort is in “unsorted” state.
func (treeModelSort *TreeModelSort) ResetDefaultSortFunc() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeModelSort).Native()))

	girepository.MustFind("Gtk", "TreeModelSort").InvokeMethod("reset_default_sort_func", _args[:], nil)

	runtime.KeepAlive(treeModelSort)
}
