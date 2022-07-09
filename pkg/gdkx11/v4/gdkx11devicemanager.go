// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// X11DeviceManagerLookup returns the Device that wraps the given device ID.
//
// The function takes the following parameters:
//
//    - deviceManager: DeviceManager.
//    - deviceId: device ID, as understood by the XInput2 protocol.
//
// The function returns the following values:
//
//    - x11DeviceXI2 (optional) wrapping the device ID, or NULL if the given ID
//      doesn’t currently represent a device.
//
func X11DeviceManagerLookup(deviceManager *X11DeviceManagerXI2, deviceId int32) *X11DeviceXI2 {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(deviceManager).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(deviceId)

	_info := girepository.MustFind("GdkX11", "x11_device_manager_lookup")
	_gret := _info.Invoke(_args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(deviceManager)
	runtime.KeepAlive(deviceId)

	var _x11DeviceXI2 *X11DeviceXI2 // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_x11DeviceXI2 = wrapX11DeviceXI2(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _x11DeviceXI2
}
