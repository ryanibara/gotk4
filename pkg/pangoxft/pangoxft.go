// Code generated by girgen. DO NOT EDIT.

package pangoxft

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/freetype2"
	"github.com/diamondburned/gotk4/pkg/pango"
	"github.com/diamondburned/gotk4/pkg/pangofc"
	"github.com/diamondburned/gotk4/pkg/xft"
	"github.com/diamondburned/gotk4/pkg/xlib"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pangoxft
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pangoxft.h>
//
// // extern void callbackDelete(gpointer);
import "C"

//export callbackDelete
func callbackDelete(ptr C.gpointer) {
	box.Delete(box.Callback, uintptr(ptr))
}

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{

		// Records
		// Skipped RendererClass.
		// Skipped RendererPrivate.

		// Classes
		{T: externglib.Type(C.pango_xft_font_get_type()), F: marshalFont},
		{T: externglib.Type(C.pango_xft_font_map_get_type()), F: marshalFontMap},
		{T: externglib.Type(C.pango_xft_renderer_get_type()), F: marshalRenderer},
	})
}

// GetContext retrieves a Context appropriate for rendering with Xft fonts on
// the given screen of the given display.
func GetContext(display *xlib.Display, screen int) pango.Context {
	var arg1 *C.Display
	var arg2 C.int

	arg1 = (*C.Display)(display.Native())
	arg2 = C.int(screen)

	ret := C.pango_xft_get_context(arg1, arg2)

	var ret0 pango.Context

	ret0 = pango.WrapContext(externglib.Take(unsafe.Pointer(ret.Native())))

	return ret0
}

// GetFontMap returns the XftFontMap for the given display and screen. The
// fontmap is owned by Pango and will be valid until the display is closed.
func GetFontMap(display *xlib.Display, screen int) pango.FontMap {
	var arg1 *C.Display
	var arg2 C.int

	arg1 = (*C.Display)(display.Native())
	arg2 = C.int(screen)

	ret := C.pango_xft_get_font_map(arg1, arg2)

	var ret0 pango.FontMap

	ret0 = pango.WrapFontMap(externglib.Take(unsafe.Pointer(ret.Native())))

	return ret0
}

