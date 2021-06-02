// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: graphene-gobject-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_triangle_get_type()), F: marshalTriangle},
	})
}

// Triangle: a triangle.
type Triangle struct {
	native C.graphene_triangle_t
}

// WrapTriangle wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTriangle(ptr unsafe.Pointer) *Triangle {
	if ptr == nil {
		return nil
	}

	return (*Triangle)(ptr)
}

func marshalTriangle(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTriangle(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *Triangle) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// NewTriangleAlloc constructs a struct Triangle.
func NewTriangleAlloc() *Triangle {
	ret := C.graphene_triangle_alloc()

	var ret0 *Triangle

	{
		ret0 = WrapTriangle(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *Triangle) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// ContainsPoint checks whether the given triangle @t contains the point @p.
func (t *Triangle) ContainsPoint(p *Point3D) bool {
	var arg0 *C.graphene_triangle_t
	var arg1 *C.graphene_point3d_t

	arg0 = (*C.graphene_triangle_t)(t.Native())
	arg1 = (*C.graphene_point3d_t)(p.Native())

	ret := C.graphene_triangle_contains_point(arg0, arg1)

	var ret0 bool

	ret0 = C.bool(ret) != C.false

	return ret0
}

// Equal checks whether the two given #graphene_triangle_t are equal.
func (a *Triangle) Equal(b *Triangle) bool {
	var arg0 *C.graphene_triangle_t
	var arg1 *C.graphene_triangle_t

	arg0 = (*C.graphene_triangle_t)(a.Native())
	arg1 = (*C.graphene_triangle_t)(b.Native())

	ret := C.graphene_triangle_equal(arg0, arg1)

	var ret0 bool

	ret0 = C.bool(ret) != C.false

	return ret0
}

// Free frees the resources allocated by graphene_triangle_alloc().
func (t *Triangle) Free() {
	var arg0 *C.graphene_triangle_t

	arg0 = (*C.graphene_triangle_t)(t.Native())

	C.graphene_triangle_free(arg0)
}

// Area computes the area of the given #graphene_triangle_t.
func (t *Triangle) Area() float32 {
	var arg0 *C.graphene_triangle_t

	arg0 = (*C.graphene_triangle_t)(t.Native())

	ret := C.graphene_triangle_get_area(arg0)

	var ret0 float32

	ret0 = float32(ret)

	return ret0
}

// Barycoords computes the barycentric coordinates
// (http://en.wikipedia.org/wiki/Barycentric_coordinate_system) of the given
// point @p.
//
// The point @p must lie on the same plane as the triangle @t; if the point is
// not coplanar, the result of this function is undefined.
//
// If we place the origin in the coordinates of the triangle's A point, the
// barycentric coordinates are `u`, which is on the AC vector; and `v` which is
// on the AB vector:
//
// ! (triangle-barycentric.png)
//
// The returned #graphene_vec2_t contains the following values, in order:
//
//    - `res.x = u`
//    - `res.y = v`
func (t *Triangle) Barycoords(p *Point3D) (res Vec2, ok bool) {
	var arg0 *C.graphene_triangle_t
	var arg1 *C.graphene_point3d_t
	var arg2 *C.graphene_vec2_t // out

	arg0 = (*C.graphene_triangle_t)(t.Native())
	arg1 = (*C.graphene_point3d_t)(p.Native())

	ret := C.graphene_triangle_get_barycoords(arg0, arg1, &arg2)

	var ret0 *Vec2
	var ret1 bool

	{
		ret0 = WrapVec2(unsafe.Pointer(arg2))
	}

	ret1 = C.bool(ret) != C.false

	return ret0, ret1
}

// BoundingBox computes the bounding box of the given #graphene_triangle_t.
func (t *Triangle) BoundingBox() Box {
	var arg0 *C.graphene_triangle_t
	var arg1 *C.graphene_box_t // out

	arg0 = (*C.graphene_triangle_t)(t.Native())

	C.graphene_triangle_get_bounding_box(arg0, &arg1)

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(arg1))
	}

	return ret0
}

// Midpoint computes the coordinates of the midpoint of the given
// #graphene_triangle_t.
//
// The midpoint G is the centroid
// (https://en.wikipedia.org/wiki/Centroid#Triangle_centroid) of the triangle,
// i.e. the intersection of its medians.
func (t *Triangle) Midpoint() Point3D {
	var arg0 *C.graphene_triangle_t
	var arg1 *C.graphene_point3d_t // out

	arg0 = (*C.graphene_triangle_t)(t.Native())

	C.graphene_triangle_get_midpoint(arg0, &arg1)

	var ret0 *Point3D

	{
		ret0 = WrapPoint3D(unsafe.Pointer(arg1))
	}

	return ret0
}

// Normal computes the normal vector of the given #graphene_triangle_t.
func (t *Triangle) Normal() Vec3 {
	var arg0 *C.graphene_triangle_t
	var arg1 *C.graphene_vec3_t // out

	arg0 = (*C.graphene_triangle_t)(t.Native())

	C.graphene_triangle_get_normal(arg0, &arg1)

	var ret0 *Vec3

	{
		ret0 = WrapVec3(unsafe.Pointer(arg1))
	}

	return ret0
}

// Plane computes the plane based on the vertices of the given
// #graphene_triangle_t.
func (t *Triangle) Plane() Plane {
	var arg0 *C.graphene_triangle_t
	var arg1 *C.graphene_plane_t // out

	arg0 = (*C.graphene_triangle_t)(t.Native())

	C.graphene_triangle_get_plane(arg0, &arg1)

	var ret0 *Plane

	{
		ret0 = WrapPlane(unsafe.Pointer(arg1))
	}

	return ret0
}

// Points retrieves the three vertices of the given #graphene_triangle_t and
// returns their coordinates as #graphene_point3d_t.
func (t *Triangle) Points() (a Point3D, b Point3D, c Point3D) {
	var arg0 *C.graphene_triangle_t
	var arg1 *C.graphene_point3d_t // out
	var arg2 *C.graphene_point3d_t // out
	var arg3 *C.graphene_point3d_t // out

	arg0 = (*C.graphene_triangle_t)(t.Native())

	C.graphene_triangle_get_points(arg0, &arg1, &arg2, &arg3)

	var ret0 *Point3D
	var ret1 *Point3D
	var ret2 *Point3D

	{
		ret0 = WrapPoint3D(unsafe.Pointer(arg1))
	}

	{
		ret1 = WrapPoint3D(unsafe.Pointer(arg2))
	}

	{
		ret2 = WrapPoint3D(unsafe.Pointer(arg3))
	}

	return ret0, ret1, ret2
}

// Uv computes the UV coordinates of the given point @p.
//
// The point @p must lie on the same plane as the triangle @t; if the point is
// not coplanar, the result of this function is undefined. If @p is nil, the
// point will be set in (0, 0, 0).
//
// The UV coordinates will be placed in the @res vector:
//
//    - `res.x = u`
//    - `res.y = v`
//
// See also: graphene_triangle_get_barycoords()
func (t *Triangle) Uv(p *Point3D, uvA *Vec2, uvB *Vec2, uvC *Vec2) (res Vec2, ok bool) {
	var arg0 *C.graphene_triangle_t
	var arg1 *C.graphene_point3d_t
	var arg2 *C.graphene_vec2_t
	var arg3 *C.graphene_vec2_t
	var arg4 *C.graphene_vec2_t
	var arg5 *C.graphene_vec2_t // out

	arg0 = (*C.graphene_triangle_t)(t.Native())
	arg1 = (*C.graphene_point3d_t)(p.Native())
	arg2 = (*C.graphene_vec2_t)(uvA.Native())
	arg3 = (*C.graphene_vec2_t)(uvB.Native())
	arg4 = (*C.graphene_vec2_t)(uvC.Native())

	ret := C.graphene_triangle_get_uv(arg0, arg1, arg2, arg3, arg4, &arg5)

	var ret0 *Vec2
	var ret1 bool

	{
		ret0 = WrapVec2(unsafe.Pointer(arg5))
	}

	ret1 = C.bool(ret) != C.false

	return ret0, ret1
}

// Vertices retrieves the three vertices of the given #graphene_triangle_t.
func (t *Triangle) Vertices() (a Vec3, b Vec3, c Vec3) {
	var arg0 *C.graphene_triangle_t
	var arg1 *C.graphene_vec3_t // out
	var arg2 *C.graphene_vec3_t // out
	var arg3 *C.graphene_vec3_t // out

	arg0 = (*C.graphene_triangle_t)(t.Native())

	C.graphene_triangle_get_vertices(arg0, &arg1, &arg2, &arg3)

	var ret0 *Vec3
	var ret1 *Vec3
	var ret2 *Vec3

	{
		ret0 = WrapVec3(unsafe.Pointer(arg1))
	}

	{
		ret1 = WrapVec3(unsafe.Pointer(arg2))
	}

	{
		ret2 = WrapVec3(unsafe.Pointer(arg3))
	}

	return ret0, ret1, ret2
}

// InitFromFloat initializes a #graphene_triangle_t using the three given arrays
// of floating point values, each representing the coordinates of a point in 3D
// space.
func (t *Triangle) InitFromFloat(a [3]float32, b [3]float32, c [3]float32) *Triangle {
	var arg0 *C.graphene_triangle_t
	var arg1 *C.float
	var arg2 *C.float
	var arg3 *C.float

	arg0 = (*C.graphene_triangle_t)(t.Native())
	arg1 = (*C.float)(&a)
	defer runtime.KeepAlive(&a)
	arg2 = (*C.float)(&b)
	defer runtime.KeepAlive(&b)
	arg3 = (*C.float)(&c)
	defer runtime.KeepAlive(&c)

	ret := C.graphene_triangle_init_from_float(arg0, arg1, arg2, arg3)

	var ret0 *Triangle

	{
		ret0 = WrapTriangle(unsafe.Pointer(ret))
	}

	return ret0
}

// InitFromPoint3D initializes a #graphene_triangle_t using the three given 3D
// points.
func (t *Triangle) InitFromPoint3D(a *Point3D, b *Point3D, c *Point3D) *Triangle {
	var arg0 *C.graphene_triangle_t
	var arg1 *C.graphene_point3d_t
	var arg2 *C.graphene_point3d_t
	var arg3 *C.graphene_point3d_t

	arg0 = (*C.graphene_triangle_t)(t.Native())
	arg1 = (*C.graphene_point3d_t)(a.Native())
	arg2 = (*C.graphene_point3d_t)(b.Native())
	arg3 = (*C.graphene_point3d_t)(c.Native())

	ret := C.graphene_triangle_init_from_point3d(arg0, arg1, arg2, arg3)

	var ret0 *Triangle

	{
		ret0 = WrapTriangle(unsafe.Pointer(ret))
	}

	return ret0
}

// InitFromVec3 initializes a #graphene_triangle_t using the three given
// vectors.
func (t *Triangle) InitFromVec3(a *Vec3, b *Vec3, c *Vec3) *Triangle {
	var arg0 *C.graphene_triangle_t
	var arg1 *C.graphene_vec3_t
	var arg2 *C.graphene_vec3_t
	var arg3 *C.graphene_vec3_t

	arg0 = (*C.graphene_triangle_t)(t.Native())
	arg1 = (*C.graphene_vec3_t)(a.Native())
	arg2 = (*C.graphene_vec3_t)(b.Native())
	arg3 = (*C.graphene_vec3_t)(c.Native())

	ret := C.graphene_triangle_init_from_vec3(arg0, arg1, arg2, arg3)

	var ret0 *Triangle

	{
		ret0 = WrapTriangle(unsafe.Pointer(ret))
	}

	return ret0
}
