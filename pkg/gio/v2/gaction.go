// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_action_get_type()), F: marshalActioner},
	})
}

// ActionOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ActionOverrider interface {
	// Activate activates the action.
	//
	// parameter must be the correct type of parameter for the action (ie: the
	// parameter type given at construction time). If the parameter type was
	// NULL then parameter must also be NULL.
	//
	// If the parameter GVariant is floating, it is consumed.
	Activate(parameter *glib.Variant)
	// ChangeState: request for the state of action to be changed to value.
	//
	// The action must be stateful and value must be of the correct type. See
	// g_action_get_state_type().
	//
	// This call merely requests a change. The action may refuse to change its
	// state or may change its state to something other than value. See
	// g_action_get_state_hint().
	//
	// If the value GVariant is floating, it is consumed.
	ChangeState(value *glib.Variant)
	// Enabled checks if action is currently enabled.
	//
	// An action must be enabled in order to be activated or in order to have
	// its state changed from outside callers.
	Enabled() bool
	// Name queries the name of action.
	Name() string
	// ParameterType queries the type of the parameter that must be given when
	// activating action.
	//
	// When activating the action using g_action_activate(), the #GVariant given
	// to that function must be of the type returned by this function.
	//
	// In the case that this function returns NULL, you must not give any
	// #GVariant, but NULL instead.
	ParameterType() *glib.VariantType
	// State queries the current state of action.
	//
	// If the action is not stateful then NULL will be returned. If the action
	// is stateful then the type of the return value is the type given by
	// g_action_get_state_type().
	//
	// The return value (if non-NULL) should be freed with g_variant_unref()
	// when it is no longer required.
	State() *glib.Variant
	// StateHint requests a hint about the valid range of values for the state
	// of action.
	//
	// If NULL is returned it either means that the action is not stateful or
	// that there is no hint about the valid range of values for the state of
	// the action.
	//
	// If a #GVariant array is returned then each item in the array is a
	// possible value for the state. If a #GVariant pair (ie: two-tuple) is
	// returned then the tuple specifies the inclusive lower and upper bound of
	// valid values for the state.
	//
	// In any case, the information is merely a hint. It may be possible to have
	// a state value outside of the hinted range and setting a value within the
	// range may fail.
	//
	// The return value (if non-NULL) should be freed with g_variant_unref()
	// when it is no longer required.
	StateHint() *glib.Variant
	// StateType queries the type of the state of action.
	//
	// If the action is stateful (e.g. created with
	// g_simple_action_new_stateful()) then this function returns the Type of
	// the state. This is the type of the initial value given as the state. All
	// calls to g_action_change_state() must give a #GVariant of this type and
	// g_action_get_state() will return a #GVariant of the same type.
	//
	// If the action is not stateful (e.g. created with g_simple_action_new())
	// then this function will return NULL. In that case, g_action_get_state()
	// will return NULL and you must not call g_action_change_state().
	StateType() *glib.VariantType
}

// Action represents a single named action.
//
// The main interface to an action is that it can be activated with
// g_action_activate(). This results in the 'activate' signal being emitted. An
// activation has a #GVariant parameter (which may be NULL). The correct type
// for the parameter is determined by a static parameter type (which is given at
// construction time).
//
// An action may optionally have a state, in which case the state may be set
// with g_action_change_state(). This call takes a #GVariant. The correct type
// for the state is determined by a static state type (which is given at
// construction time).
//
// The state may have a hint associated with it, specifying its valid range.
//
// #GAction is merely the interface to the concept of an action, as described
// above. Various implementations of actions exist, including Action.
//
// In all cases, the implementing class is responsible for storing the name of
// the action, the parameter type, the enabled state, the optional state type
// and the state and emitting the appropriate signals when these change. The
// implementor is responsible for filtering calls to g_action_activate() and
// g_action_change_state() for type safety and for the state being enabled.
//
// Probably the only useful thing to do with a #GAction is to put it inside of a
// ActionGroup.
type Action struct {
	*externglib.Object
}

