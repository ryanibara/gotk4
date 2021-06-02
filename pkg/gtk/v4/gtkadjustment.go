// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_adjustment_get_type()), F: marshalAdjustment},
	})
}

// Adjustment: `GtkAdjustment` is a model for a numeric value.
//
// The `GtkAdjustment has an associated lower and upper bound. It also contains
// step and page increments, and a page size.
//
// Adjustments are used within several GTK widgets, including
// [class@Gtk.SpinButton], [class@Gtk.Viewport], [class@Gtk.Scrollbar] and
// [class@Gtk.Scale].
//
// The `GtkAdjustment` object does not update the value itself. Instead it is
// left up to the owner of the `GtkAdjustment` to control the value.
type Adjustment interface {
	gextras.Objector

	// ClampPage updates the value property to ensure that the range between
	// @lower and @upper is in the current page.
	//
	// The current page goes from `value` to `value` + `page-size`. If the range
	// is larger than the page size, then only the start of it will be in the
	// current page.
	//
	// A [signal@Gtk.Adjustment::value-changed] signal will be emitted if the
	// value is changed.
	ClampPage(lower float64, upper float64)
	// Configure sets all properties of the adjustment at once.
	//
	// Use this function to avoid multiple emissions of the
	// [signal@Gtk.Adjustment::changed] signal. See
	// [method@Gtk.Adjustment.set_lower] for an alternative way of compressing
	// multiple emissions of [signal@Gtk.Adjustment::changed] into one.
	Configure(value float64, lower float64, upper float64, stepIncrement float64, pageIncrement float64, pageSize float64)
	// Lower retrieves the minimum value of the adjustment.
	Lower() float64
	// MinimumIncrement gets the smaller of step increment and page increment.
	MinimumIncrement() float64
	// PageIncrement retrieves the page increment of the adjustment.
	PageIncrement() float64
	// PageSize retrieves the page size of the adjustment.
	PageSize() float64
	// StepIncrement retrieves the step increment of the adjustment.
	StepIncrement() float64
	// Upper retrieves the maximum value of the adjustment.
	Upper() float64
	// Value gets the current value of the adjustment.
	Value() float64
	// SetLower sets the minimum value of the adjustment.
	//
	// When setting multiple adjustment properties via their individual setters,
	// multiple [signal@Gtk.Adjustment::changed] signals will be emitted.
	// However, since the emission of the [signal@Gtk.Adjustment::changed]
	// signal is tied to the emission of the ::notify signals of the changed
	// properties, it’s possible to compress the
	// [signal@Gtk.Adjustment::changed] signals into one by calling
	// g_object_freeze_notify() and g_object_thaw_notify() around the calls to
	// the individual setters.
	//
	// Alternatively, using a single g_object_set() for all the properties to
	// change, or using [method@Gtk.Adjustment.configure] has the same effect.
	SetLower(lower float64)
	// SetPageIncrement sets the page increment of the adjustment.
	//
	// See [method@Gtk.Adjustment.set_lower] about how to compress multiple
	// emissions of the [signal@Gtk.Adjustment::changed] signal when setting
	// multiple adjustment properties.
	SetPageIncrement(pageIncrement float64)
	// SetPageSize sets the page size of the adjustment.
	//
	// See [method@Gtk.Adjustment.set_lower] about how to compress multiple
	// emissions of the [signal@Gtk.Adjustment::changed] signal when setting
	// multiple adjustment properties.
	SetPageSize(pageSize float64)
	// SetStepIncrement sets the step increment of the adjustment.
	//
	// See [method@Gtk.Adjustment.set_lower] about how to compress multiple
	// emissions of the [signal@Gtk.Adjustment::changed] signal when setting
	// multiple adjustment properties.
	SetStepIncrement(stepIncrement float64)
	// SetUpper sets the maximum value of the adjustment.
	//
	// Note that values will be restricted by `upper - page-size` if the
	// page-size property is nonzero.
	//
	// See [method@Gtk.Adjustment.set_lower] about how to compress multiple
	// emissions of the [signal@Gtk.Adjustment::changed] signal when setting
	// multiple adjustment properties.
	SetUpper(upper float64)
	// SetValue sets the `GtkAdjustment` value.
	//
	// The value is clamped to lie between [property@Gtk.Adjustment:lower] and
	// [property@Gtk.Adjustment:upper].
	//
	// Note that for adjustments which are used in a `GtkScrollbar`, the
	// effective range of allowed values goes from
	// [property@Gtk.Adjustment:lower] to [property@Gtk.Adjustment:upper] -
	// [property@Gtk.Adjustment:page-size].
	SetValue(value float64)
}

