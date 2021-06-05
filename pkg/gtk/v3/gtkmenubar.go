// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_menu_bar_get_type()), F: marshalMenuBar},
	})
}

// MenuBar: the MenuBar is a subclass of MenuShell which contains one or more
// MenuItems. The result is a standard menu bar which can hold many menu items.
//
//
// CSS nodes
//
// GtkMenuBar has a single CSS node with name menubar.
type MenuBar interface {
	MenuShell
	Buildable

	// ChildPackDirection retrieves the current child pack direction of the
	// menubar. See gtk_menu_bar_set_child_pack_direction().
	ChildPackDirection() PackDirection
	// PackDirection retrieves the current pack direction of the menubar. See
	// gtk_menu_bar_set_pack_direction().
	PackDirection() PackDirection
	// SetChildPackDirection sets how widgets should be packed inside the
	// children of a menubar.
	SetChildPackDirection(childPackDir PackDirection)
	// SetPackDirection sets how items should be packed inside a menubar.
	SetPackDirection(packDir PackDirection)
}

// menuBar implements the MenuBar interface.
type menuBar struct {
	MenuShell
	Buildable
}

var _ MenuBar = (*menuBar)(nil)

// WrapMenuBar wraps a GObject to the right type. It is
// primarily used internally.
func WrapMenuBar(obj *externglib.Object) MenuBar {
	return MenuBar{
		MenuShell: WrapMenuShell(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalMenuBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMenuBar(obj), nil
}

// NewMenuBar constructs a class MenuBar.
func NewMenuBar() MenuBar {
	var cret C.GtkMenuBar
	var ret1 MenuBar

	cret = C.gtk_menu_bar_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(MenuBar)

	return ret1
}

// NewMenuBarFromModel constructs a class MenuBar.
func NewMenuBarFromModel(model gio.MenuModel) MenuBar {
	var arg1 *C.GMenuModel

	arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	var cret C.GtkMenuBar
	var ret1 MenuBar

	cret = C.gtk_menu_bar_new_from_model(model)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(MenuBar)

	return ret1
}

// ChildPackDirection retrieves the current child pack direction of the
// menubar. See gtk_menu_bar_set_child_pack_direction().
func (m menuBar) ChildPackDirection() PackDirection {
	var arg0 *C.GtkMenuBar

	arg0 = (*C.GtkMenuBar)(unsafe.Pointer(m.Native()))

	var cret C.GtkPackDirection
	var ret1 PackDirection

	cret = C.gtk_menu_bar_get_child_pack_direction(arg0)

	ret1 = PackDirection(cret)

	return ret1
}

// PackDirection retrieves the current pack direction of the menubar. See
// gtk_menu_bar_set_pack_direction().
func (m menuBar) PackDirection() PackDirection {
	var arg0 *C.GtkMenuBar

	arg0 = (*C.GtkMenuBar)(unsafe.Pointer(m.Native()))

	var cret C.GtkPackDirection
	var ret1 PackDirection

	cret = C.gtk_menu_bar_get_pack_direction(arg0)

	ret1 = PackDirection(cret)

	return ret1
}

// SetChildPackDirection sets how widgets should be packed inside the
// children of a menubar.
func (m menuBar) SetChildPackDirection(childPackDir PackDirection) {
	var arg0 *C.GtkMenuBar
	var arg1 C.GtkPackDirection

	arg0 = (*C.GtkMenuBar)(unsafe.Pointer(m.Native()))
	arg1 = (C.GtkPackDirection)(childPackDir)

	C.gtk_menu_bar_set_child_pack_direction(arg0, childPackDir)
}

// SetPackDirection sets how items should be packed inside a menubar.
func (m menuBar) SetPackDirection(packDir PackDirection) {
	var arg0 *C.GtkMenuBar
	var arg1 C.GtkPackDirection

	arg0 = (*C.GtkMenuBar)(unsafe.Pointer(m.Native()))
	arg1 = (C.GtkPackDirection)(packDir)

	C.gtk_menu_bar_set_pack_direction(arg0, packDir)
}
