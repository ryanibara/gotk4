// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeX11GLContext returns the GType for the type X11GLContext.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeX11GLContext() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("GdkX11", "X11GLContext").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalX11GLContext)
	return gtype
}

// GLXVersion retrieves the version of the GLX implementation.
//
// The function returns the following values:
//
//    - major: return location for the GLX major version.
//    - minor: return location for the GLX minor version.
//    - ok: TRUE if GLX is available.
//
func (display *X11Display) GLXVersion() (major, minor int32, ok bool) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_gret := girepository.MustFind("GdkX11", "X11Display").InvokeMethod("get_glx_version", _args[:], _outs[:])
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(display)

	var _major int32 // out
	var _minor int32 // out
	var _ok bool     // out

	_major = *(*int32)(unsafe.Pointer(_outs[0]))
	_minor = *(*int32)(unsafe.Pointer(_outs[1]))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _major, _minor, _ok
}

type X11GLContext struct {
	_ [0]func() // equal guard
	gdk.GLContext
}

var (
	_ gdk.GLContexter = (*X11GLContext)(nil)
)

func wrapX11GLContext(obj *coreglib.Object) *X11GLContext {
	return &X11GLContext{
		GLContext: gdk.GLContext{
			DrawContext: gdk.DrawContext{
				Object: obj,
			},
		},
	}
}

func marshalX11GLContext(p uintptr) (interface{}, error) {
	return wrapX11GLContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
