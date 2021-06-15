// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_drag_dest_get_type()), F: marshalTreeDragDest},
		{T: externglib.Type(C.gtk_tree_drag_source_get_type()), F: marshalTreeDragSource},
	})
}

// TreeCreateRowDragContent creates a content provider for dragging @path from
// @tree_model.
func TreeCreateRowDragContent(treeModel TreeModel, path *TreePath) gdk.ContentProvider {
	var _arg1 *C.GtkTreeModel       // out
	var _arg2 *C.GtkTreePath        // out
	var _cret *C.GdkContentProvider // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(treeModel.Native()))
	_arg2 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	_cret = C.gtk_tree_create_row_drag_content(_arg1, _arg2)

	var _contentProvider gdk.ContentProvider // out

	_contentProvider = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(gdk.ContentProvider)

	return _contentProvider
}

// TreeGetRowDragData obtains a @tree_model and @path from value of target type
// GTK_TYPE_TREE_ROW_DATA.
//
// The returned path must be freed with gtk_tree_path_free().
func TreeGetRowDragData(value **externglib.Value) (TreeModel, *TreePath, bool) {
	var _arg1 *C.GValue       // out
	var _arg2 *C.GtkTreeModel // in
	var _path *TreePath
	var _cret C.gboolean // in

	_arg1 = (*C.GValue)(value.GValue)

	_cret = C.gtk_tree_get_row_drag_data(_arg1, &_arg2, (**C.GtkTreePath)(unsafe.Pointer(&_path)))

	var _treeModel TreeModel // out

	var _ok bool // out

	_treeModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_arg2.Native()))).(TreeModel)

	if _cret != 0 {
		_ok = true
	}

	return _treeModel, _path, _ok
}

// TreeDragDestOverrider contains methods that are overridable. This
// interface is a subset of the interface TreeDragDest.
type TreeDragDestOverrider interface {
	// DragDataReceived asks the TreeDragDest to insert a row before the path
	// @dest, deriving the contents of the row from @value. If @dest is outside
	// the tree so that inserting before it is impossible, false will be
	// returned. Also, false may be returned if the new row is not created for
	// some model-specific reason. Should robustly handle a @dest no longer
	// found in the model!
	DragDataReceived(dest *TreePath, value **externglib.Value) bool
	// RowDropPossible determines whether a drop is possible before the given
	// @dest_path, at the same depth as @dest_path. i.e., can we drop the data
	// in @value at that location. @dest_path does not have to exist; the return
	// value will almost certainly be false if the parent of @dest_path doesn’t
	// exist, though.
	RowDropPossible(destPath *TreePath, value **externglib.Value) bool
}

// TreeDragDest: interface for Drag-and-Drop destinations in `GtkTreeView`.
type TreeDragDest interface {
	gextras.Objector
	TreeDragDestOverrider
}

// treeDragDest implements the TreeDragDest interface.
type treeDragDest struct {
	gextras.Objector
}

var _ TreeDragDest = (*treeDragDest)(nil)

// WrapTreeDragDest wraps a GObject to a type that implements interface
// TreeDragDest. It is primarily used internally.
func WrapTreeDragDest(obj *externglib.Object) TreeDragDest {
	return TreeDragDest{
		Objector: obj,
	}
}

func marshalTreeDragDest(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeDragDest(obj), nil
}

// DragDataReceived asks the TreeDragDest to insert a row before the path
// @dest, deriving the contents of the row from @value. If @dest is outside
// the tree so that inserting before it is impossible, false will be
// returned. Also, false may be returned if the new row is not created for
// some model-specific reason. Should robustly handle a @dest no longer
// found in the model!
func (d treeDragDest) DragDataReceived(dest *TreePath, value **externglib.Value) bool {
	var _arg0 *C.GtkTreeDragDest // out
	var _arg1 *C.GtkTreePath     // out
	var _arg2 *C.GValue          // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeDragDest)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(dest.Native()))
	_arg2 = (*C.GValue)(value.GValue)

	_cret = C.gtk_tree_drag_dest_drag_data_received(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowDropPossible determines whether a drop is possible before the given
// @dest_path, at the same depth as @dest_path. i.e., can we drop the data
// in @value at that location. @dest_path does not have to exist; the return
// value will almost certainly be false if the parent of @dest_path doesn’t
// exist, though.
func (d treeDragDest) RowDropPossible(destPath *TreePath, value **externglib.Value) bool {
	var _arg0 *C.GtkTreeDragDest // out
	var _arg1 *C.GtkTreePath     // out
	var _arg2 *C.GValue          // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeDragDest)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(destPath.Native()))
	_arg2 = (*C.GValue)(value.GValue)

	_cret = C.gtk_tree_drag_dest_row_drop_possible(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TreeDragSourceOverrider contains methods that are overridable. This
// interface is a subset of the interface TreeDragSource.
type TreeDragSourceOverrider interface {
	// DragDataDelete asks the TreeDragSource to delete the row at @path,
	// because it was moved somewhere else via drag-and-drop. Returns false if
	// the deletion fails because @path no longer exists, or for some
	// model-specific reason. Should robustly handle a @path no longer found in
	// the model!
	DragDataDelete(path *TreePath) bool
	// DragDataGet asks the TreeDragSource to return a ContentProvider
	// representing the row at @path. Should robustly handle a @path no longer
	// found in the model!
	DragDataGet(path *TreePath) gdk.ContentProvider
	// RowDraggable asks the TreeDragSource whether a particular row can be used
	// as the source of a DND operation. If the source doesn’t implement this
	// interface, the row is assumed draggable.
	RowDraggable(path *TreePath) bool
}

// TreeDragSource: interface for Drag-and-Drop destinations in `GtkTreeView`.
type TreeDragSource interface {
	gextras.Objector
	TreeDragSourceOverrider
}

// treeDragSource implements the TreeDragSource interface.
type treeDragSource struct {
	gextras.Objector
}

var _ TreeDragSource = (*treeDragSource)(nil)

// WrapTreeDragSource wraps a GObject to a type that implements interface
// TreeDragSource. It is primarily used internally.
func WrapTreeDragSource(obj *externglib.Object) TreeDragSource {
	return TreeDragSource{
		Objector: obj,
	}
}

func marshalTreeDragSource(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeDragSource(obj), nil
}

// DragDataDelete asks the TreeDragSource to delete the row at @path,
// because it was moved somewhere else via drag-and-drop. Returns false if
// the deletion fails because @path no longer exists, or for some
// model-specific reason. Should robustly handle a @path no longer found in
// the model!
func (d treeDragSource) DragDataDelete(path *TreePath) bool {
	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	_cret = C.gtk_tree_drag_source_drag_data_delete(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DragDataGet asks the TreeDragSource to return a ContentProvider
// representing the row at @path. Should robustly handle a @path no longer
// found in the model!
func (d treeDragSource) DragDataGet(path *TreePath) gdk.ContentProvider {
	var _arg0 *C.GtkTreeDragSource  // out
	var _arg1 *C.GtkTreePath        // out
	var _cret *C.GdkContentProvider // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	_cret = C.gtk_tree_drag_source_drag_data_get(_arg0, _arg1)

	var _contentProvider gdk.ContentProvider // out

	_contentProvider = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(gdk.ContentProvider)

	return _contentProvider
}

// RowDraggable asks the TreeDragSource whether a particular row can be used
// as the source of a DND operation. If the source doesn’t implement this
// interface, the row is assumed draggable.
func (d treeDragSource) RowDraggable(path *TreePath) bool {
	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	_cret = C.gtk_tree_drag_source_row_draggable(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
