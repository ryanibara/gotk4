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
		{T: externglib.Type(C.graphene_quaternion_get_type()), F: marshalQuaternion},
	})
}

// Quaternion: a quaternion.
//
// The contents of the #graphene_quaternion_t structure are private and should
// never be accessed directly.
type Quaternion struct {
	native C.graphene_quaternion_t
}

// WrapQuaternion wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapQuaternion(ptr unsafe.Pointer) *Quaternion {
	if ptr == nil {
		return nil
	}

	return (*Quaternion)(ptr)
}

func marshalQuaternion(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapQuaternion(unsafe.Pointer(b)), nil
}

// NewQuaternionAlloc constructs a struct Quaternion.
func NewQuaternionAlloc() *Quaternion {
	var cret *C.graphene_quaternion_t
	var ret1 *Quaternion

	cret = C.graphene_quaternion_alloc()

	ret1 = WrapQuaternion(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Quaternion) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Native returns the underlying C source pointer.
func (q *Quaternion) Native() unsafe.Pointer {
	return unsafe.Pointer(&q.native)
}

// Add adds two #graphene_quaternion_t @a and @b.
func (a *Quaternion) Add(b *Quaternion) Quaternion {
	var arg0 *C.graphene_quaternion_t
	var arg1 *C.graphene_quaternion_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_quaternion_t)(unsafe.Pointer(b.Native()))

	var arg2 C.graphene_quaternion_t
	var ret2 *Quaternion

	C.graphene_quaternion_add(arg0, b, &arg2)

	ret2 = WrapQuaternion(unsafe.Pointer(arg2))

	return ret2
}

// Dot computes the dot product of two #graphene_quaternion_t.
func (a *Quaternion) Dot(b *Quaternion) float32 {
	var arg0 *C.graphene_quaternion_t
	var arg1 *C.graphene_quaternion_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_quaternion_t)(unsafe.Pointer(b.Native()))

	var cret C.float
	var ret1 float32

	cret = C.graphene_quaternion_dot(arg0, b)

	ret1 = C.float(cret)

	return ret1
}

// Equal checks whether the given quaternions are equal.
func (a *Quaternion) Equal(b *Quaternion) bool {
	var arg0 *C.graphene_quaternion_t
	var arg1 *C.graphene_quaternion_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_quaternion_t)(unsafe.Pointer(b.Native()))

	var cret C._Bool
	var ret1 bool

	cret = C.graphene_quaternion_equal(arg0, b)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Free releases the resources allocated by graphene_quaternion_alloc().
func (q *Quaternion) Free() {
	var arg0 *C.graphene_quaternion_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))

	C.graphene_quaternion_free(arg0)
}

// Init initializes a #graphene_quaternion_t using the given four values.
func (q *Quaternion) Init(x float32, y float32, z float32, w float32) *Quaternion {
	var arg0 *C.graphene_quaternion_t
	var arg1 C.float
	var arg2 C.float
	var arg3 C.float
	var arg4 C.float

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))
	arg1 = C.float(x)
	arg2 = C.float(y)
	arg3 = C.float(z)
	arg4 = C.float(w)

	var cret *C.graphene_quaternion_t
	var ret1 *Quaternion

	cret = C.graphene_quaternion_init(arg0, x, y, z, w)

	ret1 = WrapQuaternion(unsafe.Pointer(cret))

	return ret1
}

// InitFromAngleVec3 initializes a #graphene_quaternion_t using an @angle on a
// specific @axis.
func (q *Quaternion) InitFromAngleVec3(angle float32, axis *Vec3) *Quaternion {
	var arg0 *C.graphene_quaternion_t
	var arg1 C.float
	var arg2 *C.graphene_vec3_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))
	arg1 = C.float(angle)
	arg2 = (*C.graphene_vec3_t)(unsafe.Pointer(axis.Native()))

	var cret *C.graphene_quaternion_t
	var ret1 *Quaternion

	cret = C.graphene_quaternion_init_from_angle_vec3(arg0, angle, axis)

	ret1 = WrapQuaternion(unsafe.Pointer(cret))

	return ret1
}

// InitFromAngles initializes a #graphene_quaternion_t using the values of the
// Euler angles (http://en.wikipedia.org/wiki/Euler_angles) on each axis.
//
// See also: graphene_quaternion_init_from_euler()
func (q *Quaternion) InitFromAngles(degX float32, degY float32, degZ float32) *Quaternion {
	var arg0 *C.graphene_quaternion_t
	var arg1 C.float
	var arg2 C.float
	var arg3 C.float

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))
	arg1 = C.float(degX)
	arg2 = C.float(degY)
	arg3 = C.float(degZ)

	var cret *C.graphene_quaternion_t
	var ret1 *Quaternion

	cret = C.graphene_quaternion_init_from_angles(arg0, degX, degY, degZ)

	ret1 = WrapQuaternion(unsafe.Pointer(cret))

	return ret1
}

