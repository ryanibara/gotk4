// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_state_set_get_type()), F: marshalStateSetter},
	})
}

// StateSet is a read-only representation of the full set of States that apply
// to an object at a given time. This set is not meant to be modified, but
// rather created when #atk_object_ref_state_set() is called.
type StateSet struct {
	*externglib.Object
}

func wrapStateSet(obj *externglib.Object) *StateSet {
	return &StateSet{
		Object: obj,
	}
}

func marshalStateSetter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStateSet(obj), nil
}

// NewStateSet creates a new empty state set.
func NewStateSet() *StateSet {
	var _cret *C.AtkStateSet // in

	_cret = C.atk_state_set_new()

	var _stateSet *StateSet // out

	_stateSet = wrapStateSet(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _stateSet
}

// AddState adds the state of the specified type to the state set if it is not
// already present.
//
// Note that because an StateSet is a read-only object, this method should be
// used to add a state to a newly-created set which will then be returned by
// #atk_object_ref_state_set. It should not be used to modify the existing state
// of an object. See also #atk_object_notify_state_change.
func (set *StateSet) AddState(typ StateType) bool {
	var _arg0 *C.AtkStateSet // out
	var _arg1 C.AtkStateType // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(set.Native()))
	_arg1 = C.AtkStateType(typ)

	_cret = C.atk_state_set_add_state(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AddStates adds the states of the specified types to the state set.
//
// Note that because an StateSet is a read-only object, this method should be
// used to add states to a newly-created set which will then be returned by
// #atk_object_ref_state_set. It should not be used to modify the existing state
// of an object. See also #atk_object_notify_state_change.
func (set *StateSet) AddStates(types []StateType) {
	var _arg0 *C.AtkStateSet  // out
	var _arg1 *C.AtkStateType // out
	var _arg2 C.gint

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(set.Native()))
	_arg2 = (C.gint)(len(types))
	_arg1 = (*C.AtkStateType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_AtkStateType)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.AtkStateType)(_arg1), len(types))
		for i := range types {
			out[i] = C.AtkStateType(types[i])
		}
	}

	C.atk_state_set_add_states(_arg0, _arg1, _arg2)
}

// AndSets constructs the intersection of the two sets, returning NULL if the
// intersection is empty.
func (set *StateSet) AndSets(compareSet *StateSet) *StateSet {
	var _arg0 *C.AtkStateSet // out
	var _arg1 *C.AtkStateSet // out
	var _cret *C.AtkStateSet // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(set.Native()))
	_arg1 = (*C.AtkStateSet)(unsafe.Pointer(compareSet.Native()))

	_cret = C.atk_state_set_and_sets(_arg0, _arg1)

	var _stateSet *StateSet // out

	_stateSet = wrapStateSet(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _stateSet
}

// ClearStates removes all states from the state set.
func (set *StateSet) ClearStates() {
	var _arg0 *C.AtkStateSet // out

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(set.Native()))

	C.atk_state_set_clear_states(_arg0)
}

// ContainsState checks whether the state for the specified type is in the
// specified set.
func (set *StateSet) ContainsState(typ StateType) bool {
	var _arg0 *C.AtkStateSet // out
	var _arg1 C.AtkStateType // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(set.Native()))
	_arg1 = C.AtkStateType(typ)

	_cret = C.atk_state_set_contains_state(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContainsStates checks whether the states for all the specified types are in
// the specified set.
func (set *StateSet) ContainsStates(types []StateType) bool {
	var _arg0 *C.AtkStateSet  // out
	var _arg1 *C.AtkStateType // out
	var _arg2 C.gint
	var _cret C.gboolean // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(set.Native()))
	_arg2 = (C.gint)(len(types))
	_arg1 = (*C.AtkStateType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_AtkStateType)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.AtkStateType)(_arg1), len(types))
		for i := range types {
			out[i] = C.AtkStateType(types[i])
		}
	}

	_cret = C.atk_state_set_contains_states(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsEmpty checks whether the state set is empty, i.e. has no states set.
func (set *StateSet) IsEmpty() bool {
	var _arg0 *C.AtkStateSet // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(set.Native()))

	_cret = C.atk_state_set_is_empty(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// OrSets constructs the union of the two sets.
func (set *StateSet) OrSets(compareSet *StateSet) *StateSet {
	var _arg0 *C.AtkStateSet // out
	var _arg1 *C.AtkStateSet // out
	var _cret *C.AtkStateSet // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(set.Native()))
	_arg1 = (*C.AtkStateSet)(unsafe.Pointer(compareSet.Native()))

	_cret = C.atk_state_set_or_sets(_arg0, _arg1)

	var _stateSet *StateSet // out

	if _cret != nil {
		_stateSet = wrapStateSet(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _stateSet
}

// RemoveState removes the state for the specified type from the state set.
//
// Note that because an StateSet is a read-only object, this method should be
// used to remove a state to a newly-created set which will then be returned by
// #atk_object_ref_state_set. It should not be used to modify the existing state
// of an object. See also #atk_object_notify_state_change.
func (set *StateSet) RemoveState(typ StateType) bool {
	var _arg0 *C.AtkStateSet // out
	var _arg1 C.AtkStateType // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(set.Native()))
	_arg1 = C.AtkStateType(typ)

	_cret = C.atk_state_set_remove_state(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// XorSets constructs the exclusive-or of the two sets, returning NULL is empty.
// The set returned by this operation contains the states in exactly one of the
// two sets.
func (set *StateSet) XorSets(compareSet *StateSet) *StateSet {
	var _arg0 *C.AtkStateSet // out
	var _arg1 *C.AtkStateSet // out
	var _cret *C.AtkStateSet // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(set.Native()))
	_arg1 = (*C.AtkStateSet)(unsafe.Pointer(compareSet.Native()))

	_cret = C.atk_state_set_xor_sets(_arg0, _arg1)

	var _stateSet *StateSet // out

	_stateSet = wrapStateSet(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _stateSet
}
