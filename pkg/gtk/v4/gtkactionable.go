// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_actionable_get_type()), F: marshalActionabler},
	})
}

// ActionableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ActionableOverrider interface {
	// ActionName gets the action name for actionable.
	ActionName() string
	// ActionTargetValue gets the current target value of actionable.
	ActionTargetValue() *glib.Variant
	// SetActionName specifies the name of the action with which this widget
	// should be associated.
	//
	// If action_name is NULL then the widget will be unassociated from any
	// previous action.
	//
	// Usually this function is used when the widget is located (or will be
	// located) within the hierarchy of a GtkApplicationWindow.
	//
	// Names are of the form “win.save” or “app.quit” for actions on the
	// containing GtkApplicationWindow or its associated GtkApplication,
	// respectively. This is the same form used for actions in the GMenu
	// associated with the window.
	SetActionName(actionName string)
	// SetActionTargetValue sets the target value of an actionable widget.
	//
	// If target_value is NULL then the target value is unset.
	//
	// The target value has two purposes. First, it is used as the parameter to
	// activation of the action associated with the GtkActionable widget.
	// Second, it is used to determine if the widget should be rendered as
	// “active” — the widget is active if the state is equal to the given
	// target.
	//
	// Consider the example of associating a set of buttons with a GAction with
	// string state in a typical “radio button” situation. Each button will be
	// associated with the same action, but with a different target value for
	// that action. Clicking on a particular button will activate the action
	// with the target of that button, which will typically cause the action’s
	// state to change to that value. Since the action’s state is now equal to
	// the target value of the button, the button will now be rendered as active
	// (and the other buttons, with different targets, rendered inactive).
	SetActionTargetValue(targetValue *glib.Variant)
}

// Actionable: GtkActionable interface provides a convenient way of asscociating
// widgets with actions.
//
// It primarily consists of two properties: gtk.Actionable:action-name and
// gtk.Actionable:action-target. There are also some convenience APIs for
// setting these properties.
//
// The action will be looked up in action groups that are found among the
// widgets ancestors. Most commonly, these will be the actions with the “win.”
// or “app.” prefix that are associated with the GtkApplicationWindow or
// GtkApplication, but other action groups that are added with
// gtk.Widget.InsertActionGroup() will be consulted as well.
type Actionable struct {
	Widget
}

// Actionabler describes Actionable's abstract methods.
type Actionabler interface {
	externglib.Objector

	// ActionName gets the action name for actionable.
	ActionName() string
	// ActionTargetValue gets the current target value of actionable.
	ActionTargetValue() *glib.Variant
	// SetActionName specifies the name of the action with which this widget
	// should be associated.
	SetActionName(actionName string)
	// SetActionTargetValue sets the target value of an actionable widget.
	SetActionTargetValue(targetValue *glib.Variant)
	// SetDetailedActionName sets the action-name and associated string target
	// value of an actionable widget.
	SetDetailedActionName(detailedActionName string)
}

var _ Actionabler = (*Actionable)(nil)

func wrapActionable(obj *externglib.Object) *Actionable {
	return &Actionable{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalActionabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapActionable(obj), nil
}

// ActionName gets the action name for actionable.
func (actionable *Actionable) ActionName() string {
	var _arg0 *C.GtkActionable // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkActionable)(unsafe.Pointer(actionable.Native()))

	_cret = C.gtk_actionable_get_action_name(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ActionTargetValue gets the current target value of actionable.
func (actionable *Actionable) ActionTargetValue() *glib.Variant {
	var _arg0 *C.GtkActionable // out
	var _cret *C.GVariant      // in

	_arg0 = (*C.GtkActionable)(unsafe.Pointer(actionable.Native()))

	_cret = C.gtk_actionable_get_action_target_value(_arg0)

	var _variant *glib.Variant // out

	if _cret != nil {
		_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.g_variant_ref(_cret)
		runtime.SetFinalizer(_variant, func(v *glib.Variant) {
			C.g_variant_unref((*C.GVariant)(gextras.StructNative(unsafe.Pointer(v))))
		})
	}

	return _variant
}

// SetActionName specifies the name of the action with which this widget should
// be associated.
//
// If action_name is NULL then the widget will be unassociated from any previous
// action.
//
// Usually this function is used when the widget is located (or will be located)
// within the hierarchy of a GtkApplicationWindow.
//
// Names are of the form “win.save” or “app.quit” for actions on the containing
// GtkApplicationWindow or its associated GtkApplication, respectively. This is
// the same form used for actions in the GMenu associated with the window.
func (actionable *Actionable) SetActionName(actionName string) {
	var _arg0 *C.GtkActionable // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkActionable)(unsafe.Pointer(actionable.Native()))
	if actionName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(actionName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_actionable_set_action_name(_arg0, _arg1)
}

// SetActionTargetValue sets the target value of an actionable widget.
//
// If target_value is NULL then the target value is unset.
//
// The target value has two purposes. First, it is used as the parameter to
// activation of the action associated with the GtkActionable widget. Second, it
// is used to determine if the widget should be rendered as “active” — the
// widget is active if the state is equal to the given target.
//
// Consider the example of associating a set of buttons with a GAction with
// string state in a typical “radio button” situation. Each button will be
// associated with the same action, but with a different target value for that
// action. Clicking on a particular button will activate the action with the
// target of that button, which will typically cause the action’s state to
// change to that value. Since the action’s state is now equal to the target
// value of the button, the button will now be rendered as active (and the other
// buttons, with different targets, rendered inactive).
func (actionable *Actionable) SetActionTargetValue(targetValue *glib.Variant) {
	var _arg0 *C.GtkActionable // out
	var _arg1 *C.GVariant      // out

	_arg0 = (*C.GtkActionable)(unsafe.Pointer(actionable.Native()))
	if targetValue != nil {
		_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(targetValue)))
	}

	C.gtk_actionable_set_action_target_value(_arg0, _arg1)
}

// SetDetailedActionName sets the action-name and associated string target value
// of an actionable widget.
//
// detailed_action_name is a string in the format accepted by
// g_action_parse_detailed_name().
func (actionable *Actionable) SetDetailedActionName(detailedActionName string) {
	var _arg0 *C.GtkActionable // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkActionable)(unsafe.Pointer(actionable.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(detailedActionName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_actionable_set_detailed_action_name(_arg0, _arg1)
}
