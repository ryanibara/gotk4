// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_no_selection_get_type()), F: marshalNoSelection},
	})
}

// NoSelection: `GtkNoSelection` is a `GtkSelectionModel` that does not allow
// selecting anything.
//
// This model is meant to be used as a simple wrapper around a `GListModel` when
// a `GtkSelectionModel` is required.
type NoSelection interface {
	gextras.Objector
	gio.ListModel
	SelectionModel

	// Model gets the model that @self is wrapping.
	Model() gio.ListModel
	// SetModel sets the model that @self should wrap.
	//
	// If @model is nil, this model will be empty.
	SetModel(model gio.ListModel)
}

// noSelection implements the NoSelection class.
type noSelection struct {
	gextras.Objector
	gio.ListModel
	SelectionModel
}

var _ NoSelection = (*noSelection)(nil)

// WrapNoSelection wraps a GObject to the right type. It is
// primarily used internally.
func WrapNoSelection(obj *externglib.Object) NoSelection {
	return noSelection{
		Objector:       obj,
		gio.ListModel:  gio.WrapListModel(obj),
		SelectionModel: WrapSelectionModel(obj),
	}
}

func marshalNoSelection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNoSelection(obj), nil
}

// NewNoSelection constructs a class NoSelection.
func NewNoSelection(model gio.ListModel) NoSelection {
	var _arg1 *C.GListModel    // out
	var _cret C.GtkNoSelection // in

	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_no_selection_new(_arg1)

	var _noSelection NoSelection // out

	_noSelection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(NoSelection)

	return _noSelection
}

// Model gets the model that @self is wrapping.
func (s noSelection) Model() gio.ListModel {
	var _arg0 *C.GtkNoSelection // out
	var _cret *C.GListModel     // in

	_arg0 = (*C.GtkNoSelection)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_no_selection_get_model(_arg0)

	var _listModel gio.ListModel // out

	_listModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.ListModel)

	return _listModel
}

// SetModel sets the model that @self should wrap.
//
// If @model is nil, this model will be empty.
func (s noSelection) SetModel(model gio.ListModel) {
	var _arg0 *C.GtkNoSelection // out
	var _arg1 *C.GListModel     // out

	_arg0 = (*C.GtkNoSelection)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.gtk_no_selection_set_model(_arg0, _arg1)
}
