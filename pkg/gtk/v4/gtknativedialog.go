// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_native_dialog_get_type()), F: marshalNativeDialogger},
	})
}

// NativeDialogOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type NativeDialogOverrider interface {
	// Hide hides the dialog if it is visible, aborting any interaction.
	//
	// Once this is called the gtk.NativeDialog::response signal will *not* be
	// emitted until after the next call to gtk.NativeDialog.Show().
	//
	// If the dialog is not visible this does nothing.
	Hide()
	Response(responseId int)
	// Show shows the dialog on the display.
	//
	// When the user accepts the state of the dialog the dialog will be
	// automatically hidden and the gtk.NativeDialog::response signal will be
	// emitted.
	//
	// Multiple calls while the dialog is visible will be ignored.
	Show()
}

// NativeDialog: native dialogs are platform dialogs that don't use GtkDialog.
//
// They are used in order to integrate better with a platform, by looking the
// same as other native applications and supporting platform specific features.
//
// The gtk.Dialog functions cannot be used on such objects, but we need a
// similar API in order to drive them. The GtkNativeDialog object is an API that
// allows you to do this. It allows you to set various common properties on the
// dialog, as well as show and hide it and get a gtk.NativeDialog::response
// signal when the user finished with the dialog.
//
// Note that unlike GtkDialog, GtkNativeDialog objects are not toplevel widgets,
// and GTK does not keep them alive. It is your responsibility to keep a
// reference until you are done with the object.
type NativeDialog struct {
	*externglib.Object
}

// NativeDialogger describes NativeDialog's abstract methods.
type NativeDialogger interface {
	externglib.Objector

	// Destroy destroys a dialog.
	Destroy()
	// Modal returns whether the dialog is modal.
	Modal() bool
	// Title gets the title of the GtkNativeDialog.
	Title() string
	// TransientFor fetches the transient parent for this window.
	TransientFor() *Window
	// Visible determines whether the dialog is visible.
	Visible() bool
	// Hide hides the dialog if it is visible, aborting any interaction.
	Hide()
	// SetModal sets a dialog modal or non-modal.
	SetModal(modal bool)
	// SetTitle sets the title of the GtkNativeDialog.
	SetTitle(title string)
	// SetTransientFor: dialog windows should be set transient for the main
	// application window they were spawned from.
	SetTransientFor(parent *Window)
	// Show shows the dialog on the display.
	Show()
}

var _ NativeDialogger = (*NativeDialog)(nil)

func wrapNativeDialog(obj *externglib.Object) *NativeDialog {
	return &NativeDialog{
		Object: obj,
	}
}

func marshalNativeDialogger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNativeDialog(obj), nil
}

// Destroy destroys a dialog.
//
// When a dialog is destroyed, it will break any references it holds to other
// objects.
//
// If it is visible it will be hidden and any underlying window system resources
// will be destroyed.
//
// Note that this does not release any reference to the object (as opposed to
// destroying a GtkWindow) because there is no reference from the windowing
// system to the GtkNativeDialog.
func (self *NativeDialog) Destroy() {
	var _arg0 *C.GtkNativeDialog // out

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(self.Native()))

	C.gtk_native_dialog_destroy(_arg0)
}

// Modal returns whether the dialog is modal.
func (self *NativeDialog) Modal() bool {
	var _arg0 *C.GtkNativeDialog // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_native_dialog_get_modal(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Title gets the title of the GtkNativeDialog.
func (self *NativeDialog) Title() string {
	var _arg0 *C.GtkNativeDialog // out
	var _cret *C.char            // in

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_native_dialog_get_title(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// TransientFor fetches the transient parent for this window.
func (self *NativeDialog) TransientFor() *Window {
	var _arg0 *C.GtkNativeDialog // out
	var _cret *C.GtkWindow       // in

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_native_dialog_get_transient_for(_arg0)

	var _window *Window // out

	if _cret != nil {
		_window = wrapWindow(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _window
}

// Visible determines whether the dialog is visible.
func (self *NativeDialog) Visible() bool {
	var _arg0 *C.GtkNativeDialog // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_native_dialog_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Hide hides the dialog if it is visible, aborting any interaction.
//
// Once this is called the gtk.NativeDialog::response signal will *not* be
// emitted until after the next call to gtk.NativeDialog.Show().
//
// If the dialog is not visible this does nothing.
func (self *NativeDialog) Hide() {
	var _arg0 *C.GtkNativeDialog // out

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(self.Native()))

	C.gtk_native_dialog_hide(_arg0)
}

// SetModal sets a dialog modal or non-modal.
//
// Modal dialogs prevent interaction with other windows in the same application.
// To keep modal dialogs on top of main application windows, use
// gtk.NativeDialog.SetTransientFor() to make the dialog transient for the
// parent; most window managers will then disallow lowering the dialog below the
// parent.
func (self *NativeDialog) SetModal(modal bool) {
	var _arg0 *C.GtkNativeDialog // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(self.Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.gtk_native_dialog_set_modal(_arg0, _arg1)
}

// SetTitle sets the title of the GtkNativeDialog.
func (self *NativeDialog) SetTitle(title string) {
	var _arg0 *C.GtkNativeDialog // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_native_dialog_set_title(_arg0, _arg1)
}

// SetTransientFor: dialog windows should be set transient for the main
// application window they were spawned from.
//
// This allows window managers to e.g. keep the dialog on top of the main
// window, or center the dialog over the main window.
//
// Passing NULL for parent unsets the current transient window.
func (self *NativeDialog) SetTransientFor(parent *Window) {
	var _arg0 *C.GtkNativeDialog // out
	var _arg1 *C.GtkWindow       // out

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(self.Native()))
	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	}

	C.gtk_native_dialog_set_transient_for(_arg0, _arg1)
}

// Show shows the dialog on the display.
//
// When the user accepts the state of the dialog the dialog will be
// automatically hidden and the gtk.NativeDialog::response signal will be
// emitted.
//
// Multiple calls while the dialog is visible will be ignored.
func (self *NativeDialog) Show() {
	var _arg0 *C.GtkNativeDialog // out

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(self.Native()))

	C.gtk_native_dialog_show(_arg0)
}
