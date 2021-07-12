// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_toplevel_accessible_get_type()), F: marshalToplevelAccessibler},
	})
}

// ToplevelAccessibler describes ToplevelAccessible's methods.
type ToplevelAccessibler interface {
	privateToplevelAccessible()
}

type ToplevelAccessible struct {
	atk.ObjectClass
}

var (
	_ ToplevelAccessibler = (*ToplevelAccessible)(nil)
	_ gextras.Nativer     = (*ToplevelAccessible)(nil)
)

func wrapToplevelAccessible(obj *externglib.Object) *ToplevelAccessible {
	return &ToplevelAccessible{
		ObjectClass: atk.ObjectClass{
			Object: obj,
		},
	}
}

func marshalToplevelAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapToplevelAccessible(obj), nil
}

func (*ToplevelAccessible) privateToplevelAccessible() {}
