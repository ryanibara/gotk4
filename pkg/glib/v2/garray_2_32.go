// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// Bytes: simple refcounted data type representing an immutable sequence of zero
// or more bytes from an unspecified origin.
//
// The purpose of a #GBytes is to keep the memory region that it holds alive for
// as long as anyone holds a reference to the bytes. When the last reference
// count is dropped, the memory is released. Multiple unrelated callers
// can use byte data in the #GBytes without coordinating their activities,
// resting assured that the byte data will not change or move while they hold a
// reference.
//
// A #GBytes can come from many different origins that may have different
// procedures for freeing the memory region. Examples are memory from
// g_malloc(), from memory slices, from a File or memory from other allocators.
//
// #GBytes work well as keys in Table. Use g_bytes_equal() and g_bytes_hash()
// as parameters to g_hash_table_new() or g_hash_table_new_full(). #GBytes can
// also be used as keys in a #GTree by passing the g_bytes_compare() function to
// g_tree_new().
//
// The data pointed to by this bytes must not be modified. For a mutable array
// of bytes see Array. Use g_bytes_unref_to_array() to create a mutable array
// for a #GBytes sequence. To create an immutable #GBytes from a mutable Array,
// use the g_byte_array_free_to_bytes() function.
//
// An instance of this type is always passed by reference.
type Bytes struct {
	*bytes
}

// bytes is the struct that's finalized.
type bytes struct {
	native *C.GBytes
}

// NewBytes constructs a struct Bytes.
func NewBytes(data []byte) *Bytes {
	var _arg1 C.gconstpointer // out
	var _arg2 C.gsize
	var _cret *C.GBytes // in

	_arg2 = (C.gsize)(len(data))
	if len(data) > 0 {
		_arg1 = (C.gconstpointer)(unsafe.Pointer(&data[0]))
	}

	_cret = C.g_bytes_new(_arg1, _arg2)
	runtime.KeepAlive(data)

	var _bytes *Bytes // out

	_bytes = (*Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	return _bytes
}

// Compare compares the two #GBytes values.
//
// This function can be used to sort GBytes instances in lexicographical order.
//
// If bytes1 and bytes2 have different length but the shorter one is a prefix of
// the longer one then the shorter one is considered to be less than the longer
// one. Otherwise the first byte where both differ is used for comparison.
// If bytes1 has a smaller value at that position it is considered less,
// otherwise greater than bytes2.
//
// The function takes the following parameters:
//
//   - bytes2: pointer to a #GBytes to compare with bytes1.
//
// The function returns the following values:
//
//   - gint: negative value if bytes1 is less than bytes2, a positive value if
//     bytes1 is greater than bytes2, and zero if bytes1 is equal to bytes2.
//
func (bytes1 *Bytes) Compare(bytes2 *Bytes) int {
	var _arg0 C.gconstpointer // out
	var _arg1 C.gconstpointer // out
	var _cret C.gint          // in

	_arg0 = *(*C.gconstpointer)(gextras.StructNative(unsafe.Pointer(bytes1)))
	_arg1 = *(*C.gconstpointer)(gextras.StructNative(unsafe.Pointer(bytes2)))

	_cret = C.g_bytes_compare(_arg0, _arg1)
	runtime.KeepAlive(bytes1)
	runtime.KeepAlive(bytes2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Equal compares the two #GBytes values being pointed to and returns TRUE if
// they are equal.
//
// This function can be passed to g_hash_table_new() as the key_equal_func
// parameter, when using non-NULL #GBytes pointers as keys in a Table.
//
// The function takes the following parameters:
//
//   - bytes2: pointer to a #GBytes to compare with bytes1.
//
// The function returns the following values:
//
//   - ok: TRUE if the two keys match.
//
func (bytes1 *Bytes) Equal(bytes2 *Bytes) bool {
	var _arg0 C.gconstpointer // out
	var _arg1 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg0 = *(*C.gconstpointer)(gextras.StructNative(unsafe.Pointer(bytes1)))
	_arg1 = *(*C.gconstpointer)(gextras.StructNative(unsafe.Pointer(bytes2)))

	_cret = C.g_bytes_equal(_arg0, _arg1)
	runtime.KeepAlive(bytes1)
	runtime.KeepAlive(bytes2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Data: get the byte data in the #GBytes. This data should not be modified.
//
// This function will always return the same pointer for a given #GBytes.
//
// NULL may be returned if size is 0. This is not guaranteed, as the #GBytes may
// represent an empty string with data non-NULL and size as 0. NULL will not be
// returned if size is non-zero.
//
// The function returns the following values:
//
//   - guint8s (optional): a pointer to the byte data, or NULL.
//
func (bytes *Bytes) Data() []byte {
	var _arg0 *C.GBytes       // out
	var _cret C.gconstpointer // in
	var _arg1 C.gsize         // in

	_arg0 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(bytes)))

	_cret = C.g_bytes_get_data(_arg0, &_arg1)
	runtime.KeepAlive(bytes)

	var _guint8s []byte // out

	_guint8s = make([]byte, _arg1)
	copy(_guint8s, unsafe.Slice((*byte)(unsafe.Pointer(_cret)), _arg1))

	return _guint8s
}

// Size: get the size of the byte data in the #GBytes.
//
// This function will always return the same value for a given #GBytes.
//
// The function returns the following values:
//
//   - gsize: size.
//
func (bytes *Bytes) Size() uint {
	var _arg0 *C.GBytes // out
	var _cret C.gsize   // in

	_arg0 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(bytes)))

	_cret = C.g_bytes_get_size(_arg0)
	runtime.KeepAlive(bytes)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// Hash creates an integer hash code for the byte data in the #GBytes.
//
// This function can be passed to g_hash_table_new() as the key_hash_func
// parameter, when using non-NULL #GBytes pointers as keys in a Table.
//
// The function returns the following values:
//
//   - guint: hash value corresponding to the key.
//
func (bytes *Bytes) Hash() uint {
	var _arg0 C.gconstpointer // out
	var _cret C.guint         // in

	_arg0 = *(*C.gconstpointer)(gextras.StructNative(unsafe.Pointer(bytes)))

	_cret = C.g_bytes_hash(_arg0)
	runtime.KeepAlive(bytes)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// NewFromBytes creates a #GBytes which is a subsection of another #GBytes.
// The offset + length may not be longer than the size of bytes.
//
// A reference to bytes will be held by the newly created #GBytes until the byte
// data is no longer needed.
//
// Since 2.56, if offset is 0 and length matches the size of bytes, then bytes
// will be returned with the reference count incremented by 1. If bytes is a
// slice of another #GBytes, then the resulting #GBytes will reference the same
// #GBytes instead of bytes. This allows consumers to simplify the usage of
// #GBytes when asynchronously writing to streams.
//
// The function takes the following parameters:
//
//   - offset which subsection starts at.
//   - length of subsection.
//
// The function returns the following values:
//
//   - ret: new #GBytes.
//
func (bytes *Bytes) NewFromBytes(offset uint, length uint) *Bytes {
	var _arg0 *C.GBytes // out
	var _arg1 C.gsize   // out
	var _arg2 C.gsize   // out
	var _cret *C.GBytes // in

	_arg0 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(bytes)))
	_arg1 = C.gsize(offset)
	_arg2 = C.gsize(length)

	_cret = C.g_bytes_new_from_bytes(_arg0, _arg1, _arg2)
	runtime.KeepAlive(bytes)
	runtime.KeepAlive(offset)
	runtime.KeepAlive(length)

	var _ret *Bytes // out

	_ret = (*Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_ret)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	return _ret
}
