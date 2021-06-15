// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_page_setup_unix_dialog_get_type()), F: marshalPageSetupUnixDialog},
	})
}

// PageSetupUnixDialog: `GtkPageSetupUnixDialog` implements a page setup dialog
// for platforms which don’t provide a native page setup dialog, like Unix.
//
// !An example GtkPageSetupUnixDialog (pagesetupdialog.png)
//
// It can be used very much like any other GTK dialog, at the cost of the
// portability offered by the high-level printing API in
// [class@Gtk.PrintOperation].
type PageSetupUnixDialog interface {
	Dialog
	Accessible
	Buildable
	ConstraintTarget
	Native
	Root
	ShortcutManager

	// PageSetup gets the currently selected page setup from the dialog.
	PageSetup() PageSetup
	// PrintSettings gets the current print settings from the dialog.
	PrintSettings() PrintSettings
	// SetPageSetup sets the `GtkPageSetup` from which the page setup dialog
	// takes its values.
	SetPageSetup(pageSetup PageSetup)
	// SetPrintSettings sets the `GtkPrintSettings` from which the page setup
	// dialog takes its values.
	SetPrintSettings(printSettings PrintSettings)
}

// pageSetupUnixDialog implements the PageSetupUnixDialog class.
type pageSetupUnixDialog struct {
	Dialog
	Accessible
	Buildable
	ConstraintTarget
	Native
	Root
	ShortcutManager
}

var _ PageSetupUnixDialog = (*pageSetupUnixDialog)(nil)

// WrapPageSetupUnixDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapPageSetupUnixDialog(obj *externglib.Object) PageSetupUnixDialog {
	return pageSetupUnixDialog{
		Dialog:           WrapDialog(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Native:           WrapNative(obj),
		Root:             WrapRoot(obj),
		ShortcutManager:  WrapShortcutManager(obj),
	}
}

func marshalPageSetupUnixDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPageSetupUnixDialog(obj), nil
}

// NewPageSetupUnixDialog constructs a class PageSetupUnixDialog.
func NewPageSetupUnixDialog(title string, parent Window) PageSetupUnixDialog {
	var _arg1 *C.char                  // out
	var _arg2 *C.GtkWindow             // out
	var _cret C.GtkPageSetupUnixDialog // in

	_arg1 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	_cret = C.gtk_page_setup_unix_dialog_new(_arg1, _arg2)

	var _pageSetupUnixDialog PageSetupUnixDialog // out

	_pageSetupUnixDialog = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(PageSetupUnixDialog)

	return _pageSetupUnixDialog
}

// PageSetup gets the currently selected page setup from the dialog.
func (d pageSetupUnixDialog) PageSetup() PageSetup {
	var _arg0 *C.GtkPageSetupUnixDialog // out
	var _cret *C.GtkPageSetup           // in

	_arg0 = (*C.GtkPageSetupUnixDialog)(unsafe.Pointer(d.Native()))

	_cret = C.gtk_page_setup_unix_dialog_get_page_setup(_arg0)

	var _pageSetup PageSetup // out

	_pageSetup = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(PageSetup)

	return _pageSetup
}

// PrintSettings gets the current print settings from the dialog.
func (d pageSetupUnixDialog) PrintSettings() PrintSettings {
	var _arg0 *C.GtkPageSetupUnixDialog // out
	var _cret *C.GtkPrintSettings       // in

	_arg0 = (*C.GtkPageSetupUnixDialog)(unsafe.Pointer(d.Native()))

	_cret = C.gtk_page_setup_unix_dialog_get_print_settings(_arg0)

	var _printSettings PrintSettings // out

	_printSettings = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(PrintSettings)

	return _printSettings
}

// SetPageSetup sets the `GtkPageSetup` from which the page setup dialog
// takes its values.
func (d pageSetupUnixDialog) SetPageSetup(pageSetup PageSetup) {
	var _arg0 *C.GtkPageSetupUnixDialog // out
	var _arg1 *C.GtkPageSetup           // out

	_arg0 = (*C.GtkPageSetupUnixDialog)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkPageSetup)(unsafe.Pointer(pageSetup.Native()))

	C.gtk_page_setup_unix_dialog_set_page_setup(_arg0, _arg1)
}

// SetPrintSettings sets the `GtkPrintSettings` from which the page setup
// dialog takes its values.
func (d pageSetupUnixDialog) SetPrintSettings(printSettings PrintSettings) {
	var _arg0 *C.GtkPageSetupUnixDialog // out
	var _arg1 *C.GtkPrintSettings       // out

	_arg0 = (*C.GtkPageSetupUnixDialog)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkPrintSettings)(unsafe.Pointer(printSettings.Native()))

	C.gtk_page_setup_unix_dialog_set_print_settings(_arg0, _arg1)
}
