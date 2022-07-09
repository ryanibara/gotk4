// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern GVariant* _gotk4_gio2_ActionGroupInterface_get_action_state(void*, void*);
// extern GVariant* _gotk4_gio2_ActionGroupInterface_get_action_state_hint(void*, void*);
// extern GVariantType* _gotk4_gio2_ActionGroupInterface_get_action_parameter_type(void*, void*);
// extern GVariantType* _gotk4_gio2_ActionGroupInterface_get_action_state_type(void*, void*);
// extern gboolean _gotk4_gio2_ActionGroupInterface_get_action_enabled(void*, void*);
// extern gboolean _gotk4_gio2_ActionGroupInterface_has_action(void*, void*);
// extern gchar** _gotk4_gio2_ActionGroupInterface_list_actions(void*);
// extern void _gotk4_gio2_ActionGroupInterface_action_added(void*, void*);
// extern void _gotk4_gio2_ActionGroupInterface_action_enabled_changed(void*, void*, gboolean);
// extern void _gotk4_gio2_ActionGroupInterface_action_removed(void*, void*);
// extern void _gotk4_gio2_ActionGroupInterface_action_state_changed(void*, void*, void*);
// extern void _gotk4_gio2_ActionGroupInterface_activate_action(void*, void*, void*);
// extern void _gotk4_gio2_ActionGroupInterface_change_action_state(void*, void*, void*);
// extern void _gotk4_gio2_ActionGroup_ConnectActionAdded(gpointer, void*, guintptr);
// extern void _gotk4_gio2_ActionGroup_ConnectActionEnabledChanged(gpointer, void*, gboolean, guintptr);
// extern void _gotk4_gio2_ActionGroup_ConnectActionRemoved(gpointer, void*, guintptr);
// extern void _gotk4_gio2_ActionGroup_ConnectActionStateChanged(gpointer, void*, void*, guintptr);
import "C"

// GTypeActionGroup returns the GType for the type ActionGroup.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeActionGroup() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "ActionGroup").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalActionGroup)
	return gtype
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
//
// ActionGroup wraps an interface. This means the user can get the
// underlying type by calling Cast().
type ActionGroup struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ActionGroup)(nil)
)

// ActionGrouper describes ActionGroup's interface methods.
type ActionGrouper interface {
	coreglib.Objector

	// ActionAdded emits the Group::action-added signal on action_group.
	ActionAdded(actionName string)
	// ActionEnabledChanged emits the Group::action-enabled-changed signal on
	// action_group.
	ActionEnabledChanged(actionName string, enabled bool)
	// ActionRemoved emits the Group::action-removed signal on action_group.
	ActionRemoved(actionName string)
	// ActionStateChanged emits the Group::action-state-changed signal on
	// action_group.
	ActionStateChanged(actionName string, state *glib.Variant)
	// ActivateAction: activate the named action within action_group.
	ActivateAction(actionName string, parameter *glib.Variant)
	// ChangeActionState: request for the state of the named action within
	// action_group to be changed to value.
	ChangeActionState(actionName string, value *glib.Variant)
	// ActionEnabled checks if the named action within action_group is currently
	// enabled.
	ActionEnabled(actionName string) bool
	// ActionParameterType queries the type of the parameter that must be given
	// when activating the named action within action_group.
	ActionParameterType(actionName string) *glib.VariantType
	// ActionState queries the current state of the named action within
	// action_group.
	ActionState(actionName string) *glib.Variant
	// ActionStateHint requests a hint about the valid range of values for the
	// state of the named action within action_group.
	ActionStateHint(actionName string) *glib.Variant
	// ActionStateType queries the type of the state of the named action within
	// action_group.
	ActionStateType(actionName string) *glib.VariantType
	// HasAction checks if the named action exists within action_group.
	HasAction(actionName string) bool
	// ListActions lists the actions contained within action_group.
	ListActions() []string
	// QueryAction queries all aspects of the named action within an
	// action_group.
	QueryAction(actionName string) (enabled bool, parameterType, stateType *glib.VariantType, stateHint, state *glib.Variant, ok bool)

	// Action-added signals that a new action was just added to the group.
	ConnectActionAdded(func(actionName string)) coreglib.SignalHandle
	// Action-enabled-changed signals that the enabled status of the named
	// action has changed.
	ConnectActionEnabledChanged(func(actionName string, enabled bool)) coreglib.SignalHandle
	// Action-removed signals that an action is just about to be removed from
	// the group.
	ConnectActionRemoved(func(actionName string)) coreglib.SignalHandle
	// Action-state-changed signals that the state of the named action has
	// changed.
	ConnectActionStateChanged(func(actionName string, value *glib.Variant)) coreglib.SignalHandle
}

