// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gsk_cairo_renderer_get_type()), F: marshalCairoRendererer},
	})
}

// CairoRendererer describes CairoRenderer's methods.
type CairoRendererer interface {
	privateCairoRenderer()
}

// CairoRenderer: GSK renderer that is using cairo.
//
// Since it is using cairo, this renderer cannot support 3D transformations.
type CairoRenderer struct {
	Renderer
}

var (
	_ CairoRendererer = (*CairoRenderer)(nil)
	_ gextras.Nativer = (*CairoRenderer)(nil)
)

func wrapCairoRenderer(obj *externglib.Object) CairoRendererer {
	return &CairoRenderer{
		Renderer: Renderer{
			Object: obj,
		},
	}
}

func marshalCairoRendererer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCairoRenderer(obj), nil
}

// NewCairoRenderer creates a new Cairo renderer.
//
// The Cairo renderer is the fallback renderer drawing in ways similar to how
// GTK 3 drew its content. Its primary use is as comparison tool.
//
// The Cairo renderer is incomplete. It cannot render 3D transformed content and
// will instead render an error marker. Its usage should be avoided.
func NewCairoRenderer() *CairoRenderer {
	var _cret *C.GskRenderer // in

	_cret = C.gsk_cairo_renderer_new()

	var _cairoRenderer *CairoRenderer // out

	_cairoRenderer = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*CairoRenderer)

	return _cairoRenderer
}

func (*CairoRenderer) privateCairoRenderer() {}
