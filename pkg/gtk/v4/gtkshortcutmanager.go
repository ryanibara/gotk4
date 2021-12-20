// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcut_manager_get_type()), F: marshalShortcutManagerer},
	})
}

// ShortcutManagerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ShortcutManagerOverrider interface {
	// The function takes the following parameters:
	//
	AddController(controller *ShortcutController)
	// The function takes the following parameters:
	//
	RemoveController(controller *ShortcutController)
}

// ShortcutManager: GtkShortcutManager interface is used to implement shortcut
// scopes.
//
// This is important for gtk.Native widgets that have their own surface, since
// the event controllers that are used to implement managed and global scopes
// are limited to the same native.
//
// Examples for widgets implementing GtkShortcutManager are gtk.Window and
// gtk.Popover.
//
// Every widget that implements GtkShortcutManager will be used as a
// GTK_SHORTCUT_SCOPE_MANAGED.
type ShortcutManager struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*ShortcutManager)(nil)
)

// ShortcutManagerer describes ShortcutManager's interface methods.
type ShortcutManagerer interface {
	externglib.Objector

	baseShortcutManager() *ShortcutManager
}

var _ ShortcutManagerer = (*ShortcutManager)(nil)

func wrapShortcutManager(obj *externglib.Object) *ShortcutManager {
	return &ShortcutManager{
		Object: obj,
	}
}

func marshalShortcutManagerer(p uintptr) (interface{}, error) {
	return wrapShortcutManager(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *ShortcutManager) baseShortcutManager() *ShortcutManager {
	return self
}

// BaseShortcutManager returns the underlying base object.
func BaseShortcutManager(obj ShortcutManagerer) *ShortcutManager {
	return obj.baseShortcutManager()
}
