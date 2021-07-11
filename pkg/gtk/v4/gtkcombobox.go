// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
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
		{T: externglib.Type(C.gtk_combo_box_get_type()), F: marshalComboBoxxer},
	})
}

// ComboBoxOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ComboBoxOverrider interface {
	//
	Changed()
	//
	FormatEntryText(path string) string
}

// ComboBoxxer describes ComboBox's methods.
type ComboBoxxer interface {
	// Active returns the index of the currently active item.
	Active() int
	// ActiveID returns the ID of the active row of @combo_box.
	ActiveID() string
	// ActiveIter sets @iter to point to the currently active item.
	ActiveIter() (TreeIter, bool)
	// ButtonSensitivity returns whether the combo box sets the dropdown button
	// sensitive or not when there are no items in the model.
	ButtonSensitivity() SensitivityType
	// Child gets the child widget of @combo_box.
	Child() *Widget
	// EntryTextColumn returns the column which @combo_box is using to get the
	// strings from to display in the internal entry.
	EntryTextColumn() int
	// HasEntry returns whether the combo box has an entry.
	HasEntry() bool
	// IDColumn returns the column which @combo_box is using to get string IDs
	// for values from.
	IDColumn() int
	// Model returns the `GtkTreeModel` of @combo_box.
	Model() *TreeModel
	// PopupFixedWidth gets whether the popup uses a fixed width.
	PopupFixedWidth() bool
	// Popdown hides the menu or dropdown list of @combo_box.
	Popdown()
	// Popup pops up the menu or dropdown list of @combo_box.
	Popup()
	// PopupForDevice pops up the menu of @combo_box.
	PopupForDevice(device gdk.Devicer)
	// SetActive sets the active item of @combo_box to be the item at @index.
	SetActive(index_ int)
	// SetActiveID changes the active row of @combo_box to the one that has an
	// ID equal to @active_id.
	SetActiveID(activeId string) bool
	// SetActiveIter sets the current active item to be the one referenced by
	// @iter.
	SetActiveIter(iter *TreeIter)
	// SetChild sets the child widget of @combo_box.
	SetChild(child Widgetter)
	// SetEntryTextColumn sets the model column which @combo_box should use to
	// get strings from to be @text_column.
	SetEntryTextColumn(textColumn int)
	// SetIDColumn sets the model column which @combo_box should use to get
	// string IDs for values from.
	SetIDColumn(idColumn int)
	// SetModel sets the model used by @combo_box to be @model.
	SetModel(model TreeModeller)
	// SetPopupFixedWidth specifies whether the popup’s width should be a fixed
	// width.
	SetPopupFixedWidth(fixed bool)
}

// ComboBox: `GtkComboBox` is a widget that allows the user to choose from a
// list of valid choices.
//
// !An example GtkComboBox (combo-box.png)
//
// The `GtkComboBox` displays the selected choice; when activated, the
// `GtkComboBox` displays a popup which allows the user to make a new choice.
//
// The `GtkComboBox` uses the model-view pattern; the list of valid choices is
// specified in the form of a tree model, and the display of the choices can be
// adapted to the data in the model by using cell renderers, as you would in a
// tree view. This is possible since `GtkComboBox` implements the
// [interface@Gtk.CellLayout] interface. The tree model holding the valid
// choices is not restricted to a flat list, it can be a real tree, and the
// popup will reflect the tree structure.
//
// To allow the user to enter values not in the model, the
// [property@Gtk.ComboBox:has-entry] property allows the `GtkComboBox` to
// contain a [class@Gtk.Entry]. This entry can be accessed by calling
// [method@Gtk.ComboBox.get_child] on the combo box.
//
// For a simple list of textual choices, the model-view API of `GtkComboBox` can
// be a bit overwhelming. In this case, [class@Gtk.ComboBoxText] offers a simple
// alternative. Both `GtkComboBox` and `GtkComboBoxText` can contain an entry.
//
//
// CSS nodes
//
// “` combobox ├── box.linked │ ╰── button.combo │ ╰── box │ ├── cellview │ ╰──
// arrow ╰── window.popup “`
//
// A normal combobox contains a box with the .linked class, a button with the
// .combo class and inside those buttons, there are a cellview and an arrow.
//
// “` combobox ├── box.linked │ ├── entry.combo │ ╰── button.combo │ ╰── box │
// ╰── arrow ╰── window.popup “`
//
// A `GtkComboBox` with an entry has a single CSS node with name combobox. It
// contains a box with the .linked class. That box contains an entry and a
// button, both with the .combo class added. The button also contains another
// node with name arrow.
//
//
// Accessibility
//
// `GtkComboBox` uses the GTK_ACCESSIBLE_ROLE_COMBO_BOX role.
type ComboBox struct {
	Widget

	CellEditable
	CellLayout
}

