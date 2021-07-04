// Code generated by girgen. DO NOT EDIT.

package gobject

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gobject-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_value_array_get_type()), F: marshalValueArray},
	})
}

// ValueArray: a Array contains an array of #GValue elements.
type ValueArray C.GValueArray

// WrapValueArray wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapValueArray(ptr unsafe.Pointer) *ValueArray {
	return (*ValueArray)(ptr)
}

func marshalValueArray(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*ValueArray)(unsafe.Pointer(b)), nil
}

// NewValueArray constructs a struct ValueArray.
func NewValueArray(nPrealloced uint) *ValueArray {
	var _arg1 C.guint        // out
	var _cret *C.GValueArray // in

	_arg1 = C.guint(nPrealloced)

	_cret = C.g_value_array_new(_arg1)

	var _valueArray *ValueArray // out

	_valueArray = (*ValueArray)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_valueArray, func(v **ValueArray) {
		C.free(unsafe.Pointer(v))
	})

	return _valueArray
}

// Native returns the underlying C source pointer.
func (v *ValueArray) Native() unsafe.Pointer {
	return unsafe.Pointer(v)
}

// Append: sort @value_array using @compare_func to compare the elements
// according to the semantics of DataFunc.
//
// The current implementation uses the same sorting algorithm as standard C
// qsort() function.
//
// Deprecated: since version 2.32.
func (v *ValueArray) Append(value externglib.Value) *ValueArray {
	var _arg0 *C.GValueArray // out
	var _arg1 *C.GValue      // out
	var _cret *C.GValueArray // in

	_arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	_cret = C.g_value_array_append(_arg0, _arg1)

	var _valueArray *ValueArray // out

	_valueArray = (*ValueArray)(unsafe.Pointer(_cret))

	return _valueArray
}

// Copy: sort @value_array using @compare_func to compare the elements according
// to the semantics of DataFunc.
//
// The current implementation uses the same sorting algorithm as standard C
// qsort() function.
//
// Deprecated: since version 2.32.
func (v *ValueArray) Copy() *ValueArray {
	var _arg0 *C.GValueArray // out
	var _cret *C.GValueArray // in

	_arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))

	_cret = C.g_value_array_copy(_arg0)

	var _valueArray *ValueArray // out

	_valueArray = (*ValueArray)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_valueArray, func(v **ValueArray) {
		C.free(unsafe.Pointer(v))
	})

	return _valueArray
}

// Nth: sort @value_array using @compare_func to compare the elements according
// to the semantics of DataFunc.
//
// The current implementation uses the same sorting algorithm as standard C
// qsort() function.
//
// Deprecated: since version 2.32.
func (v *ValueArray) Nth(index_ uint) externglib.Value {
	var _arg0 *C.GValueArray // out
	var _arg1 C.guint        // out
	var _cret *C.GValue      // in

	_arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))
	_arg1 = C.guint(index_)

	_cret = C.g_value_array_get_nth(_arg0, _arg1)

	var _value externglib.Value // out

	_value = externglib.ValueFromNative(unsafe.Pointer(_cret))

	return _value
}

// Insert: sort @value_array using @compare_func to compare the elements
// according to the semantics of DataFunc.
//
// The current implementation uses the same sorting algorithm as standard C
// qsort() function.
//
// Deprecated: since version 2.32.
func (v *ValueArray) Insert(index_ uint, value externglib.Value) *ValueArray {
	var _arg0 *C.GValueArray // out
	var _arg1 C.guint        // out
	var _arg2 *C.GValue      // out
	var _cret *C.GValueArray // in

	_arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))
	_arg1 = C.guint(index_)
	_arg2 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	_cret = C.g_value_array_insert(_arg0, _arg1, _arg2)

	var _valueArray *ValueArray // out

	_valueArray = (*ValueArray)(unsafe.Pointer(_cret))

	return _valueArray
}

// Prepend: sort @value_array using @compare_func to compare the elements
// according to the semantics of DataFunc.
//
// The current implementation uses the same sorting algorithm as standard C
// qsort() function.
//
// Deprecated: since version 2.32.
func (v *ValueArray) Prepend(value externglib.Value) *ValueArray {
	var _arg0 *C.GValueArray // out
	var _arg1 *C.GValue      // out
	var _cret *C.GValueArray // in

	_arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	_cret = C.g_value_array_prepend(_arg0, _arg1)

	var _valueArray *ValueArray // out

	_valueArray = (*ValueArray)(unsafe.Pointer(_cret))

	return _valueArray
}

// Remove: sort @value_array using @compare_func to compare the elements
// according to the semantics of DataFunc.
//
// The current implementation uses the same sorting algorithm as standard C
// qsort() function.
//
// Deprecated: since version 2.32.
func (v *ValueArray) Remove(index_ uint) *ValueArray {
	var _arg0 *C.GValueArray // out
	var _arg1 C.guint        // out
	var _cret *C.GValueArray // in

	_arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))
	_arg1 = C.guint(index_)

	_cret = C.g_value_array_remove(_arg0, _arg1)

	var _valueArray *ValueArray // out

	_valueArray = (*ValueArray)(unsafe.Pointer(_cret))

	return _valueArray
}

// SortWithData: sort @value_array using @compare_func to compare the elements
// according to the semantics of DataFunc.
//
// The current implementation uses the same sorting algorithm as standard C
// qsort() function.
//
// Deprecated: since version 2.32.
func (v *ValueArray) SortWithData(compareFunc glib.CompareDataFunc) *ValueArray {
	var _arg0 *C.GValueArray     // out
	var _arg1 C.GCompareDataFunc // out
	var _arg2 C.gpointer
	var _cret *C.GValueArray // in

	_arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))
	_arg1 = (*[0]byte)(C.gotk4_CompareDataFunc)
	_arg2 = C.gpointer(box.Assign(compareFunc))

	_cret = C.g_value_array_sort_with_data(_arg0, _arg1, _arg2)

	var _valueArray *ValueArray // out

	_valueArray = (*ValueArray)(unsafe.Pointer(_cret))

	return _valueArray
}
