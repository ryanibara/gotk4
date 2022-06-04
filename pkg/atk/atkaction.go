// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_atk1_ActionIface_do_action(void*, gint);
// extern gboolean _gotk4_atk1_ActionIface_set_description(void*, gint, void*);
// extern gchar* _gotk4_atk1_ActionIface_get_description(void*, gint);
// extern gchar* _gotk4_atk1_ActionIface_get_keybinding(void*, gint);
// extern gchar* _gotk4_atk1_ActionIface_get_localized_name(void*, gint);
// extern gchar* _gotk4_atk1_ActionIface_get_name(void*, gint);
// extern gint _gotk4_atk1_ActionIface_get_n_actions(void*);
import "C"

// glib.Type values for atkaction.go.
var GTypeAction = coreglib.Type(C.atk_action_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeAction, F: marshalAction},
	})
}

// ActionOverrider contains methods that are overridable.
type ActionOverrider interface {
	// DoAction: perform the specified action on the object.
	//
	// The function takes the following parameters:
	//
	//    - i: action index corresponding to the action to be performed.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if success, FALSE otherwise.
	//
	DoAction(i int32) bool
	// Description returns a description of the specified action of the object.
	//
	// The function takes the following parameters:
	//
	//    - i: action index corresponding to the action to be performed.
	//
	// The function returns the following values:
	//
	//    - utf8 (optional): description string, or NULL if action does not
	//      implement this interface.
	//
	Description(i int32) string
	// Keybinding gets the keybinding which can be used to activate this action,
	// if one exists. The string returned should contain localized,
	// human-readable, key sequences as they would appear when displayed on
	// screen. It must be in the format "mnemonic;sequence;shortcut".
	//
	// - The mnemonic key activates the object if it is presently enabled
	// onscreen. This typically corresponds to the underlined letter within the
	// widget. Example: "n" in a traditional "New..." menu item or the "a" in
	// "Apply" for a button.
	//
	// - The sequence is the full list of keys which invoke the action even if
	// the relevant element is not currently shown on screen. For instance, for
	// a menu item the sequence is the keybindings used to open the parent menus
	// before invoking. The sequence string is colon-delimited. Example:
	// "Alt+F:N" in a traditional "New..." menu item.
	//
	// - The shortcut, if it exists, will invoke the same action without showing
	// the component or its enclosing menus or dialogs. Example: "Ctrl+N" in a
	// traditional "New..." menu item.
	//
	// Example: For a traditional "New..." menu item, the expected return value
	// would be: "N;Alt+F:N;Ctrl+N" for the English locale and
	// "N;Alt+D:N;Strg+N" for the German locale. If, hypothetically, this menu
	// item lacked a mnemonic, it would be represented by ";;Ctrl+N" and
	// ";;Strg+N" respectively.
	//
	// The function takes the following parameters:
	//
	//    - i: action index corresponding to the action to be performed.
	//
	// The function returns the following values:
	//
	//    - utf8 (optional): keybinding which can be used to activate this
	//      action, or NULL if there is no keybinding for this action.
	//
	Keybinding(i int32) string
	// LocalizedName returns the localized name of the specified action of the
	// object.
	//
	// The function takes the following parameters:
	//
	//    - i: action index corresponding to the action to be performed.
	//
	// The function returns the following values:
	//
	//    - utf8 (optional): name string, or NULL if action does not implement
	//      this interface.
	//
	LocalizedName(i int32) string
	// NActions gets the number of accessible actions available on the object.
	// If there are more than one, the first one is considered the "default"
	// action of the object.
	//
	// The function returns the following values:
	//
	//    - gint: the number of actions, or 0 if action does not implement this
	//      interface.
	//
	NActions() int32
	// Name returns a non-localized string naming the specified action of the
	// object. This name is generally not descriptive of the end result of the
	// action, but instead names the 'interaction type' which the object
	// supports. By convention, the above strings should be used to represent
	// the actions which correspond to the common point-and-click interaction
	// techniques of the same name: i.e. "click", "press", "release", "drag",
	// "drop", "popup", etc. The "popup" action should be used to pop up a
	// context menu for the object, if one exists.
	//
	// For technical reasons, some toolkits cannot guarantee that the reported
	// action is actually 'bound' to a nontrivial user event; i.e. the result of
	// some actions via atk_action_do_action() may be NIL.
	//
	// The function takes the following parameters:
	//
	//    - i: action index corresponding to the action to be performed.
	//
	// The function returns the following values:
	//
	//    - utf8 (optional): name string, or NULL if action does not implement
	//      this interface.
	//
	Name(i int32) string
	// SetDescription sets a description of the specified action of the object.
	//
	// The function takes the following parameters:
	//
	//    - i: action index corresponding to the action to be performed.
	//    - desc: description to be assigned to this action.
	//
	// The function returns the following values:
	//
	//    - ok: gboolean representing if the description was successfully set;.
	//
	SetDescription(i int32, desc string) bool
}

