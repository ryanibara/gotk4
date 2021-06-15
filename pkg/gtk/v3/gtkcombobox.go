// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
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
// calling gtk_bin_get_child() on the combo box.
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
type ComboBox interface {
	Bin
	Buildable
	CellEditable
	CellLayout

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
	ActiveIter() (TreeIter, bool)
	// AddTearoffs gets the current value of the :add-tearoffs property.
	AddTearoffs() bool
	// ButtonSensitivity returns whether the combo box sets the dropdown button
	// sensitive or not when there are no items in the model.
	ButtonSensitivity() SensitivityType
	// ColumnSpanColumn returns the column with column span information for
	// @combo_box.
	ColumnSpanColumn() int
	// EntryTextColumn returns the column which @combo_box is using to get the
	// strings from to display in the internal entry.
	EntryTextColumn() int
	// FocusOnClick returns whether the combo box grabs focus when it is clicked
	// with the mouse. See gtk_combo_box_set_focus_on_click().
	FocusOnClick() bool
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
	// RowSpanColumn returns the column with row span information for
	// @combo_box.
	RowSpanColumn() int
	// Title gets the current title of the menu in tearoff mode. See
	// gtk_combo_box_set_add_tearoffs().
	Title() string
	// WrapWidth returns the wrap width which is used to determine the number of
	// columns for the popup menu. If the wrap width is larger than 1, the combo
	// box is in table mode.
	WrapWidth() int
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
	// PopupForDevice pops up the menu or dropdown list of @combo_box, the popup
	// window will be grabbed so only @device and its associated
	// pointer/keyboard are the only Devices able to send events to it.
	PopupForDevice(device gdk.Device)
	// SetActive sets the active item of @combo_box to be the item at @index.
	SetActive(index_ int)
	// SetActiveID changes the active row of @combo_box to the one that has an
	// ID equal to @active_id, or unsets the active row if @active_id is nil.
	// Rows having a nil ID string cannot be made active by this function.
	//
	// If the ComboBox:id-column property of @combo_box is unset or if no row
	// has the given ID then the function does nothing and returns false.
	SetActiveID(activeId string) bool
	// SetActiveIter sets the current active item to be the one referenced by
	// @iter, or unsets the active item if @iter is nil.
	SetActiveIter(iter *TreeIter)
	// SetAddTearoffs sets whether the popup menu should have a tearoff menu
	// item.
	SetAddTearoffs(addTearoffs bool)
	// SetButtonSensitivity sets whether the dropdown button of the combo box
	// should be always sensitive (GTK_SENSITIVITY_ON), never sensitive
	// (GTK_SENSITIVITY_OFF) or only if there is at least one item to display
	// (GTK_SENSITIVITY_AUTO).
	SetButtonSensitivity(sensitivity SensitivityType)
	// SetColumnSpanColumn sets the column with column span information for
	// @combo_box to be @column_span. The column span column contains integers
	// which indicate how many columns an item should span.
	SetColumnSpanColumn(columnSpan int)
	// SetEntryTextColumn sets the model column which @combo_box should use to
	// get strings from to be @text_column. The column @text_column in the model
	// of @combo_box must be of type G_TYPE_STRING.
	//
	// This is only relevant if @combo_box has been created with
	// ComboBox:has-entry as true.
	SetEntryTextColumn(textColumn int)
	// SetFocusOnClick sets whether the combo box will grab focus when it is
	// clicked with the mouse. Making mouse clicks not grab focus is useful in
	// places like toolbars where you don’t want the keyboard focus removed from
	// the main area of the application.
	SetFocusOnClick(focusOnClick bool)
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
	// SetRowSpanColumn sets the column with row span information for @combo_box
	// to be @row_span. The row span column contains integers which indicate how
	// many rows an item should span.
	SetRowSpanColumn(rowSpan int)
	// SetTitle sets the menu’s title in tearoff mode.
	SetTitle(title string)
	// SetWrapWidth sets the wrap width of @combo_box to be @width. The wrap
	// width is basically the preferred number of columns when you want the
	// popup to be layed out in a table.
	SetWrapWidth(width int)
}

