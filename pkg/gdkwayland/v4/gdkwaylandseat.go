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
	GTypeWaylandSeat = coreglib.Type(C.gdk_wayland_seat_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWaylandSeat, F: marshalWaylandSeat},
	})
}

// WaylandSeatOverrider contains methods that are overridable.
type WaylandSeatOverrider interface {
}

// WaylandSeat: wayland implementation of GdkSeat.
//
// Beyond the regular gdk.Seat API, the Wayland implementation provides access
// to the Wayland wl_seat object with gdkwayland.WaylandSeat.GetWlSeat().
type WaylandSeat struct {
	_ [0]func() // equal guard
	gdk.Seat
}

var (
	_ gdk.Seater = (*WaylandSeat)(nil)
)

func initClassWaylandSeat(gclass unsafe.Pointer, goval any) {
}

func wrapWaylandSeat(obj *coreglib.Object) *WaylandSeat {
	return &WaylandSeat{
		Seat: gdk.Seat{
			Object: obj,
		},
	}
}

func marshalWaylandSeat(p uintptr) (interface{}, error) {
	return wrapWaylandSeat(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
