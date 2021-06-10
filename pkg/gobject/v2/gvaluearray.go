// Code generated by girgen. DO NOT EDIT.

package gobject

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gobject-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_value_array_get_type()), F: marshalValueArray},
	})
}

// ValueArray: a Array contains an array of #GValue elements.
type ValueArray struct {
	native C.GValueArray
}

// WrapValueArray wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapValueArray(ptr unsafe.Pointer) *ValueArray {
	if ptr == nil {
		return nil
	}

	return (*ValueArray)(ptr)
}

func marshalValueArray(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapValueArray(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (v *ValueArray) Native() unsafe.Pointer {
	return unsafe.Pointer(&v.native)
}

// NValues gets the field inside the struct.
func (v *ValueArray) NValues() uint {
	var v uint
	v = (uint)(v.native.n_values)
	return v
}

// Values gets the field inside the struct.
func (v *ValueArray) Values() **externglib.Value {
	var v **externglib.Value
	v = externglib.ValueFromNative(unsafe.Pointer(v.native.values))
	return v
}

// Free: free a Array including its contents.
func (v *ValueArray) Free() {
	var _arg0 *C.GValueArray

	_arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))

	C.g_value_array_free(_arg0)
}

// Nth: return a pointer to the value at @index_ containd in @value_array.
func (v *ValueArray) Nth(index_ uint) **externglib.Value {
	var _arg0 *C.GValueArray
	var _arg1 C.guint

	_arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))
	_arg1 = C.guint(index_)

	var _cret *C.GValue

	_cret = C.g_value_array_get_nth(_arg0, _arg1)

	var _value **externglib.Value

	_value = externglib.ValueFromNative(unsafe.Pointer(_cret))

	return _value
}
