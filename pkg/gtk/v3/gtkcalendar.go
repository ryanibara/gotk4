// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_calendar_get_type()), F: marshalCalendar},
	})
}

// Calendar is a widget that displays a Gregorian calendar, one month at a time.
// It can be created with gtk_calendar_new().
//
// The month and year currently displayed can be altered with
// gtk_calendar_select_month(). The exact day can be selected from the displayed
// month using gtk_calendar_select_day().
//
// To place a visual marker on a particular day, use gtk_calendar_mark_day() and
// to remove the marker, gtk_calendar_unmark_day(). Alternative, all marks can
// be cleared with gtk_calendar_clear_marks().
//
// The way in which the calendar itself is displayed can be altered using
// gtk_calendar_set_display_options().
//
// The selected date can be retrieved from a Calendar using
// gtk_calendar_get_date().
//
// Users should be aware that, although the Gregorian calendar is the legal
// calendar in most countries, it was adopted progressively between 1582 and
// 1929. Display before these dates is likely to be historically incorrect.
type Calendar interface {
	Widget
	Buildable

	// ClearMarks: remove all visual markers.
	ClearMarks()
	// Date obtains the selected date from a Calendar.
	Date() (year uint, month uint, day uint)
	// DayIsMarked returns if the @day of the @calendar is already marked.
	DayIsMarked(day uint) bool
	// DetailHeightRows queries the height of detail cells, in rows. See
	// Calendar:detail-width-chars.
	DetailHeightRows() int
	// DetailWidthChars queries the width of detail cells, in characters. See
	// Calendar:detail-width-chars.
	DetailWidthChars() int
	// MarkDay places a visual marker on a particular day.
	MarkDay(day uint)
	// SelectDay selects a day from the current month.
	SelectDay(day uint)
	// SelectMonth shifts the calendar to a different month.
	SelectMonth(month uint, year uint)
	// SetDetailHeightRows updates the height of detail cells. See
	// Calendar:detail-height-rows.
	SetDetailHeightRows(rows int)
	// SetDetailWidthChars updates the width of detail cells. See
	// Calendar:detail-width-chars.
	SetDetailWidthChars(chars int)
	// SetDisplayOptions sets display options (whether to display the heading
	// and the month headings).
	SetDisplayOptions(flags CalendarDisplayOptions)
	// UnmarkDay removes the visual marker from a particular day.
	UnmarkDay(day uint)
}

// calendar implements the Calendar interface.
type calendar struct {
	Widget
	Buildable
}

var _ Calendar = (*calendar)(nil)

// WrapCalendar wraps a GObject to the right type. It is
// primarily used internally.
func WrapCalendar(obj *externglib.Object) Calendar {
	return Calendar{
		Widget:    WrapWidget(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalCalendar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCalendar(obj), nil
}

// ClearMarks: remove all visual markers.
func (c calendar) ClearMarks() {
	var _arg0 *C.GtkCalendar

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(c.Native()))

	C.gtk_calendar_clear_marks(_arg0)
}

// Date obtains the selected date from a Calendar.
func (c calendar) Date() (year uint, month uint, day uint) {
	var _arg0 *C.GtkCalendar

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(c.Native()))

	var _arg1 C.guint
	var _arg2 C.guint
	var _arg3 C.guint

	C.gtk_calendar_get_date(_arg0, &_arg1, &_arg2, &_arg3)

	var _year uint
	var _month uint
	var _day uint

	_year = (uint)(_arg1)
	_month = (uint)(_arg2)
	_day = (uint)(_arg3)

	return _year, _month, _day
}

// DayIsMarked returns if the @day of the @calendar is already marked.
func (c calendar) DayIsMarked(day uint) bool {
	var _arg0 *C.GtkCalendar
	var _arg1 C.guint

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(c.Native()))
	_arg1 = C.guint(day)

	var _cret C.gboolean

	_cret = C.gtk_calendar_get_day_is_marked(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// DetailHeightRows queries the height of detail cells, in rows. See
// Calendar:detail-width-chars.
func (c calendar) DetailHeightRows() int {
	var _arg0 *C.GtkCalendar

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(c.Native()))

	var _cret C.gint

	_cret = C.gtk_calendar_get_detail_height_rows(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// DetailWidthChars queries the width of detail cells, in characters. See
// Calendar:detail-width-chars.
func (c calendar) DetailWidthChars() int {
	var _arg0 *C.GtkCalendar

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(c.Native()))

	var _cret C.gint

	_cret = C.gtk_calendar_get_detail_width_chars(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// MarkDay places a visual marker on a particular day.
func (c calendar) MarkDay(day uint) {
	var _arg0 *C.GtkCalendar
	var _arg1 C.guint

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(c.Native()))
	_arg1 = C.guint(day)

	C.gtk_calendar_mark_day(_arg0, _arg1)
}

// SelectDay selects a day from the current month.
func (c calendar) SelectDay(day uint) {
	var _arg0 *C.GtkCalendar
	var _arg1 C.guint

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(c.Native()))
	_arg1 = C.guint(day)

	C.gtk_calendar_select_day(_arg0, _arg1)
}

// SelectMonth shifts the calendar to a different month.
func (c calendar) SelectMonth(month uint, year uint) {
	var _arg0 *C.GtkCalendar
	var _arg1 C.guint
	var _arg2 C.guint

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(c.Native()))
	_arg1 = C.guint(month)
	_arg2 = C.guint(year)

	C.gtk_calendar_select_month(_arg0, _arg1, _arg2)
}

// SetDetailHeightRows updates the height of detail cells. See
// Calendar:detail-height-rows.
func (c calendar) SetDetailHeightRows(rows int) {
	var _arg0 *C.GtkCalendar
	var _arg1 C.gint

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(rows)

	C.gtk_calendar_set_detail_height_rows(_arg0, _arg1)
}

// SetDetailWidthChars updates the width of detail cells. See
// Calendar:detail-width-chars.
func (c calendar) SetDetailWidthChars(chars int) {
	var _arg0 *C.GtkCalendar
	var _arg1 C.gint

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(chars)

	C.gtk_calendar_set_detail_width_chars(_arg0, _arg1)
}

// SetDisplayOptions sets display options (whether to display the heading
// and the month headings).
func (c calendar) SetDisplayOptions(flags CalendarDisplayOptions) {
	var _arg0 *C.GtkCalendar
	var _arg1 C.GtkCalendarDisplayOptions

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(c.Native()))
	_arg1 = (C.GtkCalendarDisplayOptions)(flags)

	C.gtk_calendar_set_display_options(_arg0, _arg1)
}

// UnmarkDay removes the visual marker from a particular day.
func (c calendar) UnmarkDay(day uint) {
	var _arg0 *C.GtkCalendar
	var _arg1 C.guint

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(c.Native()))
	_arg1 = C.guint(day)

	C.gtk_calendar_unmark_day(_arg0, _arg1)
}