// PictureRender renders a GlyphString onto an Xrender Picture object.
func PictureRender(display *xlib.Display, srcPicture xlib.Picture, destPicture xlib.Picture, font pango.Font, glyphs *pango.GlyphString, x int, y int) {
	var arg1 *C.Display
	var arg2 C.Picture
	var arg3 C.Picture
	var arg4 *C.PangoFont
	var arg5 *C.PangoGlyphString
	var arg6 C.gint
	var arg7 C.gint

	arg1 = (*C.Display)(display.Native())
	arg4 = (*C.PangoFont)(font.Native())
	arg5 = (*C.PangoGlyphString)(glyphs.Native())
	arg6 = C.gint(x)
	arg7 = C.gint(y)

	C.pango_xft_picture_render(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// Render renders a GlyphString onto an XftDraw object wrapping an X drawable.
func Render(draw *xft.Draw, color *xft.Color, font pango.Font, glyphs *pango.GlyphString, x int, y int) {
	var arg1 *C.XftDraw
	var arg2 *C.XftColor
	var arg3 *C.PangoFont
	var arg4 *C.PangoGlyphString
	var arg5 C.gint
	var arg6 C.gint

	arg1 = (*C.XftDraw)(draw.Native())
	arg2 = (*C.XftColor)(color.Native())
	arg3 = (*C.PangoFont)(font.Native())
	arg4 = (*C.PangoGlyphString)(glyphs.Native())
	arg5 = C.gint(x)
	arg6 = C.gint(y)

	C.pango_xft_render(arg1, arg2, arg3, arg4, arg5, arg6)
}

// RenderLayout: render a Layout onto a Draw
func RenderLayout(draw *xft.Draw, color *xft.Color, layout pango.Layout, x int, y int) {
	var arg1 *C.XftDraw
	var arg2 *C.XftColor
	var arg3 *C.PangoLayout
	var arg4 C.int
	var arg5 C.int

	arg1 = (*C.XftDraw)(draw.Native())
	arg2 = (*C.XftColor)(color.Native())
	arg3 = (*C.PangoLayout)(layout.Native())
	arg4 = C.int(x)
	arg5 = C.int(y)

	C.pango_xft_render_layout(arg1, arg2, arg3, arg4, arg5)
}

// RenderLayoutLine: render a LayoutLine onto a Draw
func RenderLayoutLine(draw *xft.Draw, color *xft.Color, line *pango.LayoutLine, x int, y int) {
	var arg1 *C.XftDraw
	var arg2 *C.XftColor
	var arg3 *C.PangoLayoutLine
	var arg4 C.int
	var arg5 C.int

	arg1 = (*C.XftDraw)(draw.Native())
	arg2 = (*C.XftColor)(color.Native())
	arg3 = (*C.PangoLayoutLine)(line.Native())
	arg4 = C.int(x)
	arg5 = C.int(y)

	C.pango_xft_render_layout_line(arg1, arg2, arg3, arg4, arg5)
}

// RenderTransformed renders a GlyphString onto a Draw, possibly transforming
// the layed-out coordinates through a transformation matrix. Note that the
// transformation matrix for @font is not changed, so to produce correct
// rendering results, the @font must have been loaded using a Context with an
// identical transformation matrix to that passed in to this function.
func RenderTransformed(draw *xft.Draw, color *xft.Color, matrix *pango.Matrix, font pango.Font, glyphs *pango.GlyphString, x int, y int) {
	var arg1 *C.XftDraw
	var arg2 *C.XftColor
	var arg3 *C.PangoMatrix
	var arg4 *C.PangoFont
	var arg5 *C.PangoGlyphString
	var arg6 C.int
	var arg7 C.int

	arg1 = (*C.XftDraw)(draw.Native())
	arg2 = (*C.XftColor)(color.Native())
	arg3 = (*C.PangoMatrix)(matrix.Native())
	arg4 = (*C.PangoFont)(font.Native())
	arg5 = (*C.PangoGlyphString)(glyphs.Native())
	arg6 = C.int(x)
	arg7 = C.int(y)

	C.pango_xft_render_transformed(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// SetDefaultSubstitute sets a function that will be called to do final
// configuration substitution on a Pattern before it is used to load the font.
// This function can be used to do things like set hinting and antialiasing
// options.
func SetDefaultSubstitute(display *xlib.Display, screen int, _func SubstituteFunc) {
	var arg1 *C.Display
	var arg2 C.int
	var arg3 C.PangoXftSubstituteFunc
	arg4 := C.gpointer(box.Assign(data))

	arg1 = (*C.Display)(display.Native())
	arg2 = C.int(screen)
	arg3 = (*[0]byte)(C.gotk4_SubstituteFunc)

	C.pango_xft_set_default_substitute(arg1, arg2, arg3, (*[0]byte)(C.callbackDelete))
}

// ShutdownDisplay: release any resources that have been cached for the
// combination of @display and @screen. Note that when the X display is closed,
// resources are released automatically, without needing to call this function.
func ShutdownDisplay(display *xlib.Display, screen int) {
	var arg1 *C.Display
	var arg2 C.int

	arg1 = (*C.Display)(display.Native())
	arg2 = C.int(screen)

	C.pango_xft_shutdown_display(arg1, arg2)
}

// SubstituteChanged: call this function any time the results of the default
// substitution function set with pango_xft_set_default_substitute() change.
// That is, if your substitution function will return different results for the
// same input pattern, you must call this function.
func SubstituteChanged(display *xlib.Display, screen int) {
	var arg1 *C.Display
	var arg2 C.int

	arg1 = (*C.Display)(display.Native())
	arg2 = C.int(screen)

	C.pango_xft_substitute_changed(arg1, arg2)
}

// Font is an implementation of FcFont using the Xft library for rendering. It
// is used in conjunction with XftFontMap.
type Font interface {
	pangofc.Font

	// Display returns the X display of the `XftFont` of a font.
	Display() *xlib.Display
	// Glyph gets the glyph index for a given Unicode character for @font. If
	// you only want to determine whether the font has the glyph, use
	// pango_xft_font_has_char().
	//
	// Use pango_fc_font_get_glyph() instead.
	Glyph(wc uint32) uint
	// UnknownGlyph returns the index of a glyph suitable for drawing @wc as an
	// unknown character.
	//
	// Use PANGO_GET_UNKNOWN_GLYPH() instead.
	UnknownGlyph(wc uint32) pango.Glyph
	// HasChar determines whether @font has a glyph for the codepoint @wc.
	//
	// Use pango_fc_font_has_char() instead.
	HasChar(wc uint32) bool
	// LockFace gets the FreeType `FT_Face` associated with a font.
	//
	// This face will be kept around until you call
	// pango_xft_font_unlock_face().
	//
	// Use pango_fc_font_lock_face() instead.
	LockFace() freetype2.Face
	// UnlockFace releases a font previously obtained with
	// pango_xft_font_lock_face().
	//
	// Use pango_fc_font_unlock_face() instead.
	UnlockFace()
}

type font struct {
	pangofc.Font
}

// WrapFont wraps a GObject to the right type. It is
// primarily used internally.
func WrapFont(obj *externglib.Object) Font {
	return font{pangofc.WrapFont(obj)}
}

func marshalFont(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFont(obj), nil
}

// Display returns the X display of the `XftFont` of a font.
func (font font) Display() *xlib.Display {
	var arg0 *C.PangoFont

	arg0 = (*C.PangoFont)(font.Native())

	ret := C.pango_xft_font_get_display(arg0)

	var ret0 *xlib.Display

	ret0 = xlib.WrapDisplay(ret)

	return ret0
}

// Glyph gets the glyph index for a given Unicode character for @font. If
// you only want to determine whether the font has the glyph, use
// pango_xft_font_has_char().
//
// Use pango_fc_font_get_glyph() instead.
func (font font) Glyph(wc uint32) uint {
	var arg0 *C.PangoFont
	var arg1 C.gunichar

	arg0 = (*C.PangoFont)(font.Native())
	arg1 = C.gunichar(wc)

	ret := C.pango_xft_font_get_glyph(arg0, arg1)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// UnknownGlyph returns the index of a glyph suitable for drawing @wc as an
// unknown character.
//
// Use PANGO_GET_UNKNOWN_GLYPH() instead.
func (font font) UnknownGlyph(wc uint32) pango.Glyph {
	var arg0 *C.PangoFont
	var arg1 C.gunichar

	arg0 = (*C.PangoFont)(font.Native())
	arg1 = C.gunichar(wc)

	ret := C.pango_xft_font_get_unknown_glyph(arg0, arg1)

	var ret0 pango.Glyph

	{
		var tmp uint32
		tmp = uint32(ret)
		ret0 = pango.Glyph(tmp)
	}

	return ret0
}

// HasChar determines whether @font has a glyph for the codepoint @wc.
//
// Use pango_fc_font_has_char() instead.
func (font font) HasChar(wc uint32) bool {
	var arg0 *C.PangoFont
	var arg1 C.gunichar

	arg0 = (*C.PangoFont)(font.Native())
	arg1 = C.gunichar(wc)

	ret := C.pango_xft_font_has_char(arg0, arg1)

	var ret0 bool

	ret0 = gextras.Gobool(ret)

	return ret0
}

// LockFace gets the FreeType `FT_Face` associated with a font.
//
// This face will be kept around until you call
// pango_xft_font_unlock_face().
//
// Use pango_fc_font_lock_face() instead.
func (font font) LockFace() freetype2.Face {
	var arg0 *C.PangoFont

	arg0 = (*C.PangoFont)(font.Native())

	ret := C.pango_xft_font_lock_face(arg0)

	var ret0 freetype2.Face

	ret0 = freetype2.WrapFace(ret)

	return ret0
}

// UnlockFace releases a font previously obtained with
// pango_xft_font_lock_face().
//
// Use pango_fc_font_unlock_face() instead.
func (font font) UnlockFace() {
	var arg0 *C.PangoFont

	arg0 = (*C.PangoFont)(font.Native())

	C.pango_xft_font_unlock_face(arg0)
}

// FontMap is an implementation of FcFontMap suitable for the Xft library as the
// renderer. It is used in to create fonts of type XftFont.
type FontMap interface {
	pangofc.FontMap
}

type fontMap struct {
	pangofc.FontMap
}

// WrapFontMap wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontMap(obj *externglib.Object) FontMap {
	return fontMap{pangofc.WrapFontMap(obj)}
}

func marshalFontMap(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontMap(obj), nil
}

// Renderer is a subclass of Renderer used for rendering with Pango's Xft
// backend. It can be used directly, or it can be further subclassed to modify
// exactly how drawing of individual elements occurs.
type Renderer interface {
	pango.Renderer

	// SetDefaultColor sets the default foreground color for a Renderer.
	SetDefaultColor(defaultColor *pango.Color)
	// SetDraw sets the Draw object that the renderer is drawing to. The
	// renderer must not be currently active.
	SetDraw(draw *xft.Draw)
}

type renderer struct {
	pango.Renderer
}

// WrapRenderer wraps a GObject to the right type. It is
// primarily used internally.
func WrapRenderer(obj *externglib.Object) Renderer {
	return renderer{pango.WrapRenderer(obj)}
}

func marshalRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRenderer(obj), nil
}

// NewPango constructs a class Renderer.
func NewPango(display *xlib.Display, screen int) pango.Renderer {
	var arg1 *C.Display
	var arg2 C.int

	arg1 = (*C.Display)(display.Native())
	arg2 = C.int(screen)

	ret := C.pango_xft_renderer_new(arg1, arg2)

	var ret0 pango.Renderer

	ret0 = pango.WrapRenderer(externglib.AssumeOwnership(unsafe.Pointer(ret.Native())))

	return ret0
}

// SetDefaultColor sets the default foreground color for a Renderer.
func (xftrenderer renderer) SetDefaultColor(defaultColor *pango.Color) {
	var arg0 *C.PangoXftRenderer
	var arg1 *C.PangoColor

	arg0 = (*C.PangoXftRenderer)(xftrenderer.Native())
	arg1 = (*C.PangoColor)(defaultColor.Native())

	C.pango_xft_renderer_set_default_color(arg0, arg1)
}

// SetDraw sets the Draw object that the renderer is drawing to. The
// renderer must not be currently active.
func (xftrenderer renderer) SetDraw(draw *xft.Draw) {
	var arg0 *C.PangoXftRenderer
	var arg1 *C.XftDraw

	arg0 = (*C.PangoXftRenderer)(xftrenderer.Native())
	arg1 = (*C.XftDraw)(draw.Native())

	C.pango_xft_renderer_set_draw(arg0, arg1)
}
