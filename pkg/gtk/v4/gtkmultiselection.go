// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_multi_selection_get_type()), F: marshalMultiSelectioner},
	})
}

// MultiSelection: GtkMultiSelection is a GtkSelectionModel that allows
// selecting multiple elements.
type MultiSelection struct {
	*externglib.Object

	SelectionModel
}

func wrapMultiSelection(obj *externglib.Object) *MultiSelection {
	return &MultiSelection{
		Object: obj,
		SelectionModel: SelectionModel{
			ListModel: gio.ListModel{
				Object: obj,
			},
		},
	}
}

func marshalMultiSelectioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMultiSelection(obj), nil
}

// NewMultiSelection creates a new selection to handle model.
func NewMultiSelection(model gio.ListModeller) *MultiSelection {
	var _arg1 *C.GListModel        // out
	var _cret *C.GtkMultiSelection // in

	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
		C.g_object_ref(C.gpointer(model.Native()))
	}

	_cret = C.gtk_multi_selection_new(_arg1)

	var _multiSelection *MultiSelection // out

	_multiSelection = wrapMultiSelection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _multiSelection
}

// Model returns the underlying model of self.
func (self *MultiSelection) Model() gio.ListModeller {
	var _arg0 *C.GtkMultiSelection // out
	var _cret *C.GListModel        // in

	_arg0 = (*C.GtkMultiSelection)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_multi_selection_get_model(_arg0)

	var _listModel gio.ListModeller // out

	_listModel = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.ListModeller)

	return _listModel
}

// SetModel sets the model that self should wrap.
//
// If model is NULL, self will be empty.
func (self *MultiSelection) SetModel(model gio.ListModeller) {
	var _arg0 *C.GtkMultiSelection // out
	var _arg1 *C.GListModel        // out

	_arg0 = (*C.GtkMultiSelection)(unsafe.Pointer(self.Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_multi_selection_set_model(_arg0, _arg1)
}
