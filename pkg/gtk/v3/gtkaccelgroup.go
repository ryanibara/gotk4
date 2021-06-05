// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_accel_group_get_type()), F: marshalAccelGroup},
	})
}

// AccelGroupsActivate finds the first accelerator in any AccelGroup attached to
// @object that matches @accel_key and @accel_mods, and activates that
// accelerator.
func AccelGroupsActivate(object gextras.Objector, accelKey uint, accelMods gdk.ModifierType) bool {
	var arg1 *C.GObject
	var arg2 C.guint
	var arg3 C.GdkModifierType

	arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))
	arg2 = C.guint(accelKey)
	arg3 = (C.GdkModifierType)(accelMods)

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_accel_groups_activate(object, accelKey, accelMods)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// AccelGroupsFromObject gets a list of all accel groups which are attached to
// @object.
func AccelGroupsFromObject(object gextras.Objector) *glib.SList {
	var arg1 *C.GObject

	arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))

	var cret *C.GSList
	var ret1 *glib.SList

	cret = C.gtk_accel_groups_from_object(object)

	ret1 = glib.WrapSList(unsafe.Pointer(cret))

	return ret1
}

// AcceleratorGetDefaultModMask gets the modifier mask.
//
// The modifier mask determines which modifiers are considered significant for
// keyboard accelerators. See gtk_accelerator_set_default_mod_mask().
func AcceleratorGetDefaultModMask() gdk.ModifierType {
	var cret C.GdkModifierType
	var ret1 gdk.ModifierType

	cret = C.gtk_accelerator_get_default_mod_mask()

	ret1 = gdk.ModifierType(cret)

	return ret1
}

