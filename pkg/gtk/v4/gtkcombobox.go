// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_combo_box_get_type()), F: marshalComboBox},
	})
}

// ComboBox: a GtkComboBox is a widget that allows the user to choose from a
// list of valid choices. The GtkComboBox displays the selected choice. When
// activated, the GtkComboBox displays a popup which allows the user to make a
// new choice. The style in which the selected value is displayed, and the style
// of the popup is determined by the current theme. It may be similar to a
// Windows-style combo box.
//
// The GtkComboBox uses the model-view pattern; the list of valid choices is
// specified in the form of a tree model, and the display of the choices can be
// adapted to the data in the model by using cell renderers, as you would in a
// tree view. This is possible since GtkComboBox implements the CellLayout
// interface. The tree model holding the valid choices is not restricted to a
// flat list, it can be a real tree, and the popup will reflect the tree
// structure.
//
// To allow the user to enter values not in the model, the “has-entry” property
// allows the GtkComboBox to contain a Entry. This entry can be accessed by
// calling gtk_combo_box_get_child() on the combo box.
//
// For a simple list of textual choices, the model-view API of GtkComboBox can
// be a bit overwhelming. In this case, ComboBoxText offers a simple
// alternative. Both GtkComboBox and ComboBoxText can contain an entry.
//
// CSS nodes
//
//    combobox
//    ├── box.linked
//    │   ├── entry.combo
//    │   ╰── button.combo
//    │       ╰── box
//    │           ╰── arrow
//    ╰── window.popup
//
// A GtkComboBox with an entry has a single CSS node with name combobox. It
// contains a box with the .linked class. That box contains an entry and a
// button, both with the .combo class added. The button also contains another
// node with name arrow.
//
//
// Accessibility
//
// GtkComboBox uses the K_ACCESSIBLE_ROLE_COMBO_BOX role.
type ComboBox interface {
	Widget
	Accessible
	Buildable
	CellEditable
	CellLayout
	ConstraintTarget

	// Active returns the index of the currently active item, or -1 if there’s
	// no active item. If the model is a non-flat treemodel, and the active item
	// is not an immediate child of the root of the tree, this function returns
	// `gtk_tree_path_get_indices (path)[0]`, where `path` is the TreePath of
	// the active item.
	Active() int
	// ActiveID returns the ID of the active row of @combo_box. This value is
	// taken from the active row and the column specified by the
	// ComboBox:id-column property of @combo_box (see
	// gtk_combo_box_set_id_column()).
	//
	// The returned value is an interned string which means that you can compare
	// the pointer by value to other interned strings and that you must not free
	// it.
	//
	// If the ComboBox:id-column property of @combo_box is not set, or if no row
	// is active, or if the active row has a nil ID value, then nil is returned.
	ActiveID() string
	// ActiveIter sets @iter to point to the currently active item, if any item
	// is active. Otherwise, @iter is left unchanged.
	ActiveIter() (iter TreeIter, ok bool)
	// ButtonSensitivity returns whether the combo box sets the dropdown button
	// sensitive or not when there are no items in the model.
	ButtonSensitivity() SensitivityType
	// Child gets the child widget of @combo_box.
	Child() Widget
	// EntryTextColumn returns the column which @combo_box is using to get the
	// strings from to display in the internal entry.
	EntryTextColumn() int
	// HasEntry returns whether the combo box has an entry.
	HasEntry() bool
	// IDColumn returns the column which @combo_box is using to get string IDs
	// for values from.
	IDColumn() int
	// Model returns the TreeModel which is acting as data source for
	// @combo_box.
	Model() TreeModel
	// PopupFixedWidth gets whether the popup uses a fixed width matching the
	// allocated width of the combo box.
	PopupFixedWidth() bool
	// Popdown hides the menu or dropdown list of @combo_box.
	//
	// This function is mostly intended for use by accessibility technologies;
	// applications should have little use for it.
	Popdown()
	// Popup pops up the menu or dropdown list of @combo_box.
	//
	// This function is mostly intended for use by accessibility technologies;
	// applications should have little use for it.
	//
	// Before calling this, @combo_box must be mapped, or nothing will happen.
	Popup()
	// PopupForDevice pops up the menu of @combo_box. Note that currently this
	// does not do anything with the device, as it was previously only used for
	// list-mode combo boxes, and those were removed in GTK 4. However, it is
	// retained in case similar functionality is added back later.
	PopupForDevice(device gdk.Device)
	// SetActive sets the active item of @combo_box to be the item at @index.
	SetActive(index_ int)
	// SetActiveID changes the active row of @combo_box to the one that has an
	// ID equal to @active_id, or unsets the active row if @active_id is nil.
	// Rows having a nil ID string cannot be made active by this function.
	//
	// If the ComboBox:id-column property of @combo_box is unset or if no row
	// has the given ID then the function does nothing and returns false.
	SetActiveID(activeID string) bool
	// SetActiveIter sets the current active item to be the one referenced by
	// @iter, or unsets the active item if @iter is nil.
	SetActiveIter(iter *TreeIter)
	// SetButtonSensitivity sets whether the dropdown button of the combo box
	// should be always sensitive (GTK_SENSITIVITY_ON), never sensitive
	// (GTK_SENSITIVITY_OFF) or only if there is at least one item to display
	// (GTK_SENSITIVITY_AUTO).
	SetButtonSensitivity(sensitivity SensitivityType)
	// SetChild sets the child widget of @combo_box.
	SetChild(child Widget)
	// SetEntryTextColumn sets the model column which @combo_box should use to
	// get strings from to be @text_column. The column @text_column in the model
	// of @combo_box must be of type G_TYPE_STRING.
	//
	// This is only relevant if @combo_box has been created with
	// ComboBox:has-entry as true.
	SetEntryTextColumn(textColumn int)
	// SetIDColumn sets the model column which @combo_box should use to get
	// string IDs for values from. The column @id_column in the model of
	// @combo_box must be of type G_TYPE_STRING.
	SetIDColumn(idColumn int)
	// SetModel sets the model used by @combo_box to be @model. Will unset a
	// previously set model (if applicable). If model is nil, then it will unset
	// the model.
	//
	// Note that this function does not clear the cell renderers, you have to
	// call gtk_cell_layout_clear() yourself if you need to set up different
	// cell renderers for the new model.
	SetModel(model TreeModel)
	// SetPopupFixedWidth specifies whether the popup’s width should be a fixed
	// width matching the allocated width of the combo box.
	SetPopupFixedWidth(fixed bool)
	// SetRowSeparatorFunc sets the row separator function, which is used to
	// determine whether a row should be drawn as a separator. If the row
	// separator function is nil, no separators are drawn. This is the default
	// value.
	SetRowSeparatorFunc(fn TreeViewRowSeparatorFunc)
}

