// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
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
		{T: externglib.Type(C.gtk_entry_buffer_get_type()), F: marshalEntryBufferer},
	})
}

// EntryBufferOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type EntryBufferOverrider interface {
	// DeleteText deletes a sequence of characters from the buffer.
	//
	// n_chars characters are deleted starting at position. If n_chars is
	// negative, then all characters until the end of the text are deleted.
	//
	// If position or n_chars are out of bounds, then they are coerced to sane
	// values.
	//
	// Note that the positions are specified in characters, not bytes.
	DeleteText(position uint, nChars uint) uint
	DeletedText(position uint, nChars uint)
	// Length retrieves the length in characters of the buffer.
	Length() uint
	Text(nBytes *uint) string
	// InsertText inserts n_chars characters of chars into the contents of the
	// buffer, at position position.
	//
	// If n_chars is negative, then characters from chars will be inserted until
	// a null-terminator is found. If position or n_chars are out of bounds, or
	// the maximum buffer text length is exceeded, then they are coerced to sane
	// values.
	//
	// Note that the position and length are in characters, not in bytes.
	InsertText(position uint, chars string, nChars uint) uint
	InsertedText(position uint, chars string, nChars uint)
}

// EntryBuffer: GtkEntryBuffer hold the text displayed in a GtkText widget.
//
// A single GtkEntryBuffer object can be shared by multiple widgets which will
// then share the same text content, but not the cursor position, visibility
// attributes, icon etc.
//
// GtkEntryBuffer may be derived from. Such a derived class might allow text to
// be stored in an alternate location, such as non-pageable memory, useful in
// the case of important passwords. Or a derived class could integrate with an
// application’s concept of undo/redo.
type EntryBuffer struct {
	*externglib.Object
}

func wrapEntryBuffer(obj *externglib.Object) *EntryBuffer {
	return &EntryBuffer{
		Object: obj,
	}
}

func marshalEntryBufferer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEntryBuffer(obj), nil
}

// NewEntryBuffer: create a new GtkEntryBuffer object.
//
// Optionally, specify initial text to set in the buffer.
func NewEntryBuffer(initialChars string, nInitialChars int) *EntryBuffer {
	var _arg1 *C.char           // out
	var _arg2 C.int             // out
	var _cret *C.GtkEntryBuffer // in

	if initialChars != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(initialChars)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = C.int(nInitialChars)

	_cret = C.gtk_entry_buffer_new(_arg1, _arg2)
	runtime.KeepAlive(initialChars)
	runtime.KeepAlive(nInitialChars)

	var _entryBuffer *EntryBuffer // out

	_entryBuffer = wrapEntryBuffer(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _entryBuffer
}

// DeleteText deletes a sequence of characters from the buffer.
//
// n_chars characters are deleted starting at position. If n_chars is negative,
// then all characters until the end of the text are deleted.
//
// If position or n_chars are out of bounds, then they are coerced to sane
// values.
//
// Note that the positions are specified in characters, not bytes.
func (buffer *EntryBuffer) DeleteText(position uint, nChars int) uint {
	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 C.guint           // out
	var _arg2 C.int             // out
	var _cret C.guint           // in

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))
	_arg1 = C.guint(position)
	_arg2 = C.int(nChars)

	_cret = C.gtk_entry_buffer_delete_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(position)
	runtime.KeepAlive(nChars)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// EmitDeletedText: used when subclassing GtkEntryBuffer.
func (buffer *EntryBuffer) EmitDeletedText(position uint, nChars uint) {
	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 C.guint           // out
	var _arg2 C.guint           // out

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))
	_arg1 = C.guint(position)
	_arg2 = C.guint(nChars)

	C.gtk_entry_buffer_emit_deleted_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(position)
	runtime.KeepAlive(nChars)
}

