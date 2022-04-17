// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// glib.Type values for gtkfilechoosernative.go.
var GTypeFileChooserNative = externglib.Type(C.gtk_file_chooser_native_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeFileChooserNative, F: marshalFileChooserNative},
	})
}

// FileChooserNativeOverrider contains methods that are overridable.
type FileChooserNativeOverrider interface {
	externglib.Objector
}

// FileChooserNative: GtkFileChooserNative is an abstraction of a dialog
// suitable for use with “File Open” or “File Save as” commands.
//
// By default, this just uses a GtkFileChooserDialog to implement the actual
// dialog. However, on some platforms, such as Windows and macOS, the native
// platform file chooser is used instead. When the application is running in a
// sandboxed environment without direct filesystem access (such as Flatpak),
// GtkFileChooserNative may call the proper APIs (portals) to let the user
// choose a file and make it available to the application.
//
// While the API of GtkFileChooserNative closely mirrors GtkFileChooserDialog,
// the main difference is that there is no access to any GtkWindow or GtkWidget
// for the dialog. This is required, as there may not be one in the case of a
// platform native dialog.
//
// Showing, hiding and running the dialog is handled by the gtk.NativeDialog
// functions.
//
// Note that unlike GtkFileChooserDialog, GtkFileChooserNative objects are not
// toplevel widgets, and GTK does not keep them alive. It is your responsibility
// to keep a reference until you are done with the object.
//
//
// Typical usage
//
// In the simplest of cases, you can the following code to use
// GtkFileChooserNative to select a file for opening:
//
//    static void
//    on_response (GtkNativeDialog *native,
//                 int              response)
//    {
//      if (response == GTK_RESPONSE_ACCEPT)
//        {
//          GtkFileChooser *chooser = GTK_FILE_CHOOSER (native);
//          GFile *file = gtk_file_chooser_get_file (chooser);
//
//          open_file (file);
//
//          g_object_unref (file);
//        }
//
//      g_object_unref (native);
//    }
//
//      // ...
//      GtkFileChooserNative *native;
//      GtkFileChooserAction action = GTK_FILE_CHOOSER_ACTION_OPEN;
//
//      native = gtk_file_chooser_native_new ("Open File",
//                                            parent_window,
//                                            action,
//                                            "_Open",
//                                            "_Cancel");
//
//      g_signal_connect (native, "response", G_CALLBACK (on_response), NULL);
//      gtk_native_dialog_show (GTK_NATIVE_DIALOG (native));
//
//
// To use a GtkFileChooserNative for saving, you can use this:
//
//    static void
//    on_response (GtkNativeDialog *native,
//                 int              response)
//    {
//      if (response == GTK_RESPONSE_ACCEPT)
//        {
//          GtkFileChooser *chooser = GTK_FILE_CHOOSER (native);
//          GFile *file = gtk_file_chooser_get_file (chooser);
//
//          save_to_file (file);
//
//          g_object_unref (file);
//        }
//
//      g_object_unref (native);
//    }
//
//      // ...
//      GtkFileChooserNative *native;
//      GtkFileChooser *chooser;
//      GtkFileChooserAction action = GTK_FILE_CHOOSER_ACTION_SAVE;
//
//      native = gtk_file_chooser_native_new ("Save File",
//                                            parent_window,
//                                            action,
//                                            "_Save",
//                                            "_Cancel");
//      chooser = GTK_FILE_CHOOSER (native);
//
//      if (user_edited_a_new_document)
//        gtk_file_chooser_set_current_name (chooser, _("Untitled document"));
//      else
//        gtk_file_chooser_set_file (chooser, existing_file, NULL);
//
//      g_signal_connect (native, "response", G_CALLBACK (on_response), NULL);
//      gtk_native_dialog_show (GTK_NATIVE_DIALOG (native));
//
//
// For more information on how to best set up a file dialog, see the
// gtk.FileChooserDialog documentation.
//
//
// Response Codes
//
// GtkFileChooserNative inherits from gtk.NativeDialog, which means it will
// return GTK_RESPONSE_ACCEPT if the user accepted, and GTK_RESPONSE_CANCEL if
// he pressed cancel. It can also return GTK_RESPONSE_DELETE_EVENT if the window
// was unexpectedly closed.
//
//
// Differences from FileChooserDialog
//
// There are a few things in the gtk.FileChooser interface that are not possible
// to use with GtkFileChooserNative, as such use would prohibit the use of a
// native dialog.
//
// No operations that change the dialog work while the dialog is visible. Set
// all the properties that are required before showing the dialog.
//
//
// Win32 details
//
// On windows the IFileDialog implementation (added in Windows Vista) is used.
// It supports many of the features that GtkFileChooser has, but there are some
// things it does not handle:
//
// * Any gtk.FileFilter added using a mimetype
//
// If any of these features are used the regular GtkFileChooserDialog will be
// used in place of the native one.
//
//
// Portal details
//
// When the org.freedesktop.portal.FileChooser portal is available on the
// session bus, it is used to bring up an out-of-process file chooser. Depending
// on the kind of session the application is running in, this may or may not be
// a GTK file chooser.
//
// macOS details
//
// On macOS the NSSavePanel and NSOpenPanel classes are used to provide native
// file chooser dialogs. Some features provided by GtkFileChooser are not
// supported:
//
// * Shortcut folders.
type FileChooserNative struct {
	_ [0]func() // equal guard
	NativeDialog

	*externglib.Object
	FileChooser
}

