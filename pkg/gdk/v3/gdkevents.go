// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_crossing_mode_get_type()), F: marshalCrossingMode},
		{T: externglib.Type(C.gdk_event_type_get_type()), F: marshalEventType},
		{T: externglib.Type(C.gdk_filter_return_get_type()), F: marshalFilterReturn},
		{T: externglib.Type(C.gdk_notify_type_get_type()), F: marshalNotifyType},
		{T: externglib.Type(C.gdk_owner_change_get_type()), F: marshalOwnerChange},
		{T: externglib.Type(C.gdk_property_state_get_type()), F: marshalPropertyState},
		{T: externglib.Type(C.gdk_scroll_direction_get_type()), F: marshalScrollDirection},
		{T: externglib.Type(C.gdk_setting_action_get_type()), F: marshalSettingAction},
		{T: externglib.Type(C.gdk_touchpad_gesture_phase_get_type()), F: marshalTouchpadGesturePhase},
		{T: externglib.Type(C.gdk_visibility_state_get_type()), F: marshalVisibilityState},
		{T: externglib.Type(C.gdk_window_state_get_type()), F: marshalWindowState},
		{T: externglib.Type(C.gdk_event_sequence_get_type()), F: marshalEventSequence},
	})
}

// XEvent: used to represent native events (XEvents for the X11 backend, MSGs
// for Win32).
type XEvent = C.void

// CrossingMode specifies the crossing mode for EventCrossing.
type CrossingMode int

const (
	// normal: crossing because of pointer motion.
	CrossingModeNormal CrossingMode = 0
	// grab: crossing because a grab is activated.
	CrossingModeGrab CrossingMode = 1
	// ungrab: crossing because a grab is deactivated.
	CrossingModeUngrab CrossingMode = 2
	// GTKGrab: crossing because a GTK+ grab is activated.
	CrossingModeGTKGrab CrossingMode = 3
	// GTKUngrab: crossing because a GTK+ grab is deactivated.
	CrossingModeGTKUngrab CrossingMode = 4
	// StateChanged: crossing because a GTK+ widget changed state (e.g.
	// sensitivity).
	CrossingModeStateChanged CrossingMode = 5
	// TouchBegin: crossing because a touch sequence has begun, this event is
	// synthetic as the pointer might have not left the window.
	CrossingModeTouchBegin CrossingMode = 6
	// TouchEnd: crossing because a touch sequence has ended, this event is
	// synthetic as the pointer might have not left the window.
	CrossingModeTouchEnd CrossingMode = 7
	// DeviceSwitch: crossing because of a device switch (i.e. a mouse taking
	// control of the pointer after a touch device), this event is synthetic as
	// the pointer didn’t leave the window.
	CrossingModeDeviceSwitch CrossingMode = 8
)

