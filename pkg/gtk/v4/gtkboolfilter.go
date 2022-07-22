// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeBoolFilter = coreglib.Type(C.gtk_bool_filter_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeBoolFilter, F: marshalBoolFilter},
	})
}

// BoolFilterOverrider contains methods that are overridable.
type BoolFilterOverrider interface {
}

// BoolFilter: GtkBoolFilter evaluates a boolean GtkExpression to determine
// whether to include items.
type BoolFilter struct {
	_ [0]func() // equal guard
	Filter
}

var (
	_ coreglib.Objector = (*BoolFilter)(nil)
)

func initClassBoolFilter(gclass unsafe.Pointer, goval any) {
}

func wrapBoolFilter(obj *coreglib.Object) *BoolFilter {
	return &BoolFilter{
		Filter: Filter{
			Object: obj,
		},
	}
}

func marshalBoolFilter(p uintptr) (interface{}, error) {
	return wrapBoolFilter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBoolFilter creates a new bool filter.
//
// The function takes the following parameters:
//
//    - expression (optional) to evaluate or NULL for none.
//
// The function returns the following values:
//
//    - boolFilter: new GtkBoolFilter.
//
func NewBoolFilter(expression Expressioner) *BoolFilter {
	var _arg1 *C.GtkExpression // out
	var _cret *C.GtkBoolFilter // in

	if expression != nil {
		_arg1 = (*C.GtkExpression)(unsafe.Pointer(coreglib.InternObject(expression).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(expression).Native()))
	}

	_cret = C.gtk_bool_filter_new(_arg1)
	runtime.KeepAlive(expression)

	var _boolFilter *BoolFilter // out

	_boolFilter = wrapBoolFilter(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _boolFilter
}

// Expression gets the expression that the filter uses to evaluate if an item
// should be filtered.
//
// The function returns the following values:
//
//    - expression (optional): GtkExpression.
//
func (self *BoolFilter) Expression() Expressioner {
	var _arg0 *C.GtkBoolFilter // out
	var _cret *C.GtkExpression // in

	_arg0 = (*C.GtkBoolFilter)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_bool_filter_get_expression(_arg0)
	runtime.KeepAlive(self)

	var _expression Expressioner // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Expressioner)
				return ok
			})
			rv, ok := casted.(Expressioner)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Expressioner")
			}
			_expression = rv
		}
	}

	return _expression
}

// Invert returns whether the filter inverts the expression.
//
// The function returns the following values:
//
//    - ok: TRUE if the filter inverts.
//
func (self *BoolFilter) Invert() bool {
	var _arg0 *C.GtkBoolFilter // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkBoolFilter)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_bool_filter_get_invert(_arg0)
	runtime.KeepAlive(self)

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
//
// The function takes the following parameters:
//
//    - expression (optional): GtkExpression.
//
func (self *BoolFilter) SetExpression(expression Expressioner) {
	var _arg0 *C.GtkBoolFilter // out
	var _arg1 *C.GtkExpression // out

	_arg0 = (*C.GtkBoolFilter)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if expression != nil {
		_arg1 = (*C.GtkExpression)(unsafe.Pointer(coreglib.InternObject(expression).Native()))
	}

	C.gtk_bool_filter_set_expression(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(expression)
}

// SetInvert sets whether the filter should invert the expression.
//
// The function takes the following parameters:
//
//    - invert: TRUE to invert.
//
func (self *BoolFilter) SetInvert(invert bool) {
	var _arg0 *C.GtkBoolFilter // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkBoolFilter)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if invert {
		_arg1 = C.TRUE
	}

	C.gtk_bool_filter_set_invert(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(invert)
}
