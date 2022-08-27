// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// NewMenuFromModel creates a Menu and populates it with menu items and submenus
// according to model.
//
// The created menu items are connected to actions found in the
// ApplicationWindow to which the menu belongs - typically by means of being
// attached to a widget (see gtk_menu_attach_to_widget()) that is contained
// within the ApplicationWindows widget hierarchy.
//
// Actions can also be added using gtk_widget_insert_action_group() on the
// menu's attach widget or on any of its parent widgets.
//
// The function takes the following parameters:
//
//    - model: Model.
//
// The function returns the following values:
//
//    - menu: new Menu.
//
func NewMenuFromModel(model gio.MenuModeller) *Menu {
	var _arg1 *C.GMenuModel // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))

	_cret = C.gtk_menu_new_from_model(_arg1)
	runtime.KeepAlive(model)

	var _menu *Menu // out

	_menu = wrapMenu(coreglib.Take(unsafe.Pointer(_cret)))

	return _menu
}
