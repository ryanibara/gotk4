// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
	GTypeToggleButtonAccessible = coreglib.Type(C.gtk_toggle_button_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeToggleButtonAccessible, F: marshalToggleButtonAccessible},
	})
}

// ToggleButtonAccessibleOverrides contains methods that are overridable.
type ToggleButtonAccessibleOverrides struct {
}

func defaultToggleButtonAccessibleOverrides(v *ToggleButtonAccessible) ToggleButtonAccessibleOverrides {
	return ToggleButtonAccessibleOverrides{}
}

type ToggleButtonAccessible struct {
	_ [0]func() // equal guard
	ButtonAccessible
}

var (
	_ coreglib.Objector = (*ToggleButtonAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ToggleButtonAccessible, *ToggleButtonAccessibleClass, ToggleButtonAccessibleOverrides](
		GTypeToggleButtonAccessible,
		initToggleButtonAccessibleClass,
		wrapToggleButtonAccessible,
		defaultToggleButtonAccessibleOverrides,
	)
}

func initToggleButtonAccessibleClass(gclass unsafe.Pointer, overrides ToggleButtonAccessibleOverrides, classInitFunc func(*ToggleButtonAccessibleClass)) {
	if classInitFunc != nil {
		class := (*ToggleButtonAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapToggleButtonAccessible(obj *coreglib.Object) *ToggleButtonAccessible {
	return &ToggleButtonAccessible{
		ButtonAccessible: ButtonAccessible{
			ContainerAccessible: ContainerAccessible{
				WidgetAccessible: WidgetAccessible{
					Accessible: Accessible{
						AtkObject: atk.AtkObject{
							Object: obj,
						},
					},
					Component: atk.Component{
						Object: obj,
					},
				},
			},
			Object: obj,
			Action: atk.Action{
				Object: obj,
			},
			Image: atk.Image{
				Object: obj,
			},
		},
	}
}

func marshalToggleButtonAccessible(p uintptr) (interface{}, error) {
	return wrapToggleButtonAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ToggleButtonAccessibleClass: instance of this type is always passed by
// reference.
type ToggleButtonAccessibleClass struct {
	*toggleButtonAccessibleClass
}

// toggleButtonAccessibleClass is the struct that's finalized.
type toggleButtonAccessibleClass struct {
	native *C.GtkToggleButtonAccessibleClass
}

func (t *ToggleButtonAccessibleClass) ParentClass() *ButtonAccessibleClass {
	valptr := &t.native.parent_class
	var _v *ButtonAccessibleClass // out
	_v = (*ButtonAccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