func marshalCrossingMode(p uintptr) (interface{}, error) {
	return CrossingMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// EventType specifies the type of the event.
//
// Do not confuse these events with the signals that GTK+ widgets emit. Although
// many of these events result in corresponding signals being emitted, the
// events are often transformed or filtered along the way.
//
// In some language bindings, the values GDK_2BUTTON_PRESS and GDK_3BUTTON_PRESS
// would translate into something syntactically invalid (eg
// `Gdk.EventType.2ButtonPress`, where a symbol is not allowed to start with a
// number). In that case, the aliases GDK_DOUBLE_BUTTON_PRESS and
// GDK_TRIPLE_BUTTON_PRESS can be used instead.
type EventType int

const (
	// nothing: special code to indicate a null event.
	EventTypeNothing EventType = -1
	// delete: the window manager has requested that the toplevel window be
	// hidden or destroyed, usually when the user clicks on a special icon in
	// the title bar.
	EventTypeDelete EventType = 0
	// destroy: the window has been destroyed.
	EventTypeDestroy EventType = 1
	// expose: all or part of the window has become visible and needs to be
	// redrawn.
	EventTypeExpose EventType = 2
	// MotionNotify: the pointer (usually a mouse) has moved.
	EventTypeMotionNotify EventType = 3
	// ButtonPress: mouse button has been pressed.
	EventTypeButtonPress EventType = 4
	// 2ButtonPress: mouse button has been double-clicked (clicked twice within
	// a short period of time). Note that each click also generates a
	// GDK_BUTTON_PRESS event.
	EventType2ButtonPress EventType = 5
	// DoubleButtonPress alias for GDK_2BUTTON_PRESS, added in 3.6.
	EventTypeDoubleButtonPress EventType = 5
	// 3ButtonPress: mouse button has been clicked 3 times in a short period of
	// time. Note that each click also generates a GDK_BUTTON_PRESS event.
	EventType3ButtonPress EventType = 6
	// TripleButtonPress alias for GDK_3BUTTON_PRESS, added in 3.6.
	EventTypeTripleButtonPress EventType = 6
	// ButtonRelease: mouse button has been released.
	EventTypeButtonRelease EventType = 7
	// KeyPress: key has been pressed.
	EventTypeKeyPress EventType = 8
	// KeyRelease: key has been released.
	EventTypeKeyRelease EventType = 9
	// EnterNotify: the pointer has entered the window.
	EventTypeEnterNotify EventType = 10
	// LeaveNotify: the pointer has left the window.
	EventTypeLeaveNotify EventType = 11
	// FocusChange: the keyboard focus has entered or left the window.
	EventTypeFocusChange EventType = 12
	// configure: the size, position or stacking order of the window has
	// changed. Note that GTK+ discards these events for GDK_WINDOW_CHILD
	// windows.
	EventTypeConfigure EventType = 13
	// map: the window has been mapped.
	EventTypeMap EventType = 14
	// unmap: the window has been unmapped.
	EventTypeUnmap EventType = 15
	// PropertyNotify: property on the window has been changed or deleted.
	EventTypePropertyNotify EventType = 16
	// SelectionClear: the application has lost ownership of a selection.
	EventTypeSelectionClear EventType = 17
	// SelectionRequest: another application has requested a selection.
	EventTypeSelectionRequest EventType = 18
	// SelectionNotify: selection has been received.
	EventTypeSelectionNotify EventType = 19
	// ProximityIn: input device has moved into contact with a sensing surface
	// (e.g. a touchscreen or graphics tablet).
	EventTypeProximityIn EventType = 20
	// ProximityOut: input device has moved out of contact with a sensing
	// surface.
	EventTypeProximityOut EventType = 21
	// DragEnter: the mouse has entered the window while a drag is in progress.
	EventTypeDragEnter EventType = 22
	// DragLeave: the mouse has left the window while a drag is in progress.
	EventTypeDragLeave EventType = 23
	// DragMotion: the mouse has moved in the window while a drag is in
	// progress.
	EventTypeDragMotion EventType = 24
	// DragStatus: the status of the drag operation initiated by the window has
	// changed.
	EventTypeDragStatus EventType = 25
	// DropStart: drop operation onto the window has started.
	EventTypeDropStart EventType = 26
	// DropFinished: the drop operation initiated by the window has completed.
	EventTypeDropFinished EventType = 27
	// ClientEvent: message has been received from another application.
	EventTypeClientEvent EventType = 28
	// VisibilityNotify: the window visibility status has changed.
	EventTypeVisibilityNotify EventType = 29
	// scroll: the scroll wheel was turned
	EventTypeScroll EventType = 31
	// WindowState: the state of a window has changed. See WindowState for the
	// possible window states
	EventTypeWindowState EventType = 32
	// setting has been modified.
	EventTypeSetting EventType = 33
	// OwnerChange: the owner of a selection has changed. This event type was
	// added in 2.6
	EventTypeOwnerChange EventType = 34
	// GrabBroken: pointer or keyboard grab was broken. This event type was
	// added in 2.8.
	EventTypeGrabBroken EventType = 35
	// damage: the content of the window has been changed. This event type was
	// added in 2.14.
	EventTypeDamage EventType = 36
	// TouchBegin: new touch event sequence has just started. This event type
	// was added in 3.4.
	EventTypeTouchBegin EventType = 37
	// TouchUpdate: touch event sequence has been updated. This event type was
	// added in 3.4.
	EventTypeTouchUpdate EventType = 38
	// TouchEnd: touch event sequence has finished. This event type was added in
	// 3.4.
	EventTypeTouchEnd EventType = 39
	// TouchCancel: touch event sequence has been canceled. This event type was
	// added in 3.4.
	EventTypeTouchCancel EventType = 40
	// TouchpadSwipe: touchpad swipe gesture event, the current state is
	// determined by its phase field. This event type was added in 3.18.
	EventTypeTouchpadSwipe EventType = 41
	// TouchpadPinch: touchpad pinch gesture event, the current state is
	// determined by its phase field. This event type was added in 3.18.
	EventTypeTouchpadPinch EventType = 42
	// PadButtonPress: tablet pad button press event. This event type was added
	// in 3.22.
	EventTypePadButtonPress EventType = 43
	// PadButtonRelease: tablet pad button release event. This event type was
	// added in 3.22.
	EventTypePadButtonRelease EventType = 44
	// PadRing: tablet pad axis event from a "ring". This event type was added
	// in 3.22.
	EventTypePadRing EventType = 45
	// PadStrip: tablet pad axis event from a "strip". This event type was added
	// in 3.22.
	EventTypePadStrip EventType = 46
	// PadGroupMode: tablet pad group mode change. This event type was added in
	// 3.22.
	EventTypePadGroupMode EventType = 47
	// EventLast marks the end of the GdkEventType enumeration. Added in 2.18
	EventTypeEventLast EventType = 48
)

func marshalEventType(p uintptr) (interface{}, error) {
	return EventType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// FilterReturn specifies the result of applying a FilterFunc to a native event.
type FilterReturn int

const (
	// continue: event not handled, continue processing.
	FilterReturnContinue FilterReturn = 0
	// translate: native event translated into a GDK event and stored in the
	// `event` structure that was passed in.
	FilterReturnTranslate FilterReturn = 1
	// remove: event handled, terminate processing.
	FilterReturnRemove FilterReturn = 2
)

func marshalFilterReturn(p uintptr) (interface{}, error) {
	return FilterReturn(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// NotifyType specifies the kind of crossing for EventCrossing.
//
// See the X11 protocol specification of LeaveNotify for full details of
// crossing event generation.
type NotifyType int

const (
	// ancestor: the window is entered from an ancestor or left towards an
	// ancestor.
	NotifyTypeAncestor NotifyType = 0
	// virtual: the pointer moves between an ancestor and an inferior of the
	// window.
	NotifyTypeVirtual NotifyType = 1
	// inferior: the window is entered from an inferior or left towards an
	// inferior.
	NotifyTypeInferior NotifyType = 2
	// nonlinear: the window is entered from or left towards a window which is
	// neither an ancestor nor an inferior.
	NotifyTypeNonlinear NotifyType = 3
	// NonlinearVirtual: the pointer moves between two windows which are not
	// ancestors of each other and the window is part of the ancestor chain
	// between one of these windows and their least common ancestor.
	NotifyTypeNonlinearVirtual NotifyType = 4
	// unknown type of enter/leave event occurred.
	NotifyTypeUnknown NotifyType = 5
)

func marshalNotifyType(p uintptr) (interface{}, error) {
	return NotifyType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// OwnerChange specifies why a selection ownership was changed.
type OwnerChange int

const (
	// NewOwner: some other app claimed the ownership
	OwnerChangeNewOwner OwnerChange = 0
	// destroy: the window was destroyed
	OwnerChangeDestroy OwnerChange = 1
	// close: the client was closed
	OwnerChangeClose OwnerChange = 2
)

func marshalOwnerChange(p uintptr) (interface{}, error) {
	return OwnerChange(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PropertyState specifies the type of a property change for a EventProperty.
type PropertyState int

const (
	// NewValue: the property value was changed.
	PropertyStateNewValue PropertyState = 0
	// delete: the property was deleted.
	PropertyStateDelete PropertyState = 1
)

func marshalPropertyState(p uintptr) (interface{}, error) {
	return PropertyState(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ScrollDirection specifies the direction for EventScroll.
type ScrollDirection int

const (
	// up: the window is scrolled up.
	ScrollDirectionUp ScrollDirection = 0
	// down: the window is scrolled down.
	ScrollDirectionDown ScrollDirection = 1
	// left: the window is scrolled to the left.
	ScrollDirectionLeft ScrollDirection = 2
	// right: the window is scrolled to the right.
	ScrollDirectionRight ScrollDirection = 3
	// smooth: the scrolling is determined by the delta values in EventScroll.
	// See gdk_event_get_scroll_deltas(). Since: 3.4
	ScrollDirectionSmooth ScrollDirection = 4
)

func marshalScrollDirection(p uintptr) (interface{}, error) {
	return ScrollDirection(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// SettingAction specifies the kind of modification applied to a setting in a
// EventSetting.
type SettingAction int

const (
	// new: setting was added.
	SettingActionNew SettingAction = 0
	// changed: setting was changed.
	SettingActionChanged SettingAction = 1
	// deleted: setting was deleted.
	SettingActionDeleted SettingAction = 2
)

func marshalSettingAction(p uintptr) (interface{}, error) {
	return SettingAction(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// TouchpadGesturePhase specifies the current state of a touchpad gesture. All
// gestures are guaranteed to begin with an event with phase
// GDK_TOUCHPAD_GESTURE_PHASE_BEGIN, followed by 0 or several events with phase
// GDK_TOUCHPAD_GESTURE_PHASE_UPDATE.
//
// A finished gesture may have 2 possible outcomes, an event with phase
// GDK_TOUCHPAD_GESTURE_PHASE_END will be emitted when the gesture is considered
// successful, this should be used as the hint to perform any permanent changes.
//
// Cancelled gestures may be so for a variety of reasons, due to hardware or the
// compositor, or due to the gesture recognition layers hinting the gesture did
// not finish resolutely (eg. a 3rd finger being added during a pinch gesture).
// In these cases, the last event will report the phase
// GDK_TOUCHPAD_GESTURE_PHASE_CANCEL, this should be used as a hint to undo any
// visible/permanent changes that were done throughout the progress of the
// gesture.
//
// See also EventTouchpadSwipe and EventTouchpadPinch.
type TouchpadGesturePhase int

const (
	// begin: the gesture has begun.
	TouchpadGesturePhaseBegin TouchpadGesturePhase = 0
	// update: the gesture has been updated.
	TouchpadGesturePhaseUpdate TouchpadGesturePhase = 1
	// end: the gesture was finished, changes should be permanently applied.
	TouchpadGesturePhaseEnd TouchpadGesturePhase = 2
	// cancel: the gesture was cancelled, all changes should be undone.
	TouchpadGesturePhaseCancel TouchpadGesturePhase = 3
)

func marshalTouchpadGesturePhase(p uintptr) (interface{}, error) {
	return TouchpadGesturePhase(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// VisibilityState specifies the visiblity status of a window for a
// EventVisibility.
type VisibilityState int

const (
	// unobscured: the window is completely visible.
	VisibilityStateUnobscured VisibilityState = 0
	// partial: the window is partially visible.
	VisibilityStatePartial VisibilityState = 1
	// FullyObscured: the window is not visible at all.
	VisibilityStateFullyObscured VisibilityState = 2
)

func marshalVisibilityState(p uintptr) (interface{}, error) {
	return VisibilityState(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// WindowState specifies the state of a toplevel window.
type WindowState int

const (
	// WindowStateWithdrawn: the window is not shown.
	WindowStateWithdrawn WindowState = 0b1
	// WindowStateIconified: the window is minimized.
	WindowStateIconified WindowState = 0b10
	// WindowStateMaximized: the window is maximized.
	WindowStateMaximized WindowState = 0b100
	// WindowStateSticky: the window is sticky.
	WindowStateSticky WindowState = 0b1000
	// WindowStateFullscreen: the window is maximized without decorations.
	WindowStateFullscreen WindowState = 0b10000
	// WindowStateAbove: the window is kept above other windows.
	WindowStateAbove WindowState = 0b100000
	// WindowStateBelow: the window is kept below other windows.
	WindowStateBelow WindowState = 0b1000000
	// WindowStateFocused: the window is presented as focused (with active
	// decorations).
	WindowStateFocused WindowState = 0b10000000
	// WindowStateTiled: the window is in a tiled state, Since 3.10. Since
	// 3.22.23, this is deprecated in favor of per-edge information.
	WindowStateTiled WindowState = 0b100000000
	// WindowStateTopTiled: whether the top edge is tiled, Since 3.22.23
	WindowStateTopTiled WindowState = 0b1000000000
	// WindowStateTopResizable: whether the top edge is resizable, Since 3.22.23
	WindowStateTopResizable WindowState = 0b10000000000
	// WindowStateRightTiled: whether the right edge is tiled, Since 3.22.23
	WindowStateRightTiled WindowState = 0b100000000000
	// WindowStateRightResizable: whether the right edge is resizable, Since
	// 3.22.23
	WindowStateRightResizable WindowState = 0b1000000000000
	// WindowStateBottomTiled: whether the bottom edge is tiled, Since 3.22.23
	WindowStateBottomTiled WindowState = 0b10000000000000
	// WindowStateBottomResizable: whether the bottom edge is resizable, Since
	// 3.22.23
	WindowStateBottomResizable WindowState = 0b100000000000000
	// WindowStateLeftTiled: whether the left edge is tiled, Since 3.22.23
	WindowStateLeftTiled WindowState = 0b1000000000000000
	// WindowStateLeftResizable: whether the left edge is resizable, Since
	// 3.22.23
	WindowStateLeftResizable WindowState = 0b10000000000000000
)

func marshalWindowState(p uintptr) (interface{}, error) {
	return WindowState(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// EventsPending checks if any events are ready to be processed for any display.
func EventsPending() bool {
	var _cret C.gboolean // in

	_cret = C.gdk_events_pending()

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GetShowEvents gets whether event debugging output is enabled.
func GetShowEvents() bool {
	var _cret C.gboolean // in

	_cret = C.gdk_get_show_events()

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetShowEvents sets whether a trace of received events is output. Note that
// GTK+ must be compiled with debugging (that is, configured using the
// `--enable-debug` option) to use this option.
func SetShowEvents(showEvents bool) {
	var _arg1 C.gboolean // out

	if showEvents {
		_arg1 = C.TRUE
	}

	C.gdk_set_show_events(_arg1)
}

// SettingGet obtains a desktop-wide setting, such as the double-click time, for
// the default screen. See gdk_screen_get_setting().
func SettingGet(name string, value externglib.Value) bool {
	var _arg1 *C.gchar   // out
	var _arg2 *C.GValue  // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	_cret = C.gdk_setting_get(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// EventAny contains the fields which are common to all event structs. Any event
// pointer can safely be cast to a pointer to a EventAny to access these fields.
type EventAny struct {
	native C.GdkEventAny
}

// WrapEventAny wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventAny(ptr unsafe.Pointer) *EventAny {
	return (*EventAny)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventAny) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventButton: used for button press and button release events. The @type field
// will be one of GDK_BUTTON_PRESS, GDK_2BUTTON_PRESS, GDK_3BUTTON_PRESS or
// GDK_BUTTON_RELEASE,
//
// Double and triple-clicks result in a sequence of events being received. For
// double-clicks the order of events will be:
//
// - GDK_BUTTON_PRESS - GDK_BUTTON_RELEASE - GDK_BUTTON_PRESS -
// GDK_2BUTTON_PRESS - GDK_BUTTON_RELEASE
//
// Note that the first click is received just like a normal button press, while
// the second click results in a GDK_2BUTTON_PRESS being received just after the
// GDK_BUTTON_PRESS.
//
// Triple-clicks are very similar to double-clicks, except that
// GDK_3BUTTON_PRESS is inserted after the third click. The order of the events
// is:
//
// - GDK_BUTTON_PRESS - GDK_BUTTON_RELEASE - GDK_BUTTON_PRESS -
// GDK_2BUTTON_PRESS - GDK_BUTTON_RELEASE - GDK_BUTTON_PRESS - GDK_3BUTTON_PRESS
// - GDK_BUTTON_RELEASE
//
// For a double click to occur, the second button press must occur within 1/4 of
// a second of the first. For a triple click to occur, the third button press
// must also occur within 1/2 second of the first button press.
type EventButton struct {
	native C.GdkEventButton
}

// WrapEventButton wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventButton(ptr unsafe.Pointer) *EventButton {
	return (*EventButton)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventButton) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventConfigure: generated when a window size or position has changed.
type EventConfigure struct {
	native C.GdkEventConfigure
}

// WrapEventConfigure wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventConfigure(ptr unsafe.Pointer) *EventConfigure {
	return (*EventConfigure)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventConfigure) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventCrossing: generated when the pointer enters or leaves a window.
type EventCrossing struct {
	native C.GdkEventCrossing
}

// WrapEventCrossing wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventCrossing(ptr unsafe.Pointer) *EventCrossing {
	return (*EventCrossing)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventCrossing) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventDND: generated during DND operations.
type EventDND struct {
	native C.GdkEventDND
}

// WrapEventDND wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventDND(ptr unsafe.Pointer) *EventDND {
	return (*EventDND)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventDND) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventExpose: generated when all or part of a window becomes visible and needs
// to be redrawn.
type EventExpose struct {
	native C.GdkEventExpose
}

// WrapEventExpose wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventExpose(ptr unsafe.Pointer) *EventExpose {
	return (*EventExpose)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventExpose) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventFocus describes a change of keyboard focus.
type EventFocus struct {
	native C.GdkEventFocus
}

// WrapEventFocus wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventFocus(ptr unsafe.Pointer) *EventFocus {
	return (*EventFocus)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventFocus) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventGrabBroken: generated when a pointer or keyboard grab is broken. On X11,
// this happens when the grab window becomes unviewable (i.e. it or one of its
// ancestors is unmapped), or if the same application grabs the pointer or
// keyboard again. Note that implicit grabs (which are initiated by button
// presses) can also cause EventGrabBroken events.
type EventGrabBroken struct {
	native C.GdkEventGrabBroken
}

// WrapEventGrabBroken wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventGrabBroken(ptr unsafe.Pointer) *EventGrabBroken {
	return (*EventGrabBroken)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventGrabBroken) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventKey describes a key press or key release event.
type EventKey struct {
	native C.GdkEventKey
}

// WrapEventKey wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventKey(ptr unsafe.Pointer) *EventKey {
	return (*EventKey)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventKey) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventMotion: generated when the pointer moves.
type EventMotion struct {
	native C.GdkEventMotion
}

// WrapEventMotion wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventMotion(ptr unsafe.Pointer) *EventMotion {
	return (*EventMotion)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventMotion) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventOwnerChange: generated when the owner of a selection changes. On X11,
// this information is only available if the X server supports the XFIXES
// extension.
type EventOwnerChange struct {
	native C.GdkEventOwnerChange
}

// WrapEventOwnerChange wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventOwnerChange(ptr unsafe.Pointer) *EventOwnerChange {
	return (*EventOwnerChange)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventOwnerChange) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventPadAxis: generated during GDK_SOURCE_TABLET_PAD interaction with tactile
// sensors.
type EventPadAxis struct {
	native C.GdkEventPadAxis
}

// WrapEventPadAxis wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventPadAxis(ptr unsafe.Pointer) *EventPadAxis {
	return (*EventPadAxis)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventPadAxis) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventPadButton: generated during GDK_SOURCE_TABLET_PAD button presses and
// releases.
type EventPadButton struct {
	native C.GdkEventPadButton
}

// WrapEventPadButton wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventPadButton(ptr unsafe.Pointer) *EventPadButton {
	return (*EventPadButton)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventPadButton) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventPadGroupMode: generated during GDK_SOURCE_TABLET_PAD mode switches in a
// group.
type EventPadGroupMode struct {
	native C.GdkEventPadGroupMode
}

// WrapEventPadGroupMode wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventPadGroupMode(ptr unsafe.Pointer) *EventPadGroupMode {
	return (*EventPadGroupMode)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventPadGroupMode) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventProperty describes a property change on a window.
type EventProperty struct {
	native C.GdkEventProperty
}

// WrapEventProperty wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventProperty(ptr unsafe.Pointer) *EventProperty {
	return (*EventProperty)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventProperty) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventProximity: proximity events are generated when using GDK’s wrapper for
// the XInput extension. The XInput extension is an add-on for standard X that
// allows you to use nonstandard devices such as graphics tablets. A proximity
// event indicates that the stylus has moved in or out of contact with the
// tablet, or perhaps that the user’s finger has moved in or out of contact with
// a touch screen.
//
// This event type will be used pretty rarely. It only is important for XInput
// aware programs that are drawing their own cursor.
type EventProximity struct {
	native C.GdkEventProximity
}

// WrapEventProximity wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventProximity(ptr unsafe.Pointer) *EventProximity {
	return (*EventProximity)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventProximity) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventScroll: generated from button presses for the buttons 4 to 7. Wheel mice
// are usually configured to generate button press events for buttons 4 and 5
// when the wheel is turned.
//
// Some GDK backends can also generate “smooth” scroll events, which can be
// recognized by the GDK_SCROLL_SMOOTH scroll direction. For these, the scroll
// deltas can be obtained with gdk_event_get_scroll_deltas().
type EventScroll struct {
	native C.GdkEventScroll
}

// WrapEventScroll wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventScroll(ptr unsafe.Pointer) *EventScroll {
	return (*EventScroll)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventScroll) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventSelection: generated when a selection is requested or ownership of a
// selection is taken over by another client application.
type EventSelection struct {
	native C.GdkEventSelection
}

// WrapEventSelection wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventSelection(ptr unsafe.Pointer) *EventSelection {
	return (*EventSelection)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventSelection) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

type EventSequence struct {
	native C.GdkEventSequence
}

// WrapEventSequence wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventSequence(ptr unsafe.Pointer) *EventSequence {
	return (*EventSequence)(ptr)
}

func marshalEventSequence(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*EventSequence)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (e *EventSequence) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventSetting: generated when a setting is modified.
type EventSetting struct {
	native C.GdkEventSetting
}

// WrapEventSetting wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventSetting(ptr unsafe.Pointer) *EventSetting {
	return (*EventSetting)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventSetting) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventTouch: used for touch events. @type field will be one of
// GDK_TOUCH_BEGIN, GDK_TOUCH_UPDATE, GDK_TOUCH_END or GDK_TOUCH_CANCEL.
//
// Touch events are grouped into sequences by means of the @sequence field,
// which can also be obtained with gdk_event_get_event_sequence(). Each sequence
// begins with a GDK_TOUCH_BEGIN event, followed by any number of
// GDK_TOUCH_UPDATE events, and ends with a GDK_TOUCH_END (or GDK_TOUCH_CANCEL)
// event. With multitouch devices, there may be several active sequences at the
// same time.
type EventTouch struct {
	native C.GdkEventTouch
}

// WrapEventTouch wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventTouch(ptr unsafe.Pointer) *EventTouch {
	return (*EventTouch)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventTouch) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventTouchpadPinch: generated during touchpad swipe gestures.
type EventTouchpadPinch struct {
	native C.GdkEventTouchpadPinch
}

// WrapEventTouchpadPinch wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventTouchpadPinch(ptr unsafe.Pointer) *EventTouchpadPinch {
	return (*EventTouchpadPinch)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventTouchpadPinch) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventTouchpadSwipe: generated during touchpad swipe gestures.
type EventTouchpadSwipe struct {
	native C.GdkEventTouchpadSwipe
}

// WrapEventTouchpadSwipe wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventTouchpadSwipe(ptr unsafe.Pointer) *EventTouchpadSwipe {
	return (*EventTouchpadSwipe)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventTouchpadSwipe) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventVisibility: generated when the window visibility status has changed.
//
// Deprecated: since version 3.12.
type EventVisibility struct {
	native C.GdkEventVisibility
}

// WrapEventVisibility wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventVisibility(ptr unsafe.Pointer) *EventVisibility {
	return (*EventVisibility)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventVisibility) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// EventWindowState: generated when the state of a toplevel window changes.
type EventWindowState struct {
	native C.GdkEventWindowState
}

// WrapEventWindowState wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventWindowState(ptr unsafe.Pointer) *EventWindowState {
	return (*EventWindowState)(ptr)
}

// Native returns the underlying C source pointer.
func (e *EventWindowState) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}
