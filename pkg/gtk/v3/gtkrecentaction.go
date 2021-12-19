// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
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

// RecentActionOverrider contains methods that are overridable.
type RecentActionOverrider interface {
}

// RecentAction represents a list of recently used files, which can be shown by
// widgets such as RecentChooserDialog or RecentChooserMenu.
//
// To construct a submenu showing recently used files, use a RecentAction as the
// action for a <menuitem>. To construct a menu toolbutton showing the recently
// used files in the popup menu, use a RecentAction as the action for a
// <toolitem> element.
type RecentAction struct {
	_ [0]func() // equal guard
	Action

	*externglib.Object
	RecentChooser
}

var (
	_ externglib.Objector = (*RecentAction)(nil)
)

func classInitRecentActioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRecentAction(obj *externglib.Object) *RecentAction {
	return &RecentAction{
		Action: Action{
			Object: obj,
			Buildable: Buildable{
				Object: obj,
			},
		},
		Object: obj,
		RecentChooser: RecentChooser{
			Object: obj,
		},
	}
}

func marshalRecentActioner(p uintptr) (interface{}, error) {
	return wrapRecentAction(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewRecentAction creates a new RecentAction object. To add the action to a
// ActionGroup and set the accelerator for the action, call
// gtk_action_group_add_action_with_accel().
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - name: unique name for the action.
//    - label (optional) displayed in menu items and on buttons, or NULL.
//    - tooltip (optional) for the action, or NULL.
//    - stockId (optional): stock icon to display in widgets representing the
//      action, or NULL.
//
// The function returns the following values:
//
//    - recentAction: newly created RecentAction.
//
func NewRecentAction(name, label, tooltip, stockId string) *RecentAction {
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _arg3 *C.gchar     // out
	var _arg4 *C.gchar     // out
	var _cret *C.GtkAction // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	if label != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	if tooltip != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(tooltip)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	if stockId != "" {
		_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
		defer C.free(unsafe.Pointer(_arg4))
	}

	_cret = C.gtk_recent_action_new(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(name)
	runtime.KeepAlive(label)
	runtime.KeepAlive(tooltip)
	runtime.KeepAlive(stockId)

	var _recentAction *RecentAction // out

	_recentAction = wrapRecentAction(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _recentAction
}

// NewRecentActionForManager creates a new RecentAction object. To add the
// action to a ActionGroup and set the accelerator for the action, call
// gtk_action_group_add_action_with_accel().
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - name: unique name for the action.
//    - label (optional) displayed in menu items and on buttons, or NULL.
//    - tooltip (optional) for the action, or NULL.
//    - stockId (optional): stock icon to display in widgets representing the
//      action, or NULL.
//    - manager (optional) or NULL for using the default RecentManager.
//
// The function returns the following values:
//
//    - recentAction: newly created RecentAction.
//
func NewRecentActionForManager(name, label, tooltip, stockId string, manager *RecentManager) *RecentAction {
	var _arg1 *C.gchar            // out
	var _arg2 *C.gchar            // out
	var _arg3 *C.gchar            // out
	var _arg4 *C.gchar            // out
	var _arg5 *C.GtkRecentManager // out
	var _cret *C.GtkAction        // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	if label != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	if tooltip != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(tooltip)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	if stockId != "" {
		_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
		defer C.free(unsafe.Pointer(_arg4))
	}
	if manager != nil {
		_arg5 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))
	}

	_cret = C.gtk_recent_action_new_for_manager(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(name)
	runtime.KeepAlive(label)
	runtime.KeepAlive(tooltip)
	runtime.KeepAlive(stockId)
	runtime.KeepAlive(manager)

	var _recentAction *RecentAction // out

	_recentAction = wrapRecentAction(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _recentAction
}

// ShowNumbers returns the value set by
// gtk_recent_chooser_menu_set_show_numbers().
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - ok: TRUE if numbers should be shown.
//
func (action *RecentAction) ShowNumbers() bool {
	var _arg0 *C.GtkRecentAction // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkRecentAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_recent_action_get_show_numbers(_arg0)
	runtime.KeepAlive(action)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetShowNumbers sets whether a number should be added to the items shown by
// the widgets representing action. The numbers are shown to provide a unique
// character for a mnemonic to be used inside the menu item's label. Only the
// first ten items get a number to avoid clashes.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - showNumbers: TRUE if the shown items should be numbered.
//
func (action *RecentAction) SetShowNumbers(showNumbers bool) {
	var _arg0 *C.GtkRecentAction // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkRecentAction)(unsafe.Pointer(action.Native()))
	if showNumbers {
		_arg1 = C.TRUE
	}

	C.gtk_recent_action_set_show_numbers(_arg0, _arg1)
	runtime.KeepAlive(action)
	runtime.KeepAlive(showNumbers)
}
