// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <gdk/wayland/gdkwayland.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeWaylandMonitor = coreglib.Type(C.gdk_wayland_monitor_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWaylandMonitor, F: marshalWaylandMonitor},
	})
}

// WaylandMonitorOverrider contains methods that are overridable.
type WaylandMonitorOverrider interface {
}

// WaylandMonitor: wayland implementation of GdkMonitor.
//
// Beyond the gdk.Monitor API, the Wayland implementation offers access to the
// Wayland wl_output object with gdkwayland.WaylandMonitor.GetWlOutput().
type WaylandMonitor struct {
	_ [0]func() // equal guard
	gdk.Monitor
}

var (
	_ coreglib.Objector = (*WaylandMonitor)(nil)
)

func initClassWaylandMonitor(gclass unsafe.Pointer, goval any) {
}

func wrapWaylandMonitor(obj *coreglib.Object) *WaylandMonitor {
	return &WaylandMonitor{
		Monitor: gdk.Monitor{
			Object: obj,
		},
	}
}

func marshalWaylandMonitor(p uintptr) (interface{}, error) {
	return wrapWaylandMonitor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
