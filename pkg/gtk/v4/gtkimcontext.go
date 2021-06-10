// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_im_context_get_type()), F: marshalIMContext},
	})
}

// IMContext: `GtkIMContext` defines the interface for GTK input methods.
//
// `GtkIMContext` is used by GTK text input widgets like `GtkText` to map from
// key events to Unicode character strings.
//
// By default, GTK uses a platform-dependent default input method. On Windows,
// the default implementation is IME-based and on Wayland, it is using the
// Wayland text protocol. The choice can be overridden programmatically via the
// [property@Gtk.Settings:gtk-im-module] setting. Users may set the
// `GTK_IM_MODULE` environment variable to override the default.
//
// Text widgets have a :im-module property (e.g.
// [property@Gtk.TextView:im-module]) that may also be used to set input methods
// for specific widget instances. For instance, a certain entry widget might be
// expected to contain certain characters which would be easier to input with a
// specific input method.
//
// An input method may consume multiple key events in sequence before finally
// outputting the composed result. This is called *preediting*, and an input
// method may provide feedback about this process by displaying the intermediate
// composition states as preedit text.
//
// For instance, the built-in GTK input method `GtkIMContextSimple` implements
// the input of arbitrary Unicode code points by holding down the
// <kbd>Control</kbd> and <kbd>Shift</kbd> keys and then typing <kbd>U</kbd>
// followed by the hexadecimal digits of the code point. When releasing the
// <kbd>Control</kbd> and <kbd>Shift</kbd> keys, preediting ends and the
// character is inserted as text. For example,
//
//    Ctrl+Shift+u 2 0 A C
//
// results in the € sign.
//
// Additional input methods can be made available for use by GTK widgets as
// loadable modules. An input method module is a small shared library which
// provides a `GIOExtension` for the extension point named "gtk-im-module".
type IMContext interface {
	gextras.Objector

	// DeleteSurrounding asks the widget that the input context is attached to
	// delete characters around the cursor position by emitting the
	// GtkIMContext::delete_surrounding signal.
	//
	// Note that @offset and @n_chars are in characters not in bytes which
	// differs from the usage other places in IMContext.
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
	// FilterKey: allow an input method to forward key press and release events
	// to another input methodm without necessarily having a `GdkEvent`
	// available.
	FilterKey(press bool, surface gdk.Surface, device gdk.Device, time uint32, keycode uint, state gdk.ModifierType, group int) bool
	// FilterKeypress: allow an input method to internally handle key press and
	// release events.
	//
	// If this function returns true, then no further processing should be done
	// for this key event.
	FilterKeypress(event gdk.Event) bool
	// FocusIn: notify the input method that the widget to which this input
	// context corresponds has gained focus.
	//
	// The input method may, for example, change the displayed feedback to
	// reflect this change.
	FocusIn()
	// FocusOut: notify the input method that the widget to which this input
	// context corresponds has lost focus.
	//
	// The input method may, for example, change the displayed feedback or reset
	// the contexts state to reflect this change.
	FocusOut()
	// PreeditString: retrieve the current preedit string for the input context,
	// and a list of attributes to apply to the string.
	//
	// This string should be displayed inserted at the insertion point.
	PreeditString() (string, *pango.AttrList, int)
	// Surrounding retrieves context around the insertion point.
	//
	// Input methods typically want context in order to constrain input text
	// based on existing text; this is important for languages such as Thai
	// where only some sequences of characters are allowed.
	//
	// This function is implemented by emitting the
	// [signal@Gtk.IMContext::retrieve-surrounding] signal on the input method;
	// in response to this signal, a widget should provide as much context as is
	// available, up to an entire paragraph, by calling
	// [method@Gtk.IMContext.set_surrounding].
	//
	// Note that there is no obligation for a widget to respond to the
	// `::retrieve-surrounding` signal, so input methods must be prepared to
	// function without context.
	Surrounding() (string, int, bool)
	// SurroundingWithSelection retrieves context around the insertion point.
	//
	// Input methods typically want context in order to constrain input text
	// based on existing text; this is important for languages such as Thai
	// where only some sequences of characters are allowed.
	//
	// This function is implemented by emitting the
	// [signal@Gtk.IMContext::retrieve-surrounding] signal on the input method;
	// in response to this signal, a widget should provide as much context as is
	// available, up to an entire paragraph, by calling
	// [method@Gtk.IMContext.set_surrounding_with_selection].
	//
	// Note that there is no obligation for a widget to respond to the
	// `::retrieve-surrounding` signal, so input methods must be prepared to
	// function without context.
	SurroundingWithSelection() (text string, cursorIndex int, anchorIndex int, ok bool)
	// Reset: notify the input method that a change such as a change in cursor
	// position has been made.
	//
	// This will typically cause the input method to clear the preedit state.
	Reset()
	// SetClientWidget: set the client widget for the input context.
	//
	// This is the `GtkWidget` holding the input focus. This widget is used in
	// order to correctly position status windows, and may also be used for
	// purposes internal to the input method.
	SetClientWidget(widget Widget)
	// SetCursorLocation: notify the input method that a change in cursor
	// position has been made.
	//
	// The location is relative to the client window.
	SetCursorLocation(area *gdk.Rectangle)
	// SetSurrounding sets surrounding context around the insertion point and
	// preedit string.
	//
	// This function is expected to be called in response to the
	// [signal@Gtk.IMContext::retrieve-surrounding] signal, and will likely have
	// no effect if called at other times.
	SetSurrounding(text string, len int, cursorIndex int)
	// SetSurroundingWithSelection sets surrounding context around the insertion
	// point and preedit string. This function is expected to be called in
	// response to the GtkIMContext::retrieve_surrounding signal, and will
	// likely have no effect if called at other times.
	SetSurroundingWithSelection(text string, len int, cursorIndex int, anchorIndex int)
	// SetUsePreedit sets whether the IM context should use the preedit string
	// to display feedback.
	//
	// If @use_preedit is false (default is true), then the IM context may use
	// some other method to display feedback, such as displaying it in a child
	// of the root window.
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
// delete characters around the cursor position by emitting the
// GtkIMContext::delete_surrounding signal.
//
// Note that @offset and @n_chars are in characters not in bytes which
// differs from the usage other places in IMContext.
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
	var _arg0 *C.GtkIMContext
	var _arg1 C.int
	var _arg2 C.int

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(offset)
	_arg2 = C.int(nChars)

	var _cret C.gboolean

	_cret = C.gtk_im_context_delete_surrounding(_arg0, _arg1, _arg2)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// FilterKey: allow an input method to forward key press and release events
// to another input methodm without necessarily having a `GdkEvent`
// available.
func (c imContext) FilterKey(press bool, surface gdk.Surface, device gdk.Device, time uint32, keycode uint, state gdk.ModifierType, group int) bool {
	var _arg0 *C.GtkIMContext
	var _arg1 C.gboolean
	var _arg2 *C.GdkSurface
	var _arg3 *C.GdkDevice
	var _arg4 C.guint32
	var _arg5 C.guint
	var _arg6 C.GdkModifierType
	var _arg7 C.int

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	if press {
		_arg1 = C.gboolean(1)
	}
	_arg2 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	_arg3 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg4 = C.guint32(time)
	_arg5 = C.guint(keycode)
	_arg6 = (C.GdkModifierType)(state)
	_arg7 = C.int(group)

	var _cret C.gboolean

	_cret = C.gtk_im_context_filter_key(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// FilterKeypress: allow an input method to internally handle key press and
// release events.
//
// If this function returns true, then no further processing should be done
// for this key event.
func (c imContext) FilterKeypress(event gdk.Event) bool {
	var _arg0 *C.GtkIMContext
	var _arg1 *C.GdkEvent

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))

	var _cret C.gboolean

	_cret = C.gtk_im_context_filter_keypress(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// FocusIn: notify the input method that the widget to which this input
// context corresponds has gained focus.
//
// The input method may, for example, change the displayed feedback to
// reflect this change.
func (c imContext) FocusIn() {
	var _arg0 *C.GtkIMContext

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))

	C.gtk_im_context_focus_in(_arg0)
}

// FocusOut: notify the input method that the widget to which this input
// context corresponds has lost focus.
//
// The input method may, for example, change the displayed feedback or reset
// the contexts state to reflect this change.
func (c imContext) FocusOut() {
	var _arg0 *C.GtkIMContext

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))

	C.gtk_im_context_focus_out(_arg0)
}

