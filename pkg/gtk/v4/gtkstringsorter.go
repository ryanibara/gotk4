// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeStringSorter = coreglib.Type(C.gtk_string_sorter_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStringSorter, F: marshalStringSorter},
	})
}

// StringSorterOverrides contains methods that are overridable.
type StringSorterOverrides struct {
}

func defaultStringSorterOverrides(v *StringSorter) StringSorterOverrides {
	return StringSorterOverrides{}
}

// StringSorter: GtkStringSorter is a GtkSorter that compares strings.
//
// It does the comparison in a linguistically correct way using the current
// locale by normalizing Unicode strings and possibly case-folding them before
// performing the comparison.
//
// To obtain the strings to compare, this sorter evaluates a gtk.Expression.
type StringSorter struct {
	_ [0]func() // equal guard
	Sorter
}

var (
	_ coreglib.Objector = (*StringSorter)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*StringSorter, *StringSorterClass, StringSorterOverrides](
		GTypeStringSorter,
		initStringSorterClass,
		wrapStringSorter,
		defaultStringSorterOverrides,
	)
}

func initStringSorterClass(gclass unsafe.Pointer, overrides StringSorterOverrides, classInitFunc func(*StringSorterClass)) {
	if classInitFunc != nil {
		class := (*StringSorterClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapStringSorter(obj *coreglib.Object) *StringSorter {
	return &StringSorter{
		Sorter: Sorter{
			Object: obj,
		},
	}
}

func marshalStringSorter(p uintptr) (interface{}, error) {
	return wrapStringSorter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStringSorter creates a new string sorter that compares items using the
// given expression.
//
// Unless an expression is set on it, this sorter will always compare items as
// invalid.
//
// The function takes the following parameters:
//
//   - expression (optional) to evaluate.
//
// The function returns the following values:
//
//   - stringSorter: new GtkStringSorter.
//
func NewStringSorter(expression Expressioner) *StringSorter {
	var _arg1 *C.GtkExpression   // out
	var _cret *C.GtkStringSorter // in

	if expression != nil {
		_arg1 = (*C.GtkExpression)(unsafe.Pointer(coreglib.InternObject(expression).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(expression).Native()))
	}

	_cret = C.gtk_string_sorter_new(_arg1)
	runtime.KeepAlive(expression)

	var _stringSorter *StringSorter // out

	_stringSorter = wrapStringSorter(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _stringSorter
}

// Expression gets the expression that is evaluated to obtain strings from
// items.
//
// The function returns the following values:
//
//   - expression (optional): GtkExpression, or NULL.
//
func (self *StringSorter) Expression() Expressioner {
	var _arg0 *C.GtkStringSorter // out
	var _cret *C.GtkExpression   // in

	_arg0 = (*C.GtkStringSorter)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_string_sorter_get_expression(_arg0)
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

// IgnoreCase gets whether the sorter ignores case differences.
//
// The function returns the following values:
//
//   - ok: TRUE if self is ignoring case differences.
//
func (self *StringSorter) IgnoreCase() bool {
	var _arg0 *C.GtkStringSorter // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkStringSorter)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_string_sorter_get_ignore_case(_arg0)
	runtime.KeepAlive(self)

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
//
// The function takes the following parameters:
//
//   - expression (optional): GtkExpression, or NULL.
//
func (self *StringSorter) SetExpression(expression Expressioner) {
	var _arg0 *C.GtkStringSorter // out
	var _arg1 *C.GtkExpression   // out

	_arg0 = (*C.GtkStringSorter)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if expression != nil {
		_arg1 = (*C.GtkExpression)(unsafe.Pointer(coreglib.InternObject(expression).Native()))
	}

	C.gtk_string_sorter_set_expression(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(expression)
}

// SetIgnoreCase sets whether the sorter will ignore case differences.
//
// The function takes the following parameters:
//
//   - ignoreCase: TRUE to ignore case differences.
//
func (self *StringSorter) SetIgnoreCase(ignoreCase bool) {
	var _arg0 *C.GtkStringSorter // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkStringSorter)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if ignoreCase {
		_arg1 = C.TRUE
	}

	C.gtk_string_sorter_set_ignore_case(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ignoreCase)
}

// StringSorterClass: instance of this type is always passed by reference.
type StringSorterClass struct {
	*stringSorterClass
}

// stringSorterClass is the struct that's finalized.
type stringSorterClass struct {
	native *C.GtkStringSorterClass
}

func (s *StringSorterClass) ParentClass() *SorterClass {
	valptr := &s.native.parent_class
	var _v *SorterClass // out
	_v = (*SorterClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
