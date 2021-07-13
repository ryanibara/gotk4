// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_action_group_get_type()), F: marshalActionGrouper},
	})
}

// ActionGroupOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ActionGroupOverrider interface {
	// Action looks up an action in the action group by name.
	//
	// Deprecated: since version 3.10.
	Action(actionName string) *Action
}

// ActionGrouper describes ActionGroup's methods.
type ActionGrouper interface {
	// AddAction adds an action object to the action group.
	AddAction(action *Action)
	// AddActionWithAccel adds an action object to the action group and sets up
	// the accelerator.
	AddActionWithAccel(action *Action, accelerator string)
	// AccelGroup gets the accelerator group.
	AccelGroup() *AccelGroup
	// Action looks up an action in the action group by name.
	Action(actionName string) *Action
	// Name gets the name of the action group.
	Name() string
	// Sensitive returns TRUE if the group is sensitive.
	Sensitive() bool
	// Visible returns TRUE if the group is visible.
	Visible() bool
	// RemoveAction removes an action object from the action group.
	RemoveAction(action *Action)
	// SetAccelGroup sets the accelerator group to be used by every action in
	// this group.
	SetAccelGroup(accelGroup *AccelGroup)
	// SetSensitive changes the sensitivity of action_group Deprecated: since
	// version 3.10.
	SetSensitive(sensitive bool)
	// SetTranslationDomain sets the translation domain and uses g_dgettext()
	// for translating the label and tooltip of ActionEntrys added by
	// gtk_action_group_add_actions().
	SetTranslationDomain(domain string)
	// SetVisible changes the visible of action_group.
	SetVisible(visible bool)
	// TranslateString translates a string using the function set with
	// gtk_action_group_set_translate_func().
	TranslateString(_string string) string
}

// ActionGroup actions are organised into groups. An action group is essentially
// a map from names to Action objects.
//
// All actions that would make sense to use in a particular context should be in
// a single group. Multiple action groups may be used for a particular user
// interface. In fact, it is expected that most nontrivial applications will
// make use of multiple groups. For example, in an application that can edit
// multiple documents, one group holding global actions (e.g. quit, about, new),
// and one group per document holding actions that act on that document (eg.
// save, cut/copy/paste, etc). Each window’s menus would be constructed from a
// combination of two action groups.
//
//
// Accelerators
//
// Accelerators are handled by the GTK+ accelerator map. All actions are
// assigned an accelerator path (which normally has the form
// <Actions>/group-name/action-name) and a shortcut is associated with this
// accelerator path. All menuitems and toolitems take on this accelerator path.
// The GTK+ accelerator map code makes sure that the correct shortcut is
// displayed next to the menu item.
//
//
// GtkActionGroup as GtkBuildable
//
// The ActionGroup implementation of the Buildable interface accepts Action
// objects as <child> elements in UI definitions.
//
// Note that it is probably more common to define actions and action groups in
// the code, since they are directly related to what the code can do.
//
// The GtkActionGroup implementation of the GtkBuildable interface supports a
// custom <accelerator> element, which has attributes named “key“ and
// “modifiers“ and allows to specify accelerators. This is similar to the
// <accelerator> element of Widget, the main difference is that it doesn’t allow
// you to specify a signal.
//
// A Dialog UI definition fragment. ##
//
//    <object class="GtkActionGroup" id="actiongroup">
//      <child>
//          <object class="GtkAction" id="About">
//              <property name="name">About</property>
//              <property name="stock_id">gtk-about</property>
//              <signal handler="about_activate" name="activate"/>
//          </object>
//          <accelerator key="F1" modifiers="GDK_CONTROL_MASK | GDK_SHIFT_MASK"/>
//      </child>
//    </object>
type ActionGroup struct {
	*externglib.Object

	Buildable
}

var (
	_ ActionGrouper   = (*ActionGroup)(nil)
	_ gextras.Nativer = (*ActionGroup)(nil)
)

func wrapActionGroup(obj *externglib.Object) *ActionGroup {
	return &ActionGroup{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalActionGrouper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapActionGroup(obj), nil
}

// NewActionGroup creates a new ActionGroup object. The name of the action group
// is used when associating [keybindings][Action-Accel] with the actions.
//
// Deprecated: since version 3.10.
func NewActionGroup(name string) *ActionGroup {
	var _arg1 *C.gchar          // out
	var _cret *C.GtkActionGroup // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))

	_cret = C.gtk_action_group_new(_arg1)

	var _actionGroup *ActionGroup // out

	_actionGroup = wrapActionGroup(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _actionGroup
}

// AddAction adds an action object to the action group. Note that this function
// does not set up the accel path of the action, which can lead to problems if a
// user tries to modify the accelerator of a menuitem associated with the
// action. Therefore you must either set the accel path yourself with
// gtk_action_set_accel_path(), or use gtk_action_group_add_action_with_accel
// (..., NULL).
//
// Deprecated: since version 3.10.
func (actionGroup *ActionGroup) AddAction(action *Action) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.GtkAction      // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_action_group_add_action(_arg0, _arg1)
}

// AddActionWithAccel adds an action object to the action group and sets up the
// accelerator.
//
// If accelerator is NULL, attempts to use the accelerator associated with the
// stock_id of the action.
//
// Accel paths are set to <Actions>/group-name/action-name.
//
// Deprecated: since version 3.10.
func (actionGroup *ActionGroup) AddActionWithAccel(action *Action, accelerator string) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.GtkAction      // out
	var _arg2 *C.gchar          // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(accelerator)))

	C.gtk_action_group_add_action_with_accel(_arg0, _arg1, _arg2)
}

