// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_message_dialog_get_type()), F: marshalMessageDialog},
	})
}

// MessageDialog presents a dialog with some message text. It’s simply a
// convenience widget; you could construct the equivalent of MessageDialog from
// Dialog without too much effort, but MessageDialog saves typing.
//
// The easiest way to do a modal message dialog is to use the GTK_DIALOG_MODAL
// flag, which will call gtk_window_set_modal() internally. The dialog will
// prevent interaction with the parent window until it's hidden or destroyed.
// You can use the Dialog::response signal to know when the user dismissed the
// dialog.
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
//     g_signal_connect (dialog, "response",
//                       G_CALLBACK (gtk_window_destroy),
//                       NULL);
//
//
// GtkMessageDialog as GtkBuildable
//
// The GtkMessageDialog implementation of the GtkBuildable interface exposes the
// message area as an internal child with the name “message_area”.
type MessageDialog interface {
	Dialog
	Accessible
	Buildable
	ConstraintTarget
	Native
	Root
	ShortcutManager

	// MessageArea returns the message area of the dialog. This is the box where
	// the dialog’s primary and secondary labels are packed. You can add your
	// own extra content to that box and it will appear below those labels. See
	// gtk_dialog_get_content_area() for the corresponding function in the
	// parent Dialog.
	MessageArea() Widget
	// SetMarkup sets the text of the message dialog to be @str, which is marked
	// up with the [Pango text markup language][PangoMarkupFormat].
	SetMarkup(str string)
}

// messageDialog implements the MessageDialog interface.
type messageDialog struct {
	Dialog
	Accessible
	Buildable
	ConstraintTarget
	Native
	Root
	ShortcutManager
}

var _ MessageDialog = (*messageDialog)(nil)

// WrapMessageDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapMessageDialog(obj *externglib.Object) MessageDialog {
	return MessageDialog{
		Dialog:           WrapDialog(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Native:           WrapNative(obj),
		Root:             WrapRoot(obj),
		ShortcutManager:  WrapShortcutManager(obj),
	}
}

func marshalMessageDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMessageDialog(obj), nil
}

// MessageArea returns the message area of the dialog. This is the box where
// the dialog’s primary and secondary labels are packed. You can add your
// own extra content to that box and it will appear below those labels. See
// gtk_dialog_get_content_area() for the corresponding function in the
// parent Dialog.
func (m messageDialog) MessageArea() Widget {
	var arg0 *C.GtkMessageDialog

	arg0 = (*C.GtkMessageDialog)(unsafe.Pointer(m.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_message_dialog_get_message_area(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// SetMarkup sets the text of the message dialog to be @str, which is marked
// up with the [Pango text markup language][PangoMarkupFormat].
func (m messageDialog) SetMarkup(str string) {
	var arg0 *C.GtkMessageDialog
	var arg1 *C.char

	arg0 = (*C.GtkMessageDialog)(unsafe.Pointer(m.Native()))
	arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_message_dialog_set_markup(arg0, str)
}
