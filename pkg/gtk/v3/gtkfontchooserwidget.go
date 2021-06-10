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
		{T: externglib.Type(C.gtk_font_chooser_widget_get_type()), F: marshalFontChooserWidget},
	})
}

// FontChooserWidget: the FontChooserWidget widget lists the available fonts,
// styles and sizes, allowing the user to select a font. It is used in the
// FontChooserDialog widget to provide a dialog box for selecting fonts.
//
// To set the font which is initially selected, use gtk_font_chooser_set_font()
// or gtk_font_chooser_set_font_desc().
//
// To get the selected font use gtk_font_chooser_get_font() or
// gtk_font_chooser_get_font_desc().
//
// To change the text which is shown in the preview area, use
// gtk_font_chooser_set_preview_text().
//
//
// CSS nodes
//
// GtkFontChooserWidget has a single CSS node with name fontchooser.
type FontChooserWidget interface {
	Box
	Buildable
	FontChooser
	Orientable
}

// fontChooserWidget implements the FontChooserWidget interface.
type fontChooserWidget struct {
	Box
	Buildable
	FontChooser
	Orientable
}

var _ FontChooserWidget = (*fontChooserWidget)(nil)

// WrapFontChooserWidget wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontChooserWidget(obj *externglib.Object) FontChooserWidget {
	return FontChooserWidget{
		Box:         WrapBox(obj),
		Buildable:   WrapBuildable(obj),
		FontChooser: WrapFontChooser(obj),
		Orientable:  WrapOrientable(obj),
	}
}

func marshalFontChooserWidget(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontChooserWidget(obj), nil
}
