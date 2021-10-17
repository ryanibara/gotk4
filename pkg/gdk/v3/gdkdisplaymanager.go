// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_display_manager_get_type()), F: marshalDisplayManagerer},
	})
}

// DisplayManager: purpose of the DisplayManager singleton object is to offer
// notification when displays appear or disappear or the default display
// changes.
//
// You can use gdk_display_manager_get() to obtain the DisplayManager singleton,
// but that should be rarely necessary. Typically, initializing GTK+ opens a
// display that you can work with without ever accessing the DisplayManager.
//
// The GDK library can be built with support for multiple backends. The
// DisplayManager object determines which backend is used at runtime.
//
// When writing backend-specific code that is supposed to work with multiple GDK
// backends, you have to consider both compile time and runtime. At compile
// time, use the K_WINDOWING_X11, K_WINDOWING_WIN32 macros, etc. to find out
// which backends are present in the GDK library you are building your
// application against. At runtime, use type-check macros like
// GDK_IS_X11_DISPLAY() to find out which backend is in use:
//
// Backend-specific code
//
//    #ifdef GDK_WINDOWING_X11
//      if (GDK_IS_X11_DISPLAY (display))
//        {
//          // make X11-specific calls here
//        }
//      else
//    #endif
//    #ifdef GDK_WINDOWING_QUARTZ
//      if (GDK_IS_QUARTZ_DISPLAY (display))
//        {
//          // make Quartz-specific calls here
//        }
//      else
//    #endif
//      g_error ("Unsupported GDK backend");.
type DisplayManager struct {
	*externglib.Object
}

func wrapDisplayManager(obj *externglib.Object) *DisplayManager {
	return &DisplayManager{
		Object: obj,
	}
}

func marshalDisplayManagerer(p uintptr) (interface{}, error) {
	return wrapDisplayManager(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DefaultDisplay gets the default Display.
func (manager *DisplayManager) DefaultDisplay() *Display {
	var _arg0 *C.GdkDisplayManager // out
	var _cret *C.GdkDisplay        // in

	_arg0 = (*C.GdkDisplayManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gdk_display_manager_get_default_display(_arg0)
	runtime.KeepAlive(manager)

	var _display *Display // out

	if _cret != nil {
		_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _display
}

// ListDisplays: list all currently open displays.
func (manager *DisplayManager) ListDisplays() []Display {
	var _arg0 *C.GdkDisplayManager // out
	var _cret *C.GSList            // in

	_arg0 = (*C.GdkDisplayManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gdk_display_manager_list_displays(_arg0)
	runtime.KeepAlive(manager)

	var _sList []Display // out

	_sList = make([]Display, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GdkDisplay)(v)
		var dst Display // out
		dst = *wrapDisplay(externglib.Take(unsafe.Pointer(src)))
		_sList = append(_sList, dst)
	})

	return _sList
}

// OpenDisplay opens a display.
//
// The function takes the following parameters:
//
//    - name of the display to open.
//
func (manager *DisplayManager) OpenDisplay(name string) *Display {
	var _arg0 *C.GdkDisplayManager // out
	var _arg1 *C.gchar             // out
	var _cret *C.GdkDisplay        // in

	_arg0 = (*C.GdkDisplayManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_display_manager_open_display(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(name)

	var _display *Display // out

	if _cret != nil {
		_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _display
}

// SetDefaultDisplay sets display as the default display.
//
// The function takes the following parameters:
//
//    - display: Display.
//
func (manager *DisplayManager) SetDefaultDisplay(display *Display) {
	var _arg0 *C.GdkDisplayManager // out
	var _arg1 *C.GdkDisplay        // out

	_arg0 = (*C.GdkDisplayManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_manager_set_default_display(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(display)
}

// ConnectDisplayOpened signal is emitted when a display is opened.
func (manager *DisplayManager) ConnectDisplayOpened(f func(display Display)) externglib.SignalHandle {
	return manager.Connect("display-opened", f)
}

// DisplayManagerGet gets the singleton DisplayManager object.
//
// When called for the first time, this function consults the GDK_BACKEND
// environment variable to find out which of the supported GDK backends to use
// (in case GDK has been compiled with multiple backends). Applications can use
// gdk_set_allowed_backends() to limit what backends can be used.
func DisplayManagerGet() *DisplayManager {
	var _cret *C.GdkDisplayManager // in

	_cret = C.gdk_display_manager_get()

	var _displayManager *DisplayManager // out

	_displayManager = wrapDisplayManager(externglib.Take(unsafe.Pointer(_cret)))

	return _displayManager
}
