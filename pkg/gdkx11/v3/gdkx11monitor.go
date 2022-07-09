// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeX11Monitor returns the GType for the type X11Monitor.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeX11Monitor() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("GdkX11", "X11Monitor").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalX11Monitor)
	return gtype
}

type X11Monitor struct {
	_ [0]func() // equal guard
	gdk.Monitor
}

var (
	_ coreglib.Objector = (*X11Monitor)(nil)
)

func wrapX11Monitor(obj *coreglib.Object) *X11Monitor {
	return &X11Monitor{
		Monitor: gdk.Monitor{
			Object: obj,
		},
	}
}

func marshalX11Monitor(p uintptr) (interface{}, error) {
	return wrapX11Monitor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
