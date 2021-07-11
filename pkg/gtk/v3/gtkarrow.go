// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_arrow_get_type()), F: marshalArrowwer},
	})
}

// Arrowwer describes Arrow's methods.
type Arrowwer interface {
	privateArrow()
}

// Arrow should be used to draw simple arrows that need to point in one of the
// four cardinal directions (up, down, left, or right). The style of the arrow
// can be one of shadow in, shadow out, etched in, or etched out. Note that
// these directions and style types may be amended in versions of GTK+ to come.
//
// GtkArrow will fill any space alloted to it, but since it is inherited from
// Misc, it can be padded and/or aligned, to fill exactly the space the
// programmer desires.
//
// Arrows are created with a call to gtk_arrow_new(). The direction or style of
// an arrow can be changed after creation by using gtk_arrow_set().
//
// GtkArrow has been deprecated; you can simply use a Image with a suitable icon
// name, such as “pan-down-symbolic“. When replacing GtkArrow by an image, pay
// attention to the fact that GtkArrow is doing automatic flipping between
// K_ARROW_LEFT and K_ARROW_RIGHT, depending on the text direction. To get the
// same effect with an image, use the icon names “pan-start-symbolic“ and
// “pan-end-symbolic“, which react to the text direction.
type Arrow struct {
	Misc
}

var (
	_ Arrowwer        = (*Arrow)(nil)
	_ gextras.Nativer = (*Arrow)(nil)
)

func wrapArrow(obj *externglib.Object) Arrowwer {
	return &Arrow{
		Misc: Misc{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
	}
}

func marshalArrowwer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapArrow(obj), nil
}

func (*Arrow) privateArrow() {}
