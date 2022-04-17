// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern GdkContentProvider* _gotk4_gtk4_TreeDragSourceIface_drag_data_get(GtkTreeDragSource*, GtkTreePath*);
// extern gboolean _gotk4_gtk4_TreeDragDestIface_drag_data_received(GtkTreeDragDest*, GtkTreePath*, GValue*);
// extern gboolean _gotk4_gtk4_TreeDragDestIface_row_drop_possible(GtkTreeDragDest*, GtkTreePath*, GValue*);
// extern gboolean _gotk4_gtk4_TreeDragSourceIface_drag_data_delete(GtkTreeDragSource*, GtkTreePath*);
// extern gboolean _gotk4_gtk4_TreeDragSourceIface_row_draggable(GtkTreeDragSource*, GtkTreePath*);
import "C"

// glib.Type values for gtktreednd.go.
var (
	GTypeTreeDragDest   = externglib.Type(C.gtk_tree_drag_dest_get_type())
	GTypeTreeDragSource = externglib.Type(C.gtk_tree_drag_source_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeTreeDragDest, F: marshalTreeDragDest},
		{T: GTypeTreeDragSource, F: marshalTreeDragSource},
	})
}

// TreeCreateRowDragContent creates a content provider for dragging path from
// tree_model.
//
// The function takes the following parameters:
//
//    - treeModel: TreeModel.
//    - path: row in tree_model.
//
// The function returns the following values:
//
//    - contentProvider: new ContentProvider.
//
func TreeCreateRowDragContent(treeModel TreeModelOverrider, path *TreePath) *gdk.ContentProvider {
	var _arg1 *C.GtkTreeModel       // out
	var _arg2 *C.GtkTreePath        // out
	var _cret *C.GdkContentProvider // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(externglib.InternObject(treeModel).Native()))
	_arg2 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_tree_create_row_drag_content(_arg1, _arg2)
	runtime.KeepAlive(treeModel)
	runtime.KeepAlive(path)

	var _contentProvider *gdk.ContentProvider // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_contentProvider = &gdk.ContentProvider{
			Object: obj,
		}
	}

	return _contentProvider
}

// TreeGetRowDragData obtains a tree_model and path from value of target type
// GTK_TYPE_TREE_ROW_DATA.
//
// The returned path must be freed with gtk_tree_path_free().
//
// The function takes the following parameters:
//
//    - value: #GValue.
//
// The function returns the following values:
//
//    - treeModel (optional): TreeModel.
//    - path (optional): row in tree_model.
//    - ok: TRUE if selection_data had target type GTK_TYPE_TREE_ROW_DATA is
//      otherwise valid.
//
func TreeGetRowDragData(value *externglib.Value) (TreeModelOverrider, *TreePath, bool) {
	var _arg1 *C.GValue       // out
	var _arg2 *C.GtkTreeModel // in
	var _arg3 *C.GtkTreePath  // in
	var _cret C.gboolean      // in

	_arg1 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C.gtk_tree_get_row_drag_data(_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(value)

	var _treeModel TreeModelOverrider // out
	var _path *TreePath               // out
	var _ok bool                      // out

	if _arg2 != nil {
		{
			objptr := unsafe.Pointer(_arg2)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(TreeModelOverrider)
				return ok
			})
			rv, ok := casted.(TreeModelOverrider)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.TreeModeller")
			}
			_treeModel = rv
		}
	}
	if _arg3 != nil {
		_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(_arg3)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_path)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gtk_tree_path_free((*C.GtkTreePath)(intern.C))
			},
		)
	}
	if _cret != 0 {
		_ok = true
	}

	return _treeModel, _path, _ok
}

