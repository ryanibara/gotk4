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
		{T: externglib.Type(C.gtk_plug_accessible_get_type()), F: marshalPlugAccessible},
	})
}

type PlugAccessible interface {
	WindowAccessible

	ID() string
}

// plugAccessible implements the PlugAccessible class.
type plugAccessible struct {
	WindowAccessible
}

// WrapPlugAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapPlugAccessible(obj *externglib.Object) PlugAccessible {
	return plugAccessible{
		WindowAccessible: WrapWindowAccessible(obj),
	}
}

func marshalPlugAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPlugAccessible(obj), nil
}

func (p plugAccessible) ID() string {
	var _arg0 *C.GtkPlugAccessible // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkPlugAccessible)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_plug_accessible_get_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
