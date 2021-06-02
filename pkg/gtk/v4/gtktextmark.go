// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
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

	// Buffer gets the buffer this mark is located inside.
	//
	// Returns nil if the mark is deleted.
	Buffer() TextBuffer
	// Deleted returns true if the mark has been removed from its buffer.
	//
	// See [method@Gtk.TextBuffer.add_mark] for a way to add it to a buffer
	// again.
	Deleted() bool
	// LeftGravity determines whether the mark has left gravity.
	LeftGravity() bool
	// Name returns the mark name.
	//
	// Returns nil for anonymous marks.
	Name() string
	// Visible returns true if the mark is visible.
	//
	// A cursor is displayed for visible marks.
	Visible() bool

	SetVisible(setting bool)
}

// textMark implements the TextMark interface.
type textMark struct {
	gextras.Objector
}

var _ TextMark = (*textMark)(nil)

// WrapTextMark wraps a GObject to the right type. It is
// primarily used internally.
func WrapTextMark(obj *externglib.Object) TextMark {
	return TextMark{
		Objector: obj,
	}
}

func marshalTextMark(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTextMark(obj), nil
}

// NewTextMark constructs a class TextMark.
func NewTextMark(name string, leftGravity bool) TextMark {
	var arg1 *C.char
	var arg2 C.gboolean

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	if leftGravity {
		arg2 = C.TRUE
	}

	ret := C.gtk_text_mark_new(arg1, arg2)

	var ret0 TextMark

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(TextMark)

	return ret0
}

// Buffer gets the buffer this mark is located inside.
//
// Returns nil if the mark is deleted.
func (m textMark) Buffer() TextBuffer {
	var arg0 *C.GtkTextMark

	arg0 = (*C.GtkTextMark)(m.Native())

	ret := C.gtk_text_mark_get_buffer(arg0)

	var ret0 TextBuffer

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(TextBuffer)

	return ret0
}

// Deleted returns true if the mark has been removed from its buffer.
//
// See [method@Gtk.TextBuffer.add_mark] for a way to add it to a buffer
// again.
func (m textMark) Deleted() bool {
	var arg0 *C.GtkTextMark

	arg0 = (*C.GtkTextMark)(m.Native())

	ret := C.gtk_text_mark_get_deleted(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// LeftGravity determines whether the mark has left gravity.
func (m textMark) LeftGravity() bool {
	var arg0 *C.GtkTextMark

	arg0 = (*C.GtkTextMark)(m.Native())

	ret := C.gtk_text_mark_get_left_gravity(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Name returns the mark name.
//
// Returns nil for anonymous marks.
func (m textMark) Name() string {
	var arg0 *C.GtkTextMark

	arg0 = (*C.GtkTextMark)(m.Native())

	ret := C.gtk_text_mark_get_name(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// Visible returns true if the mark is visible.
//
// A cursor is displayed for visible marks.
func (m textMark) Visible() bool {
	var arg0 *C.GtkTextMark

	arg0 = (*C.GtkTextMark)(m.Native())

	ret := C.gtk_text_mark_get_visible(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

func (m textMark) SetVisible(setting bool) {
	var arg0 *C.GtkTextMark
	var arg1 C.gboolean

	arg0 = (*C.GtkTextMark)(m.Native())
	if setting {
		arg1 = C.TRUE
	}

	C.gtk_text_mark_set_visible(arg0, arg1)
}
