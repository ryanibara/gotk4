// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_Window_ConnectActivateDefault
func _gotk4_gtk4_Window_ConnectActivateDefault(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_Window_ConnectActivateFocus
func _gotk4_gtk4_Window_ConnectActivateFocus(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_Window_ConnectCloseRequest
func _gotk4_gtk4_Window_ConnectCloseRequest(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (ok bool))
	}

	ok := f()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_Window_ConnectEnableDebugging
func _gotk4_gtk4_Window_ConnectEnableDebugging(arg0 C.gpointer, arg1 C.gboolean, arg2 C.guintptr) (cret C.gboolean) {
	var f func(toggle bool) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(toggle bool) (ok bool))
	}

	var _toggle bool // out

	if arg1 != 0 {
		_toggle = true
	}

	ok := f(_toggle)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_Window_ConnectKeysChanged
func _gotk4_gtk4_Window_ConnectKeysChanged(arg0 C.gpointer, arg1 C.guintptr) {
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
