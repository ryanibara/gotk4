// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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

// glib.Type values for gtkvolumebutton.go.
var GTypeVolumeButton = externglib.Type(C.gtk_volume_button_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeVolumeButton, F: marshalVolumeButton},
	})
}

// VolumeButtonOverrider contains methods that are overridable.
type VolumeButtonOverrider interface {
	externglib.Objector
}

// WrapVolumeButtonOverrider wraps the VolumeButtonOverrider
// interface implementation to access the instance methods.
func WrapVolumeButtonOverrider(obj VolumeButtonOverrider) *VolumeButton {
	return wrapVolumeButton(externglib.BaseObject(obj))
}

// VolumeButton is a subclass of ScaleButton that has been tailored for use as a
// volume control widget with suitable icons, tooltips and accessible labels.
type VolumeButton struct {
	_ [0]func() // equal guard
	ScaleButton
}

var (
	_ externglib.Objector = (*VolumeButton)(nil)
	_ Binner              = (*VolumeButton)(nil)
)

func classInitVolumeButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapVolumeButton(obj *externglib.Object) *VolumeButton {
	return &VolumeButton{
		ScaleButton: ScaleButton{
			Button: Button{
				Bin: Bin{
					Container: Container{
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
					},
				},
				Object: obj,
				Actionable: Actionable{
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
	return wrapVolumeButton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
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

	_volumeButton = wrapVolumeButton(externglib.Take(unsafe.Pointer(_cret)))

	return _volumeButton
}
