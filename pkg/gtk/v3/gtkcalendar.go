// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gchar* _gotk4_gtk3_CalendarDetailFunc(void*, guint, guint, guint, gpointer);
// extern void _gotk4_gtk3_CalendarClass_day_selected(void*);
// extern void _gotk4_gtk3_CalendarClass_day_selected_double_click(void*);
// extern void _gotk4_gtk3_CalendarClass_month_changed(void*);
// extern void _gotk4_gtk3_CalendarClass_next_month(void*);
// extern void _gotk4_gtk3_CalendarClass_next_year(void*);
// extern void _gotk4_gtk3_CalendarClass_prev_month(void*);
// extern void _gotk4_gtk3_CalendarClass_prev_year(void*);
// extern void _gotk4_gtk3_Calendar_ConnectDaySelected(gpointer, guintptr);
// extern void _gotk4_gtk3_Calendar_ConnectDaySelectedDoubleClick(gpointer, guintptr);
// extern void _gotk4_gtk3_Calendar_ConnectMonthChanged(gpointer, guintptr);
// extern void _gotk4_gtk3_Calendar_ConnectNextMonth(gpointer, guintptr);
// extern void _gotk4_gtk3_Calendar_ConnectNextYear(gpointer, guintptr);
// extern void _gotk4_gtk3_Calendar_ConnectPrevMonth(gpointer, guintptr);
// extern void _gotk4_gtk3_Calendar_ConnectPrevYear(gpointer, guintptr);
// extern void callbackDelete(gpointer);
import "C"

// glib.Type values for gtkcalendar.go.
var (
	GTypeCalendarDisplayOptions = coreglib.Type(C.gtk_calendar_display_options_get_type())
	GTypeCalendar               = coreglib.Type(C.gtk_calendar_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeCalendarDisplayOptions, F: marshalCalendarDisplayOptions},
		{T: GTypeCalendar, F: marshalCalendar},
	})
}

// CalendarDisplayOptions: these options can be used to influence the display
// and behaviour of a Calendar.
type CalendarDisplayOptions C.guint

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
	return CalendarDisplayOptions(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
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
type CalendarDetailFunc func(calendar *Calendar, year, month, day uint32) (utf8 string)

//export _gotk4_gtk3_CalendarDetailFunc
func _gotk4_gtk3_CalendarDetailFunc(arg1 *C.void, arg2 C.guint, arg3 C.guint, arg4 C.guint, arg5 C.gpointer) (cret *C.gchar) {
	var fn CalendarDetailFunc
	{
		v := gbox.Get(uintptr(arg5))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(CalendarDetailFunc)
	}

	var _calendar *Calendar // out
	var _year uint32        // out
	var _month uint32       // out
	var _day uint32         // out

	_calendar = wrapCalendar(coreglib.Take(unsafe.Pointer(arg1)))
	_year = uint32(arg2)
	_month = uint32(arg3)
	_day = uint32(arg4)

	utf8 := fn(_calendar, _year, _month, _day)

	if utf8 != "" {
		cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
	}

	return cret
}

// CalendarOverrider contains methods that are overridable.
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
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Calendar)(nil)
)

func classInitCalendarrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkCalendarClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkCalendarClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ DaySelected() }); ok {
		pclass.day_selected = (*[0]byte)(C._gotk4_gtk3_CalendarClass_day_selected)
	}

	if _, ok := goval.(interface{ DaySelectedDoubleClick() }); ok {
		pclass.day_selected_double_click = (*[0]byte)(C._gotk4_gtk3_CalendarClass_day_selected_double_click)
	}

	if _, ok := goval.(interface{ MonthChanged() }); ok {
		pclass.month_changed = (*[0]byte)(C._gotk4_gtk3_CalendarClass_month_changed)
	}

	if _, ok := goval.(interface{ NextMonth() }); ok {
		pclass.next_month = (*[0]byte)(C._gotk4_gtk3_CalendarClass_next_month)
	}

	if _, ok := goval.(interface{ NextYear() }); ok {
		pclass.next_year = (*[0]byte)(C._gotk4_gtk3_CalendarClass_next_year)
	}

	if _, ok := goval.(interface{ PrevMonth() }); ok {
		pclass.prev_month = (*[0]byte)(C._gotk4_gtk3_CalendarClass_prev_month)
	}

	if _, ok := goval.(interface{ PrevYear() }); ok {
		pclass.prev_year = (*[0]byte)(C._gotk4_gtk3_CalendarClass_prev_year)
	}
}

//export _gotk4_gtk3_CalendarClass_day_selected
func _gotk4_gtk3_CalendarClass_day_selected(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ DaySelected() })

	iface.DaySelected()
}

//export _gotk4_gtk3_CalendarClass_day_selected_double_click
func _gotk4_gtk3_CalendarClass_day_selected_double_click(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ DaySelectedDoubleClick() })

	iface.DaySelectedDoubleClick()
}

//export _gotk4_gtk3_CalendarClass_month_changed
func _gotk4_gtk3_CalendarClass_month_changed(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ MonthChanged() })

	iface.MonthChanged()
}