var _ ActionGrouper = (*ActionGroup)(nil)

func wrapActionGroup(obj *coreglib.Object) *ActionGroup {
	return &ActionGroup{
		Object: obj,
	}
}

func marshalActionGroup(p uintptr) (interface{}, error) {
	return wrapActionGroup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_ActionGroup_ConnectActionAdded
func _gotk4_gio2_ActionGroup_ConnectActionAdded(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(actionName string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(actionName string))
	}

	var _actionName string // out

	_actionName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_actionName)
}

// ConnectActionAdded signals that a new action was just added to the group.
// This signal is emitted after the action has been added and is now visible.
func (actionGroup *ActionGroup) ConnectActionAdded(f func(actionName string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(actionGroup, "action-added", false, unsafe.Pointer(C._gotk4_gio2_ActionGroup_ConnectActionAdded), f)
}

//export _gotk4_gio2_ActionGroup_ConnectActionEnabledChanged
func _gotk4_gio2_ActionGroup_ConnectActionEnabledChanged(arg0 C.gpointer, arg1 *C.void, arg2 C.gboolean, arg3 C.guintptr) {
	var f func(actionName string, enabled bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(actionName string, enabled bool))
	}

	var _actionName string // out
	var _enabled bool      // out

	_actionName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	if arg2 != 0 {
		_enabled = true
	}

	f(_actionName, _enabled)
}

// ConnectActionEnabledChanged signals that the enabled status of the named
// action has changed.
func (actionGroup *ActionGroup) ConnectActionEnabledChanged(f func(actionName string, enabled bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(actionGroup, "action-enabled-changed", false, unsafe.Pointer(C._gotk4_gio2_ActionGroup_ConnectActionEnabledChanged), f)
}

//export _gotk4_gio2_ActionGroup_ConnectActionRemoved
func _gotk4_gio2_ActionGroup_ConnectActionRemoved(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(actionName string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(actionName string))
	}

	var _actionName string // out

	_actionName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_actionName)
}

// ConnectActionRemoved signals that an action is just about to be removed from
// the group. This signal is emitted before the action is removed, so the action
// is still visible and can be queried from the signal handler.
func (actionGroup *ActionGroup) ConnectActionRemoved(f func(actionName string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(actionGroup, "action-removed", false, unsafe.Pointer(C._gotk4_gio2_ActionGroup_ConnectActionRemoved), f)
}

//export _gotk4_gio2_ActionGroup_ConnectActionStateChanged
func _gotk4_gio2_ActionGroup_ConnectActionStateChanged(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 C.guintptr) {
	var f func(actionName string, value *glib.Variant)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(actionName string, value *glib.Variant))
	}

	var _actionName string   // out
	var _value *glib.Variant // out

	_actionName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_value = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	C.g_variant_ref(arg2)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_value)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	f(_actionName, _value)
}

// ConnectActionStateChanged signals that the state of the named action has
// changed.
func (actionGroup *ActionGroup) ConnectActionStateChanged(f func(actionName string, value *glib.Variant)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(actionGroup, "action-state-changed", false, unsafe.Pointer(C._gotk4_gio2_ActionGroup_ConnectActionStateChanged), f)
}

// ActionAdded emits the Group::action-added signal on action_group.
//
// This function should only be called by Group implementations.
//
// The function takes the following parameters:
//
//    - actionName: name of an action in the group.
//
func (actionGroup *ActionGroup) ActionAdded(actionName string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[1]))

	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)
}