// TreeDragDestOverrider contains methods that are overridable.
type TreeDragDestOverrider interface {
	externglib.Objector
	// DragDataReceived asks the TreeDragDest to insert a row before the path
	// dest, deriving the contents of the row from value. If dest is outside the
	// tree so that inserting before it is impossible, FALSE will be returned.
	// Also, FALSE may be returned if the new row is not created for some
	// model-specific reason. Should robustly handle a dest no longer found in
	// the model!.
	//
	// The function takes the following parameters:
	//
	//    - dest: row to drop in front of.
	//    - value: data to drop.
	//
	// The function returns the following values:
	//
	//    - ok: whether a new row was created before position dest.
	//
	DragDataReceived(dest *TreePath, value *externglib.Value) bool
	// RowDropPossible determines whether a drop is possible before the given
	// dest_path, at the same depth as dest_path. i.e., can we drop the data in
	// value at that location. dest_path does not have to exist; the return
	// value will almost certainly be FALSE if the parent of dest_path doesn’t
	// exist, though.
	//
	// The function takes the following parameters:
	//
	//    - destPath: destination row.
	//    - value: data being dropped.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if a drop is possible before dest_path.
	//
	RowDropPossible(destPath *TreePath, value *externglib.Value) bool
}

// TreeDragDest: interface for Drag-and-Drop destinations in GtkTreeView.
type TreeDragDest struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*TreeDragDest)(nil)
)

// TreeDragDester describes TreeDragDest's interface methods.
type TreeDragDester interface {
	externglib.Objector

	// DragDataReceived asks the TreeDragDest to insert a row before the path
	// dest, deriving the contents of the row from value.
	DragDataReceived(dest *TreePath, value *externglib.Value) bool
	// RowDropPossible determines whether a drop is possible before the given
	// dest_path, at the same depth as dest_path.
	RowDropPossible(destPath *TreePath, value *externglib.Value) bool
}

var _ TreeDragDester = (*TreeDragDest)(nil)

func ifaceInitTreeDragDester(gifacePtr, data C.gpointer) {
	iface := (*C.GtkTreeDragDestIface)(unsafe.Pointer(gifacePtr))
	iface.drag_data_received = (*[0]byte)(C._gotk4_gtk4_TreeDragDestIface_drag_data_received)
	iface.row_drop_possible = (*[0]byte)(C._gotk4_gtk4_TreeDragDestIface_row_drop_possible)
}

//export _gotk4_gtk4_TreeDragDestIface_drag_data_received
func _gotk4_gtk4_TreeDragDestIface_drag_data_received(arg0 *C.GtkTreeDragDest, arg1 *C.GtkTreePath, arg2 *C.GValue) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeDragDestOverrider)

	var _dest *TreePath          // out
	var _value *externglib.Value // out

	_dest = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_value = externglib.ValueFromNative(unsafe.Pointer(arg2))

	ok := iface.DragDataReceived(_dest, _value)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_TreeDragDestIface_row_drop_possible
func _gotk4_gtk4_TreeDragDestIface_row_drop_possible(arg0 *C.GtkTreeDragDest, arg1 *C.GtkTreePath, arg2 *C.GValue) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeDragDestOverrider)

	var _destPath *TreePath      // out
	var _value *externglib.Value // out

	_destPath = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_value = externglib.ValueFromNative(unsafe.Pointer(arg2))

	ok := iface.RowDropPossible(_destPath, _value)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapTreeDragDest(obj *externglib.Object) *TreeDragDest {
	return &TreeDragDest{
		Object: obj,
	}
}

