// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_TextTagTableForEach(GtkTextTag*, gpointer);
// extern void _gotk4_gtk4_TextTagTable_ConnectTagAdded(gpointer, GtkTextTag*, guintptr);
// extern void _gotk4_gtk4_TextTagTable_ConnectTagChanged(gpointer, GtkTextTag*, gboolean, guintptr);
// extern void _gotk4_gtk4_TextTagTable_ConnectTagRemoved(gpointer, GtkTextTag*, guintptr);
import "C"

// glib.Type values for gtktexttagtable.go.
var GTypeTextTagTable = externglib.Type(C.gtk_text_tag_table_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeTextTagTable, F: marshalTextTagTable},
	})
}

// TextTagTableForEach: function used with gtk_text_tag_table_foreach(), to
// iterate over every GtkTextTag inside a GtkTextTagTable.
type TextTagTableForEach func(tag *TextTag)

//export _gotk4_gtk4_TextTagTableForEach
func _gotk4_gtk4_TextTagTableForEach(arg1 *C.GtkTextTag, arg2 C.gpointer) {
	var fn TextTagTableForEach
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TextTagTableForEach)
	}

	var _tag *TextTag // out

	_tag = wrapTextTag(externglib.Take(unsafe.Pointer(arg1)))

	fn(_tag)
}

// TextTagTable: collection of tags in a GtkTextBuffer
//
// You may wish to begin by reading the text widget conceptual overview
// (section-text-widget.html), which gives an overview of all the objects and
// data types related to the text widget and how they work together.
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
//    </object>.
type TextTagTable struct {
	_ [0]func() // equal guard
	*externglib.Object

	Buildable
}

var (
	_ externglib.Objector = (*TextTagTable)(nil)
)

func wrapTextTagTable(obj *externglib.Object) *TextTagTable {
	return &TextTagTable{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalTextTagTable(p uintptr) (interface{}, error) {
	return wrapTextTagTable(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_TextTagTable_ConnectTagAdded
func _gotk4_gtk4_TextTagTable_ConnectTagAdded(arg0 C.gpointer, arg1 *C.GtkTextTag, arg2 C.guintptr) {
	var f func(tag *TextTag)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(tag *TextTag))
	}

	var _tag *TextTag // out

	_tag = wrapTextTag(externglib.Take(unsafe.Pointer(arg1)))

	f(_tag)
}

// ConnectTagAdded: emitted every time a new tag is added in the
// GtkTextTagTable.
func (table *TextTagTable) ConnectTagAdded(f func(tag *TextTag)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(table, "tag-added", false, unsafe.Pointer(C._gotk4_gtk4_TextTagTable_ConnectTagAdded), f)
}

//export _gotk4_gtk4_TextTagTable_ConnectTagChanged
func _gotk4_gtk4_TextTagTable_ConnectTagChanged(arg0 C.gpointer, arg1 *C.GtkTextTag, arg2 C.gboolean, arg3 C.guintptr) {
	var f func(tag *TextTag, sizeChanged bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(tag *TextTag, sizeChanged bool))
	}

	var _tag *TextTag     // out
	var _sizeChanged bool // out

	_tag = wrapTextTag(externglib.Take(unsafe.Pointer(arg1)))
	if arg2 != 0 {
		_sizeChanged = true
	}

	f(_tag, _sizeChanged)
}

// ConnectTagChanged: emitted every time a tag in the GtkTextTagTable changes.
func (table *TextTagTable) ConnectTagChanged(f func(tag *TextTag, sizeChanged bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(table, "tag-changed", false, unsafe.Pointer(C._gotk4_gtk4_TextTagTable_ConnectTagChanged), f)
}

//export _gotk4_gtk4_TextTagTable_ConnectTagRemoved
func _gotk4_gtk4_TextTagTable_ConnectTagRemoved(arg0 C.gpointer, arg1 *C.GtkTextTag, arg2 C.guintptr) {
	var f func(tag *TextTag)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(tag *TextTag))
	}

	var _tag *TextTag // out

	_tag = wrapTextTag(externglib.Take(unsafe.Pointer(arg1)))

	f(_tag)
}

