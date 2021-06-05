// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_stack_sidebar_get_type()), F: marshalStackSidebar},
	})
}

// StackSidebar: a GtkStackSidebar enables you to quickly and easily provide a
// consistent "sidebar" object for your user interface.
//
// In order to use a GtkStackSidebar, you simply use a GtkStack to organize your
// UI flow, and add the sidebar to your sidebar area. You can use
// gtk_stack_sidebar_set_stack() to connect the StackSidebar to the Stack.
//
//
// CSS nodes
//
// GtkStackSidebar has a single CSS node with name stacksidebar and style class
// .sidebar.
//
// When circumstances require it, GtkStackSidebar adds the .needs-attention
// style class to the widgets representing the stack pages.
type StackSidebar interface {
	Bin
	Buildable

	// Stack retrieves the stack. See gtk_stack_sidebar_set_stack().
	Stack() Stack
	// SetStack: set the Stack associated with this StackSidebar.
	//
	// The sidebar widget will automatically update according to the order
	// (packing) and items within the given Stack.
	SetStack(stack Stack)
}

// stackSidebar implements the StackSidebar interface.
type stackSidebar struct {
	Bin
	Buildable
}

var _ StackSidebar = (*stackSidebar)(nil)

// WrapStackSidebar wraps a GObject to the right type. It is
// primarily used internally.
func WrapStackSidebar(obj *externglib.Object) StackSidebar {
	return StackSidebar{
		Bin:       WrapBin(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalStackSidebar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStackSidebar(obj), nil
}

// NewStackSidebar constructs a class StackSidebar.
func NewStackSidebar() StackSidebar {
	var cret C.GtkStackSidebar
	var ret1 StackSidebar

	cret = C.gtk_stack_sidebar_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(StackSidebar)

	return ret1
}

// Stack retrieves the stack. See gtk_stack_sidebar_set_stack().
func (s stackSidebar) Stack() Stack {
	var arg0 *C.GtkStackSidebar

	arg0 = (*C.GtkStackSidebar)(unsafe.Pointer(s.Native()))

	var cret *C.GtkStack
	var ret1 Stack

	cret = C.gtk_stack_sidebar_get_stack(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Stack)

	return ret1
}

// SetStack: set the Stack associated with this StackSidebar.
//
// The sidebar widget will automatically update according to the order
// (packing) and items within the given Stack.
func (s stackSidebar) SetStack(stack Stack) {
	var arg0 *C.GtkStackSidebar
	var arg1 *C.GtkStack

	arg0 = (*C.GtkStackSidebar)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	C.gtk_stack_sidebar_set_stack(arg0, stack)
}

type StackSidebarPrivate struct {
	native C.GtkStackSidebarPrivate
}

// WrapStackSidebarPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapStackSidebarPrivate(ptr unsafe.Pointer) *StackSidebarPrivate {
	if ptr == nil {
		return nil
	}

	return (*StackSidebarPrivate)(ptr)
}

func marshalStackSidebarPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapStackSidebarPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *StackSidebarPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}
