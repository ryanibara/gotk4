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

// MemVTable: a set of functions used to perform memory allocation. The same
// VTable must be used for all allocations in the same program; a call to
// g_mem_set_vtable(), if it exists, should be prior to any use of GLib.
//
// This functions related to this has been deprecated in 2.46, and no longer
// work.
type MemVTable C.GMemVTable

// WrapMemVTable wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapMemVTable(ptr unsafe.Pointer) *MemVTable {
	return (*MemVTable)(ptr)
}

// Native returns the underlying C source pointer.
func (m *MemVTable) Native() unsafe.Pointer {
	return unsafe.Pointer(m)
}
