// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_lock_button_get_type()), F: marshalLockButtonner},
	})
}

// LockButton: GtkLockButton is a widget to obtain and revoke authorizations
// needed to operate the controls.
//
// !An example GtkLockButton (lock-button.png)
//
// It is typically used in preference dialogs or control panels.
//
// The required authorization is represented by a GPermission object. Concrete
// implementations of GPermission may use PolicyKit or some other authorization
// framework. To obtain a PolicyKit-based GPermission, use
// polkit_permission_new().
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
// with the gtk.LockButton:text-lock, gtk.LockButton:text-unlock,
// gtk.LockButton:tooltip-lock, gtk.LockButton:tooltip-unlock and
// gtk.LockButton:tooltip-not-authorized properties.
type LockButton struct {
	Button
}

func wrapLockButton(obj *externglib.Object) *LockButton {
	return &LockButton{
		Button: Button{
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
			Actionable: Actionable{
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
			},
			Object: obj,
		},
	}
}

func marshalLockButtonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLockButton(obj), nil
}

// NewLockButton creates a new lock button which reflects the permission.
func NewLockButton(permission gio.Permissioner) *LockButton {
	var _arg1 *C.GPermission // out
	var _cret *C.GtkWidget   // in

	if permission != nil {
		_arg1 = (*C.GPermission)(unsafe.Pointer(permission.Native()))
	}

	_cret = C.gtk_lock_button_new(_arg1)
	runtime.KeepAlive(permission)

	var _lockButton *LockButton // out

	_lockButton = wrapLockButton(externglib.Take(unsafe.Pointer(_cret)))

	return _lockButton
}

// Permission obtains the GPermission object that controls button.
func (button *LockButton) Permission() gio.Permissioner {
	var _arg0 *C.GtkLockButton // out
	var _cret *C.GPermission   // in

	_arg0 = (*C.GtkLockButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_lock_button_get_permission(_arg0)
	runtime.KeepAlive(button)

	var _permission gio.Permissioner // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Permissioner is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(gio.Permissioner)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.Permissioner")
		}
		_permission = rv
	}

	return _permission
}

// SetPermission sets the GPermission object that controls button.
//
// The function takes the following parameters:
//
//    - permission: GPermission object, or NULL.
//
func (button *LockButton) SetPermission(permission gio.Permissioner) {
	var _arg0 *C.GtkLockButton // out
	var _arg1 *C.GPermission   // out

	_arg0 = (*C.GtkLockButton)(unsafe.Pointer(button.Native()))
	if permission != nil {
		_arg1 = (*C.GPermission)(unsafe.Pointer(permission.Native()))
	}

	C.gtk_lock_button_set_permission(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(permission)
}
