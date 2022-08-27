// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_TreeCellDataFunc
func _gotk4_gtk4_TreeCellDataFunc(arg1 *C.GtkTreeViewColumn, arg2 *C.GtkCellRenderer, arg3 *C.GtkTreeModel, arg4 *C.GtkTreeIter, arg5 C.gpointer) {
	var fn TreeCellDataFunc
	{
		v := gbox.Get(uintptr(arg5))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TreeCellDataFunc)
	}

	var _treeColumn *TreeViewColumn // out
	var _cell CellRendererer        // out
	var _treeModel TreeModeller     // out
	var _iter *TreeIter             // out

	_treeColumn = wrapTreeViewColumn(coreglib.Take(unsafe.Pointer(arg1)))
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gtk.CellRendererer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(CellRendererer)
			return ok
		})
		rv, ok := casted.(CellRendererer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.CellRendererer")
		}
		_cell = rv
	}
	{
		objptr := unsafe.Pointer(arg3)
		if objptr == nil {
			panic("object of type gtk.TreeModeller is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(TreeModeller)
			return ok
		})
		rv, ok := casted.(TreeModeller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.TreeModeller")
		}
		_treeModel = rv
	}
	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg4)))

	fn(_treeColumn, _cell, _treeModel, _iter)
}

//export _gotk4_gtk4_TreeViewColumn_ConnectClicked
func _gotk4_gtk4_TreeViewColumn_ConnectClicked(arg0 C.gpointer, arg1 C.guintptr) {
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