// Action should be implemented by instances of Object classes with which the
// user can interact directly, i.e. buttons, checkboxes, scrollbars, e.g.
// components which are not "passive" providers of UI information.
//
// Exceptions: when the user interaction is already covered by another
// appropriate interface such as EditableText (insert/delete text, etc.) or
// Value (set value) then these actions should not be exposed by Action as well.
//
// Though most UI interactions on components should be invocable via keyboard as
// well as mouse, there will generally be a close mapping between "mouse
// actions" that are possible on a component and the AtkActions. Where mouse and
// keyboard actions are redundant in effect, Action should expose only one
// action rather than exposing redundant actions if possible. By convention we
// have been using "mouse centric" terminology for Action names.
//
// Action wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Action struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Action)(nil)
)

// Actioner describes Action's interface methods.
type Actioner interface {
	coreglib.Objector

	// DoAction: perform the specified action on the object.
	DoAction(i int32) bool
	// Description returns a description of the specified action of the object.
	Description(i int32) string
	// Keybinding gets the keybinding which can be used to activate this action,
	// if one exists.
	Keybinding(i int32) string
	// LocalizedName returns the localized name of the specified action of the
	// object.
	LocalizedName(i int32) string
	// NActions gets the number of accessible actions available on the object.
	NActions() int32
	// Name returns a non-localized string naming the specified action of the
	// object.
	Name(i int32) string
	// SetDescription sets a description of the specified action of the object.
	SetDescription(i int32, desc string) bool
}

var _ Actioner = (*Action)(nil)

func ifaceInitActioner(gifacePtr, data C.gpointer) {
	iface := (*C.AtkActionIface)(unsafe.Pointer(gifacePtr))
	iface.do_action = (*[0]byte)(C._gotk4_atk1_ActionIface_do_action)
	iface.get_description = (*[0]byte)(C._gotk4_atk1_ActionIface_get_description)
	iface.get_keybinding = (*[0]byte)(C._gotk4_atk1_ActionIface_get_keybinding)
	iface.get_localized_name = (*[0]byte)(C._gotk4_atk1_ActionIface_get_localized_name)
	iface.get_n_actions = (*[0]byte)(C._gotk4_atk1_ActionIface_get_n_actions)
	iface.get_name = (*[0]byte)(C._gotk4_atk1_ActionIface_get_name)
	iface.set_description = (*[0]byte)(C._gotk4_atk1_ActionIface_set_description)
}

//export _gotk4_atk1_ActionIface_do_action
func _gotk4_atk1_ActionIface_do_action(arg0 *C.void, arg1 C.gint) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionOverrider)

	var _i int32 // out

	_i = int32(arg1)

	ok := iface.DoAction(_i)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_atk1_ActionIface_get_description
func _gotk4_atk1_ActionIface_get_description(arg0 *C.void, arg1 C.gint) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionOverrider)

	var _i int32 // out

	_i = int32(arg1)

	utf8 := iface.Description(_i)

	if utf8 != "" {
		cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
		defer C.free(unsafe.Pointer(cret))
	}

	return cret
}

//export _gotk4_atk1_ActionIface_get_keybinding
func _gotk4_atk1_ActionIface_get_keybinding(arg0 *C.void, arg1 C.gint) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionOverrider)

	var _i int32 // out

	_i = int32(arg1)

	utf8 := iface.Keybinding(_i)

	if utf8 != "" {
		cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
		defer C.free(unsafe.Pointer(cret))
	}

	return cret
}

//export _gotk4_atk1_ActionIface_get_localized_name
func _gotk4_atk1_ActionIface_get_localized_name(arg0 *C.void, arg1 C.gint) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionOverrider)

	var _i int32 // out

	_i = int32(arg1)

	utf8 := iface.LocalizedName(_i)

	if utf8 != "" {
		cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
		defer C.free(unsafe.Pointer(cret))
	}

	return cret
}

//export _gotk4_atk1_ActionIface_get_n_actions
func _gotk4_atk1_ActionIface_get_n_actions(arg0 *C.void) (cret C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionOverrider)

	gint := iface.NActions()

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_ActionIface_get_name
func _gotk4_atk1_ActionIface_get_name(arg0 *C.void, arg1 C.gint) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionOverrider)

	var _i int32 // out

	_i = int32(arg1)

	utf8 := iface.Name(_i)

	if utf8 != "" {
		cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
		defer C.free(unsafe.Pointer(cret))
	}

	return cret
}

