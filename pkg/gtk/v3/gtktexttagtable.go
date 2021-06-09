// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
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
		{T: externglib.Type(C.gtk_text_tag_table_get_type()), F: marshalTextTagTable},
	})
}

type TextTagTableForeach func()

//export gotk4_TextTagTableForeach
func gotk4_TextTagTableForeach(arg0 *C.GtkTextTag, arg1 C.gpointer) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(TextTagTableForeach)
	fn()
}

// TextTagTable: you may wish to begin by reading the [text widget conceptual
// overview][TextWidget] which gives an overview of all the objects and data
// types related to the text widget and how they work together.
//
//
// GtkTextTagTables as GtkBuildable
//
// The GtkTextTagTable implementation of the GtkBuildable interface supports
// adding tags by specifying “tag” as the “type” attribute of a <child> element.
//
// An example of a UI definition fragment specifying tags:
//
//    <object class="GtkTextTagTable">
//     <child type="tag">
//       <object class="GtkTextTag"/>
//     </child>
//    </object>
type TextTagTable interface {
	gextras.Objector
	Buildable

	// Add: add a tag to the table. The tag is assigned the highest priority in
	// the table.
	//
	// @tag must not be in a tag table already, and may not have the same name
	// as an already-added tag.
	Add(tag TextTag) bool
	// Foreach calls @func on each tag in @table, with user data @data. Note
	// that the table may not be modified while iterating over it (you can’t
	// add/remove tags).
	Foreach()
	// Size returns the size of the table (number of tags)
	Size() int
	// Lookup: look up a named tag.
	Lookup(name string) TextTag
	// Remove: remove a tag from the table. If a TextBuffer has @table as its
	// tag table, the tag is removed from the buffer. The table’s reference to
	// the tag is removed, so the tag will end up destroyed if you don’t have a
	// reference to it.
	Remove(tag TextTag)
}

// textTagTable implements the TextTagTable interface.
type textTagTable struct {
	gextras.Objector
	Buildable
}

var _ TextTagTable = (*textTagTable)(nil)

// WrapTextTagTable wraps a GObject to the right type. It is
// primarily used internally.
func WrapTextTagTable(obj *externglib.Object) TextTagTable {
	return TextTagTable{
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
	var _cret C.GtkTextTagTable

	cret = C.gtk_text_tag_table_new()

	var _textTagTable TextTagTable

	_textTagTable = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(TextTagTable)

	return _textTagTable
}

// Add: add a tag to the table. The tag is assigned the highest priority in
// the table.
//
// @tag must not be in a tag table already, and may not have the same name
// as an already-added tag.
func (t textTagTable) Add(tag TextTag) bool {
	var _arg0 *C.GtkTextTagTable
	var _arg1 *C.GtkTextTag

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))

	var _cret C.gboolean

	cret = C.gtk_text_tag_table_add(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Foreach calls @func on each tag in @table, with user data @data. Note
// that the table may not be modified while iterating over it (you can’t
// add/remove tags).
func (t textTagTable) Foreach() {
	var _arg0 *C.GtkTextTagTable

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))

	C.gtk_text_tag_table_foreach(_arg0)
}

// Size returns the size of the table (number of tags)
func (t textTagTable) Size() int {
	var _arg0 *C.GtkTextTagTable

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))

	var _cret C.gint

	cret = C.gtk_text_tag_table_get_size(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Lookup: look up a named tag.
func (t textTagTable) Lookup(name string) TextTag {
	var _arg0 *C.GtkTextTagTable
	var _arg1 *C.gchar

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.GtkTextTag

	cret = C.gtk_text_tag_table_lookup(_arg0, _arg1)

	var _textTag TextTag

	_textTag = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(TextTag)

	return _textTag
}

// Remove: remove a tag from the table. If a TextBuffer has @table as its
// tag table, the tag is removed from the buffer. The table’s reference to
// the tag is removed, so the tag will end up destroyed if you don’t have a
// reference to it.
func (t textTagTable) Remove(tag TextTag) {
	var _arg0 *C.GtkTextTagTable
	var _arg1 *C.GtkTextTag

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))

	C.gtk_text_tag_table_remove(_arg0, _arg1)
}