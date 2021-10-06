// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
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
	Commit(str string)
	// DeleteSurrounding asks the widget that the input context is attached to
	// delete characters around the cursor position by emitting the
	// GtkIMContext::delete_surrounding signal.
	//
	// Note that offset and n_chars are in characters not in bytes which differs
	// from the usage other places in IMContext.
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
	DeleteSurrounding(offset, nChars int) bool
	// FilterKeypress: allow an input method to internally handle key press and
	// release events.
	//
	// If this function returns TRUE, then no further processing should be done
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
	// gtk.IMContext::retrieve-surrounding signal on the input method; in
	// response to this signal, a widget should provide as much context as is
	// available, up to an entire paragraph, by calling
	// gtk.IMContext.SetSurrounding().
	//
	// Note that there is no obligation for a widget to respond to the
	// ::retrieve-surrounding signal, so input methods must be prepared to
	// function without context.
	//
	// Deprecated: Use gtk.IMContext.GetSurroundingWithSelection() instead.
	Surrounding() (string, int, bool)
	// SurroundingWithSelection retrieves context around the insertion point.
	//
	// Input methods typically want context in order to constrain input text
	// based on existing text; this is important for languages such as Thai
	// where only some sequences of characters are allowed.
	//
	// This function is implemented by emitting the
	// gtk.IMContext::retrieve-surrounding signal on the input method; in
	// response to this signal, a widget should provide as much context as is
	// available, up to an entire paragraph, by calling
	// gtk.IMContext.SetSurroundingWithSelection().
	//
	// Note that there is no obligation for a widget to respond to the
	// ::retrieve-surrounding signal, so input methods must be prepared to
	// function without context.
	SurroundingWithSelection() (text string, cursorIndex int, anchorIndex int, ok bool)
	PreeditChanged()
	PreeditEnd()
	PreeditStart()
	// Reset: notify the input method that a change such as a change in cursor
	// position has been made.
	//
	// This will typically cause the input method to clear the preedit state.
	Reset()
	RetrieveSurrounding() bool
	// SetClientWidget: set the client widget for the input context.
	//
	// This is the GtkWidget holding the input focus. This widget is used in
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
	// gtk.IMContext::retrieve-surrounding signal, and will likely have no
	// effect if called at other times.
	//
	// Deprecated: Use gtk.IMContext.SetSurroundingWithSelection() instead.
	SetSurrounding(text string, len, cursorIndex int)
	// SetSurroundingWithSelection sets surrounding context around the insertion
	// point and preedit string. This function is expected to be called in
	// response to the GtkIMContext::retrieve_surrounding signal, and will
	// likely have no effect if called at other times.
	SetSurroundingWithSelection(text string, len, cursorIndex, anchorIndex int)
	// SetUsePreedit sets whether the IM context should use the preedit string
	// to display feedback.
	//
	// If use_preedit is FALSE (default is TRUE), then the IM context may use
	// some other method to display feedback, such as displaying it in a child
	// of the root window.
	SetUsePreedit(usePreedit bool)
}

// IMContext: GtkIMContext defines the interface for GTK input methods.
//
// GtkIMContext is used by GTK text input widgets like GtkText to map from key
// events to Unicode character strings.
//
// By default, GTK uses a platform-dependent default input method. On Windows,
// the default implementation is IME-based and on Wayland, it is using the
// Wayland text protocol. The choice can be overridden programmatically via the
// gtk.Settings:gtk-im-module setting. Users may set the GTK_IM_MODULE
// environment variable to override the default.
//
// Text widgets have a :im-module property (e.g. gtk.TextView:im-module) that
// may also be used to set input methods for specific widget instances. For
// instance, a certain entry widget might be expected to contain certain
// characters which would be easier to input with a specific input method.
//
// An input method may consume multiple key events in sequence before finally
// outputting the composed result. This is called *preediting*, and an input
// method may provide feedback about this process by displaying the intermediate
// composition states as preedit text.
//
// For instance, the built-in GTK input method GtkIMContextSimple implements the
// input of arbitrary Unicode code points by holding down the <kbd>Control</kbd>
// and <kbd>Shift</kbd> keys and then typing <kbd>U</kbd> followed by the
// hexadecimal digits of the code point. When releasing the <kbd>Control</kbd>
// and <kbd>Shift</kbd> keys, preediting ends and the character is inserted as
// text. For example,
//
//    Ctrl+Shift+u 2 0 A C
//
// results in the € sign.
//
// Additional input methods can be made available for use by GTK widgets as
// loadable modules. An input method module is a small shared library which
// provides a GIOExtension for the extension point named "gtk-im-module".
type IMContext struct {
	*externglib.Object
}

