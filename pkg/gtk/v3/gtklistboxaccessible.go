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
		{T: externglib.Type(C.gtk_list_box_accessible_get_type()), F: marshalListBoxAccessible},
	})
}

type ListBoxAccessible interface {
	ContainerAccessible
}

// listBoxAccessible implements the ListBoxAccessible interface.
type listBoxAccessible struct {
	ContainerAccessible
}

var _ ListBoxAccessible = (*listBoxAccessible)(nil)

// WrapListBoxAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapListBoxAccessible(obj *externglib.Object) ListBoxAccessible {
	return ListBoxAccessible{
		ContainerAccessible: WrapContainerAccessible(obj),
	}
}

func marshalListBoxAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapListBoxAccessible(obj), nil
}
