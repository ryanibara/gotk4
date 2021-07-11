// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_volume_button_get_type()), F: marshalVolumeButtonner},
	})
}

// VolumeButtonner describes VolumeButton's methods.
type VolumeButtonner interface {
	privateVolumeButton()
}

// VolumeButton: `GtkVolumeButton` is a `GtkScaleButton` subclass tailored for
// volume control.
//
// !An example GtkVolumeButton (volumebutton.png)
type VolumeButton struct {
	ScaleButton
}

var (
	_ VolumeButtonner = (*VolumeButton)(nil)
	_ gextras.Nativer = (*VolumeButton)(nil)
)

func wrapVolumeButton(obj *externglib.Object) VolumeButtonner {
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
			},
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalVolumeButtonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapVolumeButton(obj), nil
}

// NewVolumeButton creates a `GtkVolumeButton`.
//
// The button has a range between 0.0 and 1.0, with a stepping of 0.02. Volume
// values can be obtained and modified using the functions from
// [class@Gtk.ScaleButton].
func NewVolumeButton() *VolumeButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_volume_button_new()

	var _volumeButton *VolumeButton // out

	_volumeButton = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*VolumeButton)

	return _volumeButton
}

func (*VolumeButton) privateVolumeButton() {}
