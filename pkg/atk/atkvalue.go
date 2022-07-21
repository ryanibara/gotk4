// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
// extern AtkRange* _gotk4_atk1_ValueIface_get_range(AtkValue*);
// extern GSList* _gotk4_atk1_ValueIface_get_sub_ranges(AtkValue*);
// extern gboolean _gotk4_atk1_ValueIface_set_current_value(AtkValue*, GValue*);
// extern gdouble _gotk4_atk1_ValueIface_get_increment(AtkValue*);
// extern void _gotk4_atk1_ValueIface_get_current_value(AtkValue*, GValue*);
// extern void _gotk4_atk1_ValueIface_get_maximum_value(AtkValue*, GValue*);
// extern void _gotk4_atk1_ValueIface_get_minimum_increment(AtkValue*, GValue*);
// extern void _gotk4_atk1_ValueIface_get_minimum_value(AtkValue*, GValue*);
// extern void _gotk4_atk1_ValueIface_get_value_and_text(AtkValue*, gdouble*, gchar**);
// extern void _gotk4_atk1_ValueIface_set_value(AtkValue*, gdouble);
// extern void _gotk4_atk1_Value_ConnectValueChanged(gpointer, gdouble, gchar*, guintptr);
import "C"

// GTypeValueType returns the GType for the type ValueType.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeValueType() coreglib.Type {
	gtype := coreglib.Type(C.atk_value_type_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalValueType)
	return gtype
}

// GTypeValue returns the GType for the type Value.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeValue() coreglib.Type {
	gtype := coreglib.Type(C.atk_value_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalValue)
	return gtype
}

// ValueType: default types for a given value. Those are defined in order to
// easily get localized strings to describe a given value or a given subrange,
// using atk_value_type_get_localized_name().
type ValueType C.gint

const (
	ValueVeryWeak ValueType = iota
	ValueWeak
	ValueAcceptable
	ValueStrong
	ValueVeryStrong
	ValueVeryLow
	ValueLow
	ValueMedium
	ValueHigh
	ValueVeryHigh
	ValueVeryBad
	ValueBad
	ValueGood
	ValueVeryGood
	ValueBest
	ValueLastDefined
)

