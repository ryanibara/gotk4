// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
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
		{T: externglib.Type(C.gtk_constraint_target_get_type()), F: marshalConstraintTargetter},
		{T: externglib.Type(C.gtk_constraint_get_type()), F: marshalConstrainter},
	})
}

// ConstraintTarget: GtkConstraintTarget interface is implemented by objects
// that can be used as source or target in GtkConstraints.
//
// Besides GtkWidget, it is also implemented by GtkConstraintGuide.
type ConstraintTarget struct {
	*externglib.Object
}

// ConstraintTargetter describes ConstraintTarget's abstract methods.
type ConstraintTargetter interface {
	externglib.Objector

	privateConstraintTarget()
}

var _ ConstraintTargetter = (*ConstraintTarget)(nil)

func wrapConstraintTarget(obj *externglib.Object) *ConstraintTarget {
	return &ConstraintTarget{
		Object: obj,
	}
}

func marshalConstraintTargetter(p uintptr) (interface{}, error) {
	return wrapConstraintTarget(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (*ConstraintTarget) privateConstraintTarget() {}

// Constraint: GtkConstraint describes a constraint between attributes of two
// widgets, expressed as a linear equation.
//
// The typical equation for a constraint is:
//
//      target.target_attr = source.source_attr × multiplier + constant
//
//
// Each GtkConstraint is part of a system that will be solved by a
// gtk.ConstraintLayout in order to allocate and position each child widget or
// guide.
//
// The source and target, as well as their attributes, of a GtkConstraint
// instance are immutable after creation.
type Constraint struct {
	*externglib.Object
}

func wrapConstraint(obj *externglib.Object) *Constraint {
	return &Constraint{
		Object: obj,
	}
}

func marshalConstrainter(p uintptr) (interface{}, error) {
	return wrapConstraint(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewConstraint creates a new constraint representing a relation between a
// layout attribute on a source and a layout attribute on a target.
//
// The function takes the following parameters:
//
//    - target of the constraint.
//    - targetAttribute: attribute of target to be set.
//    - relation equivalence between target_attribute and source_attribute.
//    - source of the constraint.
//    - sourceAttribute: attribute of source to be read.
//    - multiplier: multiplication factor to be applied to source_attribute.
//    - constant factor to be added to source_attribute.
//    - strength of the constraint.
//
func NewConstraint(target ConstraintTargetter, targetAttribute ConstraintAttribute, relation ConstraintRelation, source ConstraintTargetter, sourceAttribute ConstraintAttribute, multiplier, constant float64, strength int) *Constraint {
	var _arg1 C.gpointer               // out
	var _arg2 C.GtkConstraintAttribute // out
	var _arg3 C.GtkConstraintRelation  // out
	var _arg4 C.gpointer               // out
	var _arg5 C.GtkConstraintAttribute // out
	var _arg6 C.double                 // out
	var _arg7 C.double                 // out
	var _arg8 C.int                    // out
	var _cret *C.GtkConstraint         // in

	if target != nil {
		_arg1 = C.gpointer(unsafe.Pointer(target.Native()))
	}
	_arg2 = C.GtkConstraintAttribute(targetAttribute)
	_arg3 = C.GtkConstraintRelation(relation)
	if source != nil {
		_arg4 = C.gpointer(unsafe.Pointer(source.Native()))
	}
	_arg5 = C.GtkConstraintAttribute(sourceAttribute)
	_arg6 = C.double(multiplier)
	_arg7 = C.double(constant)
	_arg8 = C.int(strength)

	_cret = C.gtk_constraint_new(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
	runtime.KeepAlive(target)
	runtime.KeepAlive(targetAttribute)
	runtime.KeepAlive(relation)
	runtime.KeepAlive(source)
	runtime.KeepAlive(sourceAttribute)
	runtime.KeepAlive(multiplier)
	runtime.KeepAlive(constant)
	runtime.KeepAlive(strength)

	var _constraint *Constraint // out

	_constraint = wrapConstraint(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _constraint
}

// NewConstraintConstant creates a new constraint representing a relation
// between a layout attribute on a target and a constant value.
//
// The function takes the following parameters:
//
//    - target: the target of the constraint.
//    - targetAttribute: attribute of target to be set.
//    - relation equivalence between target_attribute and constant.
//    - constant factor to be set on target_attribute.
//    - strength of the constraint.
//
func NewConstraintConstant(target ConstraintTargetter, targetAttribute ConstraintAttribute, relation ConstraintRelation, constant float64, strength int) *Constraint {
	var _arg1 C.gpointer               // out
	var _arg2 C.GtkConstraintAttribute // out
	var _arg3 C.GtkConstraintRelation  // out
	var _arg4 C.double                 // out
	var _arg5 C.int                    // out
	var _cret *C.GtkConstraint         // in

	if target != nil {
		_arg1 = C.gpointer(unsafe.Pointer(target.Native()))
	}
	_arg2 = C.GtkConstraintAttribute(targetAttribute)
	_arg3 = C.GtkConstraintRelation(relation)
	_arg4 = C.double(constant)
	_arg5 = C.int(strength)

	_cret = C.gtk_constraint_new_constant(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(target)
	runtime.KeepAlive(targetAttribute)
	runtime.KeepAlive(relation)
	runtime.KeepAlive(constant)
	runtime.KeepAlive(strength)

	var _constraint *Constraint // out

	_constraint = wrapConstraint(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _constraint
}

// Constant retrieves the constant factor added to the source attributes' value.
func (constraint *Constraint) Constant() float64 {
	var _arg0 *C.GtkConstraint // out
	var _cret C.double         // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(constraint.Native()))

	_cret = C.gtk_constraint_get_constant(_arg0)
	runtime.KeepAlive(constraint)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Multiplier retrieves the multiplication factor applied to the source
// attribute's value.
func (constraint *Constraint) Multiplier() float64 {
	var _arg0 *C.GtkConstraint // out
	var _cret C.double         // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(constraint.Native()))

	_cret = C.gtk_constraint_get_multiplier(_arg0)
	runtime.KeepAlive(constraint)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Relation: order relation between the terms of the constraint.
func (constraint *Constraint) Relation() ConstraintRelation {
	var _arg0 *C.GtkConstraint        // out
	var _cret C.GtkConstraintRelation // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(constraint.Native()))

	_cret = C.gtk_constraint_get_relation(_arg0)
	runtime.KeepAlive(constraint)

	var _constraintRelation ConstraintRelation // out

	_constraintRelation = ConstraintRelation(_cret)

	return _constraintRelation
}

// Source retrieves the gtk.ConstraintTarget used as the source for the
// constraint.
//
// If the source is set to NULL at creation, the constraint will use the widget
// using the gtk.ConstraintLayout as the source.
func (constraint *Constraint) Source() ConstraintTargetter {
	var _arg0 *C.GtkConstraint       // out
	var _cret *C.GtkConstraintTarget // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(constraint.Native()))

	_cret = C.gtk_constraint_get_source(_arg0)
	runtime.KeepAlive(constraint)

	var _constraintTarget ConstraintTargetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(ConstraintTargetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.ConstraintTargetter")
			}
			_constraintTarget = rv
		}
	}

	return _constraintTarget
}

// SourceAttribute retrieves the attribute of the source to be read by the
// constraint.
func (constraint *Constraint) SourceAttribute() ConstraintAttribute {
	var _arg0 *C.GtkConstraint         // out
	var _cret C.GtkConstraintAttribute // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(constraint.Native()))

	_cret = C.gtk_constraint_get_source_attribute(_arg0)
	runtime.KeepAlive(constraint)

	var _constraintAttribute ConstraintAttribute // out

	_constraintAttribute = ConstraintAttribute(_cret)

	return _constraintAttribute
}

// Strength retrieves the strength of the constraint.
func (constraint *Constraint) Strength() int {
	var _arg0 *C.GtkConstraint // out
	var _cret C.int            // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(constraint.Native()))

	_cret = C.gtk_constraint_get_strength(_arg0)
	runtime.KeepAlive(constraint)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Target retrieves the gtk.ConstraintTarget used as the target for the
// constraint.
//
// If the targe is set to NULL at creation, the constraint will use the widget
// using the gtk.ConstraintLayout as the target.
func (constraint *Constraint) Target() ConstraintTargetter {
	var _arg0 *C.GtkConstraint       // out
	var _cret *C.GtkConstraintTarget // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(constraint.Native()))

	_cret = C.gtk_constraint_get_target(_arg0)
	runtime.KeepAlive(constraint)

	var _constraintTarget ConstraintTargetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(ConstraintTargetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.ConstraintTargetter")
			}
			_constraintTarget = rv
		}
	}

	return _constraintTarget
}

// TargetAttribute retrieves the attribute of the target to be set by the
// constraint.
func (constraint *Constraint) TargetAttribute() ConstraintAttribute {
	var _arg0 *C.GtkConstraint         // out
	var _cret C.GtkConstraintAttribute // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(constraint.Native()))

	_cret = C.gtk_constraint_get_target_attribute(_arg0)
	runtime.KeepAlive(constraint)

	var _constraintAttribute ConstraintAttribute // out

	_constraintAttribute = ConstraintAttribute(_cret)

	return _constraintAttribute
}

// IsAttached checks whether the constraint is attached to a
// gtk.ConstraintLayout, and it is contributing to the layout.
func (constraint *Constraint) IsAttached() bool {
	var _arg0 *C.GtkConstraint // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(constraint.Native()))

	_cret = C.gtk_constraint_is_attached(_arg0)
	runtime.KeepAlive(constraint)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsConstant checks whether the constraint describes a relation between an
// attribute on the gtk.Constraint:target and a constant value.
func (constraint *Constraint) IsConstant() bool {
	var _arg0 *C.GtkConstraint // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(constraint.Native()))

	_cret = C.gtk_constraint_is_constant(_arg0)
	runtime.KeepAlive(constraint)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsRequired checks whether the constraint is a required relation for solving
// the constraint layout.
func (constraint *Constraint) IsRequired() bool {
	var _arg0 *C.GtkConstraint // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(constraint.Native()))

	_cret = C.gtk_constraint_is_required(_arg0)
	runtime.KeepAlive(constraint)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
