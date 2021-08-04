// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_bool_filter_get_type()), F: marshalBoolFilterer},
	})
}

// BoolFilter: GtkBoolFilter evaluates a boolean GtkExpression to determine
// whether to include items.
type BoolFilter struct {
	Filter
}

func wrapBoolFilter(obj *externglib.Object) *BoolFilter {
	return &BoolFilter{
		Filter: Filter{
			Object: obj,
		},
	}
}

func marshalBoolFilterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapBoolFilter(obj), nil
}

// NewBoolFilter creates a new bool filter.
func NewBoolFilter(expression Expressioner) *BoolFilter {
	var _arg1 *C.GtkExpression // out
	var _cret *C.GtkBoolFilter // in

	if expression != nil {
		_arg1 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))
		C.g_object_ref(C.gpointer(expression.Native()))
	}

	_cret = C.gtk_bool_filter_new(_arg1)

	var _boolFilter *BoolFilter // out

	_boolFilter = wrapBoolFilter(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _boolFilter
}

// Expression gets the expression that the filter uses to evaluate if an item
// should be filtered.
func (self *BoolFilter) Expression() Expressioner {
	var _arg0 *C.GtkBoolFilter // out
	var _cret *C.GtkExpression // in

	_arg0 = (*C.GtkBoolFilter)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_bool_filter_get_expression(_arg0)

	var _expression Expressioner // out

	if _cret != nil {
		_expression = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Expressioner)
	}

	return _expression
}

// Invert returns whether the filter inverts the expression.
func (self *BoolFilter) Invert() bool {
	var _arg0 *C.GtkBoolFilter // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkBoolFilter)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_bool_filter_get_invert(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetExpression sets the expression that the filter uses to check if items
// should be filtered.
//
// The expression must have a value type of G_TYPE_BOOLEAN.
func (self *BoolFilter) SetExpression(expression Expressioner) {
	var _arg0 *C.GtkBoolFilter // out
	var _arg1 *C.GtkExpression // out

	_arg0 = (*C.GtkBoolFilter)(unsafe.Pointer(self.Native()))
	if expression != nil {
		_arg1 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))
	}

	C.gtk_bool_filter_set_expression(_arg0, _arg1)
}

// SetInvert sets whether the filter should invert the expression.
func (self *BoolFilter) SetInvert(invert bool) {
	var _arg0 *C.GtkBoolFilter // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkBoolFilter)(unsafe.Pointer(self.Native()))
	if invert {
		_arg1 = C.TRUE
	}

	C.gtk_bool_filter_set_invert(_arg0, _arg1)
}
