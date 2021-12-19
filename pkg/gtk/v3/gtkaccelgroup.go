// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_accel_flags_get_type()), F: marshalAccelFlags},
		{T: externglib.Type(C.gtk_accel_group_get_type()), F: marshalAccelGrouper},
	})
}

// AccelFlags: accelerator flags used with gtk_accel_group_connect().
type AccelFlags C.guint

const (
	// AccelVisible: accelerator is visible.
	AccelVisible AccelFlags = 0b1
	// AccelLocked: accelerator not removable.
	AccelLocked AccelFlags = 0b10
	// AccelMask: mask.
	AccelMask AccelFlags = 0b111
)

func marshalAccelFlags(p uintptr) (interface{}, error) {
	return AccelFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for AccelFlags.
func (a AccelFlags) String() string {
	if a == 0 {
		return "AccelFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(34)

	for a != 0 {
		next := a & (a - 1)
		bit := a - next

		switch bit {
		case AccelVisible:
			builder.WriteString("Visible|")
		case AccelLocked:
			builder.WriteString("Locked|")
		case AccelMask:
			builder.WriteString("Mask|")
		default:
			builder.WriteString(fmt.Sprintf("AccelFlags(0b%b)|", bit))
		}

		a = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if a contains other.
func (a AccelFlags) Has(other AccelFlags) bool {
	return (a & other) == other
}

// AccelGroupsActivate finds the first accelerator in any AccelGroup attached to
// object that matches accel_key and accel_mods, and activates that accelerator.
//
// The function takes the following parameters:
//
//    - object usually a Window, on which to activate the accelerator.
//    - accelKey: accelerator keyval from a key event.
//    - accelMods: keyboard state mask from a key event.
//
// The function returns the following values:
//
//    - ok: TRUE if an accelerator was activated and handled this keypress.
//
func AccelGroupsActivate(object *externglib.Object, accelKey uint, accelMods gdk.ModifierType) bool {
	var _arg1 *C.GObject        // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out
	var _cret C.gboolean        // in

	_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))
	_arg2 = C.guint(accelKey)
	_arg3 = C.GdkModifierType(accelMods)

	_cret = C.gtk_accel_groups_activate(_arg1, _arg2, _arg3)
	runtime.KeepAlive(object)
	runtime.KeepAlive(accelKey)
	runtime.KeepAlive(accelMods)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AccelGroupsFromObject gets a list of all accel groups which are attached to
// object.
//
// The function takes the following parameters:
//
//    - object usually a Window.
//
// The function returns the following values:
//
//    - sList: list of all accel groups which are attached to object.
//
func AccelGroupsFromObject(object *externglib.Object) []AccelGroup {
	var _arg1 *C.GObject // out
	var _cret *C.GSList  // in

	_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))

	_cret = C.gtk_accel_groups_from_object(_arg1)
	runtime.KeepAlive(object)

	var _sList []AccelGroup // out

	_sList = make([]AccelGroup, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.GtkAccelGroup)(v)
		var dst AccelGroup // out
		dst = *wrapAccelGroup(externglib.Take(unsafe.Pointer(src)))
		_sList = append(_sList, dst)
	})

	return _sList
}

// AcceleratorGetDefaultModMask gets the modifier mask.
//
// The modifier mask determines which modifiers are considered significant for
// keyboard accelerators. See gtk_accelerator_set_default_mod_mask().
//
// The function returns the following values:
//
//    - modifierType: default accelerator modifier mask.
//
func AcceleratorGetDefaultModMask() gdk.ModifierType {
	var _cret C.GdkModifierType // in

	_cret = C.gtk_accelerator_get_default_mod_mask()

	var _modifierType gdk.ModifierType // out

	_modifierType = gdk.ModifierType(_cret)

	return _modifierType
}

