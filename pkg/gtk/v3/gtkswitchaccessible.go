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
		{T: externglib.Type(C.gtk_switch_accessible_get_type()), F: marshalSwitchAccessible},
	})
}

type SwitchAccessible interface {
	WidgetAccessible
}

// switchAccessible implements the SwitchAccessible interface.
type switchAccessible struct {
	WidgetAccessible
}

var _ SwitchAccessible = (*switchAccessible)(nil)

// WrapSwitchAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapSwitchAccessible(obj *externglib.Object) SwitchAccessible {
	return SwitchAccessible{
		WidgetAccessible: WrapWidgetAccessible(obj),
	}
}

func marshalSwitchAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSwitchAccessible(obj), nil
}
