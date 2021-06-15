// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_display_get_type()), F: marshalDisplay},
	})
}

// Display objects purpose are two fold:
//
// - To manage and provide information about input devices (pointers and
// keyboards)
//
// - To manage and provide information about the available Screens
//
// GdkDisplay objects are the GDK representation of an X Display, which can be
// described as a workstation consisting of a keyboard, a pointing device (such
// as a mouse) and one or more screens. It is used to open and keep track of
// various GdkScreen objects currently instantiated by the application. It is
// also used to access the keyboard(s) and mouse pointer(s) of the display.
//
// Most of the input device handling has been factored out into the separate
// DeviceManager object. Every display has a device manager, which you can
// obtain using gdk_display_get_device_manager().
type Display interface {
	gextras.Objector

	// Beep emits a short beep on @display
	Beep()
	// Close closes the connection to the windowing system for the given
	// display, and cleans up associated resources.
	Close()
	// DeviceIsGrabbed returns true if there is an ongoing grab on @device for
	// @display.
	DeviceIsGrabbed(device Device) bool
	// Flush flushes any requests queued for the windowing system; this happens
	// automatically when the main loop blocks waiting for new events, but if
	// your application is drawing without returning control to the main loop,
	// you may need to call this function explicitly. A common case where this
	// function needs to be called is when an application is executing drawing
	// commands from a thread other than the thread where the main loop is
	// running.
	//
	// This is most useful for X11. On windowing systems where requests are
	// handled synchronously, this function will do nothing.
	Flush()
	// AppLaunchContext returns a AppLaunchContext suitable for launching
	// applications on the given display.
	AppLaunchContext() AppLaunchContext
	// DefaultCursorSize returns the default size to use for cursors on
	// @display.
	DefaultCursorSize() uint
	// DefaultGroup returns the default group leader window for all toplevel
	// windows on @display. This window is implicitly created by GDK. See
	// gdk_window_set_group().
	DefaultGroup() Window
	// DefaultScreen: get the default Screen for @display.
	DefaultScreen() Screen
	// DefaultSeat returns the default Seat for this display.
	DefaultSeat() Seat
	// DeviceManager returns the DeviceManager associated to @display.
	DeviceManager() DeviceManager
	// MaximalCursorSize gets the maximal size to use for cursors on @display.
	MaximalCursorSize() (width uint, height uint)
	// Monitor gets a monitor associated with this display.
	Monitor(monitorNum int) Monitor
	// MonitorAtPoint gets the monitor in which the point (@x, @y) is located,
	// or a nearby monitor if the point is not in any monitor.
	MonitorAtPoint(x int, y int) Monitor
	// MonitorAtWindow gets the monitor in which the largest area of @window
	// resides, or a monitor close to @window if it is outside of all monitors.
	MonitorAtWindow(window Window) Monitor
	// NMonitors gets the number of monitors that belong to @display.
	//
	// The returned number is valid until the next emission of the
	// Display::monitor-added or Display::monitor-removed signal.
	NMonitors() int
	// NScreens gets the number of screen managed by the @display.
	NScreens() int
	// Name gets the name of the display.
	Name() string
	// Pointer gets the current location of the pointer and the current modifier
	// mask for a given display.
	Pointer() (screen Screen, x int, y int, mask ModifierType)
	// PrimaryMonitor gets the primary monitor for the display.
	//
	// The primary monitor is considered the monitor where the “main desktop”
	// lives. While normal application windows typically allow the window
	// manager to place the windows, specialized desktop applications such as
	// panels should place themselves on the primary monitor.
	PrimaryMonitor() Monitor
	// Screen returns a screen object for one of the screens of the display.
	Screen(screenNum int) Screen
	// WindowAtPointer obtains the window underneath the mouse pointer,
	// returning the location of the pointer in that window in @win_x, @win_y
	// for @screen. Returns nil if the window under the mouse pointer is not
	// known to GDK (for example, belongs to another application).
	WindowAtPointer() (winX int, winY int, window Window)
	// HasPending returns whether the display has events that are waiting to be
	// processed.
	HasPending() bool
	// IsClosed finds out if the display has been closed.
	IsClosed() bool
	// KeyboardUngrab: release any keyboard grab
	KeyboardUngrab(time_ uint32)
	// NotifyStartupComplete indicates to the GUI environment that the
	// application has finished loading, using a given identifier.
	//
	// GTK+ will call this function automatically for Window with custom
	// startup-notification identifier unless
	// gtk_window_set_auto_startup_notification() is called to disable that
	// feature.
	NotifyStartupComplete(startupId string)
	// PointerIsGrabbed: test if the pointer is grabbed.
	PointerIsGrabbed() bool
	// PointerUngrab: release any pointer grab.
	PointerUngrab(time_ uint32)
	// RequestSelectionNotification: request EventOwnerChange events for
	// ownership changes of the selection named by the given atom.
	RequestSelectionNotification(selection Atom) bool
	// SetDoubleClickDistance sets the double click distance (two clicks within
	// this distance count as a double click and result in a K_2BUTTON_PRESS
	// event). See also gdk_display_set_double_click_time(). Applications should
	// not set this, it is a global user-configured setting.
	SetDoubleClickDistance(distance uint)
	// SetDoubleClickTime sets the double click time (two clicks within this
	// time interval count as a double click and result in a K_2BUTTON_PRESS
	// event). Applications should not set this, it is a global user-configured
	// setting.
	SetDoubleClickTime(msec uint)
	// StoreClipboard issues a request to the clipboard manager to store the
	// clipboard data. On X11, this is a special program that works according to
	// the FreeDesktop Clipboard Specification
	// (http://www.freedesktop.org/Standards/clipboard-manager-spec).
	StoreClipboard(clipboardWindow Window, time_ uint32, targets []Atom)
	// SupportsClipboardPersistence returns whether the speicifed display
	// supports clipboard persistance; i.e. if it’s possible to store the
	// clipboard data after an application has quit. On X11 this checks if a
	// clipboard daemon is running.
	SupportsClipboardPersistence() bool
	// SupportsComposite returns true if gdk_window_set_composited() can be used
	// to redirect drawing on the window using compositing.
	//
	// Currently this only works on X11 with XComposite and XDamage extensions
	// available.
	SupportsComposite() bool
	// SupportsCursorAlpha returns true if cursors can use an 8bit alpha channel
	// on @display. Otherwise, cursors are restricted to bilevel alpha (i.e. a
	// mask).
	SupportsCursorAlpha() bool
	// SupportsCursorColor returns true if multicolored cursors are supported on
	// @display. Otherwise, cursors have only a forground and a background
	// color.
	SupportsCursorColor() bool
	// SupportsInputShapes returns true if gdk_window_input_shape_combine_mask()
	// can be used to modify the input shape of windows on @display.
	SupportsInputShapes() bool
	// SupportsSelectionNotification returns whether EventOwnerChange events
	// will be sent when the owner of a selection changes.
	SupportsSelectionNotification() bool
	// SupportsShapes returns true if gdk_window_shape_combine_mask() can be
	// used to create shaped windows on @display.
	SupportsShapes() bool
	// Sync flushes any requests queued for the windowing system and waits until
	// all requests have been handled. This is often used for making sure that
	// the display is synchronized with the current state of the program.
	// Calling gdk_display_sync() before gdk_error_trap_pop() makes sure that
	// any errors generated from earlier requests are handled before the error
	// trap is removed.
	//
	// This is most useful for X11. On windowing systems where requests are
	// handled synchronously, this function will do nothing.
	Sync()
	// WarpPointer warps the pointer of @display to the point @x,@y on the
	// screen @screen, unless the pointer is confined to a window by a grab, in
	// which case it will be moved as far as allowed by the grab. Warping the
	// pointer creates events as if the user had moved the mouse instantaneously
	// to the destination.
	//
	// Note that the pointer should normally be under the control of the user.
	// This function was added to cover some rare use cases like keyboard
	// navigation support for the color picker in the ColorSelectionDialog.
	WarpPointer(screen Screen, x int, y int)
}

