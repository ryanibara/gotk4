// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_font_selection_get_type()), F: marshalFontSelectioner},
		{T: externglib.Type(C.gtk_font_selection_dialog_get_type()), F: marshalFontSelectionDialogger},
	})
}

type FontSelection struct {
	Box
}

func wrapFontSelection(obj *externglib.Object) *FontSelection {
	return &FontSelection{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					Object: obj,
				},
			},
			Orientable: Orientable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalFontSelectioner(p uintptr) (interface{}, error) {
	return wrapFontSelection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFontSelection creates a new FontSelection.
//
// Deprecated: Use FontChooserWidget instead.
func NewFontSelection() *FontSelection {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_font_selection_new()

	var _fontSelection *FontSelection // out

	_fontSelection = wrapFontSelection(externglib.Take(unsafe.Pointer(_cret)))

	return _fontSelection
}

// Face gets the FontFace representing the selected font group details (i.e.
// family, slant, weight, width, etc).
//
// Deprecated: Use FontChooser.
func (fontsel *FontSelection) Face() pango.FontFacer {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.PangoFontFace    // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(fontsel.Native()))

	_cret = C.gtk_font_selection_get_face(_arg0)
	runtime.KeepAlive(fontsel)

	var _fontFace pango.FontFacer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type pango.FontFacer is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(pango.FontFacer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not pango.FontFacer")
		}
		_fontFace = rv
	}

	return _fontFace
}

// FaceList: this returns the TreeView which lists all styles available for the
// selected font. For example, “Regular”, “Bold”, etc.
//
// Deprecated: Use FontChooser.
func (fontsel *FontSelection) FaceList() Widgetter {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(fontsel.Native()))

	_cret = C.gtk_font_selection_get_face_list(_arg0)
	runtime.KeepAlive(fontsel)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Widgetter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// Family gets the FontFamily representing the selected font family.
//
// Deprecated: Use FontChooser.
func (fontsel *FontSelection) Family() pango.FontFamilier {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.PangoFontFamily  // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(fontsel.Native()))

	_cret = C.gtk_font_selection_get_family(_arg0)
	runtime.KeepAlive(fontsel)

	var _fontFamily pango.FontFamilier // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type pango.FontFamilier is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(pango.FontFamilier)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not pango.FontFamilier")
		}
		_fontFamily = rv
	}

	return _fontFamily
}

// FamilyList: this returns the TreeView that lists font families, for example,
// “Sans”, “Serif”, etc.
//
// Deprecated: Use FontChooser.
func (fontsel *FontSelection) FamilyList() Widgetter {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(fontsel.Native()))

	_cret = C.gtk_font_selection_get_family_list(_arg0)
	runtime.KeepAlive(fontsel)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Widgetter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// FontName gets the currently-selected font name.
//
// Note that this can be a different string than what you set with
// gtk_font_selection_set_font_name(), as the font selection widget may
// normalize font names and thus return a string with a different structure. For
// example, “Helvetica Italic Bold 12” could be normalized to “Helvetica Bold
// Italic 12”. Use pango_font_description_equal() if you want to compare two
// font descriptions.
//
// Deprecated: Use FontChooser.
func (fontsel *FontSelection) FontName() string {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(fontsel.Native()))

	_cret = C.gtk_font_selection_get_font_name(_arg0)
	runtime.KeepAlive(fontsel)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// PreviewEntry: this returns the Entry used to display the font as a preview.
//
// Deprecated: Use FontChooser.
func (fontsel *FontSelection) PreviewEntry() Widgetter {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(fontsel.Native()))

	_cret = C.gtk_font_selection_get_preview_entry(_arg0)
	runtime.KeepAlive(fontsel)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Widgetter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// PreviewText gets the text displayed in the preview area.
//
// Deprecated: Use FontChooser.
func (fontsel *FontSelection) PreviewText() string {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(fontsel.Native()))

	_cret = C.gtk_font_selection_get_preview_text(_arg0)
	runtime.KeepAlive(fontsel)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Size: selected font size.
