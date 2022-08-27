// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

//export _gotk4_gtk3_ListBox_ConnectRowActivated
func _gotk4_gtk3_ListBox_ConnectRowActivated(arg0 C.gpointer, arg1 *C.GtkListBoxRow, arg2 C.guintptr) {
	var f func(row *ListBoxRow)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(row *ListBoxRow))
	}

	var _row *ListBoxRow // out

	_row = wrapListBoxRow(coreglib.Take(unsafe.Pointer(arg1)))

	f(_row)
}

//export _gotk4_gtk3_ListBox_ConnectRowSelected
func _gotk4_gtk3_ListBox_ConnectRowSelected(arg0 C.gpointer, arg1 *C.GtkListBoxRow, arg2 C.guintptr) {
	var f func(row *ListBoxRow)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(row *ListBoxRow))
	}

	var _row *ListBoxRow // out

	if arg1 != nil {
		_row = wrapListBoxRow(coreglib.Take(unsafe.Pointer(arg1)))
	}

	f(_row)
}

//export _gotk4_gtk3_ListBox_ConnectSelectAll
func _gotk4_gtk3_ListBox_ConnectSelectAll(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk3_ListBox_ConnectSelectedRowsChanged
func _gotk4_gtk3_ListBox_ConnectSelectedRowsChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk3_ListBox_ConnectUnselectAll
func _gotk4_gtk3_ListBox_ConnectUnselectAll(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk3_ListBoxRow_ConnectActivate
func _gotk4_gtk3_ListBoxRow_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
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
