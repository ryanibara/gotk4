// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// RandomDouble returns a random #gdouble equally distributed over the range
// [0..1).
//
// The function returns the following values:
//
//   - gdouble: random number.
//
func RandomDouble() float64 {
	var _cret C.gdouble // in

	_cret = C.g_random_double()

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// RandomDoubleRange returns a random #gdouble equally distributed over the
// range [begin..end).
//
// The function takes the following parameters:
//
//   - begin: lower closed bound of the interval.
//   - end: upper open bound of the interval.
//
// The function returns the following values:
//
//   - gdouble: random number.
//
func RandomDoubleRange(begin, end float64) float64 {
	var _arg1 C.gdouble // out
	var _arg2 C.gdouble // out
	var _cret C.gdouble // in

	_arg1 = C.gdouble(begin)
	_arg2 = C.gdouble(end)

	_cret = C.g_random_double_range(_arg1, _arg2)
	runtime.KeepAlive(begin)
	runtime.KeepAlive(end)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// RandomInt: return a random #guint32 equally distributed over the range
// [0..2^32-1].
//
// The function returns the following values:
//
//   - guint32: random number.
//
func RandomInt() uint32 {
	var _cret C.guint32 // in

	_cret = C.g_random_int()

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// RandomIntRange returns a random #gint32 equally distributed over the range
// [begin..end-1].
//
// The function takes the following parameters:
//
//   - begin: lower closed bound of the interval.
//   - end: upper open bound of the interval.
//
// The function returns the following values:
//
//   - gint32: random number.
//
func RandomIntRange(begin, end int32) int32 {
	var _arg1 C.gint32 // out
	var _arg2 C.gint32 // out
	var _cret C.gint32 // in

	_arg1 = C.gint32(begin)
	_arg2 = C.gint32(end)

	_cret = C.g_random_int_range(_arg1, _arg2)
	runtime.KeepAlive(begin)
	runtime.KeepAlive(end)

	var _gint32 int32 // out

	_gint32 = int32(_cret)

	return _gint32
}

// RandomSetSeed sets the seed for the global random number generator, which is
// used by the g_random_* functions, to seed.
//
// The function takes the following parameters:
//
//   - seed: value to reinitialize the global random number generator.
//
func RandomSetSeed(seed uint32) {
	var _arg1 C.guint32 // out

	_arg1 = C.guint32(seed)

	C.g_random_set_seed(_arg1)
	runtime.KeepAlive(seed)
}
