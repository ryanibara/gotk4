// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_scale_button_get_type()), F: marshalScaleButtonner},
	})
}

// ScaleButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ScaleButtonOverrider interface {
	ValueChanged(value float64)
}

// ScaleButton provides a button which pops up a scale widget. This kind of
// widget is commonly used for volume controls in multimedia applications, and
// GTK+ provides a VolumeButton subclass that is tailored for this use case.
//
//
// CSS nodes
//
// GtkScaleButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .scale style class.
//
// The popup widget that contains the scale has a .scale-popup style class.
type ScaleButton struct {
	Button

	Orientable
	*externglib.Object
}

func wrapScaleButton(obj *externglib.Object) *ScaleButton {
	return &ScaleButton{
		Button: Button{
			Bin: Bin{
				Container: Container{
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
						Object: obj,
					},
				},
			},
			Actionable: Actionable{
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
					Object: obj,
				},
			},
			Activatable: Activatable{
				Object: obj,
			},
			Object: obj,
		},
		Orientable: Orientable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalScaleButtonner(p uintptr) (interface{}, error) {
	return wrapScaleButton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewScaleButton creates a ScaleButton, with a range between min and max, with
// a stepping of step.
//
// The function takes the following parameters:
//
//    - size: stock icon size (IconSize).
//    - min: minimum value of the scale (usually 0).
//    - max: maximum value of the scale (usually 100).
//    - step: stepping of value when a scroll-wheel event, or up/down arrow
//    event occurs (usually 2).
//    - icons: NULL-terminated array of icon names, or NULL if you want to set
//    the list later with gtk_scale_button_set_icons().
//
func NewScaleButton(size int, min, max, step float64, icons []string) *ScaleButton {
	var _arg1 C.GtkIconSize // out
	var _arg2 C.gdouble     // out
	var _arg3 C.gdouble     // out
	var _arg4 C.gdouble     // out
	var _arg5 **C.gchar     // out
	var _cret *C.GtkWidget  // in

	_arg1 = C.GtkIconSize(size)
	_arg2 = C.gdouble(min)
	_arg3 = C.gdouble(max)
	_arg4 = C.gdouble(step)
	{
		_arg5 = (**C.gchar)(C.malloc(C.size_t(len(icons)+1) * C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg5))
		{
			out := unsafe.Slice(_arg5, len(icons)+1)
			var zero *C.gchar
			out[len(icons)] = zero
			for i := range icons {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(icons[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.gtk_scale_button_new(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(size)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
	runtime.KeepAlive(step)
	runtime.KeepAlive(icons)

	var _scaleButton *ScaleButton // out

	_scaleButton = wrapScaleButton(externglib.Take(unsafe.Pointer(_cret)))

	return _scaleButton
}

// Adjustment gets the Adjustment associated with the ScaleButton’s scale. See
// gtk_range_get_adjustment() for details.
func (button *ScaleButton) Adjustment() *Adjustment {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkAdjustment  // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_scale_button_get_adjustment(_arg0)
	runtime.KeepAlive(button)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// MinusButton retrieves the minus button of the ScaleButton.
func (button *ScaleButton) MinusButton() *Button {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_scale_button_get_minus_button(_arg0)
	runtime.KeepAlive(button)

	var _ret *Button // out

	_ret = wrapButton(externglib.Take(unsafe.Pointer(_cret)))

	return _ret
}

// PlusButton retrieves the plus button of the ScaleButton.
func (button *ScaleButton) PlusButton() *Button {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_scale_button_get_plus_button(_arg0)
	runtime.KeepAlive(button)

	var _ret *Button // out

	_ret = wrapButton(externglib.Take(unsafe.Pointer(_cret)))

	return _ret
}

// Popup retrieves the popup of the ScaleButton.
func (button *ScaleButton) Popup() Widgetter {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_scale_button_get_popup(_arg0)
	runtime.KeepAlive(button)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Widgetter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// Value gets the current value of the scale button.
func (button *ScaleButton) Value() float64 {
	var _arg0 *C.GtkScaleButton // out
	var _cret C.gdouble         // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_scale_button_get_value(_arg0)
	runtime.KeepAlive(button)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetAdjustment sets the Adjustment to be used as a model for the ScaleButton’s
// scale. See gtk_range_set_adjustment() for details.
//
// The function takes the following parameters:
//
//    - adjustment: Adjustment.
//
func (button *ScaleButton) SetAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 *C.GtkAdjustment  // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_scale_button_set_adjustment(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(adjustment)
}

// SetIcons sets the icons to be used by the scale button. For details, see the
// ScaleButton:icons property.
//
// The function takes the following parameters:
//
//    - icons: NULL-terminated array of icon names.
//
func (button *ScaleButton) SetIcons(icons []string) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 **C.gchar         // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))
	{
		_arg1 = (**C.gchar)(C.malloc(C.size_t(len(icons)+1) * C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(icons)+1)
			var zero *C.gchar
			out[len(icons)] = zero
			for i := range icons {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(icons[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.gtk_scale_button_set_icons(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(icons)
}

// SetValue sets the current value of the scale; if the value is outside the
// minimum or maximum range values, it will be clamped to fit inside them. The
// scale button emits the ScaleButton::value-changed signal if the value
// changes.
//
// The function takes the following parameters:
//
//    - value: new value of the scale button.
//
func (button *ScaleButton) SetValue(value float64) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 C.gdouble         // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))
	_arg1 = C.gdouble(value)

	C.gtk_scale_button_set_value(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(value)
}

// ConnectPopdown signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted to popdown the scale widget.
//
// The default binding for this signal is Escape.
func (button *ScaleButton) ConnectPopdown(f func()) externglib.SignalHandle {
	return button.Connect("popdown", f)
}

// ConnectPopup signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted to popup the scale widget.
//
// The default bindings for this signal are Space, Enter and Return.
func (button *ScaleButton) ConnectPopup(f func()) externglib.SignalHandle {
	return button.Connect("popup", f)
}

// ConnectValueChanged signal is emitted when the value field has changed.
func (button *ScaleButton) ConnectValueChanged(f func(value float64)) externglib.SignalHandle {
	return button.Connect("value-changed", f)
}