//export _gotk4_gtk3_CalendarClass_next_month
func _gotk4_gtk3_CalendarClass_next_month(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ NextMonth() })

	iface.NextMonth()
}

//export _gotk4_gtk3_CalendarClass_next_year
func _gotk4_gtk3_CalendarClass_next_year(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ NextYear() })

	iface.NextYear()
}

//export _gotk4_gtk3_CalendarClass_prev_month
func _gotk4_gtk3_CalendarClass_prev_month(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ PrevMonth() })

	iface.PrevMonth()
}

//export _gotk4_gtk3_CalendarClass_prev_year
func _gotk4_gtk3_CalendarClass_prev_year(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ PrevYear() })

	iface.PrevYear()
}

func wrapCalendar(obj *coreglib.Object) *Calendar {
	return &Calendar{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
		},
	}
}

func marshalCalendar(p uintptr) (interface{}, error) {
	return wrapCalendar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Calendar_ConnectDaySelected
func _gotk4_gtk3_Calendar_ConnectDaySelected(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectDaySelected is emitted when the user selects a day.
func (calendar *Calendar) ConnectDaySelected(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(calendar, "day-selected", false, unsafe.Pointer(C._gotk4_gtk3_Calendar_ConnectDaySelected), f)
}

//export _gotk4_gtk3_Calendar_ConnectDaySelectedDoubleClick
func _gotk4_gtk3_Calendar_ConnectDaySelectedDoubleClick(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectDaySelectedDoubleClick is emitted when the user double-clicks a day.
func (calendar *Calendar) ConnectDaySelectedDoubleClick(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(calendar, "day-selected-double-click", false, unsafe.Pointer(C._gotk4_gtk3_Calendar_ConnectDaySelectedDoubleClick), f)
}

//export _gotk4_gtk3_Calendar_ConnectMonthChanged
func _gotk4_gtk3_Calendar_ConnectMonthChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectMonthChanged is emitted when the user clicks a button to change the
// selected month on a calendar.
func (calendar *Calendar) ConnectMonthChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(calendar, "month-changed", false, unsafe.Pointer(C._gotk4_gtk3_Calendar_ConnectMonthChanged), f)
}

//export _gotk4_gtk3_Calendar_ConnectNextMonth
func _gotk4_gtk3_Calendar_ConnectNextMonth(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectNextMonth is emitted when the user switched to the next month.
func (calendar *Calendar) ConnectNextMonth(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(calendar, "next-month", false, unsafe.Pointer(C._gotk4_gtk3_Calendar_ConnectNextMonth), f)
}

//export _gotk4_gtk3_Calendar_ConnectNextYear
func _gotk4_gtk3_Calendar_ConnectNextYear(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectNextYear is emitted when user switched to the next year.
func (calendar *Calendar) ConnectNextYear(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(calendar, "next-year", false, unsafe.Pointer(C._gotk4_gtk3_Calendar_ConnectNextYear), f)
}

//export _gotk4_gtk3_Calendar_ConnectPrevMonth
func _gotk4_gtk3_Calendar_ConnectPrevMonth(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectPrevMonth is emitted when the user switched to the previous month.
func (calendar *Calendar) ConnectPrevMonth(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(calendar, "prev-month", false, unsafe.Pointer(C._gotk4_gtk3_Calendar_ConnectPrevMonth), f)
}

//export _gotk4_gtk3_Calendar_ConnectPrevYear
func _gotk4_gtk3_Calendar_ConnectPrevYear(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectPrevYear is emitted when user switched to the previous year.
func (calendar *Calendar) ConnectPrevYear(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(calendar, "prev-year", false, unsafe.Pointer(C._gotk4_gtk3_Calendar_ConnectPrevYear), f)
}

// NewCalendar creates a new calendar, with the current date being selected.
//
// The function returns the following values:
//
//    - calendar: newly Calendar widget.
//
func NewCalendar() *Calendar {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "Calendar").InvokeMethod("new_Calendar", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _calendar *Calendar // out

	_calendar = wrapCalendar(coreglib.Take(unsafe.Pointer(_cret)))

	return _calendar
}

// ClearMarks: remove all visual markers.
func (calendar *Calendar) ClearMarks() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(calendar).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "Calendar").InvokeMethod("clear_marks", _args[:], nil)

	runtime.KeepAlive(calendar)
}

// Date obtains the selected date from a Calendar.
//
// The function returns the following values:
//
//    - year (optional): location to store the year as a decimal number (e.g.
//      2011), or NULL.
//    - month (optional): location to store the month number (between 0 and 11),
//      or NULL.
//    - day (optional): location to store the day number (between 1 and 31), or
//      NULL.
//
func (calendar *Calendar) Date() (year uint32, month uint32, day uint32) {
	var _args [1]girepository.Argument
	var _outs [3]girepository.Argument
	var _arg0 *C.void // out
	var _out0 *C.void // in
	var _out1 *C.void // in
	var _out2 *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(calendar).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "Calendar").InvokeMethod("get_date", _args[:], _outs[:])

	runtime.KeepAlive(calendar)

	var _year uint32  // out
	var _month uint32 // out
	var _day uint32   // out
	_out0 = *(**C.void)(unsafe.Pointer(&_outs[0]))
	_out1 = *(**C.void)(unsafe.Pointer(&_outs[1]))
	_out2 = *(**C.void)(unsafe.Pointer(&_outs[2]))

	if _out0 != nil {
		_year = *(*uint32)(unsafe.Pointer(_out0))
	}
	if _out1 != nil {
		_month = *(*uint32)(unsafe.Pointer(_out1))
	}
	if _out2 != nil {
		_day = *(*uint32)(unsafe.Pointer(_out2))
	}

	return _year, _month, _day
}

// DayIsMarked returns if the day of the calendar is already marked.
//
// The function takes the following parameters:
//
//    - day number between 1 and 31.
//
// The function returns the following values:
//
//    - ok: whether the day is marked.
//
func (calendar *Calendar) DayIsMarked(day uint32) bool {
	var _args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.guint    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(calendar).Native()))
	_arg1 = C.guint(day)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.guint)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Calendar").InvokeMethod("get_day_is_marked", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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
//
// The function returns the following values:
//
//    - gint: height of detail cells, in rows.
//
func (calendar *Calendar) DetailHeightRows() int32 {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(calendar).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Calendar").InvokeMethod("get_detail_height_rows", _args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(calendar)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// DetailWidthChars queries the width of detail cells, in characters. See
// Calendar:detail-width-chars.
//
// The function returns the following values:
//
//    - gint: width of detail cells, in characters.
//
func (calendar *Calendar) DetailWidthChars() int32 {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(calendar).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Calendar").InvokeMethod("get_detail_width_chars", _args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(calendar)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// MarkDay places a visual marker on a particular day.
//
// The function takes the following parameters:
//
//    - day number to mark between 1 and 31.
//
func (calendar *Calendar) MarkDay(day uint32) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(calendar).Native()))
	_arg1 = C.guint(day)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.guint)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Calendar").InvokeMethod("mark_day", _args[:], nil)

	runtime.KeepAlive(calendar)
	runtime.KeepAlive(day)
}

