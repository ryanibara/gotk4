// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gl_area_get_type()), F: marshalGLAreaer},
	})
}

// GLAreaOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type GLAreaOverrider interface {
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	Render(context gdk.GLContexter) bool
	// The function takes the following parameters:
	//
	//    - width
	//    - height
	//
	Resize(width, height int)
}

// GLArea: GtkGLArea is a widget that allows drawing with OpenGL.
//
// !An example GtkGLArea (glarea.png)
//
// GtkGLArea sets up its own gdk.GLContext, and creates a custom GL framebuffer
// that the widget will do GL rendering onto. It also ensures that this
// framebuffer is the default GL rendering target when rendering.
//
// In order to draw, you have to connect to the gtk.GLArea::render signal, or
// subclass GtkGLArea and override the GtkGLAreaClass.render virtual function.
//
// The GtkGLArea widget ensures that the GdkGLContext is associated with the
// widget's drawing area, and it is kept updated when the size and position of
// the drawing area changes.
//
//
// Drawing with GtkGLArea
//
// The simplest way to draw using OpenGL commands in a GtkGLArea is to create a
// widget instance and connect to the gtk.GLArea::render signal:
//
// The render() function will be called when the GtkGLArea is ready for you to
// draw its content:
//
//    static gboolean
//    render (GtkGLArea *area, GdkGLContext *context)
//    {
//      // inside this function it's safe to use GL; the given
//      // GLContext has been made current to the drawable
//      // surface used by the GtkGLArea and the viewport has
//      // already been set to be the size of the allocation
//
//      // we can start by clearing the buffer
//      glClearColor (0, 0, 0, 0);
//      glClear (GL_COLOR_BUFFER_BIT);
//
//      // draw your object
//      // draw_an_object ();
//
//      // we completed our drawing; the draw commands will be
//      // flushed at the end of the signal emission chain, and
//      // the buffers will be drawn on the window
//      return TRUE;
//    }
//
//    void setup_glarea (void)
//    {
//      // create a GtkGLArea instance
//      GtkWidget *gl_area = gtk_gl_area_new ();
//
//      // connect to the "render" signal
//      g_signal_connect (gl_area, "render", G_CALLBACK (render), NULL);
//    }
//
//
// If you need to initialize OpenGL state, e.g. buffer objects or shaders, you
// should use the gtk.Widget::realize signal; you can use the
// gtk.Widget::unrealize signal to clean up. Since the GdkGLContext creation and
// initialization may fail, you will need to check for errors, using
// gtk.GLArea.GetError().
//
// An example of how to safely initialize the GL state is:
//
//    static void
//    on_realize (GtkGLarea *area)
//    {
//      // We need to make the context current if we want to
//      // call GL API
//      gtk_gl_area_make_current (area);
//
//      // If there were errors during the initialization or
//      // when trying to make the context current, this
//      // function will return a #GError for you to catch
//      if (gtk_gl_area_get_error (area) != NULL)
//        return;
//
//      // You can also use gtk_gl_area_set_error() in order
//      // to show eventual initialization errors on the
//      // GtkGLArea widget itself
//      GError *internal_error = NULL;
//      init_buffer_objects (&error);
//      if (error != NULL)
//        {
//          gtk_gl_area_set_error (area, error);
//          g_error_free (error);
//          return;
//        }
//
//      init_shaders (&error);
//      if (error != NULL)
//        {
//          gtk_gl_area_set_error (area, error);
//          g_error_free (error);
//          return;
//        }
//    }
//
//
// If you need to change the options for creating the GdkGLContext you should
// use the gtk.GLArea::create-context signal.
type GLArea struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*GLArea)(nil)
)

