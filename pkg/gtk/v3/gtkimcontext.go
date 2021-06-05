// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_im_context_get_type()), F: marshalIMContext},
	})
}

// IMContext defines the interface for GTK+ input methods. An input method is
// used by GTK+ text input widgets like Entry to map from key events to Unicode
// character strings.
//
// The default input method can be set programmatically via the
// Settings:gtk-im-module GtkSettings property. Alternatively, you may set the
// GTK_IM_MODULE environment variable as documented in [Running GTK+
// Applications][gtk-running].
//
// The Entry Entry:im-module and TextView TextView:im-module properties may also
// be used to set input methods for specific widget instances. For instance, a
// certain entry widget might be expected to contain certain characters which
// would be easier to input with a certain input method.
//
// An input method may consume multiple key events in sequence and finally
// output the composed result. This is called preediting, and an input method
// may provide feedback about this process by displaying the intermediate
// composition states as preedit text. For instance, the default GTK+ input
// method implements the input of arbitrary Unicode code points by holding down
// the Control and Shift keys and then typing “U” followed by the hexadecimal
// digits of the code point. When releasing the Control and Shift keys,
// preediting ends and the character is inserted as text. Ctrl+Shift+u20AC for
// example results in the € sign.
//
// Additional input methods can be made available for use by GTK+ widgets as
// loadable modules. An input method module is a small shared library which
// implements a subclass of IMContext or IMContextSimple and exports these four
// functions:
//
//    GtkIMContext * im_module_create(const gchar *context_id);
//
// This function should return a pointer to a newly created instance of the
// IMContext subclass identified by @context_id. The context ID is the same as
// specified in the IMContextInfo array returned by im_module_list().
//
// After a new loadable input method module has been installed on the system,
// the configuration file `gtk.immodules` needs to be regenerated by
// [gtk-query-immodules-3.0][gtk-query-immodules-3.0], in order for the new
// input method to become available to GTK+ applications.
type IMContext interface {
	gextras.Objector

	// DeleteSurrounding asks the widget that the input context is attached to
	// to delete characters around the cursor position by emitting the
	// GtkIMContext::delete_surrounding signal. Note that @offset and @n_chars
	// are in characters not in bytes which differs from the usage other places
	// in IMContext.
	//
	// In order to use this function, you should first call
	// gtk_im_context_get_surrounding() to get the current context, and call
	// this function immediately afterwards to make sure that you know what you
	// are deleting. You should also account for the fact that even if the
	// signal was handled, the input context might not have deleted all the
	// characters that were requested to be deleted.
	//
	// This function is used by an input method that wants to make subsitutions
	// in the existing text in response to new input. It is not useful for
	// applications.
	DeleteSurrounding(offset int, nChars int) bool
	// FilterKeypress: allow an input method to internally handle key press and
	// release events. If this function returns true, then no further processing
	// should be done for this key event.
	FilterKeypress(event *gdk.EventKey) bool
	// FocusIn: notify the input method that the widget to which this input
	// context corresponds has gained focus. The input method may, for example,
	// change the displayed feedback to reflect this change.
	FocusIn()
	// FocusOut: notify the input method that the widget to which this input
	// context corresponds has lost focus. The input method may, for example,
	// change the displayed feedback or reset the contexts state to reflect this
	// change.
	FocusOut()
	// PreeditString: retrieve the current preedit string for the input context,
	// and a list of attributes to apply to the string. This string should be
	// displayed inserted at the insertion point.
	PreeditString() (str string, attrs *pango.AttrList, cursorPos int)
	// Surrounding retrieves context around the insertion point. Input methods
	// typically want context in order to constrain input text based on existing
	// text; this is important for languages such as Thai where only some
	// sequences of characters are allowed.
	//
	// This function is implemented by emitting the
	// GtkIMContext::retrieve_surrounding signal on the input method; in
	// response to this signal, a widget should provide as much context as is
	// available, up to an entire paragraph, by calling
	// gtk_im_context_set_surrounding(). Note that there is no obligation for a
	// widget to respond to the ::retrieve_surrounding signal, so input methods
	// must be prepared to function without context.
	Surrounding() (text string, cursorIndex int, ok bool)
	// Reset: notify the input method that a change such as a change in cursor
	// position has been made. This will typically cause the input method to
	// clear the preedit state.
	Reset()
	// SetClientWindow: set the client window for the input context; this is the
	// Window in which the input appears. This window is used in order to
	// correctly position status windows, and may also be used for purposes
	// internal to the input method.
	SetClientWindow(window gdk.Window)
	// SetCursorLocation: notify the input method that a change in cursor
	// position has been made. The location is relative to the client window.
	SetCursorLocation(area *gdk.Rectangle)
	// SetSurrounding sets surrounding context around the insertion point and
	// preedit string. This function is expected to be called in response to the
	// GtkIMContext::retrieve_surrounding signal, and will likely have no effect
	// if called at other times.
	SetSurrounding(text string, len int, cursorIndex int)
	// SetUsePreedit sets whether the IM context should use the preedit string
	// to display feedback. If @use_preedit is FALSE (default is TRUE), then the
	// IM context may use some other method to display feedback, such as
	// displaying it in a child of the root window.
	SetUsePreedit(usePreedit bool)
}

