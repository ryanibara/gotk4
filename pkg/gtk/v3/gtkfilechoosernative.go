// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_file_chooser_native_get_type()), F: marshalFileChooserNative},
	})
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
type FileChooserNative interface {
	NativeDialog
	FileChooser

	AcceptLabel() string

	CancelLabel() string

	SetAcceptLabelFileChooserNative(acceptLabel string)

	SetCancelLabelFileChooserNative(cancelLabel string)
}

// fileChooserNative implements the FileChooserNative class.
type fileChooserNative struct {
	NativeDialog
}

// WrapFileChooserNative wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileChooserNative(obj *externglib.Object) FileChooserNative {
	return fileChooserNative{
		NativeDialog: WrapNativeDialog(obj),
	}
}

func marshalFileChooserNative(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileChooserNative(obj), nil
}

func NewFileChooserNative(title string, parent Window, action FileChooserAction, acceptLabel string, cancelLabel string) FileChooserNative {
	var _arg1 *C.gchar                // out
	var _arg2 *C.GtkWindow            // out
	var _arg3 C.GtkFileChooserAction  // out
	var _arg4 *C.gchar                // out
	var _arg5 *C.gchar                // out
	var _cret *C.GtkFileChooserNative // in

	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	_arg3 = C.GtkFileChooserAction(action)
	_arg4 = (*C.gchar)(C.CString(acceptLabel))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = (*C.gchar)(C.CString(cancelLabel))
	defer C.free(unsafe.Pointer(_arg5))

	_cret = C.gtk_file_chooser_native_new(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _fileChooserNative FileChooserNative // out

	_fileChooserNative = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(FileChooserNative)

	return _fileChooserNative
}

func (s fileChooserNative) AcceptLabel() string {
	var _arg0 *C.GtkFileChooserNative // out
	var _cret *C.char                 // in

	_arg0 = (*C.GtkFileChooserNative)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_file_chooser_native_get_accept_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s fileChooserNative) CancelLabel() string {
	var _arg0 *C.GtkFileChooserNative // out
	var _cret *C.char                 // in

	_arg0 = (*C.GtkFileChooserNative)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_file_chooser_native_get_cancel_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s fileChooserNative) SetAcceptLabelFileChooserNative(acceptLabel string) {
	var _arg0 *C.GtkFileChooserNative // out
	var _arg1 *C.char                 // out

	_arg0 = (*C.GtkFileChooserNative)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(acceptLabel))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_chooser_native_set_accept_label(_arg0, _arg1)
}

func (s fileChooserNative) SetCancelLabelFileChooserNative(cancelLabel string) {
	var _arg0 *C.GtkFileChooserNative // out
	var _arg1 *C.char                 // out

	_arg0 = (*C.GtkFileChooserNative)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(cancelLabel))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_chooser_native_set_cancel_label(_arg0, _arg1)
}

func (c fileChooserNative) AddChoice(id string, label string, options []string, optionLabels []string) {
	WrapFileChooser(gextras.InternObject(c)).AddChoice(id, label, options, optionLabels)
}

func (c fileChooserNative) AddFilter(filter FileFilter) {
	WrapFileChooser(gextras.InternObject(c)).AddFilter(filter)
}

func (c fileChooserNative) AddShortcutFolder(folder string) error {
	return WrapFileChooser(gextras.InternObject(c)).AddShortcutFolder(folder)
}

func (c fileChooserNative) AddShortcutFolderURI(uri string) error {
	return WrapFileChooser(gextras.InternObject(c)).AddShortcutFolderURI(uri)
}

func (c fileChooserNative) Action() FileChooserAction {
	return WrapFileChooser(gextras.InternObject(c)).Action()
}

func (c fileChooserNative) Choice(id string) string {
	return WrapFileChooser(gextras.InternObject(c)).Choice(id)
}

func (c fileChooserNative) CreateFolders() bool {
	return WrapFileChooser(gextras.InternObject(c)).CreateFolders()
}

func (c fileChooserNative) CurrentFolder() string {
	return WrapFileChooser(gextras.InternObject(c)).CurrentFolder()
}

func (c fileChooserNative) CurrentFolderURI() string {
	return WrapFileChooser(gextras.InternObject(c)).CurrentFolderURI()
}

func (c fileChooserNative) CurrentName() string {
	return WrapFileChooser(gextras.InternObject(c)).CurrentName()
}

func (c fileChooserNative) DoOverwriteConfirmation() bool {
	return WrapFileChooser(gextras.InternObject(c)).DoOverwriteConfirmation()
}

func (c fileChooserNative) ExtraWidget() Widget {
	return WrapFileChooser(gextras.InternObject(c)).ExtraWidget()
}

func (c fileChooserNative) Filename() string {
	return WrapFileChooser(gextras.InternObject(c)).Filename()
}

func (c fileChooserNative) Filter() FileFilter {
	return WrapFileChooser(gextras.InternObject(c)).Filter()
}

func (c fileChooserNative) LocalOnly() bool {
	return WrapFileChooser(gextras.InternObject(c)).LocalOnly()
}

func (c fileChooserNative) PreviewFilename() string {
	return WrapFileChooser(gextras.InternObject(c)).PreviewFilename()
}

func (c fileChooserNative) PreviewURI() string {
	return WrapFileChooser(gextras.InternObject(c)).PreviewURI()
}

