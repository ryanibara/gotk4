// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_ATContext_ConnectStateChange(gpointer, guintptr);
import "C"

// GTypeATContext returns the GType for the type ATContext.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeATContext() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ATContext").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalATContext)
	return gtype
}

// ATContext: GtkATContext is an abstract class provided by GTK to communicate
// to platform-specific assistive technologies API.
//
// Each platform supported by GTK implements a GtkATContext subclass, and is
// responsible for updating the accessible state in response to state changes in
// GtkAccessible.
type ATContext struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ATContext)(nil)
)

// ATContexter describes types inherited from class ATContext.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type ATContexter interface {
	coreglib.Objector
	baseATContext() *ATContext
}

var _ ATContexter = (*ATContext)(nil)

func wrapATContext(obj *coreglib.Object) *ATContext {
	return &ATContext{
		Object: obj,
	}
}

func marshalATContext(p uintptr) (interface{}, error) {
	return wrapATContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *ATContext) baseATContext() *ATContext {
	return self
}

// BaseATContext returns the underlying base object.
func BaseATContext(obj ATContexter) *ATContext {
	return obj.baseATContext()
}

//export _gotk4_gtk4_ATContext_ConnectStateChange
func _gotk4_gtk4_ATContext_ConnectStateChange(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectStateChange is emitted when the attributes of the accessible for the
// GtkATContext instance change.
func (self *ATContext) ConnectStateChange(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "state-change", false, unsafe.Pointer(C._gotk4_gtk4_ATContext_ConnectStateChange), f)
}

// Accessible retrieves the GtkAccessible using this context.
//
// The function returns the following values:
//
//    - accessible: GtkAccessible.
//
func (self *ATContext) Accessible() *Accessible {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "ATContext")
	_gret := _info.InvokeClassMethod("get_accessible", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _accessible *Accessible // out

	_accessible = wrapAccessible(coreglib.Take(unsafe.Pointer(_cret)))

	return _accessible
}
