// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/graphene"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeGLShader returns the GType for the type GLShader.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeGLShader() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gsk", "GLShader").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalGLShader)
	return gtype
}

// GTypeShaderArgsBuilder returns the GType for the type ShaderArgsBuilder.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeShaderArgsBuilder() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gsk", "ShaderArgsBuilder").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalShaderArgsBuilder)
	return gtype
}

// GLShaderOverrider contains methods that are overridable.
type GLShaderOverrider interface {
}

// GLShader: GskGLShader is a snippet of GLSL that is meant to run in the
// fragment shader of the rendering pipeline.
//
// A fragment shader gets the coordinates being rendered as input and produces
// the pixel values for that particular pixel. Additionally, the shader can
// declare a set of other input arguments, called uniforms (as they are uniform
// over all the calls to your shader in each instance of use). A shader can also
// receive up to 4 textures that it can use as input when producing the pixel
// data.
//
// GskGLShader is usually used with gtk_snapshot_push_gl_shader() to produce a
// gsk.GLShaderNode in the rendering hierarchy, and then its input textures are
// constructed by rendering the child nodes to textures before rendering the
// shader node itself. (You can pass texture nodes as children if you want to
// directly use a texture as input).
//
// The actual shader code is GLSL code that gets combined with some other code
// into the fragment shader. Since the exact capabilities of the GPU driver
// differs between different OpenGL drivers and hardware, GTK adds some defines
// that you can use to ensure your GLSL code runs on as many drivers as it can.
//
// If the OpenGL driver is GLES, then the shader language version is set to 100,
// and GSK_GLES will be defined in the shader.
//
// Otherwise, if the OpenGL driver does not support the 3.2 core profile, then
// the shader will run with language version 110 for GL2 and 130 for GL3, and
// GSK_LEGACY will be defined in the shader.
//
// If the OpenGL driver supports the 3.2 code profile, it will be used, the
// shader language version is set to 150, and GSK_GL3 will be defined in the
// shader.
//
// The main function the shader must implement is:
//
//     void mainImage(out vec4 fragColor,
//                    in vec2 fragCoord,
//                    in vec2 resolution,
//                    in vec2 uv)
//
//
// Where the input fragCoord is the coordinate of the pixel we're currently
// rendering, relative to the boundary rectangle that was specified in the
// GskGLShaderNode, and resolution is the width and height of that rectangle.
// This is in the typical GTK coordinate system with the origin in the top left.
// uv contains the u and v coordinates that can be used to index a texture at
// the corresponding point. These coordinates are in the [0..1]x[0..1] region,
// with 0, 0 being in the lower left corder (which is typical for OpenGL).
//
// The output fragColor should be a RGBA color (with premultiplied alpha) that
// will be used as the output for the specified pixel location. Note that this
// output will be automatically clipped to the clip region of the glshader node.
//
// In addition to the function arguments the shader can define up to 4 uniforms
// for textures which must be called u_textureN (i.e. u_texture1 to u_texture4)
// as well as any custom uniforms you want of types int, uint, bool, float,
// vec2, vec3 or vec4.
//
// All textures sources contain premultiplied alpha colors, but if some there
// are outer sources of colors there is a gsk_premultiply() helper to compute
// premultiplication when needed.
//
// Note that GTK parses the uniform declarations, so each uniform has to be on a
// line by itself with no other code, like so:
//
//    uniform float u_time;
//    uniform vec3 u_color;
//    uniform sampler2D u_texture1;
//    uniform sampler2D u_texture2;
//
//
// GTK uses the the "gsk" namespace in the symbols it uses in the shader, so
// your code should not use any symbols with the prefix gsk or GSK. There are
// some helper functions declared that you can use:
//
//    vec4 GskTexture(sampler2D sampler, vec2 texCoords);
//
//
// This samples a texture (e.g. u_texture1) at the specified coordinates, and
// containes some helper ifdefs to ensure that it works on all OpenGL versions.
//
// You can compile the shader yourself using gsk.GLShader.Compile(), otherwise
// the GSK renderer will do it when it handling the glshader node. If errors
// occurs, the returned error will include the glsl sources, so you can see what
// GSK was passing to the compiler. You can also set GSK_DEBUG=shaders in the
// environment to see the sources and other relevant information about all
// shaders that GSK is handling.
//
// An example shader
//
//    uniform float position;
//    uniform sampler2D u_texture1;
//    uniform sampler2D u_texture2;
//
//    void mainImage(out vec4 fragColor,
//                   in vec2 fragCoord,
//                   in vec2 resolution,
//                   in vec2 uv) {
//      vec4 source1 = GskTexture(u_texture1, uv);
//      vec4 source2 = GskTexture(u_texture2, uv);
//
//      fragColor = position * source1 + (1.0 - position) * source2;
//    }.
type GLShader struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*GLShader)(nil)
)

