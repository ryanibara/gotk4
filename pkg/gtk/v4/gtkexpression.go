// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cclosure_expression_get_type()), F: marshalCClosureExpressioner},
		{T: externglib.Type(C.gtk_closure_expression_get_type()), F: marshalClosureExpressioner},
		{T: externglib.Type(C.gtk_constant_expression_get_type()), F: marshalConstantExpressioner},
		{T: externglib.Type(C.gtk_expression_get_type()), F: marshalExpressioner},
		{T: externglib.Type(C.gtk_object_expression_get_type()), F: marshalObjectExpressioner},
		{T: externglib.Type(C.gtk_property_expression_get_type()), F: marshalPropertyExpressioner},
		{T: externglib.Type(C.gtk_expression_watch_get_type()), F: marshalExpressionWatch},
	})
}

// ExpressionNotify: callback called by gtk_expression_watch() when the
// expression value changes.
type ExpressionNotify func(userData cgo.Handle)

//export _gotk4_gtk4_ExpressionNotify
func _gotk4_gtk4_ExpressionNotify(arg0 C.gpointer) {
	v := gbox.Get(uintptr(arg0))
	if v == nil {
		panic(`callback not found`)
	}

	var userData cgo.Handle // out

	userData = (cgo.Handle)(unsafe.Pointer(arg0))

	fn := v.(ExpressionNotify)
	fn(userData)
}

// ValueDupExpression retrieves the GtkExpression stored inside the given value,
// and acquires a reference to it.
func ValueDupExpression(value *externglib.Value) *Expression {
	var _arg1 *C.GValue        // out
	var _cret *C.GtkExpression // in

	_arg1 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	_cret = C.gtk_value_dup_expression(_arg1)

	var _expression *Expression // out

	_expression = wrapExpression(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _expression
}

// ValueGetExpression retrieves the GtkExpression stored inside the given value.
func ValueGetExpression(value *externglib.Value) *Expression {
	var _arg1 *C.GValue        // out
	var _cret *C.GtkExpression // in

	_arg1 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	_cret = C.gtk_value_get_expression(_arg1)

	var _expression *Expression // out

	_expression = wrapExpression(externglib.Take(unsafe.Pointer(_cret)))

	return _expression
}

// ValueSetExpression stores the given GtkExpression inside value.
//
// The GValue will acquire a reference to the expression.
func ValueSetExpression(value *externglib.Value, expression Expressioner) {
	var _arg1 *C.GValue        // out
	var _arg2 *C.GtkExpression // out

	_arg1 = (*C.GValue)(unsafe.Pointer(&value.GValue))
	_arg2 = (*C.GtkExpression)(unsafe.Pointer((expression).(gextras.Nativer).Native()))

	C.gtk_value_set_expression(_arg1, _arg2)
}

// ValueTakeExpression stores the given GtkExpression inside value.
//
// This function transfers the ownership of the expression to the GValue.
func ValueTakeExpression(value *externglib.Value, expression Expressioner) {
	var _arg1 *C.GValue        // out
	var _arg2 *C.GtkExpression // out

	_arg1 = (*C.GValue)(unsafe.Pointer(&value.GValue))
	_arg2 = (*C.GtkExpression)(unsafe.Pointer((expression).(gextras.Nativer).Native()))

	C.gtk_value_take_expression(_arg1, _arg2)
}

// CClosureExpressioner describes CClosureExpression's methods.
type CClosureExpressioner interface {
	privateCClosureExpression()
}

// CClosureExpression: variant of GtkClosureExpression using a C closure.
type CClosureExpression struct {
	Expression
}

var (
	_ CClosureExpressioner = (*CClosureExpression)(nil)
	_ gextras.Nativer      = (*CClosureExpression)(nil)
)

func wrapCClosureExpression(obj *externglib.Object) *CClosureExpression {
	return &CClosureExpression{
		Expression: Expression{
			Object: obj,
		},
	}
}

func marshalCClosureExpressioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCClosureExpression(obj), nil
}

func (*CClosureExpression) privateCClosureExpression() {}

// ClosureExpressioner describes ClosureExpression's methods.
type ClosureExpressioner interface {
	privateClosureExpression()
}

// ClosureExpression: expression using a custom GClosure to compute the value
// from its parameters.
type ClosureExpression struct {
	Expression
}

