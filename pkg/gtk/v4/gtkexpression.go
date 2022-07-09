// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk4_ExpressionNotify(gpointer);
// extern void callbackDelete(gpointer);
import "C"

// GTypeCClosureExpression returns the GType for the type CClosureExpression.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCClosureExpression() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "CClosureExpression").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalCClosureExpression)
	return gtype
}

// GTypeClosureExpression returns the GType for the type ClosureExpression.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeClosureExpression() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ClosureExpression").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalClosureExpression)
	return gtype
}

// GTypeConstantExpression returns the GType for the type ConstantExpression.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeConstantExpression() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ConstantExpression").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalConstantExpression)
	return gtype
}

// GTypeExpression returns the GType for the type Expression.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeExpression() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Expression").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalExpression)
	return gtype
}

// GTypeObjectExpression returns the GType for the type ObjectExpression.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeObjectExpression() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ObjectExpression").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalObjectExpression)
	return gtype
}

// GTypePropertyExpression returns the GType for the type PropertyExpression.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypePropertyExpression() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "PropertyExpression").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalPropertyExpression)
	return gtype
}

// GTypeExpressionWatch returns the GType for the type ExpressionWatch.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeExpressionWatch() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ExpressionWatch").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalExpressionWatch)
	return gtype
}

// ExpressionNotify: callback called by gtk_expression_watch() when the
// expression value changes.
type ExpressionNotify func()

//export _gotk4_gtk4_ExpressionNotify
func _gotk4_gtk4_ExpressionNotify(arg1 C.gpointer) {
	var fn ExpressionNotify
	{
		v := gbox.Get(uintptr(arg1))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ExpressionNotify)
	}

	fn()
}

// ValueDupExpression retrieves the GtkExpression stored inside the given value,
// and acquires a reference to it.
//
// The function takes the following parameters:
//
//    - value: GValue initialized with type GTK_TYPE_EXPRESSION.
//
// The function returns the following values:
//
//    - expression (optional): GtkExpression.
//
func ValueDupExpression(value *coreglib.Value) Expressioner {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(value.Native()))

	_gret := girepository.MustFind("Gtk", "value_dup_expression").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(value)

	var _expression Expressioner // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
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

// ValueGetExpression retrieves the GtkExpression stored inside the given value.
//
// The function takes the following parameters:
//
//    - value: GValue initialized with type GTK_TYPE_EXPRESSION.
//
// The function returns the following values:
//
//    - expression (optional): GtkExpression.
//
func ValueGetExpression(value *coreglib.Value) Expressioner {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(value.Native()))

	_gret := girepository.MustFind("Gtk", "value_get_expression").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(value)

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

// ValueSetExpression stores the given GtkExpression inside value.
//
// The GValue will acquire a reference to the expression.
//
// The function takes the following parameters:
//
//    - value: GValue initialized with type GTK_TYPE_EXPRESSION.
//    - expression: GtkExpression.
//
func ValueSetExpression(value *coreglib.Value, expression Expressioner) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(value.Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expression).Native()))

	girepository.MustFind("Gtk", "value_set_expression").Invoke(_args[:], nil)

	runtime.KeepAlive(value)
	runtime.KeepAlive(expression)
}

// ValueTakeExpression stores the given GtkExpression inside value.
//
// This function transfers the ownership of the expression to the GValue.
//
// The function takes the following parameters:
//
//    - value: GValue initialized with type GTK_TYPE_EXPRESSION.
//    - expression (optional): GtkExpression.
//
func ValueTakeExpression(value *coreglib.Value, expression Expressioner) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(value.Native()))
	if expression != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expression).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(expression).Native()))
	}

	girepository.MustFind("Gtk", "value_take_expression").Invoke(_args[:], nil)

	runtime.KeepAlive(value)
	runtime.KeepAlive(expression)
}

// CClosureExpression: variant of GtkClosureExpression using a C closure.
type CClosureExpression struct {
	_ [0]func() // equal guard
	Expression
}

var (
	_ Expressioner = (*CClosureExpression)(nil)
)

func wrapCClosureExpression(obj *coreglib.Object) *CClosureExpression {
	return &CClosureExpression{
		Expression: Expression{
			Object: obj,
		},
	}
}

