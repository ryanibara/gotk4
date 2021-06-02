// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_monitor_get_type()), F: marshalX11Monitor},
	})
}

type X11Monitor interface {
	gdk.Monitor
}

// x11Monitor implements the X11Monitor interface.
type x11Monitor struct {
	gdk.Monitor
}

var _ X11Monitor = (*x11Monitor)(nil)

// WrapX11Monitor wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Monitor(obj *externglib.Object) X11Monitor {
	return X11Monitor{
		gdk.Monitor: gdk.WrapMonitor(obj),
	}
}

func marshalX11Monitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Monitor(obj), nil
}
