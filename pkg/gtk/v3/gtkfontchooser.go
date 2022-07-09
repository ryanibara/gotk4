// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gtk3_FontFilterFunc(void*, void*, gpointer);
// extern gint _gotk4_gtk3_FontChooserIface_get_font_size(void*);
// extern void _gotk4_gtk3_FontChooserIface_font_activated(void*, void*);
// extern void _gotk4_gtk3_FontChooserIface_set_font_map(void*, void*);
// extern void _gotk4_gtk3_FontChooser_ConnectFontActivated(gpointer, void*, guintptr);
// extern void callbackDelete(gpointer);
// extern void* _gotk4_gtk3_FontChooserIface_get_font_face(void*);
// extern void* _gotk4_gtk3_FontChooserIface_get_font_family(void*);
// extern void* _gotk4_gtk3_FontChooserIface_get_font_map(void*);
import "C"

// GTypeFontChooserLevel returns the GType for the type FontChooserLevel.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFontChooserLevel() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "FontChooserLevel").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalFontChooserLevel)
	return gtype
}

// GTypeFontChooser returns the GType for the type FontChooser.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFontChooser() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "FontChooser").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalFontChooser)
	return gtype
}

// FontChooserLevel: this enumeration specifies the granularity of font
// selection that is desired in a font chooser.
//
// This enumeration may be extended in the future; applications should ignore
// unknown values.
type FontChooserLevel C.guint

const (
	// FontChooserLevelFamily: allow selecting a font family.
	FontChooserLevelFamily FontChooserLevel = 0b0
	// FontChooserLevelStyle: allow selecting a specific font face.
	FontChooserLevelStyle FontChooserLevel = 0b1
	// FontChooserLevelSize: allow selecting a specific font size.
	FontChooserLevelSize       FontChooserLevel = 0b10
	FontChooserLevelVariations FontChooserLevel = 0b100
	// FontChooserLevelFeatures: allow selecting specific OpenType font
	// features.
	FontChooserLevelFeatures FontChooserLevel = 0b1000
)

func marshalFontChooserLevel(p uintptr) (interface{}, error) {
	return FontChooserLevel(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for FontChooserLevel.
func (f FontChooserLevel) String() string {
	if f == 0 {
		return "FontChooserLevel(0)"
	}

	var builder strings.Builder
	builder.Grow(117)

	for f != 0 {
		next := f & (f - 1)
		bit := f - next

		switch bit {
		case FontChooserLevelFamily:
			builder.WriteString("Family|")
		case FontChooserLevelStyle:
			builder.WriteString("Style|")
		case FontChooserLevelSize:
			builder.WriteString("Size|")
		case FontChooserLevelVariations:
			builder.WriteString("Variations|")
		case FontChooserLevelFeatures:
			builder.WriteString("Features|")
		default:
			builder.WriteString(fmt.Sprintf("FontChooserLevel(0b%b)|", bit))
		}

		f = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if f contains other.
func (f FontChooserLevel) Has(other FontChooserLevel) bool {
	return (f & other) == other
}

// FontFilterFunc: type of function that is used for deciding what fonts get
// shown in a FontChooser. See gtk_font_chooser_set_filter_func().
type FontFilterFunc func(family pango.FontFamilier, face pango.FontFacer) (ok bool)

//export _gotk4_gtk3_FontFilterFunc
func _gotk4_gtk3_FontFilterFunc(arg1 *C.void, arg2 *C.void, arg3 C.gpointer) (cret C.gboolean) {
	var fn FontFilterFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(FontFilterFunc)
	}

	var _family pango.FontFamilier // out
	var _face pango.FontFacer      // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type pango.FontFamilier is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(pango.FontFamilier)
			return ok
		})
		rv, ok := casted.(pango.FontFamilier)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.FontFamilier")
		}
		_family = rv
	}
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type pango.FontFacer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(pango.FontFacer)
			return ok
		})
		rv, ok := casted.(pango.FontFacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.FontFacer")
		}
		_face = rv
	}

	ok := fn(_family, _face)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// FontChooser is an interface that can be implemented by widgets displaying the