// comboBox implements the ComboBox class.
type comboBox struct {
	Bin
	Buildable
	CellEditable
	CellLayout
}

var _ ComboBox = (*comboBox)(nil)

// WrapComboBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapComboBox(obj *externglib.Object) ComboBox {
	return comboBox{
		Bin:          WrapBin(obj),
		Buildable:    WrapBuildable(obj),
		CellEditable: WrapCellEditable(obj),
		CellLayout:   WrapCellLayout(obj),
	}
}

func marshalComboBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapComboBox(obj), nil
}

// NewComboBox constructs a class ComboBox.
func NewComboBox() ComboBox {
	var _cret C.GtkComboBox // in

	_cret = C.gtk_combo_box_new()

	var _comboBox ComboBox // out

	_comboBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ComboBox)

	return _comboBox
}

// NewComboBoxWithArea constructs a class ComboBox.
func NewComboBoxWithArea(area CellArea) ComboBox {
	var _arg1 *C.GtkCellArea // out
	var _cret C.GtkComboBox  // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_combo_box_new_with_area(_arg1)

	var _comboBox ComboBox // out

	_comboBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ComboBox)

	return _comboBox
}

// NewComboBoxWithAreaAndEntry constructs a class ComboBox.
func NewComboBoxWithAreaAndEntry(area CellArea) ComboBox {
	var _arg1 *C.GtkCellArea // out
	var _cret C.GtkComboBox  // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_combo_box_new_with_area_and_entry(_arg1)

	var _comboBox ComboBox // out

	_comboBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ComboBox)

	return _comboBox
}

// NewComboBoxWithEntry constructs a class ComboBox.
func NewComboBoxWithEntry() ComboBox {
	var _cret C.GtkComboBox // in

	_cret = C.gtk_combo_box_new_with_entry()

	var _comboBox ComboBox // out

	_comboBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ComboBox)

	return _comboBox
}

// NewComboBoxWithModel constructs a class ComboBox.
func NewComboBoxWithModel(model TreeModel) ComboBox {
	var _arg1 *C.GtkTreeModel // out
	var _cret C.GtkComboBox   // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_combo_box_new_with_model(_arg1)

	var _comboBox ComboBox // out

	_comboBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ComboBox)

	return _comboBox
}

// NewComboBoxWithModelAndEntry constructs a class ComboBox.
func NewComboBoxWithModelAndEntry(model TreeModel) ComboBox {
	var _arg1 *C.GtkTreeModel // out
	var _cret C.GtkComboBox   // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_combo_box_new_with_model_and_entry(_arg1)

	var _comboBox ComboBox // out

	_comboBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ComboBox)

	return _comboBox
}

