// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypePopoverMenuBar returns the GType for the type PopoverMenuBar.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypePopoverMenuBar() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "PopoverMenuBar").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalPopoverMenuBar)
	return gtype
}

// PopoverMenuBar: GtkPopoverMenuBar presents a horizontal bar of items that pop
// up popover menus when clicked.
//
// !An example GtkPopoverMenuBar (menubar.png)
//
// The only way to create instances of GtkPopoverMenuBar is from a GMenuModel.
//
// CSS nodes
//
//    menubar
//    ├── item[.active]
//    ┊   ╰── popover
//    ╰── item
//        ╰── popover
//
//
// GtkPopoverMenuBar has a single CSS node with name menubar, below which each
// item has its CSS node, and below that the corresponding popover.
//
// The item whose popover is currently open gets the .active style class.
//
//
// Accessibility
//
// GtkPopoverMenuBar uses the GTK_ACCESSIBLE_ROLE_MENU_BAR role, the menu items
// use the GTK_ACCESSIBLE_ROLE_MENU_ITEM role and the menus use the
// GTK_ACCESSIBLE_ROLE_MENU role.
type PopoverMenuBar struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*PopoverMenuBar)(nil)
)

func wrapPopoverMenuBar(obj *coreglib.Object) *PopoverMenuBar {
	return &PopoverMenuBar{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
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

func marshalPopoverMenuBar(p uintptr) (interface{}, error) {
	return wrapPopoverMenuBar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPopoverMenuBarFromModel creates a GtkPopoverMenuBar from a GMenuModel.
//
// The function takes the following parameters:
//
//    - model (optional): GMenuModel, or NULL.
//
// The function returns the following values:
//
//    - popoverMenuBar: new GtkPopoverMenuBar.
//
func NewPopoverMenuBarFromModel(model gio.MenuModeller) *PopoverMenuBar {
	var _args [1]girepository.Argument

	if model != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	_gret := girepository.MustFind("Gtk", "PopoverMenuBar").InvokeMethod("new_PopoverMenuBar_from_model", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(model)

	var _popoverMenuBar *PopoverMenuBar // out

	_popoverMenuBar = wrapPopoverMenuBar(coreglib.Take(unsafe.Pointer(_cret)))

	return _popoverMenuBar
}

// AddChild adds a custom widget to a generated menubar.
//
// For this to work, the menu model of bar must have an item with a custom
// attribute that matches id.
//
// The function takes the following parameters:
//
//    - child: GtkWidget to add.
//    - id: ID to insert child at.
//
// The function returns the following values:
//
//    - ok: TRUE if id was found and the widget added.
//
func (bar *PopoverMenuBar) AddChild(child Widgetter, id string) bool {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(_args[2]))

	_gret := girepository.MustFind("Gtk", "PopoverMenuBar").InvokeMethod("add_child", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bar)
	runtime.KeepAlive(child)
	runtime.KeepAlive(id)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// MenuModel returns the model from which the contents of bar are taken.
//
// The function returns the following values:
//
//    - menuModel: GMenuModel.
//
func (bar *PopoverMenuBar) MenuModel() gio.MenuModeller {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_gret := girepository.MustFind("Gtk", "PopoverMenuBar").InvokeMethod("get_menu_model", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bar)

	var _menuModel gio.MenuModeller // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.MenuModeller is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gio.MenuModeller)
			return ok
		})
		rv, ok := casted.(gio.MenuModeller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.MenuModeller")
		}
		_menuModel = rv
	}

	return _menuModel
}

// RemoveChild removes a widget that has previously been added with
// gtk_popover_menu_bar_add_child().
//
// The function takes the following parameters:
//
//    - child to remove.
//
// The function returns the following values:
//
//    - ok: TRUE if the widget was removed.
//
func (bar *PopoverMenuBar) RemoveChild(child Widgetter) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_gret := girepository.MustFind("Gtk", "PopoverMenuBar").InvokeMethod("remove_child", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bar)
	runtime.KeepAlive(child)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetMenuModel sets a menu model from which bar should take its contents.
//
// The function takes the following parameters:
//
//    - model (optional): GMenuModel, or NULL.
//
func (bar *PopoverMenuBar) SetMenuModel(model gio.MenuModeller) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if model != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	girepository.MustFind("Gtk", "PopoverMenuBar").InvokeMethod("set_menu_model", _args[:], nil)

	runtime.KeepAlive(bar)
	runtime.KeepAlive(model)
}
