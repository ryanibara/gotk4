// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_stack_switcher_get_type()), F: marshalStackSwitcher},
	})
}

// StackSwitcher: the `GtkStackSwitcher` shows a row of buttons to switch
// between `GtkStack` pages.
//
// !An example GtkStackSwitcher (stackswitcher.png)
//
// It acts as a controller for the associated `GtkStack`.
//
// All the content for the buttons comes from the properties of the stacks
// [class@Gtk.StackPage] objects; the button visibility in a `GtkStackSwitcher`
// widget is controlled by the visibility of the child in the `GtkStack`.
//
// It is possible to associate multiple `GtkStackSwitcher` widgets with the same
// `GtkStack` widget.
//
//
// CSS nodes
//
// `GtkStackSwitcher` has a single CSS node named stackswitcher and style class
// .stack-switcher.
//
// When circumstances require it, `GtkStackSwitcher` adds the .needs-attention
// style class to the widgets representing the stack pages.
//
//
// Accessibility
//
// `GtkStackSwitcher` uses the GTK_ACCESSIBLE_ROLE_TAB_LIST role and uses the
// GTK_ACCESSIBLE_ROLE_TAB for its buttons.
type StackSwitcher interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget

	// Stack retrieves the stack.
	Stack() Stack
	// SetStack sets the stack to control.
	SetStack(stack Stack)
}

// stackSwitcher implements the StackSwitcher class.
type stackSwitcher struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ StackSwitcher = (*stackSwitcher)(nil)

// WrapStackSwitcher wraps a GObject to the right type. It is
// primarily used internally.
func WrapStackSwitcher(obj *externglib.Object) StackSwitcher {
	return stackSwitcher{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalStackSwitcher(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStackSwitcher(obj), nil
}

// NewStackSwitcher constructs a class StackSwitcher.
func NewStackSwitcher() StackSwitcher {
	var _cret C.GtkStackSwitcher // in

	_cret = C.gtk_stack_switcher_new()

	var _stackSwitcher StackSwitcher // out

	_stackSwitcher = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(StackSwitcher)

	return _stackSwitcher
}

// Stack retrieves the stack.
func (s stackSwitcher) Stack() Stack {
	var _arg0 *C.GtkStackSwitcher // out
	var _cret *C.GtkStack         // in

	_arg0 = (*C.GtkStackSwitcher)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_switcher_get_stack(_arg0)

	var _stack Stack // out

	_stack = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Stack)

	return _stack
}

// SetStack sets the stack to control.
func (s stackSwitcher) SetStack(stack Stack) {
	var _arg0 *C.GtkStackSwitcher // out
	var _arg1 *C.GtkStack         // out

	_arg0 = (*C.GtkStackSwitcher)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	C.gtk_stack_switcher_set_stack(_arg0, _arg1)
}
