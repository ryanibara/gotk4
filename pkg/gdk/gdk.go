// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"github.com/diamondburned/gotk4/internal/callback"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// extern void callbackDelete(gpointer);
import "C"

//export callbackDelete
func callbackDelete(ptr C.gpointer) { callback.Delete(ptr) }

type AxisUse int

const (
	// AxisUseIgnore: the axis is ignored.
	AxisUseIgnore AxisUse = 0
	// AxisUseX: the axis is used as the x axis.
	AxisUseX AxisUse = 1
	// AxisUseY: the axis is used as the y axis.
	AxisUseY AxisUse = 2
	// AxisUseDeltaX: the axis is used as the scroll x delta
	AxisUseDeltaX AxisUse = 3
	// AxisUseDeltaY: the axis is used as the scroll y delta
	AxisUseDeltaY AxisUse = 4
	// AxisUsePressure: the axis is used for pressure information.
	AxisUsePressure AxisUse = 5
	// AxisUseXTilt: the axis is used for x tilt information.
	AxisUseXTilt AxisUse = 6
	// AxisUseYTilt: the axis is used for y tilt information.
	AxisUseYTilt AxisUse = 7
	// AxisUseWheel: the axis is used for wheel information.
	AxisUseWheel AxisUse = 8
	// AxisUseDistance: the axis is used for pen/tablet distance information
	AxisUseDistance AxisUse = 9
	// AxisUseRotation: the axis is used for pen rotation information
	AxisUseRotation AxisUse = 10
	// AxisUseSlider: the axis is used for pen slider information
	AxisUseSlider AxisUse = 11
	// AxisUseLast: a constant equal to the numerically highest axis value.
	AxisUseLast AxisUse = 12
)

type CrossingMode int

const (
	// CrossingModeNormal: crossing because of pointer motion.
	CrossingModeNormal CrossingMode = 0
	// CrossingModeGrab: crossing because a grab is activated.
	CrossingModeGrab CrossingMode = 1
	// CrossingModeUngrab: crossing because a grab is deactivated.
	CrossingModeUngrab CrossingMode = 2
	// CrossingModeGTKGrab: crossing because a GTK grab is activated.
	CrossingModeGTKGrab CrossingMode = 3
	// CrossingModeGTKUngrab: crossing because a GTK grab is deactivated.
	CrossingModeGTKUngrab CrossingMode = 4
	// CrossingModeStateChanged: crossing because a GTK widget changed state
	// (e.g. sensitivity).
	CrossingModeStateChanged CrossingMode = 5
	// CrossingModeTouchBegin: crossing because a touch sequence has begun, this
	// event is synthetic as the pointer might have not left the surface.
	CrossingModeTouchBegin CrossingMode = 6
	// CrossingModeTouchEnd: crossing because a touch sequence has ended, this
	// event is synthetic as the pointer might have not left the surface.
	CrossingModeTouchEnd CrossingMode = 7
	// CrossingModeDeviceSwitch: crossing because of a device switch (i.e. a
	// mouse taking control of the pointer after a touch device), this event is
	// synthetic as the pointer didn’t leave the surface.
	CrossingModeDeviceSwitch CrossingMode = 8
)

type DevicePadFeature int

const (
	// DevicePadFeatureButton: a button
	DevicePadFeatureButton DevicePadFeature = 0
	// DevicePadFeatureRing: a ring-shaped interactive area
	DevicePadFeatureRing DevicePadFeature = 1
	// DevicePadFeatureStrip: a straight interactive area
	DevicePadFeatureStrip DevicePadFeature = 2
)

type DeviceToolType int

const (
	// DeviceToolTypeUnknown: tool is of an unknown type.
	DeviceToolTypeUnknown DeviceToolType = 0
	// DeviceToolTypePen: tool is a standard tablet stylus.
	DeviceToolTypePen DeviceToolType = 1
	// DeviceToolTypeEraser: tool is standard tablet eraser.
	DeviceToolTypeEraser DeviceToolType = 2
	// DeviceToolTypeBrush: tool is a brush stylus.
	DeviceToolTypeBrush DeviceToolType = 3
	// DeviceToolTypePencil: tool is a pencil stylus.
	DeviceToolTypePencil DeviceToolType = 4
	// DeviceToolTypeAirbrush: tool is an airbrush stylus.
	DeviceToolTypeAirbrush DeviceToolType = 5
	// DeviceToolTypeMouse: tool is a mouse.
	DeviceToolTypeMouse DeviceToolType = 6
	// DeviceToolTypeLens: tool is a lens cursor.
	DeviceToolTypeLens DeviceToolType = 7
)

type DragCancelReason int

