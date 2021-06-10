// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_expression_watch_get_type()), F: marshalExpressionWatch},
	})
}

// ExpressionNotify: callback called by gtk_expression_watch() when the
// expression value changes.
type ExpressionNotify func()

//export gotk4_ExpressionNotify
func gotk4_ExpressionNotify(arg0 C.gpointer) {
	v := box.Get(uintptr(arg0))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(ExpressionNotify)
	fn()
}

// ValueSetExpression stores the given `GtkExpression` inside `value`.
//
// The `GValue` will acquire a reference to the `expression`.
func ValueSetExpression(value **externglib.Value, expression Expression) {
	var _arg1 *C.GValue
	var _arg2 *C.GtkExpression

	_arg1 = (*C.GValue)(value.GValue)
	_arg2 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	C.gtk_value_set_expression(_arg1, _arg2)
}

// ValueTakeExpression stores the given `GtkExpression` inside `value`.
//
// This function transfers the ownership of the `expression` to the `GValue`.
func ValueTakeExpression(value **externglib.Value, expression Expression) {
	var _arg1 *C.GValue
	var _arg2 *C.GtkExpression

	_arg1 = (*C.GValue)(value.GValue)
	_arg2 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	C.gtk_value_take_expression(_arg1, _arg2)
}

// ExpressionWatch: an opaque structure representing a watched `GtkExpression`.
//
// The contents of `GtkExpressionWatch` should only be accessed through the
// provided API.
type ExpressionWatch struct {
	native C.GtkExpressionWatch
}

// WrapExpressionWatch wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapExpressionWatch(ptr unsafe.Pointer) *ExpressionWatch {
	if ptr == nil {
		return nil
	}

	return (*ExpressionWatch)(ptr)
}

func marshalExpressionWatch(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapExpressionWatch(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (e *ExpressionWatch) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// Evaluate evaluates the watched expression and on success stores the result in
// `value`.
//
// This is equivalent to calling [method@Gtk.Expression.evaluate] with the
// expression and this pointer originally used to create `watch`.
func (w *ExpressionWatch) Evaluate(value **externglib.Value) bool {
	var _arg0 *C.GtkExpressionWatch
	var _arg1 *C.GValue

	_arg0 = (*C.GtkExpressionWatch)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GValue)(value.GValue)

	var _cret C.gboolean

	_cret = C.gtk_expression_watch_evaluate(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Unref releases a reference on the given `GtkExpressionWatch`.
//
// If the reference was the last, the resources associated to `self` are freed.
func (w *ExpressionWatch) Unref() {
	var _arg0 *C.GtkExpressionWatch

	_arg0 = (*C.GtkExpressionWatch)(unsafe.Pointer(w.Native()))

	C.gtk_expression_watch_unref(_arg0)
}

// Unwatch stops watching an expression.
//
// See [method@Gtk.Expression.watch] for how the watch was established.
func (w *ExpressionWatch) Unwatch() {
	var _arg0 *C.GtkExpressionWatch

	_arg0 = (*C.GtkExpressionWatch)(unsafe.Pointer(w.Native()))

	C.gtk_expression_watch_unwatch(_arg0)
}