func marshalValueType(p uintptr) (interface{}, error) {
	return ValueType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ValueType.
func (v ValueType) String() string {
	switch v {
	case ValueVeryWeak:
		return "VeryWeak"
	case ValueWeak:
		return "Weak"
	case ValueAcceptable:
		return "Acceptable"
	case ValueStrong:
		return "Strong"
	case ValueVeryStrong:
		return "VeryStrong"
	case ValueVeryLow:
		return "VeryLow"
	case ValueLow:
		return "Low"
	case ValueMedium:
		return "Medium"
	case ValueHigh:
		return "High"
	case ValueVeryHigh:
		return "VeryHigh"
	case ValueVeryBad:
		return "VeryBad"
	case ValueBad:
		return "Bad"
	case ValueGood:
		return "Good"
	case ValueVeryGood:
		return "VeryGood"
	case ValueBest:
		return "Best"
	case ValueLastDefined:
		return "LastDefined"
	default:
		return fmt.Sprintf("ValueType(%d)", v)
	}
}

// ValueTypeGetLocalizedName gets the localized description string describing
// the ValueType value_type.
//
// The function takes the following parameters:
//
//    - valueType whose localized name is required.
//
// The function returns the following values:
//
//    - utf8: localized string describing the ValueType.
//
func ValueTypeGetLocalizedName(valueType ValueType) string {
	var _arg1 C.AtkValueType // out
	var _cret *C.gchar       // in

	_arg1 = C.AtkValueType(valueType)

	_cret = C.atk_value_type_get_localized_name(_arg1)
	runtime.KeepAlive(valueType)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ValueTypeGetName gets the description string describing the ValueType
// value_type.
//
// The function takes the following parameters:
//
//    - valueType whose name is required.
//
// The function returns the following values:
//
//    - utf8: string describing the ValueType.
//
func ValueTypeGetName(valueType ValueType) string {
	var _arg1 C.AtkValueType // out
	var _cret *C.gchar       // in

	_arg1 = C.AtkValueType(valueType)

	_cret = C.atk_value_type_get_name(_arg1)
	runtime.KeepAlive(valueType)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ValueOverrider contains methods that are overridable.
type ValueOverrider interface {
	// CurrentValue gets the value of this object.
	//
	// Deprecated: Since 2.12. Use atk_value_get_value_and_text() instead.
	//
	// The function returns the following values:
	//
	//    - value representing the current accessible value.
	//
	CurrentValue() coreglib.Value
	// Increment gets the minimum increment by which the value of this object
	// may be changed. If zero, the minimum increment is undefined, which may
	// mean that it is limited only by the floating point precision of the
	// platform.
	//
	// The function returns the following values:
	//
	//    - gdouble: minimum increment by which the value of this object may be
	//      changed. zero if undefined.
	//
	Increment() float64
	// MaximumValue gets the maximum value of this object.
	//
	// Deprecated: Since 2.12. Use atk_value_get_range() instead.
	//
	// The function returns the following values:
	//
	//    - value representing the maximum accessible value.
	//
	MaximumValue() coreglib.Value
	// MinimumIncrement gets the minimum increment by which the value of this
	// object may be changed. If zero, the minimum increment is undefined, which
	// may mean that it is limited only by the floating point precision of the
	// platform.
	//
	// Deprecated: Since 2.12. Use atk_value_get_increment() instead.
	//
	// The function returns the following values:
	//
	//    - value representing the minimum increment by which the accessible
	//      value may be changed.
	//
	MinimumIncrement() coreglib.Value
	// MinimumValue gets the minimum value of this object.
	//
	// Deprecated: Since 2.12. Use atk_value_get_range() instead.
	//
	// The function returns the following values:
	//
	//    - value representing the minimum accessible value.
	//
	MinimumValue() coreglib.Value
	// Range gets the range of this object.
	//
	// The function returns the following values:
	//
	//    - _range (optional): newly allocated Range that represents the minimum,
	//      maximum and descriptor (if available) of obj. NULL if that range is
	//      not defined.
	//
	Range() *Range
	// SubRanges gets the list of subranges defined for this object. See Value
	// introduction for examples of subranges and when to expose them.
	//
	// The function returns the following values:
	//
	//    - sList of Range which each of the subranges defined for this object.
	//      Free the returns list with g_slist_free().
	//
	SubRanges() []*Range
	// ValueAndText gets the current value and the human readable text
	// alternative of obj. text is a newly created string, that must be freed by
	// the caller. Can be NULL if no descriptor is available.
	//
	// The function returns the following values:
	//
	//    - value address of #gdouble to put the current value of obj.
	//    - text (optional) address of #gchar to put the human readable text
	//      alternative for value.
	//
	ValueAndText() (float64, string)
	// SetCurrentValue sets the value of this object.
	//
	// Deprecated: Since 2.12. Use atk_value_set_value() instead.
	//
	// The function takes the following parameters:
	//
	//    - value which is the desired new accessible value.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if new value is successfully set, FALSE otherwise.
	//
	SetCurrentValue(value *coreglib.Value) bool
	// SetValue sets the value of this object.
	//
	// This method is intended to provide a way to change the value of the
	// object. In any case, it is possible that the value can't be modified (ie:
	// a read-only component). If the value changes due this call, it is
	// possible that the text could change, and will trigger an
	// Value::value-changed signal emission.
	//
	// Note for implementors: the deprecated atk_value_set_current_value()
	// method returned TRUE or FALSE depending if the value was assigned or not.
	// In the practice several implementors were not able to decide it, and
	// returned TRUE in any case. For that reason it is not required anymore to
	// return if the value was properly assigned or not.
	//
	// The function takes the following parameters:
	//
	//    - newValue: double which is the desired new accessible value.
	//
	SetValue(newValue float64)
}

// Value should be implemented for components which either display a value from
// a bounded range, or which allow the user to specify a value from a bounded
// range, or both. For instance, most sliders and range controls, as well as
// dials, should have Object representations which implement Value on the
// component's behalf. KValues may be read-only, in which case attempts to alter
// the value return would fail.
//
// <refsect1 id="current-value-text"> <title>On the subject of current value
// text</title> <para> In addition to providing the current value, implementors
// can optionally provide an end-user-consumable textual description associated
// with this value. This description should be included when the numeric value
// fails to convey the full, on-screen representation seen by users. </para>
//
// <example> <title>Password strength</title> A password strength meter whose
// value changes as the user types their new password. Red is used for values
// less than 4.0, yellow for values between 4.0 and 7.0, and green for values
// greater than 7.0. In this instance, value text should be provided by the
// implementor. Appropriate value text would be "weak", "acceptable," and
// "strong" respectively. </example>
//
// A level bar whose value changes to reflect the battery charge. The color
// remains the same regardless of the charge and there is no on-screen text
// reflecting the fullness of the battery. In this case, because the position
// within the bar is the only indication the user has of the current charge,
// value text should not be provided by the implementor.
//
// <refsect2 id="implementor-notes"> <title>Implementor Notes</title> <para>
// Implementors should bear in mind that assistive technologies will likely
// prefer the value text provided over the numeric value when presenting a
// widget's value. As a result, strings not intended for end users should not be
// exposed in the value text, and strings which are exposed should be localized.
// In the case of widgets which display value text on screen, for instance
// through a separate label in close proximity to the value-displaying widget,
// it is still expected that implementors will expose the value text using the
// above API. </para>
//
// <para> Value should NOT be implemented for widgets whose displayed value is
// not reflective of a meaningful amount. For instance, a progress pulse
// indicator whose value alternates between 0.0 and 1.0 to indicate that some
// process is still taking place should not implement Value because the current
// value does not reflect progress towards completion. </para> </refsect2>
// </refsect1>
//
// <refsect1 id="ranges"> <title>On the subject of ranges</title> <para> In
// addition to providing the minimum and maximum values, implementors can
// optionally provide details about subranges associated with the widget. These
// details should be provided by the implementor when both of the following are
// communicated visually to the end user: </para> <itemizedlist> <listitem>The
// existence of distinct ranges such as "weak", "acceptable", and "strong"
// indicated by color, bar tick marks, and/or on-screen text.</listitem>
// <listitem>Where the current value stands within a given subrange, for
// instance illustrating progression from very "weak" towards nearly
// "acceptable" through changes in shade and/or position on the bar within the
// "weak" subrange.</listitem> </itemizedlist> <para> If both of the above do
// not apply to the widget, it should be sufficient to expose the numeric value,
// along with the value text if appropriate, to make the widget accessible.
// </para>
//
// <refsect2 id="ranges-implementor-notes"> <title>Implementor Notes</title>
// <para> If providing subrange details is deemed necessary, all possible values
// of the widget are expected to fall within one of the subranges defined by the
// implementor. </para> </refsect2> </refsect1>
//
// <refsect1 id="localization"> <title>On the subject of localization of
// end-user-consumable text values</title> <para> Because value text and
// subrange descriptors are human-consumable, implementors are expected to
// provide localized strings which can be directly presented to end users via
// their assistive technology. In order to simplify this for implementors,
// implementors can use atk_value_type_get_localized_name() with the following
// already-localized constants for commonly-needed values can be used: </para>
//
// <itemizedlist> <listitem>ATK_VALUE_VERY_WEAK</listitem>
// <listitem>ATK_VALUE_WEAK</listitem> <listitem>ATK_VALUE_ACCEPTABLE</listitem>
// <listitem>ATK_VALUE_STRONG</listitem>
// <listitem>ATK_VALUE_VERY_STRONG</listitem>
// <listitem>ATK_VALUE_VERY_LOW</listitem> <listitem>ATK_VALUE_LOW</listitem>
// <listitem>ATK_VALUE_MEDIUM</listitem> <listitem>ATK_VALUE_HIGH</listitem>
// <listitem>ATK_VALUE_VERY_HIGH</listitem>
// <listitem>ATK_VALUE_VERY_BAD</listitem> <listitem>ATK_VALUE_BAD</listitem>
// <listitem>ATK_VALUE_GOOD</listitem> <listitem>ATK_VALUE_VERY_GOOD</listitem>
// <listitem>ATK_VALUE_BEST</listitem>
// <listitem>ATK_VALUE_SUBSUBOPTIMAL</listitem>
// <listitem>ATK_VALUE_SUBOPTIMAL</listitem>
// <listitem>ATK_VALUE_OPTIMAL</listitem> </itemizedlist> <para> Proposals for
// additional constants, along with their use cases, should be submitted to the
// GNOME Accessibility Team. </para> </refsect1>
//
// <refsect1 id="changes"> <title>On the subject of changes</title> <para> Note
// that if there is a textual description associated with the new numeric value,
// that description should be included regardless of whether or not it has also
// changed. </para> </refsect1>.
//
// Value wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Value struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Value)(nil)
)

// Valueer describes Value's interface methods.
type Valueer interface {
	coreglib.Objector

	// CurrentValue gets the value of this object.
	CurrentValue() coreglib.Value
	// Increment gets the minimum increment by which the value of this object
	// may be changed.
	Increment() float64
	// MaximumValue gets the maximum value of this object.
	MaximumValue() coreglib.Value
	// MinimumIncrement gets the minimum increment by which the value of this
	// object may be changed.
	MinimumIncrement() coreglib.Value
	// MinimumValue gets the minimum value of this object.
	MinimumValue() coreglib.Value
	// Range gets the range of this object.
	Range() *Range
	// SubRanges gets the list of subranges defined for this object.
	SubRanges() []*Range
	// ValueAndText gets the current value and the human readable text
	// alternative of obj.
	ValueAndText() (float64, string)
	// SetCurrentValue sets the value of this object.
	SetCurrentValue(value *coreglib.Value) bool
	// SetValue sets the value of this object.
	SetValue(newValue float64)

	// Value-changed: 'value-changed' signal is emitted when the current value
	// that represent the object changes.
	ConnectValueChanged(func(value float64, text string)) coreglib.SignalHandle
}

var _ Valueer = (*Value)(nil)

func ifaceInitValueer(gifacePtr, data C.gpointer) {
	iface := (*C.AtkValueIface)(unsafe.Pointer(gifacePtr))
	iface.get_current_value = (*[0]byte)(C._gotk4_atk1_ValueIface_get_current_value)
	iface.get_increment = (*[0]byte)(C._gotk4_atk1_ValueIface_get_increment)
	iface.get_maximum_value = (*[0]byte)(C._gotk4_atk1_ValueIface_get_maximum_value)
	iface.get_minimum_increment = (*[0]byte)(C._gotk4_atk1_ValueIface_get_minimum_increment)
	iface.get_minimum_value = (*[0]byte)(C._gotk4_atk1_ValueIface_get_minimum_value)
	iface.get_range = (*[0]byte)(C._gotk4_atk1_ValueIface_get_range)
	iface.get_sub_ranges = (*[0]byte)(C._gotk4_atk1_ValueIface_get_sub_ranges)
	iface.get_value_and_text = (*[0]byte)(C._gotk4_atk1_ValueIface_get_value_and_text)
	iface.set_current_value = (*[0]byte)(C._gotk4_atk1_ValueIface_set_current_value)
	iface.set_value = (*[0]byte)(C._gotk4_atk1_ValueIface_set_value)
}

//export _gotk4_atk1_ValueIface_get_current_value
func _gotk4_atk1_ValueIface_get_current_value(arg0 *C.AtkValue, arg1 *C.GValue) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ValueOverrider)

	value := iface.CurrentValue()

	*arg1 = *(*C.GValue)(unsafe.Pointer((&value).Native()))
}