// InitFromEuler initializes a #graphene_quaternion_t using the given
// #graphene_euler_t.
func (q *Quaternion) InitFromEuler(e *Euler) *Quaternion {
	var arg0 *C.graphene_quaternion_t
	var arg1 *C.graphene_euler_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))
	arg1 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))

	var cret *C.graphene_quaternion_t
	var ret1 *Quaternion

	cret = C.graphene_quaternion_init_from_euler(arg0, e)

	ret1 = WrapQuaternion(unsafe.Pointer(cret))

	return ret1
}

// InitFromMatrix initializes a #graphene_quaternion_t using the rotation
// components of a transformation matrix.
func (q *Quaternion) InitFromMatrix(m *Matrix) *Quaternion {
	var arg0 *C.graphene_quaternion_t
	var arg1 *C.graphene_matrix_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))
	arg1 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var cret *C.graphene_quaternion_t
	var ret1 *Quaternion

	cret = C.graphene_quaternion_init_from_matrix(arg0, m)

	ret1 = WrapQuaternion(unsafe.Pointer(cret))

	return ret1
}

// InitFromQuaternion initializes a #graphene_quaternion_t with the values from
// @src.
func (q *Quaternion) InitFromQuaternion(src *Quaternion) *Quaternion {
	var arg0 *C.graphene_quaternion_t
	var arg1 *C.graphene_quaternion_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))
	arg1 = (*C.graphene_quaternion_t)(unsafe.Pointer(src.Native()))

	var cret *C.graphene_quaternion_t
	var ret1 *Quaternion

	cret = C.graphene_quaternion_init_from_quaternion(arg0, src)

	ret1 = WrapQuaternion(unsafe.Pointer(cret))

	return ret1
}

// InitFromRadians initializes a #graphene_quaternion_t using the values of the
// Euler angles (http://en.wikipedia.org/wiki/Euler_angles) on each axis.
//
// See also: graphene_quaternion_init_from_euler()
func (q *Quaternion) InitFromRadians(radX float32, radY float32, radZ float32) *Quaternion {
	var arg0 *C.graphene_quaternion_t
	var arg1 C.float
	var arg2 C.float
	var arg3 C.float

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))
	arg1 = C.float(radX)
	arg2 = C.float(radY)
	arg3 = C.float(radZ)

	var cret *C.graphene_quaternion_t
	var ret1 *Quaternion

	cret = C.graphene_quaternion_init_from_radians(arg0, radX, radY, radZ)

	ret1 = WrapQuaternion(unsafe.Pointer(cret))

	return ret1
}

// InitFromVec4 initializes a #graphene_quaternion_t with the values from @src.
func (q *Quaternion) InitFromVec4(src *Vec4) *Quaternion {
	var arg0 *C.graphene_quaternion_t
	var arg1 *C.graphene_vec4_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))
	arg1 = (*C.graphene_vec4_t)(unsafe.Pointer(src.Native()))

	var cret *C.graphene_quaternion_t
	var ret1 *Quaternion

	cret = C.graphene_quaternion_init_from_vec4(arg0, src)

	ret1 = WrapQuaternion(unsafe.Pointer(cret))

	return ret1
}

// InitIdentity initializes a #graphene_quaternion_t using the identity
// transformation.
func (q *Quaternion) InitIdentity() *Quaternion {
	var arg0 *C.graphene_quaternion_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))

	var cret *C.graphene_quaternion_t
	var ret1 *Quaternion

	cret = C.graphene_quaternion_init_identity(arg0)

	ret1 = WrapQuaternion(unsafe.Pointer(cret))

	return ret1
}

// Invert inverts a #graphene_quaternion_t, and returns the conjugate quaternion
// of @q.
func (q *Quaternion) Invert() Quaternion {
	var arg0 *C.graphene_quaternion_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))

	var arg1 C.graphene_quaternion_t
	var ret1 *Quaternion

	C.graphene_quaternion_invert(arg0, &arg1)

	ret1 = WrapQuaternion(unsafe.Pointer(arg1))

	return ret1
}

// Multiply multiplies two #graphene_quaternion_t @a and @b.
func (a *Quaternion) Multiply(b *Quaternion) Quaternion {
	var arg0 *C.graphene_quaternion_t
	var arg1 *C.graphene_quaternion_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_quaternion_t)(unsafe.Pointer(b.Native()))

	var arg2 C.graphene_quaternion_t
	var ret2 *Quaternion

	C.graphene_quaternion_multiply(arg0, b, &arg2)

	ret2 = WrapQuaternion(unsafe.Pointer(arg2))

	return ret2
}

