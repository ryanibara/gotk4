// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

// glib.Type values for pango-glyph.go.
var (
	GTypeShapeFlags  = coreglib.Type(C.pango_shape_flags_get_type())
	GTypeGlyphString = coreglib.Type(C.pango_glyph_string_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeShapeFlags, F: marshalShapeFlags},
		{T: GTypeGlyphString, F: marshalGlyphString},
	})
}

// GlyphUnit: PangoGlyphUnit type is used to store dimensions within Pango.
//
// Dimensions are stored in 1/PANGO_SCALE of a device unit. (A device unit might
// be a pixel for screen display, or a point on a printer.) PANGO_SCALE is
// currently 1024, and may change in the future (unlikely though), but you
// should not depend on its exact value. The PANGO_PIXELS() macro can be used to
// convert from glyph units into device units with correct rounding.
type GlyphUnit = int32

// ShapeFlags flags influencing the shaping process.
//
// PangoShapeFlags can be passed to pango_shape_with_flags().
type ShapeFlags C.guint

const (
	// ShapeNone: default value.
	ShapeNone ShapeFlags = 0b0
	// ShapeRoundPositions: round glyph positions and widths to whole device
	// units. This option should be set if the target renderer can't do subpixel
	// positioning of glyphs.
	ShapeRoundPositions ShapeFlags = 0b1
)

func marshalShapeFlags(p uintptr) (interface{}, error) {
	return ShapeFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for ShapeFlags.
func (s ShapeFlags) String() string {
	if s == 0 {
		return "ShapeFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(29)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case ShapeNone:
			builder.WriteString("None|")
		case ShapeRoundPositions:
			builder.WriteString("RoundPositions|")
		default:
			builder.WriteString(fmt.Sprintf("ShapeFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s ShapeFlags) Has(other ShapeFlags) bool {
	return (s & other) == other
}

// ReorderItems: reorder items from logical order to visual order.
//
// The visual order is determined from the associated directional levels of the
// items. The original list is unmodified.
//
// The function takes the following parameters:
//
//    - logicalItems: GList of PangoItem in logical order.
//
// The function returns the following values:
//
//    - list: #GList of Item structures in visual order.
//
//      (Please open a bug if you use this function. It is not a particularly
//      convenient interface, and the code is duplicated elsewhere in Pango for
//      that reason.).
//
func ReorderItems(logicalItems []*Item) []*Item {
	var _arg1 *C.GList // out
	var _cret *C.GList // in

	for i := len(logicalItems) - 1; i >= 0; i-- {
		src := logicalItems[i]
		var dst *C.PangoItem // out
		dst = (*C.PangoItem)(gextras.StructNative(unsafe.Pointer(src)))
		_arg1 = C.g_list_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
	}
	defer C.g_list_free(_arg1)

	_cret = C.pango_reorder_items(_arg1)
	runtime.KeepAlive(logicalItems)

	var _list []*Item // out

	_list = make([]*Item, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.PangoItem)(v)
		var dst *Item // out
		dst = (*Item)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_item_free((*C.PangoItem)(intern.C))
			},
		)
		_list = append(_list, dst)
	})

	return _list
}

// Shape: convert the characters in text into glyphs.
//
// Given a segment of text and the corresponding PangoAnalysis structure
// returned from itemize, convert the characters into glyphs. You may also pass
// in only a substring of the item from itemize.
//
// It is recommended that you use shape_full instead, since that API allows for
// shaping interaction happening across text item boundaries.
//
// Note that the extra attributes in the analyis that is returned from itemize
// have indices that are relative to the entire paragraph, so you need to
// subtract the item offset from their indices before calling shape.
//
// The function takes the following parameters:
//
//    - text to process.
//    - length (in bytes) of text.
//    - analysis: PangoAnalysis structure from itemize.
//    - glyphs: glyph string in which to store results.
//
func Shape(text string, length int32, analysis *Analysis, glyphs *GlyphString) {
	var _arg1 *C.char             // out
	var _arg2 C.int               // out
	var _arg3 *C.PangoAnalysis    // out
	var _arg4 *C.PangoGlyphString // out

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(length)
	_arg3 = (*C.PangoAnalysis)(gextras.StructNative(unsafe.Pointer(analysis)))
	_arg4 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))

	C.pango_shape(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(text)
	runtime.KeepAlive(length)
	runtime.KeepAlive(analysis)
	runtime.KeepAlive(glyphs)
}