// list of fonts. In GTK+, the main objects that implement this interface are
// FontChooserWidget, FontChooserDialog and FontButton. The GtkFontChooser
// interface has been introducted in GTK+ 3.2.
//
// FontChooser wraps an interface. This means the user can get the
// underlying type by calling Cast().
type FontChooser struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*FontChooser)(nil)
)

// FontChooserer describes FontChooser's interface methods.
type FontChooserer interface {
	coreglib.Objector

	// Font gets the currently-selected font name.
	Font() string
	// FontDesc gets the currently-selected font.
	FontDesc() *pango.FontDescription
	// FontFace gets the FontFace representing the selected font group details
	// (i.e.
	FontFace() pango.FontFacer
	// FontFamily gets the FontFamily representing the selected font family.
	FontFamily() pango.FontFamilier
	// FontFeatures gets the currently-selected font features.
	FontFeatures() string
	// FontMap gets the custom font map of this font chooser widget, or NULL if
	// it does not have one.
	FontMap() pango.FontMapper
	// FontSize: selected font size.
	FontSize() int32
	// Language gets the language that is used for font features.
	Language() string
	// PreviewText gets the text displayed in the preview area.
	PreviewText() string
	// ShowPreviewEntry returns whether the preview entry is shown or not.
	ShowPreviewEntry() bool
	// SetFilterFunc adds a filter function that decides which fonts to display
	// in the font chooser.
	SetFilterFunc(filter FontFilterFunc)
	// SetFont sets the currently-selected font.
	SetFont(fontname string)
	// SetFontDesc sets the currently-selected font from font_desc.
	SetFontDesc(fontDesc *pango.FontDescription)
	// SetFontMap sets a custom font map to use for this font chooser widget.
	SetFontMap(fontmap pango.FontMapper)
	// SetLanguage sets the language to use for font features.
	SetLanguage(language string)
	// SetPreviewText sets the text displayed in the preview area.
	SetPreviewText(text string)
	// SetShowPreviewEntry shows or hides the editable preview entry.
	SetShowPreviewEntry(showPreviewEntry bool)

	// Font-activated is emitted when a font is activated.
	ConnectFontActivated(func(fontname string)) coreglib.SignalHandle
}

var _ FontChooserer = (*FontChooser)(nil)

func wrapFontChooser(obj *coreglib.Object) *FontChooser {
	return &FontChooser{
		Object: obj,
	}
}

func marshalFontChooser(p uintptr) (interface{}, error) {
	return wrapFontChooser(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_FontChooser_ConnectFontActivated
func _gotk4_gtk3_FontChooser_ConnectFontActivated(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(fontname string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(fontname string))
	}

	var _fontname string // out

	_fontname = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_fontname)
}

// ConnectFontActivated is emitted when a font is activated. This usually
// happens when the user double clicks an item, or an item is selected and the
// user presses one of the keys Space, Shift+Space, Return or Enter.
func (fontchooser *FontChooser) ConnectFontActivated(f func(fontname string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(fontchooser, "font-activated", false, unsafe.Pointer(C._gotk4_gtk3_FontChooser_ConnectFontActivated), f)
}

// Font gets the currently-selected font name.
//
// Note that this can be a different string than what you set with
// gtk_font_chooser_set_font(), as the font chooser widget may normalize font
// names and thus return a string with a different structure. For example,
// “Helvetica Italic Bold 12” could be normalized to “Helvetica Bold Italic 12”.
//
// Use pango_font_description_equal() if you want to compare two font
// descriptions.
//
// The function returns the following values:
//
//    - utf8 (optional): string with the name of the current font, or NULL if no
//      font is selected. You must free this string with g_free().
//
func (fontchooser *FontChooser) Font() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_info := girepository.MustFind("Gtk", "FontChooser")
	_gret := _info.InvokeIfaceMethod("get_font", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// FontDesc gets the currently-selected font.
//
// Note that this can be a different string than what you set with
// gtk_font_chooser_set_font(), as the font chooser widget may normalize font
// names and thus return a string with a different structure. For example,
// “Helvetica Italic Bold 12” could be normalized to “Helvetica Bold Italic 12”.
//
// Use pango_font_description_equal() if you want to compare two font
// descriptions.
//
// The function returns the following values:
//
//    - fontDescription (optional) for the current font, or NULL if no font is
//      selected.
//
func (fontchooser *FontChooser) FontDesc() *pango.FontDescription {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_info := girepository.MustFind("Gtk", "FontChooser")
	_gret := _info.InvokeIfaceMethod("get_font_desc", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _fontDescription *pango.FontDescription // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_fontDescription = (*pango.FontDescription)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_fontDescription)),
			func(intern *struct{ C unsafe.Pointer }) {
				{
					args := [1]girepository.Argument{(*C.void)(intern.C)}
					girepository.MustFind("Pango", "FontDescription").InvokeRecordMethod("free", args[:], nil)
				}
			},
		)
	}

	return _fontDescription
}

// FontFace gets the FontFace representing the selected font group details (i.e.
// family, slant, weight, width, etc).
//
// If the selected font is not installed, returns NULL.
//
// The function returns the following values:
//
//    - fontFace (optional) representing the selected font group details, or
//      NULL. The returned object is owned by fontchooser and must not be
//      modified or freed.
//
func (fontchooser *FontChooser) FontFace() pango.FontFacer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_info := girepository.MustFind("Gtk", "FontChooser")
	_gret := _info.InvokeIfaceMethod("get_font_face", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _fontFace pango.FontFacer // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(pango.FontFacer)
				return ok
			})
			rv, ok := casted.(pango.FontFacer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.FontFacer")
			}
			_fontFace = rv
		}
	}

	return _fontFace
}

