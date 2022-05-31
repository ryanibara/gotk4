// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gtk3_ComboBox_ConnectPopdown(gpointer, guintptr);
// extern gchar* _gotk4_gtk3_ComboBoxClass_format_entry_text(GtkComboBox*, gchar*);
// extern gchar* _gotk4_gtk3_ComboBox_ConnectFormatEntryText(gpointer, gchar*, guintptr);
// extern void _gotk4_gtk3_ComboBoxClass_changed(GtkComboBox*);
// extern void _gotk4_gtk3_ComboBox_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gtk3_ComboBox_ConnectPopup(gpointer, guintptr);
import "C"

// glib.Type values for gtkcombobox.go.
var GTypeComboBox = coreglib.Type(C.gtk_combo_box_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
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

// ComboBox is a widget that allows the user to choose from a list of valid
// choices. The GtkComboBox displays the selected choice. When activated, the
// GtkComboBox displays a popup which allows the user to make a new choice. The
// style in which the selected value is displayed, and the style of the popup is
// determined by the current theme. It may be similar to a Windows-style combo
// box.
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
type ComboBox struct {
	_ [0]func() // equal guard
	Bin

	*coreglib.Object
	CellEditable
	CellLayout
}

var (
	_ Binner            = (*ComboBox)(nil)
	_ coreglib.Objector = (*ComboBox)(nil)
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
		pclass.changed = (*[0]byte)(C._gotk4_gtk3_ComboBoxClass_changed)
	}

	if _, ok := goval.(interface{ FormatEntryText(path string) string }); ok {
		pclass.format_entry_text = (*[0]byte)(C._gotk4_gtk3_ComboBoxClass_format_entry_text)
	}
}

//export _gotk4_gtk3_ComboBoxClass_changed
func _gotk4_gtk3_ComboBoxClass_changed(arg0 *C.GtkComboBox) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Changed() })

	iface.Changed()
}

//export _gotk4_gtk3_ComboBoxClass_format_entry_text
func _gotk4_gtk3_ComboBoxClass_format_entry_text(arg0 *C.GtkComboBox, arg1 *C.gchar) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ FormatEntryText(path string) string })

	var _path string // out

	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	utf8 := iface.FormatEntryText(_path)

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

func wrapComboBox(obj *coreglib.Object) *ComboBox {
	return &ComboBox{
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
	}
}

