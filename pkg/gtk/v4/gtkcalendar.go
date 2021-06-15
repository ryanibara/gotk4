// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_calendar_get_type()), F: marshalCalendar},
	})
}

// Calendar: `GtkCalendar` is a widget that displays a Gregorian calendar, one
// month at a time.
//
// !An example GtkCalendar (calendar.png)
//
// A `GtkCalendar` can be created with [ctor@Gtk.Calendar.new].
//
// The date that is currently displayed can be altered with
// [method@Gtk.Calendar.select_day].
//
// To place a visual marker on a particular day, use
// [method@Gtk.Calendar.mark_day] and to remove the marker,
// [method@Gtk.Calendar.unmark_day]. Alternative, all marks can be cleared with
// [method@Gtk.Calendar.clear_marks].
//
// The selected date can be retrieved from a `GtkCalendar` using
// [method@Gtk.Calendar.get_date].
//
// Users should be aware that, although the Gregorian calendar is the legal
// calendar in most countries, it was adopted progressively between 1582 and
// 1929. Display before these dates is likely to be historically incorrect.
//
//
// CSS nodes
//
// “` calendar.view ├── header │ ├── button │ ├── stack.month │ ├── button │ ├──
// button │ ├── label.year │ ╰── button ╰── grid ╰──
// label[.day-name][.week-number][.day-number][.other-month][.today] “`
//
// `GtkCalendar` has a main node with name calendar. It contains a subnode
// called header containing the widgets for switching between years and months.
//
// The grid subnode contains all day labels, including week numbers on the left
// (marked with the .week-number css class) and day names on top (marked with
// the .day-name css class).
//
// Day labels that belong to the previous or next month get the .other-month
// style class. The label of the current day get the .today style class.
//
// Marked day labels get the :selected state assigned.
type Calendar interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget

	// ClearMarks: remove all visual markers.
	ClearMarks()
	// DayIsMarked returns if the @day of the @calendar is already marked.
	DayIsMarked(day uint) bool
	// ShowDayNames returns whether @self is currently showing the names of the
	// week days.
	//
	// This is the value of the [property@Gtk.Calendar:show-day-names] property.
	ShowDayNames() bool
	// ShowHeading returns whether @self is currently showing the heading.
	//
	// This is the value of the [property@Gtk.Calendar:show-heading] property.
	ShowHeading() bool
	// ShowWeekNumbers returns whether @self is showing week numbers right now.
	//
	// This is the value of the [property@Gtk.Calendar:show-week-numbers]
	// property.
	ShowWeekNumbers() bool
	// MarkDay places a visual marker on a particular day.
	MarkDay(day uint)
	// SetShowDayNames sets whether the calendar shows day names.
	SetShowDayNames(value bool)
	// SetShowHeading sets whether the calendar should show a heading.
	//
	// The heading contains the current year and month as well as buttons for
	// changing both.
	SetShowHeading(value bool)
	// SetShowWeekNumbers sets whether week numbers are shown in the calendar.
	SetShowWeekNumbers(value bool)
	// UnmarkDay removes the visual marker from a particular day.
	UnmarkDay(day uint)
}

// calendar implements the Calendar class.
type calendar struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ Calendar = (*calendar)(nil)

// WrapCalendar wraps a GObject to the right type. It is
// primarily used internally.
func WrapCalendar(obj *externglib.Object) Calendar {
	return calendar{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalCalendar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCalendar(obj), nil
}

// NewCalendar constructs a class Calendar.
func NewCalendar() Calendar {
	var _cret C.GtkCalendar // in

	_cret = C.gtk_calendar_new()

	var _calendar Calendar // out

	_calendar = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Calendar)

	return _calendar
}

// ClearMarks: remove all visual markers.
func (c calendar) ClearMarks() {
	var _arg0 *C.GtkCalendar // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(c.Native()))

	C.gtk_calendar_clear_marks(_arg0)
}

// DayIsMarked returns if the @day of the @calendar is already marked.
func (c calendar) DayIsMarked(day uint) bool {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.guint        // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(c.Native()))
	_arg1 = (C.guint)(day)

	_cret = C.gtk_calendar_get_day_is_marked(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowDayNames returns whether @self is currently showing the names of the
// week days.
//
// This is the value of the [property@Gtk.Calendar:show-day-names] property.
func (s calendar) ShowDayNames() bool {
	var _arg0 *C.GtkCalendar // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_calendar_get_show_day_names(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowHeading returns whether @self is currently showing the heading.
//
// This is the value of the [property@Gtk.Calendar:show-heading] property.
func (s calendar) ShowHeading() bool {
	var _arg0 *C.GtkCalendar // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_calendar_get_show_heading(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowWeekNumbers returns whether @self is showing week numbers right now.
//
// This is the value of the [property@Gtk.Calendar:show-week-numbers]
// property.
func (s calendar) ShowWeekNumbers() bool {
	var _arg0 *C.GtkCalendar // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_calendar_get_show_week_numbers(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MarkDay places a visual marker on a particular day.
func (c calendar) MarkDay(day uint) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(c.Native()))
	_arg1 = (C.guint)(day)

	C.gtk_calendar_mark_day(_arg0, _arg1)
}

// SetShowDayNames sets whether the calendar shows day names.
func (s calendar) SetShowDayNames(value bool) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(s.Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.gtk_calendar_set_show_day_names(_arg0, _arg1)
}

// SetShowHeading sets whether the calendar should show a heading.
//
// The heading contains the current year and month as well as buttons for
// changing both.
func (s calendar) SetShowHeading(value bool) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(s.Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.gtk_calendar_set_show_heading(_arg0, _arg1)
}

// SetShowWeekNumbers sets whether week numbers are shown in the calendar.
func (s calendar) SetShowWeekNumbers(value bool) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(s.Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.gtk_calendar_set_show_week_numbers(_arg0, _arg1)
}

// UnmarkDay removes the visual marker from a particular day.
func (c calendar) UnmarkDay(day uint) {
	var _arg0 *C.GtkCalendar // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkCalendar)(unsafe.Pointer(c.Native()))
	_arg1 = (C.guint)(day)

	C.gtk_calendar_unmark_day(_arg0, _arg1)
}