var (
	_ ClosureExpressioner = (*ClosureExpression)(nil)
	_ gextras.Nativer     = (*ClosureExpression)(nil)
)

func wrapClosureExpression(obj *externglib.Object) *ClosureExpression {
	return &ClosureExpression{
		Expression: Expression{
			Object: obj,
		},
	}
}

func marshalClosureExpressioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapClosureExpression(obj), nil
}

func (*ClosureExpression) privateClosureExpression() {}

// ConstantExpressioner describes ConstantExpression's methods.
type ConstantExpressioner interface {
	// Value gets the value that a constant expression evaluates to.
	Value() *externglib.Value
}

// ConstantExpression: constant value in a GtkExpression.
type ConstantExpression struct {
	Expression
}

var (
	_ ConstantExpressioner = (*ConstantExpression)(nil)
	_ gextras.Nativer      = (*ConstantExpression)(nil)
)

func wrapConstantExpression(obj *externglib.Object) *ConstantExpression {
	return &ConstantExpression{
		Expression: Expression{
			Object: obj,
		},
	}
}

func marshalConstantExpressioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapConstantExpression(obj), nil
}

// NewConstantExpressionForValue creates an expression that always evaluates to
// the given value.
func NewConstantExpressionForValue(value *externglib.Value) *ConstantExpression {
	var _arg1 *C.GValue        // out
	var _cret *C.GtkExpression // in

	_arg1 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	_cret = C.gtk_constant_expression_new_for_value(_arg1)

	var _constantExpression *ConstantExpression // out

	_constantExpression = wrapConstantExpression(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _constantExpression
}

// Value gets the value that a constant expression evaluates to.
func (expression *ConstantExpression) Value() *externglib.Value {
	var _arg0 *C.GtkExpression // out
	var _cret *C.GValue        // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	_cret = C.gtk_constant_expression_get_value(_arg0)

	var _value *externglib.Value // out

	_value = externglib.ValueFromNative(unsafe.Pointer(_cret))

	return _value
}

// Expressioner describes Expression's methods.
type Expressioner interface {
	// Bind target's property named property to self.
	Bind(target *externglib.Object, property string, this_ *externglib.Object) *ExpressionWatch
	// Evaluate evaluates the given expression and on success stores the result
	// in value.
	Evaluate(this_ *externglib.Object, value *externglib.Value) bool
	// ValueType gets the GType that this expression evaluates to.
	ValueType() externglib.Type
	// IsStatic checks if the expression is static.
	IsStatic() bool
	// Ref acquires a reference on the given GtkExpression.
	ref() *Expression
	// Unref releases a reference on the given GtkExpression.
	unref()
}

// Expression: GtkExpression provides a way to describe references to values.
//
// An important aspect of expressions is that the value can be obtained from a
// source that is several steps away. For example, an expression may describe
// ‘the value of property A of object1, which is itself the value of a property
// of object2’. And object1 may not even exist yet at the time that the
// expression is created. This is contrast to GObject property bindings, which
// can only create direct connections between the properties of two objects that
// must both exist for the duration of the binding.
//
// An expression needs to be "evaluated" to obtain the value that it currently
// refers to. An evaluation always happens in the context of a current object
// called this (it mirrors the behavior of object-oriented languages), which may
// or may not influence the result of the evaluation. Use
// gtk.Expression.Evaluate() for evaluating an expression.
//
// Various methods for defining expressions exist, from simple constants via
// gtk.ConstantExpression.New to looking up properties in a GObject (even
// recursively) via gtk.PropertyExpression.New or providing custom functions to
// transform and combine expressions via gtk.ClosureExpression.New.
//
// Here is an example of a complex expression:
//
//      color_expr = gtk_property_expression_new (GTK_TYPE_LIST_ITEM,
//                                                NULL, "item");
//      expression = gtk_property_expression_new (GTK_TYPE_COLOR,
//                                                color_expr, "name");
//
//
// when evaluated with this being a GtkListItem, it will obtain the "item"
// property from the GtkListItem, and then obtain the "name" property from the
// resulting object (which is assumed to be of type GTK_TYPE_COLOR).
//
// A more concise way to describe this would be
//
//      this->item->name
//
//
// The most likely place where you will encounter expressions is in the context
// of list models and list widgets using them. For example, GtkDropDown is
// evaluating a GtkExpression to obtain strings from the items in its model that
// it can then use to match against the contents of its search entry.
// GtkStringFilter is using a GtkExpression for similar reasons.
//
// By default, expressions are not paying attention to changes and evaluation is
// just a snapshot of the current state at a given time. To get informed about
// changes, an expression needs to be "watched" via a gtk.ExpressionWatch, which
// will cause a callback to be called whenever the value of the expression may
// have changed; gtk.Expression.Watch() starts watching an expression, and
// gtk.ExpressionWatch.Unwatch() stops.
//
// Watches can be created for automatically updating the property of an object,
// similar to GObject's GBinding mechanism, by using gtk.Expression.Bind().
//
//
// GtkExpression in GObject properties
//
// In order to use a GtkExpression as a GObject property, you must use the
// gtk_param_spec_expression when creating a GParamSpec to install in the
// GObject class being defined; for instance:
//
//    obj_props[PROP_EXPRESSION] =
//      gtk_param_spec_expression ("expression",
//                                 "Expression",
//                                 "The expression used by the widget",
//                                 G_PARAM_READWRITE |
//                                 G_PARAM_STATIC_STRINGS |
//                                 G_PARAM_EXPLICIT_NOTIFY);
//
//
// When implementing the GObjectClass.set_property and GObjectClass.get_property
// virtual functions, you must use gtk_value_get_expression, to retrieve the
// stored GtkExpression from the GValue container, and gtk_value_set_expression,
// to store the GtkExpression into the GValue; for instance:
//
//      // in set_property()...
//      case PROP_EXPRESSION:
//        foo_widget_set_expression (foo, gtk_value_get_expression (value));
//        break;
//
//      // in get_property()...
//      case PROP_EXPRESSION:
//        gtk_value_set_expression (value, foo->expression);
//        break;
//
//
//
// GtkExpression in .ui files
//
// GtkBuilder has support for creating expressions. The syntax here can be used
// where a GtkExpression object is needed like in a <property> tag for an
// expression property, or in a <binding> tag to bind a property to an
// expression.
//
// To create an property expression, use the <lookup> element. It can have a
// type attribute to specify the object type, and a name attribute to specify
// the property to look up. The content of <lookup> can either be an element
// specfiying the expression to use the object, or a string that specifies the
// name of the object to use.
//
// Example:
//
//      <lookup name='search'>string_filter</lookup>
//
//
// To create a constant expression, use the <constant> element. If the type
// attribute is specified, the element content is interpreted as a value of that
// type. Otherwise, it is assumed to be an object. For instance:
//
//      <constant>string_filter</constant>
//      <constant type='gchararray'>Hello, world</constant>
//
//
// To create a closure expression, use the <closure> element. The type and
// function attributes specify what function to use for the closure, the content
// of the element contains the expressions for the parameters. For instance:
//
//      <closure type='gchararray' function='combine_args_somehow'>
//        <constant type='gchararray'>File size:</constant>
//        <lookup type='GFile' name='size'>myfile</lookup>
//      </closure>
type Expression struct {
	*externglib.Object
}

var (
	_ Expressioner    = (*Expression)(nil)
	_ gextras.Nativer = (*Expression)(nil)
)

func wrapExpression(obj *externglib.Object) *Expression {
	return &Expression{
		Object: obj,
	}
}

func marshalExpressioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapExpression(obj), nil
}

