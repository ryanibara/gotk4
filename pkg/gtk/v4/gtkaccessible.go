// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_accessible_get_type()), F: marshalAccessible},
	})
}

func AccessiblePropertyInitValue(property AccessibleProperty, value *externglib.Value) {
	var arg1 C.GtkAccessibleProperty
	var arg2 *C.GValue

	arg1 = (C.GtkAccessibleProperty)(property)
	arg2 = (*C.GValue)(value.GValue)

	C.gtk_accessible_property_init_value(arg1, arg2)
}

func AccessibleRelationInitValue(relation AccessibleRelation, value *externglib.Value) {
	var arg1 C.GtkAccessibleRelation
	var arg2 *C.GValue

	arg1 = (C.GtkAccessibleRelation)(relation)
	arg2 = (*C.GValue)(value.GValue)

	C.gtk_accessible_relation_init_value(arg1, arg2)
}

func AccessibleStateInitValue(state AccessibleState, value *externglib.Value) {
	var arg1 C.GtkAccessibleState
	var arg2 *C.GValue

	arg1 = (C.GtkAccessibleState)(state)
	arg2 = (*C.GValue)(value.GValue)

	C.gtk_accessible_state_init_value(arg1, arg2)
}

// Accessible: `GtkAccessible` is an interface for describing UI elements for
// Assistive Technologies.
//
// Every accessible implementation has:
//
//    - a “role”, represented by a value of the [enum@Gtk.AccessibleRole]
//      enumeration
//    - an “attribute”, represented by a set of [enum@Gtk.AccessibleState],
//      [enum@Gtk.AccessibleProperty] and [enum@Gtk.AccessibleRelation] values
//
// The role cannot be changed after instantiating a `GtkAccessible`
// implementation.
//
// The attributes are updated every time a UI element's state changes in a way
// that should be reflected by assistive technologies. For instance, if a
// `GtkWidget` visibility changes, the GTK_ACCESSIBLE_STATE_HIDDEN state will
// also change to reflect the [property@Gtk.Widget:visible] property.
type Accessible interface {
	gextras.Objector

	// AccessibleRole retrieves the `GtkAccessibleRole` for the given
	// `GtkAccessible`.
	AccessibleRole() AccessibleRole
	// ResetProperty resets the accessible @property to its default value.
	ResetProperty(property AccessibleProperty)
	// ResetRelation resets the accessible @relation to its default value.
	ResetRelation(relation AccessibleRelation)
	// ResetState resets the accessible @state to its default value.
	ResetState(state AccessibleState)
	// UpdatePropertyValue updates an array of accessible properties.
	//
	// This function should be called by `GtkWidget` types whenever an
	// accessible property change must be communicated to assistive
	// technologies.
	//
	// This function is meant to be used by language bindings.
	UpdatePropertyValue(nProperties int, properties []AccessibleProperty, values []*externglib.Value)
	// UpdateRelationValue updates an array of accessible relations.
	//
	// This function should be called by `GtkWidget` types whenever an
	// accessible relation change must be communicated to assistive
	// technologies.
	//
	// This function is meant to be used by language bindings.
	UpdateRelationValue(nRelations int, relations []AccessibleRelation, values []*externglib.Value)
	// UpdateStateValue updates an array of accessible states.
	//
	// This function should be called by `GtkWidget` types whenever an
	// accessible state change must be communicated to assistive technologies.
	//
	// This function is meant to be used by language bindings.
	UpdateStateValue(nStates int, states []AccessibleState, values []*externglib.Value)
}

// accessible implements the Accessible interface.
type accessible struct {
	gextras.Objector
}

var _ Accessible = (*accessible)(nil)

// WrapAccessible wraps a GObject to a type that implements interface
// Accessible. It is primarily used internally.
func WrapAccessible(obj *externglib.Object) Accessible {
	return Accessible{
		Objector: obj,
	}
}

func marshalAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAccessible(obj), nil
}

// AccessibleRole retrieves the `GtkAccessibleRole` for the given
// `GtkAccessible`.
func (s accessible) AccessibleRole() AccessibleRole {
	var arg0 *C.GtkAccessible

	arg0 = (*C.GtkAccessible)(s.Native())

	ret := C.gtk_accessible_get_accessible_role(arg0)

	var ret0 AccessibleRole

	ret0 = AccessibleRole(ret)

	return ret0
}

// ResetProperty resets the accessible @property to its default value.
func (s accessible) ResetProperty(property AccessibleProperty) {
	var arg0 *C.GtkAccessible
	var arg1 C.GtkAccessibleProperty

	arg0 = (*C.GtkAccessible)(s.Native())
	arg1 = (C.GtkAccessibleProperty)(property)

	C.gtk_accessible_reset_property(arg0, arg1)
}

