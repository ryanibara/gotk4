// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdbool.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_euler_get_type()), F: marshalEuler},
	})
}

// EulerOrder: specify the order of the rotations on each axis.
//
// The GRAPHENE_EULER_ORDER_DEFAULT value is special, and is used as an alias
// for one of the other orders.
type EulerOrder int

const (
	// EulerOrderDefault: rotate in the default order; the default order is one
	// of the following enumeration values
	EulerOrderDefault EulerOrder = -1
	// EulerOrderXYZ: rotate in the X, Y, and Z order. Deprecated in Graphene
	// 1.10, it's an alias for GRAPHENE_EULER_ORDER_SXYZ
	EulerOrderXYZ EulerOrder = 0
	// EulerOrderYZX: rotate in the Y, Z, and X order. Deprecated in Graphene
	// 1.10, it's an alias for GRAPHENE_EULER_ORDER_SYZX
	EulerOrderYZX EulerOrder = 1
	// EulerOrderZXY: rotate in the Z, X, and Y order. Deprecated in Graphene
	// 1.10, it's an alias for GRAPHENE_EULER_ORDER_SZXY
	EulerOrderZXY EulerOrder = 2
	// EulerOrderXZY: rotate in the X, Z, and Y order. Deprecated in Graphene
	// 1.10, it's an alias for GRAPHENE_EULER_ORDER_SXZY
	EulerOrderXZY EulerOrder = 3
	// EulerOrderYXZ: rotate in the Y, X, and Z order. Deprecated in Graphene
	// 1.10, it's an alias for GRAPHENE_EULER_ORDER_SYXZ
	EulerOrderYXZ EulerOrder = 4
	// EulerOrderZYX: rotate in the Z, Y, and X order. Deprecated in Graphene
	// 1.10, it's an alias for GRAPHENE_EULER_ORDER_SZYX
	EulerOrderZYX EulerOrder = 5
	// EulerOrderSXYZ defines a static rotation along the X, Y, and Z axes
	// (Since: 1.10)
	EulerOrderSXYZ EulerOrder = 6
	// EulerOrderSXYX defines a static rotation along the X, Y, and X axes
	// (Since: 1.10)
	EulerOrderSXYX EulerOrder = 7
	// EulerOrderSXZY defines a static rotation along the X, Z, and Y axes
	// (Since: 1.10)
	EulerOrderSXZY EulerOrder = 8
	// EulerOrderSXZX defines a static rotation along the X, Z, and X axes
	// (Since: 1.10)
	EulerOrderSXZX EulerOrder = 9
	// EulerOrderSYZX defines a static rotation along the Y, Z, and X axes
	// (Since: 1.10)
	EulerOrderSYZX EulerOrder = 10
	// EulerOrderSYZY defines a static rotation along the Y, Z, and Y axes
	// (Since: 1.10)
	EulerOrderSYZY EulerOrder = 11
	// EulerOrderSYXZ defines a static rotation along the Y, X, and Z axes
	// (Since: 1.10)
	EulerOrderSYXZ EulerOrder = 12
	// EulerOrderSYXY defines a static rotation along the Y, X, and Y axes
	// (Since: 1.10)
	EulerOrderSYXY EulerOrder = 13
	// EulerOrderSZXY defines a static rotation along the Z, X, and Y axes
	// (Since: 1.10)
	EulerOrderSZXY EulerOrder = 14
	// EulerOrderSZXZ defines a static rotation along the Z, X, and Z axes
	// (Since: 1.10)
	EulerOrderSZXZ EulerOrder = 15
	// EulerOrderSZYX defines a static rotation along the Z, Y, and X axes
	// (Since: 1.10)
	EulerOrderSZYX EulerOrder = 16
	// EulerOrderSZYZ defines a static rotation along the Z, Y, and Z axes
	// (Since: 1.10)
	EulerOrderSZYZ EulerOrder = 17
	// EulerOrderRZYX defines a relative rotation along the Z, Y, and X axes
	// (Since: 1.10)
	EulerOrderRZYX EulerOrder = 18
	// EulerOrderRXYX defines a relative rotation along the X, Y, and X axes
	// (Since: 1.10)
	EulerOrderRXYX EulerOrder = 19
	// EulerOrderRYZX defines a relative rotation along the Y, Z, and X axes
	// (Since: 1.10)
	EulerOrderRYZX EulerOrder = 20
	// EulerOrderRXZX defines a relative rotation along the X, Z, and X axes
	// (Since: 1.10)
	EulerOrderRXZX EulerOrder = 21
	// EulerOrderRXZY defines a relative rotation along the X, Z, and Y axes
	// (Since: 1.10)
	EulerOrderRXZY EulerOrder = 22
	// EulerOrderRYZY defines a relative rotation along the Y, Z, and Y axes
	// (Since: 1.10)
	EulerOrderRYZY EulerOrder = 23
	// EulerOrderRZXY defines a relative rotation along the Z, X, and Y axes
	// (Since: 1.10)
	EulerOrderRZXY EulerOrder = 24
	// EulerOrderRYXY defines a relative rotation along the Y, X, and Y axes
	// (Since: 1.10)
	EulerOrderRYXY EulerOrder = 25
	// EulerOrderRYXZ defines a relative rotation along the Y, X, and Z axes
	// (Since: 1.10)
	EulerOrderRYXZ EulerOrder = 26
	// EulerOrderRZXZ defines a relative rotation along the Z, X, and Z axes
	// (Since: 1.10)
	EulerOrderRZXZ EulerOrder = 27
	// EulerOrderRXYZ defines a relative rotation along the X, Y, and Z axes
	// (Since: 1.10)
	EulerOrderRXYZ EulerOrder = 28
	// EulerOrderRZYZ defines a relative rotation along the Z, Y, and Z axes
	// (Since: 1.10)
	EulerOrderRZYZ EulerOrder = 29
)

