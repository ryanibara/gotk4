// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// EventWindow returns the button’s event window if it is realized, NULL
// otherwise. This function should be rarely needed.
//
// The function returns the following values:
//
//    - window button’s event window.
//
func (button *Button) EventWindow() gdk.Windower {
	var _arg0 *C.GtkButton // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GtkButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_button_get_event_window(_arg0)
	runtime.KeepAlive(button)

	var _window gdk.Windower // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.Windower)
			return ok
		})
		rv, ok := casted.(gdk.Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_window = rv
	}

	return _window
}
