// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
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

type ToplevelAccessible struct {
	_ [0]func() // equal guard
	atk.ObjectClass
}

var (
	_ externglib.Objector = (*ToplevelAccessible)(nil)
)

func wrapToplevelAccessible(obj *externglib.Object) *ToplevelAccessible {
	return &ToplevelAccessible{
		ObjectClass: atk.ObjectClass{
			Object: obj,
		},
	}
}

func marshalToplevelAccessibler(p uintptr) (interface{}, error) {
	return wrapToplevelAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function returns the following values:
//
//    - list: list of children.
//
func (accessible *ToplevelAccessible) Children() []Window {
	var _arg0 *C.GtkToplevelAccessible // out
	var _cret *C.GList                 // in

	_arg0 = (*C.GtkToplevelAccessible)(unsafe.Pointer(accessible.Native()))

	_cret = C.gtk_toplevel_accessible_get_children(_arg0)
	runtime.KeepAlive(accessible)

	var _list []Window // out

	_list = make([]Window, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.GtkWindow)(v)
		var dst Window // out
		dst = *wrapWindow(externglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}
