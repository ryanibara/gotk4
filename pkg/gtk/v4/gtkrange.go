// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_RangeClass_adjust_bounds(void*, double);
// extern void _gotk4_gtk4_RangeClass_get_range_border(void*, void*);
// extern void _gotk4_gtk4_RangeClass_value_changed(void*);
// extern void _gotk4_gtk4_Range_ConnectAdjustBounds(gpointer, gdouble, guintptr);
// extern void _gotk4_gtk4_Range_ConnectValueChanged(gpointer, guintptr);
import "C"

// GTypeRange returns the GType for the type Range.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeRange() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Range").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalRange)
	return gtype
}

// RangeOverrider contains methods that are overridable.
type RangeOverrider interface {
	// The function takes the following parameters:
	//
	AdjustBounds(newValue float64)
	// The function takes the following parameters:
	//
	RangeBorder(border_ *Border)
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

	*coreglib.Object
	Orientable
}

var (
	_ Widgetter         = (*Range)(nil)
	_ coreglib.Objector = (*Range)(nil)
)

func classInitRanger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "RangeClass")

	if _, ok := goval.(interface{ AdjustBounds(newValue float64) }); ok {
		o := pclass.StructFieldOffset("adjust_bounds")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_RangeClass_adjust_bounds)
	}

	if _, ok := goval.(interface{ RangeBorder(border_ *Border) }); ok {
		o := pclass.StructFieldOffset("get_range_border")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_RangeClass_get_range_border)
	}

	if _, ok := goval.(interface{ ValueChanged() }); ok {
		o := pclass.StructFieldOffset("value_changed")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_RangeClass_value_changed)
	}
}

//export _gotk4_gtk4_RangeClass_adjust_bounds
func _gotk4_gtk4_RangeClass_adjust_bounds(arg0 *C.void, arg1 C.double) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ AdjustBounds(newValue float64) })

	var _newValue float64 // out

	_newValue = float64(arg1)

	iface.AdjustBounds(_newValue)
}

//export _gotk4_gtk4_RangeClass_get_range_border
func _gotk4_gtk4_RangeClass_get_range_border(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ RangeBorder(border_ *Border) })

	var _border_ *Border // out

	_border_ = (*Border)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	iface.RangeBorder(_border_)
}

//export _gotk4_gtk4_RangeClass_value_changed
func _gotk4_gtk4_RangeClass_value_changed(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ValueChanged() })

	iface.ValueChanged()
}

func wrapRange(obj *coreglib.Object) *Range {
	return &Range{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
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

func marshalRange(p uintptr) (interface{}, error) {
	return wrapRange(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_Range_ConnectAdjustBounds
func _gotk4_gtk4_Range_ConnectAdjustBounds(arg0 C.gpointer, arg1 C.gdouble, arg2 C.guintptr) {
	var f func(value float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
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

// ConnectAdjustBounds is emitted before clamping a value, to give the
// application a chance to adjust the bounds.
func (_range *Range) ConnectAdjustBounds(f func(value float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(_range, "adjust-bounds", false, unsafe.Pointer(C._gotk4_gtk4_Range_ConnectAdjustBounds), f)
}

//export _gotk4_gtk4_Range_ConnectValueChanged
func _gotk4_gtk4_Range_ConnectValueChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectValueChanged is emitted when the range value changes.
func (_range *Range) ConnectValueChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(_range, "value-changed", false, unsafe.Pointer(C._gotk4_gtk4_Range_ConnectValueChanged), f)
}

// Adjustment: get the adjustment which is the “model” object for GtkRange.
//
// The function returns the following values:
//
//    - adjustment: GtkAdjustment.
//
func (_range *Range) Adjustment() *Adjustment {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))

	_info := girepository.MustFind("Gtk", "Range")
	_gret := _info.InvokeClassMethod("get_adjustment", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(_range)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _adjustment
}

// FillLevel gets the current position of the fill level indicator.
//
// The function returns the following values:
//
//    - gdouble: current fill level.
//
func (_range *Range) FillLevel() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))

	_info := girepository.MustFind("Gtk", "Range")
	_gret := _info.InvokeClassMethod("get_fill_level", _args[:], nil)
	_cret := *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(_range)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))

	_info := girepository.MustFind("Gtk", "Range")
	_gret := _info.InvokeClassMethod("get_flippable", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(_range)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))

	_info := girepository.MustFind("Gtk", "Range")
	_gret := _info.InvokeClassMethod("get_inverted", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(_range)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))

	_info := girepository.MustFind("Gtk", "Range")
	_info.InvokeClassMethod("get_range_rect", _args[:], _outs[:])

	runtime.KeepAlive(_range)

	var _rangeRect *gdk.Rectangle // out

	_rangeRect = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[0])))))

	return _rangeRect
}

