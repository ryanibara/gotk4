// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_relation_set_get_type()), F: marshalRelationSeter},
	})
}

// RelationSeter describes RelationSet's methods.
type RelationSeter interface {
	// Add a new relation to the current relation set if it is not already
	// present.
	Add(relation *Relation)
	// AddRelationByType: add a new relation of the specified type with the
	// specified target to the current relation set if the relation set does not
	// contain a relation of that type.
	AddRelationByType(relationship RelationType, target *ObjectClass)
	// Contains determines whether the relation set contains a relation that
	// matches the specified type.
	Contains(relationship RelationType) bool
	// ContainsTarget determines whether the relation set contains a relation
	// that matches the specified pair formed by type relationship and object
	// target.
	ContainsTarget(relationship RelationType, target *ObjectClass) bool
	// NRelations determines the number of relations in a relation set.
	NRelations() int
	// Relation determines the relation at the specified position in the
	// relation set.
	Relation(i int) *Relation
	// RelationByType finds a relation that matches the specified type.
	RelationByType(relationship RelationType) *Relation
	// Remove removes a relation from the relation set.
	Remove(relation *Relation)
}

// RelationSet held by an object establishes its relationships with objects
// beyond the normal "parent/child" hierarchical relationships that all user
// interface objects have. AtkRelationSets establish whether objects are
// labelled or controlled by other components, share group membership with other
// components (for instance within a radio-button group), or share content which
// "flows" between them, among other types of possible relationships.
type RelationSet struct {
	*externglib.Object
}

var (
	_ RelationSeter   = (*RelationSet)(nil)
	_ gextras.Nativer = (*RelationSet)(nil)
)

func wrapRelationSet(obj *externglib.Object) *RelationSet {
	return &RelationSet{
		Object: obj,
	}
}

func marshalRelationSeter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRelationSet(obj), nil
}

// NewRelationSet creates a new empty relation set.
func NewRelationSet() *RelationSet {
	var _cret *C.AtkRelationSet // in

	_cret = C.atk_relation_set_new()

	var _relationSet *RelationSet // out

	_relationSet = wrapRelationSet(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _relationSet
}

// Add a new relation to the current relation set if it is not already present.
// This function ref's the AtkRelation so the caller of this function should
// unref it to ensure that it will be destroyed when the AtkRelationSet is
// destroyed.
func (set *RelationSet) Add(relation *Relation) {
	var _arg0 *C.AtkRelationSet // out
	var _arg1 *C.AtkRelation    // out

	_arg0 = (*C.AtkRelationSet)(unsafe.Pointer(set.Native()))
	_arg1 = (*C.AtkRelation)(unsafe.Pointer(relation.Native()))

	C.atk_relation_set_add(_arg0, _arg1)
}

// AddRelationByType: add a new relation of the specified type with the
// specified target to the current relation set if the relation set does not
// contain a relation of that type. If it is does contain a relation of that
// typea the target is added to the relation.
func (set *RelationSet) AddRelationByType(relationship RelationType, target *ObjectClass) {
	var _arg0 *C.AtkRelationSet // out
	var _arg1 C.AtkRelationType // out
	var _arg2 *C.AtkObject      // out

	_arg0 = (*C.AtkRelationSet)(unsafe.Pointer(set.Native()))
	_arg1 = C.AtkRelationType(relationship)
	_arg2 = (*C.AtkObject)(unsafe.Pointer(target.Native()))

	C.atk_relation_set_add_relation_by_type(_arg0, _arg1, _arg2)
}

// Contains determines whether the relation set contains a relation that matches
// the specified type.
func (set *RelationSet) Contains(relationship RelationType) bool {
	var _arg0 *C.AtkRelationSet // out
	var _arg1 C.AtkRelationType // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AtkRelationSet)(unsafe.Pointer(set.Native()))
	_arg1 = C.AtkRelationType(relationship)

	_cret = C.atk_relation_set_contains(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContainsTarget determines whether the relation set contains a relation that
// matches the specified pair formed by type relationship and object target.
func (set *RelationSet) ContainsTarget(relationship RelationType, target *ObjectClass) bool {
	var _arg0 *C.AtkRelationSet // out
	var _arg1 C.AtkRelationType // out
	var _arg2 *C.AtkObject      // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AtkRelationSet)(unsafe.Pointer(set.Native()))
	_arg1 = C.AtkRelationType(relationship)
	_arg2 = (*C.AtkObject)(unsafe.Pointer(target.Native()))

	_cret = C.atk_relation_set_contains_target(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NRelations determines the number of relations in a relation set.
func (set *RelationSet) NRelations() int {
	var _arg0 *C.AtkRelationSet // out
	var _cret C.gint            // in

	_arg0 = (*C.AtkRelationSet)(unsafe.Pointer(set.Native()))

	_cret = C.atk_relation_set_get_n_relations(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Relation determines the relation at the specified position in the relation
// set.
func (set *RelationSet) Relation(i int) *Relation {
	var _arg0 *C.AtkRelationSet // out
	var _arg1 C.gint            // out
	var _cret *C.AtkRelation    // in

	_arg0 = (*C.AtkRelationSet)(unsafe.Pointer(set.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_relation_set_get_relation(_arg0, _arg1)

	var _relation *Relation // out

	_relation = wrapRelation(externglib.Take(unsafe.Pointer(_cret)))

	return _relation
}

// RelationByType finds a relation that matches the specified type.
func (set *RelationSet) RelationByType(relationship RelationType) *Relation {
	var _arg0 *C.AtkRelationSet // out
	var _arg1 C.AtkRelationType // out
	var _cret *C.AtkRelation    // in

	_arg0 = (*C.AtkRelationSet)(unsafe.Pointer(set.Native()))
	_arg1 = C.AtkRelationType(relationship)

	_cret = C.atk_relation_set_get_relation_by_type(_arg0, _arg1)

	var _relation *Relation // out

	_relation = wrapRelation(externglib.Take(unsafe.Pointer(_cret)))

	return _relation
}

// Remove removes a relation from the relation set. This function unref's the
// Relation so it will be deleted unless there is another reference to it.
func (set *RelationSet) Remove(relation *Relation) {
	var _arg0 *C.AtkRelationSet // out
	var _arg1 *C.AtkRelation    // out

	_arg0 = (*C.AtkRelationSet)(unsafe.Pointer(set.Native()))
	_arg1 = (*C.AtkRelation)(unsafe.Pointer(relation.Native()))

	C.atk_relation_set_remove(_arg0, _arg1)
}