// SelectDay selects a day from the current month.
//
// The function takes the following parameters:
//
//    - day number between 1 and 31, or 0 to unselect the currently selected day.
//
func (calendar *Calendar) SelectDay(day uint32) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(calendar).Native()))
	_arg1 = C.guint(day)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.guint)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Calendar").InvokeMethod("select_day", _args[:], nil)

	runtime.KeepAlive(calendar)
	runtime.KeepAlive(day)
}

// SelectMonth shifts the calendar to a different month.
//
// The function takes the following parameters:
//
//    - month number between 0 and 11.
//    - year the month is in.
//
func (calendar *Calendar) SelectMonth(month, year uint32) {
	var _args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out
	var _arg2 C.guint // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(calendar).Native()))
	_arg1 = C.guint(month)
	_arg2 = C.guint(year)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.guint)(unsafe.Pointer(&_args[1])) = _arg1
	*(*C.guint)(unsafe.Pointer(&_args[2])) = _arg2

	girepository.MustFind("Gtk", "Calendar").InvokeMethod("select_month", _args[:], nil)

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
//
// The function takes the following parameters:
//
//    - fn: function providing details for each day.
//
func (calendar *Calendar) SetDetailFunc(fn CalendarDetailFunc) {
	var _args [4]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gpointer // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(calendar).Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_CalendarDetailFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Calendar").InvokeMethod("set_detail_func", _args[:], nil)

	runtime.KeepAlive(calendar)
	runtime.KeepAlive(fn)
}

// SetDetailHeightRows updates the height of detail cells. See
// Calendar:detail-height-rows.
//
// The function takes the following parameters:
//
//    - rows: detail height in rows.
//
func (calendar *Calendar) SetDetailHeightRows(rows int32) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(calendar).Native()))
	_arg1 = C.gint(rows)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gint)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Calendar").InvokeMethod("set_detail_height_rows", _args[:], nil)

	runtime.KeepAlive(calendar)
	runtime.KeepAlive(rows)
}

// SetDetailWidthChars updates the width of detail cells. See
// Calendar:detail-width-chars.
//
// The function takes the following parameters:
//
//    - chars: detail width in characters.
//
func (calendar *Calendar) SetDetailWidthChars(chars int32) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(calendar).Native()))
	_arg1 = C.gint(chars)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gint)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Calendar").InvokeMethod("set_detail_width_chars", _args[:], nil)

	runtime.KeepAlive(calendar)
	runtime.KeepAlive(chars)
}

// UnmarkDay removes the visual marker from a particular day.
//
// The function takes the following parameters:
//
//    - day number to unmark between 1 and 31.
//
func (calendar *Calendar) UnmarkDay(day uint32) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(calendar).Native()))
	_arg1 = C.guint(day)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.guint)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Calendar").InvokeMethod("unmark_day", _args[:], nil)

	runtime.KeepAlive(calendar)
	runtime.KeepAlive(day)
}
