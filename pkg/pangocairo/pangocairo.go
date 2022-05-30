// Code generated by girgen. DO NOT EDIT.

package pangocairo

import (
	"runtime"
	_ "runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: pangocairo pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <pango/pangocairo.h>
import "C"

// glib.Type values for pangocairo.go.
var (
	GTypeFont    = coreglib.Type(C.pango_cairo_font_get_type())
	GTypeFontMap = coreglib.Type(C.pango_cairo_font_map_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeFont, F: marshalFont},
		{T: GTypeFontMap, F: marshalFontMap},
	})
}

// ShapeRendererFunc: function type for rendering attributes of type
// PANGO_ATTR_SHAPE with Pango's Cairo renderer.
type ShapeRendererFunc func(cr *cairo.Context, attr *pango.AttrShape, doPath bool)

//export _gotk4_pangocairo1_ShapeRendererFunc
func _gotk4_pangocairo1_ShapeRendererFunc(arg1 *C.cairo_t, arg2 *C.PangoAttrShape, arg3 C.gboolean, arg4 C.gpointer) {
	var fn ShapeRendererFunc
	{
		v := gbox.Get(uintptr(arg4))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ShapeRendererFunc)
	}

	var _cr *cairo.Context     // out
	var _attr *pango.AttrShape // out
	var _doPath bool           // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.cairo_t)(unsafe.Pointer(v.Native())))
	})
	_attr = (*pango.AttrShape)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	if arg3 != 0 {
		_doPath = true
	}

	fn(_cr, _attr, _doPath)
}

