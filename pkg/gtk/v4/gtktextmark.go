// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_text_mark_get_type()), F: marshalTextMark},
	})
}

// TextMark: a `GtkTextMark` is a position in a `GtkTextbuffer` that is
// preserved across modifications.
//
// You may wish to begin by reading the text widget conceptual overview
// (section-text-widget.html), which gives an overview of all the objects and
// data types related to the text widget and how they work together.
//
// A `GtkTextMark` is like a bookmark in a text buffer; it preserves a position
// in the text. You can convert the mark to an iterator using
// [method@Gtk.TextBuffer.get_iter_at_mark]. Unlike iterators, marks remain
// valid across buffer mutations, because their behavior is defined when text is
// inserted or deleted. When text containing a mark is deleted, the mark remains
// in the position originally occupied by the deleted text. When text is
// inserted at a mark, a mark with “left gravity” will be moved to the beginning
// of the newly-inserted text, and a mark with “right gravity” will be moved to
// the end.
//
// Note that “left” and “right” here refer to logical direction (left is the
// toward the start of the buffer); in some languages such as Hebrew the
// logically-leftmost text is not actually on the left when displayed.
//
// Marks are reference counted, but the reference count only controls the
// validity of the memory; marks can be deleted from the buffer at any time with
// [method@Gtk.TextBuffer.delete_mark]. Once deleted from the buffer, a mark is
// essentially useless.
//
// Marks optionally have names; these can be convenient to avoid passing the
// `GtkTextMark` object around.
//
// Marks are typically created using the [method@Gtk.TextBuffer.create_mark]
// function.
type TextMark interface {
	gextras.Objector

	Buffer() TextBuffer

	Deleted() bool

	LeftGravity() bool

	Name() string

	Visible() bool

	SetVisibleTextMark(setting bool)
}

// textMark implements the TextMark class.
type textMark struct {
	gextras.Objector
}

// WrapTextMark wraps a GObject to the right type. It is
// primarily used internally.
func WrapTextMark(obj *externglib.Object) TextMark {
	return textMark{
		Objector: obj,
	}
}

func marshalTextMark(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTextMark(obj), nil
}

func NewTextMark(name string, leftGravity bool) TextMark {
	var _arg1 *C.char        // out
	var _arg2 C.gboolean     // out
	var _cret *C.GtkTextMark // in

	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	if leftGravity {
		_arg2 = C.TRUE
	}

	_cret = C.gtk_text_mark_new(_arg1, _arg2)

	var _textMark TextMark // out

	_textMark = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(TextMark)

	return _textMark
}

func (m textMark) Buffer() TextBuffer {
	var _arg0 *C.GtkTextMark   // out
	var _cret *C.GtkTextBuffer // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_text_mark_get_buffer(_arg0)

	var _textBuffer TextBuffer // out

	_textBuffer = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(TextBuffer)

	return _textBuffer
}

func (m textMark) Deleted() bool {
	var _arg0 *C.GtkTextMark // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_text_mark_get_deleted(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (m textMark) LeftGravity() bool {
	var _arg0 *C.GtkTextMark // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_text_mark_get_left_gravity(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (m textMark) Name() string {
	var _arg0 *C.GtkTextMark // out
	var _cret *C.char        // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_text_mark_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (m textMark) Visible() bool {
	var _arg0 *C.GtkTextMark // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_text_mark_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (m textMark) SetVisibleTextMark(setting bool) {
	var _arg0 *C.GtkTextMark // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(m.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_text_mark_set_visible(_arg0, _arg1)
}
