// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern gboolean _gotk4_gtk4_RangeClass_change_value(GtkRange*, GtkScrollType, double);
// extern gboolean _gotk4_gtk4_Range_ConnectChangeValue(gpointer, GtkScrollType, gdouble, guintptr);
// extern void _gotk4_gtk4_RangeClass_adjust_bounds(GtkRange*, double);
// extern void _gotk4_gtk4_RangeClass_get_range_border(GtkRange*, GtkBorder*);
// extern void _gotk4_gtk4_RangeClass_move_slider(GtkRange*, GtkScrollType);
// extern void _gotk4_gtk4_RangeClass_value_changed(GtkRange*);
// extern void _gotk4_gtk4_Range_ConnectAdjustBounds(gpointer, gdouble, guintptr);
// extern void _gotk4_gtk4_Range_ConnectMoveSlider(gpointer, GtkScrollType, guintptr);
// extern void _gotk4_gtk4_Range_ConnectValueChanged(gpointer, guintptr);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_range_get_type()), F: marshalRanger},
	})
}

// RangeOverrider contains methods that are overridable.
type RangeOverrider interface {
	// The function takes the following parameters:
	//
	AdjustBounds(newValue float64)
	// The function takes the following parameters:
	//
	//    - scroll
	//    - newValue
	//
	// The function returns the following values:
	//
	ChangeValue(scroll ScrollType, newValue float64) bool
	// The function takes the following parameters:
	//
	RangeBorder(border_ *Border)
	// The function takes the following parameters:
	//
	MoveSlider(scroll ScrollType)
	ValueChanged()
}

// Range: GtkRange is the common base class for widgets which visualize an
// adjustment.
//
// Widgets that are derived from GtkRange include gtk.Scale and gtk.Scrollbar.
//
// Apart from signals for monitoring the parameters of the adjustment, GtkRange
// provides properties and methods for setting a “fill level” on range widgets.
// See gtk.Range.SetFillLevel().
type Range struct {
	_ [0]func() // equal guard
	Widget

	*externglib.Object
	Orientable
}

var (
	_ Widgetter           = (*Range)(nil)
	_ externglib.Objector = (*Range)(nil)
)

func classInitRanger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkRangeClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkRangeClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ AdjustBounds(newValue float64) }); ok {
		pclass.adjust_bounds = (*[0]byte)(C._gotk4_gtk4_RangeClass_adjust_bounds)
	}

	if _, ok := goval.(interface {
		ChangeValue(scroll ScrollType, newValue float64) bool
	}); ok {
		pclass.change_value = (*[0]byte)(C._gotk4_gtk4_RangeClass_change_value)
	}

	if _, ok := goval.(interface{ RangeBorder(border_ *Border) }); ok {
		pclass.get_range_border = (*[0]byte)(C._gotk4_gtk4_RangeClass_get_range_border)
	}

	if _, ok := goval.(interface{ MoveSlider(scroll ScrollType) }); ok {
		pclass.move_slider = (*[0]byte)(C._gotk4_gtk4_RangeClass_move_slider)
	}

	if _, ok := goval.(interface{ ValueChanged() }); ok {
		pclass.value_changed = (*[0]byte)(C._gotk4_gtk4_RangeClass_value_changed)
	}
}

//export _gotk4_gtk4_RangeClass_adjust_bounds
func _gotk4_gtk4_RangeClass_adjust_bounds(arg0 *C.GtkRange, arg1 C.double) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ AdjustBounds(newValue float64) })

	var _newValue float64 // out

	_newValue = float64(arg1)

	iface.AdjustBounds(_newValue)
}

//export _gotk4_gtk4_RangeClass_change_value
func _gotk4_gtk4_RangeClass_change_value(arg0 *C.GtkRange, arg1 C.GtkScrollType, arg2 C.double) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		ChangeValue(scroll ScrollType, newValue float64) bool
	})

	var _scroll ScrollType // out
	var _newValue float64  // out

	_scroll = ScrollType(arg1)
	_newValue = float64(arg2)

	ok := iface.ChangeValue(_scroll, _newValue)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_RangeClass_get_range_border
func _gotk4_gtk4_RangeClass_get_range_border(arg0 *C.GtkRange, arg1 *C.GtkBorder) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ RangeBorder(border_ *Border) })

	var _border_ *Border // out

	_border_ = (*Border)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	iface.RangeBorder(_border_)
}

