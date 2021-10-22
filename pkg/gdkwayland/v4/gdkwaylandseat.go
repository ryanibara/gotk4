// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4-wayland gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/wayland/gdkwayland.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_wayland_seat_get_type()), F: marshalWaylandSeater},
	})
}

// WaylandSeat: wayland implementation of GdkSeat.
//
// Beyond the regular gdk.Seat API, the Wayland implementation provides access
// to the Wayland wl_seat object with gdkwayland.WaylandSeat.GetWlSeat().
type WaylandSeat struct {
	gdk.Seat
}

func wrapWaylandSeat(obj *externglib.Object) *WaylandSeat {
	return &WaylandSeat{
		Seat: gdk.Seat{
			Object: obj,
		},
	}
}

func marshalWaylandSeater(p uintptr) (interface{}, error) {
	return wrapWaylandSeat(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
