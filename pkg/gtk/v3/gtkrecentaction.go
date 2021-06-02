// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdbool.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_recent_action_get_type()), F: marshalRecentAction},
	})
}

type RecentActionPrivate struct {
	native C.GtkRecentActionPrivate
}

// WrapRecentActionPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRecentActionPrivate(ptr unsafe.Pointer) *RecentActionPrivate {
	if ptr == nil {
		return nil
	}

	return (*RecentActionPrivate)(ptr)
}

func marshalRecentActionPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRecentActionPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RecentActionPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// RecentAction: a RecentAction represents a list of recently used files, which
// can be shown by widgets such as RecentChooserDialog or RecentChooserMenu.
//
// To construct a submenu showing recently used files, use a RecentAction as the
// action for a <menuitem>. To construct a menu toolbutton showing the recently
// used files in the popup menu, use a RecentAction as the action for a
// <toolitem> element.
type RecentAction interface {
	Action
	Buildable
	RecentChooser

	// ShowNumbers returns the value set by
	// gtk_recent_chooser_menu_set_show_numbers().
	ShowNumbers() bool
	// SetShowNumbers sets whether a number should be added to the items shown
	// by the widgets representing @action. The numbers are shown to provide a
	// unique character for a mnemonic to be used inside the menu item's label.
	// Only the first ten items get a number to avoid clashes.
	SetShowNumbers(showNumbers bool)
}

// recentAction implements the RecentAction interface.
type recentAction struct {
	Action
	Buildable
	RecentChooser
}

var _ RecentAction = (*recentAction)(nil)

// WrapRecentAction wraps a GObject to the right type. It is
// primarily used internally.
func WrapRecentAction(obj *externglib.Object) RecentAction {
	return RecentAction{
		Action:        WrapAction(obj),
		Buildable:     WrapBuildable(obj),
		RecentChooser: WrapRecentChooser(obj),
	}
}

func marshalRecentAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRecentAction(obj), nil
}

// NewRecentAction constructs a class RecentAction.
func NewRecentAction(name string, label string, tooltip string, stockID string) RecentAction {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 *C.gchar

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = (*C.gchar)(C.CString(stockID))
	defer C.free(unsafe.Pointer(arg4))

	ret := C.gtk_recent_action_new(arg1, arg2, arg3, arg4)

	var ret0 RecentAction

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(RecentAction)

	return ret0
}

// NewRecentActionForManager constructs a class RecentAction.
func NewRecentActionForManager(name string, label string, tooltip string, stockID string, manager RecentManager) RecentAction {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 *C.gchar
	var arg5 *C.GtkRecentManager

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = (*C.gchar)(C.CString(stockID))
	defer C.free(unsafe.Pointer(arg4))
	arg5 = (*C.GtkRecentManager)(manager.Native())

	ret := C.gtk_recent_action_new_for_manager(arg1, arg2, arg3, arg4, arg5)

	var ret0 RecentAction

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(RecentAction)

	return ret0
}

// ShowNumbers returns the value set by
// gtk_recent_chooser_menu_set_show_numbers().
func (a recentAction) ShowNumbers() bool {
	var arg0 *C.GtkRecentAction

	arg0 = (*C.GtkRecentAction)(a.Native())

	ret := C.gtk_recent_action_get_show_numbers(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// SetShowNumbers sets whether a number should be added to the items shown
// by the widgets representing @action. The numbers are shown to provide a
// unique character for a mnemonic to be used inside the menu item's label.
// Only the first ten items get a number to avoid clashes.
func (a recentAction) SetShowNumbers(showNumbers bool) {
	var arg0 *C.GtkRecentAction
	var arg1 C.gboolean

	arg0 = (*C.GtkRecentAction)(a.Native())
	if showNumbers {
		arg1 = C.TRUE
	}

	C.gtk_recent_action_set_show_numbers(arg0, arg1)
}
