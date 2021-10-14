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
		{T: externglib.Type(C.gtk_vscrollbar_get_type()), F: marshalVScrollbarrer},
	})
}

// VScrollbar widget is a widget arranged vertically creating a scrollbar. See
// Scrollbar for details on scrollbars. Adjustment pointers may be added to
// handle the adjustment of the scrollbar or it may be left NULL in which case
// one will be created for you. See Scrollbar for a description of what the
// fields in an adjustment represent for a scrollbar.
//
// GtkVScrollbar has been deprecated, use Scrollbar instead.
type VScrollbar struct {
	Scrollbar
}

func wrapVScrollbar(obj *externglib.Object) *VScrollbar {
	return &VScrollbar{
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

func marshalVScrollbarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapVScrollbar(obj), nil
}

// NewVScrollbar creates a new vertical scrollbar.
//
// Deprecated: Use gtk_scrollbar_new() with GTK_ORIENTATION_VERTICAL instead.
//
// The function takes the following parameters:
//
//    - adjustment to use, or NULL to create a new adjustment.
//
func NewVScrollbar(adjustment *Adjustment) *VScrollbar {
	var _arg1 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	if adjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	}

	_cret = C.gtk_vscrollbar_new(_arg1)
	runtime.KeepAlive(adjustment)

	var _vScrollbar *VScrollbar // out

	_vScrollbar = wrapVScrollbar(externglib.Take(unsafe.Pointer(_cret)))

	return _vScrollbar
}

func (*VScrollbar) privateVScrollbar() {}
