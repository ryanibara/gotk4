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

// glib.Type values for gtkcolorchooserwidget.go.
var GTypeColorChooserWidget = externglib.Type(C.gtk_color_chooser_widget_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeColorChooserWidget, F: marshalColorChooserWidget},
	})
}

// ColorChooserWidget: GtkColorChooserWidget widget lets the user select a
// color.
//
// By default, the chooser presents a predefined palette of colors, plus a small
// number of settable custom colors. It is also possible to select a different
// color with the single-color editor.
//
// To enter the single-color editing mode, use the context menu of any color of
// the palette, or use the '+' button to add a new custom color.
//
// The chooser automatically remembers the last selection, as well as custom
// colors.
//
// To create a GtkColorChooserWidget, use gtk.ColorChooserWidget.New.
//
// To change the initially selected color, use gtk.ColorChooser.SetRGBA(). To
// get the selected color use gtk.ColorChooser.GetRGBA().
//
// The GtkColorChooserWidget is used in the gtk.ColorChooserDialog to provide a
// dialog for selecting colors.
//
//
// CSS names
//
// GtkColorChooserWidget has a single CSS node with name colorchooser.
type ColorChooserWidget struct {
	_ [0]func() // equal guard
	Widget

	*externglib.Object
	ColorChooser
}

var (
	_ Widgetter           = (*ColorChooserWidget)(nil)
	_ externglib.Objector = (*ColorChooserWidget)(nil)
)

func wrapColorChooserWidget(obj *externglib.Object) *ColorChooserWidget {
	return &ColorChooserWidget{
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
		ColorChooser: ColorChooser{
			Object: obj,
		},
	}
}

func marshalColorChooserWidget(p uintptr) (interface{}, error) {
	return wrapColorChooserWidget(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewColorChooserWidget creates a new GtkColorChooserWidget.
//
// The function returns the following values:
//
//    - colorChooserWidget: new GtkColorChooserWidget.
//
func NewColorChooserWidget() *ColorChooserWidget {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_color_chooser_widget_new()

	var _colorChooserWidget *ColorChooserWidget // out

	_colorChooserWidget = wrapColorChooserWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _colorChooserWidget
}
