// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
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
	var cret *C.graphene_box_t
	var ret1 *Box

	cret = C.graphene_box_empty()

	ret1 = WrapBox(unsafe.Pointer(cret))

	return ret1
}

// BoxInfinite: a degenerate #graphene_box_t that cannot be expanded.
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxInfinite() *Box {
	var cret *C.graphene_box_t
	var ret1 *Box

	cret = C.graphene_box_infinite()

	ret1 = WrapBox(unsafe.Pointer(cret))

	return ret1
}

// BoxMinusOne: a #graphene_box_t with the minimum vertex set at (-1, -1, -1)
// and the maximum vertex set at (0, 0, 0).
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxMinusOne() *Box {
	var cret *C.graphene_box_t
	var ret1 *Box

	cret = C.graphene_box_minus_one()

	ret1 = WrapBox(unsafe.Pointer(cret))

	return ret1
}

// BoxOne: a #graphene_box_t with the minimum vertex set at (0, 0, 0) and the
// maximum vertex set at (1, 1, 1).
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxOne() *Box {
	var cret *C.graphene_box_t
	var ret1 *Box

	cret = C.graphene_box_one()

	ret1 = WrapBox(unsafe.Pointer(cret))

	return ret1
}

// BoxOneMinusOne: a #graphene_box_t with the minimum vertex set at (-1, -1, -1)
// and the maximum vertex set at (1, 1, 1).
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxOneMinusOne() *Box {
	var cret *C.graphene_box_t
	var ret1 *Box

	cret = C.graphene_box_one_minus_one()

	ret1 = WrapBox(unsafe.Pointer(cret))

	return ret1
}

// BoxZero: a #graphene_box_t with both the minimum and maximum vertices set at
// (0, 0, 0).
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxZero() *Box {
	var cret *C.graphene_box_t
	var ret1 *Box

	cret = C.graphene_box_zero()

	ret1 = WrapBox(unsafe.Pointer(cret))

	return ret1
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

// NewBoxAlloc constructs a struct Box.
func NewBoxAlloc() *Box {
	var cret *C.graphene_box_t
	var ret1 *Box

	cret = C.graphene_box_alloc()

	ret1 = WrapBox(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Box) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Native returns the underlying C source pointer.
func (b *Box) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// ContainsBox checks whether the #graphene_box_t @a contains the given
// #graphene_box_t @b.
func (a *Box) ContainsBox(b *Box) bool {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var cret C._Bool
	var ret1 bool

	cret = C.graphene_box_contains_box(arg0, b)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// ContainsPoint checks whether @box contains the given @point.
func (b *Box) ContainsPoint(point *Point3D) bool {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_point3d_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))
	arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(point.Native()))

	var cret C._Bool
	var ret1 bool

	cret = C.graphene_box_contains_point(arg0, point)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Equal checks whether the two given boxes are equal.
func (a *Box) Equal(b *Box) bool {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var cret C._Bool
	var ret1 bool

	cret = C.graphene_box_equal(arg0, b)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Expand expands the dimensions of @box to include the coordinates at @point.
func (b *Box) Expand(point *Point3D) Box {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_point3d_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))
	arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(point.Native()))

	var arg2 C.graphene_box_t
	var ret2 *Box

	C.graphene_box_expand(arg0, point, &arg2)

	*ret2 = WrapBox(unsafe.Pointer(arg2))

	return ret2
}

// ExpandScalar expands the dimensions of @box by the given @scalar value.
//
// If @scalar is positive, the #graphene_box_t will grow; if @scalar is
// negative, the #graphene_box_t will shrink.
func (b *Box) ExpandScalar(scalar float32) Box {
	var arg0 *C.graphene_box_t
	var arg1 C.float

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))
	arg1 = C.float(scalar)

	var arg2 C.graphene_box_t
	var ret2 *Box

	C.graphene_box_expand_scalar(arg0, scalar, &arg2)

	*ret2 = WrapBox(unsafe.Pointer(arg2))

	return ret2
}

// ExpandVec3 expands the dimensions of @box to include the coordinates of the
// given vector.
func (b *Box) ExpandVec3(vec *Vec3) Box {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_vec3_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))
	arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(vec.Native()))

	var arg2 C.graphene_box_t
	var ret2 *Box

	C.graphene_box_expand_vec3(arg0, vec, &arg2)

	*ret2 = WrapBox(unsafe.Pointer(arg2))

	return ret2
}

// Free frees the resources allocated by graphene_box_alloc().
func (b *Box) Free() {
	var arg0 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	C.graphene_box_free(arg0)
}

// BoundingSphere computes the bounding #graphene_sphere_t capable of containing
// the given #graphene_box_t.
func (b *Box) BoundingSphere() Sphere {
	var arg0 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var arg1 C.graphene_sphere_t
	var ret1 *Sphere

	C.graphene_box_get_bounding_sphere(arg0, &arg1)

	*ret1 = WrapSphere(unsafe.Pointer(arg1))

	return ret1
}

// Center retrieves the coordinates of the center of a #graphene_box_t.
func (b *Box) Center() Point3D {
	var arg0 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var arg1 C.graphene_point3d_t
	var ret1 *Point3D

	C.graphene_box_get_center(arg0, &arg1)

	*ret1 = WrapPoint3D(unsafe.Pointer(arg1))

	return ret1
}

