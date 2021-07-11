// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_im_context_get_type()), F: marshalIMContexter},
	})
}

// IMContextOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type IMContextOverrider interface {
	//
	Commit(str string)
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
	// FilterKeypress: allow an input method to internally handle key press and
	// release events.
	//
	// If this function returns true, then no further processing should be done
	// for this key event.
	FilterKeypress(event gdk.Eventer) bool
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
	//
	// Deprecated: Use [method@Gtk.IMContext.get_surrounding_with_selection]
	// instead.
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
	//
	PreeditChanged()
	//
	PreeditEnd()
	//
	PreeditStart()
	// Reset: notify the input method that a change such as a change in cursor
	// position has been made.
	//
	// This will typically cause the input method to clear the preedit state.
	Reset()
	//
	RetrieveSurrounding() bool
	// SetClientWidget: set the client widget for the input context.
	//
	// This is the `GtkWidget` holding the input focus. This widget is used in
	// order to correctly position status windows, and may also be used for
	// purposes internal to the input method.
	SetClientWidget(widget Widgetter)
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
	//
	// Deprecated: Use [method@Gtk.IMContext.set_surrounding_with_selection]
	// instead.
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

// IMContexter describes IMContext's methods.
type IMContexter interface {
	// DeleteSurrounding asks the widget that the input context is attached to
	// delete characters around the cursor position by emitting the
	// GtkIMContext::delete_surrounding signal.
	DeleteSurrounding(offset int, nChars int) bool
	// FilterKeypress: allow an input method to internally handle key press and
	// release events.
	FilterKeypress(event gdk.Eventer) bool
	// FocusIn: notify the input method that the widget to which this input
	// context corresponds has gained focus.
	FocusIn()
	// FocusOut: notify the input method that the widget to which this input
	// context corresponds has lost focus.
	FocusOut()
	// PreeditString: retrieve the current preedit string for the input context,
	// and a list of attributes to apply to the string.
	PreeditString() (string, *pango.AttrList, int)
	// Surrounding retrieves context around the insertion point.
	Surrounding() (string, int, bool)
	// SurroundingWithSelection retrieves context around the insertion point.
	SurroundingWithSelection() (text string, cursorIndex int, anchorIndex int, ok bool)
	// Reset: notify the input method that a change such as a change in cursor
	// position has been made.
	Reset()
	// SetClientWidget: set the client widget for the input context.
	SetClientWidget(widget Widgetter)
	// SetCursorLocation: notify the input method that a change in cursor
	// position has been made.
	SetCursorLocation(area *gdk.Rectangle)
	// SetSurrounding sets surrounding context around the insertion point and
	// preedit string.
	SetSurrounding(text string, len int, cursorIndex int)
	// SetSurroundingWithSelection sets surrounding context around the insertion
	// point and preedit string.
	SetSurroundingWithSelection(text string, len int, cursorIndex int, anchorIndex int)
	// SetUsePreedit sets whether the IM context should use the preedit string
	// to display feedback.
	SetUsePreedit(usePreedit bool)
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
type IMContext struct {
	*externglib.Object
}

var (
	_ IMContexter     = (*IMContext)(nil)
	_ gextras.Nativer = (*IMContext)(nil)
)

func wrapIMContext(obj *externglib.Object) IMContexter {
	return &IMContext{
		Object: obj,
	}
}

func marshalIMContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapIMContext(obj), nil
}

