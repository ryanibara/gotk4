// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_string_sorter_get_type()), F: marshalStringSorterer},
	})
}

// StringSorterer describes StringSorter's methods.
type StringSorterer interface {
	// Expression gets the expression that is evaluated to obtain strings from
	// items.
	Expression() *Expression
	// IgnoreCase gets whether the sorter ignores case differences.
	IgnoreCase() bool
	// SetExpression sets the expression that is evaluated to obtain strings
	// from items.
	SetExpression(expression Expressioner)
	// SetIgnoreCase sets whether the sorter will ignore case differences.
	SetIgnoreCase(ignoreCase bool)
}

// StringSorter: `GtkStringSorter` is a `GtkSorter` that compares strings.
//
// It does the comparison in a linguistically correct way using the current
// locale by normalizing Unicode strings and possibly case-folding them before
// performing the comparison.
//
// To obtain the strings to compare, this sorter evaluates a
// [class@Gtk.Expression].
type StringSorter struct {
	Sorter
}

var (
	_ StringSorterer  = (*StringSorter)(nil)
	_ gextras.Nativer = (*StringSorter)(nil)
)

func wrapStringSorter(obj *externglib.Object) StringSorterer {
	return &StringSorter{
		Sorter: Sorter{
			Object: obj,
		},
	}
}

func marshalStringSorterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStringSorter(obj), nil
}

// NewStringSorter creates a new string sorter that compares items using the
// given @expression.
//
// Unless an expression is set on it, this sorter will always compare items as
// invalid.
func NewStringSorter(expression Expressioner) *StringSorter {
	var _arg1 *C.GtkExpression   // out
	var _cret *C.GtkStringSorter // in

	_arg1 = (*C.GtkExpression)(unsafe.Pointer((expression).(gextras.Nativer).Native()))

	_cret = C.gtk_string_sorter_new(_arg1)

	var _stringSorter *StringSorter // out

	_stringSorter = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*StringSorter)

	return _stringSorter
}

// Expression gets the expression that is evaluated to obtain strings from
// items.
func (self *StringSorter) Expression() *Expression {
	var _arg0 *C.GtkStringSorter // out
	var _cret *C.GtkExpression   // in

	_arg0 = (*C.GtkStringSorter)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_string_sorter_get_expression(_arg0)

	var _expression *Expression // out

	_expression = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Expression)

	return _expression
}

// IgnoreCase gets whether the sorter ignores case differences.
func (self *StringSorter) IgnoreCase() bool {
	var _arg0 *C.GtkStringSorter // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkStringSorter)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_string_sorter_get_ignore_case(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetExpression sets the expression that is evaluated to obtain strings from
// items.
//
// The expression must have the type G_TYPE_STRING.
func (self *StringSorter) SetExpression(expression Expressioner) {
	var _arg0 *C.GtkStringSorter // out
	var _arg1 *C.GtkExpression   // out

	_arg0 = (*C.GtkStringSorter)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkExpression)(unsafe.Pointer((expression).(gextras.Nativer).Native()))

	C.gtk_string_sorter_set_expression(_arg0, _arg1)
}

// SetIgnoreCase sets whether the sorter will ignore case differences.
func (self *StringSorter) SetIgnoreCase(ignoreCase bool) {
	var _arg0 *C.GtkStringSorter // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkStringSorter)(unsafe.Pointer(self.Native()))
	if ignoreCase {
		_arg1 = C.TRUE
	}

	C.gtk_string_sorter_set_ignore_case(_arg0, _arg1)
}
