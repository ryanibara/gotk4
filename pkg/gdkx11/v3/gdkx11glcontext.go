// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gdkx11glcontext.go.
var GTypeX11GLContext = externglib.Type(C.gdk_x11_gl_context_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
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
func X11DisplayGetGLXVersion(display *gdk.Display) (major int, minor int, ok bool) {
	var _arg1 *C.GdkDisplay // out
	var _arg2 C.gint        // in
	var _arg3 C.gint        // in
	var _cret C.gboolean    // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	_cret = C.gdk_x11_display_get_glx_version(_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(display)

	var _major int // out
	var _minor int // out
	var _ok bool   // out

	_major = int(_arg2)
	_minor = int(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _major, _minor, _ok
}

// X11GLContextOverrider contains methods that are overridable.
type X11GLContextOverrider interface {
	externglib.Objector
}

type X11GLContext struct {
	_ [0]func() // equal guard
	gdk.GLContext
}

var (
	_ gdk.GLContexter = (*X11GLContext)(nil)
)

func classInitX11GLContexter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapX11GLContext(obj *externglib.Object) *X11GLContext {
	return &X11GLContext{
		GLContext: gdk.GLContext{
			Object: obj,
		},
	}
}

func marshalX11GLContext(p uintptr) (interface{}, error) {
	return wrapX11GLContext(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
