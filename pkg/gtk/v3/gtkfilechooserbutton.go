// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_file_chooser_button_get_type()), F: marshalFileChooserButton},
	})
}

// FileChooserButton: the FileChooserButton is a widget that lets the user
// select a file. It implements the FileChooser interface. Visually, it is a
// file name with a button to bring up a FileChooserDialog. The user can then
// use that dialog to change the file associated with that button. This widget
// does not support setting the FileChooser:select-multiple property to true.
//
// Create a button to let the user select a file in /etc
//
//    {
//      GtkWidget *button;
//
//      button = gtk_file_chooser_button_new (_("Select a file"),
//                                            GTK_FILE_CHOOSER_ACTION_OPEN);
//      gtk_file_chooser_set_current_folder (GTK_FILE_CHOOSER (button),
//                                           "/etc");
//    }
//
// The FileChooserButton supports the FileChooserActions
// GTK_FILE_CHOOSER_ACTION_OPEN and GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER.
//
// > The FileChooserButton will ellipsize the label, and will thus > request
// little horizontal space. To give the button more space, > you should call
// gtk_widget_get_preferred_size(), > gtk_file_chooser_button_set_width_chars(),
// or pack the button in > such a way that other interface elements give space
// to the > widget.
//
//
// CSS nodes
//
// GtkFileChooserButton has a CSS node with name “filechooserbutton”, containing
// a subnode for the internal button with name “button” and style class “.file”.
type FileChooserButton interface {
	Box
	Buildable
	FileChooser
	Orientable

	// FocusOnClick returns whether the button grabs focus when it is clicked
	// with the mouse. See gtk_file_chooser_button_set_focus_on_click().
	FocusOnClick() bool
	// Title retrieves the title of the browse dialog used by @button. The
	// returned value should not be modified or freed.
	Title() string
	// WidthChars retrieves the width in characters of the @button widget’s
	// entry and/or label.
	WidthChars() int
	// SetFocusOnClick sets whether the button will grab focus when it is
	// clicked with the mouse. Making mouse clicks not grab focus is useful in
	// places like toolbars where you don’t want the keyboard focus removed from
	// the main area of the application.
	SetFocusOnClick(focusOnClick bool)
	// SetTitle modifies the @title of the browse dialog used by @button.
	SetTitle(title string)
	// SetWidthChars sets the width (in characters) that @button will use to
	// @n_chars.
	SetWidthChars(nChars int)
}

// fileChooserButton implements the FileChooserButton interface.
type fileChooserButton struct {
	Box
	Buildable
	FileChooser
	Orientable
}

var _ FileChooserButton = (*fileChooserButton)(nil)

// WrapFileChooserButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileChooserButton(obj *externglib.Object) FileChooserButton {
	return FileChooserButton{
		Box:         WrapBox(obj),
		Buildable:   WrapBuildable(obj),
		FileChooser: WrapFileChooser(obj),
		Orientable:  WrapOrientable(obj),
	}
}

func marshalFileChooserButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileChooserButton(obj), nil
}

// FocusOnClick returns whether the button grabs focus when it is clicked
// with the mouse. See gtk_file_chooser_button_set_focus_on_click().
func (b fileChooserButton) FocusOnClick() bool {
	var _arg0 *C.GtkFileChooserButton

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(b.Native()))

	var _cret C.gboolean

	_cret = C.gtk_file_chooser_button_get_focus_on_click(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Title retrieves the title of the browse dialog used by @button. The
// returned value should not be modified or freed.
func (b fileChooserButton) Title() string {
	var _arg0 *C.GtkFileChooserButton

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(b.Native()))

	var _cret *C.gchar

	_cret = C.gtk_file_chooser_button_get_title(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// WidthChars retrieves the width in characters of the @button widget’s
// entry and/or label.
func (b fileChooserButton) WidthChars() int {
	var _arg0 *C.GtkFileChooserButton

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(b.Native()))

	var _cret C.gint

	_cret = C.gtk_file_chooser_button_get_width_chars(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// SetFocusOnClick sets whether the button will grab focus when it is
// clicked with the mouse. Making mouse clicks not grab focus is useful in
// places like toolbars where you don’t want the keyboard focus removed from
// the main area of the application.
func (b fileChooserButton) SetFocusOnClick(focusOnClick bool) {
	var _arg0 *C.GtkFileChooserButton
	var _arg1 C.gboolean

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(b.Native()))
	if focusOnClick {
		_arg1 = C.gboolean(1)
	}

	C.gtk_file_chooser_button_set_focus_on_click(_arg0, _arg1)
}

// SetTitle modifies the @title of the browse dialog used by @button.
func (b fileChooserButton) SetTitle(title string) {
	var _arg0 *C.GtkFileChooserButton
	var _arg1 *C.gchar

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_chooser_button_set_title(_arg0, _arg1)
}

// SetWidthChars sets the width (in characters) that @button will use to
// @n_chars.
func (b fileChooserButton) SetWidthChars(nChars int) {
	var _arg0 *C.GtkFileChooserButton
	var _arg1 C.gint

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(b.Native()))
	_arg1 = C.gint(nChars)

	C.gtk_file_chooser_button_set_width_chars(_arg0, _arg1)
}