//export _gotk4_atk1_ValueIface_get_increment
func _gotk4_atk1_ValueIface_get_increment(arg0 *C.AtkValue) (cret C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ValueOverrider)

	gdouble := iface.Increment()

	cret = C.gdouble(gdouble)

	return cret
}

//export _gotk4_atk1_ValueIface_get_maximum_value
func _gotk4_atk1_ValueIface_get_maximum_value(arg0 *C.AtkValue, arg1 *C.GValue) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ValueOverrider)

	value := iface.MaximumValue()

	*arg1 = *(*C.GValue)(unsafe.Pointer((&value).Native()))
}

//export _gotk4_atk1_ValueIface_get_minimum_increment
func _gotk4_atk1_ValueIface_get_minimum_increment(arg0 *C.AtkValue, arg1 *C.GValue) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ValueOverrider)

	value := iface.MinimumIncrement()

	*arg1 = *(*C.GValue)(unsafe.Pointer((&value).Native()))
}

//export _gotk4_atk1_ValueIface_get_minimum_value
func _gotk4_atk1_ValueIface_get_minimum_value(arg0 *C.AtkValue, arg1 *C.GValue) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ValueOverrider)

	value := iface.MinimumValue()

	*arg1 = *(*C.GValue)(unsafe.Pointer((&value).Native()))
}

