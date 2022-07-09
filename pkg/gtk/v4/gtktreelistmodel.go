// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void callbackDelete(gpointer);
// extern void* _gotk4_gtk4_TreeListModelCreateModelFunc(gpointer, gpointer);
import "C"

// GTypeTreeListModel returns the GType for the type TreeListModel.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTreeListModel() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "TreeListModel").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalTreeListModel)
	return gtype
}

// GTypeTreeListRow returns the GType for the type TreeListRow.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTreeListRow() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "TreeListRow").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalTreeListRow)
	return gtype
}

// TreeListModelCreateModelFunc: prototype of the function called to create new
// child models when gtk_tree_list_row_set_expanded() is called.
//
// This function can return NULL to indicate that item is guaranteed to be a
// leaf node and will never have children. If it does not have children but may
// get children later, it should return an empty model that is filled once
// children arrive.
type TreeListModelCreateModelFunc func(item *coreglib.Object) (listModel *gio.ListModel)

//export _gotk4_gtk4_TreeListModelCreateModelFunc
func _gotk4_gtk4_TreeListModelCreateModelFunc(arg1 C.gpointer, arg2 C.gpointer) (cret *C.void) {
	var fn TreeListModelCreateModelFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TreeListModelCreateModelFunc)
	}

	var _item *coreglib.Object // out

	_item = coreglib.Take(unsafe.Pointer(arg1))

	listModel := fn(_item)

	if listModel != nil {
		cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(listModel).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(listModel).Native()))
	}

	return cret
}

// TreeListModelOverrider contains methods that are overridable.
type TreeListModelOverrider interface {
}

// TreeListModel: GtkTreeListModel is a list model that can create child models
// on demand.
type TreeListModel struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.ListModel
}

var (
	_ coreglib.Objector = (*TreeListModel)(nil)
)

