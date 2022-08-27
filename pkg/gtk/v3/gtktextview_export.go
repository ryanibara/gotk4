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

//export _gotk4_gtk3_TextView_ConnectBackspace
func _gotk4_gtk3_TextView_ConnectBackspace(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk3_TextView_ConnectCopyClipboard
func _gotk4_gtk3_TextView_ConnectCopyClipboard(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk3_TextView_ConnectCutClipboard
func _gotk4_gtk3_TextView_ConnectCutClipboard(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk3_TextView_ConnectDeleteFromCursor
func _gotk4_gtk3_TextView_ConnectDeleteFromCursor(arg0 C.gpointer, arg1 C.GtkDeleteType, arg2 C.gint, arg3 C.guintptr) {
	var f func(typ DeleteType, count int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(typ DeleteType, count int))
	}

	var _typ DeleteType // out
	var _count int      // out

	_typ = DeleteType(arg1)
	_count = int(arg2)

	f(_typ, _count)
}

//export _gotk4_gtk3_TextView_ConnectExtendSelection
func _gotk4_gtk3_TextView_ConnectExtendSelection(arg0 C.gpointer, arg1 C.GtkTextExtendSelection, arg2 *C.GtkTextIter, arg3 *C.GtkTextIter, arg4 *C.GtkTextIter, arg5 C.guintptr) (cret C.gboolean) {
	var f func(granularity TextExtendSelection, location, start, end *TextIter) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg5))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(granularity TextExtendSelection, location, start, end *TextIter) (ok bool))
	}

	var _granularity TextExtendSelection // out
	var _location *TextIter              // out
	var _start *TextIter                 // out
	var _end *TextIter                   // out

	_granularity = TextExtendSelection(arg1)
	_location = (*TextIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_start = (*TextIter)(gextras.NewStructNative(unsafe.Pointer(arg3)))
	_end = (*TextIter)(gextras.NewStructNative(unsafe.Pointer(arg4)))

	ok := f(_granularity, _location, _start, _end)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_TextView_ConnectInsertAtCursor
func _gotk4_gtk3_TextView_ConnectInsertAtCursor(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(str string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(str string))
	}

	var _str string // out

	_str = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_str)
}

//export _gotk4_gtk3_TextView_ConnectInsertEmoji
func _gotk4_gtk3_TextView_ConnectInsertEmoji(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk3_TextView_ConnectMoveCursor
func _gotk4_gtk3_TextView_ConnectMoveCursor(arg0 C.gpointer, arg1 C.GtkMovementStep, arg2 C.gint, arg3 C.gboolean, arg4 C.guintptr) {
	var f func(step MovementStep, count int, extendSelection bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(step MovementStep, count int, extendSelection bool))
	}

	var _step MovementStep    // out
	var _count int            // out
	var _extendSelection bool // out

	_step = MovementStep(arg1)
	_count = int(arg2)
	if arg3 != 0 {
		_extendSelection = true
	}

	f(_step, _count, _extendSelection)
}

//export _gotk4_gtk3_TextView_ConnectMoveViewport
func _gotk4_gtk3_TextView_ConnectMoveViewport(arg0 C.gpointer, arg1 C.GtkScrollStep, arg2 C.gint, arg3 C.guintptr) {
	var f func(step ScrollStep, count int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(step ScrollStep, count int))
	}

	var _step ScrollStep // out
	var _count int       // out

	_step = ScrollStep(arg1)
	_count = int(arg2)

	f(_step, _count)
}

//export _gotk4_gtk3_TextView_ConnectPasteClipboard
func _gotk4_gtk3_TextView_ConnectPasteClipboard(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk3_TextView_ConnectPopulatePopup
func _gotk4_gtk3_TextView_ConnectPopulatePopup(arg0 C.gpointer, arg1 *C.GtkWidget, arg2 C.guintptr) {
	var f func(popup Widgetter)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(popup Widgetter))
	}

	var _popup Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
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
		_popup = rv
	}

	f(_popup)
}

//export _gotk4_gtk3_TextView_ConnectPreeditChanged
func _gotk4_gtk3_TextView_ConnectPreeditChanged(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(preedit string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(preedit string))
	}

	var _preedit string // out

	_preedit = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_preedit)
}

//export _gotk4_gtk3_TextView_ConnectSelectAll
func _gotk4_gtk3_TextView_ConnectSelectAll(arg0 C.gpointer, arg1 C.gboolean, arg2 C.guintptr) {
	var f func(sel bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(sel bool))
	}

	var _sel bool // out

	if arg1 != 0 {
		_sel = true
	}

	f(_sel)
}

//export _gotk4_gtk3_TextView_ConnectSetAnchor
func _gotk4_gtk3_TextView_ConnectSetAnchor(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk3_TextView_ConnectToggleCursorVisible
func _gotk4_gtk3_TextView_ConnectToggleCursorVisible(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk3_TextView_ConnectToggleOverwrite
func _gotk4_gtk3_TextView_ConnectToggleOverwrite(arg0 C.gpointer, arg1 C.guintptr) {
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
