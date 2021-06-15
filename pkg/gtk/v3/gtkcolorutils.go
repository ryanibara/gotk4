// Code generated by girgen. DO NOT EDIT.

package gtk

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// RGBToHSV converts a color from RGB space to HSV.
//
// Input values must be in the [0.0, 1.0] range; output values will be in the
// same range.
func RGBToHSV(r float64, g float64, b float64) (h float64, s float64, v float64) {
	var _arg1 C.gdouble // out
	var _arg2 C.gdouble // out
	var _arg3 C.gdouble // out
	var _arg4 C.gdouble // in
	var _arg5 C.gdouble // in
	var _arg6 C.gdouble // in

	_arg1 = (C.gdouble)(r)
	_arg2 = (C.gdouble)(g)
	_arg3 = (C.gdouble)(b)

	C.gtk_rgb_to_hsv(_arg1, _arg2, _arg3, &_arg4, &_arg5, &_arg6)

	var _h float64 // out
	var _s float64 // out
	var _v float64 // out

	_h = (float64)(_arg4)
	_s = (float64)(_arg5)
	_v = (float64)(_arg6)

	return _h, _s, _v
}
