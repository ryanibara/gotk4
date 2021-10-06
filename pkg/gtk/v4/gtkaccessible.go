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
		{T: externglib.Type(C.gtk_accessible_get_type()), F: marshalAccessibler},
	})
}

func AccessiblePropertyInitValue(property AccessibleProperty, value *externglib.Value) {
	var _arg1 C.GtkAccessibleProperty // out
	var _arg2 *C.GValue               // out

	_arg1 = C.GtkAccessibleProperty(property)
	_arg2 = (*C.GValue)(unsafe.Pointer(value.Native()))

	C.gtk_accessible_property_init_value(_arg1, _arg2)
	runtime.KeepAlive(property)
	runtime.KeepAlive(value)
}

func AccessibleRelationInitValue(relation AccessibleRelation, value *externglib.Value) {
	var _arg1 C.GtkAccessibleRelation // out
	var _arg2 *C.GValue               // out

	_arg1 = C.GtkAccessibleRelation(relation)
	_arg2 = (*C.GValue)(unsafe.Pointer(value.Native()))

	C.gtk_accessible_relation_init_value(_arg1, _arg2)
	runtime.KeepAlive(relation)
	runtime.KeepAlive(value)
}

func AccessibleStateInitValue(state AccessibleState, value *externglib.Value) {
	var _arg1 C.GtkAccessibleState // out
	var _arg2 *C.GValue            // out

	_arg1 = C.GtkAccessibleState(state)
	_arg2 = (*C.GValue)(unsafe.Pointer(value.Native()))

	C.gtk_accessible_state_init_value(_arg1, _arg2)
	runtime.KeepAlive(state)
	runtime.KeepAlive(value)
}

// Accessible: GtkAccessible is an interface for describing UI elements for
// Assistive Technologies.
//
// Every accessible implementation has:
//
//    - a “role”, represented by a value of the gtk.AccessibleRole
//      enumeration
//    - an “attribute”, represented by a set of gtk.AccessibleState,
//      gtk.AccessibleProperty and gtk.AccessibleRelation values
//
// The role cannot be changed after instantiating a GtkAccessible
// implementation.
//
// The attributes are updated every time a UI element's state changes in a way
// that should be reflected by assistive technologies. For instance, if a
// GtkWidget visibility changes, the GTK_ACCESSIBLE_STATE_HIDDEN state will also
// change to reflect the gtk.Widget:visible property.
type Accessible struct {
	*externglib.Object
}

// Accessibler describes Accessible's abstract methods.
type Accessibler interface {
	externglib.Objector

	// AccessibleRole retrieves the GtkAccessibleRole for the given
	// GtkAccessible.
	AccessibleRole() AccessibleRole
	// ResetProperty resets the accessible property to its default value.
	ResetProperty(property AccessibleProperty)
	// ResetRelation resets the accessible relation to its default value.
	ResetRelation(relation AccessibleRelation)
	// ResetState resets the accessible state to its default value.
	ResetState(state AccessibleState)
	// UpdateProperty updates an array of accessible properties.
	UpdateProperty(properties []AccessibleProperty, values []externglib.Value)
	// UpdateRelation updates an array of accessible relations.
	UpdateRelation(relations []AccessibleRelation, values []externglib.Value)
	// UpdateState updates an array of accessible states.
	UpdateState(states []AccessibleState, values []externglib.Value)
}

var _ Accessibler = (*Accessible)(nil)

func wrapAccessible(obj *externglib.Object) *Accessible {
	return &Accessible{
		Object: obj,
	}
}

func marshalAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAccessible(obj), nil
}

// AccessibleRole retrieves the GtkAccessibleRole for the given GtkAccessible.
func (self *Accessible) AccessibleRole() AccessibleRole {
	var _arg0 *C.GtkAccessible    // out
	var _cret C.GtkAccessibleRole // in

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_accessible_get_accessible_role(_arg0)
	runtime.KeepAlive(self)

	var _accessibleRole AccessibleRole // out

	_accessibleRole = AccessibleRole(_cret)

	return _accessibleRole
}

