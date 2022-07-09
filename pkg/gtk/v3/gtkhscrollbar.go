// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeHScrollbar returns the GType for the type HScrollbar.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeHScrollbar() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "HScrollbar").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalHScrollbar)
	return gtype
}

// HScrollbarOverrider contains methods that are overridable.
type HScrollbarOverrider interface {
}

// HScrollbar widget is a widget arranged horizontally creating a scrollbar. See
// Scrollbar for details on scrollbars. Adjustment pointers may be added to
// handle the adjustment of the scrollbar or it may be left NULL in which case
// one will be created for you. See Scrollbar for a description of what the
// fields in an adjustment represent for a scrollbar.
//
// GtkHScrollbar has been deprecated, use Scrollbar instead.
type HScrollbar struct {
	_ [0]func() // equal guard
	Scrollbar
}

var (
	_ Ranger = (*HScrollbar)(nil)
)

func classInitHScrollbarrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapHScrollbar(obj *coreglib.Object) *HScrollbar {
	return &HScrollbar{
		Scrollbar: Scrollbar{
			Range: Range{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
				Object: obj,
				Orientable: Orientable{
					Object: obj,
				},
			},
		},
	}
}

func marshalHScrollbar(p uintptr) (interface{}, error) {
	return wrapHScrollbar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewHScrollbar creates a new horizontal scrollbar.
//
// Deprecated: Use gtk_scrollbar_new() with GTK_ORIENTATION_HORIZONTAL instead.
//
// The function takes the following parameters:
//
//    - adjustment (optional) to use, or NULL to create a new adjustment.
//
// The function returns the following values:
//
//    - hScrollbar: new HScrollbar.
//
func NewHScrollbar(adjustment *Adjustment) *HScrollbar {
	var _args [1]girepository.Argument

	if adjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}

	_info := girepository.MustFind("Gtk", "HScrollbar")
	_gret := _info.InvokeClassMethod("new_HScrollbar", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(adjustment)

	var _hScrollbar *HScrollbar // out

	_hScrollbar = wrapHScrollbar(coreglib.Take(unsafe.Pointer(_cret)))

	return _hScrollbar
}