func classInitGLShaderer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGLShader(obj *coreglib.Object) *GLShader {
	return &GLShader{
		Object: obj,
	}
}

func marshalGLShader(p uintptr) (interface{}, error) {
	return wrapGLShader(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewGLShaderFromBytes creates a GskGLShader that will render pixels using the
// specified code.
//
// The function takes the following parameters:
//
//    - sourcecode: GLSL sourcecode for the shader, as a GBytes.
//
// The function returns the following values:
//
//    - glShader: new GskGLShader.
//
func NewGLShaderFromBytes(sourcecode *glib.Bytes) *GLShader {
	var _args [1]girepository.Argument

	*(**C.GBytes)(unsafe.Pointer(&_args[0])) = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(sourcecode)))

	_info := girepository.MustFind("Gsk", "GLShader")
	_gret := _info.InvokeClassMethod("new_GLShader_from_bytes", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(sourcecode)

	var _glShader *GLShader // out

	_glShader = wrapGLShader(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _glShader
}

// NewGLShaderFromResource creates a GskGLShader that will render pixels using
// the specified code.
//
// The function takes the following parameters:
//
//    - resourcePath: path to a resource that contains the GLSL sourcecode for
//      the shader.
//
// The function returns the following values:
//
//    - glShader: new GskGLShader.
//
func NewGLShaderFromResource(resourcePath string) *GLShader {
	var _args [1]girepository.Argument

	*(**C.char)(unsafe.Pointer(&_args[0])) = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gsk", "GLShader")
	_gret := _info.InvokeClassMethod("new_GLShader_from_resource", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(resourcePath)

	var _glShader *GLShader // out

	_glShader = wrapGLShader(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _glShader
}

// Compile tries to compile the shader for the given renderer.
//
// If there is a problem, this function returns FALSE and reports an error. You
// should use this function before relying on the shader for rendering and use a
// fallback with a simpler shader or without shaders if it fails.
//
// Note that this will modify the rendering state (for example change the
// current GL context) and requires the renderer to be set up. This means that
// the widget has to be realized. Commonly you want to call this from the
// realize signal of a widget, or during widget snapshot.
//
// The function takes the following parameters:
//
//    - renderer: GskRenderer.
//
func (shader *GLShader) Compile(renderer Rendererer) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))

	_info := girepository.MustFind("Gsk", "GLShader")
	_info.InvokeClassMethod("compile", _args[:], nil)

	runtime.KeepAlive(shader)
	runtime.KeepAlive(renderer)

	var _goerr error // out

	if *(**C.GError)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.GError)(unsafe.Pointer(&_cerr))))
	}

	return _goerr
}

