// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_ray_get_type()), F: marshalRay},
	})
}

// RayIntersectionKind: type of intersection.
type RayIntersectionKind int

const (
	// RayIntersectionKindNone: no intersection.
	RayIntersectionKindNone RayIntersectionKind = iota
	// RayIntersectionKindEnter: ray is entering the intersected object.
	RayIntersectionKindEnter
	// RayIntersectionKindLeave: ray is leaving the intersected object.
	RayIntersectionKindLeave
)

// String returns the name in string for RayIntersectionKind.
func (r RayIntersectionKind) String() string {
	switch r {
	case RayIntersectionKindNone:
		return "None"
	case RayIntersectionKindEnter:
		return "Enter"
	case RayIntersectionKindLeave:
		return "Leave"
	default:
		return fmt.Sprintf("RayIntersectionKind(%d)", r)
	}
}

// Ray: ray emitted from an origin in a given direction.
//
// The contents of the graphene_ray_t structure are private, and should not be
// modified directly.
//
// An instance of this type is always passed by reference.
type Ray struct {
	*ray
}

// ray is the struct that's finalized.
type ray struct {
	native *C.graphene_ray_t
}

func marshalRay(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Ray{&ray{(*C.graphene_ray_t)(b)}}, nil
}

// NewRayAlloc constructs a struct Ray.
func NewRayAlloc() *Ray {
	var _cret *C.graphene_ray_t // in

	_cret = C.graphene_ray_alloc()

	var _ray *Ray // out

	_ray = (*Ray)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_ray)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.graphene_ray_free((*C.graphene_ray_t)(intern.C))
		},
	)

	return _ray
}

// Equal checks whether the two given #graphene_ray_t are equal.
func (a *Ray) Equal(b *Ray) bool {
	var _arg0 *C.graphene_ray_t // out
	var _arg1 *C.graphene_ray_t // out
	var _cret C._Bool           // in

	_arg0 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_ray_equal(_arg0, _arg1)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// ClosestPointToPoint computes the point on the given #graphene_ray_t that is
// closest to the given point p.
func (r *Ray) ClosestPointToPoint(p *Point3D) Point3D {
	var _arg0 *C.graphene_ray_t     // out
	var _arg1 *C.graphene_point3d_t // out
	var _arg2 C.graphene_point3d_t  // in

	_arg0 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(p)))

	C.graphene_ray_get_closest_point_to_point(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(r)
	runtime.KeepAlive(p)

	var _res Point3D // out

	_res = *(*Point3D)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Direction retrieves the direction of the given #graphene_ray_t.
func (r *Ray) Direction() Vec3 {
	var _arg0 *C.graphene_ray_t // out
	var _arg1 C.graphene_vec3_t // in

	_arg0 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(r)))

	C.graphene_ray_get_direction(_arg0, &_arg1)
	runtime.KeepAlive(r)

	var _direction Vec3 // out

	_direction = *(*Vec3)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _direction
}

// DistanceToPlane computes the distance of the origin of the given
// #graphene_ray_t from the given plane.
//
// If the ray does not intersect the plane, this function returns INFINITY.
func (r *Ray) DistanceToPlane(p *Plane) float32 {
	var _arg0 *C.graphene_ray_t   // out
	var _arg1 *C.graphene_plane_t // out
	var _cret C.float             // in

	_arg0 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p)))

	_cret = C.graphene_ray_get_distance_to_plane(_arg0, _arg1)
	runtime.KeepAlive(r)
	runtime.KeepAlive(p)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// DistanceToPoint computes the distance of the closest approach between the
// given #graphene_ray_t r and the point p.
//
// The closest approach to a ray from a point is the distance between the point
// and the projection of the point on the ray itself.
func (r *Ray) DistanceToPoint(p *Point3D) float32 {
	var _arg0 *C.graphene_ray_t     // out
	var _arg1 *C.graphene_point3d_t // out
	var _cret C.float               // in

	_arg0 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(p)))

	_cret = C.graphene_ray_get_distance_to_point(_arg0, _arg1)
	runtime.KeepAlive(r)
	runtime.KeepAlive(p)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Origin retrieves the origin of the given #graphene_ray_t.
