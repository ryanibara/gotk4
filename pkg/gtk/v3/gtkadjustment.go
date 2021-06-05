// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_adjustment_get_type()), F: marshalAdjustment},
	})
}

// Adjustment: the Adjustment object represents a value which has an associated
// lower and upper bound, together with step and page increments, and a page
// size. It is used within several GTK+ widgets, including SpinButton, Viewport,
// and Range (which is a base class for Scrollbar and Scale).
//
// The Adjustment object does not update the value itself. Instead it is left up
// to the owner of the Adjustment to control the value.
type Adjustment interface {
	gextras.Objector

	// Changed emits a Adjustment::changed signal from the Adjustment. This is
	// typically called by the owner of the Adjustment after it has changed any
	// of the Adjustment properties other than the value.
	Changed()
	// ClampPage updates the Adjustment:value property to ensure that the range
	// between @lower and @upper is in the current page (i.e. between
	// Adjustment:value and Adjustment:value + Adjustment:page-size). If the
	// range is larger than the page size, then only the start of it will be in
	// the current page.
	//
	// A Adjustment::value-changed signal will be emitted if the value is
	// changed.
	ClampPage(lower float64, upper float64)
	// Configure sets all properties of the adjustment at once.
	//
	// Use this function to avoid multiple emissions of the Adjustment::changed
	// signal. See gtk_adjustment_set_lower() for an alternative way of
	// compressing multiple emissions of Adjustment::changed into one.
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
	// Value gets the current value of the adjustment. See
	// gtk_adjustment_set_value().
	Value() float64
	// SetLower sets the minimum value of the adjustment.
	//
	// When setting multiple adjustment properties via their individual setters,
	// multiple Adjustment::changed signals will be emitted. However, since the
	// emission of the Adjustment::changed signal is tied to the emission of the
	// #GObject::notify signals of the changed properties, it’s possible to
	// compress the Adjustment::changed signals into one by calling
	// g_object_freeze_notify() and g_object_thaw_notify() around the calls to
	// the individual setters.
	//
	// Alternatively, using a single g_object_set() for all the properties to
	// change, or using gtk_adjustment_configure() has the same effect of
	// compressing Adjustment::changed emissions.
	SetLower(lower float64)
	// SetPageIncrement sets the page increment of the adjustment.
	//
	// See gtk_adjustment_set_lower() about how to compress multiple emissions
	// of the Adjustment::changed signal when setting multiple adjustment
	// properties.
	SetPageIncrement(pageIncrement float64)
	// SetPageSize sets the page size of the adjustment.
	//
	// See gtk_adjustment_set_lower() about how to compress multiple emissions
	// of the GtkAdjustment::changed signal when setting multiple adjustment
	// properties.
	SetPageSize(pageSize float64)
	// SetStepIncrement sets the step increment of the adjustment.
	//
	// See gtk_adjustment_set_lower() about how to compress multiple emissions
	// of the Adjustment::changed signal when setting multiple adjustment
	// properties.
	SetStepIncrement(stepIncrement float64)
	// SetUpper sets the maximum value of the adjustment.
	//
	// Note that values will be restricted by `upper - page-size` if the
	// page-size property is nonzero.
	//
	// See gtk_adjustment_set_lower() about how to compress multiple emissions
	// of the Adjustment::changed signal when setting multiple adjustment
	// properties.
	SetUpper(upper float64)
	// SetValue sets the Adjustment value. The value is clamped to lie between
	// Adjustment:lower and Adjustment:upper.
	//
	// Note that for adjustments which are used in a Scrollbar, the effective
	// range of allowed values goes from Adjustment:lower to Adjustment:upper -
	// Adjustment:page-size.
	SetValue(value float64)
	// ValueChanged emits a Adjustment::value-changed signal from the
	// Adjustment. This is typically called by the owner of the Adjustment after
	// it has changed the Adjustment:value property.
	ValueChanged()
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
	var arg1 C.gdouble
	var arg2 C.gdouble
	var arg3 C.gdouble
	var arg4 C.gdouble
	var arg5 C.gdouble
	var arg6 C.gdouble

	arg1 = C.gdouble(value)
	arg2 = C.gdouble(lower)
	arg3 = C.gdouble(upper)
	arg4 = C.gdouble(stepIncrement)
	arg5 = C.gdouble(pageIncrement)
	arg6 = C.gdouble(pageSize)

	var cret C.GtkAdjustment
	var ret1 Adjustment

	cret = C.gtk_adjustment_new(value, lower, upper, stepIncrement, pageIncrement, pageSize)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Adjustment)

	return ret1
}

