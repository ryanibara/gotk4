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

// GTypeStringSorter returns the GType for the type StringSorter.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeStringSorter() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "StringSorter").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalStringSorter)
	return gtype
}

// StringSorterOverrider contains methods that are overridable.
type StringSorterOverrider interface {
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

func classInitStringSorterer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

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
//    - expression (optional) to evaluate.
//
// The function returns the following values:
//
//    - stringSorter: new GtkStringSorter.
//
func NewStringSorter(expression Expressioner) *StringSorter {
	var _args [1]girepository.Argument

	if expression != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expression).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(expression).Native()))
	}

	_gret := girepository.MustFind("Gtk", "StringSorter").InvokeMethod("new_StringSorter", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
//    - expression (optional): GtkExpression, or NULL.
//
func (self *StringSorter) Expression() Expressioner {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "StringSorter").InvokeMethod("get_expression", _args[:], nil)
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

// IgnoreCase gets whether the sorter ignores case differences.
//
// The function returns the following values:
//
//    - ok: TRUE if self is ignoring case differences.
//
func (self *StringSorter) IgnoreCase() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "StringSorter").InvokeMethod("get_ignore_case", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
//    - expression (optional): GtkExpression, or NULL.
//
func (self *StringSorter) SetExpression(expression Expressioner) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if expression != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expression).Native()))
	}

	girepository.MustFind("Gtk", "StringSorter").InvokeMethod("set_expression", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(expression)
}

// SetIgnoreCase sets whether the sorter will ignore case differences.
//
// The function takes the following parameters:
//
//    - ignoreCase: TRUE to ignore case differences.
//
func (self *StringSorter) SetIgnoreCase(ignoreCase bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if ignoreCase {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "StringSorter").InvokeMethod("set_ignore_case", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(ignoreCase)
}
