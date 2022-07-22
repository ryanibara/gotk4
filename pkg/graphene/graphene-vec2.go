// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

// GType values.
var (
	GTypeVec2 = coreglib.Type(C.graphene_vec2_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVec2, F: marshalVec2},
	})
}

// Vec2: structure capable of holding a vector with two dimensions, x and y.
//
// The contents of the #graphene_vec2_t structure are private and should never
// be accessed directly.
//
// An instance of this type is always passed by reference.
type Vec2 struct {
	*vec2
}

// vec2 is the struct that's finalized.
type vec2 struct {
	native *C.graphene_vec2_t
}

func marshalVec2(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Vec2{&vec2{(*C.graphene_vec2_t)(b)}}, nil
}

// NewVec2Alloc constructs a struct Vec2.
func NewVec2Alloc() *Vec2 {
	var _cret *C.graphene_vec2_t // in

	_cret = C.graphene_vec2_alloc()

	var _vec2 *Vec2 // out

	_vec2 = (*Vec2)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_vec2)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.graphene_vec2_free((*C.graphene_vec2_t)(intern.C))
		},
	)

	return _vec2
}

// Add adds each component of the two passed vectors and places each result into
// the components of res.
//
// The function takes the following parameters:
//
//    - b: #graphene_vec2_t.
//
// The function returns the following values:
//
//    - res: return location for the result.
//
func (a *Vec2) Add(b *Vec2) *Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_vec2_add(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res *Vec2 // out

	_res = (*Vec2)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Divide divides each component of the first operand a by the corresponding
// component of the second operand b, and places the results into the vector
// res.
//
// The function takes the following parameters:
//
//    - b: #graphene_vec2_t.
//
// The function returns the following values:
//
//    - res: return location for the result.
//
func (a *Vec2) Divide(b *Vec2) *Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_vec2_divide(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res *Vec2 // out

	_res = (*Vec2)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Dot computes the dot product of the two given vectors.
//
// The function takes the following parameters:
//
//    - b: #graphene_vec2_t.
//
// The function returns the following values:
//
//    - gfloat: dot product of the vectors.
//
func (a *Vec2) Dot(b *Vec2) float32 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_vec2_dot(_arg0, _arg1)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Equal checks whether the two given #graphene_vec2_t are equal.
//
// The function takes the following parameters:
//
//    - v2: #graphene_vec2_t.
//
// The function returns the following values:
//
//    - ok: true if the two vectors are equal, and false otherwise.
//
func (v1 *Vec2) Equal(v2 *Vec2) bool {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _cret C._Bool            // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(v1)))
	_arg1 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(v2)))

	_cret = C.graphene_vec2_equal(_arg0, _arg1)
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// X retrieves the X component of the #graphene_vec2_t.
//
// The function returns the following values:
//
//    - gfloat: value of the X component.
//
func (v *Vec2) X() float32 {
	var _arg0 *C.graphene_vec2_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(v)))

	_cret = C.graphene_vec2_get_x(_arg0)
	runtime.KeepAlive(v)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Y retrieves the Y component of the #graphene_vec2_t.
//
// The function returns the following values:
//
//    - gfloat: value of the Y component.
//
func (v *Vec2) Y() float32 {
	var _arg0 *C.graphene_vec2_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(v)))

	_cret = C.graphene_vec2_get_y(_arg0)
	runtime.KeepAlive(v)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Init initializes a #graphene_vec2_t using the given values.
//
// This function can be called multiple times.
//
// The function takes the following parameters:
//
//    - x: x field of the vector.
//    - y: y field of the vector.
//
// The function returns the following values:
//
//    - vec2: initialized vector.
//
func (v *Vec2) Init(x float32, y float32) *Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out
	var _cret *C.graphene_vec2_t // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(v)))
	_arg1 = C.float(x)
	_arg2 = C.float(y)

	_cret = C.graphene_vec2_init(_arg0, _arg1, _arg2)
	runtime.KeepAlive(v)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _vec2 *Vec2 // out

	_vec2 = (*Vec2)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec2
}

