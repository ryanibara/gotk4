// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkshortcuttrigger.go.
var (
	GTypeAlternativeTrigger = coreglib.Type(C.gtk_alternative_trigger_get_type())
	GTypeKeyvalTrigger      = coreglib.Type(C.gtk_keyval_trigger_get_type())
	GTypeMnemonicTrigger    = coreglib.Type(C.gtk_mnemonic_trigger_get_type())
	GTypeNeverTrigger       = coreglib.Type(C.gtk_never_trigger_get_type())
	GTypeShortcutTrigger    = coreglib.Type(C.gtk_shortcut_trigger_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeAlternativeTrigger, F: marshalAlternativeTrigger},
		{T: GTypeKeyvalTrigger, F: marshalKeyvalTrigger},
		{T: GTypeMnemonicTrigger, F: marshalMnemonicTrigger},
		{T: GTypeNeverTrigger, F: marshalNeverTrigger},
		{T: GTypeShortcutTrigger, F: marshalShortcutTrigger},
	})
}

// AlternativeTrigger: GtkShortcutTrigger that combines two triggers.
//
// The GtkAlternativeTrigger triggers when either of two trigger.
//
// This can be cascaded to combine more than two triggers.
type AlternativeTrigger struct {
	_ [0]func() // equal guard
	ShortcutTrigger
}

var (
	_ ShortcutTriggerer = (*AlternativeTrigger)(nil)
)

func wrapAlternativeTrigger(obj *coreglib.Object) *AlternativeTrigger {
	return &AlternativeTrigger{
		ShortcutTrigger: ShortcutTrigger{
			Object: obj,
		},
	}
}

func marshalAlternativeTrigger(p uintptr) (interface{}, error) {
	return wrapAlternativeTrigger(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAlternativeTrigger creates a GtkShortcutTrigger that will trigger whenever
// either of the two given triggers gets triggered.
//
// Note that nesting is allowed, so if you want more than two alternative,
// create a new alternative trigger for each option.
//
// The function takes the following parameters:
//
//    - first trigger that may trigger.
//    - second trigger that may trigger.
//
// The function returns the following values:
//
//    - alternativeTrigger: new GtkShortcutTrigger.
//
func NewAlternativeTrigger(first, second ShortcutTriggerer) *AlternativeTrigger {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(first).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(first).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(second).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(second).Native()))

	_gret := girepository.MustFind("Gtk", "AlternativeTrigger").InvokeMethod("new_AlternativeTrigger", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(first)
	runtime.KeepAlive(second)

	var _alternativeTrigger *AlternativeTrigger // out

	_alternativeTrigger = wrapAlternativeTrigger(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _alternativeTrigger
}

// First gets the first of the two alternative triggers that may trigger self.
//
// gtk.AlternativeTrigger.GetSecond() will return the other one.
//
// The function returns the following values:
//
//    - shortcutTrigger: first alternative trigger.
//
func (self *AlternativeTrigger) First() ShortcutTriggerer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "AlternativeTrigger").InvokeMethod("get_first", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _shortcutTrigger ShortcutTriggerer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.ShortcutTriggerer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(ShortcutTriggerer)
			return ok
		})
		rv, ok := casted.(ShortcutTriggerer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.ShortcutTriggerer")
		}
		_shortcutTrigger = rv
	}

	return _shortcutTrigger
}

// Second gets the second of the two alternative triggers that may trigger self.
//
// gtk.AlternativeTrigger.GetFirst() will return the other one.
//
// The function returns the following values:
//
//    - shortcutTrigger: second alternative trigger.
//
func (self *AlternativeTrigger) Second() ShortcutTriggerer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "AlternativeTrigger").InvokeMethod("get_second", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _shortcutTrigger ShortcutTriggerer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.ShortcutTriggerer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(ShortcutTriggerer)
			return ok
		})
		rv, ok := casted.(ShortcutTriggerer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.ShortcutTriggerer")
		}
		_shortcutTrigger = rv
	}

	return _shortcutTrigger
}

// KeyvalTrigger: GtkShortcutTrigger that triggers when a specific keyval and
// modifiers are pressed.
type KeyvalTrigger struct {
	_ [0]func() // equal guard
	ShortcutTrigger
}

var (
	_ ShortcutTriggerer = (*KeyvalTrigger)(nil)
)

func wrapKeyvalTrigger(obj *coreglib.Object) *KeyvalTrigger {
	return &KeyvalTrigger{
		ShortcutTrigger: ShortcutTrigger{
			Object: obj,
		},
	}
}

