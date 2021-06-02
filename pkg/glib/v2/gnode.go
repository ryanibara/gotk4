// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
)

// #cgo pkg-config: glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdbool.h>
// #include <glib.h>
import "C"

// TraverseType specifies the type of traversal performed by g_tree_traverse(),
// g_node_traverse() and g_node_find(). The different orders are illustrated
// here: - In order: A, B, C, D, E, F, G, H, I !
// (Sorted_binary_tree_inorder.svg) - Pre order: F, B, A, D, C, E, G, I, H !
// (Sorted_binary_tree_preorder.svg) - Post order: A, C, E, D, B, H, I, G, F !
// (Sorted_binary_tree_postorder.svg) - Level order: F, B, G, A, D, I, C, E, H !
// (Sorted_binary_tree_breadth-first_traversal.svg)
type TraverseType int

const (
	// TraverseTypeInOrder vists a node's left child first, then the node
	// itself, then its right child. This is the one to use if you want the
	// output sorted according to the compare function.
	TraverseTypeInOrder TraverseType = 0
	// TraverseTypePreOrder visits a node, then its children.
	TraverseTypePreOrder TraverseType = 1
	// TraverseTypePostOrder visits the node's children, then the node itself.
	TraverseTypePostOrder TraverseType = 2
	// TraverseTypeLevelOrder is not implemented for [balanced binary
	// trees][glib-Balanced-Binary-Trees]. For [n-ary trees][glib-N-ary-Trees],
	// it vists the root node first, then its children, then its grandchildren,
	// and so on. Note that this is less efficient than the other orders.
	TraverseTypeLevelOrder TraverseType = 3
)

// TraverseFlags specifies which nodes are visited during several of the tree
// functions, including g_node_traverse() and g_node_find().
type TraverseFlags int

const (
	// TraverseFlagsLeaves: only leaf nodes should be visited. This name has
	// been introduced in 2.6, for older version use G_TRAVERSE_LEAFS.
	TraverseFlagsLeaves TraverseFlags = 0b1
	// TraverseFlagsNonLeaves: only non-leaf nodes should be visited. This name
	// has been introduced in 2.6, for older version use G_TRAVERSE_NON_LEAFS.
	TraverseFlagsNonLeaves TraverseFlags = 0b10
	// TraverseFlagsAll: all nodes should be visited.
	TraverseFlagsAll TraverseFlags = 0b11
	// TraverseFlagsMask: a mask of all traverse flags.
	TraverseFlagsMask TraverseFlags = 0b11
	// TraverseFlagsLeafs: identical to G_TRAVERSE_LEAVES.
	TraverseFlagsLeafs TraverseFlags = 0b1
	// TraverseFlagsNonLeafs: identical to G_TRAVERSE_NON_LEAVES.
	TraverseFlagsNonLeafs TraverseFlags = 0b10
)

// Node: the #GNode struct represents one node in a [n-ary
// tree][glib-N-ary-Trees].
type Node struct {
	native C.GNode
}

// WrapNode wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapNode(ptr unsafe.Pointer) *Node {
	if ptr == nil {
		return nil
	}

	return (*Node)(ptr)
}

