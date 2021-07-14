// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_model_filter_get_type()), F: marshalTreeModelFilterer},
	})
}

// TreeModelFilterModifyFunc: function which calculates display values from raw
// values in the model. It must fill value with the display value for the column
// column in the row indicated by iter.
//
// Since this function is called for each data access, it’s not a particularly
// efficient operation.
type TreeModelFilterModifyFunc func(model *TreeModel, iter *TreeIter, column int, data cgo.Handle) (value externglib.Value)

//export _gotk4_gtk3_TreeModelFilterModifyFunc
func _gotk4_gtk3_TreeModelFilterModifyFunc(arg0 *C.GtkTreeModel, arg1 *C.GtkTreeIter, arg2 *C.GValue, arg3 C.gint, arg4 C.gpointer) {
	v := gbox.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	var model *TreeModel // out
	var iter *TreeIter   // out
	var column int       // out
	var data cgo.Handle  // out

	model = wrapTreeModel(externglib.Take(unsafe.Pointer(arg0)))
	iter = (*TreeIter)(unsafe.Pointer(arg1))
	runtime.SetFinalizer(iter, func(v *TreeIter) {
		C.gtk_tree_iter_free((*C.GtkTreeIter)(unsafe.Pointer(v)))
	})
	column = int(arg3)
	data = (cgo.Handle)(unsafe.Pointer(arg4))

	fn := v.(TreeModelFilterModifyFunc)
	value := fn(model, iter, column, data)

	*arg2 = *(*C.GValue)(unsafe.Pointer(&(&value).GValue))
}

// TreeModelFilterVisibleFunc: function which decides whether the row indicated
// by iter is visible.
type TreeModelFilterVisibleFunc func(model *TreeModel, iter *TreeIter, data cgo.Handle) (ok bool)

//export _gotk4_gtk3_TreeModelFilterVisibleFunc
func _gotk4_gtk3_TreeModelFilterVisibleFunc(arg0 *C.GtkTreeModel, arg1 *C.GtkTreeIter, arg2 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var model *TreeModel // out
	var iter *TreeIter   // out
	var data cgo.Handle  // out

	model = wrapTreeModel(externglib.Take(unsafe.Pointer(arg0)))
	iter = (*TreeIter)(unsafe.Pointer(arg1))
	runtime.SetFinalizer(iter, func(v *TreeIter) {
		C.gtk_tree_iter_free((*C.GtkTreeIter)(unsafe.Pointer(v)))
	})
	data = (cgo.Handle)(unsafe.Pointer(arg2))

	fn := v.(TreeModelFilterVisibleFunc)
	ok := fn(model, iter, data)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// TreeModelFilterOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TreeModelFilterOverrider interface {
	Modify(childModel TreeModeler, iter *TreeIter, value *externglib.Value, column int)
	Visible(childModel TreeModeler, iter *TreeIter) bool
}

// TreeModelFilterer describes TreeModelFilter's methods.
type TreeModelFilterer interface {
	// ClearCache: this function should almost never be called.
	ClearCache()
	// ConvertChildIterToIter sets filter_iter to point to the row in filter
	// that corresponds to the row pointed at by child_iter.
	ConvertChildIterToIter(childIter *TreeIter) (TreeIter, bool)
	// ConvertChildPathToPath converts child_path to a path relative to filter.
	ConvertChildPathToPath(childPath *TreePath) *TreePath
	// ConvertIterToChildIter sets child_iter to point to the row pointed to by
	// filter_iter.
	ConvertIterToChildIter(filterIter *TreeIter) TreeIter
	// ConvertPathToChildPath converts filter_path to a path on the child model
	// of filter.
	ConvertPathToChildPath(filterPath *TreePath) *TreePath
	// Model returns a pointer to the child model of filter.
	Model() *TreeModel
	// Refilter emits ::row_changed for each row in the child model, which
	// causes the filter to re-evaluate whether a row is visible or not.
	Refilter()
	// SetVisibleColumn sets column of the child_model to be the column where
	// filter should look for visibility information.
	SetVisibleColumn(column int)
}

