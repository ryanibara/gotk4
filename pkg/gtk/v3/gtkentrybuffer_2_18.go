// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_EntryBuffer_ConnectInsertedText(gpointer, guint, gchar*, guint, guintptr);
// extern void _gotk4_gtk3_EntryBuffer_ConnectDeletedText(gpointer, guint, guint, guintptr);
// extern void _gotk4_gtk3_EntryBufferClass_inserted_text(GtkEntryBuffer*, guint, gchar*, guint);
// extern void _gotk4_gtk3_EntryBufferClass_deleted_text(GtkEntryBuffer*, guint, guint);
// extern guint _gotk4_gtk3_EntryBufferClass_insert_text(GtkEntryBuffer*, guint, gchar*, guint);
// extern guint _gotk4_gtk3_EntryBufferClass_get_length(GtkEntryBuffer*);
// extern guint _gotk4_gtk3_EntryBufferClass_delete_text(GtkEntryBuffer*, guint, guint);
// extern gchar* _gotk4_gtk3_EntryBufferClass_get_text(GtkEntryBuffer*, gsize*);
// gchar* _gotk4_gtk3_EntryBuffer_virtual_get_text(void* fnptr, GtkEntryBuffer* arg0, gsize* arg1) {
//   return ((gchar* (*)(GtkEntryBuffer*, gsize*))(fnptr))(arg0, arg1);
// };
// guint _gotk4_gtk3_EntryBuffer_virtual_delete_text(void* fnptr, GtkEntryBuffer* arg0, guint arg1, guint arg2) {
//   return ((guint (*)(GtkEntryBuffer*, guint, guint))(fnptr))(arg0, arg1, arg2);
// };
// guint _gotk4_gtk3_EntryBuffer_virtual_get_length(void* fnptr, GtkEntryBuffer* arg0) {
//   return ((guint (*)(GtkEntryBuffer*))(fnptr))(arg0);
// };
// guint _gotk4_gtk3_EntryBuffer_virtual_insert_text(void* fnptr, GtkEntryBuffer* arg0, guint arg1, gchar* arg2, guint arg3) {
//   return ((guint (*)(GtkEntryBuffer*, guint, gchar*, guint))(fnptr))(arg0, arg1, arg2, arg3);
// };
// void _gotk4_gtk3_EntryBuffer_virtual_deleted_text(void* fnptr, GtkEntryBuffer* arg0, guint arg1, guint arg2) {
//   ((void (*)(GtkEntryBuffer*, guint, guint))(fnptr))(arg0, arg1, arg2);
// };
// void _gotk4_gtk3_EntryBuffer_virtual_inserted_text(void* fnptr, GtkEntryBuffer* arg0, guint arg1, gchar* arg2, guint arg3) {
//   ((void (*)(GtkEntryBuffer*, guint, gchar*, guint))(fnptr))(arg0, arg1, arg2, arg3);
// };
import "C"

// GType values.
var (
	GTypeEntryBuffer = coreglib.Type(C.gtk_entry_buffer_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEntryBuffer, F: marshalEntryBuffer},
	})
}

// EntryBufferOverrides contains methods that are overridable.
type EntryBufferOverrides struct {
	// DeleteText deletes a sequence of characters from the buffer. n_chars
	// characters are deleted starting at position. If n_chars is negative,
	// then all characters until the end of the text are deleted.
	//
	// If position or n_chars are out of bounds, then they are coerced to sane
	// values.
	//
	// Note that the positions are specified in characters, not bytes.
	//
	// The function takes the following parameters:
	//
	//   - position at which to delete text.
	//   - nChars: number of characters to delete.
	//
	// The function returns the following values:
	//
	//   - guint: number of characters deleted.
	//
	DeleteText func(position, nChars uint) uint
	// The function takes the following parameters:
	//
	//   - position
	//   - nChars
	//
	DeletedText func(position, nChars uint)
	// Length retrieves the length in characters of the buffer.
	//
	// The function returns the following values:
	//
	//   - guint: number of characters in the buffer.
	//
	Length func() uint
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	Text func(nBytes *uint) string
	// InsertText inserts n_chars characters of chars into the contents of the
	// buffer, at position position.
	//
	// If n_chars is negative, then characters from chars will be inserted until
	// a null-terminator is found. If position or n_chars are out of bounds,
	// or the maximum buffer text length is exceeded, then they are coerced to
	// sane values.
	//
	// Note that the position and length are in characters, not in bytes.
	//
	// The function takes the following parameters:
	//
	//   - position at which to insert text.
	//   - chars: text to insert into the buffer.
	//   - nChars: length of the text in characters, or -1.
	//
	// The function returns the following values:
	//
	//   - guint: number of characters actually inserted.
	//
	InsertText func(position uint, chars string, nChars uint) uint
	// The function takes the following parameters:
	//
	//   - position
	//   - chars
	//   - nChars
	//
	InsertedText func(position uint, chars string, nChars uint)
}