//export _gotk4_atk1_ValueIface_get_range
func _gotk4_atk1_ValueIface_get_range(arg0 *C.AtkValue) (cret *C.AtkRange) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ValueOverrider)

	_range := iface.Range()

	if _range != nil {
		cret = (*C.AtkRange)(gextras.StructNative(unsafe.Pointer(_range)))
		runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(_range)), nil)
	}

	return cret
}

//export _gotk4_atk1_ValueIface_get_sub_ranges
func _gotk4_atk1_ValueIface_get_sub_ranges(arg0 *C.AtkValue) (cret *C.GSList) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ValueOverrider)

	sList := iface.SubRanges()

	for i := len(sList) - 1; i >= 0; i-- {
		src := sList[i]
		var dst *C.AtkRange // out
		dst = (*C.AtkRange)(gextras.StructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(src)), nil)
		cret = C.g_slist_prepend(cret, C.gpointer(unsafe.Pointer(dst)))
	}

	return cret
}

//export _gotk4_atk1_ValueIface_get_value_and_text
func _gotk4_atk1_ValueIface_get_value_and_text(arg0 *C.AtkValue, arg1 *C.gdouble, arg2 **C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ValueOverrider)

	value, text := iface.ValueAndText()

	*arg1 = C.gdouble(value)
	if text != "" {
		*arg2 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	}
}

