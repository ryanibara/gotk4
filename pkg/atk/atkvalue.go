// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_value_type_get_type()), F: marshalValueType},
		{T: externglib.Type(C.atk_value_get_type()), F: marshalValueer},
	})
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
	return ValueType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
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
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ValueOverrider interface {
	// CurrentValue gets the value of this object.
	//
	// Deprecated: Since 2.12. Use atk_value_get_value_and_text() instead.
	CurrentValue() externglib.Value
	// Increment gets the minimum increment by which the value of this object
	// may be changed. If zero, the minimum increment is undefined, which may
	// mean that it is limited only by the floating point precision of the
	// platform.
	Increment() float64
	// MaximumValue gets the maximum value of this object.
	//
	// Deprecated: Since 2.12. Use atk_value_get_range() instead.
	MaximumValue() externglib.Value
	// MinimumIncrement gets the minimum increment by which the value of this
	// object may be changed. If zero, the minimum increment is undefined, which
	// may mean that it is limited only by the floating point precision of the
	// platform.
	//
	// Deprecated: Since 2.12. Use atk_value_get_increment() instead.
	MinimumIncrement() externglib.Value
	// MinimumValue gets the minimum value of this object.
	//
	// Deprecated: Since 2.12. Use atk_value_get_range() instead.
	MinimumValue() externglib.Value
	// Range gets the range of this object.
	Range() *Range
	// SubRanges gets the list of subranges defined for this object. See Value
	// introduction for examples of subranges and when to expose them.
	SubRanges() []*Range
	// ValueAndText gets the current value and the human readable text
	// alternative of obj. text is a newly created string, that must be freed by
	// the caller. Can be NULL if no descriptor is available.
	ValueAndText() (float64, string)
	// SetCurrentValue sets the value of this object.
	//
	// Deprecated: Since 2.12. Use atk_value_set_value() instead.
	SetCurrentValue(value *externglib.Value) bool
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
type Value struct {
	*externglib.Object
}

// Valueer describes Value's interface methods.
type Valueer interface {
	externglib.Objector

	// CurrentValue gets the value of this object.
	CurrentValue() externglib.Value
	// Increment gets the minimum increment by which the value of this object
	// may be changed.
	Increment() float64
	// MaximumValue gets the maximum value of this object.
	MaximumValue() externglib.Value
	// MinimumIncrement gets the minimum increment by which the value of this
	// object may be changed.
	MinimumIncrement() externglib.Value
	// MinimumValue gets the minimum value of this object.
	MinimumValue() externglib.Value
	// Range gets the range of this object.
	Range() *Range
	// SubRanges gets the list of subranges defined for this object.
	SubRanges() []*Range
	// ValueAndText gets the current value and the human readable text
	// alternative of obj.
	ValueAndText() (float64, string)
	// SetCurrentValue sets the value of this object.
	SetCurrentValue(value *externglib.Value) bool
	// SetValue sets the value of this object.
	SetValue(newValue float64)
}

var _ Valueer = (*Value)(nil)

func wrapValue(obj *externglib.Object) *Value {
	return &Value{
		Object: obj,
	}
}

func marshalValueer(p uintptr) (interface{}, error) {
	return wrapValue(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CurrentValue gets the value of this object.
//
// Deprecated: Since 2.12. Use atk_value_get_value_and_text() instead.
func (obj *Value) CurrentValue() externglib.Value {
	var _arg0 *C.AtkValue // out
	var _arg1 C.GValue    // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(obj.Native()))

	C.atk_value_get_current_value(_arg0, &_arg1)
	runtime.KeepAlive(obj)

	var _value externglib.Value // out

	_value = *externglib.ValueFromNative(unsafe.Pointer((&_arg1)))

	return _value
}

// Increment gets the minimum increment by which the value of this object may be
// changed. If zero, the minimum increment is undefined, which may mean that it
// is limited only by the floating point precision of the platform.
func (obj *Value) Increment() float64 {
	var _arg0 *C.AtkValue // out
	var _cret C.gdouble   // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(obj.Native()))

	_cret = C.atk_value_get_increment(_arg0)
	runtime.KeepAlive(obj)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// MaximumValue gets the maximum value of this object.
//
// Deprecated: Since 2.12. Use atk_value_get_range() instead.
func (obj *Value) MaximumValue() externglib.Value {
	var _arg0 *C.AtkValue // out
	var _arg1 C.GValue    // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(obj.Native()))

	C.atk_value_get_maximum_value(_arg0, &_arg1)
	runtime.KeepAlive(obj)

	var _value externglib.Value // out

	_value = *externglib.ValueFromNative(unsafe.Pointer((&_arg1)))

	return _value
}

// MinimumIncrement gets the minimum increment by which the value of this object
// may be changed. If zero, the minimum increment is undefined, which may mean
// that it is limited only by the floating point precision of the platform.
//
// Deprecated: Since 2.12. Use atk_value_get_increment() instead.
func (obj *Value) MinimumIncrement() externglib.Value {
	var _arg0 *C.AtkValue // out
	var _arg1 C.GValue    // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(obj.Native()))

	C.atk_value_get_minimum_increment(_arg0, &_arg1)
	runtime.KeepAlive(obj)

	var _value externglib.Value // out

	_value = *externglib.ValueFromNative(unsafe.Pointer((&_arg1)))

	return _value
}

// MinimumValue gets the minimum value of this object.
//
// Deprecated: Since 2.12. Use atk_value_get_range() instead.
func (obj *Value) MinimumValue() externglib.Value {
	var _arg0 *C.AtkValue // out
	var _arg1 C.GValue    // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(obj.Native()))

	C.atk_value_get_minimum_value(_arg0, &_arg1)
	runtime.KeepAlive(obj)

	var _value externglib.Value // out

	_value = *externglib.ValueFromNative(unsafe.Pointer((&_arg1)))

	return _value
}

// Range gets the range of this object.
func (obj *Value) Range() *Range {
	var _arg0 *C.AtkValue // out
	var _cret *C.AtkRange // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(obj.Native()))

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
func (obj *Value) SubRanges() []*Range {
	var _arg0 *C.AtkValue // out
	var _cret *C.GSList   // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(obj.Native()))

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
func (obj *Value) ValueAndText() (float64, string) {
	var _arg0 *C.AtkValue // out
	var _arg1 C.gdouble   // in
	var _arg2 *C.gchar    // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(obj.Native()))

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
func (obj *Value) SetCurrentValue(value *externglib.Value) bool {
	var _arg0 *C.AtkValue // out
	var _arg1 *C.GValue   // out
	var _cret C.gboolean  // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(obj.Native()))
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

	_arg0 = (*C.AtkValue)(unsafe.Pointer(obj.Native()))
	_arg1 = C.gdouble(newValue)

	C.atk_value_set_value(_arg0, _arg1)
	runtime.KeepAlive(obj)
	runtime.KeepAlive(newValue)
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
func (obj *Value) ConnectValueChanged(f func(value float64, text string)) externglib.SignalHandle {
	return obj.Connect("value-changed", f)
}
