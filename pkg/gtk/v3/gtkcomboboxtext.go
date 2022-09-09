// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
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
	GTypeComboBoxText = coreglib.Type(C.gtk_combo_box_text_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeComboBoxText, F: marshalComboBoxText},
	})
}

// ComboBoxTextOverrides contains methods that are overridable.
type ComboBoxTextOverrides struct {
}

func defaultComboBoxTextOverrides(v *ComboBoxText) ComboBoxTextOverrides {
	return ComboBoxTextOverrides{}
}

// ComboBoxText is a simple variant of ComboBox that hides the model-view
// complexity for simple text-only use cases.
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
type ComboBoxText struct {
	_ [0]func() // equal guard
	ComboBox
}

var (
	_ Binner            = (*ComboBoxText)(nil)
	_ coreglib.Objector = (*ComboBoxText)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ComboBoxText, *ComboBoxTextClass, ComboBoxTextOverrides](
		GTypeComboBoxText,
		initComboBoxTextClass,
		wrapComboBoxText,
		defaultComboBoxTextOverrides,
	)
}

func initComboBoxTextClass(gclass unsafe.Pointer, overrides ComboBoxTextOverrides, classInitFunc func(*ComboBoxTextClass)) {
	if classInitFunc != nil {
		class := (*ComboBoxTextClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapComboBoxText(obj *coreglib.Object) *ComboBoxText {
	return &ComboBoxText{
		ComboBox: ComboBox{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: coreglib.InitiallyUnowned{
							Object: obj,
						},
						Object: obj,
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
					},
				},
			},
			Object: obj,
			CellEditable: CellEditable{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			CellLayout: CellLayout{
				Object: obj,
			},
		},
	}
}

func marshalComboBoxText(p uintptr) (interface{}, error) {
	return wrapComboBoxText(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewComboBoxText creates a new ComboBoxText, which is a ComboBox just
// displaying strings.
//
// The function returns the following values:
//
//    - comboBoxText: new ComboBoxText.
//
func NewComboBoxText() *ComboBoxText {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_text_new()

	var _comboBoxText *ComboBoxText // out

	_comboBoxText = wrapComboBoxText(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBoxText
}

// NewComboBoxTextWithEntry creates a new ComboBoxText, which is a ComboBox just
// displaying strings. The combo box created by this function has an entry.
//
// The function returns the following values:
//
//    - comboBoxText: new ComboBoxText.
//
func NewComboBoxTextWithEntry() *ComboBoxText {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_text_new_with_entry()

	var _comboBoxText *ComboBoxText // out

	_comboBoxText = wrapComboBoxText(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBoxText
}

// Append appends text to the list of strings stored in combo_box. If id is
// non-NULL then it is used as the ID of the row.
//
// This is the same as calling gtk_combo_box_text_insert() with a position of
// -1.
//
// The function takes the following parameters:
//
//    - id (optional): string ID for this value, or NULL.
//    - text: string.
//
func (comboBox *ComboBoxText) Append(id, text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 *C.gchar           // out
	var _arg2 *C.gchar           // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if id != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(id)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_combo_box_text_append(_arg0, _arg1, _arg2)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(id)
	runtime.KeepAlive(text)
}

// AppendText appends text to the list of strings stored in combo_box.
//
// This is the same as calling gtk_combo_box_text_insert_text() with a position
// of -1.
//
// The function takes the following parameters:
//
//    - text: string.
//
func (comboBox *ComboBoxText) AppendText(text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_combo_box_text_append_text(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(text)
}

// ActiveText returns the currently active string in combo_box, or NULL if none
// is selected. If combo_box contains an entry, this function will return its
// contents (which will not necessarily be an item from the list).
//
// The function returns the following values:
//
//    - utf8: newly allocated string containing the currently active text. Must
//      be freed with g_free().
//
func (comboBox *ComboBoxText) ActiveText() string {
	var _arg0 *C.GtkComboBoxText // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_text_get_active_text(_arg0)
	runtime.KeepAlive(comboBox)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Insert inserts text at position in the list of strings stored in combo_box.
// If id is non-NULL then it is used as the ID of the row. See
// ComboBox:id-column.
//
// If position is negative then text is appended.
//
// The function takes the following parameters:
//
//    - position: index to insert text.
//    - id (optional): string ID for this value, or NULL.
//    - text: string to display.
//
func (comboBox *ComboBoxText) Insert(position int, id, text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 C.gint             // out
	var _arg2 *C.gchar           // out
	var _arg3 *C.gchar           // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.gint(position)
	if id != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(id)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg3))

	C.gtk_combo_box_text_insert(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(position)
	runtime.KeepAlive(id)
	runtime.KeepAlive(text)
}

// InsertText inserts text at position in the list of strings stored in
// combo_box.
//
// If position is negative then text is appended.
//
// This is the same as calling gtk_combo_box_text_insert() with a NULL ID
// string.
//
// The function takes the following parameters:
//
//    - position: index to insert text.
//    - text: string.
//
func (comboBox *ComboBoxText) InsertText(position int, text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 C.gint             // out
	var _arg2 *C.gchar           // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.gint(position)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_combo_box_text_insert_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(position)
	runtime.KeepAlive(text)
}

// Prepend prepends text to the list of strings stored in combo_box. If id is
// non-NULL then it is used as the ID of the row.
//
// This is the same as calling gtk_combo_box_text_insert() with a position of 0.
//
// The function takes the following parameters:
//
//    - id (optional): string ID for this value, or NULL.
//    - text: string.
//
func (comboBox *ComboBoxText) Prepend(id, text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 *C.gchar           // out
	var _arg2 *C.gchar           // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if id != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(id)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_combo_box_text_prepend(_arg0, _arg1, _arg2)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(id)
	runtime.KeepAlive(text)
}

// PrependText prepends text to the list of strings stored in combo_box.
//
// This is the same as calling gtk_combo_box_text_insert_text() with a position
// of 0.
//
// The function takes the following parameters:
//
//    - text: string.
//
func (comboBox *ComboBoxText) PrependText(text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_combo_box_text_prepend_text(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(text)
}

// Remove removes the string at position from combo_box.
//
// The function takes the following parameters:
//
//    - position: index of the item to remove.
//
func (comboBox *ComboBoxText) Remove(position int) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 C.gint             // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.gint(position)

	C.gtk_combo_box_text_remove(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(position)
}

// RemoveAll removes all the text entries from the combo box.
func (comboBox *ComboBoxText) RemoveAll() {
	var _arg0 *C.GtkComboBoxText // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	C.gtk_combo_box_text_remove_all(_arg0)
	runtime.KeepAlive(comboBox)
}

// ComboBoxTextClass: instance of this type is always passed by reference.
type ComboBoxTextClass struct {
	*comboBoxTextClass
}

// comboBoxTextClass is the struct that's finalized.
type comboBoxTextClass struct {
	native *C.GtkComboBoxTextClass
}

func (c *ComboBoxTextClass) ParentClass() *ComboBoxClass {
	valptr := &c.native.parent_class
	var _v *ComboBoxClass // out
	_v = (*ComboBoxClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
