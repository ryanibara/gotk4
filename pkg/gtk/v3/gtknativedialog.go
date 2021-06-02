// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_native_dialog_get_type()), F: marshalNativeDialog},
	})
}

// NativeDialog: native dialogs are platform dialogs that don't use Dialog or
// Window. They are used in order to integrate better with a platform, by
// looking the same as other native applications and supporting platform
// specific features.
//
// The Dialog functions cannot be used on such objects, but we need a similar
// API in order to drive them. The NativeDialog object is an API that allows you
// to do this. It allows you to set various common properties on the dialog, as
// well as show and hide it and get a NativeDialog::response signal when the
// user finished with the dialog.
//
// There is also a gtk_native_dialog_run() helper that makes it easy to run any
// native dialog in a modal way with a recursive mainloop, similar to
// gtk_dialog_run().
type NativeDialog interface {
	gextras.Objector

	// Destroy destroys a dialog.
	//
	// When a dialog is destroyed, it will break any references it holds to
	// other objects. If it is visible it will be hidden and any underlying
	// window system resources will be destroyed.
	//
	// Note that this does not release any reference to the object (as opposed
	// to destroying a GtkWindow) because there is no reference from the
	// windowing system to the NativeDialog.
	Destroy()
	// Modal returns whether the dialog is modal. See
	// gtk_native_dialog_set_modal().
	Modal() bool
	// Title gets the title of the NativeDialog.
	Title() string
	// TransientFor fetches the transient parent for this window. See
	// gtk_native_dialog_set_transient_for().
	TransientFor() Window
	// Visible determines whether the dialog is visible.
	Visible() bool
	// Hide hides the dialog if it is visilbe, aborting any interaction. Once
	// this is called the NativeDialog::response signal will not be emitted
	// until after the next call to gtk_native_dialog_show().
	//
	// If the dialog is not visible this does nothing.
	Hide()
	// Run blocks in a recursive main loop until @self emits the
	// NativeDialog::response signal. It then returns the response ID from the
	// ::response signal emission.
	//
	// Before entering the recursive main loop, gtk_native_dialog_run() calls
	// gtk_native_dialog_show() on the dialog for you.
	//
	// After gtk_native_dialog_run() returns, then dialog will be hidden.
	//
	// Typical usage of this function might be:
	//
	//    gint result = gtk_native_dialog_run (GTK_NATIVE_DIALOG (dialog));
	//    switch (result)
	//      {
	//        case GTK_RESPONSE_ACCEPT:
	//           do_application_specific_something ();
	//           break;
	//        default:
	//           do_nothing_since_dialog_was_cancelled ();
	//           break;
	//      }
	//    g_object_unref (dialog);
	//
	// Note that even though the recursive main loop gives the effect of a modal
	// dialog (it prevents the user from interacting with other windows in the
	// same window group while the dialog is run), callbacks such as timeouts,
	// IO channel watches, DND drops, etc, will be triggered during a
	// gtk_native_dialog_run() call.
	Run() int
	// SetModal sets a dialog modal or non-modal. Modal dialogs prevent
	// interaction with other windows in the same application. To keep modal
	// dialogs on top of main application windows, use
	// gtk_native_dialog_set_transient_for() to make the dialog transient for
	// the parent; most [window managers][gtk-X11-arch] will then disallow
	// lowering the dialog below the parent.
	SetModal(modal bool)
	// SetTitle sets the title of the NativeDialog.
	SetTitle(title string)
	// SetTransientFor: dialog windows should be set transient for the main
	// application window they were spawned from. This allows [window
	// managers][gtk-X11-arch] to e.g. keep the dialog on top of the main
	// window, or center the dialog over the main window.
	//
	// Passing nil for @parent unsets the current transient window.
	SetTransientFor(parent Window)
	// Show shows the dialog on the display, allowing the user to interact with
	// it. When the user accepts the state of the dialog the dialog will be
	// automatically hidden and the NativeDialog::response signal will be
	// emitted.
	//
	// Multiple calls while the dialog is visible will be ignored.
	Show()
}

// nativeDialog implements the NativeDialog interface.
type nativeDialog struct {
	gextras.Objector
}

var _ NativeDialog = (*nativeDialog)(nil)

// WrapNativeDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapNativeDialog(obj *externglib.Object) NativeDialog {
	return NativeDialog{
		Objector: obj,
	}
}

func marshalNativeDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNativeDialog(obj), nil
}

// Destroy destroys a dialog.
//
// When a dialog is destroyed, it will break any references it holds to
// other objects. If it is visible it will be hidden and any underlying
// window system resources will be destroyed.
//
// Note that this does not release any reference to the object (as opposed
// to destroying a GtkWindow) because there is no reference from the
// windowing system to the NativeDialog.
func (s nativeDialog) Destroy() {
	var arg0 *C.GtkNativeDialog

	arg0 = (*C.GtkNativeDialog)(s.Native())

	C.gtk_native_dialog_destroy(arg0)
}