// display implements the Display class.
type display struct {
	gextras.Objector
}

var _ Display = (*display)(nil)

// WrapDisplay wraps a GObject to the right type. It is
// primarily used internally.
func WrapDisplay(obj *externglib.Object) Display {
	return display{
		Objector: obj,
	}
}

func marshalDisplay(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDisplay(obj), nil
}

// Beep emits a short beep on @display
func (d display) Beep() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_display_beep(_arg0)
}

// Close closes the connection to the windowing system for the given
// display, and cleans up associated resources.
func (d display) Close() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_display_close(_arg0)
}

// DeviceIsGrabbed returns true if there is an ongoing grab on @device for
// @display.
func (d display) DeviceIsGrabbed(device Device) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkDevice  // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_display_device_is_grabbed(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Flush flushes any requests queued for the windowing system; this happens
// automatically when the main loop blocks waiting for new events, but if
// your application is drawing without returning control to the main loop,
// you may need to call this function explicitly. A common case where this
// function needs to be called is when an application is executing drawing
// commands from a thread other than the thread where the main loop is
// running.
//
// This is most useful for X11. On windowing systems where requests are
// handled synchronously, this function will do nothing.
func (d display) Flush() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_display_flush(_arg0)
}

// AppLaunchContext returns a AppLaunchContext suitable for launching
// applications on the given display.
func (d display) AppLaunchContext() AppLaunchContext {
	var _arg0 *C.GdkDisplay          // out
	var _cret *C.GdkAppLaunchContext // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_app_launch_context(_arg0)

	var _appLaunchContext AppLaunchContext // out

	_appLaunchContext = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(AppLaunchContext)

	return _appLaunchContext
}

