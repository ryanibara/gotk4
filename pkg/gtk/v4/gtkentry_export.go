// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_Entry_ConnectActivate
func _gotk4_gtk4_Entry_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gtk4_Entry_ConnectIconPress
func _gotk4_gtk4_Entry_ConnectIconPress(arg0 C.gpointer, arg1 C.GtkEntryIconPosition, arg2 C.guintptr) {
	var f func(iconPos EntryIconPosition)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(iconPos EntryIconPosition))
	}

	var _iconPos EntryIconPosition // out

	_iconPos = EntryIconPosition(arg1)

	f(_iconPos)
}

//export _gotk4_gtk4_Entry_ConnectIconRelease
func _gotk4_gtk4_Entry_ConnectIconRelease(arg0 C.gpointer, arg1 C.GtkEntryIconPosition, arg2 C.guintptr) {
	var f func(iconPos EntryIconPosition)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(iconPos EntryIconPosition))
	}

	var _iconPos EntryIconPosition // out

	_iconPos = EntryIconPosition(arg1)

	f(_iconPos)
}