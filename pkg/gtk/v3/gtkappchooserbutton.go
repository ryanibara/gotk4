// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_app_chooser_button_get_type()), F: marshalAppChooserButton},
	})
}

// AppChooserButton: the AppChooserButton is a widget that lets the user select
// an application. It implements the AppChooser interface.
//
// Initially, a AppChooserButton selects the first application in its list,
// which will either be the most-recently used application or, if
// AppChooserButton:show-default-item is true, the default application.
//
// The list of applications shown in a AppChooserButton includes the recommended
// applications for the given content type. When
// AppChooserButton:show-default-item is set, the default application is also
// included. To let the user chooser other applications, you can set the
// AppChooserButton:show-dialog-item property, which allows to open a full
// AppChooserDialog.
//
// It is possible to add custom items to the list, using
// gtk_app_chooser_button_append_custom_item(). These items cause the
// AppChooserButton::custom-item-activated signal to be emitted when they are
// selected.
//
// To track changes in the selected application, use the ComboBox::changed
// signal.
type AppChooserButton interface {
	ComboBox
	AppChooser
	Buildable
	CellEditable
	CellLayout

	// AppendCustomItem appends a custom item to the list of applications that
	// is shown in the popup; the item name must be unique per-widget. Clients
	// can use the provided name as a detail for the
	// AppChooserButton::custom-item-activated signal, to add a callback for the
	// activation of a particular custom item in the list. See also
	// gtk_app_chooser_button_append_separator().
	AppendCustomItem(name string, label string, icon gio.Icon)
	// AppendSeparator appends a separator to the list of applications that is
	// shown in the popup.
	AppendSeparator()
	// Heading returns the text to display at the top of the dialog.
	Heading() string
	// ShowDefaultItem returns the current value of the
	// AppChooserButton:show-default-item property.
	ShowDefaultItem() bool
	// ShowDialogItem returns the current value of the
	// AppChooserButton:show-dialog-item property.
	ShowDialogItem() bool
	// SetActiveCustomItem selects a custom item previously added with
	// gtk_app_chooser_button_append_custom_item().
	//
	// Use gtk_app_chooser_refresh() to bring the selection to its initial
	// state.
	SetActiveCustomItem(name string)
	// SetHeading sets the text to display at the top of the dialog. If the
	// heading is not set, the dialog displays a default text.
	SetHeading(heading string)
	// SetShowDefaultItem sets whether the dropdown menu of this button should
	// show the default application for the given content type at top.
	SetShowDefaultItem(setting bool)
	// SetShowDialogItem sets whether the dropdown menu of this button should
	// show an entry to trigger a AppChooserDialog.
	SetShowDialogItem(setting bool)
}

// appChooserButton implements the AppChooserButton interface.
type appChooserButton struct {
	ComboBox
	AppChooser
	Buildable
	CellEditable
	CellLayout
}

var _ AppChooserButton = (*appChooserButton)(nil)

// WrapAppChooserButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapAppChooserButton(obj *externglib.Object) AppChooserButton {
	return AppChooserButton{
		ComboBox:     WrapComboBox(obj),
		AppChooser:   WrapAppChooser(obj),
		Buildable:    WrapBuildable(obj),
		CellEditable: WrapCellEditable(obj),
		CellLayout:   WrapCellLayout(obj),
	}
}

func marshalAppChooserButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppChooserButton(obj), nil
}

// NewAppChooserButton constructs a class AppChooserButton.
func NewAppChooserButton(contentType string) AppChooserButton {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(contentType))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GtkAppChooserButton

	cret = C.gtk_app_chooser_button_new(_arg1)

	var _appChooserButton AppChooserButton

	_appChooserButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(AppChooserButton)

	return _appChooserButton
}

// AppendCustomItem appends a custom item to the list of applications that
// is shown in the popup; the item name must be unique per-widget. Clients
// can use the provided name as a detail for the
// AppChooserButton::custom-item-activated signal, to add a callback for the
// activation of a particular custom item in the list. See also
// gtk_app_chooser_button_append_separator().
func (s appChooserButton) AppendCustomItem(name string, label string, icon gio.Icon) {
	var _arg0 *C.GtkAppChooserButton
	var _arg1 *C.gchar
	var _arg2 *C.gchar
	var _arg3 *C.GIcon

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	C.gtk_app_chooser_button_append_custom_item(_arg0, _arg1, _arg2, _arg3)
}

// AppendSeparator appends a separator to the list of applications that is
// shown in the popup.
func (s appChooserButton) AppendSeparator() {
	var _arg0 *C.GtkAppChooserButton

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	C.gtk_app_chooser_button_append_separator(_arg0)
}

// Heading returns the text to display at the top of the dialog.
func (s appChooserButton) Heading() string {
	var _arg0 *C.GtkAppChooserButton

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	var _cret *C.gchar

	cret = C.gtk_app_chooser_button_get_heading(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// ShowDefaultItem returns the current value of the
// AppChooserButton:show-default-item property.
func (s appChooserButton) ShowDefaultItem() bool {
	var _arg0 *C.GtkAppChooserButton

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	cret = C.gtk_app_chooser_button_get_show_default_item(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ShowDialogItem returns the current value of the
// AppChooserButton:show-dialog-item property.
func (s appChooserButton) ShowDialogItem() bool {
	var _arg0 *C.GtkAppChooserButton

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	cret = C.gtk_app_chooser_button_get_show_dialog_item(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetActiveCustomItem selects a custom item previously added with
// gtk_app_chooser_button_append_custom_item().
//
// Use gtk_app_chooser_refresh() to bring the selection to its initial
// state.
func (s appChooserButton) SetActiveCustomItem(name string) {
	var _arg0 *C.GtkAppChooserButton
	var _arg1 *C.gchar

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_button_set_active_custom_item(_arg0, _arg1)
}

// SetHeading sets the text to display at the top of the dialog. If the
// heading is not set, the dialog displays a default text.
func (s appChooserButton) SetHeading(heading string) {
	var _arg0 *C.GtkAppChooserButton
	var _arg1 *C.gchar

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(heading))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_button_set_heading(_arg0, _arg1)
}

// SetShowDefaultItem sets whether the dropdown menu of this button should
// show the default application for the given content type at top.
func (s appChooserButton) SetShowDefaultItem(setting bool) {
	var _arg0 *C.GtkAppChooserButton
	var _arg1 C.gboolean

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.gboolean(1)
	}

	C.gtk_app_chooser_button_set_show_default_item(_arg0, _arg1)
}

// SetShowDialogItem sets whether the dropdown menu of this button should
// show an entry to trigger a AppChooserDialog.
func (s appChooserButton) SetShowDialogItem(setting bool) {
	var _arg0 *C.GtkAppChooserButton
	var _arg1 C.gboolean

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.gboolean(1)
	}

	C.gtk_app_chooser_button_set_show_dialog_item(_arg0, _arg1)
}