func classInitTreeListModeller(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTreeListModel(obj *coreglib.Object) *TreeListModel {
	return &TreeListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalTreeListModel(p uintptr) (interface{}, error) {
	return wrapTreeListModel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTreeListModel creates a new empty GtkTreeListModel displaying root with
// all rows collapsed.
//
// The function takes the following parameters:
//
//    - root: GListModel to use as root.
//    - passthrough: TRUE to pass through items from the models.
//    - autoexpand: TRUE to set the autoexpand property and expand the root
//      model.
//    - createFunc: function to call to create the GListModel for the children of
//      an item.
//
// The function returns the following values:
//
//    - treeListModel: newly created GtkTreeListModel.
//
func NewTreeListModel(root gio.ListModeller, passthrough, autoexpand bool, createFunc TreeListModelCreateModelFunc) *TreeListModel {
	var _args [6]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(root).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(root).Native()))
	if passthrough {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}
	if autoexpand {
		*(*C.gboolean)(unsafe.Pointer(&_args[2])) = C.TRUE
	}
	*(*C.gpointer)(unsafe.Pointer(&_args[3])) = (*[0]byte)(C._gotk4_gtk4_TreeListModelCreateModelFunc)
	_args[4] = C.gpointer(gbox.Assign(createFunc))
	_args[5] = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	_info := girepository.MustFind("Gtk", "TreeListModel")
	_gret := _info.InvokeClassMethod("new_TreeListModel", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(root)
	runtime.KeepAlive(passthrough)
	runtime.KeepAlive(autoexpand)
	runtime.KeepAlive(createFunc)

	var _treeListModel *TreeListModel // out

	_treeListModel = wrapTreeListModel(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _treeListModel
}

// Autoexpand gets whether the model is set to automatically expand new rows
// that get added.
//
// This can be either rows added by changes to the underlying models or via
// gtk.TreeListRow.SetExpanded().
//
// The function returns the following values:
//
//    - ok: TRUE if the model is set to autoexpand.
//
func (self *TreeListModel) Autoexpand() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "TreeListModel")
	_gret := _info.InvokeClassMethod("get_autoexpand", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ChildRow gets the row item corresponding to the child at index position for
// self's root model.
//
// If position is greater than the number of children in the root model, NULL is
// returned.
//
// Do not confuse this function with gtk.TreeListModel.GetRow().
//
// The function takes the following parameters:
//
//    - position of the child to get.
//
// The function returns the following values:
//
//    - treeListRow (optional): child in position.
//
func (self *TreeListModel) ChildRow(position uint32) *TreeListRow {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(position)

	_info := girepository.MustFind("Gtk", "TreeListModel")
	_gret := _info.InvokeClassMethod("get_child_row", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(position)

	var _treeListRow *TreeListRow // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_treeListRow = wrapTreeListRow(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _treeListRow
}

// Model gets the root model that self was created with.
//
// The function returns the following values:
//
//    - listModel: root model.
//
func (self *TreeListModel) Model() *gio.ListModel {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "TreeListModel")
	_gret := _info.InvokeClassMethod("get_model", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _listModel *gio.ListModel // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_listModel = &gio.ListModel{
			Object: obj,
		}
	}

	return _listModel
}

// Passthrough gets whether the model is passing through original row items.
//
// If this function returns FALSE, the GListModel functions for self return
// custom GtkTreeListRow objects. You need to call gtk.TreeListRow.GetItem() on
// these objects to get the original item.
//
// If TRUE, the values of the child models are passed through in their original
// state. You then need to call gtk.TreeListModel.GetRow() to get the custom
// GtkTreeListRows.
//
// The function returns the following values:
//
//    - ok: TRUE if the model is passing through original row items.
//
func (self *TreeListModel) Passthrough() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "TreeListModel")
	_gret := _info.InvokeClassMethod("get_passthrough", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Row gets the row object for the given row.
//
// If position is greater than the number of items in self, NULL is returned.
//
// The row object can be used to expand and collapse rows as well as to inspect
// its position in the tree. See its documentation for details.
//
// This row object is persistent and will refer to the current item as long as
// the row is present in self, independent of other rows being added or removed.
//
// If self is set to not be passthrough, this function is equivalent to calling
// g_list_model_get_item().
//
// Do not confuse this function with gtk.TreeListModel.GetChildRow().
//
// The function takes the following parameters:
//
//    - position of the row to fetch.
//
// The function returns the following values:
//
//    - treeListRow (optional): row item.
//
func (self *TreeListModel) Row(position uint32) *TreeListRow {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(position)

	_info := girepository.MustFind("Gtk", "TreeListModel")
	_gret := _info.InvokeClassMethod("get_row", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(position)

	var _treeListRow *TreeListRow // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_treeListRow = wrapTreeListRow(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _treeListRow
}

// SetAutoexpand sets whether the model should autoexpand.
//
// If set to TRUE, the model will recursively expand all rows that get added to
// the model. This can be either rows added by changes to the underlying models
// or via gtk.TreeListRow.SetExpanded().
//
// The function takes the following parameters:
//
//    - autoexpand: TRUE to make the model autoexpand its rows.
//
func (self *TreeListModel) SetAutoexpand(autoexpand bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if autoexpand {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "TreeListModel")
	_info.InvokeClassMethod("set_autoexpand", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(autoexpand)
}

// TreeListRowOverrider contains methods that are overridable.
type TreeListRowOverrider interface {
}

// TreeListRow: GtkTreeListRow is used by GtkTreeListModel to represent items.
//
// It allows navigating the model as a tree and modify the state of rows.
//
// GtkTreeListRow instances are created by a GtkTreeListModel only when the
// gtk.TreeListModel:passthrough property is not set.
//
// There are various support objects that can make use of GtkTreeListRow
// objects, such as the gtk.TreeExpander widget that allows displaying an icon
// to expand or collapse a row or gtk.TreeListRowSorter that makes it possible
// to sort trees properly.
type TreeListRow struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TreeListRow)(nil)
)

func classInitTreeListRower(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTreeListRow(obj *coreglib.Object) *TreeListRow {
	return &TreeListRow{
		Object: obj,
	}
}

func marshalTreeListRow(p uintptr) (interface{}, error) {
	return wrapTreeListRow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ChildRow: if self is not expanded or position is greater than the number of
// children, NULL is returned.
//
// The function takes the following parameters:
//
//    - position of the child to get.
//
// The function returns the following values:
//
//    - treeListRow (optional): child in position.
//
func (self *TreeListRow) ChildRow(position uint32) *TreeListRow {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(position)

	_info := girepository.MustFind("Gtk", "TreeListRow")
	_gret := _info.InvokeClassMethod("get_child_row", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(position)

	var _treeListRow *TreeListRow // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_treeListRow = wrapTreeListRow(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _treeListRow
}

// Children: if the row is expanded, gets the model holding the children of
// self.
//
// This model is the model created by the gtk.TreeListModelCreateModelFunc and
// contains the original items, no matter what value
// gtk.TreeListModel:passthrough is set to.
//
// The function returns the following values:
//
//    - listModel (optional): model containing the children.
//
func (self *TreeListRow) Children() *gio.ListModel {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "TreeListRow")
	_gret := _info.InvokeClassMethod("get_children", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _listModel *gio.ListModel // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_listModel = &gio.ListModel{
				Object: obj,
			}
		}
	}

	return _listModel
}

// Depth gets the depth of this row.
//
// Rows that correspond to items in the root model have a depth of zero, rows
// corresponding to items of models of direct children of the root model have a
// depth of 1 and so on.
//
// The depth of a row never changes until the row is destroyed.
//
// The function returns the following values:
//
//    - guint: depth of this row.
//
func (self *TreeListRow) Depth() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "TreeListRow")
	_gret := _info.InvokeClassMethod("get_depth", _args[:], nil)
	_cret := *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// Expanded gets if a row is currently expanded.
//
// The function returns the following values:
//
//    - ok: TRUE if the row is expanded.
//
func (self *TreeListRow) Expanded() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "TreeListRow")
	_gret := _info.InvokeClassMethod("get_expanded", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Item gets the item corresponding to this row,
//
// The value returned by this function never changes until the row is destroyed.
//
// The function returns the following values:
//
//    - object (optional): item of this row or NULL when the row was destroyed.
//
func (self *TreeListRow) Item() *coreglib.Object {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "TreeListRow")
	_gret := _info.InvokeClassMethod("get_item", _args[:], nil)
	_cret := *(*C.gpointer)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _object *coreglib.Object // out

	_object = coreglib.AssumeOwnership(unsafe.Pointer(_cret))

	return _object
}

