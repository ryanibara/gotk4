// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_at_context_get_type()), F: marshalATContexter},
	})
}

// ATContext: GtkATContext is an abstract class provided by GTK to communicate
// to platform-specific assistive technologies API.
//
// Each platform supported by GTK implements a GtkATContext subclass, and is
// responsible for updating the accessible state in response to state changes in
// GtkAccessible.
type ATContext struct {
	*externglib.Object
}

// ATContexter describes ATContext's abstract methods.
type ATContexter interface {
	externglib.Objector

	// Accessible retrieves the GtkAccessible using this context.
	Accessible() Accessibler
	// AccessibleRole retrieves the accessible role of this context.
	AccessibleRole() AccessibleRole
}

var _ ATContexter = (*ATContext)(nil)

func wrapATContext(obj *externglib.Object) *ATContext {
	return &ATContext{
		Object: obj,
	}
}

func marshalATContexter(p uintptr) (interface{}, error) {
	return wrapATContext(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewATContextCreate creates a new GtkATContext instance for the given
// accessible role, accessible instance, and display connection.
//
// The GtkATContext implementation being instantiated will depend on the
// platform.
//
// The function takes the following parameters:
//
//    - accessibleRole: accessible role used by the GtkATContext.
//    - accessible: GtkAccessible implementation using the GtkATContext.
//    - display: GdkDisplay used by the GtkATContext.
//
func NewATContextCreate(accessibleRole AccessibleRole, accessible Accessibler, display *gdk.Display) *ATContext {
	var _arg1 C.GtkAccessibleRole // out
	var _arg2 *C.GtkAccessible    // out
	var _arg3 *C.GdkDisplay       // out
	var _cret *C.GtkATContext     // in

	_arg1 = C.GtkAccessibleRole(accessibleRole)
	_arg2 = (*C.GtkAccessible)(unsafe.Pointer(accessible.Native()))
	_arg3 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gtk_at_context_create(_arg1, _arg2, _arg3)
	runtime.KeepAlive(accessibleRole)
	runtime.KeepAlive(accessible)
	runtime.KeepAlive(display)

	var _atContext *ATContext // out

	if _cret != nil {
		_atContext = wrapATContext(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _atContext
}

// Accessible retrieves the GtkAccessible using this context.
func (self *ATContext) Accessible() Accessibler {
	var _arg0 *C.GtkATContext  // out
	var _cret *C.GtkAccessible // in

	_arg0 = (*C.GtkATContext)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_at_context_get_accessible(_arg0)
	runtime.KeepAlive(self)

	var _accessible Accessibler // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Accessibler is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Accessibler)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Accessibler")
		}
		_accessible = rv
	}

	return _accessible
}

// AccessibleRole retrieves the accessible role of this context.
func (self *ATContext) AccessibleRole() AccessibleRole {
	var _arg0 *C.GtkATContext     // out
	var _cret C.GtkAccessibleRole // in

	_arg0 = (*C.GtkATContext)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_at_context_get_accessible_role(_arg0)
	runtime.KeepAlive(self)

	var _accessibleRole AccessibleRole // out

	_accessibleRole = AccessibleRole(_cret)

	return _accessibleRole
}

// ConnectStateChange: emitted when the attributes of the accessible for the
// GtkATContext instance change.
func (self *ATContext) ConnectStateChange(f func()) externglib.SignalHandle {
	return self.Connect("state-change", f)
}
