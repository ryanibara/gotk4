// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_color_chooser_widget_get_type()), F: marshalColorChooserWidgetter},
	})
}

// ColorChooserWidget widget lets the user select a color. By default, the
// chooser presents a predefined palette of colors, plus a small number of
// settable custom colors. It is also possible to select a different color with
// the single-color editor. To enter the single-color editing mode, use the
// context menu of any color of the palette, or use the '+' button to add a new
// custom color.
//
// The chooser automatically remembers the last selection, as well as custom
// colors.
//
// To change the initially selected color, use gtk_color_chooser_set_rgba(). To
// get the selected color use gtk_color_chooser_get_rgba().
//
// The ColorChooserWidget is used in the ColorChooserDialog to provide a dialog
// for selecting colors.
//
//
// CSS names
//
// GtkColorChooserWidget has a single CSS node with name colorchooser.
type ColorChooserWidget struct {
	Box

	*externglib.Object
	ColorChooser
}

var (
	_ externglib.Objector = (*ColorChooserWidget)(nil)
	_ Containerer         = (*ColorChooserWidget)(nil)
)

func wrapColorChooserWidget(obj *externglib.Object) *ColorChooserWidget {
	return &ColorChooserWidget{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
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
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
		Object: obj,
		ColorChooser: ColorChooser{
			Object: obj,
		},
	}
}

func marshalColorChooserWidgetter(p uintptr) (interface{}, error) {
	return wrapColorChooserWidget(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewColorChooserWidget creates a new ColorChooserWidget.
//
// The function returns the following values:
//
//    - colorChooserWidget: new ColorChooserWidget.
//
func NewColorChooserWidget() *ColorChooserWidget {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_color_chooser_widget_new()

	var _colorChooserWidget *ColorChooserWidget // out

	_colorChooserWidget = wrapColorChooserWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _colorChooserWidget
}
