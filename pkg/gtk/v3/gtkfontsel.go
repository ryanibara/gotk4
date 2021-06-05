// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_font_selection_get_type()), F: marshalFontSelection},
		{T: externglib.Type(C.gtk_font_selection_dialog_get_type()), F: marshalFontSelectionDialog},
	})
}

type FontSelection interface {
	Box
	Buildable
	Orientable

	// Face gets the FontFace representing the selected font group details (i.e.
	// family, slant, weight, width, etc).
	Face() pango.FontFace
	// FaceList: this returns the TreeView which lists all styles available for
	// the selected font. For example, “Regular”, “Bold”, etc.
	FaceList() Widget
	// Family gets the FontFamily representing the selected font family.
	Family() pango.FontFamily
	// FamilyList: this returns the TreeView that lists font families, for
	// example, “Sans”, “Serif”, etc.
	FamilyList() Widget
	// FontName gets the currently-selected font name.
	//
	// Note that this can be a different string than what you set with
	// gtk_font_selection_set_font_name(), as the font selection widget may
	// normalize font names and thus return a string with a different structure.
	// For example, “Helvetica Italic Bold 12” could be normalized to “Helvetica
	// Bold Italic 12”. Use pango_font_description_equal() if you want to
	// compare two font descriptions.
	FontName() string
	// PreviewEntry: this returns the Entry used to display the font as a
	// preview.
	PreviewEntry() Widget
	// PreviewText gets the text displayed in the preview area.
	PreviewText() string
	// Size: the selected font size.
	Size() int
	// SizeEntry: this returns the Entry used to allow the user to edit the font
	// number manually instead of selecting it from the list of font sizes.
	SizeEntry() Widget
	// SizeList: this returns the TreeView used to list font sizes.
	SizeList() Widget
	// SetFontName sets the currently-selected font.
	//
	// Note that the @fontsel needs to know the screen in which it will appear
	// for this to work; this can be guaranteed by simply making sure that the
	// @fontsel is inserted in a toplevel window before you call this function.
	SetFontName(fontname string) bool
	// SetPreviewText sets the text displayed in the preview area. The @text is
	// used to show how the selected font looks.
	SetPreviewText(text string)
}

// fontSelection implements the FontSelection interface.
type fontSelection struct {
	Box
	Buildable
	Orientable
}

var _ FontSelection = (*fontSelection)(nil)

