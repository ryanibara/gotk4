// Code generated by girgen. DO NOT EDIT.

package pangoxft

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/fontconfig"
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
import "C"

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

type SubstituteFunc func(pattern *fontconfig.Pattern, data interface{})

//export cSubstituteFunc
func cSubstituteFunc(arg0 *C.FcPattern, arg1 C.gpointer)

// GetContext: retrieves a Context appropriate for rendering with Xft fonts on
// the given screen of the given display.
func GetContext(display *xlib.Display, screen int) pango.Context {
	var arg0 *xlib.Display
	arg0 = wrapDisplay(display)

	var arg1 int
	arg1 = int(screen)

	ret := C.pango_xft_get_context(arg0, arg1)

	var ret0 pango.Context
	ret0 = wrapContext(ret)

	return ret0
}

// GetFontMap: returns the XftFontMap for the given display and screen. The
// fontmap is owned by Pango and will be valid until the display is closed.
func GetFontMap(display *xlib.Display, screen int) pango.FontMap {
	var arg0 *xlib.Display
	arg0 = wrapDisplay(display)

	var arg1 int
	arg1 = int(screen)

	ret := C.pango_xft_get_font_map(arg0, arg1)

	var ret0 pango.FontMap
	ret0 = wrapFontMap(ret)

	return ret0
}

