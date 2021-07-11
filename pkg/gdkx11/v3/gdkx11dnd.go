// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_drag_context_get_type()), F: marshalX11DragContexter},
	})
}

// X11DragContexter describes X11DragContext's methods.
type X11DragContexter interface {
	privateX11DragContext()
}

type X11DragContext struct {
	gdk.DragContext
}

var (
	_ X11DragContexter = (*X11DragContext)(nil)
	_ gextras.Nativer  = (*X11DragContext)(nil)
)

func wrapX11DragContext(obj *externglib.Object) X11DragContexter {
	return &X11DragContext{
		DragContext: gdk.DragContext{
			Object: obj,
		},
	}
}

func marshalX11DragContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11DragContext(obj), nil
}

func (*X11DragContext) privateX11DragContext() {}
