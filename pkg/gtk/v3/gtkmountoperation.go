// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeMountOperation = coreglib.Type(C.gtk_mount_operation_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMountOperation, F: marshalMountOperation},
	})
}

// MountOperationOverrides contains methods that are overridable.
type MountOperationOverrides struct {
}

func defaultMountOperationOverrides(v *MountOperation) MountOperationOverrides {
	return MountOperationOverrides{}
}

// MountOperation: this should not be accessed directly. Use the accessor
// functions below.
type MountOperation struct {
	_ [0]func() // equal guard
	gio.MountOperation
}

var (
	_ coreglib.Objector = (*MountOperation)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*MountOperation, *MountOperationClass, MountOperationOverrides](
		GTypeMountOperation,
		initMountOperationClass,
		wrapMountOperation,
		defaultMountOperationOverrides,
	)
}

func initMountOperationClass(gclass unsafe.Pointer, overrides MountOperationOverrides, classInitFunc func(*MountOperationClass)) {
	if classInitFunc != nil {
		class := (*MountOperationClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapMountOperation(obj *coreglib.Object) *MountOperation {
	return &MountOperation{
		MountOperation: gio.MountOperation{
			Object: obj,
		},
	}
}

func marshalMountOperation(p uintptr) (interface{}, error) {
	return wrapMountOperation(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMountOperation creates a new MountOperation.
//
// The function takes the following parameters:
//
//   - parent (optional): transient parent of the window, or NULL.
//
// The function returns the following values:
//
//   - mountOperation: new MountOperation.
//
func NewMountOperation(parent *Window) *MountOperation {
	var _arg1 *C.GtkWindow       // out
	var _cret *C.GMountOperation // in

	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	}

	_cret = C.gtk_mount_operation_new(_arg1)
	runtime.KeepAlive(parent)

	var _mountOperation *MountOperation // out

	_mountOperation = wrapMountOperation(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mountOperation
}

// Parent gets the transient parent used by the MountOperation.
//
// The function returns the following values:
//
//   - window: transient parent for windows shown by op.
//
func (op *MountOperation) Parent() *Window {
	var _arg0 *C.GtkMountOperation // out
	var _cret *C.GtkWindow         // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(coreglib.InternObject(op).Native()))

	_cret = C.gtk_mount_operation_get_parent(_arg0)
	runtime.KeepAlive(op)

	var _window *Window // out

	_window = wrapWindow(coreglib.Take(unsafe.Pointer(_cret)))

	return _window
}

// Screen gets the screen on which windows of the MountOperation will be shown.
//
// The function returns the following values:
//
//   - screen on which windows of op are shown.
//
func (op *MountOperation) Screen() *gdk.Screen {
	var _arg0 *C.GtkMountOperation // out
	var _cret *C.GdkScreen         // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(coreglib.InternObject(op).Native()))

	_cret = C.gtk_mount_operation_get_screen(_arg0)
	runtime.KeepAlive(op)

	var _screen *gdk.Screen // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_screen = &gdk.Screen{
			Object: obj,
		}
	}

	return _screen
}

// IsShowing returns whether the MountOperation is currently displaying a
// window.
//
// The function returns the following values:
//
//   - ok: TRUE if op is currently displaying a window.
//
func (op *MountOperation) IsShowing() bool {
	var _arg0 *C.GtkMountOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(coreglib.InternObject(op).Native()))

	_cret = C.gtk_mount_operation_is_showing(_arg0)
	runtime.KeepAlive(op)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetParent sets the transient parent for windows shown by the MountOperation.
//
// The function takes the following parameters:
//
//   - parent (optional): transient parent of the window, or NULL.
//
func (op *MountOperation) SetParent(parent *Window) {
	var _arg0 *C.GtkMountOperation // out
	var _arg1 *C.GtkWindow         // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(coreglib.InternObject(op).Native()))
	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	}

	C.gtk_mount_operation_set_parent(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(parent)
}

// SetScreen sets the screen to show windows of the MountOperation on.
//
// The function takes the following parameters:
//
//   - screen: Screen.
//
func (op *MountOperation) SetScreen(screen *gdk.Screen) {
	var _arg0 *C.GtkMountOperation // out
	var _arg1 *C.GdkScreen         // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(coreglib.InternObject(op).Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	C.gtk_mount_operation_set_screen(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(screen)
}

// MountOperationClass: instance of this type is always passed by reference.
type MountOperationClass struct {
	*mountOperationClass
}

// mountOperationClass is the struct that's finalized.
type mountOperationClass struct {
	native *C.GtkMountOperationClass
}

// ParentClass: parent class.
func (m *MountOperationClass) ParentClass() *gio.MountOperationClass {
	valptr := &m.native.parent_class
	var _v *gio.MountOperationClass // out
	_v = (*gio.MountOperationClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
