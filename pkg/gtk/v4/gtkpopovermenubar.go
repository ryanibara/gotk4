// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_popover_menu_bar_get_type()), F: marshalPopoverMenuBarrer},
	})
}

// PopoverMenuBarrer describes PopoverMenuBar's methods.
type PopoverMenuBarrer interface {
	// AddChild adds a custom widget to a generated menubar.
	AddChild(child Widgetter, id string) bool
	// MenuModel returns the model from which the contents of @bar are taken.
	MenuModel() *gio.MenuModel
	// RemoveChild removes a widget that has previously been added with
	// gtk_popover_menu_bar_add_child().
	RemoveChild(child Widgetter) bool
	// SetMenuModel sets a menu model from which @bar should take its contents.
	SetMenuModel(model gio.MenuModeller)
}

// PopoverMenuBar: `GtkPopoverMenuBar` presents a horizontal bar of items that
// pop up popover menus when clicked.
//
// !An example GtkPopoverMenuBar (menubar.png)
//
// The only way to create instances of `GtkPopoverMenuBar` is from a
// `GMenuModel`.
//
//
// CSS nodes
//
// “` menubar ├── item[.active] ┊ ╰── popover ╰── item ╰── popover “`
//
// `GtkPopoverMenuBar` has a single CSS node with name menubar, below which each
// item has its CSS node, and below that the corresponding popover.
//
// The item whose popover is currently open gets the .active style class.
//
//
// Accessibility
//
// `GtkPopoverMenuBar` uses the GTK_ACCESSIBLE_ROLE_MENU_BAR role, the menu
// items use the GTK_ACCESSIBLE_ROLE_MENU_ITEM role and the menus use the
// GTK_ACCESSIBLE_ROLE_MENU role.
type PopoverMenuBar struct {
	Widget
}

var (
	_ PopoverMenuBarrer = (*PopoverMenuBar)(nil)
	_ gextras.Nativer   = (*PopoverMenuBar)(nil)
)

func wrapPopoverMenuBar(obj *externglib.Object) PopoverMenuBarrer {
	return &PopoverMenuBar{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalPopoverMenuBarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPopoverMenuBar(obj), nil
}

// NewPopoverMenuBarFromModel creates a `GtkPopoverMenuBar` from a `GMenuModel`.
func NewPopoverMenuBarFromModel(model gio.MenuModeller) *PopoverMenuBar {
	var _arg1 *C.GMenuModel // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GMenuModel)(unsafe.Pointer((model).(gextras.Nativer).Native()))

	_cret = C.gtk_popover_menu_bar_new_from_model(_arg1)

	var _popoverMenuBar *PopoverMenuBar // out

	_popoverMenuBar = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*PopoverMenuBar)

	return _popoverMenuBar
}

// AddChild adds a custom widget to a generated menubar.
//
// For this to work, the menu model of @bar must have an item with a `custom`
// attribute that matches @id.
func (bar *PopoverMenuBar) AddChild(child Widgetter, id string) bool {
	var _arg0 *C.GtkPopoverMenuBar // out
	var _arg1 *C.GtkWidget         // out
	var _arg2 *C.char              // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkPopoverMenuBar)(unsafe.Pointer(bar.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.char)(C.CString(id))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_popover_menu_bar_add_child(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MenuModel returns the model from which the contents of @bar are taken.
func (bar *PopoverMenuBar) MenuModel() *gio.MenuModel {
	var _arg0 *C.GtkPopoverMenuBar // out
	var _cret *C.GMenuModel        // in

	_arg0 = (*C.GtkPopoverMenuBar)(unsafe.Pointer(bar.Native()))

	_cret = C.gtk_popover_menu_bar_get_menu_model(_arg0)

	var _menuModel *gio.MenuModel // out

	_menuModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gio.MenuModel)

	return _menuModel
}

// RemoveChild removes a widget that has previously been added with
// gtk_popover_menu_bar_add_child().
func (bar *PopoverMenuBar) RemoveChild(child Widgetter) bool {
	var _arg0 *C.GtkPopoverMenuBar // out
	var _arg1 *C.GtkWidget         // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkPopoverMenuBar)(unsafe.Pointer(bar.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	_cret = C.gtk_popover_menu_bar_remove_child(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetMenuModel sets a menu model from which @bar should take its contents.
func (bar *PopoverMenuBar) SetMenuModel(model gio.MenuModeller) {
	var _arg0 *C.GtkPopoverMenuBar // out
	var _arg1 *C.GMenuModel        // out

	_arg0 = (*C.GtkPopoverMenuBar)(unsafe.Pointer(bar.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer((model).(gextras.Nativer).Native()))

	C.gtk_popover_menu_bar_set_menu_model(_arg0, _arg1)
}
