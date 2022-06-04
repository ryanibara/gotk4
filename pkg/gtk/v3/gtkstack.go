// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkstack.go.
var (
	GTypeStackTransitionType = coreglib.Type(C.gtk_stack_transition_type_get_type())
	GTypeStack               = coreglib.Type(C.gtk_stack_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeStackTransitionType, F: marshalStackTransitionType},
		{T: GTypeStack, F: marshalStack},
	})
}

// StackTransitionType: these enumeration values describe the possible
// transitions between pages in a Stack widget.
//
// New values may be added to this enumeration over time.
type StackTransitionType C.gint

const (
	// StackTransitionTypeNone: no transition.
	StackTransitionTypeNone StackTransitionType = iota
	// StackTransitionTypeCrossfade: cross-fade.
	StackTransitionTypeCrossfade
	// StackTransitionTypeSlideRight: slide from left to right.
	StackTransitionTypeSlideRight
	// StackTransitionTypeSlideLeft: slide from right to left.
	StackTransitionTypeSlideLeft
	// StackTransitionTypeSlideUp: slide from bottom up.
	StackTransitionTypeSlideUp
	// StackTransitionTypeSlideDown: slide from top down.
	StackTransitionTypeSlideDown
	// StackTransitionTypeSlideLeftRight: slide from left or right according to
	// the children order.
	StackTransitionTypeSlideLeftRight
	// StackTransitionTypeSlideUpDown: slide from top down or bottom up
	// according to the order.
	StackTransitionTypeSlideUpDown
	// StackTransitionTypeOverUp: cover the old page by sliding up. Since 3.12.
	StackTransitionTypeOverUp
	// StackTransitionTypeOverDown: cover the old page by sliding down. Since:
	// 3.12.
	StackTransitionTypeOverDown
	// StackTransitionTypeOverLeft: cover the old page by sliding to the left.
	// Since: 3.12.
	StackTransitionTypeOverLeft
	// StackTransitionTypeOverRight: cover the old page by sliding to the right.
	// Since: 3.12.
	StackTransitionTypeOverRight
	// StackTransitionTypeUnderUp: uncover the new page by sliding up. Since
	// 3.12.
	StackTransitionTypeUnderUp
	// StackTransitionTypeUnderDown: uncover the new page by sliding down.
	// Since: 3.12.
	StackTransitionTypeUnderDown
	// StackTransitionTypeUnderLeft: uncover the new page by sliding to the
	// left. Since: 3.12.
	StackTransitionTypeUnderLeft
	// StackTransitionTypeUnderRight: uncover the new page by sliding to the
	// right. Since: 3.12.
	StackTransitionTypeUnderRight
	// StackTransitionTypeOverUpDown: cover the old page sliding up or uncover
	// the new page sliding down, according to order. Since: 3.12.
	StackTransitionTypeOverUpDown
	// StackTransitionTypeOverDownUp: cover the old page sliding down or uncover
	// the new page sliding up, according to order. Since: 3.14.
	StackTransitionTypeOverDownUp
	// StackTransitionTypeOverLeftRight: cover the old page sliding left or
	// uncover the new page sliding right, according to order. Since: 3.14.
	StackTransitionTypeOverLeftRight
	// StackTransitionTypeOverRightLeft: cover the old page sliding right or
	// uncover the new page sliding left, according to order. Since: 3.14.
	StackTransitionTypeOverRightLeft
)

