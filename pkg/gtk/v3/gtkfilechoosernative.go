// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkfilechoosernative.go.
var GTypeFileChooserNative = coreglib.Type(C.gtk_file_chooser_native_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeFileChooserNative, F: marshalFileChooserNative},
	})
}

// FileChooserNativeOverrider contains methods that are overridable.
type FileChooserNativeOverrider interface {
}

// FileChooserNative is an abstraction of a dialog box suitable for use with
// “File/Open” or “File/Save as” commands. By default, this just uses a
// FileChooserDialog to implement the actual dialog. However, on certain
// platforms, such as Windows and macOS, the native platform file chooser is
// used instead. When the application is running in a sandboxed environment
// without direct filesystem access (such as Flatpak), FileChooserNative may
// call the proper APIs (portals) to let the user choose a file and make it
// available to the application.
//
// While the API of FileChooserNative closely mirrors FileChooserDialog, the
// main difference is that there is no access to any Window or Widget for the
// dialog. This is required, as there may not be one in the case of a platform
// native dialog. Showing, hiding and running the dialog is handled by the
// NativeDialog functions.
//
//
// Typical usage
//
// In the simplest of cases, you can the following code to use FileChooserDialog
// to select a file for opening:
//
//    GtkFileChooserNative *native;
//    GtkFileChooserAction action = GTK_FILE_CHOOSER_ACTION_OPEN;
//    gint res;
//
//    native = gtk_file_chooser_native_new ("Open File",
//                                          parent_window,
//                                          action,
//                                          "_Open",
//                                          "_Cancel");
//
//    res = gtk_native_dialog_run (GTK_NATIVE_DIALOG (native));
//    if (res == GTK_RESPONSE_ACCEPT)
//      {
//        char *filename;
//        GtkFileChooser *chooser = GTK_FILE_CHOOSER (native);
//        filename = gtk_file_chooser_get_filename (chooser);
//        open_file (filename);
//        g_free (filename);
//      }
//
//    g_object_unref (native);
//
// To use a dialog for saving, you can use this:
//
//    GtkFileChooserNative *native;
//    GtkFileChooser *chooser;
//    GtkFileChooserAction action = GTK_FILE_CHOOSER_ACTION_SAVE;
//    gint res;
//
//    native = gtk_file_chooser_native_new ("Save File",
//                                          parent_window,
//                                          action,
//                                          "_Save",
//                                          "_Cancel");
//    chooser = GTK_FILE_CHOOSER (native);
//
//    gtk_file_chooser_set_do_overwrite_confirmation (chooser, TRUE);
//
//    if (user_edited_a_new_document)
//      gtk_file_chooser_set_current_name (chooser,
//                                         _("Untitled document"));
//    else
//      gtk_file_chooser_set_filename (chooser,
//                                     existing_filename);
//
//    res = gtk_native_dialog_run (GTK_NATIVE_DIALOG (native));
//    if (res == GTK_RESPONSE_ACCEPT)
//      {
//        char *filename;
//
//        filename = gtk_file_chooser_get_filename (chooser);
//        save_to_file (filename);
//        g_free (filename);
//      }
//
//    g_object_unref (native);
//
// For more information on how to best set up a file dialog, see
// FileChooserDialog.
//
//
// Response Codes
//
// FileChooserNative inherits from NativeDialog, which means it will return
// K_RESPONSE_ACCEPT if the user accepted, and K_RESPONSE_CANCEL if he pressed
// cancel. It can also return K_RESPONSE_DELETE_EVENT if the window was
// unexpectedly closed.
//
// Differences from FileChooserDialog ##
// {#gtkfilechooserdialognative-differences}
//
// There are a few things in the GtkFileChooser API that are not possible to use
// with FileChooserNative, as such use would prohibit the use of a native
// dialog.
//
// There is no support for the signals that are emitted when the user navigates
// in the dialog, including: * FileChooser::current-folder-changed *
// FileChooser::selection-changed * FileChooser::file-activated *
// FileChooser::confirm-overwrite
//
// You can also not use the methods that directly control user navigation: *
// gtk_file_chooser_unselect_filename() * gtk_file_chooser_select_all() *
// gtk_file_chooser_unselect_all()
//
// If you need any of the above you will have to use FileChooserDialog directly.
//
// No operations that change the the dialog work while the dialog is visible.
// Set all the properties that are required before showing the dialog.
//
//
// Win32 details
//
// On windows the IFileDialog implementation (added in Windows Vista) is used.
// It supports many of the features that FileChooserDialog does, but there are
// some things it does not handle:
//
// * Extra widgets added with gtk_file_chooser_set_extra_widget().
//
// * Use of custom previews by connecting to FileChooser::update-preview.
//
// * Any FileFilter added using a mimetype or custom filter.
//
// If any of these features are used the regular FileChooserDialog will be used
// in place of the native one.
//
//
// Portal details
//
// When the org.freedesktop.portal.FileChooser portal is available on the
// session bus, it is used to bring up an out-of-process file chooser. Depending
// on the kind of session the application is running in, this may or may not be
// a GTK+ file chooser. In this situation, the following things are not
// supported and will be silently ignored:
//
// * Extra widgets added with gtk_file_chooser_set_extra_widget().
//
// * Use of custom previews by connecting to FileChooser::update-preview.
//
// * Any FileFilter added with a custom filter.
//
// macOS details
//
// On macOS the NSSavePanel and NSOpenPanel classes are used to provide native
// file chooser dialogs. Some features provided by FileChooserDialog are not
// supported:
//
// * Extra widgets added with gtk_file_chooser_set_extra_widget(), unless the
// widget is an instance of GtkLabel, in which case the label text will be used
// to set the NSSavePanel message instance property.
//
// * Use of custom previews by connecting to FileChooser::update-preview.
//
// * Any FileFilter added with a custom filter.
//
// * Shortcut folders.
type FileChooserNative struct {
	_ [0]func() // equal guard
	NativeDialog

	*coreglib.Object
	FileChooser
}

