// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeX11AppLaunchContext returns the GType for the type X11AppLaunchContext.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeX11AppLaunchContext() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("GdkX11", "X11AppLaunchContext").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalX11AppLaunchContext)
	return gtype
}

type X11AppLaunchContext struct {
	_ [0]func() // equal guard
	gdk.AppLaunchContext
}

var (
	_ coreglib.Objector = (*X11AppLaunchContext)(nil)
)

func wrapX11AppLaunchContext(obj *coreglib.Object) *X11AppLaunchContext {
	return &X11AppLaunchContext{
		AppLaunchContext: gdk.AppLaunchContext{
			AppLaunchContext: gio.AppLaunchContext{
				Object: obj,
			},
		},
	}
}

func marshalX11AppLaunchContext(p uintptr) (interface{}, error) {
	return wrapX11AppLaunchContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
