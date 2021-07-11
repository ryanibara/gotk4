// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
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
		{T: externglib.Type(C.gtk_drop_controller_motion_get_type()), F: marshalDropControllerMotioner},
	})
}

// DropControllerMotioner describes DropControllerMotion's methods.
type DropControllerMotioner interface {
	// ContainsPointer returns if a Drag-and-Drop operation is within the widget
	// @self or one of its children.
	ContainsPointer() bool
	// Drop returns the `GdkDrop` of a current Drag-and-Drop operation over the
	// widget of @self.
	Drop() *gdk.Drop
	// IsPointer returns if a Drag-and-Drop operation is within the widget
	// @self, not one of its children.
	IsPointer() bool
}

// DropControllerMotion: `GtkDropControllerMotion` is an event controller
// tracking the pointer during Drag-and-Drop operations.
//
// It is modeled after [class@Gtk.EventControllerMotion] so if you have used
// that, this should feel really familiar.
//
// This controller is not able to accept drops, use [class@Gtk.DropTarget] for
// that purpose.
type DropControllerMotion struct {
	EventController
}

var (
	_ DropControllerMotioner = (*DropControllerMotion)(nil)
	_ gextras.Nativer        = (*DropControllerMotion)(nil)
)

func wrapDropControllerMotion(obj *externglib.Object) DropControllerMotioner {
	return &DropControllerMotion{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalDropControllerMotioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDropControllerMotion(obj), nil
}

// NewDropControllerMotion creates a new event controller that will handle
// pointer motion events during drag and drop.
func NewDropControllerMotion() *DropControllerMotion {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_drop_controller_motion_new()

	var _dropControllerMotion *DropControllerMotion // out

	_dropControllerMotion = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*DropControllerMotion)

	return _dropControllerMotion
}

// ContainsPointer returns if a Drag-and-Drop operation is within the widget
// @self or one of its children.
func (self *DropControllerMotion) ContainsPointer() bool {
	var _arg0 *C.GtkDropControllerMotion // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkDropControllerMotion)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_controller_motion_contains_pointer(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Drop returns the `GdkDrop` of a current Drag-and-Drop operation over the
// widget of @self.
func (self *DropControllerMotion) Drop() *gdk.Drop {
	var _arg0 *C.GtkDropControllerMotion // out
	var _cret *C.GdkDrop                 // in

	_arg0 = (*C.GtkDropControllerMotion)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_controller_motion_get_drop(_arg0)

	var _drop *gdk.Drop // out

	_drop = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gdk.Drop)

	return _drop
}

// IsPointer returns if a Drag-and-Drop operation is within the widget @self,
// not one of its children.
func (self *DropControllerMotion) IsPointer() bool {
	var _arg0 *C.GtkDropControllerMotion // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkDropControllerMotion)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_controller_motion_is_pointer(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