func marshalTreeDragDest(p uintptr) (interface{}, error) {
	return wrapTreeDragDest(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DragDataReceived asks the TreeDragDest to insert a row before the path dest,
// deriving the contents of the row from value. If dest is outside the tree so
// that inserting before it is impossible, FALSE will be returned. Also, FALSE
// may be returned if the new row is not created for some model-specific reason.
// Should robustly handle a dest no longer found in the model!.
//
// The function takes the following parameters:
//
//    - dest: row to drop in front of.
//    - value: data to drop.
//
// The function returns the following values:
//
//    - ok: whether a new row was created before position dest.
//
func (dragDest *TreeDragDest) DragDataReceived(dest *TreePath, value *externglib.Value) bool {
	var _arg0 *C.GtkTreeDragDest // out
	var _arg1 *C.GtkTreePath     // out
	var _arg2 *C.GValue          // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeDragDest)(unsafe.Pointer(externglib.InternObject(dragDest).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(dest)))
	_arg2 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C.gtk_tree_drag_dest_drag_data_received(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dragDest)
	runtime.KeepAlive(dest)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowDropPossible determines whether a drop is possible before the given
// dest_path, at the same depth as dest_path. i.e., can we drop the data in
// value at that location. dest_path does not have to exist; the return value
// will almost certainly be FALSE if the parent of dest_path doesn’t exist,
// though.
//
// The function takes the following parameters:
//
//    - destPath: destination row.
//    - value: data being dropped.
//
// The function returns the following values:
//
//    - ok: TRUE if a drop is possible before dest_path.
//
func (dragDest *TreeDragDest) RowDropPossible(destPath *TreePath, value *externglib.Value) bool {
	var _arg0 *C.GtkTreeDragDest // out
	var _arg1 *C.GtkTreePath     // out
	var _arg2 *C.GValue          // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeDragDest)(unsafe.Pointer(externglib.InternObject(dragDest).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(destPath)))
	_arg2 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C.gtk_tree_drag_dest_row_drop_possible(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dragDest)
	runtime.KeepAlive(destPath)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TreeDragSourceOverrider contains methods that are overridable.
type TreeDragSourceOverrider interface {
	externglib.Objector
	// DragDataDelete asks the TreeDragSource to delete the row at path, because
	// it was moved somewhere else via drag-and-drop. Returns FALSE if the
	// deletion fails because path no longer exists, or for some model-specific
	// reason. Should robustly handle a path no longer found in the model!.
	//
	// The function takes the following parameters:
	//
	//    - path: row that was being dragged.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if the row was successfully deleted.
	//
	DragDataDelete(path *TreePath) bool
	// DragDataGet asks the TreeDragSource to return a ContentProvider
	// representing the row at path. Should robustly handle a path no longer
	// found in the model!.
	//
	// The function takes the following parameters:
	//
	//    - path: row that was dragged.
	//
	// The function returns the following values:
	//
	//    - contentProvider (optional) for the given path or NULL if none exists.
	//
	DragDataGet(path *TreePath) *gdk.ContentProvider
	// RowDraggable asks the TreeDragSource whether a particular row can be used
	// as the source of a DND operation. If the source doesn’t implement this
	// interface, the row is assumed draggable.
	//
	// The function takes the following parameters:
	//
	//    - path: row on which user is initiating a drag.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if the row can be dragged.
	//
	RowDraggable(path *TreePath) bool
}

// TreeDragSource: interface for Drag-and-Drop destinations in GtkTreeView.
type TreeDragSource struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*TreeDragSource)(nil)
)

// TreeDragSourcer describes TreeDragSource's interface methods.
type TreeDragSourcer interface {
	externglib.Objector

	// DragDataDelete asks the TreeDragSource to delete the row at path, because
	// it was moved somewhere else via drag-and-drop.
	DragDataDelete(path *TreePath) bool
	// DragDataGet asks the TreeDragSource to return a ContentProvider
	// representing the row at path.
	DragDataGet(path *TreePath) *gdk.ContentProvider
	// RowDraggable asks the TreeDragSource whether a particular row can be used
	// as the source of a DND operation.
	RowDraggable(path *TreePath) bool
}

var _ TreeDragSourcer = (*TreeDragSource)(nil)

func ifaceInitTreeDragSourcer(gifacePtr, data C.gpointer) {
	iface := (*C.GtkTreeDragSourceIface)(unsafe.Pointer(gifacePtr))
	iface.drag_data_delete = (*[0]byte)(C._gotk4_gtk4_TreeDragSourceIface_drag_data_delete)
	iface.drag_data_get = (*[0]byte)(C._gotk4_gtk4_TreeDragSourceIface_drag_data_get)
	iface.row_draggable = (*[0]byte)(C._gotk4_gtk4_TreeDragSourceIface_row_draggable)
}

//export _gotk4_gtk4_TreeDragSourceIface_drag_data_delete
func _gotk4_gtk4_TreeDragSourceIface_drag_data_delete(arg0 *C.GtkTreeDragSource, arg1 *C.GtkTreePath) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeDragSourceOverrider)

	var _path *TreePath // out

	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := iface.DragDataDelete(_path)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_TreeDragSourceIface_drag_data_get
func _gotk4_gtk4_TreeDragSourceIface_drag_data_get(arg0 *C.GtkTreeDragSource, arg1 *C.GtkTreePath) (cret *C.GdkContentProvider) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeDragSourceOverrider)

	var _path *TreePath // out

	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	contentProvider := iface.DragDataGet(_path)

	if contentProvider != nil {
		cret = (*C.GdkContentProvider)(unsafe.Pointer(externglib.InternObject(contentProvider).Native()))
		C.g_object_ref(C.gpointer(externglib.InternObject(contentProvider).Native()))
	}

	return cret
}