// adjustment implements the Adjustment interface.
type adjustment struct {
	gextras.Objector
}

var _ Adjustment = (*adjustment)(nil)

// WrapAdjustment wraps a GObject to the right type. It is
// primarily used internally.
func WrapAdjustment(obj *externglib.Object) Adjustment {
	return Adjustment{
		Objector: obj,
	}
}

func marshalAdjustment(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAdjustment(obj), nil
}

// NewAdjustment constructs a class Adjustment.
func NewAdjustment(value float64, lower float64, upper float64, stepIncrement float64, pageIncrement float64, pageSize float64) Adjustment {
	var arg1 C.double
	var arg2 C.double
	var arg3 C.double
	var arg4 C.double
	var arg5 C.double
	var arg6 C.double

	arg1 = C.double(value)
	arg2 = C.double(lower)
	arg3 = C.double(upper)
	arg4 = C.double(stepIncrement)
	arg5 = C.double(pageIncrement)
	arg6 = C.double(pageSize)

	ret := C.gtk_adjustment_new(arg1, arg2, arg3, arg4, arg5, arg6)

	var ret0 Adjustment

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Adjustment)

	return ret0
}

// ClampPage updates the value property to ensure that the range between
// @lower and @upper is in the current page.
//
// The current page goes from `value` to `value` + `page-size`. If the range
// is larger than the page size, then only the start of it will be in the
// current page.
//
// A [signal@Gtk.Adjustment::value-changed] signal will be emitted if the
// value is changed.
func (a adjustment) ClampPage(lower float64, upper float64) {
	var arg0 *C.GtkAdjustment
	var arg1 C.double
	var arg2 C.double

	arg0 = (*C.GtkAdjustment)(a.Native())
	arg1 = C.double(lower)
	arg2 = C.double(upper)

	C.gtk_adjustment_clamp_page(arg0, arg1, arg2)
}

