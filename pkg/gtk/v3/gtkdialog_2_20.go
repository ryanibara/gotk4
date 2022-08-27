// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// WidgetForResponse gets the widget button that uses the given response ID in
// the action area of a dialog.
//
// The function takes the following parameters:
//
//    - responseId: response ID used by the dialog widget.
//
// The function returns the following values:
//
//    - widget (optional) button that uses the given response_id, or NULL.
//
func (dialog *Dialog) WidgetForResponse(responseId int) Widgetter {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.gint       // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(coreglib.InternObject(dialog).Native()))
	_arg1 = C.gint(responseId)

	_cret = C.gtk_dialog_get_widget_for_response(_arg0, _arg1)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(responseId)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}
