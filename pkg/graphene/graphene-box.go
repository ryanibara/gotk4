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
		{T: externglib.Type(C.graphene_box_get_type()), F: marshalBox},
	})
}

// BoxEmpty: a degenerate #graphene_box_t that can only be expanded.
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxEmpty() *Box {
	ret := C.graphene_box_empty()

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(ret))
	}

	return ret0
}

// BoxInfinite: a degenerate #graphene_box_t that cannot be expanded.
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxInfinite() *Box {
	ret := C.graphene_box_infinite()

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(ret))
	}

	return ret0
}

// BoxMinusOne: a #graphene_box_t with the minimum vertex set at (-1, -1, -1)
// and the maximum vertex set at (0, 0, 0).
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxMinusOne() *Box {
	ret := C.graphene_box_minus_one()

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(ret))
	}

	return ret0
}

// BoxOne: a #graphene_box_t with the minimum vertex set at (0, 0, 0) and the
// maximum vertex set at (1, 1, 1).
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxOne() *Box {
	ret := C.graphene_box_one()

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(ret))
	}

	return ret0
}

// BoxOneMinusOne: a #graphene_box_t with the minimum vertex set at (-1, -1, -1)
// and the maximum vertex set at (1, 1, 1).
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxOneMinusOne() *Box {
	ret := C.graphene_box_one_minus_one()

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(ret))
	}

	return ret0
}

// BoxZero: a #graphene_box_t with both the minimum and maximum vertices set at
// (0, 0, 0).
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxZero() *Box {
	ret := C.graphene_box_zero()

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(ret))
	}

	return ret0
}

// Box: a 3D box, described as the volume between a minimum and a maximum
// vertices.
type Box struct {
	native C.graphene_box_t
}

// WrapBox wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBox(ptr unsafe.Pointer) *Box {
	if ptr == nil {
		return nil
	}

	return (*Box)(ptr)
}

func marshalBox(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapBox(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (b *Box) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// NewBoxAlloc constructs a struct Box.
func NewBoxAlloc() *Box {
	ret := C.graphene_box_alloc()

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *Box) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// ContainsBox checks whether the #graphene_box_t @a contains the given
// #graphene_box_t @b.
func (a *Box) ContainsBox(b *Box) bool {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(a.Native())
	arg1 = (*C.graphene_box_t)(b.Native())

	ret := C.graphene_box_contains_box(arg0, arg1)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// ContainsPoint checks whether @box contains the given @point.
func (b *Box) ContainsPoint(point *Point3D) bool {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_point3d_t

	arg0 = (*C.graphene_box_t)(b.Native())
	arg1 = (*C.graphene_point3d_t)(point.Native())

	ret := C.graphene_box_contains_point(arg0, arg1)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Equal checks whether the two given boxes are equal.
func (a *Box) Equal(b *Box) bool {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(a.Native())
	arg1 = (*C.graphene_box_t)(b.Native())

	ret := C.graphene_box_equal(arg0, arg1)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Expand expands the dimensions of @box to include the coordinates at @point.
func (b *Box) Expand(point *Point3D) Box {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_point3d_t
	var arg2 *C.graphene_box_t // out

	arg0 = (*C.graphene_box_t)(b.Native())
	arg1 = (*C.graphene_point3d_t)(point.Native())

	C.graphene_box_expand(arg0, arg1, &arg2)

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(arg2))
	}

	return ret0
}

// ExpandScalar expands the dimensions of @box by the given @scalar value.
//
// If @scalar is positive, the #graphene_box_t will grow; if @scalar is
// negative, the #graphene_box_t will shrink.
func (b *Box) ExpandScalar(scalar float32) Box {
	var arg0 *C.graphene_box_t
	var arg1 C.float
	var arg2 *C.graphene_box_t // out

	arg0 = (*C.graphene_box_t)(b.Native())
	arg1 = C.float(scalar)

	C.graphene_box_expand_scalar(arg0, arg1, &arg2)

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(arg2))
	}

	return ret0
}

// ExpandVec3 expands the dimensions of @box to include the coordinates of the
// given vector.
func (b *Box) ExpandVec3(vec *Vec3) Box {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_vec3_t
	var arg2 *C.graphene_box_t // out

	arg0 = (*C.graphene_box_t)(b.Native())
	arg1 = (*C.graphene_vec3_t)(vec.Native())

	C.graphene_box_expand_vec3(arg0, arg1, &arg2)

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(arg2))
	}

	return ret0
}