func marshalKeyvalTrigger(p uintptr) (interface{}, error) {
	return wrapKeyvalTrigger(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Keyval gets the keyval that must be pressed to succeed triggering self.
//
// The function returns the following values:
//
//    - guint: keyval.
//
func (self *KeyvalTrigger) Keyval() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "KeyvalTrigger").InvokeMethod("get_keyval", _args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// MnemonicTrigger: GtkShortcutTrigger that triggers when a specific mnemonic is
// pressed.
//
// Mnemonics require a *mnemonic modifier* (typically <kbd>Alt</kbd>) to be
// pressed together with the mnemonic key.
type MnemonicTrigger struct {
	_ [0]func() // equal guard
	ShortcutTrigger
}

var (
	_ ShortcutTriggerer = (*MnemonicTrigger)(nil)
)

func wrapMnemonicTrigger(obj *coreglib.Object) *MnemonicTrigger {
	return &MnemonicTrigger{
		ShortcutTrigger: ShortcutTrigger{
			Object: obj,
		},
	}
}

func marshalMnemonicTrigger(p uintptr) (interface{}, error) {
	return wrapMnemonicTrigger(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMnemonicTrigger creates a GtkShortcutTrigger that will trigger whenever
// the key with the given keyval is pressed and mnemonics have been activated.
//
// Mnemonics are activated by calling code when a key event with the right
// modifiers is detected.
//
// The function takes the following parameters:
//
//    - keyval to trigger for.
//
// The function returns the following values:
//
//    - mnemonicTrigger: new GtkShortcutTrigger.
//
func NewMnemonicTrigger(keyval uint32) *MnemonicTrigger {
	var _args [1]girepository.Argument

	*(*C.guint)(unsafe.Pointer(&_args[0])) = C.guint(keyval)

	_gret := girepository.MustFind("Gtk", "MnemonicTrigger").InvokeMethod("new_MnemonicTrigger", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(keyval)

	var _mnemonicTrigger *MnemonicTrigger // out

	_mnemonicTrigger = wrapMnemonicTrigger(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mnemonicTrigger
}

// Keyval gets the keyval that must be pressed to succeed triggering self.
//
// The function returns the following values:
//
//    - guint: keyval.
//
func (self *MnemonicTrigger) Keyval() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "MnemonicTrigger").InvokeMethod("get_keyval", _args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// NeverTrigger: GtkShortcutTrigger that never triggers.
type NeverTrigger struct {
	_ [0]func() // equal guard
	ShortcutTrigger
}

var (
	_ ShortcutTriggerer = (*NeverTrigger)(nil)
)

func wrapNeverTrigger(obj *coreglib.Object) *NeverTrigger {
	return &NeverTrigger{
		ShortcutTrigger: ShortcutTrigger{
			Object: obj,
		},
	}
}

func marshalNeverTrigger(p uintptr) (interface{}, error) {
	return wrapNeverTrigger(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NeverTriggerGet gets the never trigger.
//
// This is a singleton for a trigger that never triggers. Use this trigger
// instead of NULL because it implements all virtual functions.
//
// The function returns the following values:
//
//    - neverTrigger: never trigger.
//
func NeverTriggerGet() *NeverTrigger {
	_gret := girepository.MustFind("Gtk", "get").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _neverTrigger *NeverTrigger // out

	_neverTrigger = wrapNeverTrigger(coreglib.Take(unsafe.Pointer(_cret)))

	return _neverTrigger
}

// ShortcutTrigger: GtkShortcutTrigger tracks how a GtkShortcut should be
// activated.
//
// To find out if a GtkShortcutTrigger triggers, you can call
// gtk.ShortcutTrigger.Trigger() on a GdkEvent.
//
// GtkShortcutTriggers contain functions that allow easy presentation to end
// users as well as being printed for debugging.
//
// All GtkShortcutTriggers are immutable, you can only specify their properties
// during construction. If you want to change a trigger, you have to replace it
// with a new one.
type ShortcutTrigger struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ShortcutTrigger)(nil)
)

// ShortcutTriggerer describes types inherited from class ShortcutTrigger.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type ShortcutTriggerer interface {
	coreglib.Objector
	baseShortcutTrigger() *ShortcutTrigger
}

var _ ShortcutTriggerer = (*ShortcutTrigger)(nil)

func wrapShortcutTrigger(obj *coreglib.Object) *ShortcutTrigger {
	return &ShortcutTrigger{
		Object: obj,
	}
}

func marshalShortcutTrigger(p uintptr) (interface{}, error) {
	return wrapShortcutTrigger(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (trigger1 *ShortcutTrigger) baseShortcutTrigger() *ShortcutTrigger {
	return trigger1
}

// BaseShortcutTrigger returns the underlying base object.
func BaseShortcutTrigger(obj ShortcutTriggerer) *ShortcutTrigger {
	return obj.baseShortcutTrigger()
}

// NewShortcutTriggerParseString tries to parse the given string into a trigger.
//
// On success, the parsed trigger is returned. When parsing failed, NULL is
// returned.
//
// The accepted strings are:
//
//    - never, for GtkNeverTrigger
//    - a string parsed by gtk_accelerator_parse(), for a GtkKeyvalTrigger, e.g. <Control>C
//    - underscore, followed by a single character, for GtkMnemonicTrigger, e.g. _l
//    - two valid trigger strings, separated by a | character, for a
//      GtkAlternativeTrigger: <Control>q|<Control>w
//
// Note that you will have to escape the < and > characters when specifying
// triggers in XML files, such as GtkBuilder ui files. Use &lt; instead of < and
// &gt; instead of >.
//
// The function takes the following parameters:
//
//    - str: string to parse.
//
// The function returns the following values:
//
//    - shortcutTrigger (optional): new GtkShortcutTrigger or NULL on error.
//
func NewShortcutTriggerParseString(str string) *ShortcutTrigger {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_args[0]))

	_gret := girepository.MustFind("Gtk", "ShortcutTrigger").InvokeMethod("new_ShortcutTrigger_parse_string", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(str)

	var _shortcutTrigger *ShortcutTrigger // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_shortcutTrigger = wrapShortcutTrigger(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _shortcutTrigger
}

// Compare types of trigger1 and trigger2 are #gconstpointer only to allow use
// of this function as a Func.
//
// They must each be a GtkShortcutTrigger.
//
// The function takes the following parameters:
//
//    - trigger2: GtkShortcutTrigger.
//
// The function returns the following values:
//
//    - gint: integer less than, equal to, or greater than zero if trigger1 is
//      found, respectively, to be less than, to match, or be greater than
//      trigger2.
//
func (trigger1 *ShortcutTrigger) Compare(trigger2 ShortcutTriggerer) int32 {
	var _args [2]girepository.Argument

	*(*C.gpointer)(unsafe.Pointer(&_args[0])) = C.gpointer(unsafe.Pointer(coreglib.InternObject(trigger1).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = C.gpointer(unsafe.Pointer(coreglib.InternObject(trigger2).Native()))

	_gret := girepository.MustFind("Gtk", "ShortcutTrigger").InvokeMethod("compare", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(trigger1)
	runtime.KeepAlive(trigger2)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// Equal checks if trigger1 and trigger2 trigger under the same conditions.
//
// The types of one and two are #gconstpointer only to allow use of this
// function with Table. They must each be a GtkShortcutTrigger.
//
// The function takes the following parameters:
//
//    - trigger2: GtkShortcutTrigger.
//
// The function returns the following values:
//
//    - ok: TRUE if trigger1 and trigger2 are equal.
//
func (trigger1 *ShortcutTrigger) Equal(trigger2 ShortcutTriggerer) bool {
	var _args [2]girepository.Argument

	*(*C.gpointer)(unsafe.Pointer(&_args[0])) = C.gpointer(unsafe.Pointer(coreglib.InternObject(trigger1).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = C.gpointer(unsafe.Pointer(coreglib.InternObject(trigger2).Native()))

	_gret := girepository.MustFind("Gtk", "ShortcutTrigger").InvokeMethod("equal", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(trigger1)
	runtime.KeepAlive(trigger2)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Hash generates a hash value for a GtkShortcutTrigger.
//
// The output of this function is guaranteed to be the same for a given value
// only per-process. It may change between different processor architectures or
// even different versions of GTK. Do not use this function as a basis for
// building protocols or file formats.
//
// The types of trigger is #gconstpointer only to allow use of this function
// with Table. They must each be a GtkShortcutTrigger.
//
// The function returns the following values:
//
//    - guint: hash value corresponding to trigger.
//
func (trigger *ShortcutTrigger) Hash() uint32 {
	var _args [1]girepository.Argument

	*(*C.gpointer)(unsafe.Pointer(&_args[0])) = C.gpointer(unsafe.Pointer(coreglib.InternObject(trigger).Native()))

	_gret := girepository.MustFind("Gtk", "ShortcutTrigger").InvokeMethod("hash", _args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(trigger)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// ToLabel gets textual representation for the given trigger.
//
// This function is returning a translated string for presentation to end users
// for example in menu items or in help texts.
//
// The display in use may influence the resulting string in various forms, such
// as resolving hardware keycodes or by causing display-specific modifier names.
//
// The form of the representation may change at any time and is not guaranteed
// to stay identical.
//
// The function takes the following parameters:
//
//    - display: GdkDisplay to print for.
//
// The function returns the following values:
//
//    - utf8: new string.
//
func (self *ShortcutTrigger) ToLabel(display *gdk.Display) string {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_gret := girepository.MustFind("Gtk", "ShortcutTrigger").InvokeMethod("to_label", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(display)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// String prints the given trigger into a human-readable string.
//
// This is a small wrapper around gtk.ShortcutTrigger.Print() to help when
// debugging.
//
// The function returns the following values:
//
//    - utf8: new string.
//
func (self *ShortcutTrigger) String() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "ShortcutTrigger").InvokeMethod("to_string", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
