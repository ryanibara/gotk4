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
	GTypePopoverAccessible = coreglib.Type(C.gtk_popover_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePopoverAccessible, F: marshalPopoverAccessible},
	})
}

// PopoverAccessibleOverrider contains methods that are overridable.
type PopoverAccessibleOverrider interface {
}

type PopoverAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible
}

var (
	_ coreglib.Objector = (*PopoverAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypePopoverAccessible,
		GoType:        reflect.TypeOf((*PopoverAccessible)(nil)),
		InitClass:     initClassPopoverAccessible,
		FinalizeClass: finalizeClassPopoverAccessible,
	})
}

func initClassPopoverAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ InitPopoverAccessible(*PopoverAccessibleClass) }); ok {
		klass := (*PopoverAccessibleClass)(gextras.NewStructNative(gclass))
		goval.InitPopoverAccessible(klass)
	}
}

func finalizeClassPopoverAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ FinalizePopoverAccessible(*PopoverAccessibleClass) }); ok {
		klass := (*PopoverAccessibleClass)(gextras.NewStructNative(gclass))
		goval.FinalizePopoverAccessible(klass)
	}
}

func wrapPopoverAccessible(obj *coreglib.Object) *PopoverAccessible {
	return &PopoverAccessible{
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
	}
}

func marshalPopoverAccessible(p uintptr) (interface{}, error) {
	return wrapPopoverAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// PopoverAccessibleClass: instance of this type is always passed by reference.
type PopoverAccessibleClass struct {
	*popoverAccessibleClass
}

// popoverAccessibleClass is the struct that's finalized.
type popoverAccessibleClass struct {
	native *C.GtkPopoverAccessibleClass
}

func (p *PopoverAccessibleClass) ParentClass() *ContainerAccessibleClass {
	valptr := &p.native.parent_class
	var _v *ContainerAccessibleClass // out
	_v = (*ContainerAccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