// Free frees the resources allocated by graphene_box_alloc().
func (b *Box) Free() {
	var arg0 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(b.Native())

	C.graphene_box_free(arg0)
}

// BoundingSphere computes the bounding #graphene_sphere_t capable of containing
// the given #graphene_box_t.
func (b *Box) BoundingSphere() Sphere {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_sphere_t // out

	arg0 = (*C.graphene_box_t)(b.Native())

	C.graphene_box_get_bounding_sphere(arg0, &arg1)

	var ret0 *Sphere

	{
		ret0 = WrapSphere(unsafe.Pointer(arg1))
	}

	return ret0
}

// Center retrieves the coordinates of the center of a #graphene_box_t.
func (b *Box) Center() Point3D {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_point3d_t // out

	arg0 = (*C.graphene_box_t)(b.Native())

	C.graphene_box_get_center(arg0, &arg1)

	var ret0 *Point3D

	{
		ret0 = WrapPoint3D(unsafe.Pointer(arg1))
	}

	return ret0
}

// Depth retrieves the size of the @box on the Z axis.
func (b *Box) Depth() float32 {
	var arg0 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(b.Native())

	ret := C.graphene_box_get_depth(arg0)

	var ret0 float32

	ret0 = float32(ret)

	return ret0
}

// Height retrieves the size of the @box on the Y axis.
func (b *Box) Height() float32 {
	var arg0 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(b.Native())

	ret := C.graphene_box_get_height(arg0)

	var ret0 float32

	ret0 = float32(ret)

	return ret0
}

// Max retrieves the coordinates of the maximum point of the given
// #graphene_box_t.
func (b *Box) Max() Point3D {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_point3d_t // out

	arg0 = (*C.graphene_box_t)(b.Native())

	C.graphene_box_get_max(arg0, &arg1)

	var ret0 *Point3D

	{
		ret0 = WrapPoint3D(unsafe.Pointer(arg1))
	}

	return ret0
}

// Min retrieves the coordinates of the minimum point of the given
// #graphene_box_t.
func (b *Box) Min() Point3D {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_point3d_t // out

	arg0 = (*C.graphene_box_t)(b.Native())

	C.graphene_box_get_min(arg0, &arg1)

	var ret0 *Point3D

	{
		ret0 = WrapPoint3D(unsafe.Pointer(arg1))
	}

	return ret0
}

// Size retrieves the size of the box on all three axes, and stores it into the
// given @size vector.
func (b *Box) Size() Vec3 {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_vec3_t // out

	arg0 = (*C.graphene_box_t)(b.Native())

	C.graphene_box_get_size(arg0, &arg1)

	var ret0 *Vec3

	{
		ret0 = WrapVec3(unsafe.Pointer(arg1))
	}

	return ret0
}

// Vertices computes the vertices of the given #graphene_box_t.
func (b *Box) Vertices() [8]Vec3 {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_vec3_t // out

	arg0 = (*C.graphene_box_t)(b.Native())

	C.graphene_box_get_vertices(arg0, &arg1)

	var ret0 [8]Vec3

	{
		cArray := ([8]graphene_vec3_t)(arg1)

		for i := 0; i < 8; i++ {
			src := cArray[i]
			{
				ret0[i] = WrapVec3(unsafe.Pointer(src))
			}
		}
	}

	return ret0
}

// Width retrieves the size of the @box on the X axis.
func (b *Box) Width() float32 {
	var arg0 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(b.Native())

	ret := C.graphene_box_get_width(arg0)

	var ret0 float32

	ret0 = float32(ret)

	return ret0
}