// AcceleratorGetLabel converts an accelerator keyval and modifier mask into a
// string which can be used to represent the accelerator to the user.
func AcceleratorGetLabel(acceleratorKey uint, acceleratorMods gdk.ModifierType) string {
	var arg1 C.guint
	var arg2 C.GdkModifierType

	arg1 = C.guint(acceleratorKey)
	arg2 = (C.GdkModifierType)(acceleratorMods)

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_accelerator_get_label(acceleratorKey, acceleratorMods)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// AcceleratorGetLabelWithKeycode converts an accelerator keyval and modifier
// mask into a (possibly translated) string that can be displayed to a user,
// similarly to gtk_accelerator_get_label(), but handling keycodes.
//
// This is only useful for system-level components, applications should use
// gtk_accelerator_parse() instead.
func AcceleratorGetLabelWithKeycode(display gdk.Display, acceleratorKey uint, keycode uint, acceleratorMods gdk.ModifierType) string {
	var arg1 *C.GdkDisplay
	var arg2 C.guint
	var arg3 C.guint
	var arg4 C.GdkModifierType

	arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	arg2 = C.guint(acceleratorKey)
	arg3 = C.guint(keycode)
	arg4 = (C.GdkModifierType)(acceleratorMods)

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_accelerator_get_label_with_keycode(display, acceleratorKey, keycode, acceleratorMods)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// AcceleratorName converts an accelerator keyval and modifier mask into a
// string parseable by gtk_accelerator_parse(). For example, if you pass in
// K_KEY_q and K_CONTROL_MASK, this function returns “<Control>q”.
//
// If you need to display accelerators in the user interface, see
// gtk_accelerator_get_label().
func AcceleratorName(acceleratorKey uint, acceleratorMods gdk.ModifierType) string {
	var arg1 C.guint
	var arg2 C.GdkModifierType

	arg1 = C.guint(acceleratorKey)
	arg2 = (C.GdkModifierType)(acceleratorMods)

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_accelerator_name(acceleratorKey, acceleratorMods)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// AcceleratorNameWithKeycode converts an accelerator keyval and modifier mask
// into a string parseable by gtk_accelerator_parse_with_keycode(), similarly to
// gtk_accelerator_name() but handling keycodes. This is only useful for
// system-level components, applications should use gtk_accelerator_parse()
// instead.
func AcceleratorNameWithKeycode(display gdk.Display, acceleratorKey uint, keycode uint, acceleratorMods gdk.ModifierType) string {
	var arg1 *C.GdkDisplay
	var arg2 C.guint
	var arg3 C.guint
	var arg4 C.GdkModifierType

	arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	arg2 = C.guint(acceleratorKey)
	arg3 = C.guint(keycode)
	arg4 = (C.GdkModifierType)(acceleratorMods)

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_accelerator_name_with_keycode(display, acceleratorKey, keycode, acceleratorMods)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// AcceleratorParse parses a string representing an accelerator. The format
// looks like “<Control>a” or “<Shift><Alt>F1” or “<Release>z” (the last one is
// for key release).
//
// The parser is fairly liberal and allows lower or upper case, and also
// abbreviations such as “<Ctl>” and “<Ctrl>”. Key names are parsed using
// gdk_keyval_from_name(). For character keys the name is not the symbol, but
// the lowercase name, e.g. one would use “<Ctrl>minus” instead of “<Ctrl>-”.
//
// If the parse fails, @accelerator_key and @accelerator_mods will be set to 0
// (zero).
func AcceleratorParse(accelerator string) (acceleratorKey uint, acceleratorMods gdk.ModifierType) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(accelerator))
	defer C.free(unsafe.Pointer(arg1))

	var arg2 C.guint
	var ret2 uint
	var arg3 C.GdkModifierType
	var ret3 *gdk.ModifierType

	C.gtk_accelerator_parse(accelerator, &arg2, &arg3)

	*ret2 = C.guint(arg2)
	*ret3 = *gdk.ModifierType(arg3)

	return ret2, ret3
}

// AcceleratorParseWithKeycode parses a string representing an accelerator,
// similarly to gtk_accelerator_parse() but handles keycodes as well. This is
// only useful for system-level components, applications should use
// gtk_accelerator_parse() instead.
//
// If @accelerator_codes is given and the result stored in it is non-nil, the
// result must be freed with g_free().
//
// If a keycode is present in the accelerator and no @accelerator_codes is
// given, the parse will fail.
//
// If the parse fails, @accelerator_key, @accelerator_mods and
// @accelerator_codes will be set to 0 (zero).
func AcceleratorParseWithKeycode(accelerator string) (acceleratorKey uint, acceleratorCodes []uint, acceleratorMods gdk.ModifierType) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(accelerator))
	defer C.free(unsafe.Pointer(arg1))

	var arg2 C.guint
	var ret2 uint
	var arg3 **C.guint
	var ret3 []uint
	var arg4 C.GdkModifierType
	var ret4 *gdk.ModifierType

	C.gtk_accelerator_parse_with_keycode(accelerator, &arg2, &arg3, &arg4)

	*ret2 = C.guint(arg2)
	{
		var length int
		for p := arg3; *p != 0; p = (**C.guint)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		ret3 = make([]uint, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.guint)(ptr.Add(unsafe.Pointer(arg3), i))
			ret3[i] = *C.guint(src)
		}
	}
	*ret4 = *gdk.ModifierType(arg4)

	return ret2, ret3, ret4
}

// AcceleratorSetDefaultModMask sets the modifiers that will be considered
// significant for keyboard accelerators. The default mod mask depends on the
// GDK backend in use, but will typically include K_CONTROL_MASK | K_SHIFT_MASK
// | K_MOD1_MASK | K_SUPER_MASK | K_HYPER_MASK | K_META_MASK. In other words,
// Control, Shift, Alt, Super, Hyper and Meta. Other modifiers will by default
// be ignored by AccelGroup.
//
// You must include at least the three modifiers Control, Shift and Alt in any
// value you pass to this function.
//
// The default mod mask should be changed on application startup, before using
// any accelerator groups.
func AcceleratorSetDefaultModMask(defaultModMask gdk.ModifierType) {
	var arg1 C.GdkModifierType

	arg1 = (C.GdkModifierType)(defaultModMask)

	C.gtk_accelerator_set_default_mod_mask(defaultModMask)
}

