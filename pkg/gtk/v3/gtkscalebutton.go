// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
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
// extern void _gotk4_gtk3_ScaleButton_ConnectValueChanged(gpointer, gdouble, guintptr);
// extern void _gotk4_gtk3_ScaleButton_ConnectPopup(gpointer, guintptr);
// extern void _gotk4_gtk3_ScaleButton_ConnectPopdown(gpointer, guintptr);
// extern void _gotk4_gtk3_ScaleButtonClass_value_changed(GtkScaleButton*, gdouble);
// void _gotk4_gtk3_ScaleButton_virtual_value_changed(void* fnptr, GtkScaleButton* arg0, gdouble arg1) {
//   ((void (*)(GtkScaleButton*, gdouble))(fnptr))(arg0, arg1);
// };
import "C"

// GType values.
var (
	GTypeScaleButton = coreglib.Type(C.gtk_scale_button_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeScaleButton, F: marshalScaleButton},
	})
}

// ScaleButtonOverrides contains methods that are overridable.
type ScaleButtonOverrides struct {
	// The function takes the following parameters:
	//
	ValueChanged func(value float64)
}

func defaultScaleButtonOverrides(v *ScaleButton) ScaleButtonOverrides {
	return ScaleButtonOverrides{
		ValueChanged: v.valueChanged,
	}
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

func init() {
	coreglib.RegisterClassInfo[*ScaleButton, *ScaleButtonClass, ScaleButtonOverrides](
		GTypeScaleButton,
		initScaleButtonClass,
		wrapScaleButton,
		defaultScaleButtonOverrides,
	)
}

func initScaleButtonClass(gclass unsafe.Pointer, overrides ScaleButtonOverrides, classInitFunc func(*ScaleButtonClass)) {
	pclass := (*C.GtkScaleButtonClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeScaleButton))))

	if overrides.ValueChanged != nil {
		pclass.value_changed = (*[0]byte)(C._gotk4_gtk3_ScaleButtonClass_value_changed)
	}

	if classInitFunc != nil {
		class := (*ScaleButtonClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
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

// ConnectPopdown signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted to popdown the scale widget.
//
// The default binding for this signal is Escape.
func (button *ScaleButton) ConnectPopdown(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "popdown", false, unsafe.Pointer(C._gotk4_gtk3_ScaleButton_ConnectPopdown), f)
}

// ConnectPopup signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted to popup the scale widget.
//
// The default bindings for this signal are Space, Enter and Return.
func (button *ScaleButton) ConnectPopup(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "popup", false, unsafe.Pointer(C._gotk4_gtk3_ScaleButton_ConnectPopup), f)
}

// ConnectValueChanged signal is emitted when the value field has changed.
func (button *ScaleButton) ConnectValueChanged(f func(value float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "value-changed", false, unsafe.Pointer(C._gotk4_gtk3_ScaleButton_ConnectValueChanged), f)
}

// The function takes the following parameters:
//
func (button *ScaleButton) valueChanged(value float64) {
	gclass := (*C.GtkScaleButtonClass)(coreglib.PeekParentClass(button))
	fnarg := gclass.value_changed

	var _arg0 *C.GtkScaleButton // out
	var _arg1 C.gdouble         // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	_arg1 = C.gdouble(value)

	C._gotk4_gtk3_ScaleButton_virtual_value_changed(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(value)
}

// ScaleButtonClass: instance of this type is always passed by reference.
type ScaleButtonClass struct {
	*scaleButtonClass
}

// scaleButtonClass is the struct that's finalized.
type scaleButtonClass struct {
	native *C.GtkScaleButtonClass
}

func (s *ScaleButtonClass) ParentClass() *ButtonClass {
	valptr := &s.native.parent_class
	var _v *ButtonClass // out
	_v = (*ButtonClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
