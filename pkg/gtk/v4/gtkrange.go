// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_range_get_type()), F: marshalRange},
	})
}

// Range: `GtkRange` is the common base class for widgets which visualize an
// adjustment.
//
// Widgets that are derived from `GtkRange` include [class@Gtk.Scale] and
// [class@Gtk.Scrollbar].
//
// Apart from signals for monitoring the parameters of the adjustment,
// `GtkRange` provides properties and methods for setting a “fill level” on
// range widgets. See [method@Gtk.Range.set_fill_level].
type Range interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Orientable

	// Adjustment: get the adjustment which is the “model” object for
	// `GtkRange`.
	Adjustment() Adjustment
	// FillLevel gets the current position of the fill level indicator.
	FillLevel() float64
	// Flippable gets whether the `GtkRange` respects text direction.
	//
	// See [method@Gtk.Range.set_flippable].
	Flippable() bool
	// Inverted gets whether the range is inverted.
	//
	// See [method@Gtk.Range.set_inverted].
	Inverted() bool
	// RangeRect: this function returns the area that contains the range’s
	// trough, in coordinates relative to @range's origin.
	//
	// This function is useful mainly for `GtkRange` subclasses.
	RangeRect() gdk.Rectangle
	// RestrictToFillLevel gets whether the range is restricted to the fill
	// level.
	RestrictToFillLevel() bool
	// RoundDigits gets the number of digits to round the value to when it
	// changes.
	//
	// See [signal@Gtk.Range::change-value].
	RoundDigits() int
	// ShowFillLevel gets whether the range displays the fill level graphically.
	ShowFillLevel() bool
	// SliderRange: this function returns sliders range along the long
	// dimension, in widget->window coordinates.
	//
	// This function is useful mainly for `GtkRange` subclasses.
	SliderRange() (sliderStart int, sliderEnd int)
	// SliderSizeFixed: this function is useful mainly for `GtkRange`
	// subclasses.
	//
	// See [method@Gtk.Range.set_slider_size_fixed].
	SliderSizeFixed() bool
	// Value gets the current value of the range.
	Value() float64
	// SetAdjustment sets the adjustment to be used as the “model” object for
	// the `GtkRange`
	//
	// The adjustment indicates the current range value, the minimum and maximum
	// range values, the step/page increments used for keybindings and
	// scrolling, and the page size.
	//
	// The page size is normally 0 for `GtkScale` and nonzero for
	// `GtkScrollbar`, and indicates the size of the visible area of the widget
	// being scrolled. The page size affects the size of the scrollbar slider.
	SetAdjustment(adjustment Adjustment)
	// SetFillLevel: set the new position of the fill level indicator.
	//
	// The “fill level” is probably best described by its most prominent use
	// case, which is an indicator for the amount of pre-buffering in a
	// streaming media player. In that use case, the value of the range would
	// indicate the current play position, and the fill level would be the
	// position up to which the file/stream has been downloaded.
	//
	// This amount of prebuffering can be displayed on the range’s trough and is
	// themeable separately from the trough. To enable fill level display, use
	// [method@Gtk.Range.set_show_fill_level]. The range defaults to not showing
	// the fill level.
	//
	// Additionally, it’s possible to restrict the range’s slider position to
	// values which are smaller than the fill level. This is controlled by
	// [method@Gtk.Range.set_restrict_to_fill_level] and is by default enabled.
	SetFillLevel(fillLevel float64)
	// SetFlippable sets whether the `GtkRange` respects text direction.
	//
	// If a range is flippable, it will switch its direction if it is horizontal
	// and its direction is GTK_TEXT_DIR_RTL.
	//
	// See [method@Gtk.Widget.get_direction].
	SetFlippable(flippable bool)
	// SetIncrements sets the step and page sizes for the range.
	//
	// The step size is used when the user clicks the `GtkScrollbar` arrows or
	// moves a `GtkScale` via arrow keys. The page size is used for example when
	// moving via Page Up or Page Down keys.
	SetIncrements(step float64, page float64)
	// SetInverted sets whether to invert the range.
	//
	// Ranges normally move from lower to higher values as the slider moves from
	// top to bottom or left to right. Inverted ranges have higher values at the
	// top or on the right rather than on the bottom or left.
	SetInverted(setting bool)
	// SetRange sets the allowable values in the `GtkRange`.
	//
	// The range value is clamped to be between @min and @max. (If the range has
	// a non-zero page size, it is clamped between @min and @max - page-size.)
	SetRange(min float64, max float64)
	// SetRestrictToFillLevel sets whether the slider is restricted to the fill
	// level.
	//
	// See [method@Gtk.Range.set_fill_level] for a general description of the
	// fill level concept.
	SetRestrictToFillLevel(restrictToFillLevel bool)
	// SetRoundDigits sets the number of digits to round the value to when it
	// changes.
	//
	// See [signal@Gtk.Range::change-value].
	SetRoundDigits(roundDigits int)
	// SetShowFillLevel sets whether a graphical fill level is show on the
	// trough.
	//
	// See [method@Gtk.Range.set_fill_level] for a general description of the
	// fill level concept.
	SetShowFillLevel(showFillLevel bool)
	// SetSliderSizeFixed sets whether the range’s slider has a fixed size, or a
	// size that depends on its adjustment’s page size.
	//
	// This function is useful mainly for `GtkRange` subclasses.
	SetSliderSizeFixed(sizeFixed bool)
	// SetValue sets the current value of the range.
	//
	// If the value is outside the minimum or maximum range values, it will be
	// clamped to fit inside them. The range emits the
	// [signal@Gtk.Range::value-changed] signal if the value changes.
	SetValue(value float64)
}

