// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_Filter_ConnectChanged
func _gotk4_gtk4_Filter_ConnectChanged(arg0 C.gpointer, arg1 C.GtkFilterChange, arg2 C.guintptr) {
	var f func(change FilterChange)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(change FilterChange))
	}

	var _change FilterChange // out

	_change = FilterChange(arg1)

	f(_change)
}