// FindUniformByName looks for a uniform by the name name, and returns the index
// of the uniform, or -1 if it was not found.
//
// The function takes the following parameters:
//
//    - name: uniform name.
//
// The function returns the following values:
//
//    - gint: index of the uniform, or -1.
//
func (shader *GLShader) FindUniformByName(name string) int32 {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))
	*(**C.char)(unsafe.Pointer(&_args[1])) = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gsk", "GLShader")
	_gret := _info.InvokeClassMethod("find_uniform_by_name", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(shader)
	runtime.KeepAlive(name)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// ArgBool gets the value of the uniform idx in the args block.
//
// The uniform must be of bool type.
//
// The function takes the following parameters:
//
//    - args: uniform arguments.
//    - idx: index of the uniform.
//
// The function returns the following values:
//
//    - ok: value.
//
func (shader *GLShader) ArgBool(args *glib.Bytes, idx int32) bool {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))
	*(**C.GBytes)(unsafe.Pointer(&_args[1])) = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(args)))
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(idx)

	_info := girepository.MustFind("Gsk", "GLShader")
	_gret := _info.InvokeClassMethod("get_arg_bool", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(shader)
	runtime.KeepAlive(args)
	runtime.KeepAlive(idx)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ArgFloat gets the value of the uniform idx in the args block.
//
// The uniform must be of float type.
//
// The function takes the following parameters:
//
//    - args: uniform arguments.
//    - idx: index of the uniform.
//
// The function returns the following values:
//
//    - gfloat: value.
//
func (shader *GLShader) ArgFloat(args *glib.Bytes, idx int32) float32 {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))
	*(**C.GBytes)(unsafe.Pointer(&_args[1])) = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(args)))
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(idx)

	_info := girepository.MustFind("Gsk", "GLShader")
	_gret := _info.InvokeClassMethod("get_arg_float", _args[:], nil)
	_cret := *(*C.float)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(shader)
	runtime.KeepAlive(args)
	runtime.KeepAlive(idx)

	var _gfloat float32 // out

	_gfloat = float32(*(*C.float)(unsafe.Pointer(&_cret)))

	return _gfloat
}

// ArgInt gets the value of the uniform idx in the args block.
//
// The uniform must be of int type.
//
// The function takes the following parameters:
//
//    - args: uniform arguments.
//    - idx: index of the uniform.
//
// The function returns the following values:
//
//    - gint32: value.
//
func (shader *GLShader) ArgInt(args *glib.Bytes, idx int32) int32 {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))
	*(**C.GBytes)(unsafe.Pointer(&_args[1])) = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(args)))
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(idx)

	_info := girepository.MustFind("Gsk", "GLShader")
	_gret := _info.InvokeClassMethod("get_arg_int", _args[:], nil)
	_cret := *(*C.gint32)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(shader)
	runtime.KeepAlive(args)
	runtime.KeepAlive(idx)

	var _gint32 int32 // out

	_gint32 = int32(*(*C.gint32)(unsafe.Pointer(&_cret)))

	return _gint32
}

// ArgUint gets the value of the uniform idx in the args block.
//
// The uniform must be of uint type.
//
// The function takes the following parameters:
//
//    - args: uniform arguments.
//    - idx: index of the uniform.
//
// The function returns the following values:
//
//    - guint32: value.
//
func (shader *GLShader) ArgUint(args *glib.Bytes, idx int32) uint32 {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))
	*(**C.GBytes)(unsafe.Pointer(&_args[1])) = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(args)))
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(idx)

	_info := girepository.MustFind("Gsk", "GLShader")
	_gret := _info.InvokeClassMethod("get_arg_uint", _args[:], nil)
	_cret := *(*C.guint32)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(shader)
	runtime.KeepAlive(args)
	runtime.KeepAlive(idx)

	var _guint32 uint32 // out

	_guint32 = uint32(*(*C.guint32)(unsafe.Pointer(&_cret)))

	return _guint32
}

// ArgVec2 gets the value of the uniform idx in the args block.
//
// The uniform must be of vec2 type.
//
// The function takes the following parameters:
//
//    - args: uniform arguments.
//    - idx: index of the uniform.
//    - outValue: location to store the uniform value in.
//
func (shader *GLShader) ArgVec2(args *glib.Bytes, idx int32, outValue *graphene.Vec2) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))
	*(**C.GBytes)(unsafe.Pointer(&_args[1])) = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(args)))
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(idx)
	*(**C.graphene_vec2_t)(unsafe.Pointer(&_args[3])) = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(outValue)))

	_info := girepository.MustFind("Gsk", "GLShader")
	_info.InvokeClassMethod("get_arg_vec2", _args[:], nil)

	runtime.KeepAlive(shader)
	runtime.KeepAlive(args)
	runtime.KeepAlive(idx)
	runtime.KeepAlive(outValue)
}