// Bind target's property named property to self.
//
// The value that self evaluates to is set via g_object_set() on target. This is
// repeated whenever self changes to ensure that the object's property stays
// synchronized with self.
//
// If self's evaluation fails, target's property is not updated. You can ensure
// that this doesn't happen by using a fallback expression.
//
// Note that this function takes ownership of self. If you want to keep it
// around, you should gtk.Expression.Ref() it beforehand.
func (self *Expression) Bind(target *externglib.Object, property string, this_ *externglib.Object) *ExpressionWatch {
	var _arg0 *C.GtkExpression      // out
	var _arg1 C.gpointer            // out
	var _arg2 *C.char               // out
	var _arg3 C.gpointer            // out
	var _cret *C.GtkExpressionWatch // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(self.Native()))
	_arg1 = C.gpointer(unsafe.Pointer(target.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(property)))
	_arg3 = C.gpointer(unsafe.Pointer(this_.Native()))

	_cret = C.gtk_expression_bind(_arg0, _arg1, _arg2, _arg3)

	var _expressionWatch *ExpressionWatch // out

	_expressionWatch = (*ExpressionWatch)(unsafe.Pointer(_cret))
	C.gtk_expression_watch_ref(_cret)
	runtime.SetFinalizer(_expressionWatch, func(v *ExpressionWatch) {
		C.gtk_expression_watch_unref((*C.GtkExpressionWatch)(unsafe.Pointer(v)))
	})

	return _expressionWatch
}

