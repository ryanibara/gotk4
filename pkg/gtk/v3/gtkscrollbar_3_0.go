// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// NewScrollbar creates a new scrollbar with the given orientation.
//
// The function takes the following parameters:
//
//    - orientation scrollbar’s orientation.
//    - adjustment (optional) to use, or NULL to create a new adjustment.
//
// The function returns the following values:
//
//    - scrollbar: new Scrollbar.
//
func NewScrollbar(orientation Orientation, adjustment *Adjustment) *Scrollbar {
	var _arg1 C.GtkOrientation // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)
	if adjustment != nil {
		_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}

	_cret = C.gtk_scrollbar_new(_arg1, _arg2)
	runtime.KeepAlive(orientation)
	runtime.KeepAlive(adjustment)

	var _scrollbar *Scrollbar // out

	_scrollbar = wrapScrollbar(coreglib.Take(unsafe.Pointer(_cret)))

	return _scrollbar
}
