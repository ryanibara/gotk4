// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_action_group_get_type()), F: marshalActionGrouper},
	})
}

// ActionGroupOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ActionGroupOverrider interface {
	// ActionAdded emits the Group::action-added signal on @action_group.
	//
	// This function should only be called by Group implementations.
	ActionAdded(actionName string)
	// ActionEnabledChanged emits the Group::action-enabled-changed signal on
	// @action_group.
	//
	// This function should only be called by Group implementations.
	ActionEnabledChanged(actionName string, enabled bool)
	// ActionRemoved emits the Group::action-removed signal on @action_group.
	//
	// This function should only be called by Group implementations.
	ActionRemoved(actionName string)
	// ActionStateChanged emits the Group::action-state-changed signal on
	// @action_group.
	//
	// This function should only be called by Group implementations.
	ActionStateChanged(actionName string, state *glib.Variant)
	// ActivateAction: activate the named action within @action_group.
	//
	// If the action is expecting a parameter, then the correct type of
	// parameter must be given as @parameter. If the action is expecting no
	// parameters then @parameter must be nil. See
	// g_action_group_get_action_parameter_type().
	//
	// If the Group implementation supports asynchronous remote activation over
	// D-Bus, this call may return before the relevant D-Bus traffic has been
	// sent, or any replies have been received. In order to block on such
	// asynchronous activation calls, g_dbus_connection_flush() should be called
	// prior to the code, which depends on the result of the action activation.
	// Without flushing the D-Bus connection, there is no guarantee that the
	// action would have been activated.
	//
	// The following code which runs in a remote app instance, shows an example
	// of a "quit" action being activated on the primary app instance over
	// D-Bus. Here g_dbus_connection_flush() is called before `exit()`. Without
	// g_dbus_connection_flush(), the "quit" action may fail to be activated on
	// the primary instance.
	//
	//    // call "quit" action on primary instance
	//    g_action_group_activate_action (G_ACTION_GROUP (app), "quit", NULL);
	//
	//    // make sure the action is activated now
	//    g_dbus_connection_flush (...);
	//
	//    g_debug ("application has been terminated. exiting.");
	//
	//    exit (0);
	ActivateAction(actionName string, parameter *glib.Variant)
	// ChangeActionState: request for the state of the named action within
	// @action_group to be changed to @value.
	//
	// The action must be stateful and @value must be of the correct type. See
	// g_action_group_get_action_state_type().
	//
	// This call merely requests a change. The action may refuse to change its
	// state or may change its state to something other than @value. See
	// g_action_group_get_action_state_hint().
	//
	// If the @value GVariant is floating, it is consumed.
	ChangeActionState(actionName string, value *glib.Variant)
	// ActionEnabled checks if the named action within @action_group is
	// currently enabled.
	//
	// An action must be enabled in order to be activated or in order to have
	// its state changed from outside callers.
	ActionEnabled(actionName string) bool
	// ActionParameterType queries the type of the parameter that must be given
	// when activating the named action within @action_group.
	//
	// When activating the action using g_action_group_activate_action(), the
	// #GVariant given to that function must be of the type returned by this
	// function.
	//
	// In the case that this function returns nil, you must not give any
	// #GVariant, but nil instead.
	//
	// The parameter type of a particular action will never change but it is
	// possible for an action to be removed and for a new action to be added
	// with the same name but a different parameter type.
	ActionParameterType(actionName string) *glib.VariantType
	// ActionState queries the current state of the named action within
	// @action_group.
	//
	// If the action is not stateful then nil will be returned. If the action is
	// stateful then the type of the return value is the type given by
	// g_action_group_get_action_state_type().
	//
	// The return value (if non-nil) should be freed with g_variant_unref() when
	// it is no longer required.
	ActionState(actionName string) *glib.Variant
	// ActionStateHint requests a hint about the valid range of values for the
	// state of the named action within @action_group.
	//
	// If nil is returned it either means that the action is not stateful or
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
	// The return value (if non-nil) should be freed with g_variant_unref() when
	// it is no longer required.
	ActionStateHint(actionName string) *glib.Variant
	// ActionStateType queries the type of the state of the named action within
	// @action_group.
	//
	// If the action is stateful then this function returns the Type of the
	// state. All calls to g_action_group_change_action_state() must give a
	// #GVariant of this type and g_action_group_get_action_state() will return
	// a #GVariant of the same type.
	//
	// If the action is not stateful then this function will return nil. In that
	// case, g_action_group_get_action_state() will return nil and you must not
	// call g_action_group_change_action_state().
	//
	// The state type of a particular action will never change but it is
	// possible for an action to be removed and for a new action to be added
	// with the same name but a different state type.
	ActionStateType(actionName string) *glib.VariantType
	// HasAction checks if the named action exists within @action_group.
	HasAction(actionName string) bool
	// ListActions lists the actions contained within @action_group.
	//
	// The caller is responsible for freeing the list with g_strfreev() when it
	// is no longer required.
	ListActions() []string
	// QueryAction queries all aspects of the named action within an
	// @action_group.
	//
	// This function acquires the information available from
	// g_action_group_has_action(), g_action_group_get_action_enabled(),
	// g_action_group_get_action_parameter_type(),
	// g_action_group_get_action_state_type(),
	// g_action_group_get_action_state_hint() and
	// g_action_group_get_action_state() with a single function call.
	//
	// This provides two main benefits.
	//
	// The first is the improvement in efficiency that comes with not having to
	// perform repeated lookups of the action in order to discover different
	// things about it. The second is that implementing Group can now be done by
	// only overriding this one virtual function.
	//
	// The interface provides a default implementation of this function that
	// calls the individual functions, as required, to fetch the information.
	// The interface also provides default implementations of those functions
	// that call this function. All implementations, therefore, must override
	// either this function or all of the others.
	//
	// If the action exists, true is returned and any of the requested fields
	// (as indicated by having a non-nil reference passed in) are filled. If the
	// action doesn't exist, false is returned and the fields may or may not
	// have been modified.
	QueryAction(actionName string) (enabled bool, parameterType *glib.VariantType, stateType *glib.VariantType, stateHint *glib.Variant, state *glib.Variant, ok bool)
}

