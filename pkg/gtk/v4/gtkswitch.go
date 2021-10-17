// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_switch_get_type()), F: marshalSwitcher},
	})
}

// Switch: GtkSwitch is a "light switch" that has two states: on or off.
//
// !An example GtkSwitch (switch.png)
//
// The user can control which state should be active by clicking the empty area,
// or by dragging the handle.
//
// GtkSwitch can also handle situations where the underlying state changes with
// a delay. See gtkswitch::state-set for details.
//
// CSS nodes
//
//    switch
//    ├── label
//    ├── label
//    ╰── slider
//
//
// GtkSwitch has four css nodes, the main node with the name switch and subnodes
// for the slider and the on and off labels. Neither of them is using any style
// classes.
//
//
// Accessibility
//
// GtkSwitch uses the GTK_ACCESSIBLE_ROLE_SWITCH role.
type Switch struct {
	Widget

	Actionable
	*externglib.Object
}

func wrapSwitch(obj *externglib.Object) *Switch {
	return &Switch{
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
		Actionable: Actionable{
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
		},
		Object: obj,
	}
}

func marshalSwitcher(p uintptr) (interface{}, error) {
	return wrapSwitch(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSwitch creates a new GtkSwitch widget.
func NewSwitch() *Switch {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_switch_new()

	var __switch *Switch // out

	__switch = wrapSwitch(externglib.Take(unsafe.Pointer(_cret)))

	return __switch
}

// Active gets whether the GtkSwitch is in its “on” or “off” state.
func (self *Switch) Active() bool {
	var _arg0 *C.GtkSwitch // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_switch_get_active(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// State gets the underlying state of the GtkSwitch.
func (self *Switch) State() bool {
	var _arg0 *C.GtkSwitch // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_switch_get_state(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive changes the state of self to the desired one.
//
// The function takes the following parameters:
//
//    - isActive: TRUE if self should be active, and FALSE otherwise.
//
func (self *Switch) SetActive(isActive bool) {
	var _arg0 *C.GtkSwitch // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(self.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_switch_set_active(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(isActive)
}

// SetState sets the underlying state of the GtkSwitch.
//
// Normally, this is the same as gtk.Switch:active, unless the switch is set up
// for delayed state changes. This function is typically called from a
// gtk.Switch`::state-set signal handler.
//
// See gtk.Switch::state-set for details.
//
// The function takes the following parameters:
//
//    - state: new state.
//
func (self *Switch) SetState(state bool) {
	var _arg0 *C.GtkSwitch // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(self.Native()))
	if state {
		_arg1 = C.TRUE
	}

	C.gtk_switch_set_state(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(state)
}

// ConnectActivate: emitted to animate the switch.
//
// Applications should never connect to this signal, but use the
// gtk.Switch:active property.
func (self *Switch) ConnectActivate(f func()) externglib.SignalHandle {
	return self.Connect("activate", f)
}

// ConnectStateSet: emitted to change the underlying state.
//
// The ::state-set signal is emitted when the user changes the switch position.
// The default handler keeps the state in sync with the gtk.Switch:active
// property.
//
// To implement delayed state change, applications can connect to this signal,
// initiate the change of the underlying state, and call gtk.Switch.SetState()
// when the underlying state change is complete. The signal handler should
// return TRUE to prevent the default handler from running.
//
// Visually, the underlying state is represented by the trough color of the
// switch, while the gtk.Switch`:active property is represented by the position
// of the switch.
func (self *Switch) ConnectStateSet(f func(state bool) bool) externglib.SignalHandle {
	return self.Connect("state-set", f)
}
