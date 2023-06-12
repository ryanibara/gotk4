// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeVScale = coreglib.Type(C.gtk_vscale_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVScale, F: marshalVScale},
	})
}

// VScaleOverrides contains methods that are overridable.
type VScaleOverrides struct {
}

func defaultVScaleOverrides(v *VScale) VScaleOverrides {
	return VScaleOverrides{}
}

// VScale widget is used to allow the user to select a value using a vertical
// slider. To create one, use gtk_hscale_new_with_range().
//
// The position to show the current value, and the number of decimal places
// shown can be set using the parent Scale class’s functions.
//
// GtkVScale has been deprecated, use Scale instead.
type VScale struct {
	_ [0]func() // equal guard
	Scale
}

var (
	_ Ranger = (*VScale)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*VScale, *VScaleClass, VScaleOverrides](
		GTypeVScale,
		initVScaleClass,
		wrapVScale,
		defaultVScaleOverrides,
	)
}

func initVScaleClass(gclass unsafe.Pointer, overrides VScaleOverrides, classInitFunc func(*VScaleClass)) {
	if classInitFunc != nil {
		class := (*VScaleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapVScale(obj *coreglib.Object) *VScale {
	return &VScale{
		Scale: Scale{
			Range: Range{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
				Object: obj,
				Orientable: Orientable{
					Object: obj,
				},
			},
		},
	}
}

func marshalVScale(p uintptr) (interface{}, error) {
	return wrapVScale(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewVScale creates a new VScale.
//
// Deprecated: Use gtk_scale_new() with GTK_ORIENTATION_VERTICAL instead.
//
// The function takes the following parameters:
//
//   - adjustment which sets the range of the scale.
//
// The function returns the following values:
//
//   - vScale: new VScale.
//
func NewVScale(adjustment *Adjustment) *VScale {
	var _arg1 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	_cret = C.gtk_vscale_new(_arg1)
	runtime.KeepAlive(adjustment)

	var _vScale *VScale // out

	_vScale = wrapVScale(coreglib.Take(unsafe.Pointer(_cret)))

	return _vScale
}

// NewVScaleWithRange creates a new vertical scale widget that lets the user
// input a number between min and max (including min and max) with the increment
// step. step must be nonzero; it’s the distance the slider moves when using the
// arrow keys to adjust the scale value.
//
// Note that the way in which the precision is derived works best if step is
// a power of ten. If the resulting precision is not suitable for your needs,
// use gtk_scale_set_digits() to correct it.
//
// Deprecated: Use gtk_scale_new_with_range() with GTK_ORIENTATION_VERTICAL
// instead.
//
// The function takes the following parameters:
//
//   - min: minimum value.
//   - max: maximum value.
//   - step increment (tick size) used with keyboard shortcuts.
//
// The function returns the following values:
//
//   - vScale: new VScale.
//
func NewVScaleWithRange(min, max, step float64) *VScale {
	var _arg1 C.gdouble    // out
	var _arg2 C.gdouble    // out
	var _arg3 C.gdouble    // out
	var _cret *C.GtkWidget // in

	_arg1 = C.gdouble(min)
	_arg2 = C.gdouble(max)
	_arg3 = C.gdouble(step)

	_cret = C.gtk_vscale_new_with_range(_arg1, _arg2, _arg3)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
	runtime.KeepAlive(step)

	var _vScale *VScale // out

	_vScale = wrapVScale(coreglib.Take(unsafe.Pointer(_cret)))

	return _vScale
}

// VScaleClass: instance of this type is always passed by reference.
type VScaleClass struct {
	*vScaleClass
}

// vScaleClass is the struct that's finalized.
type vScaleClass struct {
	native *C.GtkVScaleClass
}

func (v *VScaleClass) ParentClass() *ScaleClass {
	valptr := &v.native.parent_class
	var _v *ScaleClass // out
	_v = (*ScaleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
