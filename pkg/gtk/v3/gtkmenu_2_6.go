// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// MenuGetForAttachWidget returns a list of the menus which are attached to this
// widget. This list is owned by GTK+ and must not be modified.
//
// The function takes the following parameters:
//
//    - widget: Widget.
//
// The function returns the following values:
//
//    - list: list of menus attached to his widget.
//
func MenuGetForAttachWidget(widget Widgetter) []Widgetter {
	var _arg1 *C.GtkWidget // out
	var _cret *C.GList     // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_cret = C.gtk_menu_get_for_attach_widget(_arg1)
	runtime.KeepAlive(widget)

	var _list []Widgetter // out

	_list = make([]Widgetter, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.GtkWidget)(v)
		var dst Widgetter // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gtk.Widgetter is nil")
			}

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}