//export _gotk4_atk1_ValueIface_set_current_value
func _gotk4_atk1_ValueIface_set_current_value(arg0 *C.AtkValue, arg1 *C.GValue) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ValueOverrider)

	var _value *coreglib.Value // out

	_value = coreglib.ValueFromNative(unsafe.Pointer(arg1))

	ok := iface.SetCurrentValue(_value)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_atk1_ValueIface_set_value
func _gotk4_atk1_ValueIface_set_value(arg0 *C.AtkValue, arg1 C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ValueOverrider)

	var _newValue float64 // out

	_newValue = float64(arg1)

	iface.SetValue(_newValue)
}

func wrapValue(obj *coreglib.Object) *Value {
	return &Value{
		Object: obj,
	}
}

func marshalValue(p uintptr) (interface{}, error) {
	return wrapValue(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_atk1_Value_ConnectValueChanged
func _gotk4_atk1_Value_ConnectValueChanged(arg0 C.gpointer, arg1 C.gdouble, arg2 *C.gchar, arg3 C.guintptr) {
	var f func(value float64, text string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(value float64, text string))
	}

	var _value float64 // out
	var _text string   // out

	_value = float64(arg1)
	_text = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	f(_value, _text)
}

// ConnectValueChanged: 'value-changed' signal is emitted when the current value
// that represent the object changes. value is the numerical representation of
// this new value. text is the human readable text alternative of value, and can
// be NULL if it is not available. Note that if there is a textual description
// associated with the new numeric value, that description should be included
// regardless of whether or not it has also changed.
//
// Example: a password meter whose value changes as the user types their new
// password. Appropiate value text would be "weak", "acceptable" and "strong".
func (obj *Value) ConnectValueChanged(f func(value float64, text string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(obj, "value-changed", false, unsafe.Pointer(C._gotk4_atk1_Value_ConnectValueChanged), f)
}

// CurrentValue gets the value of this object.
//
// Deprecated: Since 2.12. Use atk_value_get_value_and_text() instead.
//
// The function returns the following values:
//
//    - value representing the current accessible value.
//
func (obj *Value) CurrentValue() coreglib.Value {
	var _arg0 *C.AtkValue // out
	var _arg1 C.GValue    // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(coreglib.InternObject(obj).Native()))

	C.atk_value_get_current_value(_arg0, &_arg1)
	runtime.KeepAlive(obj)

	var _value coreglib.Value // out

	_value = *coreglib.ValueFromNative(unsafe.Pointer((&_arg1)))

	return _value
}

// Increment gets the minimum increment by which the value of this object may be
// changed. If zero, the minimum increment is undefined, which may mean that it
// is limited only by the floating point precision of the platform.
//
// The function returns the following values:
//
//    - gdouble: minimum increment by which the value of this object may be
//      changed. zero if undefined.
//
func (obj *Value) Increment() float64 {
	var _arg0 *C.AtkValue // out
	var _cret C.gdouble   // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(coreglib.InternObject(obj).Native()))

	_cret = C.atk_value_get_increment(_arg0)
	runtime.KeepAlive(obj)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// MaximumValue gets the maximum value of this object.
//
// Deprecated: Since 2.12. Use atk_value_get_range() instead.
//
// The function returns the following values:
//
//    - value representing the maximum accessible value.
//
func (obj *Value) MaximumValue() coreglib.Value {
	var _arg0 *C.AtkValue // out
	var _arg1 C.GValue    // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(coreglib.InternObject(obj).Native()))

	C.atk_value_get_maximum_value(_arg0, &_arg1)
	runtime.KeepAlive(obj)

	var _value coreglib.Value // out

	_value = *coreglib.ValueFromNative(unsafe.Pointer((&_arg1)))

	return _value
}

// MinimumIncrement gets the minimum increment by which the value of this object
// may be changed. If zero, the minimum increment is undefined, which may mean
// that it is limited only by the floating point precision of the platform.
//
// Deprecated: Since 2.12. Use atk_value_get_increment() instead.
//
// The function returns the following values:
//
//    - value representing the minimum increment by which the accessible value
//      may be changed.
//
func (obj *Value) MinimumIncrement() coreglib.Value {
	var _arg0 *C.AtkValue // out
	var _arg1 C.GValue    // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(coreglib.InternObject(obj).Native()))

	C.atk_value_get_minimum_increment(_arg0, &_arg1)
	runtime.KeepAlive(obj)

	var _value coreglib.Value // out

	_value = *coreglib.ValueFromNative(unsafe.Pointer((&_arg1)))

	return _value
}

// MinimumValue gets the minimum value of this object.
//
// Deprecated: Since 2.12. Use atk_value_get_range() instead.
//
// The function returns the following values:
//
//    - value representing the minimum accessible value.
//
func (obj *Value) MinimumValue() coreglib.Value {
	var _arg0 *C.AtkValue // out
	var _arg1 C.GValue    // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(coreglib.InternObject(obj).Native()))

	C.atk_value_get_minimum_value(_arg0, &_arg1)
	runtime.KeepAlive(obj)

	var _value coreglib.Value // out

	_value = *coreglib.ValueFromNative(unsafe.Pointer((&_arg1)))

	return _value
}

// Range gets the range of this object.
//
// The function returns the following values:
//
//    - _range (optional): newly allocated Range that represents the minimum,
//      maximum and descriptor (if available) of obj. NULL if that range is not
//      defined.
//
func (obj *Value) Range() *Range {
	var _arg0 *C.AtkValue // out
	var _cret *C.AtkRange // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(coreglib.InternObject(obj).Native()))

	_cret = C.atk_value_get_range(_arg0)
	runtime.KeepAlive(obj)

	var __range *Range // out

	if _cret != nil {
		__range = (*Range)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(__range)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.atk_range_free((*C.AtkRange)(intern.C))
			},
		)
	}

	return __range
}

// SubRanges gets the list of subranges defined for this object. See Value
// introduction for examples of subranges and when to expose them.
//
// The function returns the following values:
//
//    - sList of Range which each of the subranges defined for this object. Free
//      the returns list with g_slist_free().
//
func (obj *Value) SubRanges() []*Range {
	var _arg0 *C.AtkValue // out
	var _cret *C.GSList   // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(coreglib.InternObject(obj).Native()))

	_cret = C.atk_value_get_sub_ranges(_arg0)
	runtime.KeepAlive(obj)

	var _sList []*Range // out

	_sList = make([]*Range, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.AtkRange)(v)
		var dst *Range // out
		dst = (*Range)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.atk_range_free((*C.AtkRange)(intern.C))
			},
		)
		_sList = append(_sList, dst)
	})

	return _sList
}

