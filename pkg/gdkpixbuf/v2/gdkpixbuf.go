// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-pixbuf-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_pixbuf_simple_anim_iter_get_type()), F: marshalPixbufSimpleAnimIterer},
	})
}

// PixbufSimpleAnimIterer describes PixbufSimpleAnimIter's methods.
type PixbufSimpleAnimIterer interface {
	privatePixbufSimpleAnimIter()
}

type PixbufSimpleAnimIter struct {
	PixbufAnimationIter
}

var (
	_ PixbufSimpleAnimIterer = (*PixbufSimpleAnimIter)(nil)
	_ gextras.Nativer        = (*PixbufSimpleAnimIter)(nil)
)

func wrapPixbufSimpleAnimIter(obj *externglib.Object) *PixbufSimpleAnimIter {
	return &PixbufSimpleAnimIter{
		PixbufAnimationIter: PixbufAnimationIter{
			Object: obj,
		},
	}
}

func marshalPixbufSimpleAnimIterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPixbufSimpleAnimIter(obj), nil
}

func (*PixbufSimpleAnimIter) privatePixbufSimpleAnimIter() {}
