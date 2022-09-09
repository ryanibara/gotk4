// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

//export _gotk4_gdk3_Monitor_ConnectInvalidate
func _gotk4_gdk3_Monitor_ConnectInvalidate(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gdk3_WindowClass_create_surface
func _gotk4_gdk3_WindowClass_create_surface(arg0 *C.GdkWindow, arg1 C.gint, arg2 C.gint) (cret *C.cairo_surface_t) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[WindowOverrides](instance0)
	if overrides.CreateSurface == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected WindowOverrides.CreateSurface, got none")
	}

	var _width int  // out
	var _height int // out

	_width = int(arg1)
	_height = int(arg2)

	surface := overrides.CreateSurface(_width, _height)

	var _ *cairo.Surface

	cret = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))

	return cret
}

//export _gotk4_gdk3_WindowClass_from_embedder
func _gotk4_gdk3_WindowClass_from_embedder(arg0 *C.GdkWindow, arg1 C.gdouble, arg2 C.gdouble, arg3 *C.gdouble, arg4 *C.gdouble) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[WindowOverrides](instance0)
	if overrides.FromEmbedder == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected WindowOverrides.FromEmbedder, got none")
	}

	var _embedderX float64   // out
	var _embedderY float64   // out
	var _offscreenX *float64 // out
	var _offscreenY *float64 // out

	_embedderX = float64(arg1)
	_embedderY = float64(arg2)
	_offscreenX = (*float64)(unsafe.Pointer(arg3))
	_offscreenY = (*float64)(unsafe.Pointer(arg4))

	overrides.FromEmbedder(_embedderX, _embedderY, _offscreenX, _offscreenY)
}

//export _gotk4_gdk3_WindowClass_to_embedder
func _gotk4_gdk3_WindowClass_to_embedder(arg0 *C.GdkWindow, arg1 C.gdouble, arg2 C.gdouble, arg3 *C.gdouble, arg4 *C.gdouble) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[WindowOverrides](instance0)
	if overrides.ToEmbedder == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected WindowOverrides.ToEmbedder, got none")
	}

	var _offscreenX float64 // out
	var _offscreenY float64 // out
	var _embedderX *float64 // out
	var _embedderY *float64 // out

	_offscreenX = float64(arg1)
	_offscreenY = float64(arg2)
	_embedderX = (*float64)(unsafe.Pointer(arg3))
	_embedderY = (*float64)(unsafe.Pointer(arg4))

	overrides.ToEmbedder(_offscreenX, _offscreenY, _embedderX, _embedderY)
}
