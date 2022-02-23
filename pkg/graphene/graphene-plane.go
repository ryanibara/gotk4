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

// glib.Type values for graphene-plane.go.
var GTypePlane = externglib.Type(C.graphene_plane_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypePlane, F: marshalPlane},
	})
}

// Plane: 2D plane that extends infinitely in a 3D volume.
//
// The contents of the graphene_plane_t are private, and should not be modified
// directly.
//
// An instance of this type is always passed by reference.
type Plane struct {
	*plane
}

// plane is the struct that's finalized.
type plane struct {
	native *C.graphene_plane_t
}

func marshalPlane(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Plane{&plane{(*C.graphene_plane_t)(b)}}, nil
}

// NewPlaneAlloc constructs a struct Plane.
func NewPlaneAlloc() *Plane {
	var _cret *C.graphene_plane_t // in

	_cret = C.graphene_plane_alloc()

	var _plane *Plane // out

	_plane = (*Plane)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_plane)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.graphene_plane_free((*C.graphene_plane_t)(intern.C))
		},
	)

	return _plane
}

// Distance computes the distance of point from a #graphene_plane_t.
//
// The function takes the following parameters:
//
//    - point: #graphene_point3d_t.
//
// The function returns the following values:
//
//    - gfloat: distance of the given #graphene_point3d_t from the plane.
//
func (p *Plane) Distance(point *Point3D) float32 {
	var _arg0 *C.graphene_plane_t   // out
	var _arg1 *C.graphene_point3d_t // out
	var _cret C.float               // in

	_arg0 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(point)))

	_cret = C.graphene_plane_distance(_arg0, _arg1)
	runtime.KeepAlive(p)
	runtime.KeepAlive(point)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Equal checks whether the two given #graphene_plane_t are equal.
//
// The function takes the following parameters:
//
//    - b: #graphene_plane_t.
//
// The function returns the following values:
//
//    - ok: true if the given planes are equal.
//
func (a *Plane) Equal(b *Plane) bool {
	var _arg0 *C.graphene_plane_t // out
	var _arg1 *C.graphene_plane_t // out
	var _cret C._Bool             // in

	_arg0 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_plane_equal(_arg0, _arg1)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Constant retrieves the distance along the normal vector of the given
// #graphene_plane_t from the origin.
//
// The function returns the following values:
//
//    - gfloat: constant value of the plane.
//
func (p *Plane) Constant() float32 {
	var _arg0 *C.graphene_plane_t // out
	var _cret C.float             // in

	_arg0 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p)))

	_cret = C.graphene_plane_get_constant(_arg0)
	runtime.KeepAlive(p)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Normal retrieves the normal vector pointing towards the origin of the given
// #graphene_plane_t.
//
// The function returns the following values:
//
//    - normal: return location for the normal vector.
//
func (p *Plane) Normal() *Vec3 {
	var _arg0 *C.graphene_plane_t // out
	var _arg1 C.graphene_vec3_t   // in

	_arg0 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p)))

	C.graphene_plane_get_normal(_arg0, &_arg1)
	runtime.KeepAlive(p)

	var _normal *Vec3 // out

	_normal = (*Vec3)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _normal
}

// Init initializes the given #graphene_plane_t using the given normal vector
// and constant values.
//
// The function takes the following parameters:
//
//    - normal (optional): unit length normal vector defining the plane pointing
//      towards the origin; if unset, we use the X axis by default.
//    - constant: distance from the origin to the plane along the normal vector;
//      the sign determines the half-space occupied by the plane.
//
// The function returns the following values:
//
//    - plane: initialized plane.
//
func (p *Plane) Init(normal *Vec3, constant float32) *Plane {
	var _arg0 *C.graphene_plane_t // out
	var _arg1 *C.graphene_vec3_t  // out
	var _arg2 C.float             // out
	var _cret *C.graphene_plane_t // in

	_arg0 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p)))
	if normal != nil {
		_arg1 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(normal)))
	}
	_arg2 = C.float(constant)

	_cret = C.graphene_plane_init(_arg0, _arg1, _arg2)
	runtime.KeepAlive(p)
	runtime.KeepAlive(normal)
	runtime.KeepAlive(constant)

	var _plane *Plane // out

	_plane = (*Plane)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _plane
}

// InitFromPlane initializes the given #graphene_plane_t using the normal vector
// and constant of another #graphene_plane_t.
//
// The function takes the following parameters:
//
//    - src: #graphene_plane_t.
//
// The function returns the following values:
//
//    - plane: initialized plane.
//
func (p *Plane) InitFromPlane(src *Plane) *Plane {
	var _arg0 *C.graphene_plane_t // out
	var _arg1 *C.graphene_plane_t // out
	var _cret *C.graphene_plane_t // in

	_arg0 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p)))
	_arg1 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.graphene_plane_init_from_plane(_arg0, _arg1)
	runtime.KeepAlive(p)
	runtime.KeepAlive(src)

	var _plane *Plane // out

	_plane = (*Plane)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _plane
}

