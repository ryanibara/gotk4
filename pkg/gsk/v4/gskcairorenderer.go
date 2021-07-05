// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_cairo_renderer_get_type()), F: marshalCairoRenderer},
	})
}

// CairoRenderer: GSK renderer that is using cairo.
//
// Since it is using cairo, this renderer cannot support 3D transformations.
type CairoRenderer interface {
	Renderer
}

// cairoRenderer implements the CairoRenderer class.
type cairoRenderer struct {
	Renderer
}

// WrapCairoRenderer wraps a GObject to the right type. It is
// primarily used internally.
func WrapCairoRenderer(obj *externglib.Object) CairoRenderer {
	return cairoRenderer{
		Renderer: WrapRenderer(obj),
	}
}

func marshalCairoRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCairoRenderer(obj), nil
}

// NewCairoRenderer creates a new Cairo renderer.
//
// The Cairo renderer is the fallback renderer drawing in ways similar to how
// GTK 3 drew its content. Its primary use is as comparison tool.
//
// The Cairo renderer is incomplete. It cannot render 3D transformed content and
// will instead render an error marker. Its usage should be avoided.
func NewCairoRenderer() CairoRenderer {
	var _cret *C.GskRenderer // in

	_cret = C.gsk_cairo_renderer_new()

	var _cairoRenderer CairoRenderer // out

	_cairoRenderer = WrapCairoRenderer(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cairoRenderer
}