var (
	_ ComboBoxxer     = (*ComboBox)(nil)
	_ gextras.Nativer = (*ComboBox)(nil)
)

func wrapComboBox(obj *externglib.Object) ComboBoxxer {
	return &ComboBox{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		CellEditable: CellEditable{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
		},
		CellLayout: CellLayout{
			Object: obj,
		},
	}
}

func marshalComboBoxxer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapComboBox(obj), nil
}

// NewComboBox creates a new empty `GtkComboBox`.
func NewComboBox() *ComboBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_new()

	var _comboBox *ComboBox // out

	_comboBox = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ComboBox)

	return _comboBox
}

// NewComboBoxWithEntry creates a new empty `GtkComboBox` with an entry.
func NewComboBoxWithEntry() *ComboBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_new_with_entry()

	var _comboBox *ComboBox // out

	_comboBox = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ComboBox)

	return _comboBox
}

// NewComboBoxWithModel creates a new `GtkComboBox` with a model.
func NewComboBoxWithModel(model TreeModeller) *ComboBox {
	var _arg1 *C.GtkTreeModel // out
	var _cret *C.GtkWidget    // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer((model).(gextras.Nativer).Native()))

	_cret = C.gtk_combo_box_new_with_model(_arg1)

	var _comboBox *ComboBox // out

	_comboBox = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ComboBox)

	return _comboBox
}

// NewComboBoxWithModelAndEntry creates a new empty `GtkComboBox` with an entry
// and a model.
func NewComboBoxWithModelAndEntry(model TreeModeller) *ComboBox {
	var _arg1 *C.GtkTreeModel // out
	var _cret *C.GtkWidget    // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer((model).(gextras.Nativer).Native()))

	_cret = C.gtk_combo_box_new_with_model_and_entry(_arg1)

	var _comboBox *ComboBox // out

	_comboBox = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ComboBox)

	return _comboBox
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *ComboBox) Native() uintptr {
	return v.Widget.InitiallyUnowned.Object.Native()
}