// Changed emits a Adjustment::changed signal from the Adjustment. This is
// typically called by the owner of the Adjustment after it has changed any
// of the Adjustment properties other than the value.
func (a adjustment) Changed() {
	var arg0 *C.GtkAdjustment

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	C.gtk_adjustment_changed(arg0)
}

// ClampPage updates the Adjustment:value property to ensure that the range
// between @lower and @upper is in the current page (i.e. between
// Adjustment:value and Adjustment:value + Adjustment:page-size). If the
// range is larger than the page size, then only the start of it will be in
// the current page.
//
// A Adjustment::value-changed signal will be emitted if the value is
// changed.
func (a adjustment) ClampPage(lower float64, upper float64) {
	var arg0 *C.GtkAdjustment
	var arg1 C.gdouble
	var arg2 C.gdouble

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))
	arg1 = C.gdouble(lower)
	arg2 = C.gdouble(upper)

	C.gtk_adjustment_clamp_page(arg0, lower, upper)
}

// Configure sets all properties of the adjustment at once.
//
// Use this function to avoid multiple emissions of the Adjustment::changed
// signal. See gtk_adjustment_set_lower() for an alternative way of
// compressing multiple emissions of Adjustment::changed into one.
func (a adjustment) Configure(value float64, lower float64, upper float64, stepIncrement float64, pageIncrement float64, pageSize float64) {
	var arg0 *C.GtkAdjustment
	var arg1 C.gdouble
	var arg2 C.gdouble
	var arg3 C.gdouble
	var arg4 C.gdouble
	var arg5 C.gdouble
	var arg6 C.gdouble

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))
	arg1 = C.gdouble(value)
	arg2 = C.gdouble(lower)
	arg3 = C.gdouble(upper)
	arg4 = C.gdouble(stepIncrement)
	arg5 = C.gdouble(pageIncrement)
	arg6 = C.gdouble(pageSize)

	C.gtk_adjustment_configure(arg0, value, lower, upper, stepIncrement, pageIncrement, pageSize)
}

// Lower retrieves the minimum value of the adjustment.
func (a adjustment) Lower() float64 {
	var arg0 *C.GtkAdjustment

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	var cret C.gdouble
	var ret1 float64

	cret = C.gtk_adjustment_get_lower(arg0)

	ret1 = C.gdouble(cret)

	return ret1
}

// MinimumIncrement gets the smaller of step increment and page increment.
func (a adjustment) MinimumIncrement() float64 {
	var arg0 *C.GtkAdjustment

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	var cret C.gdouble
	var ret1 float64

	cret = C.gtk_adjustment_get_minimum_increment(arg0)

	ret1 = C.gdouble(cret)

	return ret1
}

// PageIncrement retrieves the page increment of the adjustment.
func (a adjustment) PageIncrement() float64 {
	var arg0 *C.GtkAdjustment

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	var cret C.gdouble
	var ret1 float64

	cret = C.gtk_adjustment_get_page_increment(arg0)

	ret1 = C.gdouble(cret)

	return ret1
}

// PageSize retrieves the page size of the adjustment.
func (a adjustment) PageSize() float64 {
	var arg0 *C.GtkAdjustment

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	var cret C.gdouble
	var ret1 float64

	cret = C.gtk_adjustment_get_page_size(arg0)

	ret1 = C.gdouble(cret)

	return ret1
}

// StepIncrement retrieves the step increment of the adjustment.
func (a adjustment) StepIncrement() float64 {
	var arg0 *C.GtkAdjustment

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	var cret C.gdouble
	var ret1 float64

	cret = C.gtk_adjustment_get_step_increment(arg0)

	ret1 = C.gdouble(cret)

	return ret1
}

