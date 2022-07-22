// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// BIG_ENDIAN specifies one of the possible types of byte order. See BYTE_ORDER.
const BIG_ENDIAN = 4321

// E: base of natural logarithms.
const E = 2.718282

// IEEE754_DOUBLE_BIAS bias by which exponents in double-precision floats are
// offset.
const IEEE754_DOUBLE_BIAS = 1023

// IEEE754_FLOAT_BIAS bias by which exponents in single-precision floats are
// offset.
const IEEE754_FLOAT_BIAS = 127

// LITTLE_ENDIAN specifies one of the possible types of byte order. See
// BYTE_ORDER.
const LITTLE_ENDIAN = 1234

// LN10: natural logarithm of 10.
const LN10 = 2.302585

// LN2: natural logarithm of 2.
const LN2 = 0.693147

// LOG_2_BASE_10: multiplying the base 2 exponent by this number yields the base
// 10 exponent.
const LOG_2_BASE_10 = 0.301030

// MAXINT16: maximum value which can be held in a #gint16.
const MAXINT16 = 32767

// MAXINT32: maximum value which can be held in a #gint32.
const MAXINT32 = 2147483647

// MAXINT64: maximum value which can be held in a #gint64.
const MAXINT64 = 9223372036854775807

// MAXINT8: maximum value which can be held in a #gint8.
const MAXINT8 = 127

// MAXUINT16: maximum value which can be held in a #guint16.
const MAXUINT16 = 65535

// MAXUINT32: maximum value which can be held in a #guint32.
const MAXUINT32 = 4294967295

// MAXUINT64: maximum value which can be held in a #guint64.
const MAXUINT64 = 18446744073709551615

// MAXUINT8: maximum value which can be held in a #guint8.
const MAXUINT8 = 255

// MININT16: minimum value which can be held in a #gint16.
const MININT16 = -32768

// MININT32: minimum value which can be held in a #gint32.
const MININT32 = -2147483648

// MININT64: minimum value which can be held in a #gint64.
const MININT64 = -9223372036854775808

// MININT8: minimum value which can be held in a #gint8.
const MININT8 = -128

// PDP_ENDIAN specifies one of the possible types of byte order (currently
// unused). See BYTE_ORDER.
const PDP_ENDIAN = 3412

// PI: value of pi (ratio of circle's circumference to its diameter).
const PI = 3.141593

// PI_2: pi divided by 2.
const PI_2 = 1.570796

// PI_4: pi divided by 4.
const PI_4 = 0.785398

// SQRT2: square root of two.
const SQRT2 = 1.414214

// CompareDataFunc specifies the type of a comparison function used to compare
// two values. The function should return a negative integer if the first value
// comes before the second, 0 if they are equal, or a positive integer if the
// first value comes after the second.
type CompareDataFunc func(a, b unsafe.Pointer) (gint int)

//export _gotk4_glib2_CompareDataFunc
func _gotk4_glib2_CompareDataFunc(arg1 C.gconstpointer, arg2 C.gconstpointer, arg3 C.gpointer) (cret C.gint) {
	var fn CompareDataFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(CompareDataFunc)
	}

	var _a unsafe.Pointer // out
	var _b unsafe.Pointer // out

	_a = (unsafe.Pointer)(unsafe.Pointer(arg1))
	_b = (unsafe.Pointer)(unsafe.Pointer(arg2))

	gint := fn(_a, _b)

	cret = C.gint(gint)

	return cret
}

// Func specifies the type of functions passed to g_list_foreach() and
// g_slist_foreach().
type Func func(data unsafe.Pointer)

//export _gotk4_glib2_Func
func _gotk4_glib2_Func(arg1 C.gpointer, arg2 C.gpointer) {
	var fn Func
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(Func)
	}

	var _data unsafe.Pointer // out

	_data = (unsafe.Pointer)(unsafe.Pointer(arg1))

	fn(_data)
}

// HFunc specifies the type of the function passed to g_hash_table_foreach(). It
// is called with each key/value pair, together with the user_data parameter
// which is passed to g_hash_table_foreach().
type HFunc func(key, value unsafe.Pointer)

//export _gotk4_glib2_HFunc
func _gotk4_glib2_HFunc(arg1 C.gpointer, arg2 C.gpointer, arg3 C.gpointer) {
	var fn HFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(HFunc)
	}

	var _key unsafe.Pointer   // out
	var _value unsafe.Pointer // out

	_key = (unsafe.Pointer)(unsafe.Pointer(arg1))
	_value = (unsafe.Pointer)(unsafe.Pointer(arg2))

	fn(_key, _value)
}

