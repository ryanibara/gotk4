// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_transform_get_type()), F: marshalTransform},
	})
}

// Transform: the `GskTransform` structure contains only private data.
type Transform struct {
	native C.GskTransform
}

// WrapTransform wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTransform(ptr unsafe.Pointer) *Transform {
	if ptr == nil {
		return nil
	}

	return (*Transform)(ptr)
}

func marshalTransform(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTransform(unsafe.Pointer(b)), nil
}

// NewTransform constructs a struct Transform.
func NewTransform() *Transform {
	var cret *C.GskTransform
	var ret1 *Transform

	cret = C.gsk_transform_new()

	ret1 = WrapTransform(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Transform) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Native returns the underlying C source pointer.
func (t *Transform) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Equal checks two transforms for equality.
func (f *Transform) Equal(second *Transform) bool {
	var arg0 *C.GskTransform
	var arg1 *C.GskTransform

	arg0 = (*C.GskTransform)(unsafe.Pointer(f.Native()))
	arg1 = (*C.GskTransform)(unsafe.Pointer(second.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gsk_transform_equal(arg0, second)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Category returns the category this transform belongs to.
func (s *Transform) Category() TransformCategory {
	var arg0 *C.GskTransform

	arg0 = (*C.GskTransform)(unsafe.Pointer(s.Native()))

	var cret C.GskTransformCategory
	var ret1 TransformCategory

	cret = C.gsk_transform_get_category(arg0)

	ret1 = TransformCategory(cret)

	return ret1
}

// Invert inverts the given transform.
//
// If @self is not invertible, nil is returned. Note that inverting nil also
// returns nil, which is the correct inverse of nil. If you need to
// differentiate between those cases, you should check @self is not nil before
// calling this function.
func (s *Transform) Invert() *Transform {
	var arg0 *C.GskTransform

	arg0 = (*C.GskTransform)(unsafe.Pointer(s.Native()))

	var cret *C.GskTransform
	var ret1 *Transform

	cret = C.gsk_transform_invert(arg0)

	ret1 = WrapTransform(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Transform) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Matrix multiplies @next with the given @matrix.
func (n *Transform) Matrix(matrix *graphene.Matrix) *Transform {
	var arg0 *C.GskTransform
	var arg1 *C.graphene_matrix_t

	arg0 = (*C.GskTransform)(unsafe.Pointer(n.Native()))
	arg1 = (*C.graphene_matrix_t)(unsafe.Pointer(matrix.Native()))

	var cret *C.GskTransform
	var ret1 *Transform

	cret = C.gsk_transform_matrix(arg0, matrix)

	ret1 = WrapTransform(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Transform) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Perspective applies a perspective projection transform. This transform scales
// points in X and Y based on their Z value, scaling points with positive Z
// values away from the origin, and those with negative Z values towards the
// origin. Points on the z=0 plane are unchanged.
func (n *Transform) Perspective(depth float32) *Transform {
	var arg0 *C.GskTransform
	var arg1 C.float

	arg0 = (*C.GskTransform)(unsafe.Pointer(n.Native()))
	arg1 = C.float(depth)

	var cret *C.GskTransform
	var ret1 *Transform

	cret = C.gsk_transform_perspective(arg0, depth)

	ret1 = WrapTransform(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Transform) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Print converts @self into a human-readable string representation suitable for
// printing that can later be parsed with gsk_transform_parse().
func (s *Transform) Print(string *glib.String) {
	var arg0 *C.GskTransform
	var arg1 *C.GString

	arg0 = (*C.GskTransform)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GString)(unsafe.Pointer(string.Native()))

	C.gsk_transform_print(arg0, string)
}

// Ref acquires a reference on the given Transform.
func (s *Transform) Ref() *Transform {
	var arg0 *C.GskTransform

	arg0 = (*C.GskTransform)(unsafe.Pointer(s.Native()))

	var cret *C.GskTransform
	var ret1 *Transform

	cret = C.gsk_transform_ref(arg0)

	ret1 = WrapTransform(unsafe.Pointer(cret))

	return ret1
}

// Rotate rotates @next @angle degrees in 2D - or in 3Dspeak, around the z axis.
func (n *Transform) Rotate(angle float32) *Transform {
	var arg0 *C.GskTransform
	var arg1 C.float

	arg0 = (*C.GskTransform)(unsafe.Pointer(n.Native()))
	arg1 = C.float(angle)

	var cret *C.GskTransform
	var ret1 *Transform

	cret = C.gsk_transform_rotate(arg0, angle)

	ret1 = WrapTransform(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Transform) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Rotate3D rotates @next @angle degrees around @axis.
//
// For a rotation in 2D space, use gsk_transform_rotate().
func (n *Transform) Rotate3D(angle float32, axis *graphene.Vec3) *Transform {
	var arg0 *C.GskTransform
	var arg1 C.float
	var arg2 *C.graphene_vec3_t

	arg0 = (*C.GskTransform)(unsafe.Pointer(n.Native()))
	arg1 = C.float(angle)
	arg2 = (*C.graphene_vec3_t)(unsafe.Pointer(axis.Native()))

	var cret *C.GskTransform
	var ret1 *Transform

	cret = C.gsk_transform_rotate_3d(arg0, angle, axis)

	ret1 = WrapTransform(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Transform) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Scale scales @next in 2-dimensional space by the given factors. Use
// gsk_transform_scale_3d() to scale in all 3 dimensions.
func (n *Transform) Scale(factorX float32, factorY float32) *Transform {
	var arg0 *C.GskTransform
	var arg1 C.float
	var arg2 C.float

	arg0 = (*C.GskTransform)(unsafe.Pointer(n.Native()))
	arg1 = C.float(factorX)
	arg2 = C.float(factorY)

	var cret *C.GskTransform
	var ret1 *Transform

	cret = C.gsk_transform_scale(arg0, factorX, factorY)

	ret1 = WrapTransform(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Transform) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Scale3D scales @next by the given factors.
func (n *Transform) Scale3D(factorX float32, factorY float32, factorZ float32) *Transform {
	var arg0 *C.GskTransform
	var arg1 C.float
	var arg2 C.float
	var arg3 C.float

	arg0 = (*C.GskTransform)(unsafe.Pointer(n.Native()))
	arg1 = C.float(factorX)
	arg2 = C.float(factorY)
	arg3 = C.float(factorZ)

	var cret *C.GskTransform
	var ret1 *Transform

	cret = C.gsk_transform_scale_3d(arg0, factorX, factorY, factorZ)

	ret1 = WrapTransform(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Transform) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// To2D converts a Transform to a 2D transformation matrix. @self must be a 2D
// transformation. If you are not sure, use gsk_transform_get_category() >=
// GSK_TRANSFORM_CATEGORY_2D to check.
//
// The returned values have the following layout:
//
//    | xx yx |   |  a  b  0 |
//    | xy yy | = |  c  d  0 |
//    | dx dy |   | tx ty  1 |
//
// This function can be used to convert between a Transform and a matrix type
// from other 2D drawing libraries, in particular Cairo.
func (s *Transform) To2D() (outXX float32, outYX float32, outXY float32, outYY float32, outDx float32, outDy float32) {
	var arg0 *C.GskTransform

	arg0 = (*C.GskTransform)(unsafe.Pointer(s.Native()))

	var arg1 C.float
	var ret1 float32
	var arg2 C.float
	var ret2 float32
	var arg3 C.float
	var ret3 float32
	var arg4 C.float
	var ret4 float32
	var arg5 C.float
	var ret5 float32
	var arg6 C.float
	var ret6 float32

	C.gsk_transform_to_2d(arg0, &arg1, &arg2, &arg3, &arg4, &arg5, &arg6)

	*ret1 = C.float(arg1)
	*ret2 = C.float(arg2)
	*ret3 = C.float(arg3)
	*ret4 = C.float(arg4)
	*ret5 = C.float(arg5)
	*ret6 = C.float(arg6)

	return ret1, ret2, ret3, ret4, ret5, ret6
}

// ToAffine converts a Transform to 2D affine transformation factors. @self must
// be a 2D transformation. If you are not sure, use gsk_transform_get_category()
// >= GSK_TRANSFORM_CATEGORY_2D_AFFINE to check.
func (s *Transform) ToAffine() (outScaleX float32, outScaleY float32, outDx float32, outDy float32) {
	var arg0 *C.GskTransform

	arg0 = (*C.GskTransform)(unsafe.Pointer(s.Native()))

	var arg1 C.float
	var ret1 float32
	var arg2 C.float
	var ret2 float32
	var arg3 C.float
	var ret3 float32
	var arg4 C.float
	var ret4 float32

	C.gsk_transform_to_affine(arg0, &arg1, &arg2, &arg3, &arg4)

	*ret1 = C.float(arg1)
	*ret2 = C.float(arg2)
	*ret3 = C.float(arg3)
	*ret4 = C.float(arg4)

	return ret1, ret2, ret3, ret4
}

// ToMatrix computes the actual value of @self and stores it in @out_matrix. The
// previous value of @out_matrix will be ignored.
func (s *Transform) ToMatrix() graphene.Matrix {
	var arg0 *C.GskTransform

	arg0 = (*C.GskTransform)(unsafe.Pointer(s.Native()))

	var arg1 C.graphene_matrix_t
	var ret1 *graphene.Matrix

	C.gsk_transform_to_matrix(arg0, &arg1)

	*ret1 = graphene.WrapMatrix(unsafe.Pointer(arg1))

	return ret1
}

// String converts a matrix into a string that is suitable for printing and can
// later be parsed with gsk_transform_parse().
//
// This is a wrapper around gsk_transform_print(), see that function for
// details.
func (s *Transform) String() string {
	var arg0 *C.GskTransform

	arg0 = (*C.GskTransform)(unsafe.Pointer(s.Native()))

	var cret *C.char
	var ret1 string

	cret = C.gsk_transform_to_string(arg0)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// ToTranslate converts a Transform to a translation operation. @self must be a
// 2D transformation. If you are not sure, use gsk_transform_get_category() >=
// GSK_TRANSFORM_CATEGORY_2D_TRANSLATE to check.
func (s *Transform) ToTranslate() (outDx float32, outDy float32) {
	var arg0 *C.GskTransform

	arg0 = (*C.GskTransform)(unsafe.Pointer(s.Native()))

	var arg1 C.float
	var ret1 float32
	var arg2 C.float
	var ret2 float32

	C.gsk_transform_to_translate(arg0, &arg1, &arg2)

	*ret1 = C.float(arg1)
	*ret2 = C.float(arg2)

	return ret1, ret2
}

// Transform applies all the operations from @other to @next.
func (n *Transform) Transform(other *Transform) *Transform {
	var arg0 *C.GskTransform
	var arg1 *C.GskTransform

	arg0 = (*C.GskTransform)(unsafe.Pointer(n.Native()))
	arg1 = (*C.GskTransform)(unsafe.Pointer(other.Native()))

	var cret *C.GskTransform
	var ret1 *Transform

	cret = C.gsk_transform_transform(arg0, other)

	ret1 = WrapTransform(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Transform) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// TransformBounds transforms a #graphene_rect_t using the given transform
// @self. The result is the bounding box containing the coplanar quad.
func (s *Transform) TransformBounds(rect *graphene.Rect) graphene.Rect {
	var arg0 *C.GskTransform
	var arg1 *C.graphene_rect_t

	arg0 = (*C.GskTransform)(unsafe.Pointer(s.Native()))
	arg1 = (*C.graphene_rect_t)(unsafe.Pointer(rect.Native()))

	var arg2 C.graphene_rect_t
	var ret2 *graphene.Rect

	C.gsk_transform_transform_bounds(arg0, rect, &arg2)

	*ret2 = graphene.WrapRect(unsafe.Pointer(arg2))

	return ret2
}

// TransformPoint transforms a #graphene_point_t using the given transform
// @self.
func (s *Transform) TransformPoint(point *graphene.Point) graphene.Point {
	var arg0 *C.GskTransform
	var arg1 *C.graphene_point_t

	arg0 = (*C.GskTransform)(unsafe.Pointer(s.Native()))
	arg1 = (*C.graphene_point_t)(unsafe.Pointer(point.Native()))

	var arg2 C.graphene_point_t
	var ret2 *graphene.Point

	C.gsk_transform_transform_point(arg0, point, &arg2)

	*ret2 = graphene.WrapPoint(unsafe.Pointer(arg2))

	return ret2
}

// Translate translates @next in 2dimensional space by @point.
func (n *Transform) Translate(point *graphene.Point) *Transform {
	var arg0 *C.GskTransform
	var arg1 *C.graphene_point_t

	arg0 = (*C.GskTransform)(unsafe.Pointer(n.Native()))
	arg1 = (*C.graphene_point_t)(unsafe.Pointer(point.Native()))

	var cret *C.GskTransform
	var ret1 *Transform

	cret = C.gsk_transform_translate(arg0, point)

	ret1 = WrapTransform(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Transform) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Translate3D translates @next by @point.
func (n *Transform) Translate3D(point *graphene.Point3D) *Transform {
	var arg0 *C.GskTransform
	var arg1 *C.graphene_point3d_t

	arg0 = (*C.GskTransform)(unsafe.Pointer(n.Native()))
	arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(point.Native()))

	var cret *C.GskTransform
	var ret1 *Transform

	cret = C.gsk_transform_translate_3d(arg0, point)

	ret1 = WrapTransform(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Transform) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Unref releases a reference on the given Transform.
//
// If the reference was the last, the resources associated to the @self are
// freed.
func (s *Transform) Unref() {
	var arg0 *C.GskTransform

	arg0 = (*C.GskTransform)(unsafe.Pointer(s.Native()))

	C.gsk_transform_unref(arg0)
}
