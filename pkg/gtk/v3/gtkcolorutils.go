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
	var arg1 C.gdouble
	var arg2 C.gdouble
	var arg3 C.gdouble
	var arg4 *C.gdouble // out
	var arg5 *C.gdouble // out
	var arg6 *C.gdouble // out

	arg1 = C.gdouble(r)
	arg2 = C.gdouble(g)
	arg3 = C.gdouble(b)

	C.gtk_rgb_to_hsv(arg1, arg2, arg3, &arg4, &arg5, &arg6)

	var ret0 float64
	var ret1 float64
	var ret2 float64

	ret0 = float64(arg4)

	ret1 = float64(arg5)

	ret2 = float64(arg6)

	return ret0, ret1, ret2
}
