// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void callbackDelete(gpointer);
// gchar* _gotk4_gtk3_CalendarDetailFunc(GtkCalendar*, guint, guint, guint, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_calendar_display_options_get_type()), F: marshalCalendarDisplayOptions},
		{T: externglib.Type(C.gtk_calendar_get_type()), F: marshalCalendarrer},
	})
}

// CalendarDisplayOptions: these options can be used to influence the display
// and behaviour of a Calendar.
type CalendarDisplayOptions int

const (
	// CalendarShowHeading specifies that the month and year should be
	// displayed.
	CalendarShowHeading CalendarDisplayOptions = 0b1
	// CalendarShowDayNames specifies that three letter day descriptions should
	// be present.
	CalendarShowDayNames CalendarDisplayOptions = 0b10
	// CalendarNoMonthChange prevents the user from switching months with the
	// calendar.
	CalendarNoMonthChange CalendarDisplayOptions = 0b100
	// CalendarShowWeekNumbers displays each week numbers of the current year,
	// down the left side of the calendar.
	CalendarShowWeekNumbers CalendarDisplayOptions = 0b1000
	// CalendarShowDetails: just show an indicator, not the full details text
	// when details are provided. See gtk_calendar_set_detail_func().
	CalendarShowDetails CalendarDisplayOptions = 0b100000
)

func marshalCalendarDisplayOptions(p uintptr) (interface{}, error) {
	return CalendarDisplayOptions(C.g_value_get_flags((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the names in string for CalendarDisplayOptions.
func (c CalendarDisplayOptions) String() string {
	if c == 0 {
		return "CalendarDisplayOptions(0)"
	}

	var builder strings.Builder
	builder.Grow(106)

	for c != 0 {
		next := c & (c - 1)
		bit := c - next

		switch bit {
		case CalendarShowHeading:
			builder.WriteString("ShowHeading|")
		case CalendarShowDayNames:
			builder.WriteString("ShowDayNames|")
		case CalendarNoMonthChange:
			builder.WriteString("NoMonthChange|")
		case CalendarShowWeekNumbers:
			builder.WriteString("ShowWeekNumbers|")
		case CalendarShowDetails:
			builder.WriteString("ShowDetails|")
		default:
			builder.WriteString(fmt.Sprintf("CalendarDisplayOptions(0b%b)|", bit))
		}

		c = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if c contains other.
func (c CalendarDisplayOptions) Has(other CalendarDisplayOptions) bool {
	return (c & other) == other
}

// CalendarDetailFunc: this kind of functions provide Pango markup with detail
// information for the specified day. Examples for such details are holidays or
// appointments. The function returns NULL when no information is available.
type CalendarDetailFunc func(calendar *Calendar, year uint, month uint, day uint) (utf8 string)

//export _gotk4_gtk3_CalendarDetailFunc
func _gotk4_gtk3_CalendarDetailFunc(arg0 *C.GtkCalendar, arg1 C.guint, arg2 C.guint, arg3 C.guint, arg4 C.gpointer) (cret *C.gchar) {
	v := gbox.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	var calendar *Calendar // out
	var year uint          // out
	var month uint         // out
	var day uint           // out

	calendar = wrapCalendar(externglib.Take(unsafe.Pointer(arg0)))
	year = uint(arg1)
	month = uint(arg2)
	day = uint(arg3)

	fn := v.(CalendarDetailFunc)
	utf8 := fn(calendar, year, month, day)

	if utf8 != "" {
		cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))
	}

	return cret
}

// CalendarOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CalendarOverrider interface {
	DaySelected()
	DaySelectedDoubleClick()
	MonthChanged()
	NextMonth()
	NextYear()
	PrevMonth()
	PrevYear()
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
type Calendar struct {
	Widget
}

func wrapCalendar(obj *externglib.Object) *Calendar {
	return &Calendar{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalCalendarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCalendar(obj), nil
}

// NewCalendar creates a new calendar, with the current date being selected.
func NewCalendar() *Calendar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_calendar_new()

	var _calendar *Calendar // out

	_calendar = wrapCalendar(externglib.Take(unsafe.Pointer(_cret)))

	return _calendar
}

// ClearMarks: remove all visual markers.
func (calendar *Calendar) ClearMarks() {
	var _arg0 *C.GtkCalendar // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(calendar.Native()))

	C.gtk_calendar_clear_marks(_arg0)
	runtime.KeepAlive(calendar)
}

// Date obtains the selected date from a Calendar.
func (calendar *Calendar) Date() (year uint, month uint, day uint) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.guint        // in
	var _arg2 C.guint        // in
	var _arg3 C.guint        // in

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(calendar.Native()))

	C.gtk_calendar_get_date(_arg0, &_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(calendar)

	var _year uint  // out
	var _month uint // out
	var _day uint   // out

	_year = uint(_arg1)
	_month = uint(_arg2)
	_day = uint(_arg3)

	return _year, _month, _day
}

// DayIsMarked returns if the day of the calendar is already marked.
func (calendar *Calendar) DayIsMarked(day uint) bool {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.guint        // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(calendar.Native()))
	_arg1 = C.guint(day)

	_cret = C.gtk_calendar_get_day_is_marked(_arg0, _arg1)
	runtime.KeepAlive(calendar)
	runtime.KeepAlive(day)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DetailHeightRows queries the height of detail cells, in rows. See
// Calendar:detail-width-chars.
func (calendar *Calendar) DetailHeightRows() int {
	var _arg0 *C.GtkCalendar // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(calendar.Native()))

	_cret = C.gtk_calendar_get_detail_height_rows(_arg0)
	runtime.KeepAlive(calendar)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// DetailWidthChars queries the width of detail cells, in characters. See
// Calendar:detail-width-chars.
func (calendar *Calendar) DetailWidthChars() int {
	var _arg0 *C.GtkCalendar // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(calendar.Native()))

	_cret = C.gtk_calendar_get_detail_width_chars(_arg0)
	runtime.KeepAlive(calendar)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// DisplayOptions returns the current display options of calendar.
func (calendar *Calendar) DisplayOptions() CalendarDisplayOptions {
	var _arg0 *C.GtkCalendar              // out
	var _cret C.GtkCalendarDisplayOptions // in

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(calendar.Native()))

	_cret = C.gtk_calendar_get_display_options(_arg0)
	runtime.KeepAlive(calendar)

	var _calendarDisplayOptions CalendarDisplayOptions // out

	_calendarDisplayOptions = CalendarDisplayOptions(_cret)

	return _calendarDisplayOptions
}

// MarkDay places a visual marker on a particular day.
func (calendar *Calendar) MarkDay(day uint) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(calendar.Native()))
	_arg1 = C.guint(day)

	C.gtk_calendar_mark_day(_arg0, _arg1)
	runtime.KeepAlive(calendar)
	runtime.KeepAlive(day)
}

// SelectDay selects a day from the current month.
func (calendar *Calendar) SelectDay(day uint) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(calendar.Native()))
	_arg1 = C.guint(day)

	C.gtk_calendar_select_day(_arg0, _arg1)
	runtime.KeepAlive(calendar)
	runtime.KeepAlive(day)
}

// SelectMonth shifts the calendar to a different month.
func (calendar *Calendar) SelectMonth(month uint, year uint) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.guint        // out
	var _arg2 C.guint        // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(calendar.Native()))
	_arg1 = C.guint(month)
	_arg2 = C.guint(year)

	C.gtk_calendar_select_month(_arg0, _arg1, _arg2)
	runtime.KeepAlive(calendar)
	runtime.KeepAlive(month)
	runtime.KeepAlive(year)
}

