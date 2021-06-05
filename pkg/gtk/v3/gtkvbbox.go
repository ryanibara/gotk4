// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.gtk_vbutton_box_get_type()), F: marshalVButtonBox},
	})
}

type VButtonBox interface {
	ButtonBox
	Buildable
	Orientable
}

// vButtonBox implements the VButtonBox interface.
type vButtonBox struct {
	ButtonBox
	Buildable
	Orientable
}

var _ VButtonBox = (*vButtonBox)(nil)

// WrapVButtonBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapVButtonBox(obj *externglib.Object) VButtonBox {
	return VButtonBox{
		ButtonBox:  WrapButtonBox(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalVButtonBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVButtonBox(obj), nil
}

// NewVButtonBox constructs a class VButtonBox.
func NewVButtonBox() VButtonBox {
	var cret C.GtkVButtonBox
	var ret1 VButtonBox

	cret = C.gtk_vbutton_box_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(VButtonBox)

	return ret1
}
