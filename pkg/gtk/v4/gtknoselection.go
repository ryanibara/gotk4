// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_no_selection_get_type()), F: marshalNoSelectioner},
	})
}

// NoSelectioner describes NoSelection's methods.
type NoSelectioner interface {
	// Model gets the model that @self is wrapping.
	Model() *gio.ListModel
	// SetModel sets the model that @self should wrap.
	SetModel(model gio.ListModeler)
}

// NoSelection: `GtkNoSelection` is a `GtkSelectionModel` that does not allow
// selecting anything.
//
// This model is meant to be used as a simple wrapper around a `GListModel` when
// a `GtkSelectionModel` is required.
type NoSelection struct {
	*externglib.Object

	SelectionModel
}

var (
	_ NoSelectioner   = (*NoSelection)(nil)
	_ gextras.Nativer = (*NoSelection)(nil)
)

func wrapNoSelection(obj *externglib.Object) *NoSelection {
	return &NoSelection{
		Object: obj,
		SelectionModel: SelectionModel{
			ListModel: gio.ListModel{
				Object: obj,
			},
		},
	}
}

func marshalNoSelectioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNoSelection(obj), nil
}

// NewNoSelection creates a new selection to handle @model.
func NewNoSelection(model gio.ListModeler) *NoSelection {
	var _arg1 *C.GListModel     // out
	var _cret *C.GtkNoSelection // in

	_arg1 = (*C.GListModel)(unsafe.Pointer((model).(gextras.Nativer).Native()))

	_cret = C.gtk_no_selection_new(_arg1)

	var _noSelection *NoSelection // out

	_noSelection = wrapNoSelection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _noSelection
}

// Model gets the model that @self is wrapping.
func (self *NoSelection) Model() *gio.ListModel {
	var _arg0 *C.GtkNoSelection // out
	var _cret *C.GListModel     // in

	_arg0 = (*C.GtkNoSelection)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_no_selection_get_model(_arg0)

	var _listModel *gio.ListModel // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_listModel = &gio.ListModel{
			Object: obj,
		}
	}

	return _listModel
}

// SetModel sets the model that @self should wrap.
//
// If @model is nil, this model will be empty.
func (self *NoSelection) SetModel(model gio.ListModeler) {
	var _arg0 *C.GtkNoSelection // out
	var _arg1 *C.GListModel     // out

	_arg0 = (*C.GtkNoSelection)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GListModel)(unsafe.Pointer((model).(gextras.Nativer).Native()))

	C.gtk_no_selection_set_model(_arg0, _arg1)
}
