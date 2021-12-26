// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <gdk/wayland/gdkwayland.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_wayland_popup_get_type()), F: marshalWaylandPopupper},
		{T: externglib.Type(C.gdk_wayland_surface_get_type()), F: marshalWaylandSurfacer},
		{T: externglib.Type(C.gdk_wayland_toplevel_get_type()), F: marshalWaylandTopleveller},
	})
}

// WaylandPopup: wayland implementation of GdkPopup.
type WaylandPopup struct {
	_ [0]func() // equal guard
	WaylandSurface

	*externglib.Object
	gdk.Popup
	gdk.Surface
}

var (
	_ externglib.Objector = (*WaylandPopup)(nil)
	_ gdk.Surfacer        = (*WaylandPopup)(nil)
)

func wrapWaylandPopup(obj *externglib.Object) *WaylandPopup {
	return &WaylandPopup{
		WaylandSurface: WaylandSurface{
			Surface: gdk.Surface{
				Object: obj,
			},
		},
		Object: obj,
		Popup: gdk.Popup{
			Surface: gdk.Surface{
				Object: obj,
			},
		},
		Surface: gdk.Surface{
			Object: obj,
		},
	}
}

func marshalWaylandPopupper(p uintptr) (interface{}, error) {
	return wrapWaylandPopup(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// WaylandSurface: wayland implementation of GdkSurface.
//
// Beyond the gdk.Surface API, the Wayland implementation offers access to the
// Wayland wl_surface object with gdkwayland.WaylandSurface.GetWlSurface().
type WaylandSurface struct {
	_ [0]func() // equal guard
	gdk.Surface
}

var (
	_ gdk.Surfacer = (*WaylandSurface)(nil)
)

func wrapWaylandSurface(obj *externglib.Object) *WaylandSurface {
	return &WaylandSurface{
		Surface: gdk.Surface{
			Object: obj,
		},
	}
}

func marshalWaylandSurfacer(p uintptr) (interface{}, error) {
	return wrapWaylandSurface(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// WaylandToplevel: wayland implementation of GdkToplevel.
//
// Beyond the gdk.Toplevel API, the Wayland implementation has API to set up
// cross-process parent-child relationships between surfaces with
// gdkwayland.WaylandToplevel.ExportHandle() and
// gdkwayland.WaylandToplevel.SetTransientForExported().
type WaylandToplevel struct {
	_ [0]func() // equal guard
	WaylandSurface

	*externglib.Object
	gdk.Surface
	gdk.Toplevel
}

var (
	_ externglib.Objector = (*WaylandToplevel)(nil)
	_ gdk.Surfacer        = (*WaylandToplevel)(nil)
)

func wrapWaylandToplevel(obj *externglib.Object) *WaylandToplevel {
	return &WaylandToplevel{
		WaylandSurface: WaylandSurface{
			Surface: gdk.Surface{
				Object: obj,
			},
		},
		Object: obj,
		Surface: gdk.Surface{
			Object: obj,
		},
		Toplevel: gdk.Toplevel{
			Surface: gdk.Surface{
				Object: obj,
			},
		},
	}
}

func marshalWaylandTopleveller(p uintptr) (interface{}, error) {
	return wrapWaylandToplevel(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
