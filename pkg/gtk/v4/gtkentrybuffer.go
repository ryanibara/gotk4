// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern char* _gotk4_gtk4_EntryBufferClass_get_text(void*, void*);
// extern guint _gotk4_gtk4_EntryBufferClass_delete_text(void*, guint, guint);
// extern guint _gotk4_gtk4_EntryBufferClass_get_length(void*);
// extern guint _gotk4_gtk4_EntryBufferClass_insert_text(void*, guint, void*, guint);
// extern void _gotk4_gtk4_EntryBufferClass_deleted_text(void*, guint, guint);
// extern void _gotk4_gtk4_EntryBufferClass_inserted_text(void*, guint, void*, guint);
// extern void _gotk4_gtk4_EntryBuffer_ConnectDeletedText(gpointer, guint, guint, guintptr);
// extern void _gotk4_gtk4_EntryBuffer_ConnectInsertedText(gpointer, guint, void*, guint, guintptr);
import "C"

// glib.Type values for gtkentrybuffer.go.
var GTypeEntryBuffer = coreglib.Type(C.gtk_entry_buffer_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeEntryBuffer, F: marshalEntryBuffer},
	})
}

// EntryBufferOverrider contains methods that are overridable.
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
	//
	// The function takes the following parameters:
	//
	//    - position at which to delete text.
	//    - nChars: number of characters to delete.
	//
	// The function returns the following values:
	//
	//    - guint: number of characters deleted.
	//
	DeleteText(position, nChars uint32) uint32
	// The function takes the following parameters:
	//
	//    - position
	//    - nChars
	//
	DeletedText(position, nChars uint32)
	// Length retrieves the length in characters of the buffer.
	//
	// The function returns the following values:
	//
	//    - guint: number of characters in the buffer.
	//
	Length() uint32
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
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
	//
	// The function takes the following parameters:
	//
	//    - position at which to insert text.
	//    - chars: text to insert into the buffer.
	//    - nChars: length of the text in characters, or -1.
	//
	// The function returns the following values:
	//
	//    - guint: number of characters actually inserted.
	//
	InsertText(position uint32, chars string, nChars uint32) uint32
	// The function takes the following parameters:
	//
	//    - position
	//    - chars
	//    - nChars
	//
	InsertedText(position uint32, chars string, nChars uint32)
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
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*EntryBuffer)(nil)
)

func classInitEntryBufferer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkEntryBufferClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkEntryBufferClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface {
		DeleteText(position, nChars uint32) uint32
	}); ok {
		pclass.delete_text = (*[0]byte)(C._gotk4_gtk4_EntryBufferClass_delete_text)
	}

	if _, ok := goval.(interface{ DeletedText(position, nChars uint32) }); ok {
		pclass.deleted_text = (*[0]byte)(C._gotk4_gtk4_EntryBufferClass_deleted_text)
	}

	if _, ok := goval.(interface{ Length() uint32 }); ok {
		pclass.get_length = (*[0]byte)(C._gotk4_gtk4_EntryBufferClass_get_length)
	}

	if _, ok := goval.(interface{ Text(nBytes *uint) string }); ok {
		pclass.get_text = (*[0]byte)(C._gotk4_gtk4_EntryBufferClass_get_text)
	}

	if _, ok := goval.(interface {
		InsertText(position uint32, chars string, nChars uint32) uint32
	}); ok {
		pclass.insert_text = (*[0]byte)(C._gotk4_gtk4_EntryBufferClass_insert_text)
	}

	if _, ok := goval.(interface {
		InsertedText(position uint32, chars string, nChars uint32)
	}); ok {
		pclass.inserted_text = (*[0]byte)(C._gotk4_gtk4_EntryBufferClass_inserted_text)
	}
}

//export _gotk4_gtk4_EntryBufferClass_delete_text
func _gotk4_gtk4_EntryBufferClass_delete_text(arg0 *C.void, arg1 C.guint, arg2 C.guint) (cret C.guint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		DeleteText(position, nChars uint32) uint32
	})

	var _position uint32 // out
	var _nChars uint32   // out

	_position = uint32(arg1)
	_nChars = uint32(arg2)

	guint := iface.DeleteText(_position, _nChars)

	cret = C.guint(guint)

	return cret
}

//export _gotk4_gtk4_EntryBufferClass_deleted_text
func _gotk4_gtk4_EntryBufferClass_deleted_text(arg0 *C.void, arg1 C.guint, arg2 C.guint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ DeletedText(position, nChars uint32) })

	var _position uint32 // out
	var _nChars uint32   // out

	_position = uint32(arg1)
	_nChars = uint32(arg2)

	iface.DeletedText(_position, _nChars)
}

//export _gotk4_gtk4_EntryBufferClass_get_length
func _gotk4_gtk4_EntryBufferClass_get_length(arg0 *C.void) (cret C.guint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Length() uint32 })

	guint := iface.Length()

	cret = C.guint(guint)

	return cret
}

