// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_list_box_row_accessible_get_type()), F: marshalListBoxRowAccessible},
	})
}

type ListBoxRowAccessible interface {
	ContainerAccessible
}

// listBoxRowAccessible implements the ListBoxRowAccessible class.
type listBoxRowAccessible struct {
	ContainerAccessible
}

// WrapListBoxRowAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapListBoxRowAccessible(obj *externglib.Object) ListBoxRowAccessible {
	return listBoxRowAccessible{
		ContainerAccessible: WrapContainerAccessible(obj),
	}
}

func marshalListBoxRowAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapListBoxRowAccessible(obj), nil
}
