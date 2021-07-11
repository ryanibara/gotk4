// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/x11/gdkx.h>
import "C"

// X11DeviceManagerLookup returns the Device that wraps the given device ID.
func X11DeviceManagerLookup(deviceManager X11DeviceManagerXI2er, deviceId int) *X11DeviceXI2 {
	var _arg1 *C.GdkX11DeviceManagerXI2 // out
	var _arg2 C.int                     // out
	var _cret *C.GdkDevice              // in

	_arg1 = (*C.GdkX11DeviceManagerXI2)(unsafe.Pointer((deviceManager).(gextras.Nativer).Native()))
	_arg2 = C.int(deviceId)

	_cret = C.gdk_x11_device_manager_lookup(_arg1, _arg2)

	var _x11DeviceXI2 *X11DeviceXI2 // out

	_x11DeviceXI2 = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*X11DeviceXI2)

	return _x11DeviceXI2
}
