// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_vscrollbar_get_type()), F: marshalVScrollbar},
	})
}

// VScrollbar: the VScrollbar widget is a widget arranged vertically creating a
// scrollbar. See Scrollbar for details on scrollbars. Adjustment pointers may
// be added to handle the adjustment of the scrollbar or it may be left nil in
// which case one will be created for you. See Scrollbar for a description of
// what the fields in an adjustment represent for a scrollbar.
//
// GtkVScrollbar has been deprecated, use Scrollbar instead.
type VScrollbar interface {
	Scrollbar
	Buildable
	Orientable
}

// vScrollbar implements the VScrollbar interface.
type vScrollbar struct {
	Scrollbar
	Buildable
	Orientable
}

var _ VScrollbar = (*vScrollbar)(nil)

// WrapVScrollbar wraps a GObject to the right type. It is
// primarily used internally.
func WrapVScrollbar(obj *externglib.Object) VScrollbar {
	return VScrollbar{
		Scrollbar:  WrapScrollbar(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalVScrollbar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVScrollbar(obj), nil
}

// NewVScrollbar constructs a class VScrollbar.
func NewVScrollbar(adjustment Adjustment) VScrollbar {
	var arg1 *C.GtkAdjustment

	arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	var cret C.GtkVScrollbar
	var ret1 VScrollbar

	cret = C.gtk_vscrollbar_new(adjustment)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(VScrollbar)

	return ret1
}
