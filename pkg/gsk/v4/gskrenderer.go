// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/graphene"
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
		{T: externglib.Type(C.gsk_renderer_get_type()), F: marshalRenderer},
	})
}

// Renderer: `GskRenderer` is a class that renders a scene graph defined via a
// tree of [class@Gsk.RenderNode] instances.
//
// Typically you will use a `GskRenderer` instance to repeatedly call
// [method@Gsk.Renderer.render] to update the contents of its associated
// [class@Gdk.Surface].
//
// It is necessary to realize a `GskRenderer` instance using
// [method@Gsk.Renderer.realize] before calling [method@Gsk.Renderer.render], in
// order to create the appropriate windowing system resources needed to render
// the scene.
type Renderer interface {
	gextras.Objector

	Surface() gdk.Surface

	IsRealizedRenderer() bool

	RealizeRenderer(surface gdk.Surface) error

	RenderRenderer(root RenderNode, region *cairo.Region)

	RenderTextureRenderer(root RenderNode, viewport *graphene.Rect) gdk.Texture

	UnrealizeRenderer()
}

// renderer implements the Renderer class.
type renderer struct {
	gextras.Objector
}

// WrapRenderer wraps a GObject to the right type. It is
// primarily used internally.
func WrapRenderer(obj *externglib.Object) Renderer {
	return renderer{
		Objector: obj,
	}
}

func marshalRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRenderer(obj), nil
}

func NewRendererForSurface(surface gdk.Surface) Renderer {
	var _arg1 *C.GdkSurface  // out
	var _cret *C.GskRenderer // in

	_arg1 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gsk_renderer_new_for_surface(_arg1)

	var _renderer Renderer // out

	_renderer = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Renderer)

	return _renderer
}

func (r renderer) Surface() gdk.Surface {
	var _arg0 *C.GskRenderer // out
	var _cret *C.GdkSurface  // in

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(r.Native()))

	_cret = C.gsk_renderer_get_surface(_arg0)

	var _surface gdk.Surface // out

	_surface = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Surface)

	return _surface
}

func (r renderer) IsRealizedRenderer() bool {
	var _arg0 *C.GskRenderer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(r.Native()))

	_cret = C.gsk_renderer_is_realized(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (r renderer) RealizeRenderer(surface gdk.Surface) error {
	var _arg0 *C.GskRenderer // out
	var _arg1 *C.GdkSurface  // out
	var _cerr *C.GError      // in

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	C.gsk_renderer_realize(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (r renderer) RenderRenderer(root RenderNode, region *cairo.Region) {
	var _arg0 *C.GskRenderer    // out
	var _arg1 *C.GskRenderNode  // out
	var _arg2 *C.cairo_region_t // out

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GskRenderNode)(unsafe.Pointer(root.Native()))
	_arg2 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))

	C.gsk_renderer_render(_arg0, _arg1, _arg2)
}

func (r renderer) RenderTextureRenderer(root RenderNode, viewport *graphene.Rect) gdk.Texture {
	var _arg0 *C.GskRenderer     // out
	var _arg1 *C.GskRenderNode   // out
	var _arg2 *C.graphene_rect_t // out
	var _cret *C.GdkTexture      // in

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GskRenderNode)(unsafe.Pointer(root.Native()))
	_arg2 = (*C.graphene_rect_t)(unsafe.Pointer(viewport.Native()))

	_cret = C.gsk_renderer_render_texture(_arg0, _arg1, _arg2)

	var _texture gdk.Texture // out

	_texture = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(gdk.Texture)

	return _texture
}

func (r renderer) UnrealizeRenderer() {
	var _arg0 *C.GskRenderer // out

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(r.Native()))

	C.gsk_renderer_unrealize(_arg0)
}
