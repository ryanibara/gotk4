// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gdkx11glcontext.go.
var GTypeX11GLContext = coreglib.Type(C.gdk_x11_gl_context_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeX11GLContext, F: marshalX11GLContext},
	})
}

// X11DisplayGetGLXVersion retrieves the version of the GLX implementation.
//
// The function takes the following parameters:
//
//    - display: Display.
//
// The function returns the following values:
//
//    - major: return location for the GLX major version.
//    - minor: return location for the GLX minor version.
//    - ok: TRUE if GLX is available.
//
func X11DisplayGetGLXVersion(display *gdk.Display) (major, minor int32, ok bool) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_gret := girepository.MustFind("GdkX11", "get_glx_version").Invoke(_args[:], _outs[:])
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
			Object: obj,
		},
	}
}

func marshalX11GLContext(p uintptr) (interface{}, error) {
	return wrapX11GLContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
