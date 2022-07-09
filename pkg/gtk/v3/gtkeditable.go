// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gtk3_EditableInterface_get_selection_bounds(void*, void*, void*);
// extern gint _gotk4_gtk3_EditableInterface_get_position(void*);
// extern void _gotk4_gtk3_EditableInterface_changed(void*);
// extern void _gotk4_gtk3_EditableInterface_delete_text(void*, gint, gint);
// extern void _gotk4_gtk3_EditableInterface_set_position(void*, gint);
// extern void _gotk4_gtk3_Editable_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gtk3_Editable_ConnectDeleteText(gpointer, gint, gint, guintptr);
// extern void* _gotk4_gtk3_EditableInterface_get_chars(void*, gint, gint);
import "C"

// GTypeEditable returns the GType for the type Editable.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeEditable() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Editable").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalEditable)
	return gtype
}

// Editable interface is an interface which should be implemented by text
// editing widgets, such as Entry and SpinButton. It contains functions for
// generically manipulating an editable widget, a large number of action signals
// used for key bindings, and several signals that an application can connect to
// to modify the behavior of a widget.
//
// As an example of the latter usage, by connecting the following handler to
// Editable::insert-text, an application can convert all entry into a widget
// into uppercase.
//
// Forcing entry to uppercase.
//
//    #include <ctype.h>;
//
//    void
//    insert_text_handler (GtkEditable *editable,
//                         const gchar *text,
//                         gint         length,
//                         gint        *position,
//                         gpointer     data)
//    {
//      gchar *result = g_utf8_strup (text, length);
//
//      g_signal_handlers_block_by_func (editable,
//                                   (gpointer) insert_text_handler, data);
//      gtk_editable_insert_text (editable, result, length, position);
//      g_signal_handlers_unblock_by_func (editable,
//                                         (gpointer) insert_text_handler, data);
//
//      g_signal_stop_emission_by_name (editable, "insert_text");
//
//      g_free (result);
//    }.
//
// Editable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Editable struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Editable)(nil)
)

// Editabler describes Editable's interface methods.
type Editabler interface {
	coreglib.Objector

	// CopyClipboard copies the contents of the currently selected content in
	// the editable and puts it on the clipboard.
	CopyClipboard()
	// CutClipboard removes the contents of the currently selected content in
	// the editable and puts it on the clipboard.
	CutClipboard()
	// DeleteSelection deletes the currently selected text of the editable.
	DeleteSelection()
	// DeleteText deletes a sequence of characters.
	DeleteText(startPos, endPos int32)
	// Chars retrieves a sequence of characters.
	Chars(startPos, endPos int32) string
	// Editable retrieves whether editable is editable.
	Editable() bool
	// Position retrieves the current position of the cursor relative to the
	// start of the content of the editable.
	Position() int32
	// SelectionBounds retrieves the selection bound of the editable.
	SelectionBounds() (startPos, endPos int32, ok bool)
	// PasteClipboard pastes the content of the clipboard to the current
	// position of the cursor in the editable.
	PasteClipboard()
	// SelectRegion selects a region of text.
	SelectRegion(startPos, endPos int32)
	// SetEditable determines if the user can edit the text in the editable
	// widget or not.
	SetEditable(isEditable bool)
	// SetPosition sets the cursor position in the editable to the given value.
	SetPosition(position int32)

	// Changed signal is emitted at the end of a single user-visible operation
	// on the contents of the Editable.
	ConnectChanged(func()) coreglib.SignalHandle
	// Delete-text: this signal is emitted when text is deleted from the widget
	// by the user.
	ConnectDeleteText(func(startPos, endPos int32)) coreglib.SignalHandle
}

var _ Editabler = (*Editable)(nil)

func wrapEditable(obj *coreglib.Object) *Editable {
	return &Editable{
		Object: obj,
	}
}

