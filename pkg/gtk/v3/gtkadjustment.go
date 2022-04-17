// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_AdjustmentClass_changed(GtkAdjustment*);
// extern void _gotk4_gtk3_AdjustmentClass_value_changed(GtkAdjustment*);
// extern void _gotk4_gtk3_Adjustment_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gtk3_Adjustment_ConnectValueChanged(gpointer, guintptr);
import "C"

// glib.Type values for gtkadjustment.go.
var GTypeAdjustment = externglib.Type(C.gtk_adjustment_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeAdjustment, F: marshalAdjustment},
	})
}

// AdjustmentOverrider contains methods that are overridable.
type AdjustmentOverrider interface {
	externglib.Objector
	// Changed emits a Adjustment::changed signal from the Adjustment. This is
	// typically called by the owner of the Adjustment after it has changed any
	// of the Adjustment properties other than the value.
	//
	// Deprecated: GTK+ emits Adjustment::changed itself whenever any of the
	// properties (other than value) change.
	Changed()
	// ValueChanged emits a Adjustment::value-changed signal from the
	// Adjustment. This is typically called by the owner of the Adjustment after
	// it has changed the Adjustment:value property.
	//
	// Deprecated: GTK+ emits Adjustment::value-changed itself whenever the
	// value changes.
	ValueChanged()
}

// Adjustment object represents a value which has an associated lower and upper
// bound, together with step and page increments, and a page size. It is used
// within several GTK+ widgets, including SpinButton, Viewport, and Range (which
// is a base class for Scrollbar and Scale).
//
// The Adjustment object does not update the value itself. Instead it is left up
// to the owner of the Adjustment to control the value.
type Adjustment struct {
	_ [0]func() // equal guard
	externglib.InitiallyUnowned
}

var ()

func classInitAdjustmenter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkAdjustmentClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkAdjustmentClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Changed() }); ok {
		pclass.changed = (*[0]byte)(C._gotk4_gtk3_AdjustmentClass_changed)
	}

	if _, ok := goval.(interface{ ValueChanged() }); ok {
		pclass.value_changed = (*[0]byte)(C._gotk4_gtk3_AdjustmentClass_value_changed)
	}
}

//export _gotk4_gtk3_AdjustmentClass_changed
func _gotk4_gtk3_AdjustmentClass_changed(arg0 *C.GtkAdjustment) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Changed() })

	iface.Changed()
}

//export _gotk4_gtk3_AdjustmentClass_value_changed
func _gotk4_gtk3_AdjustmentClass_value_changed(arg0 *C.GtkAdjustment) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ValueChanged() })

	iface.ValueChanged()
}

func wrapAdjustment(obj *externglib.Object) *Adjustment {
	return &Adjustment{
		InitiallyUnowned: externglib.InitiallyUnowned{
			Object: obj,
		},
	}
}