// RestrictToFillLevel gets whether the range is restricted to the fill level.
//
// The function returns the following values:
//
//    - ok: TRUE if range is restricted to the fill level.
//
func (_range *Range) RestrictToFillLevel() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))

	_info := girepository.MustFind("Gtk", "Range")
	_gret := _info.InvokeClassMethod("get_restrict_to_fill_level", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(_range)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
func (_range *Range) RoundDigits() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))

	_info := girepository.MustFind("Gtk", "Range")
	_gret := _info.InvokeClassMethod("get_round_digits", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(_range)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// ShowFillLevel gets whether the range displays the fill level graphically.
//
// The function returns the following values:
//
//    - ok: TRUE if range shows the fill level.
//
func (_range *Range) ShowFillLevel() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))

	_info := girepository.MustFind("Gtk", "Range")
	_gret := _info.InvokeClassMethod("get_show_fill_level", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(_range)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
func (_range *Range) SliderRange() (sliderStart, sliderEnd int32) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))

	_info := girepository.MustFind("Gtk", "Range")
	_info.InvokeClassMethod("get_slider_range", _args[:], _outs[:])

	runtime.KeepAlive(_range)

	var _sliderStart int32 // out
	var _sliderEnd int32   // out

	_sliderStart = int32(*(*C.int)(unsafe.Pointer(&_outs[0])))
	_sliderEnd = int32(*(*C.int)(unsafe.Pointer(&_outs[1])))

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))

	_info := girepository.MustFind("Gtk", "Range")
	_gret := _info.InvokeClassMethod("get_slider_size_fixed", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(_range)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))

	_info := girepository.MustFind("Gtk", "Range")
	_gret := _info.InvokeClassMethod("get_value", _args[:], nil)
	_cret := *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(_range)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	_info := girepository.MustFind("Gtk", "Range")
	_info.InvokeClassMethod("set_adjustment", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(fillLevel)

	_info := girepository.MustFind("Gtk", "Range")
	_info.InvokeClassMethod("set_fill_level", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))
	if flippable {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Range")
	_info.InvokeClassMethod("set_flippable", _args[:], nil)

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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(step)
	*(*C.double)(unsafe.Pointer(&_args[2])) = C.double(page)

	_info := girepository.MustFind("Gtk", "Range")
	_info.InvokeClassMethod("set_increments", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))
	if setting {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Range")
	_info.InvokeClassMethod("set_inverted", _args[:], nil)

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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(min)
	*(*C.double)(unsafe.Pointer(&_args[2])) = C.double(max)

	_info := girepository.MustFind("Gtk", "Range")
	_info.InvokeClassMethod("set_range", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))
	if restrictToFillLevel {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Range")
	_info.InvokeClassMethod("set_restrict_to_fill_level", _args[:], nil)

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
func (_range *Range) SetRoundDigits(roundDigits int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(roundDigits)

	_info := girepository.MustFind("Gtk", "Range")
	_info.InvokeClassMethod("set_round_digits", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))
	if showFillLevel {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Range")
	_info.InvokeClassMethod("set_show_fill_level", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))
	if sizeFixed {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Range")
	_info.InvokeClassMethod("set_slider_size_fixed", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(_range).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(value)

	_info := girepository.MustFind("Gtk", "Range")
	_info.InvokeClassMethod("set_value", _args[:], nil)

	runtime.KeepAlive(_range)
	runtime.KeepAlive(value)
}
