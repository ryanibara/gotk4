// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_renderer_get_type()), F: marshalRenderer},
	})
}

// Renderer: `PangoRenderer` is a base class for objects that can render text
// provided as `PangoGlyphString` or `PangoLayout`.
//
// By subclassing `PangoRenderer` and overriding operations such as @draw_glyphs
// and @draw_rectangle, renderers for particular font backends and destinations
// can be created.
type Renderer interface {
	gextras.Objector

	// Activate does initial setup before rendering operations on @renderer.
	//
	// [method@Pango.Renderer.deactivate] should be called when done drawing.
	// Calls such as [method@Pango.Renderer.draw_layout] automatically activate
	// the layout before drawing on it. Calls to `pango_renderer_activate()` and
	// `pango_renderer_deactivate()` can be nested and the renderer will only be
	// initialized and deinitialized once.
	Activate()
	// Deactivate cleans up after rendering operations on @renderer.
	//
	// See docs for [method@Pango.Renderer.activate].
	Deactivate()
	// DrawErrorUnderline: draw a squiggly line that approximately covers the
	// given rectangle in the style of an underline used to indicate a spelling
	// error.
	//
	// The width of the underline is rounded to an integer number of up/down
	// segments and the resulting rectangle is centered in the original
	// rectangle.
	//
	// This should be called while @renderer is already active. Use
	// [method@Pango.Renderer.activate] to activate a renderer.
	DrawErrorUnderline(x int, y int, width int, height int)
	// DrawGlyphItem draws the glyphs in @glyph_item with the specified
	// `PangoRenderer`, embedding the text associated with the glyphs in the
	// output if the output format supports it.
	//
	// This is useful for rendering text in PDF.
	//
	// Note that @text is the start of the text for layout, which is then
	// indexed by `glyph_item->item->offset`.
	//
	// If @text is nil, this simply calls [method@Pango.Renderer.draw_glyphs].
	//
	// The default implementation of this method simply falls back to
	// [method@Pango.Renderer.draw_glyphs].
	DrawGlyphItem(text string, glyphItem *GlyphItem, x int, y int)
	// DrawGlyphs draws the glyphs in @glyphs with the specified
	// `PangoRenderer`.
	DrawGlyphs(font Font, glyphs *GlyphString, x int, y int)
	// DrawLayout draws @layout with the specified `PangoRenderer`.
	DrawLayout(layout Layout, x int, y int)
	// DrawLayoutLine draws @line with the specified `PangoRenderer`.
	DrawLayoutLine(line *LayoutLine, x int, y int)
	// DrawRectangle draws an axis-aligned rectangle in user space coordinates
	// with the specified `PangoRenderer`.
	//
	// This should be called while @renderer is already active. Use
	// [method@Pango.Renderer.activate] to activate a renderer.
	DrawRectangle(part RenderPart, x int, y int, width int, height int)
	// DrawTrapezoid draws a trapezoid with the parallel sides aligned with the
	// X axis using the given `PangoRenderer`; coordinates are in device space.
	DrawTrapezoid(part RenderPart, y1 float64, x11 float64, x21 float64, y2 float64, x12 float64, x22 float64)
	// Alpha gets the current alpha for the specified part.
	Alpha(part RenderPart) uint16
	// Color gets the current rendering color for the specified part.
	Color(part RenderPart) *Color
	// Layout gets the layout currently being rendered using @renderer.
	//
	// Calling this function only makes sense from inside a subclass's methods,
	// like in its draw_shape vfunc, for example.
	//
	// The returned layout should not be modified while still being rendered.
	Layout() Layout
	// LayoutLine gets the layout line currently being rendered using @renderer.
	//
	// Calling this function only makes sense from inside a subclass's methods,
	// like in its draw_shape vfunc, for example.
	//
	// The returned layout line should not be modified while still being
	// rendered.
	LayoutLine() *LayoutLine
	// Matrix gets the transformation matrix that will be applied when
	// rendering.
	//
	// See [method@Pango.Renderer.set_matrix].
	Matrix() *Matrix
	// PartChanged informs Pango that the way that the rendering is done for
	// @part has changed.
	//
	// This should be called if the rendering changes in a way that would
	// prevent multiple pieces being joined together into one drawing call. For
	// instance, if a subclass of `PangoRenderer` was to add a stipple option
	// for drawing underlines, it needs to call
	//
	// “` pango_renderer_part_changed (render, PANGO_RENDER_PART_UNDERLINE); “`
	//
	// When the stipple changes or underlines with different stipples might be
	// joined together. Pango automatically calls this for changes to colors.
	// (See [method@Pango.Renderer.set_color])
	PartChanged(part RenderPart)
	// SetAlpha sets the alpha for part of the rendering.
	//
	// Note that the alpha may only be used if a color is specified for @part as
	// well.
	SetAlpha(part RenderPart, alpha uint16)
	// SetColor sets the color for part of the rendering.
	//
	// Also see [method@Pango.Renderer.set_alpha].
	SetColor(part RenderPart, color *Color)
	// SetMatrix sets the transformation matrix that will be applied when
	// rendering.
	SetMatrix(matrix *Matrix)
}

