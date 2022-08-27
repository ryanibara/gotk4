// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_GLArea_ConnectCreateContext
func _gotk4_gtk4_GLArea_ConnectCreateContext(arg0 C.gpointer, arg1 C.guintptr) (cret *C.GdkGLContext) {
	var f func() (glContext gdk.GLContexter)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (glContext gdk.GLContexter))
	}

	glContext := f()

	cret = (*C.GdkGLContext)(unsafe.Pointer(coreglib.InternObject(glContext).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(glContext).Native()))

	return cret
}

//export _gotk4_gtk4_GLArea_ConnectRender
func _gotk4_gtk4_GLArea_ConnectRender(arg0 C.gpointer, arg1 *C.GdkGLContext, arg2 C.guintptr) (cret C.gboolean) {
	var f func(context gdk.GLContexter) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(context gdk.GLContexter) (ok bool))
	}

	var _context gdk.GLContexter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.GLContexter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.GLContexter)
			return ok
		})
		rv, ok := casted.(gdk.GLContexter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.GLContexter")
		}
		_context = rv
	}

	ok := f(_context)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_GLArea_ConnectResize
func _gotk4_gtk4_GLArea_ConnectResize(arg0 C.gpointer, arg1 C.gint, arg2 C.gint, arg3 C.guintptr) {
	var f func(width, height int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(width, height int))
	}

	var _width int  // out
	var _height int // out

	_width = int(arg1)
	_height = int(arg2)

	f(_width, _height)
}
