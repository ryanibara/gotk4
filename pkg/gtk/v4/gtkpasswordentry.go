// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_PasswordEntry_ConnectActivate(gpointer, guintptr);
import "C"

// GTypePasswordEntry returns the GType for the type PasswordEntry.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypePasswordEntry() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "PasswordEntry").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalPasswordEntry)
	return gtype
}

// PasswordEntry: GtkPasswordEntry is an entry that has been tailored for
// entering secrets.
//
// !An example GtkPasswordEntry (password-entry.png)
//
// It does not show its contents in clear text, does not allow to copy it to the
// clipboard, and it shows a warning when Caps Lock is engaged. If the
// underlying platform allows it, GtkPasswordEntry will also place the text in a
// non-pageable memory area, to avoid it being written out to disk by the
// operating system.
//
// Optionally, it can offer a way to reveal the contents in clear text.
//
// GtkPasswordEntry provides only minimal API and should be used with the
// gtk.Editable API.
//
// CSS Nodes
//
//    entry.password
//    ╰── text
//        ├── image.caps-lock-indicator
//        ┊
//
//
// GtkPasswordEntry has a single CSS node with name entry that carries a
// .passwordstyle class. The text Css node below it has a child with name image
// and style class .caps-lock-indicator for the Caps Lock icon, and possibly
// other children.
//
//
// Accessibility
//
// GtkPasswordEntry uses the GTK_ACCESSIBLE_ROLE_TEXT_BOX role.
type PasswordEntry struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Editable
}

var (
	_ Widgetter         = (*PasswordEntry)(nil)
	_ coreglib.Objector = (*PasswordEntry)(nil)
)

func wrapPasswordEntry(obj *coreglib.Object) *PasswordEntry {
	return &PasswordEntry{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		Editable: Editable{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
		},
	}
}

func marshalPasswordEntry(p uintptr) (interface{}, error) {
	return wrapPasswordEntry(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_PasswordEntry_ConnectActivate
func _gotk4_gtk4_PasswordEntry_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectActivate is emitted when the entry is activated.
//
// The keybindings for this signal are all forms of the Enter key.
func (entry *PasswordEntry) ConnectActivate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(entry, "activate", false, unsafe.Pointer(C._gotk4_gtk4_PasswordEntry_ConnectActivate), f)
}

// NewPasswordEntry creates a GtkPasswordEntry.
//
// The function returns the following values:
//
//    - passwordEntry: new GtkPasswordEntry.
//
func NewPasswordEntry() *PasswordEntry {
	_info := girepository.MustFind("Gtk", "PasswordEntry")
	_gret := _info.InvokeClassMethod("new_PasswordEntry", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _passwordEntry *PasswordEntry // out

	_passwordEntry = wrapPasswordEntry(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _passwordEntry
}

// ExtraMenu gets the menu model set with gtk_password_entry_set_extra_menu().
//
// The function returns the following values:
//
//    - menuModel: (nullable): the menu model.
//
func (entry *PasswordEntry) ExtraMenu() gio.MenuModeller {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_info := girepository.MustFind("Gtk", "PasswordEntry")
	_gret := _info.InvokeClassMethod("get_extra_menu", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _menuModel gio.MenuModeller // out

	{
		objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))
		if objptr == nil {
			panic("object of type gio.MenuModeller is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gio.MenuModeller)
			return ok
		})
		rv, ok := casted.(gio.MenuModeller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.MenuModeller")
		}
		_menuModel = rv
	}

	return _menuModel
}

// ShowPeekIcon returns whether the entry is showing an icon to reveal the
// contents.
//
// The function returns the following values:
//
//    - ok: TRUE if an icon is shown.
//
func (entry *PasswordEntry) ShowPeekIcon() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_info := girepository.MustFind("Gtk", "PasswordEntry")
	_gret := _info.InvokeClassMethod("get_show_peek_icon", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetExtraMenu sets a menu model to add when constructing the context menu for
// entry.
//
// The function takes the following parameters:
//
//    - model (optional): GMenuModel.
//
func (entry *PasswordEntry) SetExtraMenu(model gio.MenuModeller) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	if model != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	_info := girepository.MustFind("Gtk", "PasswordEntry")
	_info.InvokeClassMethod("set_extra_menu", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(model)
}

// SetShowPeekIcon sets whether the entry should have a clickable icon to reveal
// the contents.
//
// Setting this to FALSE also hides the text again.
//
// The function takes the following parameters:
//
//    - showPeekIcon: whether to show the peek icon.
//
func (entry *PasswordEntry) SetShowPeekIcon(showPeekIcon bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	if showPeekIcon {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "PasswordEntry")
	_info.InvokeClassMethod("set_show_peek_icon", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(showPeekIcon)
}
