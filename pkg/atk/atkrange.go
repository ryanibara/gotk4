// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_range_get_type()), F: marshalRange},
	})
}

// Range are used on Value, in order to represent the full range of a given
// component (for example an slider or a range control), or to define each
// individual subrange this full range is splitted if available. See Value
// documentation for further details.
type Range struct {
	native C.AtkRange
}

func marshalRange(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Range)(unsafe.Pointer(b)), nil
}

// NewRange constructs a struct Range.
func NewRange(lowerLimit float64, upperLimit float64, description string) *Range {
	var _arg1 C.gdouble   // out
	var _arg2 C.gdouble   // out
	var _arg3 *C.gchar    // out
	var _cret *C.AtkRange // in

	_arg1 = C.gdouble(lowerLimit)
	_arg2 = C.gdouble(upperLimit)
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(description)))

	_cret = C.atk_range_new(_arg1, _arg2, _arg3)

	var __range *Range // out

	__range = (*Range)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(__range, func(v *Range) {
		C.atk_range_free((*C.AtkRange)(unsafe.Pointer(v)))
	})

	return __range
}

// Native returns the underlying C source pointer.
func (r *Range) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Copy returns a new Range that is a exact copy of @src
func (src *Range) Copy() *Range {
	var _arg0 *C.AtkRange // out
	var _cret *C.AtkRange // in

	_arg0 = (*C.AtkRange)(unsafe.Pointer(src))

	_cret = C.atk_range_copy(_arg0)

	var __range *Range // out

	__range = (*Range)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(__range, func(v *Range) {
		C.atk_range_free((*C.AtkRange)(unsafe.Pointer(v)))
	})

	return __range
}

// Free: free @range
func (_range *Range) free() {
	var _arg0 *C.AtkRange // out

	_arg0 = (*C.AtkRange)(unsafe.Pointer(_range))

	C.atk_range_free(_arg0)
}

// Description returns the human readable description of @range
func (_range *Range) Description() string {
	var _arg0 *C.AtkRange // out
	var _cret *C.gchar    // in

	_arg0 = (*C.AtkRange)(unsafe.Pointer(_range))

	_cret = C.atk_range_get_description(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// LowerLimit returns the lower limit of @range
func (_range *Range) LowerLimit() float64 {
	var _arg0 *C.AtkRange // out
	var _cret C.gdouble   // in

	_arg0 = (*C.AtkRange)(unsafe.Pointer(_range))

	_cret = C.atk_range_get_lower_limit(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// UpperLimit returns the upper limit of @range
func (_range *Range) UpperLimit() float64 {
	var _arg0 *C.AtkRange // out
	var _cret C.gdouble   // in

	_arg0 = (*C.AtkRange)(unsafe.Pointer(_range))

	_cret = C.atk_range_get_upper_limit(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}
