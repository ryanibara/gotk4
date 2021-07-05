// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_switch_get_type()), F: marshalSwitch},
	})
}

// Switch: `GtkSwitch` is a "light switch" that has two states: on or off.
//
// !An example GtkSwitch (switch.png)
//
// The user can control which state should be active by clicking the empty area,
// or by dragging the handle.
//
// `GtkSwitch` can also handle situations where the underlying state changes
// with a delay. See [signal@GtkSwitch::state-set] for details.
//
//
// CSS nodes
//
// “` switch ├── label ├── label ╰── slider “`
//
// `GtkSwitch` has four css nodes, the main node with the name switch and
// subnodes for the slider and the on and off labels. Neither of them is using
// any style classes.
//
//
// Accessibility
//
// `GtkSwitch` uses the GTK_ACCESSIBLE_ROLE_SWITCH role.
type Switch interface {
	Widget

	// AsAccessible casts the class to the Accessible interface.
	AsAccessible() Accessible
	// AsActionable casts the class to the Actionable interface.
	AsActionable() Actionable
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsConstraintTarget casts the class to the ConstraintTarget interface.
	AsConstraintTarget() ConstraintTarget

	// Active gets whether the `GtkSwitch` is in its “on” or “off” state.
	Active() bool
	// State gets the underlying state of the `GtkSwitch`.
	State() bool
	// SetActiveSwitch changes the state of @self to the desired one.
	SetActiveSwitch(isActive bool)
	// SetStateSwitch sets the underlying state of the `GtkSwitch`.
	//
	// Normally, this is the same as [property@Gtk.Switch:active], unless the
	// switch is set up for delayed state changes. This function is typically
	// called from a [signal@Gtk.Switch`::state-set] signal handler.
	//
	// See [signal@Gtk.Switch::state-set] for details.
	SetStateSwitch(state bool)
}

// _switch implements the Switch class.
type _switch struct {
	Widget
}

// WrapSwitch wraps a GObject to the right type. It is
// primarily used internally.
func WrapSwitch(obj *externglib.Object) Switch {
	return _switch{
		Widget: WrapWidget(obj),
	}
}

func marshalSwitch(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSwitch(obj), nil
}

// NewSwitch creates a new `GtkSwitch` widget.
func NewSwitch() Switch {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_switch_new()

	var __switch Switch // out

	__switch = WrapSwitch(externglib.Take(unsafe.Pointer(_cret)))

	return __switch
}

func (s _switch) Active() bool {
	var _arg0 *C.GtkSwitch // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_switch_get_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s _switch) State() bool {
	var _arg0 *C.GtkSwitch // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_switch_get_state(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s _switch) SetActiveSwitch(isActive bool) {
	var _arg0 *C.GtkSwitch // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(s.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_switch_set_active(_arg0, _arg1)
}

func (s _switch) SetStateSwitch(state bool) {
	var _arg0 *C.GtkSwitch // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(s.Native()))
	if state {
		_arg1 = C.TRUE
	}

	C.gtk_switch_set_state(_arg0, _arg1)
}

func (_ _switch) AsAccessible() Accessible {
	return WrapAccessible(gextras.InternObject(_))
}

func (_ _switch) AsActionable() Actionable {
	return WrapActionable(gextras.InternObject(_))
}

func (_ _switch) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(_))
}

func (_ _switch) AsConstraintTarget() ConstraintTarget {
	return WrapConstraintTarget(gextras.InternObject(_))
}
