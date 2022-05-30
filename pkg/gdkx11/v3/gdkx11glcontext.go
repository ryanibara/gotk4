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
import "C"

// glib.Type values for gdkx11glcontext.go.
var GTypeX11GLContext = coreglib.Type(C.gdk_x11_gl_context_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeX11GLContext, F: marshalX11GLContext},
	})
}

// X11GLContextOverrider contains methods that are overridable.
type X11GLContextOverrider interface {
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