// ResetProperty resets the accessible property to its default value.
//
// The function takes the following parameters:
//
//    - property: GtkAccessibleProperty.
//
func (self *Accessible) ResetProperty(property AccessibleProperty) {
	var _arg0 *C.GtkAccessible        // out
	var _arg1 C.GtkAccessibleProperty // out

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkAccessibleProperty(property)

	C.gtk_accessible_reset_property(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(property)
}

// ResetRelation resets the accessible relation to its default value.
//
// The function takes the following parameters:
//
//    - relation: GtkAccessibleRelation.
//
func (self *Accessible) ResetRelation(relation AccessibleRelation) {
	var _arg0 *C.GtkAccessible        // out
	var _arg1 C.GtkAccessibleRelation // out

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkAccessibleRelation(relation)

	C.gtk_accessible_reset_relation(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(relation)
}

// ResetState resets the accessible state to its default value.
//
// The function takes the following parameters:
//
//    - state: GtkAccessibleState.
//
func (self *Accessible) ResetState(state AccessibleState) {
	var _arg0 *C.GtkAccessible     // out
	var _arg1 C.GtkAccessibleState // out

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkAccessibleState(state)

	C.gtk_accessible_reset_state(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(state)
}

// UpdateProperty updates an array of accessible properties.
//
// This function should be called by GtkWidget types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
//
// The function takes the following parameters:
//
//    - properties: array of GtkAccessibleProperty.
//    - values: array of GValues, one for each property.
//
func (self *Accessible) UpdateProperty(properties []AccessibleProperty, values []externglib.Value) {
	var _arg0 *C.GtkAccessible         // out
	var _arg2 *C.GtkAccessibleProperty // out
	var _arg1 C.int
	var _arg3 *C.GValue // out

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(self.Native()))
	_arg1 = (C.int)(len(properties))
	_arg2 = (*C.GtkAccessibleProperty)(C.malloc(C.ulong(len(properties)) * C.ulong(C.sizeof_GtkAccessibleProperty)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GtkAccessibleProperty)(_arg2), len(properties))
		for i := range properties {
			out[i] = C.GtkAccessibleProperty(properties[i])
		}
	}
	_arg1 = (C.int)(len(values))
	_arg3 = (*C.GValue)(C.malloc(C.ulong(len(values)) * C.ulong(C.sizeof_GValue)))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice((*C.GValue)(_arg3), len(values))
		for i := range values {
			out[i] = *(*C.GValue)(unsafe.Pointer((&values[i]).Native()))
		}
	}

	C.gtk_accessible_update_property_value(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(properties)
	runtime.KeepAlive(values)
}

// UpdateRelation updates an array of accessible relations.
//
// This function should be called by GtkWidget types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
//
// The function takes the following parameters:
//
//    - relations: array of GtkAccessibleRelation.
//    - values: array of GValues, one for each relation.
//
func (self *Accessible) UpdateRelation(relations []AccessibleRelation, values []externglib.Value) {
	var _arg0 *C.GtkAccessible         // out
	var _arg2 *C.GtkAccessibleRelation // out
	var _arg1 C.int
	var _arg3 *C.GValue // out

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(self.Native()))
	_arg1 = (C.int)(len(relations))
	_arg2 = (*C.GtkAccessibleRelation)(C.malloc(C.ulong(len(relations)) * C.ulong(C.sizeof_GtkAccessibleRelation)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GtkAccessibleRelation)(_arg2), len(relations))
		for i := range relations {
			out[i] = C.GtkAccessibleRelation(relations[i])
		}
	}
	_arg1 = (C.int)(len(values))
	_arg3 = (*C.GValue)(C.malloc(C.ulong(len(values)) * C.ulong(C.sizeof_GValue)))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice((*C.GValue)(_arg3), len(values))
		for i := range values {
			out[i] = *(*C.GValue)(unsafe.Pointer((&values[i]).Native()))
		}
	}

	C.gtk_accessible_update_relation_value(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(relations)
	runtime.KeepAlive(values)
}

// UpdateState updates an array of accessible states.
//
// This function should be called by GtkWidget types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
//
// The function takes the following parameters:
//
//    - states: array of GtkAccessibleState.
//    - values: array of GValues, one for each state.
//
func (self *Accessible) UpdateState(states []AccessibleState, values []externglib.Value) {
	var _arg0 *C.GtkAccessible      // out
	var _arg2 *C.GtkAccessibleState // out
	var _arg1 C.int
	var _arg3 *C.GValue // out

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(self.Native()))
	_arg1 = (C.int)(len(states))
	_arg2 = (*C.GtkAccessibleState)(C.malloc(C.ulong(len(states)) * C.ulong(C.sizeof_GtkAccessibleState)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GtkAccessibleState)(_arg2), len(states))
		for i := range states {
			out[i] = C.GtkAccessibleState(states[i])
		}
	}
	_arg1 = (C.int)(len(values))
	_arg3 = (*C.GValue)(C.malloc(C.ulong(len(values)) * C.ulong(C.sizeof_GValue)))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice((*C.GValue)(_arg3), len(values))
		for i := range values {
			out[i] = *(*C.GValue)(unsafe.Pointer((&values[i]).Native()))
		}
	}

	C.gtk_accessible_update_state_value(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(states)
	runtime.KeepAlive(values)
}
