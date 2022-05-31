// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtklistitem.go.
var GTypeListItem = coreglib.Type(C.gtk_list_item_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeListItem, F: marshalListItem},
	})
}

// ListItemOverrider contains methods that are overridable.
type ListItemOverrider interface {
}

// ListItem: GtkListItem is used by list widgets to represent items in a
// GListModel.
//
// The GtkListItems are managed by the list widget (with its factory) and cannot
// be created by applications, but they need to be populated by application
// code. This is done by calling gtk.ListItem.SetChild().
//
// GtkListItems exist in 2 stages:
//
// 1. The unbound stage where the listitem is not currently connected to an item
// in the list. In that case, the gtk.ListItem:item property is set to NULL.
//
// 2. The bound stage where the listitem references an item from the list. The
// gtk.ListItem:item property is not NULL.
type ListItem struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ListItem)(nil)
)

func classInitListItemmer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapListItem(obj *coreglib.Object) *ListItem {
	return &ListItem{
		Object: obj,
	}
}

func marshalListItem(p uintptr) (interface{}, error) {
	return wrapListItem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Activatable checks if a list item has been set to be activatable via
// gtk_list_item_set_activatable().
//
// The function returns the following values:
//
//    - ok: TRUE if the item is activatable.
//
func (self *ListItem) Activatable() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**ListItem)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ListItem").InvokeMethod("get_activatable", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Child gets the child previously set via gtk_list_item_set_child() or NULL if
// none was set.
//
// The function returns the following values:
//
//    - widget (optional): child.
//
func (self *ListItem) Child() Widgetter {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**ListItem)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ListItem").InvokeMethod("get_child", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Position gets the position in the model that self currently displays.
//
// If self is unbound, GTK_INVALID_LIST_POSITION is returned.
//
// The function returns the following values:
//
//    - guint: position of this item.
//
func (self *ListItem) Position() uint32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**ListItem)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ListItem").InvokeMethod("get_position", args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// Selectable checks if a list item has been set to be selectable via
// gtk_list_item_set_selectable().
//
// Do not confuse this function with gtk.ListItem.GetSelected().
//
// The function returns the following values:
//
//    - ok: TRUE if the item is selectable.
//
func (self *ListItem) Selectable() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**ListItem)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ListItem").InvokeMethod("get_selectable", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Selected checks if the item is displayed as selected.
//
// The selected state is maintained by the liste widget and its model and cannot
// be set otherwise.
//
// The function returns the following values:
//
//    - ok: TRUE if the item is selected.
//
func (self *ListItem) Selected() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**ListItem)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ListItem").InvokeMethod("get_selected", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActivatable sets self to be activatable.
//
// If an item is activatable, double-clicking on the item, using the Return key
// or calling gtk_widget_activate() will activate the item. Activating instructs
// the containing view to handle activation. GtkListView for example will be
// emitting the gtk.ListView::activate signal.
//
// By default, list items are activatable.
//
// The function takes the following parameters:
//
//    - activatable: if the item should be activatable.
//
func (self *ListItem) SetActivatable(activatable bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if activatable {
		_arg1 = C.TRUE
	}
	*(**ListItem)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ListItem").InvokeMethod("set_activatable", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(activatable)
}

// SetChild sets the child to be used for this listitem.
//
// This function is typically called by applications when setting up a listitem
// so that the widget can be reused when binding it multiple times.
//
// The function takes the following parameters:
//
//    - child (optional): list item's child or NULL to unset.
//
func (self *ListItem) SetChild(child Widgetter) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}
	*(**ListItem)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ListItem").InvokeMethod("set_child", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetSelectable sets self to be selectable.
//
// If an item is selectable, clicking on the item or using the keyboard will try
// to select or unselect the item. If this succeeds is up to the model to
// determine, as it is managing the selected state.
//
// Note that this means that making an item non-selectable has no influence on
// the selected state at all. A non-selectable item may still be selected.
//
// By default, list items are selectable. When rebinding them to a new item,
// they will also be reset to be selectable by GTK.
//
// The function takes the following parameters:
//
//    - selectable: if the item should be selectable.
//
func (self *ListItem) SetSelectable(selectable bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if selectable {
		_arg1 = C.TRUE
	}
	*(**ListItem)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ListItem").InvokeMethod("set_selectable", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(selectable)
}
