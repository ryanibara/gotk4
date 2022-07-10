// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// BitsetIter: opaque, stack-allocated struct for iterating over the elements of
// a GtkBitset.
//
// Before a GtkBitsetIter can be used, it needs to be initialized with
// gtk.BitsetIter().InitFirst, gtk.BitsetIter().InitLast or
// gtk.BitsetIter().InitAt.
//
// An instance of this type is always passed by reference.
type BitsetIter struct {
	*bitsetIter
}

// bitsetIter is the struct that's finalized.
type bitsetIter struct {
	native unsafe.Pointer
}

// Value gets the current value that iter points to.
//
// If iter is not valid and gtk.BitsetIter.IsValid() returns FALSE, this
// function returns 0.
//
// The function returns the following values:
//
//    - guint: current value pointer to by iter.
//
func (iter *BitsetIter) Value() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))

	_info := girepository.MustFind("Gtk", "BitsetIter")
	_gret := _info.InvokeRecordMethod("get_value", _args[:], nil)
	_cret := *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(iter)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// IsValid checks if iter points to a valid value.
//
// The function returns the following values:
//
//    - ok: TRUE if iter points to a valid value.
//
func (iter *BitsetIter) IsValid() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))

	_info := girepository.MustFind("Gtk", "BitsetIter")
	_gret := _info.InvokeRecordMethod("is_valid", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(iter)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Next moves iter to the next value in the set.
//
// If it was already pointing to the last value in the set, FALSE is returned
// and iter is invalidated.
//
// The function returns the following values:
//
//    - value (optional): set to the next value.
//    - ok: TRUE if a next value existed.
//
func (iter *BitsetIter) Next() (uint32, bool) {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))

	_info := girepository.MustFind("Gtk", "BitsetIter")
	_gret := _info.InvokeRecordMethod("next", _args[:], _outs[:])
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(iter)

	var _value uint32 // out
	var _ok bool      // out

	_value = uint32(*(*C.guint)(unsafe.Pointer(&_outs[0])))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _value, _ok
}

// Previous moves iter to the previous value in the set.
//
// If it was already pointing to the first value in the set, FALSE is returned
// and iter is invalidated.
//
// The function returns the following values:
//
//    - value (optional): set to the previous value.
//    - ok: TRUE if a previous value existed.
//
func (iter *BitsetIter) Previous() (uint32, bool) {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))

	_info := girepository.MustFind("Gtk", "BitsetIter")
	_gret := _info.InvokeRecordMethod("previous", _args[:], _outs[:])
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(iter)

	var _value uint32 // out
	var _ok bool      // out

	_value = uint32(*(*C.guint)(unsafe.Pointer(&_outs[0])))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _value, _ok
}

// BitsetIterInitAt initializes iter to point to target.
//
// If target is not found, finds the next value after it. If no value >= target
// exists in set, this function returns FALSE.
//
// The function takes the following parameters:
//
//    - set: GtkBitset.
//    - target value to start iterating at.
//
// The function returns the following values:
//
//    - iter: pointer to an uninitialized GtkBitsetIter.
//    - value (optional): set to the found value in set.
//    - ok: TRUE if a value was found.
//
func BitsetIterInitAt(set *Bitset, target uint32) (*BitsetIter, uint32, bool) {
	var _args [2]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(set)))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(target)

	_info := girepository.MustFind("Gtk", "init_at")
	_gret := _info.InvokeFunction(_args[:], _outs[:])
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(set)
	runtime.KeepAlive(target)

	var _iter *BitsetIter // out
	var _value uint32     // out
	var _ok bool          // out

	_iter = (*BitsetIter)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[0])))))
	_value = uint32(*(*C.guint)(unsafe.Pointer(&_outs[1])))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _iter, _value, _ok
}

// BitsetIterInitFirst initializes an iterator for set and points it to the
// first value in set.
//
// If set is empty, FALSE is returned and value is set to G_MAXUINT.
//
// The function takes the following parameters:
//
//    - set: GtkBitset.
//
// The function returns the following values:
//
//    - iter: pointer to an uninitialized GtkBitsetIter.
//    - value (optional): set to the first value in set.
//    - ok: TRUE if set isn't empty.
//
func BitsetIterInitFirst(set *Bitset) (*BitsetIter, uint32, bool) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(set)))

	_info := girepository.MustFind("Gtk", "init_first")
	_gret := _info.InvokeFunction(_args[:], _outs[:])
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(set)

	var _iter *BitsetIter // out
	var _value uint32     // out
	var _ok bool          // out

	_iter = (*BitsetIter)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[0])))))
	_value = uint32(*(*C.guint)(unsafe.Pointer(&_outs[1])))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _iter, _value, _ok
}

// BitsetIterInitLast initializes an iterator for set and points it to the last
// value in set.
//
// If set is empty, FALSE is returned.
//
// The function takes the following parameters:
//
//    - set: GtkBitset.
//
// The function returns the following values:
//
//    - iter: pointer to an uninitialized GtkBitsetIter.
//    - value (optional): set to the last value in set.
//    - ok: TRUE if set isn't empty.
//
func BitsetIterInitLast(set *Bitset) (*BitsetIter, uint32, bool) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(set)))

	_info := girepository.MustFind("Gtk", "init_last")
	_gret := _info.InvokeFunction(_args[:], _outs[:])
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(set)

	var _iter *BitsetIter // out
	var _value uint32     // out
	var _ok bool          // out

	_iter = (*BitsetIter)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[0])))))
	_value = uint32(*(*C.guint)(unsafe.Pointer(&_outs[1])))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _iter, _value, _ok
}
