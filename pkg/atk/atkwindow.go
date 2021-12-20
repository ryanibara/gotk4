// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_window_get_type()), F: marshalWindower},
	})
}

// Window should be implemented by the UI elements that represent a top-level
// window, such as the main window of an application or dialog.
type Window struct {
	ObjectClass

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ externglib.Objector = (*Window)(nil)
)

// Windower describes Window's interface methods.
type Windower interface {
	externglib.Objector

	baseWindow() *Window
}

var _ Windower = (*Window)(nil)

func wrapWindow(obj *externglib.Object) *Window {
	return &Window{
		ObjectClass: ObjectClass{
			Object: obj,
		},
	}
}

func marshalWindower(p uintptr) (interface{}, error) {
	return wrapWindow(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Window) baseWindow() *Window {
	return v
}

// BaseWindow returns the underlying base object.
func BaseWindow(obj Windower) *Window {
	return obj.baseWindow()
}

// ConnectActivate: signal Window::activate is emitted when a window becomes the
// active window of the application or session.
func (v *Window) ConnectActivate(f func()) externglib.SignalHandle {
	return v.Connect("activate", f)
}

// ConnectCreate: signal Window::create is emitted when a new window is created.
func (v *Window) ConnectCreate(f func()) externglib.SignalHandle {
	return v.Connect("create", f)
}

// ConnectDeactivate: signal Window::deactivate is emitted when a window is no
// longer the active window of the application or session.
func (v *Window) ConnectDeactivate(f func()) externglib.SignalHandle {
	return v.Connect("deactivate", f)
}

// ConnectDestroy: signal Window::destroy is emitted when a window is destroyed.
func (v *Window) ConnectDestroy(f func()) externglib.SignalHandle {
	return v.Connect("destroy", f)
}

// ConnectMaximize: signal Window::maximize is emitted when a window is
// maximized.
func (v *Window) ConnectMaximize(f func()) externglib.SignalHandle {
	return v.Connect("maximize", f)
}

// ConnectMinimize: signal Window::minimize is emitted when a window is
// minimized.
func (v *Window) ConnectMinimize(f func()) externglib.SignalHandle {
	return v.Connect("minimize", f)
}

// ConnectMove: signal Window::move is emitted when a window is moved.
func (v *Window) ConnectMove(f func()) externglib.SignalHandle {
	return v.Connect("move", f)
}

// ConnectResize: signal Window::resize is emitted when a window is resized.
func (v *Window) ConnectResize(f func()) externglib.SignalHandle {
	return v.Connect("resize", f)
}

// ConnectRestore: signal Window::restore is emitted when a window is restored.
func (v *Window) ConnectRestore(f func()) externglib.SignalHandle {
	return v.Connect("restore", f)
}