// DeleteSurrounding asks the widget that the input context is attached to
// delete characters around the cursor position by emitting the
// GtkIMContext::delete_surrounding signal.
//
// Note that @offset and @n_chars are in characters not in bytes which differs
// from the usage other places in IMContext.
//
// In order to use this function, you should first call
// gtk_im_context_get_surrounding() to get the current context, and call this
// function immediately afterwards to make sure that you know what you are
// deleting. You should also account for the fact that even if the signal was
// handled, the input context might not have deleted all the characters that
// were requested to be deleted.
//
// This function is used by an input method that wants to make subsitutions in
// the existing text in response to new input. It is not useful for
// applications.
func (context *IMContext) DeleteSurrounding(offset int, nChars int) bool {
	var _arg0 *C.GtkIMContext // out
	var _arg1 C.int           // out
	var _arg2 C.int           // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.int(offset)
	_arg2 = C.int(nChars)

	_cret = C.gtk_im_context_delete_surrounding(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FilterKeypress: allow an input method to internally handle key press and
// release events.
//
// If this function returns true, then no further processing should be done for
// this key event.
func (context *IMContext) FilterKeypress(event gdk.Eventer) bool {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.GdkEvent     // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer((event).(gextras.Nativer).Native()))

	_cret = C.gtk_im_context_filter_keypress(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FocusIn: notify the input method that the widget to which this input context
// corresponds has gained focus.
//
// The input method may, for example, change the displayed feedback to reflect
// this change.
func (context *IMContext) FocusIn() {
	var _arg0 *C.GtkIMContext // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))

	C.gtk_im_context_focus_in(_arg0)
}

// FocusOut: notify the input method that the widget to which this input context
// corresponds has lost focus.
//
// The input method may, for example, change the displayed feedback or reset the
// contexts state to reflect this change.
func (context *IMContext) FocusOut() {
	var _arg0 *C.GtkIMContext // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))

	C.gtk_im_context_focus_out(_arg0)
}

// PreeditString: retrieve the current preedit string for the input context, and
// a list of attributes to apply to the string.
//
// This string should be displayed inserted at the insertion point.
func (context *IMContext) PreeditString() (string, *pango.AttrList, int) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.char         // in
	var _attrs *pango.AttrList
	var _arg3 C.int // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))

	C.gtk_im_context_get_preedit_string(_arg0, &_arg1, (**C.PangoAttrList)(unsafe.Pointer(&_attrs)), &_arg3)

	var _str string // out

	var _cursorPos int // out

	_str = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	defer C.free(unsafe.Pointer(_arg1))

	_cursorPos = int(_arg3)

	return _str, _attrs, _cursorPos
}

// Surrounding retrieves context around the insertion point.
//
// Input methods typically want context in order to constrain input text based
// on existing text; this is important for languages such as Thai where only
// some sequences of characters are allowed.
//
// This function is implemented by emitting the
// [signal@Gtk.IMContext::retrieve-surrounding] signal on the input method; in
// response to this signal, a widget should provide as much context as is
// available, up to an entire paragraph, by calling
// [method@Gtk.IMContext.set_surrounding].
//
// Note that there is no obligation for a widget to respond to the
// `::retrieve-surrounding` signal, so input methods must be prepared to
// function without context.
//
// Deprecated: Use [method@Gtk.IMContext.get_surrounding_with_selection]
// instead.
func (context *IMContext) Surrounding() (string, int, bool) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.char         // in
	var _arg2 C.int           // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_im_context_get_surrounding(_arg0, &_arg1, &_arg2)

	var _text string     // out
	var _cursorIndex int // out
	var _ok bool         // out

	_text = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	defer C.free(unsafe.Pointer(_arg1))
	_cursorIndex = int(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _text, _cursorIndex, _ok
}

// SurroundingWithSelection retrieves context around the insertion point.
//
// Input methods typically want context in order to constrain input text based
// on existing text; this is important for languages such as Thai where only
// some sequences of characters are allowed.
//
// This function is implemented by emitting the
// [signal@Gtk.IMContext::retrieve-surrounding] signal on the input method; in
// response to this signal, a widget should provide as much context as is
// available, up to an entire paragraph, by calling
// [method@Gtk.IMContext.set_surrounding_with_selection].
//
// Note that there is no obligation for a widget to respond to the
// `::retrieve-surrounding` signal, so input methods must be prepared to
// function without context.
func (context *IMContext) SurroundingWithSelection() (text string, cursorIndex int, anchorIndex int, ok bool) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.char         // in
	var _arg2 C.int           // in
	var _arg3 C.int           // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_im_context_get_surrounding_with_selection(_arg0, &_arg1, &_arg2, &_arg3)

	var _text string     // out
	var _cursorIndex int // out
	var _anchorIndex int // out
	var _ok bool         // out

	_text = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	defer C.free(unsafe.Pointer(_arg1))
	_cursorIndex = int(_arg2)
	_anchorIndex = int(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _text, _cursorIndex, _anchorIndex, _ok
}

// Reset: notify the input method that a change such as a change in cursor
// position has been made.
//
// This will typically cause the input method to clear the preedit state.
func (context *IMContext) Reset() {
	var _arg0 *C.GtkIMContext // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))

	C.gtk_im_context_reset(_arg0)
}

// SetClientWidget: set the client widget for the input context.
//
// This is the `GtkWidget` holding the input focus. This widget is used in order
// to correctly position status windows, and may also be used for purposes
// internal to the input method.
func (context *IMContext) SetClientWidget(widget Widgetter) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))

	C.gtk_im_context_set_client_widget(_arg0, _arg1)
}

// SetCursorLocation: notify the input method that a change in cursor position
// has been made.
//
// The location is relative to the client window.
func (context *IMContext) SetCursorLocation(area *gdk.Rectangle) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.GdkRectangle // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(area))

	C.gtk_im_context_set_cursor_location(_arg0, _arg1)
}

// SetSurrounding sets surrounding context around the insertion point and
// preedit string.
//
// This function is expected to be called in response to the
// [signal@Gtk.IMContext::retrieve-surrounding] signal, and will likely have no
// effect if called at other times.
//
// Deprecated: Use [method@Gtk.IMContext.set_surrounding_with_selection]
// instead.
func (context *IMContext) SetSurrounding(text string, len int, cursorIndex int) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.char         // out
	var _arg2 C.int           // out
	var _arg3 C.int           // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(len)
	_arg3 = C.int(cursorIndex)

	C.gtk_im_context_set_surrounding(_arg0, _arg1, _arg2, _arg3)
}

// SetSurroundingWithSelection sets surrounding context around the insertion
// point and preedit string. This function is expected to be called in response
// to the GtkIMContext::retrieve_surrounding signal, and will likely have no
// effect if called at other times.
func (context *IMContext) SetSurroundingWithSelection(text string, len int, cursorIndex int, anchorIndex int) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.char         // out
	var _arg2 C.int           // out
	var _arg3 C.int           // out
	var _arg4 C.int           // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(len)
	_arg3 = C.int(cursorIndex)
	_arg4 = C.int(anchorIndex)

	C.gtk_im_context_set_surrounding_with_selection(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// SetUsePreedit sets whether the IM context should use the preedit string to
// display feedback.
//
// If @use_preedit is false (default is true), then the IM context may use some
// other method to display feedback, such as displaying it in a child of the
// root window.
func (context *IMContext) SetUsePreedit(usePreedit bool) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	if usePreedit {
		_arg1 = C.TRUE
	}

	C.gtk_im_context_set_use_preedit(_arg0, _arg1)
}
