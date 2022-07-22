// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeX11AppLaunchContext = coreglib.Type(C.gdk_x11_app_launch_context_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeX11AppLaunchContext, F: marshalX11AppLaunchContext},
	})
}

// X11AppLaunchContextOverrider contains methods that are overridable.
type X11AppLaunchContextOverrider interface {
}

type X11AppLaunchContext struct {
	_ [0]func() // equal guard
	gdk.AppLaunchContext
}

var (
	_ coreglib.Objector = (*X11AppLaunchContext)(nil)
)

func classInitX11AppLaunchContexter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

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