// EmitInsertedText: used when subclassing GtkEntryBuffer.
func (buffer *EntryBuffer) EmitInsertedText(position uint, chars string, nChars uint) {
	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 C.guint           // out
	var _arg2 *C.char           // out
	var _arg3 C.guint           // out

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))
	_arg1 = C.guint(position)
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(chars)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.guint(nChars)

	C.gtk_entry_buffer_emit_inserted_text(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(position)
	runtime.KeepAlive(chars)
	runtime.KeepAlive(nChars)
}

// Bytes retrieves the length in bytes of the buffer.
//
// See gtk.EntryBuffer.GetLength().
func (buffer *EntryBuffer) Bytes() uint {
	var _arg0 *C.GtkEntryBuffer // out
	var _cret C.gsize           // in

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))

	_cret = C.gtk_entry_buffer_get_bytes(_arg0)
	runtime.KeepAlive(buffer)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// Length retrieves the length in characters of the buffer.
func (buffer *EntryBuffer) Length() uint {
	var _arg0 *C.GtkEntryBuffer // out
	var _cret C.guint           // in

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))

	_cret = C.gtk_entry_buffer_get_length(_arg0)
	runtime.KeepAlive(buffer)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// MaxLength retrieves the maximum allowed length of the text in buffer.
func (buffer *EntryBuffer) MaxLength() int {
	var _arg0 *C.GtkEntryBuffer // out
	var _cret C.int             // in

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))

	_cret = C.gtk_entry_buffer_get_max_length(_arg0)
	runtime.KeepAlive(buffer)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Text retrieves the contents of the buffer.
//
// The memory pointer returned by this call will not change unless this object
// emits a signal, or is finalized.
func (buffer *EntryBuffer) Text() string {
	var _arg0 *C.GtkEntryBuffer // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))

	_cret = C.gtk_entry_buffer_get_text(_arg0)
	runtime.KeepAlive(buffer)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// InsertText inserts n_chars characters of chars into the contents of the
// buffer, at position position.
//
// If n_chars is negative, then characters from chars will be inserted until a
// null-terminator is found. If position or n_chars are out of bounds, or the
// maximum buffer text length is exceeded, then they are coerced to sane values.
//
// Note that the position and length are in characters, not in bytes.
func (buffer *EntryBuffer) InsertText(position uint, chars string, nChars int) uint {
	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 C.guint           // out
	var _arg2 *C.char           // out
	var _arg3 C.int             // out
	var _cret C.guint           // in

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))
	_arg1 = C.guint(position)
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(chars)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.int(nChars)

	_cret = C.gtk_entry_buffer_insert_text(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(position)
	runtime.KeepAlive(chars)
	runtime.KeepAlive(nChars)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SetMaxLength sets the maximum allowed length of the contents of the buffer.
//
// If the current contents are longer than the given length, then they will be
// truncated to fit.
func (buffer *EntryBuffer) SetMaxLength(maxLength int) {
	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 C.int             // out

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))
	_arg1 = C.int(maxLength)

	C.gtk_entry_buffer_set_max_length(_arg0, _arg1)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(maxLength)
}

// SetText sets the text in the buffer.
//
// This is roughly equivalent to calling gtk.EntryBuffer.DeleteText() and
// gtk.EntryBuffer.InsertText().
//
// Note that n_chars is in characters, not in bytes.
func (buffer *EntryBuffer) SetText(chars string, nChars int) {
	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 *C.char           // out
	var _arg2 C.int             // out

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(chars)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(nChars)

	C.gtk_entry_buffer_set_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(chars)
	runtime.KeepAlive(nChars)
}

// ConnectDeletedText: text is altered in the default handler for this signal.
//
// If you want access to the text after the text has been modified, use
// G_CONNECT_AFTER.
func (e *EntryBuffer) ConnectDeletedText(f func(position, nChars uint)) glib.SignalHandle {
	return e.Connect("deleted-text", f)
}

// ConnectInsertedText: this signal is emitted after text is inserted into the
// buffer.
func (e *EntryBuffer) ConnectInsertedText(f func(position uint, chars string, nChars uint)) glib.SignalHandle {
	return e.Connect("inserted-text", f)
}
