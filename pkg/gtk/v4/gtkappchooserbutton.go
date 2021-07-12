// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_app_chooser_button_get_type()), F: marshalAppChooserButtoner},
	})
}

// AppChooserButtoner describes AppChooserButton's methods.
type AppChooserButtoner interface {
	// AppendCustomItem appends a custom item to the list of applications that
	// is shown in the popup.
	AppendCustomItem(name string, label string, icon gio.Iconer)
	// AppendSeparator appends a separator to the list of applications that is
	// shown in the popup.
	AppendSeparator()
	// Heading returns the text to display at the top of the dialog.
	Heading() string
	// Modal gets whether the dialog is modal.
	Modal() bool
	// ShowDefaultItem returns whether the dropdown menu should show the default
	// application at the top.
	ShowDefaultItem() bool
	// ShowDialogItem returns whether the dropdown menu shows an item for a
	// `GtkAppChooserDialog`.
	ShowDialogItem() bool
	// SetActiveCustomItem selects a custom item.
	SetActiveCustomItem(name string)
	// SetHeading sets the text to display at the top of the dialog.
	SetHeading(heading string)
	// SetModal sets whether the dialog should be modal.
	SetModal(modal bool)
	// SetShowDefaultItem sets whether the dropdown menu of this button should
	// show the default application for the given content type at top.
	SetShowDefaultItem(setting bool)
	// SetShowDialogItem sets whether the dropdown menu of this button should
	// show an entry to trigger a `GtkAppChooserDialog`.
	SetShowDialogItem(setting bool)
}

// AppChooserButton: `GtkAppChooserButton` lets the user select an application.
//
// !An example GtkAppChooserButton (appchooserbutton.png)
//
// Initially, a `GtkAppChooserButton` selects the first application in its list,
// which will either be the most-recently used application or, if
// [property@Gtk.AppChooserButton:show-default-item] is true, the default
// application.
//
// The list of applications shown in a `GtkAppChooserButton` includes the
// recommended applications for the given content type. When
// [property@Gtk.AppChooserButton:show-default-item] is set, the default
// application is also included. To let the user chooser other applications, you
// can set the [property@Gtk.AppChooserButton:show-dialog-item] property, which
// allows to open a full [class@Gtk.AppChooserDialog].
//
// It is possible to add custom items to the list, using
// [method@Gtk.AppChooserButton.append_custom_item]. These items cause the
// [signal@Gtk.AppChooserButton::custom-item-activated] signal to be emitted
// when they are selected.
//
// To track changes in the selected application, use the
// [signal@Gtk.AppChooserButton::changed] signal.
//
//
// CSS nodes
//
// `GtkAppChooserButton` has a single CSS node with the name “appchooserbutton”.
type AppChooserButton struct {
	Widget

	AppChooser
}

var (
	_ AppChooserButtoner = (*AppChooserButton)(nil)
	_ gextras.Nativer    = (*AppChooserButton)(nil)
)

func wrapAppChooserButton(obj *externglib.Object) *AppChooserButton {
	return &AppChooserButton{
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
		},
		AppChooser: AppChooser{
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
			},
		},
	}
}

func marshalAppChooserButtoner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAppChooserButton(obj), nil
}

// NewAppChooserButton creates a new `GtkAppChooserButton` for applications that
// can handle content of the given type.
func NewAppChooserButton(contentType string) *AppChooserButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(contentType)))

	_cret = C.gtk_app_chooser_button_new(_arg1)

	var _appChooserButton *AppChooserButton // out

	_appChooserButton = wrapAppChooserButton(externglib.Take(unsafe.Pointer(_cret)))

	return _appChooserButton
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *AppChooserButton) Native() uintptr {
	return v.Widget.InitiallyUnowned.Object.Native()
}

// AppendCustomItem appends a custom item to the list of applications that is
// shown in the popup.
//
// The item name must be unique per-widget. Clients can use the provided name as
// a detail for the [signal@Gtk.AppChooserButton::custom-item-activated] signal,
// to add a callback for the activation of a particular custom item in the list.
//
// See also [method@Gtk.AppChooserButton.append_separator].
func (self *AppChooserButton) AppendCustomItem(name string, label string, icon gio.Iconer) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 *C.char                // out
	var _arg2 *C.char                // out
	var _arg3 *C.GIcon               // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(label)))
	_arg3 = (*C.GIcon)(unsafe.Pointer((icon).(gextras.Nativer).Native()))

	C.gtk_app_chooser_button_append_custom_item(_arg0, _arg1, _arg2, _arg3)
}

// AppendSeparator appends a separator to the list of applications that is shown
// in the popup.
func (self *AppChooserButton) AppendSeparator() {
	var _arg0 *C.GtkAppChooserButton // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))

	C.gtk_app_chooser_button_append_separator(_arg0)
}

// Heading returns the text to display at the top of the dialog.
func (self *AppChooserButton) Heading() string {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret *C.char                // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_button_get_heading(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Modal gets whether the dialog is modal.
func (self *AppChooserButton) Modal() bool {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_button_get_modal(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowDefaultItem returns whether the dropdown menu should show the default
// application at the top.
func (self *AppChooserButton) ShowDefaultItem() bool {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_button_get_show_default_item(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowDialogItem returns whether the dropdown menu shows an item for a
// `GtkAppChooserDialog`.
func (self *AppChooserButton) ShowDialogItem() bool {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_button_get_show_dialog_item(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActiveCustomItem selects a custom item.
//
// See [method@Gtk.AppChooserButton.append_custom_item].
//
// Use [method@Gtk.AppChooser.refresh] to bring the selection to its initial
// state.
func (self *AppChooserButton) SetActiveCustomItem(name string) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 *C.char                // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))

	C.gtk_app_chooser_button_set_active_custom_item(_arg0, _arg1)
}

// SetHeading sets the text to display at the top of the dialog.
//
// If the heading is not set, the dialog displays a default text.
func (self *AppChooserButton) SetHeading(heading string) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 *C.char                // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(heading)))

	C.gtk_app_chooser_button_set_heading(_arg0, _arg1)
}

// SetModal sets whether the dialog should be modal.
func (self *AppChooserButton) SetModal(modal bool) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_button_set_modal(_arg0, _arg1)
}

// SetShowDefaultItem sets whether the dropdown menu of this button should show
// the default application for the given content type at top.
func (self *AppChooserButton) SetShowDefaultItem(setting bool) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_button_set_show_default_item(_arg0, _arg1)
}

// SetShowDialogItem sets whether the dropdown menu of this button should show
// an entry to trigger a `GtkAppChooserDialog`.
func (self *AppChooserButton) SetShowDialogItem(setting bool) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_button_set_show_dialog_item(_arg0, _arg1)
}