func marshalComboBox(p uintptr) (interface{}, error) {
	return wrapComboBox(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_ComboBox_ConnectChanged
func _gotk4_gtk3_ComboBox_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectChanged signal is emitted when the active item is changed. The can be
// due to the user selecting a different item from the list, or due to a call to
// gtk_combo_box_set_active_iter(). It will also be emitted while typing into
// the entry of a combo box with an entry.
func (comboBox *ComboBox) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(comboBox, "changed", false, unsafe.Pointer(C._gotk4_gtk3_ComboBox_ConnectChanged), f)
}

//export _gotk4_gtk3_ComboBox_ConnectFormatEntryText
func _gotk4_gtk3_ComboBox_ConnectFormatEntryText(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) (cret *C.gchar) {
	var f func(path string) (utf8 string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(path string) (utf8 string))
	}

	var _path string // out

	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	utf8 := f(_path)

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

// ConnectFormatEntryText: for combo boxes that are created with an entry (See
// GtkComboBox:has-entry).
//
// A signal which allows you to change how the text displayed in a combo box's
// entry is displayed.
//
// Connect a signal handler which returns an allocated string representing path.
// That string will then be used to set the text in the combo box's entry. The
// default signal handler uses the text from the GtkComboBox::entry-text-column
// model column.
//
// Here's an example signal handler which fetches data from the model and
// displays it in the entry.
//
//    static gchar*
//    format_entry_text_callback (GtkComboBox *combo,
//                                const gchar *path,
//                                gpointer     user_data)
//    {
//      GtkTreeIter iter;
//      GtkTreeModel model;
//      gdouble      value;
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
func (comboBox *ComboBox) ConnectFormatEntryText(f func(path string) (utf8 string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(comboBox, "format-entry-text", false, unsafe.Pointer(C._gotk4_gtk3_ComboBox_ConnectFormatEntryText), f)
}

//export _gotk4_gtk3_ComboBox_ConnectPopdown
func _gotk4_gtk3_ComboBox_ConnectPopdown(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
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

// ConnectPopdown signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted to popdown the combo box list.
//
// The default bindings for this signal are Alt+Up and Escape.
func (comboBox *ComboBox) ConnectPopdown(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(comboBox, "popdown", false, unsafe.Pointer(C._gotk4_gtk3_ComboBox_ConnectPopdown), f)
}

//export _gotk4_gtk3_ComboBox_ConnectPopup
func _gotk4_gtk3_ComboBox_ConnectPopup(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectPopup signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted to popup the combo box list.
//
// The default binding for this signal is Alt+Down.
func (comboBox *ComboBox) ConnectPopup(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(comboBox, "popup", false, unsafe.Pointer(C._gotk4_gtk3_ComboBox_ConnectPopup), f)
}

// NewComboBox creates a new empty ComboBox.
//
// The function returns the following values:
//
//    - comboBox: new ComboBox.
//
func NewComboBox() *ComboBox {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("new_ComboBox", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithArea creates a new empty ComboBox using area to layout cells.
//
// The function takes the following parameters:
//
//    - area to use to layout cell renderers.
//
// The function returns the following values:
//
//    - comboBox: new ComboBox.
//
func NewComboBoxWithArea(area CellAreaer) *ComboBox {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(area).Native()))
	*(*CellAreaer)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("new_ComboBox_with_area", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(area)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithAreaAndEntry creates a new empty ComboBox with an entry.
//
// The new combo box will use area to layout cells.
//
// The function takes the following parameters:
//
//    - area to use to layout cell renderers.
//
// The function returns the following values:
//
//    - comboBox: new ComboBox.
//
func NewComboBoxWithAreaAndEntry(area CellAreaer) *ComboBox {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(area).Native()))
	*(*CellAreaer)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("new_ComboBox_with_area_and_entry", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(area)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithEntry creates a new empty ComboBox with an entry.
//
// The function returns the following values:
//
//    - comboBox: new ComboBox.
//
func NewComboBoxWithEntry() *ComboBox {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("new_ComboBox_with_entry", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithModel creates a new ComboBox with the model initialized to
// model.
//
// The function takes the following parameters:
//
//    - model: TreeModel.
//
// The function returns the following values:
//
//    - comboBox: new ComboBox.
//
func NewComboBoxWithModel(model TreeModeller) *ComboBox {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	*(*TreeModeller)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("new_ComboBox_with_model", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(model)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithModelAndEntry creates a new empty ComboBox with an entry and
// with the model initialized to model.
//
// The function takes the following parameters:
//
//    - model: TreeModel.
//
// The function returns the following values:
//
//    - comboBox: new ComboBox.
//
func NewComboBoxWithModelAndEntry(model TreeModeller) *ComboBox {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	*(*TreeModeller)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("new_ComboBox_with_model_and_entry", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(model)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// Active returns the index of the currently active item, or -1 if there’s no
// active item. If the model is a non-flat treemodel, and the active item is not
// an immediate child of the root of the tree, this function returns
// gtk_tree_path_get_indices (path)[0], where path is the TreePath of the active
// item.
//
// The function returns the following values:
//
//    - gint: integer which is the index of the currently active item, or -1 if
//      there’s no active item.
//
func (comboBox *ComboBox) Active() int32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_active", args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// ActiveID returns the ID of the active row of combo_box. This value is taken
// from the active row and the column specified by the ComboBox:id-column
// property of combo_box (see gtk_combo_box_set_id_column()).
//
// The returned value is an interned string which means that you can compare the
// pointer by value to other interned strings and that you must not free it.
//
// If the ComboBox:id-column property of combo_box is not set, or if no row is
// active, or if the active row has a NULL ID value, then NULL is returned.
//
// The function returns the following values:
//
//    - utf8 (optional): ID of the active row, or NULL.
//
func (comboBox *ComboBox) ActiveID() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_active_id", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// AddTearoffs gets the current value of the :add-tearoffs property.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - ok: current value of the :add-tearoffs property.
//
func (comboBox *ComboBox) AddTearoffs() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_add_tearoffs", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ColumnSpanColumn returns the column with column span information for
// combo_box.
//
// The function returns the following values:
//
//    - gint: column span column.
//
func (comboBox *ComboBox) ColumnSpanColumn() int32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_column_span_column", args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// EntryTextColumn returns the column which combo_box is using to get the
// strings from to display in the internal entry.
//
// The function returns the following values:
//
//    - gint: column in the data source model of combo_box.
//
func (comboBox *ComboBox) EntryTextColumn() int32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_entry_text_column", args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// FocusOnClick returns whether the combo box grabs focus when it is clicked
// with the mouse. See gtk_combo_box_set_focus_on_click().
//
// Deprecated: Use gtk_widget_get_focus_on_click() instead.
//
// The function returns the following values:
//
//    - ok: TRUE if the combo box grabs focus when it is clicked with the mouse.
//
func (combo *ComboBox) FocusOnClick() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(combo).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_focus_on_click", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(combo)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasEntry returns whether the combo box has an entry.
//
// The function returns the following values:
//
//    - ok: whether there is an entry in combo_box.
//
func (comboBox *ComboBox) HasEntry() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_has_entry", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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
func (comboBox *ComboBox) IDColumn() int32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_id_column", args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// Model returns the TreeModel which is acting as data source for combo_box.
//
// The function returns the following values:
//
//    - treeModel which was passed during construction.
//
func (comboBox *ComboBox) Model() *TreeModel {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_model", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _treeModel *TreeModel // out

	_treeModel = wrapTreeModel(coreglib.Take(unsafe.Pointer(_cret)))

	return _treeModel
}

// PopupAccessible gets the accessible object corresponding to the combo box’s
// popup.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
//
// The function returns the following values:
//
//    - object: accessible object corresponding to the combo box’s popup.
//
func (comboBox *ComboBox) PopupAccessible() *atk.ObjectClass {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_popup_accessible", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _object *atk.ObjectClass // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_object = &atk.ObjectClass{
			Object: obj,
		}
	}

	return _object
}

// PopupFixedWidth gets whether the popup uses a fixed width matching the
// allocated width of the combo box.
//
// The function returns the following values:
//
//    - ok: TRUE if the popup uses a fixed width.
//
func (comboBox *ComboBox) PopupFixedWidth() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_popup_fixed_width", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowSpanColumn returns the column with row span information for combo_box.
//
// The function returns the following values:
//
//    - gint: row span column.
//
func (comboBox *ComboBox) RowSpanColumn() int32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_row_span_column", args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// Title gets the current title of the menu in tearoff mode. See
// gtk_combo_box_set_add_tearoffs().
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - utf8 menu’s title in tearoff mode. This is an internal copy of the string
//      which must not be freed.
//
func (comboBox *ComboBox) Title() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_title", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// WrapWidth returns the wrap width which is used to determine the number of
// columns for the popup menu. If the wrap width is larger than 1, the combo box
// is in table mode.
//
// The function returns the following values:
//
//    - gint: wrap width.
//
func (comboBox *ComboBox) WrapWidth() int32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_wrap_width", args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// Popdown hides the menu or dropdown list of combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
func (comboBox *ComboBox) Popdown() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("popdown", args[:], nil)

	runtime.KeepAlive(comboBox)
}

// Popup pops up the menu or dropdown list of combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
//
// Before calling this, combo_box must be mapped, or nothing will happen.
func (comboBox *ComboBox) Popup() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("popup", args[:], nil)

	runtime.KeepAlive(comboBox)
}

// PopupForDevice pops up the menu or dropdown list of combo_box, the popup
// window will be grabbed so only device and its associated pointer/keyboard are
// the only Devices able to send events to it.
//
// The function takes the following parameters:
//
//    - device: Device.
//
func (comboBox *ComboBox) PopupForDevice(device gdk.Devicer) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("popup_for_device", args[:], nil)

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
func (comboBox *ComboBox) SetActive(index_ int32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.gint(index_)
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_active", args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(index_)
}

// SetActiveID changes the active row of combo_box to the one that has an ID
// equal to active_id, or unsets the active row if active_id is NULL. Rows
// having a NULL ID string cannot be made active by this function.
//
// If the ComboBox:id-column property of combo_box is unset or if no row has the
// given ID then the function does nothing and returns FALSE.
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
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if activeId != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(activeId)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_active_id", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(activeId)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActiveIter sets the current active item to be the one referenced by iter,
// or unsets the active item if iter is NULL.
//
// The function takes the following parameters:
//
//    - iter (optional) or NULL.
//
func (comboBox *ComboBox) SetActiveIter(iter *TreeIter) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if iter != nil {
		_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	}
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_active_iter", args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(iter)
}

// SetAddTearoffs sets whether the popup menu should have a tearoff menu item.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - addTearoffs: TRUE to add tearoff menu items.
//
func (comboBox *ComboBox) SetAddTearoffs(addTearoffs bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if addTearoffs {
		_arg1 = C.TRUE
	}
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_add_tearoffs", args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(addTearoffs)
}

// SetColumnSpanColumn sets the column with column span information for
// combo_box to be column_span. The column span column contains integers which
// indicate how many columns an item should span.
//
// The function takes the following parameters:
//
//    - columnSpan: column in the model passed during construction.
//
func (comboBox *ComboBox) SetColumnSpanColumn(columnSpan int32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.gint(columnSpan)
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_column_span_column", args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(columnSpan)
}

// SetEntryTextColumn sets the model column which combo_box should use to get
// strings from to be text_column. The column text_column in the model of
// combo_box must be of type G_TYPE_STRING.
//
// This is only relevant if combo_box has been created with ComboBox:has-entry
// as TRUE.
//
// The function takes the following parameters:
//
//    - textColumn: column in model to get the strings from for the internal
//      entry.
//
func (comboBox *ComboBox) SetEntryTextColumn(textColumn int32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.gint(textColumn)
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_entry_text_column", args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(textColumn)
}

// SetFocusOnClick sets whether the combo box will grab focus when it is clicked
// with the mouse. Making mouse clicks not grab focus is useful in places like
// toolbars where you don’t want the keyboard focus removed from the main area
// of the application.
//
// Deprecated: Use gtk_widget_set_focus_on_click() instead.
//
// The function takes the following parameters:
//
//    - focusOnClick: whether the combo box grabs focus when clicked with the
//      mouse.
//
func (combo *ComboBox) SetFocusOnClick(focusOnClick bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(combo).Native()))
	if focusOnClick {
		_arg1 = C.TRUE
	}
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_focus_on_click", args[:], nil)

	runtime.KeepAlive(combo)
	runtime.KeepAlive(focusOnClick)
}

// SetIDColumn sets the model column which combo_box should use to get string
// IDs for values from. The column id_column in the model of combo_box must be
// of type G_TYPE_STRING.
//
// The function takes the following parameters:
//
//    - idColumn: column in model to get string IDs for values from.
//
func (comboBox *ComboBox) SetIDColumn(idColumn int32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.gint(idColumn)
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_id_column", args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(idColumn)
}

// SetModel sets the model used by combo_box to be model. Will unset a
// previously set model (if applicable). If model is NULL, then it will unset
// the model.
//
// Note that this function does not clear the cell renderers, you have to call
// gtk_cell_layout_clear() yourself if you need to set up different cell
// renderers for the new model.
//
// The function takes the following parameters:
//
//    - model (optional): TreeModel.
//
func (comboBox *ComboBox) SetModel(model TreeModeller) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if model != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_model", args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(model)
}

// SetPopupFixedWidth specifies whether the popup’s width should be a fixed
// width matching the allocated width of the combo box.
//
// The function takes the following parameters:
//
//    - fixed: whether to use a fixed popup width.
//
func (comboBox *ComboBox) SetPopupFixedWidth(fixed bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if fixed {
		_arg1 = C.TRUE
	}
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_popup_fixed_width", args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(fixed)
}

// SetRowSpanColumn sets the column with row span information for combo_box to
// be row_span. The row span column contains integers which indicate how many
// rows an item should span.
//
// The function takes the following parameters:
//
//    - rowSpan: column in the model passed during construction.
//
func (comboBox *ComboBox) SetRowSpanColumn(rowSpan int32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.gint(rowSpan)
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_row_span_column", args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(rowSpan)
}

// SetTitle sets the menu’s title in tearoff mode.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - title for the menu in tearoff mode.
//
func (comboBox *ComboBox) SetTitle(title string) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_title", args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(title)
}

// SetWrapWidth sets the wrap width of combo_box to be width. The wrap width is
// basically the preferred number of columns when you want the popup to be layed
// out in a table.
//
// The function takes the following parameters:
//
//    - width: preferred number of columns.
//
func (comboBox *ComboBox) SetWrapWidth(width int32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.gint(width)
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_wrap_width", args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(width)
}
