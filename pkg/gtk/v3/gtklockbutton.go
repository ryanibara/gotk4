// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_lock_button_get_type()), F: marshalLockButton},
	})
}

// LockButton: gtkLockButton is a widget that can be used in control panels or
// preference dialogs to allow users to obtain and revoke authorizations needed
// to operate the controls. The required authorization is represented by a
// #GPermission object. Concrete implementations of #GPermission may use
// PolicyKit or some other authorization framework. To obtain a PolicyKit-based
// #GPermission, use polkit_permission_new().
//
// If the user is not currently allowed to perform the action, but can obtain
// the permission, the widget looks like this:
//
// ! (lockbutton-locked.png)
//
// and the user can click the button to request the permission. Depending on the
// platform, this may pop up an authentication dialog or ask the user to
// authenticate in some other way. Once the user has obtained the permission,
// the widget changes to this:
//
// ! (lockbutton-unlocked.png)
//
// and the permission can be dropped again by clicking the button. If the user
// is not able to obtain the permission at all, the widget looks like this:
//
// ! (lockbutton-sorry.png)
//
// If the user has the permission and cannot drop it, the button is hidden.
//
// The text (and tooltips) that are shown in the various cases can be adjusted
// with the LockButton:text-lock, LockButton:text-unlock,
// LockButton:tooltip-lock, LockButton:tooltip-unlock and
// LockButton:tooltip-not-authorized properties.
type LockButton interface {
	Button
	Actionable
	Activatable
	Buildable

	// Permission obtains the #GPermission object that controls @button.
	Permission() gio.Permission
	// SetPermission sets the #GPermission object that controls @button.
	SetPermission(permission gio.Permission)
}

// lockButton implements the LockButton class.
type lockButton struct {
	Button
	Actionable
	Activatable
	Buildable
}

var _ LockButton = (*lockButton)(nil)

// WrapLockButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapLockButton(obj *externglib.Object) LockButton {
	return lockButton{
		Button:      WrapButton(obj),
		Actionable:  WrapActionable(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalLockButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLockButton(obj), nil
}

// NewLockButton constructs a class LockButton.
func NewLockButton(permission gio.Permission) LockButton {
	var _arg1 *C.GPermission  // out
	var _cret C.GtkLockButton // in

	_arg1 = (*C.GPermission)(unsafe.Pointer(permission.Native()))

	_cret = C.gtk_lock_button_new(_arg1)

	var _lockButton LockButton // out

	_lockButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(LockButton)

	return _lockButton
}

// Permission obtains the #GPermission object that controls @button.
func (b lockButton) Permission() gio.Permission {
	var _arg0 *C.GtkLockButton // out
	var _cret *C.GPermission   // in

	_arg0 = (*C.GtkLockButton)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_lock_button_get_permission(_arg0)

	var _permission gio.Permission // out

	_permission = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.Permission)

	return _permission
}

// SetPermission sets the #GPermission object that controls @button.
func (b lockButton) SetPermission(permission gio.Permission) {
	var _arg0 *C.GtkLockButton // out
	var _arg1 *C.GPermission   // out

	_arg0 = (*C.GtkLockButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GPermission)(unsafe.Pointer(permission.Native()))

	C.gtk_lock_button_set_permission(_arg0, _arg1)
}
