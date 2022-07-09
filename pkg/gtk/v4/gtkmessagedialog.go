// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// GtkWidget* _gotk4_gtk_message_dialog_new2(GtkWindow* parent, GtkDialogFlags flags, GtkMessageType type, GtkButtonsType buttons) {
// 	return gtk_message_dialog_new_with_markup(parent, flags, type, buttons, NULL);
// }
import "C"

// GTypeButtonsType returns the GType for the type ButtonsType.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeButtonsType() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ButtonsType").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalButtonsType)
	return gtype
}

// GTypeMessageDialog returns the GType for the type MessageDialog.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeMessageDialog() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "MessageDialog").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalMessageDialog)
	return gtype
}

// ButtonsType: prebuilt sets of buttons for GtkDialog.
//
// If none of these choices are appropriate, simply use GTK_BUTTONS_NONE and
// call gtk.Dialog.AddButtons().
//
// > Please note that GTK_BUTTONS_OK, GTK_BUTTONS_YES_NO > and
// GTK_BUTTONS_OK_CANCEL are discouraged by the > GNOME Human Interface
// Guidelines (http://library.gnome.org/devel/hig-book/stable/).
type ButtonsType C.gint

const (
	// ButtonsNone: no buttons at all.
	ButtonsNone ButtonsType = iota
	// ButtonsOK: OK button.
	ButtonsOK
	// ButtonsClose: close button.
	ButtonsClose
	// ButtonsCancel: cancel button.
	ButtonsCancel
	// ButtonsYesNo yes and No buttons.
	ButtonsYesNo
	// ButtonsOKCancel: OK and Cancel buttons.
	ButtonsOKCancel
)

func marshalButtonsType(p uintptr) (interface{}, error) {
	return ButtonsType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
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

// MessageDialog: GtkMessageDialog presents a dialog with some message text.
//
// !An example GtkMessageDialog (messagedialog.png)
//
// It’s simply a convenience widget; you could construct the equivalent of
// GtkMessageDialog from GtkDialog without too much effort, but GtkMessageDialog
// saves typing.
//
// The easiest way to do a modal message dialog is to use the GTK_DIALOG_MODAL
// flag, which will call gtk.Window.SetModal() internally. The dialog will
// prevent interaction with the parent window until it's hidden or destroyed.
// You can use the gtk.Dialog::response signal to know when the user dismissed
// the dialog.
//
// An example for using a modal dialog:
//
//    GtkDialogFlags flags = GTK_DIALOG_DESTROY_WITH_PARENT | GTK_DIALOG_MODAL;
//    dialog = gtk_message_dialog_new (parent_window,
//                                     flags,
//                                     GTK_MESSAGE_ERROR,
//                                     GTK_BUTTONS_CLOSE,
//                                     "Error reading “s”: s",
//                                     filename,
//                                     g_strerror (errno));
//    // Destroy the dialog when the user responds to it
//    // (e.g. clicks a button)
//
//    g_signal_connect (dialog, "response",
//                      G_CALLBACK (gtk_window_destroy),
//                      NULL);
//
//
// You might do a non-modal GtkMessageDialog simply by omitting the
// GTK_DIALOG_MODAL flag:
//
//    GtkDialogFlags flags = GTK_DIALOG_DESTROY_WITH_PARENT;
//    dialog = gtk_message_dialog_new (parent_window,
//                                     flags,
//                                     GTK_MESSAGE_ERROR,
//                                     GTK_BUTTONS_CLOSE,
//                                     "Error reading “s”: s",
//                                     filename,
//                                     g_strerror (errno));
//
//    // Destroy the dialog when the user responds to it
//    // (e.g. clicks a button)
//    g_signal_connect (dialog, "response",
//                      G_CALLBACK (gtk_window_destroy),
//                      NULL);
//
//
//
// GtkMessageDialog as GtkBuildable
//
// The GtkMessageDialog implementation of the GtkBuildable interface exposes the
// message area as an internal child with the name “message_area”.
type MessageDialog struct {
	_ [0]func() // equal guard
	Dialog
}

var (
	_ Widgetter         = (*MessageDialog)(nil)
	_ coreglib.Objector = (*MessageDialog)(nil)
)

func wrapMessageDialog(obj *coreglib.Object) *MessageDialog {
	return &MessageDialog{
		Dialog: Dialog{
			Window: Window{
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
				Root: Root{
					NativeSurface: NativeSurface{
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
				},
				ShortcutManager: ShortcutManager{
					Object: obj,
				},
			},
		},
	}
}

func marshalMessageDialog(p uintptr) (interface{}, error) {
	return wrapMessageDialog(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// MessageArea returns the message area of the dialog.
//
// This is the box where the dialog’s primary and secondary labels are packed.
// You can add your own extra content to that box and it will appear below those
// labels. See gtk.Dialog.GetContentArea() for the corresponding function in the
// parent gtk.Dialog.
//
// The function returns the following values:
//
//    - widget: GtkBox corresponding to the “message area” in the message_dialog.
//
func (messageDialog *MessageDialog) MessageArea() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(messageDialog).Native()))

	_info := girepository.MustFind("Gtk", "MessageDialog")
	_gret := _info.InvokeClassMethod("get_message_area", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(messageDialog)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// SetMarkup sets the text of the message dialog.
//
// The function takes the following parameters:
//
//    - str: string with Pango markup.
//
func (messageDialog *MessageDialog) SetMarkup(str string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(messageDialog).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gtk", "MessageDialog")
	_info.InvokeClassMethod("set_markup", _args[:], nil)

	runtime.KeepAlive(messageDialog)
	runtime.KeepAlive(str)
}

// NewMessageDialog creates a new message dialog. This is a simple
// dialog with some text taht the user may want to see. When the user
// clicks a button, a "response" signal is emitted with response IDs
// from ResponseType.
func NewMessageDialog(parent *Window, flags DialogFlags, typ MessageType, buttons ButtonsType) *MessageDialog {
	w := C._gotk4_gtk_message_dialog_new2(
		(*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(parent).Native())),
		(C.GtkDialogFlags)(flags),
		(C.GtkMessageType)(typ),
		(C.GtkButtonsType)(buttons),
	)
	runtime.KeepAlive(parent)

	return wrapMessageDialog(coreglib.Take(unsafe.Pointer(w)))
}
