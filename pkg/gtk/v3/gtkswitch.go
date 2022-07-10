// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gtk3_SwitchClass_state_set(void*, gboolean);
// extern gboolean _gotk4_gtk3_Switch_ConnectStateSet(gpointer, gboolean, guintptr);
// extern void _gotk4_gtk3_SwitchClass_activate(void*);
// extern void _gotk4_gtk3_Switch_ConnectActivate(gpointer, guintptr);
import "C"

// GTypeSwitch returns the GType for the type Switch.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeSwitch() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Switch").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalSwitch)
	return gtype
}

// SwitchOverrider contains methods that are overridable.
type SwitchOverrider interface {
	Activate()
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	StateSet(state bool) bool
}

// Switch is a widget that has two states: on or off. The user can control which
// state should be active by clicking the empty area, or by dragging the handle.
//
// GtkSwitch can also handle situations where the underlying state changes with
// a delay. See Switch::state-set for details.
//
// CSS nodes
//
//    switch
//    ╰── slider
//
// GtkSwitch has two css nodes, the main node with the name switch and a subnode
// named slider. Neither of them is using any style classes.
type Switch struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Actionable
	Activatable
}

var (
	_ Widgetter         = (*Switch)(nil)
	_ coreglib.Objector = (*Switch)(nil)
)

func classInitSwitcher(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "SwitchClass")

	if _, ok := goval.(interface{ Activate() }); ok {
		o := pclass.StructFieldOffset("activate")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_SwitchClass_activate)
	}

	if _, ok := goval.(interface{ StateSet(state bool) bool }); ok {
		o := pclass.StructFieldOffset("state_set")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_SwitchClass_state_set)
	}
}

//export _gotk4_gtk3_SwitchClass_activate
func _gotk4_gtk3_SwitchClass_activate(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Activate() })

	iface.Activate()
}

//export _gotk4_gtk3_SwitchClass_state_set
func _gotk4_gtk3_SwitchClass_state_set(arg0 *C.void, arg1 C.gboolean) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ StateSet(state bool) bool })

	var _state bool // out

	if arg1 != 0 {
		_state = true
	}

	ok := iface.StateSet(_state)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapSwitch(obj *coreglib.Object) *Switch {
	return &Switch{
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
		Object: obj,
		Actionable: Actionable{
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
		Activatable: Activatable{
			Object: obj,
		},
	}
}

func marshalSwitch(p uintptr) (interface{}, error) {
	return wrapSwitch(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Switch_ConnectActivate
func _gotk4_gtk3_Switch_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectActivate signal on GtkSwitch is an action signal and emitting it
// causes the switch to animate. Applications should never connect to this
// signal, but use the notify::active signal.
func (sw *Switch) ConnectActivate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(sw, "activate", false, unsafe.Pointer(C._gotk4_gtk3_Switch_ConnectActivate), f)
}

//export _gotk4_gtk3_Switch_ConnectStateSet
func _gotk4_gtk3_Switch_ConnectStateSet(arg0 C.gpointer, arg1 C.gboolean, arg2 C.guintptr) (cret C.gboolean) {
	var f func(state bool) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(state bool) (ok bool))
	}

	var _state bool // out

	if arg1 != 0 {
		_state = true
	}

	ok := f(_state)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectStateSet signal on GtkSwitch is emitted to change the underlying
// state. It is emitted when the user changes the switch position. The default
// handler keeps the state in sync with the Switch:active property.
//
// To implement delayed state change, applications can connect to this signal,
// initiate the change of the underlying state, and call gtk_switch_set_state()
// when the underlying state change is complete. The signal handler should
// return TRUE to prevent the default handler from running.
//
// Visually, the underlying state is represented by the trough color of the
// switch, while the Switch:active property is represented by the position of
// the switch.
func (sw *Switch) ConnectStateSet(f func(state bool) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(sw, "state-set", false, unsafe.Pointer(C._gotk4_gtk3_Switch_ConnectStateSet), f)
}

// NewSwitch creates a new Switch widget.
//
// The function returns the following values:
//
//    - _switch: newly created Switch instance.
//
func NewSwitch() *Switch {
	_info := girepository.MustFind("Gtk", "Switch")
	_gret := _info.InvokeClassMethod("new_Switch", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var __switch *Switch // out

	__switch = wrapSwitch(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return __switch
}

// Active gets whether the Switch is in its “on” or “off” state.
//
// The function returns the following values:
//
//    - ok: TRUE if the Switch is active, and FALSE otherwise.
//
func (sw *Switch) Active() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(sw).Native()))

	_info := girepository.MustFind("Gtk", "Switch")
	_gret := _info.InvokeClassMethod("get_active", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(sw)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// State gets the underlying state of the Switch.
//
// The function returns the following values:
//
//    - ok: underlying state.
//
func (sw *Switch) State() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(sw).Native()))

	_info := girepository.MustFind("Gtk", "Switch")
	_gret := _info.InvokeClassMethod("get_state", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(sw)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetActive changes the state of sw to the desired one.
//
// The function takes the following parameters:
//
//    - isActive: TRUE if sw should be active, and FALSE otherwise.
//
func (sw *Switch) SetActive(isActive bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(sw).Native()))
	if isActive {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Switch")
	_info.InvokeClassMethod("set_active", _args[:], nil)

	runtime.KeepAlive(sw)
	runtime.KeepAlive(isActive)
}

// SetState sets the underlying state of the Switch.
//
// Normally, this is the same as Switch:active, unless the switch is set up for
// delayed state changes. This function is typically called from a
// Switch::state-set signal handler.
//
// See Switch::state-set for details.
//
// The function takes the following parameters:
//
//    - state: new state.
//
func (sw *Switch) SetState(state bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(sw).Native()))
	if state {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Switch")
	_info.InvokeClassMethod("set_state", _args[:], nil)

	runtime.KeepAlive(sw)
	runtime.KeepAlive(state)
}
