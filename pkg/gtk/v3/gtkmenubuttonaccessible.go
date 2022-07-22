// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
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
	GTypeMenuButtonAccessible = coreglib.Type(C.gtk_menu_button_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMenuButtonAccessible, F: marshalMenuButtonAccessible},
	})
}

// MenuButtonAccessibleOverrider contains methods that are overridable.
type MenuButtonAccessibleOverrider interface {
}

type MenuButtonAccessible struct {
	_ [0]func() // equal guard
	ToggleButtonAccessible
}

var (
	_ coreglib.Objector = (*MenuButtonAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:     GTypeMenuButtonAccessible,
		GoType:    reflect.TypeOf((*MenuButtonAccessible)(nil)),
		InitClass: initClassMenuButtonAccessible,
	})
}

func initClassMenuButtonAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		InitMenuButtonAccessible(*MenuButtonAccessibleClass)
	}); ok {
		klass := (*MenuButtonAccessibleClass)(gextras.NewStructNative(gclass))
		goval.InitMenuButtonAccessible(klass)
	}
}

func wrapMenuButtonAccessible(obj *coreglib.Object) *MenuButtonAccessible {
	return &MenuButtonAccessible{
		ToggleButtonAccessible: ToggleButtonAccessible{
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
		},
	}
}

func marshalMenuButtonAccessible(p uintptr) (interface{}, error) {
	return wrapMenuButtonAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// MenuButtonAccessibleClass: instance of this type is always passed by
// reference.
type MenuButtonAccessibleClass struct {
	*menuButtonAccessibleClass
}

// menuButtonAccessibleClass is the struct that's finalized.
type menuButtonAccessibleClass struct {
	native *C.GtkMenuButtonAccessibleClass
}

func (m *MenuButtonAccessibleClass) ParentClass() *ToggleButtonAccessibleClass {
	valptr := &m.native.parent_class
	var v *ToggleButtonAccessibleClass // out
	v = (*ToggleButtonAccessibleClass)(gextras.NewStructNative(unsafe.Pointer((&*valptr))))
	return v
}
