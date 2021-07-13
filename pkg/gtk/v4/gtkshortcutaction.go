// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcut_action_flags_get_type()), F: marshalShortcutActionFlags},
		{T: externglib.Type(C.gtk_activate_action_get_type()), F: marshalActivateActioner},
		{T: externglib.Type(C.gtk_callback_action_get_type()), F: marshalCallbackActioner},
		{T: externglib.Type(C.gtk_mnemonic_action_get_type()), F: marshalMnemonicActioner},
		{T: externglib.Type(C.gtk_named_action_get_type()), F: marshalNamedActioner},
		{T: externglib.Type(C.gtk_nothing_action_get_type()), F: marshalNothingActioner},
		{T: externglib.Type(C.gtk_shortcut_action_get_type()), F: marshalShortcutActioner},
		{T: externglib.Type(C.gtk_signal_action_get_type()), F: marshalSignalActioner},
	})
}

// ShortcutActionFlags: list of flags that can be passed to action activation.
//
// More flags may be added in the future.
type ShortcutActionFlags int

const (
	// ShortcutActionFlagsExclusive: action is the only action that can be
	// activated. If this flag is not set, a future activation may select a
	// different action.
	ShortcutActionFlagsExclusive ShortcutActionFlags = 0b1
)

func marshalShortcutActionFlags(p uintptr) (interface{}, error) {
	return ShortcutActionFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ShortcutFunc: prototype for shortcuts based on user callbacks.
type ShortcutFunc func(widget *Widget, args *glib.Variant, userData cgo.Handle) (ok bool)

//export gotk4_ShortcutFunc
func gotk4_ShortcutFunc(arg0 *C.GtkWidget, arg1 *C.GVariant, arg2 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var widget *Widget      // out
	var args *glib.Variant  // out
	var userData cgo.Handle // out

	widget = wrapWidget(externglib.Take(unsafe.Pointer(arg0)))
	args = (*glib.Variant)(unsafe.Pointer(arg1))
	C.g_variant_ref(arg1)
	runtime.SetFinalizer(args, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(unsafe.Pointer(v)))
	})
	userData = (cgo.Handle)(unsafe.Pointer(arg2))

	fn := v.(ShortcutFunc)
	ok := fn(widget, args, userData)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ActivateActioner describes ActivateAction's methods.
type ActivateActioner interface {
	privateActivateAction()
}

// ActivateAction: GtkShortcutAction that calls gtk_widget_activate().
type ActivateAction struct {
	ShortcutAction
}

var (
	_ ActivateActioner = (*ActivateAction)(nil)
	_ gextras.Nativer  = (*ActivateAction)(nil)
)

func wrapActivateAction(obj *externglib.Object) *ActivateAction {
	return &ActivateAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalActivateActioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapActivateAction(obj), nil
}

func (*ActivateAction) privateActivateAction() {}

// ActivateActionGet gets the activate action.
//
// This is an action that calls gtk_widget_activate() on the given widget upon
// activation.
func ActivateActionGet() *ActivateAction {
	var _cret *C.GtkShortcutAction // in

	_cret = C.gtk_activate_action_get()

	var _activateAction *ActivateAction // out

	_activateAction = wrapActivateAction(externglib.Take(unsafe.Pointer(_cret)))

	return _activateAction
}

// CallbackActioner describes CallbackAction's methods.
type CallbackActioner interface {
	privateCallbackAction()
}

// CallbackAction: GtkShortcutAction that invokes a callback.
type CallbackAction struct {
	ShortcutAction
}

var (
	_ CallbackActioner = (*CallbackAction)(nil)
	_ gextras.Nativer  = (*CallbackAction)(nil)
)

func wrapCallbackAction(obj *externglib.Object) *CallbackAction {
	return &CallbackAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalCallbackActioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCallbackAction(obj), nil
}

func (*CallbackAction) privateCallbackAction() {}

// MnemonicActioner describes MnemonicAction's methods.
type MnemonicActioner interface {
	privateMnemonicAction()
}

// MnemonicAction: GtkShortcutAction that calls gtk_widget_mnemonic_activate().
type MnemonicAction struct {
	ShortcutAction
}

