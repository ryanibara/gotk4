// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/graphene"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_renderer_get_type()), F: marshalRenderer},
	})
}

// Renderer: base type for the object managing the rendering pipeline for a
// Surface.
type Renderer interface {
	gextras.Objector

	// Surface retrieves the Surface set using gsk_renderer_realize(). If the
	// renderer has not been realized yet, nil will be returned.
	Surface() gdk.Surface
	// IsRealized checks whether the @renderer is realized or not.
	IsRealized() bool
	// Realize creates the resources needed by the @renderer to render the scene
	// graph.
	Realize(surface gdk.Surface) error
	// Render renders the scene graph, described by a tree of RenderNode
	// instances, ensuring that the given @region gets redrawn.
	//
	// Renderers must ensure that changes of the contents given by the @root
	// node as well as the area given by @region are redrawn. They are however
	// free to not redraw any pixel outside of @region if they can guarantee
	// that it didn't change.
	//
	// The @renderer will acquire a reference on the RenderNode tree while the
	// rendering is in progress.
	Render(root RenderNode, region *cairo.Region)
	// RenderTexture renders the scene graph, described by a tree of RenderNode
	// instances, to a Texture.
	//
	// The @renderer will acquire a reference on the RenderNode tree while the
	// rendering is in progress.
	//
	// If you want to apply any transformations to @root, you should put it into
	// a transform node and pass that node instead.
	RenderTexture(root RenderNode, viewport *graphene.Rect) gdk.Texture
	// Unrealize releases all the resources created by gsk_renderer_realize().
	Unrealize()
}

// renderer implements the Renderer interface.
type renderer struct {
	gextras.Objector
}

var _ Renderer = (*renderer)(nil)

// WrapRenderer wraps a GObject to the right type. It is
// primarily used internally.
func WrapRenderer(obj *externglib.Object) Renderer {
	return Renderer{
		Objector: obj,
	}
}

func marshalRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRenderer(obj), nil
}

// NewRendererForSurface constructs a class Renderer.
func NewRendererForSurface(surface gdk.Surface) Renderer {
	var arg1 *C.GdkSurface

	arg1 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	var cret C.GskRenderer
	var ret1 Renderer

	cret = C.gsk_renderer_new_for_surface(surface)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Renderer)

	return ret1
}

// Surface retrieves the Surface set using gsk_renderer_realize(). If the
// renderer has not been realized yet, nil will be returned.
func (r renderer) Surface() gdk.Surface {
	var arg0 *C.GskRenderer

	arg0 = (*C.GskRenderer)(unsafe.Pointer(r.Native()))

	var cret *C.GdkSurface
	var ret1 gdk.Surface

	cret = C.gsk_renderer_get_surface(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Surface)

	return ret1
}

// IsRealized checks whether the @renderer is realized or not.
func (r renderer) IsRealized() bool {
	var arg0 *C.GskRenderer

	arg0 = (*C.GskRenderer)(unsafe.Pointer(r.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gsk_renderer_is_realized(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Realize creates the resources needed by the @renderer to render the scene
// graph.
func (r renderer) Realize(surface gdk.Surface) error {
	var arg0 *C.GskRenderer
	var arg1 *C.GdkSurface
	var errout *C.GError

	arg0 = (*C.GskRenderer)(unsafe.Pointer(r.Native()))
	arg1 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	var goerr error

	C.gsk_renderer_realize(arg0, surface, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// Render renders the scene graph, described by a tree of RenderNode
// instances, ensuring that the given @region gets redrawn.
//
// Renderers must ensure that changes of the contents given by the @root
// node as well as the area given by @region are redrawn. They are however
// free to not redraw any pixel outside of @region if they can guarantee
// that it didn't change.
//
// The @renderer will acquire a reference on the RenderNode tree while the
// rendering is in progress.
func (r renderer) Render(root RenderNode, region *cairo.Region) {
	var arg0 *C.GskRenderer
	var arg1 *C.GskRenderNode
	var arg2 *C.cairo_region_t

	arg0 = (*C.GskRenderer)(unsafe.Pointer(r.Native()))
	arg1 = (*C.GskRenderNode)(unsafe.Pointer(root.Native()))
	arg2 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))

	C.gsk_renderer_render(arg0, root, region)
}

// RenderTexture renders the scene graph, described by a tree of RenderNode
// instances, to a Texture.
//
// The @renderer will acquire a reference on the RenderNode tree while the
// rendering is in progress.
//
// If you want to apply any transformations to @root, you should put it into
// a transform node and pass that node instead.
func (r renderer) RenderTexture(root RenderNode, viewport *graphene.Rect) gdk.Texture {
	var arg0 *C.GskRenderer
	var arg1 *C.GskRenderNode
	var arg2 *C.graphene_rect_t

	arg0 = (*C.GskRenderer)(unsafe.Pointer(r.Native()))
	arg1 = (*C.GskRenderNode)(unsafe.Pointer(root.Native()))
	arg2 = (*C.graphene_rect_t)(unsafe.Pointer(viewport.Native()))

	var cret *C.GdkTexture
	var ret1 gdk.Texture

	cret = C.gsk_renderer_render_texture(arg0, root, viewport)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(gdk.Texture)

	return ret1
}

// Unrealize releases all the resources created by gsk_renderer_realize().
func (r renderer) Unrealize() {
	var arg0 *C.GskRenderer

	arg0 = (*C.GskRenderer)(unsafe.Pointer(r.Native()))

	C.gsk_renderer_unrealize(arg0)
}
