// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

// glib.Type values for graphene-triangle.go.
var GTypeTriangle = externglib.Type(C.graphene_triangle_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeTriangle, F: marshalTriangle},
	})
}

// Triangle: triangle.
//
// An instance of this type is always passed by reference.
type Triangle struct {
	*triangle
}

// triangle is the struct that's finalized.
type triangle struct {
	native *C.graphene_triangle_t
}

func marshalTriangle(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Triangle{&triangle{(*C.graphene_triangle_t)(b)}}, nil
}

// NewTriangleAlloc constructs a struct Triangle.
func NewTriangleAlloc() *Triangle {
	var _cret *C.graphene_triangle_t // in

	_cret = C.graphene_triangle_alloc()

	var _triangle *Triangle // out

	_triangle = (*Triangle)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_triangle)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.graphene_triangle_free((*C.graphene_triangle_t)(intern.C))
		},
	)

	return _triangle
}

// ContainsPoint checks whether the given triangle t contains the point p.
//
// The function takes the following parameters:
//
//    - p: #graphene_point3d_t.
//
// The function returns the following values:
//
//    - ok: true if the point is inside the triangle.
//
func (t *Triangle) ContainsPoint(p *Point3D) bool {
	var _arg0 *C.graphene_triangle_t // out
	var _arg1 *C.graphene_point3d_t  // out
	var _cret C._Bool                // in

	_arg0 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(t)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(p)))

	_cret = C.graphene_triangle_contains_point(_arg0, _arg1)
	runtime.KeepAlive(t)
	runtime.KeepAlive(p)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Equal checks whether the two given #graphene_triangle_t are equal.
//
// The function takes the following parameters:
//
//    - b: #graphene_triangle_t.
//
// The function returns the following values:
//
//    - ok: true if the triangles are equal.
//
func (a *Triangle) Equal(b *Triangle) bool {
	var _arg0 *C.graphene_triangle_t // out
	var _arg1 *C.graphene_triangle_t // out
	var _cret C._Bool                // in

	_arg0 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_triangle_equal(_arg0, _arg1)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Area computes the area of the given #graphene_triangle_t.
//
// The function returns the following values:
//
//    - gfloat: area of the triangle.
//
func (t *Triangle) Area() float32 {
	var _arg0 *C.graphene_triangle_t // out
	var _cret C.float                // in

	_arg0 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(t)))

	_cret = C.graphene_triangle_get_area(_arg0)
	runtime.KeepAlive(t)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Barycoords computes the barycentric coordinates
// (http://en.wikipedia.org/wiki/Barycentric_coordinate_system) of the given
// point p.
//
// The point p must lie on the same plane as the triangle t; if the point is not
// coplanar, the result of this function is undefined.
//
// If we place the origin in the coordinates of the triangle's A point, the
// barycentric coordinates are u, which is on the AC vector; and v which is on
// the AB vector:
//
// ! (triangle-barycentric.png)
//
// The returned #graphene_vec2_t contains the following values, in order:
//
//    - res.x = u
//    - res.y = v.
//
// The function takes the following parameters:
//
//    - p (optional): #graphene_point3d_t.
//
// The function returns the following values:
//
//    - res: return location for the vector with the barycentric coordinates.
//    - ok: true if the barycentric coordinates are valid.
//
func (t *Triangle) Barycoords(p *Point3D) (*Vec2, bool) {
	var _arg0 *C.graphene_triangle_t // out
	var _arg1 *C.graphene_point3d_t  // out
	var _arg2 C.graphene_vec2_t      // in
	var _cret C._Bool                // in

	_arg0 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(t)))
	if p != nil {
		_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(p)))
	}

	_cret = C.graphene_triangle_get_barycoords(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(t)
	runtime.KeepAlive(p)

	var _res *Vec2 // out
	var _ok bool   // out

	_res = (*Vec2)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret {
		_ok = true
	}

	return _res, _ok
}

// BoundingBox computes the bounding box of the given #graphene_triangle_t.
//
// The function returns the following values:
//
//    - res: return location for the box.
//
func (t *Triangle) BoundingBox() *Box {
	var _arg0 *C.graphene_triangle_t // out
	var _arg1 C.graphene_box_t       // in

	_arg0 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(t)))

	C.graphene_triangle_get_bounding_box(_arg0, &_arg1)
	runtime.KeepAlive(t)

	var _res *Box // out

	_res = (*Box)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// Midpoint computes the coordinates of the midpoint of the given