var (
	_ MnemonicActioner = (*MnemonicAction)(nil)
	_ gextras.Nativer  = (*MnemonicAction)(nil)
)

func wrapMnemonicAction(obj *externglib.Object) *MnemonicAction {
	return &MnemonicAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalMnemonicActioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMnemonicAction(obj), nil
}

func (*MnemonicAction) privateMnemonicAction() {}

// MnemonicActionGet gets the mnemonic action.
//
// This is an action that calls gtk_widget_mnemonic_activate() on the given
// widget upon activation.
func MnemonicActionGet() *MnemonicAction {
	var _cret *C.GtkShortcutAction // in

	_cret = C.gtk_mnemonic_action_get()

	var _mnemonicAction *MnemonicAction // out

	_mnemonicAction = wrapMnemonicAction(externglib.Take(unsafe.Pointer(_cret)))

	return _mnemonicAction
}

// NamedActioner describes NamedAction's methods.
type NamedActioner interface {
	// ActionName returns the name of the action that will be activated.
	ActionName() string
}

// NamedAction: GtkShortcutAction that activates an action by name.
type NamedAction struct {
	ShortcutAction
}

var (
	_ NamedActioner   = (*NamedAction)(nil)
	_ gextras.Nativer = (*NamedAction)(nil)
)

func wrapNamedAction(obj *externglib.Object) *NamedAction {
	return &NamedAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalNamedActioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNamedAction(obj), nil
}

