// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

// RandomDouble returns a random #gdouble equally distributed over the range
// [0..1).
func RandomDouble() float64 {
	var cret C.gdouble
	var ret1 float64

	cret = C.g_random_double()

	ret1 = C.gdouble(cret)

	return ret1
}

// RandomDoubleRange returns a random #gdouble equally distributed over the
// range [@begin..@end).
func RandomDoubleRange(begin float64, end float64) float64 {
	var arg1 C.gdouble
	var arg2 C.gdouble

	arg1 = C.gdouble(begin)
	arg2 = C.gdouble(end)

	var cret C.gdouble
	var ret1 float64

	cret = C.g_random_double_range(begin, end)

	ret1 = C.gdouble(cret)

	return ret1
}

// RandomInt: return a random #guint32 equally distributed over the range
// [0..2^32-1].
func RandomInt() uint32 {
	var cret C.guint32
	var ret1 uint32

	cret = C.g_random_int()

	ret1 = C.guint32(cret)

	return ret1
}

// RandomIntRange returns a random #gint32 equally distributed over the range
// [@begin..@end-1].
func RandomIntRange(begin int32, end int32) int32 {
	var arg1 C.gint32
	var arg2 C.gint32

	arg1 = C.gint32(begin)
	arg2 = C.gint32(end)

	var cret C.gint32
	var ret1 int32

	cret = C.g_random_int_range(begin, end)

	ret1 = C.gint32(cret)

	return ret1
}

// RandomSetSeed sets the seed for the global random number generator, which is
// used by the g_random_* functions, to @seed.
func RandomSetSeed(seed uint32) {
	var arg1 C.guint32

	arg1 = C.guint32(seed)

	C.g_random_set_seed(seed)
}

// Rand: the GRand struct is an opaque data structure. It should only be
// accessed through the g_rand_* functions.
type Rand struct {
	native C.GRand
}

// WrapRand wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRand(ptr unsafe.Pointer) *Rand {
	if ptr == nil {
		return nil
	}

	return (*Rand)(ptr)
}

func marshalRand(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRand(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *Rand) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Copy copies a #GRand into a new one with the same exact state as before. This
// way you can take a snapshot of the random number generator for replaying
// later.
func (r *Rand) Copy() *Rand {
	var arg0 *C.GRand

	arg0 = (*C.GRand)(unsafe.Pointer(r.Native()))

	var cret *C.GRand
	var ret1 *Rand

	cret = C.g_rand_copy(arg0)

	ret1 = WrapRand(unsafe.Pointer(cret))

	return ret1
}

// Double returns the next random #gdouble from @rand_ equally distributed over
// the range [0..1).
func (r *Rand) Double() float64 {
	var arg0 *C.GRand

	arg0 = (*C.GRand)(unsafe.Pointer(r.Native()))

	var cret C.gdouble
	var ret1 float64

	cret = C.g_rand_double(arg0)

	ret1 = C.gdouble(cret)

	return ret1
}

// DoubleRange returns the next random #gdouble from @rand_ equally distributed
// over the range [@begin..@end).
func (r *Rand) DoubleRange(begin float64, end float64) float64 {
	var arg0 *C.GRand
	var arg1 C.gdouble
	var arg2 C.gdouble

	arg0 = (*C.GRand)(unsafe.Pointer(r.Native()))
	arg1 = C.gdouble(begin)
	arg2 = C.gdouble(end)

	var cret C.gdouble
	var ret1 float64

	cret = C.g_rand_double_range(arg0, begin, end)

	ret1 = C.gdouble(cret)

	return ret1
}

// Free frees the memory allocated for the #GRand.
func (r *Rand) Free() {
	var arg0 *C.GRand

	arg0 = (*C.GRand)(unsafe.Pointer(r.Native()))

	C.g_rand_free(arg0)
}

// Int returns the next random #guint32 from @rand_ equally distributed over the
// range [0..2^32-1].
func (r *Rand) Int() uint32 {
	var arg0 *C.GRand

	arg0 = (*C.GRand)(unsafe.Pointer(r.Native()))

	var cret C.guint32
	var ret1 uint32

	cret = C.g_rand_int(arg0)

	ret1 = C.guint32(cret)

	return ret1
}

// IntRange returns the next random #gint32 from @rand_ equally distributed over
// the range [@begin..@end-1].
func (r *Rand) IntRange(begin int32, end int32) int32 {
	var arg0 *C.GRand
	var arg1 C.gint32
	var arg2 C.gint32

	arg0 = (*C.GRand)(unsafe.Pointer(r.Native()))
	arg1 = C.gint32(begin)
	arg2 = C.gint32(end)

	var cret C.gint32
	var ret1 int32

	cret = C.g_rand_int_range(arg0, begin, end)

	ret1 = C.gint32(cret)

	return ret1
}

// SetSeed sets the seed for the random number generator #GRand to @seed.
func (r *Rand) SetSeed(seed uint32) {
	var arg0 *C.GRand
	var arg1 C.guint32

	arg0 = (*C.GRand)(unsafe.Pointer(r.Native()))
	arg1 = C.guint32(seed)

	C.g_rand_set_seed(arg0, seed)
}

// SetSeedArray initializes the random number generator by an array of longs.
// Array can be of arbitrary size, though only the first 624 values are taken.
// This function is useful if you have many low entropy seeds, or if you require
// more then 32 bits of actual entropy for your application.
func (r *Rand) SetSeedArray(seed uint32, seedLength uint) {
	var arg0 *C.GRand
	var arg1 *C.guint32
	var arg2 C.guint

	arg0 = (*C.GRand)(unsafe.Pointer(r.Native()))
	arg1 = *C.guint32(seed)
	arg2 = C.guint(seedLength)

	C.g_rand_set_seed_array(arg0, seed, seedLength)
}
