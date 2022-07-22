// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeNoSelection = coreglib.Type(C.gtk_no_selection_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeNoSelection, F: marshalNoSelection},
	})
}

// NoSelectionOverrider contains methods that are overridable.
type NoSelectionOverrider interface {
}

// NoSelection: GtkNoSelection is a GtkSelectionModel that does not allow
// selecting anything.
//
// This model is meant to be used as a simple wrapper around a GListModel when a
// GtkSelectionModel is required.
type NoSelection struct {
	_ [0]func() // equal guard
	*coreglib.Object

	SelectionModel
}

var (
	_ coreglib.Objector = (*NoSelection)(nil)
)

func initClassNoSelection(gclass unsafe.Pointer, goval any) {
}

func wrapNoSelection(obj *coreglib.Object) *NoSelection {
	return &NoSelection{
		Object: obj,
		SelectionModel: SelectionModel{
			ListModel: gio.ListModel{
				Object: obj,
			},
		},
	}
}

func marshalNoSelection(p uintptr) (interface{}, error) {
	return wrapNoSelection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewNoSelection creates a new selection to handle model.
//
// The function takes the following parameters:
//
//    - model (optional): GListModel to manage, or NULL.
//
// The function returns the following values:
//
//    - noSelection: new GtkNoSelection.
//
func NewNoSelection(model gio.ListModeller) *NoSelection {
	var _arg1 *C.GListModel     // out
	var _cret *C.GtkNoSelection // in

	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(model).Native()))
	}

	_cret = C.gtk_no_selection_new(_arg1)
	runtime.KeepAlive(model)

	var _noSelection *NoSelection // out

	_noSelection = wrapNoSelection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _noSelection
}

// Model gets the model that self is wrapping.
//
// The function returns the following values:
//
//    - listModel: model being wrapped.
//
func (self *NoSelection) Model() *gio.ListModel {
	var _arg0 *C.GtkNoSelection // out
	var _cret *C.GListModel     // in

	_arg0 = (*C.GtkNoSelection)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_no_selection_get_model(_arg0)
	runtime.KeepAlive(self)

	var _listModel *gio.ListModel // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_listModel = &gio.ListModel{
			Object: obj,
		}
	}

	return _listModel
}

// SetModel sets the model that self should wrap.
//
// If model is NULL, this model will be empty.
//
// The function takes the following parameters:
//
//    - model (optional): GListModel to wrap.
//
func (self *NoSelection) SetModel(model gio.ListModeller) {
	var _arg0 *C.GtkNoSelection // out
	var _arg1 *C.GListModel     // out

	_arg0 = (*C.GtkNoSelection)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	C.gtk_no_selection_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}
