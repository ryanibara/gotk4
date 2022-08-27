// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_CellRendererText_ConnectEdited
func _gotk4_gtk4_CellRendererText_ConnectEdited(arg0 C.gpointer, arg1 *C.gchar, arg2 *C.gchar, arg3 C.guintptr) {
	var f func(path, newText string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(path, newText string))
	}

	var _path string    // out
	var _newText string // out

	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_newText = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	f(_path, _newText)
}
