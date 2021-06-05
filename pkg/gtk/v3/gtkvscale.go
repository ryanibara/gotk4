// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.gtk_vscale_get_type()), F: marshalVScale},
	})
}

// VScale: the VScale widget is used to allow the user to select a value using a
// vertical slider. To create one, use gtk_hscale_new_with_range().
//
// The position to show the current value, and the number of decimal places
// shown can be set using the parent Scale class’s functions.
//
// GtkVScale has been deprecated, use Scale instead.
type VScale interface {
	Scale
	Buildable
	Orientable
}

// vScale implements the VScale interface.
type vScale struct {
	Scale
	Buildable
	Orientable
}

var _ VScale = (*vScale)(nil)

// WrapVScale wraps a GObject to the right type. It is
// primarily used internally.
func WrapVScale(obj *externglib.Object) VScale {
	return VScale{
		Scale:      WrapScale(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalVScale(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVScale(obj), nil
}

// NewVScale constructs a class VScale.
func NewVScale(adjustment Adjustment) VScale {
	var arg1 *C.GtkAdjustment

	arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	var cret C.GtkVScale
	var ret1 VScale

	cret = C.gtk_vscale_new(adjustment)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(VScale)

	return ret1
}

// NewVScaleWithRange constructs a class VScale.
func NewVScaleWithRange(min float64, max float64, step float64) VScale {
	var arg1 C.gdouble
	var arg2 C.gdouble
	var arg3 C.gdouble

	arg1 = C.gdouble(min)
	arg2 = C.gdouble(max)
	arg3 = C.gdouble(step)

	var cret C.GtkVScale
	var ret1 VScale

	cret = C.gtk_vscale_new_with_range(min, max, step)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(VScale)

	return ret1
}