// Euler: describe a rotation using Euler angles.
//
// The contents of the #graphene_euler_t structure are private and should never
// be accessed directly.
type Euler struct {
	native C.graphene_euler_t
}

// WrapEuler wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEuler(ptr unsafe.Pointer) *Euler {
	if ptr == nil {
		return nil
	}

	return (*Euler)(ptr)
}

func marshalEuler(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapEuler(unsafe.Pointer(b)), nil
}

// NewEulerAlloc constructs a struct Euler.
func NewEulerAlloc() *Euler {
	var cret *C.graphene_euler_t
	var ret1 *Euler

	cret = C.graphene_euler_alloc()

	ret1 = WrapEuler(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Euler) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Native returns the underlying C source pointer.
func (e *Euler) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// Equal checks if two #graphene_euler_t are equal.
func (a *Euler) Equal(b *Euler) bool {
	var arg0 *C.graphene_euler_t
	var arg1 *C.graphene_euler_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_euler_t)(unsafe.Pointer(b.Native()))

	var cret C._Bool
	var ret1 bool

	cret = C.graphene_euler_equal(arg0, b)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Free frees the resources allocated by graphene_euler_alloc().
func (e *Euler) Free() {
	var arg0 *C.graphene_euler_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))

	C.graphene_euler_free(arg0)
}

// Alpha retrieves the first component of the Euler angle vector, depending on
// the order of rotation.
//
// See also: graphene_euler_get_x()
func (e *Euler) Alpha() float32 {
	var arg0 *C.graphene_euler_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))

	var cret C.float
	var ret1 float32

	cret = C.graphene_euler_get_alpha(arg0)

	ret1 = C.float(cret)

	return ret1
}

// Beta retrieves the second component of the Euler angle vector, depending on
// the order of rotation.
//
// See also: graphene_euler_get_y()
func (e *Euler) Beta() float32 {
	var arg0 *C.graphene_euler_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))

	var cret C.float
	var ret1 float32

	cret = C.graphene_euler_get_beta(arg0)

	ret1 = C.float(cret)

	return ret1
}

// Gamma retrieves the third component of the Euler angle vector, depending on
// the order of rotation.
//
// See also: graphene_euler_get_z()
func (e *Euler) Gamma() float32 {
	var arg0 *C.graphene_euler_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))

	var cret C.float
	var ret1 float32

	cret = C.graphene_euler_get_gamma(arg0)

	ret1 = C.float(cret)

	return ret1
}

// Order retrieves the order used to apply the rotations described in the
// #graphene_euler_t structure, when converting to and from other structures,
// like #graphene_quaternion_t and #graphene_matrix_t.
//
// This function does not return the GRAPHENE_EULER_ORDER_DEFAULT enumeration
// value; it will return the effective order of rotation instead.
func (e *Euler) Order() EulerOrder {
	var arg0 *C.graphene_euler_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))

	var cret C.graphene_euler_order_t
	var ret1 EulerOrder

	cret = C.graphene_euler_get_order(arg0)

	ret1 = EulerOrder(cret)

	return ret1
}

