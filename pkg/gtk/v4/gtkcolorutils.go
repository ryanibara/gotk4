// Code generated by girgen. DO NOT EDIT.

package gtk

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

// HSVToRGB converts a color from HSV space to RGB.
//
// Input values must be in the [0.0, 1.0] range; output values will be in the
// same range.
func HSVToRGB(h float32, s float32, v float32) (r float32, g float32, b float32) {
	var arg1 C.float
	var arg2 C.float
	var arg3 C.float

	arg1 = C.float(h)
	arg2 = C.float(s)
	arg3 = C.float(v)

	var arg4 C.float
	var ret4 float32
	var arg5 C.float
	var ret5 float32
	var arg6 C.float
	var ret6 float32

	C.gtk_hsv_to_rgb(h, s, v, &arg4, &arg5, &arg6)

	*ret4 = C.float(arg4)
	*ret5 = C.float(arg5)
	*ret6 = C.float(arg6)

	return ret4, ret5, ret6
}

// RGBToHSV converts a color from RGB space to HSV.
//
// Input values must be in the [0.0, 1.0] range; output values will be in the
// same range.
func RGBToHSV(r float32, g float32, b float32) (h float32, s float32, v float32) {
	var arg1 C.float
	var arg2 C.float
	var arg3 C.float

	arg1 = C.float(r)
	arg2 = C.float(g)
	arg3 = C.float(b)

	var arg4 C.float
	var ret4 float32
	var arg5 C.float
	var ret5 float32
	var arg6 C.float
	var ret6 float32

	C.gtk_rgb_to_hsv(r, g, b, &arg4, &arg5, &arg6)

	*ret4 = C.float(arg4)
	*ret5 = C.float(arg5)
	*ret6 = C.float(arg6)

	return ret4, ret5, ret6
}
