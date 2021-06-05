// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_entry_buffer_get_type()), F: marshalEntryBuffer},
	})
}

// EntryBuffer: the EntryBuffer class contains the actual text displayed in a
// Entry widget.
//
// A single EntryBuffer object can be shared by multiple Entry widgets which
// will then share the same text content, but not the cursor position,
// visibility attributes, icon etc.
//
// EntryBuffer may be derived from. Such a derived class might allow text to be
// stored in an alternate location, such as non-pageable memory, useful in the
// case of important passwords. Or a derived class could integrate with an
// application’s concept of undo/redo.
type EntryBuffer interface {
	gextras.Objector

	// DeleteText deletes a sequence of characters from the buffer. @n_chars
	// characters are deleted starting at @position. If @n_chars is negative,
	// then all characters until the end of the text are deleted.
	//
	// If @position or @n_chars are out of bounds, then they are coerced to sane
	// values.
	//
	// Note that the positions are specified in characters, not bytes.
	DeleteText(position uint, nChars int) uint
	// EmitDeletedText: used when subclassing EntryBuffer
	EmitDeletedText(position uint, nChars uint)
	// EmitInsertedText: used when subclassing EntryBuffer
	EmitInsertedText(position uint, chars string, nChars uint)
	// Bytes retrieves the length in bytes of the buffer. See
	// gtk_entry_buffer_get_length().
	Bytes() uint
	// Length retrieves the length in characters of the buffer.
	Length() uint
	// MaxLength retrieves the maximum allowed length of the text in @buffer.
	// See gtk_entry_buffer_set_max_length().
	MaxLength() int
	// Text retrieves the contents of the buffer.
	//
	// The memory pointer returned by this call will not change unless this
	// object emits a signal, or is finalized.
	Text() string
	// InsertText inserts @n_chars characters of @chars into the contents of the
	// buffer, at position @position.
	//
	// If @n_chars is negative, then characters from chars will be inserted
	// until a null-terminator is found. If @position or @n_chars are out of
	// bounds, or the maximum buffer text length is exceeded, then they are
	// coerced to sane values.
	//
	// Note that the position and length are in characters, not in bytes.
	InsertText(position uint, chars string, nChars int) uint
	// SetMaxLength sets the maximum allowed length of the contents of the
	// buffer. If the current contents are longer than the given length, then
	// they will be truncated to fit.
	SetMaxLength(maxLength int)
	// SetText sets the text in the buffer.
	//
	// This is roughly equivalent to calling gtk_entry_buffer_delete_text() and
	// gtk_entry_buffer_insert_text().
	//
	// Note that @n_chars is in characters, not in bytes.
	SetText(chars string, nChars int)
}

// entryBuffer implements the EntryBuffer interface.
type entryBuffer struct {
	gextras.Objector
}

var _ EntryBuffer = (*entryBuffer)(nil)

// WrapEntryBuffer wraps a GObject to the right type. It is
// primarily used internally.
func WrapEntryBuffer(obj *externglib.Object) EntryBuffer {
	return EntryBuffer{
		Objector: obj,
	}
}

func marshalEntryBuffer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEntryBuffer(obj), nil
}

// NewEntryBuffer constructs a class EntryBuffer.
func NewEntryBuffer(initialChars string, nInitialChars int) EntryBuffer {
	var arg1 *C.gchar
	var arg2 C.gint

	arg1 = (*C.gchar)(C.CString(initialChars))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gint(nInitialChars)

	var cret C.GtkEntryBuffer
	var ret1 EntryBuffer

	cret = C.gtk_entry_buffer_new(initialChars, nInitialChars)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(EntryBuffer)

	return ret1
}

// DeleteText deletes a sequence of characters from the buffer. @n_chars
// characters are deleted starting at @position. If @n_chars is negative,
// then all characters until the end of the text are deleted.
//
// If @position or @n_chars are out of bounds, then they are coerced to sane
// values.
//
// Note that the positions are specified in characters, not bytes.
func (b entryBuffer) DeleteText(position uint, nChars int) uint {
	var arg0 *C.GtkEntryBuffer
	var arg1 C.guint
	var arg2 C.gint

	arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(b.Native()))
	arg1 = C.guint(position)
	arg2 = C.gint(nChars)

	var cret C.guint
	var ret1 uint

	cret = C.gtk_entry_buffer_delete_text(arg0, position, nChars)

	ret1 = C.guint(cret)

	return ret1
}