//
// Deprecated: Use FontChooser.
func (fontsel *FontSelection) Size() int {
	var _arg0 *C.GtkFontSelection // out
	var _cret C.gint              // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(fontsel.Native()))

	_cret = C.gtk_font_selection_get_size(_arg0)
	runtime.KeepAlive(fontsel)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SizeEntry: this returns the Entry used to allow the user to edit the font
// number manually instead of selecting it from the list of font sizes.
//
// Deprecated: Use FontChooser.
func (fontsel *FontSelection) SizeEntry() Widgetter {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(fontsel.Native()))

	_cret = C.gtk_font_selection_get_size_entry(_arg0)
	runtime.KeepAlive(fontsel)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Widgetter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// SizeList: this returns the TreeView used to list font sizes.
//
// Deprecated: Use FontChooser.
func (fontsel *FontSelection) SizeList() Widgetter {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(fontsel.Native()))

	_cret = C.gtk_font_selection_get_size_list(_arg0)
	runtime.KeepAlive(fontsel)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Widgetter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// SetFontName sets the currently-selected font.
//
// Note that the fontsel needs to know the screen in which it will appear for
// this to work; this can be guaranteed by simply making sure that the fontsel
// is inserted in a toplevel window before you call this function.
//
// Deprecated: Use FontChooser.
//
// The function takes the following parameters:
//
//    - fontname: font name like “Helvetica 12” or “Times Bold 18”.
//
func (fontsel *FontSelection) SetFontName(fontname string) bool {
	var _arg0 *C.GtkFontSelection // out
	var _arg1 *C.gchar            // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(fontsel.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fontname)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_font_selection_set_font_name(_arg0, _arg1)
	runtime.KeepAlive(fontsel)
	runtime.KeepAlive(fontname)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetPreviewText sets the text displayed in the preview area. The text is used
// to show how the selected font looks.
//
// Deprecated: Use FontChooser.
//
// The function takes the following parameters:
//
//    - text to display in the preview area.
//
func (fontsel *FontSelection) SetPreviewText(text string) {
	var _arg0 *C.GtkFontSelection // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(fontsel.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_selection_set_preview_text(_arg0, _arg1)
	runtime.KeepAlive(fontsel)
	runtime.KeepAlive(text)
}

type FontSelectionDialog struct {
	Dialog
}

func wrapFontSelectionDialog(obj *externglib.Object) *FontSelectionDialog {
	return &FontSelectionDialog{
		Dialog: Dialog{
			Window: Window{
				Bin: Bin{
					Container: Container{
						Widget: Widget{
							InitiallyUnowned: externglib.InitiallyUnowned{
								Object: obj,
							},
							ImplementorIface: atk.ImplementorIface{
								Object: obj,
							},
							Buildable: Buildable{
								Object: obj,
							},
							Object: obj,
						},
					},
				},
			},
		},
	}
}

func marshalFontSelectionDialogger(p uintptr) (interface{}, error) {
	return wrapFontSelectionDialog(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFontSelectionDialog creates a new FontSelectionDialog.
//
// Deprecated: Use FontChooserDialog.
//
// The function takes the following parameters:
//
//    - title of the dialog window.
//
func NewFontSelectionDialog(title string) *FontSelectionDialog {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_font_selection_dialog_new(_arg1)
	runtime.KeepAlive(title)

	var _fontSelectionDialog *FontSelectionDialog // out

	_fontSelectionDialog = wrapFontSelectionDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _fontSelectionDialog
}

// CancelButton gets the “Cancel” button.
//
// Deprecated: Use FontChooserDialog.
func (fsd *FontSelectionDialog) CancelButton() Widgetter {
	var _arg0 *C.GtkFontSelectionDialog // out
	var _cret *C.GtkWidget              // in

	_arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(fsd.Native()))

	_cret = C.gtk_font_selection_dialog_get_cancel_button(_arg0)
	runtime.KeepAlive(fsd)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Widgetter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// FontName gets the currently-selected font name.
//
// Note that this can be a different string than what you set with
// gtk_font_selection_dialog_set_font_name(), as the font selection widget may
// normalize font names and thus return a string with a different structure. For
// example, “Helvetica Italic Bold 12” could be normalized to “Helvetica Bold
// Italic 12”. Use pango_font_description_equal() if you want to compare two
// font descriptions.
//
// Deprecated: Use FontChooserDialog.
func (fsd *FontSelectionDialog) FontName() string {
	var _arg0 *C.GtkFontSelectionDialog // out
	var _cret *C.gchar                  // in

	_arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(fsd.Native()))

	_cret = C.gtk_font_selection_dialog_get_font_name(_arg0)
	runtime.KeepAlive(fsd)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FontSelection retrieves the FontSelection widget embedded in the dialog.
//
// Deprecated: Use FontChooserDialog.
func (fsd *FontSelectionDialog) FontSelection() Widgetter {
	var _arg0 *C.GtkFontSelectionDialog // out
	var _cret *C.GtkWidget              // in

	_arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(fsd.Native()))

	_cret = C.gtk_font_selection_dialog_get_font_selection(_arg0)
	runtime.KeepAlive(fsd)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Widgetter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// OKButton gets the “OK” button.
//
// Deprecated: Use FontChooserDialog.
func (fsd *FontSelectionDialog) OKButton() Widgetter {
	var _arg0 *C.GtkFontSelectionDialog // out
	var _cret *C.GtkWidget              // in

	_arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(fsd.Native()))

	_cret = C.gtk_font_selection_dialog_get_ok_button(_arg0)
	runtime.KeepAlive(fsd)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Widgetter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// PreviewText gets the text displayed in the preview area.
//
// Deprecated: Use FontChooserDialog.
func (fsd *FontSelectionDialog) PreviewText() string {
	var _arg0 *C.GtkFontSelectionDialog // out
	var _cret *C.gchar                  // in

	_arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(fsd.Native()))

	_cret = C.gtk_font_selection_dialog_get_preview_text(_arg0)
	runtime.KeepAlive(fsd)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetFontName sets the currently selected font.
//
// Deprecated: Use FontChooserDialog.
//
// The function takes the following parameters:
//
//    - fontname: font name like “Helvetica 12” or “Times Bold 18”.
//
func (fsd *FontSelectionDialog) SetFontName(fontname string) bool {
	var _arg0 *C.GtkFontSelectionDialog // out
	var _arg1 *C.gchar                  // out
	var _cret C.gboolean                // in

	_arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(fsd.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fontname)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_font_selection_dialog_set_font_name(_arg0, _arg1)
	runtime.KeepAlive(fsd)
	runtime.KeepAlive(fontname)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetPreviewText sets the text displayed in the preview area.
//
// Deprecated: Use FontChooserDialog.
//
// The function takes the following parameters:
//
//    - text to display in the preview area.
//
func (fsd *FontSelectionDialog) SetPreviewText(text string) {
	var _arg0 *C.GtkFontSelectionDialog // out
	var _arg1 *C.gchar                  // out

	_arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(fsd.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_selection_dialog_set_preview_text(_arg0, _arg1)
	runtime.KeepAlive(fsd)
	runtime.KeepAlive(text)
}
