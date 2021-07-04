// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_bin_layout_get_type()), F: marshalBinLayout},
	})
}

// BinLayout: `GtkBinLayout` is a `GtkLayoutManager` subclass useful for create
// "bins" of widgets.
//
// `GtkBinLayout` will stack each child of a widget on top of each other, using
// the [property@Gtk.Widget:hexpand], [property@Gtk.Widget:vexpand],
// [property@Gtk.Widget:halign], and [property@Gtk.Widget:valign] properties of
// each child to determine where they should be positioned.
type BinLayout interface {
	LayoutManager
}

// binLayout implements the BinLayout class.
type binLayout struct {
	LayoutManager
}

// WrapBinLayout wraps a GObject to the right type. It is
// primarily used internally.
func WrapBinLayout(obj *externglib.Object) BinLayout {
	return binLayout{
		LayoutManager: WrapLayoutManager(obj),
	}
}

func marshalBinLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBinLayout(obj), nil
}

func NewBinLayout() BinLayout {
	var _cret *C.GtkLayoutManager // in

	_cret = C.gtk_bin_layout_new()

	var _binLayout BinLayout // out

	_binLayout = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(BinLayout)

	return _binLayout
}