func (c fileChooserNative) PreviewWidget() Widget {
	return WrapFileChooser(gextras.InternObject(c)).PreviewWidget()
}

func (c fileChooserNative) PreviewWidgetActive() bool {
	return WrapFileChooser(gextras.InternObject(c)).PreviewWidgetActive()
}

func (c fileChooserNative) SelectMultiple() bool {
	return WrapFileChooser(gextras.InternObject(c)).SelectMultiple()
}

func (c fileChooserNative) ShowHidden() bool {
	return WrapFileChooser(gextras.InternObject(c)).ShowHidden()
}

func (c fileChooserNative) URI() string {
	return WrapFileChooser(gextras.InternObject(c)).URI()
}

func (c fileChooserNative) UsePreviewLabel() bool {
	return WrapFileChooser(gextras.InternObject(c)).UsePreviewLabel()
}

func (c fileChooserNative) RemoveChoice(id string) {
	WrapFileChooser(gextras.InternObject(c)).RemoveChoice(id)
}

func (c fileChooserNative) RemoveFilter(filter FileFilter) {
	WrapFileChooser(gextras.InternObject(c)).RemoveFilter(filter)
}

func (c fileChooserNative) RemoveShortcutFolder(folder string) error {
	return WrapFileChooser(gextras.InternObject(c)).RemoveShortcutFolder(folder)
}

func (c fileChooserNative) RemoveShortcutFolderURI(uri string) error {
	return WrapFileChooser(gextras.InternObject(c)).RemoveShortcutFolderURI(uri)
}

func (c fileChooserNative) SelectAll() {
	WrapFileChooser(gextras.InternObject(c)).SelectAll()
}

func (c fileChooserNative) SelectFilename(filename string) bool {
	return WrapFileChooser(gextras.InternObject(c)).SelectFilename(filename)
}

func (c fileChooserNative) SelectURI(uri string) bool {
	return WrapFileChooser(gextras.InternObject(c)).SelectURI(uri)
}

func (c fileChooserNative) SetAction(action FileChooserAction) {
	WrapFileChooser(gextras.InternObject(c)).SetAction(action)
}

func (c fileChooserNative) SetChoice(id string, option string) {
	WrapFileChooser(gextras.InternObject(c)).SetChoice(id, option)
}

func (c fileChooserNative) SetCreateFolders(createFolders bool) {
	WrapFileChooser(gextras.InternObject(c)).SetCreateFolders(createFolders)
}

func (c fileChooserNative) SetCurrentFolder(filename string) bool {
	return WrapFileChooser(gextras.InternObject(c)).SetCurrentFolder(filename)
}

func (c fileChooserNative) SetCurrentFolderURI(uri string) bool {
	return WrapFileChooser(gextras.InternObject(c)).SetCurrentFolderURI(uri)
}

func (c fileChooserNative) SetCurrentName(name string) {
	WrapFileChooser(gextras.InternObject(c)).SetCurrentName(name)
}

func (c fileChooserNative) SetDoOverwriteConfirmation(doOverwriteConfirmation bool) {
	WrapFileChooser(gextras.InternObject(c)).SetDoOverwriteConfirmation(doOverwriteConfirmation)
}

func (c fileChooserNative) SetExtraWidget(extraWidget Widget) {
	WrapFileChooser(gextras.InternObject(c)).SetExtraWidget(extraWidget)
}

func (c fileChooserNative) SetFilename(filename string) bool {
	return WrapFileChooser(gextras.InternObject(c)).SetFilename(filename)
}

func (c fileChooserNative) SetFilter(filter FileFilter) {
	WrapFileChooser(gextras.InternObject(c)).SetFilter(filter)
}

func (c fileChooserNative) SetLocalOnly(localOnly bool) {
	WrapFileChooser(gextras.InternObject(c)).SetLocalOnly(localOnly)
}

func (c fileChooserNative) SetPreviewWidget(previewWidget Widget) {
	WrapFileChooser(gextras.InternObject(c)).SetPreviewWidget(previewWidget)
}

func (c fileChooserNative) SetPreviewWidgetActive(active bool) {
	WrapFileChooser(gextras.InternObject(c)).SetPreviewWidgetActive(active)
}

func (c fileChooserNative) SetSelectMultiple(selectMultiple bool) {
	WrapFileChooser(gextras.InternObject(c)).SetSelectMultiple(selectMultiple)
}

func (c fileChooserNative) SetShowHidden(showHidden bool) {
	WrapFileChooser(gextras.InternObject(c)).SetShowHidden(showHidden)
}

func (c fileChooserNative) SetURI(uri string) bool {
	return WrapFileChooser(gextras.InternObject(c)).SetURI(uri)
}

func (c fileChooserNative) SetUsePreviewLabel(useLabel bool) {
	WrapFileChooser(gextras.InternObject(c)).SetUsePreviewLabel(useLabel)
}

func (c fileChooserNative) UnselectAll() {
	WrapFileChooser(gextras.InternObject(c)).UnselectAll()
}

func (c fileChooserNative) UnselectFilename(filename string) {
	WrapFileChooser(gextras.InternObject(c)).UnselectFilename(filename)
}

func (c fileChooserNative) UnselectURI(uri string) {
	WrapFileChooser(gextras.InternObject(c)).UnselectURI(uri)
}
