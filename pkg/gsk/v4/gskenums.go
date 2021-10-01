// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"fmt"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_blend_mode_get_type()), F: marshalBlendMode},
		{T: externglib.Type(C.gsk_corner_get_type()), F: marshalCorner},
		{T: externglib.Type(C.gsk_gl_uniform_type_get_type()), F: marshalGLUniformType},
		{T: externglib.Type(C.gsk_render_node_type_get_type()), F: marshalRenderNodeType},
		{T: externglib.Type(C.gsk_scaling_filter_get_type()), F: marshalScalingFilter},
		{T: externglib.Type(C.gsk_serialization_error_get_type()), F: marshalSerializationError},
		{T: externglib.Type(C.gsk_transform_category_get_type()), F: marshalTransformCategory},
	})
}

// BlendMode: blend modes available for render nodes.
//
// The implementation of each blend mode is deferred to the rendering pipeline.
//
// See https://www.w3.org/TR/compositing-1/#blending for more information on
// blending and blend modes.
type BlendMode int

const (
	// BlendModeDefault: default blend mode, which specifies no blending.
	BlendModeDefault BlendMode = iota
	// BlendModeMultiply: source color is multiplied by the destination and
	// replaces the destination.
	BlendModeMultiply
	// BlendModeScreen multiplies the complements of the destination and source
	// color values, then complements the result.
	BlendModeScreen
	// BlendModeOverlay multiplies or screens the colors, depending on the
	// destination color value. This is the inverse of hard-list.
	BlendModeOverlay
	// BlendModeDarken selects the darker of the destination and source colors.
	BlendModeDarken
	// BlendModeLighten selects the lighter of the destination and source
	// colors.
	BlendModeLighten
	// BlendModeColorDodge brightens the destination color to reflect the source
	// color.
	BlendModeColorDodge
	// BlendModeColorBurn darkens the destination color to reflect the source
	// color.
	BlendModeColorBurn
	// BlendModeHardLight multiplies or screens the colors, depending on the
	// source color value.
	BlendModeHardLight
	// BlendModeSoftLight darkens or lightens the colors, depending on the
	// source color value.
	BlendModeSoftLight
	// BlendModeDifference subtracts the darker of the two constituent colors
	// from the lighter color.
	BlendModeDifference
	// BlendModeExclusion produces an effect similar to that of the difference
	// mode but lower in contrast.
	BlendModeExclusion
	// BlendModeColor creates a color with the hue and saturation of the source
	// color and the luminosity of the destination color.
	BlendModeColor
	// BlendModeHue creates a color with the hue of the source color and the
	// saturation and luminosity of the destination color.
	BlendModeHue
	// BlendModeSaturation creates a color with the saturation of the source
	// color and the hue and luminosity of the destination color.
	BlendModeSaturation
	// BlendModeLuminosity creates a color with the luminosity of the source
	// color and the hue and saturation of the destination color.
	BlendModeLuminosity
)

