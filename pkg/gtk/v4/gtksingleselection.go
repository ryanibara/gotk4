// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
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
	var arg1 *C.GListModel

	arg1 = (*C.GListModel)(model.Native())

	ret := C.gtk_single_selection_new(arg1)

	var ret0 SingleSelection

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(SingleSelection)

	return ret0
}

// Autoselect checks if autoselect has been enabled or disabled via
// gtk_single_selection_set_autoselect().
func (s singleSelection) Autoselect() bool {
	var arg0 *C.GtkSingleSelection

	arg0 = (*C.GtkSingleSelection)(s.Native())

	ret := C.gtk_single_selection_get_autoselect(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// CanUnselect: if true, gtk_selection_model_unselect_item() is supported
// and allows unselecting the selected item.
func (s singleSelection) CanUnselect() bool {
	var arg0 *C.GtkSingleSelection

	arg0 = (*C.GtkSingleSelection)(s.Native())

	ret := C.gtk_single_selection_get_can_unselect(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Model gets the model that @self is wrapping.
func (s singleSelection) Model() gio.ListModel {
	var arg0 *C.GtkSingleSelection

	arg0 = (*C.GtkSingleSelection)(s.Native())

	ret := C.gtk_single_selection_get_model(arg0)

	var ret0 gio.ListModel

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(gio.ListModel)

	return ret0
}

// Selected gets the position of the selected item.
//
// If no item is selected, GTK_INVALID_LIST_POSITION is returned.
func (s singleSelection) Selected() uint {
	var arg0 *C.GtkSingleSelection

	arg0 = (*C.GtkSingleSelection)(s.Native())

	ret := C.gtk_single_selection_get_selected(arg0)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// SelectedItem gets the selected item.
//
// If no item is selected, nil is returned.
func (s singleSelection) SelectedItem() gextras.Objector {
	var arg0 *C.GtkSingleSelection

	arg0 = (*C.GtkSingleSelection)(s.Native())

	ret := C.gtk_single_selection_get_selected_item(arg0)

	var ret0 gextras.Objector

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(gextras.Objector)

	return ret0
}

// SetAutoselect enables or disables autoselect.
//
// If @autoselect is true, @self will enforce that an item is always
// selected. It will select a new item when the currently selected item is
// deleted and it will disallow unselecting the current item.
func (s singleSelection) SetAutoselect(autoselect bool) {
	var arg0 *C.GtkSingleSelection
	var arg1 C.gboolean

	arg0 = (*C.GtkSingleSelection)(s.Native())
	if autoselect {
		arg1 = C.TRUE
	}

	C.gtk_single_selection_set_autoselect(arg0, arg1)
}

// SetCanUnselect: if true, unselecting the current item via
// gtk_selection_model_unselect_item() is supported.
//
// Note that setting [property@Gtk.SingleSelection:autoselect] will cause
// unselecting to not work, so it practically makes no sense to set both at
// the same time the same time.
func (s singleSelection) SetCanUnselect(canUnselect bool) {
	var arg0 *C.GtkSingleSelection
	var arg1 C.gboolean

	arg0 = (*C.GtkSingleSelection)(s.Native())
	if canUnselect {
		arg1 = C.TRUE
	}

	C.gtk_single_selection_set_can_unselect(arg0, arg1)
}

// SetModel sets the model that @self should wrap.
//
// If @model is nil, @self will be empty.
func (s singleSelection) SetModel(model gio.ListModel) {
	var arg0 *C.GtkSingleSelection
	var arg1 *C.GListModel

	arg0 = (*C.GtkSingleSelection)(s.Native())
	arg1 = (*C.GListModel)(model.Native())

	C.gtk_single_selection_set_model(arg0, arg1)
}

// SetSelected selects the item at the given position.
//
// If the list does not have an item at @position or
// GTK_INVALID_LIST_POSITION is given, the behavior depends on the value of
// the [property@Gtk.SingleSelection:autoselect] property: If it is set, no
// change will occur and the old item will stay selected. If it is unset,
// the selection will be unset and no item will be selected.
func (s singleSelection) SetSelected(position uint) {
	var arg0 *C.GtkSingleSelection
	var arg1 C.guint

	arg0 = (*C.GtkSingleSelection)(s.Native())
	arg1 = C.guint(position)

	C.gtk_single_selection_set_selected(arg0, arg1)
}