var (
	_ NativeDialogger     = (*FileChooserNative)(nil)
	_ externglib.Objector = (*FileChooserNative)(nil)
)

func classInitFileChooserNativer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFileChooserNative(obj *externglib.Object) *FileChooserNative {
	return &FileChooserNative{
		NativeDialog: NativeDialog{
			Object: obj,
		},
		Object: obj,
		FileChooser: FileChooser{
			Object: obj,
		},
	}
}

func marshalFileChooserNative(p uintptr) (interface{}, error) {
	return wrapFileChooserNative(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFileChooserNative creates a new GtkFileChooserNative.
//
// The function takes the following parameters:
//
//    - title (optional): title of the native, or NULL.
//    - parent (optional): transient parent of the native, or NULL.
//    - action: open or save mode for the dialog.
//    - acceptLabel (optional): text to go in the accept button, or NULL for the
//      default.
//    - cancelLabel (optional): text to go in the cancel button, or NULL for the
//      default.
//
// The function returns the following values:
//
//    - fileChooserNative: new GtkFileChooserNative.
//
func NewFileChooserNative(title string, parent *Window, action FileChooserAction, acceptLabel, cancelLabel string) *FileChooserNative {
	var _arg1 *C.char                 // out
	var _arg2 *C.GtkWindow            // out
	var _arg3 C.GtkFileChooserAction  // out
	var _arg4 *C.char                 // out
	var _arg5 *C.char                 // out
	var _cret *C.GtkFileChooserNative // in

	if title != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	if parent != nil {
		_arg2 = (*C.GtkWindow)(unsafe.Pointer(externglib.InternObject(parent).Native()))
	}
	_arg3 = C.GtkFileChooserAction(action)
	if acceptLabel != "" {
		_arg4 = (*C.char)(unsafe.Pointer(C.CString(acceptLabel)))
		defer C.free(unsafe.Pointer(_arg4))
	}
	if cancelLabel != "" {
		_arg5 = (*C.char)(unsafe.Pointer(C.CString(cancelLabel)))
		defer C.free(unsafe.Pointer(_arg5))
	}

	_cret = C.gtk_file_chooser_native_new(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(title)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(action)
	runtime.KeepAlive(acceptLabel)
	runtime.KeepAlive(cancelLabel)

	var _fileChooserNative *FileChooserNative // out

	_fileChooserNative = wrapFileChooserNative(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _fileChooserNative
}

// AcceptLabel retrieves the custom label text for the accept button.
//
// The function returns the following values:
//
//    - utf8 (optional): custom label, or NULL for the default. This string is
//      owned by GTK and should not be modified or freed.
//
func (self *FileChooserNative) AcceptLabel() string {
	var _arg0 *C.GtkFileChooserNative // out
	var _cret *C.char                 // in

	_arg0 = (*C.GtkFileChooserNative)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_file_chooser_native_get_accept_label(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// CancelLabel retrieves the custom label text for the cancel button.
//
// The function returns the following values:
//
//    - utf8 (optional): custom label, or NULL for the default. This string is
//      owned by GTK and should not be modified or freed.
//
func (self *FileChooserNative) CancelLabel() string {
	var _arg0 *C.GtkFileChooserNative // out
	var _cret *C.char                 // in

	_arg0 = (*C.GtkFileChooserNative)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_file_chooser_native_get_cancel_label(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SetAcceptLabel sets the custom label text for the accept button.
//
// If characters in label are preceded by an underscore, they are underlined. If
// you need a literal underscore character in a label, use “__” (two
// underscores). The first underlined character represents a keyboard
// accelerator called a mnemonic.
//
// Pressing Alt and that key should activate the button.
//
// The function takes the following parameters:
//
//    - acceptLabel (optional): custom label or NULL for the default.
//
func (self *FileChooserNative) SetAcceptLabel(acceptLabel string) {
	var _arg0 *C.GtkFileChooserNative // out
	var _arg1 *C.char                 // out

	_arg0 = (*C.GtkFileChooserNative)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if acceptLabel != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(acceptLabel)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_file_chooser_native_set_accept_label(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(acceptLabel)
}

// SetCancelLabel sets the custom label text for the cancel button.
//
// If characters in label are preceded by an underscore, they are underlined. If
// you need a literal underscore character in a label, use “__” (two
// underscores). The first underlined character represents a keyboard
// accelerator called a mnemonic.
//
// Pressing Alt and that key should activate the button.
//
// The function takes the following parameters:
//
//    - cancelLabel (optional): custom label or NULL for the default.
//
func (self *FileChooserNative) SetCancelLabel(cancelLabel string) {
	var _arg0 *C.GtkFileChooserNative // out
	var _arg1 *C.char                 // out

	_arg0 = (*C.GtkFileChooserNative)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if cancelLabel != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(cancelLabel)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_file_chooser_native_set_cancel_label(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(cancelLabel)
}
