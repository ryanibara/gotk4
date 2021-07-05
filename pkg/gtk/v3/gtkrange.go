// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_range_get_type()), F: marshalRange},
	})
}

// Range is the common base class for widgets which visualize an adjustment, e.g
// Scale or Scrollbar.
//
// Apart from signals for monitoring the parameters of the adjustment, Range
// provides properties and methods for influencing the sensitivity of the
// “steppers”. It also provides properties and methods for setting a “fill
// level” on range widgets. See gtk_range_set_fill_level().
type Range interface {
	Widget

	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsOrientable casts the class to the Orientable interface.
	AsOrientable() Orientable

	// Adjustment: get the Adjustment which is the “model” object for Range. See
	// gtk_range_set_adjustment() for details. The return value does not have a
	// reference added, so should not be unreferenced.
	Adjustment() Adjustment
	// FillLevel gets the current position of the fill level indicator.
	FillLevel() float64
	// Flippable gets the value set by gtk_range_set_flippable().
	Flippable() bool
	// Inverted gets the value set by gtk_range_set_inverted().
	Inverted() bool
	// LowerStepperSensitivity gets the sensitivity policy for the stepper that
	// points to the 'lower' end of the GtkRange’s adjustment.
	LowerStepperSensitivity() SensitivityType
	// MinSliderSize: this function is useful mainly for Range subclasses.
	//
	// See gtk_range_set_min_slider_size().
	//
	// Deprecated: since version 3.20.
	MinSliderSize() int
	// RangeRect: this function returns the area that contains the range’s
	// trough and its steppers, in widget->window coordinates.
	//
	// This function is useful mainly for Range subclasses.
	RangeRect() gdk.Rectangle
	// RestrictToFillLevel gets whether the range is restricted to the fill
	// level.
	RestrictToFillLevel() bool
	// RoundDigits gets the number of digits to round the value to when it
	// changes. See Range::change-value.
	RoundDigits() int
	// ShowFillLevel gets whether the range displays the fill level graphically.
	ShowFillLevel() bool
	// SliderRange: this function returns sliders range along the long
	// dimension, in widget->window coordinates.
	//
	// This function is useful mainly for Range subclasses.
	SliderRange() (sliderStart int, sliderEnd int)
	// SliderSizeFixed: this function is useful mainly for Range subclasses.
	//
	// See gtk_range_set_slider_size_fixed().
	SliderSizeFixed() bool
	// UpperStepperSensitivity gets the sensitivity policy for the stepper that
	// points to the 'upper' end of the GtkRange’s adjustment.
	UpperStepperSensitivity() SensitivityType
	// Value gets the current value of the range.
	Value() float64
	// SetAdjustmentRange sets the adjustment to be used as the “model” object
	// for this range widget. The adjustment indicates the current range value,
	// the minimum and maximum range values, the step/page increments used for
	// keybindings and scrolling, and the page size. The page size is normally 0
	// for Scale and nonzero for Scrollbar, and indicates the size of the
	// visible area of the widget being scrolled. The page size affects the size
	// of the scrollbar slider.
	SetAdjustmentRange(adjustment Adjustment)
	// SetFillLevelRange: set the new position of the fill level indicator.
	//
	// The “fill level” is probably best described by its most prominent use
	// case, which is an indicator for the amount of pre-buffering in a
	// streaming media player. In that use case, the value of the range would
	// indicate the current play position, and the fill level would be the
	// position up to which the file/stream has been downloaded.
	//
	// This amount of prebuffering can be displayed on the range’s trough and is
	// themeable separately from the trough. To enable fill level display, use
	// gtk_range_set_show_fill_level(). The range defaults to not showing the
	// fill level.
	//
	// Additionally, it’s possible to restrict the range’s slider position to
	// values which are smaller than the fill level. This is controller by
	// gtk_range_set_restrict_to_fill_level() and is by default enabled.
	SetFillLevelRange(fillLevel float64)
	// SetFlippableRange: if a range is flippable, it will switch its direction
	// if it is horizontal and its direction is GTK_TEXT_DIR_RTL.
	//
	// See gtk_widget_get_direction().
	SetFlippableRange(flippable bool)
	// SetIncrementsRange sets the step and page sizes for the range. The step
	// size is used when the user clicks the Scrollbar arrows or moves Scale via
	// arrow keys. The page size is used for example when moving via Page Up or
	// Page Down keys.
	SetIncrementsRange(step float64, page float64)
	// SetInvertedRange ranges normally move from lower to higher values as the
	// slider moves from top to bottom or left to right. Inverted ranges have
	// higher values at the top or on the right rather than on the bottom or
	// left.
	SetInvertedRange(setting bool)
	// SetLowerStepperSensitivityRange sets the sensitivity policy for the
	// stepper that points to the 'lower' end of the GtkRange’s adjustment.
	SetLowerStepperSensitivityRange(sensitivity SensitivityType)
	// SetMinSliderSizeRange sets the minimum size of the range’s slider.
	//
	// This function is useful mainly for Range subclasses.
	//
	// Deprecated: since version 3.20.
	SetMinSliderSizeRange(minSize int)
	// SetRangeRange sets the allowable values in the Range, and clamps the
	// range value to be between @min and @max. (If the range has a non-zero
	// page size, it is clamped between @min and @max - page-size.)
	SetRangeRange(min float64, max float64)
	// SetRestrictToFillLevelRange sets whether the slider is restricted to the
	// fill level. See gtk_range_set_fill_level() for a general description of
	// the fill level concept.
	SetRestrictToFillLevelRange(restrictToFillLevel bool)
	// SetRoundDigitsRange sets the number of digits to round the value to when
	// it changes. See Range::change-value.
	SetRoundDigitsRange(roundDigits int)
	// SetShowFillLevelRange sets whether a graphical fill level is show on the
	// trough. See gtk_range_set_fill_level() for a general description of the
	// fill level concept.
	SetShowFillLevelRange(showFillLevel bool)
	// SetSliderSizeFixedRange sets whether the range’s slider has a fixed size,
	// or a size that depends on its adjustment’s page size.
	//
	// This function is useful mainly for Range subclasses.
	SetSliderSizeFixedRange(sizeFixed bool)
	// SetUpperStepperSensitivityRange sets the sensitivity policy for the
	// stepper that points to the 'upper' end of the GtkRange’s adjustment.
	SetUpperStepperSensitivityRange(sensitivity SensitivityType)
	// SetValueRange sets the current value of the range; if the value is
	// outside the minimum or maximum range values, it will be clamped to fit
	// inside them. The range emits the Range::value-changed signal if the value
	// changes.
	SetValueRange(value float64)
}

