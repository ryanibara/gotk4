// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_volume_button_get_type()), F: marshalVolumeButtonner},
	})
}

// VolumeButton: GtkVolumeButton is a GtkScaleButton subclass tailored for
// volume control.
//
// !An example GtkVolumeButton (volumebutton.png).
type VolumeButton struct {
	ScaleButton
}

var (
	_ Widgetter           = (*VolumeButton)(nil)
	_ externglib.Objector = (*VolumeButton)(nil)
)

func wrapVolumeButton(obj *externglib.Object) *VolumeButton {
	return &VolumeButton{
		ScaleButton: ScaleButton{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
				Object: obj,
			},
			Orientable: Orientable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalVolumeButtonner(p uintptr) (interface{}, error) {
	return wrapVolumeButton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewVolumeButton creates a GtkVolumeButton.
//
// The button has a range between 0.0 and 1.0, with a stepping of 0.02. Volume
// values can be obtained and modified using the functions from gtk.ScaleButton.
func NewVolumeButton() *VolumeButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_volume_button_new()

	var _volumeButton *VolumeButton // out

	_volumeButton = wrapVolumeButton(externglib.Take(unsafe.Pointer(_cret)))

	return _volumeButton
}
