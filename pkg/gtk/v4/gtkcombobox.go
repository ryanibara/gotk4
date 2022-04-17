// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern char* _gotk4_gtk4_ComboBoxClass_format_entry_text(GtkComboBox*, char*);
// extern gboolean _gotk4_gtk4_ComboBox_ConnectPopdown(gpointer, guintptr);
// extern gboolean _gotk4_gtk4_TreeViewRowSeparatorFunc(GtkTreeModel*, GtkTreeIter*, gpointer);
// extern gchar* _gotk4_gtk4_ComboBox_ConnectFormatEntryText(gpointer, gchar*, guintptr);
// extern void _gotk4_gtk4_ComboBoxClass_changed(GtkComboBox*);
// extern void _gotk4_gtk4_ComboBox_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gtk4_ComboBox_ConnectMoveActive(gpointer, GtkScrollType, guintptr);
// extern void _gotk4_gtk4_ComboBox_ConnectPopup(gpointer, guintptr);
// extern void callbackDelete(gpointer);
import "C"

// glib.Type values for gtkcombobox.go.
var GTypeComboBox = externglib.Type(C.gtk_combo_box_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeComboBox, F: marshalComboBox},
	})
}

// ComboBoxOverrider contains methods that are overridable.
type ComboBoxOverrider interface {
	Changed()
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	FormatEntryText(path string) string
}

// ComboBox: GtkComboBox is a widget that allows the user to choose from a list
// of valid choices.
//
// !An example GtkComboBox (combo-box.png)
//
// The GtkComboBox displays the selected choice; when activated, the GtkComboBox
// displays a popup which allows the user to make a new choice.
//
// The GtkComboBox uses the model-view pattern; the list of valid choices is
// specified in the form of a tree model, and the display of the choices can be
// adapted to the data in the model by using cell renderers, as you would in a
// tree view. This is possible since GtkComboBox implements the gtk.CellLayout
// interface. The tree model holding the valid choices is not restricted to a
// flat list, it can be a real tree, and the popup will reflect the tree
// structure.
//
// To allow the user to enter values not in the model, the
// gtk.ComboBox:has-entry property allows the GtkComboBox to contain a
// gtk.Entry. This entry can be accessed by calling gtk.ComboBox.GetChild() on
// the combo box.
//
// For a simple list of textual choices, the model-view API of GtkComboBox can
// be a bit overwhelming. In this case, gtk.ComboBoxText offers a simple
// alternative. Both GtkComboBox and GtkComboBoxText can contain an entry.
//
// CSS nodes
//
//    combobox
//    ├── box.linked
//    │   ╰── button.combo
//    │       ╰── box
//    │           ├── cellview
//    │           ╰── arrow
//    ╰── window.popup
//
//
// A normal combobox contains a box with the .linked class, a button with the
// .combo class and inside those buttons, there are a cellview and an arrow.
//
//    combobox
//    ├── box.linked
//    │   ├── entry.combo
//    │   ╰── button.combo
//    │       ╰── box
//    │           ╰── arrow
//    ╰── window.popup
//
//
// A GtkComboBox with an entry has a single CSS node with name combobox. It
// contains a box with the .linked class. That box contains an entry and a
// button, both with the .combo class added. The button also contains another
// node with name arrow.
//
//
// Accessibility
//
// GtkComboBox uses the GTK_ACCESSIBLE_ROLE_COMBO_BOX role.
type ComboBox struct {
	_ [0]func() // equal guard
	Widget

	*externglib.Object
	CellEditable
	CellLayout
}

var (
	_ Widgetter           = (*ComboBox)(nil)
	_ externglib.Objector = (*ComboBox)(nil)
)

func classInitComboBoxer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkComboBoxClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkComboBoxClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Changed() }); ok {
		pclass.changed = (*[0]byte)(C._gotk4_gtk4_ComboBoxClass_changed)
	}

	if _, ok := goval.(interface{ FormatEntryText(path string) string }); ok {
		pclass.format_entry_text = (*[0]byte)(C._gotk4_gtk4_ComboBoxClass_format_entry_text)
	}
}

//export _gotk4_gtk4_ComboBoxClass_changed
func _gotk4_gtk4_ComboBoxClass_changed(arg0 *C.GtkComboBox) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Changed() })

	iface.Changed()
}