// ValueAndText gets the current value and the human readable text alternative
// of obj. text is a newly created string, that must be freed by the caller. Can
// be NULL if no descriptor is available.
//
// The function returns the following values:
//
//    - value address of #gdouble to put the current value of obj.
//    - text (optional) address of #gchar to put the human readable text
//      alternative for value.
//
func (obj *Value) ValueAndText() (float64, string) {
	var _arg0 *C.AtkValue // out
	var _arg1 C.gdouble   // in
	var _arg2 *C.gchar    // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(coreglib.InternObject(obj).Native()))

	C.atk_value_get_value_and_text(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(obj)

	var _value float64 // out
	var _text string   // out

	_value = float64(_arg1)
	if _arg2 != nil {
		_text = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	return _value, _text
}

// SetCurrentValue sets the value of this object.
//
// Deprecated: Since 2.12. Use atk_value_set_value() instead.
//
// The function takes the following parameters:
//
//    - value which is the desired new accessible value.
//
// The function returns the following values:
//
//    - ok: TRUE if new value is successfully set, FALSE otherwise.
//
func (obj *Value) SetCurrentValue(value *coreglib.Value) bool {
	var _arg0 *C.AtkValue // out
	var _arg1 *C.GValue   // out
	var _cret C.gboolean  // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(coreglib.InternObject(obj).Native()))
	_arg1 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C.atk_value_set_current_value(_arg0, _arg1)
	runtime.KeepAlive(obj)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetValue sets the value of this object.
//
// This method is intended to provide a way to change the value of the object.
// In any case, it is possible that the value can't be modified (ie: a read-only
// component). If the value changes due this call, it is possible that the text
// could change, and will trigger an Value::value-changed signal emission.
//
// Note for implementors: the deprecated atk_value_set_current_value() method
// returned TRUE or FALSE depending if the value was assigned or not. In the
// practice several implementors were not able to decide it, and returned TRUE
// in any case. For that reason it is not required anymore to return if the
// value was properly assigned or not.
//
// The function takes the following parameters:
//
//    - newValue: double which is the desired new accessible value.
//
func (obj *Value) SetValue(newValue float64) {
	var _arg0 *C.AtkValue // out
	var _arg1 C.gdouble   // out

	_arg0 = (*C.AtkValue)(unsafe.Pointer(coreglib.InternObject(obj).Native()))
	_arg1 = C.gdouble(newValue)

	C.atk_value_set_value(_arg0, _arg1)
	runtime.KeepAlive(obj)
	runtime.KeepAlive(newValue)
}
