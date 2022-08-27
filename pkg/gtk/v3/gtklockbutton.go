// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeLockButton = coreglib.Type(C.gtk_lock_button_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeLockButton, F: marshalLockButton},
	})
}

// LockButtonOverrides contains methods that are overridable.
type LockButtonOverrides struct {
}

func defaultLockButtonOverrides(v *LockButton) LockButtonOverrides {
	return LockButtonOverrides{}
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

func init() {
	coreglib.RegisterClassInfo[*LockButton, *LockButtonClass, LockButtonOverrides](
		GTypeLockButton,
		initLockButtonClass,
		wrapLockButton,
		defaultLockButtonOverrides,
	)
}

func initLockButtonClass(gclass unsafe.Pointer, overrides LockButtonOverrides, classInitFunc func(*LockButtonClass)) {
	if classInitFunc != nil {
		class := (*LockButtonClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
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

// LockButtonClass: instance of this type is always passed by reference.
type LockButtonClass struct {
	*lockButtonClass
}

// lockButtonClass is the struct that's finalized.
type lockButtonClass struct {
	native *C.GtkLockButtonClass
}

// ParentClass: parent class.
func (l *LockButtonClass) ParentClass() *ButtonClass {
	valptr := &l.native.parent_class
	var _v *ButtonClass // out
	_v = (*ButtonClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
