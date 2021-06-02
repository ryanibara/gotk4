// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// BitsetIterInitAt initializes @iter to point to @target.
//
// If @target is not found, finds the next value after it. If no value >=
// @target exists in @set, this function returns false.
func BitsetIterInitAt(set *Bitset, target uint) (iter BitsetIter, value uint, ok bool) {
	var arg1 *C.GtkBitsetIter // out
	var arg2 *C.GtkBitset
	var arg3 C.guint
	var arg4 *C.guint // out

	arg2 = (*C.GtkBitset)(set.Native())
	arg3 = C.guint(target)

	ret := C.gtk_bitset_iter_init_at(&arg1, arg2, arg3, &arg4)

	var ret0 *BitsetIter
	var ret1 uint
	var ret2 bool

	{
		ret0 = WrapBitsetIter(unsafe.Pointer(arg1))
	}

	ret1 = uint(arg4)

	ret2 = C.bool(ret) != 0

	return ret0, ret1, ret2
}

// BitsetIterInitFirst initializes an iterator for @set and points it to the
// first value in @set.
//
// If @set is empty, false is returned and @value is set to G_MAXUINT.
func BitsetIterInitFirst(set *Bitset) (iter BitsetIter, value uint, ok bool) {
	var arg1 *C.GtkBitsetIter // out
	var arg2 *C.GtkBitset
	var arg3 *C.guint // out

	arg2 = (*C.GtkBitset)(set.Native())

	ret := C.gtk_bitset_iter_init_first(&arg1, arg2, &arg3)

	var ret0 *BitsetIter
	var ret1 uint
	var ret2 bool

	{
		ret0 = WrapBitsetIter(unsafe.Pointer(arg1))
	}

	ret1 = uint(arg3)

	ret2 = C.bool(ret) != 0

	return ret0, ret1, ret2
}

// BitsetIterInitLast initializes an iterator for @set and points it to the last
// value in @set.
//
// If @set is empty, false is returned.
func BitsetIterInitLast(set *Bitset) (iter BitsetIter, value uint, ok bool) {
	var arg1 *C.GtkBitsetIter // out
	var arg2 *C.GtkBitset
	var arg3 *C.guint // out

	arg2 = (*C.GtkBitset)(set.Native())

	ret := C.gtk_bitset_iter_init_last(&arg1, arg2, &arg3)

	var ret0 *BitsetIter
	var ret1 uint
	var ret2 bool

	{
		ret0 = WrapBitsetIter(unsafe.Pointer(arg1))
	}

	ret1 = uint(arg3)

	ret2 = C.bool(ret) != 0

	return ret0, ret1, ret2
}

// BitsetIter: an opaque, stack-allocated struct for iterating over the elements
// of a `GtkBitset`.
//
// Before a `GtkBitsetIter` can be used, it needs to be initialized with
// [func@Gtk.BitsetIter.init_first], [func@Gtk.BitsetIter.init_last] or
// [func@Gtk.BitsetIter.init_at].
type BitsetIter struct {
	native C.GtkBitsetIter
}

// WrapBitsetIter wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBitsetIter(ptr unsafe.Pointer) *BitsetIter {
	if ptr == nil {
		return nil
	}

	return (*BitsetIter)(ptr)
}

func marshalBitsetIter(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapBitsetIter(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (b *BitsetIter) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// Value gets the current value that @iter points to.
//
// If @iter is not valid and [method@Gtk.BitsetIter.is_valid] returns false,
// this function returns 0.
func (i *BitsetIter) Value() uint {
	var arg0 *C.GtkBitsetIter

	arg0 = (*C.GtkBitsetIter)(i.Native())

	ret := C.gtk_bitset_iter_get_value(arg0)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// IsValid checks if @iter points to a valid value.
func (i *BitsetIter) IsValid() bool {
	var arg0 *C.GtkBitsetIter

	arg0 = (*C.GtkBitsetIter)(i.Native())

	ret := C.gtk_bitset_iter_is_valid(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Next moves @iter to the next value in the set.
//
// If it was already pointing to the last value in the set, false is returned
// and @iter is invalidated.
func (i *BitsetIter) Next() (value uint, ok bool) {
	var arg0 *C.GtkBitsetIter
	var arg1 *C.guint // out

	arg0 = (*C.GtkBitsetIter)(i.Native())

	ret := C.gtk_bitset_iter_next(arg0, &arg1)

	var ret0 uint
	var ret1 bool

	ret0 = uint(arg1)

	ret1 = C.bool(ret) != 0

	return ret0, ret1
}

// Previous moves @iter to the previous value in the set.
//
// If it was already pointing to the first value in the set, false is returned
// and @iter is invalidated.
func (i *BitsetIter) Previous() (value uint, ok bool) {
	var arg0 *C.GtkBitsetIter
	var arg1 *C.guint // out

	arg0 = (*C.GtkBitsetIter)(i.Native())

	ret := C.gtk_bitset_iter_previous(arg0, &arg1)

	var ret0 uint
	var ret1 bool

	ret0 = uint(arg1)

	ret1 = C.bool(ret) != 0

	return ret0, ret1
}