// AcceleratorGetLabel converts an accelerator keyval and modifier mask into a
// string which can be used to represent the accelerator to the user.
//
// The function takes the following parameters:
//
//    - acceleratorKey: accelerator keyval.
//    - acceleratorMods: accelerator modifier mask.
//
// The function returns the following values:
//
//    - utf8: newly-allocated string representing the accelerator.
//
func AcceleratorGetLabel(acceleratorKey uint, acceleratorMods gdk.ModifierType) string {
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _cret *C.gchar          // in

	_arg1 = C.guint(acceleratorKey)
	_arg2 = C.GdkModifierType(acceleratorMods)

	_cret = C.gtk_accelerator_get_label(_arg1, _arg2)
	runtime.KeepAlive(acceleratorKey)
	runtime.KeepAlive(acceleratorMods)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// AcceleratorGetLabelWithKeycode converts an accelerator keyval and modifier
// mask into a (possibly translated) string that can be displayed to a user,
// similarly to gtk_accelerator_get_label(), but handling keycodes.
//
// This is only useful for system-level components, applications should use
// gtk_accelerator_parse() instead.
//
// The function takes the following parameters:
//
//    - display (optional) or NULL to use the default display.
//    - acceleratorKey: accelerator keyval.
//    - keycode: accelerator keycode.
//    - acceleratorMods: accelerator modifier mask.
//
// The function returns the following values:
//
//    - utf8: newly-allocated string representing the accelerator.
//
func AcceleratorGetLabelWithKeycode(display *gdk.Display, acceleratorKey, keycode uint, acceleratorMods gdk.ModifierType) string {
	var _arg1 *C.GdkDisplay     // out
	var _arg2 C.guint           // out
	var _arg3 C.guint           // out
	var _arg4 C.GdkModifierType // out
	var _cret *C.gchar          // in

	if display != nil {
		_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	}
	_arg2 = C.guint(acceleratorKey)
	_arg3 = C.guint(keycode)
	_arg4 = C.GdkModifierType(acceleratorMods)

	_cret = C.gtk_accelerator_get_label_with_keycode(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(display)
	runtime.KeepAlive(acceleratorKey)
	runtime.KeepAlive(keycode)
	runtime.KeepAlive(acceleratorMods)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// AcceleratorName converts an accelerator keyval and modifier mask into a
// string parseable by gtk_accelerator_parse(). For example, if you pass in
// K_KEY_q and K_CONTROL_MASK, this function returns “<Control>q”.
//
// If you need to display accelerators in the user interface, see
// gtk_accelerator_get_label().
//
// The function takes the following parameters:
//
//    - acceleratorKey: accelerator keyval.
//    - acceleratorMods: accelerator modifier mask.
//
// The function returns the following values:
//
//    - utf8: newly-allocated accelerator name.
//
func AcceleratorName(acceleratorKey uint, acceleratorMods gdk.ModifierType) string {
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _cret *C.gchar          // in

	_arg1 = C.guint(acceleratorKey)
	_arg2 = C.GdkModifierType(acceleratorMods)

	_cret = C.gtk_accelerator_name(_arg1, _arg2)
	runtime.KeepAlive(acceleratorKey)
	runtime.KeepAlive(acceleratorMods)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// AcceleratorNameWithKeycode converts an accelerator keyval and modifier mask
// into a string parseable by gtk_accelerator_parse_with_keycode(), similarly to
// gtk_accelerator_name() but handling keycodes. This is only useful for
// system-level components, applications should use gtk_accelerator_parse()
// instead.
//
// The function takes the following parameters:
//
//    - display (optional) or NULL to use the default display.
//    - acceleratorKey: accelerator keyval.
//    - keycode: accelerator keycode.
//    - acceleratorMods: accelerator modifier mask.
//
// The function returns the following values:
//
//    - utf8: newly allocated accelerator name.
//
func AcceleratorNameWithKeycode(display *gdk.Display, acceleratorKey, keycode uint, acceleratorMods gdk.ModifierType) string {
	var _arg1 *C.GdkDisplay     // out
	var _arg2 C.guint           // out
	var _arg3 C.guint           // out
	var _arg4 C.GdkModifierType // out
	var _cret *C.gchar          // in

	if display != nil {
		_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	}
	_arg2 = C.guint(acceleratorKey)
	_arg3 = C.guint(keycode)
	_arg4 = C.GdkModifierType(acceleratorMods)

	_cret = C.gtk_accelerator_name_with_keycode(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(display)
	runtime.KeepAlive(acceleratorKey)
	runtime.KeepAlive(keycode)
	runtime.KeepAlive(acceleratorMods)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
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
// If the parse fails, accelerator_key and accelerator_mods will be set to 0
// (zero).
//
// The function takes the following parameters:
//
//    - accelerator: string representing an accelerator.
//
// The function returns the following values:
//
//    - acceleratorKey (optional): return location for accelerator keyval, or
//      NULL.
//    - acceleratorMods (optional): return location for accelerator modifier
//      mask, NULL.
//
func AcceleratorParse(accelerator string) (uint, gdk.ModifierType) {
	var _arg1 *C.gchar          // out
	var _arg2 C.guint           // in
	var _arg3 C.GdkModifierType // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelerator)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_accelerator_parse(_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(accelerator)

	var _acceleratorKey uint              // out
	var _acceleratorMods gdk.ModifierType // out

	_acceleratorKey = uint(_arg2)
	_acceleratorMods = gdk.ModifierType(_arg3)

	return _acceleratorKey, _acceleratorMods
}

// AcceleratorParseWithKeycode parses a string representing an accelerator,
// similarly to gtk_accelerator_parse() but handles keycodes as well. This is
// only useful for system-level components, applications should use
// gtk_accelerator_parse() instead.
//
// If accelerator_codes is given and the result stored in it is non-NULL, the
// result must be freed with g_free().
//
// If a keycode is present in the accelerator and no accelerator_codes is given,
// the parse will fail.
//
// If the parse fails, accelerator_key, accelerator_mods and accelerator_codes
// will be set to 0 (zero).
//
// The function takes the following parameters:
//
//    - accelerator: string representing an accelerator.
//
// The function returns the following values:
//
//    - acceleratorKey (optional): return location for accelerator keyval, or
//      NULL.
//    - acceleratorCodes (optional): return location for accelerator keycodes, or
//      NULL.
//    - acceleratorMods (optional): return location for accelerator modifier
//      mask, NULL.
//
func AcceleratorParseWithKeycode(accelerator string) (uint, []uint, gdk.ModifierType) {
	var _arg1 *C.gchar          // out
	var _arg2 C.guint           // in
	var _arg3 *C.guint          // in
	var _arg4 C.GdkModifierType // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelerator)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_accelerator_parse_with_keycode(_arg1, &_arg2, &_arg3, &_arg4)
	runtime.KeepAlive(accelerator)

	var _acceleratorKey uint              // out
	var _acceleratorCodes []uint          // out
	var _acceleratorMods gdk.ModifierType // out

	_acceleratorKey = uint(_arg2)
	if _arg3 != nil {
		defer C.free(unsafe.Pointer(_arg3))
		{
			var i int
			var z C.guint
			for p := _arg3; *p != z; p = &unsafe.Slice(p, 2)[1] {
				i++
			}

			src := unsafe.Slice(_arg3, i)
			_acceleratorCodes = make([]uint, i)
			for i := range src {
				_acceleratorCodes[i] = uint(src[i])
			}
		}
	}
	_acceleratorMods = gdk.ModifierType(_arg4)

	return _acceleratorKey, _acceleratorCodes, _acceleratorMods
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
//
// The function takes the following parameters:
//
//    - defaultModMask: accelerator modifier mask.
//
func AcceleratorSetDefaultModMask(defaultModMask gdk.ModifierType) {
	var _arg1 C.GdkModifierType // out

	_arg1 = C.GdkModifierType(defaultModMask)

	C.gtk_accelerator_set_default_mod_mask(_arg1)
	runtime.KeepAlive(defaultModMask)
}

// AcceleratorValid determines whether a given keyval and modifier mask
// constitute a valid keyboard accelerator. For example, the K_KEY_a keyval plus
// K_CONTROL_MASK is valid - this is a “Ctrl+a” accelerator. But, you can't, for
// instance, use the K_KEY_Control_L keyval as an accelerator.
//
// The function takes the following parameters:
//
//    - keyval: GDK keyval.
//    - modifiers: modifier mask.
//
// The function returns the following values:
//
//    - ok: TRUE if the accelerator is valid.
//
func AcceleratorValid(keyval uint, modifiers gdk.ModifierType) bool {
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _cret C.gboolean        // in

	_arg1 = C.guint(keyval)
	_arg2 = C.GdkModifierType(modifiers)

	_cret = C.gtk_accelerator_valid(_arg1, _arg2)
	runtime.KeepAlive(keyval)
	runtime.KeepAlive(modifiers)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AccelGroupOverrider contains methods that are overridable.
type AccelGroupOverrider interface {
}

// AccelGroup represents a group of keyboard accelerators, typically attached to
// a toplevel Window (with gtk_window_add_accel_group()). Usually you won’t need
// to create a AccelGroup directly; instead, when using UIManager, GTK+
// automatically sets up the accelerators for your menus in the ui manager’s
// AccelGroup.
//
// Note that “accelerators” are different from “mnemonics”. Accelerators are
// shortcuts for activating a menu item; they appear alongside the menu item
// they’re a shortcut for. For example “Ctrl+Q” might appear alongside the
// “Quit” menu item. Mnemonics are shortcuts for GUI elements such as text
// entries or buttons; they appear as underlined characters. See
// gtk_label_new_with_mnemonic(). Menu items can have both accelerators and
// mnemonics, of course.
type AccelGroup struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*AccelGroup)(nil)
)

func classInitAccelGrouper(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapAccelGroup(obj *externglib.Object) *AccelGroup {
	return &AccelGroup{
		Object: obj,
	}
}

func marshalAccelGrouper(p uintptr) (interface{}, error) {
	return wrapAccelGroup(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectAccelActivate signal is an implementation detail of AccelGroup and not
// meant to be used by applications.
func (accelGroup *AccelGroup) ConnectAccelActivate(f func(acceleratable *externglib.Object, keyval uint, modifier gdk.ModifierType) bool) externglib.SignalHandle {
	return accelGroup.Connect("accel-activate", externglib.GeneratedClosure{Func: f})
}

// NewAccelGroup creates a new AccelGroup.
//
// The function returns the following values:
//
//    - accelGroup: new AccelGroup object.
//
func NewAccelGroup() *AccelGroup {
	var _cret *C.GtkAccelGroup // in

	_cret = C.gtk_accel_group_new()

	var _accelGroup *AccelGroup // out

	_accelGroup = wrapAccelGroup(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _accelGroup
}

// Activate finds the first accelerator in accel_group that matches accel_key
// and accel_mods, and activates it.
//
// The function takes the following parameters:
//
//    - accelQuark: quark for the accelerator name.
//    - acceleratable usually a Window, on which to activate the accelerator.
//    - accelKey: accelerator keyval from a key event.
//    - accelMods: keyboard state mask from a key event.
//
// The function returns the following values:
//
//    - ok: TRUE if an accelerator was activated and handled this keypress.
//
func (accelGroup *AccelGroup) Activate(accelQuark glib.Quark, acceleratable *externglib.Object, accelKey uint, accelMods gdk.ModifierType) bool {
	var _arg0 *C.GtkAccelGroup  // out
	var _arg1 C.GQuark          // out
	var _arg2 *C.GObject        // out
	var _arg3 C.guint           // out
	var _arg4 C.GdkModifierType // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))
	_arg1 = C.guint32(accelQuark)
	_arg2 = (*C.GObject)(unsafe.Pointer(acceleratable.Native()))
	_arg3 = C.guint(accelKey)
	_arg4 = C.GdkModifierType(accelMods)

	_cret = C.gtk_accel_group_activate(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(accelGroup)
	runtime.KeepAlive(accelQuark)
	runtime.KeepAlive(acceleratable)
	runtime.KeepAlive(accelKey)
	runtime.KeepAlive(accelMods)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ConnectAccelGroup installs an accelerator in this group. When accel_group is
// being activated in response to a call to gtk_accel_groups_activate(), closure
// will be invoked if the accel_key and accel_mods from
// gtk_accel_groups_activate() match those of this connection.
//
// The signature used for the closure is that of AccelGroupActivate.
//
// Note that, due to implementation details, a single closure can only be
// connected to one accelerator group.
//
// The function takes the following parameters:
//
//    - accelKey: key value of the accelerator.
//    - accelMods: modifier combination of the accelerator.
//    - accelFlags: flag mask to configure this accelerator.
//    - closure to be executed upon accelerator activation.
//
func (accelGroup *AccelGroup) ConnectAccelGroup(accelKey uint, accelMods gdk.ModifierType, accelFlags AccelFlags, closure externglib.AnyClosure) {
	var _arg0 *C.GtkAccelGroup  // out
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _arg3 C.GtkAccelFlags   // out
	var _arg4 *C.GClosure       // out

	_arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))
	_arg1 = C.guint(accelKey)
	_arg2 = C.GdkModifierType(accelMods)
	_arg3 = C.GtkAccelFlags(accelFlags)
	_arg4 = (*C.GClosure)(externglib.NewClosure(externglib.InternObject(accelGroup), closure))

	C.gtk_accel_group_connect(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(accelGroup)
	runtime.KeepAlive(accelKey)
	runtime.KeepAlive(accelMods)
	runtime.KeepAlive(accelFlags)
	runtime.KeepAlive(closure)
}

// ConnectByPath installs an accelerator in this group, using an accelerator
// path to look up the appropriate key and modifiers (see
// gtk_accel_map_add_entry()). When accel_group is being activated in response
// to a call to gtk_accel_groups_activate(), closure will be invoked if the
// accel_key and accel_mods from gtk_accel_groups_activate() match the key and
// modifiers for the path.
//
// The signature used for the closure is that of AccelGroupActivate.
//
// Note that accel_path string will be stored in a #GQuark. Therefore, if you
// pass a static string, you can save some memory by interning it first with
// g_intern_static_string().
//
// The function takes the following parameters:
//
//    - accelPath: path used for determining key and modifiers.
//    - closure to be executed upon accelerator activation.
//
func (accelGroup *AccelGroup) ConnectByPath(accelPath string, closure externglib.AnyClosure) {
	var _arg0 *C.GtkAccelGroup // out
	var _arg1 *C.gchar         // out
	var _arg2 *C.GClosure      // out

	_arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelPath)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GClosure)(externglib.NewClosure(externglib.InternObject(accelGroup), closure))

	C.gtk_accel_group_connect_by_path(_arg0, _arg1, _arg2)
	runtime.KeepAlive(accelGroup)
	runtime.KeepAlive(accelPath)
	runtime.KeepAlive(closure)
}

// Disconnect removes an accelerator previously installed through
// gtk_accel_group_connect().
//
// Since 2.20 closure can be NULL.
//
// The function takes the following parameters:
//
//    - closure (optional) to remove from this accelerator group, or NULL to
//      remove all closures.
//
// The function returns the following values:
//
//    - ok: TRUE if the closure was found and got disconnected.
//
func (accelGroup *AccelGroup) Disconnect(closure externglib.AnyClosure) bool {
	var _arg0 *C.GtkAccelGroup // out
	var _arg1 *C.GClosure      // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))
	_arg1 = (*C.GClosure)(externglib.NewClosure(externglib.InternObject(accelGroup), closure))

	_cret = C.gtk_accel_group_disconnect(_arg0, _arg1)
	runtime.KeepAlive(accelGroup)
	runtime.KeepAlive(closure)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DisconnectKey removes an accelerator previously installed through
// gtk_accel_group_connect().
//
// The function takes the following parameters:
//
//    - accelKey: key value of the accelerator.
//    - accelMods: modifier combination of the accelerator.
//
// The function returns the following values:
//
//    - ok: TRUE if there was an accelerator which could be removed, FALSE
//      otherwise.
//
func (accelGroup *AccelGroup) DisconnectKey(accelKey uint, accelMods gdk.ModifierType) bool {
	var _arg0 *C.GtkAccelGroup  // out
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))
	_arg1 = C.guint(accelKey)
	_arg2 = C.GdkModifierType(accelMods)

	_cret = C.gtk_accel_group_disconnect_key(_arg0, _arg1, _arg2)
	runtime.KeepAlive(accelGroup)
	runtime.KeepAlive(accelKey)
	runtime.KeepAlive(accelMods)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsLocked locks are added and removed using gtk_accel_group_lock() and
// gtk_accel_group_unlock().
//
// The function returns the following values:
//
//    - ok: TRUE if there are 1 or more locks on the accel_group, FALSE
//      otherwise.
//
func (accelGroup *AccelGroup) IsLocked() bool {
	var _arg0 *C.GtkAccelGroup // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))

	_cret = C.gtk_accel_group_get_is_locked(_arg0)
	runtime.KeepAlive(accelGroup)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ModifierMask gets a ModifierType representing the mask for this accel_group.
// For example, K_CONTROL_MASK, K_SHIFT_MASK, etc.
//
// The function returns the following values:
//
//    - modifierType: modifier mask for this accel group.
//
func (accelGroup *AccelGroup) ModifierMask() gdk.ModifierType {
	var _arg0 *C.GtkAccelGroup  // out
	var _cret C.GdkModifierType // in

	_arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))

	_cret = C.gtk_accel_group_get_modifier_mask(_arg0)
	runtime.KeepAlive(accelGroup)

	var _modifierType gdk.ModifierType // out

	_modifierType = gdk.ModifierType(_cret)

	return _modifierType
}

// Lock locks the given accelerator group.
//
// Locking an acelerator group prevents the accelerators contained within it to
// be changed during runtime. Refer to gtk_accel_map_change_entry() about
// runtime accelerator changes.
//
// If called more than once, accel_group remains locked until
// gtk_accel_group_unlock() has been called an equivalent number of times.
func (accelGroup *AccelGroup) Lock() {
	var _arg0 *C.GtkAccelGroup // out

	_arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))

	C.gtk_accel_group_lock(_arg0)
	runtime.KeepAlive(accelGroup)
}

// Query queries an accelerator group for all entries matching accel_key and
// accel_mods.
//
// The function takes the following parameters:
//
//    - accelKey: key value of the accelerator.
//    - accelMods: modifier combination of the accelerator.
//
// The function returns the following values:
//
//    - accelGroupEntrys (optional): array of n_entries AccelGroupEntry elements,
//      or NULL. The array is owned by GTK+ and must not be freed.
//
func (accelGroup *AccelGroup) Query(accelKey uint, accelMods gdk.ModifierType) []AccelGroupEntry {
	var _arg0 *C.GtkAccelGroup      // out
	var _arg1 C.guint               // out
	var _arg2 C.GdkModifierType     // out
	var _cret *C.GtkAccelGroupEntry // in
	var _arg3 C.guint               // in

	_arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))
	_arg1 = C.guint(accelKey)
	_arg2 = C.GdkModifierType(accelMods)

	_cret = C.gtk_accel_group_query(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(accelGroup)
	runtime.KeepAlive(accelKey)
	runtime.KeepAlive(accelMods)

	var _accelGroupEntrys []AccelGroupEntry // out

	if _cret != nil {
		{
			src := unsafe.Slice(_cret, _arg3)
			_accelGroupEntrys = make([]AccelGroupEntry, _arg3)
			for i := 0; i < int(_arg3); i++ {
				_accelGroupEntrys[i] = *(*AccelGroupEntry)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
			}
		}
	}

	return _accelGroupEntrys
}

// Unlock undoes the last call to gtk_accel_group_lock() on this accel_group.
func (accelGroup *AccelGroup) Unlock() {
	var _arg0 *C.GtkAccelGroup // out

	_arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))

	C.gtk_accel_group_unlock(_arg0)
	runtime.KeepAlive(accelGroup)
}

// AccelGroupEntry: instance of this type is always passed by reference.
type AccelGroupEntry struct {
	*accelGroupEntry
}

// accelGroupEntry is the struct that's finalized.
type accelGroupEntry struct {
	native *C.GtkAccelGroupEntry
}

func (a *AccelGroupEntry) Key() *AccelKey {
	var v *AccelKey // out
	v = (*AccelKey)(gextras.NewStructNative(unsafe.Pointer((&a.native.key))))
	return v
}

func (a *AccelGroupEntry) AccelPathQuark() glib.Quark {
	var v glib.Quark // out
	v = uint32(a.native.accel_path_quark)
	return v
}

// AccelKey: instance of this type is always passed by reference.
type AccelKey struct {
	*accelKey
}

// accelKey is the struct that's finalized.
type accelKey struct {
	native *C.GtkAccelKey
}

// AccelKey: accelerator keyval.
func (a *AccelKey) AccelKey() uint {
	var v uint // out
	v = uint(a.native.accel_key)
	return v
}

// AccelMods: accelerator modifiers.
func (a *AccelKey) AccelMods() gdk.ModifierType {
	var v gdk.ModifierType // out
	v = gdk.ModifierType(a.native.accel_mods)
	return v
}