func marshalNode(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapNode(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (n *Node) Native() unsafe.Pointer {
	return unsafe.Pointer(&n.native)
}

// Data gets the field inside the struct.
func (n *Node) Data() interface{} {
	var ret interface{}
	ret = box.Get(uintptr(n.native.data))
	return ret
}

// Next gets the field inside the struct.
func (n *Node) Next() *Node {
	var ret *Node
	{
		ret = WrapNode(unsafe.Pointer(n.native.next))
	}
	return ret
}

// Prev gets the field inside the struct.
func (n *Node) Prev() *Node {
	var ret *Node
	{
		ret = WrapNode(unsafe.Pointer(n.native.prev))
	}
	return ret
}

// Parent gets the field inside the struct.
func (n *Node) Parent() *Node {
	var ret *Node
	{
		ret = WrapNode(unsafe.Pointer(n.native.parent))
	}
	return ret
}

// Children gets the field inside the struct.
func (n *Node) Children() *Node {
	var ret *Node
	{
		ret = WrapNode(unsafe.Pointer(n.native.children))
	}
	return ret
}

// ChildIndex gets the position of the first child of a #GNode which contains
// the given data.
func (n *Node) ChildIndex(data interface{}) int {
	var arg0 *C.GNode
	var arg1 C.gpointer

	arg0 = (*C.GNode)(n.Native())
	arg1 = C.gpointer(box.Assign(data))

	ret := C.g_node_child_index(arg0, arg1)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// ChildPosition gets the position of a #GNode with respect to its siblings.
// @child must be a child of @node. The first child is numbered 0, the second 1,
// and so on.
func (n *Node) ChildPosition(child *Node) int {
	var arg0 *C.GNode
	var arg1 *C.GNode

	arg0 = (*C.GNode)(n.Native())
	arg1 = (*C.GNode)(child.Native())

	ret := C.g_node_child_position(arg0, arg1)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// ChildrenForeach calls a function for each of the children of a #GNode. Note
// that it doesn't descend beneath the child nodes. @func must not do anything
// that would modify the structure of the tree.
func (n *Node) ChildrenForeach(flags TraverseFlags, fn NodeForeachFunc) {
	var arg0 *C.GNode
	var arg1 C.GTraverseFlags
	var arg2 C.GNodeForeachFunc
	var arg3 C.gpointer

	arg0 = (*C.GNode)(n.Native())
	arg1 = (C.GTraverseFlags)(flags)
	arg2 = (*[0]byte)(C.gotk4_NodeForeachFunc)
	arg3 = C.gpointer(box.Assign(fn))

	C.g_node_children_foreach(arg0, arg1, arg2, arg3)
}

// Copy: recursively copies a #GNode (but does not deep-copy the data inside the
// nodes, see g_node_copy_deep() if you need that).
func (n *Node) Copy() *Node {
	var arg0 *C.GNode

	arg0 = (*C.GNode)(n.Native())

	ret := C.g_node_copy(arg0)

	var ret0 *Node

	{
		ret0 = WrapNode(unsafe.Pointer(ret))
	}

	return ret0
}

// CopyDeep: recursively copies a #GNode and its data.
func (n *Node) CopyDeep(copyFunc CopyFunc) *Node {
	var arg0 *C.GNode
	var arg1 C.GCopyFunc
	var arg2 C.gpointer

	arg0 = (*C.GNode)(n.Native())
	arg1 = (*[0]byte)(C.gotk4_CopyFunc)
	arg2 = C.gpointer(box.Assign(copyFunc))

	ret := C.g_node_copy_deep(arg0, arg1, arg2)

	var ret0 *Node

	{
		ret0 = WrapNode(unsafe.Pointer(ret))
	}

	return ret0
}

// Depth gets the depth of a #GNode.
//
// If @node is nil the depth is 0. The root node has a depth of 1. For the
// children of the root node the depth is 2. And so on.
func (n *Node) Depth() uint {
	var arg0 *C.GNode

	arg0 = (*C.GNode)(n.Native())

	ret := C.g_node_depth(arg0)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// Destroy removes @root and its children from the tree, freeing any memory
// allocated.
func (r *Node) Destroy() {
	var arg0 *C.GNode

	arg0 = (*C.GNode)(r.Native())

	C.g_node_destroy(arg0)
}

// Find finds a #GNode in a tree.
func (r *Node) Find(order TraverseType, flags TraverseFlags, data interface{}) *Node {
	var arg0 *C.GNode
	var arg1 C.GTraverseType
	var arg2 C.GTraverseFlags
	var arg3 C.gpointer

	arg0 = (*C.GNode)(r.Native())
	arg1 = (C.GTraverseType)(order)
	arg2 = (C.GTraverseFlags)(flags)
	arg3 = C.gpointer(box.Assign(data))

	ret := C.g_node_find(arg0, arg1, arg2, arg3)

	var ret0 *Node

	{
		ret0 = WrapNode(unsafe.Pointer(ret))
	}

	return ret0
}

// FindChild finds the first child of a #GNode with the given data.
func (n *Node) FindChild(flags TraverseFlags, data interface{}) *Node {
	var arg0 *C.GNode
	var arg1 C.GTraverseFlags
	var arg2 C.gpointer

	arg0 = (*C.GNode)(n.Native())
	arg1 = (C.GTraverseFlags)(flags)
	arg2 = C.gpointer(box.Assign(data))

	ret := C.g_node_find_child(arg0, arg1, arg2)

	var ret0 *Node

	{
		ret0 = WrapNode(unsafe.Pointer(ret))
	}

	return ret0
}

// FirstSibling gets the first sibling of a #GNode. This could possibly be the
// node itself.
func (n *Node) FirstSibling() *Node {
	var arg0 *C.GNode

	arg0 = (*C.GNode)(n.Native())

	ret := C.g_node_first_sibling(arg0)

	var ret0 *Node

	{
		ret0 = WrapNode(unsafe.Pointer(ret))
	}

	return ret0
}

// Root gets the root of a tree.
func (n *Node) Root() *Node {
	var arg0 *C.GNode

	arg0 = (*C.GNode)(n.Native())

	ret := C.g_node_get_root(arg0)

	var ret0 *Node

	{
		ret0 = WrapNode(unsafe.Pointer(ret))
	}

	return ret0
}

// Insert inserts a #GNode beneath the parent at the given position.
func (p *Node) Insert(position int, node *Node) *Node {
	var arg0 *C.GNode
	var arg1 C.gint
	var arg2 *C.GNode

	arg0 = (*C.GNode)(p.Native())
	arg1 = C.gint(position)
	arg2 = (*C.GNode)(node.Native())

	ret := C.g_node_insert(arg0, arg1, arg2)

	var ret0 *Node

	{
		ret0 = WrapNode(unsafe.Pointer(ret))
	}

	return ret0
}

// InsertAfter inserts a #GNode beneath the parent after the given sibling.
func (p *Node) InsertAfter(sibling *Node, node *Node) *Node {
	var arg0 *C.GNode
	var arg1 *C.GNode
	var arg2 *C.GNode

	arg0 = (*C.GNode)(p.Native())
	arg1 = (*C.GNode)(sibling.Native())
	arg2 = (*C.GNode)(node.Native())

	ret := C.g_node_insert_after(arg0, arg1, arg2)

	var ret0 *Node

	{
		ret0 = WrapNode(unsafe.Pointer(ret))
	}

	return ret0
}

// InsertBefore inserts a #GNode beneath the parent before the given sibling.
func (p *Node) InsertBefore(sibling *Node, node *Node) *Node {
	var arg0 *C.GNode
	var arg1 *C.GNode
	var arg2 *C.GNode

	arg0 = (*C.GNode)(p.Native())
	arg1 = (*C.GNode)(sibling.Native())
	arg2 = (*C.GNode)(node.Native())

	ret := C.g_node_insert_before(arg0, arg1, arg2)

	var ret0 *Node

	{
		ret0 = WrapNode(unsafe.Pointer(ret))
	}

	return ret0
}

// IsAncestor returns true if @node is an ancestor of @descendant. This is true
// if node is the parent of @descendant, or if node is the grandparent of
// @descendant etc.
func (n *Node) IsAncestor(descendant *Node) bool {
	var arg0 *C.GNode
	var arg1 *C.GNode

	arg0 = (*C.GNode)(n.Native())
	arg1 = (*C.GNode)(descendant.Native())

	ret := C.g_node_is_ancestor(arg0, arg1)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// LastChild gets the last child of a #GNode.
func (n *Node) LastChild() *Node {
	var arg0 *C.GNode

	arg0 = (*C.GNode)(n.Native())

	ret := C.g_node_last_child(arg0)

	var ret0 *Node

	{
		ret0 = WrapNode(unsafe.Pointer(ret))
	}

	return ret0
}

// LastSibling gets the last sibling of a #GNode. This could possibly be the
// node itself.
func (n *Node) LastSibling() *Node {
	var arg0 *C.GNode

	arg0 = (*C.GNode)(n.Native())

	ret := C.g_node_last_sibling(arg0)

	var ret0 *Node

	{
		ret0 = WrapNode(unsafe.Pointer(ret))
	}

	return ret0
}

// MaxHeight gets the maximum height of all branches beneath a #GNode. This is
// the maximum distance from the #GNode to all leaf nodes.
//
// If @root is nil, 0 is returned. If @root has no children, 1 is returned. If
// @root has children, 2 is returned. And so on.
func (r *Node) MaxHeight() uint {
	var arg0 *C.GNode

	arg0 = (*C.GNode)(r.Native())

	ret := C.g_node_max_height(arg0)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// NChildren gets the number of children of a #GNode.
func (n *Node) NChildren() uint {
	var arg0 *C.GNode

	arg0 = (*C.GNode)(n.Native())

	ret := C.g_node_n_children(arg0)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// NNodes gets the number of nodes in a tree.
func (r *Node) NNodes(flags TraverseFlags) uint {
	var arg0 *C.GNode
	var arg1 C.GTraverseFlags

	arg0 = (*C.GNode)(r.Native())
	arg1 = (C.GTraverseFlags)(flags)

	ret := C.g_node_n_nodes(arg0, arg1)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// NthChild gets a child of a #GNode, using the given index. The first child is
// at index 0. If the index is too big, nil is returned.
func (n *Node) NthChild(n uint) *Node {
	var arg0 *C.GNode
	var arg1 C.guint

	arg0 = (*C.GNode)(n.Native())
	arg1 = C.guint(n)

	ret := C.g_node_nth_child(arg0, arg1)

	var ret0 *Node

	{
		ret0 = WrapNode(unsafe.Pointer(ret))
	}

	return ret0
}

// Prepend inserts a #GNode as the first child of the given parent.
func (p *Node) Prepend(node *Node) *Node {
	var arg0 *C.GNode
	var arg1 *C.GNode

	arg0 = (*C.GNode)(p.Native())
	arg1 = (*C.GNode)(node.Native())

	ret := C.g_node_prepend(arg0, arg1)

	var ret0 *Node

	{
		ret0 = WrapNode(unsafe.Pointer(ret))
	}

	return ret0
}

// ReverseChildren reverses the order of the children of a #GNode. (It doesn't
// change the order of the grandchildren.)
func (n *Node) ReverseChildren() {
	var arg0 *C.GNode

	arg0 = (*C.GNode)(n.Native())

	C.g_node_reverse_children(arg0)
}

// Traverse traverses a tree starting at the given root #GNode. It calls the
// given function for each node visited. The traversal can be halted at any
// point by returning true from @func. @func must not do anything that would
// modify the structure of the tree.
func (r *Node) Traverse(order TraverseType, flags TraverseFlags, maxDepth int, fn NodeTraverseFunc) {
	var arg0 *C.GNode
	var arg1 C.GTraverseType
	var arg2 C.GTraverseFlags
	var arg3 C.gint
	var arg4 C.GNodeTraverseFunc
	var arg5 C.gpointer

	arg0 = (*C.GNode)(r.Native())
	arg1 = (C.GTraverseType)(order)
	arg2 = (C.GTraverseFlags)(flags)
	arg3 = C.gint(maxDepth)
	arg4 = (*[0]byte)(C.gotk4_NodeTraverseFunc)
	arg5 = C.gpointer(box.Assign(fn))

	C.g_node_traverse(arg0, arg1, arg2, arg3, arg4, arg5)
}

// Unlink unlinks a #GNode from a tree, resulting in two separate trees.
func (n *Node) Unlink() {
	var arg0 *C.GNode

	arg0 = (*C.GNode)(n.Native())

	C.g_node_unlink(arg0)
}