func defaultEntryBufferOverrides(v *EntryBuffer) EntryBufferOverrides {
	return EntryBufferOverrides{
		DeleteText:   v.deleteText,
		DeletedText:  v.deletedText,
		Length:       v.length,
		Text:         v.text,
		InsertText:   v.insertText,
		InsertedText: v.insertedText,
	}
}

// EntryBuffer class contains the actual text displayed in a Entry widget.
//
// A single EntryBuffer object can be shared by multiple Entry widgets
// which will then share the same text content, but not the cursor position,
// visibility attributes, icon etc.
//
// EntryBuffer may be derived from. Such a derived class might allow text to
// be stored in an alternate location, such as non-pageable memory, useful in
// the case of important passwords. Or a derived class could integrate with an
// application’s concept of undo/redo.
type EntryBuffer struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*EntryBuffer)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*EntryBuffer, *EntryBufferClass, EntryBufferOverrides](
		GTypeEntryBuffer,
		initEntryBufferClass,
		wrapEntryBuffer,
		defaultEntryBufferOverrides,
	)
}

func initEntryBufferClass(gclass unsafe.Pointer, overrides EntryBufferOverrides, classInitFunc func(*EntryBufferClass)) {
	pclass := (*C.GtkEntryBufferClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeEntryBuffer))))

	if overrides.DeleteText != nil {
		pclass.delete_text = (*[0]byte)(C._gotk4_gtk3_EntryBufferClass_delete_text)
	}

	if overrides.DeletedText != nil {
		pclass.deleted_text = (*[0]byte)(C._gotk4_gtk3_EntryBufferClass_deleted_text)
	}

	if overrides.Length != nil {
		pclass.get_length = (*[0]byte)(C._gotk4_gtk3_EntryBufferClass_get_length)
	}

	if overrides.Text != nil {
		pclass.get_text = (*[0]byte)(C._gotk4_gtk3_EntryBufferClass_get_text)
	}

	if overrides.InsertText != nil {
		pclass.insert_text = (*[0]byte)(C._gotk4_gtk3_EntryBufferClass_insert_text)
	}

	if overrides.InsertedText != nil {
		pclass.inserted_text = (*[0]byte)(C._gotk4_gtk3_EntryBufferClass_inserted_text)
	}

	if classInitFunc != nil {
		class := (*EntryBufferClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapEntryBuffer(obj *coreglib.Object) *EntryBuffer {
	return &EntryBuffer{
		Object: obj,
	}
}

func marshalEntryBuffer(p uintptr) (interface{}, error) {
	return wrapEntryBuffer(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectDeletedText: this signal is emitted after text is deleted from the
// buffer.
func (buffer *EntryBuffer) ConnectDeletedText(f func(position, nChars uint)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(buffer, "deleted-text", false, unsafe.Pointer(C._gotk4_gtk3_EntryBuffer_ConnectDeletedText), f)
}

// ConnectInsertedText: this signal is emitted after text is inserted into the
// buffer.
func (buffer *EntryBuffer) ConnectInsertedText(f func(position uint, chars string, nChars uint)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(buffer, "inserted-text", false, unsafe.Pointer(C._gotk4_gtk3_EntryBuffer_ConnectInsertedText), f)
}

// NewEntryBuffer: create a new GtkEntryBuffer object.
//
// Optionally, specify initial text to set in the buffer.
//
// The function takes the following parameters:
//
//   - initialChars (optional): initial buffer text, or NULL.
//   - nInitialChars: number of characters in initial_chars, or -1.
//
// The function returns the following values:
//
//   - entryBuffer: new GtkEntryBuffer object.
//
func NewEntryBuffer(initialChars string, nInitialChars int) *EntryBuffer {
	var _arg1 *C.gchar          // out
	var _arg2 C.gint            // out
	var _cret *C.GtkEntryBuffer // in

	if initialChars != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(initialChars)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = C.gint(nInitialChars)

	_cret = C.gtk_entry_buffer_new(_arg1, _arg2)
	runtime.KeepAlive(initialChars)
	runtime.KeepAlive(nInitialChars)

	var _entryBuffer *EntryBuffer // out

	_entryBuffer = wrapEntryBuffer(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _entryBuffer
}

// DeleteText deletes a sequence of characters from the buffer. n_chars
// characters are deleted starting at position. If n_chars is negative, then all
// characters until the end of the text are deleted.
//
// If position or n_chars are out of bounds, then they are coerced to sane
// values.
//
// Note that the positions are specified in characters, not bytes.
//
// The function takes the following parameters:
//
//   - position at which to delete text.
//   - nChars: number of characters to delete.
//
// The function returns the following values:
//
//   - guint: number of characters deleted.
//
func (buffer *EntryBuffer) DeleteText(position uint, nChars int) uint {
	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 C.guint           // out
	var _arg2 C.gint            // out
	var _cret C.guint           // in

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = C.guint(position)
	_arg2 = C.gint(nChars)

	_cret = C.gtk_entry_buffer_delete_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(position)
	runtime.KeepAlive(nChars)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// EmitDeletedText: used when subclassing EntryBuffer.
//
// The function takes the following parameters:
//
//   - position at which text was deleted.
//   - nChars: number of characters deleted.
//
func (buffer *EntryBuffer) EmitDeletedText(position, nChars uint) {
	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 C.guint           // out
	var _arg2 C.guint           // out

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = C.guint(position)
	_arg2 = C.guint(nChars)

	C.gtk_entry_buffer_emit_deleted_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(position)
	runtime.KeepAlive(nChars)
}

// EmitInsertedText: used when subclassing EntryBuffer.
//
// The function takes the following parameters:
//
//   - position at which text was inserted.
//   - chars: text that was inserted.
//   - nChars: number of characters inserted.
//
func (buffer *EntryBuffer) EmitInsertedText(position uint, chars string, nChars uint) {
	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 C.guint           // out
	var _arg2 *C.gchar          // out
	var _arg3 C.guint           // out

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = C.guint(position)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(chars)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.guint(nChars)

	C.gtk_entry_buffer_emit_inserted_text(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(position)
	runtime.KeepAlive(chars)
	runtime.KeepAlive(nChars)
}

// Bytes retrieves the length in bytes of the buffer. See
// gtk_entry_buffer_get_length().
//
// The function returns the following values:
//
//   - gsize: byte length of the buffer.
//
func (buffer *EntryBuffer) Bytes() uint {
	var _arg0 *C.GtkEntryBuffer // out
	var _cret C.gsize           // in

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_cret = C.gtk_entry_buffer_get_bytes(_arg0)
	runtime.KeepAlive(buffer)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// Length retrieves the length in characters of the buffer.
//
// The function returns the following values:
//
//   - guint: number of characters in the buffer.
//
func (buffer *EntryBuffer) Length() uint {
	var _arg0 *C.GtkEntryBuffer // out
	var _cret C.guint           // in

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_cret = C.gtk_entry_buffer_get_length(_arg0)
	runtime.KeepAlive(buffer)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// MaxLength retrieves the maximum allowed length of the text in buffer.
// See gtk_entry_buffer_set_max_length().
//
// The function returns the following values:
//
//   - gint: maximum allowed number of characters in EntryBuffer, or 0 if there
//     is no maximum.
//
func (buffer *EntryBuffer) MaxLength() int {
	var _arg0 *C.GtkEntryBuffer // out
	var _cret C.gint            // in

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

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
//
// The function returns the following values:
//
//   - utf8: pointer to the contents of the widget as a string. This string
//     points to internally allocated storage in the buffer and must not be
//     freed, modified or stored.
//
func (buffer *EntryBuffer) Text() string {
	var _arg0 *C.GtkEntryBuffer // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

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
//
// The function takes the following parameters:
//
//   - position at which to insert text.
//   - chars: text to insert into the buffer.
//   - nChars: length of the text in characters, or -1.
//
// The function returns the following values:
//
//   - guint: number of characters actually inserted.
//
func (buffer *EntryBuffer) InsertText(position uint, chars string, nChars int) uint {
	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 C.guint           // out
	var _arg2 *C.gchar          // out
	var _arg3 C.gint            // out
	var _cret C.guint           // in

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = C.guint(position)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(chars)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gint(nChars)

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
// If the current contents are longer than the given length, then they will be
// truncated to fit.
//
// The function takes the following parameters:
//
//   - maxLength: maximum length of the entry buffer, or 0 for no maximum.
//     (other than the maximum length of entries.) The value passed in will be
//     clamped to the range 0-65536.
//
func (buffer *EntryBuffer) SetMaxLength(maxLength int) {
	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 C.gint            // out

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = C.gint(maxLength)

	C.gtk_entry_buffer_set_max_length(_arg0, _arg1)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(maxLength)
}

// SetText sets the text in the buffer.
//
// This is roughly equivalent to calling gtk_entry_buffer_delete_text() and
// gtk_entry_buffer_insert_text().
//
// Note that n_chars is in characters, not in bytes.
//
// The function takes the following parameters:
//
//   - chars: new text.
//   - nChars: number of characters in text, or -1.
//
func (buffer *EntryBuffer) SetText(chars string, nChars int) {
	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 *C.gchar          // out
	var _arg2 C.gint            // out

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(chars)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(nChars)

	C.gtk_entry_buffer_set_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(chars)
	runtime.KeepAlive(nChars)
}

// deleteText deletes a sequence of characters from the buffer. n_chars
// characters are deleted starting at position. If n_chars is negative, then all
// characters until the end of the text are deleted.
//
// If position or n_chars are out of bounds, then they are coerced to sane
// values.
//
// Note that the positions are specified in characters, not bytes.
//
// The function takes the following parameters:
//
//   - position at which to delete text.
//   - nChars: number of characters to delete.
//
// The function returns the following values:
//
//   - guint: number of characters deleted.
//
func (buffer *EntryBuffer) deleteText(position, nChars uint) uint {
	gclass := (*C.GtkEntryBufferClass)(coreglib.PeekParentClass(buffer))
	fnarg := gclass.delete_text

	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 C.guint           // out
	var _arg2 C.guint           // out
	var _cret C.guint           // in

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = C.guint(position)
	_arg2 = C.guint(nChars)

	_cret = C._gotk4_gtk3_EntryBuffer_virtual_delete_text(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(position)
	runtime.KeepAlive(nChars)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// The function takes the following parameters:
//
//   - position
//   - nChars
//
func (buffer *EntryBuffer) deletedText(position, nChars uint) {
	gclass := (*C.GtkEntryBufferClass)(coreglib.PeekParentClass(buffer))
	fnarg := gclass.deleted_text

	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 C.guint           // out
	var _arg2 C.guint           // out

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = C.guint(position)
	_arg2 = C.guint(nChars)

	C._gotk4_gtk3_EntryBuffer_virtual_deleted_text(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(position)
	runtime.KeepAlive(nChars)
}

// Length retrieves the length in characters of the buffer.
//
// The function returns the following values:
//
//   - guint: number of characters in the buffer.
//
func (buffer *EntryBuffer) length() uint {
	gclass := (*C.GtkEntryBufferClass)(coreglib.PeekParentClass(buffer))
	fnarg := gclass.get_length

	var _arg0 *C.GtkEntryBuffer // out
	var _cret C.guint           // in

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_cret = C._gotk4_gtk3_EntryBuffer_virtual_get_length(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(buffer)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (buffer *EntryBuffer) text(nBytes *uint) string {
	gclass := (*C.GtkEntryBufferClass)(coreglib.PeekParentClass(buffer))
	fnarg := gclass.get_text

	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 *C.gsize          // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = (*C.gsize)(unsafe.Pointer(nBytes))

	_cret = C._gotk4_gtk3_EntryBuffer_virtual_get_text(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(nBytes)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// insertText inserts n_chars characters of chars into the contents of the
// buffer, at position position.
//
// If n_chars is negative, then characters from chars will be inserted until a
// null-terminator is found. If position or n_chars are out of bounds, or the
// maximum buffer text length is exceeded, then they are coerced to sane values.
//
// Note that the position and length are in characters, not in bytes.
//
// The function takes the following parameters:
//
//   - position at which to insert text.
//   - chars: text to insert into the buffer.
//   - nChars: length of the text in characters, or -1.
//
// The function returns the following values:
//
//   - guint: number of characters actually inserted.
//
func (buffer *EntryBuffer) insertText(position uint, chars string, nChars uint) uint {
	gclass := (*C.GtkEntryBufferClass)(coreglib.PeekParentClass(buffer))
	fnarg := gclass.insert_text

	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 C.guint           // out
	var _arg2 *C.gchar          // out
	var _arg3 C.guint           // out
	var _cret C.guint           // in

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = C.guint(position)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(chars)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.guint(nChars)

	_cret = C._gotk4_gtk3_EntryBuffer_virtual_insert_text(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(position)
	runtime.KeepAlive(chars)
	runtime.KeepAlive(nChars)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// The function takes the following parameters:
//
//   - position
//   - chars
//   - nChars
//
func (buffer *EntryBuffer) insertedText(position uint, chars string, nChars uint) {
	gclass := (*C.GtkEntryBufferClass)(coreglib.PeekParentClass(buffer))
	fnarg := gclass.inserted_text

	var _arg0 *C.GtkEntryBuffer // out
	var _arg1 C.guint           // out
	var _arg2 *C.gchar          // out
	var _arg3 C.guint           // out

	_arg0 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = C.guint(position)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(chars)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.guint(nChars)

	C._gotk4_gtk3_EntryBuffer_virtual_inserted_text(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(position)
	runtime.KeepAlive(chars)
	runtime.KeepAlive(nChars)
}