func marshalCClosureExpression(p uintptr) (interface{}, error) {
	return wrapCClosureExpression(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ClosureExpression: expression using a custom GClosure to compute the value
// from its parameters.
type ClosureExpression struct {
	_ [0]func() // equal guard
	Expression
}

var (
	_ Expressioner = (*ClosureExpression)(nil)
)

func wrapClosureExpression(obj *coreglib.Object) *ClosureExpression {
	return &ClosureExpression{
		Expression: Expression{
			Object: obj,
		},
	}
}

func marshalClosureExpression(p uintptr) (interface{}, error) {
	return wrapClosureExpression(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConstantExpression: constant value in a GtkExpression.
type ConstantExpression struct {
	_ [0]func() // equal guard
	Expression
}

var (
	_ Expressioner = (*ConstantExpression)(nil)
)

func wrapConstantExpression(obj *coreglib.Object) *ConstantExpression {
	return &ConstantExpression{
		Expression: Expression{
			Object: obj,
		},
	}
}

func marshalConstantExpression(p uintptr) (interface{}, error) {
	return wrapConstantExpression(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewConstantExpressionForValue creates an expression that always evaluates to
// the given value.
//
// The function takes the following parameters:
//
//    - value: GValue.
//
// The function returns the following values:
//
//    - constantExpression: new GtkExpression.
//
func NewConstantExpressionForValue(value *coreglib.Value) *ConstantExpression {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(value.Native()))

	_gret := girepository.MustFind("Gtk", "ConstantExpression").InvokeMethod("new_ConstantExpression_for_value", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(value)

	var _constantExpression *ConstantExpression // out

	_constantExpression = wrapConstantExpression(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _constantExpression
}

// Value gets the value that a constant expression evaluates to.
//
// The function returns the following values:
//
//    - value: value.
//
func (expression *ConstantExpression) Value() *coreglib.Value {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expression).Native()))

	_gret := girepository.MustFind("Gtk", "ConstantExpression").InvokeMethod("get_value", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(expression)

	var _value *coreglib.Value // out

	_value = coreglib.ValueFromNative(unsafe.Pointer(_cret))

	return _value
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
//    <closure type='gchararray' function='combine_args_somehow'>
//      <constant type='gchararray'>File size:</constant>
//      <lookup type='GFile' name='size'>myfile</lookup>
//    </closure>.
type Expression struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Expression)(nil)
)

// Expressioner describes types inherited from class Expression.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Expressioner interface {
	coreglib.Objector
	baseExpression() *Expression
}

var _ Expressioner = (*Expression)(nil)

func wrapExpression(obj *coreglib.Object) *Expression {
	return &Expression{
		Object: obj,
	}
}

func marshalExpression(p uintptr) (interface{}, error) {
	return wrapExpression(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *Expression) baseExpression() *Expression {
	return self
}

// BaseExpression returns the underlying base object.
func BaseExpression(obj Expressioner) *Expression {
	return obj.baseExpression()
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
//
// The function takes the following parameters:
//
//    - target object to bind to.
//    - property: name of the property on target to bind to.
//    - this_ (optional): argument for the evaluation of self.
//
// The function returns the following values:
//
//    - expressionWatch: GtkExpressionWatch.
//
func (self *Expression) Bind(target *coreglib.Object, property string, this_ *coreglib.Object) *ExpressionWatch {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(self).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = C.gpointer(unsafe.Pointer(target.Native()))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(property)))
	defer C.free(unsafe.Pointer(_args[2]))
	*(*C.gpointer)(unsafe.Pointer(&_args[3])) = C.gpointer(unsafe.Pointer(this_.Native()))

	_gret := girepository.MustFind("Gtk", "Expression").InvokeMethod("bind", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(target)
	runtime.KeepAlive(property)
	runtime.KeepAlive(this_)

	var _expressionWatch *ExpressionWatch // out

	_expressionWatch = (*ExpressionWatch)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gtk_expression_watch_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_expressionWatch)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

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
//
// The function takes the following parameters:
//
//    - this_ (optional): argument for the evaluation.
//    - value: empty GValue.
//
// The function returns the following values:
//
//    - ok: TRUE if the expression could be evaluated.
//
func (self *Expression) Evaluate(this_ *coreglib.Object, value *coreglib.Value) bool {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = C.gpointer(unsafe.Pointer(this_.Native()))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(value.Native()))

	_gret := girepository.MustFind("Gtk", "Expression").InvokeMethod("evaluate", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(this_)
	runtime.KeepAlive(value)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IsStatic checks if the expression is static.
//
// A static expression will never change its result when
// gtk.Expression.Evaluate() is called on it with the same arguments.
//
// That means a call to gtk.Expression.Watch() is not necessary because it will
// never trigger a notify.
//
// The function returns the following values:
//
//    - ok: TRUE if the expression is static.
//
func (self *Expression) IsStatic() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "Expression").InvokeMethod("is_static", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Watch installs a watch for the given expression that calls the notify
// function whenever the evaluation of self may have changed.
//
// GTK cannot guarantee that the evaluation did indeed change when the notify
// gets invoked, but it guarantees the opposite: When it did in fact change, the
// notify will be invoked.
//
// The function takes the following parameters:
//
//    - this_ (optional): this argument to watch.
//    - notify: callback to invoke when the expression changes.
//
// The function returns the following values:
//
//    - expressionWatch: newly installed watch. Note that the only reference held
//      to the watch will be released when the watch is unwatched which can
//      happen automatically, and not just via gtk.ExpressionWatch.Unwatch(). You
//      should call gtk.ExpressionWatch.Ref() if you want to keep the watch
//      around.
//
func (self *Expression) Watch(this_ *coreglib.Object, notify ExpressionNotify) *ExpressionWatch {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = C.gpointer(unsafe.Pointer(this_.Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[2])) = (*[0]byte)(C._gotk4_gtk4_ExpressionNotify)
	_args[3] = C.gpointer(gbox.Assign(notify))
	_args[4] = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	_gret := girepository.MustFind("Gtk", "Expression").InvokeMethod("watch", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(this_)
	runtime.KeepAlive(notify)

	var _expressionWatch *ExpressionWatch // out

	_expressionWatch = (*ExpressionWatch)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gtk_expression_watch_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_expressionWatch)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _expressionWatch
}

// ObjectExpression: GObject value in a GtkExpression.
type ObjectExpression struct {
	_ [0]func() // equal guard
	Expression
}

var (
	_ Expressioner = (*ObjectExpression)(nil)
)

func wrapObjectExpression(obj *coreglib.Object) *ObjectExpression {
	return &ObjectExpression{
		Expression: Expression{
			Object: obj,
		},
	}
}

func marshalObjectExpression(p uintptr) (interface{}, error) {
	return wrapObjectExpression(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewObjectExpression creates an expression evaluating to the given object with
// a weak reference.
//
// Once the object is disposed, it will fail to evaluate.
//
// This expression is meant to break reference cycles.
//
// If you want to keep a reference to object, use gtk.ConstantExpression.New.
//
// The function takes the following parameters:
//
//    - object to watch.
//
// The function returns the following values:
//
//    - objectExpression: new GtkExpression.
//
func NewObjectExpression(object *coreglib.Object) *ObjectExpression {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(object.Native()))

	_gret := girepository.MustFind("Gtk", "ObjectExpression").InvokeMethod("new_ObjectExpression", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(object)

	var _objectExpression *ObjectExpression // out

	_objectExpression = wrapObjectExpression(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _objectExpression
}

// Object gets the object that the expression evaluates to.
//
// The function returns the following values:
//
//    - object (optional): object, or NULL.
//
func (expression *ObjectExpression) Object() *coreglib.Object {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expression).Native()))

	_gret := girepository.MustFind("Gtk", "ObjectExpression").InvokeMethod("get_object", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(expression)

	var _object *coreglib.Object // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_object = coreglib.Take(unsafe.Pointer(_cret))
	}

	return _object
}

// PropertyExpression: GObject property value in a GtkExpression.
type PropertyExpression struct {
	_ [0]func() // equal guard
	Expression
}

var (
	_ Expressioner = (*PropertyExpression)(nil)
)

func wrapPropertyExpression(obj *coreglib.Object) *PropertyExpression {
	return &PropertyExpression{
		Expression: Expression{
			Object: obj,
		},
	}
}

func marshalPropertyExpression(p uintptr) (interface{}, error) {
	return wrapPropertyExpression(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// GetExpression gets the expression specifying the object of a property
// expression.
//
// The function returns the following values:
//
//    - ret: object expression.
//
func (expression *PropertyExpression) GetExpression() Expressioner {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(expression).Native()))

	_gret := girepository.MustFind("Gtk", "PropertyExpression").InvokeMethod("get_expression", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(expression)

	var _ret Expressioner // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Expressioner is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Expressioner)
			return ok
		})
		rv, ok := casted.(Expressioner)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Expressioner")
		}
		_ret = rv
	}

	return _ret
}

// ExpressionWatch: opaque structure representing a watched GtkExpression.
//
// The contents of GtkExpressionWatch should only be accessed through the
// provided API.
//
// An instance of this type is always passed by reference.
type ExpressionWatch struct {
	*expressionWatch
}

// expressionWatch is the struct that's finalized.
type expressionWatch struct {
	native unsafe.Pointer
}

func marshalExpressionWatch(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &ExpressionWatch{&expressionWatch{(unsafe.Pointer)(b)}}, nil
}

// Evaluate evaluates the watched expression and on success stores the result in
// value.
//
// This is equivalent to calling gtk.Expression.Evaluate() with the expression
// and this pointer originally used to create watch.
//
// The function takes the following parameters:
//
//    - value: empty GValue to be set.
//
// The function returns the following values:
//
//    - ok: TRUE if the expression could be evaluated and value was set.
//
func (watch *ExpressionWatch) Evaluate(value *coreglib.Value) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(watch)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(value.Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(watch)
	runtime.KeepAlive(value)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Unwatch stops watching an expression.
//
// See gtk.Expression.Watch() for how the watch was established.
func (watch *ExpressionWatch) Unwatch() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(watch)))

	runtime.KeepAlive(watch)
}
