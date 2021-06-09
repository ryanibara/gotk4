// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gobject/v2"
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

// ParamSpecExpression creates a new `GParamSpec` instance for a property
// holding a `GtkExpression`.
//
// See `g_param_spec_internal()` for details on the property strings.
func ParamSpecExpression(name string, nick string, blurb string, flags gobject.ParamFlags) gobject.ParamSpec {
	var _arg1 *C.char
	var _arg2 *C.char
	var _arg3 *C.char
	var _arg4 C.GParamFlags

	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.char)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec

	cret = C.gtk_param_spec_expression(_arg1, _arg2, _arg3, _arg4)

	var _paramSpec gobject.ParamSpec

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(gobject.ParamSpec)

	return _paramSpec
}

// ValueDupExpression retrieves the `GtkExpression` stored inside the given
// `value`, and acquires a reference to it.
func ValueDupExpression(value **externglib.Value) Expression {
	var _arg1 *C.GValue

	_arg1 = (*C.GValue)(value.GValue)

	var _cret *C.GtkExpression

	cret = C.gtk_value_dup_expression(_arg1)

	var _expression Expression

	_expression = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(Expression)

	return _expression
}

// ValueGetExpression retrieves the `GtkExpression` stored inside the given
// `value`.
func ValueGetExpression(value **externglib.Value) Expression {
	var _arg1 *C.GValue

	_arg1 = (*C.GValue)(value.GValue)

	var _cret *C.GtkExpression

	cret = C.gtk_value_get_expression(_arg1)

	var _expression Expression

	_expression = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Expression)

	return _expression
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

	cret = C.gtk_expression_watch_evaluate(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Ref acquires a reference on the given `GtkExpressionWatch`.
func (w *ExpressionWatch) Ref() *ExpressionWatch {
	var _arg0 *C.GtkExpressionWatch

	_arg0 = (*C.GtkExpressionWatch)(unsafe.Pointer(w.Native()))

	var _cret *C.GtkExpressionWatch

	cret = C.gtk_expression_watch_ref(_arg0)

	var _expressionWatch *ExpressionWatch

	_expressionWatch = WrapExpressionWatch(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_expressionWatch, func(v *ExpressionWatch) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _expressionWatch
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