func marshalStackTransitionType(p uintptr) (interface{}, error) {
	return StackTransitionType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for StackTransitionType.
func (s StackTransitionType) String() string {
	switch s {
	case StackTransitionTypeNone:
		return "None"
	case StackTransitionTypeCrossfade:
		return "Crossfade"
	case StackTransitionTypeSlideRight:
		return "SlideRight"
	case StackTransitionTypeSlideLeft:
		return "SlideLeft"
	case StackTransitionTypeSlideUp:
		return "SlideUp"
	case StackTransitionTypeSlideDown:
		return "SlideDown"
	case StackTransitionTypeSlideLeftRight:
		return "SlideLeftRight"
	case StackTransitionTypeSlideUpDown:
		return "SlideUpDown"
	case StackTransitionTypeOverUp:
		return "OverUp"
	case StackTransitionTypeOverDown:
		return "OverDown"
	case StackTransitionTypeOverLeft:
		return "OverLeft"
	case StackTransitionTypeOverRight:
		return "OverRight"
	case StackTransitionTypeUnderUp:
		return "UnderUp"
	case StackTransitionTypeUnderDown:
		return "UnderDown"
	case StackTransitionTypeUnderLeft:
		return "UnderLeft"
	case StackTransitionTypeUnderRight:
		return "UnderRight"
	case StackTransitionTypeOverUpDown:
		return "OverUpDown"
	case StackTransitionTypeOverDownUp:
		return "OverDownUp"
	case StackTransitionTypeOverLeftRight:
		return "OverLeftRight"
	case StackTransitionTypeOverRightLeft:
		return "OverRightLeft"
	default:
		return fmt.Sprintf("StackTransitionType(%d)", s)
	}
}

// StackOverrider contains methods that are overridable.
type StackOverrider interface {
}

// Stack widget is a container which only shows one of its children at a time.
// In contrast to GtkNotebook, GtkStack does not provide a means for users to
// change the visible child. Instead, the StackSwitcher widget can be used with
// GtkStack to provide this functionality.
//
// Transitions between pages can be animated as slides or fades. This can be
// controlled with gtk_stack_set_transition_type(). These animations respect the
// Settings:gtk-enable-animations setting.
//
// The GtkStack widget was added in GTK+ 3.10.
//
//
// CSS nodes
//
// GtkStack has a single CSS node named stack.
type Stack struct {
	_ [0]func() // equal guard
	Container
}

var (
	_ Containerer = (*Stack)(nil)
)

func classInitStacker(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapStack(obj *coreglib.Object) *Stack {
	return &Stack{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
	}
}

func marshalStack(p uintptr) (interface{}, error) {
	return wrapStack(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStack creates a new Stack container.
//
// The function returns the following values:
//
//    - stack: new Stack.
//
func NewStack() *Stack {
	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("new_Stack", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _stack *Stack // out

	_stack = wrapStack(coreglib.Take(unsafe.Pointer(_cret)))

	return _stack
}

// AddNamed adds a child to stack. The child is identified by the name.
//
// The function takes the following parameters:
//
//    - child: widget to add.
//    - name for child.
//
func (stack *Stack) AddNamed(child Widgetter, name string) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[2]))

	girepository.MustFind("Gtk", "Stack").InvokeMethod("add_named", _args[:], nil)

	runtime.KeepAlive(stack)
	runtime.KeepAlive(child)
	runtime.KeepAlive(name)
}

// AddTitled adds a child to stack. The child is identified by the name. The
// title will be used by StackSwitcher to represent child in a tab bar, so it
// should be short.
//
// The function takes the following parameters:
//
//    - child: widget to add.
//    - name for child.
//    - title: human-readable title for child.
//
func (stack *Stack) AddTitled(child Widgetter, name, title string) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[2]))
	*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_args[3]))

	girepository.MustFind("Gtk", "Stack").InvokeMethod("add_titled", _args[:], nil)

	runtime.KeepAlive(stack)
	runtime.KeepAlive(child)
	runtime.KeepAlive(name)
	runtime.KeepAlive(title)
}

// ChildByName finds the child of the Stack with the name given as the argument.
// Returns NULL if there is no child with this name.
//
// The function takes the following parameters:
//
//    - name of the child to find.
//
// The function returns the following values:
//
//    - widget (optional): requested child of the Stack.
//
func (stack *Stack) ChildByName(name string) Widgetter {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_child_by_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)
	runtime.KeepAlive(name)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Hhomogeneous gets whether stack is horizontally homogeneous. See
