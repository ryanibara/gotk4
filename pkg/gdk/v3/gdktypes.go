// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_axis_use_get_type()), F: marshalAxisUse},
		{T: externglib.Type(C.gdk_byte_order_get_type()), F: marshalByteOrder},
		{T: externglib.Type(C.gdk_gl_error_get_type()), F: marshalGLError},
		{T: externglib.Type(C.gdk_grab_ownership_get_type()), F: marshalGrabOwnership},
		{T: externglib.Type(C.gdk_grab_status_get_type()), F: marshalGrabStatus},
		{T: externglib.Type(C.gdk_modifier_intent_get_type()), F: marshalModifierIntent},
		{T: externglib.Type(C.gdk_window_type_hint_get_type()), F: marshalWindowTypeHint},
		{T: externglib.Type(C.gdk_axis_flags_get_type()), F: marshalAxisFlags},
		{T: externglib.Type(C.gdk_event_mask_get_type()), F: marshalEventMask},
		{T: externglib.Type(C.gdk_modifier_type_get_type()), F: marshalModifierType},
		{T: externglib.Type(C.gdk_rectangle_get_type()), F: marshalRectangle},
	})
}

// CURRENT_TIME represents the current time, and can be used anywhere a time is
// expected.
const CURRENT_TIME = 0

// PARENT_RELATIVE: special value, indicating that the background for a window
// should be inherited from the parent window.
const PARENT_RELATIVE = 1

// AxisUse: enumeration describing the way in which a device axis (valuator)
// maps onto the predefined valuator types that GTK+ understands.
//
// Note that the X and Y axes are not really needed; pointer devices report
// their location via the x/y members of events regardless. Whether X and Y are
// present as axes depends on the GDK backend.
type AxisUse C.gint

const (
	// AxisIgnore axis is ignored.
	AxisIgnore AxisUse = iota
	// AxisX axis is used as the x axis.
	AxisX
	// AxisY axis is used as the y axis.
	AxisY
	// AxisPressure axis is used for pressure information.
	AxisPressure
	// AxisXtilt axis is used for x tilt information.
	AxisXtilt
	// AxisYtilt axis is used for y tilt information.
	AxisYtilt
	// AxisWheel axis is used for wheel information.
	AxisWheel
	// AxisDistance axis is used for pen/tablet distance information. (Since:
	// 3.22).
	AxisDistance
	// AxisRotation axis is used for pen rotation information. (Since: 3.22).
	AxisRotation
	// AxisSlider axis is used for pen slider information. (Since: 3.22).
	AxisSlider
	// AxisLast: constant equal to the numerically highest axis value.
	AxisLast
)