// PreeditString: retrieve the current preedit string for the input context,
// and a list of attributes to apply to the string.
//
// This string should be displayed inserted at the insertion point.
func (c imContext) PreeditString() (string, *pango.AttrList, int) {
	var _arg0 *C.GtkIMContext

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))

	var _arg1 *C.char
	var _attrs *pango.AttrList
	var _arg3 C.int

	C.gtk_im_context_get_preedit_string(_arg0, &_arg1, (**C.PangoAttrList)(unsafe.Pointer(&_attrs)), &_arg3)

	var _str string

	var _cursorPos int

	_str = C.GoString(_arg1)
	defer C.free(unsafe.Pointer(_arg1))

	_cursorPos = (int)(_arg3)

	return _str, _attrs, _cursorPos
}

// Surrounding retrieves context around the insertion point.
//
// Input methods typically want context in order to constrain input text
// based on existing text; this is important for languages such as Thai
// where only some sequences of characters are allowed.
//
// This function is implemented by emitting the
// [signal@Gtk.IMContext::retrieve-surrounding] signal on the input method;
// in response to this signal, a widget should provide as much context as is
// available, up to an entire paragraph, by calling
// [method@Gtk.IMContext.set_surrounding].
//
// Note that there is no obligation for a widget to respond to the
// `::retrieve-surrounding` signal, so input methods must be prepared to
// function without context.
func (c imContext) Surrounding() (string, int, bool) {
	var _arg0 *C.GtkIMContext

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))

	var _arg1 *C.char
	var _arg2 C.int
	var _cret C.gboolean

	_cret = C.gtk_im_context_get_surrounding(_arg0, &_arg1, &_arg2)

	var _text string
	var _cursorIndex int
	var _ok bool

	_text = C.GoString(_arg1)
	defer C.free(unsafe.Pointer(_arg1))
	_cursorIndex = (int)(_arg2)
	if _cret {
		_ok = true
	}

	return _text, _cursorIndex, _ok
}