// Modal returns whether the dialog is modal. See
// gtk_native_dialog_set_modal().
func (s nativeDialog) Modal() bool {
	var arg0 *C.GtkNativeDialog

	arg0 = (*C.GtkNativeDialog)(s.Native())

	ret := C.gtk_native_dialog_get_modal(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Title gets the title of the NativeDialog.
func (s nativeDialog) Title() string {
	var arg0 *C.GtkNativeDialog

	arg0 = (*C.GtkNativeDialog)(s.Native())

	ret := C.gtk_native_dialog_get_title(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// TransientFor fetches the transient parent for this window. See
// gtk_native_dialog_set_transient_for().
func (s nativeDialog) TransientFor() Window {
	var arg0 *C.GtkNativeDialog

	arg0 = (*C.GtkNativeDialog)(s.Native())

	ret := C.gtk_native_dialog_get_transient_for(arg0)

	var ret0 Window

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Window)

	return ret0
}

// Visible determines whether the dialog is visible.
func (s nativeDialog) Visible() bool {
	var arg0 *C.GtkNativeDialog

	arg0 = (*C.GtkNativeDialog)(s.Native())

	ret := C.gtk_native_dialog_get_visible(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Hide hides the dialog if it is visilbe, aborting any interaction. Once
// this is called the NativeDialog::response signal will not be emitted
// until after the next call to gtk_native_dialog_show().
//
// If the dialog is not visible this does nothing.
func (s nativeDialog) Hide() {
	var arg0 *C.GtkNativeDialog

	arg0 = (*C.GtkNativeDialog)(s.Native())

	C.gtk_native_dialog_hide(arg0)
}

// Run blocks in a recursive main loop until @self emits the
// NativeDialog::response signal. It then returns the response ID from the
// ::response signal emission.
//
// Before entering the recursive main loop, gtk_native_dialog_run() calls
// gtk_native_dialog_show() on the dialog for you.
//
// After gtk_native_dialog_run() returns, then dialog will be hidden.
//
// Typical usage of this function might be:
//
//    gint result = gtk_native_dialog_run (GTK_NATIVE_DIALOG (dialog));
//    switch (result)
//      {
//        case GTK_RESPONSE_ACCEPT:
//           do_application_specific_something ();
//           break;
//        default:
//           do_nothing_since_dialog_was_cancelled ();
//           break;
//      }
//    g_object_unref (dialog);
//
// Note that even though the recursive main loop gives the effect of a modal
// dialog (it prevents the user from interacting with other windows in the
// same window group while the dialog is run), callbacks such as timeouts,
// IO channel watches, DND drops, etc, will be triggered during a
// gtk_native_dialog_run() call.
func (s nativeDialog) Run() int {
	var arg0 *C.GtkNativeDialog

	arg0 = (*C.GtkNativeDialog)(s.Native())

	ret := C.gtk_native_dialog_run(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// SetModal sets a dialog modal or non-modal. Modal dialogs prevent
// interaction with other windows in the same application. To keep modal
// dialogs on top of main application windows, use
// gtk_native_dialog_set_transient_for() to make the dialog transient for
// the parent; most [window managers][gtk-X11-arch] will then disallow
// lowering the dialog below the parent.
func (s nativeDialog) SetModal(modal bool) {
	var arg0 *C.GtkNativeDialog
	var arg1 C.gboolean

	arg0 = (*C.GtkNativeDialog)(s.Native())
	if modal {
		arg1 = C.TRUE
	}

	C.gtk_native_dialog_set_modal(arg0, arg1)
}

// SetTitle sets the title of the NativeDialog.
func (s nativeDialog) SetTitle(title string) {
	var arg0 *C.GtkNativeDialog
	var arg1 *C.char

	arg0 = (*C.GtkNativeDialog)(s.Native())
	arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_native_dialog_set_title(arg0, arg1)
}

// SetTransientFor: dialog windows should be set transient for the main
// application window they were spawned from. This allows [window
// managers][gtk-X11-arch] to e.g. keep the dialog on top of the main
// window, or center the dialog over the main window.
//
// Passing nil for @parent unsets the current transient window.
func (s nativeDialog) SetTransientFor(parent Window) {
	var arg0 *C.GtkNativeDialog
	var arg1 *C.GtkWindow

	arg0 = (*C.GtkNativeDialog)(s.Native())
	arg1 = (*C.GtkWindow)(parent.Native())

	C.gtk_native_dialog_set_transient_for(arg0, arg1)
}

// Show shows the dialog on the display, allowing the user to interact with
// it. When the user accepts the state of the dialog the dialog will be
// automatically hidden and the NativeDialog::response signal will be
// emitted.
//
// Multiple calls while the dialog is visible will be ignored.
func (s nativeDialog) Show() {
	var arg0 *C.GtkNativeDialog

	arg0 = (*C.GtkNativeDialog)(s.Native())

	C.gtk_native_dialog_show(arg0)
}
