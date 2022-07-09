// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_AdjustmentClass_changed(void*);
// extern void _gotk4_gtk4_AdjustmentClass_value_changed(void*);
// extern void _gotk4_gtk4_Adjustment_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gtk4_Adjustment_ConnectValueChanged(gpointer, guintptr);
import "C"

// GTypeAdjustment returns the GType for the type Adjustment.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeAdjustment() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Adjustment").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalAdjustment)
	return gtype
}

// AdjustmentOverrider contains methods that are overridable.
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
	_ [0]func() // equal guard
	coreglib.InitiallyUnowned
}

var ()

func classInitAdjustmenter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "AdjustmentClass")

	if _, ok := goval.(interface{ Changed() }); ok {
		o := pclass.StructFieldOffset("changed")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_AdjustmentClass_changed)
	}

	if _, ok := goval.(interface{ ValueChanged() }); ok {
		o := pclass.StructFieldOffset("value_changed")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_AdjustmentClass_value_changed)
	}
}

//export _gotk4_gtk4_AdjustmentClass_changed
func _gotk4_gtk4_AdjustmentClass_changed(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Changed() })

	iface.Changed()
}

//export _gotk4_gtk4_AdjustmentClass_value_changed
func _gotk4_gtk4_AdjustmentClass_value_changed(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ValueChanged() })

	iface.ValueChanged()
}

func wrapAdjustment(obj *coreglib.Object) *Adjustment {
	return &Adjustment{
		InitiallyUnowned: coreglib.InitiallyUnowned{
			Object: obj,
		},
	}
}

func marshalAdjustment(p uintptr) (interface{}, error) {
	return wrapAdjustment(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_Adjustment_ConnectChanged
func _gotk4_gtk4_Adjustment_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectChanged is emitted when one or more of the GtkAdjustment properties
// have been changed.
//
// Note that the gtk.Adjustment:value property is covered by the
// gtk.Adjustment::value-changed signal.
func (adjustment *Adjustment) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(adjustment, "changed", false, unsafe.Pointer(C._gotk4_gtk4_Adjustment_ConnectChanged), f)
}

//export _gotk4_gtk4_Adjustment_ConnectValueChanged
func _gotk4_gtk4_Adjustment_ConnectValueChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectValueChanged is emitted when the value has been changed.
func (adjustment *Adjustment) ConnectValueChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(adjustment, "value-changed", false, unsafe.Pointer(C._gotk4_gtk4_Adjustment_ConnectValueChanged), f)
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
// The function returns the following values:
//
//    - adjustment: new GtkAdjustment.
//
func NewAdjustment(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) *Adjustment {
	var _args [6]girepository.Argument

	*(*C.double)(unsafe.Pointer(&_args[0])) = C.double(value)
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(lower)
	*(*C.double)(unsafe.Pointer(&_args[2])) = C.double(upper)
	*(*C.double)(unsafe.Pointer(&_args[3])) = C.double(stepIncrement)
	*(*C.double)(unsafe.Pointer(&_args[4])) = C.double(pageIncrement)
	*(*C.double)(unsafe.Pointer(&_args[5])) = C.double(pageSize)

	_gret := girepository.MustFind("Gtk", "Adjustment").InvokeMethod("new_Adjustment", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(value)
	runtime.KeepAlive(lower)
	runtime.KeepAlive(upper)
	runtime.KeepAlive(stepIncrement)
	runtime.KeepAlive(pageIncrement)
	runtime.KeepAlive(pageSize)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(_cret)))

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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(lower)
	*(*C.double)(unsafe.Pointer(&_args[2])) = C.double(upper)

	girepository.MustFind("Gtk", "Adjustment").InvokeMethod("clamp_page", _args[:], nil)

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
	var _args [7]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(value)
	*(*C.double)(unsafe.Pointer(&_args[2])) = C.double(lower)
	*(*C.double)(unsafe.Pointer(&_args[3])) = C.double(upper)
	*(*C.double)(unsafe.Pointer(&_args[4])) = C.double(stepIncrement)
	*(*C.double)(unsafe.Pointer(&_args[5])) = C.double(pageIncrement)
	*(*C.double)(unsafe.Pointer(&_args[6])) = C.double(pageSize)

	girepository.MustFind("Gtk", "Adjustment").InvokeMethod("configure", _args[:], nil)

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	_gret := girepository.MustFind("Gtk", "Adjustment").InvokeMethod("get_lower", _args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// MinimumIncrement gets the smaller of step increment and page increment.
//
// The function returns the following values:
//
//    - gdouble: minimum increment of adjustment.
//
func (adjustment *Adjustment) MinimumIncrement() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	_gret := girepository.MustFind("Gtk", "Adjustment").InvokeMethod("get_minimum_increment", _args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// PageIncrement retrieves the page increment of the adjustment.
//
// The function returns the following values:
//
//    - gdouble: current page increment of the adjustment.
//
func (adjustment *Adjustment) PageIncrement() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	_gret := girepository.MustFind("Gtk", "Adjustment").InvokeMethod("get_page_increment", _args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// PageSize retrieves the page size of the adjustment.
//
// The function returns the following values:
//
//    - gdouble: current page size of the adjustment.
//
func (adjustment *Adjustment) PageSize() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	_gret := girepository.MustFind("Gtk", "Adjustment").InvokeMethod("get_page_size", _args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// StepIncrement retrieves the step increment of the adjustment.
//
// The function returns the following values:
//
//    - gdouble: current step increment of the adjustment.
//
func (adjustment *Adjustment) StepIncrement() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	_gret := girepository.MustFind("Gtk", "Adjustment").InvokeMethod("get_step_increment", _args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// Upper retrieves the maximum value of the adjustment.
//
// The function returns the following values:
//
//    - gdouble: current maximum value of the adjustment.
//
func (adjustment *Adjustment) Upper() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	_gret := girepository.MustFind("Gtk", "Adjustment").InvokeMethod("get_upper", _args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// Value gets the current value of the adjustment.
//
// The function returns the following values:
//
//    - gdouble: current value of the adjustment.
//
func (adjustment *Adjustment) Value() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	_gret := girepository.MustFind("Gtk", "Adjustment").InvokeMethod("get_value", _args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(lower)

	girepository.MustFind("Gtk", "Adjustment").InvokeMethod("set_lower", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(pageIncrement)

	girepository.MustFind("Gtk", "Adjustment").InvokeMethod("set_page_increment", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(pageSize)

	girepository.MustFind("Gtk", "Adjustment").InvokeMethod("set_page_size", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(stepIncrement)

	girepository.MustFind("Gtk", "Adjustment").InvokeMethod("set_step_increment", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(upper)

	girepository.MustFind("Gtk", "Adjustment").InvokeMethod("set_upper", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(value)

	girepository.MustFind("Gtk", "Adjustment").InvokeMethod("set_value", _args[:], nil)

	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(value)
}
