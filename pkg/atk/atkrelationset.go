// Code generated by girgen. DO NOT EDIT.

package atk

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

// GTypeRelationSet returns the GType for the type RelationSet.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeRelationSet() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Atk", "RelationSet").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalRelationSet)
	return gtype
}

// RelationSetOverrider contains methods that are overridable.
type RelationSetOverrider interface {
}

// RelationSet held by an object establishes its relationships with objects
// beyond the normal "parent/child" hierarchical relationships that all user
// interface objects have. AtkRelationSets establish whether objects are
// labelled or controlled by other components, share group membership with other
// components (for instance within a radio-button group), or share content which
// "flows" between them, among other types of possible relationships.
type RelationSet struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*RelationSet)(nil)
)

func classInitRelationSetter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRelationSet(obj *coreglib.Object) *RelationSet {
	return &RelationSet{
		Object: obj,
	}
}

func marshalRelationSet(p uintptr) (interface{}, error) {
	return wrapRelationSet(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewRelationSet creates a new empty relation set.
//
// The function returns the following values:
//
//    - relationSet: new RelationSet.
//
func NewRelationSet() *RelationSet {
	_info := girepository.MustFind("Atk", "RelationSet")
	_gret := _info.InvokeClassMethod("new_RelationSet", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _relationSet *RelationSet // out

	_relationSet = wrapRelationSet(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _relationSet
}

// Add a new relation to the current relation set if it is not already present.
// This function ref's the AtkRelation so the caller of this function should
// unref it to ensure that it will be destroyed when the AtkRelationSet is
// destroyed.
//
// The function takes the following parameters:
//
//    - relation: Relation.
//
func (set *RelationSet) Add(relation *Relation) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(set).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(relation).Native()))

	_info := girepository.MustFind("Atk", "RelationSet")
	_info.InvokeClassMethod("add", _args[:], nil)

	runtime.KeepAlive(set)
	runtime.KeepAlive(relation)
}

// NRelations determines the number of relations in a relation set.
//
// The function returns the following values:
//
//    - gint: integer representing the number of relations in the set.
//
func (set *RelationSet) NRelations() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(set).Native()))

	_info := girepository.MustFind("Atk", "RelationSet")
	_gret := _info.InvokeClassMethod("get_n_relations", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(set)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// Relation determines the relation at the specified position in the relation
// set.
//
// The function takes the following parameters:
//
//    - i: gint representing a position in the set, starting from 0.
//
// The function returns the following values:
//
//    - relation which is the relation at position i in the set.
//
func (set *RelationSet) Relation(i int32) *Relation {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(set).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(i)

	_info := girepository.MustFind("Atk", "RelationSet")
	_gret := _info.InvokeClassMethod("get_relation", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(set)
	runtime.KeepAlive(i)

	var _relation *Relation // out

	_relation = wrapRelation(coreglib.Take(unsafe.Pointer(_cret)))

	return _relation
}

// Remove removes a relation from the relation set. This function unref's the
// Relation so it will be deleted unless there is another reference to it.
//
// The function takes the following parameters:
//
//    - relation: Relation.
//
func (set *RelationSet) Remove(relation *Relation) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(set).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(relation).Native()))

	_info := girepository.MustFind("Atk", "RelationSet")
	_info.InvokeClassMethod("remove", _args[:], nil)

	runtime.KeepAlive(set)
	runtime.KeepAlive(relation)
}