// ShapeFull: convert the characters in text into glyphs.
//
// Given a segment of text and the corresponding PangoAnalysis structure
// returned from itemize, convert the characters into glyphs. You may also pass
// in only a substring of the item from itemize.
//
// This is similar to shape, except it also can optionally take the full
// paragraph text as input, which will then be used to perform certain
// cross-item shaping interactions. If you have access to the broader text of
// which item_text is part of, provide the broader text as paragraph_text. If
// paragraph_text is NULL, item text is used instead.
//
// Note that the extra attributes in the analyis that is returned from itemize
// have indices that are relative to the entire paragraph, so you do not pass
// the full paragraph text as paragraph_text, you need to subtract the item
// offset from their indices before calling shape_full.
//
// The function takes the following parameters:
//
//    - itemText: valid UTF-8 text to shape.
//    - itemLength: length (in bytes) of item_text. -1 means nul-terminated text.
//    - paragraphText (optional): text of the paragraph (see details). May be
//      NULL.
//    - paragraphLength: length (in bytes) of paragraph_text. -1 means
//      nul-terminated text.
//    - analysis: PangoAnalysis structure from itemize.
//    - glyphs: glyph string in which to store results.
//
func ShapeFull(itemText string, itemLength int32, paragraphText string, paragraphLength int32, analysis *Analysis, glyphs *GlyphString) {
	var _arg1 *C.char             // out
	var _arg2 C.int               // out
	var _arg3 *C.char             // out
	var _arg4 C.int               // out
	var _arg5 *C.PangoAnalysis    // out
	var _arg6 *C.PangoGlyphString // out

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(itemText)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(itemLength)
	if paragraphText != "" {
		_arg3 = (*C.char)(unsafe.Pointer(C.CString(paragraphText)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	_arg4 = C.int(paragraphLength)
	_arg5 = (*C.PangoAnalysis)(gextras.StructNative(unsafe.Pointer(analysis)))
	_arg6 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))

	C.pango_shape_full(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(itemText)
	runtime.KeepAlive(itemLength)
	runtime.KeepAlive(paragraphText)
	runtime.KeepAlive(paragraphLength)
	runtime.KeepAlive(analysis)
	runtime.KeepAlive(glyphs)
}

// ShapeWithFlags: convert the characters in text into glyphs.
//
// Given a segment of text and the corresponding PangoAnalysis structure
// returned from itemize, convert the characters into glyphs. You may also pass
// in only a substring of the item from itemize.
//
// This is similar to shape_full, except it also takes flags that can influence
// the shaping process.
//
// Note that the extra attributes in the analyis that is returned from itemize
// have indices that are relative to the entire paragraph, so you do not pass
// the full paragraph text as paragraph_text, you need to subtract the item
// offset from their indices before calling shape_with_flags.
//
// The function takes the following parameters:
//
//    - itemText: valid UTF-8 text to shape.
//    - itemLength: length (in bytes) of item_text. -1 means nul-terminated text.
//    - paragraphText (optional): text of the paragraph (see details). May be
//      NULL.
//    - paragraphLength: length (in bytes) of paragraph_text. -1 means
//      nul-terminated text.
//    - analysis: PangoAnalysis structure from itemize.
//    - glyphs: glyph string in which to store results.
//    - flags influencing the shaping process.
//
func ShapeWithFlags(itemText string, itemLength int32, paragraphText string, paragraphLength int32, analysis *Analysis, glyphs *GlyphString, flags ShapeFlags) {
	var _arg1 *C.char             // out
	var _arg2 C.int               // out
	var _arg3 *C.char             // out
	var _arg4 C.int               // out
	var _arg5 *C.PangoAnalysis    // out
	var _arg6 *C.PangoGlyphString // out
	var _arg7 C.PangoShapeFlags   // out

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(itemText)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(itemLength)
	if paragraphText != "" {
		_arg3 = (*C.char)(unsafe.Pointer(C.CString(paragraphText)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	_arg4 = C.int(paragraphLength)
	_arg5 = (*C.PangoAnalysis)(gextras.StructNative(unsafe.Pointer(analysis)))
	_arg6 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))
	_arg7 = C.PangoShapeFlags(flags)

	C.pango_shape_with_flags(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(itemText)
	runtime.KeepAlive(itemLength)
	runtime.KeepAlive(paragraphText)
	runtime.KeepAlive(paragraphLength)
	runtime.KeepAlive(analysis)
	runtime.KeepAlive(glyphs)
	runtime.KeepAlive(flags)
}

// GlyphGeometry: PangoGlyphGeometry structure contains width and positioning
// information for a single glyph.
//
// An instance of this type is always passed by reference.
type GlyphGeometry struct {
	*glyphGeometry
}

// glyphGeometry is the struct that's finalized.
type glyphGeometry struct {
	native *C.PangoGlyphGeometry
}

// Width: logical width to use for the the character.
func (g *GlyphGeometry) Width() GlyphUnit {
	var v GlyphUnit // out
	v = int32(g.native.width)
	return v
}

// XOffset: horizontal offset from nominal character position.
func (g *GlyphGeometry) XOffset() GlyphUnit {
	var v GlyphUnit // out
	v = int32(g.native.x_offset)
	return v
}

// YOffset: vertical offset from nominal character position.
func (g *GlyphGeometry) YOffset() GlyphUnit {
	var v GlyphUnit // out
	v = int32(g.native.y_offset)
	return v
}

// GlyphInfo: PangoGlyphInfo structure represents a single glyph with
// positioning information and visual attributes.
//
// An instance of this type is always passed by reference.
type GlyphInfo struct {
	*glyphInfo
}

// glyphInfo is the struct that's finalized.
type glyphInfo struct {
	native *C.PangoGlyphInfo
}

// Glyph: glyph itself.
func (g *GlyphInfo) Glyph() Glyph {
	var v Glyph // out
	v = uint32(g.native.glyph)
	return v
}

// Geometry: positional information about the glyph.
func (g *GlyphInfo) Geometry() *GlyphGeometry {
	var v *GlyphGeometry // out
	v = (*GlyphGeometry)(gextras.NewStructNative(unsafe.Pointer((&g.native.geometry))))
	return v
}

// Attr: visual attributes of the glyph.
func (g *GlyphInfo) Attr() *GlyphVisAttr {
	var v *GlyphVisAttr // out
	v = (*GlyphVisAttr)(gextras.NewStructNative(unsafe.Pointer((&g.native.attr))))
	return v
}

// GlyphString: PangoGlyphString is used to store strings of glyphs with
// geometry and visual attribute information.
//
// The storage for the glyph information is owned by the structure which
// simplifies memory management.
//
// An instance of this type is always passed by reference.
type GlyphString struct {
	*glyphString
}

// glyphString is the struct that's finalized.
type glyphString struct {
	native *C.PangoGlyphString
}

func marshalGlyphString(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &GlyphString{&glyphString{(*C.PangoGlyphString)(b)}}, nil
}

// NewGlyphString constructs a struct GlyphString.
func NewGlyphString() *GlyphString {
	var _cret *C.PangoGlyphString // in

	_cret = C.pango_glyph_string_new()

	var _glyphString *GlyphString // out

	_glyphString = (*GlyphString)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_glyphString)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_glyph_string_free((*C.PangoGlyphString)(intern.C))
		},
	)

	return _glyphString
}

// Copy a glyph string and associated storage.
//
// The function returns the following values:
//
//    - glyphString (optional): newly allocated PangoGlyphString, which should be
//      freed with pango.GlyphString.Free(), or NULL if string was NULL.
//
func (str *GlyphString) Copy() *GlyphString {
	var _arg0 *C.PangoGlyphString // out
	var _cret *C.PangoGlyphString // in

	if str != nil {
		_arg0 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(str)))
	}

	_cret = C.pango_glyph_string_copy(_arg0)
	runtime.KeepAlive(str)

	var _glyphString *GlyphString // out

	if _cret != nil {
		_glyphString = (*GlyphString)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_glyphString)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_glyph_string_free((*C.PangoGlyphString)(intern.C))
			},
		)
	}

	return _glyphString
}