//export _gotk4_gtk4_EntryBufferClass_get_text
func _gotk4_gtk4_EntryBufferClass_get_text(arg0 *C.void, arg1 *C.void) (cret *C.char) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Text(nBytes *uint) string })

	var _nBytes *uint // out

	_nBytes = (*uint)(unsafe.Pointer(arg1))

	utf8 := iface.Text(_nBytes)

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_gtk4_EntryBufferClass_insert_text
func _gotk4_gtk4_EntryBufferClass_insert_text(arg0 *C.void, arg1 C.guint, arg2 *C.void, arg3 C.guint) (cret C.guint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		InsertText(position uint32, chars string, nChars uint32) uint32
	})

	var _position uint32 // out
	var _chars string    // out
	var _nChars uint32   // out

	_position = uint32(arg1)
	_chars = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	_nChars = uint32(arg3)

	guint := iface.InsertText(_position, _chars, _nChars)

	cret = C.guint(guint)

	return cret
}

//export _gotk4_gtk4_EntryBufferClass_inserted_text
func _gotk4_gtk4_EntryBufferClass_inserted_text(arg0 *C.void, arg1 C.guint, arg2 *C.void, arg3 C.guint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		InsertedText(position uint32, chars string, nChars uint32)
	})

	var _position uint32 // out
	var _chars string    // out
	var _nChars uint32   // out

	_position = uint32(arg1)
	_chars = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	_nChars = uint32(arg3)

	iface.InsertedText(_position, _chars, _nChars)
}

func wrapEntryBuffer(obj *coreglib.Object) *EntryBuffer {
	return &EntryBuffer{
		Object: obj,
	}
}

func marshalEntryBuffer(p uintptr) (interface{}, error) {
	return wrapEntryBuffer(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_EntryBuffer_ConnectDeletedText
func _gotk4_gtk4_EntryBuffer_ConnectDeletedText(arg0 C.gpointer, arg1 C.guint, arg2 C.guint, arg3 C.guintptr) {
	var f func(position, nChars uint32)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(position, nChars uint32))
	}

	var _position uint32 // out
	var _nChars uint32   // out

	_position = uint32(arg1)
	_nChars = uint32(arg2)

	f(_position, _nChars)
}

// ConnectDeletedText: text is altered in the default handler for this signal.
//
// If you want access to the text after the text has been modified, use
// G_CONNECT_AFTER.
func (buffer *EntryBuffer) ConnectDeletedText(f func(position, nChars uint32)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(buffer, "deleted-text", false, unsafe.Pointer(C._gotk4_gtk4_EntryBuffer_ConnectDeletedText), f)
}

//export _gotk4_gtk4_EntryBuffer_ConnectInsertedText
func _gotk4_gtk4_EntryBuffer_ConnectInsertedText(arg0 C.gpointer, arg1 C.guint, arg2 *C.void, arg3 C.guint, arg4 C.guintptr) {
	var f func(position uint32, chars string, nChars uint32)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(position uint32, chars string, nChars uint32))
	}

	var _position uint32 // out
	var _chars string    // out
	var _nChars uint32   // out

	_position = uint32(arg1)
	_chars = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	_nChars = uint32(arg3)

	f(_position, _chars, _nChars)
}

// ConnectInsertedText: this signal is emitted after text is inserted into the
// buffer.
func (buffer *EntryBuffer) ConnectInsertedText(f func(position uint32, chars string, nChars uint32)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(buffer, "inserted-text", false, unsafe.Pointer(C._gotk4_gtk4_EntryBuffer_ConnectInsertedText), f)
}

