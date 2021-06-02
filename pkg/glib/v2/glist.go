// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
)

// #cgo pkg-config: glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

// ClearList clears a pointer to a #GList, freeing it and, optionally, freeing
// its elements using @destroy.
//
// @list_ptr must be a valid pointer. If @list_ptr points to a null #GList, this
// does nothing.
func ClearList(listPtr **List) {
	var arg1 **C.GList
	var arg2 C.GDestroyNotify

	arg1 = (**C.GList)(listPtr.Native())

	C.g_clear_list(arg1, arg2)
}

// List: the #GList struct is used for each element in a doubly-linked list.
type List struct {
	native C.GList
}

// WrapList wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapList(ptr unsafe.Pointer) *List {
	if ptr == nil {
		return nil
	}

	return (*List)(ptr)
}

func marshalList(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapList(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (l *List) Native() unsafe.Pointer {
	return unsafe.Pointer(&l.native)
}

// Data gets the field inside the struct.
func (l *List) Data() interface{} {
	var ret interface{}
	ret = box.Get(uintptr(l.native.data))
	return ret
}

// Next gets the field inside the struct.
func (l *List) Next() *List {
	var ret *List
	{
		ret = WrapList(unsafe.Pointer(l.native.next))
	}
	return ret
}

// Prev gets the field inside the struct.
func (l *List) Prev() *List {
	var ret *List
	{
		ret = WrapList(unsafe.Pointer(l.native.prev))
	}
	return ret
}
