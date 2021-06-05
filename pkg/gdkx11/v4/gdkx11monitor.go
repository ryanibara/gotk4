// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/x11/gdkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_monitor_get_type()), F: marshalX11Monitor},
	})
}

type X11Monitor interface {
	gdk.Monitor

	// Workarea retrieves the size and position of the “work area” on a monitor
	// within the display coordinate space. The returned geometry is in
	// ”application pixels”, not in ”device pixels” (see
	// gdk_monitor_get_scale_factor()).
	Workarea() gdk.Rectangle
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

// Workarea retrieves the size and position of the “work area” on a monitor
// within the display coordinate space. The returned geometry is in
// ”application pixels”, not in ”device pixels” (see
// gdk_monitor_get_scale_factor()).
func (m x11Monitor) Workarea() gdk.Rectangle {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var arg1 C.GdkRectangle
	var ret1 *gdk.Rectangle

	C.gdk_x11_monitor_get_workarea(arg0, &arg1)

	ret1 = gdk.WrapRectangle(unsafe.Pointer(arg1))

	return ret1
}