// ActionGrouper describes ActionGroup's methods.
type ActionGrouper interface {
	// ActionAdded emits the Group::action-added signal on @action_group.
	ActionAdded(actionName string)
	// ActionEnabledChanged emits the Group::action-enabled-changed signal on
	// @action_group.
	ActionEnabledChanged(actionName string, enabled bool)
	// ActionRemoved emits the Group::action-removed signal on @action_group.
	ActionRemoved(actionName string)
	// ActionStateChanged emits the Group::action-state-changed signal on
	// @action_group.
	ActionStateChanged(actionName string, state *glib.Variant)
	// ActivateAction: activate the named action within @action_group.
	ActivateAction(actionName string, parameter *glib.Variant)
	// ChangeActionState: request for the state of the named action within
	// @action_group to be changed to @value.
	ChangeActionState(actionName string, value *glib.Variant)
	// ActionEnabled checks if the named action within @action_group is
	// currently enabled.
	ActionEnabled(actionName string) bool
	// ActionParameterType queries the type of the parameter that must be given
	// when activating the named action within @action_group.
	ActionParameterType(actionName string) *glib.VariantType
	// ActionState queries the current state of the named action within
	// @action_group.
	ActionState(actionName string) *glib.Variant
	// ActionStateHint requests a hint about the valid range of values for the
	// state of the named action within @action_group.
	ActionStateHint(actionName string) *glib.Variant
	// ActionStateType queries the type of the state of the named action within
	// @action_group.
	ActionStateType(actionName string) *glib.VariantType
	// HasAction checks if the named action exists within @action_group.
	HasAction(actionName string) bool
	// ListActions lists the actions contained within @action_group.
	ListActions() []string
	// QueryAction queries all aspects of the named action within an
	// @action_group.
	QueryAction(actionName string) (enabled bool, parameterType *glib.VariantType, stateType *glib.VariantType, stateHint *glib.Variant, state *glib.Variant, ok bool)
}

