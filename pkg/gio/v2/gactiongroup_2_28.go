// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// GVariant* _gotk4_gio2_ActionGroup_virtual_get_action_state(void* fnptr, GActionGroup* arg0, gchar* arg1) {
//   return ((GVariant* (*)(GActionGroup*, gchar*))(fnptr))(arg0, arg1);
// };
// GVariant* _gotk4_gio2_ActionGroup_virtual_get_action_state_hint(void* fnptr, GActionGroup* arg0, gchar* arg1) {
//   return ((GVariant* (*)(GActionGroup*, gchar*))(fnptr))(arg0, arg1);
// };
// GVariantType* _gotk4_gio2_ActionGroup_virtual_get_action_parameter_type(void* fnptr, GActionGroup* arg0, gchar* arg1) {
//   return ((GVariantType* (*)(GActionGroup*, gchar*))(fnptr))(arg0, arg1);
// };
// GVariantType* _gotk4_gio2_ActionGroup_virtual_get_action_state_type(void* fnptr, GActionGroup* arg0, gchar* arg1) {
//   return ((GVariantType* (*)(GActionGroup*, gchar*))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gio2_ActionGroup_virtual_get_action_enabled(void* fnptr, GActionGroup* arg0, gchar* arg1) {
//   return ((gboolean (*)(GActionGroup*, gchar*))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gio2_ActionGroup_virtual_has_action(void* fnptr, GActionGroup* arg0, gchar* arg1) {
//   return ((gboolean (*)(GActionGroup*, gchar*))(fnptr))(arg0, arg1);
// };
// gchar** _gotk4_gio2_ActionGroup_virtual_list_actions(void* fnptr, GActionGroup* arg0) {
//   return ((gchar** (*)(GActionGroup*))(fnptr))(arg0);
// };
// void _gotk4_gio2_ActionGroup_virtual_action_added(void* fnptr, GActionGroup* arg0, gchar* arg1) {
//   ((void (*)(GActionGroup*, gchar*))(fnptr))(arg0, arg1);
// };
// void _gotk4_gio2_ActionGroup_virtual_action_enabled_changed(void* fnptr, GActionGroup* arg0, gchar* arg1, gboolean arg2) {
//   ((void (*)(GActionGroup*, gchar*, gboolean))(fnptr))(arg0, arg1, arg2);
// };
// void _gotk4_gio2_ActionGroup_virtual_action_removed(void* fnptr, GActionGroup* arg0, gchar* arg1) {
//   ((void (*)(GActionGroup*, gchar*))(fnptr))(arg0, arg1);
// };
// void _gotk4_gio2_ActionGroup_virtual_action_state_changed(void* fnptr, GActionGroup* arg0, gchar* arg1, GVariant* arg2) {
//   ((void (*)(GActionGroup*, gchar*, GVariant*))(fnptr))(arg0, arg1, arg2);
// };
// void _gotk4_gio2_ActionGroup_virtual_activate_action(void* fnptr, GActionGroup* arg0, gchar* arg1, GVariant* arg2) {
//   ((void (*)(GActionGroup*, gchar*, GVariant*))(fnptr))(arg0, arg1, arg2);
// };
// void _gotk4_gio2_ActionGroup_virtual_change_action_state(void* fnptr, GActionGroup* arg0, gchar* arg1, GVariant* arg2) {
//   ((void (*)(GActionGroup*, gchar*, GVariant*))(fnptr))(arg0, arg1, arg2);
// };
import "C"

