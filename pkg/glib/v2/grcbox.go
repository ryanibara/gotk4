// Code generated by girgen. DO NOT EDIT.

package glib

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// AtomicRCBoxAcquire: atomically acquires a reference on the data pointed by
// @mem_block.
func AtomicRCBoxAcquire(memBlock interface{}) interface{} {
	var _arg1 C.gpointer

	_arg1 = C.gpointer(memBlock)

	var _cret C.gpointer

	cret = C.g_atomic_rc_box_acquire(_arg1)

	var _gpointer interface{}

	_gpointer = (interface{})(_cret)

	return _gpointer
}

// AtomicRCBoxAlloc allocates @block_size bytes of memory, and adds atomic
// reference counting semantics to it.
//
// The data will be freed when its reference count drops to zero.
//
// The allocated data is guaranteed to be suitably aligned for any built-in
// type.
func AtomicRCBoxAlloc(blockSize uint) interface{} {
	var _arg1 C.gsize

	_arg1 = C.gsize(blockSize)

	var _cret C.gpointer

	cret = C.g_atomic_rc_box_alloc(_arg1)

	var _gpointer interface{}

	_gpointer = (interface{})(_cret)

	return _gpointer
}

// AtomicRCBoxAlloc0 allocates @block_size bytes of memory, and adds atomic
// reference counting semantics to it.
//
// The contents of the returned data is set to zero.
//
// The data will be freed when its reference count drops to zero.
//
// The allocated data is guaranteed to be suitably aligned for any built-in
// type.
func AtomicRCBoxAlloc0(blockSize uint) interface{} {
	var _arg1 C.gsize

	_arg1 = C.gsize(blockSize)

	var _cret C.gpointer

	cret = C.g_atomic_rc_box_alloc0(_arg1)

	var _gpointer interface{}

	_gpointer = (interface{})(_cret)

	return _gpointer
}

// AtomicRCBoxDup allocates a new block of data with atomic reference counting
// semantics, and copies @block_size bytes of @mem_block into it.
func AtomicRCBoxDup(blockSize uint, memBlock interface{}) interface{} {
	var _arg1 C.gsize
	var _arg2 C.gpointer

	_arg1 = C.gsize(blockSize)
	_arg2 = C.gpointer(memBlock)

	var _cret C.gpointer

	cret = C.g_atomic_rc_box_dup(_arg1, _arg2)

	var _gpointer interface{}

	_gpointer = (interface{})(_cret)

	return _gpointer
}

// AtomicRCBoxGetSize retrieves the size of the reference counted data pointed
// by @mem_block.
func AtomicRCBoxGetSize(memBlock interface{}) uint {
	var _arg1 C.gpointer

	_arg1 = C.gpointer(memBlock)

	var _cret C.gsize

	cret = C.g_atomic_rc_box_get_size(_arg1)

	var _gsize uint

	_gsize = (uint)(_cret)

	return _gsize
}

// AtomicRCBoxRelease: atomically releases a reference on the data pointed by
// @mem_block.
//
// If the reference was the last one, it will free the resources allocated for
// @mem_block.
func AtomicRCBoxRelease(memBlock interface{}) {
	var _arg1 C.gpointer

	_arg1 = C.gpointer(memBlock)

	C.g_atomic_rc_box_release(_arg1)
}

// RCBoxAcquire acquires a reference on the data pointed by @mem_block.
func RCBoxAcquire(memBlock interface{}) interface{} {
	var _arg1 C.gpointer

	_arg1 = C.gpointer(memBlock)

	var _cret C.gpointer

	cret = C.g_rc_box_acquire(_arg1)

	var _gpointer interface{}

	_gpointer = (interface{})(_cret)

	return _gpointer
}

// RCBoxAlloc allocates @block_size bytes of memory, and adds reference counting
// semantics to it.
//
// The data will be freed when its reference count drops to zero.
//
// The allocated data is guaranteed to be suitably aligned for any built-in
// type.
func RCBoxAlloc(blockSize uint) interface{} {
	var _arg1 C.gsize

	_arg1 = C.gsize(blockSize)

	var _cret C.gpointer

	cret = C.g_rc_box_alloc(_arg1)

	var _gpointer interface{}

	_gpointer = (interface{})(_cret)

	return _gpointer
}

// RCBoxAlloc0 allocates @block_size bytes of memory, and adds reference
// counting semantics to it.
//
// The contents of the returned data is set to zero.
//
// The data will be freed when its reference count drops to zero.
//
// The allocated data is guaranteed to be suitably aligned for any built-in
// type.
func RCBoxAlloc0(blockSize uint) interface{} {
	var _arg1 C.gsize

	_arg1 = C.gsize(blockSize)

	var _cret C.gpointer

	cret = C.g_rc_box_alloc0(_arg1)

	var _gpointer interface{}

	_gpointer = (interface{})(_cret)

	return _gpointer
}

// RCBoxDup allocates a new block of data with reference counting semantics, and
// copies @block_size bytes of @mem_block into it.
func RCBoxDup(blockSize uint, memBlock interface{}) interface{} {
	var _arg1 C.gsize
	var _arg2 C.gpointer

	_arg1 = C.gsize(blockSize)
	_arg2 = C.gpointer(memBlock)

	var _cret C.gpointer

	cret = C.g_rc_box_dup(_arg1, _arg2)

	var _gpointer interface{}

	_gpointer = (interface{})(_cret)

	return _gpointer
}

// RCBoxGetSize retrieves the size of the reference counted data pointed by
// @mem_block.
func RCBoxGetSize(memBlock interface{}) uint {
	var _arg1 C.gpointer

	_arg1 = C.gpointer(memBlock)

	var _cret C.gsize

	cret = C.g_rc_box_get_size(_arg1)

	var _gsize uint

	_gsize = (uint)(_cret)

	return _gsize
}

// RCBoxRelease releases a reference on the data pointed by @mem_block.
//
// If the reference was the last one, it will free the resources allocated for
// @mem_block.
func RCBoxRelease(memBlock interface{}) {
	var _arg1 C.gpointer

	_arg1 = C.gpointer(memBlock)

	C.g_rc_box_release(_arg1)
}