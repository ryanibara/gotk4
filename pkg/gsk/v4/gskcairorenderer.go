// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

// GType values.
var (
	GTypeCairoRenderer = coreglib.Type(C.gsk_cairo_renderer_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCairoRenderer, F: marshalCairoRenderer},
	})
}

// CairoRendererOverrider contains methods that are overridable.
type CairoRendererOverrider interface {
}

// CairoRenderer: GSK renderer that is using cairo.
//
// Since it is using cairo, this renderer cannot support 3D transformations.
type CairoRenderer struct {
	_ [0]func() // equal guard
	Renderer
}

var (
	_ Rendererer = (*CairoRenderer)(nil)
)

func initClassCairoRenderer(gclass unsafe.Pointer, goval any) {
}

func wrapCairoRenderer(obj *coreglib.Object) *CairoRenderer {
	return &CairoRenderer{
		Renderer: Renderer{
			Object: obj,
		},
	}
}

func marshalCairoRenderer(p uintptr) (interface{}, error) {
	return wrapCairoRenderer(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCairoRenderer creates a new Cairo renderer.
//
// The Cairo renderer is the fallback renderer drawing in ways similar to how
// GTK 3 drew its content. Its primary use is as comparison tool.
//
// The Cairo renderer is incomplete. It cannot render 3D transformed content and
// will instead render an error marker. Its usage should be avoided.
//
// The function returns the following values:
//
//    - cairoRenderer: new Cairo renderer.
//
func NewCairoRenderer() *CairoRenderer {
	var _cret *C.GskRenderer // in

	_cret = C.gsk_cairo_renderer_new()

	var _cairoRenderer *CairoRenderer // out

	_cairoRenderer = wrapCairoRenderer(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cairoRenderer
}
