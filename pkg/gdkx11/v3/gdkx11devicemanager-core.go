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

// GTypeX11DeviceManagerCore returns the GType for the type X11DeviceManagerCore.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeX11DeviceManagerCore() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("GdkX11", "X11DeviceManagerCore").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalX11DeviceManagerCore)
	return gtype
}

type X11DeviceManagerCore struct {
	_ [0]func() // equal guard
	gdk.DeviceManager
}

var (
	_ gdk.DeviceManagerer = (*X11DeviceManagerCore)(nil)
)

func wrapX11DeviceManagerCore(obj *coreglib.Object) *X11DeviceManagerCore {
	return &X11DeviceManagerCore{
		DeviceManager: gdk.DeviceManager{
			Object: obj,
		},
	}
}

func marshalX11DeviceManagerCore(p uintptr) (interface{}, error) {
	return wrapX11DeviceManagerCore(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