//export _gotk4_gtk4_RangeClass_move_slider
func _gotk4_gtk4_RangeClass_move_slider(arg0 *C.GtkRange, arg1 C.GtkScrollType) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ MoveSlider(scroll ScrollType) })

	var _scroll ScrollType // out

	_scroll = ScrollType(arg1)

	iface.MoveSlider(_scroll)
}

//export _gotk4_gtk4_RangeClass_value_changed
func _gotk4_gtk4_RangeClass_value_changed(arg0 *C.GtkRange) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ValueChanged() })

	iface.ValueChanged()
}

func wrapRange(obj *externglib.Object) *Range {
	return &Range{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalRanger(p uintptr) (interface{}, error) {
	return wrapRange(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_Range_ConnectAdjustBounds
func _gotk4_gtk4_Range_ConnectAdjustBounds(arg0 C.gpointer, arg1 C.gdouble, arg2 C.guintptr) {
	var f func(value float64)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(value float64))
	}

	var _value float64 // out

	_value = float64(arg1)

	f(_value)
}

// ConnectAdjustBounds: emitted before clamping a value, to give the application
// a chance to adjust the bounds.
func (_range *Range) ConnectAdjustBounds(f func(value float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(_range, "adjust-bounds", false, unsafe.Pointer(C._gotk4_gtk4_Range_ConnectAdjustBounds), f)
}

//export _gotk4_gtk4_Range_ConnectChangeValue
func _gotk4_gtk4_Range_ConnectChangeValue(arg0 C.gpointer, arg1 C.GtkScrollType, arg2 C.gdouble, arg3 C.guintptr) (cret C.gboolean) {
	var f func(scroll ScrollType, value float64) (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(scroll ScrollType, value float64) (ok bool))
	}

	var _scroll ScrollType // out
	var _value float64     // out

	_scroll = ScrollType(arg1)
	_value = float64(arg2)

	ok := f(_scroll, _value)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectChangeValue: emitted when a scroll action is performed on a range.
//
// It allows an application to determine the type of scroll event that occurred
// and the resultant new value. The application can handle the event itself and
// return TRUE to prevent further processing. Or, by returning FALSE, it can
// pass the event to other handlers until the default GTK handler is reached.
//
// The value parameter is unrounded. An application that overrides the
// ::change-value signal is responsible for clamping the value to the desired
// number of decimal digits; the default GTK handler clamps the value based on
// gtk.Range:round-digits.
func (_range *Range) ConnectChangeValue(f func(scroll ScrollType, value float64) (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(_range, "change-value", false, unsafe.Pointer(C._gotk4_gtk4_Range_ConnectChangeValue), f)
}

//export _gotk4_gtk4_Range_ConnectMoveSlider
func _gotk4_gtk4_Range_ConnectMoveSlider(arg0 C.gpointer, arg1 C.GtkScrollType, arg2 C.guintptr) {
	var f func(step ScrollType)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(step ScrollType))
	}

	var _step ScrollType // out

	_step = ScrollType(arg1)

	f(_step)
}

// ConnectMoveSlider: virtual function that moves the slider.
//
// Used for keybindings.
func (_range *Range) ConnectMoveSlider(f func(step ScrollType)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(_range, "move-slider", false, unsafe.Pointer(C._gotk4_gtk4_Range_ConnectMoveSlider), f)
}

//export _gotk4_gtk4_Range_ConnectValueChanged
func _gotk4_gtk4_Range_ConnectValueChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectValueChanged: emitted when the range value changes.
func (_range *Range) ConnectValueChanged(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(_range, "value-changed", false, unsafe.Pointer(C._gotk4_gtk4_Range_ConnectValueChanged), f)
}

// Adjustment: get the adjustment which is the “model” object for GtkRange.
//
// The function returns the following values:
//
//    - adjustment: GtkAdjustment.
//
func (_range *Range) Adjustment() *Adjustment {
	var _arg0 *C.GtkRange      // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_adjustment(_arg0)
	runtime.KeepAlive(_range)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// FillLevel gets the current position of the fill level indicator.
//
// The function returns the following values:
//
//    - gdouble: current fill level.
//
func (_range *Range) FillLevel() float64 {
	var _arg0 *C.GtkRange // out
	var _cret C.double    // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_fill_level(_arg0)
	runtime.KeepAlive(_range)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Flippable gets whether the GtkRange respects text direction.
//
// See gtk.Range.SetFlippable().
//
// The function returns the following values:
//
//    - ok: TRUE if the range is flippable.
//
func (_range *Range) Flippable() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_flippable(_arg0)
	runtime.KeepAlive(_range)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Inverted gets whether the range is inverted.
//
// See gtk.Range.SetInverted().
//
// The function returns the following values:
//
//    - ok: TRUE if the range is inverted.
//
func (_range *Range) Inverted() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_inverted(_arg0)
	runtime.KeepAlive(_range)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RangeRect: this function returns the area that contains the range’s trough,
// in coordinates relative to range's origin.
//
// This function is useful mainly for GtkRange subclasses.
//
// The function returns the following values:
//
//    - rangeRect: return location for the range rectangle.
//
func (_range *Range) RangeRect() *gdk.Rectangle {
	var _arg0 *C.GtkRange    // out
	var _arg1 C.GdkRectangle // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	C.gtk_range_get_range_rect(_arg0, &_arg1)
	runtime.KeepAlive(_range)

	var _rangeRect *gdk.Rectangle // out

	_rangeRect = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _rangeRect
}

// RestrictToFillLevel gets whether the range is restricted to the fill level.
//
// The function returns the following values:
//
//    - ok: TRUE if range is restricted to the fill level.
//
func (_range *Range) RestrictToFillLevel() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_restrict_to_fill_level(_arg0)
	runtime.KeepAlive(_range)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RoundDigits gets the number of digits to round the value to when it changes.
//
// See gtk.Range::change-value.
//
// The function returns the following values:
//
//    - gint: number of digits to round to.
//
func (_range *Range) RoundDigits() int {
	var _arg0 *C.GtkRange // out
	var _cret C.int       // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_round_digits(_arg0)
	runtime.KeepAlive(_range)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ShowFillLevel gets whether the range displays the fill level graphically.
//
// The function returns the following values:
//
//    - ok: TRUE if range shows the fill level.
//
func (_range *Range) ShowFillLevel() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_show_fill_level(_arg0)
	runtime.KeepAlive(_range)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SliderRange: this function returns sliders range along the long dimension, in
// widget->window coordinates.
//
// This function is useful mainly for GtkRange subclasses.
//
// The function returns the following values:
//
//    - sliderStart (optional): return location for the slider's start, or NULL.
//    - sliderEnd (optional): return location for the slider's end, or NULL.
//
func (_range *Range) SliderRange() (sliderStart int, sliderEnd int) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.int       // in
	var _arg2 C.int       // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	C.gtk_range_get_slider_range(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(_range)

	var _sliderStart int // out
	var _sliderEnd int   // out

	_sliderStart = int(_arg1)
	_sliderEnd = int(_arg2)

	return _sliderStart, _sliderEnd
}

// SliderSizeFixed: this function is useful mainly for GtkRange subclasses.
//
// See gtk.Range.SetSliderSizeFixed().
//
// The function returns the following values:
//
//    - ok: whether the range’s slider has a fixed size.
//
func (_range *Range) SliderSizeFixed() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_slider_size_fixed(_arg0)
	runtime.KeepAlive(_range)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Value gets the current value of the range.
//
// The function returns the following values:
//
//    - gdouble: current value of the range.
//
func (_range *Range) Value() float64 {
	var _arg0 *C.GtkRange // out
	var _cret C.double    // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_value(_arg0)
	runtime.KeepAlive(_range)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetAdjustment sets the adjustment to be used as the “model” object for the
// GtkRange
//
// The adjustment indicates the current range value, the minimum and maximum
// range values, the step/page increments used for keybindings and scrolling,
// and the page size.
//
// The page size is normally 0 for GtkScale and nonzero for GtkScrollbar, and
// indicates the size of the visible area of the widget being scrolled. The page
// size affects the size of the scrollbar slider.
//
// The function takes the following parameters:
//
//    - adjustment: GtkAdjustment.
//
func (_range *Range) SetAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkRange      // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_range_set_adjustment(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(adjustment)
}

// SetFillLevel: set the new position of the fill level indicator.
//
// The “fill level” is probably best described by its most prominent use case,
// which is an indicator for the amount of pre-buffering in a streaming media
// player. In that use case, the value of the range would indicate the current
// play position, and the fill level would be the position up to which the
// file/stream has been downloaded.
//
// This amount of prebuffering can be displayed on the range’s trough and is
// themeable separately from the trough. To enable fill level display, use
// gtk.Range.SetShowFillLevel(). The range defaults to not showing the fill
// level.
//
// Additionally, it’s possible to restrict the range’s slider position to values
// which are smaller than the fill level. This is controlled by
// gtk.Range.SetRestrictToFillLevel() and is by default enabled.
//
// The function takes the following parameters:
//
//    - fillLevel: new position of the fill level indicator.
//
func (_range *Range) SetFillLevel(fillLevel float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.double    // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.double(fillLevel)

	C.gtk_range_set_fill_level(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(fillLevel)
}

// SetFlippable sets whether the GtkRange respects text direction.
//
// If a range is flippable, it will switch its direction if it is horizontal and
// its direction is GTK_TEXT_DIR_RTL.
//
// See gtk.Widget.GetDirection().
//
// The function takes the following parameters:
//
//    - flippable: TRUE to make the range flippable.
//
func (_range *Range) SetFlippable(flippable bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	if flippable {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_flippable(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(flippable)
}

// SetIncrements sets the step and page sizes for the range.
//
// The step size is used when the user clicks the GtkScrollbar arrows or moves a
// GtkScale via arrow keys. The page size is used for example when moving via
// Page Up or Page Down keys.
//
// The function takes the following parameters:
//
//    - step size.
//    - page size.
//
func (_range *Range) SetIncrements(step, page float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.double    // out
	var _arg2 C.double    // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.double(step)
	_arg2 = C.double(page)

	C.gtk_range_set_increments(_arg0, _arg1, _arg2)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(step)
	runtime.KeepAlive(page)
}

// SetInverted sets whether to invert the range.
//
// Ranges normally move from lower to higher values as the slider moves from top
// to bottom or left to right. Inverted ranges have higher values at the top or
// on the right rather than on the bottom or left.
//
// The function takes the following parameters:
//
//    - setting: TRUE to invert the range.
//
func (_range *Range) SetInverted(setting bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_inverted(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(setting)
}

// SetRange sets the allowable values in the GtkRange.
//
// The range value is clamped to be between min and max. (If the range has a
// non-zero page size, it is clamped between min and max - page-size.).
//
// The function takes the following parameters:
//
//    - min: minimum range value.
//    - max: maximum range value.
//
func (_range *Range) SetRange(min, max float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.double    // out
	var _arg2 C.double    // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.double(min)
	_arg2 = C.double(max)

	C.gtk_range_set_range(_arg0, _arg1, _arg2)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
}

// SetRestrictToFillLevel sets whether the slider is restricted to the fill
// level.
//
// See gtk.Range.SetFillLevel() for a general description of the fill level
// concept.
//
// The function takes the following parameters:
//
//    - restrictToFillLevel: whether the fill level restricts slider movement.
//
func (_range *Range) SetRestrictToFillLevel(restrictToFillLevel bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	if restrictToFillLevel {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_restrict_to_fill_level(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(restrictToFillLevel)
}

// SetRoundDigits sets the number of digits to round the value to when it
// changes.
//
// See gtk.Range::change-value.
//
// The function takes the following parameters:
//
//    - roundDigits: precision in digits, or -1.
//
func (_range *Range) SetRoundDigits(roundDigits int) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.int(roundDigits)

	C.gtk_range_set_round_digits(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(roundDigits)
}

// SetShowFillLevel sets whether a graphical fill level is show on the trough.
//
// See gtk.Range.SetFillLevel() for a general description of the fill level
// concept.
//
// The function takes the following parameters:
//
//    - showFillLevel: whether a fill level indicator graphics is shown.
//
func (_range *Range) SetShowFillLevel(showFillLevel bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	if showFillLevel {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_show_fill_level(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(showFillLevel)
}

// SetSliderSizeFixed sets whether the range’s slider has a fixed size, or a
// size that depends on its adjustment’s page size.
//
// This function is useful mainly for GtkRange subclasses.
//
// The function takes the following parameters:
//
//    - sizeFixed: TRUE to make the slider size constant.
//
func (_range *Range) SetSliderSizeFixed(sizeFixed bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	if sizeFixed {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_slider_size_fixed(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(sizeFixed)
}

// SetValue sets the current value of the range.
//
// If the value is outside the minimum or maximum range values, it will be
// clamped to fit inside them. The range emits the gtk.Range::value-changed
// signal if the value changes.
//
// The function takes the following parameters:
//
//    - value: new value of the range.
//
func (_range *Range) SetValue(value float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.double    // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.double(value)

	C.gtk_range_set_value(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(value)
}
