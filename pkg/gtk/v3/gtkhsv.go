// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_hsv_get_type()), F: marshalHSV},
	})
}

// HSV is the “color wheel” part of a complete color selector widget. It allows
// to select a color by determining its HSV components in an intuitive way.
// Moving the selection around the outer ring changes the hue, and moving the
// selection point inside the inner triangle changes value and saturation.
//
// HSV has been deprecated together with ColorSelection, where it was used.
type HSV interface {
	Widget
	Buildable

	// Color queries the current color in an HSV color selector. Returned values
	// will be in the [0.0, 1.0] range.
	Color() (h float64, s float64, v float64)
	// Metrics queries the size and ring width of an HSV color selector.
	Metrics() (size int, ringWidth int)
	// IsAdjusting: an HSV color selector can be said to be adjusting if
	// multiple rapid changes are being made to its value, for example, when the
	// user is adjusting the value with the mouse. This function queries whether
	// the HSV color selector is being adjusted or not.
	IsAdjusting() bool
	// SetColor sets the current color in an HSV color selector. Color component
	// values must be in the [0.0, 1.0] range.
	SetColor(h float64, s float64, v float64)
	// SetMetrics sets the size and ring width of an HSV color selector.
	SetMetrics(size int, ringWidth int)
}

// hsV implements the HSV interface.
type hsV struct {
	Widget
	Buildable
}

var _ HSV = (*hsV)(nil)

// WrapHSV wraps a GObject to the right type. It is
// primarily used internally.
func WrapHSV(obj *externglib.Object) HSV {
	return HSV{
		Widget:    WrapWidget(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalHSV(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapHSV(obj), nil
}

// Color queries the current color in an HSV color selector. Returned values
// will be in the [0.0, 1.0] range.
func (h hsV) Color() (h float64, s float64, v float64) {
	var _arg0 *C.GtkHSV

	_arg0 = (*C.GtkHSV)(unsafe.Pointer(h.Native()))

	var _arg1 C.gdouble
	var _arg2 C.gdouble
	var _arg3 C.gdouble

	C.gtk_hsv_get_color(_arg0, &_arg1, &_arg2, &_arg3)

	var _h float64
	var _s float64
	var _v float64

	_h = (float64)(_arg1)
	_s = (float64)(_arg2)
	_v = (float64)(_arg3)

	return _h, _s, _v
}

// Metrics queries the size and ring width of an HSV color selector.
func (h hsV) Metrics() (size int, ringWidth int) {
	var _arg0 *C.GtkHSV

	_arg0 = (*C.GtkHSV)(unsafe.Pointer(h.Native()))

	var _arg1 C.gint
	var _arg2 C.gint

	C.gtk_hsv_get_metrics(_arg0, &_arg1, &_arg2)

	var _size int
	var _ringWidth int

	_size = (int)(_arg1)
	_ringWidth = (int)(_arg2)

	return _size, _ringWidth
}

// IsAdjusting: an HSV color selector can be said to be adjusting if
// multiple rapid changes are being made to its value, for example, when the
// user is adjusting the value with the mouse. This function queries whether
// the HSV color selector is being adjusted or not.
func (h hsV) IsAdjusting() bool {
	var _arg0 *C.GtkHSV

	_arg0 = (*C.GtkHSV)(unsafe.Pointer(h.Native()))

	var _cret C.gboolean

	_cret = C.gtk_hsv_is_adjusting(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetColor sets the current color in an HSV color selector. Color component
// values must be in the [0.0, 1.0] range.
func (h hsV) SetColor(h float64, s float64, v float64) {
	var _arg0 *C.GtkHSV
	var _arg1 C.double
	var _arg2 C.double
	var _arg3 C.double

	_arg0 = (*C.GtkHSV)(unsafe.Pointer(h.Native()))
	_arg1 = C.double(h)
	_arg2 = C.double(s)
	_arg3 = C.double(v)

	C.gtk_hsv_set_color(_arg0, _arg1, _arg2, _arg3)
}

// SetMetrics sets the size and ring width of an HSV color selector.
func (h hsV) SetMetrics(size int, ringWidth int) {
	var _arg0 *C.GtkHSV
	var _arg1 C.gint
	var _arg2 C.gint

	_arg0 = (*C.GtkHSV)(unsafe.Pointer(h.Native()))
	_arg1 = C.gint(size)
	_arg2 = C.gint(ringWidth)

	C.gtk_hsv_set_metrics(_arg0, _arg1, _arg2)
}