func marshalEditable(p uintptr) (interface{}, error) {
	return wrapEditable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Editable_ConnectChanged
func _gotk4_gtk3_Editable_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectChanged signal is emitted at the end of a single user-visible
// operation on the contents of the Editable.
//
// E.g., a paste operation that replaces the contents of the selection will
// cause only one signal emission (even though it is implemented by first
// deleting the selection, then inserting the new content, and may cause
// multiple ::notify::text signals to be emitted).
func (editable *Editable) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(editable, "changed", false, unsafe.Pointer(C._gotk4_gtk3_Editable_ConnectChanged), f)
}

//export _gotk4_gtk3_Editable_ConnectDeleteText
func _gotk4_gtk3_Editable_ConnectDeleteText(arg0 C.gpointer, arg1 C.gint, arg2 C.gint, arg3 C.guintptr) {
	var f func(startPos, endPos int32)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(startPos, endPos int32))
	}

	var _startPos int32 // out
	var _endPos int32   // out

	_startPos = int32(arg1)
	_endPos = int32(arg2)

	f(_startPos, _endPos)
}

// ConnectDeleteText: this signal is emitted when text is deleted from the
// widget by the user. The default handler for this signal will normally be
// responsible for deleting the text, so by connecting to this signal and then
// stopping the signal with g_signal_stop_emission(), it is possible to modify
// the range of deleted text, or prevent it from being deleted entirely. The
// start_pos and end_pos parameters are interpreted as for
// gtk_editable_delete_text().
func (editable *Editable) ConnectDeleteText(f func(startPos, endPos int32)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(editable, "delete-text", false, unsafe.Pointer(C._gotk4_gtk3_Editable_ConnectDeleteText), f)
}

// CopyClipboard copies the contents of the currently selected content in the
// editable and puts it on the clipboard.
func (editable *Editable) CopyClipboard() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(editable).Native()))

	_info := girepository.MustFind("Gtk", "Editable")
	_info.InvokeIfaceMethod("copy_clipboard", _args[:], nil)

	runtime.KeepAlive(editable)
}

// CutClipboard removes the contents of the currently selected content in the
// editable and puts it on the clipboard.
func (editable *Editable) CutClipboard() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(editable).Native()))

	_info := girepository.MustFind("Gtk", "Editable")
	_info.InvokeIfaceMethod("cut_clipboard", _args[:], nil)

	runtime.KeepAlive(editable)
}

// DeleteSelection deletes the currently selected text of the editable. This
// call doesn’t do anything if there is no selected text.
func (editable *Editable) DeleteSelection() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(editable).Native()))

	_info := girepository.MustFind("Gtk", "Editable")
	_info.InvokeIfaceMethod("delete_selection", _args[:], nil)

	runtime.KeepAlive(editable)
}

// DeleteText deletes a sequence of characters. The characters that are deleted
// are those characters at positions from start_pos up to, but not including
// end_pos. If end_pos is negative, then the characters deleted are those from
// start_pos to the end of the text.
//
// Note that the positions are specified in characters, not bytes.
//
// The function takes the following parameters:
//
//    - startPos: start position.
//    - endPos: end position.
//
func (editable *Editable) DeleteText(startPos, endPos int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(editable).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(startPos)
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(endPos)

	_info := girepository.MustFind("Gtk", "Editable")
	_info.InvokeIfaceMethod("delete_text", _args[:], nil)

	runtime.KeepAlive(editable)
	runtime.KeepAlive(startPos)
	runtime.KeepAlive(endPos)
}