const (
	// DragCancelReasonNoTarget: there is no suitable drop target.
	DragCancelReasonNoTarget DragCancelReason = 0
	// DragCancelReasonUserCancelled: drag cancelled by the user
	DragCancelReasonUserCancelled DragCancelReason = 1
	// DragCancelReasonError: unspecified error.
	DragCancelReasonError DragCancelReason = 2
)

type EventType int

const (
	// EventTypeDelete: the window manager has requested that the toplevel
	// surface be hidden or destroyed, usually when the user clicks on a special
	// icon in the title bar.
	EventTypeDelete EventType = 0
	// EventTypeMotionNotify: the pointer (usually a mouse) has moved.
	EventTypeMotionNotify EventType = 1
	// EventTypeButtonPress: a mouse button has been pressed.
	EventTypeButtonPress EventType = 2
	// EventTypeButtonRelease: a mouse button has been released.
	EventTypeButtonRelease EventType = 3
	// EventTypeKeyPress: a key has been pressed.
	EventTypeKeyPress EventType = 4
	// EventTypeKeyRelease: a key has been released.
	EventTypeKeyRelease EventType = 5
	// EventTypeEnterNotify: the pointer has entered the surface.
	EventTypeEnterNotify EventType = 6
	// EventTypeLeaveNotify: the pointer has left the surface.
	EventTypeLeaveNotify EventType = 7
	// EventTypeFocusChange: the keyboard focus has entered or left the surface.
	EventTypeFocusChange EventType = 8
	// EventTypeProximityIn: an input device has moved into contact with a
	// sensing surface (e.g. a touchscreen or graphics tablet).
	EventTypeProximityIn EventType = 9
	// EventTypeProximityOut: an input device has moved out of contact with a
	// sensing surface.
	EventTypeProximityOut EventType = 10
	// EventTypeDragEnter: the mouse has entered the surface while a drag is in
	// progress.
	EventTypeDragEnter EventType = 11
	// EventTypeDragLeave: the mouse has left the surface while a drag is in
	// progress.
	EventTypeDragLeave EventType = 12
	// EventTypeDragMotion: the mouse has moved in the surface while a drag is
	// in progress.
	EventTypeDragMotion EventType = 13
	// EventTypeDropStart: a drop operation onto the surface has started.
	EventTypeDropStart EventType = 14
	// EventTypeScroll: the scroll wheel was turned
	EventTypeScroll EventType = 15
	// EventTypeGrabBroken: a pointer or keyboard grab was broken.
	EventTypeGrabBroken EventType = 16
	// EventTypeTouchBegin: a new touch event sequence has just started.
	EventTypeTouchBegin EventType = 17
	// EventTypeTouchUpdate: a touch event sequence has been updated.
	EventTypeTouchUpdate EventType = 18
	// EventTypeTouchEnd: a touch event sequence has finished.
	EventTypeTouchEnd EventType = 19
	// EventTypeTouchCancel: a touch event sequence has been canceled.
	EventTypeTouchCancel EventType = 20
	// EventTypeTouchpadSwipe: a touchpad swipe gesture event, the current state
	// is determined by its phase field.
	EventTypeTouchpadSwipe EventType = 21
	// EventTypeTouchpadPinch: a touchpad pinch gesture event, the current state
	// is determined by its phase field.
	EventTypeTouchpadPinch EventType = 22
	// EventTypePadButtonPress: a tablet pad button press event.
	EventTypePadButtonPress EventType = 23
	// EventTypePadButtonRelease: a tablet pad button release event.
	EventTypePadButtonRelease EventType = 24
	// EventTypePadRing: a tablet pad axis event from a "ring".
	EventTypePadRing EventType = 25
	// EventTypePadStrip: a tablet pad axis event from a "strip".
	EventTypePadStrip EventType = 26
	// EventTypePadGroupMode: a tablet pad group mode change.
	EventTypePadGroupMode EventType = 27
	// EventTypeEventLast: marks the end of the GdkEventType enumeration.
	EventTypeEventLast EventType = 28
)

type FullscreenMode int

const (
	// FullscreenModeCurrentMonitor: fullscreen on current monitor only.
	FullscreenModeCurrentMonitor FullscreenMode = 0
	// FullscreenModeAllMonitors: span across all monitors when fullscreen.
	FullscreenModeAllMonitors FullscreenMode = 1
)

type GLError int

const (
	// GLErrorNotAvailable: openGL support is not available
	GLErrorNotAvailable GLError = 0
	// GLErrorUnsupportedFormat: the requested visual format is not supported
	GLErrorUnsupportedFormat GLError = 1
	// GLErrorUnsupportedProfile: the requested profile is not supported
	GLErrorUnsupportedProfile GLError = 2
	// GLErrorCompilationFailed: the shader compilation failed
	GLErrorCompilationFailed GLError = 3
	// GLErrorLinkFailed: the shader linking failed
	GLErrorLinkFailed GLError = 4
)

