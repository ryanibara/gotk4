// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_matrix_get_type()), F: marshalMatrix},
	})
}

// Matrix: a `PangoMatrix` specifies a transformation between user-space and
// device coordinates.
//
//
// The transformation is given by
//
// “` x_device = x_user * matrix->xx + y_user * matrix->xy + matrix->x0;
// y_device = x_user * matrix->yx + y_user * matrix->yy + matrix->y0; “`
type Matrix struct {
	native C.PangoMatrix
}

// WrapMatrix wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapMatrix(ptr unsafe.Pointer) *Matrix {
	if ptr == nil {
		return nil
	}

	return (*Matrix)(ptr)
}

func marshalMatrix(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapMatrix(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (m *Matrix) Native() unsafe.Pointer {
	return unsafe.Pointer(&m.native)
}

// XX gets the field inside the struct.
func (m *Matrix) XX() float64 {
	v = C.double(m.native.xx)
}

// XY gets the field inside the struct.
func (m *Matrix) XY() float64 {
	v = C.double(m.native.xy)
}

// YX gets the field inside the struct.
func (m *Matrix) YX() float64 {
	v = C.double(m.native.yx)
}

// YY gets the field inside the struct.
func (m *Matrix) YY() float64 {
	v = C.double(m.native.yy)
}

// X0 gets the field inside the struct.
func (m *Matrix) X0() float64 {
	v = C.double(m.native.x0)
}

// Y0 gets the field inside the struct.
func (m *Matrix) Y0() float64 {
	v = C.double(m.native.y0)
}

// Concat changes the transformation represented by @matrix to be the
// transformation given by first applying transformation given by @new_matrix
// then applying the original transformation.
func (m *Matrix) Concat(newMatrix *Matrix) {
	var arg0 *C.PangoMatrix
	var arg1 *C.PangoMatrix

	arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))
	arg1 = (*C.PangoMatrix)(unsafe.Pointer(newMatrix.Native()))

	C.pango_matrix_concat(arg0, newMatrix)
}

// Copy copies a `PangoMatrix`.
func (m *Matrix) Copy() *Matrix {
	var arg0 *C.PangoMatrix

	arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))

	var cret *C.PangoMatrix
	var ret1 *Matrix

	cret = C.pango_matrix_copy(arg0)

	ret1 = WrapMatrix(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Matrix) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Free: free a `PangoMatrix`.
func (m *Matrix) Free() {
	var arg0 *C.PangoMatrix

	arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))

	C.pango_matrix_free(arg0)
}

// FontScaleFactor returns the scale factor of a matrix on the height of the
// font.
//
// That is, the scale factor in the direction perpendicular to the vector that
// the X coordinate is mapped to. If the scale in the X coordinate is needed as
// well, use [method@Pango.Matrix.get_font_scale_factors].
func (m *Matrix) FontScaleFactor() float64 {
	var arg0 *C.PangoMatrix

	arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))

	var cret C.double
	var ret1 float64

	cret = C.pango_matrix_get_font_scale_factor(arg0)

	ret1 = C.double(cret)

	return ret1
}

// FontScaleFactors calculates the scale factor of a matrix on the width and
// height of the font.
//
// That is, @xscale is the scale factor in the direction of the X coordinate,
// and @yscale is the scale factor in the direction perpendicular to the vector
// that the X coordinate is mapped to.
//
// Note that output numbers will always be non-negative.
func (m *Matrix) FontScaleFactors() (xscale float64, yscale float64) {
	var arg0 *C.PangoMatrix

	arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))

	var arg1 C.double
	var ret1 float64
	var arg2 C.double
	var ret2 float64

	C.pango_matrix_get_font_scale_factors(arg0, &arg1, &arg2)

	ret1 = C.double(arg1)
	ret2 = C.double(arg2)

	return ret1, ret2
}

// Rotate changes the transformation represented by @matrix to be the
// transformation given by first rotating by @degrees degrees counter-clockwise
// then applying the original transformation.
func (m *Matrix) Rotate(degrees float64) {
	var arg0 *C.PangoMatrix
	var arg1 C.double

	arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))
	arg1 = C.double(degrees)

	C.pango_matrix_rotate(arg0, degrees)
}