// ActionGroup represents a group of actions. Actions can be used to expose
// functionality in a structured way, either from one part of a program to
// another, or to the outside world. Action groups are often used together with
// a Model that provides additional representation data for displaying the
// actions to the user, e.g. in a menu.
//
// The main way to interact with the actions in a GActionGroup is to activate
// them with g_action_group_activate_action(). Activating an action may require
// a #GVariant parameter. The required type of the parameter can be inquired
// with g_action_group_get_action_parameter_type(). Actions may be disabled, see
// g_action_group_get_action_enabled(). Activating a disabled action has no
// effect.
//
// Actions may optionally have a state in the form of a #GVariant. The current
// state of an action can be inquired with g_action_group_get_action_state().
// Activating a stateful action may change its state, but it is also possible to
// set the state by calling g_action_group_change_action_state().
//
// As typical example, consider a text editing application which has an option
// to change the current font to 'bold'. A good way to represent this would be a
// stateful action, with a boolean state. Activating the action would toggle the
// state.
//
// Each action in the group has a unique name (which is a string). All method
// calls, except g_action_group_list_actions() take the name of an action as an
// argument.
//
// The Group API is meant to be the 'public' API to the action group. The calls
// here are exactly the interaction that 'external forces' (eg: UI, incoming
// D-Bus messages, etc.) are supposed to have with actions. 'Internal' APIs (ie:
// ones meant only to be accessed by the action group implementation) are found
// on subclasses. This is why you will find - for example -
// g_action_group_get_action_enabled() but not an equivalent set() call.
//
// Signals are emitted on the action group in response to state changes on
// individual actions.
//
// Implementations of Group should provide implementations for the virtual
// functions g_action_group_list_actions() and g_action_group_query_action().
// The other virtual functions should not be implemented - their "wrappers" are
// actually implemented with calls to g_action_group_query_action().
type ActionGroup struct {
	*externglib.Object
}

var (
	_ ActionGrouper   = (*ActionGroup)(nil)
	_ gextras.Nativer = (*ActionGroup)(nil)
)

func wrapActionGroup(obj *externglib.Object) *ActionGroup {
	return &ActionGroup{
		Object: obj,
	}
}

func marshalActionGrouper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapActionGroup(obj), nil
}

// ActionAdded emits the Group::action-added signal on @action_group.
//
// This function should only be called by Group implementations.
func (actionGroup *ActionGroup) ActionAdded(actionName string) {
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))

	C.g_action_group_action_added(_arg0, _arg1)
}

// ActionEnabledChanged emits the Group::action-enabled-changed signal on
// @action_group.
//
// This function should only be called by Group implementations.
func (actionGroup *ActionGroup) ActionEnabledChanged(actionName string, enabled bool) {
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	if enabled {
		_arg2 = C.TRUE
	}

	C.g_action_group_action_enabled_changed(_arg0, _arg1, _arg2)
}

// ActionRemoved emits the Group::action-removed signal on @action_group.
//
// This function should only be called by Group implementations.
func (actionGroup *ActionGroup) ActionRemoved(actionName string) {
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))

	C.g_action_group_action_removed(_arg0, _arg1)
}

// ActionStateChanged emits the Group::action-state-changed signal on
// @action_group.
//
// This function should only be called by Group implementations.
func (actionGroup *ActionGroup) ActionStateChanged(actionName string, state *glib.Variant) {
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.GVariant     // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	_arg2 = (*C.GVariant)(unsafe.Pointer(state))

	C.g_action_group_action_state_changed(_arg0, _arg1, _arg2)
}

