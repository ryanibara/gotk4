// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk"
	"github.com/diamondburned/gotk4/pkg/glib"
	"github.com/diamondburned/gotk4/pkg/graphene"
	"github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gsk/gsk.h>
import "C"

func init() {
	glib.RegisterGValueMarshalers([]glib.TypeMarshaler{
		// Enums
		{T: glib.Type(C.gsk_blend_mode_get_type()), F: marshalBlendMode},
		{T: glib.Type(C.gsk_corner_get_type()), F: marshalCorner},
		{T: glib.Type(C.gsk_gl_uniform_type_get_type()), F: marshalGLUniformType},
		{T: glib.Type(C.gsk_render_node_type_get_type()), F: marshalRenderNodeType},
		{T: glib.Type(C.gsk_scaling_filter_get_type()), F: marshalScalingFilter},
		{T: glib.Type(C.gsk_serialization_error_get_type()), F: marshalSerializationError},
		{T: glib.Type(C.gsk_transform_category_get_type()), F: marshalTransformCategory},

		// Objects/Classes
	})
}

// BlendMode: the blend modes available for render nodes.
//
// The implementation of each blend mode is deferred to the rendering pipeline.
//
// See https://www.w3.org/TR/compositing-1/#blending for more information on
// blending and blend modes.
type BlendMode int

const (
	// BlendModeDefault: the default blend mode, which specifies no blending
	BlendModeDefault BlendMode = 0
	// BlendModeMultiply: the source color is multiplied by the destination and
	// replaces the destination
	BlendModeMultiply BlendMode = 1
	// BlendModeScreen: multiplies the complements of the destination and source
	// color values, then complements the result.
	BlendModeScreen BlendMode = 2
	// BlendModeOverlay: multiplies or screens the colors, depending on the
	// destination color value. This is the inverse of hard-list
	BlendModeOverlay BlendMode = 3
	// BlendModeDarken: selects the darker of the destination and source colors
	BlendModeDarken BlendMode = 4
	// BlendModeLighten: selects the lighter of the destination and source
	// colors
	BlendModeLighten BlendMode = 5
	// BlendModeColorDodge: brightens the destination color to reflect the
	// source color
	BlendModeColorDodge BlendMode = 6
	// BlendModeColorBurn: darkens the destination color to reflect the source
	// color
	BlendModeColorBurn BlendMode = 7
	// BlendModeHardLight: multiplies or screens the colors, depending on the
	// source color value
	BlendModeHardLight BlendMode = 8
	// BlendModeSoftLight: darkens or lightens the colors, depending on the
	// source color value
	BlendModeSoftLight BlendMode = 9
	// BlendModeDifference: subtracts the darker of the two constituent colors
	// from the lighter color
	BlendModeDifference BlendMode = 10
	// BlendModeExclusion: produces an effect similar to that of the difference
	// mode but lower in contrast
	BlendModeExclusion BlendMode = 11
	// BlendModeColor: creates a color with the hue and saturation of the source
	// color and the luminosity of the destination color
	BlendModeColor BlendMode = 12
	// BlendModeHue: creates a color with the hue of the source color and the
	// saturation and luminosity of the destination color
	BlendModeHue BlendMode = 13
	// BlendModeSaturation: creates a color with the saturation of the source
	// color and the hue and luminosity of the destination color
	BlendModeSaturation BlendMode = 14
	// BlendModeLuminosity: creates a color with the luminosity of the source
	// color and the hue and saturation of the destination color
	BlendModeLuminosity BlendMode = 15
)