// Parent gets the row representing the parent for self.
//
// That is the row that would need to be collapsed to make this row disappear.
//
// If self is a row corresponding to the root model, NULL is returned.
//
// The value returned by this function never changes until the row is destroyed.
//
// The function returns the following values:
//
//    - treeListRow (optional): parent of self.
//
func (self *TreeListRow) Parent() *TreeListRow {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "TreeListRow")
	_gret := _info.InvokeClassMethod("get_parent", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _treeListRow *TreeListRow // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_treeListRow = wrapTreeListRow(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _treeListRow
}

// Position returns the position in the GtkTreeListModel that self occupies at
// the moment.
//
// The function returns the following values:
//
//    - guint: position in the model.
//
func (self *TreeListRow) Position() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "TreeListRow")
	_gret := _info.InvokeClassMethod("get_position", _args[:], nil)
	_cret := *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// IsExpandable checks if a row can be expanded.
//
// This does not mean that the row is actually expanded, this can be checked
// with gtk.TreeListRow.GetExpanded().
//
// If a row is expandable never changes until the row is destroyed.
//
// The function returns the following values:
//
//    - ok: TRUE if the row is expandable.
//
func (self *TreeListRow) IsExpandable() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "TreeListRow")
	_gret := _info.InvokeClassMethod("is_expandable", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetExpanded expands or collapses a row.
//
// If a row is expanded, the model of calling the
// gtk.TreeListModelCreateModelFunc for the row's item will be inserted after
// this row. If a row is collapsed, those items will be removed from the model.
//
// If the row is not expandable, this function does nothing.
//
// The function takes the following parameters:
//
//    - expanded: TRUE if the row should be expanded.
//
func (self *TreeListRow) SetExpanded(expanded bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if expanded {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "TreeListRow")
	_info.InvokeClassMethod("set_expanded", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(expanded)
}