// FontFamily gets the FontFamily representing the selected font family. Font
// families are a collection of font faces.
//
// If the selected font is not installed, returns NULL.
//
// The function returns the following values:
//
//    - fontFamily (optional) representing the selected font family, or NULL. The
//      returned object is owned by fontchooser and must not be modified or
//      freed.
//
func (fontchooser *FontChooser) FontFamily() pango.FontFamilier {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_info := girepository.MustFind("Gtk", "FontChooser")
	_gret := _info.InvokeIfaceMethod("get_font_family", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _fontFamily pango.FontFamilier // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(pango.FontFamilier)
				return ok
			})
			rv, ok := casted.(pango.FontFamilier)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.FontFamilier")
			}
			_fontFamily = rv
		}
	}

	return _fontFamily
}

// FontFeatures gets the currently-selected font features.
//
// The function returns the following values:
//
//    - utf8: currently selected font features.
//
func (fontchooser *FontChooser) FontFeatures() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_info := girepository.MustFind("Gtk", "FontChooser")
	_gret := _info.InvokeIfaceMethod("get_font_features", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FontMap gets the custom font map of this font chooser widget, or NULL if it
// does not have one.
//
// The function returns the following values:
//
//    - fontMap (optional) or NULL.
//
func (fontchooser *FontChooser) FontMap() pango.FontMapper {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_info := girepository.MustFind("Gtk", "FontChooser")
	_gret := _info.InvokeIfaceMethod("get_font_map", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _fontMap pango.FontMapper // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(pango.FontMapper)
				return ok
			})
			rv, ok := casted.(pango.FontMapper)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.FontMapper")
			}
			_fontMap = rv
		}
	}

	return _fontMap
}

