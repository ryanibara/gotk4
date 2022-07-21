// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

// GTypeDragSurface returns the GType for the type DragSurface.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDragSurface() coreglib.Type {
	gtype := coreglib.Type(C.gdk_drag_surface_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalDragSurface)
	return gtype
}

// DragSurfaceOverrider contains methods that are overridable.
type DragSurfaceOverrider interface {
}

// DragSurface is an interface for surfaces used during DND.
//
// DragSurface wraps an interface. This means the user can get the
// underlying type by calling Cast().
type DragSurface struct {
	_ [0]func() // equal guard
	Surface
}

var (
	_ Surfacer = (*DragSurface)(nil)
)

// DragSurfacer describes DragSurface's interface methods.
type DragSurfacer interface {
	coreglib.Objector

	// Present drag_surface.
	Present(width, height int32) bool
}

var _ DragSurfacer = (*DragSurface)(nil)

func ifaceInitDragSurfacer(gifacePtr, data C.gpointer) {
}

func wrapDragSurface(obj *coreglib.Object) *DragSurface {
	return &DragSurface{
		Surface: Surface{
			Object: obj,
		},
	}
}

func marshalDragSurface(p uintptr) (interface{}, error) {
	return wrapDragSurface(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Present drag_surface.
//
// The function takes the following parameters:
//
//    - width: unconstrained drag_surface width to layout.
//    - height: unconstrained drag_surface height to layout.
//
// The function returns the following values:
//
//    - ok: FALSE if it failed to be presented, otherwise TRUE.
//
func (dragSurface *DragSurface) Present(width, height int32) bool {
	var _arg0 *C.GdkDragSurface // out
	var _arg1 C.int             // out
	var _arg2 C.int             // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkDragSurface)(unsafe.Pointer(coreglib.InternObject(dragSurface).Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	_cret = C.gdk_drag_surface_present(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dragSurface)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
