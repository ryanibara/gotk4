// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_no_selection_get_type()), F: marshalNoSelection},
	})
}

// NoSelection: `GtkNoSelection` is a `GtkSelectionModel` that does not allow
// selecting anything.
//
// This model is meant to be used as a simple wrapper around a `GListModel` when
// a `GtkSelectionModel` is required.
type NoSelection interface {
	gextras.Objector
}

// noSelection implements the NoSelection class.
type noSelection struct {
	gextras.Objector
}

// WrapNoSelection wraps a GObject to the right type. It is
// primarily used internally.
func WrapNoSelection(obj *externglib.Object) NoSelection {
	return noSelection{
		Objector: obj,
	}
}

func marshalNoSelection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNoSelection(obj), nil
}
