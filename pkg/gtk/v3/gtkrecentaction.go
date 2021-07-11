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
		{T: externglib.Type(C.gtk_recent_action_get_type()), F: marshalRecentActioner},
	})
}

// RecentActioner describes RecentAction's methods.
type RecentActioner interface {
	// ShowNumbers returns the value set by
	// gtk_recent_chooser_menu_set_show_numbers().
	ShowNumbers() bool
	// SetShowNumbers sets whether a number should be added to the items shown
	// by the widgets representing @action.
	SetShowNumbers(showNumbers bool)
}

// RecentAction represents a list of recently used files, which can be shown by
// widgets such as RecentChooserDialog or RecentChooserMenu.
//
// To construct a submenu showing recently used files, use a RecentAction as the
// action for a <menuitem>. To construct a menu toolbutton showing the recently
// used files in the popup menu, use a RecentAction as the action for a
// <toolitem> element.
type RecentAction struct {
	Action

	RecentChooser
}

var (
	_ RecentActioner  = (*RecentAction)(nil)
	_ gextras.Nativer = (*RecentAction)(nil)
)

func wrapRecentAction(obj *externglib.Object) RecentActioner {
	return &RecentAction{
		Action: Action{
			Object: obj,
			Buildable: Buildable{
				Object: obj,
			},
		},
		RecentChooser: RecentChooser{
			Object: obj,
		},
	}
}

func marshalRecentActioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRecentAction(obj), nil
}

// NewRecentAction creates a new RecentAction object. To add the action to a
// ActionGroup and set the accelerator for the action, call
// gtk_action_group_add_action_with_accel().
//
// Deprecated: since version 3.10.
func NewRecentAction(name string, label string, tooltip string, stockId string) *RecentAction {
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _arg3 *C.gchar     // out
	var _arg4 *C.gchar     // out
	var _cret *C.GtkAction // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(tooltip)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg4))

	_cret = C.gtk_recent_action_new(_arg1, _arg2, _arg3, _arg4)

	var _recentAction *RecentAction // out

	_recentAction = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*RecentAction)

	return _recentAction
}

// NewRecentActionForManager creates a new RecentAction object. To add the
// action to a ActionGroup and set the accelerator for the action, call
// gtk_action_group_add_action_with_accel().
//
// Deprecated: since version 3.10.
func NewRecentActionForManager(name string, label string, tooltip string, stockId string, manager RecentManagerer) *RecentAction {
	var _arg1 *C.gchar            // out
	var _arg2 *C.gchar            // out
	var _arg3 *C.gchar            // out
	var _arg4 *C.gchar            // out
	var _arg5 *C.GtkRecentManager // out
	var _cret *C.GtkAction        // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(tooltip)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = (*C.GtkRecentManager)(unsafe.Pointer((manager).(gextras.Nativer).Native()))

	_cret = C.gtk_recent_action_new_for_manager(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _recentAction *RecentAction // out

	_recentAction = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*RecentAction)

	return _recentAction
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *RecentAction) Native() uintptr {
	return v.Action.Object.Native()
}

// ShowNumbers returns the value set by
// gtk_recent_chooser_menu_set_show_numbers().
//
// Deprecated: since version 3.10.
func (action *RecentAction) ShowNumbers() bool {
	var _arg0 *C.GtkRecentAction // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkRecentAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_recent_action_get_show_numbers(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetShowNumbers sets whether a number should be added to the items shown by
// the widgets representing @action. The numbers are shown to provide a unique
// character for a mnemonic to be used inside the menu item's label. Only the
// first ten items get a number to avoid clashes.
//
// Deprecated: since version 3.10.
func (action *RecentAction) SetShowNumbers(showNumbers bool) {
	var _arg0 *C.GtkRecentAction // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkRecentAction)(unsafe.Pointer(action.Native()))
	if showNumbers {
		_arg1 = C.TRUE
	}

	C.gtk_recent_action_set_show_numbers(_arg0, _arg1)
}
