// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeTearoffMenuItem returns the GType for the type TearoffMenuItem.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTearoffMenuItem() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "TearoffMenuItem").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalTearoffMenuItem)
	return gtype
}

// TearoffMenuItemOverrider contains methods that are overridable.
type TearoffMenuItemOverrider interface {
}

// TearoffMenuItem is a special MenuItem which is used to tear off and reattach
// its menu.
//
// When its menu is shown normally, the TearoffMenuItem is drawn as a dotted
// line indicating that the menu can be torn off. Activating it causes its menu
// to be torn off and displayed in its own window as a tearoff menu.
//
// When its menu is shown as a tearoff menu, the TearoffMenuItem is drawn as a
// dotted line which has a left pointing arrow graphic indicating that the
// tearoff menu can be reattached. Activating it will erase the tearoff menu
// window.
//
// > TearoffMenuItem is deprecated and should not be used in newly > written
// code. Menus are not meant to be torn around.
type TearoffMenuItem struct {
	_ [0]func() // equal guard
	MenuItem
}

var (
	_ Binner            = (*TearoffMenuItem)(nil)
	_ coreglib.Objector = (*TearoffMenuItem)(nil)
)

func classInitTearoffMenuItemmer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTearoffMenuItem(obj *coreglib.Object) *TearoffMenuItem {
	return &TearoffMenuItem{
		MenuItem: MenuItem{
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

func marshalTearoffMenuItem(p uintptr) (interface{}, error) {
	return wrapTearoffMenuItem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTearoffMenuItem creates a new TearoffMenuItem.
//
// Deprecated: TearoffMenuItem is deprecated and should not be used in newly
// written code.
//
// The function returns the following values:
//
//    - tearoffMenuItem: new TearoffMenuItem.
//
func NewTearoffMenuItem() *TearoffMenuItem {
	_gret := girepository.MustFind("Gtk", "TearoffMenuItem").InvokeMethod("new_TearoffMenuItem", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _tearoffMenuItem *TearoffMenuItem // out

	_tearoffMenuItem = wrapTearoffMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _tearoffMenuItem
}