// X retrieves the rotation angle on the X axis, in degrees.
func (e *Euler) X() float32 {
	var arg0 *C.graphene_euler_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))

	var cret C.float
	var ret1 float32

	cret = C.graphene_euler_get_x(arg0)

	ret1 = C.float(cret)

	return ret1
}

// Y retrieves the rotation angle on the Y axis, in degrees.
func (e *Euler) Y() float32 {
	var arg0 *C.graphene_euler_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))

	var cret C.float
	var ret1 float32

	cret = C.graphene_euler_get_y(arg0)

	ret1 = C.float(cret)

	return ret1
}

// Z retrieves the rotation angle on the Z axis, in degrees.
func (e *Euler) Z() float32 {
	var arg0 *C.graphene_euler_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))

	var cret C.float
	var ret1 float32

	cret = C.graphene_euler_get_z(arg0)

	ret1 = C.float(cret)

	return ret1
}

// Init initializes a #graphene_euler_t using the given angles.
//
// The order of the rotations is GRAPHENE_EULER_ORDER_DEFAULT.
func (e *Euler) Init(x float32, y float32, z float32) *Euler {
	var arg0 *C.graphene_euler_t
	var arg1 C.float
	var arg2 C.float
	var arg3 C.float

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))
	arg1 = C.float(x)
	arg2 = C.float(y)
	arg3 = C.float(z)

	var cret *C.graphene_euler_t
	var ret1 *Euler

	cret = C.graphene_euler_init(arg0, x, y, z)

	ret1 = WrapEuler(unsafe.Pointer(cret))

	return ret1
}

// InitFromEuler initializes a #graphene_euler_t using the angles and order of
// another #graphene_euler_t.
//
// If the #graphene_euler_t @src is nil, this function is equivalent to calling
// graphene_euler_init() with all angles set to 0.
func (e *Euler) InitFromEuler(src *Euler) *Euler {
	var arg0 *C.graphene_euler_t
	var arg1 *C.graphene_euler_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))
	arg1 = (*C.graphene_euler_t)(unsafe.Pointer(src.Native()))

	var cret *C.graphene_euler_t
	var ret1 *Euler

	cret = C.graphene_euler_init_from_euler(arg0, src)

	ret1 = WrapEuler(unsafe.Pointer(cret))

	return ret1
}

// InitFromMatrix initializes a #graphene_euler_t using the given rotation
// matrix.
//
// If the #graphene_matrix_t @m is nil, the #graphene_euler_t will be
// initialized with all angles set to 0.
func (e *Euler) InitFromMatrix(m *Matrix, order EulerOrder) *Euler {
	var arg0 *C.graphene_euler_t
	var arg1 *C.graphene_matrix_t
	var arg2 C.graphene_euler_order_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))
	arg1 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	arg2 = (C.graphene_euler_order_t)(order)

	var cret *C.graphene_euler_t
	var ret1 *Euler

	cret = C.graphene_euler_init_from_matrix(arg0, m, order)

	ret1 = WrapEuler(unsafe.Pointer(cret))

	return ret1
}

// InitFromQuaternion initializes a #graphene_euler_t using the given normalized
// quaternion.
//
// If the #graphene_quaternion_t @q is nil, the #graphene_euler_t will be
// initialized with all angles set to 0.
func (e *Euler) InitFromQuaternion(q *Quaternion, order EulerOrder) *Euler {
	var arg0 *C.graphene_euler_t
	var arg1 *C.graphene_quaternion_t
	var arg2 C.graphene_euler_order_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))
	arg1 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))
	arg2 = (C.graphene_euler_order_t)(order)

	var cret *C.graphene_euler_t
	var ret1 *Euler

	cret = C.graphene_euler_init_from_quaternion(arg0, q, order)

	ret1 = WrapEuler(unsafe.Pointer(cret))

	return ret1
}

// InitFromRadians initializes a #graphene_euler_t using the given angles and
// order of rotation.
func (e *Euler) InitFromRadians(x float32, y float32, z float32, order EulerOrder) *Euler {
	var arg0 *C.graphene_euler_t
	var arg1 C.float
	var arg2 C.float
	var arg3 C.float
	var arg4 C.graphene_euler_order_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))
	arg1 = C.float(x)
	arg2 = C.float(y)
	arg3 = C.float(z)
	arg4 = (C.graphene_euler_order_t)(order)

	var cret *C.graphene_euler_t
	var ret1 *Euler

	cret = C.graphene_euler_init_from_radians(arg0, x, y, z, order)

	ret1 = WrapEuler(unsafe.Pointer(cret))

	return ret1
}