// Scale changes the transformation represented by @matrix to be the
// transformation given by first scaling by @sx in the X direction and @sy in
// the Y direction then applying the original transformation.
func (m *Matrix) Scale(scaleX float64, scaleY float64) {
	var arg0 *C.PangoMatrix
	var arg1 C.double
	var arg2 C.double

	arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))
	arg1 = C.double(scaleX)
	arg2 = C.double(scaleY)

	C.pango_matrix_scale(arg0, scaleX, scaleY)
}

// TransformDistance transforms the distance vector (@dx,@dy) by @matrix.
//
// This is similar to [method@Pango.Matrix.transform_point], except that the
// translation components of the transformation are ignored. The calculation of
// the returned vector is as follows:
//
// “` dx2 = dx1 * xx + dy1 * xy; dy2 = dx1 * yx + dy1 * yy; “`
//
// Affine transformations are position invariant, so the same vector always
// transforms to the same vector. If (@x1,@y1) transforms to (@x2,@y2) then
// (@x1+@dx1,@y1+@dy1) will transform to (@x1+@dx2,@y1+@dy2) for all values of
// @x1 and @x2.
func (m *Matrix) TransformDistance(dx float64, dy float64) {
	var arg0 *C.PangoMatrix
	var arg1 *C.double
	var arg2 *C.double

	arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))
	arg1 = *C.double(dx)
	arg2 = *C.double(dy)

	C.pango_matrix_transform_distance(arg0, dx, dy)
}

// TransformPixelRectangle: first transforms the @rect using @matrix, then
// calculates the bounding box of the transformed rectangle.
//
// This function is useful for example when you want to draw a rotated
// @PangoLayout to an image buffer, and want to know how large the image should
// be and how much you should shift the layout when rendering.
//
// For better accuracy, you should use [method@Pango.Matrix.transform_rectangle]
// on original rectangle in Pango units and convert to pixels afterward using
// [func@extents_to_pixels]'s first argument.
func (m *Matrix) TransformPixelRectangle(rect *Rectangle) {
	var arg0 *C.PangoMatrix
	var arg1 *C.PangoRectangle

	arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))
	arg1 = (*C.PangoRectangle)(unsafe.Pointer(rect.Native()))

	C.pango_matrix_transform_pixel_rectangle(arg0, rect)
}

// TransformPoint transforms the point (@x, @y) by @matrix.
func (m *Matrix) TransformPoint(x float64, y float64) {
	var arg0 *C.PangoMatrix
	var arg1 *C.double
	var arg2 *C.double

	arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))
	arg1 = *C.double(x)
	arg2 = *C.double(y)

	C.pango_matrix_transform_point(arg0, x, y)
}

// TransformRectangle: first transforms @rect using @matrix, then calculates the
// bounding box of the transformed rectangle.
//
// This function is useful for example when you want to draw a rotated
// @PangoLayout to an image buffer, and want to know how large the image should
// be and how much you should shift the layout when rendering.
//
// If you have a rectangle in device units (pixels), use
// [method@Pango.Matrix.transform_pixel_rectangle].
//
// If you have the rectangle in Pango units and want to convert to transformed
// pixel bounding box, it is more accurate to transform it first (using this
// function) and pass the result to pango_extents_to_pixels(), first argument,
// for an inclusive rounded rectangle. However, there are valid reasons that you
// may want to convert to pixels first and then transform, for example when the
// transformed coordinates may overflow in Pango units (large matrix translation
// for example).
func (m *Matrix) TransformRectangle(rect *Rectangle) {
	var arg0 *C.PangoMatrix
	var arg1 *C.PangoRectangle

	arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))
	arg1 = (*C.PangoRectangle)(unsafe.Pointer(rect.Native()))

	C.pango_matrix_transform_rectangle(arg0, rect)
}

// Translate changes the transformation represented by @matrix to be the
// transformation given by first translating by (@tx, @ty) then applying the
// original transformation.
func (m *Matrix) Translate(tx float64, ty float64) {
	var arg0 *C.PangoMatrix
	var arg1 C.double
	var arg2 C.double

	arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))
	arg1 = C.double(tx)
	arg2 = C.double(ty)

	C.pango_matrix_translate(arg0, tx, ty)
}