// NewEntryBuffer: create a new GtkEntryBuffer object.
//
// Optionally, specify initial text to set in the buffer.
//
// The function takes the following parameters:
//
//    - initialChars (optional): initial buffer text, or NULL.
//    - nInitialChars: number of characters in initial_chars, or -1.
//
// The function returns the following values:
//
//    - entryBuffer: new GtkEntryBuffer object.
//
func NewEntryBuffer(initialChars string, nInitialChars int32) *EntryBuffer {
	var _args [2]girepository.Argument

	if initialChars != "" {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(initialChars)))
		defer C.free(unsafe.Pointer(_args[0]))
	}
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(nInitialChars)

	_gret := girepository.MustFind("Gtk", "EntryBuffer").InvokeMethod("new_EntryBuffer", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(initialChars)
	runtime.KeepAlive(nInitialChars)

	var _entryBuffer *EntryBuffer // out

	_entryBuffer = wrapEntryBuffer(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

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
//
// The function takes the following parameters:
//
//    - position at which to delete text.
//    - nChars: number of characters to delete.
//
// The function returns the following values:
//
//    - guint: number of characters deleted.
//
func (buffer *EntryBuffer) DeleteText(position uint32, nChars int32) uint32 {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(position)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(nChars)

	_gret := girepository.MustFind("Gtk", "EntryBuffer").InvokeMethod("delete_text", _args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(buffer)
	runtime.KeepAlive(position)
	runtime.KeepAlive(nChars)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// EmitDeletedText: used when subclassing GtkEntryBuffer.
//
// The function takes the following parameters:
//
//    - position at which text was deleted.
//    - nChars: number of characters deleted.
//
func (buffer *EntryBuffer) EmitDeletedText(position, nChars uint32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(position)
	*(*C.guint)(unsafe.Pointer(&_args[2])) = C.guint(nChars)

	girepository.MustFind("Gtk", "EntryBuffer").InvokeMethod("emit_deleted_text", _args[:], nil)

	runtime.KeepAlive(buffer)
	runtime.KeepAlive(position)
	runtime.KeepAlive(nChars)
}

// EmitInsertedText: used when subclassing GtkEntryBuffer.
//
// The function takes the following parameters:
//
//    - position at which text was inserted.
//    - chars: text that was inserted.
//    - nChars: number of characters inserted.
//
func (buffer *EntryBuffer) EmitInsertedText(position uint32, chars string, nChars uint32) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(position)
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(chars)))
	defer C.free(unsafe.Pointer(_args[2]))
	*(*C.guint)(unsafe.Pointer(&_args[3])) = C.guint(nChars)

	girepository.MustFind("Gtk", "EntryBuffer").InvokeMethod("emit_inserted_text", _args[:], nil)

	runtime.KeepAlive(buffer)
	runtime.KeepAlive(position)
	runtime.KeepAlive(chars)
	runtime.KeepAlive(nChars)
}

// Bytes retrieves the length in bytes of the buffer.
//
// See gtk.EntryBuffer.GetLength().
//
// The function returns the following values:
//
//    - gsize: byte length of the buffer.
//
func (buffer *EntryBuffer) Bytes() uint {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_gret := girepository.MustFind("Gtk", "EntryBuffer").InvokeMethod("get_bytes", _args[:], nil)
	_cret = *(*C.gsize)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(buffer)

	var _gsize uint // out

	_gsize = uint(*(*C.gsize)(unsafe.Pointer(&_cret)))

	return _gsize
}

// Length retrieves the length in characters of the buffer.
//
// The function returns the following values:
//
//    - guint: number of characters in the buffer.
//
func (buffer *EntryBuffer) Length() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_gret := girepository.MustFind("Gtk", "EntryBuffer").InvokeMethod("get_length", _args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(buffer)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// MaxLength retrieves the maximum allowed length of the text in buffer.
//
// The function returns the following values:
//
//    - gint: maximum allowed number of characters in EntryBuffer, or 0 if there
//      is no maximum.
//
func (buffer *EntryBuffer) MaxLength() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_gret := girepository.MustFind("Gtk", "EntryBuffer").InvokeMethod("get_max_length", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(buffer)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// Text retrieves the contents of the buffer.
//
// The memory pointer returned by this call will not change unless this object
// emits a signal, or is finalized.
//
// The function returns the following values:
//
//    - utf8: pointer to the contents of the widget as a string. This string
//      points to internally allocated storage in the buffer and must not be
//      freed, modified or stored.
//
func (buffer *EntryBuffer) Text() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_gret := girepository.MustFind("Gtk", "EntryBuffer").InvokeMethod("get_text", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
//    - position at which to insert text.
//    - chars: text to insert into the buffer.
//    - nChars: length of the text in characters, or -1.
//
// The function returns the following values:
//
//    - guint: number of characters actually inserted.
//
func (buffer *EntryBuffer) InsertText(position uint32, chars string, nChars int32) uint32 {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(position)
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(chars)))
	defer C.free(unsafe.Pointer(_args[2]))
	*(*C.int)(unsafe.Pointer(&_args[3])) = C.int(nChars)

	_gret := girepository.MustFind("Gtk", "EntryBuffer").InvokeMethod("insert_text", _args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(buffer)
	runtime.KeepAlive(position)
	runtime.KeepAlive(chars)
	runtime.KeepAlive(nChars)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// SetMaxLength sets the maximum allowed length of the contents of the buffer.
//
// If the current contents are longer than the given length, then they will be
// truncated to fit.
//
// The function takes the following parameters:
//
//    - maxLength: maximum length of the entry buffer, or 0 for no maximum.
//      (other than the maximum length of entries.) The value passed in will be
//      clamped to the range 0-65536.
//
func (buffer *EntryBuffer) SetMaxLength(maxLength int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(maxLength)

	girepository.MustFind("Gtk", "EntryBuffer").InvokeMethod("set_max_length", _args[:], nil)

	runtime.KeepAlive(buffer)
	runtime.KeepAlive(maxLength)
}

// SetText sets the text in the buffer.
//
// This is roughly equivalent to calling gtk.EntryBuffer.DeleteText() and
// gtk.EntryBuffer.InsertText().
//
// Note that n_chars is in characters, not in bytes.
//
// The function takes the following parameters:
//
//    - chars: new text.
//    - nChars: number of characters in text, or -1.
//
func (buffer *EntryBuffer) SetText(chars string, nChars int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(chars)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(nChars)

	girepository.MustFind("Gtk", "EntryBuffer").InvokeMethod("set_text", _args[:], nil)

	runtime.KeepAlive(buffer)
	runtime.KeepAlive(chars)
	runtime.KeepAlive(nChars)
}