// Active returns the index of the currently active item, or -1 if there’s
// no active item. If the model is a non-flat treemodel, and the active item
// is not an immediate child of the root of the tree, this function returns
// `gtk_tree_path_get_indices (path)[0]`, where `path` is the TreePath of
// the active item.
func (c comboBox) Active() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_active(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
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
	var _arg0 *C.GtkComboBox // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_active_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// ActiveIter sets @iter to point to the currently active item, if any item
// is active. Otherwise, @iter is left unchanged.
func (c comboBox) ActiveIter() (TreeIter, bool) {
	var _arg0 *C.GtkComboBox // out
	var _iter TreeIter
	var _cret C.gboolean // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_active_iter(_arg0, (*C.GtkTreeIter)(unsafe.Pointer(&_iter)))

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _iter, _ok
}

// AddTearoffs gets the current value of the :add-tearoffs property.
func (c comboBox) AddTearoffs() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_add_tearoffs(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ButtonSensitivity returns whether the combo box sets the dropdown button
// sensitive or not when there are no items in the model.
func (c comboBox) ButtonSensitivity() SensitivityType {
	var _arg0 *C.GtkComboBox       // out
	var _cret C.GtkSensitivityType // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_button_sensitivity(_arg0)

	var _sensitivityType SensitivityType // out

	_sensitivityType = SensitivityType(_cret)

	return _sensitivityType
}

// ColumnSpanColumn returns the column with column span information for
// @combo_box.
func (c comboBox) ColumnSpanColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_column_span_column(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// EntryTextColumn returns the column which @combo_box is using to get the
// strings from to display in the internal entry.
func (c comboBox) EntryTextColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_entry_text_column(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// FocusOnClick returns whether the combo box grabs focus when it is clicked
// with the mouse. See gtk_combo_box_set_focus_on_click().
func (c comboBox) FocusOnClick() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_focus_on_click(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasEntry returns whether the combo box has an entry.
func (c comboBox) HasEntry() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_has_entry(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IDColumn returns the column which @combo_box is using to get string IDs
// for values from.
func (c comboBox) IDColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_id_column(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Model returns the TreeModel which is acting as data source for
// @combo_box.
func (c comboBox) Model() TreeModel {
	var _arg0 *C.GtkComboBox  // out
	var _cret *C.GtkTreeModel // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_model(_arg0)

	var _treeModel TreeModel // out

	_treeModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(TreeModel)

	return _treeModel
}

// PopupFixedWidth gets whether the popup uses a fixed width matching the
// allocated width of the combo box.
func (c comboBox) PopupFixedWidth() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_popup_fixed_width(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowSpanColumn returns the column with row span information for
// @combo_box.
func (c comboBox) RowSpanColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_row_span_column(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Title gets the current title of the menu in tearoff mode. See
// gtk_combo_box_set_add_tearoffs().
func (c comboBox) Title() string {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// WrapWidth returns the wrap width which is used to determine the number of
// columns for the popup menu. If the wrap width is larger than 1, the combo
// box is in table mode.
func (c comboBox) WrapWidth() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_wrap_width(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Popdown hides the menu or dropdown list of @combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
func (c comboBox) Popdown() {
	var _arg0 *C.GtkComboBox // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	C.gtk_combo_box_popdown(_arg0)
}

// Popup pops up the menu or dropdown list of @combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
//
// Before calling this, @combo_box must be mapped, or nothing will happen.
func (c comboBox) Popup() {
	var _arg0 *C.GtkComboBox // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	C.gtk_combo_box_popup(_arg0)
}

// PopupForDevice pops up the menu or dropdown list of @combo_box, the popup
// window will be grabbed so only @device and its associated
// pointer/keyboard are the only Devices able to send events to it.
func (c comboBox) PopupForDevice(device gdk.Device) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GdkDevice   // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	C.gtk_combo_box_popup_for_device(_arg0, _arg1)
}

// SetActive sets the active item of @combo_box to be the item at @index.
func (c comboBox) SetActive(index_ int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (C.gint)(index_)

	C.gtk_combo_box_set_active(_arg0, _arg1)
}

// SetActiveID changes the active row of @combo_box to the one that has an
// ID equal to @active_id, or unsets the active row if @active_id is nil.
// Rows having a nil ID string cannot be made active by this function.
//
// If the ComboBox:id-column property of @combo_box is unset or if no row
// has the given ID then the function does nothing and returns false.
func (c comboBox) SetActiveID(activeId string) bool {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.gchar       // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(activeId))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_combo_box_set_active_id(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActiveIter sets the current active item to be the one referenced by
// @iter, or unsets the active item if @iter is nil.
func (c comboBox) SetActiveIter(iter *TreeIter) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GtkTreeIter // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	C.gtk_combo_box_set_active_iter(_arg0, _arg1)
}

// SetAddTearoffs sets whether the popup menu should have a tearoff menu
// item.
func (c comboBox) SetAddTearoffs(addTearoffs bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	if addTearoffs {
		_arg1 = C.TRUE
	}

	C.gtk_combo_box_set_add_tearoffs(_arg0, _arg1)
}

// SetButtonSensitivity sets whether the dropdown button of the combo box
// should be always sensitive (GTK_SENSITIVITY_ON), never sensitive
// (GTK_SENSITIVITY_OFF) or only if there is at least one item to display
// (GTK_SENSITIVITY_AUTO).
func (c comboBox) SetButtonSensitivity(sensitivity SensitivityType) {
	var _arg0 *C.GtkComboBox       // out
	var _arg1 C.GtkSensitivityType // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (C.GtkSensitivityType)(sensitivity)

	C.gtk_combo_box_set_button_sensitivity(_arg0, _arg1)
}

// SetColumnSpanColumn sets the column with column span information for
// @combo_box to be @column_span. The column span column contains integers
// which indicate how many columns an item should span.
func (c comboBox) SetColumnSpanColumn(columnSpan int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (C.gint)(columnSpan)

	C.gtk_combo_box_set_column_span_column(_arg0, _arg1)
}

// SetEntryTextColumn sets the model column which @combo_box should use to
// get strings from to be @text_column. The column @text_column in the model
// of @combo_box must be of type G_TYPE_STRING.
//
// This is only relevant if @combo_box has been created with
// ComboBox:has-entry as true.
func (c comboBox) SetEntryTextColumn(textColumn int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (C.gint)(textColumn)

	C.gtk_combo_box_set_entry_text_column(_arg0, _arg1)
}

// SetFocusOnClick sets whether the combo box will grab focus when it is
// clicked with the mouse. Making mouse clicks not grab focus is useful in
// places like toolbars where you don’t want the keyboard focus removed from
// the main area of the application.
func (c comboBox) SetFocusOnClick(focusOnClick bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	if focusOnClick {
		_arg1 = C.TRUE
	}

	C.gtk_combo_box_set_focus_on_click(_arg0, _arg1)
}

// SetIDColumn sets the model column which @combo_box should use to get
// string IDs for values from. The column @id_column in the model of
// @combo_box must be of type G_TYPE_STRING.
func (c comboBox) SetIDColumn(idColumn int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (C.gint)(idColumn)

	C.gtk_combo_box_set_id_column(_arg0, _arg1)
}

// SetModel sets the model used by @combo_box to be @model. Will unset a
// previously set model (if applicable). If model is nil, then it will unset
// the model.
//
// Note that this function does not clear the cell renderers, you have to
// call gtk_cell_layout_clear() yourself if you need to set up different
// cell renderers for the new model.
func (c comboBox) SetModel(model TreeModel) {
	var _arg0 *C.GtkComboBox  // out
	var _arg1 *C.GtkTreeModel // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	C.gtk_combo_box_set_model(_arg0, _arg1)
}

// SetPopupFixedWidth specifies whether the popup’s width should be a fixed
// width matching the allocated width of the combo box.
func (c comboBox) SetPopupFixedWidth(fixed bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	if fixed {
		_arg1 = C.TRUE
	}

	C.gtk_combo_box_set_popup_fixed_width(_arg0, _arg1)
}

// SetRowSpanColumn sets the column with row span information for @combo_box
// to be @row_span. The row span column contains integers which indicate how
// many rows an item should span.
func (c comboBox) SetRowSpanColumn(rowSpan int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (C.gint)(rowSpan)

	C.gtk_combo_box_set_row_span_column(_arg0, _arg1)
}

// SetTitle sets the menu’s title in tearoff mode.
func (c comboBox) SetTitle(title string) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_combo_box_set_title(_arg0, _arg1)
}

// SetWrapWidth sets the wrap width of @combo_box to be @width. The wrap
// width is basically the preferred number of columns when you want the
// popup to be layed out in a table.
func (c comboBox) SetWrapWidth(width int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (C.gint)(width)

	C.gtk_combo_box_set_wrap_width(_arg0, _arg1)
}
