// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_slice_list_model_get_type()), F: marshalSliceListModel},
	})
}

// SliceListModel: `GtkSliceListModel` is a list model that presents a slice of
// another model.
//
// This is useful when implementing paging by setting the size to the number of
// elements per page and updating the offset whenever a different page is
// opened.
type SliceListModel interface {
	gextras.Objector

	Offset() uint

	Size() uint

	SetOffsetSliceListModel(offset uint)

	SetSizeSliceListModel(size uint)
}

// sliceListModel implements the SliceListModel class.
type sliceListModel struct {
	gextras.Objector
}

// WrapSliceListModel wraps a GObject to the right type. It is
// primarily used internally.
func WrapSliceListModel(obj *externglib.Object) SliceListModel {
	return sliceListModel{
		Objector: obj,
	}
}

func marshalSliceListModel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSliceListModel(obj), nil
}

func (s sliceListModel) Offset() uint {
	var _arg0 *C.GtkSliceListModel // out
	var _cret C.guint              // in

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_slice_list_model_get_offset(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (s sliceListModel) Size() uint {
	var _arg0 *C.GtkSliceListModel // out
	var _cret C.guint              // in

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_slice_list_model_get_size(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (s sliceListModel) SetOffsetSliceListModel(offset uint) {
	var _arg0 *C.GtkSliceListModel // out
	var _arg1 C.guint              // out

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(offset)

	C.gtk_slice_list_model_set_offset(_arg0, _arg1)
}

func (s sliceListModel) SetSizeSliceListModel(size uint) {
	var _arg0 *C.GtkSliceListModel // out
	var _arg1 C.guint              // out

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(size)

	C.gtk_slice_list_model_set_size(_arg0, _arg1)
}