// Normalize normalizes a #graphene_quaternion_t.
func (q *Quaternion) Normalize() Quaternion {
	var arg0 *C.graphene_quaternion_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))

	var arg1 C.graphene_quaternion_t
	var ret1 *Quaternion

	C.graphene_quaternion_normalize(arg0, &arg1)

	ret1 = WrapQuaternion(unsafe.Pointer(arg1))

	return ret1
}

// Scale scales all the elements of a #graphene_quaternion_t @q using the given
// scalar factor.
func (q *Quaternion) Scale(factor float32) Quaternion {
	var arg0 *C.graphene_quaternion_t
	var arg1 C.float

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))
	arg1 = C.float(factor)

	var arg2 C.graphene_quaternion_t
	var ret2 *Quaternion

	C.graphene_quaternion_scale(arg0, factor, &arg2)

	ret2 = WrapQuaternion(unsafe.Pointer(arg2))

	return ret2
}

// Slerp interpolates between the two given quaternions using a spherical linear
// interpolation, or SLERP (http://en.wikipedia.org/wiki/Slerp), using the given
// interpolation @factor.
func (a *Quaternion) Slerp(b *Quaternion, factor float32) Quaternion {
	var arg0 *C.graphene_quaternion_t
	var arg1 *C.graphene_quaternion_t
	var arg2 C.float

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_quaternion_t)(unsafe.Pointer(b.Native()))
	arg2 = C.float(factor)

	var arg3 C.graphene_quaternion_t
	var ret3 *Quaternion

	C.graphene_quaternion_slerp(arg0, b, factor, &arg3)

	ret3 = WrapQuaternion(unsafe.Pointer(arg3))

	return ret3
}

// ToAngleVec3 converts a quaternion into an @angle, @axis pair.
func (q *Quaternion) ToAngleVec3() (angle float32, axis Vec3) {
	var arg0 *C.graphene_quaternion_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))

	var arg1 C.float
	var ret1 float32
	var arg2 C.graphene_vec3_t
	var ret2 *Vec3

	C.graphene_quaternion_to_angle_vec3(arg0, &arg1, &arg2)

	ret1 = C.float(arg1)
	ret2 = WrapVec3(unsafe.Pointer(arg2))

	return ret1, ret2
}

// ToAngles converts a #graphene_quaternion_t to its corresponding rotations on
// the Euler angles (http://en.wikipedia.org/wiki/Euler_angles) on each axis.
func (q *Quaternion) ToAngles() (degX float32, degY float32, degZ float32) {
	var arg0 *C.graphene_quaternion_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))

	var arg1 C.float
	var ret1 float32
	var arg2 C.float
	var ret2 float32
	var arg3 C.float
	var ret3 float32

	C.graphene_quaternion_to_angles(arg0, &arg1, &arg2, &arg3)

	ret1 = C.float(arg1)
	ret2 = C.float(arg2)
	ret3 = C.float(arg3)

	return ret1, ret2, ret3
}

// ToMatrix converts a quaternion into a transformation matrix expressing the
// rotation defined by the #graphene_quaternion_t.
func (q *Quaternion) ToMatrix() Matrix {
	var arg0 *C.graphene_quaternion_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))

	var arg1 C.graphene_matrix_t
	var ret1 *Matrix

	C.graphene_quaternion_to_matrix(arg0, &arg1)

	ret1 = WrapMatrix(unsafe.Pointer(arg1))

	return ret1
}

// ToRadians converts a #graphene_quaternion_t to its corresponding rotations on
// the Euler angles (http://en.wikipedia.org/wiki/Euler_angles) on each axis.
func (q *Quaternion) ToRadians() (radX float32, radY float32, radZ float32) {
	var arg0 *C.graphene_quaternion_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))

	var arg1 C.float
	var ret1 float32
	var arg2 C.float
	var ret2 float32
	var arg3 C.float
	var ret3 float32

	C.graphene_quaternion_to_radians(arg0, &arg1, &arg2, &arg3)

	ret1 = C.float(arg1)
	ret2 = C.float(arg2)
	ret3 = C.float(arg3)

	return ret1, ret2, ret3
}

// ToVec4 copies the components of a #graphene_quaternion_t into a
// #graphene_vec4_t.
func (q *Quaternion) ToVec4() Vec4 {
	var arg0 *C.graphene_quaternion_t

	arg0 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))

	var arg1 C.graphene_vec4_t
	var ret1 *Vec4

	C.graphene_quaternion_to_vec4(arg0, &arg1)

	ret1 = WrapVec4(unsafe.Pointer(arg1))

	return ret1
}