//export _gotk4_gtk4_ComboBoxClass_format_entry_text
func _gotk4_gtk4_ComboBoxClass_format_entry_text(arg0 *C.GtkComboBox, arg1 *C.char) (cret *C.char) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ FormatEntryText(path string) string })

	var _path string // out

	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	utf8 := iface.FormatEntryText(_path)

	cret = (*C.char)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

func wrapComboBox(obj *externglib.Object) *ComboBox {
	return &ComboBox{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
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
		Object: obj,
		CellEditable: CellEditable{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
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

func marshalComboBox(p uintptr) (interface{}, error) {
	return wrapComboBox(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_ComboBox_ConnectChanged
func _gotk4_gtk4_ComboBox_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectChanged is emitted when the active item is changed.
//
// The can be due to the user selecting a different item from the list, or due
// to a call to gtk.ComboBox.SetActiveIter(). It will also be emitted while
// typing into the entry of a combo box with an entry.
func (comboBox *ComboBox) ConnectChanged(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(comboBox, "changed", false, unsafe.Pointer(C._gotk4_gtk4_ComboBox_ConnectChanged), f)
}

//export _gotk4_gtk4_ComboBox_ConnectFormatEntryText
func _gotk4_gtk4_ComboBox_ConnectFormatEntryText(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) (cret *C.gchar) {
	var f func(path string) (utf8 string)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(path string) (utf8 string))
	}

	var _path string // out

	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	utf8 := f(_path)

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

// ConnectFormatEntryText is emitted to allow changing how the text in a combo
// box's entry is displayed.
//
// See gtk.ComboBox:has-entry.
//
// Connect a signal handler which returns an allocated string representing path.
// That string will then be used to set the text in the combo box's entry. The
// default signal handler uses the text from the gtk.ComboBox:entry-text-column
// model column.
//
// Here's an example signal handler which fetches data from the model and
// displays it in the entry.
//
//    static char *
//    format_entry_text_callback (GtkComboBox *combo,
//                                const char *path,
//                                gpointer     user_data)
//    {
//      GtkTreeIter iter;
//      GtkTreeModel model;
//      double       value;
//
//      model = gtk_combo_box_get_model (combo);
//
//      gtk_tree_model_get_iter_from_string (model, &iter, path);
//      gtk_tree_model_get (model, &iter,
//                          THE_DOUBLE_VALUE_COLUMN, &value,
//                          -1);
//
//      return g_strdup_printf ("g", value);
//    }.
func (comboBox *ComboBox) ConnectFormatEntryText(f func(path string) (utf8 string)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(comboBox, "format-entry-text", false, unsafe.Pointer(C._gotk4_gtk4_ComboBox_ConnectFormatEntryText), f)
}

//export _gotk4_gtk4_ComboBox_ConnectMoveActive
func _gotk4_gtk4_ComboBox_ConnectMoveActive(arg0 C.gpointer, arg1 C.GtkScrollType, arg2 C.guintptr) {
	var f func(scrollType ScrollType)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(scrollType ScrollType))
	}

	var _scrollType ScrollType // out

	_scrollType = ScrollType(arg1)

	f(_scrollType)
}

// ConnectMoveActive is emitted to move the active selection.
//
// This is an keybinding signal (class.SignalAction.html).
func (comboBox *ComboBox) ConnectMoveActive(f func(scrollType ScrollType)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(comboBox, "move-active", false, unsafe.Pointer(C._gotk4_gtk4_ComboBox_ConnectMoveActive), f)
}

//export _gotk4_gtk4_ComboBox_ConnectPopdown
func _gotk4_gtk4_ComboBox_ConnectPopdown(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (ok bool))
	}

	ok := f()

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectPopdown is emitted to popdown the combo box list.
//
// This is an keybinding signal (class.SignalAction.html).
//
// The default bindings for this signal are Alt+Up and Escape.
func (comboBox *ComboBox) ConnectPopdown(f func() (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(comboBox, "popdown", false, unsafe.Pointer(C._gotk4_gtk4_ComboBox_ConnectPopdown), f)
}

//export _gotk4_gtk4_ComboBox_ConnectPopup
func _gotk4_gtk4_ComboBox_ConnectPopup(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectPopup is emitted to popup the combo box list.
//
// This is an keybinding signal (class.SignalAction.html).
//
// The default binding for this signal is Alt+Down.
func (comboBox *ComboBox) ConnectPopup(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(comboBox, "popup", false, unsafe.Pointer(C._gotk4_gtk4_ComboBox_ConnectPopup), f)
}

// NewComboBox creates a new empty GtkComboBox.
//
// The function returns the following values:
//
//    - comboBox: new GtkComboBox.
//
func NewComboBox() *ComboBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_new()

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithEntry creates a new empty GtkComboBox with an entry.
//
// The function returns the following values:
//
//    - comboBox: new GtkComboBox.
//
func NewComboBoxWithEntry() *ComboBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_new_with_entry()

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithModel creates a new GtkComboBox with a model.
//
// The function takes the following parameters:
//
//    - model: GtkTreeModel.
//
// The function returns the following values:
//
//    - comboBox: new GtkComboBox.
//
func NewComboBoxWithModel(model TreeModeller) *ComboBox {
	var _arg1 *C.GtkTreeModel // out
	var _cret *C.GtkWidget    // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(externglib.InternObject(model).Native()))

	_cret = C.gtk_combo_box_new_with_model(_arg1)
	runtime.KeepAlive(model)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithModelAndEntry creates a new empty GtkComboBox with an entry
// and a model.
//
// The function takes the following parameters:
//
//    - model: GtkTreeModel.
//
// The function returns the following values:
//
//    - comboBox: new GtkComboBox.
//
func NewComboBoxWithModelAndEntry(model TreeModeller) *ComboBox {
	var _arg1 *C.GtkTreeModel // out
	var _cret *C.GtkWidget    // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(externglib.InternObject(model).Native()))

	_cret = C.gtk_combo_box_new_with_model_and_entry(_arg1)
	runtime.KeepAlive(model)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// Active returns the index of the currently active item.
//
// If the model is a non-flat treemodel, and the active item is not an immediate
// child of the root of the tree, this function returns
// gtk_tree_path_get_indices (path)[0], where path is the gtk.TreePath of the
// active item.
//
// The function returns the following values:
//
//    - gint: integer which is the index of the currently active item, or -1 if
//      there’s no active item.
//
func (comboBox *ComboBox) Active() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.int          // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_active(_arg0)
	runtime.KeepAlive(comboBox)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ActiveID returns the ID of the active row of combo_box.
//
// This value is taken from the active row and the column specified by the
// gtk.ComboBox:id-column property of combo_box (see
// gtk.ComboBox.SetIDColumn()).
//
// The returned value is an interned string which means that you can compare the
// pointer by value to other interned strings and that you must not free it.
//
// If the gtk.ComboBox:id-column property of combo_box is not set, or if no row
// is active, or if the active row has a NULL ID value, then NULL is returned.
//
// The function returns the following values:
//
//    - utf8 (optional): ID of the active row, or NULL.
//
func (comboBox *ComboBox) ActiveID() string {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.char        // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_active_id(_arg0)
	runtime.KeepAlive(comboBox)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ActiveIter sets iter to point to the currently active item.
//
// If no item is active, iter is left unchanged.
//
// The function returns the following values:
//
//    - iter: GtkTreeIter.
//    - ok: TRUE if iter was set, FALSE otherwise.
//
func (comboBox *ComboBox) ActiveIter() (*TreeIter, bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.GtkTreeIter  // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_active_iter(_arg0, &_arg1)
	runtime.KeepAlive(comboBox)

	var _iter *TreeIter // out
	var _ok bool        // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	if _cret != 0 {
		_ok = true
	}

	return _iter, _ok
}

// ButtonSensitivity returns whether the combo box sets the dropdown button
// sensitive or not when there are no items in the model.
//
// The function returns the following values:
//
//    - sensitivityType: GTK_SENSITIVITY_ON if the dropdown button is sensitive
//      when the model is empty, GTK_SENSITIVITY_OFF if the button is always
//      insensitive or GTK_SENSITIVITY_AUTO if it is only sensitive as long as
//      the model has one item to be selected.
//
func (comboBox *ComboBox) ButtonSensitivity() SensitivityType {
	var _arg0 *C.GtkComboBox       // out
	var _cret C.GtkSensitivityType // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_button_sensitivity(_arg0)
	runtime.KeepAlive(comboBox)

	var _sensitivityType SensitivityType // out

	_sensitivityType = SensitivityType(_cret)

	return _sensitivityType
}

// Child gets the child widget of combo_box.
//
// The function returns the following values:
//
//    - widget (optional): child widget of combo_box.
//
func (comboBox *ComboBox) Child() Widgetter {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_child(_arg0)
	runtime.KeepAlive(comboBox)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// EntryTextColumn returns the column which combo_box is using to get the
// strings from to display in the internal entry.
//
// The function returns the following values:
//
//    - gint: column in the data source model of combo_box.
//
func (comboBox *ComboBox) EntryTextColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.int          // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_entry_text_column(_arg0)
	runtime.KeepAlive(comboBox)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// HasEntry returns whether the combo box has an entry.
//
// The function returns the following values:
//
//    - ok: whether there is an entry in combo_box.
//
func (comboBox *ComboBox) HasEntry() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_has_entry(_arg0)
	runtime.KeepAlive(comboBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IDColumn returns the column which combo_box is using to get string IDs for
// values from.
//
// The function returns the following values:
//
//    - gint: column in the data source model of combo_box.
//
func (comboBox *ComboBox) IDColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.int          // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_id_column(_arg0)
	runtime.KeepAlive(comboBox)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Model returns the GtkTreeModel of combo_box.
//
// The function returns the following values:
//
//    - treeModel (optional): GtkTreeModel which was passed during construction.
//
func (comboBox *ComboBox) Model() *TreeModel {
	var _arg0 *C.GtkComboBox  // out
	var _cret *C.GtkTreeModel // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_model(_arg0)
	runtime.KeepAlive(comboBox)

	var _treeModel *TreeModel // out

	if _cret != nil {
		_treeModel = wrapTreeModel(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _treeModel
}

// PopupFixedWidth gets whether the popup uses a fixed width.
//
// The function returns the following values:
//
//    - ok: TRUE if the popup uses a fixed width.
//
func (comboBox *ComboBox) PopupFixedWidth() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_popup_fixed_width(_arg0)
	runtime.KeepAlive(comboBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Popdown hides the menu or dropdown list of combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
func (comboBox *ComboBox) Popdown() {
	var _arg0 *C.GtkComboBox // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))

	C.gtk_combo_box_popdown(_arg0)
	runtime.KeepAlive(comboBox)
}

// Popup pops up the menu or dropdown list of combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
//
// Before calling this, combo_box must be mapped, or nothing will happen.
func (comboBox *ComboBox) Popup() {
	var _arg0 *C.GtkComboBox // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))

	C.gtk_combo_box_popup(_arg0)
	runtime.KeepAlive(comboBox)
}

// PopupForDevice pops up the menu of combo_box.
//
// Note that currently this does not do anything with the device, as it was
// previously only used for list-mode combo boxes, and those were removed in GTK
// 4. However, it is retained in case similar functionality is added back later.
//
// The function takes the following parameters:
//
//    - device: GdkDevice.
//
func (comboBox *ComboBox) PopupForDevice(device gdk.Devicer) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GdkDevice   // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(externglib.InternObject(device).Native()))

	C.gtk_combo_box_popup_for_device(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(device)
}

// SetActive sets the active item of combo_box to be the item at index.
//
// The function takes the following parameters:
//
//    - index_: index in the model passed during construction, or -1 to have no
//      active item.
//
func (comboBox *ComboBox) SetActive(index_ int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))
	_arg1 = C.int(index_)

	C.gtk_combo_box_set_active(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(index_)
}

// SetActiveID changes the active row of combo_box to the one that has an ID
// equal to active_id.
//
// If active_id is NULL, the active row is unset. Rows having a NULL ID string
// cannot be made active by this function.
//
// If the gtk.ComboBox:id-column property of combo_box is unset or if no row has
// the given ID then the function does nothing and returns FALSE.
//
// The function takes the following parameters:
//
//    - activeId (optional): ID of the row to select, or NULL.
//
// The function returns the following values:
//
//    - ok: TRUE if a row with a matching ID was found. If a NULL active_id was
//      given to unset the active row, the function always returns TRUE.
//
func (comboBox *ComboBox) SetActiveID(activeId string) bool {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.char        // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))
	if activeId != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(activeId)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_combo_box_set_active_id(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(activeId)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActiveIter sets the current active item to be the one referenced by iter.
//
// If iter is NULL, the active item is unset.
//
// The function takes the following parameters:
//
//    - iter (optional): GtkTreeIter, or NULL.
//
func (comboBox *ComboBox) SetActiveIter(iter *TreeIter) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GtkTreeIter // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))
	if iter != nil {
		_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	}

	C.gtk_combo_box_set_active_iter(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(iter)
}

// SetButtonSensitivity sets whether the dropdown button of the combo box should
// update its sensitivity depending on the model contents.
//
// The function takes the following parameters:
//
//    - sensitivity: specify the sensitivity of the dropdown button.
//
func (comboBox *ComboBox) SetButtonSensitivity(sensitivity SensitivityType) {
	var _arg0 *C.GtkComboBox       // out
	var _arg1 C.GtkSensitivityType // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))
	_arg1 = C.GtkSensitivityType(sensitivity)

	C.gtk_combo_box_set_button_sensitivity(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(sensitivity)
}

// SetChild sets the child widget of combo_box.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (comboBox *ComboBox) SetChild(child Widgetter) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(child).Native()))
	}

	C.gtk_combo_box_set_child(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(child)
}

// SetEntryTextColumn sets the model column which combo_box should use to get
// strings from to be text_column.
//
// The column text_column in the model of combo_box must be of type
// G_TYPE_STRING.
//
// This is only relevant if combo_box has been created with
// gtk.ComboBox:has-entry as TRUE.
//
// The function takes the following parameters:
//
//    - textColumn: column in model to get the strings from for the internal
//      entry.
//
func (comboBox *ComboBox) SetEntryTextColumn(textColumn int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))
	_arg1 = C.int(textColumn)

	C.gtk_combo_box_set_entry_text_column(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(textColumn)
}

// SetIDColumn sets the model column which combo_box should use to get string
// IDs for values from.
//
// The column id_column in the model of combo_box must be of type G_TYPE_STRING.
//
// The function takes the following parameters:
//
//    - idColumn: column in model to get string IDs for values from.
//
func (comboBox *ComboBox) SetIDColumn(idColumn int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))
	_arg1 = C.int(idColumn)

	C.gtk_combo_box_set_id_column(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(idColumn)
}

// SetModel sets the model used by combo_box to be model.
//
// Will unset a previously set model (if applicable). If model is NULL, then it
// will unset the model.
//
// Note that this function does not clear the cell renderers, you have to call
// gtk.CellLayout.Clear() yourself if you need to set up different cell
// renderers for the new model.
//
// The function takes the following parameters:
//
//    - model (optional): GtkTreeModel.
//
func (comboBox *ComboBox) SetModel(model TreeModeller) {
	var _arg0 *C.GtkComboBox  // out
	var _arg1 *C.GtkTreeModel // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))
	if model != nil {
		_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(externglib.InternObject(model).Native()))
	}

	C.gtk_combo_box_set_model(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(model)
}

// SetPopupFixedWidth specifies whether the popup’s width should be a fixed
// width.
//
// If fixed is TRUE, the popup's width is set to match the allocated width of
// the combo box.
//
// The function takes the following parameters:
//
//    - fixed: whether to use a fixed popup width.
//
func (comboBox *ComboBox) SetPopupFixedWidth(fixed bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))
	if fixed {
		_arg1 = C.TRUE
	}

	C.gtk_combo_box_set_popup_fixed_width(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(fixed)
}

// SetRowSeparatorFunc sets the row separator function, which is used to
// determine whether a row should be drawn as a separator.
//
// If the row separator function is NULL, no separators are drawn. This is the
// default value.
//
// The function takes the following parameters:
//
//    - fn (optional): GtkTreeViewRowSeparatorFunc.
//
func (comboBox *ComboBox) SetRowSeparatorFunc(fn TreeViewRowSeparatorFunc) {
	var _arg0 *C.GtkComboBox                // out
	var _arg1 C.GtkTreeViewRowSeparatorFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(externglib.InternObject(comboBox).Native()))
	if fn != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk4_TreeViewRowSeparatorFunc)
		_arg2 = C.gpointer(gbox.Assign(fn))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_combo_box_set_row_separator_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(fn)
}