func marshalBlendMode(p uintptr) (interface{}, error) {
	return BlendMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for BlendMode.
func (b BlendMode) String() string {
	switch b {
	case BlendModeDefault:
		return "Default"
	case BlendModeMultiply:
		return "Multiply"
	case BlendModeScreen:
		return "Screen"
	case BlendModeOverlay:
		return "Overlay"
	case BlendModeDarken:
		return "Darken"
	case BlendModeLighten:
		return "Lighten"
	case BlendModeColorDodge:
		return "ColorDodge"
	case BlendModeColorBurn:
		return "ColorBurn"
	case BlendModeHardLight:
		return "HardLight"
	case BlendModeSoftLight:
		return "SoftLight"
	case BlendModeDifference:
		return "Difference"
	case BlendModeExclusion:
		return "Exclusion"
	case BlendModeColor:
		return "Color"
	case BlendModeHue:
		return "Hue"
	case BlendModeSaturation:
		return "Saturation"
	case BlendModeLuminosity:
		return "Luminosity"
	default:
		return fmt.Sprintf("BlendMode(%d)", b)
	}
}

// Corner: corner indices used by RoundedRect.
type Corner int

const (
	// CornerTopLeft: top left corner.
	CornerTopLeft Corner = iota
	// CornerTopRight: top right corner.
	CornerTopRight
	// CornerBottomRight: bottom right corner.
	CornerBottomRight
	// CornerBottomLeft: bottom left corner.
	CornerBottomLeft
)

func marshalCorner(p uintptr) (interface{}, error) {
	return Corner(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for Corner.
func (c Corner) String() string {
	switch c {
	case CornerTopLeft:
		return "TopLeft"
	case CornerTopRight:
		return "TopRight"
	case CornerBottomRight:
		return "BottomRight"
	case CornerBottomLeft:
		return "BottomLeft"
	default:
		return fmt.Sprintf("Corner(%d)", c)
	}
}

// GLUniformType: this defines the types of the uniforms that GskGLShaders
// declare.
//
// It defines both what the type is called in the GLSL shader code, and what the
// corresponding C type is on the Gtk side.
type GLUniformType int

const (
	// GLUniformTypeNone: no type, used for uninitialized or unspecified values.
	GLUniformTypeNone GLUniformType = iota
	// GLUniformTypeFloat: float uniform.
	GLUniformTypeFloat
	// GLUniformTypeInt: GLSL int / gint32 uniform.
	GLUniformTypeInt
	// GLUniformTypeUint: GLSL uint / guint32 uniform.
	GLUniformTypeUint
	// GLUniformTypeBool: GLSL bool / gboolean uniform.
	GLUniformTypeBool
	// GLUniformTypeVec2: GLSL vec2 / graphene_vec2_t uniform.
	GLUniformTypeVec2
	// GLUniformTypeVec3: GLSL vec3 / graphene_vec3_t uniform.
	GLUniformTypeVec3
	// GLUniformTypeVec4: GLSL vec4 / graphene_vec4_t uniform.
	GLUniformTypeVec4
)

func marshalGLUniformType(p uintptr) (interface{}, error) {
	return GLUniformType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for GLUniformType.
func (g GLUniformType) String() string {
	switch g {
	case GLUniformTypeNone:
		return "None"
	case GLUniformTypeFloat:
		return "Float"
	case GLUniformTypeInt:
		return "Int"
	case GLUniformTypeUint:
		return "Uint"
	case GLUniformTypeBool:
		return "Bool"
	case GLUniformTypeVec2:
		return "Vec2"
	case GLUniformTypeVec3:
		return "Vec3"
	case GLUniformTypeVec4:
		return "Vec4"
	default:
		return fmt.Sprintf("GLUniformType(%d)", g)
	}
}

// RenderNodeType: type of a node determines what the node is rendering.
type RenderNodeType int

const (
	// NotARenderNodeType: error type. No node will ever have this type.
	NotARenderNodeType RenderNodeType = iota
	// ContainerNodeType: node containing a stack of children.
	ContainerNodeType
	// CairoNodeType: node drawing a #cairo_surface_t.
	CairoNodeType
	// ColorNodeType: node drawing a single color rectangle.
	ColorNodeType
	// LinearGradientNodeType: node drawing a linear gradient.
	LinearGradientNodeType
	// RepeatingLinearGradientNodeType: node drawing a repeating linear
	// gradient.
	RepeatingLinearGradientNodeType
	// RadialGradientNodeType: node drawing a radial gradient.
	RadialGradientNodeType
	// RepeatingRadialGradientNodeType: node drawing a repeating radial
	// gradient.
	RepeatingRadialGradientNodeType
	// ConicGradientNodeType: node drawing a conic gradient.
	ConicGradientNodeType
	// BorderNodeType: node stroking a border around an area.
	BorderNodeType
	// TextureNodeType: node drawing a Texture.
	TextureNodeType
	// InsetShadowNodeType: node drawing an inset shadow.
	InsetShadowNodeType
	// OutsetShadowNodeType: node drawing an outset shadow.
	OutsetShadowNodeType
	// TransformNodeType: node that renders its child after applying a matrix
	// transform.
	TransformNodeType
	// OpacityNodeType: node that changes the opacity of its child.
	OpacityNodeType
	// ColorMatrixNodeType: node that applies a color matrix to every pixel.
	ColorMatrixNodeType
	// RepeatNodeType: node that repeats the child's contents.
	RepeatNodeType
	// ClipNodeType: node that clips its child to a rectangular area.
	ClipNodeType
	// RoundedClipNodeType: node that clips its child to a rounded rectangle.
	RoundedClipNodeType
	// ShadowNodeType: node that draws a shadow below its child.
	ShadowNodeType
	// BlendNodeType: node that blends two children together.
	BlendNodeType
	// CrossFadeNodeType: node that cross-fades between two children.
	CrossFadeNodeType
	// TextNodeType: node containing a glyph string.
	TextNodeType
	// BlurNodeType: node that applies a blur.
	BlurNodeType
	// DebugNodeType: debug information that does not affect the rendering.
	DebugNodeType
	// GLShaderNodeType: node that uses OpenGL fragment shaders to render.
	GLShaderNodeType
)

func marshalRenderNodeType(p uintptr) (interface{}, error) {
	return RenderNodeType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for RenderNodeType.
func (r RenderNodeType) String() string {
	switch r {
	case NotARenderNodeType:
		return "NotARenderNode"
	case ContainerNodeType:
		return "ContainerNode"
	case CairoNodeType:
		return "CairoNode"
	case ColorNodeType:
		return "ColorNode"
	case LinearGradientNodeType:
		return "LinearGradientNode"
	case RepeatingLinearGradientNodeType:
		return "RepeatingLinearGradientNode"
	case RadialGradientNodeType:
		return "RadialGradientNode"
	case RepeatingRadialGradientNodeType:
		return "RepeatingRadialGradientNode"
	case ConicGradientNodeType:
		return "ConicGradientNode"
	case BorderNodeType:
		return "BorderNode"
	case TextureNodeType:
		return "TextureNode"
	case InsetShadowNodeType:
		return "InsetShadowNode"
	case OutsetShadowNodeType:
		return "OutsetShadowNode"
	case TransformNodeType:
		return "TransformNode"
	case OpacityNodeType:
		return "OpacityNode"
	case ColorMatrixNodeType:
		return "ColorMatrixNode"
	case RepeatNodeType:
		return "RepeatNode"
	case ClipNodeType:
		return "ClipNode"
	case RoundedClipNodeType:
		return "RoundedClipNode"
	case ShadowNodeType:
		return "ShadowNode"
	case BlendNodeType:
		return "BlendNode"
	case CrossFadeNodeType:
		return "CrossFadeNode"
	case TextNodeType:
		return "TextNode"
	case BlurNodeType:
		return "BlurNode"
	case DebugNodeType:
		return "DebugNode"
	case GLShaderNodeType:
		return "GLShaderNode"
	default:
		return fmt.Sprintf("RenderNodeType(%d)", r)
	}
}

// ScalingFilter filters used when scaling texture data.
//
// The actual implementation of each filter is deferred to the rendering
// pipeline.
type ScalingFilter int

const (
	// ScalingFilterLinear: linear interpolation filter.
	ScalingFilterLinear ScalingFilter = iota
	// ScalingFilterNearest: nearest neighbor interpolation filter.
	ScalingFilterNearest
	// ScalingFilterTrilinear: linear interpolation along each axis, plus mipmap
	// generation, with linear interpolation along the mipmap levels.
	ScalingFilterTrilinear
)

func marshalScalingFilter(p uintptr) (interface{}, error) {
	return ScalingFilter(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for ScalingFilter.
func (s ScalingFilter) String() string {
	switch s {
	case ScalingFilterLinear:
		return "Linear"
	case ScalingFilterNearest:
		return "Nearest"
	case ScalingFilterTrilinear:
		return "Trilinear"
	default:
		return fmt.Sprintf("ScalingFilter(%d)", s)
	}
}

// SerializationError errors that can happen during (de)serialization.
type SerializationError int

const (
	// SerializationUnsupportedFormat: format can not be identified.
	SerializationUnsupportedFormat SerializationError = iota
	// SerializationUnsupportedVersion: version of the data is not understood.
	SerializationUnsupportedVersion
	// SerializationInvalidData: given data may not exist in a proper
	// serialization.
	SerializationInvalidData
)

func marshalSerializationError(p uintptr) (interface{}, error) {
	return SerializationError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for SerializationError.
func (s SerializationError) String() string {
	switch s {
	case SerializationUnsupportedFormat:
		return "UnsupportedFormat"
	case SerializationUnsupportedVersion:
		return "UnsupportedVersion"
	case SerializationInvalidData:
		return "InvalidData"
	default:
		return fmt.Sprintf("SerializationError(%d)", s)
	}
}

// TransformCategory categories of matrices relevant for GSK and GTK.
//
// Note that any category includes matrices of all later categories. So if you
// want to for example check if a matrix is a 2D matrix, category >=
// GSK_TRANSFORM_CATEGORY_2D is the way to do this.
//
// Also keep in mind that rounding errors may cause matrices to not conform to
// their categories. Otherwise, matrix operations done via multiplication will
// not worsen categories. So for the matrix multiplication C = A * B,
// category(C) = MIN (category(A), category(B)).
type TransformCategory int

const (
	// TransformCategoryUnknown: category of the matrix has not been determined.
	TransformCategoryUnknown TransformCategory = iota
	// TransformCategoryAny: analyzing the matrix concluded that it does not fit
	// in any other category.
	TransformCategoryAny
	// TransformCategory3D: matrix is a 3D matrix. This means that the w column
	// (the last column) has the values (0, 0, 0, 1).
	TransformCategory3D
	// TransformCategory2D: matrix is a 2D matrix. This is equivalent to
	// graphene_matrix_is_2d() returning TRUE. In particular, this means that
	// Cairo can deal with the matrix.
	TransformCategory2D
	// TransformCategory2DAffine: matrix is a combination of 2D scale and 2D
	// translation operations. In particular, this means that any rectangle can
	// be transformed exactly using this matrix.
	TransformCategory2DAffine
	// TransformCategory2DTranslate: matrix is a 2D translation.
	TransformCategory2DTranslate
	// TransformCategoryIdentity: matrix is the identity matrix.
	TransformCategoryIdentity
)

func marshalTransformCategory(p uintptr) (interface{}, error) {
	return TransformCategory(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for TransformCategory.
func (t TransformCategory) String() string {
	switch t {
	case TransformCategoryUnknown:
		return "Unknown"
	case TransformCategoryAny:
		return "Any"
	case TransformCategory3D:
		return "3D"
	case TransformCategory2D:
		return "2D"
	case TransformCategory2DAffine:
		return "2DAffine"
	case TransformCategory2DTranslate:
		return "2DTranslate"
	case TransformCategoryIdentity:
		return "Identity"
	default:
		return fmt.Sprintf("TransformCategory(%d)", t)
	}
}
