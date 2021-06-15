// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_toggle_action_get_type()), F: marshalToggleAction},
	})
}

// ToggleAction: a ToggleAction corresponds roughly to a CheckMenuItem. It has
// an “active” state specifying whether the action has been checked or not.
type ToggleAction interface {
	Action
	Buildable

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

// toggleAction implements the ToggleAction class.
type toggleAction struct {
	Action
	Buildable
}

var _ ToggleAction = (*toggleAction)(nil)

// WrapToggleAction wraps a GObject to the right type. It is
// primarily used internally.
func WrapToggleAction(obj *externglib.Object) ToggleAction {
	return toggleAction{
		Action:    WrapAction(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalToggleAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToggleAction(obj), nil
}

// NewToggleAction constructs a class ToggleAction.
func NewToggleAction(name string, label string, tooltip string, stockId string) ToggleAction {
	var _arg1 *C.gchar          // out
	var _arg2 *C.gchar          // out
	var _arg3 *C.gchar          // out
	var _arg4 *C.gchar          // out
	var _cret C.GtkToggleAction // in

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg4))

	_cret = C.gtk_toggle_action_new(_arg1, _arg2, _arg3, _arg4)

	var _toggleAction ToggleAction // out

	_toggleAction = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ToggleAction)

	return _toggleAction
}

// Active returns the checked state of the toggle action.
func (a toggleAction) Active() bool {
	var _arg0 *C.GtkToggleAction // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_toggle_action_get_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DrawAsRadio returns whether the action should have proxies like a radio
// action.
func (a toggleAction) DrawAsRadio() bool {
	var _arg0 *C.GtkToggleAction // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_toggle_action_get_draw_as_radio(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive sets the checked state on the toggle action.
func (a toggleAction) SetActive(isActive bool) {
	var _arg0 *C.GtkToggleAction // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(a.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_action_set_active(_arg0, _arg1)
}

// SetDrawAsRadio sets whether the action should have proxies like a radio
// action.
func (a toggleAction) SetDrawAsRadio(drawAsRadio bool) {
	var _arg0 *C.GtkToggleAction // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(a.Native()))
	if drawAsRadio {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_action_set_draw_as_radio(_arg0, _arg1)
}

// Toggled emits the “toggled” signal on the toggle action.
func (a toggleAction) Toggled() {
	var _arg0 *C.GtkToggleAction // out

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(a.Native()))

	C.gtk_toggle_action_toggled(_arg0)
}