// ActionAdded emits the Group::action-added signal on action_group.
//
// This function should only be called by Group implementations.
//
// The function takes the following parameters:
//
//    - actionName: name of an action in the group.
//
func (actionGroup *ActionGroup) ActionAdded(actionName string) {
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_action_group_action_added(_arg0, _arg1)
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
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	if enabled {
		_arg2 = C.TRUE
	}

	C.g_action_group_action_enabled_changed(_arg0, _arg1, _arg2)
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
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_action_group_action_removed(_arg0, _arg1)
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
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.GVariant     // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(state)))

	C.g_action_group_action_state_changed(_arg0, _arg1, _arg2)
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
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.GVariant     // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	if parameter != nil {
		_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(parameter)))
	}

	C.g_action_group_activate_action(_arg0, _arg1, _arg2)
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
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.GVariant     // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))

	C.g_action_group_change_action_state(_arg0, _arg1, _arg2)
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
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_action_group_get_action_enabled(_arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

	var _ok bool // out

	if _cret != 0 {
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
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret *C.GVariantType // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_action_group_get_action_parameter_type(_arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

	var _variantType *glib.VariantType // out

	if _cret != nil {
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
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_action_group_get_action_state(_arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

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
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_action_group_get_action_state_hint(_arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

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
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret *C.GVariantType // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_action_group_get_action_state_type(_arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

	var _variantType *glib.VariantType // out

	if _cret != nil {
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
	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_action_group_has_action(_arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

	var _ok bool // out

	if _cret != 0 {
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
	var _arg0 *C.GActionGroup // out
	var _cret **C.gchar       // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))

	_cret = C.g_action_group_list_actions(_arg0)
	runtime.KeepAlive(actionGroup)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
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

// actionAdded emits the Group::action-added signal on action_group.
//
// This function should only be called by Group implementations.
//
// The function takes the following parameters:
//
//    - actionName: name of an action in the group.
//
func (actionGroup *ActionGroup) actionAdded(actionName string) {
	gclass := (*C.GActionGroupInterface)(coreglib.PeekParentClass(actionGroup))
	fnarg := gclass.action_added

	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	C._gotk4_gio2_ActionGroup_virtual_action_added(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)
}

// actionEnabledChanged emits the Group::action-enabled-changed signal on
// action_group.
//
// This function should only be called by Group implementations.
//
// The function takes the following parameters:
//
//    - actionName: name of an action in the group.
//    - enabled: whether or not the action is now enabled.
//
func (actionGroup *ActionGroup) actionEnabledChanged(actionName string, enabled bool) {
	gclass := (*C.GActionGroupInterface)(coreglib.PeekParentClass(actionGroup))
	fnarg := gclass.action_enabled_changed

	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	if enabled {
		_arg2 = C.TRUE
	}

	C._gotk4_gio2_ActionGroup_virtual_action_enabled_changed(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(enabled)
}

// actionRemoved emits the Group::action-removed signal on action_group.
//
// This function should only be called by Group implementations.
//
// The function takes the following parameters:
//
//    - actionName: name of an action in the group.
//
func (actionGroup *ActionGroup) actionRemoved(actionName string) {
	gclass := (*C.GActionGroupInterface)(coreglib.PeekParentClass(actionGroup))
	fnarg := gclass.action_removed

	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	C._gotk4_gio2_ActionGroup_virtual_action_removed(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)
}

// actionStateChanged emits the Group::action-state-changed signal on
// action_group.
//
// This function should only be called by Group implementations.
//
// The function takes the following parameters:
//
//    - actionName: name of an action in the group.
//    - state: new state of the named action.
//
func (actionGroup *ActionGroup) actionStateChanged(actionName string, state *glib.Variant) {
	gclass := (*C.GActionGroupInterface)(coreglib.PeekParentClass(actionGroup))
	fnarg := gclass.action_state_changed

	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.GVariant     // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(state)))

	C._gotk4_gio2_ActionGroup_virtual_action_state_changed(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(state)
}

// activateAction: activate the named action within action_group.
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
func (actionGroup *ActionGroup) activateAction(actionName string, parameter *glib.Variant) {
	gclass := (*C.GActionGroupInterface)(coreglib.PeekParentClass(actionGroup))
	fnarg := gclass.activate_action

	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.GVariant     // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	if parameter != nil {
		_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(parameter)))
	}

	C._gotk4_gio2_ActionGroup_virtual_activate_action(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(parameter)
}

// changeActionState: request for the state of the named action within
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
func (actionGroup *ActionGroup) changeActionState(actionName string, value *glib.Variant) {
	gclass := (*C.GActionGroupInterface)(coreglib.PeekParentClass(actionGroup))
	fnarg := gclass.change_action_state

	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.GVariant     // out

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))

	C._gotk4_gio2_ActionGroup_virtual_change_action_state(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(value)
}

// actionEnabled checks if the named action within action_group is currently
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
func (actionGroup *ActionGroup) actionEnabled(actionName string) bool {
	gclass := (*C.GActionGroupInterface)(coreglib.PeekParentClass(actionGroup))
	fnarg := gclass.get_action_enabled

	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C._gotk4_gio2_ActionGroup_virtual_get_action_enabled(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// actionParameterType queries the type of the parameter that must be given when
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
func (actionGroup *ActionGroup) actionParameterType(actionName string) *glib.VariantType {
	gclass := (*C.GActionGroupInterface)(coreglib.PeekParentClass(actionGroup))
	fnarg := gclass.get_action_parameter_type

	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret *C.GVariantType // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C._gotk4_gio2_ActionGroup_virtual_get_action_parameter_type(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

	var _variantType *glib.VariantType // out

	if _cret != nil {
		_variantType = (*glib.VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _variantType
}

// actionState queries the current state of the named action within
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
func (actionGroup *ActionGroup) actionState(actionName string) *glib.Variant {
	gclass := (*C.GActionGroupInterface)(coreglib.PeekParentClass(actionGroup))
	fnarg := gclass.get_action_state

	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C._gotk4_gio2_ActionGroup_virtual_get_action_state(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

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

// actionStateHint requests a hint about the valid range of values for the state
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
func (actionGroup *ActionGroup) actionStateHint(actionName string) *glib.Variant {
	gclass := (*C.GActionGroupInterface)(coreglib.PeekParentClass(actionGroup))
	fnarg := gclass.get_action_state_hint

	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C._gotk4_gio2_ActionGroup_virtual_get_action_state_hint(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

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

// actionStateType queries the type of the state of the named action within
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
func (actionGroup *ActionGroup) actionStateType(actionName string) *glib.VariantType {
	gclass := (*C.GActionGroupInterface)(coreglib.PeekParentClass(actionGroup))
	fnarg := gclass.get_action_state_type

	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret *C.GVariantType // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C._gotk4_gio2_ActionGroup_virtual_get_action_state_type(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

	var _variantType *glib.VariantType // out

	if _cret != nil {
		_variantType = (*glib.VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _variantType
}

// hasAction checks if the named action exists within action_group.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to check for.
//
// The function returns the following values:
//
//    - ok: whether the named action exists.
//
func (actionGroup *ActionGroup) hasAction(actionName string) bool {
	gclass := (*C.GActionGroupInterface)(coreglib.PeekParentClass(actionGroup))
	fnarg := gclass.has_action

	var _arg0 *C.GActionGroup // out
	var _arg1 *C.gchar        // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C._gotk4_gio2_ActionGroup_virtual_has_action(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// listActions lists the actions contained within action_group.
//
// The caller is responsible for freeing the list with g_strfreev() when it is
// no longer required.
//
// The function returns the following values:
//
//    - utf8s: NULL-terminated array of the names of the actions in the group.
//
func (actionGroup *ActionGroup) listActions() []string {
	gclass := (*C.GActionGroupInterface)(coreglib.PeekParentClass(actionGroup))
	fnarg := gclass.list_actions

	var _arg0 *C.GActionGroup // out
	var _cret **C.gchar       // in

	_arg0 = (*C.GActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))

	_cret = C._gotk4_gio2_ActionGroup_virtual_list_actions(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(actionGroup)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
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

// ActionGroupInterface: virtual function table for Group.
//
// An instance of this type is always passed by reference.
type ActionGroupInterface struct {
	*actionGroupInterface
}

// actionGroupInterface is the struct that's finalized.
type actionGroupInterface struct {
	native *C.GActionGroupInterface
}
