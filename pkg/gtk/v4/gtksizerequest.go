// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// DistributeNaturalAllocation distributes @extra_space to child @sizes by
// bringing smaller children up to natural size first.
//
// The remaining space will be added to the @minimum_size member of the
// `GtkRequestedSize` struct. If all sizes reach their natural size then the
// remaining space is returned.
func DistributeNaturalAllocation(extraSpace int, sizes []RequestedSize) int {
	var _arg1 C.int
	var _arg3 *C.GtkRequestedSize
	var _arg2 C.guint

	_arg1 = C.int(extraSpace)
	_arg2 = C.guint(len(sizes))
	_arg3 = (*C.GtkRequestedSize)(unsafe.Pointer(&sizes[0]))

	var _cret C.int

	_cret = C.gtk_distribute_natural_allocation(_arg1, _arg2, _arg3)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// RequestedSize represents a request of a screen object in a given orientation.
// These are primarily used in container implementations when allocating a
// natural size for children calling. See gtk_distribute_natural_allocation().
type RequestedSize struct {
	native C.GtkRequestedSize
}

// WrapRequestedSize wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRequestedSize(ptr unsafe.Pointer) *RequestedSize {
	if ptr == nil {
		return nil
	}

	return (*RequestedSize)(ptr)
}

func marshalRequestedSize(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRequestedSize(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RequestedSize) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Data gets the field inside the struct.
func (r *RequestedSize) Data() interface{} {
	var v interface{}
	v = (interface{})(r.native.data)
	return v
}

// MinimumSize gets the field inside the struct.
func (r *RequestedSize) MinimumSize() int {
	var v int
	v = (int)(r.native.minimum_size)
	return v
}

// NaturalSize gets the field inside the struct.
func (r *RequestedSize) NaturalSize() int {
	var v int
	v = (int)(r.native.natural_size)
	return v
}