// ActionEnabledChanged emits the Group::action-enabled-changed signal on
// action_group.
//
// This function should only be called by Group implementations.
//
// The function takes the following parameters:
//
//    - actionName: name of an action in the group.
//    - enabled: whether or not the action is now enabled.
//
func (actionGroup *ActionGroup) ActionEnabledChanged(actionName string, enabled bool) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[1]))
	if enabled {
		*(*C.gboolean)(unsafe.Pointer(&_args[2])) = C.TRUE
	}

	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(enabled)
}

// ActionRemoved emits the Group::action-removed signal on action_group.
//
// This function should only be called by Group implementations.
//
// The function takes the following parameters:
//
//    - actionName: name of an action in the group.
//
func (actionGroup *ActionGroup) ActionRemoved(actionName string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[1]))

	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)
}

// ActionStateChanged emits the Group::action-state-changed signal on
// action_group.
//
// This function should only be called by Group implementations.
//
// The function takes the following parameters:
//
//    - actionName: name of an action in the group.
//    - state: new state of the named action.
//
func (actionGroup *ActionGroup) ActionStateChanged(actionName string, state *glib.Variant) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(state)))

	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(state)
}

// ActivateAction: activate the named action within action_group.
//
// If the action is expecting a parameter, then the correct type of parameter
// must be given as parameter. If the action is expecting no parameters then
// parameter must be NULL. See g_action_group_get_action_parameter_type().
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
// g_dbus_connection_flush() is called before exit(). Without
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
//    exit (0);.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to activate.
//    - parameter (optional) parameters to the activation.
//
func (actionGroup *ActionGroup) ActivateAction(actionName string, parameter *glib.Variant) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[1]))
	if parameter != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(parameter)))
	}

	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(parameter)
}

// ChangeActionState: request for the state of the named action within
// action_group to be changed to value.
//
// The action must be stateful and value must be of the correct type. See
// g_action_group_get_action_state_type().
//
// This call merely requests a change. The action may refuse to change its state
// or may change its state to something other than value. See
// g_action_group_get_action_state_hint().
//
// If the value GVariant is floating, it is consumed.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to request the change on.
//    - value: new state.
//
func (actionGroup *ActionGroup) ChangeActionState(actionName string, value *glib.Variant) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(value)))

	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(value)
}

// ActionEnabled checks if the named action within action_group is currently
// enabled.
//
// An action must be enabled in order to be activated or in order to have its
// state changed from outside callers.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to query.
//
// The function returns the following values:
//
//    - ok: whether or not the action is currently enabled.
//
func (actionGroup *ActionGroup) ActionEnabled(actionName string) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[1]))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ActionParameterType queries the type of the parameter that must be given when
// activating the named action within action_group.
//
// When activating the action using g_action_group_activate_action(), the
// #GVariant given to that function must be of the type returned by this
// function.
//
// In the case that this function returns NULL, you must not give any #GVariant,
// but NULL instead.
//
// The parameter type of a particular action will never change but it is
// possible for an action to be removed and for a new action to be added with
// the same name but a different parameter type.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to query.
//
// The function returns the following values:
//
//    - variantType (optional): parameter type.
//
func (actionGroup *ActionGroup) ActionParameterType(actionName string) *glib.VariantType {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[1]))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

	var _variantType *glib.VariantType // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_variantType = (*glib.VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _variantType
}

// ActionState queries the current state of the named action within
// action_group.
//
// If the action is not stateful then NULL will be returned. If the action is
// stateful then the type of the return value is the type given by
// g_action_group_get_action_state_type().
//
// The return value (if non-NULL) should be freed with g_variant_unref() when it
// is no longer required.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to query.
//
// The function returns the following values:
//
//    - variant (optional): current state of the action.
//
func (actionGroup *ActionGroup) ActionState(actionName string) *glib.Variant {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[1]))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

	var _variant *glib.Variant // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_variant)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _variant
}

// ActionStateHint requests a hint about the valid range of values for the state
// of the named action within action_group.
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
//
// The function takes the following parameters:
//
//    - actionName: name of the action to query.
//
// The function returns the following values:
//
//    - variant (optional): state range hint.
//
func (actionGroup *ActionGroup) ActionStateHint(actionName string) *glib.Variant {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[1]))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

	var _variant *glib.Variant // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_variant)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _variant
}

