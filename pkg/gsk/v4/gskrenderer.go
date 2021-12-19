// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/graphene"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_renderer_get_type()), F: marshalRendererer},
	})
}

// RendererOverrider contains methods that are overridable.
type RendererOverrider interface {
}

// Renderer: GskRenderer is a class that renders a scene graph defined via a
// tree of gsk.RenderNode instances.
//
// Typically you will use a GskRenderer instance to repeatedly call
// gsk.Renderer.Render() to update the contents of its associated gdk.Surface.
//
// It is necessary to realize a GskRenderer instance using
// gsk.Renderer.Realize() before calling gsk.Renderer.Render(), in order to
// create the appropriate windowing system resources needed to render the scene.
type Renderer struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Renderer)(nil)
)

// Rendererer describes types inherited from class Renderer.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Rendererer interface {
	externglib.Objector
	baseRenderer() *Renderer
}

var _ Rendererer = (*Renderer)(nil)

func classInitRendererer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRenderer(obj *externglib.Object) *Renderer {
	return &Renderer{
		Object: obj,
	}
}

func marshalRendererer(p uintptr) (interface{}, error) {
	return wrapRenderer(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (renderer *Renderer) baseRenderer() *Renderer {
	return renderer
}

// BaseRenderer returns the underlying base object.
func BaseRenderer(obj Rendererer) *Renderer {
	return obj.baseRenderer()
}

// NewRendererForSurface creates an appropriate GskRenderer instance for the
// given surface.
//
// If the GSK_RENDERER environment variable is set, GSK will try that renderer
// first, before trying the backend-specific default. The ultimate fallback is
// the cairo renderer.
//
// The renderer will be realized before it is returned.
//
// The function takes the following parameters:
//
//    - surface: GdkSurface.
//
// The function returns the following values:
//
//    - renderer (optional): GskRenderer.
//
func NewRendererForSurface(surface gdk.Surfacer) *Renderer {
	var _arg1 *C.GdkSurface  // out
	var _cret *C.GskRenderer // in

	_arg1 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gsk_renderer_new_for_surface(_arg1)
	runtime.KeepAlive(surface)

	var _renderer *Renderer // out

	if _cret != nil {
		_renderer = wrapRenderer(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _renderer
}

// Surface retrieves the GdkSurface set using gsk_enderer_realize().
//
// If the renderer has not been realized yet, NULL will be returned.
//
// The function returns the following values:
//
//    - surface (optional): GdkSurface.
//
func (renderer *Renderer) Surface() gdk.Surfacer {
	var _arg0 *C.GskRenderer // out
	var _cret *C.GdkSurface  // in

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(renderer.Native()))

	_cret = C.gsk_renderer_get_surface(_arg0)
	runtime.KeepAlive(renderer)

	var _surface gdk.Surfacer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gdk.Surfacer)
				return ok
			})
			rv, ok := casted.(gdk.Surfacer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Surfacer")
			}
			_surface = rv
		}
	}

	return _surface
}

// IsRealized checks whether the renderer is realized or not.
//
// The function returns the following values:
//
//    - ok: TRUE if the GskRenderer was realized, and FALSE otherwise.
//
func (renderer *Renderer) IsRealized() bool {
	var _arg0 *C.GskRenderer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(renderer.Native()))

	_cret = C.gsk_renderer_is_realized(_arg0)
	runtime.KeepAlive(renderer)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Realize creates the resources needed by the renderer to render the scene
// graph.
//
// The function takes the following parameters:
//
//    - surface: GdkSurface renderer will be used on.
//
func (renderer *Renderer) Realize(surface gdk.Surfacer) error {
	var _arg0 *C.GskRenderer // out
	var _arg1 *C.GdkSurface  // out
	var _cerr *C.GError      // in

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(renderer.Native()))
	_arg1 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	C.gsk_renderer_realize(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(surface)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Render renders the scene graph, described by a tree of GskRenderNode
// instances, ensuring that the given region gets redrawn.
//
// Renderers must ensure that changes of the contents given by the root node as
// well as the area given by region are redrawn. They are however free to not
// redraw any pixel outside of region if they can guarantee that it didn't
// change.
//
// The renderer will acquire a reference on the GskRenderNode tree while the
// rendering is in progress.
//
// The function takes the following parameters:
//
//    - root: GskRenderNode.
//    - region (optional): cairo_region_t that must be redrawn or NULL for the
//      whole window.
//
func (renderer *Renderer) Render(root RenderNoder, region *cairo.Region) {
	var _arg0 *C.GskRenderer    // out
	var _arg1 *C.GskRenderNode  // out
	var _arg2 *C.cairo_region_t // out

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(renderer.Native()))
	_arg1 = (*C.GskRenderNode)(unsafe.Pointer(root.Native()))
	if region != nil {
		_arg2 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))
	}

	C.gsk_renderer_render(_arg0, _arg1, _arg2)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(root)
	runtime.KeepAlive(region)
}

// RenderTexture renders the scene graph, described by a tree of GskRenderNode
// instances, to a GdkTexture.
//
// The renderer will acquire a reference on the GskRenderNode tree while the
// rendering is in progress.
//
// If you want to apply any transformations to root, you should put it into a
// transform node and pass that node instead.
//
// The function takes the following parameters:
//
//    - root: GskRenderNode.
//    - viewport (optional): section to draw or NULL to use root's bounds.
//
// The function returns the following values:
//
//    - texture: GdkTexture with the rendered contents of root.
//
func (renderer *Renderer) RenderTexture(root RenderNoder, viewport *graphene.Rect) gdk.Texturer {
	var _arg0 *C.GskRenderer     // out
	var _arg1 *C.GskRenderNode   // out
	var _arg2 *C.graphene_rect_t // out
	var _cret *C.GdkTexture      // in

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(renderer.Native()))
	_arg1 = (*C.GskRenderNode)(unsafe.Pointer(root.Native()))
	if viewport != nil {
		_arg2 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(viewport)))
	}

	_cret = C.gsk_renderer_render_texture(_arg0, _arg1, _arg2)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(root)
	runtime.KeepAlive(viewport)

	var _texture gdk.Texturer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Texturer is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gdk.Texturer)
			return ok
		})
		rv, ok := casted.(gdk.Texturer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Texturer")
		}
		_texture = rv
	}

	return _texture
}

// Unrealize releases all the resources created by gsk_renderer_realize().
func (renderer *Renderer) Unrealize() {
	var _arg0 *C.GskRenderer // out

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(renderer.Native()))

	C.gsk_renderer_unrealize(_arg0)
	runtime.KeepAlive(renderer)
}