// TreeModelFilter is a tree model which wraps another tree model, and can do
// the following things:
//
// - Filter specific rows, based on data from a “visible column”, a column
// storing booleans indicating whether the row should be filtered or not, or
// based on the return value of a “visible function”, which gets a model, iter
// and user_data and returns a boolean indicating whether the row should be
// filtered or not.
//
// - Modify the “appearance” of the model, using a modify function. This is
// extremely powerful and allows for just changing some values and also for
// creating a completely different model based on the given child model.
//
// - Set a different root node, also known as a “virtual root”. You can pass in
// a TreePath indicating the root node for the filter at construction time.
//
// The basic API is similar to TreeModelSort. For an example on its usage, see
// the section on TreeModelSort.
//
// When using TreeModelFilter, it is important to realize that TreeModelFilter
// maintains an internal cache of all nodes which are visible in its clients.
// The cache is likely to be a subtree of the tree exposed by the child model.
// TreeModelFilter will not cache the entire child model when unnecessary to not
// compromise the caching mechanism that is exposed by the reference counting
// scheme. If the child model implements reference counting, unnecessary signals
// may not be emitted because of reference counting rule 3, see the TreeModel
// documentation. (Note that e.g. TreeStore does not implement reference
// counting and will always emit all signals, even when the receiving node is
// not visible).
//
// Because of this, limitations for possible visible functions do apply. In
// general, visible functions should only use data or properties from the node
// for which the visibility state must be determined, its siblings or its
// parents. Usually, having a dependency on the state of any child node is not
// possible, unless references are taken on these explicitly. When no such
// reference exists, no signals may be received for these child nodes (see
// reference couting rule number 3 in the TreeModel section).
//
// Determining the visibility state of a given node based on the state of its
// child nodes is a frequently occurring use case. Therefore, TreeModelFilter
// explicitly supports this. For example, when a node does not have any
// children, you might not want the node to be visible. As soon as the first row
// is added to the node’s child level (or the last row removed), the node’s
// visibility should be updated.
//
// This introduces a dependency from the node on its child nodes. In order to
// accommodate this, TreeModelFilter must make sure the necessary signals are
// received from the child model. This is achieved by building, for all nodes
// which are exposed as visible nodes to TreeModelFilter's clients, the child
// level (if any) and take a reference on the first node in this level.
// Furthermore, for every row-inserted, row-changed or row-deleted signal (also
// these which were not handled because the node was not cached),
// TreeModelFilter will check if the visibility state of any parent node has
// changed.
//
// Beware, however, that this explicit support is limited to these two cases.
// For example, if you want a node to be visible only if two nodes in a child’s
// child level (2 levels deeper) are visible, you are on your own. In this case,
// either rely on TreeStore to emit all signals because it does not implement
// reference counting, or for models that do implement reference counting,
// obtain references on these child levels yourself.
type TreeModelFilter struct {
	*externglib.Object

	TreeDragSource
	TreeModel
}

var (
	_ TreeModelFilterer = (*TreeModelFilter)(nil)
	_ gextras.Nativer   = (*TreeModelFilter)(nil)
)

func wrapTreeModelFilter(obj *externglib.Object) *TreeModelFilter {
	return &TreeModelFilter{
		Object: obj,
		TreeDragSource: TreeDragSource{
			Object: obj,
		},
		TreeModel: TreeModel{
			Object: obj,
		},
	}
}

func marshalTreeModelFilterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTreeModelFilter(obj), nil
}

// ClearCache: this function should almost never be called. It clears the filter
// of any cached iterators that haven’t been reffed with
// gtk_tree_model_ref_node(). This might be useful if the child model being
// filtered is static (and doesn’t change often) and there has been a lot of
// unreffed access to nodes. As a side effect of this function, all unreffed
// iters will be invalid.
func (filter *TreeModelFilter) ClearCache() {
	var _arg0 *C.GtkTreeModelFilter // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_tree_model_filter_clear_cache(_arg0)
}