// Chars retrieves a sequence of characters. The characters that are retrieved
// are those characters at positions from start_pos up to, but not including
// end_pos. If end_pos is negative, then the characters retrieved are those
// characters from start_pos to the end of the text.
//
// Note that positions are specified in characters, not bytes.
//
// The function takes the following parameters:
//
//    - startPos: start of text.
//    - endPos: end of text.
//
// The function returns the following values:
//
//    - utf8: pointer to the contents of the widget as a string. This string is
//      allocated by the Editable implementation and should be freed by the
//      caller.
//
func (editable *Editable) Chars(startPos, endPos int32) string {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(editable).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(startPos)
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(endPos)

	_info := girepository.MustFind("Gtk", "Editable")
	_gret := _info.InvokeIfaceMethod("get_chars", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(editable)
	runtime.KeepAlive(startPos)
	runtime.KeepAlive(endPos)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Editable retrieves whether editable is editable. See
// gtk_editable_set_editable().
//
// The function returns the following values:
//
//    - ok: TRUE if editable is editable.
//
func (editable *Editable) Editable() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(editable).Native()))

	_info := girepository.MustFind("Gtk", "Editable")
	_gret := _info.InvokeIfaceMethod("get_editable", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(editable)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Position retrieves the current position of the cursor relative to the start
// of the content of the editable.
//
// Note that this position is in characters, not in bytes.
//
// The function returns the following values:
//
//    - gint: cursor position.
//
func (editable *Editable) Position() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(editable).Native()))

	_info := girepository.MustFind("Gtk", "Editable")
	_gret := _info.InvokeIfaceMethod("get_position", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(editable)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// SelectionBounds retrieves the selection bound of the editable. start_pos will
// be filled with the start of the selection and end_pos with end. If no text
// was selected both will be identical and FALSE will be returned.
//
// Note that positions are specified in characters, not bytes.
//
// The function returns the following values:
//
//    - startPos (optional): location to store the starting position, or NULL.
//    - endPos (optional): location to store the end position, or NULL.
//    - ok: TRUE if an area is selected, FALSE otherwise.
//
func (editable *Editable) SelectionBounds() (startPos, endPos int32, ok bool) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(editable).Native()))

	_info := girepository.MustFind("Gtk", "Editable")
	_gret := _info.InvokeIfaceMethod("get_selection_bounds", _args[:], _outs[:])
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(editable)

	var _startPos int32 // out
	var _endPos int32   // out
	var _ok bool        // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_startPos = *(*int32)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_endPos = *(*int32)(unsafe.Pointer(_outs[1]))
	}
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _startPos, _endPos, _ok
}

// PasteClipboard pastes the content of the clipboard to the current position of
// the cursor in the editable.
func (editable *Editable) PasteClipboard() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(editable).Native()))

	_info := girepository.MustFind("Gtk", "Editable")
	_info.InvokeIfaceMethod("paste_clipboard", _args[:], nil)

	runtime.KeepAlive(editable)
}

// SelectRegion selects a region of text. The characters that are selected are
// those characters at positions from start_pos up to, but not including
// end_pos. If end_pos is negative, then the characters selected are those
// characters from start_pos to the end of the text.
//
// Note that positions are specified in characters, not bytes.
//
// The function takes the following parameters:
//
//    - startPos: start of region.
//    - endPos: end of region.
//
func (editable *Editable) SelectRegion(startPos, endPos int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(editable).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(startPos)
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(endPos)

	_info := girepository.MustFind("Gtk", "Editable")
	_info.InvokeIfaceMethod("select_region", _args[:], nil)

	runtime.KeepAlive(editable)
	runtime.KeepAlive(startPos)
	runtime.KeepAlive(endPos)
}

// SetEditable determines if the user can edit the text in the editable widget
// or not.
//
// The function takes the following parameters:
//
//    - isEditable: TRUE if the user is allowed to edit the text in the widget.
//
func (editable *Editable) SetEditable(isEditable bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(editable).Native()))
	if isEditable {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Editable")
	_info.InvokeIfaceMethod("set_editable", _args[:], nil)

	runtime.KeepAlive(editable)
	runtime.KeepAlive(isEditable)
}

// SetPosition sets the cursor position in the editable to the given value.
//
// The cursor is displayed before the character with the given (base 0) index in
// the contents of the editable. The value must be less than or equal to the
// number of characters in the editable. A value of -1 indicates that the
// position should be set after the last character of the editable. Note that
// position is in characters, not in bytes.
//
// The function takes the following parameters:
//
//    - position of the cursor.
//
func (editable *Editable) SetPosition(position int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(editable).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(position)

	_info := girepository.MustFind("Gtk", "Editable")
	_info.InvokeIfaceMethod("set_position", _args[:], nil)

	runtime.KeepAlive(editable)
	runtime.KeepAlive(position)
}
