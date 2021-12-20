// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_relation_get_type()), F: marshalRelationer},
	})
}

// RelationTypeForName: get the RelationType type corresponding to a relation
// name.
//
// The function takes the following parameters:
//
//    - name: string which is the (non-localized) name of an ATK relation type.
//
// The function returns the following values:
//
//    - relationType enumerated type corresponding to the specified name, or
//      K_RELATION_NULL if no matching relation type is found.
//
func RelationTypeForName(name string) RelationType {
	var _arg1 *C.gchar          // out
	var _cret C.AtkRelationType // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.atk_relation_type_for_name(_arg1)
	runtime.KeepAlive(name)

	var _relationType RelationType // out

	_relationType = RelationType(_cret)

	return _relationType
}

// RelationTypeGetName gets the description string describing the RelationType
// type.
//
// The function takes the following parameters:
//
//    - typ whose name is required.
//
// The function returns the following values:
//
//    - utf8: string describing the AtkRelationType.
//
func RelationTypeGetName(typ RelationType) string {
	var _arg1 C.AtkRelationType // out
	var _cret *C.gchar          // in

	_arg1 = C.AtkRelationType(typ)

	_cret = C.atk_relation_type_get_name(_arg1)
	runtime.KeepAlive(typ)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// RelationTypeRegister: associate name with a new RelationType.
//
// The function takes the following parameters:
//
//    - name string.
//
// The function returns the following values:
//
//    - relationType associated with name.
//
func RelationTypeRegister(name string) RelationType {
	var _arg1 *C.gchar          // out
	var _cret C.AtkRelationType // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.atk_relation_type_register(_arg1)
	runtime.KeepAlive(name)

	var _relationType RelationType // out

	_relationType = RelationType(_cret)

	return _relationType
}

// Relation describes a relation between an object and one or more other
// objects. The actual relations that an object has with other objects are
// defined as an AtkRelationSet, which is a set of AtkRelations.
type Relation struct {
	*externglib.Object

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ externglib.Objector = (*Relation)(nil)
)

func wrapRelation(obj *externglib.Object) *Relation {
	return &Relation{
		Object: obj,
	}
}

func marshalRelationer(p uintptr) (interface{}, error) {
	return wrapRelation(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewRelation: create a new relation for the specified key and the specified
// list of targets. See also atk_object_add_relationship().
//
// The function takes the following parameters:
//
//    - targets: array of pointers to Objects.
//    - relationship with which to create the new Relation.
//
// The function returns the following values:
//
//    - relation: pointer to a new Relation.
//
func NewRelation(targets []*ObjectClass, relationship RelationType) *Relation {
	var _arg1 **C.AtkObject // out
	var _arg2 C.gint
	var _arg3 C.AtkRelationType // out
	var _cret *C.AtkRelation    // in

	_arg2 = (C.gint)(len(targets))
	_arg1 = (**C.AtkObject)(C.calloc(C.size_t(len(targets)), C.size_t(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((**C.AtkObject)(_arg1), len(targets))
		for i := range targets {
			out[i] = (*C.AtkObject)(unsafe.Pointer(targets[i].Native()))
		}
	}
	_arg3 = C.AtkRelationType(relationship)

	_cret = C.atk_relation_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(targets)
	runtime.KeepAlive(relationship)

	var _relation *Relation // out

	_relation = wrapRelation(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _relation
}

// AddTarget adds the specified AtkObject to the target for the relation, if it
// is not already present. See also atk_object_add_relationship().
//
// The function takes the following parameters:
//
//    - target: Object.
//
func (relation *Relation) AddTarget(target *ObjectClass) {
	var _arg0 *C.AtkRelation // out
	var _arg1 *C.AtkObject   // out

	_arg0 = (*C.AtkRelation)(unsafe.Pointer(relation.Native()))
	_arg1 = (*C.AtkObject)(unsafe.Pointer(target.Native()))

	C.atk_relation_add_target(_arg0, _arg1)
	runtime.KeepAlive(relation)
	runtime.KeepAlive(target)
}

// RelationType gets the type of relation.
//
// The function returns the following values:
//
//    - relationType: type of relation.
//
func (relation *Relation) RelationType() RelationType {
	var _arg0 *C.AtkRelation    // out
	var _cret C.AtkRelationType // in

	_arg0 = (*C.AtkRelation)(unsafe.Pointer(relation.Native()))

	_cret = C.atk_relation_get_relation_type(_arg0)
	runtime.KeepAlive(relation)

	var _relationType RelationType // out

	_relationType = RelationType(_cret)

	return _relationType
}

// RemoveTarget: remove the specified AtkObject from the target for the
// relation.
//
// The function takes the following parameters:
//
//    - target: Object.
//
// The function returns the following values:
//
//    - ok: TRUE if the removal is successful.
//
func (relation *Relation) RemoveTarget(target *ObjectClass) bool {
	var _arg0 *C.AtkRelation // out
	var _arg1 *C.AtkObject   // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkRelation)(unsafe.Pointer(relation.Native()))
	_arg1 = (*C.AtkObject)(unsafe.Pointer(target.Native()))

	_cret = C.atk_relation_remove_target(_arg0, _arg1)
	runtime.KeepAlive(relation)
	runtime.KeepAlive(target)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