// Active returns the index of the currently active item.
//
// If the model is a non-flat treemodel, and the active item is not an immediate
// child of the root of the tree, this function returns
// `gtk_tree_path_get_indices (path)[0]`, where `path` is the
// [struct@Gtk.TreePath] of the active item.
func (comboBox *ComboBox) Active() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.int          // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_active(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ActiveID returns the ID of the active row of @combo_box.
//
// This value is taken from the active row and the column specified by the
// [property@Gtk.ComboBox:id-column] property of @combo_box (see
// [method@Gtk.ComboBox.set_id_column]).
//
// The returned value is an interned string which means that you can compare the
// pointer by value to other interned strings and that you must not free it.
//
// If the [property@Gtk.ComboBox:id-column] property of @combo_box is not set,
// or if no row is active, or if the active row has a nil ID value, then nil is
// returned.
func (comboBox *ComboBox) ActiveID() string {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.char        // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_active_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ActiveIter sets @iter to point to the currently active item.
//
// If no item is active, @iter is left unchanged.
func (comboBox *ComboBox) ActiveIter() (TreeIter, bool) {
	var _arg0 *C.GtkComboBox // out
	var _iter TreeIter
	var _cret C.gboolean // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_active_iter(_arg0, (*C.GtkTreeIter)(unsafe.Pointer(&_iter)))

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _iter, _ok
}

// ButtonSensitivity returns whether the combo box sets the dropdown button
// sensitive or not when there are no items in the model.
func (comboBox *ComboBox) ButtonSensitivity() SensitivityType {
	var _arg0 *C.GtkComboBox       // out
	var _cret C.GtkSensitivityType // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_button_sensitivity(_arg0)

	var _sensitivityType SensitivityType // out

	_sensitivityType = SensitivityType(_cret)

	return _sensitivityType
}

// Child gets the child widget of @combo_box.
func (comboBox *ComboBox) Child() *Widget {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_child(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// EntryTextColumn returns the column which @combo_box is using to get the
// strings from to display in the internal entry.
func (comboBox *ComboBox) EntryTextColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.int          // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_entry_text_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// HasEntry returns whether the combo box has an entry.
func (comboBox *ComboBox) HasEntry() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_has_entry(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IDColumn returns the column which @combo_box is using to get string IDs for
// values from.
func (comboBox *ComboBox) IDColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.int          // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_id_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Model returns the `GtkTreeModel` of @combo_box.
func (comboBox *ComboBox) Model() *TreeModel {
	var _arg0 *C.GtkComboBox  // out
	var _cret *C.GtkTreeModel // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_model(_arg0)

	var _treeModel *TreeModel // out

	_treeModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*TreeModel)

	return _treeModel
}

// PopupFixedWidth gets whether the popup uses a fixed width.
func (comboBox *ComboBox) PopupFixedWidth() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_popup_fixed_width(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Popdown hides the menu or dropdown list of @combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
func (comboBox *ComboBox) Popdown() {
	var _arg0 *C.GtkComboBox // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	C.gtk_combo_box_popdown(_arg0)
}

// Popup pops up the menu or dropdown list of @combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
//
// Before calling this, @combo_box must be mapped, or nothing will happen.
func (comboBox *ComboBox) Popup() {
	var _arg0 *C.GtkComboBox // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	C.gtk_combo_box_popup(_arg0)
}

// PopupForDevice pops up the menu of @combo_box.
//
// Note that currently this does not do anything with the device, as it was
// previously only used for list-mode combo boxes, and those were removed in GTK
// 4. However, it is retained in case similar functionality is added back later.
func (comboBox *ComboBox) PopupForDevice(device gdk.Devicer) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GdkDevice   // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer((device).(gextras.Nativer).Native()))

	C.gtk_combo_box_popup_for_device(_arg0, _arg1)
}

// SetActive sets the active item of @combo_box to be the item at @index.
func (comboBox *ComboBox) SetActive(index_ int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.int(index_)

	C.gtk_combo_box_set_active(_arg0, _arg1)
}

// SetActiveID changes the active row of @combo_box to the one that has an ID
// equal to @active_id.
//
// If @active_id is nil, the active row is unset. Rows having a nil ID string
// cannot be made active by this function.
//
// If the [property@Gtk.ComboBox:id-column] property of @combo_box is unset or
// if no row has the given ID then the function does nothing and returns false.
func (comboBox *ComboBox) SetActiveID(activeId string) bool {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.char        // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(activeId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_combo_box_set_active_id(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActiveIter sets the current active item to be the one referenced by @iter.
//
// If @iter is nil, the active item is unset.
func (comboBox *ComboBox) SetActiveIter(iter *TreeIter) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GtkTreeIter // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.gtk_combo_box_set_active_iter(_arg0, _arg1)
}

// SetChild sets the child widget of @combo_box.
func (comboBox *ComboBox) SetChild(child Widgetter) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_combo_box_set_child(_arg0, _arg1)
}

// SetEntryTextColumn sets the model column which @combo_box should use to get
// strings from to be @text_column.
//
// The column @text_column in the model of @combo_box must be of type
// G_TYPE_STRING.
//
// This is only relevant if @combo_box has been created with
// [property@Gtk.ComboBox:has-entry] as true.
func (comboBox *ComboBox) SetEntryTextColumn(textColumn int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.int(textColumn)

	C.gtk_combo_box_set_entry_text_column(_arg0, _arg1)
}

// SetIDColumn sets the model column which @combo_box should use to get string
// IDs for values from.
//
// The column @id_column in the model of @combo_box must be of type
// G_TYPE_STRING.
func (comboBox *ComboBox) SetIDColumn(idColumn int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.int(idColumn)

	C.gtk_combo_box_set_id_column(_arg0, _arg1)
}

// SetModel sets the model used by @combo_box to be @model.
//
// Will unset a previously set model (if applicable). If model is nil, then it
// will unset the model.
//
// Note that this function does not clear the cell renderers, you have to call
// [method@Gtk.CellLayout.clear] yourself if you need to set up different cell
// renderers for the new model.
func (comboBox *ComboBox) SetModel(model TreeModeller) {
	var _arg0 *C.GtkComboBox  // out
	var _arg1 *C.GtkTreeModel // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer((model).(gextras.Nativer).Native()))

	C.gtk_combo_box_set_model(_arg0, _arg1)
}

// SetPopupFixedWidth specifies whether the popup’s width should be a fixed
// width.
//
// If @fixed is true, the popup's width is set to match the allocated width of
// the combo box.
func (comboBox *ComboBox) SetPopupFixedWidth(fixed bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	if fixed {
		_arg1 = C.TRUE
	}

	C.gtk_combo_box_set_popup_fixed_width(_arg0, _arg1)
}