// _range implements the Range interface.
type _range struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Orientable
}

var _ Range = (*_range)(nil)

// WrapRange wraps a GObject to the right type. It is
// primarily used internally.
func WrapRange(obj *externglib.Object) Range {
	return Range{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Orientable:       WrapOrientable(obj),
	}
}

func marshalRange(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRange(obj), nil
}

// Adjustment: get the adjustment which is the “model” object for
// `GtkRange`.
func (r _range) Adjustment() Adjustment {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(r.Native())

	ret := C.gtk_range_get_adjustment(arg0)

	var ret0 Adjustment

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Adjustment)

	return ret0
}

// FillLevel gets the current position of the fill level indicator.
func (r _range) FillLevel() float64 {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(r.Native())

	ret := C.gtk_range_get_fill_level(arg0)

	var ret0 float64

	ret0 = float64(ret)

	return ret0
}

// Flippable gets whether the `GtkRange` respects text direction.
//
// See [method@Gtk.Range.set_flippable].
func (r _range) Flippable() bool {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(r.Native())

	ret := C.gtk_range_get_flippable(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Inverted gets whether the range is inverted.
//
// See [method@Gtk.Range.set_inverted].
func (r _range) Inverted() bool {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(r.Native())

	ret := C.gtk_range_get_inverted(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// RangeRect: this function returns the area that contains the range’s
// trough, in coordinates relative to @range's origin.
//
// This function is useful mainly for `GtkRange` subclasses.
func (r _range) RangeRect() gdk.Rectangle {
	var arg0 *C.GtkRange
	var arg1 *C.GdkRectangle // out

	arg0 = (*C.GtkRange)(r.Native())

	C.gtk_range_get_range_rect(arg0, &arg1)

	var ret0 *gdk.Rectangle

	{
		ret0 = gdk.WrapRectangle(unsafe.Pointer(arg1))
	}

	return ret0
}

// RestrictToFillLevel gets whether the range is restricted to the fill
// level.
func (r _range) RestrictToFillLevel() bool {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(r.Native())

	ret := C.gtk_range_get_restrict_to_fill_level(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// RoundDigits gets the number of digits to round the value to when it
// changes.
//
// See [signal@Gtk.Range::change-value].
func (r _range) RoundDigits() int {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(r.Native())

	ret := C.gtk_range_get_round_digits(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// ShowFillLevel gets whether the range displays the fill level graphically.
func (r _range) ShowFillLevel() bool {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(r.Native())

	ret := C.gtk_range_get_show_fill_level(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// SliderRange: this function returns sliders range along the long
// dimension, in widget->window coordinates.
//
// This function is useful mainly for `GtkRange` subclasses.
func (r _range) SliderRange() (sliderStart int, sliderEnd int) {
	var arg0 *C.GtkRange
	var arg1 *C.int // out
	var arg2 *C.int // out

	arg0 = (*C.GtkRange)(r.Native())

	C.gtk_range_get_slider_range(arg0, &arg1, &arg2)

	var ret0 int
	var ret1 int

	ret0 = int(arg1)

	ret1 = int(arg2)

	return ret0, ret1
}

// SliderSizeFixed: this function is useful mainly for `GtkRange`
// subclasses.
//
// See [method@Gtk.Range.set_slider_size_fixed].
func (r _range) SliderSizeFixed() bool {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(r.Native())

	ret := C.gtk_range_get_slider_size_fixed(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Value gets the current value of the range.
func (r _range) Value() float64 {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(r.Native())

	ret := C.gtk_range_get_value(arg0)

	var ret0 float64

	ret0 = float64(ret)

	return ret0
}

// SetAdjustment sets the adjustment to be used as the “model” object for
// the `GtkRange`
//
// The adjustment indicates the current range value, the minimum and maximum
// range values, the step/page increments used for keybindings and
// scrolling, and the page size.
//
// The page size is normally 0 for `GtkScale` and nonzero for
// `GtkScrollbar`, and indicates the size of the visible area of the widget
// being scrolled. The page size affects the size of the scrollbar slider.
func (r _range) SetAdjustment(adjustment Adjustment) {
	var arg0 *C.GtkRange
	var arg1 *C.GtkAdjustment

	arg0 = (*C.GtkRange)(r.Native())
	arg1 = (*C.GtkAdjustment)(adjustment.Native())

	C.gtk_range_set_adjustment(arg0, arg1)
}

// SetFillLevel: set the new position of the fill level indicator.
//
// The “fill level” is probably best described by its most prominent use
// case, which is an indicator for the amount of pre-buffering in a
// streaming media player. In that use case, the value of the range would
// indicate the current play position, and the fill level would be the
// position up to which the file/stream has been downloaded.
//
// This amount of prebuffering can be displayed on the range’s trough and is
// themeable separately from the trough. To enable fill level display, use
// [method@Gtk.Range.set_show_fill_level]. The range defaults to not showing
// the fill level.
//
// Additionally, it’s possible to restrict the range’s slider position to
// values which are smaller than the fill level. This is controlled by
// [method@Gtk.Range.set_restrict_to_fill_level] and is by default enabled.
func (r _range) SetFillLevel(fillLevel float64) {
	var arg0 *C.GtkRange
	var arg1 C.double

	arg0 = (*C.GtkRange)(r.Native())
	arg1 = C.double(fillLevel)

	C.gtk_range_set_fill_level(arg0, arg1)
}

// SetFlippable sets whether the `GtkRange` respects text direction.
//
// If a range is flippable, it will switch its direction if it is horizontal
// and its direction is GTK_TEXT_DIR_RTL.
//
// See [method@Gtk.Widget.get_direction].
func (r _range) SetFlippable(flippable bool) {
	var arg0 *C.GtkRange
	var arg1 C.gboolean

	arg0 = (*C.GtkRange)(r.Native())
	if flippable {
		arg1 = C.TRUE
	}

	C.gtk_range_set_flippable(arg0, arg1)
}

// SetIncrements sets the step and page sizes for the range.
//
// The step size is used when the user clicks the `GtkScrollbar` arrows or
// moves a `GtkScale` via arrow keys. The page size is used for example when
// moving via Page Up or Page Down keys.
func (r _range) SetIncrements(step float64, page float64) {
	var arg0 *C.GtkRange
	var arg1 C.double
	var arg2 C.double

	arg0 = (*C.GtkRange)(r.Native())
	arg1 = C.double(step)
	arg2 = C.double(page)

	C.gtk_range_set_increments(arg0, arg1, arg2)
}

// SetInverted sets whether to invert the range.
//
// Ranges normally move from lower to higher values as the slider moves from
// top to bottom or left to right. Inverted ranges have higher values at the
// top or on the right rather than on the bottom or left.
func (r _range) SetInverted(setting bool) {
	var arg0 *C.GtkRange
	var arg1 C.gboolean

	arg0 = (*C.GtkRange)(r.Native())
	if setting {
		arg1 = C.TRUE
	}

	C.gtk_range_set_inverted(arg0, arg1)
}

// SetRange sets the allowable values in the `GtkRange`.
//
// The range value is clamped to be between @min and @max. (If the range has
// a non-zero page size, it is clamped between @min and @max - page-size.)
func (r _range) SetRange(min float64, max float64) {
	var arg0 *C.GtkRange
	var arg1 C.double
	var arg2 C.double

	arg0 = (*C.GtkRange)(r.Native())
	arg1 = C.double(min)
	arg2 = C.double(max)

	C.gtk_range_set_range(arg0, arg1, arg2)
}

// SetRestrictToFillLevel sets whether the slider is restricted to the fill
// level.
//
// See [method@Gtk.Range.set_fill_level] for a general description of the
// fill level concept.
func (r _range) SetRestrictToFillLevel(restrictToFillLevel bool) {
	var arg0 *C.GtkRange
	var arg1 C.gboolean

	arg0 = (*C.GtkRange)(r.Native())
	if restrictToFillLevel {
		arg1 = C.TRUE
	}

	C.gtk_range_set_restrict_to_fill_level(arg0, arg1)
}

// SetRoundDigits sets the number of digits to round the value to when it
// changes.
//
// See [signal@Gtk.Range::change-value].
func (r _range) SetRoundDigits(roundDigits int) {
	var arg0 *C.GtkRange
	var arg1 C.int

	arg0 = (*C.GtkRange)(r.Native())
	arg1 = C.int(roundDigits)

	C.gtk_range_set_round_digits(arg0, arg1)
}

// SetShowFillLevel sets whether a graphical fill level is show on the
// trough.
//
// See [method@Gtk.Range.set_fill_level] for a general description of the
// fill level concept.
func (r _range) SetShowFillLevel(showFillLevel bool) {
	var arg0 *C.GtkRange
	var arg1 C.gboolean

	arg0 = (*C.GtkRange)(r.Native())
	if showFillLevel {
		arg1 = C.TRUE
	}

	C.gtk_range_set_show_fill_level(arg0, arg1)
}

// SetSliderSizeFixed sets whether the range’s slider has a fixed size, or a
// size that depends on its adjustment’s page size.
//
// This function is useful mainly for `GtkRange` subclasses.
func (r _range) SetSliderSizeFixed(sizeFixed bool) {
	var arg0 *C.GtkRange
	var arg1 C.gboolean

	arg0 = (*C.GtkRange)(r.Native())
	if sizeFixed {
		arg1 = C.TRUE
	}

	C.gtk_range_set_slider_size_fixed(arg0, arg1)
}

// SetValue sets the current value of the range.
//
// If the value is outside the minimum or maximum range values, it will be
// clamped to fit inside them. The range emits the
// [signal@Gtk.Range::value-changed] signal if the value changes.
func (r _range) SetValue(value float64) {
	var arg0 *C.GtkRange
	var arg1 C.double

	arg0 = (*C.GtkRange)(r.Native())
	arg1 = C.double(value)

	C.gtk_range_set_value(arg0, arg1)
}