type Gravity int

const (
	// GravityNorthWest: the reference point is at the top left corner.
	GravityNorthWest Gravity = 1
	// GravityNorth: the reference point is in the middle of the top edge.
	GravityNorth Gravity = 2
	// GravityNorthEast: the reference point is at the top right corner.
	GravityNorthEast Gravity = 3
	// GravityWest: the reference point is at the middle of the left edge.
	GravityWest Gravity = 4
	// GravityCenter: the reference point is at the center of the surface.
	GravityCenter Gravity = 5
	// GravityEast: the reference point is at the middle of the right edge.
	GravityEast Gravity = 6
	// GravitySouthWest: the reference point is at the lower left corner.
	GravitySouthWest Gravity = 7
	// GravitySouth: the reference point is at the middle of the lower edge.
	GravitySouth Gravity = 8
	// GravitySouthEast: the reference point is at the lower right corner.
	GravitySouthEast Gravity = 9
	// GravityStatic: the reference point is at the top left corner of the
	// surface itself, ignoring window manager decorations.
	GravityStatic Gravity = 10
)

type InputSource int

const (
	// InputSourceMouse: the device is a mouse. (This will be reported for the
	// core pointer, even if it is something else, such as a trackball.)
	InputSourceMouse InputSource = 0
	// InputSourcePen: the device is a stylus of a graphics tablet or similar
	// device.
	InputSourcePen InputSource = 1
	// InputSourceKeyboard: the device is a keyboard.
	InputSourceKeyboard InputSource = 2
	// InputSourceTouchscreen: the device is a direct-input touch device, such
	// as a touchscreen or tablet
	InputSourceTouchscreen InputSource = 3
	// InputSourceTouchpad: the device is an indirect touch device, such as a
	// touchpad
	InputSourceTouchpad InputSource = 4
	// InputSourceTrackpoint: the device is a trackpoint
	InputSourceTrackpoint InputSource = 5
	// InputSourceTabletPad: the device is a "pad", a collection of buttons,
	// rings and strips found in drawing tablets
	InputSourceTabletPad InputSource = 6
)

type KeyMatch int

const (
	// KeyMatchNone: the key event does not match
	KeyMatchNone KeyMatch = 0
	// KeyMatchPartial: the key event matches if keyboard state (specifically,
	// the currently active group) is ignored
	KeyMatchPartial KeyMatch = 1
	// KeyMatchExact: the key event matches
	KeyMatchExact KeyMatch = 2
)

type MemoryFormat int

const (
	// MemoryFormatB8G8R8A8Premultiplied: 4 bytes; for blue, green, red, alpha.
	// The color values are premultiplied with the alpha value.
	MemoryFormatB8G8R8A8Premultiplied MemoryFormat = 0
	// MemoryFormatA8R8G8B8Premultiplied: 4 bytes; for alpha, red, green, blue.
	// The color values are premultiplied with the alpha value.
	MemoryFormatA8R8G8B8Premultiplied MemoryFormat = 1
	// MemoryFormatR8G8B8A8Premultiplied: 4 bytes; for red, green, blue, alpha
	// The color values are premultiplied with the alpha value.
	MemoryFormatR8G8B8A8Premultiplied MemoryFormat = 2
	// MemoryFormatB8G8R8A8: 4 bytes; for blue, green, red, alpha.
	MemoryFormatB8G8R8A8 MemoryFormat = 3
	// MemoryFormatA8R8G8B8: 4 bytes; for alpha, red, green, blue.
	MemoryFormatA8R8G8B8 MemoryFormat = 4
	// MemoryFormatR8G8B8A8: 4 bytes; for red, green, blue, alpha.
	MemoryFormatR8G8B8A8 MemoryFormat = 5
	// MemoryFormatA8B8G8R8: 4 bytes; for alpha, blue, green, red.
	MemoryFormatA8B8G8R8 MemoryFormat = 6
	// MemoryFormatR8G8B8: 3 bytes; for red, green, blue. The data is opaque.
	MemoryFormatR8G8B8 MemoryFormat = 7
	// MemoryFormatB8G8R8: 3 bytes; for blue, green, red. The data is opaque.
	MemoryFormatB8G8R8 MemoryFormat = 8
	// MemoryFormatNFormats: the number of formats. This value will change as
	// more formats get added, so do not rely on its concrete integer.
	MemoryFormatNFormats MemoryFormat = 9
)

type NotifyType int