// AcceleratorValid determines whether a given keyval and modifier mask
// constitute a valid keyboard accelerator. For example, the K_KEY_a keyval plus
// K_CONTROL_MASK is valid - this is a “Ctrl+a” accelerator. But, you can't, for
// instance, use the K_KEY_Control_L keyval as an accelerator.
func AcceleratorValid(keyval uint, modifiers gdk.ModifierType) bool {
	var arg1 C.guint
	var arg2 C.GdkModifierType

	arg1 = C.guint(keyval)
	arg2 = (C.GdkModifierType)(modifiers)

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_accelerator_valid(keyval, modifiers)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// AccelGroup: a AccelGroup represents a group of keyboard accelerators,
// typically attached to a toplevel Window (with gtk_window_add_accel_group()).
// Usually you won’t need to create a AccelGroup directly; instead, when using
// UIManager, GTK+ automatically sets up the accelerators for your menus in the
// ui manager’s AccelGroup.
//
// Note that “accelerators” are different from “mnemonics”. Accelerators are
// shortcuts for activating a menu item; they appear alongside the menu item
// they’re a shortcut for. For example “Ctrl+Q” might appear alongside the
// “Quit” menu item. Mnemonics are shortcuts for GUI elements such as text
// entries or buttons; they appear as underlined characters. See
// gtk_label_new_with_mnemonic(). Menu items can have both accelerators and
// mnemonics, of course.
type AccelGroup interface {
	gextras.Objector

	// DisconnectKey removes an accelerator previously installed through
	// gtk_accel_group_connect().
	DisconnectKey(accelKey uint, accelMods gdk.ModifierType) bool
	// Find finds the first entry in an accelerator group for which @find_func
	// returns true and returns its AccelKey.
	Find(findFunc AccelGroupFindFunc) *AccelKey
	// IsLocked locks are added and removed using gtk_accel_group_lock() and
	// gtk_accel_group_unlock().
	IsLocked() bool
	// ModifierMask gets a ModifierType representing the mask for this
	// @accel_group. For example, K_CONTROL_MASK, K_SHIFT_MASK, etc.
	ModifierMask() gdk.ModifierType
	// Lock locks the given accelerator group.
	//
	// Locking an acelerator group prevents the accelerators contained within it
	// to be changed during runtime. Refer to gtk_accel_map_change_entry() about
	// runtime accelerator changes.
	//
	// If called more than once, @accel_group remains locked until
	// gtk_accel_group_unlock() has been called an equivalent number of times.
	Lock()
	// Query queries an accelerator group for all entries matching @accel_key
	// and @accel_mods.
	Query(accelKey uint, accelMods gdk.ModifierType) (nEntries uint, accelGroupEntrys []AccelGroupEntry)
	// Unlock undoes the last call to gtk_accel_group_lock() on this
	// @accel_group.
	Unlock()
}

// accelGroup implements the AccelGroup interface.
type accelGroup struct {
	gextras.Objector
}

var _ AccelGroup = (*accelGroup)(nil)

// WrapAccelGroup wraps a GObject to the right type. It is
// primarily used internally.
func WrapAccelGroup(obj *externglib.Object) AccelGroup {
	return AccelGroup{
		Objector: obj,
	}
}

func marshalAccelGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAccelGroup(obj), nil
}

// NewAccelGroup constructs a class AccelGroup.
func NewAccelGroup() AccelGroup {
	var cret C.GtkAccelGroup
	var ret1 AccelGroup

	cret = C.gtk_accel_group_new()

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(AccelGroup)

	return ret1
}

// DisconnectKey removes an accelerator previously installed through
// gtk_accel_group_connect().
func (a accelGroup) DisconnectKey(accelKey uint, accelMods gdk.ModifierType) bool {
	var arg0 *C.GtkAccelGroup
	var arg1 C.guint
	var arg2 C.GdkModifierType

	arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(a.Native()))
	arg1 = C.guint(accelKey)
	arg2 = (C.GdkModifierType)(accelMods)

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_accel_group_disconnect_key(arg0, accelKey, accelMods)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Find finds the first entry in an accelerator group for which @find_func
// returns true and returns its AccelKey.
func (a accelGroup) Find(findFunc AccelGroupFindFunc) *AccelKey {
	var arg0 *C.GtkAccelGroup

	arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(a.Native()))

	var cret *C.GtkAccelKey
	var ret1 *AccelKey

	cret = C.gtk_accel_group_find(arg0, findFunc, data)

	ret1 = WrapAccelKey(unsafe.Pointer(cret))

	return ret1
}

