// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gdk-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
//
// void gotk4_SeatGrabPrepareFunc(GdkSeat*, GdkWindow*, gpointer);
import "C"

// SeatGrabPrepareFunc: type of the callback used to set up @window so it can be
// grabbed. A typical action would be ensuring the window is visible, although
// there's room for other initialization actions.
type SeatGrabPrepareFunc func(seat Seat, window Window)

//export gotk4_SeatGrabPrepareFunc
func gotk4_SeatGrabPrepareFunc(arg0 *C.GdkSeat, arg1 *C.GdkWindow, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var seat Seat
	var window Window

	seat = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0.Native()))).(Seat)

	window = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1.Native()))).(Window)

	v.(SeatGrabPrepareFunc)(seat, window)
}