// Actioner describes Action's interface methods.
type Actioner interface {
	externglib.Objector

	// Activate activates the action.
	Activate(parameter *glib.Variant)
	// ChangeState: request for the state of action to be changed to value.
	ChangeState(value *glib.Variant)
	// Enabled checks if action is currently enabled.
	Enabled() bool
	// Name queries the name of action.
	Name() string
	// ParameterType queries the type of the parameter that must be given when
	// activating action.
	ParameterType() *glib.VariantType
	// State queries the current state of action.
	State() *glib.Variant
	// StateHint requests a hint about the valid range of values for the state
	// of action.
	StateHint() *glib.Variant
	// StateType queries the type of the state of action.
	StateType() *glib.VariantType
}

var _ Actioner = (*Action)(nil)

func wrapAction(obj *externglib.Object) *Action {
	return &Action{
		Object: obj,
	}
}

func marshalActioner(p uintptr) (interface{}, error) {
	return wrapAction(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Activate activates the action.
//
// parameter must be the correct type of parameter for the action (ie: the
// parameter type given at construction time). If the parameter type was NULL
// then parameter must also be NULL.
//
// If the parameter GVariant is floating, it is consumed.
//
// The function takes the following parameters:
//
//    - parameter to the activation.
//
func (action *Action) Activate(parameter *glib.Variant) {
	var _arg0 *C.GAction  // out
	var _arg1 *C.GVariant // out

	_arg0 = (*C.GAction)(unsafe.Pointer(action.Native()))
	if parameter != nil {
		_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(parameter)))
	}

	C.g_action_activate(_arg0, _arg1)
	runtime.KeepAlive(action)
	runtime.KeepAlive(parameter)
}

// ChangeState: request for the state of action to be changed to value.
//
// The action must be stateful and value must be of the correct type. See
// g_action_get_state_type().
//
// This call merely requests a change. The action may refuse to change its state
// or may change its state to something other than value. See
// g_action_get_state_hint().
//
// If the value GVariant is floating, it is consumed.
//
// The function takes the following parameters:
//
//    - value: new state.
//
func (action *Action) ChangeState(value *glib.Variant) {
	var _arg0 *C.GAction  // out
	var _arg1 *C.GVariant // out

	_arg0 = (*C.GAction)(unsafe.Pointer(action.Native()))
	_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))

	C.g_action_change_state(_arg0, _arg1)
	runtime.KeepAlive(action)
	runtime.KeepAlive(value)
}

