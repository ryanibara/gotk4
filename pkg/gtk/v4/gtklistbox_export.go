// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_ListBoxCreateWidgetFunc
func _gotk4_gtk4_ListBoxCreateWidgetFunc(arg1 C.gpointer, arg2 C.gpointer) (cret *C.GtkWidget) {
	var fn ListBoxCreateWidgetFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ListBoxCreateWidgetFunc)
	}

	var _item *coreglib.Object // out

	_item = coreglib.Take(unsafe.Pointer(arg1))

	widget := fn(_item)

	cret = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(widget).Native()))

	return cret
}

//export _gotk4_gtk4_ListBoxFilterFunc
func _gotk4_gtk4_ListBoxFilterFunc(arg1 *C.GtkListBoxRow, arg2 C.gpointer) (cret C.gboolean) {
	var fn ListBoxFilterFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ListBoxFilterFunc)
	}

	var _row *ListBoxRow // out

	_row = wrapListBoxRow(coreglib.Take(unsafe.Pointer(arg1)))

	ok := fn(_row)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_ListBoxForEachFunc
func _gotk4_gtk4_ListBoxForEachFunc(arg1 *C.GtkListBox, arg2 *C.GtkListBoxRow, arg3 C.gpointer) {
	var fn ListBoxForEachFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ListBoxForEachFunc)
	}

	var _box *ListBox    // out
	var _row *ListBoxRow // out

	_box = wrapListBox(coreglib.Take(unsafe.Pointer(arg1)))
	_row = wrapListBoxRow(coreglib.Take(unsafe.Pointer(arg2)))

	fn(_box, _row)
}

//export _gotk4_gtk4_ListBoxSortFunc
func _gotk4_gtk4_ListBoxSortFunc(arg1 *C.GtkListBoxRow, arg2 *C.GtkListBoxRow, arg3 C.gpointer) (cret C.int) {
	var fn ListBoxSortFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ListBoxSortFunc)
	}

	var _row1 *ListBoxRow // out
	var _row2 *ListBoxRow // out

	_row1 = wrapListBoxRow(coreglib.Take(unsafe.Pointer(arg1)))
	_row2 = wrapListBoxRow(coreglib.Take(unsafe.Pointer(arg2)))

	gint := fn(_row1, _row2)

	cret = C.int(gint)

	return cret
}

//export _gotk4_gtk4_ListBoxUpdateHeaderFunc
func _gotk4_gtk4_ListBoxUpdateHeaderFunc(arg1 *C.GtkListBoxRow, arg2 *C.GtkListBoxRow, arg3 C.gpointer) {
	var fn ListBoxUpdateHeaderFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ListBoxUpdateHeaderFunc)
	}

	var _row *ListBoxRow    // out
	var _before *ListBoxRow // out

	_row = wrapListBoxRow(coreglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_before = wrapListBoxRow(coreglib.Take(unsafe.Pointer(arg2)))
	}

	fn(_row, _before)
}

//export _gotk4_gtk4_ListBox_ConnectRowActivated
func _gotk4_gtk4_ListBox_ConnectRowActivated(arg0 C.gpointer, arg1 *C.GtkListBoxRow, arg2 C.guintptr) {
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

//export _gotk4_gtk4_ListBox_ConnectRowSelected
func _gotk4_gtk4_ListBox_ConnectRowSelected(arg0 C.gpointer, arg1 *C.GtkListBoxRow, arg2 C.guintptr) {
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

//export _gotk4_gtk4_ListBox_ConnectSelectAll
func _gotk4_gtk4_ListBox_ConnectSelectAll(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_ListBox_ConnectSelectedRowsChanged
func _gotk4_gtk4_ListBox_ConnectSelectedRowsChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_ListBox_ConnectUnselectAll
func _gotk4_gtk4_ListBox_ConnectUnselectAll(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_ListBoxRow_ConnectActivate
func _gotk4_gtk4_ListBoxRow_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
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
