// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_stack_sidebar_get_type()), F: marshalStackSidebarrer},
	})
}

// StackSidebar: GtkStackSidebar uses a sidebar to switch between GtkStack
// pages.
//
// In order to use a GtkStackSidebar, you simply use a GtkStack to organize your
// UI flow, and add the sidebar to your sidebar area. You can use
// gtk.StackSidebar.SetStack() to connect the GtkStackSidebar to the GtkStack.
//
//
// CSS nodes
//
// GtkStackSidebar has a single CSS node with name stacksidebar and style class
// .sidebar.
//
// When circumstances require it, GtkStackSidebar adds the .needs-attention
// style class to the widgets representing the stack pages.
type StackSidebar struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*StackSidebar)(nil)
)

func wrapStackSidebar(obj *externglib.Object) *StackSidebar {
	return &StackSidebar{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalStackSidebarrer(p uintptr) (interface{}, error) {
	return wrapStackSidebar(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStackSidebar creates a new GtkStackSidebar.
//
// The function returns the following values:
//
//    - stackSidebar: new GtkStackSidebar.
//
func NewStackSidebar() *StackSidebar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_stack_sidebar_new()

	var _stackSidebar *StackSidebar // out

	_stackSidebar = wrapStackSidebar(externglib.Take(unsafe.Pointer(_cret)))

	return _stackSidebar
}

// Stack retrieves the stack.
//
// The function returns the following values:
//
//    - stack (optional): associated Stack or NULL if none has been set
//      explicitly.
//
func (self *StackSidebar) Stack() *Stack {
	var _arg0 *C.GtkStackSidebar // out
	var _cret *C.GtkStack        // in

	_arg0 = (*C.GtkStackSidebar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_stack_sidebar_get_stack(_arg0)
	runtime.KeepAlive(self)

	var _stack *Stack // out

	if _cret != nil {
		_stack = wrapStack(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _stack
}

// SetStack: set the GtkStack associated with this GtkStackSidebar.
//
// The sidebar widget will automatically update according to the order and items
// within the given GtkStack.
//
// The function takes the following parameters:
//
//    - stack: GtkStack.
//
func (self *StackSidebar) SetStack(stack *Stack) {
	var _arg0 *C.GtkStackSidebar // out
	var _arg1 *C.GtkStack        // out

	_arg0 = (*C.GtkStackSidebar)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	C.gtk_stack_sidebar_set_stack(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(stack)
}
