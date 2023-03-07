// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

//export _gotk4_gio2_SocketSourceFunc
func _gotk4_gio2_SocketSourceFunc(arg1 *C.GSocket, arg2 C.GIOCondition, arg3 C.gpointer) (cret C.gboolean) {
	var fn SocketSourceFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(SocketSourceFunc)
	}

	var _socket *Socket             // out
	var _condition glib.IOCondition // out

	_socket = wrapSocket(coreglib.Take(unsafe.Pointer(arg1)))
	_condition = glib.IOCondition(arg2)

	ok := fn(_socket, _condition)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}