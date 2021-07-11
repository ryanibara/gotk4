// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_relation_set_get_type()), F: marshalRelationSetter},
	})
}

// RelationSetter describes RelationSet's methods.
type RelationSetter interface {
	// Add a new relation to the current relation set if it is not already
	// present.
	Add(relation Relationer)
	// NRelations determines the number of relations in a relation set.
	NRelations() int
	// Relation determines the relation at the specified position in the
	// relation set.
	Relation(i int) *Relation
	// Remove removes a relation from the relation set.
	Remove(relation Relationer)
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
	_ RelationSetter  = (*RelationSet)(nil)
	_ gextras.Nativer = (*RelationSet)(nil)
)

func wrapRelationSet(obj *externglib.Object) RelationSetter {
	return &RelationSet{
		Object: obj,
	}
}

func marshalRelationSetter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRelationSet(obj), nil
}

// NewRelationSet creates a new empty relation set.
func NewRelationSet() *RelationSet {
	var _cret *C.AtkRelationSet // in

	_cret = C.atk_relation_set_new()

	var _relationSet *RelationSet // out

	_relationSet = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*RelationSet)

	return _relationSet
}

// Add a new relation to the current relation set if it is not already present.
// This function ref's the AtkRelation so the caller of this function should
// unref it to ensure that it will be destroyed when the AtkRelationSet is
// destroyed.
func (set *RelationSet) Add(relation Relationer) {
	var _arg0 *C.AtkRelationSet // out
	var _arg1 *C.AtkRelation    // out

	_arg0 = (*C.AtkRelationSet)(unsafe.Pointer(set.Native()))
	_arg1 = (*C.AtkRelation)(unsafe.Pointer((relation).(gextras.Nativer).Native()))

	C.atk_relation_set_add(_arg0, _arg1)
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

	_relation = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Relation)

	return _relation
}

// Remove removes a relation from the relation set. This function unref's the
// Relation so it will be deleted unless there is another reference to it.
func (set *RelationSet) Remove(relation Relationer) {
	var _arg0 *C.AtkRelationSet // out
	var _arg1 *C.AtkRelation    // out

	_arg0 = (*C.AtkRelationSet)(unsafe.Pointer(set.Native()))
	_arg1 = (*C.AtkRelation)(unsafe.Pointer((relation).(gextras.Nativer).Native()))

	C.atk_relation_set_remove(_arg0, _arg1)
}