// InitFromVec3 initializes a #graphene_euler_t using the angles contained in a
// #graphene_vec3_t.
//
// If the #graphene_vec3_t @v is nil, the #graphene_euler_t will be initialized
// with all angles set to 0.
func (e *Euler) InitFromVec3(v *Vec3, order EulerOrder) *Euler {
	var arg0 *C.graphene_euler_t
	var arg1 *C.graphene_vec3_t
	var arg2 C.graphene_euler_order_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))
	arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(v.Native()))
	arg2 = (C.graphene_euler_order_t)(order)

	var cret *C.graphene_euler_t
	var ret1 *Euler

	cret = C.graphene_euler_init_from_vec3(arg0, v, order)

	ret1 = WrapEuler(unsafe.Pointer(cret))

	return ret1
}

// InitWithOrder initializes a #graphene_euler_t with the given angles and
// @order.
func (e *Euler) InitWithOrder(x float32, y float32, z float32, order EulerOrder) *Euler {
	var arg0 *C.graphene_euler_t
	var arg1 C.float
	var arg2 C.float
	var arg3 C.float
	var arg4 C.graphene_euler_order_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))
	arg1 = C.float(x)
	arg2 = C.float(y)
	arg3 = C.float(z)
	arg4 = (C.graphene_euler_order_t)(order)

	var cret *C.graphene_euler_t
	var ret1 *Euler

	cret = C.graphene_euler_init_with_order(arg0, x, y, z, order)

	ret1 = WrapEuler(unsafe.Pointer(cret))

	return ret1
}

// Reorder reorders a #graphene_euler_t using @order.
//
// This function is equivalent to creating a #graphene_quaternion_t from the
// given #graphene_euler_t, and then converting the quaternion into another
// #graphene_euler_t.
func (e *Euler) Reorder(order EulerOrder) Euler {
	var arg0 *C.graphene_euler_t
	var arg1 C.graphene_euler_order_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))
	arg1 = (C.graphene_euler_order_t)(order)

	var arg2 C.graphene_euler_t
	var ret2 *Euler

	C.graphene_euler_reorder(arg0, order, &arg2)

	*ret2 = WrapEuler(unsafe.Pointer(arg2))

	return ret2
}

// ToMatrix converts a #graphene_euler_t into a transformation matrix expressing
// the extrinsic composition of rotations described by the Euler angles.
//
// The rotations are applied over the reference frame axes in the order
// associated with the #graphene_euler_t; for instance, if the order used to
// initialize @e is GRAPHENE_EULER_ORDER_XYZ:
//
//    * the first rotation moves the body around the X axis with
//      an angle φ
//    * the second rotation moves the body around the Y axis with
//      an angle of ϑ
//    * the third rotation moves the body around the Z axis with
//      an angle of ψ
//
// The rotation sign convention is right-handed, to preserve compatibility
// between Euler-based, quaternion-based, and angle-axis-based rotations.
func (e *Euler) ToMatrix() Matrix {
	var arg0 *C.graphene_euler_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))

	var arg1 C.graphene_matrix_t
	var ret1 *Matrix

	C.graphene_euler_to_matrix(arg0, &arg1)

	*ret1 = WrapMatrix(unsafe.Pointer(arg1))

	return ret1
}

// ToQuaternion converts a #graphene_euler_t into a #graphene_quaternion_t.
func (e *Euler) ToQuaternion() Quaternion {
	var arg0 *C.graphene_euler_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))

	var arg1 C.graphene_quaternion_t
	var ret1 *Quaternion

	C.graphene_euler_to_quaternion(arg0, &arg1)

	*ret1 = WrapQuaternion(unsafe.Pointer(arg1))

	return ret1
}

// ToVec3 retrieves the angles of a #graphene_euler_t and initializes a
// #graphene_vec3_t with them.
func (e *Euler) ToVec3() Vec3 {
	var arg0 *C.graphene_euler_t

	arg0 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))

	var arg1 C.graphene_vec3_t
	var ret1 *Vec3

	C.graphene_euler_to_vec3(arg0, &arg1)

	*ret1 = WrapVec3(unsafe.Pointer(arg1))

	return ret1
}
