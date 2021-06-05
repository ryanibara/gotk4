// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_combo_box_text_get_type()), F: marshalComboBoxText},
	})
}

// ComboBoxText: a GtkComboBoxText is a simple variant of ComboBox that hides
// the model-view complexity for simple text-only use cases.
//
// To create a GtkComboBoxText, use gtk_combo_box_text_new() or
// gtk_combo_box_text_new_with_entry().
//
// You can add items to a GtkComboBoxText with gtk_combo_box_text_append_text(),
// gtk_combo_box_text_insert_text() or gtk_combo_box_text_prepend_text() and
// remove options with gtk_combo_box_text_remove().
//
// If the GtkComboBoxText contains an entry (via the “has-entry” property), its
// contents can be retrieved using gtk_combo_box_text_get_active_text(). The
// entry itself can be accessed by calling gtk_bin_get_child() on the combo box.
//
// You should not call gtk_combo_box_set_model() or attempt to pack more cells
// into this combo box via its GtkCellLayout interface.
//
//
// GtkComboBoxText as GtkBuildable
//
// The GtkComboBoxText implementation of the GtkBuildable interface supports
// adding items directly using the <items> element and specifying <item>
// elements for each item. Each <item> element can specify the “id”
// corresponding to the appended text and also supports the regular translation
// attributes “translatable”, “context” and “comments”.
//
// Here is a UI definition fragment specifying GtkComboBoxText items:
//
//    <object class="GtkComboBoxText">
//      <items>
//        <item translatable="yes" id="factory">Factory</item>
//        <item translatable="yes" id="home">Home</item>
//        <item translatable="yes" id="subway">Subway</item>
//      </items>
//    </object>
//
// CSS nodes
//
//    combobox
//    ╰── box.linked
//        ├── entry.combo
//        ├── button.combo
//        ╰── window.popup
//
// GtkComboBoxText has a single CSS node with name combobox. It adds the style
// class .combo to the main CSS nodes of its entry and button children, and the
// .linked class to the node of its internal box.
type ComboBoxText interface {
	ComboBox
	Buildable
	CellEditable
	CellLayout

	// Append appends @text to the list of strings stored in @combo_box. If @id
	// is non-nil then it is used as the ID of the row.
	//
	// This is the same as calling gtk_combo_box_text_insert() with a position
	// of -1.
	Append(id string, text string)
	// AppendText appends @text to the list of strings stored in @combo_box.
	//
	// This is the same as calling gtk_combo_box_text_insert_text() with a
	// position of -1.
	AppendText(text string)
	// ActiveText returns the currently active string in @combo_box, or nil if
	// none is selected. If @combo_box contains an entry, this function will
	// return its contents (which will not necessarily be an item from the
	// list).
	ActiveText() string
	// Insert inserts @text at @position in the list of strings stored in
	// @combo_box. If @id is non-nil then it is used as the ID of the row. See
	// ComboBox:id-column.
	//
	// If @position is negative then @text is appended.
	Insert(position int, id string, text string)
	// InsertText inserts @text at @position in the list of strings stored in
	// @combo_box.
	//
	// If @position is negative then @text is appended.
	//
	// This is the same as calling gtk_combo_box_text_insert() with a nil ID
	// string.
	InsertText(position int, text string)
	// Prepend prepends @text to the list of strings stored in @combo_box. If
	// @id is non-nil then it is used as the ID of the row.
	//
	// This is the same as calling gtk_combo_box_text_insert() with a position
	// of 0.
	Prepend(id string, text string)
	// PrependText prepends @text to the list of strings stored in @combo_box.
	//
	// This is the same as calling gtk_combo_box_text_insert_text() with a
	// position of 0.
	PrependText(text string)
	// Remove removes the string at @position from @combo_box.
	Remove(position int)
	// RemoveAll removes all the text entries from the combo box.
	RemoveAll()
}

// comboBoxText implements the ComboBoxText interface.
type comboBoxText struct {
	ComboBox
	Buildable
	CellEditable
	CellLayout
}

var _ ComboBoxText = (*comboBoxText)(nil)

// WrapComboBoxText wraps a GObject to the right type. It is
// primarily used internally.
func WrapComboBoxText(obj *externglib.Object) ComboBoxText {
	return ComboBoxText{
		ComboBox:     WrapComboBox(obj),
		Buildable:    WrapBuildable(obj),
		CellEditable: WrapCellEditable(obj),
		CellLayout:   WrapCellLayout(obj),
	}
}

func marshalComboBoxText(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapComboBoxText(obj), nil
}

