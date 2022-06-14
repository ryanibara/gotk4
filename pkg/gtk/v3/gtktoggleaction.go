// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_ToggleActionClass_toggled(void*);
// extern void _gotk4_gtk3_ToggleAction_ConnectToggled(gpointer, guintptr);
import "C"

// glib.Type values for gtktoggleaction.go.
var GTypeToggleAction = coreglib.Type(C.gtk_toggle_action_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeToggleAction, F: marshalToggleAction},
	})
}

// ToggleActionOverrider contains methods that are overridable.
type ToggleActionOverrider interface {
	// Toggled emits the “toggled” signal on the toggle action.
	//
	// Deprecated: since version 3.10.
	Toggled()
}

// ToggleAction corresponds roughly to a CheckMenuItem. It has an “active” state
// specifying whether the action has been checked or not.
type ToggleAction struct {
	_ [0]func() // equal guard
	Action
}

var (
	_ coreglib.Objector = (*ToggleAction)(nil)
)

func classInitToggleActioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "ToggleActionClass")

	if _, ok := goval.(interface{ Toggled() }); ok {
		o := pclass.StructFieldOffset("toggled")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ToggleActionClass_toggled)
	}
}

//export _gotk4_gtk3_ToggleActionClass_toggled
func _gotk4_gtk3_ToggleActionClass_toggled(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Toggled() })

	iface.Toggled()
}

func wrapToggleAction(obj *coreglib.Object) *ToggleAction {
	return &ToggleAction{
		Action: Action{
			Object: obj,
			Buildable: Buildable{
				Object: obj,
			},
		},
	}
}

func marshalToggleAction(p uintptr) (interface{}, error) {
	return wrapToggleAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_ToggleAction_ConnectToggled
func _gotk4_gtk3_ToggleAction_ConnectToggled(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectToggled: should be connected if you wish to perform an action whenever
// the ToggleAction state is changed.
func (action *ToggleAction) ConnectToggled(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(action, "toggled", false, unsafe.Pointer(C._gotk4_gtk3_ToggleAction_ConnectToggled), f)
}

// NewToggleAction creates a new ToggleAction object. To add the action to a
// ActionGroup and set the accelerator for the action, call
// gtk_action_group_add_action_with_accel().
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - name: unique name for the action.
//    - label (optional) displayed in menu items and on buttons, or NULL.
//    - tooltip (optional) for the action, or NULL.
//    - stockId (optional): stock icon to display in widgets representing the
//      action, or NULL.
//
// The function returns the following values:
//
//    - toggleAction: new ToggleAction.
//
func NewToggleAction(name, label, tooltip, stockId string) *ToggleAction {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[0]))
	if label != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_args[1]))
	}
	if tooltip != "" {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(tooltip)))
		defer C.free(unsafe.Pointer(_args[2]))
	}
	if stockId != "" {
		*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(C.CString(stockId)))
		defer C.free(unsafe.Pointer(_args[3]))
	}

	_gret := girepository.MustFind("Gtk", "ToggleAction").InvokeMethod("new_ToggleAction", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(name)
	runtime.KeepAlive(label)
	runtime.KeepAlive(tooltip)
	runtime.KeepAlive(stockId)

	var _toggleAction *ToggleAction // out

	_toggleAction = wrapToggleAction(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _toggleAction
}

// Active returns the checked state of the toggle action.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - ok: checked state of the toggle action.
//
func (action *ToggleAction) Active() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_gret := girepository.MustFind("Gtk", "ToggleAction").InvokeMethod("get_active", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(action)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// DrawAsRadio returns whether the action should have proxies like a radio
// action.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - ok: whether the action should have proxies like a radio action.
//
func (action *ToggleAction) DrawAsRadio() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_gret := girepository.MustFind("Gtk", "ToggleAction").InvokeMethod("get_draw_as_radio", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(action)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetActive sets the checked state on the toggle action.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - isActive: whether the action should be checked or not.
//
func (action *ToggleAction) SetActive(isActive bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	if isActive {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "ToggleAction").InvokeMethod("set_active", _args[:], nil)

	runtime.KeepAlive(action)
	runtime.KeepAlive(isActive)
}

// SetDrawAsRadio sets whether the action should have proxies like a radio
// action.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - drawAsRadio: whether the action should have proxies like a radio action.
//
func (action *ToggleAction) SetDrawAsRadio(drawAsRadio bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	if drawAsRadio {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "ToggleAction").InvokeMethod("set_draw_as_radio", _args[:], nil)

	runtime.KeepAlive(action)
	runtime.KeepAlive(drawAsRadio)
}

// Toggled emits the “toggled” signal on the toggle action.
//
// Deprecated: since version 3.10.
func (action *ToggleAction) Toggled() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	girepository.MustFind("Gtk", "ToggleAction").InvokeMethod("toggled", _args[:], nil)

	runtime.KeepAlive(action)
}