// comboBox implements the ComboBox interface.
type comboBox struct {
	Widget
	Accessible
	Buildable
	CellEditable
	CellLayout
	ConstraintTarget
}

var _ ComboBox = (*comboBox)(nil)

// WrapComboBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapComboBox(obj *externglib.Object) ComboBox {
	return ComboBox{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		CellEditable:     WrapCellEditable(obj),
		CellLayout:       WrapCellLayout(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalComboBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapComboBox(obj), nil
}

// NewComboBox constructs a class ComboBox.
func NewComboBox() ComboBox {
	var cret C.GtkComboBox
	var ret1 ComboBox

	cret = C.gtk_combo_box_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ComboBox)

	return ret1
}

// NewComboBoxWithEntry constructs a class ComboBox.
func NewComboBoxWithEntry() ComboBox {
	var cret C.GtkComboBox
	var ret1 ComboBox

	cret = C.gtk_combo_box_new_with_entry()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ComboBox)

	return ret1
}

// NewComboBoxWithModel constructs a class ComboBox.
func NewComboBoxWithModel(model TreeModel) ComboBox {
	var arg1 *C.GtkTreeModel

	arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	var cret C.GtkComboBox
	var ret1 ComboBox

	cret = C.gtk_combo_box_new_with_model(model)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ComboBox)

	return ret1
}

// NewComboBoxWithModelAndEntry constructs a class ComboBox.
func NewComboBoxWithModelAndEntry(model TreeModel) ComboBox {
	var arg1 *C.GtkTreeModel

	arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	var cret C.GtkComboBox
	var ret1 ComboBox

	cret = C.gtk_combo_box_new_with_model_and_entry(model)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ComboBox)

	return ret1
}