// IMContexter describes IMContext's abstract methods.
type IMContexter interface {
	externglib.Objector

	// DeleteSurrounding asks the widget that the input context is attached to
	// delete characters around the cursor position by emitting the
	// GtkIMContext::delete_surrounding signal.
	DeleteSurrounding(offset, nChars int) bool
	// FilterKey: allow an input method to forward key press and release events
	// to another input methodm without necessarily having a GdkEvent available.
	FilterKey(press bool, surface gdk.Surfacer, device gdk.Devicer, time uint32, keycode uint, state gdk.ModifierType, group int) bool
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
	SetSurrounding(text string, len, cursorIndex int)
	// SetSurroundingWithSelection sets surrounding context around the insertion
	// point and preedit string.
	SetSurroundingWithSelection(text string, len, cursorIndex, anchorIndex int)
	// SetUsePreedit sets whether the IM context should use the preedit string
	// to display feedback.
	SetUsePreedit(usePreedit bool)
}

var _ IMContexter = (*IMContext)(nil)

func wrapIMContext(obj *externglib.Object) *IMContext {
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
// Note that offset and n_chars are in characters not in bytes which differs
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
//
// The function takes the following parameters:
//
//    - offset from cursor position in chars; a negative value means start
//    before the cursor.
//    - nChars: number of characters to delete.
//
func (context *IMContext) DeleteSurrounding(offset, nChars int) bool {
	var _arg0 *C.GtkIMContext // out
	var _arg1 C.int           // out
	var _arg2 C.int           // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.int(offset)
	_arg2 = C.int(nChars)

	_cret = C.gtk_im_context_delete_surrounding(_arg0, _arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(offset)
	runtime.KeepAlive(nChars)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FilterKey: allow an input method to forward key press and release events to
// another input methodm without necessarily having a GdkEvent available.
//
// The function takes the following parameters:
//
//    - press: whether to forward a key press or release event.
//    - surface the event is for.
//    - device that the event is for.
//    - time: timestamp for the event.
//    - keycode for the event.
//    - state: modifier state for the event.
//    - group: active keyboard group for the event.
//
func (context *IMContext) FilterKey(press bool, surface gdk.Surfacer, device gdk.Devicer, time uint32, keycode uint, state gdk.ModifierType, group int) bool {
	var _arg0 *C.GtkIMContext   // out
	var _arg1 C.gboolean        // out
	var _arg2 *C.GdkSurface     // out
	var _arg3 *C.GdkDevice      // out
	var _arg4 C.guint32         // out
	var _arg5 C.guint           // out
	var _arg6 C.GdkModifierType // out
	var _arg7 C.int             // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	if press {
		_arg1 = C.TRUE
	}
	_arg2 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	_arg3 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg4 = C.guint32(time)
	_arg5 = C.guint(keycode)
	_arg6 = C.GdkModifierType(state)
	_arg7 = C.int(group)

	_cret = C.gtk_im_context_filter_key(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(context)
	runtime.KeepAlive(press)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(device)
	runtime.KeepAlive(time)
	runtime.KeepAlive(keycode)
	runtime.KeepAlive(state)
	runtime.KeepAlive(group)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FilterKeypress: allow an input method to internally handle key press and
// release events.
//
// If this function returns TRUE, then no further processing should be done for
// this key event.
//
// The function takes the following parameters:
//
//    - event: key event.
//
func (context *IMContext) FilterKeypress(event gdk.Eventer) bool {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.GdkEvent     // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))

	_cret = C.gtk_im_context_filter_keypress(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(event)

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
	runtime.KeepAlive(context)
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
	runtime.KeepAlive(context)
}

// PreeditString: retrieve the current preedit string for the input context, and
// a list of attributes to apply to the string.
//
// This string should be displayed inserted at the insertion point.
func (context *IMContext) PreeditString() (string, *pango.AttrList, int) {
	var _arg0 *C.GtkIMContext  // out
	var _arg1 *C.char          // in
	var _arg2 *C.PangoAttrList // in
	var _arg3 C.int            // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))

	C.gtk_im_context_get_preedit_string(_arg0, &_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(context)

	var _str string            // out
	var _attrs *pango.AttrList // out
	var _cursorPos int         // out

	_str = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	defer C.free(unsafe.Pointer(_arg1))
	_attrs = (*pango.AttrList)(gextras.NewStructNative(unsafe.Pointer(_arg2)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_attrs)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_attr_list_unref((*C.PangoAttrList)(intern.C))
		},
	)
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
// gtk.IMContext::retrieve-surrounding signal on the input method; in response
// to this signal, a widget should provide as much context as is available, up
// to an entire paragraph, by calling gtk.IMContext.SetSurrounding().
//
// Note that there is no obligation for a widget to respond to the
// ::retrieve-surrounding signal, so input methods must be prepared to function
// without context.
//
// Deprecated: Use gtk.IMContext.GetSurroundingWithSelection() instead.
func (context *IMContext) Surrounding() (string, int, bool) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.char         // in
	var _arg2 C.int           // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_im_context_get_surrounding(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(context)

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
// gtk.IMContext::retrieve-surrounding signal on the input method; in response
// to this signal, a widget should provide as much context as is available, up
// to an entire paragraph, by calling
// gtk.IMContext.SetSurroundingWithSelection().
//
// Note that there is no obligation for a widget to respond to the
// ::retrieve-surrounding signal, so input methods must be prepared to function
// without context.
func (context *IMContext) SurroundingWithSelection() (text string, cursorIndex int, anchorIndex int, ok bool) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.char         // in
	var _arg2 C.int           // in
	var _arg3 C.int           // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_im_context_get_surrounding_with_selection(_arg0, &_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(context)

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
	runtime.KeepAlive(context)
}

// SetClientWidget: set the client widget for the input context.
//
// This is the GtkWidget holding the input focus. This widget is used in order
// to correctly position status windows, and may also be used for purposes
// internal to the input method.
//
// The function takes the following parameters:
//
//    - widget: client widget. This may be NULL to indicate that the previous
//    client widget no longer exists.
//
func (context *IMContext) SetClientWidget(widget Widgetter) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.gtk_im_context_set_client_widget(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(widget)
}

// SetCursorLocation: notify the input method that a change in cursor position
// has been made.
//
// The location is relative to the client window.
//
// The function takes the following parameters:
//
//    - area: new location.
//
func (context *IMContext) SetCursorLocation(area *gdk.Rectangle) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.GdkRectangle // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(area)))

	C.gtk_im_context_set_cursor_location(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(area)
}

// SetSurrounding sets surrounding context around the insertion point and
// preedit string.
//
// This function is expected to be called in response to the
// gtk.IMContext::retrieve-surrounding signal, and will likely have no effect if
// called at other times.
//
// Deprecated: Use gtk.IMContext.SetSurroundingWithSelection() instead.
//
// The function takes the following parameters:
//
//    - text surrounding the insertion point, as UTF-8. the preedit string
//    should not be included within text.
//    - len: length of text, or -1 if text is nul-terminated.
//    - cursorIndex: byte index of the insertion cursor within text.
//
func (context *IMContext) SetSurrounding(text string, len, cursorIndex int) {
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
	runtime.KeepAlive(context)
	runtime.KeepAlive(text)
	runtime.KeepAlive(len)
	runtime.KeepAlive(cursorIndex)
}

// SetSurroundingWithSelection sets surrounding context around the insertion
// point and preedit string. This function is expected to be called in response
// to the GtkIMContext::retrieve_surrounding signal, and will likely have no
// effect if called at other times.
//
// The function takes the following parameters:
//
//    - text surrounding the insertion point, as UTF-8. the preedit string
//    should not be included within text.
//    - len: length of text, or -1 if text is nul-terminated.
//    - cursorIndex: byte index of the insertion cursor within text.
//    - anchorIndex: byte index of the selection bound within text.
//
func (context *IMContext) SetSurroundingWithSelection(text string, len, cursorIndex, anchorIndex int) {
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
	runtime.KeepAlive(context)
	runtime.KeepAlive(text)
	runtime.KeepAlive(len)
	runtime.KeepAlive(cursorIndex)
	runtime.KeepAlive(anchorIndex)
}

// SetUsePreedit sets whether the IM context should use the preedit string to
// display feedback.
//
// If use_preedit is FALSE (default is TRUE), then the IM context may use some
// other method to display feedback, such as displaying it in a child of the
// root window.
//
// The function takes the following parameters:
//
//    - usePreedit: whether the IM context should use the preedit string.
//
func (context *IMContext) SetUsePreedit(usePreedit bool) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(context.Native()))
	if usePreedit {
		_arg1 = C.TRUE
	}

	C.gtk_im_context_set_use_preedit(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(usePreedit)
}

// ConnectCommit signal is emitted when a complete input sequence has been
// entered by the user. This can be a single character immediately after a key
// press or the final result of preediting.
func (context *IMContext) ConnectCommit(f func(str string)) externglib.SignalHandle {
	return context.Connect("commit", f)
}

// ConnectDeleteSurrounding signal is emitted when the input method needs to
// delete all or part of the context surrounding the cursor.
func (context *IMContext) ConnectDeleteSurrounding(f func(offset, nChars int) bool) externglib.SignalHandle {
	return context.Connect("delete-surrounding", f)
}

// ConnectPreeditChanged signal is emitted whenever the preedit sequence
// currently being entered has changed. It is also emitted at the end of a
// preedit sequence, in which case gtk_im_context_get_preedit_string() returns
// the empty string.
func (context *IMContext) ConnectPreeditChanged(f func()) externglib.SignalHandle {
	return context.Connect("preedit-changed", f)
}

// ConnectPreeditEnd signal is emitted when a preediting sequence has been
// completed or canceled.
func (context *IMContext) ConnectPreeditEnd(f func()) externglib.SignalHandle {
	return context.Connect("preedit-end", f)
}

// ConnectPreeditStart signal is emitted when a new preediting sequence starts.
func (context *IMContext) ConnectPreeditStart(f func()) externglib.SignalHandle {
	return context.Connect("preedit-start", f)
}

// ConnectRetrieveSurrounding signal is emitted when the input method requires
// the context surrounding the cursor. The callback should set the input method
// surrounding context by calling the gtk_im_context_set_surrounding() method.
func (context *IMContext) ConnectRetrieveSurrounding(f func() bool) externglib.SignalHandle {
	return context.Connect("retrieve-surrounding", f)
}
