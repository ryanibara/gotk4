// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

//export _gotk4_gdk4_Drag_ConnectCancel
func _gotk4_gdk4_Drag_ConnectCancel(arg0 C.gpointer, arg1 C.GdkDragCancelReason, arg2 C.guintptr) {
	var f func(reason DragCancelReason)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(reason DragCancelReason))
	}

	var _reason DragCancelReason // out

	_reason = DragCancelReason(arg1)

	f(_reason)
}

//export _gotk4_gdk4_Drag_ConnectDNDFinished
func _gotk4_gdk4_Drag_ConnectDNDFinished(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gdk4_Drag_ConnectDropPerformed
func _gotk4_gdk4_Drag_ConnectDropPerformed(arg0 C.gpointer, arg1 C.guintptr) {
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