//export _gotk4_atk1_ActionIface_set_description
func _gotk4_atk1_ActionIface_set_description(arg0 *C.void, arg1 C.gint, arg2 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionOverrider)

	var _i int32     // out
	var _desc string // out

	_i = int32(arg1)
	_desc = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	ok := iface.SetDescription(_i, _desc)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapAction(obj *coreglib.Object) *Action {
	return &Action{
		Object: obj,
	}
}

func marshalAction(p uintptr) (interface{}, error) {
	return wrapAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DoAction: perform the specified action on the object.
//
// The function takes the following parameters:
//
//    - i: action index corresponding to the action to be performed.
//
// The function returns the following values:
//
//    - ok: TRUE if success, FALSE otherwise.
//
func (action *Action) DoAction(i int32) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(i)

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(action)
	runtime.KeepAlive(i)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Description returns a description of the specified action of the object.
//
// The function takes the following parameters:
//
//    - i: action index corresponding to the action to be performed.
//
// The function returns the following values:
//
//    - utf8 (optional): description string, or NULL if action does not implement
//      this interface.
//
func (action *Action) Description(i int32) string {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(i)

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(action)
	runtime.KeepAlive(i)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Keybinding gets the keybinding which can be used to activate this action, if
// one exists. The string returned should contain localized, human-readable, key
// sequences as they would appear when displayed on screen. It must be in the
// format "mnemonic;sequence;shortcut".
//
// - The mnemonic key activates the object if it is presently enabled onscreen.
// This typically corresponds to the underlined letter within the widget.
// Example: "n" in a traditional "New..." menu item or the "a" in "Apply" for a
// button.
//
// - The sequence is the full list of keys which invoke the action even if the
// relevant element is not currently shown on screen. For instance, for a menu
// item the sequence is the keybindings used to open the parent menus before
// invoking. The sequence string is colon-delimited. Example: "Alt+F:N" in a
// traditional "New..." menu item.
//
// - The shortcut, if it exists, will invoke the same action without showing the
// component or its enclosing menus or dialogs. Example: "Ctrl+N" in a
// traditional "New..." menu item.
//
// Example: For a traditional "New..." menu item, the expected return value
// would be: "N;Alt+F:N;Ctrl+N" for the English locale and "N;Alt+D:N;Strg+N"
// for the German locale. If, hypothetically, this menu item lacked a mnemonic,
// it would be represented by ";;Ctrl+N" and ";;Strg+N" respectively.
//
// The function takes the following parameters:
//
//    - i: action index corresponding to the action to be performed.
//
// The function returns the following values:
//
//    - utf8 (optional): keybinding which can be used to activate this action, or
//      NULL if there is no keybinding for this action.
//
func (action *Action) Keybinding(i int32) string {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(i)

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(action)
	runtime.KeepAlive(i)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// LocalizedName returns the localized name of the specified action of the
// object.
//
// The function takes the following parameters:
//
//    - i: action index corresponding to the action to be performed.
//
// The function returns the following values:
//
//    - utf8 (optional): name string, or NULL if action does not implement this
//      interface.
//
func (action *Action) LocalizedName(i int32) string {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(i)

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(action)
	runtime.KeepAlive(i)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// NActions gets the number of accessible actions available on the object. If
// there are more than one, the first one is considered the "default" action of
// the object.
//
// The function returns the following values:
//
//    - gint: the number of actions, or 0 if action does not implement this
//      interface.
//
func (action *Action) NActions() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(action)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// Name returns a non-localized string naming the specified action of the
// object. This name is generally not descriptive of the end result of the
// action, but instead names the 'interaction type' which the object supports.
// By convention, the above strings should be used to represent the actions
// which correspond to the common point-and-click interaction techniques of the
// same name: i.e. "click", "press", "release", "drag", "drop", "popup", etc.
// The "popup" action should be used to pop up a context menu for the object, if
// one exists.
//
// For technical reasons, some toolkits cannot guarantee that the reported
// action is actually 'bound' to a nontrivial user event; i.e. the result of
// some actions via atk_action_do_action() may be NIL.
//
// The function takes the following parameters:
//
//    - i: action index corresponding to the action to be performed.
//
// The function returns the following values:
//
//    - utf8 (optional): name string, or NULL if action does not implement this
//      interface.
//
func (action *Action) Name(i int32) string {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(i)

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(action)
	runtime.KeepAlive(i)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SetDescription sets a description of the specified action of the object.
//
// The function takes the following parameters:
//
//    - i: action index corresponding to the action to be performed.
//    - desc: description to be assigned to this action.
//
// The function returns the following values:
//
//    - ok: gboolean representing if the description was successfully set;.
//
func (action *Action) SetDescription(i int32, desc string) bool {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(i)
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(desc)))
	defer C.free(unsafe.Pointer(_args[2]))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(action)
	runtime.KeepAlive(i)
	runtime.KeepAlive(desc)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}