const (
	// NotifyTypeAncestor: the surface is entered from an ancestor or left
	// towards an ancestor.
	NotifyTypeAncestor NotifyType = 0
	// NotifyTypeVirtual: the pointer moves between an ancestor and an inferior
	// of the surface.
	NotifyTypeVirtual NotifyType = 1
	// NotifyTypeInferior: the surface is entered from an inferior or left
	// towards an inferior.
	NotifyTypeInferior NotifyType = 2
	// NotifyTypeNonlinear: the surface is entered from or left towards a
	// surface which is neither an ancestor nor an inferior.
	NotifyTypeNonlinear NotifyType = 3
	// NotifyTypeNonlinearVirtual: the pointer moves between two surfaces which
	// are not ancestors of each other and the surface is part of the ancestor
	// chain between one of these surfaces and their least common ancestor.
	NotifyTypeNonlinearVirtual NotifyType = 4
	// NotifyTypeUnknown: an unknown type of enter/leave event occurred.
	NotifyTypeUnknown NotifyType = 5
)

type ScrollDirection int

const (
	// ScrollDirectionUp: the surface is scrolled up.
	ScrollDirectionUp ScrollDirection = 0
	// ScrollDirectionDown: the surface is scrolled down.
	ScrollDirectionDown ScrollDirection = 1
	// ScrollDirectionLeft: the surface is scrolled to the left.
	ScrollDirectionLeft ScrollDirection = 2
	// ScrollDirectionRight: the surface is scrolled to the right.
	ScrollDirectionRight ScrollDirection = 3
	// ScrollDirectionSmooth: the scrolling is determined by the delta values in
	// scroll events. See gdk_scroll_event_get_deltas()
	ScrollDirectionSmooth ScrollDirection = 4
)

type SubpixelLayout int

const (
	// SubpixelLayoutUnknown: the layout is not known
	SubpixelLayoutUnknown SubpixelLayout = 0
	// SubpixelLayoutNone: not organized in this way
	SubpixelLayoutNone SubpixelLayout = 1
	// SubpixelLayoutHorizontalRgb: the layout is horizontal, the order is RGB
	SubpixelLayoutHorizontalRgb SubpixelLayout = 2
	// SubpixelLayoutHorizontalBgr: the layout is horizontal, the order is BGR
	SubpixelLayoutHorizontalBgr SubpixelLayout = 3
	// SubpixelLayoutVerticalRgb: the layout is vertical, the order is RGB
	SubpixelLayoutVerticalRgb SubpixelLayout = 4
	// SubpixelLayoutVerticalBgr: the layout is vertical, the order is BGR
	SubpixelLayoutVerticalBgr SubpixelLayout = 5
)

type SurfaceEdge int

const (
	// SurfaceEdgeNorthWest: the top left corner.
	SurfaceEdgeNorthWest SurfaceEdge = 0
	// SurfaceEdgeNorth: the top edge.
	SurfaceEdgeNorth SurfaceEdge = 1
	// SurfaceEdgeNorthEast: the top right corner.
	SurfaceEdgeNorthEast SurfaceEdge = 2
	// SurfaceEdgeWest: the left edge.
	SurfaceEdgeWest SurfaceEdge = 3
	// SurfaceEdgeEast: the right edge.
	SurfaceEdgeEast SurfaceEdge = 4
	// SurfaceEdgeSouthWest: the lower left corner.
	SurfaceEdgeSouthWest SurfaceEdge = 5
	// SurfaceEdgeSouth: the lower edge.
	SurfaceEdgeSouth SurfaceEdge = 6
	// SurfaceEdgeSouthEast: the lower right corner.
	SurfaceEdgeSouthEast SurfaceEdge = 7
)

type TouchpadGesturePhase int

const (
	// TouchpadGesturePhaseBegin: the gesture has begun.
	TouchpadGesturePhaseBegin TouchpadGesturePhase = 0
	// TouchpadGesturePhaseUpdate: the gesture has been updated.
	TouchpadGesturePhaseUpdate TouchpadGesturePhase = 1
	// TouchpadGesturePhaseEnd: the gesture was finished, changes should be
	// permanently applied.
	TouchpadGesturePhaseEnd TouchpadGesturePhase = 2
	// TouchpadGesturePhaseCancel: the gesture was cancelled, all changes should
	// be undone.
	TouchpadGesturePhaseCancel TouchpadGesturePhase = 3
)

type VulkanError int

const (
	// VulkanErrorUnsupported: vulkan is not supported on this backend or has
	// not been compiled in.
	VulkanErrorUnsupported VulkanError = 0
	// VulkanErrorNotAvailable: vulkan support is not available on this Surface
	VulkanErrorNotAvailable VulkanError = 1
)
