// Code generated by girgen. DO NOT EDIT.

package glib

// #cgo pkg-config: glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib.h>
import "C"

// AtomicRefCountCompare: atomically compares the current value of @arc with
// @val.
func AtomicRefCountCompare(arc int, val int) bool {
	var arg1 *C.gatomicrefcount
	var arg2 C.gint

	arg1 = (*C.gatomicrefcount)(arc)
	arg2 = C.gint(val)

	ret := C.g_atomic_ref_count_compare(arg1, arg2)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// AtomicRefCountDec: atomically decreases the reference count.
func AtomicRefCountDec(arc int) bool {
	var arg1 *C.gatomicrefcount

	arg1 = (*C.gatomicrefcount)(arc)

	ret := C.g_atomic_ref_count_dec(arg1)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// AtomicRefCountInc: atomically increases the reference count.
func AtomicRefCountInc(arc int) {
	var arg1 *C.gatomicrefcount

	arg1 = (*C.gatomicrefcount)(arc)

	C.g_atomic_ref_count_inc(arg1)
}

// AtomicRefCountInit initializes a reference count variable.
func AtomicRefCountInit(arc int) {
	var arg1 *C.gatomicrefcount

	arg1 = (*C.gatomicrefcount)(arc)

	C.g_atomic_ref_count_init(arg1)
}

// RefCountCompare compares the current value of @rc with @val.
func RefCountCompare(rc int, val int) bool {
	var arg1 *C.grefcount
	var arg2 C.gint

	arg1 = (*C.grefcount)(rc)
	arg2 = C.gint(val)

	ret := C.g_ref_count_compare(arg1, arg2)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// RefCountDec decreases the reference count.
func RefCountDec(rc int) bool {
	var arg1 *C.grefcount

	arg1 = (*C.grefcount)(rc)

	ret := C.g_ref_count_dec(arg1)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// RefCountInc increases the reference count.
func RefCountInc(rc int) {
	var arg1 *C.grefcount

	arg1 = (*C.grefcount)(rc)

	C.g_ref_count_inc(arg1)
}

// RefCountInit initializes a reference count variable.
func RefCountInit(rc int) {
	var arg1 *C.grefcount

	arg1 = (*C.grefcount)(rc)

	C.g_ref_count_init(arg1)
}