// #graphene_triangle_t.
//
// The midpoint G is the centroid
// (https://en.wikipedia.org/wiki/Centroid#Triangle_centroid) of the triangle,
// i.e. the intersection of its medians.
//
// The function returns the following values:
//
//    - res: return location for the coordinates of the midpoint.
//
func (t *Triangle) Midpoint() *Point3D {
	var _arg0 *C.graphene_triangle_t // out
	var _arg1 C.graphene_point3d_t   // in

	_arg0 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(t)))

	C.graphene_triangle_get_midpoint(_arg0, &_arg1)
	runtime.KeepAlive(t)

	var _res *Point3D // out

	_res = (*Point3D)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// Normal computes the normal vector of the given #graphene_triangle_t.
//
// The function returns the following values:
//
//    - res: return location for the normal vector.
//
func (t *Triangle) Normal() *Vec3 {
	var _arg0 *C.graphene_triangle_t // out
	var _arg1 C.graphene_vec3_t      // in

	_arg0 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(t)))

	C.graphene_triangle_get_normal(_arg0, &_arg1)
	runtime.KeepAlive(t)

	var _res *Vec3 // out

	_res = (*Vec3)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// Plane computes the plane based on the vertices of the given
// #graphene_triangle_t.
//
// The function returns the following values:
//
//    - res: return location for the plane.
//
func (t *Triangle) Plane() *Plane {
	var _arg0 *C.graphene_triangle_t // out
	var _arg1 C.graphene_plane_t     // in

	_arg0 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(t)))

	C.graphene_triangle_get_plane(_arg0, &_arg1)
	runtime.KeepAlive(t)

	var _res *Plane // out

	_res = (*Plane)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// Points retrieves the three vertices of the given #graphene_triangle_t and
// returns their coordinates as #graphene_point3d_t.
//
// The function returns the following values:
//
//    - a (optional): return location for the coordinates of the first vertex.
//    - b (optional): return location for the coordinates of the second vertex.
//    - c (optional): return location for the coordinates of the third vertex.
//
func (t *Triangle) Points() (a *Point3D, b *Point3D, c *Point3D) {
	var _arg0 *C.graphene_triangle_t // out
	var _arg1 C.graphene_point3d_t   // in
	var _arg2 C.graphene_point3d_t   // in
	var _arg3 C.graphene_point3d_t   // in

	_arg0 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(t)))

	C.graphene_triangle_get_points(_arg0, &_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(t)

	var _a *Point3D // out
	var _b *Point3D // out
	var _c *Point3D // out

	_a = (*Point3D)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	_b = (*Point3D)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	_c = (*Point3D)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _a, _b, _c
}

// Uv computes the UV coordinates of the given point p.
//
// The point p must lie on the same plane as the triangle t; if the point is not
// coplanar, the result of this function is undefined. If p is NULL, the point
// will be set in (0, 0, 0).
//
// The UV coordinates will be placed in the res vector:
//
//    - res.x = u
//    - res.y = v
//
// See also: graphene_triangle_get_barycoords().
//
// The function takes the following parameters:
//
//    - p (optional): #graphene_point3d_t.
//    - uvA: UV coordinates of the first point.
//    - uvB: UV coordinates of the second point.
//    - uvC: UV coordinates of the third point.
//
// The function returns the following values:
//
//    - res: vector containing the UV coordinates of the given point p.
//    - ok: true if the coordinates are valid.
//
func (t *Triangle) Uv(p *Point3D, uvA *Vec2, uvB *Vec2, uvC *Vec2) (*Vec2, bool) {
	var _arg0 *C.graphene_triangle_t // out
	var _arg1 *C.graphene_point3d_t  // out
	var _arg2 *C.graphene_vec2_t     // out
	var _arg3 *C.graphene_vec2_t     // out
	var _arg4 *C.graphene_vec2_t     // out
	var _arg5 C.graphene_vec2_t      // in
	var _cret C._Bool                // in

	_arg0 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(t)))
	if p != nil {
		_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(p)))
	}
	_arg2 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(uvA)))
	_arg3 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(uvB)))
	_arg4 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(uvC)))

	_cret = C.graphene_triangle_get_uv(_arg0, _arg1, _arg2, _arg3, _arg4, &_arg5)
	runtime.KeepAlive(t)
	runtime.KeepAlive(p)
	runtime.KeepAlive(uvA)
	runtime.KeepAlive(uvB)
	runtime.KeepAlive(uvC)

	var _res *Vec2 // out
	var _ok bool   // out

	_res = (*Vec2)(gextras.NewStructNative(unsafe.Pointer((&_arg5))))
	if _cret {
		_ok = true
	}

	return _res, _ok
}

