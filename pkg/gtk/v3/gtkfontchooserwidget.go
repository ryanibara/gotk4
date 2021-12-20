// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"sync"
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
		{T: externglib.Type(C.gtk_font_chooser_widget_get_type()), F: marshalFontChooserWidgetter},
	})
}

// FontChooserWidget widget lists the available fonts, styles and sizes,
// allowing the user to select a font. It is used in the FontChooserDialog
// widget to provide a dialog box for selecting fonts.
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
type FontChooserWidget struct {
	Box

	*externglib.Object
	FontChooser

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ externglib.Objector = (*FontChooserWidget)(nil)
	_ Containerer         = (*FontChooserWidget)(nil)
)

func wrapFontChooserWidget(obj *externglib.Object) *FontChooserWidget {
	return &FontChooserWidget{
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
		FontChooser: FontChooser{
			Object: obj,
		},
	}
}

func marshalFontChooserWidgetter(p uintptr) (interface{}, error) {
	return wrapFontChooserWidget(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFontChooserWidget creates a new FontChooserWidget.
//
// The function returns the following values:
//
//    - fontChooserWidget: new FontChooserWidget.
//
func NewFontChooserWidget() *FontChooserWidget {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_font_chooser_widget_new()

	var _fontChooserWidget *FontChooserWidget // out

	_fontChooserWidget = wrapFontChooserWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _fontChooserWidget
}