// Enabled checks if action is currently enabled.
//
// An action must be enabled in order to be activated or in order to have its
// state changed from outside callers.
func (action *Action) Enabled() bool {
	var _arg0 *C.GAction // out
	var _cret C.gboolean // in

	_arg0 = (*C.GAction)(unsafe.Pointer(action.Native()))

	_cret = C.g_action_get_enabled(_arg0)
	runtime.KeepAlive(action)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Name queries the name of action.
func (action *Action) Name() string {
	var _arg0 *C.GAction // out
	var _cret *C.gchar   // in

	_arg0 = (*C.GAction)(unsafe.Pointer(action.Native()))

	_cret = C.g_action_get_name(_arg0)
	runtime.KeepAlive(action)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ParameterType queries the type of the parameter that must be given when
// activating action.
//
// When activating the action using g_action_activate(), the #GVariant given to
// that function must be of the type returned by this function.
//
// In the case that this function returns NULL, you must not give any #GVariant,
// but NULL instead.
func (action *Action) ParameterType() *glib.VariantType {
	var _arg0 *C.GAction      // out
	var _cret *C.GVariantType // in

	_arg0 = (*C.GAction)(unsafe.Pointer(action.Native()))

	_cret = C.g_action_get_parameter_type(_arg0)
	runtime.KeepAlive(action)

	var _variantType *glib.VariantType // out

	if _cret != nil {
		_variantType = (*glib.VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _variantType
}

// State queries the current state of action.
//
// If the action is not stateful then NULL will be returned. If the action is
// stateful then the type of the return value is the type given by
// g_action_get_state_type().
//
// The return value (if non-NULL) should be freed with g_variant_unref() when it
// is no longer required.
func (action *Action) State() *glib.Variant {
	var _arg0 *C.GAction  // out
	var _cret *C.GVariant // in

	_arg0 = (*C.GAction)(unsafe.Pointer(action.Native()))

	_cret = C.g_action_get_state(_arg0)
	runtime.KeepAlive(action)

	var _variant *glib.Variant // out

	if _cret != nil {
		_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_variant)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}

	return _variant
}

// StateHint requests a hint about the valid range of values for the state of
// action.
//
// If NULL is returned it either means that the action is not stateful or that
// there is no hint about the valid range of values for the state of the action.
//
// If a #GVariant array is returned then each item in the array is a possible
// value for the state. If a #GVariant pair (ie: two-tuple) is returned then the
// tuple specifies the inclusive lower and upper bound of valid values for the
// state.
//
// In any case, the information is merely a hint. It may be possible to have a
// state value outside of the hinted range and setting a value within the range
// may fail.
//
// The return value (if non-NULL) should be freed with g_variant_unref() when it
// is no longer required.
func (action *Action) StateHint() *glib.Variant {
	var _arg0 *C.GAction  // out
	var _cret *C.GVariant // in

	_arg0 = (*C.GAction)(unsafe.Pointer(action.Native()))

	_cret = C.g_action_get_state_hint(_arg0)
	runtime.KeepAlive(action)

	var _variant *glib.Variant // out

	if _cret != nil {
		_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_variant)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}

	return _variant
}

// StateType queries the type of the state of action.
//
// If the action is stateful (e.g. created with g_simple_action_new_stateful())
// then this function returns the Type of the state. This is the type of the
// initial value given as the state. All calls to g_action_change_state() must
// give a #GVariant of this type and g_action_get_state() will return a
// #GVariant of the same type.
//
// If the action is not stateful (e.g. created with g_simple_action_new()) then
// this function will return NULL. In that case, g_action_get_state() will
// return NULL and you must not call g_action_change_state().
func (action *Action) StateType() *glib.VariantType {
	var _arg0 *C.GAction      // out
	var _cret *C.GVariantType // in

	_arg0 = (*C.GAction)(unsafe.Pointer(action.Native()))

	_cret = C.g_action_get_state_type(_arg0)
	runtime.KeepAlive(action)

	var _variantType *glib.VariantType // out

	if _cret != nil {
		_variantType = (*glib.VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _variantType
}

// ActionNameIsValid checks if action_name is valid.
//
// action_name is valid if it consists only of alphanumeric characters, plus '-'
// and '.'. The empty string is not a valid action name.
//
// It is an error to call this function with a non-utf8 action_name. action_name
// must not be NULL.
//
// The function takes the following parameters:
//
//    - actionName: potential action name.
//
func ActionNameIsValid(actionName string) bool {
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_action_name_is_valid(_arg1)
	runtime.KeepAlive(actionName)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ActionPrintDetailedName formats a detailed action name from action_name and
// target_value.
//
// It is an error to call this function with an invalid action name.
//
// This function is the opposite of g_action_parse_detailed_name(). It will
// produce a string that can be parsed back to the action_name and target_value
// by that function.
//
// See that function for the types of strings that will be printed by this
// function.
//
// The function takes the following parameters:
//
//    - actionName: valid action name.
//    - targetValue target value, or NULL.
//
func ActionPrintDetailedName(actionName string, targetValue *glib.Variant) string {
	var _arg1 *C.gchar    // out
	var _arg2 *C.GVariant // out
	var _cret *C.gchar    // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	if targetValue != nil {
		_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(targetValue)))
	}

	_cret = C.g_action_print_detailed_name(_arg1, _arg2)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(targetValue)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
