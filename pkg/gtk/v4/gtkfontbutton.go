// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_font_button_get_type()), F: marshalFontButton},
	})
}

// FontButton: the `GtkFontButton` allows to open a font chooser dialog to
// change the font.
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
type FontButton interface {
	Widget
	FontChooser

	Modal() bool

	Title() string

	UseFont() bool

	UseSize() bool

	SetModalFontButton(modal bool)

	SetTitleFontButton(title string)

	SetUseFontFontButton(useFont bool)

	SetUseSizeFontButton(useSize bool)
}

// fontButton implements the FontButton class.
type fontButton struct {
	Widget
}

// WrapFontButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontButton(obj *externglib.Object) FontButton {
	return fontButton{
		Widget: WrapWidget(obj),
	}
}

func marshalFontButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontButton(obj), nil
}

func NewFontButton() FontButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_font_button_new()

	var _fontButton FontButton // out

	_fontButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(FontButton)

	return _fontButton
}

func NewFontButtonWithFont(fontname string) FontButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(fontname))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_font_button_new_with_font(_arg1)

	var _fontButton FontButton // out

	_fontButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(FontButton)

	return _fontButton
}

func (f fontButton) Modal() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_button_get_modal(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (f fontButton) Title() string {
	var _arg0 *C.GtkFontButton // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_button_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (f fontButton) UseFont() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_button_get_use_font(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (f fontButton) UseSize() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_button_get_use_size(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (f fontButton) SetModalFontButton(modal bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_modal(_arg0, _arg1)
}

func (f fontButton) SetTitleFontButton(title string) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_button_set_title(_arg0, _arg1)
}

func (f fontButton) SetUseFontFontButton(useFont bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))
	if useFont {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_use_font(_arg0, _arg1)
}

func (f fontButton) SetUseSizeFontButton(useSize bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))
	if useSize {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_use_size(_arg0, _arg1)
}

func (s fontButton) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s fontButton) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s fontButton) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s fontButton) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s fontButton) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s fontButton) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s fontButton) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b fontButton) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}

func (f fontButton) Font() string {
	return WrapFontChooser(gextras.InternObject(f)).Font()
}

func (f fontButton) FontDesc() *pango.FontDescription {
	return WrapFontChooser(gextras.InternObject(f)).FontDesc()
}

func (f fontButton) FontFace() pango.FontFace {
	return WrapFontChooser(gextras.InternObject(f)).FontFace()
}

func (f fontButton) FontFamily() pango.FontFamily {
	return WrapFontChooser(gextras.InternObject(f)).FontFamily()
}

func (f fontButton) FontFeatures() string {
	return WrapFontChooser(gextras.InternObject(f)).FontFeatures()
}

func (f fontButton) FontMap() pango.FontMap {
	return WrapFontChooser(gextras.InternObject(f)).FontMap()
}

func (f fontButton) FontSize() int {
	return WrapFontChooser(gextras.InternObject(f)).FontSize()
}

func (f fontButton) Language() string {
	return WrapFontChooser(gextras.InternObject(f)).Language()
}

func (f fontButton) Level() FontChooserLevel {
	return WrapFontChooser(gextras.InternObject(f)).Level()
}

func (f fontButton) PreviewText() string {
	return WrapFontChooser(gextras.InternObject(f)).PreviewText()
}

func (f fontButton) ShowPreviewEntry() bool {
	return WrapFontChooser(gextras.InternObject(f)).ShowPreviewEntry()
}

func (f fontButton) SetFont(fontname string) {
	WrapFontChooser(gextras.InternObject(f)).SetFont(fontname)
}

func (f fontButton) SetFontDesc(fontDesc *pango.FontDescription) {
	WrapFontChooser(gextras.InternObject(f)).SetFontDesc(fontDesc)
}

func (f fontButton) SetFontMap(fontmap pango.FontMap) {
	WrapFontChooser(gextras.InternObject(f)).SetFontMap(fontmap)
}

func (f fontButton) SetLanguage(language string) {
	WrapFontChooser(gextras.InternObject(f)).SetLanguage(language)
}

func (f fontButton) SetLevel(level FontChooserLevel) {
	WrapFontChooser(gextras.InternObject(f)).SetLevel(level)
}

func (f fontButton) SetPreviewText(text string) {
	WrapFontChooser(gextras.InternObject(f)).SetPreviewText(text)
}

func (f fontButton) SetShowPreviewEntry(showPreviewEntry bool) {
	WrapFontChooser(gextras.InternObject(f)).SetShowPreviewEntry(showPreviewEntry)
}
