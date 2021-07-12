// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_hscrollbar_get_type()), F: marshalHScrollbarer},
	})
}

// HScrollbarer describes HScrollbar's methods.
type HScrollbarer interface {
	privateHScrollbar()
}

// HScrollbar widget is a widget arranged horizontally creating a scrollbar. See
// Scrollbar for details on scrollbars. Adjustment pointers may be added to
// handle the adjustment of the scrollbar or it may be left nil in which case
// one will be created for you. See Scrollbar for a description of what the
// fields in an adjustment represent for a scrollbar.
//
// GtkHScrollbar has been deprecated, use Scrollbar instead.
type HScrollbar struct {
	Scrollbar
}

var (
	_ HScrollbarer    = (*HScrollbar)(nil)
	_ gextras.Nativer = (*HScrollbar)(nil)
)

func wrapHScrollbar(obj *externglib.Object) *HScrollbar {
	return &HScrollbar{
		Scrollbar: Scrollbar{
			Range: Range{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
				Orientable: Orientable{
					Object: obj,
				},
			},
		},
	}
}

func marshalHScrollbarer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapHScrollbar(obj), nil
}

// NewHScrollbar creates a new horizontal scrollbar.
//
// Deprecated: Use gtk_scrollbar_new() with GTK_ORIENTATION_HORIZONTAL instead.
func NewHScrollbar(adjustment Adjustmenter) *HScrollbar {
	var _arg1 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer((adjustment).(gextras.Nativer).Native()))

	_cret = C.gtk_hscrollbar_new(_arg1)

	var _hScrollbar *HScrollbar // out

	_hScrollbar = wrapHScrollbar(externglib.Take(unsafe.Pointer(_cret)))

	return _hScrollbar
}

func (*HScrollbar) privateHScrollbar() {}
