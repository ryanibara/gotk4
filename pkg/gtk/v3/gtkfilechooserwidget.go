// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_file_chooser_widget_get_type()), F: marshalFileChooserWidget},
	})
}

// FileChooserWidget is a widget for choosing files. It exposes the FileChooser
// interface, and you should use the methods of this interface to interact with
// the widget.
//
//
// CSS nodes
//
// GtkFileChooserWidget has a single CSS node with name filechooser.
type FileChooserWidget interface {
	Box
	Buildable
	FileChooser
	Orientable
}

// fileChooserWidget implements the FileChooserWidget interface.
type fileChooserWidget struct {
	Box
	Buildable
	FileChooser
	Orientable
}

var _ FileChooserWidget = (*fileChooserWidget)(nil)

// WrapFileChooserWidget wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileChooserWidget(obj *externglib.Object) FileChooserWidget {
	return FileChooserWidget{
		Box:         WrapBox(obj),
		Buildable:   WrapBuildable(obj),
		FileChooser: WrapFileChooser(obj),
		Orientable:  WrapOrientable(obj),
	}
}

func marshalFileChooserWidget(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileChooserWidget(obj), nil
}
