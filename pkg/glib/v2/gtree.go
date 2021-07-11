// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tree_get_type()), F: marshalTree},
	})
}

// Tree struct is an opaque data structure representing a [balanced binary
// tree][glib-Balanced-Binary-Trees]. It should be accessed only by using the
// following functions.
type Tree struct {
	native C.GTree
}

func marshalTree(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Tree)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *Tree) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Destroy removes all keys and values from the #GTree and decreases its
// reference count by one. If keys and/or values are dynamically allocated, you
// should either free them first or create the #GTree using g_tree_new_full().
// In the latter case the destroy functions you supplied will be called on all
// keys and values before destroying the #GTree.
func (tree *Tree) Destroy() {
	var _arg0 *C.GTree // out

	_arg0 = (*C.GTree)(unsafe.Pointer(tree))

	C.g_tree_destroy(_arg0)
}

// Height gets the height of a #GTree.
//
// If the #GTree contains no nodes, the height is 0. If the #GTree contains only
// one root node the height is 1. If the root node has children the height is 2,
// etc.
func (tree *Tree) Height() int {
	var _arg0 *C.GTree // out
	var _cret C.gint   // in

	_arg0 = (*C.GTree)(unsafe.Pointer(tree))

	_cret = C.g_tree_height(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Insert inserts a key/value pair into a #GTree.
//
// Inserts a new key and value into a #GTree as g_tree_insert_node() does, only
// this function does not return the inserted or set node.
func (tree *Tree) Insert(key cgo.Handle, value cgo.Handle) {
	var _arg0 *C.GTree   // out
	var _arg1 C.gpointer // out
	var _arg2 C.gpointer // out

	_arg0 = (*C.GTree)(unsafe.Pointer(tree))
	_arg1 = (C.gpointer)(unsafe.Pointer(key))
	_arg2 = (C.gpointer)(unsafe.Pointer(value))

	C.g_tree_insert(_arg0, _arg1, _arg2)
}

// Lookup gets the value corresponding to the given key. Since a #GTree is
// automatically balanced as key/value pairs are added, key lookup is O(log n)
// (where n is the number of key/value pairs in the tree).
func (tree *Tree) Lookup(key cgo.Handle) cgo.Handle {
	var _arg0 *C.GTree        // out
	var _arg1 C.gconstpointer // out
	var _cret C.gpointer      // in

	_arg0 = (*C.GTree)(unsafe.Pointer(tree))
	_arg1 = (C.gconstpointer)(unsafe.Pointer(key))

	_cret = C.g_tree_lookup(_arg0, _arg1)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}

// LookupExtended looks up a key in the #GTree, returning the original key and
// the associated value. This is useful if you need to free the memory allocated
// for the original key, for example before calling g_tree_remove().
func (tree *Tree) LookupExtended(lookupKey cgo.Handle) (origKey cgo.Handle, value cgo.Handle, ok bool) {
	var _arg0 *C.GTree        // out
	var _arg1 C.gconstpointer // out
	var _arg2 C.gpointer      // in
	var _arg3 C.gpointer      // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GTree)(unsafe.Pointer(tree))
	_arg1 = (C.gconstpointer)(unsafe.Pointer(lookupKey))

	_cret = C.g_tree_lookup_extended(_arg0, _arg1, &_arg2, &_arg3)

	var _origKey cgo.Handle // out
	var _value cgo.Handle   // out
	var _ok bool            // out

	_origKey = (cgo.Handle)(unsafe.Pointer(_arg2))
	_value = (cgo.Handle)(unsafe.Pointer(_arg3))
	if _cret != 0 {
		_ok = true
	}

	return _origKey, _value, _ok
}

// Nnodes gets the number of nodes in a #GTree.
func (tree *Tree) Nnodes() int {
	var _arg0 *C.GTree // out
	var _cret C.gint   // in

	_arg0 = (*C.GTree)(unsafe.Pointer(tree))

	_cret = C.g_tree_nnodes(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Ref increments the reference count of @tree by one.
//
// It is safe to call this function from any thread.
func (tree *Tree) ref() *Tree {
	var _arg0 *C.GTree // out
	var _cret *C.GTree // in

	_arg0 = (*C.GTree)(unsafe.Pointer(tree))

	_cret = C.g_tree_ref(_arg0)

	var _ret *Tree // out

	_ret = (*Tree)(unsafe.Pointer(_cret))
	C.g_tree_ref(_cret)
	runtime.SetFinalizer(_ret, func(v *Tree) {
		C.g_tree_unref((*C.GTree)(unsafe.Pointer(v)))
	})

	return _ret
}

// Remove removes a key/value pair from a #GTree.
//
// If the #GTree was created using g_tree_new_full(), the key and value are
// freed using the supplied destroy functions, otherwise you have to make sure
// that any dynamically allocated values are freed yourself. If the key does not
// exist in the #GTree, the function does nothing.
//
// The cost of maintaining a balanced tree while removing a key/value result in
// a O(n log(n)) operation where most of the other operations are O(log(n)).
func (tree *Tree) Remove(key cgo.Handle) bool {
	var _arg0 *C.GTree        // out
	var _arg1 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GTree)(unsafe.Pointer(tree))
	_arg1 = (C.gconstpointer)(unsafe.Pointer(key))

	_cret = C.g_tree_remove(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Replace inserts a new key and value into a #GTree as g_tree_replace_node()
// does, only this function does not return the inserted or set node.
func (tree *Tree) Replace(key cgo.Handle, value cgo.Handle) {
	var _arg0 *C.GTree   // out
	var _arg1 C.gpointer // out
	var _arg2 C.gpointer // out

	_arg0 = (*C.GTree)(unsafe.Pointer(tree))
	_arg1 = (C.gpointer)(unsafe.Pointer(key))
	_arg2 = (C.gpointer)(unsafe.Pointer(value))

	C.g_tree_replace(_arg0, _arg1, _arg2)
}

// Steal removes a key and its associated value from a #GTree without calling
// the key and value destroy functions.
//
// If the key does not exist in the #GTree, the function does nothing.
func (tree *Tree) Steal(key cgo.Handle) bool {
	var _arg0 *C.GTree        // out
	var _arg1 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GTree)(unsafe.Pointer(tree))
	_arg1 = (C.gconstpointer)(unsafe.Pointer(key))

	_cret = C.g_tree_steal(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Unref decrements the reference count of @tree by one. If the reference count
// drops to 0, all keys and values will be destroyed (if destroy functions were
// specified) and all memory allocated by @tree will be released.
//
// It is safe to call this function from any thread.
func (tree *Tree) unref() {
	var _arg0 *C.GTree // out

	_arg0 = (*C.GTree)(unsafe.Pointer(tree))

	C.g_tree_unref(_arg0)
}