// IsLocked locks are added and removed using gtk_accel_group_lock() and
// gtk_accel_group_unlock().
func (a accelGroup) IsLocked() bool {
	var arg0 *C.GtkAccelGroup

	arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(a.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_accel_group_get_is_locked(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// ModifierMask gets a ModifierType representing the mask for this
// @accel_group. For example, K_CONTROL_MASK, K_SHIFT_MASK, etc.
func (a accelGroup) ModifierMask() gdk.ModifierType {
	var arg0 *C.GtkAccelGroup

	arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(a.Native()))

	var cret C.GdkModifierType
	var ret1 gdk.ModifierType

	cret = C.gtk_accel_group_get_modifier_mask(arg0)

	ret1 = gdk.ModifierType(cret)

	return ret1
}

// Lock locks the given accelerator group.
//
// Locking an acelerator group prevents the accelerators contained within it
// to be changed during runtime. Refer to gtk_accel_map_change_entry() about
// runtime accelerator changes.
//
// If called more than once, @accel_group remains locked until
// gtk_accel_group_unlock() has been called an equivalent number of times.
func (a accelGroup) Lock() {
	var arg0 *C.GtkAccelGroup

	arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(a.Native()))

	C.gtk_accel_group_lock(arg0)
}

// Query queries an accelerator group for all entries matching @accel_key
// and @accel_mods.
func (a accelGroup) Query(accelKey uint, accelMods gdk.ModifierType) (nEntries uint, accelGroupEntrys []AccelGroupEntry) {
	var arg0 *C.GtkAccelGroup
	var arg1 C.guint
	var arg2 C.GdkModifierType

	arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(a.Native()))
	arg1 = C.guint(accelKey)
	arg2 = (C.GdkModifierType)(accelMods)

	var cret *C.GtkAccelGroupEntry
	var arg3 *C.guint
	var ret2 []AccelGroupEntry

	cret = C.gtk_accel_group_query(arg0, accelKey, accelMods, &arg3)

	ret2 = make([]AccelGroupEntry, arg3)
	for i := 0; i < uintptr(arg3); i++ {
		src := (C.GtkAccelGroupEntry)(ptr.Add(unsafe.Pointer(cret), i))
		ret2[i] = WrapAccelGroupEntry(unsafe.Pointer(src))
	}

	return ret3, ret2
}

// Unlock undoes the last call to gtk_accel_group_lock() on this
// @accel_group.
func (a accelGroup) Unlock() {
	var arg0 *C.GtkAccelGroup

	arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(a.Native()))

	C.gtk_accel_group_unlock(arg0)
}

type AccelGroupEntry struct {
	native C.GtkAccelGroupEntry
}

// WrapAccelGroupEntry wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAccelGroupEntry(ptr unsafe.Pointer) *AccelGroupEntry {
	if ptr == nil {
		return nil
	}

	return (*AccelGroupEntry)(ptr)
}

func marshalAccelGroupEntry(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAccelGroupEntry(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AccelGroupEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// Key gets the field inside the struct.
func (a *AccelGroupEntry) Key() AccelKey {
	v = WrapAccelKey(unsafe.Pointer(a.native.key))
}

type AccelKey struct {
	native C.GtkAccelKey
}

// WrapAccelKey wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAccelKey(ptr unsafe.Pointer) *AccelKey {
	if ptr == nil {
		return nil
	}

	return (*AccelKey)(ptr)
}

func marshalAccelKey(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAccelKey(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AccelKey) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// AccelKey gets the field inside the struct.
func (a *AccelKey) AccelKey() uint {
	v = C.guint(a.native.accel_key)
}

// AccelMods gets the field inside the struct.
func (a *AccelKey) AccelMods() gdk.ModifierType {
	v = gdk.ModifierType(a.native.accel_mods)
}