func marshalBlendMode(p uintptr) (interface{}, error) {
	return BlendMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Corner: the corner indices used by RoundedRect.
type Corner int

const (
	// CornerTopLeft: the top left corner
	CornerTopLeft Corner = 0
	// CornerTopRight: the top right corner
	CornerTopRight Corner = 1
	// CornerBottomRight: the bottom right corner
	CornerBottomRight Corner = 2
	// CornerBottomLeft: the bottom left corner
	CornerBottomLeft Corner = 3
)

func marshalCorner(p uintptr) (interface{}, error) {
	return Corner(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// GLUniformType: this defines the types of the uniforms that GLShaders declare.
// It defines both what the type is called in the GLSL shader code, and what the
// corresponding C type is on the Gtk side.
type GLUniformType int

const (
	// GLUniformTypeNone: no type, used for uninitialized or unspecified values.
	GLUniformTypeNone GLUniformType = 0
	// GLUniformTypeFloat: a float uniform
	GLUniformTypeFloat GLUniformType = 1
	// GLUniformTypeInt: a GLSL int / gint32 uniform
	GLUniformTypeInt GLUniformType = 2
	// GLUniformTypeUint: a GLSL uint / guint32 uniform
	GLUniformTypeUint GLUniformType = 3
	// GLUniformTypeBool: a GLSL bool / gboolean uniform
	GLUniformTypeBool GLUniformType = 4
	// GLUniformTypeVec2: a GLSL vec2 / graphene_vec2_t uniform
	GLUniformTypeVec2 GLUniformType = 5
	// GLUniformTypeVec3: a GLSL vec3 / graphene_vec3_t uniform
	GLUniformTypeVec3 GLUniformType = 6
	// GLUniformTypeVec4: a GLSL vec4 / graphene_vec4_t uniform
	GLUniformTypeVec4 GLUniformType = 7
)

func marshalGLUniformType(p uintptr) (interface{}, error) {
	return GLUniformType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// RenderNodeType: the type of a node determines what the node is rendering.
type RenderNodeType int

const (
	// RenderNodeTypeNotARenderNode: error type. No node will ever have this
	// type.
	RenderNodeTypeNotARenderNode RenderNodeType = 0
	// RenderNodeTypeContainerNode: a node containing a stack of children
	RenderNodeTypeContainerNode RenderNodeType = 1
	// RenderNodeTypeCairoNode: a node drawing a #cairo_surface_t
	RenderNodeTypeCairoNode RenderNodeType = 2
	// RenderNodeTypeColorNode: a node drawing a single color rectangle
	RenderNodeTypeColorNode RenderNodeType = 3
	// RenderNodeTypeLinearGradientNode: a node drawing a linear gradient
	RenderNodeTypeLinearGradientNode RenderNodeType = 4
	// RenderNodeTypeRepeatingLinearGradientNode: a node drawing a repeating
	// linear gradient
	RenderNodeTypeRepeatingLinearGradientNode RenderNodeType = 5
	// RenderNodeTypeRadialGradientNode: a node drawing a radial gradient
	RenderNodeTypeRadialGradientNode RenderNodeType = 6
	// RenderNodeTypeRepeatingRadialGradientNode: a node drawing a repeating
	// radial gradient
	RenderNodeTypeRepeatingRadialGradientNode RenderNodeType = 7
	// RenderNodeTypeConicGradientNode: a node drawing a conic gradient
	RenderNodeTypeConicGradientNode RenderNodeType = 8
	// RenderNodeTypeBorderNode: a node stroking a border around an area
	RenderNodeTypeBorderNode RenderNodeType = 9
	// RenderNodeTypeTextureNode: a node drawing a Texture
	RenderNodeTypeTextureNode RenderNodeType = 10
	// RenderNodeTypeInsetShadowNode: a node drawing an inset shadow
	RenderNodeTypeInsetShadowNode RenderNodeType = 11
	// RenderNodeTypeOutsetShadowNode: a node drawing an outset shadow
	RenderNodeTypeOutsetShadowNode RenderNodeType = 12
	// RenderNodeTypeTransformNode: a node that renders its child after applying
	// a matrix transform
	RenderNodeTypeTransformNode RenderNodeType = 13
	// RenderNodeTypeOpacityNode: a node that changes the opacity of its child
	RenderNodeTypeOpacityNode RenderNodeType = 14
	// RenderNodeTypeColorMatrixNode: a node that applies a color matrix to
	// every pixel
	RenderNodeTypeColorMatrixNode RenderNodeType = 15
	// RenderNodeTypeRepeatNode: a node that repeats the child's contents
	RenderNodeTypeRepeatNode RenderNodeType = 16
	// RenderNodeTypeClipNode: a node that clips its child to a rectangular area
	RenderNodeTypeClipNode RenderNodeType = 17
	// RenderNodeTypeRoundedClipNode: a node that clips its child to a rounded
	// rectangle
	RenderNodeTypeRoundedClipNode RenderNodeType = 18
	// RenderNodeTypeShadowNode: a node that draws a shadow below its child
	RenderNodeTypeShadowNode RenderNodeType = 19
	// RenderNodeTypeBlendNode: a node that blends two children together
	RenderNodeTypeBlendNode RenderNodeType = 20
	// RenderNodeTypeCrossFadeNode: a node that cross-fades between two children
	RenderNodeTypeCrossFadeNode RenderNodeType = 21
	// RenderNodeTypeTextNode: a node containing a glyph string
	RenderNodeTypeTextNode RenderNodeType = 22
	// RenderNodeTypeBlurNode: a node that applies a blur
	RenderNodeTypeBlurNode RenderNodeType = 23
	// RenderNodeTypeDebugNode: debug information that does not affect the
	// rendering
	RenderNodeTypeDebugNode RenderNodeType = 24
	// RenderNodeTypeGlShaderNode: a node that uses OpenGL fragment shaders to
	// render
	RenderNodeTypeGlShaderNode RenderNodeType = 25
)

func marshalRenderNodeType(p uintptr) (interface{}, error) {
	return RenderNodeType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ScalingFilter: the filters used when scaling texture data.
//
// The actual implementation of each filter is deferred to the rendering
// pipeline.
type ScalingFilter int

const (
	// ScalingFilterLinear: linear interpolation filter
	ScalingFilterLinear ScalingFilter = 0
	// ScalingFilterNearest: nearest neighbor interpolation filter
	ScalingFilterNearest ScalingFilter = 1
	// ScalingFilterTrilinear: linear interpolation along each axis, plus mipmap
	// generation, with linear interpolation along the mipmap levels
	ScalingFilterTrilinear ScalingFilter = 2
)

func marshalScalingFilter(p uintptr) (interface{}, error) {
	return ScalingFilter(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// SerializationError: errors that can happen during (de)serialization.
type SerializationError int

const (
	// SerializationErrorUnsupportedFormat: the format can not be identified
	SerializationErrorUnsupportedFormat SerializationError = 0
	// SerializationErrorUnsupportedVersion: the version of the data is not
	// understood
	SerializationErrorUnsupportedVersion SerializationError = 1
	// SerializationErrorInvalidData: the given data may not exist in a proper
	// serialization
	SerializationErrorInvalidData SerializationError = 2
)

func marshalSerializationError(p uintptr) (interface{}, error) {
	return SerializationError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// TransformCategory: the categories of matrices relevant for GSK and GTK. Note
// that any category includes matrices of all later categories. So if you want
// to for example check if a matrix is a 2D matrix, `category >=
// GSK_TRANSFORM_CATEGORY_2D` is the way to do this.
//
// Also keep in mind that rounding errors may cause matrices to not conform to
// their categories. Otherwise, matrix operations done via mutliplication will
// not worsen categories. So for the matrix multiplication `C = A * B`,
// `category(C) = MIN (category(A), category(B))`.
type TransformCategory int

const (
	// TransformCategoryUnknown: the category of the matrix has not been
	// determined.
	TransformCategoryUnknown TransformCategory = 0
	// TransformCategoryAny: analyzing the matrix concluded that it does not fit
	// in any other category.
	TransformCategoryAny TransformCategory = 1
	// TransformCategory3D: the matrix is a 3D matrix. This means that the w
	// column (the last column) has the values (0, 0, 0, 1).
	TransformCategory3D TransformCategory = 2
	// TransformCategory2D: the matrix is a 2D matrix. This is equivalent to
	// graphene_matrix_is_2d() returning true. In particular, this means that
	// Cairo can deal with the matrix.
	TransformCategory2D TransformCategory = 3
	// TransformCategory2DAffine: the matrix is a combination of 2D scale and 2D
	// translation operations. In particular, this means that any rectangle can
	// be transformed exactly using this matrix.
	TransformCategory2DAffine TransformCategory = 4
	// TransformCategory2DTranslate: the matrix is a 2D translation.
	TransformCategory2DTranslate TransformCategory = 5
	// TransformCategoryIdentity: the matrix is the identity matrix.
	TransformCategoryIdentity TransformCategory = 6
)

func marshalTransformCategory(p uintptr) (interface{}, error) {
	return TransformCategory(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

func SerializationErrorQuark() glib.Quark

// TransformParse: parses the given @string into a transform and puts it in
// @out_transform. Strings printed via gsk_transform_to_string() can be read in
// again successfully using this function.
//
// If @string does not describe a valid transform, false is returned and nil is
// put in @out_transform.
func TransformParse(string string) (*Transform, bool)

// ColorStop: a color stop in a gradient node.
type ColorStop struct {
	// Offset: the offset of the color stop
	Offset float32
	// Color: the color at the given offset
	Color gdk.RGBA

	native *C.GskColorStop
}

func wrapColorStop(p *C.GskColorStop) *ColorStop {
	var v ColorStop

	v.Offset = float32(p.offset)
	v.Color = wrapRGBA(p.color)

	return &v
}

func marshalColorStop(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.GskColorStop)(unsafe.Pointer(b))

	return wrapColorStop(c)
}

// Native returns the pointer to *C.GskColorStop. The caller is expected to
// cast.
func (c *ColorStop) Native() unsafe.Pointer {
	return unsafe.Pointer(c.native)
}

// ParseLocation: a location in a parse buffer.
type ParseLocation struct {
	// Bytes: the offset of the location in the parse buffer, as bytes
	Bytes uint
	// Chars: the offset of the location in the parse buffer, as characters
	Chars uint
	// Lines: the line of the location in the parse buffer
	Lines uint
	// LineBytes: the position in the line, as bytes
	LineBytes uint
	// LineChars: the position in the line, as characters
	LineChars uint

	native *C.GskParseLocation
}

func wrapParseLocation(p *C.GskParseLocation) *ParseLocation {
	var v ParseLocation

	v.Bytes = uint(p.bytes)
	v.Chars = uint(p.chars)
	v.Lines = uint(p.lines)
	v.LineBytes = uint(p.line_bytes)
	v.LineChars = uint(p.line_chars)

	return &v
}

func marshalParseLocation(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.GskParseLocation)(unsafe.Pointer(b))

	return wrapParseLocation(c)
}

// Native returns the pointer to *C.GskParseLocation. The caller is expected to
// cast.
func (p *ParseLocation) Native() unsafe.Pointer {
	return unsafe.Pointer(p.native)
}

// RoundedRect: a rectangular region with rounded corners.
//
// Application code should normalize rectangles using
// gsk_rounded_rect_normalize(); this function will ensure that the bounds of
// the rectangle are normalized and ensure that the corner values are positive
// and the corners do not overlap. All functions taking a RoundedRect as an
// argument will internally operate on a normalized copy; all functions
// returning a RoundedRect will always return a normalized one.
type RoundedRect struct {
	// Bounds: the bounds of the rectangle
	Bounds graphene.Rect
	// Corner: the size of the 4 rounded corners
	Corner [4]graphene.Size

	native *C.GskRoundedRect
}

func wrapRoundedRect(p *C.GskRoundedRect) *RoundedRect {
	var v RoundedRect

	v.Bounds = wrapRect(p.bounds)

	return &v
}

func marshalRoundedRect(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.GskRoundedRect)(unsafe.Pointer(b))

	return wrapRoundedRect(c)
}

// Native returns the pointer to *C.GskRoundedRect. The caller is expected to
// cast.
func (r *RoundedRect) Native() unsafe.Pointer {
	return unsafe.Pointer(r.native)
}

// ShaderArgsBuilder: an object to build the uniforms data for a GLShader.
type ShaderArgsBuilder struct {
	native *C.GskShaderArgsBuilder
}

func wrapShaderArgsBuilder(p *C.GskShaderArgsBuilder) *ShaderArgsBuilder {
	v := ShaderArgsBuilder{native: p}

	runtime.SetFinalizer(v, nil)
	runtime.SetFinalizer(v, (*ShaderArgsBuilder).free)

	return &v
}

func marshalShaderArgsBuilder(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.GskShaderArgsBuilder)(unsafe.Pointer(b))

	return wrapShaderArgsBuilder(c)
}

func (s *ShaderArgsBuilder) free() {}

// Native returns the pointer to *C.GskShaderArgsBuilder. The caller is expected to
// cast.
func (s *ShaderArgsBuilder) Native() unsafe.Pointer {
	return unsafe.Pointer(s.native)
}

func NewShaderArgsBuilder(shader *GLShader, initialValues *glib.Bytes) *ShaderArgsBuilder

// Shadow: the shadow parameters in a shadow node.
type Shadow struct {
	// Color: the color of the shadow
	Color gdk.RGBA
	// Dx: the horizontal offset of the shadow
	Dx float32
	// Dy: the vertical offset of the shadow
	Dy float32
	// Radius: the radius of the shadow
	Radius float32

	native *C.GskShadow
}

func wrapShadow(p *C.GskShadow) *Shadow {
	var v Shadow

	v.Color = wrapRGBA(p.color)
	v.Dx = float32(p.dx)
	v.Dy = float32(p.dy)
	v.Radius = float32(p.radius)

	return &v
}

func marshalShadow(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.GskShadow)(unsafe.Pointer(b))

	return wrapShadow(c)
}

// Native returns the pointer to *C.GskShadow. The caller is expected to
// cast.
func (s *Shadow) Native() unsafe.Pointer {
	return unsafe.Pointer(s.native)
}

// Transform: the `GskTransform` structure contains only private data.
type Transform struct {
	native *C.GskTransform
}

func wrapTransform(p *C.GskTransform) *Transform {
	v := Transform{native: p}

	runtime.SetFinalizer(v, nil)
	runtime.SetFinalizer(v, (*Transform).free)

	return &v
}

func marshalTransform(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.GskTransform)(unsafe.Pointer(b))

	return wrapTransform(c)
}

func (t *Transform) free() {}

// Native returns the pointer to *C.GskTransform. The caller is expected to
// cast.
func (t *Transform) Native() unsafe.Pointer {
	return unsafe.Pointer(t.native)
}

func NewTransform() *Transform

// BlendNode: a render node applying a blending function between its two child
// nodes.
type BlendNode struct {
	RenderNode
}

func wrapBlendNode(obj *glib.Object) *BlendNode {
	return &BlendNode{RenderNode{obj}}
}

func marshalBlendNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// BlurNode: a render node applying a blur effect to its single child.
type BlurNode struct {
	RenderNode
}

func wrapBlurNode(obj *glib.Object) *BlurNode {
	return &BlurNode{RenderNode{obj}}
}

func marshalBlurNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// BorderNode: a render node for a border.
type BorderNode struct {
	RenderNode
}

func wrapBorderNode(obj *glib.Object) *BorderNode {
	return &BorderNode{RenderNode{obj}}
}

func marshalBorderNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// CairoNode: a render node for a Cairo surface.
type CairoNode struct {
	RenderNode
}

func wrapCairoNode(obj *glib.Object) *CairoNode {
	return &CairoNode{RenderNode{obj}}
}

func marshalCairoNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

type CairoRenderer struct {
	Renderer
}

func wrapCairoRenderer(obj *glib.Object) *CairoRenderer {
	return &CairoRenderer{Renderer{*glib.Object{obj}}}
}

func marshalCairoRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// ClipNode: a render node applying a rectangular clip to its single child node.
type ClipNode struct {
	RenderNode
}

func wrapClipNode(obj *glib.Object) *ClipNode {
	return &ClipNode{RenderNode{obj}}
}

func marshalClipNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// ColorMatrixNode: a render node controlling the color matrix of its single
// child node.
type ColorMatrixNode struct {
	RenderNode
}

func wrapColorMatrixNode(obj *glib.Object) *ColorMatrixNode {
	return &ColorMatrixNode{RenderNode{obj}}
}

func marshalColorMatrixNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// ColorNode: a render node for a solid color.
type ColorNode struct {
	RenderNode
}

func wrapColorNode(obj *glib.Object) *ColorNode {
	return &ColorNode{RenderNode{obj}}
}

func marshalColorNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// ConicGradientNode: a render node for a conic gradient.
type ConicGradientNode struct {
	RenderNode
}

func wrapConicGradientNode(obj *glib.Object) *ConicGradientNode {
	return &ConicGradientNode{RenderNode{obj}}
}

func marshalConicGradientNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// ContainerNode: a render node that can contain other render nodes.
type ContainerNode struct {
	RenderNode
}

func wrapContainerNode(obj *glib.Object) *ContainerNode {
	return &ContainerNode{RenderNode{obj}}
}

func marshalContainerNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// CrossFadeNode: a render node cross fading between two child nodes.
type CrossFadeNode struct {
	RenderNode
}

func wrapCrossFadeNode(obj *glib.Object) *CrossFadeNode {
	return &CrossFadeNode{RenderNode{obj}}
}

func marshalCrossFadeNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// DebugNode: a render node that emits a debugging message when drawing its
// child node.
type DebugNode struct {
	RenderNode
}

func wrapDebugNode(obj *glib.Object) *DebugNode {
	return &DebugNode{RenderNode{obj}}
}

func marshalDebugNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

type GLRenderer struct {
	Renderer
}

func wrapGLRenderer(obj *glib.Object) *GLRenderer {
	return &GLRenderer{Renderer{*glib.Object{obj}}}
}

func marshalGLRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// GLShader: an object representing a GL shader program.
type GLShader struct {
	*glib.Object
}

func wrapGLShader(obj *glib.Object) *GLShader {
	return &GLShader{*glib.Object{obj}}
}

func marshalGLShader(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// GLShaderNode: a render node using a GL shader when drawing its children
// nodes.
type GLShaderNode struct {
	RenderNode
}

func wrapGLShaderNode(obj *glib.Object) *GLShaderNode {
	return &GLShaderNode{RenderNode{obj}}
}

func marshalGLShaderNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// InsetShadowNode: a render node for an inset shadow.
type InsetShadowNode struct {
	RenderNode
}

func wrapInsetShadowNode(obj *glib.Object) *InsetShadowNode {
	return &InsetShadowNode{RenderNode{obj}}
}

func marshalInsetShadowNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// LinearGradientNode: a render node for a linear gradient.
type LinearGradientNode struct {
	RenderNode
}

func wrapLinearGradientNode(obj *glib.Object) *LinearGradientNode {
	return &LinearGradientNode{RenderNode{obj}}
}

func marshalLinearGradientNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// OpacityNode: a render node controlling the opacity of its single child node.
type OpacityNode struct {
	RenderNode
}

func wrapOpacityNode(obj *glib.Object) *OpacityNode {
	return &OpacityNode{RenderNode{obj}}
}

func marshalOpacityNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// OutsetShadowNode: a render node for an outset shadow.
type OutsetShadowNode struct {
	RenderNode
}

func wrapOutsetShadowNode(obj *glib.Object) *OutsetShadowNode {
	return &OutsetShadowNode{RenderNode{obj}}
}

func marshalOutsetShadowNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// RadialGradientNode: a render node for a radial gradient.
type RadialGradientNode struct {
	RenderNode
}

func wrapRadialGradientNode(obj *glib.Object) *RadialGradientNode {
	return &RadialGradientNode{RenderNode{obj}}
}

func marshalRadialGradientNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// Renderer: base type for the object managing the rendering pipeline for a
// Surface.
type Renderer struct {
	*glib.Object
}

func wrapRenderer(obj *glib.Object) *Renderer {
	return &Renderer{*glib.Object{obj}}
}

func marshalRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// RepeatNode: a render node repeating its single child node.
type RepeatNode struct {
	RenderNode
}

func wrapRepeatNode(obj *glib.Object) *RepeatNode {
	return &RepeatNode{RenderNode{obj}}
}

func marshalRepeatNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// RepeatingLinearGradientNode: a render node for a repeating linear gradient.
type RepeatingLinearGradientNode struct {
	RenderNode
}

func wrapRepeatingLinearGradientNode(obj *glib.Object) *RepeatingLinearGradientNode {
	return &RepeatingLinearGradientNode{RenderNode{obj}}
}

func marshalRepeatingLinearGradientNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// RepeatingRadialGradientNode: a render node for a repeating radial gradient.
type RepeatingRadialGradientNode struct {
	RenderNode
}

func wrapRepeatingRadialGradientNode(obj *glib.Object) *RepeatingRadialGradientNode {
	return &RepeatingRadialGradientNode{RenderNode{obj}}
}

func marshalRepeatingRadialGradientNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// RoundedClipNode: a render node applying a rounded rectangle clip to its
// single child.
type RoundedClipNode struct {
	RenderNode
}

func wrapRoundedClipNode(obj *glib.Object) *RoundedClipNode {
	return &RoundedClipNode{RenderNode{obj}}
}

func marshalRoundedClipNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// ShadowNode: a render node drawing one or more shadows behind its single child
// node.
type ShadowNode struct {
	RenderNode
}

func wrapShadowNode(obj *glib.Object) *ShadowNode {
	return &ShadowNode{RenderNode{obj}}
}

func marshalShadowNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// TextNode: a render node drawing a set of glyphs.
type TextNode struct {
	RenderNode
}

func wrapTextNode(obj *glib.Object) *TextNode {
	return &TextNode{RenderNode{obj}}
}

func marshalTextNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// TextureNode: a render node for a Texture.
type TextureNode struct {
	RenderNode
}

func wrapTextureNode(obj *glib.Object) *TextureNode {
	return &TextureNode{RenderNode{obj}}
}

func marshalTextureNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

// TransformNode: a render node applying a Transform to its single child node.
type TransformNode struct {
	RenderNode
}

func wrapTransformNode(obj *glib.Object) *TransformNode {
	return &TransformNode{RenderNode{obj}}
}

func marshalTransformNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

type VulkanRenderer struct {
	Renderer
}

func wrapVulkanRenderer(obj *glib.Object) *VulkanRenderer {
	return &VulkanRenderer{Renderer{*glib.Object{obj}}}
}

func marshalVulkanRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}