// Evaluate evaluates the given expression and on success stores the result in
// value.
//
// The GType of value will be the type given by gtk.Expression.GetValueType().
//
// It is possible that expressions cannot be evaluated - for example when the
// expression references objects that have been destroyed or set to NULL. In
// that case value will remain empty and FALSE will be returned.
func (self *Expression) Evaluate(this_ *externglib.Object, value *externglib.Value) bool {
	var _arg0 *C.GtkExpression // out
	var _arg1 C.gpointer       // out
	var _arg2 *C.GValue        // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(self.Native()))
	_arg1 = C.gpointer(unsafe.Pointer(this_.Native()))
	_arg2 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	_cret = C.gtk_expression_evaluate(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ValueType gets the GType that this expression evaluates to.
//
// This type is constant and will not change over the lifetime of this
// expression.
func (self *Expression) ValueType() externglib.Type {
	var _arg0 *C.GtkExpression // out
	var _cret C.GType          // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_expression_get_value_type(_arg0)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// IsStatic checks if the expression is static.
//
// A static expression will never change its result when
// gtk.Expression.Evaluate() is called on it with the same arguments.
//
// That means a call to gtk.Expression.Watch() is not necessary because it will
// never trigger a notify.
func (self *Expression) IsStatic() bool {
	var _arg0 *C.GtkExpression // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_expression_is_static(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Ref acquires a reference on the given GtkExpression.
func (self *Expression) ref() *Expression {
	var _arg0 *C.GtkExpression // out
	var _cret *C.GtkExpression // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_expression_ref(_arg0)

	var _expression *Expression // out

	_expression = wrapExpression(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _expression
}

// Unref releases a reference on the given GtkExpression.
//
// If the reference was the last, the resources associated to the self are
// freed.
func (self *Expression) unref() {
	var _arg0 *C.GtkExpression // out

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(self.Native()))

	C.gtk_expression_unref(_arg0)
}

// ObjectExpressioner describes ObjectExpression's methods.
type ObjectExpressioner interface {
	// Object gets the object that the expression evaluates to.
	Object() *externglib.Object
}

// ObjectExpression: GObject value in a GtkExpression.
type ObjectExpression struct {
	Expression
}

var (
	_ ObjectExpressioner = (*ObjectExpression)(nil)
	_ gextras.Nativer    = (*ObjectExpression)(nil)
)

func wrapObjectExpression(obj *externglib.Object) *ObjectExpression {
	return &ObjectExpression{
		Expression: Expression{
			Object: obj,
		},
	}
}

func marshalObjectExpressioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapObjectExpression(obj), nil
}

// NewObjectExpression creates an expression evaluating to the given object with
// a weak reference.
//
// Once the object is disposed, it will fail to evaluate.
//
// This expression is meant to break reference cycles.
//
// If you want to keep a reference to object, use gtk.ConstantExpression.New.
func NewObjectExpression(object *externglib.Object) *ObjectExpression {
	var _arg1 *C.GObject       // out
	var _cret *C.GtkExpression // in

	_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))

	_cret = C.gtk_object_expression_new(_arg1)

	var _objectExpression *ObjectExpression // out

	_objectExpression = wrapObjectExpression(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _objectExpression
}

// Object gets the object that the expression evaluates to.
func (expression *ObjectExpression) Object() *externglib.Object {
	var _arg0 *C.GtkExpression // out
	var _cret *C.GObject       // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	_cret = C.gtk_object_expression_get_object(_arg0)

	var _object *externglib.Object // out

	_object = externglib.Take(unsafe.Pointer(_cret))

	return _object
}

// PropertyExpressioner describes PropertyExpression's methods.
type PropertyExpressioner interface {
	// GetExpression gets the expression specifying the object of a property
	// expression.
	GetExpression() *Expression
}

// PropertyExpression: GObject property value in a GtkExpression.
type PropertyExpression struct {
	Expression
}

var (
	_ PropertyExpressioner = (*PropertyExpression)(nil)
	_ gextras.Nativer      = (*PropertyExpression)(nil)
)

func wrapPropertyExpression(obj *externglib.Object) *PropertyExpression {
	return &PropertyExpression{
		Expression: Expression{
			Object: obj,
		},
	}
}

func marshalPropertyExpressioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPropertyExpression(obj), nil
}

// NewPropertyExpression creates an expression that looks up a property via the
// given expression or the this argument when expression is NULL.
//
// If the resulting object conforms to this_type, its property named
// property_name will be queried. Otherwise, this expression's evaluation will
// fail.
//
// The given this_type must have a property with property_name.
func NewPropertyExpression(thisType externglib.Type, expression Expressioner, propertyName string) *PropertyExpression {
	var _arg1 C.GType          // out
	var _arg2 *C.GtkExpression // out
	var _arg3 *C.char          // out
	var _cret *C.GtkExpression // in

	_arg1 = C.GType(thisType)
	_arg2 = (*C.GtkExpression)(unsafe.Pointer((expression).(gextras.Nativer).Native()))
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(propertyName)))

	_cret = C.gtk_property_expression_new(_arg1, _arg2, _arg3)

	var _propertyExpression *PropertyExpression // out

	_propertyExpression = wrapPropertyExpression(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _propertyExpression
}

