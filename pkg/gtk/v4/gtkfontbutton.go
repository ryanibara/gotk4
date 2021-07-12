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
		{T: externglib.Type(C.gtk_font_button_get_type()), F: marshalFontButtoner},
	})
}

// FontButtoner describes FontButton's methods.
type FontButtoner interface {
	// Modal gets whether the dialog is modal.
	Modal() bool
	// Title retrieves the title of the font chooser dialog.
	Title() string
	// UseFont returns whether the selected font is used in the label.
	UseFont() bool
	// UseSize returns whether the selected size is used in the label.
	UseSize() bool
	// SetModal sets whether the dialog should be modal.
	SetModal(modal bool)
	// SetTitle sets the title for the font chooser dialog.
	SetTitle(title string)
	// SetUseFont: if @use_font is true, the font name will be written using the
	// selected font.
	SetUseFont(useFont bool)
	// SetUseSize: if @use_size is true, the font name will be written using the
	// selected size.
	SetUseSize(useSize bool)
}

// FontButton: `GtkFontButton` allows to open a font chooser dialog to change
// the font.
//
// !An example GtkFontButton (font-button.png)
//
// It is suitable widget for selecting a font in a preference dialog.
//
//
// CSS nodes
//
// “` fontbutton ╰── button.font ╰── [content] “`
//
// `GtkFontButton` has a single CSS node with name fontbutton which contains a
// button node with the .font style class.
type FontButton struct {
	Widget

	FontChooser
}

var (
	_ FontButtoner    = (*FontButton)(nil)
	_ gextras.Nativer = (*FontButton)(nil)
)

func wrapFontButton(obj *externglib.Object) *FontButton {
	return &FontButton{
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

func marshalFontButtoner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFontButton(obj), nil
}

// NewFontButton creates a new font picker widget.
func NewFontButton() *FontButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_font_button_new()

	var _fontButton *FontButton // out

	_fontButton = wrapFontButton(externglib.Take(unsafe.Pointer(_cret)))

	return _fontButton
}

// NewFontButtonWithFont creates a new font picker widget showing the given
// font.
func NewFontButtonWithFont(fontname string) *FontButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(fontname)))

	_cret = C.gtk_font_button_new_with_font(_arg1)

	var _fontButton *FontButton // out

	_fontButton = wrapFontButton(externglib.Take(unsafe.Pointer(_cret)))

	return _fontButton
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *FontButton) Native() uintptr {
	return v.Widget.InitiallyUnowned.Object.Native()
}

// Modal gets whether the dialog is modal.
func (fontButton *FontButton) Modal() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))

	_cret = C.gtk_font_button_get_modal(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Title retrieves the title of the font chooser dialog.
func (fontButton *FontButton) Title() string {
	var _arg0 *C.GtkFontButton // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))

	_cret = C.gtk_font_button_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UseFont returns whether the selected font is used in the label.
func (fontButton *FontButton) UseFont() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))

	_cret = C.gtk_font_button_get_use_font(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UseSize returns whether the selected size is used in the label.
func (fontButton *FontButton) UseSize() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))

	_cret = C.gtk_font_button_get_use_size(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetModal sets whether the dialog should be modal.
func (fontButton *FontButton) SetModal(modal bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_modal(_arg0, _arg1)
}

// SetTitle sets the title for the font chooser dialog.
func (fontButton *FontButton) SetTitle(title string) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))

	C.gtk_font_button_set_title(_arg0, _arg1)
}

// SetUseFont: if @use_font is true, the font name will be written using the
// selected font.
func (fontButton *FontButton) SetUseFont(useFont bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))
	if useFont {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_use_font(_arg0, _arg1)
}

// SetUseSize: if @use_size is true, the font name will be written using the
// selected size.
func (fontButton *FontButton) SetUseSize(useSize bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))
	if useSize {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_use_size(_arg0, _arg1)
}
