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
		{T: externglib.Type(C.gtk_no_selection_get_type()), F: marshalNoSelectioner},
	})
}

// NoSelection: GtkNoSelection is a GtkSelectionModel that does not allow
// selecting anything.
//
// This model is meant to be used as a simple wrapper around a GListModel when a
// GtkSelectionModel is required.
type NoSelection struct {
	*externglib.Object

	SelectionModel
}

var (
	_ externglib.Objector = (*NoSelection)(nil)
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
	return wrapNoSelection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewNoSelection creates a new selection to handle model.
//
// The function takes the following parameters:
//
//    - model: GListModel to manage, or NULL.
//
func NewNoSelection(model gio.ListModeller) *NoSelection {
	var _arg1 *C.GListModel     // out
	var _cret *C.GtkNoSelection // in

	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
		C.g_object_ref(C.gpointer(model.Native()))
	}

	_cret = C.gtk_no_selection_new(_arg1)
	runtime.KeepAlive(model)

	var _noSelection *NoSelection // out

	_noSelection = wrapNoSelection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _noSelection
}

// Model gets the model that self is wrapping.
func (self *NoSelection) Model() gio.ListModeller {
	var _arg0 *C.GtkNoSelection // out
	var _cret *C.GListModel     // in

	_arg0 = (*C.GtkNoSelection)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_no_selection_get_model(_arg0)
	runtime.KeepAlive(self)

	var _listModel gio.ListModeller // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.ListModeller is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(gio.ListModeller)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.ListModeller")
		}
		_listModel = rv
	}

	return _listModel
}

// SetModel sets the model that self should wrap.
//
// If model is NULL, this model will be empty.
//
// The function takes the following parameters:
//
//    - model: GListModel to wrap.
//
func (self *NoSelection) SetModel(model gio.ListModeller) {
	var _arg0 *C.GtkNoSelection // out
	var _arg1 *C.GListModel     // out

	_arg0 = (*C.GtkNoSelection)(unsafe.Pointer(self.Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_no_selection_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}