// Extents: compute the logical and ink extents of a glyph string.
//
// See the documentation for pango.Font.GetGlyphExtents() for details about the
// interpretation of the rectangles.
//
// Examples of logical (red) and ink (green) rects:
//
// ! (rects1.png) ! (rects2.png).
//
// The function takes the following parameters:
//
//    - font: PangoFont.
//
// The function returns the following values:
//
//    - inkRect (optional): rectangle used to store the extents of the glyph
//      string as drawn or NULL to indicate that the result is not needed.
//    - logicalRect (optional): rectangle used to store the logical extents of
//      the glyph string or NULL to indicate that the result is not needed.
//
func (glyphs *GlyphString) Extents(font Fonter) (inkRect *Rectangle, logicalRect *Rectangle) {
	var _arg0 *C.PangoGlyphString // out
	var _arg1 *C.PangoFont        // out
	var _arg2 C.PangoRectangle    // in
	var _arg3 C.PangoRectangle    // in

	_arg0 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))
	_arg1 = (*C.PangoFont)(unsafe.Pointer(coreglib.InternObject(font).Native()))

	C.pango_glyph_string_extents(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(glyphs)
	runtime.KeepAlive(font)

	var _inkRect *Rectangle     // out
	var _logicalRect *Rectangle // out

	_inkRect = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	_logicalRect = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _inkRect, _logicalRect
}

