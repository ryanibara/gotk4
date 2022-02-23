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

// glib.Type values for gtkfontchooserwidget.go.
var GTypeFontChooserWidget = externglib.Type(C.gtk_font_chooser_widget_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeFontChooserWidget, F: marshalFontChooserWidget},
	})
}

// FontChooserWidget: GtkFontChooserWidget widget lets the user select a font.
//
// It is used in the GtkFontChooserDialog widget to provide a dialog for
// selecting fonts.
//
// To set the font which is initially selected, use gtk.FontChooser.SetFont() or
// gtk.FontChooser.SetFontDesc().
//
// To get the selected font use gtk.FontChooser.GetFont() or
// gtk.FontChooser.GetFontDesc().
//
// To change the text which is shown in the preview area, use
// gtk.FontChooser.SetPreviewText().
//
//
// CSS nodes
//
// GtkFontChooserWidget has a single CSS node with name fontchooser.
type FontChooserWidget struct {
	_ [0]func() // equal guard
	Widget

	*externglib.Object
	FontChooser
}

var (
	_ Widgetter           = (*FontChooserWidget)(nil)
	_ externglib.Objector = (*FontChooserWidget)(nil)
)

func wrapFontChooserWidget(obj *externglib.Object) *FontChooserWidget {
	return &FontChooserWidget{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		FontChooser: FontChooser{
			Object: obj,
		},
	}
}

func marshalFontChooserWidget(p uintptr) (interface{}, error) {
	return wrapFontChooserWidget(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFontChooserWidget creates a new GtkFontChooserWidget.
//
// The function returns the following values:
//
//    - fontChooserWidget: new GtkFontChooserWidget.
//
func NewFontChooserWidget() *FontChooserWidget {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_font_chooser_widget_new()

	var _fontChooserWidget *FontChooserWidget // out

	_fontChooserWidget = wrapFontChooserWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _fontChooserWidget
}
