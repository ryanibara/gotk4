// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

// Axes returns the axes currently available on the device.
//
// The function returns the following values:
//
func (device *Device) Axes() AxisFlags {
	var _arg0 *C.GdkDevice   // out
	var _cret C.GdkAxisFlags // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_cret = C.gdk_device_get_axes(_arg0)
	runtime.KeepAlive(device)

	var _axisFlags AxisFlags // out

	_axisFlags = AxisFlags(_cret)

	return _axisFlags
}