// DefaultCursorSize returns the default size to use for cursors on
// @display.
func (d display) DefaultCursorSize() uint {
	var _arg0 *C.GdkDisplay // out
	var _cret C.guint       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_default_cursor_size(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// DefaultGroup returns the default group leader window for all toplevel
// windows on @display. This window is implicitly created by GDK. See
// gdk_window_set_group().
func (d display) DefaultGroup() Window {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkWindow  // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_default_group(_arg0)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Window)

	return _window
}

// DefaultScreen: get the default Screen for @display.
func (d display) DefaultScreen() Screen {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkScreen  // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_default_screen(_arg0)

	var _screen Screen // out

	_screen = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Screen)

	return _screen
}

// DefaultSeat returns the default Seat for this display.
func (d display) DefaultSeat() Seat {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkSeat    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_default_seat(_arg0)

	var _seat Seat // out

	_seat = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Seat)

	return _seat
}

// DeviceManager returns the DeviceManager associated to @display.
func (d display) DeviceManager() DeviceManager {
	var _arg0 *C.GdkDisplay       // out
	var _cret *C.GdkDeviceManager // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_device_manager(_arg0)

	var _deviceManager DeviceManager // out

	_deviceManager = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(DeviceManager)

	return _deviceManager
}

// MaximalCursorSize gets the maximal size to use for cursors on @display.
func (d display) MaximalCursorSize() (width uint, height uint) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint       // in
	var _arg2 C.guint       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_display_get_maximal_cursor_size(_arg0, &_arg1, &_arg2)

	var _width uint  // out
	var _height uint // out

	_width = (uint)(_arg1)
	_height = (uint)(_arg2)

	return _width, _height
}

// Monitor gets a monitor associated with this display.
func (d display) Monitor(monitorNum int) Monitor {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.int         // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (C.int)(monitorNum)

	_cret = C.gdk_display_get_monitor(_arg0, _arg1)

	var _monitor Monitor // out

	_monitor = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Monitor)

	return _monitor
}

// MonitorAtPoint gets the monitor in which the point (@x, @y) is located,
// or a nearby monitor if the point is not in any monitor.
func (d display) MonitorAtPoint(x int, y int) Monitor {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.int         // out
	var _arg2 C.int         // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (C.int)(x)
	_arg2 = (C.int)(y)

	_cret = C.gdk_display_get_monitor_at_point(_arg0, _arg1, _arg2)

	var _monitor Monitor // out

	_monitor = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Monitor)

	return _monitor
}

// MonitorAtWindow gets the monitor in which the largest area of @window
// resides, or a monitor close to @window if it is outside of all monitors.
func (d display) MonitorAtWindow(window Window) Monitor {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkWindow  // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gdk_display_get_monitor_at_window(_arg0, _arg1)

	var _monitor Monitor // out

	_monitor = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Monitor)

	return _monitor
}