// ContextGetFontOptions retrieves any font rendering options previously set
// with pangocairo.ContextSetFontOptions().
//
// This function does not report options that are derived from the target
// surface by update_context.
//
// The function takes the following parameters:
//
//    - context: PangoContext, from a pangocairo font map.
//
// The function returns the following values:
//
//    - fontOptions (optional): font options previously set on the context, or
//      NULL if no options have been set. This value is owned by the context and
//      must not be modified or freed.
//
func ContextGetFontOptions(context *pango.Context) *cairo.FontOptions {
	var _arg1 *C.PangoContext         // out
	var _cret *C.cairo_font_options_t // in

	_arg1 = (*C.PangoContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.pango_cairo_context_get_font_options(_arg1)
	runtime.KeepAlive(context)

	var _fontOptions *cairo.FontOptions // out

	if _cret != nil {
		_fontOptions = (*cairo.FontOptions)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _fontOptions
}

// ContextGetResolution gets the resolution for the context. See
// pangocairo.ContextSetResolution().
//
// The function takes the following parameters:
//
//    - context: PangoContext, from a pangocairo font map.
//
// The function returns the following values:
//
//    - gdouble: resolution in "dots per inch". A negative value will be returned
//      if no resolution has previously been set.
//
func ContextGetResolution(context *pango.Context) float64 {
	var _arg1 *C.PangoContext // out
	var _cret C.double        // in

	_arg1 = (*C.PangoContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.pango_cairo_context_get_resolution(_arg1)
	runtime.KeepAlive(context)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// ContextSetFontOptions sets the font options used when rendering text with
// this context.
//
// These options override any options that update_context derives from the
// target surface.
//
// The function takes the following parameters:
//
//    - context: PangoContext, from a pangocairo font map.
//    - options (optional): cairo_font_options_t, or NULL to unset any previously
//      set options. A copy is made.
//
func ContextSetFontOptions(context *pango.Context, options *cairo.FontOptions) {
	var _arg1 *C.PangoContext         // out
	var _arg2 *C.cairo_font_options_t // out

	_arg1 = (*C.PangoContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	if options != nil {
		_arg2 = (*C.cairo_font_options_t)(gextras.StructNative(unsafe.Pointer(options)))
	}

	C.pango_cairo_context_set_font_options(_arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(options)
}

// ContextSetResolution sets the resolution for the context.
//
// This is a scale factor between points specified in a PangoFontDescription and
// Cairo units. The default value is 96, meaning that a 10 point font will be 13
// units high. (10 * 96. / 72. = 13.3).
//
// The function takes the following parameters:
//
//    - context: PangoContext, from a pangocairo font map.
//    - dpi: resolution in "dots per inch". (Physical inches aren't actually
//      involved; the terminology is conventional.) A 0 or negative value means
//      to use the resolution from the font map.
//
func ContextSetResolution(context *pango.Context, dpi float64) {
	var _arg1 *C.PangoContext // out
	var _arg2 C.double        // out

	_arg1 = (*C.PangoContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = C.double(dpi)

	C.pango_cairo_context_set_resolution(_arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(dpi)
}

// CreateContext creates a context object set up to match the current
// transformation and target surface of the Cairo context.
//
// This context can then be used to create a layout using pango.Layout.New.
//
// This function is a convenience function that creates a context using the
// default font map, then updates it to cr. If you just need to create a layout
// for use with cr and do not need to access PangoContext directly, you can use
// create_layout instead.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//
// The function returns the following values:
//
//    - context: newly created PangoContext. Free with g_object_unref().
//
func CreateContext(cr *cairo.Context) *pango.Context {
	var _arg1 *C.cairo_t      // out
	var _cret *C.PangoContext // in

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))

	_cret = C.pango_cairo_create_context(_arg1)
	runtime.KeepAlive(cr)

	var _context *pango.Context // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_context = &pango.Context{
			Object: obj,
		}
	}

	return _context
}

// CreateLayout creates a layout object set up to match the current
// transformation and target surface of the Cairo context.
//
// This layout can then be used for text measurement with functions like
// pango.Layout.GetSize() or drawing with functions like show_layout. If you
// change the transformation or target surface for cr, you need to call
// update_layout.
//
// This function is the most convenient way to use Cairo with Pango, however it
// is slightly inefficient since it creates a separate PangoContext object for
// each layout. This might matter in an application that was laying out large
// amounts of text.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//
// The function returns the following values:
//
//    - layout: newly created PangoLayout. Free with g_object_unref().
//
func CreateLayout(cr *cairo.Context) *pango.Layout {
	var _arg1 *C.cairo_t     // out
	var _cret *C.PangoLayout // in

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))

	_cret = C.pango_cairo_create_layout(_arg1)
	runtime.KeepAlive(cr)

	var _layout *pango.Layout // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_layout = &pango.Layout{
			Object: obj,
		}
	}

	return _layout
}

// ErrorUnderlinePath: add a squiggly line to the current path in the specified
// cairo context that approximately covers the given rectangle in the style of
// an underline used to indicate a spelling error.
//
// The width of the underline is rounded to an integer number of up/down
// segments and the resulting rectangle is centered in the original rectangle.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - x: x coordinate of one corner of the rectangle.
//    - y: y coordinate of one corner of the rectangle.
//    - width: non-negative width of the rectangle.
//    - height: non-negative height of the rectangle.
//
func ErrorUnderlinePath(cr *cairo.Context, x, y, width, height float64) {
	var _arg1 *C.cairo_t // out
	var _arg2 C.double   // out
	var _arg3 C.double   // out
	var _arg4 C.double   // out
	var _arg5 C.double   // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.pango_cairo_error_underline_path(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// GlyphStringPath adds the glyphs in glyphs to the current path in the
// specified cairo context.
//
// The origin of the glyphs (the left edge of the baseline) will be at the
// current point of the cairo context.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - font: PangoFont from a PangoCairoFontMap.
//    - glyphs: PangoGlyphString.
//
func GlyphStringPath(cr *cairo.Context, font pango.Fonter, glyphs *pango.GlyphString) {
	var _arg1 *C.cairo_t          // out
	var _arg2 *C.PangoFont        // out
	var _arg3 *C.PangoGlyphString // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoFont)(unsafe.Pointer(coreglib.InternObject(font).Native()))
	_arg3 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))

	C.pango_cairo_glyph_string_path(_arg1, _arg2, _arg3)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(font)
	runtime.KeepAlive(glyphs)
}