// _range implements the Range class.
type _range struct {
	Widget
}

// WrapRange wraps a GObject to the right type. It is
// primarily used internally.
func WrapRange(obj *externglib.Object) Range {
	return _range{
		Widget: WrapWidget(obj),
	}
}

func marshalRange(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRange(obj), nil
}

func (r _range) Adjustment() Adjustment {
	var _arg0 *C.GtkRange      // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_adjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Adjustment)

	return _adjustment
}

func (r _range) FillLevel() float64 {
	var _arg0 *C.GtkRange // out
	var _cret C.gdouble   // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_fill_level(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (r _range) Flippable() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_flippable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (r _range) Inverted() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_inverted(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (r _range) LowerStepperSensitivity() SensitivityType {
	var _arg0 *C.GtkRange          // out
	var _cret C.GtkSensitivityType // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_lower_stepper_sensitivity(_arg0)

	var _sensitivityType SensitivityType // out

	_sensitivityType = SensitivityType(_cret)

	return _sensitivityType
}

func (r _range) MinSliderSize() int {
	var _arg0 *C.GtkRange // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_min_slider_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (r _range) RangeRect() gdk.Rectangle {
	var _arg0 *C.GtkRange     // out
	var _arg1 *C.GdkRectangle // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	C.gtk_range_get_range_rect(_arg0, &_arg1)

	var _rangeRect gdk.Rectangle // out

	_rangeRect = (gdk.Rectangle)(unsafe.Pointer(_arg1))

	return _rangeRect
}

