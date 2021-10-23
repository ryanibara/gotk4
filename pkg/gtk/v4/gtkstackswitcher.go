// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_stack_switcher_get_type()), F: marshalStackSwitcherer},
	})
}

// StackSwitcher: GtkStackSwitcher shows a row of buttons to switch between
// GtkStack pages.
//
// !An example GtkStackSwitcher (stackswitcher.png)
//
// It acts as a controller for the associated GtkStack.
//
// All the content for the buttons comes from the properties of the stacks
// gtk.StackPage objects; the button visibility in a GtkStackSwitcher widget is
// controlled by the visibility of the child in the GtkStack.
//
// It is possible to associate multiple GtkStackSwitcher widgets with the same
// GtkStack widget.
//
//
// CSS nodes
//
// GtkStackSwitcher has a single CSS node named stackswitcher and style class
// .stack-switcher.
//
// When circumstances require it, GtkStackSwitcher adds the .needs-attention
// style class to the widgets representing the stack pages.
//
//
// Accessibility
//
// GtkStackSwitcher uses the GTK_ACCESSIBLE_ROLE_TAB_LIST role and uses the
// GTK_ACCESSIBLE_ROLE_TAB for its buttons.
type StackSwitcher struct {
	Widget
}

func wrapStackSwitcher(obj *externglib.Object) *StackSwitcher {
	return &StackSwitcher{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalStackSwitcherer(p uintptr) (interface{}, error) {
	return wrapStackSwitcher(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStackSwitcher: create a new GtkStackSwitcher.
func NewStackSwitcher() *StackSwitcher {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_stack_switcher_new()

	var _stackSwitcher *StackSwitcher // out

	_stackSwitcher = wrapStackSwitcher(externglib.Take(unsafe.Pointer(_cret)))

	return _stackSwitcher
}

// Stack retrieves the stack.
func (switcher *StackSwitcher) Stack() *Stack {
	var _arg0 *C.GtkStackSwitcher // out
	var _cret *C.GtkStack         // in

	_arg0 = (*C.GtkStackSwitcher)(unsafe.Pointer(switcher.Native()))

	_cret = C.gtk_stack_switcher_get_stack(_arg0)
	runtime.KeepAlive(switcher)

	var _stack *Stack // out

	if _cret != nil {
		_stack = wrapStack(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _stack
}

// SetStack sets the stack to control.
//
// The function takes the following parameters:
//
//    - stack: GtkStack.
//
func (switcher *StackSwitcher) SetStack(stack *Stack) {
	var _arg0 *C.GtkStackSwitcher // out
	var _arg1 *C.GtkStack         // out

	_arg0 = (*C.GtkStackSwitcher)(unsafe.Pointer(switcher.Native()))
	if stack != nil {
		_arg1 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	}

	C.gtk_stack_switcher_set_stack(_arg0, _arg1)
	runtime.KeepAlive(switcher)
	runtime.KeepAlive(stack)
}