// AccelGroup gets the accelerator group.
//
// Deprecated: since version 3.10.
func (actionGroup *ActionGroup) AccelGroup() *AccelGroup {
	var _arg0 *C.GtkActionGroup // out
	var _cret *C.GtkAccelGroup  // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))

	_cret = C.gtk_action_group_get_accel_group(_arg0)

	var _accelGroup *AccelGroup // out

	_accelGroup = wrapAccelGroup(externglib.Take(unsafe.Pointer(_cret)))

	return _accelGroup
}

// Action looks up an action in the action group by name.
//
// Deprecated: since version 3.10.
func (actionGroup *ActionGroup) Action(actionName string) *Action {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.gchar          // out
	var _cret *C.GtkAction      // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))

	_cret = C.gtk_action_group_get_action(_arg0, _arg1)

	var _action *Action // out

	_action = wrapAction(externglib.Take(unsafe.Pointer(_cret)))

	return _action
}

// Name gets the name of the action group.
//
// Deprecated: since version 3.10.
func (actionGroup *ActionGroup) Name() string {
	var _arg0 *C.GtkActionGroup // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))

	_cret = C.gtk_action_group_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Sensitive returns TRUE if the group is sensitive. The constituent actions can
// only be logically sensitive (see gtk_action_is_sensitive()) if they are
// sensitive (see gtk_action_get_sensitive()) and their group is sensitive.
//
// Deprecated: since version 3.10.
func (actionGroup *ActionGroup) Sensitive() bool {
	var _arg0 *C.GtkActionGroup // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))

	_cret = C.gtk_action_group_get_sensitive(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Visible returns TRUE if the group is visible. The constituent actions can
// only be logically visible (see gtk_action_is_visible()) if they are visible
// (see gtk_action_get_visible()) and their group is visible.
//
// Deprecated: since version 3.10.
func (actionGroup *ActionGroup) Visible() bool {
	var _arg0 *C.GtkActionGroup // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))

	_cret = C.gtk_action_group_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RemoveAction removes an action object from the action group.
//
// Deprecated: since version 3.10.
func (actionGroup *ActionGroup) RemoveAction(action *Action) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.GtkAction      // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_action_group_remove_action(_arg0, _arg1)
}

// SetAccelGroup sets the accelerator group to be used by every action in this
// group.
//
// Deprecated: since version 3.10.
func (actionGroup *ActionGroup) SetAccelGroup(accelGroup *AccelGroup) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.GtkAccelGroup  // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))

	C.gtk_action_group_set_accel_group(_arg0, _arg1)
}

// SetSensitive changes the sensitivity of action_group
//
// Deprecated: since version 3.10.
func (actionGroup *ActionGroup) SetSensitive(sensitive bool) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	if sensitive {
		_arg1 = C.TRUE
	}

	C.gtk_action_group_set_sensitive(_arg0, _arg1)
}

// SetTranslationDomain sets the translation domain and uses g_dgettext() for
// translating the label and tooltip of ActionEntrys added by
// gtk_action_group_add_actions().
//
// If you’re not using gettext() for localization, see
// gtk_action_group_set_translate_func().
//
// Deprecated: since version 3.10.
func (actionGroup *ActionGroup) SetTranslationDomain(domain string) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(domain)))

	C.gtk_action_group_set_translation_domain(_arg0, _arg1)
}

// SetVisible changes the visible of action_group.
//
// Deprecated: since version 3.10.
func (actionGroup *ActionGroup) SetVisible(visible bool) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_action_group_set_visible(_arg0, _arg1)
}

// TranslateString translates a string using the function set with
// gtk_action_group_set_translate_func(). This is mainly intended for language
// bindings.
//
// Deprecated: since version 3.10.
func (actionGroup *ActionGroup) TranslateString(_string string) string {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.gchar          // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))

	_cret = C.gtk_action_group_translate_string(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ActionEntry structs are used with gtk_action_group_add_actions() to construct
// actions.
//
// Deprecated: since version 3.10.
type ActionEntry struct {
	native C.GtkActionEntry
}

// Native returns the underlying C source pointer.
func (a *ActionEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// RadioActionEntry structs are used with gtk_action_group_add_radio_actions()
// to construct groups of radio actions.
//
// Deprecated: since version 3.10.
type RadioActionEntry struct {
	native C.GtkRadioActionEntry
}

// Native returns the underlying C source pointer.
func (r *RadioActionEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// ToggleActionEntry structs are used with gtk_action_group_add_toggle_actions()
// to construct toggle actions.
//
// Deprecated: since version 3.10.
type ToggleActionEntry struct {
	native C.GtkToggleActionEntry
}

// Native returns the underlying C source pointer.
func (t *ToggleActionEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}
