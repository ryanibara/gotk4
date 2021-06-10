// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_activate_action_get_type()), F: marshalActivateAction},
		{T: externglib.Type(C.gtk_callback_action_get_type()), F: marshalCallbackAction},
		{T: externglib.Type(C.gtk_mnemonic_action_get_type()), F: marshalMnemonicAction},
		{T: externglib.Type(C.gtk_named_action_get_type()), F: marshalNamedAction},
		{T: externglib.Type(C.gtk_nothing_action_get_type()), F: marshalNothingAction},
		{T: externglib.Type(C.gtk_shortcut_action_get_type()), F: marshalShortcutAction},
		{T: externglib.Type(C.gtk_signal_action_get_type()), F: marshalSignalAction},
	})
}

// ActivateAction: a `GtkShortcutAction` that calls gtk_widget_activate().
type ActivateAction interface {
	ShortcutAction
}

// activateAction implements the ActivateAction interface.
type activateAction struct {
	ShortcutAction
}

var _ ActivateAction = (*activateAction)(nil)

// WrapActivateAction wraps a GObject to the right type. It is
// primarily used internally.
func WrapActivateAction(obj *externglib.Object) ActivateAction {
	return ActivateAction{
		ShortcutAction: WrapShortcutAction(obj),
	}
}

func marshalActivateAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapActivateAction(obj), nil
}

// CallbackAction: a `GtkShortcutAction` that invokes a callback.
type CallbackAction interface {
	ShortcutAction
}

// callbackAction implements the CallbackAction interface.
type callbackAction struct {
	ShortcutAction
}

var _ CallbackAction = (*callbackAction)(nil)

// WrapCallbackAction wraps a GObject to the right type. It is
// primarily used internally.
func WrapCallbackAction(obj *externglib.Object) CallbackAction {
	return CallbackAction{
		ShortcutAction: WrapShortcutAction(obj),
	}
}

func marshalCallbackAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCallbackAction(obj), nil
}

// MnemonicAction: a `GtkShortcutAction` that calls
// gtk_widget_mnemonic_activate().
type MnemonicAction interface {
	ShortcutAction
}

// mnemonicAction implements the MnemonicAction interface.
type mnemonicAction struct {
	ShortcutAction
}

var _ MnemonicAction = (*mnemonicAction)(nil)

// WrapMnemonicAction wraps a GObject to the right type. It is
// primarily used internally.
func WrapMnemonicAction(obj *externglib.Object) MnemonicAction {
	return MnemonicAction{
		ShortcutAction: WrapShortcutAction(obj),
	}
}

func marshalMnemonicAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMnemonicAction(obj), nil
}

// NamedAction: a `GtkShortcutAction` that activates an action by name.
type NamedAction interface {
	ShortcutAction

	// ActionName returns the name of the action that will be activated.
	ActionName() string
}

// namedAction implements the NamedAction interface.
type namedAction struct {
	ShortcutAction
}

var _ NamedAction = (*namedAction)(nil)

// WrapNamedAction wraps a GObject to the right type. It is
// primarily used internally.
func WrapNamedAction(obj *externglib.Object) NamedAction {
	return NamedAction{
		ShortcutAction: WrapShortcutAction(obj),
	}
}

func marshalNamedAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNamedAction(obj), nil
}

