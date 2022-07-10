// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk3_CheckMenuItemClass_draw_indicator(void*, void*);
// extern void _gotk4_gtk3_CheckMenuItemClass_toggled(void*);
// extern void _gotk4_gtk3_CheckMenuItem_ConnectToggled(gpointer, guintptr);
import "C"

// GTypeCheckMenuItem returns the GType for the type CheckMenuItem.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCheckMenuItem() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "CheckMenuItem").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalCheckMenuItem)
	return gtype
}

// CheckMenuItemOverrider contains methods that are overridable.
type CheckMenuItemOverrider interface {
	// The function takes the following parameters:
	//
	DrawIndicator(cr *cairo.Context)
	// Toggled emits the CheckMenuItem::toggled signal.
	Toggled()
}

// CheckMenuItem is a menu item that maintains the state of a boolean value in
// addition to a MenuItem usual role in activating application code.
//
// A check box indicating the state of the boolean value is displayed at the
// left side of the MenuItem. Activating the MenuItem toggles the value.
//
// CSS nodes
//
//    menuitem
//    ├── check.left
//    ╰── <child>
//
// GtkCheckMenuItem has a main CSS node with name menuitem, and a subnode with
// name check, which gets the .left or .right style class.
type CheckMenuItem struct {
	_ [0]func() // equal guard
	MenuItem
}

var (
	_ Binner            = (*CheckMenuItem)(nil)
	_ coreglib.Objector = (*CheckMenuItem)(nil)
)

func classInitCheckMenuItemmer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "CheckMenuItemClass")

	if _, ok := goval.(interface{ DrawIndicator(cr *cairo.Context) }); ok {
		o := pclass.StructFieldOffset("draw_indicator")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_CheckMenuItemClass_draw_indicator)
	}

	if _, ok := goval.(interface{ Toggled() }); ok {
		o := pclass.StructFieldOffset("toggled")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_CheckMenuItemClass_toggled)
	}
}

//export _gotk4_gtk3_CheckMenuItemClass_draw_indicator
func _gotk4_gtk3_CheckMenuItemClass_draw_indicator(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ DrawIndicator(cr *cairo.Context) })

	var _cr *cairo.Context // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})

	iface.DrawIndicator(_cr)
}

//export _gotk4_gtk3_CheckMenuItemClass_toggled
func _gotk4_gtk3_CheckMenuItemClass_toggled(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Toggled() })

	iface.Toggled()
}