// NMonitors gets the number of monitors that belong to @display.
//
// The returned number is valid until the next emission of the
// Display::monitor-added or Display::monitor-removed signal.
func (d display) NMonitors() int {
	var _arg0 *C.GdkDisplay // out
	var _cret C.int         // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_n_monitors(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// NScreens gets the number of screen managed by the @display.
func (d display) NScreens() int {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gint        // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_n_screens(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Name gets the name of the display.
func (d display) Name() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Pointer gets the current location of the pointer and the current modifier
// mask for a given display.
func (d display) Pointer() (screen Screen, x int, y int, mask ModifierType) {
	var _arg0 *C.GdkDisplay     // out
	var _arg1 *C.GdkScreen      // in
	var _arg2 C.gint            // in
	var _arg3 C.gint            // in
	var _arg4 C.GdkModifierType // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_display_get_pointer(_arg0, &_arg1, &_arg2, &_arg3, &_arg4)

	var _screen Screen     // out
	var _x int             // out
	var _y int             // out
	var _mask ModifierType // out

	_screen = gextras.CastObject(externglib.Take(unsafe.Pointer(_arg1.Native()))).(Screen)
	_x = (int)(_arg2)
	_y = (int)(_arg3)
	_mask = ModifierType(_arg4)

	return _screen, _x, _y, _mask
}

// PrimaryMonitor gets the primary monitor for the display.
//
// The primary monitor is considered the monitor where the “main desktop”
// lives. While normal application windows typically allow the window
// manager to place the windows, specialized desktop applications such as
// panels should place themselves on the primary monitor.
func (d display) PrimaryMonitor() Monitor {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_primary_monitor(_arg0)

	var _monitor Monitor // out

	_monitor = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Monitor)

	return _monitor
}

// Screen returns a screen object for one of the screens of the display.
func (d display) Screen(screenNum int) Screen {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.gint        // out
	var _cret *C.GdkScreen  // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (C.gint)(screenNum)

	_cret = C.gdk_display_get_screen(_arg0, _arg1)

	var _screen Screen // out

	_screen = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Screen)

	return _screen
}

// WindowAtPointer obtains the window underneath the mouse pointer,
// returning the location of the pointer in that window in @win_x, @win_y
// for @screen. Returns nil if the window under the mouse pointer is not
// known to GDK (for example, belongs to another application).
func (d display) WindowAtPointer() (winX int, winY int, window Window) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.gint        // in
	var _arg2 C.gint        // in
	var _cret *C.GdkWindow  // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_window_at_pointer(_arg0, &_arg1, &_arg2)

	var _winX int      // out
	var _winY int      // out
	var _window Window // out

	_winX = (int)(_arg1)
	_winY = (int)(_arg2)
	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Window)

	return _winX, _winY, _window
}

// HasPending returns whether the display has events that are waiting to be
// processed.
func (d display) HasPending() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_has_pending(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsClosed finds out if the display has been closed.
func (d display) IsClosed() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_is_closed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// KeyboardUngrab: release any keyboard grab
func (d display) KeyboardUngrab(time_ uint32) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint32     // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (C.guint32)(time_)

	C.gdk_display_keyboard_ungrab(_arg0, _arg1)
}

// NotifyStartupComplete indicates to the GUI environment that the
// application has finished loading, using a given identifier.
//
// GTK+ will call this function automatically for Window with custom
// startup-notification identifier unless
// gtk_window_set_auto_startup_notification() is called to disable that
// feature.
func (d display) NotifyStartupComplete(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.gchar      // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.gchar)(C.CString(startupId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_display_notify_startup_complete(_arg0, _arg1)
}

// PointerIsGrabbed: test if the pointer is grabbed.
func (d display) PointerIsGrabbed() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_pointer_is_grabbed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PointerUngrab: release any pointer grab.
func (d display) PointerUngrab(time_ uint32) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint32     // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (C.guint32)(time_)

	C.gdk_display_pointer_ungrab(_arg0, _arg1)
}

