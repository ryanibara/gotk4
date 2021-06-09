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
		{T: externglib.Type(C.gtk_flatten_list_model_get_type()), F: marshalFlattenListModel},
	})
}

// FlattenListModel: `GtkFlattenListModel` is a list model that concatenates
// other list models.
//
// `GtkFlattenListModel` takes a list model containing list models, and flattens
// it into a single model.
type FlattenListModel interface {
	gextras.Objector
	gio.ListModel

	// Model gets the model set via gtk_flatten_list_model_set_model().
	Model() gio.ListModel
	// ModelForItem returns the model containing the item at the given position.
	ModelForItem(position uint) gio.ListModel
	// SetModel sets a new model to be flattened.
	SetModel(model gio.ListModel)
}

// flattenListModel implements the FlattenListModel interface.
type flattenListModel struct {
	gextras.Objector
	gio.ListModel
}

var _ FlattenListModel = (*flattenListModel)(nil)

// WrapFlattenListModel wraps a GObject to the right type. It is
// primarily used internally.
func WrapFlattenListModel(obj *externglib.Object) FlattenListModel {
	return FlattenListModel{
		Objector:      obj,
		gio.ListModel: gio.WrapListModel(obj),
	}
}

func marshalFlattenListModel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFlattenListModel(obj), nil
}

// NewFlattenListModel constructs a class FlattenListModel.
func NewFlattenListModel(model gio.ListModel) FlattenListModel {
	var _arg1 *C.GListModel

	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	var _cret C.GtkFlattenListModel

	cret = C.gtk_flatten_list_model_new(_arg1)

	var _flattenListModel FlattenListModel

	_flattenListModel = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(FlattenListModel)

	return _flattenListModel
}

// Model gets the model set via gtk_flatten_list_model_set_model().
func (s flattenListModel) Model() gio.ListModel {
	var _arg0 *C.GtkFlattenListModel

	_arg0 = (*C.GtkFlattenListModel)(unsafe.Pointer(s.Native()))

	var _cret *C.GListModel

	cret = C.gtk_flatten_list_model_get_model(_arg0)

	var _listModel gio.ListModel

	_listModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.ListModel)

	return _listModel
}

// ModelForItem returns the model containing the item at the given position.
func (s flattenListModel) ModelForItem(position uint) gio.ListModel {
	var _arg0 *C.GtkFlattenListModel
	var _arg1 C.guint

	_arg0 = (*C.GtkFlattenListModel)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(position)

	var _cret *C.GListModel

	cret = C.gtk_flatten_list_model_get_model_for_item(_arg0, _arg1)

	var _listModel gio.ListModel

	_listModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.ListModel)

	return _listModel
}

// SetModel sets a new model to be flattened.
func (s flattenListModel) SetModel(model gio.ListModel) {
	var _arg0 *C.GtkFlattenListModel
	var _arg1 *C.GListModel

	_arg0 = (*C.GtkFlattenListModel)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.gtk_flatten_list_model_set_model(_arg0, _arg1)
}