// ActivateAction: activate the named action within @action_group.
//
// If the action is expecting a parameter, then the correct type of parameter
// must be given as @parameter. If the action is expecting no parameters then
// @parameter must be nil. See g_action_group_get_action_parameter_type().
//
// If the Group implementation supports asynchronous remote activation over
// D-Bus, this call may return before the relevant D-Bus traffic has been sent,
// or any replies have been received. In order to block on such asynchronous
// activation calls, g_dbus_connection_flush() should be called prior to the
// code, which depends on the result of the action activation. Without flushing
// the D-Bus connection, there is no guarantee that the action would have been
// activated.
//
// The following code which runs in a remote app instance, shows an example of a
// "quit" action being activated on the primary app instance over D-Bus. Here
// g_dbus_connection_flush() is called before `exit()`. Without
// g_dbus_connection_flush(), the "quit" action may fail to be activated on the
// primary instance.
//
//    // call "quit" action on primary instance
//    g_action_group_activate_action (G_ACTION_GROUP (app), "quit", NULL);
//
//    // make sure the action is activated now
//    g_dbus_connection_flush (...);
//
//    g_debug ("application has been terminated. exiting.");
//
//    exit (0);
func (actionGroup *ActionGroup) ActivateAction(actionName string, parameter *glib.Variant) {
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.GVariant     // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	_arg2 = (*C.GVariant)(unsafe.Pointer(parameter))

	C.g_action_group_activate_action(_arg0, _arg1, _arg2)
}

// ChangeActionState: request for the state of the named action within
// @action_group to be changed to @value.
//
// The action must be stateful and @value must be of the correct type. See
// g_action_group_get_action_state_type().
//
// This call merely requests a change. The action may refuse to change its state
// or may change its state to something other than @value. See
// g_action_group_get_action_state_hint().
//
// If the @value GVariant is floating, it is consumed.
func (actionGroup *ActionGroup) ChangeActionState(actionName string, value *glib.Variant) {
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.GVariant     // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	_arg2 = (*C.GVariant)(unsafe.Pointer(value))

	C.g_action_group_change_action_state(_arg0, _arg1, _arg2)
}

// ActionEnabled checks if the named action within @action_group is currently
// enabled.
//
// An action must be enabled in order to be activated or in order to have its
// state changed from outside callers.
func (actionGroup *ActionGroup) ActionEnabled(actionName string) bool {
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))

	_cret = C.g_action_group_get_action_enabled(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ActionParameterType queries the type of the parameter that must be given when
// activating the named action within @action_group.
//
// When activating the action using g_action_group_activate_action(), the
// #GVariant given to that function must be of the type returned by this
// function.
//
// In the case that this function returns nil, you must not give any #GVariant,
// but nil instead.
//
// The parameter type of a particular action will never change but it is
// possible for an action to be removed and for a new action to be added with
// the same name but a different parameter type.
func (actionGroup *ActionGroup) ActionParameterType(actionName string) *glib.VariantType {
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret *C.GVariantType // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))

	_cret = C.g_action_group_get_action_parameter_type(_arg0, _arg1)

	var _variantType *glib.VariantType // out

	_variantType = (*glib.VariantType)(unsafe.Pointer(_cret))

	return _variantType
}

// ActionState queries the current state of the named action within
// @action_group.
//
// If the action is not stateful then nil will be returned. If the action is
// stateful then the type of the return value is the type given by
// g_action_group_get_action_state_type().
//
// The return value (if non-nil) should be freed with g_variant_unref() when it
// is no longer required.
func (actionGroup *ActionGroup) ActionState(actionName string) *glib.Variant {
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))

	_cret = C.g_action_group_get_action_state(_arg0, _arg1)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(unsafe.Pointer(v)))
	})

	return _variant
}