// ArgVec3 gets the value of the uniform idx in the args block.
//
// The uniform must be of vec3 type.
//
// The function takes the following parameters:
//
//    - args: uniform arguments.
//    - idx: index of the uniform.
//    - outValue: location to store the uniform value in.
//
func (shader *GLShader) ArgVec3(args *glib.Bytes, idx int32, outValue *graphene.Vec3) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))
	*(**C.GBytes)(unsafe.Pointer(&_args[1])) = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(args)))
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(idx)
	*(**C.graphene_vec3_t)(unsafe.Pointer(&_args[3])) = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(outValue)))

	_info := girepository.MustFind("Gsk", "GLShader")
	_info.InvokeClassMethod("get_arg_vec3", _args[:], nil)

	runtime.KeepAlive(shader)
	runtime.KeepAlive(args)
	runtime.KeepAlive(idx)
	runtime.KeepAlive(outValue)
}

// ArgVec4 gets the value of the uniform idx in the args block.
//
// The uniform must be of vec4 type.
//
// The function takes the following parameters:
//
//    - args: uniform arguments.
//    - idx: index of the uniform.
//    - outValue: location to store set the uniform value in.
//
func (shader *GLShader) ArgVec4(args *glib.Bytes, idx int32, outValue *graphene.Vec4) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))
	*(**C.GBytes)(unsafe.Pointer(&_args[1])) = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(args)))
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(idx)
	*(**C.graphene_vec4_t)(unsafe.Pointer(&_args[3])) = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(outValue)))

	_info := girepository.MustFind("Gsk", "GLShader")
	_info.InvokeClassMethod("get_arg_vec4", _args[:], nil)

	runtime.KeepAlive(shader)
	runtime.KeepAlive(args)
	runtime.KeepAlive(idx)
	runtime.KeepAlive(outValue)
}

// ArgsSize: get the size of the data block used to specify arguments for this
// shader.
//
// The function returns the following values:
//
//    - gsize: size of the data block.
//
func (shader *GLShader) ArgsSize() uint {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))

	_info := girepository.MustFind("Gsk", "GLShader")
	_gret := _info.InvokeClassMethod("get_args_size", _args[:], nil)
	_cret := *(*C.gsize)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(shader)

	var _gsize uint // out

	_gsize = uint(*(*C.gsize)(unsafe.Pointer(&_cret)))

	return _gsize
}

// NTextures returns the number of textures that the shader requires.
//
// This can be used to check that the a passed shader works in your usecase. It
// is determined by looking at the highest u_textureN value that the shader
// defines.
//
// The function returns the following values:
//
//    - gint: number of texture inputs required by shader.
//
func (shader *GLShader) NTextures() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))

	_info := girepository.MustFind("Gsk", "GLShader")
	_gret := _info.InvokeClassMethod("get_n_textures", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(shader)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// NUniforms: get the number of declared uniforms for this shader.
//
// The function returns the following values:
//
//    - gint: number of declared uniforms.
//
func (shader *GLShader) NUniforms() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))

	_info := girepository.MustFind("Gsk", "GLShader")
	_gret := _info.InvokeClassMethod("get_n_uniforms", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(shader)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// Resource gets the resource path for the GLSL sourcecode being used to render
// this shader.
//
// The function returns the following values:
//
//    - utf8: resource path for the shader, or NULL if none.
//
func (shader *GLShader) Resource() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))

	_info := girepository.MustFind("Gsk", "GLShader")
	_gret := _info.InvokeClassMethod("get_resource", _args[:], nil)
	_cret := *(**C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(shader)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_cret)))))

	return _utf8
}