// ResetRelation resets the accessible @relation to its default value.
func (s accessible) ResetRelation(relation AccessibleRelation) {
	var arg0 *C.GtkAccessible
	var arg1 C.GtkAccessibleRelation

	arg0 = (*C.GtkAccessible)(s.Native())
	arg1 = (C.GtkAccessibleRelation)(relation)

	C.gtk_accessible_reset_relation(arg0, arg1)
}

// ResetState resets the accessible @state to its default value.
func (s accessible) ResetState(state AccessibleState) {
	var arg0 *C.GtkAccessible
	var arg1 C.GtkAccessibleState

	arg0 = (*C.GtkAccessible)(s.Native())
	arg1 = (C.GtkAccessibleState)(state)

	C.gtk_accessible_reset_state(arg0, arg1)
}

// UpdatePropertyValue updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an
// accessible property change must be communicated to assistive
// technologies.
//
// This function is meant to be used by language bindings.
func (s accessible) UpdatePropertyValue(nProperties int, properties []AccessibleProperty, values []*externglib.Value) {
	var arg0 *C.GtkAccessible
	var arg1 C.int
	var arg2 *C.GtkAccessibleProperty
	var arg3 *C.GValue

	arg0 = (*C.GtkAccessible)(s.Native())
	arg2 = (*C.GtkAccessibleProperty)(unsafe.Pointer(&properties[0]))
	arg1 = len(properties)
	defer runtime.KeepAlive(properties)
	{
		var dst []C.GValue
		ptr := C.malloc(C.sizeof_GValue * len(values))
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&dst))
		sliceHeader.Data = uintptr(unsafe.Pointer(ptr))
		sliceHeader.Len = len(values)
		sliceHeader.Cap = len(values)
		defer C.free(unsafe.Pointer(ptr))

		for i := 0; i < len(values); i++ {
			src := values[i]
			dst[i] = (*C.GValue)(src.GValue)
		}

		arg3 = (*C.GValue)(unsafe.Pointer(ptr))
		arg1 = len(values)
	}

	C.gtk_accessible_update_property_value(arg0, arg1, arg2, arg3)
}

// UpdateRelationValue updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an
// accessible relation change must be communicated to assistive
// technologies.
//
// This function is meant to be used by language bindings.
func (s accessible) UpdateRelationValue(nRelations int, relations []AccessibleRelation, values []*externglib.Value) {
	var arg0 *C.GtkAccessible
	var arg1 C.int
	var arg2 *C.GtkAccessibleRelation
	var arg3 *C.GValue

	arg0 = (*C.GtkAccessible)(s.Native())
	arg2 = (*C.GtkAccessibleRelation)(unsafe.Pointer(&relations[0]))
	arg1 = len(relations)
	defer runtime.KeepAlive(relations)
	{
		var dst []C.GValue
		ptr := C.malloc(C.sizeof_GValue * len(values))
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&dst))
		sliceHeader.Data = uintptr(unsafe.Pointer(ptr))
		sliceHeader.Len = len(values)
		sliceHeader.Cap = len(values)
		defer C.free(unsafe.Pointer(ptr))

		for i := 0; i < len(values); i++ {
			src := values[i]
			dst[i] = (*C.GValue)(src.GValue)
		}

		arg3 = (*C.GValue)(unsafe.Pointer(ptr))
		arg1 = len(values)
	}

	C.gtk_accessible_update_relation_value(arg0, arg1, arg2, arg3)
}

// UpdateStateValue updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an
// accessible state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (s accessible) UpdateStateValue(nStates int, states []AccessibleState, values []*externglib.Value) {
	var arg0 *C.GtkAccessible
	var arg1 C.int
	var arg2 *C.GtkAccessibleState
	var arg3 *C.GValue

	arg0 = (*C.GtkAccessible)(s.Native())
	arg2 = (*C.GtkAccessibleState)(unsafe.Pointer(&states[0]))
	arg1 = len(states)
	defer runtime.KeepAlive(states)
	{
		var dst []C.GValue
		ptr := C.malloc(C.sizeof_GValue * len(values))
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&dst))
		sliceHeader.Data = uintptr(unsafe.Pointer(ptr))
		sliceHeader.Len = len(values)
		sliceHeader.Cap = len(values)
		defer C.free(unsafe.Pointer(ptr))

		for i := 0; i < len(values); i++ {
			src := values[i]
			dst[i] = (*C.GValue)(src.GValue)
		}

		arg3 = (*C.GValue)(unsafe.Pointer(ptr))
		arg1 = len(values)
	}

	C.gtk_accessible_update_state_value(arg0, arg1, arg2, arg3)
}
