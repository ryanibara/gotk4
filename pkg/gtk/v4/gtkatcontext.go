// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_ATContext_ConnectStateChange(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeATContext = coreglib.Type(C.gtk_at_context_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeATContext, F: marshalATContext},
	})
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

// ConnectStateChange is emitted when the attributes of the accessible for the
// GtkATContext instance change.
func (self *ATContext) ConnectStateChange(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "state-change", false, unsafe.Pointer(C._gotk4_gtk4_ATContext_ConnectStateChange), f)
}

// NewATContextCreate creates a new GtkATContext instance for the given
// accessible role, accessible instance, and display connection.
//
// The GtkATContext implementation being instantiated will depend on the
// platform.
//
// The function takes the following parameters:
//
//   - accessibleRole: accessible role used by the GtkATContext.
//   - accessible: GtkAccessible implementation using the GtkATContext.
//   - display: GdkDisplay used by the GtkATContext.
//
// The function returns the following values:
//
//   - atContext (optional): GtkATContext.
//
func NewATContextCreate(accessibleRole AccessibleRole, accessible Accessibler, display *gdk.Display) *ATContext {
	var _arg1 C.GtkAccessibleRole // out
	var _arg2 *C.GtkAccessible    // out
	var _arg3 *C.GdkDisplay       // out
	var _cret *C.GtkATContext     // in

	_arg1 = C.GtkAccessibleRole(accessibleRole)
	_arg2 = (*C.GtkAccessible)(unsafe.Pointer(coreglib.InternObject(accessible).Native()))
	_arg3 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_cret = C.gtk_at_context_create(_arg1, _arg2, _arg3)
	runtime.KeepAlive(accessibleRole)
	runtime.KeepAlive(accessible)
	runtime.KeepAlive(display)

	var _atContext *ATContext // out

	if _cret != nil {
		_atContext = wrapATContext(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _atContext
}

// Accessible retrieves the GtkAccessible using this context.
//
// The function returns the following values:
//
//   - accessible: GtkAccessible.
//
func (self *ATContext) Accessible() *Accessible {
	var _arg0 *C.GtkATContext  // out
	var _cret *C.GtkAccessible // in

	_arg0 = (*C.GtkATContext)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_at_context_get_accessible(_arg0)
	runtime.KeepAlive(self)

	var _accessible *Accessible // out

	_accessible = wrapAccessible(coreglib.Take(unsafe.Pointer(_cret)))

	return _accessible
}

// AccessibleRole retrieves the accessible role of this context.
//
// The function returns the following values:
//
//   - accessibleRole: GtkAccessibleRole.
//
func (self *ATContext) AccessibleRole() AccessibleRole {
	var _arg0 *C.GtkATContext     // out
	var _cret C.GtkAccessibleRole // in

	_arg0 = (*C.GtkATContext)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_at_context_get_accessible_role(_arg0)
	runtime.KeepAlive(self)

	var _accessibleRole AccessibleRole // out

	_accessibleRole = AccessibleRole(_cret)

	return _accessibleRole
}