// LayoutLinePath adds the text in PangoLayoutLine to the current path in the
// specified cairo context.
//
// The origin of the glyphs (the left edge of the line) will be at the current
// point of the cairo context.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - line: PangoLayoutLine.
//
func LayoutLinePath(cr *cairo.Context, line *pango.LayoutLine) {
	var _arg1 *C.cairo_t         // out
	var _arg2 *C.PangoLayoutLine // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoLayoutLine)(gextras.StructNative(unsafe.Pointer(line)))

	C.pango_cairo_layout_line_path(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(line)
}

// LayoutPath adds the text in a PangoLayout to the current path in the
// specified cairo context.
//
// The top-left corner of the PangoLayout will be at the current point of the
// cairo context.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - layout: pango layout.
//
func LayoutPath(cr *cairo.Context, layout *pango.Layout) {
	var _arg1 *C.cairo_t     // out
	var _arg2 *C.PangoLayout // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))

	C.pango_cairo_layout_path(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(layout)
}

// ShowErrorUnderline: draw a squiggly line in the specified cairo context that
// approximately covers the given rectangle in the style of an underline used to
// indicate a spelling error.
//
// The width of the underline is rounded to an integer number of up/down
// segments and the resulting rectangle is centered in the original rectangle.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - x: x coordinate of one corner of the rectangle.
//    - y: y coordinate of one corner of the rectangle.
//    - width: non-negative width of the rectangle.
//    - height: non-negative height of the rectangle.
//
func ShowErrorUnderline(cr *cairo.Context, x, y, width, height float64) {
	var _arg1 *C.cairo_t // out
	var _arg2 C.double   // out
	var _arg3 C.double   // out
	var _arg4 C.double   // out
	var _arg5 C.double   // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.pango_cairo_show_error_underline(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// ShowGlyphItem draws the glyphs in glyph_item in the specified cairo context,
//
// embedding the text associated with the glyphs in the output if the output
// format supports it (PDF for example), otherwise it acts similar to
// show_glyph_string.
//
// The origin of the glyphs (the left edge of the baseline) will be drawn at the
// current point of the cairo context.
//
// Note that text is the start of the text for layout, which is then indexed by
// glyph_item->item->offset.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - text: UTF-8 text that glyph_item refers to.
//    - glyphItem: PangoGlyphItem.
//
func ShowGlyphItem(cr *cairo.Context, text string, glyphItem *pango.GlyphItem) {
	var _arg1 *C.cairo_t        // out
	var _arg2 *C.char           // out
	var _arg3 *C.PangoGlyphItem // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.PangoGlyphItem)(gextras.StructNative(unsafe.Pointer(glyphItem)))

	C.pango_cairo_show_glyph_item(_arg1, _arg2, _arg3)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(text)
	runtime.KeepAlive(glyphItem)
}

// ShowGlyphString draws the glyphs in glyphs in the specified cairo context.
//
// The origin of the glyphs (the left edge of the baseline) will be drawn at the
// current point of the cairo context.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - font: PangoFont from a PangoCairoFontMap.
//    - glyphs: PangoGlyphString.
//
func ShowGlyphString(cr *cairo.Context, font pango.Fonter, glyphs *pango.GlyphString) {
	var _arg1 *C.cairo_t          // out
	var _arg2 *C.PangoFont        // out
	var _arg3 *C.PangoGlyphString // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoFont)(unsafe.Pointer(coreglib.InternObject(font).Native()))
	_arg3 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))

	C.pango_cairo_show_glyph_string(_arg1, _arg2, _arg3)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(font)
	runtime.KeepAlive(glyphs)
}

// ShowLayout draws a PangoLayout in the specified cairo context.
//
// The top-left corner of the PangoLayout will be drawn at the current point of
// the cairo context.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - layout: pango layout.
//
func ShowLayout(cr *cairo.Context, layout *pango.Layout) {
	var _arg1 *C.cairo_t     // out
	var _arg2 *C.PangoLayout // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))

	C.pango_cairo_show_layout(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(layout)
}

// ShowLayoutLine draws a PangoLayoutLine in the specified cairo context.
//
// The origin of the glyphs (the left edge of the line) will be drawn at the
// current point of the cairo context.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - line: PangoLayoutLine.
//
func ShowLayoutLine(cr *cairo.Context, line *pango.LayoutLine) {
	var _arg1 *C.cairo_t         // out
	var _arg2 *C.PangoLayoutLine // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoLayoutLine)(gextras.StructNative(unsafe.Pointer(line)))

	C.pango_cairo_show_layout_line(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(line)
}

// UpdateContext updates a PangoContext previously created for use with Cairo to
// match the current transformation and target surface of a Cairo context.
//
// If any layouts have been created for the context, it's necessary to call
// pango.Layout.ContextChanged() on those layouts.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - context: PangoContext, from a pangocairo font map.
//
func UpdateContext(cr *cairo.Context, context *pango.Context) {
	var _arg1 *C.cairo_t      // out
	var _arg2 *C.PangoContext // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	C.pango_cairo_update_context(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(context)
}

// UpdateLayout updates the private PangoContext of a PangoLayout created with
// create_layout to match the current transformation and target surface of a
// Cairo context.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - layout: PangoLayout, from create_layout.
//
func UpdateLayout(cr *cairo.Context, layout *pango.Layout) {
	var _arg1 *C.cairo_t     // out
	var _arg2 *C.PangoLayout // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))

	C.pango_cairo_update_layout(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(layout)
}

// Font: PangoCairoFont is an interface exported by fonts for use with Cairo.
//
// The actual type of the font will depend on the particular font technology
// Cairo was compiled to use.
//
// Font wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Font struct {
	_ [0]func() // equal guard
	pango.Font
}

var (
	_ pango.Fonter = (*Font)(nil)
)

// Fonter describes Font's interface methods.
type Fonter interface {
	coreglib.Objector

	baseFont() *Font
}

var _ Fonter = (*Font)(nil)

func wrapFont(obj *coreglib.Object) *Font {
	return &Font{
		Font: pango.Font{
			Object: obj,
		},
	}
}

func marshalFont(p uintptr) (interface{}, error) {
	return wrapFont(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Font) baseFont() *Font {
	return v
}

// BaseFont returns the underlying base object.
func BaseFont(obj Fonter) *Font {
	return obj.baseFont()
}

// FontMap: PangoCairoFontMap is an interface exported by font maps for use with
// Cairo.
//
// The actual type of the font map will depend on the particular font technology
// Cairo was compiled to use.
//
// FontMap wraps an interface. This means the user can get the
// underlying type by calling Cast().
type FontMap struct {
	_ [0]func() // equal guard
	pango.FontMap
}

var (
	_ pango.FontMapper = (*FontMap)(nil)
)

// FontMapper describes FontMap's interface methods.
type FontMapper interface {
	coreglib.Objector
}

var _ FontMapper = (*FontMap)(nil)

func wrapFontMap(obj *coreglib.Object) *FontMap {
	return &FontMap{
		FontMap: pango.FontMap{
			Object: obj,
		},
	}
}

func marshalFontMap(p uintptr) (interface{}, error) {
	return wrapFontMap(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Resolution gets the resolution for the fontmap.
//
// See pangocairo.FontMap.SetResolution().
//
// The function returns the following values:
//
//    - gdouble: resolution in "dots per inch".
//
func (fontmap *FontMap) Resolution() float64 {
	var _arg0 *C.PangoCairoFontMap // out
	var _cret C.double             // in

	_arg0 = (*C.PangoCairoFontMap)(unsafe.Pointer(coreglib.InternObject(fontmap).Native()))

	_cret = C.pango_cairo_font_map_get_resolution(_arg0)
	runtime.KeepAlive(fontmap)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetDefault sets a default PangoCairoFontMap to use with Cairo.
//
// This can be used to change the Cairo font backend that the default fontmap
// uses for example. The old default font map is unreffed and the new font map
// referenced.
//
// Note that since Pango 1.32.6, the default fontmap is per-thread. This
// function only changes the default fontmap for the current thread. Default
// fontmaps of existing threads are not changed. Default fontmaps of any new
// threads will still be created using pangocairo.FontMap.New.
//
// A value of NULL for fontmap will cause the current default font map to be
// released and a new default font map to be created on demand, using
// pangocairo.FontMap.New.
func (fontmap *FontMap) SetDefault() {
	var _arg0 *C.PangoCairoFontMap // out

	if fontmap != nil {
		_arg0 = (*C.PangoCairoFontMap)(unsafe.Pointer(coreglib.InternObject(fontmap).Native()))
	}

	C.pango_cairo_font_map_set_default(_arg0)
	runtime.KeepAlive(fontmap)
}

// SetResolution sets the resolution for the fontmap.
//
// This is a scale factor between points specified in a PangoFontDescription and
// Cairo units. The default value is 96, meaning that a 10 point font will be 13
// units high. (10 * 96. / 72. = 13.3).
//
// The function takes the following parameters:
//
//    - dpi: resolution in "dots per inch". (Physical inches aren't actually
//      involved; the terminology is conventional.).
//
func (fontmap *FontMap) SetResolution(dpi float64) {
	var _arg0 *C.PangoCairoFontMap // out
	var _arg1 C.double             // out

	_arg0 = (*C.PangoCairoFontMap)(unsafe.Pointer(coreglib.InternObject(fontmap).Native()))
	_arg1 = C.double(dpi)

	C.pango_cairo_font_map_set_resolution(_arg0, _arg1)
	runtime.KeepAlive(fontmap)
	runtime.KeepAlive(dpi)
}

// FontMapGetDefault gets a default PangoCairoFontMap to use with Cairo.
//
// Note that the type of the returned object will depend on the particular font
// backend Cairo was compiled to use; you generally should only use the
// PangoFontMap and PangoCairoFontMap interfaces on the returned object.
//
// The default Cairo fontmap can be changed by using
// pangocairo.FontMap.SetDefault(). This can be used to change the Cairo font
// backend that the default fontmap uses for example.
//
// Note that since Pango 1.32.6, the default fontmap is per-thread. Each thread
// gets its own default fontmap. In this way, PangoCairo can be used safely from
// multiple threads.
//
// The function returns the following values:
//
//    - fontMap: default PangoCairo fontmap for the current thread. This object
//      is owned by Pango and must not be freed.
//
func FontMapGetDefault() pango.FontMapper {
	var _cret *C.PangoFontMap // in

	_cret = C.pango_cairo_font_map_get_default()

	var _fontMap pango.FontMapper // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type pango.FontMapper is nil")
		}

		object := coreglib.Take(objptr)
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

	return _fontMap
}

// NewFontMap creates a new PangoCairoFontMap object.
//
// A fontmap is used to cache information about available fonts, and holds
// certain global parameters such as the resolution. In most cases, you can use
// `funcPangoCairo.font_map_get_default] instead.
//
// Note that the type of the returned object will depend on the particular font
// backend Cairo was compiled to use; You generally should only use the
// PangoFontMap and PangoCairoFontMap interfaces on the returned object.
//
// You can override the type of backend returned by using an environment
// variable PANGOCAIRO_BACKEND. Supported types, based on your build, are fc
// (fontconfig), win32, and coretext. If requested type is not available, NULL
// is returned. Ie. this is only useful for testing, when at least two backends
// are compiled in.
//
// The function returns the following values:
//
//    - fontMap: newly allocated PangoFontMap, which should be freed with
//      g_object_unref().
//
func NewFontMap() pango.FontMapper {
	var _cret *C.PangoFontMap // in

	_cret = C.pango_cairo_font_map_new()

	var _fontMap pango.FontMapper // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type pango.FontMapper is nil")
		}

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

	return _fontMap
}
