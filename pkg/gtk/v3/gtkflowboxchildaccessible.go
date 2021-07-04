// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_flow_box_child_accessible_get_type()), F: marshalFlowBoxChildAccessible},
	})
}

type FlowBoxChildAccessible interface {
	ContainerAccessible
}

// flowBoxChildAccessible implements the FlowBoxChildAccessible class.
type flowBoxChildAccessible struct {
	ContainerAccessible
}

// WrapFlowBoxChildAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapFlowBoxChildAccessible(obj *externglib.Object) FlowBoxChildAccessible {
	return flowBoxChildAccessible{
		ContainerAccessible: WrapContainerAccessible(obj),
	}
}

func marshalFlowBoxChildAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFlowBoxChildAccessible(obj), nil
}