func marshalAdjustment(p uintptr) (interface{}, error) {
	return wrapAdjustment(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Adjustment_ConnectChanged
func _gotk4_gtk3_Adjustment_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectChanged is emitted when one or more of the Adjustment properties have
// been changed, other than the Adjustment:value property.
func (adjustment *Adjustment) ConnectChanged(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(adjustment, "changed", false, unsafe.Pointer(C._gotk4_gtk3_Adjustment_ConnectChanged), f)
}

//export _gotk4_gtk3_Adjustment_ConnectValueChanged
func _gotk4_gtk3_Adjustment_ConnectValueChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectValueChanged is emitted when the Adjustment:value property has been
// changed.
func (adjustment *Adjustment) ConnectValueChanged(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(adjustment, "value-changed", false, unsafe.Pointer(C._gotk4_gtk3_Adjustment_ConnectValueChanged), f)
}

// NewAdjustment creates a new Adjustment.
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
// The function returns the following values:
//
//    - adjustment: new Adjustment.
//
func NewAdjustment(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) *Adjustment {
	var _arg1 C.gdouble        // out
	var _arg2 C.gdouble        // out
	var _arg3 C.gdouble        // out
	var _arg4 C.gdouble        // out
	var _arg5 C.gdouble        // out
	var _arg6 C.gdouble        // out
	var _cret *C.GtkAdjustment // in

	_arg1 = C.gdouble(value)
	_arg2 = C.gdouble(lower)
	_arg3 = C.gdouble(upper)
	_arg4 = C.gdouble(stepIncrement)
	_arg5 = C.gdouble(pageIncrement)
	_arg6 = C.gdouble(pageSize)

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

// Changed emits a Adjustment::changed signal from the Adjustment. This is
// typically called by the owner of the Adjustment after it has changed any of
// the Adjustment properties other than the value.
//
// Deprecated: GTK+ emits Adjustment::changed itself whenever any of the
// properties (other than value) change.
func (adjustment *Adjustment) Changed() {
	var _arg0 *C.GtkAdjustment // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))

	C.gtk_adjustment_changed(_arg0)
	runtime.KeepAlive(adjustment)
}

// ClampPage updates the Adjustment:value property to ensure that the range
// between lower and upper is in the current page (i.e. between Adjustment:value
// and Adjustment:value + Adjustment:page-size). If the range is larger than the
// page size, then only the start of it will be in the current page.
//
// A Adjustment::value-changed signal will be emitted if the value is changed.
//
// The function takes the following parameters:
//
//    - lower value.
//    - upper value.
//
func (adjustment *Adjustment) ClampPage(lower, upper float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out
	var _arg2 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))
	_arg1 = C.gdouble(lower)
	_arg2 = C.gdouble(upper)

	C.gtk_adjustment_clamp_page(_arg0, _arg1, _arg2)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(lower)
	runtime.KeepAlive(upper)
}

// Configure sets all properties of the adjustment at once.
//
// Use this function to avoid multiple emissions of the Adjustment::changed
// signal. See gtk_adjustment_set_lower() for an alternative way of compressing
// multiple emissions of Adjustment::changed into one.
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
	var _arg1 C.gdouble        // out
	var _arg2 C.gdouble        // out
	var _arg3 C.gdouble        // out
	var _arg4 C.gdouble        // out
	var _arg5 C.gdouble        // out
	var _arg6 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))
	_arg1 = C.gdouble(value)
	_arg2 = C.gdouble(lower)
	_arg3 = C.gdouble(upper)
	_arg4 = C.gdouble(stepIncrement)
	_arg5 = C.gdouble(pageIncrement)
	_arg6 = C.gdouble(pageSize)

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
//
// The function returns the following values:
//
//    - gdouble: current minimum value of the adjustment.
//
func (adjustment *Adjustment) Lower() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))

	_cret = C.gtk_adjustment_get_lower(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// MinimumIncrement gets the smaller of step increment and page increment.
//
// The function returns the following values:
//
//    - gdouble: minimum increment of adjustment.
//
func (adjustment *Adjustment) MinimumIncrement() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))

	_cret = C.gtk_adjustment_get_minimum_increment(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// PageIncrement retrieves the page increment of the adjustment.
//
// The function returns the following values:
//
//    - gdouble: current page increment of the adjustment.
//
func (adjustment *Adjustment) PageIncrement() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))

	_cret = C.gtk_adjustment_get_page_increment(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// PageSize retrieves the page size of the adjustment.
//
// The function returns the following values:
//
//    - gdouble: current page size of the adjustment.
//
func (adjustment *Adjustment) PageSize() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))

	_cret = C.gtk_adjustment_get_page_size(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// StepIncrement retrieves the step increment of the adjustment.
//
// The function returns the following values:
//
//    - gdouble: current step increment of the adjustment.
//
func (adjustment *Adjustment) StepIncrement() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))

	_cret = C.gtk_adjustment_get_step_increment(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Upper retrieves the maximum value of the adjustment.
//
// The function returns the following values:
//
//    - gdouble: current maximum value of the adjustment.
//
func (adjustment *Adjustment) Upper() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))

	_cret = C.gtk_adjustment_get_upper(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Value gets the current value of the adjustment. See
// gtk_adjustment_set_value().
//
// The function returns the following values:
//
//    - gdouble: current value of the adjustment.
//
func (adjustment *Adjustment) Value() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))

	_cret = C.gtk_adjustment_get_value(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetLower sets the minimum value of the adjustment.
//
// When setting multiple adjustment properties via their individual setters,
// multiple Adjustment::changed signals will be emitted. However, since the
// emission of the Adjustment::changed signal is tied to the emission of the
// #GObject::notify signals of the changed properties, it’s possible to compress
// the Adjustment::changed signals into one by calling g_object_freeze_notify()
// and g_object_thaw_notify() around the calls to the individual setters.
//
// Alternatively, using a single g_object_set() for all the properties to
// change, or using gtk_adjustment_configure() has the same effect of
// compressing Adjustment::changed emissions.
//
// The function takes the following parameters:
//
//    - lower: new minimum value.
//
func (adjustment *Adjustment) SetLower(lower float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))
	_arg1 = C.gdouble(lower)

	C.gtk_adjustment_set_lower(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(lower)
}

