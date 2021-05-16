// Code generated by girgen. DO NOT EDIT.

package pangoxft

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/pango"
	"github.com/diamondburned/gotk4/pkg/pangofc"
	"github.com/diamondburned/gotk4/pkg/xft"
	"github.com/diamondburned/gotk4/pkg/xlib"
	"github.com/gotk3/gotk3/glib"
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

// GetContext: retrieves a Context appropriate for rendering with Xft fonts on
// the given screen of the given display.
func GetContext(display *xlib.Display, screen int) *pango.Context

// GetFontMap: returns the XftFontMap for the given display and screen. The
// fontmap is owned by Pango and will be valid until the display is closed.
func GetFontMap(display *xlib.Display, screen int) *pango.FontMap

// PictureRender: renders a GlyphString onto an Xrender Picture object.
func PictureRender(display *xlib.Display, srcPicture xlib.Picture, destPicture xlib.Picture, font *pango.Font, glyphs *pango.GlyphString, x int, y int)

// Render: renders a GlyphString onto an XftDraw object wrapping an X drawable.
func Render(draw *xft.Draw, color *xft.Color, font *pango.Font, glyphs *pango.GlyphString, x int, y int)

// RenderLayout: render a Layout onto a Draw
func RenderLayout(draw *xft.Draw, color *xft.Color, layout *pango.Layout, x int, y int)

// RenderLayoutLine: render a LayoutLine onto a Draw
func RenderLayoutLine(draw *xft.Draw, color *xft.Color, line *pango.LayoutLine, x int, y int)

// RenderTransformed: renders a GlyphString onto a Draw, possibly transforming
// the layed-out coordinates through a transformation matrix. Note that the
// transformation matrix for @font is not changed, so to produce correct
// rendering results, the @font must have been loaded using a Context with an
// identical transformation matrix to that passed in to this function.
func RenderTransformed(draw *xft.Draw, color *xft.Color, matrix *pango.Matrix, font *pango.Font, glyphs *pango.GlyphString, x int, y int)

// SetDefaultSubstitute: sets a function that will be called to do final
// configuration substitution on a Pattern before it is used to load the font.
// This function can be used to do things like set hinting and antialiasing
// options.
func SetDefaultSubstitute(display *xlib.Display, screen int, _func SubstituteFunc, data unsafe.Pointer, notify unsafe.Pointer)

// ShutdownDisplay: release any resources that have been cached for the
// combination of @display and @screen. Note that when the X display is closed,
// resources are released automatically, without needing to call this function.
func ShutdownDisplay(display *xlib.Display, screen int)

// SubstituteChanged: call this function any time the results of the default
// substitution function set with pango_xft_set_default_substitute() change.
// That is, if your substitution function will return different results for the
// same input pattern, you must call this function.
func SubstituteChanged(display *xlib.Display, screen int)

// Font: pangoXftFont is an implementation of FcFont using the Xft library for
// rendering. It is used in conjunction with XftFontMap.
type Font struct {
	pangofc.Font
}

func wrapFont(obj *glib.Object) *Font {
	return &Font{Font{Font{*externglib.Object{obj}}}}
}

func marshalFont(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// FontMap: pangoXftFontMap is an implementation of FcFontMap suitable for the
// Xft library as the renderer. It is used in to create fonts of type XftFont.
type FontMap struct {
	pangofc.FontMap
}

func wrapFontMap(obj *glib.Object) *FontMap {
	return &FontMap{FontMap{FontMap{*externglib.Object{obj}}}}
}

func marshalFontMap(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// Renderer: pangoXftRenderer is a subclass of Renderer used for rendering with
// Pango's Xft backend. It can be used directly, or it can be further subclassed
// to modify exactly how drawing of individual elements occurs.
type Renderer struct {
	pango.Renderer
}

func wrapRenderer(obj *glib.Object) *Renderer {
	return &Renderer{Renderer{*externglib.Object{obj}}}
}

func marshalRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}
