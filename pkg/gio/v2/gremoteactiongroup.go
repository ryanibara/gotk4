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
// extern void _gotk4_gio2_RemoteActionGroupInterface_activate_action_full(void*, void*, void*, void*);
// extern void _gotk4_gio2_RemoteActionGroupInterface_change_action_state_full(void*, void*, void*, void*);
import "C"

// GTypeRemoteActionGroup returns the GType for the type RemoteActionGroup.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeRemoteActionGroup() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "RemoteActionGroup").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalRemoteActionGroup)
	return gtype
}

// RemoteActionGroupOverrider contains methods that are overridable.
type RemoteActionGroupOverrider interface {
	// ActivateActionFull activates the remote action.
	//
	// This is the same as g_action_group_activate_action() except that it
	// allows for provision of "platform data" to be sent along with the
	// activation request. This typically contains details such as the user
	// interaction timestamp or startup notification information.
	//
	// platform_data must be non-NULL and must have the type
	// G_VARIANT_TYPE_VARDICT. If it is floating, it will be consumed.
	//
	// The function takes the following parameters:
	//
	//    - actionName: name of the action to activate.
	//    - parameter (optional): optional parameter to the activation.
	//    - platformData: platform data to send.
	//
	ActivateActionFull(actionName string, parameter, platformData *glib.Variant)
	// ChangeActionStateFull changes the state of a remote action.
	//
	// This is the same as g_action_group_change_action_state() except that it
	// allows for provision of "platform data" to be sent along with the state
	// change request. This typically contains details such as the user
	// interaction timestamp or startup notification information.
	//
	// platform_data must be non-NULL and must have the type
	// G_VARIANT_TYPE_VARDICT. If it is floating, it will be consumed.
	//
	// The function takes the following parameters:
	//
	//    - actionName: name of the action to change the state of.
	//    - value: new requested value for the state.
	//    - platformData: platform data to send.
	//
	ChangeActionStateFull(actionName string, value, platformData *glib.Variant)
}

// RemoteActionGroup interface is implemented by Group instances that either
// transmit action invocations to other processes or receive action invocations
// in the local process from other processes.
//
// The interface has _full variants of the two methods on Group used to activate
// actions: g_action_group_activate_action() and
// g_action_group_change_action_state(). These variants allow a "platform data"
// #GVariant to be specified: a dictionary providing context for the action
// invocation (for example: timestamps, startup notification IDs, etc).
//
// BusActionGroup implements ActionGroup. This provides a mechanism to send
// platform data for action invocations over D-Bus.
//
// Additionally, g_dbus_connection_export_action_group() will check if the
// exported Group implements ActionGroup and use the _full variants of the calls
// if available. This provides a mechanism by which to receive platform data for
// action invocations that arrive by way of D-Bus.
//
// RemoteActionGroup wraps an interface. This means the user can get the
// underlying type by calling Cast().
type RemoteActionGroup struct {
	_ [0]func() // equal guard
	ActionGroup
}

var ()

// RemoteActionGrouper describes RemoteActionGroup's interface methods.
type RemoteActionGrouper interface {
	coreglib.Objector

	// ActivateActionFull activates the remote action.
	ActivateActionFull(actionName string, parameter, platformData *glib.Variant)
	// ChangeActionStateFull changes the state of a remote action.
	ChangeActionStateFull(actionName string, value, platformData *glib.Variant)
}

var _ RemoteActionGrouper = (*RemoteActionGroup)(nil)

func ifaceInitRemoteActionGrouper(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Gio", "RemoteActionGroupInterface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("activate_action_full"))) = unsafe.Pointer(C._gotk4_gio2_RemoteActionGroupInterface_activate_action_full)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("change_action_state_full"))) = unsafe.Pointer(C._gotk4_gio2_RemoteActionGroupInterface_change_action_state_full)
}

//export _gotk4_gio2_RemoteActionGroupInterface_activate_action_full
func _gotk4_gio2_RemoteActionGroupInterface_activate_action_full(arg0 *C.void, arg1 *C.void, arg2 *C.void, arg3 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(RemoteActionGroupOverrider)

	var _actionName string          // out
	var _parameter *glib.Variant    // out
	var _platformData *glib.Variant // out

	_actionName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_parameter = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg2)))
		C.g_variant_ref(arg2)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_parameter)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}
	_platformData = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg3)))
	C.g_variant_ref(arg3)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_platformData)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	iface.ActivateActionFull(_actionName, _parameter, _platformData)
}

//export _gotk4_gio2_RemoteActionGroupInterface_change_action_state_full
func _gotk4_gio2_RemoteActionGroupInterface_change_action_state_full(arg0 *C.void, arg1 *C.void, arg2 *C.void, arg3 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(RemoteActionGroupOverrider)

	var _actionName string          // out
	var _value *glib.Variant        // out
	var _platformData *glib.Variant // out

	_actionName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_value = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	C.g_variant_ref(arg2)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_value)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)
	_platformData = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg3)))
	C.g_variant_ref(arg3)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_platformData)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	iface.ChangeActionStateFull(_actionName, _value, _platformData)
}

func wrapRemoteActionGroup(obj *coreglib.Object) *RemoteActionGroup {
	return &RemoteActionGroup{
		ActionGroup: ActionGroup{
			Object: obj,
		},
	}
}

func marshalRemoteActionGroup(p uintptr) (interface{}, error) {
	return wrapRemoteActionGroup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ActivateActionFull activates the remote action.
//
// This is the same as g_action_group_activate_action() except that it allows
// for provision of "platform data" to be sent along with the activation
// request. This typically contains details such as the user interaction
// timestamp or startup notification information.
//
// platform_data must be non-NULL and must have the type G_VARIANT_TYPE_VARDICT.
// If it is floating, it will be consumed.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to activate.
//    - parameter (optional): optional parameter to the activation.
//    - platformData: platform data to send.
//
func (remote *RemoteActionGroup) ActivateActionFull(actionName string, parameter, platformData *glib.Variant) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(remote).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[1]))
	if parameter != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(parameter)))
	}
	*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(gextras.StructNative(unsafe.Pointer(platformData)))

	runtime.KeepAlive(remote)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(parameter)
	runtime.KeepAlive(platformData)
}

// ChangeActionStateFull changes the state of a remote action.
//
// This is the same as g_action_group_change_action_state() except that it
// allows for provision of "platform data" to be sent along with the state
// change request. This typically contains details such as the user interaction
// timestamp or startup notification information.
//
// platform_data must be non-NULL and must have the type G_VARIANT_TYPE_VARDICT.
// If it is floating, it will be consumed.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to change the state of.
//    - value: new requested value for the state.
//    - platformData: platform data to send.
//
func (remote *RemoteActionGroup) ChangeActionStateFull(actionName string, value, platformData *glib.Variant) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(remote).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(value)))
	*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(gextras.StructNative(unsafe.Pointer(platformData)))

	runtime.KeepAlive(remote)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(value)
	runtime.KeepAlive(platformData)
}
