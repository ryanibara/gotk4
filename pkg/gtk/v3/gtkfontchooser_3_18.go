// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// PangoFontMap* _gotk4_gtk3_FontChooser_virtual_get_font_map(void* fnptr, GtkFontChooser* arg0) {
//   return ((PangoFontMap* (*)(GtkFontChooser*))(fnptr))(arg0);
// };
// void _gotk4_gtk3_FontChooser_virtual_set_font_map(void* fnptr, GtkFontChooser* arg0, PangoFontMap* arg1) {
//   ((void (*)(GtkFontChooser*, PangoFontMap*))(fnptr))(arg0, arg1);
// };
import "C"

// FontMap gets the custom font map of this font chooser widget, or NULL if it
// does not have one.
//
// The function returns the following values:
//
//    - fontMap (optional) or NULL.
//
func (fontchooser *FontChooser) FontMap() pango.FontMapper {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.PangoFontMap   // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_cret = C.gtk_font_chooser_get_font_map(_arg0)
	runtime.KeepAlive(fontchooser)

	var _fontMap pango.FontMapper // out

	if _cret != nil {
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
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.PangoFontMap   // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	if fontmap != nil {
		_arg1 = (*C.PangoFontMap)(unsafe.Pointer(coreglib.InternObject(fontmap).Native()))
	}

	C.gtk_font_chooser_set_font_map(_arg0, _arg1)
	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(fontmap)
}

// fontMap gets the custom font map of this font chooser widget, or NULL if it
// does not have one.
//
// The function returns the following values:
//
//    - fontMap (optional) or NULL.
//
func (fontchooser *FontChooser) fontMap() pango.FontMapper {
	gclass := (*C.GtkFontChooserIface)(coreglib.PeekParentClass(fontchooser))
	fnarg := gclass.get_font_map

	var _arg0 *C.GtkFontChooser // out
	var _cret *C.PangoFontMap   // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_cret = C._gotk4_gtk3_FontChooser_virtual_get_font_map(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(fontchooser)

	var _fontMap pango.FontMapper // out

	if _cret != nil {
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

// setFontMap sets a custom font map to use for this font chooser widget. A
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
func (fontchooser *FontChooser) setFontMap(fontmap pango.FontMapper) {
	gclass := (*C.GtkFontChooserIface)(coreglib.PeekParentClass(fontchooser))
	fnarg := gclass.set_font_map

	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.PangoFontMap   // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	if fontmap != nil {
		_arg1 = (*C.PangoFontMap)(unsafe.Pointer(coreglib.InternObject(fontmap).Native()))
	}

	C._gotk4_gtk3_FontChooser_virtual_set_font_map(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(fontmap)
}