// gtk_stack_set_hhomogeneous().
//
// The function returns the following values:
//
//    - ok: whether stack is horizontally homogeneous.
//
func (stack *Stack) Hhomogeneous() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_hhomogeneous", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Homogeneous gets whether stack is homogeneous. See
// gtk_stack_set_homogeneous().
//
// The function returns the following values:
//
//    - ok: whether stack is homogeneous.
//
func (stack *Stack) Homogeneous() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_homogeneous", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// InterpolateSize returns wether the Stack is set up to interpolate between the
// sizes of children on page switch.
//
// The function returns the following values:
//
//    - ok: TRUE if child sizes are interpolated.
//
func (stack *Stack) InterpolateSize() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_interpolate_size", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// TransitionDuration returns the amount of time (in milliseconds) that
// transitions between pages in stack will take.
//
// The function returns the following values:
//
//    - guint: transition duration.
//
func (stack *Stack) TransitionDuration() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_transition_duration", _args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// TransitionRunning returns whether the stack is currently in a transition from
// one page to another.
//
// The function returns the following values:
//
//    - ok: TRUE if the transition is currently running, FALSE otherwise.
//
func (stack *Stack) TransitionRunning() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_transition_running", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Vhomogeneous gets whether stack is vertically homogeneous. See
// gtk_stack_set_vhomogeneous().
//
// The function returns the following values:
//
//    - ok: whether stack is vertically homogeneous.
//
func (stack *Stack) Vhomogeneous() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_vhomogeneous", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// VisibleChild gets the currently visible child of stack, or NULL if there are
// no visible children.
//
// The function returns the following values:
//
//    - widget (optional): visible child of the Stack.
//
func (stack *Stack) VisibleChild() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_visible_child", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// VisibleChildName returns the name of the currently visible child of stack, or
// NULL if there is no visible child.
//
// The function returns the following values:
//
//    - utf8 (optional): name of the visible child of the Stack.
//
func (stack *Stack) VisibleChildName() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_visible_child_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SetHhomogeneous sets the Stack to be horizontally homogeneous or not. If it
// is homogeneous, the Stack will request the same width for all its children.
// If it isn't, the stack may change width when a different child becomes
// visible.
//
// The function takes the following parameters:
//
//    - hhomogeneous: TRUE to make stack horizontally homogeneous.
//
func (stack *Stack) SetHhomogeneous(hhomogeneous bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	if hhomogeneous {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Stack").InvokeMethod("set_hhomogeneous", _args[:], nil)

	runtime.KeepAlive(stack)
	runtime.KeepAlive(hhomogeneous)
}

// SetHomogeneous sets the Stack to be homogeneous or not. If it is homogeneous,
// the Stack will request the same size for all its children. If it isn't, the
// stack may change size when a different child becomes visible.
//
// Since 3.16, homogeneity can be controlled separately for horizontal and
// vertical size, with the Stack:hhomogeneous and Stack:vhomogeneous.
//
// The function takes the following parameters:
//
//    - homogeneous: TRUE to make stack homogeneous.
//
func (stack *Stack) SetHomogeneous(homogeneous bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	if homogeneous {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Stack").InvokeMethod("set_homogeneous", _args[:], nil)

	runtime.KeepAlive(stack)
	runtime.KeepAlive(homogeneous)
}

// SetInterpolateSize sets whether or not stack will interpolate its size when
// changing the visible child. If the Stack:interpolate-size property is set to
// TRUE, stack will interpolate its size between the current one and the one
// it'll take after changing the visible child, according to the set transition
// duration.
//
// The function takes the following parameters:
//
//    - interpolateSize: new value.
//
func (stack *Stack) SetInterpolateSize(interpolateSize bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	if interpolateSize {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Stack").InvokeMethod("set_interpolate_size", _args[:], nil)

	runtime.KeepAlive(stack)
	runtime.KeepAlive(interpolateSize)
}

// SetTransitionDuration sets the duration that transitions between pages in
// stack will take.
//
// The function takes the following parameters:
//
//    - duration: new duration, in milliseconds.
//
func (stack *Stack) SetTransitionDuration(duration uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(duration)

	girepository.MustFind("Gtk", "Stack").InvokeMethod("set_transition_duration", _args[:], nil)

	runtime.KeepAlive(stack)
	runtime.KeepAlive(duration)
}

// SetVhomogeneous sets the Stack to be vertically homogeneous or not. If it is
// homogeneous, the Stack will request the same height for all its children. If
// it isn't, the stack may change height when a different child becomes visible.
//
// The function takes the following parameters:
//
//    - vhomogeneous: TRUE to make stack vertically homogeneous.
//
func (stack *Stack) SetVhomogeneous(vhomogeneous bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	if vhomogeneous {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Stack").InvokeMethod("set_vhomogeneous", _args[:], nil)

	runtime.KeepAlive(stack)
	runtime.KeepAlive(vhomogeneous)
}

// SetVisibleChild makes child the visible child of stack.
//
// If child is different from the currently visible child, the transition
// between the two will be animated with the current transition type of stack.
//
// Note that the child widget has to be visible itself (see gtk_widget_show())
// in order to become the visible child of stack.
//
// The function takes the following parameters:
//
//    - child of stack.
//
func (stack *Stack) SetVisibleChild(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	girepository.MustFind("Gtk", "Stack").InvokeMethod("set_visible_child", _args[:], nil)

	runtime.KeepAlive(stack)
	runtime.KeepAlive(child)
}

// SetVisibleChildName makes the child with the given name visible.
//
// If child is different from the currently visible child, the transition
// between the two will be animated with the current transition type of stack.
//
// Note that the child widget has to be visible itself (see gtk_widget_show())
// in order to become the visible child of stack.
//
// The function takes the following parameters:
//
//    - name of the child to make visible.
//
func (stack *Stack) SetVisibleChildName(name string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "Stack").InvokeMethod("set_visible_child_name", _args[:], nil)

	runtime.KeepAlive(stack)
	runtime.KeepAlive(name)
}