// Configure sets all properties of the adjustment at once.
//
// Use this function to avoid multiple emissions of the
// [signal@Gtk.Adjustment::changed] signal. See
// [method@Gtk.Adjustment.set_lower] for an alternative way of compressing
// multiple emissions of [signal@Gtk.Adjustment::changed] into one.
func (a adjustment) Configure(value float64, lower float64, upper float64, stepIncrement float64, pageIncrement float64, pageSize float64) {
	var arg0 *C.GtkAdjustment
	var arg1 C.double
	var arg2 C.double
	var arg3 C.double
	var arg4 C.double
	var arg5 C.double
	var arg6 C.double

	arg0 = (*C.GtkAdjustment)(a.Native())
	arg1 = C.double(value)
	arg2 = C.double(lower)
	arg3 = C.double(upper)
	arg4 = C.double(stepIncrement)
	arg5 = C.double(pageIncrement)
	arg6 = C.double(pageSize)

	C.gtk_adjustment_configure(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// Lower retrieves the minimum value of the adjustment.
func (a adjustment) Lower() float64 {
	var arg0 *C.GtkAdjustment

	arg0 = (*C.GtkAdjustment)(a.Native())

	ret := C.gtk_adjustment_get_lower(arg0)

	var ret0 float64

	ret0 = float64(ret)

	return ret0
}

// MinimumIncrement gets the smaller of step increment and page increment.
func (a adjustment) MinimumIncrement() float64 {
	var arg0 *C.GtkAdjustment

	arg0 = (*C.GtkAdjustment)(a.Native())

	ret := C.gtk_adjustment_get_minimum_increment(arg0)

	var ret0 float64

	ret0 = float64(ret)

	return ret0
}

// PageIncrement retrieves the page increment of the adjustment.
func (a adjustment) PageIncrement() float64 {
	var arg0 *C.GtkAdjustment

	arg0 = (*C.GtkAdjustment)(a.Native())

	ret := C.gtk_adjustment_get_page_increment(arg0)

	var ret0 float64

	ret0 = float64(ret)

	return ret0
}

// PageSize retrieves the page size of the adjustment.
func (a adjustment) PageSize() float64 {
	var arg0 *C.GtkAdjustment

	arg0 = (*C.GtkAdjustment)(a.Native())

	ret := C.gtk_adjustment_get_page_size(arg0)

	var ret0 float64

	ret0 = float64(ret)

	return ret0
}

// StepIncrement retrieves the step increment of the adjustment.
func (a adjustment) StepIncrement() float64 {
	var arg0 *C.GtkAdjustment

	arg0 = (*C.GtkAdjustment)(a.Native())

	ret := C.gtk_adjustment_get_step_increment(arg0)

	var ret0 float64

	ret0 = float64(ret)

	return ret0
}

// Upper retrieves the maximum value of the adjustment.
func (a adjustment) Upper() float64 {
	var arg0 *C.GtkAdjustment

	arg0 = (*C.GtkAdjustment)(a.Native())

	ret := C.gtk_adjustment_get_upper(arg0)

	var ret0 float64

	ret0 = float64(ret)

	return ret0
}

// Value gets the current value of the adjustment.
func (a adjustment) Value() float64 {
	var arg0 *C.GtkAdjustment

	arg0 = (*C.GtkAdjustment)(a.Native())

	ret := C.gtk_adjustment_get_value(arg0)

	var ret0 float64

	ret0 = float64(ret)

	return ret0
}

// SetLower sets the minimum value of the adjustment.
//
// When setting multiple adjustment properties via their individual setters,
// multiple [signal@Gtk.Adjustment::changed] signals will be emitted.
// However, since the emission of the [signal@Gtk.Adjustment::changed]
// signal is tied to the emission of the ::notify signals of the changed
// properties, it’s possible to compress the
// [signal@Gtk.Adjustment::changed] signals into one by calling
// g_object_freeze_notify() and g_object_thaw_notify() around the calls to
// the individual setters.
//
// Alternatively, using a single g_object_set() for all the properties to
// change, or using [method@Gtk.Adjustment.configure] has the same effect.
func (a adjustment) SetLower(lower float64) {
	var arg0 *C.GtkAdjustment
	var arg1 C.double

	arg0 = (*C.GtkAdjustment)(a.Native())
	arg1 = C.double(lower)

	C.gtk_adjustment_set_lower(arg0, arg1)
}

// SetPageIncrement sets the page increment of the adjustment.
//
// See [method@Gtk.Adjustment.set_lower] about how to compress multiple
// emissions of the [signal@Gtk.Adjustment::changed] signal when setting
// multiple adjustment properties.
func (a adjustment) SetPageIncrement(pageIncrement float64) {
	var arg0 *C.GtkAdjustment
	var arg1 C.double

	arg0 = (*C.GtkAdjustment)(a.Native())
	arg1 = C.double(pageIncrement)

	C.gtk_adjustment_set_page_increment(arg0, arg1)
}

// SetPageSize sets the page size of the adjustment.
//
// See [method@Gtk.Adjustment.set_lower] about how to compress multiple
// emissions of the [signal@Gtk.Adjustment::changed] signal when setting
// multiple adjustment properties.
func (a adjustment) SetPageSize(pageSize float64) {
	var arg0 *C.GtkAdjustment
	var arg1 C.double

	arg0 = (*C.GtkAdjustment)(a.Native())
	arg1 = C.double(pageSize)

	C.gtk_adjustment_set_page_size(arg0, arg1)
}

// SetStepIncrement sets the step increment of the adjustment.
//
// See [method@Gtk.Adjustment.set_lower] about how to compress multiple
// emissions of the [signal@Gtk.Adjustment::changed] signal when setting
// multiple adjustment properties.
func (a adjustment) SetStepIncrement(stepIncrement float64) {
	var arg0 *C.GtkAdjustment
	var arg1 C.double

	arg0 = (*C.GtkAdjustment)(a.Native())
	arg1 = C.double(stepIncrement)

	C.gtk_adjustment_set_step_increment(arg0, arg1)
}

// SetUpper sets the maximum value of the adjustment.
//
// Note that values will be restricted by `upper - page-size` if the
// page-size property is nonzero.
//
// See [method@Gtk.Adjustment.set_lower] about how to compress multiple
// emissions of the [signal@Gtk.Adjustment::changed] signal when setting
// multiple adjustment properties.
func (a adjustment) SetUpper(upper float64) {
	var arg0 *C.GtkAdjustment
	var arg1 C.double

	arg0 = (*C.GtkAdjustment)(a.Native())
	arg1 = C.double(upper)

	C.gtk_adjustment_set_upper(arg0, arg1)
}

// SetValue sets the `GtkAdjustment` value.
//
// The value is clamped to lie between [property@Gtk.Adjustment:lower] and
// [property@Gtk.Adjustment:upper].
//
// Note that for adjustments which are used in a `GtkScrollbar`, the
// effective range of allowed values goes from
// [property@Gtk.Adjustment:lower] to [property@Gtk.Adjustment:upper] -
// [property@Gtk.Adjustment:page-size].
func (a adjustment) SetValue(value float64) {
	var arg0 *C.GtkAdjustment
	var arg1 C.double

	arg0 = (*C.GtkAdjustment)(a.Native())
	arg1 = C.double(value)

	C.gtk_adjustment_set_value(arg0, arg1)
}
