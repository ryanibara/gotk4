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
// #include <glib-object.h>
import "C"

// GTypeConstraintTarget returns the GType for the type ConstraintTarget.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeConstraintTarget() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ConstraintTarget").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalConstraintTarget)
	return gtype
}

// GTypeConstraint returns the GType for the type Constraint.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeConstraint() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Constraint").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalConstraint)
	return gtype
}

// ConstraintTarget: GtkConstraintTarget interface is implemented by objects
// that can be used as source or target in GtkConstraints.
//
// Besides GtkWidget, it is also implemented by GtkConstraintGuide.
//
// ConstraintTarget wraps an interface. This means the user can get the
// underlying type by calling Cast().
type ConstraintTarget struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ConstraintTarget)(nil)
)

// ConstraintTargetter describes ConstraintTarget's interface methods.
type ConstraintTargetter interface {
	coreglib.Objector

	baseConstraintTarget() *ConstraintTarget
}

var _ ConstraintTargetter = (*ConstraintTarget)(nil)

func wrapConstraintTarget(obj *coreglib.Object) *ConstraintTarget {
	return &ConstraintTarget{
		Object: obj,
	}
}

func marshalConstraintTarget(p uintptr) (interface{}, error) {
	return wrapConstraintTarget(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *ConstraintTarget) baseConstraintTarget() *ConstraintTarget {
	return v
}

// BaseConstraintTarget returns the underlying base object.
func BaseConstraintTarget(obj ConstraintTargetter) *ConstraintTarget {
	return obj.baseConstraintTarget()
}

// ConstraintOverrider contains methods that are overridable.
type ConstraintOverrider interface {
}

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
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Constraint)(nil)
)

func classInitConstrainter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapConstraint(obj *coreglib.Object) *Constraint {
	return &Constraint{
		Object: obj,
	}
}

func marshalConstraint(p uintptr) (interface{}, error) {
	return wrapConstraint(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Constant retrieves the constant factor added to the source attributes' value.
//
// The function returns the following values:
//
//    - gdouble: constant factor.
//
func (constraint *Constraint) Constant() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_gret := girepository.MustFind("Gtk", "Constraint").InvokeMethod("get_constant", _args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(constraint)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// Multiplier retrieves the multiplication factor applied to the source
// attribute's value.
//
// The function returns the following values:
//
//    - gdouble: multiplication factor.
//
func (constraint *Constraint) Multiplier() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_gret := girepository.MustFind("Gtk", "Constraint").InvokeMethod("get_multiplier", _args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(constraint)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// Source retrieves the gtk.ConstraintTarget used as the source for the
// constraint.
//
// If the source is set to NULL at creation, the constraint will use the widget
// using the gtk.ConstraintLayout as the source.
//
// The function returns the following values:
//
//    - constraintTarget (optional): source of the constraint.
//
func (constraint *Constraint) Source() *ConstraintTarget {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_gret := girepository.MustFind("Gtk", "Constraint").InvokeMethod("get_source", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(constraint)

	var _constraintTarget *ConstraintTarget // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_constraintTarget = wrapConstraintTarget(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _constraintTarget
}

// Strength retrieves the strength of the constraint.
//
// The function returns the following values:
//
//    - gint: strength value.
//
func (constraint *Constraint) Strength() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_gret := girepository.MustFind("Gtk", "Constraint").InvokeMethod("get_strength", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(constraint)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// Target retrieves the gtk.ConstraintTarget used as the target for the
// constraint.
//
// If the targe is set to NULL at creation, the constraint will use the widget
// using the gtk.ConstraintLayout as the target.
//
// The function returns the following values:
//
//    - constraintTarget (optional): ConstraintTarget.
//
func (constraint *Constraint) Target() *ConstraintTarget {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_gret := girepository.MustFind("Gtk", "Constraint").InvokeMethod("get_target", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(constraint)

	var _constraintTarget *ConstraintTarget // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_constraintTarget = wrapConstraintTarget(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _constraintTarget
}

// IsAttached checks whether the constraint is attached to a
// gtk.ConstraintLayout, and it is contributing to the layout.
//
// The function returns the following values:
//
//    - ok: TRUE if the constraint is attached.
//
func (constraint *Constraint) IsAttached() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_gret := girepository.MustFind("Gtk", "Constraint").InvokeMethod("is_attached", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(constraint)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IsConstant checks whether the constraint describes a relation between an
// attribute on the gtk.Constraint:target and a constant value.
//
// The function returns the following values:
//
//    - ok: TRUE if the constraint is a constant relation.
//
func (constraint *Constraint) IsConstant() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_gret := girepository.MustFind("Gtk", "Constraint").InvokeMethod("is_constant", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(constraint)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IsRequired checks whether the constraint is a required relation for solving
// the constraint layout.
//
// The function returns the following values:
//
//    - ok: TRUE if the constraint is required.
//
func (constraint *Constraint) IsRequired() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_gret := girepository.MustFind("Gtk", "Constraint").InvokeMethod("is_required", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(constraint)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}
