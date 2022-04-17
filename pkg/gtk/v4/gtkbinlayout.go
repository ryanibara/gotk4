// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// glib.Type values for gtkbinlayout.go.
var GTypeBinLayout = externglib.Type(C.gtk_bin_layout_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeBinLayout, F: marshalBinLayout},
	})
}

// BinLayoutOverrider contains methods that are overridable.
type BinLayoutOverrider interface {
	externglib.Objector
}

// BinLayout: GtkBinLayout is a GtkLayoutManager subclass useful for create
// "bins" of widgets.
//
// GtkBinLayout will stack each child of a widget on top of each other, using
// the gtk.Widget:hexpand, gtk.Widget:vexpand, gtk.Widget:halign, and
// gtk.Widget:valign properties of each child to determine where they should be
// positioned.
type BinLayout struct {
	_ [0]func() // equal guard
	LayoutManager
}

var (
	_ LayoutManagerer = (*BinLayout)(nil)
)

func classInitBinLayouter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapBinLayout(obj *externglib.Object) *BinLayout {
	return &BinLayout{
		LayoutManager: LayoutManager{
			Object: obj,
		},
	}
}

func marshalBinLayout(p uintptr) (interface{}, error) {
	return wrapBinLayout(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBinLayout creates a new GtkBinLayout instance.
//
// The function returns the following values:
//
//    - binLayout: newly created GtkBinLayout.
//
func NewBinLayout() *BinLayout {
	var _cret *C.GtkLayoutManager // in

	_cret = C.gtk_bin_layout_new()

	var _binLayout *BinLayout // out

	_binLayout = wrapBinLayout(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _binLayout
}
