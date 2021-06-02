// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
)

// #cgo pkg-config: gtk4-x11
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
import "C"

// X11DeviceManagerLookup returns the Device that wraps the given device ID.
func X11DeviceManagerLookup(deviceManager X11DeviceManagerXI2, deviceID int) X11DeviceXI2 {
	var arg1 *C.GdkX11DeviceManagerXI2
	var arg2 C.int

	arg1 = (*C.GdkX11DeviceManagerXI2)(deviceManager.Native())
	arg2 = C.int(deviceID)

	ret := C.gdk_x11_device_manager_lookup(arg1, arg2)

	var ret0 X11DeviceXI2

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(X11DeviceXI2)

	return ret0
}