func wrapGLArea(obj *externglib.Object) *GLArea {
	return &GLArea{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalGLAreaer(p uintptr) (interface{}, error) {
	return wrapGLArea(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectCreateContext: emitted when the widget is being realized.
//
// This allows you to override how the GL context is created. This is useful
// when you want to reuse an existing GL context, or if you want to try creating
// different kinds of GL options.
//
// If context creation fails then the signal handler can use
// gtk.GLArea.SetError() to register a more detailed error of how the
// construction failed.
func (area *GLArea) ConnectCreateContext(f func() gdk.GLContexter) externglib.SignalHandle {
	return area.Connect("create-context", f)
}

// ConnectRender: emitted every time the contents of the GtkGLArea should be
// redrawn.
//
// The context is bound to the area prior to emitting this function, and the
// buffers are painted to the window once the emission terminates.
func (area *GLArea) ConnectRender(f func(context gdk.GLContexter) bool) externglib.SignalHandle {
	return area.Connect("render", f)
}

// ConnectResize: emitted once when the widget is realized, and then each time
// the widget is changed while realized.
//
// This is useful in order to keep GL state up to date with the widget size,
// like for instance camera properties which may depend on the width/height
// ratio.
//
// The GL context for the area is guaranteed to be current when this signal is
// emitted.
//
// The default handler sets up the GL viewport.
func (area *GLArea) ConnectResize(f func(width, height int)) externglib.SignalHandle {
	return area.Connect("resize", f)
}

// NewGLArea creates a new GtkGLArea widget.
//
// The function returns the following values:
//
//    - glArea: new GtkGLArea.
//
func NewGLArea() *GLArea {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_gl_area_new()

	var _glArea *GLArea // out

	_glArea = wrapGLArea(externglib.Take(unsafe.Pointer(_cret)))

	return _glArea
}

// AttachBuffers binds buffers to the framebuffer.
//
// Ensures that the area framebuffer object is made the current draw and read
// target, and that all the required buffers for the area are created and bound
// to the framebuffer.
//
// This function is automatically called before emitting the gtk.GLArea::render
// signal, and doesn't normally need to be called by application code.
func (area *GLArea) AttachBuffers() {
	var _arg0 *C.GtkGLArea // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	C.gtk_gl_area_attach_buffers(_arg0)
	runtime.KeepAlive(area)
}

// AutoRender returns whether the area is in auto render mode or not.
//
// The function returns the following values:
//
//    - ok: TRUE if the area is auto rendering, FALSE otherwise.
//
func (area *GLArea) AutoRender() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_gl_area_get_auto_render(_arg0)
	runtime.KeepAlive(area)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Context retrieves the GdkGLContext used by area.
//
// The function returns the following values:
//
//    - glContext: GdkGLContext.
//
func (area *GLArea) Context() gdk.GLContexter {
	var _arg0 *C.GtkGLArea    // out
	var _cret *C.GdkGLContext // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_gl_area_get_context(_arg0)
	runtime.KeepAlive(area)

	var _glContext gdk.GLContexter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.GLContexter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(gdk.GLContexter)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gdk.GLContexter")
		}
		_glContext = rv
	}

	return _glContext
}

// Error gets the current error set on the area.
//
// The function returns the following values:
//
//    - err (optional) or NULL.
//
func (area *GLArea) Error() error {
	var _arg0 *C.GtkGLArea // out
	var _cret *C.GError    // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_gl_area_get_error(_arg0)
	runtime.KeepAlive(area)

	var _err error // out

	if _cret != nil {
		_err = gerror.Take(unsafe.Pointer(_cret))
	}

	return _err
}

// HasDepthBuffer returns whether the area has a depth buffer.
//
// The function returns the following values:
//
//    - ok: TRUE if the area has a depth buffer, FALSE otherwise.
//
func (area *GLArea) HasDepthBuffer() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_gl_area_get_has_depth_buffer(_arg0)
	runtime.KeepAlive(area)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasStencilBuffer returns whether the area has a stencil buffer.
//
// The function returns the following values:
//
//    - ok: TRUE if the area has a stencil buffer, FALSE otherwise.
//
func (area *GLArea) HasStencilBuffer() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_gl_area_get_has_stencil_buffer(_arg0)
	runtime.KeepAlive(area)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RequiredVersion retrieves the required version of OpenGL.
//
// See gtk.GLArea.SetRequiredVersion().
//
// The function returns the following values:
//
//    - major: return location for the required major version.
//    - minor: return location for the required minor version.
//
func (area *GLArea) RequiredVersion() (major int, minor int) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.int        // in
	var _arg2 C.int        // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	C.gtk_gl_area_get_required_version(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(area)

	var _major int // out
	var _minor int // out

	_major = int(_arg1)
	_minor = int(_arg2)

	return _major, _minor
}

// UseES returns whether the GtkGLArea should use OpenGL ES.
//
// See gtk.GLArea.SetUseES().
//
// The function returns the following values:
//
//    - ok: TRUE if the GtkGLArea should create an OpenGL ES context and FALSE
//      otherwise.
//
func (area *GLArea) UseES() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_gl_area_get_use_es(_arg0)
	runtime.KeepAlive(area)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MakeCurrent ensures that the GdkGLContext used by area is associated with the
// GtkGLArea.
//
// This function is automatically called before emitting the gtk.GLArea::render
// signal, and doesn't normally need to be called by application code.
func (area *GLArea) MakeCurrent() {
	var _arg0 *C.GtkGLArea // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	C.gtk_gl_area_make_current(_arg0)
	runtime.KeepAlive(area)
}

// QueueRender marks the currently rendered data (if any) as invalid, and queues
// a redraw of the widget.
//
// This ensures that the gtk.GLArea::render signal is emitted during the draw.
//
// This is only needed when gtk.GLArea.SetAutoRender() has been called with a
// FALSE value. The default behaviour is to emit gtk.GLArea::render on each
// draw.
func (area *GLArea) QueueRender() {
	var _arg0 *C.GtkGLArea // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	C.gtk_gl_area_queue_render(_arg0)
	runtime.KeepAlive(area)
}

// SetAutoRender sets whether the GtkGLArea is in auto render mode.
//
// If auto_render is TRUE the gtk.GLArea::render signal will be emitted every
// time the widget draws. This is the default and is useful if drawing the
// widget is faster.
//
// If auto_render is FALSE the data from previous rendering is kept around and
// will be used for drawing the widget the next time, unless the window is
// resized. In order to force a rendering gtk.GLArea.QueueRender() must be
// called. This mode is useful when the scene changes seldom, but takes a long
// time to redraw.
//
// The function takes the following parameters:
//
//    - autoRender: boolean.
//
func (area *GLArea) SetAutoRender(autoRender bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))
	if autoRender {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_auto_render(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(autoRender)
}

// SetError sets an error on the area which will be shown instead of the GL
// rendering.
//
// This is useful in the gtk.GLArea::create-context signal if GL context
// creation fails.
//
// The function takes the following parameters:
//
//    - err (optional): new GError, or NULL to unset the error.
//
func (area *GLArea) SetError(err error) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 *C.GError    // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))
	_arg1 = (*C.GError)(gerror.New(err))

	C.gtk_gl_area_set_error(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(err)
}

// SetHasDepthBuffer sets whether the GtkGLArea should use a depth buffer.
//
// If has_depth_buffer is TRUE the widget will allocate and enable a depth
// buffer for the target framebuffer. Otherwise there will be none.
//
// The function takes the following parameters:
//
//    - hasDepthBuffer: TRUE to add a depth buffer.
//
func (area *GLArea) SetHasDepthBuffer(hasDepthBuffer bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))
	if hasDepthBuffer {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_has_depth_buffer(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(hasDepthBuffer)
}

// SetHasStencilBuffer sets whether the GtkGLArea should use a stencil buffer.
//
// If has_stencil_buffer is TRUE the widget will allocate and enable a stencil
// buffer for the target framebuffer. Otherwise there will be none.
//
// The function takes the following parameters:
//
//    - hasStencilBuffer: TRUE to add a stencil buffer.
//
func (area *GLArea) SetHasStencilBuffer(hasStencilBuffer bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))
	if hasStencilBuffer {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_has_stencil_buffer(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(hasStencilBuffer)
}

// SetRequiredVersion sets the required version of OpenGL to be used when
// creating the context for the widget.
//
// This function must be called before the area has been realized.
//
// The function takes the following parameters:
//
//    - major version.
//    - minor version.
//
func (area *GLArea) SetRequiredVersion(major, minor int) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.int        // out
	var _arg2 C.int        // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))
	_arg1 = C.int(major)
	_arg2 = C.int(minor)

	C.gtk_gl_area_set_required_version(_arg0, _arg1, _arg2)
	runtime.KeepAlive(area)
	runtime.KeepAlive(major)
	runtime.KeepAlive(minor)
}

// SetUseES sets whether the area should create an OpenGL or an OpenGL ES
// context.
//
// You should check the capabilities of the GLContext before drawing with either
// API.
//
// The function takes the following parameters:
//
//    - useEs: whether to use OpenGL or OpenGL ES.
//
func (area *GLArea) SetUseES(useEs bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))
	if useEs {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_use_es(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(useEs)
}