// ActionStateType queries the type of the state of the named action within
// action_group.
//
// If the action is stateful then this function returns the Type of the state.
// All calls to g_action_group_change_action_state() must give a #GVariant of
// this type and g_action_group_get_action_state() will return a #GVariant of
// the same type.
//
// If the action is not stateful then this function will return NULL. In that
// case, g_action_group_get_action_state() will return NULL and you must not
// call g_action_group_change_action_state().
//
// The state type of a particular action will never change but it is possible
// for an action to be removed and for a new action to be added with the same
// name but a different state type.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to query.
//
// The function returns the following values:
//
//    - variantType (optional): state type, if the action is stateful.
//
func (actionGroup *ActionGroup) ActionStateType(actionName string) *glib.VariantType {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[1]))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

	var _variantType *glib.VariantType // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_variantType = (*glib.VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _variantType
}

// HasAction checks if the named action exists within action_group.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to check for.
//
// The function returns the following values:
//
//    - ok: whether the named action exists.
//
func (actionGroup *ActionGroup) HasAction(actionName string) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[1]))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ListActions lists the actions contained within action_group.
//
// The caller is responsible for freeing the list with g_strfreev() when it is
// no longer required.
//
// The function returns the following values:
//
//    - utf8s: NULL-terminated array of the names of the actions in the group.
//
func (actionGroup *ActionGroup) ListActions() []string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))

	_cret = *(***C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(actionGroup)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.void
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// QueryAction queries all aspects of the named action within an action_group.
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
// If the action exists, TRUE is returned and any of the requested fields (as
// indicated by having a non-NULL reference passed in) are filled. If the action
// doesn't exist, FALSE is returned and the fields may or may not have been
// modified.
//
// The function takes the following parameters:
//
//    - actionName: name of an action in the group.
//
// The function returns the following values:
//
//    - enabled: if the action is presently enabled.
//    - parameterType (optional): parameter type, or NULL if none needed.
//    - stateType (optional): state type, or NULL if stateless.
//    - stateHint (optional): state hint, or NULL if none.
//    - state (optional): current state, or NULL if stateless.
//    - ok: TRUE if the action exists, else FALSE.
//
func (actionGroup *ActionGroup) QueryAction(actionName string) (enabled bool, parameterType, stateType *glib.VariantType, stateHint, state *glib.Variant, ok bool) {
	var _args [2]girepository.Argument
	var _outs [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[1]))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

	var _enabled bool                    // out
	var _parameterType *glib.VariantType // out
	var _stateType *glib.VariantType     // out
	var _stateHint *glib.Variant         // out
	var _state *glib.Variant             // out
	var _ok bool                         // out

	if **(**C.void)(unsafe.Pointer(&_outs[0])) != 0 {
		_enabled = true
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_parameterType = (*glib.VariantType)(gextras.NewStructNative(unsafe.Pointer(_outs[1])))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_parameterType)),
			func(intern *struct{ C unsafe.Pointer }) {
				{
					args := [1]girepository.Argument{(*C.void)(intern.C)}
					girepository.MustFind("GLib", "VariantType").InvokeMethod("free", args[:], nil)
				}
			},
		)
	}
	if *(**C.void)(unsafe.Pointer(&_outs[2])) != nil {
		_stateType = (*glib.VariantType)(gextras.NewStructNative(unsafe.Pointer(_outs[2])))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_stateType)),
			func(intern *struct{ C unsafe.Pointer }) {
				{
					args := [1]girepository.Argument{(*C.void)(intern.C)}
					girepository.MustFind("GLib", "VariantType").InvokeMethod("free", args[:], nil)
				}
			},
		)
	}
	if *(**C.void)(unsafe.Pointer(&_outs[3])) != nil {
		_stateHint = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_outs[3])))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_stateHint)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}
	if *(**C.void)(unsafe.Pointer(&_outs[4])) != nil {
		_state = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_outs[4])))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_state)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _enabled, _parameterType, _stateType, _stateHint, _state, _ok
}