// Source gets the GLSL sourcecode being used to render this shader.
//
// The function returns the following values:
//
//    - bytes: source code for the shader.
//
func (shader *GLShader) Source() *glib.Bytes {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))

	_info := girepository.MustFind("Gsk", "GLShader")
	_gret := _info.InvokeClassMethod("get_source", _args[:], nil)
	_cret := *(**C.GBytes)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(shader)

	var _bytes *glib.Bytes // out

	_bytes = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(*(**C.GBytes)(unsafe.Pointer(&_cret)))))
	C.g_bytes_ref(*(**C.GBytes)(unsafe.Pointer(&_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _bytes
}

// UniformName: get the name of the declared uniform for this shader at index
// idx.
//
// The function takes the following parameters:
//
//    - idx: index of the uniform.
//
// The function returns the following values:
//
//    - utf8: name of the declared uniform.
//
func (shader *GLShader) UniformName(idx int32) string {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(idx)

	_info := girepository.MustFind("Gsk", "GLShader")
	_gret := _info.InvokeClassMethod("get_uniform_name", _args[:], nil)
	_cret := *(**C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(shader)
	runtime.KeepAlive(idx)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_cret)))))

	return _utf8
}

// UniformOffset: get the offset into the data block where data for this
// uniforms is stored.
//
// The function takes the following parameters:
//
//    - idx: index of the uniform.
//
// The function returns the following values:
//
//    - gint: data offset.
//
func (shader *GLShader) UniformOffset(idx int32) int32 {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(idx)

	_info := girepository.MustFind("Gsk", "GLShader")
	_gret := _info.InvokeClassMethod("get_uniform_offset", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(shader)
	runtime.KeepAlive(idx)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// ShaderArgsBuilder: object to build the uniforms data for a GLShader.
//
// An instance of this type is always passed by reference.
type ShaderArgsBuilder struct {
	*shaderArgsBuilder
}

// shaderArgsBuilder is the struct that's finalized.
type shaderArgsBuilder struct {
	native unsafe.Pointer
}

func marshalShaderArgsBuilder(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &ShaderArgsBuilder{&shaderArgsBuilder{(unsafe.Pointer)(b)}}, nil
}

// NewShaderArgsBuilder constructs a struct ShaderArgsBuilder.
func NewShaderArgsBuilder(shader *GLShader, initialValues *glib.Bytes) *ShaderArgsBuilder {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shader).Native()))
	if initialValues != nil {
		*(**C.GBytes)(unsafe.Pointer(&_args[1])) = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(initialValues)))
	}

	_info := girepository.MustFind("Gsk", "ShaderArgsBuilder")
	_gret := _info.InvokeRecordMethod("new", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(shader)
	runtime.KeepAlive(initialValues)

	var _shaderArgsBuilder *ShaderArgsBuilder // out

	_shaderArgsBuilder = (*ShaderArgsBuilder)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_shaderArgsBuilder)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _shaderArgsBuilder
}

// SetBool sets the value of the uniform idx.
//
// The uniform must be of bool type.
//
// The function takes the following parameters:
//
//    - idx: index of the uniform.
//    - value to set the uniform to.
//
func (builder *ShaderArgsBuilder) SetBool(idx int32, value bool) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(builder)))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(idx)
	if value {
		*(*C.gboolean)(unsafe.Pointer(&_args[2])) = C.TRUE
	}

	_info := girepository.MustFind("Gsk", "ShaderArgsBuilder")
	_info.InvokeRecordMethod("set_bool", _args[:], nil)

	runtime.KeepAlive(builder)
	runtime.KeepAlive(idx)
	runtime.KeepAlive(value)
}

// SetFloat sets the value of the uniform idx.
//
// The uniform must be of float type.
//
// The function takes the following parameters:
//
//    - idx: index of the uniform.
//    - value to set the uniform to.
//
func (builder *ShaderArgsBuilder) SetFloat(idx int32, value float32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(builder)))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(idx)
	*(*C.float)(unsafe.Pointer(&_args[2])) = C.float(value)

	_info := girepository.MustFind("Gsk", "ShaderArgsBuilder")
	_info.InvokeRecordMethod("set_float", _args[:], nil)

	runtime.KeepAlive(builder)
	runtime.KeepAlive(idx)
	runtime.KeepAlive(value)
}

// SetInt sets the value of the uniform idx.
//
// The uniform must be of int type.
//
// The function takes the following parameters:
//
//    - idx: index of the uniform.
//    - value to set the uniform to.
//
func (builder *ShaderArgsBuilder) SetInt(idx int32, value int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(builder)))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(idx)
	*(*C.gint32)(unsafe.Pointer(&_args[2])) = C.gint32(value)

	_info := girepository.MustFind("Gsk", "ShaderArgsBuilder")
	_info.InvokeRecordMethod("set_int", _args[:], nil)

	runtime.KeepAlive(builder)
	runtime.KeepAlive(idx)
	runtime.KeepAlive(value)
}

