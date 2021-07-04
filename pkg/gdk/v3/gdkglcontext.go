// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_gl_context_get_type()), F: marshalGLContext},
	})
}

// GLContext is an object representing the platform-specific OpenGL drawing
// context.
//
// GLContexts are created for a Window using gdk_window_create_gl_context(), and
// the context will match the Visual of the window.
//
// A GLContext is not tied to any particular normal framebuffer. For instance,
// it cannot draw to the Window back buffer. The GDK repaint system is in full
// control of the painting to that. Instead, you can create render buffers or
// textures and use gdk_cairo_draw_from_gl() in the draw function of your widget
// to draw them. Then GDK will handle the integration of your rendering with
// that of other widgets.
//
// Support for GLContext is platform-specific, context creation can fail,
// returning nil context.
//
// A GLContext has to be made "current" in order to start using it, otherwise
// any OpenGL call will be ignored.
//
// Creating a new OpenGL context ##
//
// In order to create a new GLContext instance you need a Window, which you
// typically get during the realize call of a widget.
//
// A GLContext is not realized until either gdk_gl_context_make_current(), or
// until it is realized using gdk_gl_context_realize(). It is possible to
// specify details of the GL context like the OpenGL version to be used, or
// whether the GL context should have extra state validation enabled after
// calling gdk_window_create_gl_context() by calling gdk_gl_context_realize().
// If the realization fails you have the option to change the settings of the
// GLContext and try again.
//
// Using a GdkGLContext ##
//
// You will need to make the GLContext the current context before issuing OpenGL
// calls; the system sends OpenGL commands to whichever context is current. It
// is possible to have multiple contexts, so you always need to ensure that the
// one which you want to draw with is the current one before issuing commands:
//
//    gdk_gl_context_make_current (context);
//
// You can now perform your drawing using OpenGL commands.
//
// You can check which GLContext is the current one by using
// gdk_gl_context_get_current(); you can also unset any GLContext that is
// currently set by calling gdk_gl_context_clear_current().
type GLContext interface {
	gextras.Objector

	DebugEnabled() bool

	Display() Display

	ForwardCompatible() bool

	RequiredVersion() (major int, minor int)

	SharedContext() GLContext

	UseES() bool

	Version() (major int, minor int)

	Window() Window

	IsLegacyGLContext() bool

	MakeCurrentGLContext()

	RealizeGLContext() error

	SetDebugEnabledGLContext(enabled bool)

	SetForwardCompatibleGLContext(compatible bool)

	SetRequiredVersionGLContext(major int, minor int)

	SetUseESGLContext(useEs int)
}

// glContext implements the GLContext class.
type glContext struct {
	gextras.Objector
}

// WrapGLContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapGLContext(obj *externglib.Object) GLContext {
	return glContext{
		Objector: obj,
	}
}

func marshalGLContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGLContext(obj), nil
}

func (c glContext) DebugEnabled() bool {
	var _arg0 *C.GdkGLContext // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_gl_context_get_debug_enabled(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c glContext) Display() Display {
	var _arg0 *C.GdkGLContext // out
	var _cret *C.GdkDisplay   // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_gl_context_get_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Display)

	return _display
}

func (c glContext) ForwardCompatible() bool {
	var _arg0 *C.GdkGLContext // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_gl_context_get_forward_compatible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c glContext) RequiredVersion() (major int, minor int) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.int           // in
	var _arg2 C.int           // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(c.Native()))

	C.gdk_gl_context_get_required_version(_arg0, &_arg1, &_arg2)

	var _major int // out
	var _minor int // out

	_major = int(_arg1)
	_minor = int(_arg2)

	return _major, _minor
}

func (c glContext) SharedContext() GLContext {
	var _arg0 *C.GdkGLContext // out
	var _cret *C.GdkGLContext // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_gl_context_get_shared_context(_arg0)

	var _glContext GLContext // out

	_glContext = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(GLContext)

	return _glContext
}

func (c glContext) UseES() bool {
	var _arg0 *C.GdkGLContext // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_gl_context_get_use_es(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c glContext) Version() (major int, minor int) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.int           // in
	var _arg2 C.int           // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(c.Native()))

	C.gdk_gl_context_get_version(_arg0, &_arg1, &_arg2)

	var _major int // out
	var _minor int // out

	_major = int(_arg1)
	_minor = int(_arg2)

	return _major, _minor
}

func (c glContext) Window() Window {
	var _arg0 *C.GdkGLContext // out
	var _cret *C.GdkWindow    // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_gl_context_get_window(_arg0)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _window
}

func (c glContext) IsLegacyGLContext() bool {
	var _arg0 *C.GdkGLContext // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_gl_context_is_legacy(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c glContext) MakeCurrentGLContext() {
	var _arg0 *C.GdkGLContext // out

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(c.Native()))

	C.gdk_gl_context_make_current(_arg0)
}

func (c glContext) RealizeGLContext() error {
	var _arg0 *C.GdkGLContext // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(c.Native()))

	C.gdk_gl_context_realize(_arg0, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (c glContext) SetDebugEnabledGLContext(enabled bool) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(c.Native()))
	if enabled {
		_arg1 = C.TRUE
	}

	C.gdk_gl_context_set_debug_enabled(_arg0, _arg1)
}

func (c glContext) SetForwardCompatibleGLContext(compatible bool) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(c.Native()))
	if compatible {
		_arg1 = C.TRUE
	}

	C.gdk_gl_context_set_forward_compatible(_arg0, _arg1)
}

func (c glContext) SetRequiredVersionGLContext(major int, minor int) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.int           // out
	var _arg2 C.int           // out

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(major)
	_arg2 = C.int(minor)

	C.gdk_gl_context_set_required_version(_arg0, _arg1, _arg2)
}

func (c glContext) SetUseESGLContext(useEs int) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.int           // out

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(useEs)

	C.gdk_gl_context_set_use_es(_arg0, _arg1)
}