// PictureRender: renders a GlyphString onto an Xrender Picture object.
func PictureRender(display *xlib.Display, srcPicture xlib.Picture, destPicture xlib.Picture, font pango.Font, glyphs *pango.GlyphString, x int, y int) {
	var arg0 *xlib.Display
	arg0 = wrapDisplay(display)

	var arg1 xlib.Picture
	{
		tmp := uint32(srcPicture)
		arg1 = Picture(tmp)
	}

	var arg2 xlib.Picture
	{
		tmp := uint32(destPicture)
		arg2 = Picture(tmp)
	}

	var arg3 pango.Font
	arg3 = wrapFont(font)

	var arg4 *pango.GlyphString
	arg4 = wrapGlyphString(glyphs)

	var arg5 int
	arg5 = int(x)

	var arg6 int
	arg6 = int(y)

	C.pango_xft_picture_render(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// Render: renders a GlyphString onto an XftDraw object wrapping an X drawable.
func Render(draw *xft.Draw, color *xft.Color, font pango.Font, glyphs *pango.GlyphString, x int, y int) {
	var arg0 *xft.Draw
	arg0 = wrapDraw(draw)

	var arg1 *xft.Color
	arg1 = wrapColor(color)

	var arg2 pango.Font
	arg2 = wrapFont(font)

	var arg3 *pango.GlyphString
	arg3 = wrapGlyphString(glyphs)

	var arg4 int
	arg4 = int(x)

	var arg5 int
	arg5 = int(y)

	C.pango_xft_render(arg0, arg1, arg2, arg3, arg4, arg5)
}

// RenderLayout: render a Layout onto a Draw
func RenderLayout(draw *xft.Draw, color *xft.Color, layout pango.Layout, x int, y int) {
	var arg0 *xft.Draw
	arg0 = wrapDraw(draw)

	var arg1 *xft.Color
	arg1 = wrapColor(color)

	var arg2 pango.Layout
	arg2 = wrapLayout(layout)

	var arg3 int
	arg3 = int(x)

	var arg4 int
	arg4 = int(y)

	C.pango_xft_render_layout(arg0, arg1, arg2, arg3, arg4)
}

// RenderLayoutLine: render a LayoutLine onto a Draw
func RenderLayoutLine(draw *xft.Draw, color *xft.Color, line *pango.LayoutLine, x int, y int) {
	var arg0 *xft.Draw
	arg0 = wrapDraw(draw)

	var arg1 *xft.Color
	arg1 = wrapColor(color)

	var arg2 *pango.LayoutLine
	arg2 = wrapLayoutLine(line)

	var arg3 int
	arg3 = int(x)

	var arg4 int
	arg4 = int(y)

	C.pango_xft_render_layout_line(arg0, arg1, arg2, arg3, arg4)
}

// RenderTransformed: renders a GlyphString onto a Draw, possibly transforming
// the layed-out coordinates through a transformation matrix. Note that the
// transformation matrix for @font is not changed, so to produce correct
// rendering results, the @font must have been loaded using a Context with an
// identical transformation matrix to that passed in to this function.
func RenderTransformed(draw *xft.Draw, color *xft.Color, matrix *pango.Matrix, font pango.Font, glyphs *pango.GlyphString, x int, y int) {
	var arg0 *xft.Draw
	arg0 = wrapDraw(draw)

	var arg1 *xft.Color
	arg1 = wrapColor(color)

	var arg2 *pango.Matrix
	arg2 = wrapMatrix(matrix)

	var arg3 pango.Font
	arg3 = wrapFont(font)

	var arg4 *pango.GlyphString
	arg4 = wrapGlyphString(glyphs)

	var arg5 int
	arg5 = int(x)

	var arg6 int
	arg6 = int(y)

	C.pango_xft_render_transformed(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// SetDefaultSubstitute: sets a function that will be called to do final
// configuration substitution on a Pattern before it is used to load the font.
// This function can be used to do things like set hinting and antialiasing
// options.
func SetDefaultSubstitute(display *xlib.Display, screen int, _func SubstituteFunc) {
	var arg0 *xlib.Display
	arg0 = wrapDisplay(display)

	var arg1 int
	arg1 = int(screen)

	var arg2 SubstituteFunc
	arg2 = wrapSubstituteFunc(_func)

	C.pango_xft_set_default_substitute(arg0, arg1, arg2)
}

// ShutdownDisplay: release any resources that have been cached for the
// combination of @display and @screen. Note that when the X display is closed,
// resources are released automatically, without needing to call this function.
func ShutdownDisplay(display *xlib.Display, screen int) {
	var arg0 *xlib.Display
	arg0 = wrapDisplay(display)

	var arg1 int
	arg1 = int(screen)

	C.pango_xft_shutdown_display(arg0, arg1)
}

// SubstituteChanged: call this function any time the results of the default
// substitution function set with pango_xft_set_default_substitute() change.
// That is, if your substitution function will return different results for the
// same input pattern, you must call this function.
func SubstituteChanged(display *xlib.Display, screen int) {
	var arg0 *xlib.Display
	arg0 = wrapDisplay(display)

	var arg1 int
	arg1 = int(screen)

	C.pango_xft_substitute_changed(arg0, arg1)
}

// Font is an implementation of FcFont using the Xft library for rendering. It
// is used in conjunction with XftFontMap.
type Font interface {
	pangofc.Font

	// GetDisplay: returns the X display of the `XftFont` of a font.
	GetDisplay() *xlib.Display
	// GetGlyph: gets the glyph index for a given Unicode character for @font.
	// If you only want to determine whether the font has the glyph, use
	// pango_xft_font_has_char().
	//
	// Use pango_fc_font_get_glyph() instead.
	GetGlyph(wc uint32) uint
	// GetUnknownGlyph: returns the index of a glyph suitable for drawing @wc as
	// an unknown character.
	//
	// Use PANGO_GET_UNKNOWN_GLYPH() instead.
	GetUnknownGlyph(wc uint32) pango.Glyph
	// HasChar: determines whether @font has a glyph for the codepoint @wc.
	//
	// Use pango_fc_font_has_char() instead.
	HasChar(wc uint32) bool
	// LockFace: gets the FreeType `FT_Face` associated with a font.
	//
	// This face will be kept around until you call
	// pango_xft_font_unlock_face().
	//
	// Use pango_fc_font_lock_face() instead.
	LockFace() freetype2.Face
	// UnlockFace: releases a font previously obtained with
	// pango_xft_font_lock_face().
	//
	// Use pango_fc_font_unlock_face() instead.
	UnlockFace()
}

type font struct {
	pangofc.font
}

func wrapFont(obj *externglib.Object) Font {
	return font{pangofc.font{pango.font{*externglib.Object{obj}}}}
}

func marshalFont(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

func (f font) GetDisplay() *xlib.Display

func (f font) GetGlyph(wc uint32) uint

func (f font) GetUnknownGlyph(wc uint32) pango.Glyph

func (f font) HasChar(wc uint32) bool

func (f font) LockFace() freetype2.Face

func (f font) UnlockFace()

// FontMap is an implementation of FcFontMap suitable for the Xft library as the
// renderer. It is used in to create fonts of type XftFont.
type FontMap interface {
	pangofc.FontMap
}

type fontMap struct {
	pangofc.fontMap
}

func wrapFontMap(obj *externglib.Object) FontMap {
	return fontMap{pangofc.fontMap{pango.fontMap{*externglib.Object{obj}}}}
}

func marshalFontMap(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// Renderer is a subclass of Renderer used for rendering with Pango's Xft
// backend. It can be used directly, or it can be further subclassed to modify
// exactly how drawing of individual elements occurs.
type Renderer interface {
	pango.Renderer

	// SetDefaultColor: sets the default foreground color for a Renderer.
	SetDefaultColor(defaultColor *pango.Color)
	// SetDraw: sets the Draw object that the renderer is drawing to. The
	// renderer must not be currently active.
	SetDraw(draw *xft.Draw)
}

type renderer struct {
	pango.renderer
}

func wrapRenderer(obj *externglib.Object) Renderer {
	return renderer{pango.renderer{*externglib.Object{obj}}}
}

func marshalRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

func NewRenderer(display *xlib.Display, screen int) Renderer

func (r renderer) SetDefaultColor(defaultColor *pango.Color)

func (r renderer) SetDraw(draw *xft.Draw)
