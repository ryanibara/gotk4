// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_monitor_get_type()), F: marshalX11Monitorrer},
	})
}

type X11Monitor struct {
	gdk.Monitor
}

func wrapX11Monitor(obj *externglib.Object) *X11Monitor {
	return &X11Monitor{
		Monitor: gdk.Monitor{
			Object: obj,
		},
	}
}

func marshalX11Monitorrer(p uintptr) (interface{}, error) {
	return wrapX11Monitor(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Workarea retrieves the size and position of the “work area” on a monitor
// within the display coordinate space. The returned geometry is in ”application
// pixels”, not in ”device pixels” (see gdk_monitor_get_scale_factor()).
func (monitor *X11Monitor) Workarea() *gdk.Rectangle {
	var _arg0 *C.GdkMonitor  // out
	var _arg1 C.GdkRectangle // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	C.gdk_x11_monitor_get_workarea(_arg0, &_arg1)
	runtime.KeepAlive(monitor)

	var _workarea *gdk.Rectangle // out

	_workarea = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _workarea
}
