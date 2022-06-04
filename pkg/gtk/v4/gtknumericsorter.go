// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtknumericsorter.go.
var GTypeNumericSorter = coreglib.Type(C.gtk_numeric_sorter_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeNumericSorter, F: marshalNumericSorter},
	})
}

// NumericSorterOverrider contains methods that are overridable.
type NumericSorterOverrider interface {
}

// NumericSorter: GtkNumericSorter is a GtkSorter that compares numbers.
//
// To obtain the numbers to compare, this sorter evaluates a gtk.Expression.
type NumericSorter struct {
	_ [0]func() // equal guard
	Sorter
}

var (
	_ coreglib.Objector = (*NumericSorter)(nil)
)

func classInitNumericSorterer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapNumericSorter(obj *coreglib.Object) *NumericSorter {
	return &NumericSorter{
		Sorter: Sorter{
			Object: obj,
		},
	}
}

func marshalNumericSorter(p uintptr) (interface{}, error) {
	return wrapNumericSorter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewNumericSorter creates a new numeric sorter using the given expression.
//
// Smaller numbers will be sorted first. You can call
// gtk.NumericSorter.SetSortOrder() to change this.
//
// The function takes the following parameters:
//
//    - expression (optional) to evaluate.
//
// The function returns the following values:
//
//    - numericSorter: new GtkNumericSorter.
//
func NewNumericSorter(expression Expressioner) *NumericSorter {
	var _args [1]girepository.Argument

	if expression != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expression).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(expression).Native()))
	}

	_gret := girepository.MustFind("Gtk", "NumericSorter").InvokeMethod("new_NumericSorter", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(expression)

	var _numericSorter *NumericSorter // out

	_numericSorter = wrapNumericSorter(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _numericSorter
}

// Expression gets the expression that is evaluated to obtain numbers from
// items.
//
// The function returns the following values:
//
//    - expression (optional): GtkExpression, or NULL.
//
func (self *NumericSorter) Expression() Expressioner {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "NumericSorter").InvokeMethod("get_expression", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _expression Expressioner // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
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

// SetExpression sets the expression that is evaluated to obtain numbers from
// items.
//
// Unless an expression is set on self, the sorter will always compare items as
// invalid.
//
// The expression must have a return type that can be compared numerically, such
// as G_TYPE_INT or G_TYPE_DOUBLE.
//
// The function takes the following parameters:
//
//    - expression (optional): GtkExpression, or NULL.
//
func (self *NumericSorter) SetExpression(expression Expressioner) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if expression != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expression).Native()))
	}

	girepository.MustFind("Gtk", "NumericSorter").InvokeMethod("set_expression", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(expression)
}