// imContext implements the IMContext interface.
type imContext struct {
	gextras.Objector
}

var _ IMContext = (*imContext)(nil)

// WrapIMContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapIMContext(obj *externglib.Object) IMContext {
	return IMContext{
		Objector: obj,
	}
}

func marshalIMContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapIMContext(obj), nil
}

// DeleteSurrounding asks the widget that the input context is attached to
// to delete characters around the cursor position by emitting the
// GtkIMContext::delete_surrounding signal. Note that @offset and @n_chars
// are in characters not in bytes which differs from the usage other places
// in IMContext.
//
// In order to use this function, you should first call
// gtk_im_context_get_surrounding() to get the current context, and call
// this function immediately afterwards to make sure that you know what you
// are deleting. You should also account for the fact that even if the
// signal was handled, the input context might not have deleted all the
// characters that were requested to be deleted.
//
// This function is used by an input method that wants to make subsitutions
// in the existing text in response to new input. It is not useful for
// applications.
func (c imContext) DeleteSurrounding(offset int, nChars int) bool {
	var arg0 *C.GtkIMContext
	var arg1 C.gint
	var arg2 C.gint

	arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	arg1 = C.gint(offset)
	arg2 = C.gint(nChars)

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_im_context_delete_surrounding(arg0, offset, nChars)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// FilterKeypress: allow an input method to internally handle key press and
// release events. If this function returns true, then no further processing
// should be done for this key event.
func (c imContext) FilterKeypress(event *gdk.EventKey) bool {
	var arg0 *C.GtkIMContext
	var arg1 *C.GdkEventKey

	arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GdkEventKey)(unsafe.Pointer(event.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_im_context_filter_keypress(arg0, event)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// FocusIn: notify the input method that the widget to which this input
// context corresponds has gained focus. The input method may, for example,
// change the displayed feedback to reflect this change.
func (c imContext) FocusIn() {
	var arg0 *C.GtkIMContext

	arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))

	C.gtk_im_context_focus_in(arg0)
}

// FocusOut: notify the input method that the widget to which this input
// context corresponds has lost focus. The input method may, for example,
// change the displayed feedback or reset the contexts state to reflect this
// change.
func (c imContext) FocusOut() {
	var arg0 *C.GtkIMContext

	arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))

	C.gtk_im_context_focus_out(arg0)
}