//export _gotk4_gtk4_TreeDragSourceIface_row_draggable
func _gotk4_gtk4_TreeDragSourceIface_row_draggable(arg0 *C.GtkTreeDragSource, arg1 *C.GtkTreePath) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeDragSourceOverrider)

	var _path *TreePath // out

	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := iface.RowDraggable(_path)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapTreeDragSource(obj *externglib.Object) *TreeDragSource {
	return &TreeDragSource{
		Object: obj,
	}
}

func marshalTreeDragSource(p uintptr) (interface{}, error) {
	return wrapTreeDragSource(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DragDataDelete asks the TreeDragSource to delete the row at path, because it
// was moved somewhere else via drag-and-drop. Returns FALSE if the deletion
// fails because path no longer exists, or for some model-specific reason.
// Should robustly handle a path no longer found in the model!.
//
// The function takes the following parameters:
//
//    - path: row that was being dragged.
//
// The function returns the following values:
//
//    - ok: TRUE if the row was successfully deleted.
//
func (dragSource *TreeDragSource) DragDataDelete(path *TreePath) bool {
	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(externglib.InternObject(dragSource).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_tree_drag_source_drag_data_delete(_arg0, _arg1)
	runtime.KeepAlive(dragSource)
	runtime.KeepAlive(path)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DragDataGet asks the TreeDragSource to return a ContentProvider representing
// the row at path. Should robustly handle a path no longer found in the model!.
//
// The function takes the following parameters:
//
//    - path: row that was dragged.
//
// The function returns the following values:
//
//    - contentProvider (optional) for the given path or NULL if none exists.
//
func (dragSource *TreeDragSource) DragDataGet(path *TreePath) *gdk.ContentProvider {
	var _arg0 *C.GtkTreeDragSource  // out
	var _arg1 *C.GtkTreePath        // out
	var _cret *C.GdkContentProvider // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(externglib.InternObject(dragSource).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_tree_drag_source_drag_data_get(_arg0, _arg1)
	runtime.KeepAlive(dragSource)
	runtime.KeepAlive(path)

	var _contentProvider *gdk.ContentProvider // out

	if _cret != nil {
		{
			obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
			_contentProvider = &gdk.ContentProvider{
				Object: obj,
			}
		}
	}

	return _contentProvider
}

// RowDraggable asks the TreeDragSource whether a particular row can be used as
// the source of a DND operation. If the source doesn’t implement this
// interface, the row is assumed draggable.
//
// The function takes the following parameters:
//
//    - path: row on which user is initiating a drag.
//
// The function returns the following values:
//
//    - ok: TRUE if the row can be dragged.
//
func (dragSource *TreeDragSource) RowDraggable(path *TreePath) bool {
	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(externglib.InternObject(dragSource).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_tree_drag_source_row_draggable(_arg0, _arg1)
	runtime.KeepAlive(dragSource)
	runtime.KeepAlive(path)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