// renderer implements the Renderer interface.
type renderer struct {
	gextras.Objector
}

var _ Renderer = (*renderer)(nil)

// WrapRenderer wraps a GObject to the right type. It is
// primarily used internally.
func WrapRenderer(obj *externglib.Object) Renderer {
	return Renderer{
		Objector: obj,
	}
}

func marshalRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRenderer(obj), nil
}

// Activate does initial setup before rendering operations on @renderer.
//
// [method@Pango.Renderer.deactivate] should be called when done drawing.
// Calls such as [method@Pango.Renderer.draw_layout] automatically activate
// the layout before drawing on it. Calls to `pango_renderer_activate()` and
// `pango_renderer_deactivate()` can be nested and the renderer will only be
// initialized and deinitialized once.
func (r renderer) Activate() {
	var _arg0 *C.PangoRenderer

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))

	C.pango_renderer_activate(_arg0)
}

// Deactivate cleans up after rendering operations on @renderer.
//
// See docs for [method@Pango.Renderer.activate].
func (r renderer) Deactivate() {
	var _arg0 *C.PangoRenderer

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))

	C.pango_renderer_deactivate(_arg0)
}

// DrawErrorUnderline: draw a squiggly line that approximately covers the
// given rectangle in the style of an underline used to indicate a spelling
// error.
//
// The width of the underline is rounded to an integer number of up/down
// segments and the resulting rectangle is centered in the original
// rectangle.
//
// This should be called while @renderer is already active. Use
// [method@Pango.Renderer.activate] to activate a renderer.
func (r renderer) DrawErrorUnderline(x int, y int, width int, height int) {
	var _arg0 *C.PangoRenderer
	var _arg1 C.int
	var _arg2 C.int
	var _arg3 C.int
	var _arg4 C.int

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = C.int(x)
	_arg2 = C.int(y)
	_arg3 = C.int(width)
	_arg4 = C.int(height)

	C.pango_renderer_draw_error_underline(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// DrawGlyphItem draws the glyphs in @glyph_item with the specified
// `PangoRenderer`, embedding the text associated with the glyphs in the
// output if the output format supports it.
//
// This is useful for rendering text in PDF.
//
// Note that @text is the start of the text for layout, which is then
// indexed by `glyph_item->item->offset`.
//
// If @text is nil, this simply calls [method@Pango.Renderer.draw_glyphs].
//
// The default implementation of this method simply falls back to
// [method@Pango.Renderer.draw_glyphs].
func (r renderer) DrawGlyphItem(text string, glyphItem *GlyphItem, x int, y int) {
	var _arg0 *C.PangoRenderer
	var _arg1 *C.char
	var _arg2 *C.PangoGlyphItem
	var _arg3 C.int
	var _arg4 C.int

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.PangoGlyphItem)(unsafe.Pointer(glyphItem.Native()))
	_arg3 = C.int(x)
	_arg4 = C.int(y)

	C.pango_renderer_draw_glyph_item(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// DrawGlyphs draws the glyphs in @glyphs with the specified
// `PangoRenderer`.
func (r renderer) DrawGlyphs(font Font, glyphs *GlyphString, x int, y int) {
	var _arg0 *C.PangoRenderer
	var _arg1 *C.PangoFont
	var _arg2 *C.PangoGlyphString
	var _arg3 C.int
	var _arg4 C.int

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.PangoFont)(unsafe.Pointer(font.Native()))
	_arg2 = (*C.PangoGlyphString)(unsafe.Pointer(glyphs.Native()))
	_arg3 = C.int(x)
	_arg4 = C.int(y)

	C.pango_renderer_draw_glyphs(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// DrawLayout draws @layout with the specified `PangoRenderer`.
func (r renderer) DrawLayout(layout Layout, x int, y int) {
	var _arg0 *C.PangoRenderer
	var _arg1 *C.PangoLayout
	var _arg2 C.int
	var _arg3 C.int

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))
	_arg2 = C.int(x)
	_arg3 = C.int(y)

	C.pango_renderer_draw_layout(_arg0, _arg1, _arg2, _arg3)
}

// DrawLayoutLine draws @line with the specified `PangoRenderer`.
func (r renderer) DrawLayoutLine(line *LayoutLine, x int, y int) {
	var _arg0 *C.PangoRenderer
	var _arg1 *C.PangoLayoutLine
	var _arg2 C.int
	var _arg3 C.int

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.PangoLayoutLine)(unsafe.Pointer(line.Native()))
	_arg2 = C.int(x)
	_arg3 = C.int(y)

	C.pango_renderer_draw_layout_line(_arg0, _arg1, _arg2, _arg3)
}

// DrawRectangle draws an axis-aligned rectangle in user space coordinates
// with the specified `PangoRenderer`.
//
// This should be called while @renderer is already active. Use
// [method@Pango.Renderer.activate] to activate a renderer.
func (r renderer) DrawRectangle(part RenderPart, x int, y int, width int, height int) {
	var _arg0 *C.PangoRenderer
	var _arg1 C.PangoRenderPart
	var _arg2 C.int
	var _arg3 C.int
	var _arg4 C.int
	var _arg5 C.int

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (C.PangoRenderPart)(part)
	_arg2 = C.int(x)
	_arg3 = C.int(y)
	_arg4 = C.int(width)
	_arg5 = C.int(height)

	C.pango_renderer_draw_rectangle(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// DrawTrapezoid draws a trapezoid with the parallel sides aligned with the
// X axis using the given `PangoRenderer`; coordinates are in device space.
func (r renderer) DrawTrapezoid(part RenderPart, y1 float64, x11 float64, x21 float64, y2 float64, x12 float64, x22 float64) {
	var _arg0 *C.PangoRenderer
	var _arg1 C.PangoRenderPart
	var _arg2 C.double
	var _arg3 C.double
	var _arg4 C.double
	var _arg5 C.double
	var _arg6 C.double
	var _arg7 C.double

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (C.PangoRenderPart)(part)
	_arg2 = C.double(y1)
	_arg3 = C.double(x11)
	_arg4 = C.double(x21)
	_arg5 = C.double(y2)
	_arg6 = C.double(x12)
	_arg7 = C.double(x22)

	C.pango_renderer_draw_trapezoid(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
}

// Alpha gets the current alpha for the specified part.
func (r renderer) Alpha(part RenderPart) uint16 {
	var _arg0 *C.PangoRenderer
	var _arg1 C.PangoRenderPart

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (C.PangoRenderPart)(part)

	var _cret C.guint16

	cret = C.pango_renderer_get_alpha(_arg0, _arg1)

	var _guint16 uint16

	_guint16 = (uint16)(_cret)

	return _guint16
}

// Color gets the current rendering color for the specified part.
func (r renderer) Color(part RenderPart) *Color {
	var _arg0 *C.PangoRenderer
	var _arg1 C.PangoRenderPart

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (C.PangoRenderPart)(part)

	var _cret *C.PangoColor

	cret = C.pango_renderer_get_color(_arg0, _arg1)

	var _color *Color

	_color = WrapColor(unsafe.Pointer(_cret))

	return _color
}

// Layout gets the layout currently being rendered using @renderer.
//
// Calling this function only makes sense from inside a subclass's methods,
// like in its draw_shape vfunc, for example.
//
// The returned layout should not be modified while still being rendered.
func (r renderer) Layout() Layout {
	var _arg0 *C.PangoRenderer

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))

	var _cret *C.PangoLayout

	cret = C.pango_renderer_get_layout(_arg0)

	var _layout Layout

	_layout = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Layout)

	return _layout
}

// LayoutLine gets the layout line currently being rendered using @renderer.
//
// Calling this function only makes sense from inside a subclass's methods,
// like in its draw_shape vfunc, for example.
//
// The returned layout line should not be modified while still being
// rendered.
func (r renderer) LayoutLine() *LayoutLine {
	var _arg0 *C.PangoRenderer

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))

	var _cret *C.PangoLayoutLine

	cret = C.pango_renderer_get_layout_line(_arg0)

	var _layoutLine *LayoutLine

	_layoutLine = WrapLayoutLine(unsafe.Pointer(_cret))

	return _layoutLine
}