// FontSize: selected font size.
//
// The function returns the following values:
//
//    - gint: n integer representing the selected font size, or -1 if no font
//      size is selected.
//
func (fontchooser *FontChooser) FontSize() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_info := girepository.MustFind("Gtk", "FontChooser")
	_gret := _info.InvokeIfaceMethod("get_font_size", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// Language gets the language that is used for font features.
//
// The function returns the following values:
//
//    - utf8: currently selected language.
//
func (fontchooser *FontChooser) Language() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_info := girepository.MustFind("Gtk", "FontChooser")
	_gret := _info.InvokeIfaceMethod("get_language", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// PreviewText gets the text displayed in the preview area.
//
// The function returns the following values:
//
//    - utf8: text displayed in the preview area.
//
func (fontchooser *FontChooser) PreviewText() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_info := girepository.MustFind("Gtk", "FontChooser")
	_gret := _info.InvokeIfaceMethod("get_preview_text", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ShowPreviewEntry returns whether the preview entry is shown or not.
//
// The function returns the following values:
//
//    - ok: TRUE if the preview entry is shown or FALSE if it is hidden.
//
func (fontchooser *FontChooser) ShowPreviewEntry() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_info := girepository.MustFind("Gtk", "FontChooser")
	_gret := _info.InvokeIfaceMethod("get_show_preview_entry", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetFilterFunc adds a filter function that decides which fonts to display in
// the font chooser.
//
// The function takes the following parameters:
//
//    - filter (optional) or NULL.
//
func (fontchooser *FontChooser) SetFilterFunc(filter FontFilterFunc) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	if filter != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[1])) = (*[0]byte)(C._gotk4_gtk3_FontFilterFunc)
		_args[2] = C.gpointer(gbox.Assign(filter))
		_args[3] = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_info := girepository.MustFind("Gtk", "FontChooser")
	_info.InvokeIfaceMethod("set_filter_func", _args[:], nil)

	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(filter)
}

// SetFont sets the currently-selected font.
//
// The function takes the following parameters:
//
//    - fontname: font name like “Helvetica 12” or “Times Bold 18”.
//
func (fontchooser *FontChooser) SetFont(fontname string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(fontname)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gtk", "FontChooser")
	_info.InvokeIfaceMethod("set_font", _args[:], nil)

	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(fontname)
}

// SetFontDesc sets the currently-selected font from font_desc.
//
// The function takes the following parameters:
//
//    - fontDesc: FontDescription.
//
func (fontchooser *FontChooser) SetFontDesc(fontDesc *pango.FontDescription) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(fontDesc)))

	_info := girepository.MustFind("Gtk", "FontChooser")
	_info.InvokeIfaceMethod("set_font_desc", _args[:], nil)

	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(fontDesc)
}

// SetFontMap sets a custom font map to use for this font chooser widget. A
// custom font map can be used to present application-specific fonts instead of
// or in addition to the normal system fonts.
//
//    FcConfig *config;
//    PangoFontMap *fontmap;
//
//    config = FcInitLoadConfigAndFonts ();
//    FcConfigAppFontAddFile (config, my_app_font_file);
//
//    fontmap = pango_cairo_font_map_new_for_font_type (CAIRO_FONT_TYPE_FT);
//    pango_fc_font_map_set_config (PANGO_FC_FONT_MAP (fontmap), config);
//
//    gtk_font_chooser_set_font_map (font_chooser, fontmap);
//
// Note that other GTK+ widgets will only be able to use the
// application-specific font if it is present in the font map they use:
//
//    context = gtk_widget_get_pango_context (label);
//    pango_context_set_font_map (context, fontmap);.
//
// The function takes the following parameters:
//
//    - fontmap (optional): FontMap.
//
func (fontchooser *FontChooser) SetFontMap(fontmap pango.FontMapper) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	if fontmap != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontmap).Native()))
	}

	_info := girepository.MustFind("Gtk", "FontChooser")
	_info.InvokeIfaceMethod("set_font_map", _args[:], nil)

	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(fontmap)
}

// SetLanguage sets the language to use for font features.
//
// The function takes the following parameters:
//
//    - language: language.
//
func (fontchooser *FontChooser) SetLanguage(language string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(language)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gtk", "FontChooser")
	_info.InvokeIfaceMethod("set_language", _args[:], nil)

	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(language)
}

// SetPreviewText sets the text displayed in the preview area. The text is used
// to show how the selected font looks.
//
// The function takes the following parameters:
//
//    - text to display in the preview area.
//
func (fontchooser *FontChooser) SetPreviewText(text string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gtk", "FontChooser")
	_info.InvokeIfaceMethod("set_preview_text", _args[:], nil)

	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(text)
}

// SetShowPreviewEntry shows or hides the editable preview entry.
//
// The function takes the following parameters:
//
//    - showPreviewEntry: whether to show the editable preview entry or not.
//
func (fontchooser *FontChooser) SetShowPreviewEntry(showPreviewEntry bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	if showPreviewEntry {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "FontChooser")
	_info.InvokeIfaceMethod("set_show_preview_entry", _args[:], nil)

	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(showPreviewEntry)
}