func wrapCheckMenuItem(obj *coreglib.Object) *CheckMenuItem {
	return &CheckMenuItem{
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

func marshalCheckMenuItem(p uintptr) (interface{}, error) {
	return wrapCheckMenuItem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_CheckMenuItem_ConnectToggled
func _gotk4_gtk3_CheckMenuItem_ConnectToggled(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectToggled: this signal is emitted when the state of the check box is
// changed.
//
// A signal handler can use gtk_check_menu_item_get_active() to discover the new
// state.
func (checkMenuItem *CheckMenuItem) ConnectToggled(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(checkMenuItem, "toggled", false, unsafe.Pointer(C._gotk4_gtk3_CheckMenuItem_ConnectToggled), f)
}

// NewCheckMenuItem creates a new CheckMenuItem.
//
// The function returns the following values:
//
//    - checkMenuItem: new CheckMenuItem.
//
func NewCheckMenuItem() *CheckMenuItem {
	_info := girepository.MustFind("Gtk", "CheckMenuItem")
	_gret := _info.InvokeClassMethod("new_CheckMenuItem", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _checkMenuItem *CheckMenuItem // out

	_checkMenuItem = wrapCheckMenuItem(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _checkMenuItem
}

// NewCheckMenuItemWithLabel creates a new CheckMenuItem with a label.
//
// The function takes the following parameters:
//
//    - label: string to use for the label.
//
// The function returns the following values:
//
//    - checkMenuItem: new CheckMenuItem.
//
func NewCheckMenuItemWithLabel(label string) *CheckMenuItem {
	var _args [1]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "CheckMenuItem")
	_gret := _info.InvokeClassMethod("new_CheckMenuItem_with_label", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(label)

	var _checkMenuItem *CheckMenuItem // out

	_checkMenuItem = wrapCheckMenuItem(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _checkMenuItem
}

// NewCheckMenuItemWithMnemonic creates a new CheckMenuItem containing a label.
// The label will be created using gtk_label_new_with_mnemonic(), so underscores
// in label indicate the mnemonic for the menu item.
//
// The function takes the following parameters:
//
//    - label: text of the button, with an underscore in front of the character.
//
// The function returns the following values:
//
//    - checkMenuItem: new CheckMenuItem.
//
func NewCheckMenuItemWithMnemonic(label string) *CheckMenuItem {
	var _args [1]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "CheckMenuItem")
	_gret := _info.InvokeClassMethod("new_CheckMenuItem_with_mnemonic", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(label)

	var _checkMenuItem *CheckMenuItem // out

	_checkMenuItem = wrapCheckMenuItem(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _checkMenuItem
}

// Active returns whether the check menu item is active. See
// gtk_check_menu_item_set_active ().
//
// The function returns the following values:
//
//    - ok: TRUE if the menu item is checked.
//
func (checkMenuItem *CheckMenuItem) Active() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(checkMenuItem).Native()))

	_info := girepository.MustFind("Gtk", "CheckMenuItem")
	_gret := _info.InvokeClassMethod("get_active", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(checkMenuItem)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// DrawAsRadio returns whether check_menu_item looks like a RadioMenuItem.
//
// The function returns the following values:
//
//    - ok: whether check_menu_item looks like a RadioMenuItem.
//
func (checkMenuItem *CheckMenuItem) DrawAsRadio() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(checkMenuItem).Native()))

	_info := girepository.MustFind("Gtk", "CheckMenuItem")
	_gret := _info.InvokeClassMethod("get_draw_as_radio", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(checkMenuItem)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Inconsistent retrieves the value set by
// gtk_check_menu_item_set_inconsistent().
//
// The function returns the following values:
//
//    - ok: TRUE if inconsistent.
//
func (checkMenuItem *CheckMenuItem) Inconsistent() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(checkMenuItem).Native()))

	_info := girepository.MustFind("Gtk", "CheckMenuItem")
	_gret := _info.InvokeClassMethod("get_inconsistent", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(checkMenuItem)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetActive sets the active state of the menu item’s check box.
//
// The function takes the following parameters:
//
//    - isActive: boolean value indicating whether the check box is active.
//
func (checkMenuItem *CheckMenuItem) SetActive(isActive bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(checkMenuItem).Native()))
	if isActive {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "CheckMenuItem")
	_info.InvokeClassMethod("set_active", _args[:], nil)

	runtime.KeepAlive(checkMenuItem)
	runtime.KeepAlive(isActive)
}

// SetDrawAsRadio sets whether check_menu_item is drawn like a RadioMenuItem.
//
// The function takes the following parameters:
//
//    - drawAsRadio: whether check_menu_item is drawn like a RadioMenuItem.
//
func (checkMenuItem *CheckMenuItem) SetDrawAsRadio(drawAsRadio bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(checkMenuItem).Native()))
	if drawAsRadio {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "CheckMenuItem")
	_info.InvokeClassMethod("set_draw_as_radio", _args[:], nil)

	runtime.KeepAlive(checkMenuItem)
	runtime.KeepAlive(drawAsRadio)
}

// SetInconsistent: if the user has selected a range of elements (such as some
// text or spreadsheet cells) that are affected by a boolean setting, and the
// current values in that range are inconsistent, you may want to display the
// check in an “in between” state. This function turns on “in between” display.
// Normally you would turn off the inconsistent state again if the user
// explicitly selects a setting. This has to be done manually,
// gtk_check_menu_item_set_inconsistent() only affects visual appearance, it
// doesn’t affect the semantics of the widget.
//
// The function takes the following parameters:
//
//    - setting: TRUE to display an “inconsistent” third state check.
//
func (checkMenuItem *CheckMenuItem) SetInconsistent(setting bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(checkMenuItem).Native()))
	if setting {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "CheckMenuItem")
	_info.InvokeClassMethod("set_inconsistent", _args[:], nil)

	runtime.KeepAlive(checkMenuItem)
	runtime.KeepAlive(setting)
}

// Toggled emits the CheckMenuItem::toggled signal.
func (checkMenuItem *CheckMenuItem) Toggled() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(checkMenuItem).Native()))

	_info := girepository.MustFind("Gtk", "CheckMenuItem")
	_info.InvokeClassMethod("toggled", _args[:], nil)

	runtime.KeepAlive(checkMenuItem)
}