// Init initializes the given #graphene_box_t with two vertices.
func (b *Box) Init(min *Point3D, max *Point3D) *Box {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_point3d_t
	var arg2 *C.graphene_point3d_t

	arg0 = (*C.graphene_box_t)(b.Native())
	arg1 = (*C.graphene_point3d_t)(min.Native())
	arg2 = (*C.graphene_point3d_t)(max.Native())

	ret := C.graphene_box_init(arg0, arg1, arg2)

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(ret))
	}

	return ret0
}

// InitFromBox initializes the given #graphene_box_t with the vertices of
// another #graphene_box_t.
func (b *Box) InitFromBox(src *Box) *Box {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(b.Native())
	arg1 = (*C.graphene_box_t)(src.Native())

	ret := C.graphene_box_init_from_box(arg0, arg1)

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(ret))
	}

	return ret0
}

// InitFromPoints initializes the given #graphene_box_t with the given array of
// vertices.
//
// If @n_points is 0, the returned box is initialized with graphene_box_empty().
func (b *Box) InitFromPoints(nPoints uint, points []Point3D) *Box {
	var arg0 *C.graphene_box_t
	var arg1 C.uint
	var arg2 *C.graphene_point3d_t

	arg0 = (*C.graphene_box_t)(b.Native())
	arg2 = (*C.graphene_point3d_t)(unsafe.Pointer(&points[0]))
	arg1 = len(points)
	defer runtime.KeepAlive(points)

	ret := C.graphene_box_init_from_points(arg0, arg1, arg2)

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(ret))
	}

	return ret0
}

// InitFromVec3 initializes the given #graphene_box_t with two vertices stored
// inside #graphene_vec3_t.
func (b *Box) InitFromVec3(min *Vec3, max *Vec3) *Box {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_vec3_t
	var arg2 *C.graphene_vec3_t

	arg0 = (*C.graphene_box_t)(b.Native())
	arg1 = (*C.graphene_vec3_t)(min.Native())
	arg2 = (*C.graphene_vec3_t)(max.Native())

	ret := C.graphene_box_init_from_vec3(arg0, arg1, arg2)

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(ret))
	}

	return ret0
}

// InitFromVectors initializes the given #graphene_box_t with the given array of
// vertices.
//
// If @n_vectors is 0, the returned box is initialized with
// graphene_box_empty().
func (b *Box) InitFromVectors(nVectors uint, vectors []Vec3) *Box {
	var arg0 *C.graphene_box_t
	var arg1 C.uint
	var arg2 *C.graphene_vec3_t

	arg0 = (*C.graphene_box_t)(b.Native())
	arg2 = (*C.graphene_vec3_t)(unsafe.Pointer(&vectors[0]))
	arg1 = len(vectors)
	defer runtime.KeepAlive(vectors)

	ret := C.graphene_box_init_from_vectors(arg0, arg1, arg2)

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(ret))
	}

	return ret0
}

// Intersection intersects the two given #graphene_box_t.
//
// If the two boxes do not intersect, @res will contain a degenerate box
// initialized with graphene_box_empty().
func (a *Box) Intersection(b *Box) (res Box, ok bool) {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_box_t
	var arg2 *C.graphene_box_t // out

	arg0 = (*C.graphene_box_t)(a.Native())
	arg1 = (*C.graphene_box_t)(b.Native())

	ret := C.graphene_box_intersection(arg0, arg1, &arg2)

	var ret0 *Box
	var ret1 bool

	{
		ret0 = WrapBox(unsafe.Pointer(arg2))
	}

	ret1 = C.bool(ret) != 0

	return ret0, ret1
}

// Union unions the two given #graphene_box_t.
func (a *Box) Union(b *Box) Box {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_box_t
	var arg2 *C.graphene_box_t // out

	arg0 = (*C.graphene_box_t)(a.Native())
	arg1 = (*C.graphene_box_t)(b.Native())

	C.graphene_box_union(arg0, arg1, &arg2)

	var ret0 *Box

	{
		ret0 = WrapBox(unsafe.Pointer(arg2))
	}

	return ret0
}
