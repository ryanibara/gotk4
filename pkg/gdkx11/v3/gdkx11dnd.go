// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

// GTypeX11DragContext returns the GType for the type X11DragContext.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeX11DragContext() coreglib.Type {
	gtype := coreglib.Type(C.gdk_x11_drag_context_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalX11DragContext)
	return gtype
}

// X11DragContextOverrider contains methods that are overridable.
type X11DragContextOverrider interface {
}

type X11DragContext struct {
	_ [0]func() // equal guard
	gdk.DragContext
}

var (
	_ coreglib.Objector = (*X11DragContext)(nil)
)

func classInitX11DragContexter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapX11DragContext(obj *coreglib.Object) *X11DragContext {
	return &X11DragContext{
		DragContext: gdk.DragContext{
			Object: obj,
		},
	}
}

func marshalX11DragContext(p uintptr) (interface{}, error) {
	return wrapX11DragContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