// WrapFontSelection wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontSelection(obj *externglib.Object) FontSelection {
	return FontSelection{
		Box:        WrapBox(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalFontSelection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontSelection(obj), nil
}

// NewFontSelection constructs a class FontSelection.
func NewFontSelection() FontSelection {
	var cret C.GtkFontSelection
	var ret1 FontSelection

	cret = C.gtk_font_selection_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(FontSelection)

	return ret1
}

// Face gets the FontFace representing the selected font group details (i.e.
// family, slant, weight, width, etc).
func (f fontSelection) Face() pango.FontFace {
	var arg0 *C.GtkFontSelection

	arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	var cret *C.PangoFontFace
	var ret1 pango.FontFace

	cret = C.gtk_font_selection_get_face(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(pango.FontFace)

	return ret1
}

// FaceList: this returns the TreeView which lists all styles available for
// the selected font. For example, “Regular”, “Bold”, etc.
func (f fontSelection) FaceList() Widget {
	var arg0 *C.GtkFontSelection

	arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_font_selection_get_face_list(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// Family gets the FontFamily representing the selected font family.
func (f fontSelection) Family() pango.FontFamily {
	var arg0 *C.GtkFontSelection

	arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	var cret *C.PangoFontFamily
	var ret1 pango.FontFamily

	cret = C.gtk_font_selection_get_family(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(pango.FontFamily)

	return ret1
}

// FamilyList: this returns the TreeView that lists font families, for
// example, “Sans”, “Serif”, etc.
func (f fontSelection) FamilyList() Widget {
	var arg0 *C.GtkFontSelection

	arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_font_selection_get_family_list(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// FontName gets the currently-selected font name.
//
// Note that this can be a different string than what you set with
// gtk_font_selection_set_font_name(), as the font selection widget may
// normalize font names and thus return a string with a different structure.
// For example, “Helvetica Italic Bold 12” could be normalized to “Helvetica
// Bold Italic 12”. Use pango_font_description_equal() if you want to
// compare two font descriptions.
func (f fontSelection) FontName() string {
	var arg0 *C.GtkFontSelection

	arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_font_selection_get_font_name(arg0)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// PreviewEntry: this returns the Entry used to display the font as a
// preview.
func (f fontSelection) PreviewEntry() Widget {
	var arg0 *C.GtkFontSelection

	arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_font_selection_get_preview_entry(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// PreviewText gets the text displayed in the preview area.
func (f fontSelection) PreviewText() string {
	var arg0 *C.GtkFontSelection

	arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_font_selection_get_preview_text(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Size: the selected font size.
func (f fontSelection) Size() int {
	var arg0 *C.GtkFontSelection

	arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	var cret C.gint
	var ret1 int

	cret = C.gtk_font_selection_get_size(arg0)

	ret1 = C.gint(cret)

	return ret1
}

// SizeEntry: this returns the Entry used to allow the user to edit the font
// number manually instead of selecting it from the list of font sizes.
func (f fontSelection) SizeEntry() Widget {
	var arg0 *C.GtkFontSelection

	arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_font_selection_get_size_entry(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// SizeList: this returns the TreeView used to list font sizes.
func (f fontSelection) SizeList() Widget {
	var arg0 *C.GtkFontSelection

	arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_font_selection_get_size_list(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// SetFontName sets the currently-selected font.
//
// Note that the @fontsel needs to know the screen in which it will appear
// for this to work; this can be guaranteed by simply making sure that the
// @fontsel is inserted in a toplevel window before you call this function.
func (f fontSelection) SetFontName(fontname string) bool {
	var arg0 *C.GtkFontSelection
	var arg1 *C.gchar

	arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))
	arg1 = (*C.gchar)(C.CString(fontname))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_font_selection_set_font_name(arg0, fontname)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// SetPreviewText sets the text displayed in the preview area. The @text is
// used to show how the selected font looks.
func (f fontSelection) SetPreviewText(text string) {
	var arg0 *C.GtkFontSelection
	var arg1 *C.gchar

	arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))
	arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_font_selection_set_preview_text(arg0, text)
}

type FontSelectionDialog interface {
	Dialog
	Buildable

	// CancelButton gets the “Cancel” button.
	CancelButton() Widget
	// FontName gets the currently-selected font name.
	//
	// Note that this can be a different string than what you set with
	// gtk_font_selection_dialog_set_font_name(), as the font selection widget
	// may normalize font names and thus return a string with a different
	// structure. For example, “Helvetica Italic Bold 12” could be normalized to
	// “Helvetica Bold Italic 12”. Use pango_font_description_equal() if you
	// want to compare two font descriptions.
	FontName() string
	// FontSelection retrieves the FontSelection widget embedded in the dialog.
	FontSelection() Widget
	// OkButton gets the “OK” button.
	OkButton() Widget
	// PreviewText gets the text displayed in the preview area.
	PreviewText() string
	// SetFontName sets the currently selected font.
	SetFontName(fontname string) bool
	// SetPreviewText sets the text displayed in the preview area.
	SetPreviewText(text string)
}

// fontSelectionDialog implements the FontSelectionDialog interface.
type fontSelectionDialog struct {
	Dialog
	Buildable
}

var _ FontSelectionDialog = (*fontSelectionDialog)(nil)

// WrapFontSelectionDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontSelectionDialog(obj *externglib.Object) FontSelectionDialog {
	return FontSelectionDialog{
		Dialog:    WrapDialog(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalFontSelectionDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontSelectionDialog(obj), nil
}

// NewFontSelectionDialog constructs a class FontSelectionDialog.
func NewFontSelectionDialog(title string) FontSelectionDialog {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkFontSelectionDialog
	var ret1 FontSelectionDialog

	cret = C.gtk_font_selection_dialog_new(title)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(FontSelectionDialog)

	return ret1
}

// CancelButton gets the “Cancel” button.
func (f fontSelectionDialog) CancelButton() Widget {
	var arg0 *C.GtkFontSelectionDialog

	arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(f.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_font_selection_dialog_get_cancel_button(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// FontName gets the currently-selected font name.
//
// Note that this can be a different string than what you set with
// gtk_font_selection_dialog_set_font_name(), as the font selection widget
// may normalize font names and thus return a string with a different
// structure. For example, “Helvetica Italic Bold 12” could be normalized to
// “Helvetica Bold Italic 12”. Use pango_font_description_equal() if you
// want to compare two font descriptions.
func (f fontSelectionDialog) FontName() string {
	var arg0 *C.GtkFontSelectionDialog

	arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(f.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_font_selection_dialog_get_font_name(arg0)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// FontSelection retrieves the FontSelection widget embedded in the dialog.
func (f fontSelectionDialog) FontSelection() Widget {
	var arg0 *C.GtkFontSelectionDialog

	arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(f.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_font_selection_dialog_get_font_selection(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// OkButton gets the “OK” button.
func (f fontSelectionDialog) OkButton() Widget {
	var arg0 *C.GtkFontSelectionDialog

	arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(f.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_font_selection_dialog_get_ok_button(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// PreviewText gets the text displayed in the preview area.
func (f fontSelectionDialog) PreviewText() string {
	var arg0 *C.GtkFontSelectionDialog

	arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(f.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_font_selection_dialog_get_preview_text(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// SetFontName sets the currently selected font.
func (f fontSelectionDialog) SetFontName(fontname string) bool {
	var arg0 *C.GtkFontSelectionDialog
	var arg1 *C.gchar

	arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(f.Native()))
	arg1 = (*C.gchar)(C.CString(fontname))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_font_selection_dialog_set_font_name(arg0, fontname)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// SetPreviewText sets the text displayed in the preview area.
func (f fontSelectionDialog) SetPreviewText(text string) {
	var arg0 *C.GtkFontSelectionDialog
	var arg1 *C.gchar

	arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(f.Native()))
	arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_font_selection_dialog_set_preview_text(arg0, text)
}