// NewNamedAction creates an action that when activated, activates the named
// action on the widget.
//
// It also passes the given arguments to it.
//
// See gtk.Widget.InsertActionGroup() for how to add actions to widgets.
func NewNamedAction(name string) *NamedAction {
	var _arg1 *C.char              // out
	var _cret *C.GtkShortcutAction // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))

	_cret = C.gtk_named_action_new(_arg1)

	var _namedAction *NamedAction // out

	_namedAction = wrapNamedAction(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _namedAction
}

// ActionName returns the name of the action that will be activated.
func (self *NamedAction) ActionName() string {
	var _arg0 *C.GtkNamedAction // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkNamedAction)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_named_action_get_action_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// NothingActioner describes NothingAction's methods.
type NothingActioner interface {
	privateNothingAction()
}

// NothingAction: GtkShortcutAction that does nothing.
type NothingAction struct {
	ShortcutAction
}

var (
	_ NothingActioner = (*NothingAction)(nil)
	_ gextras.Nativer = (*NothingAction)(nil)
)

func wrapNothingAction(obj *externglib.Object) *NothingAction {
	return &NothingAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalNothingActioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNothingAction(obj), nil
}

func (*NothingAction) privateNothingAction() {}

// NothingActionGet gets the nothing action.
//
// This is an action that does nothing and where activating it always fails.
func NothingActionGet() *NothingAction {
	var _cret *C.GtkShortcutAction // in

	_cret = C.gtk_nothing_action_get()

	var _nothingAction *NothingAction // out

	_nothingAction = wrapNothingAction(externglib.Take(unsafe.Pointer(_cret)))

	return _nothingAction
}

// ShortcutActioner describes ShortcutAction's methods.
type ShortcutActioner interface {
	// Activate activates the action on the widget with the given args.
	Activate(flags ShortcutActionFlags, widget Widgeter, args *glib.Variant) bool
	// String prints the given action into a human-readable string.
	String() string
}

// ShortcutAction: GtkShortcutAction encodes an action that can be triggered by
// a keyboard shortcut.
//
// GtkShortcutActions contain functions that allow easy presentation to end
// users as well as being printed for debugging.
//
// All GtkShortcutActions are immutable, you can only specify their properties
// during construction. If you want to change a action, you have to replace it
// with a new one. If you need to pass arguments to an action, these are
// specified by the higher-level GtkShortcut object.
//
// To activate a GtkShortcutAction manually, gtk.ShortcutAction.Activate() can
// be called.
//
// GTK provides various actions:
//
//    - gtk.MnemonicAction: a shortcut action that calls
//      gtk_widget_mnemonic_activate()
//    - gtk.CallbackAction: a shortcut action that invokes
//      a given callback
//    - gtk.SignalAction: a shortcut action that emits a
//      given signal
//    - gtk.ActivateAction: a shortcut action that calls
//      gtk_widget_activate()
//    - gtk.NamedAction: a shortcut action that calls
//      gtk_widget_activate_action()
//    - gtk.NothingAction: a shortcut action that does nothing
type ShortcutAction struct {
	*externglib.Object
}

var (
	_ ShortcutActioner = (*ShortcutAction)(nil)
	_ gextras.Nativer  = (*ShortcutAction)(nil)
)

func wrapShortcutAction(obj *externglib.Object) *ShortcutAction {
	return &ShortcutAction{
		Object: obj,
	}
}

func marshalShortcutActioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapShortcutAction(obj), nil
}

// NewShortcutActionParseString tries to parse the given string into an action.
//
// On success, the parsed action is returned. When parsing failed, NULL is
// returned.
//
// The accepted strings are:
//
// - nothing, for GtkNothingAction
//
// - activate, for GtkActivateAction
//
// - mnemonic-activate, for GtkMnemonicAction
//
// - action(NAME), for a GtkNamedAction for the action named NAME
//
// - signal(NAME), for a GtkSignalAction for the signal NAME
func NewShortcutActionParseString(_string string) *ShortcutAction {
	var _arg1 *C.char              // out
	var _cret *C.GtkShortcutAction // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(_string)))

	_cret = C.gtk_shortcut_action_parse_string(_arg1)

	var _shortcutAction *ShortcutAction // out

	_shortcutAction = wrapShortcutAction(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _shortcutAction
}

// Activate activates the action on the widget with the given args.
//
// Note that some actions ignore the passed in flags, widget or args.
//
// Activation of an action can fail for various reasons. If the action is not
// supported by the widget, if the args don't match the action or if the
// activation otherwise had no effect, FALSE will be returned.
func (self *ShortcutAction) Activate(flags ShortcutActionFlags, widget Widgeter, args *glib.Variant) bool {
	var _arg0 *C.GtkShortcutAction     // out
	var _arg1 C.GtkShortcutActionFlags // out
	var _arg2 *C.GtkWidget             // out
	var _arg3 *C.GVariant              // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkShortcutAction)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkShortcutActionFlags(flags)
	_arg2 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))
	_arg3 = (*C.GVariant)(unsafe.Pointer(args))

	_cret = C.gtk_shortcut_action_activate(_arg0, _arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// String prints the given action into a human-readable string.
//
// This is a small wrapper around gtk.ShortcutAction.Print() to help when
// debugging.
func (self *ShortcutAction) String() string {
	var _arg0 *C.GtkShortcutAction // out
	var _cret *C.char              // in

	_arg0 = (*C.GtkShortcutAction)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_shortcut_action_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SignalActioner describes SignalAction's methods.
type SignalActioner interface {
	// SignalName returns the name of the signal that will be emitted.
	SignalName() string
}

// SignalAction: GtkShortcutAction that emits a signal.
//
// Signals that are used in this way are referred to as keybinding signals, and
// they are expected to be defined with the G_SIGNAL_ACTION flag.
type SignalAction struct {
	ShortcutAction
}

var (
	_ SignalActioner  = (*SignalAction)(nil)
	_ gextras.Nativer = (*SignalAction)(nil)
)

func wrapSignalAction(obj *externglib.Object) *SignalAction {
	return &SignalAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalSignalActioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSignalAction(obj), nil
}

// NewSignalAction creates an action that when activated, emits the given action
// signal on the provided widget.
//
// It will also unpack the args into arguments passed to the signal.
func NewSignalAction(signalName string) *SignalAction {
	var _arg1 *C.char              // out
	var _cret *C.GtkShortcutAction // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(signalName)))

	_cret = C.gtk_signal_action_new(_arg1)

	var _signalAction *SignalAction // out

	_signalAction = wrapSignalAction(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _signalAction
}

// SignalName returns the name of the signal that will be emitted.
func (self *SignalAction) SignalName() string {
	var _arg0 *C.GtkSignalAction // out
	var _cret *C.char            // in

	_arg0 = (*C.GtkSignalAction)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_signal_action_get_signal_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