// Matrix gets the transformation matrix that will be applied when
// rendering.
//
// See [method@Pango.Renderer.set_matrix].
func (r renderer) Matrix() *Matrix {
	var _arg0 *C.PangoRenderer

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))

	var _cret *C.PangoMatrix

	cret = C.pango_renderer_get_matrix(_arg0)

	var _matrix *Matrix

	_matrix = WrapMatrix(unsafe.Pointer(_cret))

	return _matrix
}

// PartChanged informs Pango that the way that the rendering is done for
// @part has changed.
//
// This should be called if the rendering changes in a way that would
// prevent multiple pieces being joined together into one drawing call. For
// instance, if a subclass of `PangoRenderer` was to add a stipple option
// for drawing underlines, it needs to call
//
// “` pango_renderer_part_changed (render, PANGO_RENDER_PART_UNDERLINE); “`
//
// When the stipple changes or underlines with different stipples might be
// joined together. Pango automatically calls this for changes to colors.
// (See [method@Pango.Renderer.set_color])
func (r renderer) PartChanged(part RenderPart) {
	var _arg0 *C.PangoRenderer
	var _arg1 C.PangoRenderPart

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (C.PangoRenderPart)(part)

	C.pango_renderer_part_changed(_arg0, _arg1)
}

// SetAlpha sets the alpha for part of the rendering.
//
// Note that the alpha may only be used if a color is specified for @part as
// well.
func (r renderer) SetAlpha(part RenderPart, alpha uint16) {
	var _arg0 *C.PangoRenderer
	var _arg1 C.PangoRenderPart
	var _arg2 C.guint16

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (C.PangoRenderPart)(part)
	_arg2 = C.guint16(alpha)

	C.pango_renderer_set_alpha(_arg0, _arg1, _arg2)
}

// SetColor sets the color for part of the rendering.
//
// Also see [method@Pango.Renderer.set_alpha].
func (r renderer) SetColor(part RenderPart, color *Color) {
	var _arg0 *C.PangoRenderer
	var _arg1 C.PangoRenderPart
	var _arg2 *C.PangoColor

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (C.PangoRenderPart)(part)
	_arg2 = (*C.PangoColor)(unsafe.Pointer(color.Native()))

	C.pango_renderer_set_color(_arg0, _arg1, _arg2)
}

// SetMatrix sets the transformation matrix that will be applied when
// rendering.
func (r renderer) SetMatrix(matrix *Matrix) {
	var _arg0 *C.PangoRenderer
	var _arg1 *C.PangoMatrix

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.PangoMatrix)(unsafe.Pointer(matrix.Native()))

	C.pango_renderer_set_matrix(_arg0, _arg1)
}