// InitFromPoint initializes the given #graphene_plane_t using the given normal
// vector and an arbitrary co-planar point.
//
// The function takes the following parameters:
//
//    - normal vector defining the plane pointing towards the origin.
//    - point: #graphene_point3d_t.
//
// The function returns the following values:
//
//    - plane: initialized plane.
//
func (p *Plane) InitFromPoint(normal *Vec3, point *Point3D) *Plane {
	var _arg0 *C.graphene_plane_t   // out
	var _arg1 *C.graphene_vec3_t    // out
	var _arg2 *C.graphene_point3d_t // out
	var _cret *C.graphene_plane_t   // in

	_arg0 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p)))
	_arg1 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(normal)))
	_arg2 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(point)))

	_cret = C.graphene_plane_init_from_point(_arg0, _arg1, _arg2)
	runtime.KeepAlive(p)
	runtime.KeepAlive(normal)
	runtime.KeepAlive(point)

	var _plane *Plane // out

	_plane = (*Plane)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _plane
}

// InitFromPoints initializes the given #graphene_plane_t using the 3 provided
// co-planar points.
//
// The winding order is counter-clockwise, and determines which direction the
// normal vector will point.
//
// The function takes the following parameters:
//
//    - a: #graphene_point3d_t.
//    - b: #graphene_point3d_t.
//    - c: #graphene_point3d_t.
//
// The function returns the following values:
//
//    - plane: initialized plane.
//
func (p *Plane) InitFromPoints(a *Point3D, b *Point3D, c *Point3D) *Plane {
	var _arg0 *C.graphene_plane_t   // out
	var _arg1 *C.graphene_point3d_t // out
	var _arg2 *C.graphene_point3d_t // out
	var _arg3 *C.graphene_point3d_t // out
	var _cret *C.graphene_plane_t   // in

	_arg0 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg2 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(b)))
	_arg3 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(c)))

	_cret = C.graphene_plane_init_from_points(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(p)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
	runtime.KeepAlive(c)

	var _plane *Plane // out

	_plane = (*Plane)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _plane
}

// InitFromVec4 initializes the given #graphene_plane_t using the components of
// the given #graphene_vec4_t vector.
//
// The function takes the following parameters:
//
//    - src containing the normal vector in its first three components, and the
//      distance in its fourth component.
//
// The function returns the following values:
//
//    - plane: initialized plane.
//
func (p *Plane) InitFromVec4(src *Vec4) *Plane {
	var _arg0 *C.graphene_plane_t // out
	var _arg1 *C.graphene_vec4_t  // out
	var _cret *C.graphene_plane_t // in

	_arg0 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p)))
	_arg1 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.graphene_plane_init_from_vec4(_arg0, _arg1)
	runtime.KeepAlive(p)
	runtime.KeepAlive(src)

	var _plane *Plane // out

	_plane = (*Plane)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _plane
}

// Negate negates the normal vector and constant of a #graphene_plane_t,
// effectively mirroring the plane across the origin.
//
// The function returns the following values:
//
//    - res: return location for the negated plane.
//
func (p *Plane) Negate() *Plane {
	var _arg0 *C.graphene_plane_t // out
	var _arg1 C.graphene_plane_t  // in

	_arg0 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p)))

	C.graphene_plane_negate(_arg0, &_arg1)
	runtime.KeepAlive(p)

	var _res *Plane // out

	_res = (*Plane)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// Normalize normalizes the vector of the given #graphene_plane_t, and adjusts
// the constant accordingly.
//
// The function returns the following values:
//
//    - res: return location for the normalized plane.
//
func (p *Plane) Normalize() *Plane {
	var _arg0 *C.graphene_plane_t // out
	var _arg1 C.graphene_plane_t  // in

	_arg0 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p)))

	C.graphene_plane_normalize(_arg0, &_arg1)
	runtime.KeepAlive(p)

	var _res *Plane // out

	_res = (*Plane)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// Transform transforms a #graphene_plane_t p using the given matrix and
// normal_matrix.
//
// If normal_matrix is NULL, a transformation matrix for the plane normal will
// be computed from matrix. If you are transforming multiple planes using the
// same matrix it's recommended to compute the normal matrix beforehand to avoid
// incurring in the cost of recomputing it every time.
//
// The function takes the following parameters:
//
//    - matrix: #graphene_matrix_t.
//    - normalMatrix (optional): #graphene_matrix_t.
//
// The function returns the following values:
//
//    - res: transformed plane.
//
func (p *Plane) Transform(matrix *Matrix, normalMatrix *Matrix) *Plane {
	var _arg0 *C.graphene_plane_t  // out
	var _arg1 *C.graphene_matrix_t // out
	var _arg2 *C.graphene_matrix_t // out
	var _arg3 C.graphene_plane_t   // in

	_arg0 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p)))
	_arg1 = (*C.graphene_matrix_t)(gextras.StructNative(unsafe.Pointer(matrix)))
	if normalMatrix != nil {
		_arg2 = (*C.graphene_matrix_t)(gextras.StructNative(unsafe.Pointer(normalMatrix)))
	}

	C.graphene_plane_transform(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(p)
	runtime.KeepAlive(matrix)
	runtime.KeepAlive(normalMatrix)

	var _res *Plane // out

	_res = (*Plane)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _res
}
