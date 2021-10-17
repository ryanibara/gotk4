// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_adjustment_get_type()), F: marshalAdjustmenter},
	})
}

// AdjustmentOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type AdjustmentOverrider interface {
	Changed()
	ValueChanged()
}

// Adjustment: GtkAdjustment is a model for a numeric value.
//
// The `GtkAdjustment has an associated lower and upper bound. It also contains
// step and page increments, and a page size.
//
// Adjustments are used within several GTK widgets, including gtk.SpinButton,
// gtk.Viewport, gtk.Scrollbar and gtk.Scale.
//
// The GtkAdjustment object does not update the value itself. Instead it is left
// up to the owner of the GtkAdjustment to control the value.
type Adjustment struct {
	externglib.InitiallyUnowned
}

func wrapAdjustment(obj *externglib.Object) *Adjustment {
	return &Adjustment{
		InitiallyUnowned: externglib.InitiallyUnowned{
			Object: obj,
		},
	}
}

func marshalAdjustmenter(p uintptr) (interface{}, error) {
	return wrapAdjustment(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAdjustment creates a new GtkAdjustment.
//
// The function takes the following parameters:
//
//    - value: initial value.
//    - lower: minimum value.
//    - upper: maximum value.
//    - stepIncrement: step increment.
//    - pageIncrement: page increment.
//    - pageSize: page size.
//
func NewAdjustment(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) *Adjustment {
	var _arg1 C.double         // out
	var _arg2 C.double         // out
	var _arg3 C.double         // out
	var _arg4 C.double         // out
	var _arg5 C.double         // out
	var _arg6 C.double         // out
	var _cret *C.GtkAdjustment // in

	_arg1 = C.double(value)
	_arg2 = C.double(lower)
	_arg3 = C.double(upper)
	_arg4 = C.double(stepIncrement)
	_arg5 = C.double(pageIncrement)
	_arg6 = C.double(pageSize)

	_cret = C.gtk_adjustment_new(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(value)
	runtime.KeepAlive(lower)
	runtime.KeepAlive(upper)
	runtime.KeepAlive(stepIncrement)
	runtime.KeepAlive(pageIncrement)
	runtime.KeepAlive(pageSize)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// ClampPage updates the value property to ensure that the range between lower
// and upper is in the current page.
//
// The current page goes from value to value + page-size. If the range is larger
// than the page size, then only the start of it will be in the current page.
//
// A gtk.Adjustment::value-changed signal will be emitted if the value is
// changed.
//
// The function takes the following parameters:
//
//    - lower value.
//    - upper value.
//
func (adjustment *Adjustment) ClampPage(lower, upper float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.double         // out
	var _arg2 C.double         // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	_arg1 = C.double(lower)
	_arg2 = C.double(upper)

	C.gtk_adjustment_clamp_page(_arg0, _arg1, _arg2)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(lower)
	runtime.KeepAlive(upper)
}

// Configure sets all properties of the adjustment at once.
//
// Use this function to avoid multiple emissions of the gtk.Adjustment::changed
// signal. See gtk.Adjustment.SetLower() for an alternative way of compressing
// multiple emissions of gtk.Adjustment::changed into one.
//
// The function takes the following parameters:
//
//    - value: new value.
//    - lower: new minimum value.
//    - upper: new maximum value.
//    - stepIncrement: new step increment.
//    - pageIncrement: new page increment.
//    - pageSize: new page size.
//
func (adjustment *Adjustment) Configure(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.double         // out
	var _arg2 C.double         // out
	var _arg3 C.double         // out
	var _arg4 C.double         // out
	var _arg5 C.double         // out
	var _arg6 C.double         // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	_arg1 = C.double(value)
	_arg2 = C.double(lower)
	_arg3 = C.double(upper)
	_arg4 = C.double(stepIncrement)
	_arg5 = C.double(pageIncrement)
	_arg6 = C.double(pageSize)

	C.gtk_adjustment_configure(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(value)
	runtime.KeepAlive(lower)
	runtime.KeepAlive(upper)
	runtime.KeepAlive(stepIncrement)
	runtime.KeepAlive(pageIncrement)
	runtime.KeepAlive(pageSize)
}

// Lower retrieves the minimum value of the adjustment.
func (adjustment *Adjustment) Lower() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.double         // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_adjustment_get_lower(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// MinimumIncrement gets the smaller of step increment and page increment.
func (adjustment *Adjustment) MinimumIncrement() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.double         // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_adjustment_get_minimum_increment(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// PageIncrement retrieves the page increment of the adjustment.
func (adjustment *Adjustment) PageIncrement() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.double         // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_adjustment_get_page_increment(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// PageSize retrieves the page size of the adjustment.
func (adjustment *Adjustment) PageSize() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.double         // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_adjustment_get_page_size(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// StepIncrement retrieves the step increment of the adjustment.
func (adjustment *Adjustment) StepIncrement() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.double         // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_adjustment_get_step_increment(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Upper retrieves the maximum value of the adjustment.
func (adjustment *Adjustment) Upper() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.double         // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_adjustment_get_upper(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Value gets the current value of the adjustment.
func (adjustment *Adjustment) Value() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.double         // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_adjustment_get_value(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetLower sets the minimum value of the adjustment.
//
// When setting multiple adjustment properties via their individual setters,
// multiple gtk.Adjustment::changed signals will be emitted. However, since the
// emission of the gtk.Adjustment::changed signal is tied to the emission of the
// ::notify signals of the changed properties, it’s possible to compress the
// gtk.Adjustment::changed signals into one by calling g_object_freeze_notify()
// and g_object_thaw_notify() around the calls to the individual setters.
//
// Alternatively, using a single g_object_set() for all the properties to
// change, or using gtk.Adjustment.Configure() has the same effect.
//
// The function takes the following parameters:
//
//    - lower: new minimum value.
//
func (adjustment *Adjustment) SetLower(lower float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.double         // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	_arg1 = C.double(lower)

	C.gtk_adjustment_set_lower(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(lower)
}

// SetPageIncrement sets the page increment of the adjustment.
//
// See gtk.Adjustment.SetLower() about how to compress multiple emissions of the
// gtk.Adjustment::changed signal when setting multiple adjustment properties.
//
// The function takes the following parameters:
//
//    - pageIncrement: new page increment.
//
func (adjustment *Adjustment) SetPageIncrement(pageIncrement float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.double         // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	_arg1 = C.double(pageIncrement)

	C.gtk_adjustment_set_page_increment(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(pageIncrement)
}

// SetPageSize sets the page size of the adjustment.
//
// See gtk.Adjustment.SetLower() about how to compress multiple emissions of the
// gtk.Adjustment::changed signal when setting multiple adjustment properties.
//
// The function takes the following parameters:
//
//    - pageSize: new page size.
//
func (adjustment *Adjustment) SetPageSize(pageSize float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.double         // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	_arg1 = C.double(pageSize)

	C.gtk_adjustment_set_page_size(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(pageSize)
}

// SetStepIncrement sets the step increment of the adjustment.
//
// See gtk.Adjustment.SetLower() about how to compress multiple emissions of the
// gtk.Adjustment::changed signal when setting multiple adjustment properties.
//
// The function takes the following parameters:
//
//    - stepIncrement: new step increment.
//
func (adjustment *Adjustment) SetStepIncrement(stepIncrement float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.double         // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	_arg1 = C.double(stepIncrement)

	C.gtk_adjustment_set_step_increment(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(stepIncrement)
}

// SetUpper sets the maximum value of the adjustment.
//
// Note that values will be restricted by upper - page-size if the page-size
// property is nonzero.
//
// See gtk.Adjustment.SetLower() about how to compress multiple emissions of the
// gtk.Adjustment::changed signal when setting multiple adjustment properties.
//
// The function takes the following parameters:
//
//    - upper: new maximum value.
//
func (adjustment *Adjustment) SetUpper(upper float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.double         // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	_arg1 = C.double(upper)

	C.gtk_adjustment_set_upper(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(upper)
}

// SetValue sets the GtkAdjustment value.
//
// The value is clamped to lie between gtk.Adjustment:lower and
// gtk.Adjustment:upper.
//
// Note that for adjustments which are used in a GtkScrollbar, the effective
// range of allowed values goes from gtk.Adjustment:lower to
// gtk.Adjustment:upper - gtk.Adjustment:page-size.
//
// The function takes the following parameters:
//
//    - value: new value.
//
func (adjustment *Adjustment) SetValue(value float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.double         // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	_arg1 = C.double(value)

	C.gtk_adjustment_set_value(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(value)
}

// ConnectChanged: emitted when one or more of the GtkAdjustment properties have
// been changed.
//
// Note that the gtk.Adjustment:value property is covered by the
// gtk.Adjustment::value-changed signal.
func (adjustment *Adjustment) ConnectChanged(f func()) externglib.SignalHandle {
	return adjustment.Connect("changed", f)
}

// ConnectValueChanged: emitted when the value has been changed.
func (adjustment *Adjustment) ConnectValueChanged(f func()) externglib.SignalHandle {
	return adjustment.Connect("value-changed", f)
}
