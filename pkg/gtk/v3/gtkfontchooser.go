// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_font_chooser_level_get_type()), F: marshalFontChooserLevel},
		{T: externglib.Type(C.gtk_font_chooser_get_type()), F: marshalFontChooser},
	})
}

// FontChooserLevel: this enumeration specifies the granularity of font
// selection that is desired in a font chooser.
//
// This enumeration may be extended in the future; applications should ignore
// unknown values.
type FontChooserLevel int

const (
	// FontChooserLevelFamily: allow selecting a font family
	FontChooserLevelFamily FontChooserLevel = 0b0
	// FontChooserLevelStyle: allow selecting a specific font face
	FontChooserLevelStyle FontChooserLevel = 0b1
	// FontChooserLevelSize: allow selecting a specific font size
	FontChooserLevelSize       FontChooserLevel = 0b10
	FontChooserLevelVariations FontChooserLevel = 0b100
	// FontChooserLevelFeatures: allow selecting specific OpenType font features
	FontChooserLevelFeatures FontChooserLevel = 0b1000
)

func marshalFontChooserLevel(p uintptr) (interface{}, error) {
	return FontChooserLevel(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// FontFilterFunc: the type of function that is used for deciding what fonts get
// shown in a FontChooser. See gtk_font_chooser_set_filter_func().
type FontFilterFunc func(family pango.FontFamily, face pango.FontFace, ok bool)

//export gotk4_FontFilterFunc
func _FontFilterFunc(arg0 *C.PangoFontFamily, arg1 *C.PangoFontFace, arg2 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var family pango.FontFamily // out
	var face pango.FontFace     // out

	family = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(pango.FontFamily)
	face = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1))).(pango.FontFace)

	fn := v.(FontFilterFunc)
	ok := fn(family, face)

	var cret C.gboolean // out

	if ok {
		cret = C.TRUE
	}

	return cret
}

// FontChooser is an interface that can be implemented by widgets displaying the
// list of fonts. In GTK+, the main objects that implement this interface are
// FontChooserWidget, FontChooserDialog and FontButton. The GtkFontChooser
// interface has been introducted in GTK+ 3.2.
type FontChooser interface {
	gextras.Objector

	// Font shows or hides the editable preview entry.
	Font() string
	// FontDesc shows or hides the editable preview entry.
	FontDesc() *pango.FontDescription
	// FontFace shows or hides the editable preview entry.
	FontFace() pango.FontFace
	// FontFamily shows or hides the editable preview entry.
	FontFamily() pango.FontFamily
	// FontFeatures shows or hides the editable preview entry.
	FontFeatures() string
	// FontMap shows or hides the editable preview entry.
	FontMap() pango.FontMap
	// FontSize shows or hides the editable preview entry.
	FontSize() int
	// Language shows or hides the editable preview entry.
	Language() string
	// Level shows or hides the editable preview entry.
	Level() FontChooserLevel
	// PreviewText shows or hides the editable preview entry.
	PreviewText() string
	// ShowPreviewEntry shows or hides the editable preview entry.
	ShowPreviewEntry() bool
	// SetFont shows or hides the editable preview entry.
	SetFont(fontname string)
	// SetFontDesc shows or hides the editable preview entry.
	SetFontDesc(fontDesc *pango.FontDescription)
	// SetFontMap shows or hides the editable preview entry.
	SetFontMap(fontmap pango.FontMap)
	// SetLanguage shows or hides the editable preview entry.
	SetLanguage(language string)
	// SetLevel shows or hides the editable preview entry.
	SetLevel(level FontChooserLevel)
	// SetPreviewText shows or hides the editable preview entry.
	SetPreviewText(text string)
	// SetShowPreviewEntry shows or hides the editable preview entry.
	SetShowPreviewEntry(showPreviewEntry bool)
}

// fontChooser implements the FontChooser interface.
type fontChooser struct {
	gextras.Objector
}

var _ FontChooser = (*fontChooser)(nil)

