// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_buttons_type_get_type()), F: marshalButtonsType},
		{T: externglib.Type(C.gtk_message_dialog_get_type()), F: marshalMessageDialogger},
	})
}

// ButtonsType: prebuilt sets of buttons for the dialog. If none of these
// choices are appropriate, simply use GTK_BUTTONS_NONE then call
// gtk_dialog_add_buttons().
//
// > Please note that GTK_BUTTONS_OK, GTK_BUTTONS_YES_NO > and
// GTK_BUTTONS_OK_CANCEL are discouraged by the > GNOME Human Interface
// Guidelines (http://library.gnome.org/devel/hig-book/stable/).
type ButtonsType int

const (
	// ButtonsNone: no buttons at all
	ButtonsNone ButtonsType = iota
	// ButtonsOK: OK button
	ButtonsOK
	// ButtonsClose: close button
	ButtonsClose
	// ButtonsCancel: cancel button
	ButtonsCancel
	// ButtonsYesNo yes and No buttons
	ButtonsYesNo
	// ButtonsOKCancel: OK and Cancel buttons
	ButtonsOKCancel
)

func marshalButtonsType(p uintptr) (interface{}, error) {
	return ButtonsType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for ButtonsType.
func (b ButtonsType) String() string {
	switch b {
	case ButtonsNone:
		return "None"
	case ButtonsOK:
		return "OK"
	case ButtonsClose:
		return "Close"
	case ButtonsCancel:
		return "Cancel"
	case ButtonsYesNo:
		return "YesNo"
	case ButtonsOKCancel:
		return "OKCancel"
	default:
		return fmt.Sprintf("ButtonsType(%d)", b)
	}
}

// MessageDialog presents a dialog with some message text. It’s simply a
// convenience widget; you could construct the equivalent of MessageDialog from
// Dialog without too much effort, but MessageDialog saves typing.
//
// One difference from Dialog is that MessageDialog sets the
// Window:skip-taskbar-hint property to TRUE, so that the dialog is hidden from
// the taskbar by default.
//
// The easiest way to do a modal message dialog is to use gtk_dialog_run(),
// though you can also pass in the GTK_DIALOG_MODAL flag, gtk_dialog_run()
// automatically makes the dialog modal and waits for the user to respond to it.
// gtk_dialog_run() returns when any dialog button is clicked.
//
// An example for using a modal dialog:
//
//     GtkDialogFlags flags = GTK_DIALOG_DESTROY_WITH_PARENT;
//     dialog = gtk_message_dialog_new (parent_window,
//                                      flags,
//                                      GTK_MESSAGE_ERROR,
//                                      GTK_BUTTONS_CLOSE,
//                                      "Error reading “s”: s",
//                                      filename,
//                                      g_strerror (errno));
//
//     // Destroy the dialog when the user responds to it
//     // (e.g. clicks a button)
//
//     g_signal_connect_swapped (dialog, "response",
//                               G_CALLBACK (gtk_widget_destroy),
//                               dialog);
//
//
// GtkMessageDialog as GtkBuildable
//
// The GtkMessageDialog implementation of the GtkBuildable interface exposes the
// message area as an internal child with the name “message_area”.
type MessageDialog struct {
	Dialog
}

func wrapMessageDialog(obj *externglib.Object) *MessageDialog {
	return &MessageDialog{
		Dialog: Dialog{
			Window: Window{
				Bin: Bin{
					Container: Container{
						Widget: Widget{
							InitiallyUnowned: externglib.InitiallyUnowned{
								Object: obj,
							},
							ImplementorIface: atk.ImplementorIface{
								Object: obj,
							},
							Buildable: Buildable{
								Object: obj,
							},
							Object: obj,
						},
					},
				},
			},
		},
	}
}

func marshalMessageDialogger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMessageDialog(obj), nil
}

// Image gets the dialog’s image.
//
// Deprecated: Use Dialog for dialogs with images.
func (dialog *MessageDialog) Image() Widgetter {
	var _arg0 *C.GtkMessageDialog // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkMessageDialog)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_message_dialog_get_image(_arg0)

	var _widget Widgetter // out

	_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// MessageArea returns the message area of the dialog. This is the box where the
// dialog’s primary and secondary labels are packed. You can add your own extra
// content to that box and it will appear below those labels. See
// gtk_dialog_get_content_area() for the corresponding function in the parent
// Dialog.
func (messageDialog *MessageDialog) MessageArea() Widgetter {
	var _arg0 *C.GtkMessageDialog // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkMessageDialog)(unsafe.Pointer(messageDialog.Native()))

	_cret = C.gtk_message_dialog_get_message_area(_arg0)

	var _widget Widgetter // out

	_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// SetImage sets the dialog’s image to image.
//
// Deprecated: Use Dialog to create dialogs with images.
func (dialog *MessageDialog) SetImage(image Widgetter) {
	var _arg0 *C.GtkMessageDialog // out
	var _arg1 *C.GtkWidget        // out

	_arg0 = (*C.GtkMessageDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(image.Native()))

	C.gtk_message_dialog_set_image(_arg0, _arg1)
}

// SetMarkup sets the text of the message dialog to be str, which is marked up
// with the [Pango text markup language][PangoMarkupFormat].
func (messageDialog *MessageDialog) SetMarkup(str string) {
	var _arg0 *C.GtkMessageDialog // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkMessageDialog)(unsafe.Pointer(messageDialog.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_message_dialog_set_markup(_arg0, _arg1)
}