// ExtentsRange computes the extents of a sub-portion of a glyph string.
//
// The extents are relative to the start of the glyph string range (the origin
// of their coordinate system is at the start of the range, not at the start of
// the entire glyph string).
//
// The function takes the following parameters:
//
//    - start index.
//    - end index (the range is the set of bytes with indices such that start <=
//      index < end).
//    - font: PangoFont.
//
// The function returns the following values:
//
//    - inkRect (optional): rectangle used to store the extents of the glyph
//      string range as drawn or NULL to indicate that the result is not needed.
//    - logicalRect (optional): rectangle used to store the logical extents of
//      the glyph string range or NULL to indicate that the result is not needed.
//
func (glyphs *GlyphString) ExtentsRange(start int32, end int32, font Fonter) (inkRect *Rectangle, logicalRect *Rectangle) {
	var _arg0 *C.PangoGlyphString // out
	var _arg1 C.int               // out
	var _arg2 C.int               // out
	var _arg3 *C.PangoFont        // out
	var _arg4 C.PangoRectangle    // in
	var _arg5 C.PangoRectangle    // in

	_arg0 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))
	_arg1 = C.int(start)
	_arg2 = C.int(end)
	_arg3 = (*C.PangoFont)(unsafe.Pointer(coreglib.InternObject(font).Native()))

	C.pango_glyph_string_extents_range(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5)
	runtime.KeepAlive(glyphs)
	runtime.KeepAlive(start)
	runtime.KeepAlive(end)
	runtime.KeepAlive(font)

	var _inkRect *Rectangle     // out
	var _logicalRect *Rectangle // out

	_inkRect = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg4))))
	_logicalRect = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg5))))

	return _inkRect, _logicalRect
}

// Width computes the logical width of the glyph string.
//
// This can also be computed using pango.GlyphString.Extents(). However, since
// this only computes the width, it's much faster. This is in fact only a
// convenience function that computes the sum of geometry.width for each glyph
// in the glyphs.
//
// The function returns the following values:
//
//    - gint: logical width of the glyph string.
//
func (glyphs *GlyphString) Width() int32 {
	var _arg0 *C.PangoGlyphString // out
	var _cret C.int               // in

	_arg0 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))

	_cret = C.pango_glyph_string_get_width(_arg0)
	runtime.KeepAlive(glyphs)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// IndexToX converts from character position to x position.