// ConvertChildIterToIter sets filter_iter to point to the row in filter that
// corresponds to the row pointed at by child_iter. If filter_iter was not set,
// FALSE is returned.
func (filter *TreeModelFilter) ConvertChildIterToIter(childIter *TreeIter) (TreeIter, bool) {
	var _arg0 *C.GtkTreeModelFilter // out
	var _filterIter TreeIter
	var _arg2 *C.GtkTreeIter // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(childIter))

	_cret = C.gtk_tree_model_filter_convert_child_iter_to_iter(_arg0, (*C.GtkTreeIter)(unsafe.Pointer(&_filterIter)), _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _filterIter, _ok
}

// ConvertChildPathToPath converts child_path to a path relative to filter. That
// is, child_path points to a path in the child model. The rerturned path will
// point to the same row in the filtered model. If child_path isn’t a valid path
// on the child model or points to a row which is not visible in filter, then
// NULL is returned.
func (filter *TreeModelFilter) ConvertChildPathToPath(childPath *TreePath) *TreePath {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg1 *C.GtkTreePath        // out
	var _cret *C.GtkTreePath        // in

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(childPath))

	_cret = C.gtk_tree_model_filter_convert_child_path_to_path(_arg0, _arg1)

	var _treePath *TreePath // out

	_treePath = (*TreePath)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_treePath, func(v *TreePath) {
		C.gtk_tree_path_free((*C.GtkTreePath)(unsafe.Pointer(v)))
	})

	return _treePath
}

// ConvertIterToChildIter sets child_iter to point to the row pointed to by
// filter_iter.
func (filter *TreeModelFilter) ConvertIterToChildIter(filterIter *TreeIter) TreeIter {
	var _arg0 *C.GtkTreeModelFilter // out
	var _childIter TreeIter
	var _arg2 *C.GtkTreeIter // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(filterIter))

	C.gtk_tree_model_filter_convert_iter_to_child_iter(_arg0, (*C.GtkTreeIter)(unsafe.Pointer(&_childIter)), _arg2)

	return _childIter
}

// ConvertPathToChildPath converts filter_path to a path on the child model of
// filter. That is, filter_path points to a location in filter. The returned
// path will point to the same location in the model not being filtered. If
// filter_path does not point to a location in the child model, NULL is
// returned.
func (filter *TreeModelFilter) ConvertPathToChildPath(filterPath *TreePath) *TreePath {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg1 *C.GtkTreePath        // out
	var _cret *C.GtkTreePath        // in

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(filterPath))

	_cret = C.gtk_tree_model_filter_convert_path_to_child_path(_arg0, _arg1)

	var _treePath *TreePath // out

	_treePath = (*TreePath)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_treePath, func(v *TreePath) {
		C.gtk_tree_path_free((*C.GtkTreePath)(unsafe.Pointer(v)))
	})

	return _treePath
}

// Model returns a pointer to the child model of filter.
func (filter *TreeModelFilter) Model() *TreeModel {
	var _arg0 *C.GtkTreeModelFilter // out
	var _cret *C.GtkTreeModel       // in

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))

	_cret = C.gtk_tree_model_filter_get_model(_arg0)

	var _treeModel *TreeModel // out

	_treeModel = wrapTreeModel(externglib.Take(unsafe.Pointer(_cret)))

	return _treeModel
}

// Refilter emits ::row_changed for each row in the child model, which causes
// the filter to re-evaluate whether a row is visible or not.
func (filter *TreeModelFilter) Refilter() {
	var _arg0 *C.GtkTreeModelFilter // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_tree_model_filter_refilter(_arg0)
}

// SetVisibleColumn sets column of the child_model to be the column where filter
// should look for visibility information. columns should be a column of type
// G_TYPE_BOOLEAN, where TRUE means that a row is visible, and FALSE if not.
//
// Note that gtk_tree_model_filter_set_visible_func() or
// gtk_tree_model_filter_set_visible_column() can only be called once for a
// given filter model.
func (filter *TreeModelFilter) SetVisibleColumn(column int) {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg1 C.gint                // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = C.gint(column)

	C.gtk_tree_model_filter_set_visible_column(_arg0, _arg1)
}