// PreeditString: retrieve the current preedit string for the input context,
// and a list of attributes to apply to the string. This string should be
// displayed inserted at the insertion point.
func (c imContext) PreeditString() (str string, attrs *pango.AttrList, cursorPos int) {
	var arg0 *C.GtkIMContext

	arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))

	var arg1 *C.gchar
	var ret1 string
	var arg2 *C.PangoAttrList
	var ret2 **pango.AttrList
	var arg3 C.gint
	var ret3 int

	C.gtk_im_context_get_preedit_string(arg0, &arg1, &arg2, &arg3)

	*ret1 = C.GoString(arg1)
	defer C.free(unsafe.Pointer(arg1))
	*ret2 = pango.WrapAttrList(unsafe.Pointer(arg2))
	runtime.SetFinalizer(*ret2, func(v **pango.AttrList) {
		C.free(unsafe.Pointer(v.Native()))
	})
	*ret3 = C.gint(arg3)

	return ret1, ret2, ret3
}

// Surrounding retrieves context around the insertion point. Input methods
// typically want context in order to constrain input text based on existing
// text; this is important for languages such as Thai where only some
// sequences of characters are allowed.
//
// This function is implemented by emitting the
// GtkIMContext::retrieve_surrounding signal on the input method; in
// response to this signal, a widget should provide as much context as is
// available, up to an entire paragraph, by calling
// gtk_im_context_set_surrounding(). Note that there is no obligation for a
// widget to respond to the ::retrieve_surrounding signal, so input methods
// must be prepared to function without context.
func (c imContext) Surrounding() (text string, cursorIndex int, ok bool) {
	var arg0 *C.GtkIMContext

	arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))

	var arg1 *C.gchar
	var ret1 string
	var arg2 C.gint
	var ret2 int
	var cret C.gboolean
	var ret3 bool

	cret = C.gtk_im_context_get_surrounding(arg0, &arg1, &arg2)

	*ret1 = C.GoString(arg1)
	defer C.free(unsafe.Pointer(arg1))
	*ret2 = C.gint(arg2)
	ret3 = C.bool(cret) != C.false

	return ret1, ret2, ret3
}

// Reset: notify the input method that a change such as a change in cursor
// position has been made. This will typically cause the input method to
// clear the preedit state.
func (c imContext) Reset() {
	var arg0 *C.GtkIMContext

	arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))

	C.gtk_im_context_reset(arg0)
}

// SetClientWindow: set the client window for the input context; this is the
// Window in which the input appears. This window is used in order to
// correctly position status windows, and may also be used for purposes
// internal to the input method.
func (c imContext) SetClientWindow(window gdk.Window) {
	var arg0 *C.GtkIMContext
	var arg1 *C.GdkWindow

	arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_im_context_set_client_window(arg0, window)
}

// SetCursorLocation: notify the input method that a change in cursor
// position has been made. The location is relative to the client window.
func (c imContext) SetCursorLocation(area *gdk.Rectangle) {
	var arg0 *C.GtkIMContext
	var arg1 *C.GdkRectangle

	arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GdkRectangle)(unsafe.Pointer(area.Native()))

	C.gtk_im_context_set_cursor_location(arg0, area)
}

// SetSurrounding sets surrounding context around the insertion point and
// preedit string. This function is expected to be called in response to the
// GtkIMContext::retrieve_surrounding signal, and will likely have no effect
// if called at other times.
func (c imContext) SetSurrounding(text string, len int, cursorIndex int) {
	var arg0 *C.GtkIMContext
	var arg1 *C.gchar
	var arg2 C.gint
	var arg3 C.gint

	arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gint(len)
	arg3 = C.gint(cursorIndex)

	C.gtk_im_context_set_surrounding(arg0, text, len, cursorIndex)
}

// SetUsePreedit sets whether the IM context should use the preedit string
// to display feedback. If @use_preedit is FALSE (default is TRUE), then the
// IM context may use some other method to display feedback, such as
// displaying it in a child of the root window.
func (c imContext) SetUsePreedit(usePreedit bool) {
	var arg0 *C.GtkIMContext
	var arg1 C.gboolean

	arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	if usePreedit {
		arg1 = C.gboolean(1)
	}

	C.gtk_im_context_set_use_preedit(arg0, usePreedit)
}
