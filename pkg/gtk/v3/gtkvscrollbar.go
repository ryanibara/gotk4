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
//
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

// VScrollbarrer describes VScrollbar's methods.
type VScrollbarrer interface {
	privateVScrollbar()
}

// VScrollbar widget is a widget arranged vertically creating a scrollbar. See
// Scrollbar for details on scrollbars. Adjustment pointers may be added to
// handle the adjustment of the scrollbar or it may be left nil in which case
// one will be created for you. See Scrollbar for a description of what the
// fields in an adjustment represent for a scrollbar.
//
// GtkVScrollbar has been deprecated, use Scrollbar instead.
type VScrollbar struct {
	Scrollbar
}

var (
	_ VScrollbarrer   = (*VScrollbar)(nil)
	_ gextras.Nativer = (*VScrollbar)(nil)
)

func wrapVScrollbar(obj *externglib.Object) VScrollbarrer {
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
				},
				Orientable: Orientable{
					Object: obj,
				},
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
func NewVScrollbar(adjustment Adjustmenter) *VScrollbar {
	var _arg1 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer((adjustment).(gextras.Nativer).Native()))

	_cret = C.gtk_vscrollbar_new(_arg1)

	var _vScrollbar *VScrollbar // out

	_vScrollbar = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*VScrollbar)

	return _vScrollbar
}

func (*VScrollbar) privateVScrollbar() {}
