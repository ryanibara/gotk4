// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_password_entry_get_type()), F: marshalPasswordEntry},
	})
}

// PasswordEntry: `GtkPasswordEntry` is an entry that has been tailored for
// entering secrets.
//
// !An example GtkPasswordEntry (password-entry.png)
//
// It does not show its contents in clear text, does not allow to copy it to the
// clipboard, and it shows a warning when Caps Lock is engaged. If the
// underlying platform allows it, `GtkPasswordEntry` will also place the text in
// a non-pageable memory area, to avoid it being written out to disk by the
// operating system.
//
// Optionally, it can offer a way to reveal the contents in clear text.
//
// `GtkPasswordEntry` provides only minimal API and should be used with the
// [iface@Gtk.Editable] API.
//
//
// CSS Nodes
//
// “` entry.password ╰── text ├── image.caps-lock-indicator ┊ “`
//
// `GtkPasswordEntry` has a single CSS node with name entry that carries a
// .passwordstyle class. The text Css node below it has a child with name image
// and style class .caps-lock-indicator for the Caps Lock icon, and possibly
// other children.
//
//
// Accessibility
//
// `GtkPasswordEntry` uses the GTK_ACCESSIBLE_ROLE_TEXT_BOX role.
type PasswordEntry interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Editable

	// ExtraMenu gets the menu model set with
	// gtk_password_entry_set_extra_menu().
	ExtraMenu() gio.MenuModel
	// ShowPeekIcon returns whether the entry is showing an icon to reveal the
	// contents.
	ShowPeekIcon() bool
	// SetExtraMenu sets a menu model to add when constructing the context menu
	// for @entry.
	SetExtraMenu(model gio.MenuModel)
	// SetShowPeekIcon sets whether the entry should have a clickable icon to
	// reveal the contents.
	//
	// Setting this to false also hides the text again.
	SetShowPeekIcon(showPeekIcon bool)
}

// passwordEntry implements the PasswordEntry interface.
type passwordEntry struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Editable
}

var _ PasswordEntry = (*passwordEntry)(nil)

// WrapPasswordEntry wraps a GObject to the right type. It is
// primarily used internally.
func WrapPasswordEntry(obj *externglib.Object) PasswordEntry {
	return PasswordEntry{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Editable:         WrapEditable(obj),
	}
}

func marshalPasswordEntry(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPasswordEntry(obj), nil
}

// NewPasswordEntry constructs a class PasswordEntry.
func NewPasswordEntry() PasswordEntry {
	ret := C.gtk_password_entry_new()

	var ret0 PasswordEntry

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(PasswordEntry)

	return ret0
}

// ExtraMenu gets the menu model set with
// gtk_password_entry_set_extra_menu().
func (e passwordEntry) ExtraMenu() gio.MenuModel {
	var arg0 *C.GtkPasswordEntry

	arg0 = (*C.GtkPasswordEntry)(e.Native())

	ret := C.gtk_password_entry_get_extra_menu(arg0)

	var ret0 gio.MenuModel

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(gio.MenuModel)

	return ret0
}

// ShowPeekIcon returns whether the entry is showing an icon to reveal the
// contents.
func (e passwordEntry) ShowPeekIcon() bool {
	var arg0 *C.GtkPasswordEntry

	arg0 = (*C.GtkPasswordEntry)(e.Native())

	ret := C.gtk_password_entry_get_show_peek_icon(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// SetExtraMenu sets a menu model to add when constructing the context menu
// for @entry.
func (e passwordEntry) SetExtraMenu(model gio.MenuModel) {
	var arg0 *C.GtkPasswordEntry
	var arg1 *C.GMenuModel

	arg0 = (*C.GtkPasswordEntry)(e.Native())
	arg1 = (*C.GMenuModel)(model.Native())

	C.gtk_password_entry_set_extra_menu(arg0, arg1)
}

// SetShowPeekIcon sets whether the entry should have a clickable icon to
// reveal the contents.
//
// Setting this to false also hides the text again.
func (e passwordEntry) SetShowPeekIcon(showPeekIcon bool) {
	var arg0 *C.GtkPasswordEntry
	var arg1 C.gboolean

	arg0 = (*C.GtkPasswordEntry)(e.Native())
	if showPeekIcon {
		arg1 = C.TRUE
	}

	C.gtk_password_entry_set_show_peek_icon(arg0, arg1)
}
