// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeFlattenListModel = coreglib.Type(C.gtk_flatten_list_model_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFlattenListModel, F: marshalFlattenListModel},
	})
}

// FlattenListModelOverrider contains methods that are overridable.
type FlattenListModelOverrider interface {
}

// FlattenListModel: GtkFlattenListModel is a list model that concatenates other
// list models.
//
// GtkFlattenListModel takes a list model containing list models, and flattens
// it into a single model.
type FlattenListModel struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.ListModel
}

var (
	_ coreglib.Objector = (*FlattenListModel)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeFlattenListModel,
		GoType:        reflect.TypeOf((*FlattenListModel)(nil)),
		InitClass:     initClassFlattenListModel,
		FinalizeClass: finalizeClassFlattenListModel,
	})
}

func initClassFlattenListModel(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ InitFlattenListModel(*FlattenListModelClass) }); ok {
		klass := (*FlattenListModelClass)(gextras.NewStructNative(gclass))
		goval.InitFlattenListModel(klass)
	}
}

func finalizeClassFlattenListModel(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ FinalizeFlattenListModel(*FlattenListModelClass) }); ok {
		klass := (*FlattenListModelClass)(gextras.NewStructNative(gclass))
		goval.FinalizeFlattenListModel(klass)
	}
}

func wrapFlattenListModel(obj *coreglib.Object) *FlattenListModel {
	return &FlattenListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalFlattenListModel(p uintptr) (interface{}, error) {
	return wrapFlattenListModel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFlattenListModel creates a new GtkFlattenListModel that flattens list.
//
// The function takes the following parameters:
//
//    - model (optional) to be flattened.
//
// The function returns the following values:
//
//    - flattenListModel: new GtkFlattenListModel.
//
func NewFlattenListModel(model gio.ListModeller) *FlattenListModel {
	var _arg1 *C.GListModel          // out
	var _cret *C.GtkFlattenListModel // in

	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(model).Native()))
	}

	_cret = C.gtk_flatten_list_model_new(_arg1)
	runtime.KeepAlive(model)

	var _flattenListModel *FlattenListModel // out

	_flattenListModel = wrapFlattenListModel(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _flattenListModel
}

// Model gets the model set via gtk_flatten_list_model_set_model().
//
// The function returns the following values:
//
//    - listModel (optional): model flattened by self.
//
func (self *FlattenListModel) Model() *gio.ListModel {
	var _arg0 *C.GtkFlattenListModel // out
	var _cret *C.GListModel          // in

	_arg0 = (*C.GtkFlattenListModel)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_flatten_list_model_get_model(_arg0)
	runtime.KeepAlive(self)

	var _listModel *gio.ListModel // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_listModel = &gio.ListModel{
				Object: obj,
			}
		}
	}

	return _listModel
}

// ModelForItem returns the model containing the item at the given position.
//
// The function takes the following parameters:
//
//    - position: position.
//
// The function returns the following values:
//
//    - listModel: model containing the item at position.
//
func (self *FlattenListModel) ModelForItem(position uint) *gio.ListModel {
	var _arg0 *C.GtkFlattenListModel // out
	var _arg1 C.guint                // out
	var _cret *C.GListModel          // in

	_arg0 = (*C.GtkFlattenListModel)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(position)

	_cret = C.gtk_flatten_list_model_get_model_for_item(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)

	var _listModel *gio.ListModel // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_listModel = &gio.ListModel{
			Object: obj,
		}
	}

	return _listModel
}

// SetModel sets a new model to be flattened.
//
// The function takes the following parameters:
//
//    - model (optional): new model or NULL.
//
func (self *FlattenListModel) SetModel(model gio.ListModeller) {
	var _arg0 *C.GtkFlattenListModel // out
	var _arg1 *C.GListModel          // out

	_arg0 = (*C.GtkFlattenListModel)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	C.gtk_flatten_list_model_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}

// FlattenListModelClass: instance of this type is always passed by reference.
type FlattenListModelClass struct {
	*flattenListModelClass
}

// flattenListModelClass is the struct that's finalized.
type flattenListModelClass struct {
	native *C.GtkFlattenListModelClass
}