// GetExpression gets the expression specifying the object of a property
// expression.
func (expression *PropertyExpression) GetExpression() *Expression {
	var _arg0 *C.GtkExpression // out
	var _cret *C.GtkExpression // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	_cret = C.gtk_property_expression_get_expression(_arg0)

	var _ret *Expression // out

	_ret = wrapExpression(externglib.Take(unsafe.Pointer(_cret)))

	return _ret
}

// ExpressionWatch: opaque structure representing a watched GtkExpression.
//
// The contents of GtkExpressionWatch should only be accessed through the
// provided API.
type ExpressionWatch struct {
	native C.GtkExpressionWatch
}

func marshalExpressionWatch(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*ExpressionWatch)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (e *ExpressionWatch) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// Evaluate evaluates the watched expression and on success stores the result in
// value.
//
// This is equivalent to calling gtk.Expression.Evaluate() with the expression
// and this pointer originally used to create watch.
func (watch *ExpressionWatch) Evaluate(value *externglib.Value) bool {
	var _arg0 *C.GtkExpressionWatch // out
	var _arg1 *C.GValue             // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkExpressionWatch)(unsafe.Pointer(watch))
	_arg1 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	_cret = C.gtk_expression_watch_evaluate(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Ref acquires a reference on the given GtkExpressionWatch.
func (watch *ExpressionWatch) ref() *ExpressionWatch {
	var _arg0 *C.GtkExpressionWatch // out
	var _cret *C.GtkExpressionWatch // in

	_arg0 = (*C.GtkExpressionWatch)(unsafe.Pointer(watch))

	_cret = C.gtk_expression_watch_ref(_arg0)

	var _expressionWatch *ExpressionWatch // out

	_expressionWatch = (*ExpressionWatch)(unsafe.Pointer(_cret))
	C.gtk_expression_watch_ref(_cret)
	runtime.SetFinalizer(_expressionWatch, func(v *ExpressionWatch) {
		C.gtk_expression_watch_unref((*C.GtkExpressionWatch)(unsafe.Pointer(v)))
	})

	return _expressionWatch
}

// Unref releases a reference on the given GtkExpressionWatch.
//
// If the reference was the last, the resources associated to self are freed.
func (watch *ExpressionWatch) unref() {
	var _arg0 *C.GtkExpressionWatch // out

	_arg0 = (*C.GtkExpressionWatch)(unsafe.Pointer(watch))

	C.gtk_expression_watch_unref(_arg0)
}

// Unwatch stops watching an expression.
//
// See gtk.Expression.Watch() for how the watch was established.
func (watch *ExpressionWatch) Unwatch() {
	var _arg0 *C.GtkExpressionWatch // out

	_arg0 = (*C.GtkExpressionWatch)(unsafe.Pointer(watch))

	C.gtk_expression_watch_unwatch(_arg0)
}
