// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gsk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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

	PageSetup() PageSetup

	PrintSettings() PrintSettings

	SetPageSetupPageSetupUnixDialog(pageSetup PageSetup)

	SetPrintSettingsPageSetupUnixDialog(printSettings PrintSettings)
}

// pageSetupUnixDialog implements the PageSetupUnixDialog class.
type pageSetupUnixDialog struct {
	Dialog
}

// WrapPageSetupUnixDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapPageSetupUnixDialog(obj *externglib.Object) PageSetupUnixDialog {
	return pageSetupUnixDialog{
		Dialog: WrapDialog(obj),
	}
}

func marshalPageSetupUnixDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPageSetupUnixDialog(obj), nil
}

func NewPageSetupUnixDialog(title string, parent Window) PageSetupUnixDialog {
	var _arg1 *C.char      // out
	var _arg2 *C.GtkWindow // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	_cret = C.gtk_page_setup_unix_dialog_new(_arg1, _arg2)

	var _pageSetupUnixDialog PageSetupUnixDialog // out

	_pageSetupUnixDialog = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(PageSetupUnixDialog)

	return _pageSetupUnixDialog
}

func (d pageSetupUnixDialog) PageSetup() PageSetup {
	var _arg0 *C.GtkPageSetupUnixDialog // out
	var _cret *C.GtkPageSetup           // in

	_arg0 = (*C.GtkPageSetupUnixDialog)(unsafe.Pointer(d.Native()))

	_cret = C.gtk_page_setup_unix_dialog_get_page_setup(_arg0)

	var _pageSetup PageSetup // out

	_pageSetup = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(PageSetup)

	return _pageSetup
}

func (d pageSetupUnixDialog) PrintSettings() PrintSettings {
	var _arg0 *C.GtkPageSetupUnixDialog // out
	var _cret *C.GtkPrintSettings       // in

	_arg0 = (*C.GtkPageSetupUnixDialog)(unsafe.Pointer(d.Native()))

	_cret = C.gtk_page_setup_unix_dialog_get_print_settings(_arg0)

	var _printSettings PrintSettings // out

	_printSettings = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(PrintSettings)

	return _printSettings
}

func (d pageSetupUnixDialog) SetPageSetupPageSetupUnixDialog(pageSetup PageSetup) {
	var _arg0 *C.GtkPageSetupUnixDialog // out
	var _arg1 *C.GtkPageSetup           // out

	_arg0 = (*C.GtkPageSetupUnixDialog)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkPageSetup)(unsafe.Pointer(pageSetup.Native()))

	C.gtk_page_setup_unix_dialog_set_page_setup(_arg0, _arg1)
}

func (d pageSetupUnixDialog) SetPrintSettingsPageSetupUnixDialog(printSettings PrintSettings) {
	var _arg0 *C.GtkPageSetupUnixDialog // out
	var _arg1 *C.GtkPrintSettings       // out

	_arg0 = (*C.GtkPageSetupUnixDialog)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkPrintSettings)(unsafe.Pointer(printSettings.Native()))

	C.gtk_page_setup_unix_dialog_set_print_settings(_arg0, _arg1)
}

func (s pageSetupUnixDialog) Display() gdk.Display {
	return WrapRoot(gextras.InternObject(s)).Display()
}

func (s pageSetupUnixDialog) Focus() Widget {
	return WrapRoot(gextras.InternObject(s)).Focus()
}

func (s pageSetupUnixDialog) SetFocus(focus Widget) {
	WrapRoot(gextras.InternObject(s)).SetFocus(focus)
}

func (s pageSetupUnixDialog) Renderer() gsk.Renderer {
	return WrapNative(gextras.InternObject(s)).Renderer()
}

func (s pageSetupUnixDialog) Surface() gdk.Surface {
	return WrapNative(gextras.InternObject(s)).Surface()
}

func (s pageSetupUnixDialog) SurfaceTransform() (x float64, y float64) {
	return WrapNative(gextras.InternObject(s)).SurfaceTransform()
}

func (s pageSetupUnixDialog) Realize() {
	WrapNative(gextras.InternObject(s)).Realize()
}

func (s pageSetupUnixDialog) Unrealize() {
	WrapNative(gextras.InternObject(s)).Unrealize()
}

func (s pageSetupUnixDialog) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s pageSetupUnixDialog) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s pageSetupUnixDialog) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s pageSetupUnixDialog) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s pageSetupUnixDialog) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s pageSetupUnixDialog) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s pageSetupUnixDialog) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b pageSetupUnixDialog) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
