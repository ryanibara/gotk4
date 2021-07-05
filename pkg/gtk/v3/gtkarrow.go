// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_arrow_get_type()), F: marshalArrow},
	})
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
type Arrow interface {
	Misc

	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable

	// SetArrow sets the direction and style of the Arrow, @arrow.
	//
	// Deprecated: since version 3.14.
	SetArrow(arrowType ArrowType, shadowType ShadowType)
}

// arrow implements the Arrow class.
type arrow struct {
	Misc
}

// WrapArrow wraps a GObject to the right type. It is
// primarily used internally.
func WrapArrow(obj *externglib.Object) Arrow {
	return arrow{
		Misc: WrapMisc(obj),
	}
}

func marshalArrow(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapArrow(obj), nil
}

// NewArrow creates a new Arrow widget.
//
// Deprecated: since version 3.14.
func NewArrow(arrowType ArrowType, shadowType ShadowType) Arrow {
	var _arg1 C.GtkArrowType  // out
	var _arg2 C.GtkShadowType // out
	var _cret *C.GtkWidget    // in

	_arg1 = C.GtkArrowType(arrowType)
	_arg2 = C.GtkShadowType(shadowType)

	_cret = C.gtk_arrow_new(_arg1, _arg2)

	var _arrow Arrow // out

	_arrow = WrapArrow(externglib.Take(unsafe.Pointer(_cret)))

	return _arrow
}

func (a arrow) SetArrow(arrowType ArrowType, shadowType ShadowType) {
	var _arg0 *C.GtkArrow     // out
	var _arg1 C.GtkArrowType  // out
	var _arg2 C.GtkShadowType // out

	_arg0 = (*C.GtkArrow)(unsafe.Pointer(a.Native()))
	_arg1 = C.GtkArrowType(arrowType)
	_arg2 = C.GtkShadowType(shadowType)

	C.gtk_arrow_set(_arg0, _arg1, _arg2)
}

func (a arrow) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(a))
}
