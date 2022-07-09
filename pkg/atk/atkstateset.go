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
import "C"

// GTypeStateSet returns the GType for the type StateSet.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeStateSet() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Atk", "StateSet").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalStateSet)
	return gtype
}

// StateSetOverrider contains methods that are overridable.
type StateSetOverrider interface {
}

// StateSet is a read-only representation of the full set of States that apply
// to an object at a given time. This set is not meant to be modified, but
// rather created when #atk_object_ref_state_set() is called.
type StateSet struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*StateSet)(nil)
)

func classInitStateSetter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapStateSet(obj *coreglib.Object) *StateSet {
	return &StateSet{
		Object: obj,
	}
}

func marshalStateSet(p uintptr) (interface{}, error) {
	return wrapStateSet(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStateSet creates a new empty state set.
//
// The function returns the following values:
//
//    - stateSet: new StateSet.
//
func NewStateSet() *StateSet {
	_gret := girepository.MustFind("Atk", "StateSet").InvokeMethod("new_StateSet", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _stateSet *StateSet // out

	_stateSet = wrapStateSet(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _stateSet
}

// AndSets constructs the intersection of the two sets, returning NULL if the
// intersection is empty.
//
// The function takes the following parameters:
//
//    - compareSet: another StateSet.
//
// The function returns the following values:
//
//    - stateSet: new StateSet which is the intersection of the two sets.
//
func (set *StateSet) AndSets(compareSet *StateSet) *StateSet {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(set).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(compareSet).Native()))

	_gret := girepository.MustFind("Atk", "StateSet").InvokeMethod("and_sets", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(set)
	runtime.KeepAlive(compareSet)

	var _stateSet *StateSet // out

	_stateSet = wrapStateSet(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _stateSet
}

// ClearStates removes all states from the state set.
func (set *StateSet) ClearStates() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(set).Native()))

	girepository.MustFind("Atk", "StateSet").InvokeMethod("clear_states", _args[:], nil)

	runtime.KeepAlive(set)
}

// IsEmpty checks whether the state set is empty, i.e. has no states set.
//
// The function returns the following values:
//
//    - ok: TRUE if set has no states set, otherwise FALSE.
//
func (set *StateSet) IsEmpty() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(set).Native()))

	_gret := girepository.MustFind("Atk", "StateSet").InvokeMethod("is_empty", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(set)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// OrSets constructs the union of the two sets.
//
// The function takes the following parameters:
//
//    - compareSet: another StateSet.
//
// The function returns the following values:
//
//    - stateSet (optional): new StateSet which is the union of the two sets,
//      returning NULL is empty.
//
func (set *StateSet) OrSets(compareSet *StateSet) *StateSet {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(set).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(compareSet).Native()))

	_gret := girepository.MustFind("Atk", "StateSet").InvokeMethod("or_sets", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(set)
	runtime.KeepAlive(compareSet)

	var _stateSet *StateSet // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_stateSet = wrapStateSet(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _stateSet
}

// XorSets constructs the exclusive-or of the two sets, returning NULL is empty.
// The set returned by this operation contains the states in exactly one of the
// two sets.
//
// The function takes the following parameters:
//
//    - compareSet: another StateSet.
//
// The function returns the following values:
//
//    - stateSet: new StateSet which contains the states which are in exactly one
//      of the two sets.
//
func (set *StateSet) XorSets(compareSet *StateSet) *StateSet {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(set).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(compareSet).Native()))

	_gret := girepository.MustFind("Atk", "StateSet").InvokeMethod("xor_sets", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(set)
	runtime.KeepAlive(compareSet)

	var _stateSet *StateSet // out

	_stateSet = wrapStateSet(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _stateSet
}
