// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gdk3_DisplayManager_ConnectDisplayOpened(gpointer, GdkDisplay*, guintptr);
import "C"

// glib.Type values for gdkdisplaymanager.go.
var GTypeDisplayManager = coreglib.Type(C.gdk_display_manager_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeDisplayManager, F: marshalDisplayManager},
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
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DisplayManager)(nil)
)

func wrapDisplayManager(obj *coreglib.Object) *DisplayManager {
	return &DisplayManager{
		Object: obj,
	}
}

func marshalDisplayManager(p uintptr) (interface{}, error) {
	return wrapDisplayManager(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gdk3_DisplayManager_ConnectDisplayOpened
func _gotk4_gdk3_DisplayManager_ConnectDisplayOpened(arg0 C.gpointer, arg1 *C.GdkDisplay, arg2 C.guintptr) {
	var f func(display *Display)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(display *Display))
	}

	var _display *Display // out

	_display = wrapDisplay(coreglib.Take(unsafe.Pointer(arg1)))

	f(_display)
}

// ConnectDisplayOpened signal is emitted when a display is opened.
func (manager *DisplayManager) ConnectDisplayOpened(f func(display *Display)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(manager, "display-opened", false, unsafe.Pointer(C._gotk4_gdk3_DisplayManager_ConnectDisplayOpened), f)
}

// DefaultDisplay gets the default Display.
//
// The function returns the following values:
//
//    - display (optional) or NULL if there is no default display.
//
func (manager *DisplayManager) DefaultDisplay() *Display {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	*(**DisplayManager)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "DisplayManager").InvokeMethod("get_default_display", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(manager)

	var _display *Display // out

	if _cret != nil {
		_display = wrapDisplay(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _display
}

// ListDisplays: list all currently open displays.
//
// The function returns the following values:
//
//    - sList: newly allocated List of Display objects. Free with g_slist_free()
//      when you are done with it.
//
func (manager *DisplayManager) ListDisplays() []*Display {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	*(**DisplayManager)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "DisplayManager").InvokeMethod("list_displays", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(manager)

	var _sList []*Display // out

	_sList = make([]*Display, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst *Display // out
		dst = wrapDisplay(coreglib.Take(unsafe.Pointer(src)))
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
// The function returns the following values:
//
//    - display (optional) or NULL if the display could not be opened.
//
func (manager *DisplayManager) OpenDisplay(name string) *Display {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**DisplayManager)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gdk", "DisplayManager").InvokeMethod("open_display", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(manager)
	runtime.KeepAlive(name)

	var _display *Display // out

	if _cret != nil {
		_display = wrapDisplay(coreglib.Take(unsafe.Pointer(_cret)))
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
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(display).Native()))
	*(**DisplayManager)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gdk", "DisplayManager").InvokeMethod("set_default_display", args[:], nil)

	runtime.KeepAlive(manager)
	runtime.KeepAlive(display)
}

// DisplayManagerGet gets the singleton DisplayManager object.
//
// When called for the first time, this function consults the GDK_BACKEND
// environment variable to find out which of the supported GDK backends to use
// (in case GDK has been compiled with multiple backends). Applications can use
// gdk_set_allowed_backends() to limit what backends can be used.
//
// The function returns the following values:
//
//    - displayManager: global DisplayManager singleton; gdk_parse_args(),
//      gdk_init(), or gdk_init_check() must have been called first.
//
func DisplayManagerGet() *DisplayManager {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gdk", "get").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _displayManager *DisplayManager // out

	_displayManager = wrapDisplayManager(coreglib.Take(unsafe.Pointer(_cret)))

	return _displayManager
}
