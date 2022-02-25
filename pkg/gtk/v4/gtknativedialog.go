// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_NativeDialogClass_hide(GtkNativeDialog*);
// extern void _gotk4_gtk4_NativeDialogClass_response(GtkNativeDialog*, int);
// extern void _gotk4_gtk4_NativeDialogClass_show(GtkNativeDialog*);
// extern void _gotk4_gtk4_NativeDialog_ConnectResponse(gpointer, gint, guintptr);
import "C"

// glib.Type values for gtknativedialog.go.
var GTypeNativeDialog = externglib.Type(C.gtk_native_dialog_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeNativeDialog, F: marshalNativeDialog},
	})
}

// NativeDialogOverrider contains methods that are overridable.
type NativeDialogOverrider interface {
	// Hide hides the dialog if it is visible, aborting any interaction.
	//
	// Once this is called the gtk.NativeDialog::response signal will *not* be
	// emitted until after the next call to gtk.NativeDialog.Show().
	//
	// If the dialog is not visible this does nothing.
	Hide()
	// The function takes the following parameters:
	//
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
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*NativeDialog)(nil)
)

// NativeDialogger describes types inherited from class NativeDialog.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type NativeDialogger interface {
	externglib.Objector
	baseNativeDialog() *NativeDialog
}

var _ NativeDialogger = (*NativeDialog)(nil)

func classInitNativeDialogger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkNativeDialogClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkNativeDialogClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Hide() }); ok {
		pclass.hide = (*[0]byte)(C._gotk4_gtk4_NativeDialogClass_hide)
	}

	if _, ok := goval.(interface{ Response(responseId int) }); ok {
		pclass.response = (*[0]byte)(C._gotk4_gtk4_NativeDialogClass_response)
	}

	if _, ok := goval.(interface{ Show() }); ok {
		pclass.show = (*[0]byte)(C._gotk4_gtk4_NativeDialogClass_show)
	}
}

//export _gotk4_gtk4_NativeDialogClass_hide
func _gotk4_gtk4_NativeDialogClass_hide(arg0 *C.GtkNativeDialog) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Hide() })

	iface.Hide()
}

//export _gotk4_gtk4_NativeDialogClass_response
func _gotk4_gtk4_NativeDialogClass_response(arg0 *C.GtkNativeDialog, arg1 C.int) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Response(responseId int) })

	var _responseId int // out

	_responseId = int(arg1)

	iface.Response(_responseId)
}

//export _gotk4_gtk4_NativeDialogClass_show
func _gotk4_gtk4_NativeDialogClass_show(arg0 *C.GtkNativeDialog) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Show() })

	iface.Show()
}

func wrapNativeDialog(obj *externglib.Object) *NativeDialog {
	return &NativeDialog{
		Object: obj,
	}
}

func marshalNativeDialog(p uintptr) (interface{}, error) {
	return wrapNativeDialog(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *NativeDialog) baseNativeDialog() *NativeDialog {
	return self
}

// BaseNativeDialog returns the underlying base object.
func BaseNativeDialog(obj NativeDialogger) *NativeDialog {
	return obj.baseNativeDialog()
}

//export _gotk4_gtk4_NativeDialog_ConnectResponse
func _gotk4_gtk4_NativeDialog_ConnectResponse(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) {
	var f func(responseId int)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(responseId int))
	}

	var _responseId int // out

	_responseId = int(arg1)

	f(_responseId)
}

// ConnectResponse: emitted when the user responds to the dialog.
//
// When this is called the dialog has been hidden.
//
// If you call gtk.NativeDialog.Hide() before the user responds to the dialog
// this signal will not be emitted.
func (self *NativeDialog) ConnectResponse(f func(responseId int)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "response", false, unsafe.Pointer(C._gotk4_gtk4_NativeDialog_ConnectResponse), f)
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

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(externglib.InternObject(self).Native()))

	C.gtk_native_dialog_destroy(_arg0)
	runtime.KeepAlive(self)
}

// Modal returns whether the dialog is modal.
//
// The function returns the following values:
//
//    - ok: TRUE if the dialog is set to be modal.
//
func (self *NativeDialog) Modal() bool {
	var _arg0 *C.GtkNativeDialog // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_native_dialog_get_modal(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Title gets the title of the GtkNativeDialog.
//
// The function returns the following values:
//
//    - utf8 (optional): title of the dialog, or NULL if none has been set
//      explicitly. The returned string is owned by the widget and must not be
//      modified or freed.
//
func (self *NativeDialog) Title() string {
	var _arg0 *C.GtkNativeDialog // out
	var _cret *C.char            // in

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_native_dialog_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// TransientFor fetches the transient parent for this window.
//
// The function returns the following values:
//
//    - window (optional): transient parent for this window, or NULL if no
//      transient parent has been set.
//
func (self *NativeDialog) TransientFor() *Window {
	var _arg0 *C.GtkNativeDialog // out
	var _cret *C.GtkWindow       // in

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_native_dialog_get_transient_for(_arg0)
	runtime.KeepAlive(self)

	var _window *Window // out

	if _cret != nil {
		_window = wrapWindow(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _window
}

// Visible determines whether the dialog is visible.
//
// The function returns the following values:
//
//    - ok: TRUE if the dialog is visible.
//
func (self *NativeDialog) Visible() bool {
	var _arg0 *C.GtkNativeDialog // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_native_dialog_get_visible(_arg0)
	runtime.KeepAlive(self)

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

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(externglib.InternObject(self).Native()))

	C.gtk_native_dialog_hide(_arg0)
	runtime.KeepAlive(self)
}

// SetModal sets a dialog modal or non-modal.
//
// Modal dialogs prevent interaction with other windows in the same application.
// To keep modal dialogs on top of main application windows, use
// gtk.NativeDialog.SetTransientFor() to make the dialog transient for the
// parent; most window managers will then disallow lowering the dialog below the
// parent.
//
// The function takes the following parameters:
//
//    - modal: whether the window is modal.
//
func (self *NativeDialog) SetModal(modal bool) {
	var _arg0 *C.GtkNativeDialog // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.gtk_native_dialog_set_modal(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(modal)
}

// SetTitle sets the title of the GtkNativeDialog.
//
// The function takes the following parameters:
//
//    - title of the dialog.
//
func (self *NativeDialog) SetTitle(title string) {
	var _arg0 *C.GtkNativeDialog // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_native_dialog_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}

// SetTransientFor: dialog windows should be set transient for the main
// application window they were spawned from.
//
// This allows window managers to e.g. keep the dialog on top of the main
// window, or center the dialog over the main window.
//
// Passing NULL for parent unsets the current transient window.
//
// The function takes the following parameters:
//
//    - parent (optional) window, or NULL.
//
func (self *NativeDialog) SetTransientFor(parent *Window) {
	var _arg0 *C.GtkNativeDialog // out
	var _arg1 *C.GtkWindow       // out

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(externglib.InternObject(parent).Native()))
	}

	C.gtk_native_dialog_set_transient_for(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(parent)
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

	_arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(externglib.InternObject(self).Native()))

	C.gtk_native_dialog_show(_arg0)
	runtime.KeepAlive(self)
}