// Active returns the index of the currently active item, or -1 if there’s
// no active item. If the model is a non-flat treemodel, and the active item
// is not an immediate child of the root of the tree, this function returns
// `gtk_tree_path_get_indices (path)[0]`, where `path` is the TreePath of
// the active item.
func (c comboBox) Active() int {
	var arg0 *C.GtkComboBox

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	var cret C.int
	var ret1 int

	cret = C.gtk_combo_box_get_active(arg0)

	ret1 = C.int(cret)

	return ret1
}

// ActiveID returns the ID of the active row of @combo_box. This value is
// taken from the active row and the column specified by the
// ComboBox:id-column property of @combo_box (see
// gtk_combo_box_set_id_column()).
//
// The returned value is an interned string which means that you can compare
// the pointer by value to other interned strings and that you must not free
// it.
//
// If the ComboBox:id-column property of @combo_box is not set, or if no row
// is active, or if the active row has a nil ID value, then nil is returned.
func (c comboBox) ActiveID() string {
	var arg0 *C.GtkComboBox

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	var cret *C.char
	var ret1 string

	cret = C.gtk_combo_box_get_active_id(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// ActiveIter sets @iter to point to the currently active item, if any item
// is active. Otherwise, @iter is left unchanged.
func (c comboBox) ActiveIter() (iter TreeIter, ok bool) {
	var arg0 *C.GtkComboBox

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	var arg1 C.GtkTreeIter
	var ret1 *TreeIter
	var cret C.gboolean
	var ret2 bool

	cret = C.gtk_combo_box_get_active_iter(arg0, &arg1)

	*ret1 = WrapTreeIter(unsafe.Pointer(arg1))
	ret2 = C.bool(cret) != C.false

	return ret1, ret2
}

// ButtonSensitivity returns whether the combo box sets the dropdown button
// sensitive or not when there are no items in the model.
func (c comboBox) ButtonSensitivity() SensitivityType {
	var arg0 *C.GtkComboBox

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	var cret C.GtkSensitivityType
	var ret1 SensitivityType

	cret = C.gtk_combo_box_get_button_sensitivity(arg0)

	ret1 = SensitivityType(cret)

	return ret1
}

// Child gets the child widget of @combo_box.
func (c comboBox) Child() Widget {
	var arg0 *C.GtkComboBox

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_combo_box_get_child(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// EntryTextColumn returns the column which @combo_box is using to get the
// strings from to display in the internal entry.
func (c comboBox) EntryTextColumn() int {
	var arg0 *C.GtkComboBox

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	var cret C.int
	var ret1 int

	cret = C.gtk_combo_box_get_entry_text_column(arg0)

	ret1 = C.int(cret)

	return ret1
}

// HasEntry returns whether the combo box has an entry.
func (c comboBox) HasEntry() bool {
	var arg0 *C.GtkComboBox

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_combo_box_get_has_entry(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// IDColumn returns the column which @combo_box is using to get string IDs
// for values from.
func (c comboBox) IDColumn() int {
	var arg0 *C.GtkComboBox

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	var cret C.int
	var ret1 int

	cret = C.gtk_combo_box_get_id_column(arg0)

	ret1 = C.int(cret)

	return ret1
}

// Model returns the TreeModel which is acting as data source for
// @combo_box.
func (c comboBox) Model() TreeModel {
	var arg0 *C.GtkComboBox

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	var cret *C.GtkTreeModel
	var ret1 TreeModel

	cret = C.gtk_combo_box_get_model(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(TreeModel)

	return ret1
}

// PopupFixedWidth gets whether the popup uses a fixed width matching the
// allocated width of the combo box.
func (c comboBox) PopupFixedWidth() bool {
	var arg0 *C.GtkComboBox

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_combo_box_get_popup_fixed_width(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Popdown hides the menu or dropdown list of @combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
func (c comboBox) Popdown() {
	var arg0 *C.GtkComboBox

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	C.gtk_combo_box_popdown(arg0)
}

// Popup pops up the menu or dropdown list of @combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
//
// Before calling this, @combo_box must be mapped, or nothing will happen.
func (c comboBox) Popup() {
	var arg0 *C.GtkComboBox

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	C.gtk_combo_box_popup(arg0)
}

// PopupForDevice pops up the menu of @combo_box. Note that currently this
// does not do anything with the device, as it was previously only used for
// list-mode combo boxes, and those were removed in GTK 4. However, it is
// retained in case similar functionality is added back later.
func (c comboBox) PopupForDevice(device gdk.Device) {
	var arg0 *C.GtkComboBox
	var arg1 *C.GdkDevice

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	C.gtk_combo_box_popup_for_device(arg0, device)
}

// SetActive sets the active item of @combo_box to be the item at @index.
func (c comboBox) SetActive(index_ int) {
	var arg0 *C.GtkComboBox
	var arg1 C.int

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	arg1 = C.int(index_)

	C.gtk_combo_box_set_active(arg0, index_)
}

// SetActiveID changes the active row of @combo_box to the one that has an
// ID equal to @active_id, or unsets the active row if @active_id is nil.
// Rows having a nil ID string cannot be made active by this function.
//
// If the ComboBox:id-column property of @combo_box is unset or if no row
// has the given ID then the function does nothing and returns false.
func (c comboBox) SetActiveID(activeID string) bool {
	var arg0 *C.GtkComboBox
	var arg1 *C.char

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	arg1 = (*C.char)(C.CString(activeID))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_combo_box_set_active_id(arg0, activeID)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// SetActiveIter sets the current active item to be the one referenced by
// @iter, or unsets the active item if @iter is nil.
func (c comboBox) SetActiveIter(iter *TreeIter) {
	var arg0 *C.GtkComboBox
	var arg1 *C.GtkTreeIter

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	C.gtk_combo_box_set_active_iter(arg0, iter)
}

// SetButtonSensitivity sets whether the dropdown button of the combo box
// should be always sensitive (GTK_SENSITIVITY_ON), never sensitive
// (GTK_SENSITIVITY_OFF) or only if there is at least one item to display
// (GTK_SENSITIVITY_AUTO).
func (c comboBox) SetButtonSensitivity(sensitivity SensitivityType) {
	var arg0 *C.GtkComboBox
	var arg1 C.GtkSensitivityType

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	arg1 = (C.GtkSensitivityType)(sensitivity)

	C.gtk_combo_box_set_button_sensitivity(arg0, sensitivity)
}

// SetChild sets the child widget of @combo_box.
func (c comboBox) SetChild(child Widget) {
	var arg0 *C.GtkComboBox
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_combo_box_set_child(arg0, child)
}

// SetEntryTextColumn sets the model column which @combo_box should use to
// get strings from to be @text_column. The column @text_column in the model
// of @combo_box must be of type G_TYPE_STRING.
//
// This is only relevant if @combo_box has been created with
// ComboBox:has-entry as true.
func (c comboBox) SetEntryTextColumn(textColumn int) {
	var arg0 *C.GtkComboBox
	var arg1 C.int

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	arg1 = C.int(textColumn)

	C.gtk_combo_box_set_entry_text_column(arg0, textColumn)
}

// SetIDColumn sets the model column which @combo_box should use to get
// string IDs for values from. The column @id_column in the model of
// @combo_box must be of type G_TYPE_STRING.
func (c comboBox) SetIDColumn(idColumn int) {
	var arg0 *C.GtkComboBox
	var arg1 C.int

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	arg1 = C.int(idColumn)

	C.gtk_combo_box_set_id_column(arg0, idColumn)
}

// SetModel sets the model used by @combo_box to be @model. Will unset a
// previously set model (if applicable). If model is nil, then it will unset
// the model.
//
// Note that this function does not clear the cell renderers, you have to
// call gtk_cell_layout_clear() yourself if you need to set up different
// cell renderers for the new model.
func (c comboBox) SetModel(model TreeModel) {
	var arg0 *C.GtkComboBox
	var arg1 *C.GtkTreeModel

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	C.gtk_combo_box_set_model(arg0, model)
}

// SetPopupFixedWidth specifies whether the popup’s width should be a fixed
// width matching the allocated width of the combo box.
func (c comboBox) SetPopupFixedWidth(fixed bool) {
	var arg0 *C.GtkComboBox
	var arg1 C.gboolean

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	if fixed {
		arg1 = C.gboolean(1)
	}

	C.gtk_combo_box_set_popup_fixed_width(arg0, fixed)
}

// SetRowSeparatorFunc sets the row separator function, which is used to
// determine whether a row should be drawn as a separator. If the row
// separator function is nil, no separators are drawn. This is the default
// value.
func (c comboBox) SetRowSeparatorFunc(fn TreeViewRowSeparatorFunc) {
	var arg0 *C.GtkComboBox

	arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	C.gtk_combo_box_set_row_separator_func(arg0, fn, data, destroy)
}