// TimeVal represents a precise time, with seconds and microseconds. Similar to
// the struct timeval returned by the gettimeofday() UNIX system call.
//
// GLib is attempting to unify around the use of 64-bit integers to represent
// microsecond-precision time. As such, this type will be removed from a future
// version of GLib. A consequence of using glong for tv_sec is that on 32-bit
// systems GTimeVal is subject to the year 2038 problem.
//
// Deprecated: Use Time or #guint64 instead.
//
// An instance of this type is always passed by reference.
type TimeVal struct {
	*timeVal
}

// timeVal is the struct that's finalized.
type timeVal struct {
	native *C.GTimeVal
}

// NewTimeVal creates a new TimeVal instance from the given
// fields. Beware that this function allocates on the Go heap; be careful
// when using it!
func NewTimeVal(tvSec, tvUsec int32) TimeVal {
	var f0 C.glong // out
	f0 = C.glong(tvSec)
	var f1 C.glong // out
	f1 = C.glong(tvUsec)

	v := C.GTimeVal{
		tv_sec:  f0,
		tv_usec: f1,
	}

	return *(*TimeVal)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

// TvSec: seconds.
func (t *TimeVal) TvSec() int32 {
	valptr := &t.native.tv_sec
	var _v int32 // out
	_v = int32(*valptr)
	return _v
}

// TvUsec: microseconds.
func (t *TimeVal) TvUsec() int32 {
	valptr := &t.native.tv_usec
	var _v int32 // out
	_v = int32(*valptr)
	return _v
}

// TvSec: seconds.
func (t *TimeVal) SetTvSec(tvSec int32) {
	valptr := &t.native.tv_sec
	*valptr = C.glong(tvSec)
}

// TvUsec: microseconds.
func (t *TimeVal) SetTvUsec(tvUsec int32) {
	valptr := &t.native.tv_usec
	*valptr = C.glong(tvUsec)
}

// Add adds the given number of microseconds to time_. microseconds can also be
// negative to decrease the value of time_.
//
// Deprecated: Val is not year-2038-safe. Use guint64 for representing
// microseconds since the epoch, or use Time.
//
// The function takes the following parameters:
//
//    - microseconds: number of microseconds to add to time.
//
func (time_ *TimeVal) Add(microseconds int32) {
	var _arg0 *C.GTimeVal // out
	var _arg1 C.glong     // out

	_arg0 = (*C.GTimeVal)(gextras.StructNative(unsafe.Pointer(time_)))
	_arg1 = C.glong(microseconds)

	C.g_time_val_add(_arg0, _arg1)
	runtime.KeepAlive(time_)
	runtime.KeepAlive(microseconds)
}

// ToISO8601 converts time_ into an RFC 3339 encoded string, relative to the
// Coordinated Universal Time (UTC). This is one of the many formats allowed by
// ISO 8601.
//
// ISO 8601 allows a large number of date/time formats, with or without
// punctuation and optional elements. The format returned by this function is a
// complete date and time, with optional punctuation included, the UTC time zone
// represented as "Z", and the tv_usec part included if and only if it is
// nonzero, i.e. either "YYYY-MM-DDTHH:MM:SSZ" or "YYYY-MM-DDTHH:MM:SS.fffffZ".
//
// This corresponds to the Internet date/time format defined by RFC 3339
// (https://www.ietf.org/rfc/rfc3339.txt), and to either of the two most-precise
// formats defined by the W3C Note Date and Time Formats
// (http://www.w3.org/TR/NOTE-datetime-19980827). Both of these documents are
// profiles of ISO 8601.
//
// Use g_date_time_format() or g_strdup_printf() if a different variation of ISO
// 8601 format is required.
//
// If time_ represents a date which is too large to fit into a struct tm, NULL
// will be returned. This is platform dependent. Note also that since GTimeVal
// stores the number of seconds as a glong, on 32-bit systems it is subject to
// the year 2038 problem. Accordingly, since GLib 2.62, this function has been
// deprecated. Equivalent functionality is available using:
//
//    GDateTime *dt = g_date_time_new_from_unix_utc (time_val);
//    iso8601_string = g_date_time_format_iso8601 (dt);
//    g_date_time_unref (dt);
//
// The return value of g_time_val_to_iso8601() has been nullable since GLib
// 2.54; before then, GLib would crash under the same conditions.
//
// Deprecated: Val is not year-2038-safe. Use g_date_time_format_iso8601(dt)
// instead.
//
// The function returns the following values:
//
//    - utf8 (optional): newly allocated string containing an ISO 8601 date, or
//      NULL if time_ was too large.
//
func (time_ *TimeVal) ToISO8601() string {
	var _arg0 *C.GTimeVal // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GTimeVal)(gextras.StructNative(unsafe.Pointer(time_)))

	_cret = C.g_time_val_to_iso8601(_arg0)
	runtime.KeepAlive(time_)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}
