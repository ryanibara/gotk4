// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_action_get_type()), F: marshalAction},
	})
}

// Action: Action should be implemented by instances of Object classes with
// which the user can interact directly, i.e. buttons, checkboxes, scrollbars,
// e.g. components which are not "passive" providers of UI information.
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
type Action interface {
	gextras.Objector

	// DoAction sets a description of the specified action of the object.
	DoAction(i int) bool
	// Description sets a description of the specified action of the object.
	Description(i int) string
	// Keybinding sets a description of the specified action of the object.
	Keybinding(i int) string
	// LocalizedName sets a description of the specified action of the object.
	LocalizedName(i int) string
	// NActions sets a description of the specified action of the object.
	NActions() int
	// Name sets a description of the specified action of the object.
	Name(i int) string
	// SetDescription sets a description of the specified action of the object.
	SetDescription(i int, desc string) bool
}

// action implements the Action interface.
type action struct {
	gextras.Objector
}

var _ Action = (*action)(nil)

// WrapAction wraps a GObject to a type that implements
// interface Action. It is primarily used internally.
func WrapAction(obj *externglib.Object) Action {
	return action{
		Objector: obj,
	}
}

func marshalAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAction(obj), nil
}

func (a action) DoAction(i int) bool {
	var _arg0 *C.AtkAction // out
	var _arg1 C.gint       // out
	var _cret C.gboolean   // in

	_arg0 = (*C.AtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_action_do_action(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) Description(i int) string {
	var _arg0 *C.AtkAction // out
	var _arg1 C.gint       // out
	var _cret *C.gchar     // in

	_arg0 = (*C.AtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_action_get_description(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) Keybinding(i int) string {
	var _arg0 *C.AtkAction // out
	var _arg1 C.gint       // out
	var _cret *C.gchar     // in

	_arg0 = (*C.AtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_action_get_keybinding(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) LocalizedName(i int) string {
	var _arg0 *C.AtkAction // out
	var _arg1 C.gint       // out
	var _cret *C.gchar     // in

	_arg0 = (*C.AtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_action_get_localized_name(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) NActions() int {
	var _arg0 *C.AtkAction // out
	var _cret C.gint       // in

	_arg0 = (*C.AtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.atk_action_get_n_actions(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (a action) Name(i int) string {
	var _arg0 *C.AtkAction // out
	var _arg1 C.gint       // out
	var _cret *C.gchar     // in

	_arg0 = (*C.AtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_action_get_name(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) SetDescription(i int, desc string) bool {
	var _arg0 *C.AtkAction // out
	var _arg1 C.gint       // out
	var _arg2 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.AtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = C.gint(i)
	_arg2 = (*C.gchar)(C.CString(desc))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.atk_action_set_description(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
