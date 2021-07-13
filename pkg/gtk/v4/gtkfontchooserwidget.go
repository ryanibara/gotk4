// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_font_chooser_widget_get_type()), F: marshalFontChooserWidgeter},
	})
}

// FontChooserWidgeter describes FontChooserWidget's methods.
type FontChooserWidgeter interface {
	privateFontChooserWidget()
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
	Widget

	FontChooser
}

var (
	_ FontChooserWidgeter = (*FontChooserWidget)(nil)
	_ gextras.Nativer     = (*FontChooserWidget)(nil)
)

func wrapFontChooserWidget(obj *externglib.Object) *FontChooserWidget {
	return &FontChooserWidget{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
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
		FontChooser: FontChooser{
			Object: obj,
		},
	}
}

func marshalFontChooserWidgeter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFontChooserWidget(obj), nil
}

// NewFontChooserWidget creates a new GtkFontChooserWidget.
func NewFontChooserWidget() *FontChooserWidget {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_font_chooser_widget_new()

	var _fontChooserWidget *FontChooserWidget // out

	_fontChooserWidget = wrapFontChooserWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _fontChooserWidget
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *FontChooserWidget) Native() uintptr {
	return v.Widget.InitiallyUnowned.Object.Native()
}

func (*FontChooserWidget) privateFontChooserWidget() {}
