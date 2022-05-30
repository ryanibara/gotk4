// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern GAction* _gotk4_gio2_ActionMapInterface_lookup_action(GActionMap*, gchar*);
// extern void _gotk4_gio2_ActionMapInterface_add_action(GActionMap*, GAction*);
// extern void _gotk4_gio2_ActionMapInterface_remove_action(GActionMap*, gchar*);
import "C"

// glib.Type values for gactionmap.go.
var GTypeActionMap = coreglib.Type(C.g_action_map_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeActionMap, F: marshalActionMap},
	})
}

// ActionMapOverrider contains methods that are overridable.
type ActionMapOverrider interface {
	// AddAction adds an action to the action_map.
	//
	// If the action map already contains an action with the same name as action
	// then the old action is dropped from the action map.
	//
	// The action map takes its own reference on action.
	//
	// The function takes the following parameters:
	//
	//    - action: #GAction.
	//
	AddAction(action Actioner)
	// LookupAction looks up the action with the name action_name in action_map.
	//
	// If no such action exists, returns NULL.
	//
	// The function takes the following parameters:
	//
	//    - actionName: name of an action.
	//
	// The function returns the following values:
	//
	//    - action (optional) or NULL.
	//
	LookupAction(actionName string) *Action
	// RemoveAction removes the named action from the action map.
	//
	// If no action of this name is in the map then nothing happens.
	//
	// The function takes the following parameters:
	//
	//    - actionName: name of the action.
	//
	RemoveAction(actionName string)
}

// ActionMap interface is implemented by Group implementations that operate by
// containing a number of named #GAction instances, such as ActionGroup.
//
// One useful application of this interface is to map the names of actions from
// various action groups to unique, prefixed names (e.g. by prepending "app." or
// "win."). This is the motivation for the 'Map' part of the interface name.
//
// ActionMap wraps an interface. This means the user can get the
// underlying type by calling Cast().
type ActionMap struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ActionMap)(nil)
)

// ActionMapper describes ActionMap's interface methods.
type ActionMapper interface {
	coreglib.Objector

	// AddAction adds an action to the action_map.
	AddAction(action Actioner)
	// LookupAction looks up the action with the name action_name in action_map.
	LookupAction(actionName string) *Action
	// RemoveAction removes the named action from the action map.
	RemoveAction(actionName string)
}

var _ ActionMapper = (*ActionMap)(nil)

func ifaceInitActionMapper(gifacePtr, data C.gpointer) {
	iface := (*C.GActionMapInterface)(unsafe.Pointer(gifacePtr))
	iface.add_action = (*[0]byte)(C._gotk4_gio2_ActionMapInterface_add_action)
	iface.lookup_action = (*[0]byte)(C._gotk4_gio2_ActionMapInterface_lookup_action)
	iface.remove_action = (*[0]byte)(C._gotk4_gio2_ActionMapInterface_remove_action)
}

//export _gotk4_gio2_ActionMapInterface_add_action
func _gotk4_gio2_ActionMapInterface_add_action(arg0 *C.GActionMap, arg1 *C.GAction) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionMapOverrider)

	var _action Actioner // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.Actioner is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Actioner)
			return ok
		})
		rv, ok := casted.(Actioner)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Actioner")
		}
		_action = rv
	}

	iface.AddAction(_action)
}

//export _gotk4_gio2_ActionMapInterface_lookup_action
func _gotk4_gio2_ActionMapInterface_lookup_action(arg0 *C.GActionMap, arg1 *C.gchar) (cret *C.GAction) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionMapOverrider)

	var _actionName string // out

	_actionName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	action := iface.LookupAction(_actionName)

	if action != nil {
		cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	}

	return cret
}

//export _gotk4_gio2_ActionMapInterface_remove_action
func _gotk4_gio2_ActionMapInterface_remove_action(arg0 *C.GActionMap, arg1 *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionMapOverrider)

	var _actionName string // out

	_actionName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	iface.RemoveAction(_actionName)
}

func wrapActionMap(obj *coreglib.Object) *ActionMap {
	return &ActionMap{
		Object: obj,
	}
}

func marshalActionMap(p uintptr) (interface{}, error) {
	return wrapActionMap(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AddAction adds an action to the action_map.
//
// If the action map already contains an action with the same name as action
// then the old action is dropped from the action map.
//
// The action map takes its own reference on action.
//
// The function takes the following parameters:
//
//    - action: #GAction.
//
func (actionMap *ActionMap) AddAction(action Actioner) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionMap).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	*(**ActionMap)(unsafe.Pointer(&args[1])) = _arg1

	runtime.KeepAlive(actionMap)
	runtime.KeepAlive(action)
}

// LookupAction looks up the action with the name action_name in action_map.
//
// If no such action exists, returns NULL.
//
// The function takes the following parameters:
//
//    - actionName: name of an action.
//
// The function returns the following values:
//
//    - action (optional) or NULL.
//
func (actionMap *ActionMap) LookupAction(actionName string) *Action {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionMap).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**ActionMap)(unsafe.Pointer(&args[1])) = _arg1

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(actionMap)
	runtime.KeepAlive(actionName)

	var _action *Action // out

	if _cret != nil {
		_action = wrapAction(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _action
}

// RemoveAction removes the named action from the action map.
//
// If no action of this name is in the map then nothing happens.
//
// The function takes the following parameters:
//
//    - actionName: name of the action.
//
func (actionMap *ActionMap) RemoveAction(actionName string) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionMap).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**ActionMap)(unsafe.Pointer(&args[1])) = _arg1

	runtime.KeepAlive(actionMap)
	runtime.KeepAlive(actionName)
}

// ActionEntry: this struct defines a single action. It is for use with
// g_action_map_add_action_entries().
//
// The order of the items in the structure are intended to reflect frequency of
// use. It is permissible to use an incomplete initialiser in order to leave
// some of the later values as NULL. All values after name are optional.
// Additional optional fields may be added in the future.
//
// See g_action_map_add_action_entries() for an example.
//
// An instance of this type is always passed by reference.
type ActionEntry struct {
	*actionEntry
}

// actionEntry is the struct that's finalized.
type actionEntry struct {
	native *C.GActionEntry
}

// Name: name of the action.
func (a *ActionEntry) Name() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(a.native.name)))
	return v
}

// ParameterType: type of the parameter that must be passed to the activate
// function for this action, given as a single GVariant type string (or NULL for
// no parameter).
func (a *ActionEntry) ParameterType() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(a.native.parameter_type)))
	return v
}

// State: initial state for this action, given in [GVariant text
// format][gvariant-text]. The state is parsed with no extra type information,
// so type tags must be added to the string if they are necessary. Stateless
// actions should give NULL here.
func (a *ActionEntry) State() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(a.native.state)))
	return v
}
