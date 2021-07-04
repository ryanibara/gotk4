// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_mount_operation_get_type()), F: marshalMountOperation},
	})
}

// MountOperation: `GtkMountOperation` is an implementation of
// `GMountOperation`.
//
// The functions and objects described here make working with GTK and GIO more
// convenient.
//
// `GtkMountOperation` is needed when mounting volumes: It is an implementation
// of `GMountOperation` that can be used with GIO functions for mounting volumes
// such as g_file_mount_enclosing_volume(), g_file_mount_mountable(),
// g_volume_mount(), g_mount_unmount_with_operation() and others.
//
// When necessary, `GtkMountOperation` shows dialogs to let the user enter
// passwords, ask questions or show processes blocking unmount.
type MountOperation interface {
	gio.MountOperation

	Display() gdk.Display

	Parent() Window

	IsShowingMountOperation() bool

	SetDisplayMountOperation(display gdk.Display)

	SetParentMountOperation(parent Window)
}

// mountOperation implements the MountOperation class.
type mountOperation struct {
	gio.MountOperation
}

// WrapMountOperation wraps a GObject to the right type. It is
// primarily used internally.
func WrapMountOperation(obj *externglib.Object) MountOperation {
	return mountOperation{
		MountOperation: gio.WrapMountOperation(obj),
	}
}

func marshalMountOperation(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMountOperation(obj), nil
}

func NewMountOperation(parent Window) MountOperation {
	var _arg1 *C.GtkWindow       // out
	var _cret *C.GMountOperation // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	_cret = C.gtk_mount_operation_new(_arg1)

	var _mountOperation MountOperation // out

	_mountOperation = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(MountOperation)

	return _mountOperation
}

func (o mountOperation) Display() gdk.Display {
	var _arg0 *C.GtkMountOperation // out
	var _cret *C.GdkDisplay        // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_mount_operation_get_display(_arg0)

	var _display gdk.Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Display)

	return _display
}

func (o mountOperation) Parent() Window {
	var _arg0 *C.GtkMountOperation // out
	var _cret *C.GtkWindow         // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_mount_operation_get_parent(_arg0)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _window
}

func (o mountOperation) IsShowingMountOperation() bool {
	var _arg0 *C.GtkMountOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_mount_operation_is_showing(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (o mountOperation) SetDisplayMountOperation(display gdk.Display) {
	var _arg0 *C.GtkMountOperation // out
	var _arg1 *C.GdkDisplay        // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gtk_mount_operation_set_display(_arg0, _arg1)
}

func (o mountOperation) SetParentMountOperation(parent Window) {
	var _arg0 *C.GtkMountOperation // out
	var _arg1 *C.GtkWindow         // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	C.gtk_mount_operation_set_parent(_arg0, _arg1)
}
