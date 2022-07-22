// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
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
import "C"

// GType values.
var (
	GTypeTextMark = coreglib.Type(C.gtk_text_mark_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTextMark, F: marshalTextMark},
	})
}

// TextMarkOverrider contains methods that are overridable.
type TextMarkOverrider interface {
}

// TextMark: you may wish to begin by reading the [text widget conceptual
// overview][TextWidget] which gives an overview of all the objects and data
// types related to the text widget and how they work together.
//
// A TextMark is like a bookmark in a text buffer; it preserves a position in
// the text. You can convert the mark to an iterator using
// gtk_text_buffer_get_iter_at_mark(). Unlike iterators, marks remain valid
// across buffer mutations, because their behavior is defined when text is
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
// gtk_text_buffer_delete_mark(). Once deleted from the buffer, a mark is
// essentially useless.
//
// Marks optionally have names; these can be convenient to avoid passing the
// TextMark object around.
//
// Marks are typically created using the gtk_text_buffer_create_mark() function.
type TextMark struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TextMark)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:     GTypeTextMark,
		GoType:    reflect.TypeOf((*TextMark)(nil)),
		InitClass: initClassTextMark,
	})
}

func initClassTextMark(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ InitTextMark(*TextMarkClass) }); ok {
		klass := (*TextMarkClass)(gextras.NewStructNative(gclass))
		goval.InitTextMark(klass)
	}
}

func wrapTextMark(obj *coreglib.Object) *TextMark {
	return &TextMark{
		Object: obj,
	}
}

func marshalTextMark(p uintptr) (interface{}, error) {
	return wrapTextMark(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTextMark creates a text mark. Add it to a buffer using
// gtk_text_buffer_add_mark(). If name is NULL, the mark is anonymous;
// otherwise, the mark can be retrieved by name using
// gtk_text_buffer_get_mark(). If a mark has left gravity, and text is inserted
// at the mark’s current location, the mark will be moved to the left of the
// newly-inserted text. If the mark has right gravity (left_gravity = FALSE),
// the mark will end up on the right of newly-inserted text. The standard
// left-to-right cursor is a mark with right gravity (when you type, the cursor
// stays on the right side of the text you’re typing).
//
// The function takes the following parameters:
//
//    - name (optional): mark name or NULL.
//    - leftGravity: whether the mark should have left gravity.
//
// The function returns the following values:
//
//    - textMark: new TextMark.
//
func NewTextMark(name string, leftGravity bool) *TextMark {
	var _arg1 *C.gchar       // out
	var _arg2 C.gboolean     // out
	var _cret *C.GtkTextMark // in

	if name != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	if leftGravity {
		_arg2 = C.TRUE
	}

	_cret = C.gtk_text_mark_new(_arg1, _arg2)
	runtime.KeepAlive(name)
	runtime.KeepAlive(leftGravity)

	var _textMark *TextMark // out

	_textMark = wrapTextMark(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _textMark
}

// Buffer gets the buffer this mark is located inside, or NULL if the mark is
// deleted.
//
// The function returns the following values:
//
//    - textBuffer mark’s TextBuffer.
//
func (mark *TextMark) Buffer() *TextBuffer {
	var _arg0 *C.GtkTextMark   // out
	var _cret *C.GtkTextBuffer // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(coreglib.InternObject(mark).Native()))

	_cret = C.gtk_text_mark_get_buffer(_arg0)
	runtime.KeepAlive(mark)

	var _textBuffer *TextBuffer // out

	_textBuffer = wrapTextBuffer(coreglib.Take(unsafe.Pointer(_cret)))

	return _textBuffer
}

// Deleted returns TRUE if the mark has been removed from its buffer with
// gtk_text_buffer_delete_mark(). See gtk_text_buffer_add_mark() for a way to
// add it to a buffer again.
//
// The function returns the following values:
//
//    - ok: whether the mark is deleted.
//
func (mark *TextMark) Deleted() bool {
	var _arg0 *C.GtkTextMark // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(coreglib.InternObject(mark).Native()))

	_cret = C.gtk_text_mark_get_deleted(_arg0)
	runtime.KeepAlive(mark)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LeftGravity determines whether the mark has left gravity.
//
// The function returns the following values:
//
//    - ok: TRUE if the mark has left gravity, FALSE otherwise.
//
func (mark *TextMark) LeftGravity() bool {
	var _arg0 *C.GtkTextMark // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(coreglib.InternObject(mark).Native()))

	_cret = C.gtk_text_mark_get_left_gravity(_arg0)
	runtime.KeepAlive(mark)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Name returns the mark name; returns NULL for anonymous marks.
//
// The function returns the following values:
//
//    - utf8 (optional): mark name.
//
func (mark *TextMark) Name() string {
	var _arg0 *C.GtkTextMark // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(coreglib.InternObject(mark).Native()))

	_cret = C.gtk_text_mark_get_name(_arg0)
	runtime.KeepAlive(mark)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Visible returns TRUE if the mark is visible (i.e. a cursor is displayed for
// it).
//
// The function returns the following values:
//
//    - ok: TRUE if visible.
//
func (mark *TextMark) Visible() bool {
	var _arg0 *C.GtkTextMark // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(coreglib.InternObject(mark).Native()))

	_cret = C.gtk_text_mark_get_visible(_arg0)
	runtime.KeepAlive(mark)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetVisible sets the visibility of mark; the insertion point is normally
// visible, i.e. you can see it as a vertical bar. Also, the text widget uses a
// visible mark to indicate where a drop will occur when dragging-and-dropping
// text. Most other marks are not visible. Marks are not visible by default.
//
// The function takes the following parameters:
//
//    - setting: visibility of mark.
//
func (mark *TextMark) SetVisible(setting bool) {
	var _arg0 *C.GtkTextMark // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(coreglib.InternObject(mark).Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_text_mark_set_visible(_arg0, _arg1)
	runtime.KeepAlive(mark)
	runtime.KeepAlive(setting)
}

// TextMarkClass: instance of this type is always passed by reference.
type TextMarkClass struct {
	*textMarkClass
}

// textMarkClass is the struct that's finalized.
type textMarkClass struct {
	native *C.GtkTextMarkClass
}