func marshalAxisUse(p uintptr) (interface{}, error) {
	return AxisUse(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for AxisUse.
func (a AxisUse) String() string {
	switch a {
	case AxisIgnore:
		return "Ignore"
	case AxisX:
		return "X"
	case AxisY:
		return "Y"
	case AxisPressure:
		return "Pressure"
	case AxisXtilt:
		return "Xtilt"
	case AxisYtilt:
		return "Ytilt"
	case AxisWheel:
		return "Wheel"
	case AxisDistance:
		return "Distance"
	case AxisRotation:
		return "Rotation"
	case AxisSlider:
		return "Slider"
	case AxisLast:
		return "Last"
	default:
		return fmt.Sprintf("AxisUse(%d)", a)
	}
}

// ByteOrder: set of values describing the possible byte-orders for storing
// pixel values in memory.
type ByteOrder C.gint

const (
	// LsbFirst values are stored with the least-significant byte first. For
	// instance, the 32-bit value 0xffeecc would be stored in memory as 0xcc,
	// 0xee, 0xff, 0x00.
	LsbFirst ByteOrder = iota
	// MsbFirst values are stored with the most-significant byte first. For
	// instance, the 32-bit value 0xffeecc would be stored in memory as 0x00,
	// 0xff, 0xee, 0xcc.
	MsbFirst
)

func marshalByteOrder(p uintptr) (interface{}, error) {
	return ByteOrder(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ByteOrder.
func (b ByteOrder) String() string {
	switch b {
	case LsbFirst:
		return "LsbFirst"
	case MsbFirst:
		return "MsbFirst"
	default:
		return fmt.Sprintf("ByteOrder(%d)", b)
	}
}

// GLError: error enumeration for GLContext.
type GLError C.gint

const (
	// GLErrorNotAvailable: openGL support is not available.
	GLErrorNotAvailable GLError = iota
	// GLErrorUnsupportedFormat: requested visual format is not supported.
	GLErrorUnsupportedFormat
	// GLErrorUnsupportedProfile: requested profile is not supported.
	GLErrorUnsupportedProfile
)

func marshalGLError(p uintptr) (interface{}, error) {
	return GLError(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for GLError.
func (g GLError) String() string {
	switch g {
	case GLErrorNotAvailable:
		return "NotAvailable"
	case GLErrorUnsupportedFormat:
		return "UnsupportedFormat"
	case GLErrorUnsupportedProfile:
		return "UnsupportedProfile"
	default:
		return fmt.Sprintf("GLError(%d)", g)
	}
}

// GrabOwnership defines how device grabs interact with other devices.
type GrabOwnership C.gint

const (
	// OwnershipNone: all other devices’ events are allowed.
	OwnershipNone GrabOwnership = iota
	// OwnershipWindow: other devices’ events are blocked for the grab window.
	OwnershipWindow
	// OwnershipApplication: other devices’ events are blocked for the whole
	// application.
	OwnershipApplication
)

func marshalGrabOwnership(p uintptr) (interface{}, error) {
	return GrabOwnership(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for GrabOwnership.
func (g GrabOwnership) String() string {
	switch g {
	case OwnershipNone:
		return "None"
	case OwnershipWindow:
		return "Window"
	case OwnershipApplication:
		return "Application"
	default:
		return fmt.Sprintf("GrabOwnership(%d)", g)
	}
}

// GrabStatus: returned by gdk_device_grab(), gdk_pointer_grab() and
// gdk_keyboard_grab() to indicate success or the reason for the failure of the
// grab attempt.
type GrabStatus C.gint

const (
	// GrabSuccess: resource was successfully grabbed.
	GrabSuccess GrabStatus = iota
	// GrabAlreadyGrabbed: resource is actively grabbed by another client.
	GrabAlreadyGrabbed
	// GrabInvalidTime: resource was grabbed more recently than the specified
	// time.
	GrabInvalidTime
	// GrabNotViewable: grab window or the confine_to window are not viewable.
	GrabNotViewable
	// GrabFrozen: resource is frozen by an active grab of another client.
	GrabFrozen
	// GrabFailed: grab failed for some other reason. Since 3.16.
	GrabFailed
)

func marshalGrabStatus(p uintptr) (interface{}, error) {
	return GrabStatus(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for GrabStatus.
func (g GrabStatus) String() string {
	switch g {
	case GrabSuccess:
		return "Success"
	case GrabAlreadyGrabbed:
		return "AlreadyGrabbed"
	case GrabInvalidTime:
		return "InvalidTime"
	case GrabNotViewable:
		return "NotViewable"
	case GrabFrozen:
		return "Frozen"
	case GrabFailed:
		return "Failed"
	default:
		return fmt.Sprintf("GrabStatus(%d)", g)
	}
}

// ModifierIntent: this enum is used with gdk_keymap_get_modifier_mask() in
// order to determine what modifiers the currently used windowing system backend
// uses for particular purposes. For example, on X11/Windows, the Control key is
// used for invoking menu shortcuts (accelerators), whereas on Apple computers
// it’s the Command key (which correspond to GDK_CONTROL_MASK and GDK_MOD2_MASK,
// respectively).
type ModifierIntent C.gint

const (
	// ModifierIntentPrimaryAccelerator: primary modifier used to invoke menu
	// accelerators.
	ModifierIntentPrimaryAccelerator ModifierIntent = iota
	// ModifierIntentContextMenu: modifier used to invoke context menus. Note
	// that mouse button 3 always triggers context menus. When this modifier is
	// not 0, it additionally triggers context menus when used with mouse button
	// 1.
	ModifierIntentContextMenu
	// ModifierIntentExtendSelection: modifier used to extend selections using
	// modifier-click or modifier-cursor-key.
	ModifierIntentExtendSelection
	// ModifierIntentModifySelection: modifier used to modify selections, which
	// in most cases means toggling the clicked item into or out of the
	// selection.
	ModifierIntentModifySelection
	// ModifierIntentNoTextInput: when any of these modifiers is pressed, the
	// key event cannot produce a symbol directly. This is meant to be used for
	// input methods, and for use cases like typeahead search.
	ModifierIntentNoTextInput
	// ModifierIntentShiftGroup: modifier that switches between keyboard groups
	// (AltGr on X11/Windows and Option/Alt on OS X).
	ModifierIntentShiftGroup
	// ModifierIntentDefaultModMask: set of modifier masks accepted as modifiers
	// in accelerators. Needed because Command is mapped to MOD2 on OSX, which
	// is widely used, but on X11 MOD2 is NumLock and using that for a mod key
	// is problematic at best. Ref:
	// https://bugzilla.gnome.org/show_bug.cgi?id=736125.
	ModifierIntentDefaultModMask
)

func marshalModifierIntent(p uintptr) (interface{}, error) {
	return ModifierIntent(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ModifierIntent.
func (m ModifierIntent) String() string {
	switch m {
	case ModifierIntentPrimaryAccelerator:
		return "PrimaryAccelerator"
	case ModifierIntentContextMenu:
		return "ContextMenu"
	case ModifierIntentExtendSelection:
		return "ExtendSelection"
	case ModifierIntentModifySelection:
		return "ModifySelection"
	case ModifierIntentNoTextInput:
		return "NoTextInput"
	case ModifierIntentShiftGroup:
		return "ShiftGroup"
	case ModifierIntentDefaultModMask:
		return "DefaultModMask"
	default:
		return fmt.Sprintf("ModifierIntent(%d)", m)
	}
}

// WindowTypeHint: these are hints for the window manager that indicate what
// type of function the window has. The window manager can use this when
// determining decoration and behaviour of the window. The hint must be set
// before mapping the window.
//
// See the Extended Window Manager Hints
// (http://www.freedesktop.org/Standards/wm-spec) specification for more details
// about window types.
type WindowTypeHint C.gint

const (
	// WindowTypeHintNormal: normal toplevel window.
	WindowTypeHintNormal WindowTypeHint = iota
	// WindowTypeHintDialog: dialog window.
	WindowTypeHintDialog
	// WindowTypeHintMenu: window used to implement a menu; GTK+ uses this hint
	// only for torn-off menus, see TearoffMenuItem.
	WindowTypeHintMenu
	// WindowTypeHintToolbar: window used to implement toolbars.
	WindowTypeHintToolbar
	// WindowTypeHintSplashscreen: window used to display a splash screen during
	// application startup.
	WindowTypeHintSplashscreen
	// WindowTypeHintUtility: utility windows which are not detached toolbars or
	// dialogs.
	WindowTypeHintUtility
	// WindowTypeHintDock: used for creating dock or panel windows.
	WindowTypeHintDock
	// WindowTypeHintDesktop: used for creating the desktop background window.
	WindowTypeHintDesktop
	// WindowTypeHintDropdownMenu: menu that belongs to a menubar.
	WindowTypeHintDropdownMenu
	// WindowTypeHintPopupMenu: menu that does not belong to a menubar, e.g. a
	// context menu.
	WindowTypeHintPopupMenu
	// WindowTypeHintTooltip: tooltip.
	WindowTypeHintTooltip
	// WindowTypeHintNotification: notification - typically a “bubble” that
	// belongs to a status icon.
	WindowTypeHintNotification
	// WindowTypeHintCombo: popup from a combo box.
	WindowTypeHintCombo
	// WindowTypeHintDND: window that is used to implement a DND cursor.
	WindowTypeHintDND
)

func marshalWindowTypeHint(p uintptr) (interface{}, error) {
	return WindowTypeHint(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for WindowTypeHint.
func (w WindowTypeHint) String() string {
	switch w {
	case WindowTypeHintNormal:
		return "Normal"
	case WindowTypeHintDialog:
		return "Dialog"
	case WindowTypeHintMenu:
		return "Menu"
	case WindowTypeHintToolbar:
		return "Toolbar"
	case WindowTypeHintSplashscreen:
		return "Splashscreen"
	case WindowTypeHintUtility:
		return "Utility"
	case WindowTypeHintDock:
		return "Dock"
	case WindowTypeHintDesktop:
		return "Desktop"
	case WindowTypeHintDropdownMenu:
		return "DropdownMenu"
	case WindowTypeHintPopupMenu:
		return "PopupMenu"
	case WindowTypeHintTooltip:
		return "Tooltip"
	case WindowTypeHintNotification:
		return "Notification"
	case WindowTypeHintCombo:
		return "Combo"
	case WindowTypeHintDND:
		return "DND"
	default:
		return fmt.Sprintf("WindowTypeHint(%d)", w)
	}
}

// AxisFlags flags describing the current capabilities of a device/tool.
type AxisFlags C.guint

const (
	// AxisFlagX: x axis is present.
	AxisFlagX AxisFlags = 0b10
	// AxisFlagY: y axis is present.
	AxisFlagY AxisFlags = 0b100
	// AxisFlagPressure: pressure axis is present.
	AxisFlagPressure AxisFlags = 0b1000
	// AxisFlagXtilt: x tilt axis is present.
	AxisFlagXtilt AxisFlags = 0b10000
	// AxisFlagYtilt: y tilt axis is present.
	AxisFlagYtilt AxisFlags = 0b100000
	// AxisFlagWheel: wheel axis is present.
	AxisFlagWheel AxisFlags = 0b1000000
	// AxisFlagDistance: distance axis is present.
	AxisFlagDistance AxisFlags = 0b10000000
	// AxisFlagRotation z-axis rotation is present.
	AxisFlagRotation AxisFlags = 0b100000000
	// AxisFlagSlider: slider axis is present.
	AxisFlagSlider AxisFlags = 0b1000000000
)

func marshalAxisFlags(p uintptr) (interface{}, error) {
	return AxisFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for AxisFlags.
func (a AxisFlags) String() string {
	if a == 0 {
		return "AxisFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(127)

	for a != 0 {
		next := a & (a - 1)
		bit := a - next

		switch bit {
		case AxisFlagX:
			builder.WriteString("X|")
		case AxisFlagY:
			builder.WriteString("Y|")
		case AxisFlagPressure:
			builder.WriteString("Pressure|")
		case AxisFlagXtilt:
			builder.WriteString("Xtilt|")
		case AxisFlagYtilt:
			builder.WriteString("Ytilt|")
		case AxisFlagWheel:
			builder.WriteString("Wheel|")
		case AxisFlagDistance:
			builder.WriteString("Distance|")
		case AxisFlagRotation:
			builder.WriteString("Rotation|")
		case AxisFlagSlider:
			builder.WriteString("Slider|")
		default:
			builder.WriteString(fmt.Sprintf("AxisFlags(0b%b)|", bit))
		}

		a = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if a contains other.
func (a AxisFlags) Has(other AxisFlags) bool {
	return (a & other) == other
}

// EventMask: set of bit-flags to indicate which events a window is to receive.
// Most of these masks map onto one or more of the EventType event types above.
//
// See the [input handling overview][chap-input-handling] for details of [event
// masks][event-masks] and [event propagation][event-propagation].
//
// GDK_POINTER_MOTION_HINT_MASK is deprecated. It is a special mask to reduce
// the number of GDK_MOTION_NOTIFY events received. When using
// GDK_POINTER_MOTION_HINT_MASK, fewer GDK_MOTION_NOTIFY events will be sent,
// some of which are marked as a hint (the is_hint member is TRUE). To receive
// more motion events after a motion hint event, the application needs to asks
// for more, by calling gdk_event_request_motions().
//
// Since GTK 3.8, motion events are already compressed by default, independent
// of this mechanism. This compression can be disabled with
// gdk_window_set_event_compression(). See the documentation of that function
// for details.
//
// If GDK_TOUCH_MASK is enabled, the window will receive touch events from
// touch-enabled devices. Those will come as sequences of EventTouch with type
// GDK_TOUCH_UPDATE, enclosed by two events with type GDK_TOUCH_BEGIN and
// GDK_TOUCH_END (or GDK_TOUCH_CANCEL). gdk_event_get_event_sequence() returns
// the event sequence for these events, so different sequences may be
// distinguished.
type EventMask C.guint

const (
	// ExposureMask: receive expose events.
	ExposureMask EventMask = 0b10
	// PointerMotionMask: receive all pointer motion events.
	PointerMotionMask EventMask = 0b100
	// PointerMotionHintMask: deprecated. see the explanation above.
	PointerMotionHintMask EventMask = 0b1000
	// ButtonMotionMask: receive pointer motion events while any button is
	// pressed.
	ButtonMotionMask EventMask = 0b10000
	// Button1MotionMask: receive pointer motion events while 1 button is
	// pressed.
	Button1MotionMask EventMask = 0b100000
	// Button2MotionMask: receive pointer motion events while 2 button is
	// pressed.
	Button2MotionMask EventMask = 0b1000000
	// Button3MotionMask: receive pointer motion events while 3 button is
	// pressed.
	Button3MotionMask EventMask = 0b10000000
	// ButtonPressMask: receive button press events.
	ButtonPressMask EventMask = 0b100000000
	// ButtonReleaseMask: receive button release events.
	ButtonReleaseMask EventMask = 0b1000000000
	// KeyPressMask: receive key press events.
	KeyPressMask EventMask = 0b10000000000
	// KeyReleaseMask: receive key release events.
	KeyReleaseMask EventMask = 0b100000000000
	// EnterNotifyMask: receive window enter events.
	EnterNotifyMask EventMask = 0b1000000000000
	// LeaveNotifyMask: receive window leave events.
	LeaveNotifyMask EventMask = 0b10000000000000
	// FocusChangeMask: receive focus change events.
	FocusChangeMask EventMask = 0b100000000000000
	// StructureMask: receive events about window configuration change.
	StructureMask EventMask = 0b1000000000000000
	// PropertyChangeMask: receive property change events.
	PropertyChangeMask EventMask = 0b10000000000000000
	// VisibilityNotifyMask: receive visibility change events.
	VisibilityNotifyMask EventMask = 0b100000000000000000
	// ProximityInMask: receive proximity in events.
	ProximityInMask EventMask = 0b1000000000000000000
	// ProximityOutMask: receive proximity out events.
	ProximityOutMask EventMask = 0b10000000000000000000
	// SubstructureMask: receive events about window configuration changes of
	// child windows.
	SubstructureMask EventMask = 0b100000000000000000000
	// ScrollMask: receive scroll events.
	ScrollMask EventMask = 0b1000000000000000000000
	// TouchMask: receive touch events. Since 3.4.
	TouchMask EventMask = 0b10000000000000000000000
	// SmoothScrollMask: receive smooth scrolling events. Since 3.4.
	SmoothScrollMask EventMask = 0b100000000000000000000000
	// TouchpadGestureMask: receive touchpad gesture events. Since 3.18.
	TouchpadGestureMask EventMask = 0b1000000000000000000000000
	// TabletPadMask: receive tablet pad events. Since 3.22.
	TabletPadMask EventMask = 0b10000000000000000000000000
	// AllEventsMask: combination of all the above event masks.
	AllEventsMask EventMask = 0b11111111111111111111111110
)

func marshalEventMask(p uintptr) (interface{}, error) {
	return EventMask(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for EventMask.
func (e EventMask) String() string {
	if e == 0 {
		return "EventMask(0)"
	}

	var builder strings.Builder
	builder.Grow(256)

	for e != 0 {
		next := e & (e - 1)
		bit := e - next

		switch bit {
		case ExposureMask:
			builder.WriteString("ExposureMask|")
		case PointerMotionMask:
			builder.WriteString("PointerMotionMask|")
		case PointerMotionHintMask:
			builder.WriteString("PointerMotionHintMask|")
		case ButtonMotionMask:
			builder.WriteString("ButtonMotionMask|")
		case Button1MotionMask:
			builder.WriteString("Button1MotionMask|")
		case Button2MotionMask:
			builder.WriteString("Button2MotionMask|")
		case Button3MotionMask:
			builder.WriteString("Button3MotionMask|")
		case ButtonPressMask:
			builder.WriteString("ButtonPressMask|")
		case ButtonReleaseMask:
			builder.WriteString("ButtonReleaseMask|")
		case KeyPressMask:
			builder.WriteString("KeyPressMask|")
		case KeyReleaseMask:
			builder.WriteString("KeyReleaseMask|")
		case EnterNotifyMask:
			builder.WriteString("EnterNotifyMask|")
		case LeaveNotifyMask:
			builder.WriteString("LeaveNotifyMask|")
		case FocusChangeMask:
			builder.WriteString("FocusChangeMask|")
		case StructureMask:
			builder.WriteString("StructureMask|")
		case PropertyChangeMask:
			builder.WriteString("PropertyChangeMask|")
		case VisibilityNotifyMask:
			builder.WriteString("VisibilityNotifyMask|")
		case ProximityInMask:
			builder.WriteString("ProximityInMask|")
		case ProximityOutMask:
			builder.WriteString("ProximityOutMask|")
		case SubstructureMask:
			builder.WriteString("SubstructureMask|")
		case ScrollMask:
			builder.WriteString("ScrollMask|")
		case TouchMask:
			builder.WriteString("TouchMask|")
		case SmoothScrollMask:
			builder.WriteString("SmoothScrollMask|")
		case TouchpadGestureMask:
			builder.WriteString("TouchpadGestureMask|")
		case TabletPadMask:
			builder.WriteString("TabletPadMask|")
		case AllEventsMask:
			builder.WriteString("AllEventsMask|")
		default:
			builder.WriteString(fmt.Sprintf("EventMask(0b%b)|", bit))
		}

		e = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if e contains other.
func (e EventMask) Has(other EventMask) bool {
	return (e & other) == other
}

// ModifierType: set of bit-flags to indicate the state of modifier keys and
// mouse buttons in various event types. Typical modifier keys are Shift,
// Control, Meta, Super, Hyper, Alt, Compose, Apple, CapsLock or ShiftLock.
//
// Like the X Window System, GDK supports 8 modifier keys and 5 mouse buttons.
//
// Since 2.10, GDK recognizes which of the Meta, Super or Hyper keys are mapped
// to Mod2 - Mod5, and indicates this by setting GDK_SUPER_MASK, GDK_HYPER_MASK
// or GDK_META_MASK in the state field of key events.
//
// Note that GDK may add internal values to events which include reserved values
// such as GDK_MODIFIER_RESERVED_13_MASK. Your code should preserve and ignore
// them. You can use GDK_MODIFIER_MASK to remove all reserved values.
//
// Also note that the GDK X backend interprets button press events for button
// 4-7 as scroll events, so GDK_BUTTON4_MASK and GDK_BUTTON5_MASK will never be
// set.
type ModifierType C.guint

const (
	// ShiftMask: shift key.
	ShiftMask ModifierType = 0b1
	// LockMask: lock key (depending on the modifier mapping of the X server
	// this may either be CapsLock or ShiftLock).
	LockMask ModifierType = 0b10
	// ControlMask: control key.
	ControlMask ModifierType = 0b100
	// Mod1Mask: fourth modifier key (it depends on the modifier mapping of the
	// X server which key is interpreted as this modifier, but normally it is
	// the Alt key).
	Mod1Mask ModifierType = 0b1000
	// Mod2Mask: fifth modifier key (it depends on the modifier mapping of the X
	// server which key is interpreted as this modifier).
	Mod2Mask ModifierType = 0b10000
	// Mod3Mask: sixth modifier key (it depends on the modifier mapping of the X
	// server which key is interpreted as this modifier).
	Mod3Mask ModifierType = 0b100000
	// Mod4Mask: seventh modifier key (it depends on the modifier mapping of the
	// X server which key is interpreted as this modifier).
	Mod4Mask ModifierType = 0b1000000
	// Mod5Mask: eighth modifier key (it depends on the modifier mapping of the
	// X server which key is interpreted as this modifier).
	Mod5Mask ModifierType = 0b10000000
	// Button1Mask: first mouse button.
	Button1Mask ModifierType = 0b100000000
	// Button2Mask: second mouse button.
	Button2Mask ModifierType = 0b1000000000
	// Button3Mask: third mouse button.
	Button3Mask ModifierType = 0b10000000000
	// Button4Mask: fourth mouse button.
	Button4Mask ModifierType = 0b100000000000
	// Button5Mask: fifth mouse button.
	Button5Mask ModifierType = 0b1000000000000
	// ModifierReserved13Mask: reserved bit flag; do not use in your own code.
	ModifierReserved13Mask ModifierType = 0b10000000000000
	// ModifierReserved14Mask: reserved bit flag; do not use in your own code.
	ModifierReserved14Mask ModifierType = 0b100000000000000
	// ModifierReserved15Mask: reserved bit flag; do not use in your own code.
	ModifierReserved15Mask ModifierType = 0b1000000000000000
	// ModifierReserved16Mask: reserved bit flag; do not use in your own code.
	ModifierReserved16Mask ModifierType = 0b10000000000000000
	// ModifierReserved17Mask: reserved bit flag; do not use in your own code.
	ModifierReserved17Mask ModifierType = 0b100000000000000000
	// ModifierReserved18Mask: reserved bit flag; do not use in your own code.
	ModifierReserved18Mask ModifierType = 0b1000000000000000000
	// ModifierReserved19Mask: reserved bit flag; do not use in your own code.
	ModifierReserved19Mask ModifierType = 0b10000000000000000000
	// ModifierReserved20Mask: reserved bit flag; do not use in your own code.
	ModifierReserved20Mask ModifierType = 0b100000000000000000000
	// ModifierReserved21Mask: reserved bit flag; do not use in your own code.
	ModifierReserved21Mask ModifierType = 0b1000000000000000000000
	// ModifierReserved22Mask: reserved bit flag; do not use in your own code.
	ModifierReserved22Mask ModifierType = 0b10000000000000000000000
	// ModifierReserved23Mask: reserved bit flag; do not use in your own code.
	ModifierReserved23Mask ModifierType = 0b100000000000000000000000
	// ModifierReserved24Mask: reserved bit flag; do not use in your own code.
	ModifierReserved24Mask ModifierType = 0b1000000000000000000000000
	// ModifierReserved25Mask: reserved bit flag; do not use in your own code.
	ModifierReserved25Mask ModifierType = 0b10000000000000000000000000
	// SuperMask: super modifier. Since 2.10.
	SuperMask ModifierType = 0b100000000000000000000000000
	// HyperMask: hyper modifier. Since 2.10.
	HyperMask ModifierType = 0b1000000000000000000000000000
	// MetaMask: meta modifier. Since 2.10.
	MetaMask ModifierType = 0b10000000000000000000000000000
	// ModifierReserved29Mask: reserved bit flag; do not use in your own code.
	ModifierReserved29Mask ModifierType = 0b100000000000000000000000000000
	// ReleaseMask: not used in GDK itself. GTK+ uses it to differentiate
	// between (keyval, modifiers) pairs from key press and release events.
	ReleaseMask ModifierType = 0b1000000000000000000000000000000
	// ModifierMask: mask covering all modifier types.
	ModifierMask ModifierType = 0b1011100000000000001111111111111
)

func marshalModifierType(p uintptr) (interface{}, error) {
	return ModifierType(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for ModifierType.
func (m ModifierType) String() string {
	if m == 0 {
		return "ModifierType(0)"
	}

	var builder strings.Builder
	builder.Grow(256)

	for m != 0 {
		next := m & (m - 1)
		bit := m - next

		switch bit {
		case ShiftMask:
			builder.WriteString("ShiftMask|")
		case LockMask:
			builder.WriteString("LockMask|")
		case ControlMask:
			builder.WriteString("ControlMask|")
		case Mod1Mask:
			builder.WriteString("Mod1Mask|")
		case Mod2Mask:
			builder.WriteString("Mod2Mask|")
		case Mod3Mask:
			builder.WriteString("Mod3Mask|")
		case Mod4Mask:
			builder.WriteString("Mod4Mask|")
		case Mod5Mask:
			builder.WriteString("Mod5Mask|")
		case Button1Mask:
			builder.WriteString("Button1Mask|")
		case Button2Mask:
			builder.WriteString("Button2Mask|")
		case Button3Mask:
			builder.WriteString("Button3Mask|")
		case Button4Mask:
			builder.WriteString("Button4Mask|")
		case Button5Mask:
			builder.WriteString("Button5Mask|")
		case ModifierReserved13Mask:
			builder.WriteString("ModifierReserved13Mask|")
		case ModifierReserved14Mask:
			builder.WriteString("ModifierReserved14Mask|")
		case ModifierReserved15Mask:
			builder.WriteString("ModifierReserved15Mask|")
		case ModifierReserved16Mask:
			builder.WriteString("ModifierReserved16Mask|")
		case ModifierReserved17Mask:
			builder.WriteString("ModifierReserved17Mask|")
		case ModifierReserved18Mask:
			builder.WriteString("ModifierReserved18Mask|")
		case ModifierReserved19Mask:
			builder.WriteString("ModifierReserved19Mask|")
		case ModifierReserved20Mask:
			builder.WriteString("ModifierReserved20Mask|")
		case ModifierReserved21Mask:
			builder.WriteString("ModifierReserved21Mask|")
		case ModifierReserved22Mask:
			builder.WriteString("ModifierReserved22Mask|")
		case ModifierReserved23Mask:
			builder.WriteString("ModifierReserved23Mask|")
		case ModifierReserved24Mask:
			builder.WriteString("ModifierReserved24Mask|")
		case ModifierReserved25Mask:
			builder.WriteString("ModifierReserved25Mask|")
		case SuperMask:
			builder.WriteString("SuperMask|")
		case HyperMask:
			builder.WriteString("HyperMask|")
		case MetaMask:
			builder.WriteString("MetaMask|")
		case ModifierReserved29Mask:
			builder.WriteString("ModifierReserved29Mask|")
		case ReleaseMask:
			builder.WriteString("ReleaseMask|")
		case ModifierMask:
			builder.WriteString("ModifierMask|")
		default:
			builder.WriteString(fmt.Sprintf("ModifierType(0b%b)|", bit))
		}

		m = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if m contains other.
func (m ModifierType) Has(other ModifierType) bool {
	return (m & other) == other
}

// Point defines the x and y coordinates of a point.
//
// An instance of this type is always passed by reference.
type Point struct {
	*point
}

// point is the struct that's finalized.
type point struct {
	native *C.GdkPoint
}

// NewPoint creates a new Point instance from the given
// fields.
func NewPoint(x, y int) Point {
	var f0 C.gint // out
	f0 = C.gint(x)
	var f1 C.gint // out
	f1 = C.gint(y)

	v := C.GdkPoint{
		x: f0,
		y: f1,
	}

	return *(*Point)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

// X: x coordinate of the point.
func (p *Point) X() int {
	var v int // out
	v = int(p.native.x)
	return v
}

// Y: y coordinate of the point.
func (p *Point) Y() int {
	var v int // out
	v = int(p.native.y)
	return v
}

// Rectangle defines the position and size of a rectangle. It is identical to
// #cairo_rectangle_int_t.
//
// An instance of this type is always passed by reference.
type Rectangle struct {
	*rectangle
}

// rectangle is the struct that's finalized.
type rectangle struct {
	native *C.GdkRectangle
}

func marshalRectangle(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Rectangle{&rectangle{(*C.GdkRectangle)(b)}}, nil
}

// NewRectangle creates a new Rectangle instance from the given
// fields.
func NewRectangle(x, y, width, height int) Rectangle {
	var f0 C.int // out
	f0 = C.int(x)
	var f1 C.int // out
	f1 = C.int(y)
	var f2 C.int // out
	f2 = C.int(width)
	var f3 C.int // out
	f3 = C.int(height)

	v := C.GdkRectangle{
		x:      f0,
		y:      f1,
		width:  f2,
		height: f3,
	}

	return *(*Rectangle)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

func (r *Rectangle) X() int {
	var v int // out
	v = int(r.native.x)
	return v
}

func (r *Rectangle) Y() int {
	var v int // out
	v = int(r.native.y)
	return v
}

func (r *Rectangle) Width() int {
	var v int // out
	v = int(r.native.width)
	return v
}

func (r *Rectangle) Height() int {
	var v int // out
	v = int(r.native.height)
	return v
}

// Equal checks if the two given rectangles are equal.
//
// The function takes the following parameters:
//
//    - rect2: Rectangle.
//
// The function returns the following values:
//
//    - ok: TRUE if the rectangles are equal.
//
func (rect1 *Rectangle) Equal(rect2 *Rectangle) bool {
	var _arg0 *C.GdkRectangle // out
	var _arg1 *C.GdkRectangle // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(rect1)))
	_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(rect2)))

	_cret = C.gdk_rectangle_equal(_arg0, _arg1)
	runtime.KeepAlive(rect1)
	runtime.KeepAlive(rect2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Intersect calculates the intersection of two rectangles. It is allowed for
// dest to be the same as either src1 or src2. If the rectangles do not
// intersect, dest’s width and height is set to 0 and its x and y values are
// undefined. If you are only interested in whether the rectangles intersect,
// but not in the intersecting area itself, pass NULL for dest.
//
// The function takes the following parameters:
//
//    - src2: Rectangle.
//
// The function returns the following values:
//
//    - dest (optional): return location for the intersection of src1 and src2,
//      or NULL.
//    - ok: TRUE if the rectangles intersect.
//
func (src1 *Rectangle) Intersect(src2 *Rectangle) (*Rectangle, bool) {
	var _arg0 *C.GdkRectangle // out
	var _arg1 *C.GdkRectangle // out
	var _arg2 C.GdkRectangle  // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(src1)))
	_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(src2)))

	_cret = C.gdk_rectangle_intersect(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(src1)
	runtime.KeepAlive(src2)

	var _dest *Rectangle // out
	var _ok bool         // out

	_dest = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret != 0 {
		_ok = true
	}

	return _dest, _ok
}

// Union calculates the union of two rectangles. The union of rectangles src1
// and src2 is the smallest rectangle which includes both src1 and src2 within
// it. It is allowed for dest to be the same as either src1 or src2.
//
// Note that this function does not ignore 'empty' rectangles (ie. with zero
// width or height).
//
// The function takes the following parameters:
//
//    - src2: Rectangle.
//
// The function returns the following values:
//
//    - dest: return location for the union of src1 and src2.
//
func (src1 *Rectangle) Union(src2 *Rectangle) *Rectangle {
	var _arg0 *C.GdkRectangle // out
	var _arg1 *C.GdkRectangle // out
	var _arg2 C.GdkRectangle  // in

	_arg0 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(src1)))
	_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(src2)))

	C.gdk_rectangle_union(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(src1)
	runtime.KeepAlive(src2)

	var _dest *Rectangle // out

	_dest = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _dest
}