// Depth retrieves the size of the @box on the Z axis.
func (b *Box) Depth() float32 {
	var arg0 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var cret C.float
	var ret1 float32

	cret = C.graphene_box_get_depth(arg0)

	ret1 = C.float(cret)

	return ret1
}

// Height retrieves the size of the @box on the Y axis.
func (b *Box) Height() float32 {
	var arg0 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var cret C.float
	var ret1 float32

	cret = C.graphene_box_get_height(arg0)

	ret1 = C.float(cret)

	return ret1
}

// Max retrieves the coordinates of the maximum point of the given
// #graphene_box_t.
func (b *Box) Max() Point3D {
	var arg0 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var arg1 C.graphene_point3d_t
	var ret1 *Point3D

	C.graphene_box_get_max(arg0, &arg1)

	*ret1 = WrapPoint3D(unsafe.Pointer(arg1))

	return ret1
}

// Min retrieves the coordinates of the minimum point of the given
// #graphene_box_t.
func (b *Box) Min() Point3D {
	var arg0 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var arg1 C.graphene_point3d_t
	var ret1 *Point3D

	C.graphene_box_get_min(arg0, &arg1)

	*ret1 = WrapPoint3D(unsafe.Pointer(arg1))

	return ret1
}

// Size retrieves the size of the box on all three axes, and stores it into the
// given @size vector.
func (b *Box) Size() Vec3 {
	var arg0 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var arg1 C.graphene_vec3_t
	var ret1 *Vec3

	C.graphene_box_get_size(arg0, &arg1)

	*ret1 = WrapVec3(unsafe.Pointer(arg1))

	return ret1
}

// Vertices computes the vertices of the given #graphene_box_t.
func (b *Box) Vertices() [8]Vec3 {
	var arg0 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var arg1 [8]C.graphene_vec3_t
	var ret1 [8]Vec3

	C.graphene_box_get_vertices(arg0, &arg1)

	{
		tmp := *(*[8]Vec3)(unsafe.Pointer(&arg1))
		for i := 0; i < 8; i++ {
			src := tmp[i]
			ret1[i] = WrapVec3(unsafe.Pointer(src))
		}
	}

	return ret1
}

// Width retrieves the size of the @box on the X axis.
func (b *Box) Width() float32 {
	var arg0 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var cret C.float
	var ret1 float32

	cret = C.graphene_box_get_width(arg0)

	ret1 = C.float(cret)

	return ret1
}

// Init initializes the given #graphene_box_t with two vertices.
func (b *Box) Init(min *Point3D, max *Point3D) *Box {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_point3d_t
	var arg2 *C.graphene_point3d_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))
	arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(min.Native()))
	arg2 = (*C.graphene_point3d_t)(unsafe.Pointer(max.Native()))

	var cret *C.graphene_box_t
	var ret1 *Box

	cret = C.graphene_box_init(arg0, min, max)

	ret1 = WrapBox(unsafe.Pointer(cret))

	return ret1
}

// InitFromBox initializes the given #graphene_box_t with the vertices of
// another #graphene_box_t.
func (b *Box) InitFromBox(src *Box) *Box {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))
	arg1 = (*C.graphene_box_t)(unsafe.Pointer(src.Native()))

	var cret *C.graphene_box_t
	var ret1 *Box

	cret = C.graphene_box_init_from_box(arg0, src)

	ret1 = WrapBox(unsafe.Pointer(cret))

	return ret1
}

// InitFromVec3 initializes the given #graphene_box_t with two vertices stored
// inside #graphene_vec3_t.
func (b *Box) InitFromVec3(min *Vec3, max *Vec3) *Box {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_vec3_t
	var arg2 *C.graphene_vec3_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))
	arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(min.Native()))
	arg2 = (*C.graphene_vec3_t)(unsafe.Pointer(max.Native()))

	var cret *C.graphene_box_t
	var ret1 *Box

	cret = C.graphene_box_init_from_vec3(arg0, min, max)

	ret1 = WrapBox(unsafe.Pointer(cret))

	return ret1
}

// Intersection intersects the two given #graphene_box_t.
//
// If the two boxes do not intersect, @res will contain a degenerate box
// initialized with graphene_box_empty().
func (a *Box) Intersection(b *Box) (res Box, ok bool) {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var arg2 C.graphene_box_t
	var ret2 *Box
	var cret C._Bool
	var ret2 bool

	cret = C.graphene_box_intersection(arg0, b, &arg2)

	*ret2 = WrapBox(unsafe.Pointer(arg2))
	ret2 = C.bool(cret) != C.false

	return ret2, ret2
}

// Union unions the two given #graphene_box_t.
func (a *Box) Union(b *Box) Box {
	var arg0 *C.graphene_box_t
	var arg1 *C.graphene_box_t

	arg0 = (*C.graphene_box_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var arg2 C.graphene_box_t
	var ret2 *Box

	C.graphene_box_union(arg0, b, &arg2)

	*ret2 = WrapBox(unsafe.Pointer(arg2))

	return ret2
}
