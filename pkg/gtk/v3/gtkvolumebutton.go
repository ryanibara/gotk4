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
	GTypeVolumeButton = coreglib.Type(C.gtk_volume_button_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVolumeButton, F: marshalVolumeButton},
	})
}

// VolumeButtonOverrider contains methods that are overridable.
type VolumeButtonOverrider interface {
}

// VolumeButton is a subclass of ScaleButton that has been tailored for use as a
// volume control widget with suitable icons, tooltips and accessible labels.
type VolumeButton struct {
	_ [0]func() // equal guard
	ScaleButton
}

var (
	_ coreglib.Objector = (*VolumeButton)(nil)
	_ Binner            = (*VolumeButton)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeVolumeButton,
		GoType:        reflect.TypeOf((*VolumeButton)(nil)),
		InitClass:     initClassVolumeButton,
		FinalizeClass: finalizeClassVolumeButton,
	})
}

func initClassVolumeButton(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ InitVolumeButton(*VolumeButtonClass) }); ok {
		klass := (*VolumeButtonClass)(gextras.NewStructNative(gclass))
		goval.InitVolumeButton(klass)
	}
}

func finalizeClassVolumeButton(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ FinalizeVolumeButton(*VolumeButtonClass) }); ok {
		klass := (*VolumeButtonClass)(gextras.NewStructNative(gclass))
		goval.FinalizeVolumeButton(klass)
	}
}

func wrapVolumeButton(obj *coreglib.Object) *VolumeButton {
	return &VolumeButton{
		ScaleButton: ScaleButton{
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
		},
	}
}

func marshalVolumeButton(p uintptr) (interface{}, error) {
	return wrapVolumeButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewVolumeButton creates a VolumeButton, with a range between 0.0 and 1.0,
// with a stepping of 0.02. Volume values can be obtained and modified using the
// functions from ScaleButton.
//
// The function returns the following values:
//
//    - volumeButton: new VolumeButton.
//
func NewVolumeButton() *VolumeButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_volume_button_new()

	var _volumeButton *VolumeButton // out

	_volumeButton = wrapVolumeButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _volumeButton
}

// VolumeButtonClass: instance of this type is always passed by reference.
type VolumeButtonClass struct {
	*volumeButtonClass
}

// volumeButtonClass is the struct that's finalized.
type volumeButtonClass struct {
	native *C.GtkVolumeButtonClass
}

func (v *VolumeButtonClass) ParentClass() *ScaleButtonClass {
	valptr := &v.native.parent_class
	var _v *ScaleButtonClass // out
	_v = (*ScaleButtonClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