// Upper retrieves the maximum value of the adjustment.
func (a adjustment) Upper() float64 {
	var arg0 *C.GtkAdjustment

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	var cret C.gdouble
	var ret1 float64

	cret = C.gtk_adjustment_get_upper(arg0)

	ret1 = C.gdouble(cret)

	return ret1
}

// Value gets the current value of the adjustment. See
// gtk_adjustment_set_value().
func (a adjustment) Value() float64 {
	var arg0 *C.GtkAdjustment

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	var cret C.gdouble
	var ret1 float64

	cret = C.gtk_adjustment_get_value(arg0)

	ret1 = C.gdouble(cret)

	return ret1
}

// SetLower sets the minimum value of the adjustment.
//
// When setting multiple adjustment properties via their individual setters,
// multiple Adjustment::changed signals will be emitted. However, since the
// emission of the Adjustment::changed signal is tied to the emission of the
// #GObject::notify signals of the changed properties, it’s possible to
// compress the Adjustment::changed signals into one by calling
// g_object_freeze_notify() and g_object_thaw_notify() around the calls to
// the individual setters.
//
// Alternatively, using a single g_object_set() for all the properties to
// change, or using gtk_adjustment_configure() has the same effect of
// compressing Adjustment::changed emissions.
func (a adjustment) SetLower(lower float64) {
	var arg0 *C.GtkAdjustment
	var arg1 C.gdouble

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))
	arg1 = C.gdouble(lower)

	C.gtk_adjustment_set_lower(arg0, lower)
}

// SetPageIncrement sets the page increment of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions
// of the Adjustment::changed signal when setting multiple adjustment
// properties.
func (a adjustment) SetPageIncrement(pageIncrement float64) {
	var arg0 *C.GtkAdjustment
	var arg1 C.gdouble

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))
	arg1 = C.gdouble(pageIncrement)

	C.gtk_adjustment_set_page_increment(arg0, pageIncrement)
}

// SetPageSize sets the page size of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions
// of the GtkAdjustment::changed signal when setting multiple adjustment
// properties.
func (a adjustment) SetPageSize(pageSize float64) {
	var arg0 *C.GtkAdjustment
	var arg1 C.gdouble

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))
	arg1 = C.gdouble(pageSize)

	C.gtk_adjustment_set_page_size(arg0, pageSize)
}

// SetStepIncrement sets the step increment of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions
// of the Adjustment::changed signal when setting multiple adjustment
// properties.
func (a adjustment) SetStepIncrement(stepIncrement float64) {
	var arg0 *C.GtkAdjustment
	var arg1 C.gdouble

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))
	arg1 = C.gdouble(stepIncrement)

	C.gtk_adjustment_set_step_increment(arg0, stepIncrement)
}

// SetUpper sets the maximum value of the adjustment.
//
// Note that values will be restricted by `upper - page-size` if the
// page-size property is nonzero.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions
// of the Adjustment::changed signal when setting multiple adjustment
// properties.
func (a adjustment) SetUpper(upper float64) {
	var arg0 *C.GtkAdjustment
	var arg1 C.gdouble

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))
	arg1 = C.gdouble(upper)

	C.gtk_adjustment_set_upper(arg0, upper)
}

// SetValue sets the Adjustment value. The value is clamped to lie between
// Adjustment:lower and Adjustment:upper.
//
// Note that for adjustments which are used in a Scrollbar, the effective
// range of allowed values goes from Adjustment:lower to Adjustment:upper -
// Adjustment:page-size.
func (a adjustment) SetValue(value float64) {
	var arg0 *C.GtkAdjustment
	var arg1 C.gdouble

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))
	arg1 = C.gdouble(value)

	C.gtk_adjustment_set_value(arg0, value)
}

// ValueChanged emits a Adjustment::value-changed signal from the
// Adjustment. This is typically called by the owner of the Adjustment after
// it has changed the Adjustment:value property.
func (a adjustment) ValueChanged() {
	var arg0 *C.GtkAdjustment

	arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	C.gtk_adjustment_value_changed(arg0)
}