// ActionName returns the name of the action that will be activated.
func (s namedAction) ActionName() string {
	var _arg0 *C.GtkNamedAction

	_arg0 = (*C.GtkNamedAction)(unsafe.Pointer(s.Native()))

	var _cret *C.char

	_cret = C.gtk_named_action_get_action_name(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// NothingAction: a `GtkShortcutAction` that does nothing.
type NothingAction interface {
	ShortcutAction
}

// nothingAction implements the NothingAction interface.
type nothingAction struct {
	ShortcutAction
}

var _ NothingAction = (*nothingAction)(nil)

// WrapNothingAction wraps a GObject to the right type. It is
// primarily used internally.
func WrapNothingAction(obj *externglib.Object) NothingAction {
	return NothingAction{
		ShortcutAction: WrapShortcutAction(obj),
	}
}

func marshalNothingAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNothingAction(obj), nil
}

// ShortcutAction: `GtkShortcutAction` encodes an action that can be triggered
// by a keyboard shortcut.
//
// `GtkShortcutActions` contain functions that allow easy presentation to end
// users as well as being printed for debugging.
//
// All `GtkShortcutActions` are immutable, you can only specify their properties
// during construction. If you want to change a action, you have to replace it
// with a new one. If you need to pass arguments to an action, these are
// specified by the higher-level `GtkShortcut` object.
//
// To activate a `GtkShortcutAction` manually,
// [method@Gtk.ShortcutAction.activate] can be called.
//
// GTK provides various actions:
//
//    - [class@Gtk.MnemonicAction]: a shortcut action that calls
//      gtk_widget_mnemonic_activate()
//    - [class@Gtk.CallbackAction]: a shortcut action that invokes
//      a given callback
//    - [class@Gtk.SignalAction]: a shortcut action that emits a
//      given signal
//    - [class@Gtk.ActivateAction]: a shortcut action that calls
//      gtk_widget_activate()
//    - [class@Gtk.NamedAction]: a shortcut action that calls
//      gtk_widget_activate_action()
//    - [class@Gtk.NothingAction]: a shortcut action that does nothing
type ShortcutAction interface {
	gextras.Objector

	// Activate activates the action on the @widget with the given @args.
	//
	// Note that some actions ignore the passed in @flags, @widget or @args.
	//
	// Activation of an action can fail for various reasons. If the action is
	// not supported by the @widget, if the @args don't match the action or if
	// the activation otherwise had no effect, false will be returned.
	Activate(flags ShortcutActionFlags, widget Widget, args *glib.Variant) bool
	// Print prints the given action into a string for the developer.
	//
	// This is meant for debugging and logging.
	//
	// The form of the representation may change at any time and is not
	// guaranteed to stay identical.
	Print(string *glib.String)
	// String prints the given action into a human-readable string.
	//
	// This is a small wrapper around [method@Gtk.ShortcutAction.print] to help
	// when debugging.
	String() string
}

// shortcutAction implements the ShortcutAction interface.
type shortcutAction struct {
	gextras.Objector
}

var _ ShortcutAction = (*shortcutAction)(nil)

// WrapShortcutAction wraps a GObject to the right type. It is
// primarily used internally.
func WrapShortcutAction(obj *externglib.Object) ShortcutAction {
	return ShortcutAction{
		Objector: obj,
	}
}

func marshalShortcutAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapShortcutAction(obj), nil
}

// Activate activates the action on the @widget with the given @args.
//
// Note that some actions ignore the passed in @flags, @widget or @args.
//
// Activation of an action can fail for various reasons. If the action is
// not supported by the @widget, if the @args don't match the action or if
// the activation otherwise had no effect, false will be returned.
func (s shortcutAction) Activate(flags ShortcutActionFlags, widget Widget, args *glib.Variant) bool {
	var _arg0 *C.GtkShortcutAction
	var _arg1 C.GtkShortcutActionFlags
	var _arg2 *C.GtkWidget
	var _arg3 *C.GVariant

	_arg0 = (*C.GtkShortcutAction)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkShortcutActionFlags)(flags)
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.GVariant)(unsafe.Pointer(args.Native()))

	var _cret C.gboolean

	_cret = C.gtk_shortcut_action_activate(_arg0, _arg1, _arg2, _arg3)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Print prints the given action into a string for the developer.
//
// This is meant for debugging and logging.
//
// The form of the representation may change at any time and is not
// guaranteed to stay identical.
func (s shortcutAction) Print(string *glib.String) {
	var _arg0 *C.GtkShortcutAction
	var _arg1 *C.GString

	_arg0 = (*C.GtkShortcutAction)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GString)(unsafe.Pointer(string.Native()))

	C.gtk_shortcut_action_print(_arg0, _arg1)
}

// String prints the given action into a human-readable string.
//
// This is a small wrapper around [method@Gtk.ShortcutAction.print] to help
// when debugging.
func (s shortcutAction) String() string {
	var _arg0 *C.GtkShortcutAction

	_arg0 = (*C.GtkShortcutAction)(unsafe.Pointer(s.Native()))

	var _cret *C.char

	_cret = C.gtk_shortcut_action_to_string(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SignalAction: a `GtkShortcut`Action that emits a signal.
//
// Signals that are used in this way are referred to as keybinding signals, and
// they are expected to be defined with the G_SIGNAL_ACTION flag.
type SignalAction interface {
	ShortcutAction

	// SignalName returns the name of the signal that will be emitted.
	SignalName() string
}

// signalAction implements the SignalAction interface.
type signalAction struct {
	ShortcutAction
}

var _ SignalAction = (*signalAction)(nil)

// WrapSignalAction wraps a GObject to the right type. It is
// primarily used internally.
func WrapSignalAction(obj *externglib.Object) SignalAction {
	return SignalAction{
		ShortcutAction: WrapShortcutAction(obj),
	}
}

func marshalSignalAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSignalAction(obj), nil
}

// SignalName returns the name of the signal that will be emitted.
func (s signalAction) SignalName() string {
	var _arg0 *C.GtkSignalAction

	_arg0 = (*C.GtkSignalAction)(unsafe.Pointer(s.Native()))

	var _cret *C.char

	_cret = C.gtk_signal_action_get_signal_name(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}
