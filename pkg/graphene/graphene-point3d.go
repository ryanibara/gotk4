// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_point3d_get_type()), F: marshalPoint3D},
	})
}

// Point3D: a point with three components: X, Y, and Z.
type Point3D struct {
	native C.graphene_point3d_t
}

// WrapPoint3D wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPoint3D(ptr unsafe.Pointer) *Point3D {
	if ptr == nil {
		return nil
	}

	return (*Point3D)(ptr)
}

func marshalPoint3D(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPoint3D(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *Point3D) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// X gets the field inside the struct.
func (p *Point3D) X() float32 {
	var v float32
	v = (float32)(p.native.x)
	return v
}

// Y gets the field inside the struct.
func (p *Point3D) Y() float32 {
	var v float32
	v = (float32)(p.native.y)
	return v
}

// Z gets the field inside the struct.
func (p *Point3D) Z() float32 {
	var v float32
	v = (float32)(p.native.z)
	return v
}

// Cross computes the cross product of the two given #graphene_point3d_t.
func (a *Point3D) Cross(b *Point3D) Point3D {
	var _arg0 *C.graphene_point3d_t
	var _arg1 *C.graphene_point3d_t

	_arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(b.Native()))

	var _res Point3D

	C.graphene_point3d_cross(_arg0, _arg1, (*C.graphene_point3d_t)(unsafe.Pointer(&_res)))

	return _res
}

// Distance computes the distance between the two given #graphene_point3d_t.
func (a *Point3D) Distance(b *Point3D) (Vec3, float32) {
	var _arg0 *C.graphene_point3d_t
	var _arg1 *C.graphene_point3d_t

	_arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(b.Native()))

	var _delta Vec3
	var _cret C.float

	_cret = C.graphene_point3d_distance(_arg0, _arg1, (*C.graphene_vec3_t)(unsafe.Pointer(&_delta)))

	var _gfloat float32

	_gfloat = (float32)(_cret)

	return _delta, _gfloat
}

// Dot computes the dot product of the two given #graphene_point3d_t.
func (a *Point3D) Dot(b *Point3D) float32 {
	var _arg0 *C.graphene_point3d_t
	var _arg1 *C.graphene_point3d_t

	_arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(b.Native()))

	var _cret C.float

	_cret = C.graphene_point3d_dot(_arg0, _arg1)

	var _gfloat float32

	_gfloat = (float32)(_cret)

	return _gfloat
}

// Equal checks whether two given points are equal.
func (a *Point3D) Equal(b *Point3D) bool {
	var _arg0 *C.graphene_point3d_t
	var _arg1 *C.graphene_point3d_t

	_arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(b.Native()))

	var _cret C._Bool

	_cret = C.graphene_point3d_equal(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Free frees the resources allocated via graphene_point3d_alloc().
func (p *Point3D) Free() {
	var _arg0 *C.graphene_point3d_t

	_arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))

	C.graphene_point3d_free(_arg0)
}

// Interpolate: linearly interpolates each component of @a and @b using the
// provided @factor, and places the result in @res.
func (a *Point3D) Interpolate(b *Point3D, factor float64) Point3D {
	var _arg0 *C.graphene_point3d_t
	var _arg1 *C.graphene_point3d_t
	var _arg2 C.double

	_arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(b.Native()))
	_arg2 = C.double(factor)

	var _res Point3D

	C.graphene_point3d_interpolate(_arg0, _arg1, _arg2, (*C.graphene_point3d_t)(unsafe.Pointer(&_res)))

	return _res
}

// Length computes the length of the vector represented by the coordinates of
// the given #graphene_point3d_t.
func (p *Point3D) Length() float32 {
	var _arg0 *C.graphene_point3d_t

	_arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))

	var _cret C.float

	_cret = C.graphene_point3d_length(_arg0)

	var _gfloat float32

	_gfloat = (float32)(_cret)

	return _gfloat
}

// Near checks whether the two points are near each other, within an @epsilon
// factor.
func (a *Point3D) Near(b *Point3D, epsilon float32) bool {
	var _arg0 *C.graphene_point3d_t
	var _arg1 *C.graphene_point3d_t
	var _arg2 C.float

	_arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(b.Native()))
	_arg2 = C.float(epsilon)

	var _cret C._Bool

	_cret = C.graphene_point3d_near(_arg0, _arg1, _arg2)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Normalize computes the normalization of the vector represented by the
// coordinates of the given #graphene_point3d_t.
func (p *Point3D) Normalize() Point3D {
	var _arg0 *C.graphene_point3d_t

	_arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))

	var _res Point3D

	C.graphene_point3d_normalize(_arg0, (*C.graphene_point3d_t)(unsafe.Pointer(&_res)))

	return _res
}

// NormalizeViewport normalizes the coordinates of a #graphene_point3d_t using
// the given viewport and clipping planes.
//
// The coordinates of the resulting #graphene_point3d_t will be in the [ -1, 1 ]
// range.
func (p *Point3D) NormalizeViewport(viewport *Rect, zNear float32, zFar float32) Point3D {
	var _arg0 *C.graphene_point3d_t
	var _arg1 *C.graphene_rect_t
	var _arg2 C.float
	var _arg3 C.float

	_arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(viewport.Native()))
	_arg2 = C.float(zNear)
	_arg3 = C.float(zFar)

	var _res Point3D

	C.graphene_point3d_normalize_viewport(_arg0, _arg1, _arg2, _arg3, (*C.graphene_point3d_t)(unsafe.Pointer(&_res)))

	return _res
}

// Scale scales the coordinates of the given #graphene_point3d_t by the given
// @factor.
func (p *Point3D) Scale(factor float32) Point3D {
	var _arg0 *C.graphene_point3d_t
	var _arg1 C.float

	_arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))
	_arg1 = C.float(factor)

	var _res Point3D

	C.graphene_point3d_scale(_arg0, _arg1, (*C.graphene_point3d_t)(unsafe.Pointer(&_res)))

	return _res
}

// ToVec3 stores the coordinates of a #graphene_point3d_t into a
// #graphene_vec3_t.
func (p *Point3D) ToVec3() Vec3 {
	var _arg0 *C.graphene_point3d_t

	_arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))

	var _v Vec3

	C.graphene_point3d_to_vec3(_arg0, (*C.graphene_vec3_t)(unsafe.Pointer(&_v)))

	return _v
}