func (r _range) RestrictToFillLevel() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_restrict_to_fill_level(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (r _range) RoundDigits() int {
	var _arg0 *C.GtkRange // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_round_digits(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (r _range) ShowFillLevel() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_show_fill_level(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (r _range) SliderRange() (sliderStart int, sliderEnd int) {
	var _arg0 *C.GtkRange // out
	var _arg1 *C.gint     // in
	var _arg2 *C.gint     // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	C.gtk_range_get_slider_range(_arg0, &_arg1, &_arg2)

	var _sliderStart int // out
	var _sliderEnd int   // out

	_sliderStart = int(_arg1)
	_sliderEnd = int(_arg2)

	return _sliderStart, _sliderEnd
}

func (r _range) SliderSizeFixed() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_slider_size_fixed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (r _range) UpperStepperSensitivity() SensitivityType {
	var _arg0 *C.GtkRange          // out
	var _cret C.GtkSensitivityType // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_upper_stepper_sensitivity(_arg0)

	var _sensitivityType SensitivityType // out

	_sensitivityType = SensitivityType(_cret)

	return _sensitivityType
}

func (r _range) Value() float64 {
	var _arg0 *C.GtkRange // out
	var _cret C.gdouble   // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_value(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (r _range) SetAdjustmentRange(adjustment Adjustment) {
	var _arg0 *C.GtkRange      // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_range_set_adjustment(_arg0, _arg1)
}

func (r _range) SetFillLevelRange(fillLevel float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gdouble   // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = C.gdouble(fillLevel)

	C.gtk_range_set_fill_level(_arg0, _arg1)
}

func (r _range) SetFlippableRange(flippable bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	if flippable {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_flippable(_arg0, _arg1)
}

func (r _range) SetIncrementsRange(step float64, page float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gdouble   // out
	var _arg2 C.gdouble   // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = C.gdouble(step)
	_arg2 = C.gdouble(page)

	C.gtk_range_set_increments(_arg0, _arg1, _arg2)
}

func (r _range) SetInvertedRange(setting bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_inverted(_arg0, _arg1)
}

func (r _range) SetLowerStepperSensitivityRange(sensitivity SensitivityType) {
	var _arg0 *C.GtkRange          // out
	var _arg1 C.GtkSensitivityType // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = C.GtkSensitivityType(sensitivity)

	C.gtk_range_set_lower_stepper_sensitivity(_arg0, _arg1)
}

func (r _range) SetMinSliderSizeRange(minSize int) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = C.gint(minSize)

	C.gtk_range_set_min_slider_size(_arg0, _arg1)
}

func (r _range) SetRangeRange(min float64, max float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gdouble   // out
	var _arg2 C.gdouble   // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = C.gdouble(min)
	_arg2 = C.gdouble(max)

	C.gtk_range_set_range(_arg0, _arg1, _arg2)
}

func (r _range) SetRestrictToFillLevelRange(restrictToFillLevel bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	if restrictToFillLevel {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_restrict_to_fill_level(_arg0, _arg1)
}

func (r _range) SetRoundDigitsRange(roundDigits int) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = C.gint(roundDigits)

	C.gtk_range_set_round_digits(_arg0, _arg1)
}

func (r _range) SetShowFillLevelRange(showFillLevel bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	if showFillLevel {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_show_fill_level(_arg0, _arg1)
}

func (r _range) SetSliderSizeFixedRange(sizeFixed bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	if sizeFixed {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_slider_size_fixed(_arg0, _arg1)
}

func (r _range) SetUpperStepperSensitivityRange(sensitivity SensitivityType) {
	var _arg0 *C.GtkRange          // out
	var _arg1 C.GtkSensitivityType // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = C.GtkSensitivityType(sensitivity)

	C.gtk_range_set_upper_stepper_sensitivity(_arg0, _arg1)
}

func (r _range) SetValueRange(value float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gdouble   // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = C.gdouble(value)

	C.gtk_range_set_value(_arg0, _arg1)
}

func (_ _range) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(_))
}

func (_ _range) AsOrientable() Orientable {
	return WrapOrientable(gextras.InternObject(_))
}
