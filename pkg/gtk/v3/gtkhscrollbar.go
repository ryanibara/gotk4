// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
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
		{T: externglib.Type(C.gtk_hscrollbar_get_type()), F: marshalHScrollbarrer},
	})
}

// HScrollbar widget is a widget arranged horizontally creating a scrollbar. See
// Scrollbar for details on scrollbars. Adjustment pointers may be added to
// handle the adjustment of the scrollbar or it may be left NULL in which case
// one will be created for you. See Scrollbar for a description of what the
// fields in an adjustment represent for a scrollbar.
//
// GtkHScrollbar has been deprecated, use Scrollbar instead.
type HScrollbar struct {
	Scrollbar
}

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
					Object: obj,
				},
				Orientable: Orientable{
					Object: obj,
				},
				Object: obj,
			},
		},
	}
}

func marshalHScrollbarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapHScrollbar(obj), nil
}

// NewHScrollbar creates a new horizontal scrollbar.
//
// Deprecated: Use gtk_scrollbar_new() with GTK_ORIENTATION_HORIZONTAL instead.
//
// The function takes the following parameters:
//
//    - adjustment to use, or NULL to create a new adjustment.
//
func NewHScrollbar(adjustment *Adjustment) *HScrollbar {
	var _arg1 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	if adjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	}

	_cret = C.gtk_hscrollbar_new(_arg1)
	runtime.KeepAlive(adjustment)

	var _hScrollbar *HScrollbar // out

	_hScrollbar = wrapHScrollbar(externglib.Take(unsafe.Pointer(_cret)))

	return _hScrollbar
}

func (*HScrollbar) privateHScrollbar() {}
