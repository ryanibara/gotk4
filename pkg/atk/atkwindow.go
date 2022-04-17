// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
// extern void _gotk4_atk1_Window_ConnectActivate(gpointer, guintptr);
// extern void _gotk4_atk1_Window_ConnectCreate(gpointer, guintptr);
// extern void _gotk4_atk1_Window_ConnectDeactivate(gpointer, guintptr);
// extern void _gotk4_atk1_Window_ConnectDestroy(gpointer, guintptr);
// extern void _gotk4_atk1_Window_ConnectMaximize(gpointer, guintptr);
// extern void _gotk4_atk1_Window_ConnectMinimize(gpointer, guintptr);
// extern void _gotk4_atk1_Window_ConnectMove(gpointer, guintptr);
// extern void _gotk4_atk1_Window_ConnectResize(gpointer, guintptr);
// extern void _gotk4_atk1_Window_ConnectRestore(gpointer, guintptr);
import "C"

// glib.Type values for atkwindow.go.
var GTypeWindow = externglib.Type(C.atk_window_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeWindow, F: marshalWindow},
	})
}

// WindowOverrider contains methods that are overridable.
type WindowOverrider interface {
	externglib.Objector
}

// Window should be implemented by the UI elements that represent a top-level
// window, such as the main window of an application or dialog.
type Window struct {
	_ [0]func() // equal guard
	ObjectClass
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

func ifaceInitWindower(gifacePtr, data C.gpointer) {
}

func wrapWindow(obj *externglib.Object) *Window {
	return &Window{
		ObjectClass: ObjectClass{
			Object: obj,
		},
	}
}

func marshalWindow(p uintptr) (interface{}, error) {
	return wrapWindow(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Window) baseWindow() *Window {
	return v
}

// BaseWindow returns the underlying base object.
func BaseWindow(obj Windower) *Window {
	return obj.baseWindow()
}

//export _gotk4_atk1_Window_ConnectActivate
func _gotk4_atk1_Window_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectActivate: signal Window::activate is emitted when a window becomes the
// active window of the application or session.
func (v *Window) ConnectActivate(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "activate", false, unsafe.Pointer(C._gotk4_atk1_Window_ConnectActivate), f)
}

//export _gotk4_atk1_Window_ConnectCreate
func _gotk4_atk1_Window_ConnectCreate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectCreate: signal Window::create is emitted when a new window is created.
func (v *Window) ConnectCreate(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "create", false, unsafe.Pointer(C._gotk4_atk1_Window_ConnectCreate), f)
}

//export _gotk4_atk1_Window_ConnectDeactivate
func _gotk4_atk1_Window_ConnectDeactivate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectDeactivate: signal Window::deactivate is emitted when a window is no
// longer the active window of the application or session.
func (v *Window) ConnectDeactivate(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "deactivate", false, unsafe.Pointer(C._gotk4_atk1_Window_ConnectDeactivate), f)
}

//export _gotk4_atk1_Window_ConnectDestroy
func _gotk4_atk1_Window_ConnectDestroy(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectDestroy: signal Window::destroy is emitted when a window is destroyed.
func (v *Window) ConnectDestroy(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "destroy", false, unsafe.Pointer(C._gotk4_atk1_Window_ConnectDestroy), f)
}

//export _gotk4_atk1_Window_ConnectMaximize
func _gotk4_atk1_Window_ConnectMaximize(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectMaximize: signal Window::maximize is emitted when a window is
// maximized.
func (v *Window) ConnectMaximize(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "maximize", false, unsafe.Pointer(C._gotk4_atk1_Window_ConnectMaximize), f)
}

//export _gotk4_atk1_Window_ConnectMinimize
func _gotk4_atk1_Window_ConnectMinimize(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectMinimize: signal Window::minimize is emitted when a window is
// minimized.
func (v *Window) ConnectMinimize(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "minimize", false, unsafe.Pointer(C._gotk4_atk1_Window_ConnectMinimize), f)
}

//export _gotk4_atk1_Window_ConnectMove
func _gotk4_atk1_Window_ConnectMove(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectMove: signal Window::move is emitted when a window is moved.
func (v *Window) ConnectMove(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "move", false, unsafe.Pointer(C._gotk4_atk1_Window_ConnectMove), f)
}

//export _gotk4_atk1_Window_ConnectResize
func _gotk4_atk1_Window_ConnectResize(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectResize: signal Window::resize is emitted when a window is resized.
func (v *Window) ConnectResize(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "resize", false, unsafe.Pointer(C._gotk4_atk1_Window_ConnectResize), f)
}

//export _gotk4_atk1_Window_ConnectRestore
func _gotk4_atk1_Window_ConnectRestore(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectRestore: signal Window::restore is emitted when a window is restored.
func (v *Window) ConnectRestore(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "restore", false, unsafe.Pointer(C._gotk4_atk1_Window_ConnectRestore), f)
}
