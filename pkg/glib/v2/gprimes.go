// Code generated by girgen. DO NOT EDIT.

package glib

// #cgo pkg-config: glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// SpacedPrimesClosest gets the smallest prime number from a built-in array of
// primes which is larger than @num. This is used within GLib to calculate the
// optimum size of a Table.
//
// The built-in array of primes ranges from 11 to 13845163 such that each prime
// is approximately 1.5-2 times the previous prime.
func SpacedPrimesClosest(num uint) uint {
	var arg1 C.guint

	arg1 = C.guint(num)

	ret := C.g_spaced_primes_closest(arg1)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}