//
// The X position is measured from the left edge of the run. Character positions
// are computed by dividing up each cluster into equal portions.
//
// The function takes the following parameters:
//
//    - text for the run.
//    - length: number of bytes (not characters) in text.
//    - analysis information return from itemize.
//    - index_: byte index within text.
//    - trailing: whether we should compute the result for the beginning (FALSE)
//      or end (TRUE) of the character.
//
// The function returns the following values:
//
//    - xPos: location to store result.
//
func (glyphs *GlyphString) IndexToX(text string, length int32, analysis *Analysis, index_ int32, trailing bool) int32 {
	var _arg0 *C.PangoGlyphString // out
	var _arg1 *C.char             // out
	var _arg2 C.int               // out
	var _arg3 *C.PangoAnalysis    // out
	var _arg4 C.int               // out
	var _arg5 C.gboolean          // out
	var _arg6 C.int               // in

	_arg0 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(length)
	_arg3 = (*C.PangoAnalysis)(gextras.StructNative(unsafe.Pointer(analysis)))
	_arg4 = C.int(index_)
	if trailing {
		_arg5 = C.TRUE
	}

	C.pango_glyph_string_index_to_x(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, &_arg6)
	runtime.KeepAlive(glyphs)
	runtime.KeepAlive(text)
	runtime.KeepAlive(length)
	runtime.KeepAlive(analysis)
	runtime.KeepAlive(index_)
	runtime.KeepAlive(trailing)

	var _xPos int32 // out

	_xPos = int32(_arg6)

	return _xPos
}

// SetSize: resize a glyph string to the given length.
//
// The function takes the following parameters:
//
//    - newLen: new length of the string.
//
func (str *GlyphString) SetSize(newLen int32) {
	var _arg0 *C.PangoGlyphString // out
	var _arg1 C.gint              // out

	_arg0 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(str)))
	_arg1 = C.gint(newLen)

	C.pango_glyph_string_set_size(_arg0, _arg1)
	runtime.KeepAlive(str)
	runtime.KeepAlive(newLen)
}

// XToIndex: convert from x offset to character position.
//
// Character positions are computed by dividing up each cluster into equal
// portions. In scripts where positioning within a cluster is not allowed (such
// as Thai), the returned value may not be a valid cursor position; the caller
// must combine the result with the logical attributes for the text to compute
// the valid cursor position.
//
// The function takes the following parameters:
//
//    - text for the run.
//    - length: number of bytes (not characters) in text.
//    - analysis information return from itemize.
//    - xPos: x offset (in Pango units).
//
// The function returns the following values:
//
//    - index_: location to store calculated byte index within text.
//    - trailing: location to store a boolean indicating whether the user clicked
//      on the leading or trailing edge of the character.
//
func (glyphs *GlyphString) XToIndex(text string, length int32, analysis *Analysis, xPos int32) (index_ int32, trailing int32) {
	var _arg0 *C.PangoGlyphString // out
	var _arg1 *C.char             // out
	var _arg2 C.int               // out
	var _arg3 *C.PangoAnalysis    // out
	var _arg4 C.int               // out
	var _arg5 C.int               // in
	var _arg6 C.int               // in

	_arg0 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(length)
	_arg3 = (*C.PangoAnalysis)(gextras.StructNative(unsafe.Pointer(analysis)))
	_arg4 = C.int(xPos)

	C.pango_glyph_string_x_to_index(_arg0, _arg1, _arg2, _arg3, _arg4, &_arg5, &_arg6)
	runtime.KeepAlive(glyphs)
	runtime.KeepAlive(text)
	runtime.KeepAlive(length)
	runtime.KeepAlive(analysis)
	runtime.KeepAlive(xPos)

	var _index_ int32   // out
	var _trailing int32 // out

	_index_ = int32(_arg5)
	_trailing = int32(_arg6)

	return _index_, _trailing
}

// GlyphVisAttr: PangoGlyphVisAttr structure communicates information between
// the shaping and rendering phases.
//
// Currently, it contains only cluster start information. yMore attributes may
// be added in the future.
//
// An instance of this type is always passed by reference.
type GlyphVisAttr struct {
	*glyphVisAttr
}

// glyphVisAttr is the struct that's finalized.
type glyphVisAttr struct {
	native *C.PangoGlyphVisAttr
}
