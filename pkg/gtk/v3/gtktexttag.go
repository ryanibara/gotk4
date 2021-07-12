// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_text_tag_get_type()), F: marshalTextTager},
	})
}

// TextTager describes TextTag's methods.
type TextTager interface {
	// Changed emits the TextTagTable::tag-changed signal on the TextTagTable
	// where the tag is included.
	Changed(sizeChanged bool)
	// Priority: get the tag priority.
	Priority() int
	// SetPriority sets the priority of a TextTag.
	SetPriority(priority int)
}

// TextTag: you may wish to begin by reading the [text widget conceptual
// overview][TextWidget] which gives an overview of all the objects and data
// types related to the text widget and how they work together.
//
// Tags should be in the TextTagTable for a given TextBuffer before using them
// with that buffer.
//
// gtk_text_buffer_create_tag() is the best way to create tags. See “gtk3-demo”
// for numerous examples.
//
// For each property of TextTag, there is a “set” property, e.g. “font-set”
// corresponds to “font”. These “set” properties reflect whether a property has
// been set or not. They are maintained by GTK+ and you should not set them
// independently.
type TextTag struct {
	*externglib.Object
}

var (
	_ TextTager       = (*TextTag)(nil)
	_ gextras.Nativer = (*TextTag)(nil)
)

func wrapTextTag(obj *externglib.Object) *TextTag {
	return &TextTag{
		Object: obj,
	}
}

func marshalTextTager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTextTag(obj), nil
}

// NewTextTag creates a TextTag. Configure the tag using object arguments, i.e.
// using g_object_set().
func NewTextTag(name string) *TextTag {
	var _arg1 *C.gchar      // out
	var _cret *C.GtkTextTag // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))

	_cret = C.gtk_text_tag_new(_arg1)

	var _textTag *TextTag // out

	_textTag = wrapTextTag(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _textTag
}

// Changed emits the TextTagTable::tag-changed signal on the TextTagTable where
// the tag is included.
//
// The signal is already emitted when setting a TextTag property. This function
// is useful for a TextTag subclass.
func (tag *TextTag) Changed(sizeChanged bool) {
	var _arg0 *C.GtkTextTag // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))
	if sizeChanged {
		_arg1 = C.TRUE
	}

	C.gtk_text_tag_changed(_arg0, _arg1)
}

// Priority: get the tag priority.
func (tag *TextTag) Priority() int {
	var _arg0 *C.GtkTextTag // out
	var _cret C.gint        // in

	_arg0 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))

	_cret = C.gtk_text_tag_get_priority(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SetPriority sets the priority of a TextTag. Valid priorities start at 0 and
// go to one less than gtk_text_tag_table_get_size(). Each tag in a table has a
// unique priority; setting the priority of one tag shifts the priorities of all
// the other tags in the table to maintain a unique priority for each tag.
// Higher priority tags “win” if two tags both set the same text attribute. When
// adding a tag to a tag table, it will be assigned the highest priority in the
// table by default; so normally the precedence of a set of tags is the order in
// which they were added to the table, or created with
// gtk_text_buffer_create_tag(), which adds the tag to the buffer’s table
// automatically.
func (tag *TextTag) SetPriority(priority int) {
	var _arg0 *C.GtkTextTag // out
	var _arg1 C.gint        // out

	_arg0 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))
	_arg1 = C.gint(priority)

	C.gtk_text_tag_set_priority(_arg0, _arg1)
}
