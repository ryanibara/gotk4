// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_list_item_get_type()), F: marshalListItemer},
	})
}

// ListItemer describes ListItem's methods.
type ListItemer interface {
	// Activatable checks if a list item has been set to be activatable via
	// gtk_list_item_set_activatable().
	Activatable() bool
	// Child gets the child previously set via gtk_list_item_set_child() or nil
	// if none was set.
	Child() *Widget
	// Item gets the model item that associated with @self.
	Item() *externglib.Object
	// Position gets the position in the model that @self currently displays.
	Position() uint
	// Selectable checks if a list item has been set to be selectable via
	// gtk_list_item_set_selectable().
	Selectable() bool
	// Selected checks if the item is displayed as selected.
	Selected() bool
	// SetActivatable sets @self to be activatable.
	SetActivatable(activatable bool)
	// SetChild sets the child to be used for this listitem.
	SetChild(child Widgeter)
	// SetSelectable sets @self to be selectable.
	SetSelectable(selectable bool)
}

// ListItem: `GtkListItem` is used by list widgets to represent items in a
// `GListModel`.
//
// The `GtkListItem`s are managed by the list widget (with its factory) and
// cannot be created by applications, but they need to be populated by
// application code. This is done by calling [method@Gtk.ListItem.set_child].
//
// `GtkListItem`s exist in 2 stages:
//
// 1. The unbound stage where the listitem is not currently connected to an item
// in the list. In that case, the [property@Gtk.ListItem:item] property is set
// to nil.
//
// 2. The bound stage where the listitem references an item from the list. The
// [property@Gtk.ListItem:item] property is not nil.
type ListItem struct {
	*externglib.Object
}

var (
	_ ListItemer      = (*ListItem)(nil)
	_ gextras.Nativer = (*ListItem)(nil)
)

func wrapListItem(obj *externglib.Object) *ListItem {
	return &ListItem{
		Object: obj,
	}
}

func marshalListItemer(p uintptr) (interface{}, error) {
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

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Child gets the child previously set via gtk_list_item_set_child() or nil if
// none was set.
func (self *ListItem) Child() *Widget {
	var _arg0 *C.GtkListItem // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_list_item_get_child(_arg0)

	var _widget *Widget // out

	_widget = wrapWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _widget
}

// Item gets the model item that associated with @self.
//
// If @self is unbound, this function returns nil.
func (self *ListItem) Item() *externglib.Object {
	var _arg0 *C.GtkListItem // out
	var _cret C.gpointer     // in

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_list_item_get_item(_arg0)

	var _object *externglib.Object // out

	_object = externglib.Take(unsafe.Pointer(_cret))

	return _object
}

// Position gets the position in the model that @self currently displays.
//
// If @self is unbound, GTK_INVALID_LIST_POSITION is returned.
func (self *ListItem) Position() uint {
	var _arg0 *C.GtkListItem // out
	var _cret C.guint        // in

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_list_item_get_position(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Selectable checks if a list item has been set to be selectable via
// gtk_list_item_set_selectable().
//
// Do not confuse this function with [method@Gtk.ListItem.get_selected].
func (self *ListItem) Selectable() bool {
	var _arg0 *C.GtkListItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_list_item_get_selectable(_arg0)

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

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActivatable sets @self to be activatable.
//
// If an item is activatable, double-clicking on the item, using the Return key
// or calling gtk_widget_activate() will activate the item. Activating instructs
// the containing view to handle activation. `GtkListView` for example will be
// emitting the [signal@Gtk.ListView::activate] signal.
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
}

// SetChild sets the child to be used for this listitem.
//
// This function is typically called by applications when setting up a listitem
// so that the widget can be reused when binding it multiple times.
func (self *ListItem) SetChild(child Widgeter) {
	var _arg0 *C.GtkListItem // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_list_item_set_child(_arg0, _arg1)
}

// SetSelectable sets @self to be selectable.
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
}
