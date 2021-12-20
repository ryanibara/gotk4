// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_rect_get_type()), F: marshalRect},
	})
}

// Rect: location and size of a rectangle region.
//
// The width and height of a #graphene_rect_t can be negative; for instance, a
// #graphene_rect_t with an origin of [ 0, 0 ] and a size of [ 10, 10 ] is
// equivalent to a #graphene_rect_t with an origin of [ 10, 10 ] and a size of [
// -10, -10 ].
//
// Application code can normalize rectangles using graphene_rect_normalize();
// this function will ensure that the width and height of a rectangle are
// positive values. All functions taking a #graphene_rect_t as an argument will
// internally operate on a normalized copy; all functions returning a
// #graphene_rect_t will always return a normalized rectangle.
//
// An instance of this type is always passed by reference.
type Rect struct {
	*rect
}

// rect is the struct that's finalized.
type rect struct {
	native *C.graphene_rect_t
}

func marshalRect(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Rect{&rect{(*C.graphene_rect_t)(b)}}, nil
}

// Origin coordinates of the origin of the rectangle.
func (r *Rect) Origin() *Point {
	var v *Point // out
	v = (*Point)(gextras.NewStructNative(unsafe.Pointer((&r.native.origin))))
	return v
}

// Size: size of the rectangle.
func (r *Rect) Size() *Size {
	var v *Size // out
	v = (*Size)(gextras.NewStructNative(unsafe.Pointer((&r.native.size))))
	return v
}