// InitFromFloat initializes v with the contents of the given array.
//
// The function takes the following parameters:
//
//    - src: array of floating point values with at least two elements.
//
// The function returns the following values:
//
//    - vec2: initialized vector.
//
func (v *Vec2) InitFromFloat(src [2]float32) *Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.float           // out
	var _cret *C.graphene_vec2_t // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(v)))
	_arg1 = (*C.float)(unsafe.Pointer(&src))

	_cret = C.graphene_vec2_init_from_float(_arg0, _arg1)
	runtime.KeepAlive(v)
	runtime.KeepAlive(src)

	var _vec2 *Vec2 // out

	_vec2 = (*Vec2)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec2
}

// InitFromVec2 copies the contents of src into v.
//
// The function takes the following parameters:
//
//    - src: #graphene_vec2_t.
//
// The function returns the following values:
//
//    - vec2: initialized vector.
//
func (v *Vec2) InitFromVec2(src *Vec2) *Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _cret *C.graphene_vec2_t // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(v)))
	_arg1 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.graphene_vec2_init_from_vec2(_arg0, _arg1)
	runtime.KeepAlive(v)
	runtime.KeepAlive(src)

	var _vec2 *Vec2 // out

	_vec2 = (*Vec2)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec2
}

// Interpolate: linearly interpolates v1 and v2 using the given factor.
//
// The function takes the following parameters:
//
//    - v2: #graphene_vec2_t.
//    - factor: interpolation factor.
//
// The function returns the following values:
//
//    - res: interpolated vector.
//
func (v1 *Vec2) Interpolate(v2 *Vec2, factor float64) *Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.double           // out
	var _arg3 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(v1)))
	_arg1 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(v2)))
	_arg2 = C.double(factor)

	C.graphene_vec2_interpolate(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)
	runtime.KeepAlive(factor)

	var _res *Vec2 // out

	_res = (*Vec2)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _res
}

// Length computes the length of the given vector.
//
// The function returns the following values:
//
//    - gfloat: length of the vector.
//
func (v *Vec2) Length() float32 {
	var _arg0 *C.graphene_vec2_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(v)))

	_cret = C.graphene_vec2_length(_arg0)
	runtime.KeepAlive(v)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Max compares the two given vectors and places the maximum values of each
// component into res.
//
// The function takes the following parameters:
//
//    - b: #graphene_vec2_t.
//
// The function returns the following values:
//
//    - res: resulting vector.
//
func (a *Vec2) Max(b *Vec2) *Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_vec2_max(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res *Vec2 // out

	_res = (*Vec2)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Min compares the two given vectors and places the minimum values of each
// component into res.
//
// The function takes the following parameters:
//
//    - b: #graphene_vec2_t.
//
// The function returns the following values:
//
//    - res: resulting vector.
//
func (a *Vec2) Min(b *Vec2) *Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_vec2_min(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res *Vec2 // out

	_res = (*Vec2)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Multiply multiplies each component of the two passed vectors and places each
// result into the components of res.
//
// The function takes the following parameters:
//
//    - b: #graphene_vec2_t.
//
// The function returns the following values:
//
//    - res: return location for the result.
//
func (a *Vec2) Multiply(b *Vec2) *Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_vec2_multiply(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res *Vec2 // out

	_res = (*Vec2)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Near compares the two given #graphene_vec2_t vectors and checks whether their
// values are within the given epsilon.
//
// The function takes the following parameters:
//
//    - v2: #graphene_vec2_t.
//    - epsilon: threshold between the two vectors.
//
// The function returns the following values:
//
//    - ok: true if the two vectors are near each other.
//
func (v1 *Vec2) Near(v2 *Vec2, epsilon float32) bool {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.float            // out
	var _cret C._Bool            // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(v1)))
	_arg1 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(v2)))
	_arg2 = C.float(epsilon)

	_cret = C.graphene_vec2_near(_arg0, _arg1, _arg2)
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)
	runtime.KeepAlive(epsilon)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Negate negates the given #graphene_vec2_t.
//
// The function returns the following values:
//
//    - res: return location for the result vector.
//
func (v *Vec2) Negate() *Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(v)))

	C.graphene_vec2_negate(_arg0, &_arg1)
	runtime.KeepAlive(v)

	var _res *Vec2 // out

	_res = (*Vec2)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// Normalize computes the normalized vector for the given vector v.