// SetPageIncrement sets the page increment of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions of
// the Adjustment::changed signal when setting multiple adjustment properties.
//
// The function takes the following parameters:
//
//    - pageIncrement: new page increment.
//
func (adjustment *Adjustment) SetPageIncrement(pageIncrement float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))
	_arg1 = C.gdouble(pageIncrement)

	C.gtk_adjustment_set_page_increment(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(pageIncrement)
}

// SetPageSize sets the page size of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions of
// the GtkAdjustment::changed signal when setting multiple adjustment
// properties.
//
// The function takes the following parameters:
//
//    - pageSize: new page size.
//
func (adjustment *Adjustment) SetPageSize(pageSize float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))
	_arg1 = C.gdouble(pageSize)

	C.gtk_adjustment_set_page_size(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(pageSize)
}

// SetStepIncrement sets the step increment of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions of
// the Adjustment::changed signal when setting multiple adjustment properties.
//
// The function takes the following parameters:
//
//    - stepIncrement: new step increment.
//
func (adjustment *Adjustment) SetStepIncrement(stepIncrement float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))
	_arg1 = C.gdouble(stepIncrement)

	C.gtk_adjustment_set_step_increment(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(stepIncrement)
}

// SetUpper sets the maximum value of the adjustment.
//
// Note that values will be restricted by upper - page-size if the page-size
// property is nonzero.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions of
// the Adjustment::changed signal when setting multiple adjustment properties.
//
// The function takes the following parameters:
//
//    - upper: new maximum value.
//
func (adjustment *Adjustment) SetUpper(upper float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))
	_arg1 = C.gdouble(upper)

	C.gtk_adjustment_set_upper(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(upper)
}

// SetValue sets the Adjustment value. The value is clamped to lie between
// Adjustment:lower and Adjustment:upper.
//
// Note that for adjustments which are used in a Scrollbar, the effective range
// of allowed values goes from Adjustment:lower to Adjustment:upper -
// Adjustment:page-size.
//
// The function takes the following parameters:
//
//    - value: new value.
//
func (adjustment *Adjustment) SetValue(value float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))
	_arg1 = C.gdouble(value)

	C.gtk_adjustment_set_value(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(value)
}

// ValueChanged emits a Adjustment::value-changed signal from the Adjustment.
// This is typically called by the owner of the Adjustment after it has changed
// the Adjustment:value property.
//
// Deprecated: GTK+ emits Adjustment::value-changed itself whenever the value
// changes.
func (adjustment *Adjustment) ValueChanged() {
	var _arg0 *C.GtkAdjustment // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(adjustment).Native()))

	C.gtk_adjustment_value_changed(_arg0)
	runtime.KeepAlive(adjustment)
}
