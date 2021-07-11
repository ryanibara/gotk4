// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_hscale_get_type()), F: marshalHScaler},
	})
}

// HScaler describes HScale's methods.
type HScaler interface {
	privateHScale()
}

// HScale widget is used to allow the user to select a value using a horizontal
// slider. To create one, use gtk_hscale_new_with_range().
//
// The position to show the current value, and the number of decimal places
// shown can be set using the parent Scale class’s functions.
//
// GtkHScale has been deprecated, use Scale instead.
type HScale struct {
	Scale
}

var (
	_ HScaler         = (*HScale)(nil)
	_ gextras.Nativer = (*HScale)(nil)
)

func wrapHScale(obj *externglib.Object) HScaler {
	return &HScale{
		Scale: Scale{
			Range: Range{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
				Orientable: Orientable{
					Object: obj,
				},
			},
		},
	}
}

func marshalHScaler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapHScale(obj), nil
}

// NewHScale creates a new HScale.
//
// Deprecated: Use gtk_scale_new() with GTK_ORIENTATION_HORIZONTAL instead.
func NewHScale(adjustment Adjustmenter) *HScale {
	var _arg1 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer((adjustment).(gextras.Nativer).Native()))

	_cret = C.gtk_hscale_new(_arg1)

	var _hScale *HScale // out

	_hScale = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*HScale)

	return _hScale
}

// NewHScaleWithRange creates a new horizontal scale widget that lets the user
// input a number between @min and @max (including @min and @max) with the
// increment @step. @step must be nonzero; it’s the distance the slider moves
// when using the arrow keys to adjust the scale value.
//
// Note that the way in which the precision is derived works best if @step is a
// power of ten. If the resulting precision is not suitable for your needs, use
// gtk_scale_set_digits() to correct it.
//
// Deprecated: Use gtk_scale_new_with_range() with GTK_ORIENTATION_HORIZONTAL
// instead.
func NewHScaleWithRange(min float64, max float64, step float64) *HScale {
	var _arg1 C.gdouble    // out
	var _arg2 C.gdouble    // out
	var _arg3 C.gdouble    // out
	var _cret *C.GtkWidget // in

	_arg1 = C.gdouble(min)
	_arg2 = C.gdouble(max)
	_arg3 = C.gdouble(step)

	_cret = C.gtk_hscale_new_with_range(_arg1, _arg2, _arg3)

	var _hScale *HScale // out

	_hScale = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*HScale)

	return _hScale
}

func (*HScale) privateHScale() {}
