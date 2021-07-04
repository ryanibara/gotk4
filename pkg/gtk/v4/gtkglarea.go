// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gl_area_get_type()), F: marshalGLArea},
	})
}

// GLArea: `GtkGLArea` is a widget that allows drawing with OpenGL.
//
// !An example GtkGLArea (glarea.png)
//
// `GtkGLArea` sets up its own [class@Gdk.GLContext], and creates a custom GL
// framebuffer that the widget will do GL rendering onto. It also ensures that
// this framebuffer is the default GL rendering target when rendering.
//
// In order to draw, you have to connect to the [signal@Gtk.GLArea::render]
// signal, or subclass `GtkGLArea` and override the GtkGLAreaClass.render
// virtual function.
//
// The `GtkGLArea` widget ensures that the `GdkGLContext` is associated with the
// widget's drawing area, and it is kept updated when the size and position of
// the drawing area changes.
//
//
// Drawing with GtkGLArea
//
// The simplest way to draw using OpenGL commands in a `GtkGLArea` is to create
// a widget instance and connect to the [signal@Gtk.GLArea::render] signal:
//
// The `render()` function will be called when the `GtkGLArea` is ready for you
// to draw its content:
//
// “`c static gboolean render (GtkGLArea *area, GdkGLContext *context) { //
// inside this function it's safe to use GL; the given // GLContext has been
// made current to the drawable // surface used by the `GtkGLArea` and the
// viewport has // already been set to be the size of the allocation
//
//    // we can start by clearing the buffer
//    glClearColor (0, 0, 0, 0);
//    glClear (GL_COLOR_BUFFER_BIT);
//
//    // draw your object
//    // draw_an_object ();
//
//    // we completed our drawing; the draw commands will be
//    // flushed at the end of the signal emission chain, and
//    // the buffers will be drawn on the window
//    return TRUE;
//
// }
//
// void setup_glarea (void) { // create a GtkGLArea instance GtkWidget *gl_area
// = gtk_gl_area_new ();
//
//    // connect to the "render" signal
//    g_signal_connect (gl_area, "render", G_CALLBACK (render), NULL);
//
// } “`
//
// If you need to initialize OpenGL state, e.g. buffer objects or shaders, you
// should use the [signal@Gtk.Widget::realize] signal; you can use the
// [signal@Gtk.Widget::unrealize] signal to clean up. Since the `GdkGLContext`
// creation and initialization may fail, you will need to check for errors,
// using [method@Gtk.GLArea.get_error].
//
// An example of how to safely initialize the GL state is:
//
// “`c static void on_realize (GtkGLarea *area) { // We need to make the context
// current if we want to // call GL API gtk_gl_area_make_current (area);
//
//    // If there were errors during the initialization or
//    // when trying to make the context current, this
//    // function will return a #GError for you to catch
//    if (gtk_gl_area_get_error (area) != NULL)
//      return;
//
//    // You can also use gtk_gl_area_set_error() in order
//    // to show eventual initialization errors on the
//    // GtkGLArea widget itself
//    GError *internal_error = NULL;
//    init_buffer_objects (&error);
//    if (error != NULL)
//      {
//        gtk_gl_area_set_error (area, error);
//        g_error_free (error);
//        return;
//      }
//
//    init_shaders (&error);
//    if (error != NULL)
//      {
//        gtk_gl_area_set_error (area, error);
//        g_error_free (error);
//        return;
//      }
//
// } “`
//
// If you need to change the options for creating the `GdkGLContext` you should
// use the [signal@Gtk.GLArea::create-context] signal.
type GLArea interface {
	Widget

	AttachBuffersGLArea()

	AutoRender() bool

	Context() gdk.GLContext

	Error() error

	HasDepthBuffer() bool

	HasStencilBuffer() bool

	RequiredVersion() (major int, minor int)

	UseES() bool

	MakeCurrentGLArea()

	QueueRenderGLArea()

	SetAutoRenderGLArea(autoRender bool)

	SetErrorGLArea(err error)

	SetHasDepthBufferGLArea(hasDepthBuffer bool)

	SetHasStencilBufferGLArea(hasStencilBuffer bool)

	SetRequiredVersionGLArea(major int, minor int)

	SetUseESGLArea(useEs bool)
}

// glArea implements the GLArea class.
type glArea struct {
	Widget
}

// WrapGLArea wraps a GObject to the right type. It is
// primarily used internally.
func WrapGLArea(obj *externglib.Object) GLArea {
	return glArea{
		Widget: WrapWidget(obj),
	}
}

func marshalGLArea(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGLArea(obj), nil
}

func NewGLArea() GLArea {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_gl_area_new()

	var _glArea GLArea // out

	_glArea = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(GLArea)

	return _glArea
}

func (a glArea) AttachBuffersGLArea() {
	var _arg0 *C.GtkGLArea // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	C.gtk_gl_area_attach_buffers(_arg0)
}

func (a glArea) AutoRender() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_gl_area_get_auto_render(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a glArea) Context() gdk.GLContext {
	var _arg0 *C.GtkGLArea    // out
	var _cret *C.GdkGLContext // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_gl_area_get_context(_arg0)

	var _glContext gdk.GLContext // out

	_glContext = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.GLContext)

	return _glContext
}

func (a glArea) Error() error {
	var _arg0 *C.GtkGLArea // out
	var _cret *C.GError    // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_gl_area_get_error(_arg0)

	var _err error // out

	_err = gerror.Take(unsafe.Pointer(_cret))

	return _err
}

func (a glArea) HasDepthBuffer() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_gl_area_get_has_depth_buffer(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a glArea) HasStencilBuffer() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_gl_area_get_has_stencil_buffer(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a glArea) RequiredVersion() (major int, minor int) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.int        // in
	var _arg2 C.int        // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	C.gtk_gl_area_get_required_version(_arg0, &_arg1, &_arg2)

	var _major int // out
	var _minor int // out

	_major = int(_arg1)
	_minor = int(_arg2)

	return _major, _minor
}

func (a glArea) UseES() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_gl_area_get_use_es(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a glArea) MakeCurrentGLArea() {
	var _arg0 *C.GtkGLArea // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	C.gtk_gl_area_make_current(_arg0)
}

func (a glArea) QueueRenderGLArea() {
	var _arg0 *C.GtkGLArea // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	C.gtk_gl_area_queue_render(_arg0)
}

func (a glArea) SetAutoRenderGLArea(autoRender bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))
	if autoRender {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_auto_render(_arg0, _arg1)
}

func (a glArea) SetErrorGLArea(err error) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 *C.GError    // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GError)(gerror.New(err))
	defer C.g_error_free(_arg1)

	C.gtk_gl_area_set_error(_arg0, _arg1)
}

func (a glArea) SetHasDepthBufferGLArea(hasDepthBuffer bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))
	if hasDepthBuffer {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_has_depth_buffer(_arg0, _arg1)
}

func (a glArea) SetHasStencilBufferGLArea(hasStencilBuffer bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))
	if hasStencilBuffer {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_has_stencil_buffer(_arg0, _arg1)
}

func (a glArea) SetRequiredVersionGLArea(major int, minor int) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.int        // out
	var _arg2 C.int        // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))
	_arg1 = C.int(major)
	_arg2 = C.int(minor)

	C.gtk_gl_area_set_required_version(_arg0, _arg1, _arg2)
}

func (a glArea) SetUseESGLArea(useEs bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))
	if useEs {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_use_es(_arg0, _arg1)
}

func (s glArea) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s glArea) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s glArea) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s glArea) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s glArea) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s glArea) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s glArea) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b glArea) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
