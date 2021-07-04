// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_mount_operation_get_type()), F: marshalMountOperation},
	})
}

// MountOperation: this should not be accessed directly. Use the accessor
// functions below.
type MountOperation interface {
	gio.MountOperation

	Parent() Window

	Screen() gdk.Screen

	IsShowingMountOperation() bool

	SetParentMountOperation(parent Window)

	SetScreenMountOperation(screen gdk.Screen)
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

func (o mountOperation) Parent() Window {
	var _arg0 *C.GtkMountOperation // out
	var _cret *C.GtkWindow         // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_mount_operation_get_parent(_arg0)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _window
}

func (o mountOperation) Screen() gdk.Screen {
	var _arg0 *C.GtkMountOperation // out
	var _cret *C.GdkScreen         // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_mount_operation_get_screen(_arg0)

	var _screen gdk.Screen // out

	_screen = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Screen)

	return _screen
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

func (o mountOperation) SetParentMountOperation(parent Window) {
	var _arg0 *C.GtkMountOperation // out
	var _arg1 *C.GtkWindow         // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	C.gtk_mount_operation_set_parent(_arg0, _arg1)
}

func (o mountOperation) SetScreenMountOperation(screen gdk.Screen) {
	var _arg0 *C.GtkMountOperation // out
	var _arg1 *C.GdkScreen         // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	C.gtk_mount_operation_set_screen(_arg0, _arg1)
}