func (r *Ray) Origin() Point3D {
	var _arg0 *C.graphene_ray_t    // out
	var _arg1 C.graphene_point3d_t // in

	_arg0 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(r)))

	C.graphene_ray_get_origin(_arg0, &_arg1)
	runtime.KeepAlive(r)

	var _origin Point3D // out

	_origin = *(*Point3D)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _origin
}

// PositionAt retrieves the coordinates of a point at the distance t along the
// given #graphene_ray_t.
func (r *Ray) PositionAt(t float32) Point3D {
	var _arg0 *C.graphene_ray_t    // out
	var _arg1 C.float              // out
	var _arg2 C.graphene_point3d_t // in

	_arg0 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = C.float(t)

	C.graphene_ray_get_position_at(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(r)
	runtime.KeepAlive(t)

	var _position Point3D // out

	_position = *(*Point3D)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _position
}

// Init initializes the given #graphene_ray_t using the given origin and
// direction values.
func (r *Ray) Init(origin *Point3D, direction *Vec3) *Ray {
	var _arg0 *C.graphene_ray_t     // out
	var _arg1 *C.graphene_point3d_t // out
	var _arg2 *C.graphene_vec3_t    // out
	var _cret *C.graphene_ray_t     // in

	_arg0 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(r)))
	if origin != nil {
		_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(origin)))
	}
	if direction != nil {
		_arg2 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(direction)))
	}

	_cret = C.graphene_ray_init(_arg0, _arg1, _arg2)
	runtime.KeepAlive(r)
	runtime.KeepAlive(origin)
	runtime.KeepAlive(direction)

	var _ray *Ray // out

	_ray = (*Ray)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _ray
}

// InitFromRay initializes the given #graphene_ray_t using the origin and
// direction values of another #graphene_ray_t.
func (r *Ray) InitFromRay(src *Ray) *Ray {
	var _arg0 *C.graphene_ray_t // out
	var _arg1 *C.graphene_ray_t // out
	var _cret *C.graphene_ray_t // in

	_arg0 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.graphene_ray_init_from_ray(_arg0, _arg1)
	runtime.KeepAlive(r)
	runtime.KeepAlive(src)

	var _ray *Ray // out

	_ray = (*Ray)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _ray
}

// InitFromVec3 initializes the given #graphene_ray_t using the given vectors.
func (r *Ray) InitFromVec3(origin *Vec3, direction *Vec3) *Ray {
	var _arg0 *C.graphene_ray_t  // out
	var _arg1 *C.graphene_vec3_t // out
	var _arg2 *C.graphene_vec3_t // out
	var _cret *C.graphene_ray_t  // in

	_arg0 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(r)))
	if origin != nil {
		_arg1 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(origin)))
	}
	if direction != nil {
		_arg2 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(direction)))
	}

	_cret = C.graphene_ray_init_from_vec3(_arg0, _arg1, _arg2)
	runtime.KeepAlive(r)
	runtime.KeepAlive(origin)
	runtime.KeepAlive(direction)

	var _ray *Ray // out

	_ray = (*Ray)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _ray
}