// Vertices retrieves the three vertices of the given #graphene_triangle_t.
//
// The function returns the following values:
//
//    - a (optional): return location for the first vertex.
//    - b (optional): return location for the second vertex.
//    - c (optional): return location for the third vertex.
//
func (t *Triangle) Vertices() (a *Vec3, b *Vec3, c *Vec3) {
	var _arg0 *C.graphene_triangle_t // out
	var _arg1 C.graphene_vec3_t      // in
	var _arg2 C.graphene_vec3_t      // in
	var _arg3 C.graphene_vec3_t      // in

	_arg0 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(t)))

	C.graphene_triangle_get_vertices(_arg0, &_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(t)

	var _a *Vec3 // out
	var _b *Vec3 // out
	var _c *Vec3 // out

	_a = (*Vec3)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	_b = (*Vec3)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	_c = (*Vec3)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _a, _b, _c
}

// InitFromFloat initializes a #graphene_triangle_t using the three given arrays
// of floating point values, each representing the coordinates of a point in 3D
// space.
//
// The function takes the following parameters:
//
//    - a: array of 3 floating point values.
//    - b: array of 3 floating point values.
//    - c: array of 3 floating point values.
//
// The function returns the following values:
//
//    - triangle: initialized #graphene_triangle_t.
//
func (t *Triangle) InitFromFloat(a [3]float32, b [3]float32, c [3]float32) *Triangle {
	var _arg0 *C.graphene_triangle_t // out
	var _arg1 *C.float               // out
	var _arg2 *C.float               // out
	var _arg3 *C.float               // out
	var _cret *C.graphene_triangle_t // in

	_arg0 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(t)))
	_arg1 = (*C.float)(unsafe.Pointer(&a))
	_arg2 = (*C.float)(unsafe.Pointer(&b))
	_arg3 = (*C.float)(unsafe.Pointer(&c))

	_cret = C.graphene_triangle_init_from_float(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(t)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
	runtime.KeepAlive(c)

	var _triangle *Triangle // out

	_triangle = (*Triangle)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _triangle
}

// InitFromPoint3D initializes a #graphene_triangle_t using the three given 3D
// points.
//
// The function takes the following parameters:
//
//    - a (optional): #graphene_point3d_t.
//    - b (optional): #graphene_point3d_t.
//    - c (optional): #graphene_point3d_t.
//
// The function returns the following values:
//
//    - triangle: initialized #graphene_triangle_t.
//
func (t *Triangle) InitFromPoint3D(a *Point3D, b *Point3D, c *Point3D) *Triangle {
	var _arg0 *C.graphene_triangle_t // out
	var _arg1 *C.graphene_point3d_t  // out
	var _arg2 *C.graphene_point3d_t  // out
	var _arg3 *C.graphene_point3d_t  // out
	var _cret *C.graphene_triangle_t // in

	_arg0 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(t)))
	if a != nil {
		_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(a)))
	}
	if b != nil {
		_arg2 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(b)))
	}
	if c != nil {
		_arg3 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(c)))
	}

	_cret = C.graphene_triangle_init_from_point3d(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(t)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
	runtime.KeepAlive(c)

	var _triangle *Triangle // out

	_triangle = (*Triangle)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _triangle
}

// InitFromVec3 initializes a #graphene_triangle_t using the three given
// vectors.
//
// The function takes the following parameters:
//
//    - a (optional): #graphene_vec3_t.
//    - b (optional): #graphene_vec3_t.
//    - c (optional): #graphene_vec3_t.
//
// The function returns the following values:
//
//    - triangle: initialized #graphene_triangle_t.
//
func (t *Triangle) InitFromVec3(a *Vec3, b *Vec3, c *Vec3) *Triangle {
	var _arg0 *C.graphene_triangle_t // out
	var _arg1 *C.graphene_vec3_t     // out
	var _arg2 *C.graphene_vec3_t     // out
	var _arg3 *C.graphene_vec3_t     // out
	var _cret *C.graphene_triangle_t // in

	_arg0 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(t)))
	if a != nil {
		_arg1 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(a)))
	}
	if b != nil {
		_arg2 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(b)))
	}
	if c != nil {
		_arg3 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(c)))
	}

	_cret = C.graphene_triangle_init_from_vec3(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(t)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
	runtime.KeepAlive(c)

	var _triangle *Triangle // out

	_triangle = (*Triangle)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _triangle
}
