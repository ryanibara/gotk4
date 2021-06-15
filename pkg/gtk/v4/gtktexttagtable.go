// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_text_tag_table_get_type()), F: marshalTextTagTable},
	})
}

// TextTagTable: the collection of tags in a `GtkTextBuffer`
//
// You may wish to begin by reading the text widget conceptual overview
// (section-text-widget.html), which gives an overview of all the objects and
// data types related to the text widget and how they work together.
//
//
// GtkTextTagTables as GtkBuildable
//
// The `GtkTextTagTable` implementation of the `GtkBuildable` interface supports
// adding tags by specifying “tag” as the “type” attribute of a <child> element.
//
// An example of a UI definition fragment specifying tags: “`xml <object
// class="GtkTextTagTable"> <child type="tag"> <object class="GtkTextTag"/>
// </child> </object> “`
type TextTagTable interface {
	gextras.Objector
	Buildable

	// Add: add a tag to the table.
	//
	// The tag is assigned the highest priority in the table.
	//
	// @tag must not be in a tag table already, and may not have the same name
	// as an already-added tag.
	Add(tag TextTag) bool
	// Size returns the size of the table (number of tags)
	Size() int
	// Lookup: look up a named tag.
	Lookup(name string) TextTag
	// Remove: remove a tag from the table.
	//
	// If a `GtkTextBuffer` has @table as its tag table, the tag is removed from
	// the buffer. The table’s reference to the tag is removed, so the tag will
	// end up destroyed if you don’t have a reference to it.
	Remove(tag TextTag)
}

// textTagTable implements the TextTagTable class.
type textTagTable struct {
	gextras.Objector
	Buildable
}

var _ TextTagTable = (*textTagTable)(nil)

// WrapTextTagTable wraps a GObject to the right type. It is
// primarily used internally.
func WrapTextTagTable(obj *externglib.Object) TextTagTable {
	return textTagTable{
		Objector:  obj,
		Buildable: WrapBuildable(obj),
	}
}

func marshalTextTagTable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTextTagTable(obj), nil
}

// NewTextTagTable constructs a class TextTagTable.
func NewTextTagTable() TextTagTable {
	var _cret C.GtkTextTagTable // in

	_cret = C.gtk_text_tag_table_new()

	var _textTagTable TextTagTable // out

	_textTagTable = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(TextTagTable)

	return _textTagTable
}

// Add: add a tag to the table.
//
// The tag is assigned the highest priority in the table.
//
// @tag must not be in a tag table already, and may not have the same name
// as an already-added tag.
func (t textTagTable) Add(tag TextTag) bool {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.GtkTextTag      // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))

	_cret = C.gtk_text_tag_table_add(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Size returns the size of the table (number of tags)
func (t textTagTable) Size() int {
	var _arg0 *C.GtkTextTagTable // out
	var _cret C.int              // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_text_tag_table_get_size(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Lookup: look up a named tag.
func (t textTagTable) Lookup(name string) TextTag {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.char            // out
	var _cret *C.GtkTextTag      // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_text_tag_table_lookup(_arg0, _arg1)

	var _textTag TextTag // out

	_textTag = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(TextTag)

	return _textTag
}

// Remove: remove a tag from the table.
//
// If a `GtkTextBuffer` has @table as its tag table, the tag is removed from
// the buffer. The table’s reference to the tag is removed, so the tag will
// end up destroyed if you don’t have a reference to it.
func (t textTagTable) Remove(tag TextTag) {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.GtkTextTag      // out

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))

	C.gtk_text_tag_table_remove(_arg0, _arg1)
}