var (
	_ NativeDialogger   = (*FileChooserNative)(nil)
	_ coreglib.Objector = (*FileChooserNative)(nil)
)

func classInitFileChooserNativer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFileChooserNative(obj *coreglib.Object) *FileChooserNative {
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
	return wrapFileChooserNative(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AcceptLabel retrieves the custom label text for the accept button.
//
// The function returns the following values:
//
//    - utf8 (optional): custom label, or NULL for the default. This string is
//      owned by GTK+ and should not be modified or freed.
//
func (self *FileChooserNative) AcceptLabel() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**FileChooserNative)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "FileChooserNative").InvokeMethod("get_accept_label", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
//      owned by GTK+ and should not be modified or freed.
//
func (self *FileChooserNative) CancelLabel() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**FileChooserNative)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "FileChooserNative").InvokeMethod("get_cancel_label", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
// accelerator called a mnemonic. Pressing Alt and that key activates the
// button.
//
// The function takes the following parameters:
//
//    - acceptLabel (optional): custom label or NULL for the default.
//
func (self *FileChooserNative) SetAcceptLabel(acceptLabel string) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if acceptLabel != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(acceptLabel)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	*(**FileChooserNative)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "FileChooserNative").InvokeMethod("set_accept_label", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(acceptLabel)
}

// SetCancelLabel sets the custom label text for the cancel button.
//
// If characters in label are preceded by an underscore, they are underlined. If
// you need a literal underscore character in a label, use “__” (two
// underscores). The first underlined character represents a keyboard
// accelerator called a mnemonic. Pressing Alt and that key activates the
// button.
//
// The function takes the following parameters:
//
//    - cancelLabel (optional): custom label or NULL for the default.
//
func (self *FileChooserNative) SetCancelLabel(cancelLabel string) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if cancelLabel != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(cancelLabel)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	*(**FileChooserNative)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "FileChooserNative").InvokeMethod("set_cancel_label", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(cancelLabel)
}
