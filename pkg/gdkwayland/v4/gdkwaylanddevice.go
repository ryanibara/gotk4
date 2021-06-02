// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-wayland
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/wayland/gdkwayland.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_wayland_device_get_type()), F: marshalWaylandDevice},
	})
}

// WaylandDevice: the Wayland implementation of `GdkDevice`.
//
// Beyond the regular [class@Gdk.Device] API, the Wayland implementation
// provides access to Wayland objects such as the `wl_seat` with
// [method@GdkWayland.WaylandDevice.get_wl_seat], the `wl_keyboard` with
// [method@GdkWayland.WaylandDevice.get_wl_keyboard] and the `wl_pointer` with
// [method@GdkWayland.WaylandDevice.get_wl_pointer].
type WaylandDevice interface {
	gdk.Device

	// NodePath returns the `/dev/input/event*` path of this device.
	//
	// For `GdkDevice`s that possibly coalesce multiple hardware devices (eg.
	// mouse, keyboard, touch,...), this function will return nil.
	//
	// This is most notably implemented for devices of type GDK_SOURCE_PEN,
	// GDK_SOURCE_TABLET_PAD.
	NodePath() string
	// WlKeyboard returns the Wayland `wl_keyboard` of a `GdkDevice`.
	WlKeyboard() interface{}
	// WlPointer returns the Wayland `wl_pointer` of a `GdkDevice`.
	WlPointer() interface{}
	// WlSeat returns the Wayland `wl_seat` of a `GdkDevice`.
	WlSeat() interface{}
}

// waylandDevice implements the WaylandDevice interface.
type waylandDevice struct {
	gdk.Device
}

var _ WaylandDevice = (*waylandDevice)(nil)

// WrapWaylandDevice wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandDevice(obj *externglib.Object) WaylandDevice {
	return WaylandDevice{
		gdk.Device: gdk.WrapDevice(obj),
	}
}

func marshalWaylandDevice(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandDevice(obj), nil
}

// NodePath returns the `/dev/input/event*` path of this device.
//
// For `GdkDevice`s that possibly coalesce multiple hardware devices (eg.
// mouse, keyboard, touch,...), this function will return nil.
//
// This is most notably implemented for devices of type GDK_SOURCE_PEN,
// GDK_SOURCE_TABLET_PAD.
func (d waylandDevice) NodePath() string {
	var arg0 *C.GdkDevice

	arg0 = (*C.GdkDevice)(d.Native())

	ret := C.gdk_wayland_device_get_node_path(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// WlKeyboard returns the Wayland `wl_keyboard` of a `GdkDevice`.
func (d waylandDevice) WlKeyboard() interface{} {
	var arg0 *C.GdkDevice

	arg0 = (*C.GdkDevice)(d.Native())

	ret := C.gdk_wayland_device_get_wl_keyboard(arg0)

	var ret0 interface{}

	ret0 = box.Get(uintptr(ret)).(interface{})

	return ret0
}

// WlPointer returns the Wayland `wl_pointer` of a `GdkDevice`.
func (d waylandDevice) WlPointer() interface{} {
	var arg0 *C.GdkDevice

	arg0 = (*C.GdkDevice)(d.Native())

	ret := C.gdk_wayland_device_get_wl_pointer(arg0)

	var ret0 interface{}

	ret0 = box.Get(uintptr(ret)).(interface{})

	return ret0
}

// WlSeat returns the Wayland `wl_seat` of a `GdkDevice`.
func (d waylandDevice) WlSeat() interface{} {
	var arg0 *C.GdkDevice

	arg0 = (*C.GdkDevice)(d.Native())

	ret := C.gdk_wayland_device_get_wl_seat(arg0)

	var ret0 interface{}

	ret0 = box.Get(uintptr(ret)).(interface{})

	return ret0
}
