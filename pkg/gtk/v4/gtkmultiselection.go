// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
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
	_ [0]func() // equal guard
	*externglib.Object

	SelectionModel
}

var (
	_ externglib.Objector = (*MultiSelection)(nil)
)

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
	return wrapMultiSelection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMultiSelection creates a new selection to handle model.
//
// The function takes the following parameters:
//
//    - model (optional): GListModel to manage, or NULL.
//
// The function returns the following values:
//
//    - multiSelection: new GtkMultiSelection.
//
func NewMultiSelection(model gio.ListModeller) *MultiSelection {
	var _arg1 *C.GListModel        // out
	var _cret *C.GtkMultiSelection // in

	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
		C.g_object_ref(C.gpointer(model.Native()))
	}

	_cret = C.gtk_multi_selection_new(_arg1)
	runtime.KeepAlive(model)

	var _multiSelection *MultiSelection // out

	_multiSelection = wrapMultiSelection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _multiSelection
}

// Model returns the underlying model of self.
//
// The function returns the following values:
//
//    - listModel: underlying model.
//
func (self *MultiSelection) Model() gio.ListModeller {
	var _arg0 *C.GtkMultiSelection // out
	var _cret *C.GListModel        // in

	_arg0 = (*C.GtkMultiSelection)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_multi_selection_get_model(_arg0)
	runtime.KeepAlive(self)

	var _listModel gio.ListModeller // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.ListModeller is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gio.ListModeller)
			return ok
		})
		rv, ok := casted.(gio.ListModeller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.ListModeller")
		}
		_listModel = rv
	}

	return _listModel
}

// SetModel sets the model that self should wrap.
//
// If model is NULL, self will be empty.
//
// The function takes the following parameters:
//
//    - model (optional): GListModel to wrap.
//
func (self *MultiSelection) SetModel(model gio.ListModeller) {
	var _arg0 *C.GtkMultiSelection // out
	var _arg1 *C.GListModel        // out

	_arg0 = (*C.GtkMultiSelection)(unsafe.Pointer(self.Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_multi_selection_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}
