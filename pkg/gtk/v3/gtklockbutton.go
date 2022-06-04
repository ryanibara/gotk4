// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtklockbutton.go.
var GTypeLockButton = coreglib.Type(C.gtk_lock_button_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeLockButton, F: marshalLockButton},
	})
}

// LockButtonOverrider contains methods that are overridable.
type LockButtonOverrider interface {
}

// LockButton is a widget that can be used in control panels or preference
// dialogs to allow users to obtain and revoke authorizations needed to operate
// the controls. The required authorization is represented by a #GPermission
// object. Concrete implementations of #GPermission may use PolicyKit or some
// other authorization framework. To obtain a PolicyKit-based #GPermission, use
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
// with the LockButton:text-lock, LockButton:text-unlock,
// LockButton:tooltip-lock, LockButton:tooltip-unlock and
// LockButton:tooltip-not-authorized properties.
type LockButton struct {
	_ [0]func() // equal guard
	Button
}

var (
	_ Binner            = (*LockButton)(nil)
	_ coreglib.Objector = (*LockButton)(nil)
)

func classInitLockButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapLockButton(obj *coreglib.Object) *LockButton {
	return &LockButton{
		Button: Button{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: coreglib.InitiallyUnowned{
							Object: obj,
						},
						Object: obj,
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
					},
				},
			},
			Object: obj,
			Actionable: Actionable{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Activatable: Activatable{
				Object: obj,
			},
		},
	}
}

func marshalLockButton(p uintptr) (interface{}, error) {
	return wrapLockButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewLockButton creates a new lock button which reflects the permission.
//
// The function takes the following parameters:
//
//    - permission (optional): #GPermission.
//
// The function returns the following values:
//
//    - lockButton: new LockButton.
//
func NewLockButton(permission gio.Permissioner) *LockButton {
	var _args [1]girepository.Argument

	if permission != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(permission).Native()))
	}

	_gret := girepository.MustFind("Gtk", "LockButton").InvokeMethod("new_LockButton", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(permission)

	var _lockButton *LockButton // out

	_lockButton = wrapLockButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _lockButton
}

// Permission obtains the #GPermission object that controls button.
//
// The function returns the following values:
//
//    - permission of button.
//
func (button *LockButton) Permission() gio.Permissioner {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "LockButton").InvokeMethod("get_permission", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _permission gio.Permissioner // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Permissioner is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gio.Permissioner)
			return ok
		})
		rv, ok := casted.(gio.Permissioner)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Permissioner")
		}
		_permission = rv
	}

	return _permission
}

// SetPermission sets the #GPermission object that controls button.
//
// The function takes the following parameters:
//
//    - permission (optional) object, or NULL.
//
func (button *LockButton) SetPermission(permission gio.Permissioner) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if permission != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(permission).Native()))
	}

	girepository.MustFind("Gtk", "LockButton").InvokeMethod("set_permission", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(permission)
}
