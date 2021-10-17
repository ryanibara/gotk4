// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_simple_action_group_get_type()), F: marshalSimpleActionGrouper},
	})
}

// SimpleActionGroup is a hash table filled with #GAction objects, implementing
// the Group and Map interfaces.
type SimpleActionGroup struct {
	*externglib.Object

	ActionGroup
	ActionMap
}

func wrapSimpleActionGroup(obj *externglib.Object) *SimpleActionGroup {
	return &SimpleActionGroup{
		Object: obj,
		ActionGroup: ActionGroup{
			Object: obj,
		},
		ActionMap: ActionMap{
			Object: obj,
		},
	}
}

func marshalSimpleActionGrouper(p uintptr) (interface{}, error) {
	return wrapSimpleActionGroup(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSimpleActionGroup creates a new, empty, ActionGroup.
func NewSimpleActionGroup() *SimpleActionGroup {
	var _cret *C.GSimpleActionGroup // in

	_cret = C.g_simple_action_group_new()

	var _simpleActionGroup *SimpleActionGroup // out

	_simpleActionGroup = wrapSimpleActionGroup(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _simpleActionGroup
}

// AddEntries: convenience function for creating multiple Action instances and
// adding them to the action group.
//
// Deprecated: Use g_action_map_add_action_entries().
//
// The function takes the following parameters:
//
//    - entries: pointer to the first item in an array of Entry structs.
//    - userData: user data for signal connections.
//
func (simple *SimpleActionGroup) AddEntries(entries []ActionEntry, userData cgo.Handle) {
	var _arg0 *C.GSimpleActionGroup // out
	var _arg1 *C.GActionEntry       // out
	var _arg2 C.gint
	var _arg3 C.gpointer // out

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(simple.Native()))
	_arg2 = (C.gint)(len(entries))
	_arg1 = (*C.GActionEntry)(C.malloc(C.ulong(len(entries)) * C.ulong(C.sizeof_GActionEntry)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.GActionEntry)(_arg1), len(entries))
		for i := range entries {
			out[i] = *(*C.GActionEntry)(gextras.StructNative(unsafe.Pointer((&entries[i]))))
		}
	}
	_arg3 = (C.gpointer)(unsafe.Pointer(userData))

	C.g_simple_action_group_add_entries(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(simple)
	runtime.KeepAlive(entries)
	runtime.KeepAlive(userData)
}

// Insert adds an action to the action group.
//
// If the action group already contains an action with the same name as action
// then the old action is dropped from the group.
//
// The action group takes its own reference on action.
//
// Deprecated: Use g_action_map_add_action().
//
// The function takes the following parameters:
//
//    - action: #GAction.
//
func (simple *SimpleActionGroup) Insert(action Actioner) {
	var _arg0 *C.GSimpleActionGroup // out
	var _arg1 *C.GAction            // out

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(simple.Native()))
	_arg1 = (*C.GAction)(unsafe.Pointer(action.Native()))

	C.g_simple_action_group_insert(_arg0, _arg1)
	runtime.KeepAlive(simple)
	runtime.KeepAlive(action)
}

// Lookup looks up the action with the name action_name in the group.
//
// If no such action exists, returns NULL.
//
// Deprecated: Use g_action_map_lookup_action().
//
// The function takes the following parameters:
//
//    - actionName: name of an action.
//
func (simple *SimpleActionGroup) Lookup(actionName string) Actioner {
	var _arg0 *C.GSimpleActionGroup // out
	var _arg1 *C.gchar              // out
	var _cret *C.GAction            // in

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(simple.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_simple_action_group_lookup(_arg0, _arg1)
	runtime.KeepAlive(simple)
	runtime.KeepAlive(actionName)

	var _action Actioner // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Actioner is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Actioner)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.Actioner")
		}
		_action = rv
	}

	return _action
}

// Remove removes the named action from the action group.
//
// If no action of this name is in the group then nothing happens.
//
// Deprecated: Use g_action_map_remove_action().
//
// The function takes the following parameters:
//
//    - actionName: name of the action.
//
func (simple *SimpleActionGroup) Remove(actionName string) {
	var _arg0 *C.GSimpleActionGroup // out
	var _arg1 *C.gchar              // out

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(simple.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_simple_action_group_remove(_arg0, _arg1)
	runtime.KeepAlive(simple)
	runtime.KeepAlive(actionName)
}