// SurroundingWithSelection retrieves context around the insertion point.
//
// Input methods typically want context in order to constrain input text
// based on existing text; this is important for languages such as Thai
// where only some sequences of characters are allowed.
//
// This function is implemented by emitting the
// [signal@Gtk.IMContext::retrieve-surrounding] signal on the input method;
// in response to this signal, a widget should provide as much context as is
// available, up to an entire paragraph, by calling
// [method@Gtk.IMContext.set_surrounding_with_selection].
//
// Note that there is no obligation for a widget to respond to the
// `::retrieve-surrounding` signal, so input methods must be prepared to
// function without context.
func (c imContext) SurroundingWithSelection() (text string, cursorIndex int, anchorIndex int, ok bool) {
	var _arg0 *C.GtkIMContext

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))

	var _arg1 *C.char
	var _arg2 C.int
	var _arg3 C.int
	var _cret C.gboolean

	_cret = C.gtk_im_context_get_surrounding_with_selection(_arg0, &_arg1, &_arg2, &_arg3)

	var _text string
	var _cursorIndex int
	var _anchorIndex int
	var _ok bool

	_text = C.GoString(_arg1)
	defer C.free(unsafe.Pointer(_arg1))
	_cursorIndex = (int)(_arg2)
	_anchorIndex = (int)(_arg3)
	if _cret {
		_ok = true
	}

	return _text, _cursorIndex, _anchorIndex, _ok
}

// Reset: notify the input method that a change such as a change in cursor
// position has been made.
//
// This will typically cause the input method to clear the preedit state.
func (c imContext) Reset() {
	var _arg0 *C.GtkIMContext

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))

	C.gtk_im_context_reset(_arg0)
}

// SetClientWidget: set the client widget for the input context.
//
// This is the `GtkWidget` holding the input focus. This widget is used in
// order to correctly position status windows, and may also be used for
// purposes internal to the input method.
func (c imContext) SetClientWidget(widget Widget) {
	var _arg0 *C.GtkIMContext
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_im_context_set_client_widget(_arg0, _arg1)
}

// SetCursorLocation: notify the input method that a change in cursor
// position has been made.
//
// The location is relative to the client window.
func (c imContext) SetCursorLocation(area *gdk.Rectangle) {
	var _arg0 *C.GtkIMContext
	var _arg1 *C.GdkRectangle

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(area.Native()))

	C.gtk_im_context_set_cursor_location(_arg0, _arg1)
}

// SetSurrounding sets surrounding context around the insertion point and
// preedit string.
//
// This function is expected to be called in response to the
// [signal@Gtk.IMContext::retrieve-surrounding] signal, and will likely have
// no effect if called at other times.
func (c imContext) SetSurrounding(text string, len int, cursorIndex int) {
	var _arg0 *C.GtkIMContext
	var _arg1 *C.char
	var _arg2 C.int
	var _arg3 C.int

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(len)
	_arg3 = C.int(cursorIndex)

	C.gtk_im_context_set_surrounding(_arg0, _arg1, _arg2, _arg3)
}

// SetSurroundingWithSelection sets surrounding context around the insertion
// point and preedit string. This function is expected to be called in
// response to the GtkIMContext::retrieve_surrounding signal, and will
// likely have no effect if called at other times.
func (c imContext) SetSurroundingWithSelection(text string, len int, cursorIndex int, anchorIndex int) {
	var _arg0 *C.GtkIMContext
	var _arg1 *C.char
	var _arg2 C.int
	var _arg3 C.int
	var _arg4 C.int

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(len)
	_arg3 = C.int(cursorIndex)
	_arg4 = C.int(anchorIndex)

	C.gtk_im_context_set_surrounding_with_selection(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// SetUsePreedit sets whether the IM context should use the preedit string
// to display feedback.
//
// If @use_preedit is false (default is true), then the IM context may use
// some other method to display feedback, such as displaying it in a child
// of the root window.
func (c imContext) SetUsePreedit(usePreedit bool) {
	var _arg0 *C.GtkIMContext
	var _arg1 C.gboolean

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	if usePreedit {
		_arg1 = C.gboolean(1)
	}

	C.gtk_im_context_set_use_preedit(_arg0, _arg1)
}