// SetDetailFunc installs a function which provides Pango markup with detail
// information for each day. Examples for such details are holidays or
// appointments. That information is shown below each day when
// Calendar:show-details is set. A tooltip containing with full detail
// information is provided, if the entire text should not fit into the details
// area, or if Calendar:show-details is not set.
//
// The size of the details area can be restricted by setting the
// Calendar:detail-width-chars and Calendar:detail-height-rows properties.
func (calendar *Calendar) SetDetailFunc(fn CalendarDetailFunc) {
	var _arg0 *C.GtkCalendar          // out
	var _arg1 C.GtkCalendarDetailFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(calendar.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_CalendarDetailFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_calendar_set_detail_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(calendar)
	runtime.KeepAlive(fn)
}

// SetDetailHeightRows updates the height of detail cells. See
// Calendar:detail-height-rows.
func (calendar *Calendar) SetDetailHeightRows(rows int) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(calendar.Native()))
	_arg1 = C.gint(rows)

	C.gtk_calendar_set_detail_height_rows(_arg0, _arg1)
	runtime.KeepAlive(calendar)
	runtime.KeepAlive(rows)
}

// SetDetailWidthChars updates the width of detail cells. See
// Calendar:detail-width-chars.
func (calendar *Calendar) SetDetailWidthChars(chars int) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(calendar.Native()))
	_arg1 = C.gint(chars)

	C.gtk_calendar_set_detail_width_chars(_arg0, _arg1)
	runtime.KeepAlive(calendar)
	runtime.KeepAlive(chars)
}

// SetDisplayOptions sets display options (whether to display the heading and
// the month headings).
func (calendar *Calendar) SetDisplayOptions(flags CalendarDisplayOptions) {
	var _arg0 *C.GtkCalendar              // out
	var _arg1 C.GtkCalendarDisplayOptions // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(calendar.Native()))
	_arg1 = C.GtkCalendarDisplayOptions(flags)

	C.gtk_calendar_set_display_options(_arg0, _arg1)
	runtime.KeepAlive(calendar)
	runtime.KeepAlive(flags)
}

// UnmarkDay removes the visual marker from a particular day.
func (calendar *Calendar) UnmarkDay(day uint) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(calendar.Native()))
	_arg1 = C.guint(day)

	C.gtk_calendar_unmark_day(_arg0, _arg1)
	runtime.KeepAlive(calendar)
	runtime.KeepAlive(day)
}

// ConnectDaySelected: emitted when the user selects a day.
func (c *Calendar) ConnectDaySelected(f func()) glib.SignalHandle {
	return c.Connect("day-selected", f)
}

// ConnectDaySelectedDoubleClick: emitted when the user double-clicks a day.
func (c *Calendar) ConnectDaySelectedDoubleClick(f func()) glib.SignalHandle {
	return c.Connect("day-selected-double-click", f)
}

// ConnectMonthChanged: emitted when the user clicks a button to change the
// selected month on a calendar.
func (c *Calendar) ConnectMonthChanged(f func()) glib.SignalHandle {
	return c.Connect("month-changed", f)
}

// ConnectNextMonth: emitted when the user switched to the next month.
func (c *Calendar) ConnectNextMonth(f func()) glib.SignalHandle {
	return c.Connect("next-month", f)
}

// ConnectNextYear: emitted when user switched to the next year.
func (c *Calendar) ConnectNextYear(f func()) glib.SignalHandle {
	return c.Connect("next-year", f)
}

// ConnectPrevMonth: emitted when the user switched to the previous month.
func (c *Calendar) ConnectPrevMonth(f func()) glib.SignalHandle {
	return c.Connect("prev-month", f)
}

// ConnectPrevYear: emitted when user switched to the previous year.
func (c *Calendar) ConnectPrevYear(f func()) glib.SignalHandle {
	return c.Connect("prev-year", f)
}
