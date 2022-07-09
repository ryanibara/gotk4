// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_ScaleButtonClass_value_changed(void*, gdouble);
// extern void _gotk4_gtk3_ScaleButton_ConnectPopdown(gpointer, guintptr);
// extern void _gotk4_gtk3_ScaleButton_ConnectPopup(gpointer, guintptr);
// extern void _gotk4_gtk3_ScaleButton_ConnectValueChanged(gpointer, gdouble, guintptr);
import "C"

// GTypeScaleButton returns the GType for the type ScaleButton.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeScaleButton() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ScaleButton").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalScaleButton)
	return gtype
}

// ScaleButtonOverrider contains methods that are overridable.
type ScaleButtonOverrider interface {
	// The function takes the following parameters:
	//
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
	_ [0]func() // equal guard
	Button

	*coreglib.Object
	Orientable
}

var (
	_ coreglib.Objector = (*ScaleButton)(nil)
	_ Binner            = (*ScaleButton)(nil)
)

func classInitScaleButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "ScaleButtonClass")

	if _, ok := goval.(interface{ ValueChanged(value float64) }); ok {
		o := pclass.StructFieldOffset("value_changed")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ScaleButtonClass_value_changed)
	}
}

//export _gotk4_gtk3_ScaleButtonClass_value_changed
func _gotk4_gtk3_ScaleButtonClass_value_changed(arg0 *C.void, arg1 C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ValueChanged(value float64) })

	var _value float64 // out

	_value = float64(arg1)

	iface.ValueChanged(_value)
}

func wrapScaleButton(obj *coreglib.Object) *ScaleButton {
	return &ScaleButton{
		Button: Button{
			Bin: Bin{
				Container: Container{
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
				},
			},
			Object: obj,
			Actionable: Actionable{
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
			},
			Activatable: Activatable{
				Object: obj,
			},
		},
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalScaleButton(p uintptr) (interface{}, error) {
	return wrapScaleButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_ScaleButton_ConnectPopdown
func _gotk4_gtk3_ScaleButton_ConnectPopdown(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectPopdown signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted to popdown the scale widget.
//
// The default binding for this signal is Escape.
func (button *ScaleButton) ConnectPopdown(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "popdown", false, unsafe.Pointer(C._gotk4_gtk3_ScaleButton_ConnectPopdown), f)
}

//export _gotk4_gtk3_ScaleButton_ConnectPopup
func _gotk4_gtk3_ScaleButton_ConnectPopup(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectPopup signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted to popup the scale widget.
//
// The default bindings for this signal are Space, Enter and Return.
func (button *ScaleButton) ConnectPopup(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "popup", false, unsafe.Pointer(C._gotk4_gtk3_ScaleButton_ConnectPopup), f)
}

//export _gotk4_gtk3_ScaleButton_ConnectValueChanged
func _gotk4_gtk3_ScaleButton_ConnectValueChanged(arg0 C.gpointer, arg1 C.gdouble, arg2 C.guintptr) {
	var f func(value float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(value float64))
	}

	var _value float64 // out

	_value = float64(arg1)

	f(_value)
}

// ConnectValueChanged signal is emitted when the value field has changed.
func (button *ScaleButton) ConnectValueChanged(f func(value float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "value-changed", false, unsafe.Pointer(C._gotk4_gtk3_ScaleButton_ConnectValueChanged), f)
}

// Adjustment gets the Adjustment associated with the ScaleButton’s scale. See
// gtk_range_get_adjustment() for details.
//
// The function returns the following values:
//
//    - adjustment associated with the scale.
//
func (button *ScaleButton) Adjustment() *Adjustment {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "ScaleButton").InvokeMethod("get_adjustment", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// MinusButton retrieves the minus button of the ScaleButton.
//
// The function returns the following values:
//
//    - ret minus button of the ScaleButton as a Button.
//
func (button *ScaleButton) MinusButton() *Button {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "ScaleButton").InvokeMethod("get_minus_button", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _ret *Button // out

	_ret = wrapButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _ret
}

// PlusButton retrieves the plus button of the ScaleButton.
//
// The function returns the following values:
//
//    - ret plus button of the ScaleButton as a Button.
//
func (button *ScaleButton) PlusButton() *Button {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "ScaleButton").InvokeMethod("get_plus_button", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _ret *Button // out

	_ret = wrapButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _ret
}

// Popup retrieves the popup of the ScaleButton.
//
// The function returns the following values:
//
//    - widget: popup of the ScaleButton.
//
func (button *ScaleButton) Popup() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "ScaleButton").InvokeMethod("get_popup", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// Value gets the current value of the scale button.
//
// The function returns the following values:
//
//    - gdouble: current value of the scale button.
//
func (button *ScaleButton) Value() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "ScaleButton").InvokeMethod("get_value", _args[:], nil)
	_cret = *(*C.gdouble)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.gdouble)(unsafe.Pointer(&_cret)))

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	girepository.MustFind("Gtk", "ScaleButton").InvokeMethod("set_adjustment", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	{
		*(***C.void)(unsafe.Pointer(&_args[1])) = (**C.void)(C.calloc(C.size_t((len(icons) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_args[1]))
		{
			out := unsafe.Slice(_args[1], len(icons)+1)
			var zero *C.void
			out[len(icons)] = zero
			for i := range icons {
				*(**C.void)(unsafe.Pointer(&out[i])) = (*C.void)(unsafe.Pointer(C.CString(icons[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	girepository.MustFind("Gtk", "ScaleButton").InvokeMethod("set_icons", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	*(*C.gdouble)(unsafe.Pointer(&_args[1])) = C.gdouble(value)

	girepository.MustFind("Gtk", "ScaleButton").InvokeMethod("set_value", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(value)
}
