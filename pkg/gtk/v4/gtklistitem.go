// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_list_item_get_type()), F: marshalListItemmer},
	})
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
	*externglib.Object
}

func wrapListItem(obj *externglib.Object) *ListItem {
	return &ListItem{
		Object: obj,
	}
}

func marshalListItemmer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapListItem(obj), nil
}

// Activatable checks if a list item has been set to be activatable via
// gtk_list_item_set_activatable().
func (self *ListItem) Activatable() bool {
	var _arg0 *C.GtkListItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_list_item_get_activatable(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Child gets the child previously set via gtk_list_item_set_child() or NULL if
// none was set.
func (self *ListItem) Child() Widgetter {
	var _arg0 *C.GtkListItem // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_list_item_get_child(_arg0)
	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if _cret != nil {
		{
			object := externglib.Take(unsafe.Pointer(_cret))
			rv, ok := (externglib.CastObject(object)).(Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Item gets the model item that associated with self.
//
// If self is unbound, this function returns NULL.
func (self *ListItem) Item() *externglib.Object {
	var _arg0 *C.GtkListItem // out
	var _cret C.gpointer     // in

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_list_item_get_item(_arg0)
	runtime.KeepAlive(self)

	var _object *externglib.Object // out

	_object = externglib.Take(unsafe.Pointer(_cret))

	return _object
}

// Position gets the position in the model that self currently displays.
//
// If self is unbound, GTK_INVALID_LIST_POSITION is returned.
func (self *ListItem) Position() uint {
	var _arg0 *C.GtkListItem // out
	var _cret C.guint        // in

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_list_item_get_position(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Selectable checks if a list item has been set to be selectable via
// gtk_list_item_set_selectable().
//
// Do not confuse this function with gtk.ListItem.GetSelected().
func (self *ListItem) Selectable() bool {
	var _arg0 *C.GtkListItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_list_item_get_selectable(_arg0)
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
func (self *ListItem) Selected() bool {
	var _arg0 *C.GtkListItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_list_item_get_selected(_arg0)
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
func (self *ListItem) SetActivatable(activatable bool) {
	var _arg0 *C.GtkListItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(self.Native()))
	if activatable {
		_arg1 = C.TRUE
	}

	C.gtk_list_item_set_activatable(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(activatable)
}

// SetChild sets the child to be used for this listitem.
//
// This function is typically called by applications when setting up a listitem
// so that the widget can be reused when binding it multiple times.
func (self *ListItem) SetChild(child Widgetter) {
	var _arg0 *C.GtkListItem // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(self.Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	}

	C.gtk_list_item_set_child(_arg0, _arg1)
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
func (self *ListItem) SetSelectable(selectable bool) {
	var _arg0 *C.GtkListItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(self.Native()))
	if selectable {
		_arg1 = C.TRUE
	}

	C.gtk_list_item_set_selectable(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(selectable)
}
