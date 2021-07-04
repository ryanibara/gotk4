// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib.h>
import "C"

// Array contains the public fields of a GArray.
type Array C.GArray

// WrapArray wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapArray(ptr unsafe.Pointer) *Array {
	return (*Array)(ptr)
}

func marshalArray(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Array)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *Array) Native() unsafe.Pointer {
	return unsafe.Pointer(a)
}

// ByteArray contains the public fields of a GByteArray.
type ByteArray C.GByteArray

// WrapByteArray wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapByteArray(ptr unsafe.Pointer) *ByteArray {
	return (*ByteArray)(ptr)
}

func marshalByteArray(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*ByteArray)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (b *ByteArray) Native() unsafe.Pointer {
	return unsafe.Pointer(b)
}

// Bytes: a simple refcounted data type representing an immutable sequence of
// zero or more bytes from an unspecified origin.
//
// The purpose of a #GBytes is to keep the memory region that it holds alive for
// as long as anyone holds a reference to the bytes. When the last reference
// count is dropped, the memory is released. Multiple unrelated callers can use
// byte data in the #GBytes without coordinating their activities, resting
// assured that the byte data will not change or move while they hold a
// reference.
//
// A #GBytes can come from many different origins that may have different
// procedures for freeing the memory region. Examples are memory from
// g_malloc(), from memory slices, from a File or memory from other allocators.
//
// #GBytes work well as keys in Table. Use g_bytes_equal() and g_bytes_hash() as
// parameters to g_hash_table_new() or g_hash_table_new_full(). #GBytes can also
// be used as keys in a #GTree by passing the g_bytes_compare() function to
// g_tree_new().
//
// The data pointed to by this bytes must not be modified. For a mutable array
// of bytes see Array. Use g_bytes_unref_to_array() to create a mutable array
// for a #GBytes sequence. To create an immutable #GBytes from a mutable Array,
// use the g_byte_array_free_to_bytes() function.
type Bytes C.GBytes

// WrapBytes wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBytes(ptr unsafe.Pointer) *Bytes {
	return (*Bytes)(ptr)
}

func marshalBytes(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Bytes)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (b *Bytes) Native() unsafe.Pointer {
	return unsafe.Pointer(b)
}

// PtrArray contains the public fields of a pointer array.
type PtrArray C.GPtrArray

// WrapPtrArray wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPtrArray(ptr unsafe.Pointer) *PtrArray {
	return (*PtrArray)(ptr)
}

func marshalPtrArray(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*PtrArray)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *PtrArray) Native() unsafe.Pointer {
	return unsafe.Pointer(p)
}