// WrapFontChooser wraps a GObject to a type that implements
// interface FontChooser. It is primarily used internally.
func WrapFontChooser(obj *externglib.Object) FontChooser {
	return fontChooser{
		Objector: obj,
	}
}

func marshalFontChooser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontChooser(obj), nil
}

func (f fontChooser) Font() string {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_chooser_get_font(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (f fontChooser) FontDesc() *pango.FontDescription {
	var _arg0 *C.GtkFontChooser       // out
	var _cret *C.PangoFontDescription // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_chooser_get_font_desc(_arg0)

	var _fontDescription *pango.FontDescription // out

	_fontDescription = (*pango.FontDescription)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_fontDescription, func(v **pango.FontDescription) {
		C.free(unsafe.Pointer(v))
	})

	return _fontDescription
}

func (f fontChooser) FontFace() pango.FontFace {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.PangoFontFace  // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_chooser_get_font_face(_arg0)

	var _fontFace pango.FontFace // out

	_fontFace = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(pango.FontFace)

	return _fontFace
}

func (f fontChooser) FontFamily() pango.FontFamily {
	var _arg0 *C.GtkFontChooser  // out
	var _cret *C.PangoFontFamily // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_chooser_get_font_family(_arg0)

	var _fontFamily pango.FontFamily // out

	_fontFamily = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(pango.FontFamily)

	return _fontFamily
}

func (f fontChooser) FontFeatures() string {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_chooser_get_font_features(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (f fontChooser) FontMap() pango.FontMap {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.PangoFontMap   // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_chooser_get_font_map(_arg0)

	var _fontMap pango.FontMap // out

	_fontMap = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(pango.FontMap)

	return _fontMap
}

func (f fontChooser) FontSize() int {
	var _arg0 *C.GtkFontChooser // out
	var _cret C.gint            // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_chooser_get_font_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (f fontChooser) Language() string {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_chooser_get_language(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (f fontChooser) Level() FontChooserLevel {
	var _arg0 *C.GtkFontChooser     // out
	var _cret C.GtkFontChooserLevel // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_chooser_get_level(_arg0)

	var _fontChooserLevel FontChooserLevel // out

	_fontChooserLevel = FontChooserLevel(_cret)

	return _fontChooserLevel
}

func (f fontChooser) PreviewText() string {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_chooser_get_preview_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (f fontChooser) ShowPreviewEntry() bool {
	var _arg0 *C.GtkFontChooser // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_chooser_get_show_preview_entry(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (f fontChooser) SetFont(fontname string) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(fontname))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_chooser_set_font(_arg0, _arg1)
}

func (f fontChooser) SetFontDesc(fontDesc *pango.FontDescription) {
	var _arg0 *C.GtkFontChooser       // out
	var _arg1 *C.PangoFontDescription // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.PangoFontDescription)(unsafe.Pointer(fontDesc.Native()))

	C.gtk_font_chooser_set_font_desc(_arg0, _arg1)
}

func (f fontChooser) SetFontMap(fontmap pango.FontMap) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.PangoFontMap   // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.PangoFontMap)(unsafe.Pointer(fontmap.Native()))

	C.gtk_font_chooser_set_font_map(_arg0, _arg1)
}

func (f fontChooser) SetLanguage(language string) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.char)(C.CString(language))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_chooser_set_language(_arg0, _arg1)
}

func (f fontChooser) SetLevel(level FontChooserLevel) {
	var _arg0 *C.GtkFontChooser     // out
	var _arg1 C.GtkFontChooserLevel // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	_arg1 = C.GtkFontChooserLevel(level)

	C.gtk_font_chooser_set_level(_arg0, _arg1)
}

func (f fontChooser) SetPreviewText(text string) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_chooser_set_preview_text(_arg0, _arg1)
}

func (f fontChooser) SetShowPreviewEntry(showPreviewEntry bool) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	if showPreviewEntry {
		_arg1 = C.TRUE
	}

	C.gtk_font_chooser_set_show_preview_entry(_arg0, _arg1)
}
