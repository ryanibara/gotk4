// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_date_time_get_type()), F: marshalDateTime},
	})
}

// TimeSpan: a value representing an interval of time, in microseconds.
type TimeSpan int64

// DateTime: `GDateTime` is an opaque structure whose members cannot be accessed
// directly.
type DateTime struct {
	native C.GDateTime
}

// WrapDateTime wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDateTime(ptr unsafe.Pointer) *DateTime {
	if ptr == nil {
		return nil
	}

	return (*DateTime)(ptr)
}

func marshalDateTime(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDateTime(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DateTime) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// Compare: a comparison function for Times that is suitable as a Func. Both
// Times must be non-nil.
func (d *DateTime) Compare(dt2 DateTime) int {
	var _arg0 C.gpointer
	var _arg1 C.gpointer

	_arg0 = (C.gpointer)(unsafe.Pointer(d.Native()))
	_arg1 = (C.gpointer)(unsafe.Pointer(dt2.Native()))

	var _cret C.gint

	_cret = C.g_date_time_compare(_arg0, _arg1)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Equal checks to see if @dt1 and @dt2 are equal.
//
// Equal here means that they represent the same moment after converting them to
// the same time zone.
func (d *DateTime) Equal(dt2 DateTime) bool {
	var _arg0 C.gpointer
	var _arg1 C.gpointer

	_arg0 = (C.gpointer)(unsafe.Pointer(d.Native()))
	_arg1 = (C.gpointer)(unsafe.Pointer(dt2.Native()))

	var _cret C.gboolean

	_cret = C.g_date_time_equal(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Format creates a newly allocated string representing the requested @format.
//
// The format strings understood by this function are a subset of the strftime()
// format language as specified by C99. The \D, \U and \W conversions are not
// supported, nor is the 'E' modifier. The GNU extensions \k, \l, \s and \P are
// supported, however, as are the '0', '_' and '-' modifiers. The Python
// extension \f is also supported.
//
// In contrast to strftime(), this function always produces a UTF-8 string,
// regardless of the current locale. Note that the rendering of many formats is
// locale-dependent and may not match the strftime() output exactly.
//
// The following format specifiers are supported:
//
// - \a: the abbreviated weekday name according to the current locale - \A: the
// full weekday name according to the current locale - \b: the abbreviated month
// name according to the current locale - \B: the full month name according to
// the current locale - \c: the preferred date and time representation for the
// current locale - \C: the century number (year/100) as a 2-digit integer
// (00-99) - \d: the day of the month as a decimal number (range 01 to 31) - \e:
// the day of the month as a decimal number (range 1 to 31) - \F: equivalent to
// `Y-m-d` (the ISO 8601 date format) - \g: the last two digits of the ISO 8601
// week-based year as a decimal number (00-99). This works well with \V and \u.
// - \G: the ISO 8601 week-based year as a decimal number. This works well with
// \V and \u. - \h: equivalent to \b - \H: the hour as a decimal number using a
// 24-hour clock (range 00 to 23) - \I: the hour as a decimal number using a
// 12-hour clock (range 01 to 12) - \j: the day of the year as a decimal number
// (range 001 to 366) - \k: the hour (24-hour clock) as a decimal number (range
// 0 to 23); single digits are preceded by a blank - \l: the hour (12-hour
// clock) as a decimal number (range 1 to 12); single digits are preceded by a
// blank - \m: the month as a decimal number (range 01 to 12) - \M: the minute
// as a decimal number (range 00 to 59) - \f: the microsecond as a decimal
// number (range 000000 to 999999) - \p: either "AM" or "PM" according to the
// given time value, or the corresponding strings for the current locale. Noon
// is treated as "PM" and midnight as "AM". Use of this format specifier is
// discouraged, as many locales have no concept of AM/PM formatting. Use \c or
// \X instead. - \P: like \p but lowercase: "am" or "pm" or a corresponding
// string for the current locale. Use of this format specifier is discouraged,
// as many locales have no concept of AM/PM formatting. Use \c or \X instead. -
// \r: the time in a.m. or p.m. notation. Use of this format specifier is
// discouraged, as many locales have no concept of AM/PM formatting. Use \c or
// \X instead. - \R: the time in 24-hour notation (\H:\M) - \s: the number of
// seconds since the Epoch, that is, since 1970-01-01 00:00:00 UTC - \S: the
// second as a decimal number (range 00 to 60) - \t: a tab character - \T: the
// time in 24-hour notation with seconds (\H:\M:\S) - \u: the ISO 8601 standard
// day of the week as a decimal, range 1 to 7, Monday being 1. This works well
// with \G and \V. - \V: the ISO 8601 standard week number of the current year
// as a decimal number, range 01 to 53, where week 1 is the first week that has
// at least 4 days in the new year. See g_date_time_get_week_of_year(). This
// works well with \G and \u. - \w: the day of the week as a decimal, range 0 to
// 6, Sunday being 0. This is not the ISO 8601 standard format -- use \u
// instead. - \x: the preferred date representation for the current locale
// without the time - \X: the preferred time representation for the current
// locale without the date - \y: the year as a decimal number without the
// century - \Y: the year as a decimal number including the century - \z: the
// time zone as an offset from UTC (+hhmm) - \%:z: the time zone as an offset
// from UTC (+hh:mm). This is a gnulib strftime() extension. Since: 2.38 -
// \%::z: the time zone as an offset from UTC (+hh:mm:ss). This is a gnulib
// strftime() extension. Since: 2.38 - \%:::z: the time zone as an offset from
// UTC, with : to necessary precision (e.g., -04, +05:30). This is a gnulib
// strftime() extension. Since: 2.38 - \Z: the time zone or name or abbreviation
// - \%\%: a literal \% character
//
// Some conversion specifications can be modified by preceding the conversion
// specifier by one or more modifier characters. The following modifiers are
// supported for many of the numeric conversions:
//
// - O: Use alternative numeric symbols, if the current locale supports those. -
// _: Pad a numeric result with spaces. This overrides the default padding for
// the specifier. - -: Do not pad a numeric result. This overrides the default
// padding for the specifier. - 0: Pad a numeric result with zeros. This
// overrides the default padding for the specifier.
//
// Additionally, when O is used with B, b, or h, it produces the alternative
// form of a month name. The alternative form should be used when the month name
// is used without a day number (e.g., standalone). It is required in some
// languages (Baltic, Slavic, Greek, and more) due to their grammatical rules.
// For other languages there is no difference. \OB is a GNU and BSD strftime()
// extension expected to be added to the future POSIX specification, \Ob and \Oh
// are GNU strftime() extensions. Since: 2.56
func (d *DateTime) Format(format string) string {
	var _arg0 *C.GDateTime
	var _arg1 *C.gchar

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.gchar)(C.CString(format))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.gchar

	_cret = C.g_date_time_format(_arg0, _arg1)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FormatISO8601: format @datetime in ISO 8601 format
// (https://en.wikipedia.org/wiki/ISO_8601), including the date, time and time
// zone, and return that as a UTF-8 encoded string.
//
// Since GLib 2.66, this will output to sub-second precision if needed.
func (d *DateTime) FormatISO8601() string {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _cret *C.gchar

	_cret = C.g_date_time_format_iso8601(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// DayOfMonth retrieves the day of the month represented by @datetime in the
// gregorian calendar.
func (d *DateTime) DayOfMonth() int {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _cret C.gint

	_cret = C.g_date_time_get_day_of_month(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// DayOfWeek retrieves the ISO 8601 day of the week on which @datetime falls (1
// is Monday, 2 is Tuesday... 7 is Sunday).
func (d *DateTime) DayOfWeek() int {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _cret C.gint

	_cret = C.g_date_time_get_day_of_week(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// DayOfYear retrieves the day of the year represented by @datetime in the
// Gregorian calendar.
func (d *DateTime) DayOfYear() int {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _cret C.gint

	_cret = C.g_date_time_get_day_of_year(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Hour retrieves the hour of the day represented by @datetime
func (d *DateTime) Hour() int {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _cret C.gint

	_cret = C.g_date_time_get_hour(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Microsecond retrieves the microsecond of the date represented by @datetime
func (d *DateTime) Microsecond() int {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _cret C.gint

	_cret = C.g_date_time_get_microsecond(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Minute retrieves the minute of the hour represented by @datetime
func (d *DateTime) Minute() int {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _cret C.gint

	_cret = C.g_date_time_get_minute(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Month retrieves the month of the year represented by @datetime in the
// Gregorian calendar.
func (d *DateTime) Month() int {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _cret C.gint

	_cret = C.g_date_time_get_month(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Second retrieves the second of the minute represented by @datetime
func (d *DateTime) Second() int {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _cret C.gint

	_cret = C.g_date_time_get_second(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Seconds retrieves the number of seconds since the start of the last minute,
// including the fractional part.
func (d *DateTime) Seconds() float64 {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _cret C.gdouble

	_cret = C.g_date_time_get_seconds(_arg0)

	var _gdouble float64

	_gdouble = (float64)(_cret)

	return _gdouble
}

// TimezoneAbbreviation determines the time zone abbreviation to be used at the
// time and in the time zone of @datetime.
//
// For example, in Toronto this is currently "EST" during the winter months and
// "EDT" during the summer months when daylight savings time is in effect.
func (d *DateTime) TimezoneAbbreviation() string {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _cret *C.gchar

	_cret = C.g_date_time_get_timezone_abbreviation(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// WeekNumberingYear returns the ISO 8601 week-numbering year in which the week
// containing @datetime falls.
//
// This function, taken together with g_date_time_get_week_of_year() and
// g_date_time_get_day_of_week() can be used to determine the full ISO week date
// on which @datetime falls.
//
// This is usually equal to the normal Gregorian year (as returned by
// g_date_time_get_year()), except as detailed below:
//
// For Thursday, the week-numbering year is always equal to the usual calendar
// year. For other days, the number is such that every day within a complete
// week (Monday to Sunday) is contained within the same week-numbering year.
//
// For Monday, Tuesday and Wednesday occurring near the end of the year, this
// may mean that the week-numbering year is one greater than the calendar year
// (so that these days have the same week-numbering year as the Thursday
// occurring early in the next year).
//
// For Friday, Saturday and Sunday occurring near the start of the year, this
// may mean that the week-numbering year is one less than the calendar year (so
// that these days have the same week-numbering year as the Thursday occurring
// late in the previous year).
//
// An equivalent description is that the week-numbering year is equal to the
// calendar year containing the majority of the days in the current week (Monday
// to Sunday).
//
// Note that January 1 0001 in the proleptic Gregorian calendar is a Monday, so
// this function never returns 0.
func (d *DateTime) WeekNumberingYear() int {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _cret C.gint

	_cret = C.g_date_time_get_week_numbering_year(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// WeekOfYear returns the ISO 8601 week number for the week containing
// @datetime. The ISO 8601 week number is the same for every day of the week
// (from Moday through Sunday). That can produce some unusual results (described
// below).
//
// The first week of the year is week 1. This is the week that contains the
// first Thursday of the year. Equivalently, this is the first week that has
// more than 4 of its days falling within the calendar year.
//
// The value 0 is never returned by this function. Days contained within a year
// but occurring before the first ISO 8601 week of that year are considered as
// being contained in the last week of the previous year. Similarly, the final
// days of a calendar year may be considered as being part of the first ISO 8601
// week of the next year if 4 or more days of that week are contained within the
// new year.
func (d *DateTime) WeekOfYear() int {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _cret C.gint

	_cret = C.g_date_time_get_week_of_year(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Year retrieves the year represented by @datetime in the Gregorian calendar.
func (d *DateTime) Year() int {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _cret C.gint

	_cret = C.g_date_time_get_year(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Ymd retrieves the Gregorian day, month, and year of a given Time.
func (d *DateTime) Ymd() (year int, month int, day int) {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _arg1 C.gint
	var _arg2 C.gint
	var _arg3 C.gint

	C.g_date_time_get_ymd(_arg0, &_arg1, &_arg2, &_arg3)

	var _year int
	var _month int
	var _day int

	_year = (int)(_arg1)
	_month = (int)(_arg2)
	_day = (int)(_arg3)

	return _year, _month, _day
}

// Hash hashes @datetime into a #guint, suitable for use within Table.
func (d *DateTime) Hash() uint {
	var _arg0 C.gpointer

	_arg0 = (C.gpointer)(unsafe.Pointer(d.Native()))

	var _cret C.guint

	_cret = C.g_date_time_hash(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// IsDaylightSavings determines if daylight savings time is in effect at the
// time and in the time zone of @datetime.
func (d *DateTime) IsDaylightSavings() bool {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _cret C.gboolean

	_cret = C.g_date_time_is_daylight_savings(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ToTimeval stores the instant in time that @datetime represents into @tv.
//
// The time contained in a Val is always stored in the form of seconds elapsed
// since 1970-01-01 00:00:00 UTC, regardless of the time zone associated with
// @datetime.
//
// On systems where 'long' is 32bit (ie: all 32bit systems and all Windows
// systems), a Val is incapable of storing the entire range of values that Time
// is capable of expressing. On those systems, this function returns false to
// indicate that the time is out of range.
//
// On systems where 'long' is 64bit, this function never fails.
func (d *DateTime) ToTimeval(tv *TimeVal) bool {
	var _arg0 *C.GDateTime
	var _arg1 *C.GTimeVal

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GTimeVal)(unsafe.Pointer(tv.Native()))

	var _cret C.gboolean

	_cret = C.g_date_time_to_timeval(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ToUnix gives the Unix time corresponding to @datetime, rounding down to the
// nearest second.
//
// Unix time is the number of seconds that have elapsed since 1970-01-01
// 00:00:00 UTC, regardless of the time zone associated with @datetime.
func (d *DateTime) ToUnix() int64 {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	var _cret C.gint64

	_cret = C.g_date_time_to_unix(_arg0)

	var _gint64 int64

	_gint64 = (int64)(_cret)

	return _gint64
}

// Unref: atomically decrements the reference count of @datetime by one.
//
// When the reference count reaches zero, the resources allocated by @datetime
// are freed
func (d *DateTime) Unref() {
	var _arg0 *C.GDateTime

	_arg0 = (*C.GDateTime)(unsafe.Pointer(d.Native()))

	C.g_date_time_unref(_arg0)
}
