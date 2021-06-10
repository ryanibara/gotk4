// Code generated by girgen. DO NOT EDIT.

package pango

import (
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
	var v float64
	v = (float64)(m.native.xx)
	return v
}

// XY gets the field inside the struct.
func (m *Matrix) XY() float64 {
	var v float64
	v = (float64)(m.native.xy)
	return v
}

// YX gets the field inside the struct.
func (m *Matrix) YX() float64 {
	var v float64
	v = (float64)(m.native.yx)
	return v
}

// YY gets the field inside the struct.
func (m *Matrix) YY() float64 {
	var v float64
	v = (float64)(m.native.yy)
	return v
}

// X0 gets the field inside the struct.
func (m *Matrix) X0() float64 {
	var v float64
	v = (float64)(m.native.x0)
	return v
}

// Y0 gets the field inside the struct.
func (m *Matrix) Y0() float64 {
	var v float64
	v = (float64)(m.native.y0)
	return v
}

// Concat changes the transformation represented by @matrix to be the
// transformation given by first applying transformation given by @new_matrix
// then applying the original transformation.
func (m *Matrix) Concat(newMatrix *Matrix) {
	var _arg0 *C.PangoMatrix
	var _arg1 *C.PangoMatrix

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.PangoMatrix)(unsafe.Pointer(newMatrix.Native()))

	C.pango_matrix_concat(_arg0, _arg1)
}

// Free: free a `PangoMatrix`.
func (m *Matrix) Free() {
	var _arg0 *C.PangoMatrix

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))

	C.pango_matrix_free(_arg0)
}

// FontScaleFactor returns the scale factor of a matrix on the height of the
// font.
//
// That is, the scale factor in the direction perpendicular to the vector that
// the X coordinate is mapped to. If the scale in the X coordinate is needed as
// well, use [method@Pango.Matrix.get_font_scale_factors].
func (m *Matrix) FontScaleFactor() float64 {
	var _arg0 *C.PangoMatrix

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))

	var _cret C.double

	_cret = C.pango_matrix_get_font_scale_factor(_arg0)

	var _gdouble float64

	_gdouble = (float64)(_cret)

	return _gdouble
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
	var _arg0 *C.PangoMatrix

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))

	var _arg1 C.double
	var _arg2 C.double

	C.pango_matrix_get_font_scale_factors(_arg0, &_arg1, &_arg2)

	var _xscale float64
	var _yscale float64

	_xscale = (float64)(_arg1)
	_yscale = (float64)(_arg2)

	return _xscale, _yscale
}

// Rotate changes the transformation represented by @matrix to be the
// transformation given by first rotating by @degrees degrees counter-clockwise
// then applying the original transformation.
func (m *Matrix) Rotate(degrees float64) {
	var _arg0 *C.PangoMatrix
	var _arg1 C.double

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))
	_arg1 = C.double(degrees)

	C.pango_matrix_rotate(_arg0, _arg1)
}

// Scale changes the transformation represented by @matrix to be the
// transformation given by first scaling by @sx in the X direction and @sy in
// the Y direction then applying the original transformation.
func (m *Matrix) Scale(scaleX float64, scaleY float64) {
	var _arg0 *C.PangoMatrix
	var _arg1 C.double
	var _arg2 C.double

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))
	_arg1 = C.double(scaleX)
	_arg2 = C.double(scaleY)

	C.pango_matrix_scale(_arg0, _arg1, _arg2)
}

// Translate changes the transformation represented by @matrix to be the
// transformation given by first translating by (@tx, @ty) then applying the
// original transformation.
func (m *Matrix) Translate(tx float64, ty float64) {
	var _arg0 *C.PangoMatrix
	var _arg1 C.double
	var _arg2 C.double

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(m.Native()))
	_arg1 = C.double(tx)
	_arg2 = C.double(ty)

	C.pango_matrix_translate(_arg0, _arg1, _arg2)
}
