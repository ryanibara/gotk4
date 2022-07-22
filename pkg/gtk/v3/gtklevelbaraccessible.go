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
	GTypeLevelBarAccessible = coreglib.Type(C.gtk_level_bar_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeLevelBarAccessible, F: marshalLevelBarAccessible},
	})
}

// LevelBarAccessibleOverrider contains methods that are overridable.
type LevelBarAccessibleOverrider interface {
}

type LevelBarAccessible struct {
	_ [0]func() // equal guard
	WidgetAccessible

	atk.Value
}

var (
	_ coreglib.Objector = (*LevelBarAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeLevelBarAccessible,
		GoType:        reflect.TypeOf((*LevelBarAccessible)(nil)),
		InitClass:     initClassLevelBarAccessible,
		FinalizeClass: finalizeClassLevelBarAccessible,
	})
}

func initClassLevelBarAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		InitLevelBarAccessible(*LevelBarAccessibleClass)
	}); ok {
		klass := (*LevelBarAccessibleClass)(gextras.NewStructNative(gclass))
		goval.InitLevelBarAccessible(klass)
	}
}

func finalizeClassLevelBarAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		FinalizeLevelBarAccessible(*LevelBarAccessibleClass)
	}); ok {
		klass := (*LevelBarAccessibleClass)(gextras.NewStructNative(gclass))
		goval.FinalizeLevelBarAccessible(klass)
	}
}

func wrapLevelBarAccessible(obj *coreglib.Object) *LevelBarAccessible {
	return &LevelBarAccessible{
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
		Value: atk.Value{
			Object: obj,
		},
	}
}

func marshalLevelBarAccessible(p uintptr) (interface{}, error) {
	return wrapLevelBarAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// LevelBarAccessibleClass: instance of this type is always passed by reference.
type LevelBarAccessibleClass struct {
	*levelBarAccessibleClass
}

// levelBarAccessibleClass is the struct that's finalized.
type levelBarAccessibleClass struct {
	native *C.GtkLevelBarAccessibleClass
}

func (l *LevelBarAccessibleClass) ParentClass() *WidgetAccessibleClass {
	valptr := &l.native.parent_class
	var _v *WidgetAccessibleClass // out
	_v = (*WidgetAccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
