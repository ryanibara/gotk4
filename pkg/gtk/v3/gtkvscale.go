// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtkvscale.go.
var GTypeVScale = externglib.Type(C.gtk_vscale_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeVScale, F: marshalVScale},
	})
}

// VScaleOverrider contains methods that are overridable.
type VScaleOverrider interface {
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

func classInitVScaler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapVScale(obj *externglib.Object) *VScale {
	return &VScale{
		Scale: Scale{
			Range: Range{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
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
	return wrapVScale(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewVScale creates a new VScale.
//
// Deprecated: Use gtk_scale_new() with GTK_ORIENTATION_VERTICAL instead.
//
// The function takes the following parameters:
//
//    - adjustment which sets the range of the scale.
//
// The function returns the following values:
//
//    - vScale: new VScale.
//
func NewVScale(adjustment *Adjustment) *VScale {
	var _arg1 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))

	_cret = C.gtk_vscale_new(_arg1)
	runtime.KeepAlive(adjustment)

	var _vScale *VScale // out

	_vScale = wrapVScale(externglib.Take(unsafe.Pointer(_cret)))

	return _vScale
}

// NewVScaleWithRange creates a new vertical scale widget that lets the user
// input a number between min and max (including min and max) with the increment
// step. step must be nonzero; it’s the distance the slider moves when using the
// arrow keys to adjust the scale value.
//
// Note that the way in which the precision is derived works best if step is a
// power of ten. If the resulting precision is not suitable for your needs, use
// gtk_scale_set_digits() to correct it.
//
// Deprecated: Use gtk_scale_new_with_range() with GTK_ORIENTATION_VERTICAL
// instead.
//
// The function takes the following parameters:
//
//    - min: minimum value.
//    - max: maximum value.
//    - step increment (tick size) used with keyboard shortcuts.
//
// The function returns the following values:
//
//    - vScale: new VScale.
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

	_vScale = wrapVScale(externglib.Take(unsafe.Pointer(_cret)))

	return _vScale
}
