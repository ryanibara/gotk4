// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_text_mark_get_type()), F: marshalTextMarker},
	})
}

// TextMarker describes TextMark's methods.
type TextMarker interface {
	// Buffer gets the buffer this mark is located inside.
	Buffer() *TextBuffer
	// Deleted returns true if the mark has been removed from its buffer.
	Deleted() bool
	// LeftGravity determines whether the mark has left gravity.
	LeftGravity() bool
	// Name returns the mark name.
	Name() string
	// Visible returns true if the mark is visible.
	Visible() bool

	SetVisible(setting bool)
}

// TextMark: `GtkTextMark` is a position in a `GtkTextbuffer` that is preserved
// across modifications.
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
type TextMark struct {
	*externglib.Object
}

var (
	_ TextMarker      = (*TextMark)(nil)
	_ gextras.Nativer = (*TextMark)(nil)
)

func wrapTextMark(obj *externglib.Object) TextMarker {
	return &TextMark{
		Object: obj,
	}
}

func marshalTextMarker(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTextMark(obj), nil
}

// NewTextMark creates a text mark.
//
// Add it to a buffer using [method@Gtk.TextBuffer.add_mark]. If @name is nil,
// the mark is anonymous; otherwise, the mark can be retrieved by name using
// [method@Gtk.TextBuffer.get_mark]. If a mark has left gravity, and text is
// inserted at the mark’s current location, the mark will be moved to the left
// of the newly-inserted text. If the mark has right gravity (@left_gravity =
// false), the mark will end up on the right of newly-inserted text. The
// standard left-to-right cursor is a mark with right gravity (when you type,
// the cursor stays on the right side of the text you’re typing).
func NewTextMark(name string, leftGravity bool) *TextMark {
	var _arg1 *C.char        // out
	var _arg2 C.gboolean     // out
	var _cret *C.GtkTextMark // in

	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	if leftGravity {
		_arg2 = C.TRUE
	}

	_cret = C.gtk_text_mark_new(_arg1, _arg2)

	var _textMark *TextMark // out

	_textMark = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*TextMark)

	return _textMark
}

// Buffer gets the buffer this mark is located inside.
//
// Returns nil if the mark is deleted.
func (mark *TextMark) Buffer() *TextBuffer {
	var _arg0 *C.GtkTextMark   // out
	var _cret *C.GtkTextBuffer // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(mark.Native()))

	_cret = C.gtk_text_mark_get_buffer(_arg0)

	var _textBuffer *TextBuffer // out

	_textBuffer = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*TextBuffer)

	return _textBuffer
}

// Deleted returns true if the mark has been removed from its buffer.
//
// See [method@Gtk.TextBuffer.add_mark] for a way to add it to a buffer again.
func (mark *TextMark) Deleted() bool {
	var _arg0 *C.GtkTextMark // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(mark.Native()))

	_cret = C.gtk_text_mark_get_deleted(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LeftGravity determines whether the mark has left gravity.
func (mark *TextMark) LeftGravity() bool {
	var _arg0 *C.GtkTextMark // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(mark.Native()))

	_cret = C.gtk_text_mark_get_left_gravity(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Name returns the mark name.
//
// Returns nil for anonymous marks.
func (mark *TextMark) Name() string {
	var _arg0 *C.GtkTextMark // out
	var _cret *C.char        // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(mark.Native()))

	_cret = C.gtk_text_mark_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Visible returns true if the mark is visible.
//
// A cursor is displayed for visible marks.
func (mark *TextMark) Visible() bool {
	var _arg0 *C.GtkTextMark // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(mark.Native()))

	_cret = C.gtk_text_mark_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (mark *TextMark) SetVisible(setting bool) {
	var _arg0 *C.GtkTextMark // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(mark.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_text_mark_set_visible(_arg0, _arg1)
}