// ActionStateHint requests a hint about the valid range of values for the state
// of the named action within @action_group.
//
// If nil is returned it either means that the action is not stateful or that
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
// The return value (if non-nil) should be freed with g_variant_unref() when it
// is no longer required.
func (actionGroup *ActionGroup) ActionStateHint(actionName string) *glib.Variant {
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))

	_cret = C.g_action_group_get_action_state_hint(_arg0, _arg1)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(unsafe.Pointer(v)))
	})

	return _variant
}

// ActionStateType queries the type of the state of the named action within
// @action_group.
//
// If the action is stateful then this function returns the Type of the state.
// All calls to g_action_group_change_action_state() must give a #GVariant of
// this type and g_action_group_get_action_state() will return a #GVariant of
// the same type.
//
// If the action is not stateful then this function will return nil. In that
// case, g_action_group_get_action_state() will return nil and you must not call
// g_action_group_change_action_state().
//
// The state type of a particular action will never change but it is possible
// for an action to be removed and for a new action to be added with the same
// name but a different state type.
func (actionGroup *ActionGroup) ActionStateType(actionName string) *glib.VariantType {
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret *C.GVariantType // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))

	_cret = C.g_action_group_get_action_state_type(_arg0, _arg1)

	var _variantType *glib.VariantType // out

	_variantType = (*glib.VariantType)(unsafe.Pointer(_cret))

	return _variantType
}

// HasAction checks if the named action exists within @action_group.
func (actionGroup *ActionGroup) HasAction(actionName string) bool {
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))

	_cret = C.g_action_group_has_action(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ListActions lists the actions contained within @action_group.
//
// The caller is responsible for freeing the list with g_strfreev() when it is
// no longer required.
func (actionGroup *ActionGroup) ListActions() []string {
	var _arg0 *C.GActionGroup // out
	var _cret **C.gchar

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(actionGroup.Native()))

	_cret = C.g_action_group_list_actions(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// QueryAction queries all aspects of the named action within an @action_group.
//
// This function acquires the information available from
// g_action_group_has_action(), g_action_group_get_action_enabled(),
// g_action_group_get_action_parameter_type(),
// g_action_group_get_action_state_type(),
// g_action_group_get_action_state_hint() and g_action_group_get_action_state()
// with a single function call.
//
// This provides two main benefits.
//
// The first is the improvement in efficiency that comes with not having to
// perform repeated lookups of the action in order to discover different things
// about it. The second is that implementing Group can now be done by only
// overriding this one virtual function.
//
// The interface provides a default implementation of this function that calls
// the individual functions, as required, to fetch the information. The
// interface also provides default implementations of those functions that call
// this function. All implementations, therefore, must override either this
// function or all of the others.
//
// If the action exists, true is returned and any of the requested fields (as
// indicated by having a non-nil reference passed in) are filled. If the action
// doesn't exist, false is returned and the fields may or may not have been
// modified.
func (actionGroup *ActionGroup) QueryAction(actionName string) (enabled bool, parameterType *glib.VariantType, stateType *glib.VariantType, stateHint *glib.Variant, state *glib.Variant, ok bool) {
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _arg2 C.gboolean      // in
	var _parameterType *glib.VariantType
	var _stateType *glib.VariantType
	var _stateHint *glib.Variant
	var _state *glib.Variant
	var _cret C.gboolean // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))

	_cret = C.g_action_group_query_action(_arg0, _arg1, &_arg2, (**C.GVariantType)(unsafe.Pointer(&_parameterType)), (**C.GVariantType)(unsafe.Pointer(&_stateType)), (**C.GVariant)(unsafe.Pointer(&_stateHint)), (**C.GVariant)(unsafe.Pointer(&_state)))

	var _enabled bool // out

	var _ok bool // out

	if _arg2 != 0 {
		_enabled = true
	}

	if _cret != 0 {
		_ok = true
	}

	return _enabled, _parameterType, _stateType, _stateHint, _state, _ok
}
