// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeMountOperation returns the GType for the type MountOperation.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeMountOperation() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "MountOperation").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalMountOperation)
	return gtype
}

// MountOperationOverrider contains methods that are overridable.
type MountOperationOverrider interface {
}

// MountOperation: GtkMountOperation is an implementation of GMountOperation.
//
// The functions and objects described here make working with GTK and GIO more
// convenient.
//
// GtkMountOperation is needed when mounting volumes: It is an implementation of
// GMountOperation that can be used with GIO functions for mounting volumes such
// as g_file_mount_enclosing_volume(), g_file_mount_mountable(),
// g_volume_mount(), g_mount_unmount_with_operation() and others.
//
// When necessary, GtkMountOperation shows dialogs to let the user enter
// passwords, ask questions or show processes blocking unmount.
type MountOperation struct {
	_ [0]func() // equal guard
	gio.MountOperation
}

var (
	_ coreglib.Objector = (*MountOperation)(nil)
)

func classInitMountOperationer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

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

// NewMountOperation creates a new GtkMountOperation.
//
// The function takes the following parameters:
//
//    - parent (optional): transient parent of the window, or NULL.
//
// The function returns the following values:
//
//    - mountOperation: new GtkMountOperation.
//
func NewMountOperation(parent *Window) *MountOperation {
	var _args [1]girepository.Argument

	if parent != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	}

	_info := girepository.MustFind("Gtk", "MountOperation")
	_gret := _info.InvokeClassMethod("new_MountOperation", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(parent)

	var _mountOperation *MountOperation // out

	_mountOperation = wrapMountOperation(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _mountOperation
}

// Display gets the display on which windows of the GtkMountOperation will be
// shown.
//
// The function returns the following values:
//
//    - display on which windows of op are shown.
//
func (op *MountOperation) Display() *gdk.Display {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(op).Native()))

	_info := girepository.MustFind("Gtk", "MountOperation")
	_gret := _info.InvokeClassMethod("get_display", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(op)

	var _display *gdk.Display // out

	{
		obj := coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret))))
		_display = &gdk.Display{
			Object: obj,
		}
	}

	return _display
}

// Parent gets the transient parent used by the GtkMountOperation.
//
// The function returns the following values:
//
//    - window: transient parent for windows shown by op.
//
func (op *MountOperation) Parent() *Window {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(op).Native()))

	_info := girepository.MustFind("Gtk", "MountOperation")
	_gret := _info.InvokeClassMethod("get_parent", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(op)

	var _window *Window // out

	_window = wrapWindow(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _window
}

// IsShowing returns whether the GtkMountOperation is currently displaying a
// window.
//
// The function returns the following values:
//
//    - ok: TRUE if op is currently displaying a window.
//
func (op *MountOperation) IsShowing() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(op).Native()))

	_info := girepository.MustFind("Gtk", "MountOperation")
	_gret := _info.InvokeClassMethod("is_showing", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(op)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetDisplay sets the display to show windows of the GtkMountOperation on.
//
// The function takes the following parameters:
//
//    - display: GdkDisplay.
//
func (op *MountOperation) SetDisplay(display *gdk.Display) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(op).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_info := girepository.MustFind("Gtk", "MountOperation")
	_info.InvokeClassMethod("set_display", _args[:], nil)

	runtime.KeepAlive(op)
	runtime.KeepAlive(display)
}

// SetParent sets the transient parent for windows shown by the
// GtkMountOperation.
//
// The function takes the following parameters:
//
//    - parent (optional): transient parent of the window, or NULL.
//
func (op *MountOperation) SetParent(parent *Window) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(op).Native()))
	if parent != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	}

	_info := girepository.MustFind("Gtk", "MountOperation")
	_info.InvokeClassMethod("set_parent", _args[:], nil)

	runtime.KeepAlive(op)
	runtime.KeepAlive(parent)
}
