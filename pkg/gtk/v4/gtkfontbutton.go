// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_font_button_get_type()), F: marshalFontButtonner},
	})
}

// FontButton: GtkFontButton allows to open a font chooser dialog to change the
// font.
//
// !An example GtkFontButton (font-button.png)
//
// It is suitable widget for selecting a font in a preference dialog.
//
// CSS nodes
//
//    fontbutton
//    ╰── button.font
//        ╰── [content]
//
//
// GtkFontButton has a single CSS node with name fontbutton which contains a
// button node with the .font style class.
type FontButton struct {
	Widget

	*externglib.Object
	FontChooser
}

var (
	_ Widgetter           = (*FontButton)(nil)
	_ externglib.Objector = (*FontButton)(nil)
)

func wrapFontButton(obj *externglib.Object) *FontButton {
	return &FontButton{
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

func marshalFontButtonner(p uintptr) (interface{}, error) {
	return wrapFontButton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectFontSet: emitted when the user selects a font.
//
// When handling this signal, use gtk.FontChooser.GetFont() to find out which
// font was just selected.
//
// Note that this signal is only emitted when the user changes the font. If you
// need to react to programmatic font changes as well, use the notify::font
// signal.
func (fontButton *FontButton) ConnectFontSet(f func()) externglib.SignalHandle {
	return fontButton.Connect("font-set", f)
}

// NewFontButton creates a new font picker widget.
//
// The function returns the following values:
//
//    - fontButton: new font picker widget.
//
func NewFontButton() *FontButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_font_button_new()

	var _fontButton *FontButton // out

	_fontButton = wrapFontButton(externglib.Take(unsafe.Pointer(_cret)))

	return _fontButton
}

// NewFontButtonWithFont creates a new font picker widget showing the given
// font.
//
// The function takes the following parameters:
//
//    - fontname: name of font to display in font chooser dialog.
//
// The function returns the following values:
//
//    - fontButton: new font picker widget.
//
func NewFontButtonWithFont(fontname string) *FontButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(fontname)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_font_button_new_with_font(_arg1)
	runtime.KeepAlive(fontname)

	var _fontButton *FontButton // out

	_fontButton = wrapFontButton(externglib.Take(unsafe.Pointer(_cret)))

	return _fontButton
}

// Modal gets whether the dialog is modal.
//
// The function returns the following values:
//
//    - ok: TRUE if the dialog is modal.
//
func (fontButton *FontButton) Modal() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))

	_cret = C.gtk_font_button_get_modal(_arg0)
	runtime.KeepAlive(fontButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Title retrieves the title of the font chooser dialog.
//
// The function returns the following values:
//
//    - utf8: internal copy of the title string which must not be freed.
//
func (fontButton *FontButton) Title() string {
	var _arg0 *C.GtkFontButton // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))

	_cret = C.gtk_font_button_get_title(_arg0)
	runtime.KeepAlive(fontButton)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UseFont returns whether the selected font is used in the label.
//
// The function returns the following values:
//
//    - ok: whether the selected font is used in the label.
//
func (fontButton *FontButton) UseFont() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))

	_cret = C.gtk_font_button_get_use_font(_arg0)
	runtime.KeepAlive(fontButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UseSize returns whether the selected size is used in the label.
//
// The function returns the following values:
//
//    - ok: whether the selected size is used in the label.
//
func (fontButton *FontButton) UseSize() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))

	_cret = C.gtk_font_button_get_use_size(_arg0)
	runtime.KeepAlive(fontButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetModal sets whether the dialog should be modal.
//
// The function takes the following parameters:
//
//    - modal: TRUE to make the dialog modal.
//
func (fontButton *FontButton) SetModal(modal bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_modal(_arg0, _arg1)
	runtime.KeepAlive(fontButton)
	runtime.KeepAlive(modal)
}

// SetTitle sets the title for the font chooser dialog.
//
// The function takes the following parameters:
//
//    - title: string containing the font chooser dialog title.
//
func (fontButton *FontButton) SetTitle(title string) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_button_set_title(_arg0, _arg1)
	runtime.KeepAlive(fontButton)
	runtime.KeepAlive(title)
}

// SetUseFont: if use_font is TRUE, the font name will be written using the
// selected font.
//
// The function takes the following parameters:
//
//    - useFont: if TRUE, font name will be written using font chosen.
//
func (fontButton *FontButton) SetUseFont(useFont bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))
	if useFont {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_use_font(_arg0, _arg1)
	runtime.KeepAlive(fontButton)
	runtime.KeepAlive(useFont)
}

// SetUseSize: if use_size is TRUE, the font name will be written using the
// selected size.
//
// The function takes the following parameters:
//
//    - useSize: if TRUE, font name will be written using the selected size.
//
func (fontButton *FontButton) SetUseSize(useSize bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))
	if useSize {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_use_size(_arg0, _arg1)
	runtime.KeepAlive(fontButton)
	runtime.KeepAlive(useSize)
}