// RequestSelectionNotification: request EventOwnerChange events for
// ownership changes of the selection named by the given atom.
func (d display) RequestSelectionNotification(selection Atom) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.GdkAtom     // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = *(*C.GdkAtom)(unsafe.Pointer(selection.Native()))

	_cret = C.gdk_display_request_selection_notification(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDoubleClickDistance sets the double click distance (two clicks within
// this distance count as a double click and result in a K_2BUTTON_PRESS
// event). See also gdk_display_set_double_click_time(). Applications should
// not set this, it is a global user-configured setting.
func (d display) SetDoubleClickDistance(distance uint) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (C.guint)(distance)

	C.gdk_display_set_double_click_distance(_arg0, _arg1)
}

// SetDoubleClickTime sets the double click time (two clicks within this
// time interval count as a double click and result in a K_2BUTTON_PRESS
// event). Applications should not set this, it is a global user-configured
// setting.
func (d display) SetDoubleClickTime(msec uint) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (C.guint)(msec)

	C.gdk_display_set_double_click_time(_arg0, _arg1)
}

// StoreClipboard issues a request to the clipboard manager to store the
// clipboard data. On X11, this is a special program that works according to
// the FreeDesktop Clipboard Specification
// (http://www.freedesktop.org/Standards/clipboard-manager-spec).
func (d display) StoreClipboard(clipboardWindow Window, time_ uint32, targets []Atom) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkWindow  // out
	var _arg2 C.guint32     // out
	var _arg3 *C.GdkAtom
	var _arg4 C.gint

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(clipboardWindow.Native()))
	_arg2 = (C.guint32)(time_)
	_arg4 = C.gint(len(targets))
	_arg3 = (*C.GdkAtom)(unsafe.Pointer(&targets[0]))

	C.gdk_display_store_clipboard(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// SupportsClipboardPersistence returns whether the speicifed display
// supports clipboard persistance; i.e. if it’s possible to store the
// clipboard data after an application has quit. On X11 this checks if a
// clipboard daemon is running.
func (d display) SupportsClipboardPersistence() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_supports_clipboard_persistence(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsComposite returns true if gdk_window_set_composited() can be used
// to redirect drawing on the window using compositing.
//
// Currently this only works on X11 with XComposite and XDamage extensions
// available.
func (d display) SupportsComposite() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_supports_composite(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsCursorAlpha returns true if cursors can use an 8bit alpha channel
// on @display. Otherwise, cursors are restricted to bilevel alpha (i.e. a
// mask).
func (d display) SupportsCursorAlpha() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_supports_cursor_alpha(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsCursorColor returns true if multicolored cursors are supported on
// @display. Otherwise, cursors have only a forground and a background
// color.
func (d display) SupportsCursorColor() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_supports_cursor_color(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsInputShapes returns true if gdk_window_input_shape_combine_mask()
// can be used to modify the input shape of windows on @display.
func (d display) SupportsInputShapes() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_supports_input_shapes(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsSelectionNotification returns whether EventOwnerChange events
// will be sent when the owner of a selection changes.
func (d display) SupportsSelectionNotification() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_supports_selection_notification(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsShapes returns true if gdk_window_shape_combine_mask() can be
// used to create shaped windows on @display.
func (d display) SupportsShapes() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_supports_shapes(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Sync flushes any requests queued for the windowing system and waits until
// all requests have been handled. This is often used for making sure that
// the display is synchronized with the current state of the program.
// Calling gdk_display_sync() before gdk_error_trap_pop() makes sure that
// any errors generated from earlier requests are handled before the error
// trap is removed.
//
// This is most useful for X11. On windowing systems where requests are
// handled synchronously, this function will do nothing.
func (d display) Sync() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_display_sync(_arg0)
}

// WarpPointer warps the pointer of @display to the point @x,@y on the
// screen @screen, unless the pointer is confined to a window by a grab, in
// which case it will be moved as far as allowed by the grab. Warping the
// pointer creates events as if the user had moved the mouse instantaneously
// to the destination.
//
// Note that the pointer should normally be under the control of the user.
// This function was added to cover some rare use cases like keyboard
// navigation support for the color picker in the ColorSelectionDialog.
func (d display) WarpPointer(screen Screen, x int, y int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkScreen  // out
	var _arg2 C.gint        // out
	var _arg3 C.gint        // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg2 = (C.gint)(x)
	_arg3 = (C.gint)(y)

	C.gdk_display_warp_pointer(_arg0, _arg1, _arg2, _arg3)
}