// SetUint sets the value of the uniform idx.
//
// The uniform must be of uint type.
//
// The function takes the following parameters:
//
//    - idx: index of the uniform.
//    - value to set the uniform to.
//
func (builder *ShaderArgsBuilder) SetUint(idx int32, value uint32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(builder)))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(idx)
	*(*C.guint32)(unsafe.Pointer(&_args[2])) = C.guint32(value)

	_info := girepository.MustFind("Gsk", "ShaderArgsBuilder")
	_info.InvokeRecordMethod("set_uint", _args[:], nil)

	runtime.KeepAlive(builder)
	runtime.KeepAlive(idx)
	runtime.KeepAlive(value)
}

// SetVec2 sets the value of the uniform idx.
//
// The uniform must be of vec2 type.
//
// The function takes the following parameters:
//
//    - idx: index of the uniform.
//    - value to set the uniform too.
//
func (builder *ShaderArgsBuilder) SetVec2(idx int32, value *graphene.Vec2) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(builder)))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(idx)
	*(**C.graphene_vec2_t)(unsafe.Pointer(&_args[2])) = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(value)))

	_info := girepository.MustFind("Gsk", "ShaderArgsBuilder")
	_info.InvokeRecordMethod("set_vec2", _args[:], nil)

	runtime.KeepAlive(builder)
	runtime.KeepAlive(idx)
	runtime.KeepAlive(value)
}

// SetVec3 sets the value of the uniform idx.
//
// The uniform must be of vec3 type.
//
// The function takes the following parameters:
//
//    - idx: index of the uniform.
//    - value to set the uniform too.
//
func (builder *ShaderArgsBuilder) SetVec3(idx int32, value *graphene.Vec3) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(builder)))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(idx)
	*(**C.graphene_vec3_t)(unsafe.Pointer(&_args[2])) = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(value)))

	_info := girepository.MustFind("Gsk", "ShaderArgsBuilder")
	_info.InvokeRecordMethod("set_vec3", _args[:], nil)

	runtime.KeepAlive(builder)
	runtime.KeepAlive(idx)
	runtime.KeepAlive(value)
}

// SetVec4 sets the value of the uniform idx.
//
// The uniform must be of vec4 type.
//
// The function takes the following parameters:
//
//    - idx: index of the uniform.
//    - value to set the uniform too.
//
func (builder *ShaderArgsBuilder) SetVec4(idx int32, value *graphene.Vec4) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(builder)))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(idx)
	*(**C.graphene_vec4_t)(unsafe.Pointer(&_args[2])) = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(value)))

	_info := girepository.MustFind("Gsk", "ShaderArgsBuilder")
	_info.InvokeRecordMethod("set_vec4", _args[:], nil)

	runtime.KeepAlive(builder)
	runtime.KeepAlive(idx)
	runtime.KeepAlive(value)
}

// ToArgs creates a new GBytes args from the current state of the given builder.
//
// Any uniforms of the shader that have not been explicitly set on the builder
// are zero-initialized.
//
// The given GskShaderArgsBuilder is reset once this function returns; you
// cannot call this function multiple times on the same builder instance.
//
// This function is intended primarily for bindings. C code should use
// gsk.ShaderArgsBuilder.FreeToArgs().
//
// The function returns the following values:
//
//    - bytes: newly allocated buffer with all the args added to builder.
//
func (builder *ShaderArgsBuilder) ToArgs() *glib.Bytes {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(builder)))

	_info := girepository.MustFind("Gsk", "ShaderArgsBuilder")
	_gret := _info.InvokeRecordMethod("to_args", _args[:], nil)
	_cret := *(**C.GBytes)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(builder)

	var _bytes *glib.Bytes // out

	_bytes = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(*(**C.GBytes)(unsafe.Pointer(&_cret)))))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _bytes
}
