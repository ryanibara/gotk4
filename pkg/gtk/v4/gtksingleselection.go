// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_single_selection_get_type()), F: marshalSingleSelection},
	})
}

// SingleSelection: `GtkSingleSelection` is a `GtkSelectionModel` that allows
// selecting a single item.
//
// Note that the selection is *persistent* -- if the selected item is removed
// and re-added in the same ::items-changed emission, it stays selected. In
// particular, this means that changing the sort order of an underlying sort
// model will preserve the selection.
type SingleSelection interface {
	gextras.Objector
	gio.ListModel
	SelectionModel

	// Autoselect checks if autoselect has been enabled or disabled via
	// gtk_single_selection_set_autoselect().
	Autoselect() bool
	// CanUnselect: if true, gtk_selection_model_unselect_item() is supported
	// and allows unselecting the selected item.
	CanUnselect() bool
	// Model gets the model that @self is wrapping.
	Model() gio.ListModel
	// Selected gets the position of the selected item.
	//
	// If no item is selected, GTK_INVALID_LIST_POSITION is returned.
	Selected() uint
	// SelectedItem gets the selected item.
	//
	// If no item is selected, nil is returned.
	SelectedItem() gextras.Objector
	// SetAutoselect enables or disables autoselect.
	//
	// If @autoselect is true, @self will enforce that an item is always
	// selected. It will select a new item when the currently selected item is
	// deleted and it will disallow unselecting the current item.
	SetAutoselect(autoselect bool)
	// SetCanUnselect: if true, unselecting the current item via
	// gtk_selection_model_unselect_item() is supported.
	//
	// Note that setting [property@Gtk.SingleSelection:autoselect] will cause
	// unselecting to not work, so it practically makes no sense to set both at
	// the same time the same time.
	SetCanUnselect(canUnselect bool)
	// SetModel sets the model that @self should wrap.
	//
	// If @model is nil, @self will be empty.
	SetModel(model gio.ListModel)
	// SetSelected selects the item at the given position.
	//
	// If the list does not have an item at @position or
	// GTK_INVALID_LIST_POSITION is given, the behavior depends on the value of
	// the [property@Gtk.SingleSelection:autoselect] property: If it is set, no
	// change will occur and the old item will stay selected. If it is unset,
	// the selection will be unset and no item will be selected.
	SetSelected(position uint)
}

// singleSelection implements the SingleSelection interface.
type singleSelection struct {
	gextras.Objector
	gio.ListModel
	SelectionModel
}

var _ SingleSelection = (*singleSelection)(nil)

// WrapSingleSelection wraps a GObject to the right type. It is
// primarily used internally.
func WrapSingleSelection(obj *externglib.Object) SingleSelection {
	return SingleSelection{
		Objector:       obj,
		gio.ListModel:  gio.WrapListModel(obj),
		SelectionModel: WrapSelectionModel(obj),
	}
}

func marshalSingleSelection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSingleSelection(obj), nil
}

// NewSingleSelection constructs a class SingleSelection.
func NewSingleSelection(model gio.ListModel) SingleSelection {
	var _arg1 *C.GListModel

	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	var _cret C.GtkSingleSelection

	cret = C.gtk_single_selection_new(_arg1)

	var _singleSelection SingleSelection

	_singleSelection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(SingleSelection)

	return _singleSelection
}

// Autoselect checks if autoselect has been enabled or disabled via
// gtk_single_selection_set_autoselect().
func (s singleSelection) Autoselect() bool {
	var _arg0 *C.GtkSingleSelection

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	cret = C.gtk_single_selection_get_autoselect(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// CanUnselect: if true, gtk_selection_model_unselect_item() is supported
// and allows unselecting the selected item.
func (s singleSelection) CanUnselect() bool {
	var _arg0 *C.GtkSingleSelection

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	cret = C.gtk_single_selection_get_can_unselect(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Model gets the model that @self is wrapping.
func (s singleSelection) Model() gio.ListModel {
	var _arg0 *C.GtkSingleSelection

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))

	var _cret *C.GListModel

	cret = C.gtk_single_selection_get_model(_arg0)

	var _listModel gio.ListModel

	_listModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.ListModel)

	return _listModel
}

// Selected gets the position of the selected item.
//
// If no item is selected, GTK_INVALID_LIST_POSITION is returned.
func (s singleSelection) Selected() uint {
	var _arg0 *C.GtkSingleSelection

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))

	var _cret C.guint

	cret = C.gtk_single_selection_get_selected(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// SelectedItem gets the selected item.
//
// If no item is selected, nil is returned.
func (s singleSelection) SelectedItem() gextras.Objector {
	var _arg0 *C.GtkSingleSelection

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))

	var _cret C.gpointer

	cret = C.gtk_single_selection_get_selected_item(_arg0)

	var _object gextras.Objector

	_object = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gextras.Objector)

	return _object
}

// SetAutoselect enables or disables autoselect.
//
// If @autoselect is true, @self will enforce that an item is always
// selected. It will select a new item when the currently selected item is
// deleted and it will disallow unselecting the current item.
func (s singleSelection) SetAutoselect(autoselect bool) {
	var _arg0 *C.GtkSingleSelection
	var _arg1 C.gboolean

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))
	if autoselect {
		_arg1 = C.gboolean(1)
	}

	C.gtk_single_selection_set_autoselect(_arg0, _arg1)
}

// SetCanUnselect: if true, unselecting the current item via
// gtk_selection_model_unselect_item() is supported.
//
// Note that setting [property@Gtk.SingleSelection:autoselect] will cause
// unselecting to not work, so it practically makes no sense to set both at
// the same time the same time.
func (s singleSelection) SetCanUnselect(canUnselect bool) {
	var _arg0 *C.GtkSingleSelection
	var _arg1 C.gboolean

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))
	if canUnselect {
		_arg1 = C.gboolean(1)
	}

	C.gtk_single_selection_set_can_unselect(_arg0, _arg1)
}

// SetModel sets the model that @self should wrap.
//
// If @model is nil, @self will be empty.
func (s singleSelection) SetModel(model gio.ListModel) {
	var _arg0 *C.GtkSingleSelection
	var _arg1 *C.GListModel

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.gtk_single_selection_set_model(_arg0, _arg1)
}

// SetSelected selects the item at the given position.
//
// If the list does not have an item at @position or
// GTK_INVALID_LIST_POSITION is given, the behavior depends on the value of
// the [property@Gtk.SingleSelection:autoselect] property: If it is set, no
// change will occur and the old item will stay selected. If it is unset,
// the selection will be unset and no item will be selected.
func (s singleSelection) SetSelected(position uint) {
	var _arg0 *C.GtkSingleSelection
	var _arg1 C.guint

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(position)

	C.gtk_single_selection_set_selected(_arg0, _arg1)
}