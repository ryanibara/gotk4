// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_calendar_get_type()), F: marshalCalendarrer},
	})
}

// Calendar: GtkCalendar is a widget that displays a Gregorian calendar, one
// month at a time.
//
// !An example GtkCalendar (calendar.png)
//
// A GtkCalendar can be created with gtk.Calendar.New.
//
// The date that is currently displayed can be altered with
// gtk.Calendar.SelectDay().
//
// To place a visual marker on a particular day, use gtk.Calendar.MarkDay() and
// to remove the marker, gtk.Calendar.UnmarkDay(). Alternative, all marks can be
// cleared with gtk.Calendar.ClearMarks().
//
// The selected date can be retrieved from a GtkCalendar using
// gtk.Calendar.GetDate().
//
// Users should be aware that, although the Gregorian calendar is the legal
// calendar in most countries, it was adopted progressively between 1582 and
// 1929. Display before these dates is likely to be historically incorrect.
//
// CSS nodes
//
//    calendar.view
//    ├── header
//    │   ├── button
//    │   ├── stack.month
//    │   ├── button
//    │   ├── button
//    │   ├── label.year
//    │   ╰── button
//    ╰── grid
//        ╰── label[.day-name][.week-number][.day-number][.other-month][.today]
//
//
// GtkCalendar has a main node with name calendar. It contains a subnode called
// header containing the widgets for switching between years and months.
//
// The grid subnode contains all day labels, including week numbers on the left
// (marked with the .week-number css class) and day names on top (marked with
// the .day-name css class).
//
// Day labels that belong to the previous or next month get the .other-month
// style class. The label of the current day get the .today style class.
//
// Marked day labels get the :selected state assigned.
type Calendar struct {
	Widget
}

func wrapCalendar(obj *externglib.Object) *Calendar {
	return &Calendar{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalCalendarrer(p uintptr) (interface{}, error) {
	return wrapCalendar(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
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

// Date returns a Time representing the shown year, month and the selected day.
//
// The returned date is in the local time zone.
func (self *Calendar) Date() *glib.DateTime {
	var _arg0 *C.GtkCalendar // out
	var _cret *C.GDateTime   // in

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_calendar_get_date(_arg0)
	runtime.KeepAlive(self)

	var _dateTime *glib.DateTime // out

	_dateTime = (*glib.DateTime)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_dateTime)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_date_time_unref((*C.GDateTime)(intern.C))
		},
	)

	return _dateTime
}

// DayIsMarked returns if the day of the calendar is already marked.
//
// The function takes the following parameters:
//
//    - day number between 1 and 31.
//
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

// ShowDayNames returns whether self is currently showing the names of the week
// days.
//
// This is the value of the gtk.Calendar:show-day-names property.
func (self *Calendar) ShowDayNames() bool {
	var _arg0 *C.GtkCalendar // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_calendar_get_show_day_names(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowHeading returns whether self is currently showing the heading.
//
// This is the value of the gtk.Calendar:show-heading property.
func (self *Calendar) ShowHeading() bool {
	var _arg0 *C.GtkCalendar // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_calendar_get_show_heading(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowWeekNumbers returns whether self is showing week numbers right now.
//
// This is the value of the gtk.Calendar:show-week-numbers property.
func (self *Calendar) ShowWeekNumbers() bool {
	var _arg0 *C.GtkCalendar // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_calendar_get_show_week_numbers(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MarkDay places a visual marker on a particular day.
//
// The function takes the following parameters:
//
//    - day number to mark between 1 and 31.
//
func (calendar *Calendar) MarkDay(day uint) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(calendar.Native()))
	_arg1 = C.guint(day)

	C.gtk_calendar_mark_day(_arg0, _arg1)
	runtime.KeepAlive(calendar)
	runtime.KeepAlive(day)
}

// SelectDay switches to date's year and month and select its day.
//
// The function takes the following parameters:
//
//    - date representing the day to select.
//
func (self *Calendar) SelectDay(date *glib.DateTime) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 *C.GDateTime   // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GDateTime)(gextras.StructNative(unsafe.Pointer(date)))

	C.gtk_calendar_select_day(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(date)
}

// SetShowDayNames sets whether the calendar shows day names.
//
// The function takes the following parameters:
//
//    - value: whether to show day names above the day numbers.
//
func (self *Calendar) SetShowDayNames(value bool) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(self.Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.gtk_calendar_set_show_day_names(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetShowHeading sets whether the calendar should show a heading.
//
// The heading contains the current year and month as well as buttons for
// changing both.
//
// The function takes the following parameters:
//
//    - value: whether to show the heading in the calendar.
//
func (self *Calendar) SetShowHeading(value bool) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(self.Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.gtk_calendar_set_show_heading(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetShowWeekNumbers sets whether week numbers are shown in the calendar.
//
// The function takes the following parameters:
//
//    - value: whether to show week numbers on the left of the days.
//
func (self *Calendar) SetShowWeekNumbers(value bool) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(self.Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.gtk_calendar_set_show_week_numbers(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// UnmarkDay removes the visual marker from a particular day.
//
// The function takes the following parameters:
//
//    - day number to unmark between 1 and 31.
//
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
func (calendar *Calendar) ConnectDaySelected(f func()) externglib.SignalHandle {
	return calendar.Connect("day-selected", f)
}

// ConnectNextMonth: emitted when the user switched to the next month.
func (calendar *Calendar) ConnectNextMonth(f func()) externglib.SignalHandle {
	return calendar.Connect("next-month", f)
}

// ConnectNextYear: emitted when user switched to the next year.
func (calendar *Calendar) ConnectNextYear(f func()) externglib.SignalHandle {
	return calendar.Connect("next-year", f)
}

// ConnectPrevMonth: emitted when the user switched to the previous month.
func (calendar *Calendar) ConnectPrevMonth(f func()) externglib.SignalHandle {
	return calendar.Connect("prev-month", f)
}

// ConnectPrevYear: emitted when user switched to the previous year.
func (calendar *Calendar) ConnectPrevYear(f func()) externglib.SignalHandle {
	return calendar.Connect("prev-year", f)
}
