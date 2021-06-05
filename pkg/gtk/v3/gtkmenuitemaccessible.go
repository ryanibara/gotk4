// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_menu_item_accessible_get_type()), F: marshalMenuItemAccessible},
	})
}

type MenuItemAccessible interface {
	ContainerAccessible
}

// menuItemAccessible implements the MenuItemAccessible interface.
type menuItemAccessible struct {
	ContainerAccessible
}

var _ MenuItemAccessible = (*menuItemAccessible)(nil)

// WrapMenuItemAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapMenuItemAccessible(obj *externglib.Object) MenuItemAccessible {
	return MenuItemAccessible{
		ContainerAccessible: WrapContainerAccessible(obj),
	}
}

func marshalMenuItemAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMenuItemAccessible(obj), nil
}
