// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/atk"
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
		{T: externglib.Type(C.gtk_cell_accessible_get_type()), F: marshalCellAccessible},
	})
}

type CellAccessible interface {
	Accessible
	atk.Action
}

// cellAccessible implements the CellAccessible class.
type cellAccessible struct {
	Accessible
}

// WrapCellAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellAccessible(obj *externglib.Object) CellAccessible {
	return cellAccessible{
		Accessible: WrapAccessible(obj),
	}
}

func marshalCellAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellAccessible(obj), nil
}

func (a cellAccessible) DoAction(i int) bool {
	return atk.WrapAction(gextras.InternObject(a)).DoAction(i)
}

func (a cellAccessible) Description(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Description(i)
}

func (a cellAccessible) Keybinding(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Keybinding(i)
}

func (a cellAccessible) LocalizedName(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).LocalizedName(i)
}

func (a cellAccessible) NActions() int {
	return atk.WrapAction(gextras.InternObject(a)).NActions()
}

func (a cellAccessible) Name(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Name(i)
}

func (a cellAccessible) SetDescription(i int, desc string) bool {
	return atk.WrapAction(gextras.InternObject(a)).SetDescription(i, desc)
}