// IntersectBox intersects the given #graphene_ray_t r with the given
// #graphene_box_t b.
func (r *Ray) IntersectBox(b *Box) (float32, RayIntersectionKind) {
	var _arg0 *C.graphene_ray_t                  // out
	var _arg1 *C.graphene_box_t                  // out
	var _arg2 C.float                            // in
	var _cret C.graphene_ray_intersection_kind_t // in

	_arg0 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_ray_intersect_box(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(r)
	runtime.KeepAlive(b)

	var _tOut float32                            // out
	var _rayIntersectionKind RayIntersectionKind // out

	_tOut = float32(_arg2)
	_rayIntersectionKind = RayIntersectionKind(_cret)

	return _tOut, _rayIntersectionKind
}

// IntersectSphere intersects the given #graphene_ray_t r with the given
// #graphene_sphere_t s.
func (r *Ray) IntersectSphere(s *Sphere) (float32, RayIntersectionKind) {
	var _arg0 *C.graphene_ray_t                  // out
	var _arg1 *C.graphene_sphere_t               // out
	var _arg2 C.float                            // in
	var _cret C.graphene_ray_intersection_kind_t // in

	_arg0 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = (*C.graphene_sphere_t)(gextras.StructNative(unsafe.Pointer(s)))

	_cret = C.graphene_ray_intersect_sphere(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(r)
	runtime.KeepAlive(s)

	var _tOut float32                            // out
	var _rayIntersectionKind RayIntersectionKind // out

	_tOut = float32(_arg2)
	_rayIntersectionKind = RayIntersectionKind(_cret)

	return _tOut, _rayIntersectionKind
}

// IntersectTriangle intersects the given #graphene_ray_t r with the given
// #graphene_triangle_t t.
func (r *Ray) IntersectTriangle(t *Triangle) (float32, RayIntersectionKind) {
	var _arg0 *C.graphene_ray_t                  // out
	var _arg1 *C.graphene_triangle_t             // out
	var _arg2 C.float                            // in
	var _cret C.graphene_ray_intersection_kind_t // in

	_arg0 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(t)))

	_cret = C.graphene_ray_intersect_triangle(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(r)
	runtime.KeepAlive(t)

	var _tOut float32                            // out
	var _rayIntersectionKind RayIntersectionKind // out

	_tOut = float32(_arg2)
	_rayIntersectionKind = RayIntersectionKind(_cret)

	return _tOut, _rayIntersectionKind
}

// IntersectsBox checks whether the given #graphene_ray_t r intersects the given
// #graphene_box_t b.
//
// See also: graphene_ray_intersect_box().
func (r *Ray) IntersectsBox(b *Box) bool {
	var _arg0 *C.graphene_ray_t // out
	var _arg1 *C.graphene_box_t // out
	var _cret C._Bool           // in

	_arg0 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_ray_intersects_box(_arg0, _arg1)
	runtime.KeepAlive(r)
	runtime.KeepAlive(b)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// IntersectsSphere checks if the given #graphene_ray_t r intersects the given
// #graphene_sphere_t s.
//
// See also: graphene_ray_intersect_sphere().
func (r *Ray) IntersectsSphere(s *Sphere) bool {
	var _arg0 *C.graphene_ray_t    // out
	var _arg1 *C.graphene_sphere_t // out
	var _cret C._Bool              // in

	_arg0 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = (*C.graphene_sphere_t)(gextras.StructNative(unsafe.Pointer(s)))

	_cret = C.graphene_ray_intersects_sphere(_arg0, _arg1)
	runtime.KeepAlive(r)
	runtime.KeepAlive(s)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// IntersectsTriangle checks whether the given #graphene_ray_t r intersects the
// given #graphene_triangle_t b.
//
// See also: graphene_ray_intersect_triangle().
func (r *Ray) IntersectsTriangle(t *Triangle) bool {
	var _arg0 *C.graphene_ray_t      // out
	var _arg1 *C.graphene_triangle_t // out
	var _cret C._Bool                // in

	_arg0 = (*C.graphene_ray_t)(gextras.StructNative(unsafe.Pointer(r)))
	_arg1 = (*C.graphene_triangle_t)(gextras.StructNative(unsafe.Pointer(t)))

	_cret = C.graphene_ray_intersects_triangle(_arg0, _arg1)
	runtime.KeepAlive(r)
	runtime.KeepAlive(t)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}
