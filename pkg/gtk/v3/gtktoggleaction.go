// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_toggle_action_get_type()), F: marshalToggleActioner},
	})
}

// ToggleActionOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ToggleActionOverrider interface {
	// Toggled emits the “toggled” signal on the toggle action.
	//
	// Deprecated: since version 3.10.
	Toggled()
}

// ToggleActioner describes ToggleAction's methods.
type ToggleActioner interface {
	// Active returns the checked state of the toggle action.
	Active() bool
	// DrawAsRadio returns whether the action should have proxies like a radio
	// action.
	DrawAsRadio() bool
	// SetActive sets the checked state on the toggle action.
	SetActive(isActive bool)
	// SetDrawAsRadio sets whether the action should have proxies like a radio
	// action.
	SetDrawAsRadio(drawAsRadio bool)
	// Toggled emits the “toggled” signal on the toggle action.
	Toggled()
}

// ToggleAction corresponds roughly to a CheckMenuItem. It has an “active” state
// specifying whether the action has been checked or not.
type ToggleAction struct {
	Action
}

var (
	_ ToggleActioner  = (*ToggleAction)(nil)
	_ gextras.Nativer = (*ToggleAction)(nil)
)

func wrapToggleAction(obj *externglib.Object) ToggleActioner {
	return &ToggleAction{
		Action: Action{
			Object: obj,
			Buildable: Buildable{
				Object: obj,
			},
		},
	}
}

func marshalToggleActioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapToggleAction(obj), nil
}

// NewToggleAction creates a new ToggleAction object. To add the action to a
// ActionGroup and set the accelerator for the action, call
// gtk_action_group_add_action_with_accel().
//
// Deprecated: since version 3.10.
func NewToggleAction(name string, label string, tooltip string, stockId string) *ToggleAction {
	var _arg1 *C.gchar           // out
	var _arg2 *C.gchar           // out
	var _arg3 *C.gchar           // out
	var _arg4 *C.gchar           // out
	var _cret *C.GtkToggleAction // in

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg4))

	_cret = C.gtk_toggle_action_new(_arg1, _arg2, _arg3, _arg4)

	var _toggleAction *ToggleAction // out

	_toggleAction = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*ToggleAction)

	return _toggleAction
}

// Active returns the checked state of the toggle action.
//
// Deprecated: since version 3.10.
func (action *ToggleAction) Active() bool {
	var _arg0 *C.GtkToggleAction // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_toggle_action_get_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DrawAsRadio returns whether the action should have proxies like a radio
// action.
//
// Deprecated: since version 3.10.
func (action *ToggleAction) DrawAsRadio() bool {
	var _arg0 *C.GtkToggleAction // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_toggle_action_get_draw_as_radio(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive sets the checked state on the toggle action.
//
// Deprecated: since version 3.10.
func (action *ToggleAction) SetActive(isActive bool) {
	var _arg0 *C.GtkToggleAction // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(action.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_action_set_active(_arg0, _arg1)
}

// SetDrawAsRadio sets whether the action should have proxies like a radio
// action.
//
// Deprecated: since version 3.10.
func (action *ToggleAction) SetDrawAsRadio(drawAsRadio bool) {
	var _arg0 *C.GtkToggleAction // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(action.Native()))
	if drawAsRadio {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_action_set_draw_as_radio(_arg0, _arg1)
}

// Toggled emits the “toggled” signal on the toggle action.
//
// Deprecated: since version 3.10.
func (action *ToggleAction) Toggled() {
	var _arg0 *C.GtkToggleAction // out

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(action.Native()))

	C.gtk_toggle_action_toggled(_arg0)
}