// EmitDeletedText: used when subclassing EntryBuffer
func (b entryBuffer) EmitDeletedText(position uint, nChars uint) {
	var arg0 *C.GtkEntryBuffer
	var arg1 C.guint
	var arg2 C.guint

	arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(b.Native()))
	arg1 = C.guint(position)
	arg2 = C.guint(nChars)

	C.gtk_entry_buffer_emit_deleted_text(arg0, position, nChars)
}

// EmitInsertedText: used when subclassing EntryBuffer
func (b entryBuffer) EmitInsertedText(position uint, chars string, nChars uint) {
	var arg0 *C.GtkEntryBuffer
	var arg1 C.guint
	var arg2 *C.gchar
	var arg3 C.guint

	arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(b.Native()))
	arg1 = C.guint(position)
	arg2 = (*C.gchar)(C.CString(chars))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.guint(nChars)

	C.gtk_entry_buffer_emit_inserted_text(arg0, position, chars, nChars)
}

// Bytes retrieves the length in bytes of the buffer. See
// gtk_entry_buffer_get_length().
func (b entryBuffer) Bytes() uint {
	var arg0 *C.GtkEntryBuffer

	arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(b.Native()))

	var cret C.gsize
	var ret1 uint

	cret = C.gtk_entry_buffer_get_bytes(arg0)

	ret1 = C.gsize(cret)

	return ret1
}

// Length retrieves the length in characters of the buffer.
func (b entryBuffer) Length() uint {
	var arg0 *C.GtkEntryBuffer

	arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(b.Native()))

	var cret C.guint
	var ret1 uint

	cret = C.gtk_entry_buffer_get_length(arg0)

	ret1 = C.guint(cret)

	return ret1
}

// MaxLength retrieves the maximum allowed length of the text in @buffer.
// See gtk_entry_buffer_set_max_length().
func (b entryBuffer) MaxLength() int {
	var arg0 *C.GtkEntryBuffer

	arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(b.Native()))

	var cret C.gint
	var ret1 int

	cret = C.gtk_entry_buffer_get_max_length(arg0)

	ret1 = C.gint(cret)

	return ret1
}

// Text retrieves the contents of the buffer.
//
// The memory pointer returned by this call will not change unless this
// object emits a signal, or is finalized.
func (b entryBuffer) Text() string {
	var arg0 *C.GtkEntryBuffer

	arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(b.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_entry_buffer_get_text(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// InsertText inserts @n_chars characters of @chars into the contents of the
// buffer, at position @position.
//
// If @n_chars is negative, then characters from chars will be inserted
// until a null-terminator is found. If @position or @n_chars are out of
// bounds, or the maximum buffer text length is exceeded, then they are
// coerced to sane values.
//
// Note that the position and length are in characters, not in bytes.
func (b entryBuffer) InsertText(position uint, chars string, nChars int) uint {
	var arg0 *C.GtkEntryBuffer
	var arg1 C.guint
	var arg2 *C.gchar
	var arg3 C.gint

	arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(b.Native()))
	arg1 = C.guint(position)
	arg2 = (*C.gchar)(C.CString(chars))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.gint(nChars)

	var cret C.guint
	var ret1 uint

	cret = C.gtk_entry_buffer_insert_text(arg0, position, chars, nChars)

	ret1 = C.guint(cret)

	return ret1
}

// SetMaxLength sets the maximum allowed length of the contents of the
// buffer. If the current contents are longer than the given length, then
// they will be truncated to fit.
func (b entryBuffer) SetMaxLength(maxLength int) {
	var arg0 *C.GtkEntryBuffer
	var arg1 C.gint

	arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(b.Native()))
	arg1 = C.gint(maxLength)

	C.gtk_entry_buffer_set_max_length(arg0, maxLength)
}

// SetText sets the text in the buffer.
//
// This is roughly equivalent to calling gtk_entry_buffer_delete_text() and
// gtk_entry_buffer_insert_text().
//
// Note that @n_chars is in characters, not in bytes.
func (b entryBuffer) SetText(chars string, nChars int) {
	var arg0 *C.GtkEntryBuffer
	var arg1 *C.gchar
	var arg2 C.gint

	arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(chars))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gint(nChars)

	C.gtk_entry_buffer_set_text(arg0, chars, nChars)
}

type EntryBufferPrivate struct {
	native C.GtkEntryBufferPrivate
}

// WrapEntryBufferPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEntryBufferPrivate(ptr unsafe.Pointer) *EntryBufferPrivate {
	if ptr == nil {
		return nil
	}

	return (*EntryBufferPrivate)(ptr)
}

func marshalEntryBufferPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapEntryBufferPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (e *EntryBufferPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}