// NewComboBoxText constructs a class ComboBoxText.
func NewComboBoxText() ComboBoxText {
	var cret C.GtkComboBoxText
	var ret1 ComboBoxText

	cret = C.gtk_combo_box_text_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ComboBoxText)

	return ret1
}

// NewComboBoxTextWithEntry constructs a class ComboBoxText.
func NewComboBoxTextWithEntry() ComboBoxText {
	var cret C.GtkComboBoxText
	var ret1 ComboBoxText

	cret = C.gtk_combo_box_text_new_with_entry()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ComboBoxText)

	return ret1
}

// Append appends @text to the list of strings stored in @combo_box. If @id
// is non-nil then it is used as the ID of the row.
//
// This is the same as calling gtk_combo_box_text_insert() with a position
// of -1.
func (c comboBoxText) Append(id string, text string) {
	var arg0 *C.GtkComboBoxText
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))
	arg1 = (*C.gchar)(C.CString(id))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg2))

	C.gtk_combo_box_text_append(arg0, id, text)
}

// AppendText appends @text to the list of strings stored in @combo_box.
//
// This is the same as calling gtk_combo_box_text_insert_text() with a
// position of -1.
func (c comboBoxText) AppendText(text string) {
	var arg0 *C.GtkComboBoxText
	var arg1 *C.gchar

	arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))
	arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_combo_box_text_append_text(arg0, text)
}

// ActiveText returns the currently active string in @combo_box, or nil if
// none is selected. If @combo_box contains an entry, this function will
// return its contents (which will not necessarily be an item from the
// list).
func (c comboBoxText) ActiveText() string {
	var arg0 *C.GtkComboBoxText

	arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_combo_box_text_get_active_text(arg0)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// Insert inserts @text at @position in the list of strings stored in
// @combo_box. If @id is non-nil then it is used as the ID of the row. See
// ComboBox:id-column.
//
// If @position is negative then @text is appended.
func (c comboBoxText) Insert(position int, id string, text string) {
	var arg0 *C.GtkComboBoxText
	var arg1 C.gint
	var arg2 *C.gchar
	var arg3 *C.gchar

	arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))
	arg1 = C.gint(position)
	arg2 = (*C.gchar)(C.CString(id))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg3))

	C.gtk_combo_box_text_insert(arg0, position, id, text)
}

// InsertText inserts @text at @position in the list of strings stored in
// @combo_box.
//
// If @position is negative then @text is appended.
//
// This is the same as calling gtk_combo_box_text_insert() with a nil ID
// string.
func (c comboBoxText) InsertText(position int, text string) {
	var arg0 *C.GtkComboBoxText
	var arg1 C.gint
	var arg2 *C.gchar

	arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))
	arg1 = C.gint(position)
	arg2 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg2))

	C.gtk_combo_box_text_insert_text(arg0, position, text)
}

// Prepend prepends @text to the list of strings stored in @combo_box. If
// @id is non-nil then it is used as the ID of the row.
//
// This is the same as calling gtk_combo_box_text_insert() with a position
// of 0.
func (c comboBoxText) Prepend(id string, text string) {
	var arg0 *C.GtkComboBoxText
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))
	arg1 = (*C.gchar)(C.CString(id))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg2))

	C.gtk_combo_box_text_prepend(arg0, id, text)
}

// PrependText prepends @text to the list of strings stored in @combo_box.
//
// This is the same as calling gtk_combo_box_text_insert_text() with a
// position of 0.
func (c comboBoxText) PrependText(text string) {
	var arg0 *C.GtkComboBoxText
	var arg1 *C.gchar

	arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))
	arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_combo_box_text_prepend_text(arg0, text)
}

// Remove removes the string at @position from @combo_box.
func (c comboBoxText) Remove(position int) {
	var arg0 *C.GtkComboBoxText
	var arg1 C.gint

	arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))
	arg1 = C.gint(position)

	C.gtk_combo_box_text_remove(arg0, position)
}

// RemoveAll removes all the text entries from the combo box.
func (c comboBoxText) RemoveAll() {
	var arg0 *C.GtkComboBoxText

	arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))

	C.gtk_combo_box_text_remove_all(arg0)
}

type ComboBoxTextPrivate struct {
	native C.GtkComboBoxTextPrivate
}

// WrapComboBoxTextPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapComboBoxTextPrivate(ptr unsafe.Pointer) *ComboBoxTextPrivate {
	if ptr == nil {
		return nil
	}

	return (*ComboBoxTextPrivate)(ptr)
}

func marshalComboBoxTextPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapComboBoxTextPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *ComboBoxTextPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}