//
// The function returns the following values:
//
//    - res: return location for the normalized vector.
//
func (v *Vec2) Normalize() *Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(v)))

	C.graphene_vec2_normalize(_arg0, &_arg1)
	runtime.KeepAlive(v)

	var _res *Vec2 // out

	_res = (*Vec2)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// Scale multiplies all components of the given vector with the given scalar
// factor.
//
// The function takes the following parameters:
//
//    - factor: scalar factor.
//
// The function returns the following values:
//
//    - res: return location for the result vector.
//
func (v *Vec2) Scale(factor float32) *Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 C.float            // out
	var _arg2 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(v)))
	_arg1 = C.float(factor)

	C.graphene_vec2_scale(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(v)
	runtime.KeepAlive(factor)

	var _res *Vec2 // out

	_res = (*Vec2)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Subtract subtracts from each component of the first operand a the
// corresponding component of the second operand b and places each result into
// the components of res.
//
// The function takes the following parameters:
//
//    - b: #graphene_vec2_t.
//
// The function returns the following values:
//
//    - res: return location for the result.
//
func (a *Vec2) Subtract(b *Vec2) *Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_vec2_subtract(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res *Vec2 // out

	_res = (*Vec2)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// ToFloat stores the components of v into an array.
//
// The function returns the following values:
//
//    - dest: return location for an array of floating point values with at least
//      2 elements.
//
func (v *Vec2) ToFloat() [2]float32 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 [2]C.float         // in

	_arg0 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(v)))

	C.graphene_vec2_to_float(_arg0, &_arg1[0])
	runtime.KeepAlive(v)

	var _dest [2]float32 // out

	_dest = *(*[2]float32)(unsafe.Pointer(&_arg1))

	return _dest
}

// Vec2One retrieves a constant vector with (1, 1) components.
//
// The function returns the following values:
//
//    - vec2: one vector.
//
func Vec2One() *Vec2 {
	var _cret *C.graphene_vec2_t // in

	_cret = C.graphene_vec2_one()

	var _vec2 *Vec2 // out

	_vec2 = (*Vec2)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec2
}

// Vec2XAxis retrieves a constant vector with (1, 0) components.
//
// The function returns the following values:
//
//    - vec2: x axis vector.
//
func Vec2XAxis() *Vec2 {
	var _cret *C.graphene_vec2_t // in

	_cret = C.graphene_vec2_x_axis()

	var _vec2 *Vec2 // out

	_vec2 = (*Vec2)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec2
}

// Vec2YAxis retrieves a constant vector with (0, 1) components.
//
// The function returns the following values:
//
//    - vec2: y axis vector.
//
func Vec2YAxis() *Vec2 {
	var _cret *C.graphene_vec2_t // in

	_cret = C.graphene_vec2_y_axis()

	var _vec2 *Vec2 // out

	_vec2 = (*Vec2)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec2
}

// Vec2Zero retrieves a constant vector with (0, 0) components.
//
// The function returns the following values:
//
//    - vec2: zero vector.
//
func Vec2Zero() *Vec2 {
	var _cret *C.graphene_vec2_t // in

	_cret = C.graphene_vec2_zero()

	var _vec2 *Vec2 // out

	_vec2 = (*Vec2)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec2
}
