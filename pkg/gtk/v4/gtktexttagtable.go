// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
//
// void gotk4_TextTagTableForeach(GtkTextTag*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_text_tag_table_get_type()), F: marshalTextTagTable},
	})
}

// TextTagTableForeach: a function used with gtk_text_tag_table_foreach(), to
// iterate over every `GtkTextTag` inside a `GtkTextTagTable`.
type TextTagTableForeach func(tag TextTag)

//export gotk4_TextTagTableForeach
func _TextTagTableForeach(arg0 *C.GtkTextTag, arg1 C.gpointer) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var tag TextTag // out

	tag = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(TextTag)

	fn := v.(TextTagTableForeach)
	fn(tag)
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
	Buildable

	AddTextTagTable(tag TextTag) bool

	ForeachTextTagTable(fn TextTagTableForeach)

	Size() int

	LookupTextTagTable(name string) TextTag

	RemoveTextTagTable(tag TextTag)
}

// textTagTable implements the TextTagTable class.
type textTagTable struct {
	gextras.Objector
}

// WrapTextTagTable wraps a GObject to the right type. It is
// primarily used internally.
func WrapTextTagTable(obj *externglib.Object) TextTagTable {
	return textTagTable{
		Objector: obj,
	}
}

func marshalTextTagTable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTextTagTable(obj), nil
}

func NewTextTagTable() TextTagTable {
	var _cret *C.GtkTextTagTable // in

	_cret = C.gtk_text_tag_table_new()

	var _textTagTable TextTagTable // out

	_textTagTable = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(TextTagTable)

	return _textTagTable
}

func (t textTagTable) AddTextTagTable(tag TextTag) bool {
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

func (t textTagTable) ForeachTextTagTable(fn TextTagTableForeach) {
	var _arg0 *C.GtkTextTagTable       // out
	var _arg1 C.GtkTextTagTableForeach // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*[0]byte)(C.gotk4_TextTagTableForeach)
	_arg2 = C.gpointer(box.Assign(fn))

	C.gtk_text_tag_table_foreach(_arg0, _arg1, _arg2)
}

func (t textTagTable) Size() int {
	var _arg0 *C.GtkTextTagTable // out
	var _cret C.int              // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_text_tag_table_get_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (t textTagTable) LookupTextTagTable(name string) TextTag {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.char            // out
	var _cret *C.GtkTextTag      // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_text_tag_table_lookup(_arg0, _arg1)

	var _textTag TextTag // out

	_textTag = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(TextTag)

	return _textTag
}

func (t textTagTable) RemoveTextTagTable(tag TextTag) {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.GtkTextTag      // out

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))

	C.gtk_text_tag_table_remove(_arg0, _arg1)
}

func (b textTagTable) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