// ConnectTagRemoved: emitted every time a tag is removed from the
// GtkTextTagTable.
//
// The tag is still valid by the time the signal is emitted, but it is not
// associated with a tag table any more.
func (table *TextTagTable) ConnectTagRemoved(f func(tag *TextTag)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(table, "tag-removed", false, unsafe.Pointer(C._gotk4_gtk4_TextTagTable_ConnectTagRemoved), f)
}

// NewTextTagTable creates a new GtkTextTagTable.
//
// The table contains no tags by default.
//
// The function returns the following values:
//
//    - textTagTable: new GtkTextTagTable.
//
func NewTextTagTable() *TextTagTable {
	var _cret *C.GtkTextTagTable // in

	_cret = C.gtk_text_tag_table_new()

	var _textTagTable *TextTagTable // out

	_textTagTable = wrapTextTagTable(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _textTagTable
}

// Add a tag to the table.
//
// The tag is assigned the highest priority in the table.
//
// tag must not be in a tag table already, and may not have the same name as an
// already-added tag.
//
// The function takes the following parameters:
//
//    - tag: GtkTextTag.
//
// The function returns the following values:
//
//    - ok: TRUE on success.
//
func (table *TextTagTable) Add(tag *TextTag) bool {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.GtkTextTag      // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(externglib.InternObject(table).Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(externglib.InternObject(tag).Native()))

	_cret = C.gtk_text_tag_table_add(_arg0, _arg1)
	runtime.KeepAlive(table)
	runtime.KeepAlive(tag)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ForEach calls func on each tag in table, with user data data.
//
// Note that the table may not be modified while iterating over it (you can’t
// add/remove tags).
//
// The function takes the following parameters:
//
//    - fn: function to call on each tag.
//
func (table *TextTagTable) ForEach(fn TextTagTableForEach) {
	var _arg0 *C.GtkTextTagTable       // out
	var _arg1 C.GtkTextTagTableForeach // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(externglib.InternObject(table).Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk4_TextTagTableForEach)
	_arg2 = C.gpointer(gbox.Assign(fn))
	defer gbox.Delete(uintptr(_arg2))

	C.gtk_text_tag_table_foreach(_arg0, _arg1, _arg2)
	runtime.KeepAlive(table)
	runtime.KeepAlive(fn)
}

// Size returns the size of the table (number of tags).
//
// The function returns the following values:
//
//    - gint: number of tags in table.
//
func (table *TextTagTable) Size() int {
	var _arg0 *C.GtkTextTagTable // out
	var _cret C.int              // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(externglib.InternObject(table).Native()))

	_cret = C.gtk_text_tag_table_get_size(_arg0)
	runtime.KeepAlive(table)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Lookup: look up a named tag.
//
// The function takes the following parameters:
//
//    - name of a tag.
//
// The function returns the following values:
//
//    - textTag (optional): tag, or NULL if none by that name is in the table.
//
func (table *TextTagTable) Lookup(name string) *TextTag {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.char            // out
	var _cret *C.GtkTextTag      // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(externglib.InternObject(table).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_text_tag_table_lookup(_arg0, _arg1)
	runtime.KeepAlive(table)
	runtime.KeepAlive(name)

	var _textTag *TextTag // out

	if _cret != nil {
		_textTag = wrapTextTag(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _textTag
}

// Remove a tag from the table.
//
// If a GtkTextBuffer has table as its tag table, the tag is removed from the
// buffer. The table’s reference to the tag is removed, so the tag will end up
// destroyed if you don’t have a reference to it.
//
// The function takes the following parameters:
//
//    - tag: GtkTextTag.
//
func (table *TextTagTable) Remove(tag *TextTag) {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.GtkTextTag      // out

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(externglib.InternObject(table).Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(externglib.InternObject(tag).Native()))

	C.gtk_text_tag_table_remove(_arg0, _arg1)
	runtime.KeepAlive(table)
	runtime.KeepAlive(tag)
}