// ContainsPoint checks whether a #graphene_rect_t contains the given
// coordinates.
//
// The function takes the following parameters:
//
//    - p: #graphene_point_t.
//
// The function returns the following values:
//
//    - ok: true if the rectangle contains the point.
//
func (r *Rect) ContainsPoint(p *Point) bool {
	var _arg0 *C.graphene_rect_t  // out
	var _arg1 *C.graphene_point_t // out
	var _cret C._Bool             // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(p)))

	_cret = C.graphene_rect_contains_point(_arg0, _arg1)
	runtime.KeepAlive(r)
	runtime.KeepAlive(p)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// ContainsRect checks whether a #graphene_rect_t fully contains the given
// rectangle.
//
// The function takes the following parameters:
//
//    - b: #graphene_rect_t.
//
// The function returns the following values:
//
//    - ok: true if the rectangle a fully contains b.
//
func (a *Rect) ContainsRect(b *Rect) bool {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 *C.graphene_rect_t // out
	var _cret C._Bool            // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_rect_contains_rect(_arg0, _arg1)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Equal checks whether the two given rectangle are equal.
//
// The function takes the following parameters:
//
//    - b: #graphene_rect_t.
//
// The function returns the following values:
//
//    - ok: true if the rectangles are equal.
//
func (a *Rect) Equal(b *Rect) bool {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 *C.graphene_rect_t // out
	var _cret C._Bool            // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_rect_equal(_arg0, _arg1)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Expand expands a #graphene_rect_t to contain the given #graphene_point_t.
//
// The function takes the following parameters:
//
//    - p: #graphene_point_t.
//
// The function returns the following values:
//
//    - res: return location for the expanded rectangle.
//
func (r *Rect) Expand(p *Point) *Rect {
	var _arg0 *C.graphene_rect_t  // out
	var _arg1 *C.graphene_point_t // out
	var _arg2 C.graphene_rect_t   // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(p)))

	C.graphene_rect_expand(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(r)
	runtime.KeepAlive(p)

	var _res *Rect // out

	_res = (*Rect)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Area: compute the area of given normalized rectangle.
//
// The function returns the following values:
//
//    - gfloat: area of the normalized rectangle.
//
func (r *Rect) Area() float32 {
	var _arg0 *C.graphene_rect_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	_cret = C.graphene_rect_get_area(_arg0)
	runtime.KeepAlive(r)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// BottomLeft retrieves the coordinates of the bottom-left corner of the given
// rectangle.
//
// The function returns the following values:
//
//    - p: return location for a #graphene_point_t.
//
func (r *Rect) BottomLeft() *Point {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 C.graphene_point_t // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	C.graphene_rect_get_bottom_left(_arg0, &_arg1)
	runtime.KeepAlive(r)

	var _p *Point // out

	_p = (*Point)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _p
}

// BottomRight retrieves the coordinates of the bottom-right corner of the given
// rectangle.
//
// The function returns the following values:
//
//    - p: return location for a #graphene_point_t.
//
func (r *Rect) BottomRight() *Point {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 C.graphene_point_t // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	C.graphene_rect_get_bottom_right(_arg0, &_arg1)
	runtime.KeepAlive(r)

	var _p *Point // out

	_p = (*Point)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _p
}

// Center retrieves the coordinates of the center of the given rectangle.
//
// The function returns the following values:
//
//    - p: return location for a #graphene_point_t.
//
func (r *Rect) Center() *Point {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 C.graphene_point_t // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	C.graphene_rect_get_center(_arg0, &_arg1)
	runtime.KeepAlive(r)

	var _p *Point // out

	_p = (*Point)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _p
}

// Height retrieves the normalized height of the given rectangle.
//
// The function returns the following values:
//
//    - gfloat: normalized height of the rectangle.
//
func (r *Rect) Height() float32 {
	var _arg0 *C.graphene_rect_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	_cret = C.graphene_rect_get_height(_arg0)
	runtime.KeepAlive(r)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// TopLeft retrieves the coordinates of the top-left corner of the given
// rectangle.
//
// The function returns the following values:
//
//    - p: return location for a #graphene_point_t.
//
func (r *Rect) TopLeft() *Point {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 C.graphene_point_t // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	C.graphene_rect_get_top_left(_arg0, &_arg1)
	runtime.KeepAlive(r)

	var _p *Point // out

	_p = (*Point)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _p
}

// TopRight retrieves the coordinates of the top-right corner of the given
// rectangle.
//
// The function returns the following values:
//
//    - p: return location for a #graphene_point_t.
//
func (r *Rect) TopRight() *Point {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 C.graphene_point_t // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	C.graphene_rect_get_top_right(_arg0, &_arg1)
	runtime.KeepAlive(r)

	var _p *Point // out

	_p = (*Point)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _p
}

// Vertices computes the four vertices of a #graphene_rect_t.
//
// The function returns the following values:
//
//    - vertices: return location for an array of 4 #graphene_vec2_t.
//
func (r *Rect) Vertices() [4]Vec2 {
	var _arg0 *C.graphene_rect_t   // out
	var _arg1 [4]C.graphene_vec2_t // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	C.graphene_rect_get_vertices(_arg0, &_arg1[0])
	runtime.KeepAlive(r)

	var _vertices [4]Vec2 // out

	{
		src := &_arg1
		for i := 0; i < 4; i++ {
			_vertices[i] = *(*Vec2)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
		}
	}

	return _vertices
}

// Width retrieves the normalized width of the given rectangle.
//
// The function returns the following values:
//
//    - gfloat: normalized width of the rectangle.
//
func (r *Rect) Width() float32 {
	var _arg0 *C.graphene_rect_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	_cret = C.graphene_rect_get_width(_arg0)
	runtime.KeepAlive(r)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// X retrieves the normalized X coordinate of the origin of the given rectangle.
//
// The function returns the following values:
//
//    - gfloat: normalized X coordinate of the rectangle.
//
func (r *Rect) X() float32 {
	var _arg0 *C.graphene_rect_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	_cret = C.graphene_rect_get_x(_arg0)
	runtime.KeepAlive(r)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Y retrieves the normalized Y coordinate of the origin of the given rectangle.
//
// The function returns the following values:
//
//    - gfloat: normalized Y coordinate of the rectangle.
//
func (r *Rect) Y() float32 {
	var _arg0 *C.graphene_rect_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	_cret = C.graphene_rect_get_y(_arg0)
	runtime.KeepAlive(r)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Init initializes the given #graphene_rect_t with the given values.
//
// This function will implicitly normalize the #graphene_rect_t before
// returning.
//
// The function takes the following parameters:
//
//    - x: x coordinate of the graphene_rect_t.origin.
//    - y: y coordinate of the graphene_rect_t.origin.
//    - width of the graphene_rect_t.size.
//    - height of the graphene_rect_t.size.
//
// The function returns the following values:
//
//    - rect: initialized rectangle.
//
func (r *Rect) Init(x float32, y float32, width float32, height float32) *Rect {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out
	var _arg3 C.float            // out
	var _arg4 C.float            // out
	var _cret *C.graphene_rect_t // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = C.float(x)
	_arg2 = C.float(y)
	_arg3 = C.float(width)
	_arg4 = C.float(height)

	_cret = C.graphene_rect_init(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(r)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _rect *Rect // out

	_rect = (*Rect)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _rect
}

// InitFromRect initializes r using the given src rectangle.
//
// This function will implicitly normalize the #graphene_rect_t before
// returning.
//
// The function takes the following parameters:
//
//    - src: #graphene_rect_t.
//
// The function returns the following values:
//
//    - rect: initialized rectangle.
//
func (r *Rect) InitFromRect(src *Rect) *Rect {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 *C.graphene_rect_t // out
	var _cret *C.graphene_rect_t // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.graphene_rect_init_from_rect(_arg0, _arg1)
	runtime.KeepAlive(r)
	runtime.KeepAlive(src)

	var _rect *Rect // out

	_rect = (*Rect)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _rect
}

// Inset changes the given rectangle to be smaller, or larger depending on the
// given inset parameters.
//
// To create an inset rectangle, use positive d_x or d_y values; to create a
// larger, encompassing rectangle, use negative d_x or d_y values.
//
// The origin of the rectangle is offset by d_x and d_y, while the size is
// adjusted by (2 * d_x, 2 * d_y). If d_x and d_y are positive values, the size
// of the rectangle is decreased; if d_x and d_y are negative values, the size
// of the rectangle is increased.
//
// If the size of the resulting inset rectangle has a negative width or height
// then the size will be set to zero.
//
// The function takes the following parameters:
//
//    - dX: horizontal inset.
//    - dY: vertical inset.
//
// The function returns the following values:
//
//    - rect: inset rectangle.
//
func (r *Rect) Inset(dX float32, dY float32) *Rect {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out
	var _cret *C.graphene_rect_t // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = C.float(dX)
	_arg2 = C.float(dY)

	_cret = C.graphene_rect_inset(_arg0, _arg1, _arg2)
	runtime.KeepAlive(r)
	runtime.KeepAlive(dX)
	runtime.KeepAlive(dY)

	var _rect *Rect // out

	_rect = (*Rect)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _rect
}

// InsetR changes the given rectangle to be smaller, or larger depending on the
// given inset parameters.
//
// To create an inset rectangle, use positive d_x or d_y values; to create a
// larger, encompassing rectangle, use negative d_x or d_y values.
//
// The origin of the rectangle is offset by d_x and d_y, while the size is
// adjusted by (2 * d_x, 2 * d_y). If d_x and d_y are positive values, the size
// of the rectangle is decreased; if d_x and d_y are negative values, the size
// of the rectangle is increased.
//
// If the size of the resulting inset rectangle has a negative width or height
// then the size will be set to zero.
//
// The function takes the following parameters:
//
//    - dX: horizontal inset.
//    - dY: vertical inset.
//
// The function returns the following values:
//
//    - res: return location for the inset rectangle.
//
func (r *Rect) InsetR(dX float32, dY float32) *Rect {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out
	var _arg3 C.graphene_rect_t  // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = C.float(dX)
	_arg2 = C.float(dY)

	C.graphene_rect_inset_r(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(r)
	runtime.KeepAlive(dX)
	runtime.KeepAlive(dY)

	var _res *Rect // out

	_res = (*Rect)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _res
}

// Interpolate: linearly interpolates the origin and size of the two given
// rectangles.
//
// The function takes the following parameters:
//
//    - b: #graphene_rect_t.
//    - factor: linear interpolation factor.
//
// The function returns the following values:
//
//    - res: return location for the interpolated rectangle.
//
func (a *Rect) Interpolate(b *Rect, factor float64) *Rect {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 *C.graphene_rect_t // out
	var _arg2 C.double           // out
	var _arg3 C.graphene_rect_t  // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(b)))
	_arg2 = C.double(factor)

	C.graphene_rect_interpolate(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
	runtime.KeepAlive(factor)

	var _res *Rect // out

	_res = (*Rect)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _res
}

// Intersection computes the intersection of the two given rectangles.
//
// ! (rectangle-intersection.png)
//
// The intersection in the image above is the blue outline.
//
// If the two rectangles do not intersect, res will contain a degenerate
// rectangle with origin in (0, 0) and a size of 0.
//
// The function takes the following parameters:
//
//    - b: #graphene_rect_t.
//
// The function returns the following values:
//
//    - res (optional): return location for a #graphene_rect_t.
//    - ok: true if the two rectangles intersect.
//
func (a *Rect) Intersection(b *Rect) (*Rect, bool) {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 *C.graphene_rect_t // out
	var _arg2 C.graphene_rect_t  // in
	var _cret C._Bool            // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_rect_intersection(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res *Rect // out
	var _ok bool   // out

	_res = (*Rect)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret {
		_ok = true
	}

	return _res, _ok
}

// Normalize normalizes the passed rectangle.
//
// This function ensures that the size of the rectangle is made of positive
// values, and that the origin is the top-left corner of the rectangle.
//
// The function returns the following values:
//
//    - rect: normalized rectangle.
//
func (r *Rect) Normalize() *Rect {
	var _arg0 *C.graphene_rect_t // out
	var _cret *C.graphene_rect_t // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	_cret = C.graphene_rect_normalize(_arg0)
	runtime.KeepAlive(r)

	var _rect *Rect // out

	_rect = (*Rect)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _rect
}

// NormalizeR normalizes the passed rectangle.
//
// This function ensures that the size of the rectangle is made of positive
// values, and that the origin is in the top-left corner of the rectangle.
//
// The function returns the following values:
//
//    - res: return location for the normalized rectangle.
//
func (r *Rect) NormalizeR() *Rect {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 C.graphene_rect_t  // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	C.graphene_rect_normalize_r(_arg0, &_arg1)
	runtime.KeepAlive(r)

	var _res *Rect // out

	_res = (*Rect)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// Offset offsets the origin by d_x and d_y.
//
// The size of the rectangle is unchanged.
//
// The function takes the following parameters:
//
//    - dX: horizontal offset.
//    - dY: vertical offset.
//
// The function returns the following values:
//
//    - rect: offset rectangle.
//
func (r *Rect) Offset(dX float32, dY float32) *Rect {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out
	var _cret *C.graphene_rect_t // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = C.float(dX)
	_arg2 = C.float(dY)

	_cret = C.graphene_rect_offset(_arg0, _arg1, _arg2)
	runtime.KeepAlive(r)
	runtime.KeepAlive(dX)
	runtime.KeepAlive(dY)

	var _rect *Rect // out

	_rect = (*Rect)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _rect
}

// OffsetR offsets the origin of the given rectangle by d_x and d_y.
//
// The size of the rectangle is left unchanged.
//
// The function takes the following parameters:
//
//    - dX: horizontal offset.
//    - dY: vertical offset.
//
// The function returns the following values:
//
//    - res: return location for the offset rectangle.
//
func (r *Rect) OffsetR(dX float32, dY float32) *Rect {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out
	var _arg3 C.graphene_rect_t  // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = C.float(dX)
	_arg2 = C.float(dY)

	C.graphene_rect_offset_r(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(r)
	runtime.KeepAlive(dX)
	runtime.KeepAlive(dY)

	var _res *Rect // out

	_res = (*Rect)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _res
}

// Round rounds the origin and size of the given rectangle to their nearest
// integer values; the rounding is guaranteed to be large enough to have an area
// bigger or equal to the original rectangle, but might not fully contain its
// extents. Use graphene_rect_round_extents() in case you need to round to a
// rectangle that covers fully the original one.
//
// This function is the equivalent of calling floor on the coordinates of the
// origin, and ceil on the size.
//
// Deprecated: Use graphene_rect_round_extents() instead.
//
// The function returns the following values:
//
//    - res: return location for the rounded rectangle.
//
func (r *Rect) Round() *Rect {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 C.graphene_rect_t  // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	C.graphene_rect_round(_arg0, &_arg1)
	runtime.KeepAlive(r)

	var _res *Rect // out

	_res = (*Rect)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// RoundExtents rounds the origin of the given rectangle to its nearest integer
// value and and recompute the size so that the rectangle is large enough to
// contain all the conrners of the original rectangle.
//
// This function is the equivalent of calling floor on the coordinates of the
// origin, and recomputing the size calling ceil on the bottom-right
// coordinates.
//
// If you want to be sure that the rounded rectangle completely covers the area
// that was covered by the original rectangle — i.e. you want to cover the area
// including all its corners — this function will make sure that the size is
// recomputed taking into account the ceiling of the coordinates of the
// bottom-right corner. If the difference between the original coordinates and
// the coordinates of the rounded rectangle is greater than the difference
// between the original size and and the rounded size, then the move of the
// origin would not be compensated by a move in the anti-origin, leaving the
// corners of the original rectangle outside the rounded one.
//
// The function returns the following values:
//
//    - res: return location for the rectangle with rounded extents.
//
func (r *Rect) RoundExtents() *Rect {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 C.graphene_rect_t  // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	C.graphene_rect_round_extents(_arg0, &_arg1)
	runtime.KeepAlive(r)

	var _res *Rect // out

	_res = (*Rect)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// RoundToPixel rounds the origin and the size of the given rectangle to their
// nearest integer values; the rounding is guaranteed to be large enough to
// contain the original rectangle.
//
// Deprecated: Use graphene_rect_round() instead.
//
// The function returns the following values:
//
//    - rect: pixel-aligned rectangle.
//
func (r *Rect) RoundToPixel() *Rect {
	var _arg0 *C.graphene_rect_t // out
	var _cret *C.graphene_rect_t // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	_cret = C.graphene_rect_round_to_pixel(_arg0)
	runtime.KeepAlive(r)

	var _rect *Rect // out

	_rect = (*Rect)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _rect
}

// Scale scales the size and origin of a rectangle horizontaly by s_h, and
// vertically by s_v. The result res is normalized.
//
// The function takes the following parameters:
//
//    - sH: horizontal scale factor.
//    - sV: vertical scale factor.
//
// The function returns the following values:
//
//    - res: return location for the scaled rectangle.
//
func (r *Rect) Scale(sH float32, sV float32) *Rect {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out
	var _arg3 C.graphene_rect_t  // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = C.float(sH)
	_arg2 = C.float(sV)

	C.graphene_rect_scale(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(r)
	runtime.KeepAlive(sH)
	runtime.KeepAlive(sV)

	var _res *Rect // out

	_res = (*Rect)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _res
}

// Union computes the union of the two given rectangles.
//
// ! (rectangle-union.png)
//
// The union in the image above is the blue outline.
//
// The function takes the following parameters:
//
//    - b: #graphene_rect_t.
//
// The function returns the following values:
//
//    - res: return location for a #graphene_rect_t.
//
func (a *Rect) Union(b *Rect) *Rect {
	var _arg0 *C.graphene_rect_t // out
	var _arg1 *C.graphene_rect_t // out
	var _arg2 C.graphene_rect_t  // in

	_arg0 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_rect_union(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res *Rect // out

	_res = (*Rect)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// RectAlloc allocates a new #graphene_rect_t.
//
// The contents of the returned rectangle are undefined.
//
// The function returns the following values:
//
//    - rect: newly allocated rectangle.
//
func RectAlloc() *Rect {
	var _cret *C.graphene_rect_t // in

	_cret = C.graphene_rect_alloc()

	var _rect *Rect // out

	_rect = (*Rect)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_rect)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.graphene_rect_free((*C.graphene_rect_t)(intern.C))
		},
	)

	return _rect
}

// RectZero returns a degenerate rectangle with origin fixed at (0, 0) and a
// size of 0, 0.
//
// The function returns the following values:
//
//    - rect: fixed rectangle.
//
func RectZero() *Rect {
	var _cret *C.graphene_rect_t // in

	_cret = C.graphene_rect_zero()

	var _rect *Rect // out

	_rect = (*Rect)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _rect
}
