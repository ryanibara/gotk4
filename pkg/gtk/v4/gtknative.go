// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gsk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_native_get_type()), F: marshalNative},
	})
}

// NativeGetForSurface finds the `GtkNative` associated with the surface.
func NativeGetForSurface(surface gdk.Surface) Native {
	var arg1 *C.GdkSurface

	arg1 = (*C.GdkSurface)(surface.Native())

	ret := C.gtk_native_get_for_surface(arg1)

	var ret0 Native

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Native)

	return ret0
}

// Native: `GtkNative` is the interface implemented by all widgets that have
// their own `GdkSurface`.
//
// The obvious example of a `GtkNative` is `GtkWindow`.
//
// Every widget that is not itself a `GtkNative` is contained in one, and you
// can get it with [method@Gtk.Widget.get_native].
//
// To get the surface of a `GtkNative`, use [method@Gtk.Native.get_surface]. It
// is also possible to find the `GtkNative` to which a surface belongs, with
// [func@Gtk.Native.get_for_surface].
//
// In addition to a [class@Gdk.Surface], a `GtkNative` also provides a
// [class@Gsk.Renderer] for rendering on that surface. To get the renderer, use
// [method@Gtk.Native.get_renderer].
type Native interface {
	Widget

	// Renderer returns the renderer that is used for this `GtkNative`.
	Renderer() gsk.Renderer
	// Surface returns the surface of this `GtkNative`.
	Surface() gdk.Surface
	// SurfaceTransform retrieves the surface transform of @self.
	//
	// This is the translation from @self's surface coordinates into @self's
	// widget coordinates.
	SurfaceTransform() (x float64, y float64)
	// Realize realizes a `GtkNative`.
	//
	// This should only be used by subclasses.
	Realize()
	// Unrealize unrealizes a `GtkNative`.
	//
	// This should only be used by subclasses.
	Unrealize()
}

// native implements the Native interface.
type native struct {
	Widget
}

var _ Native = (*native)(nil)

// WrapNative wraps a GObject to a type that implements interface
// Native. It is primarily used internally.
func WrapNative(obj *externglib.Object) Native {
	return Native{
		Widget: WrapWidget(obj),
	}
}

func marshalNative(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNative(obj), nil
}

// Renderer returns the renderer that is used for this `GtkNative`.
func (s native) Renderer() gsk.Renderer {
	var arg0 *C.GtkNative

	arg0 = (*C.GtkNative)(s.Native())

	ret := C.gtk_native_get_renderer(arg0)

	var ret0 gsk.Renderer

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(gsk.Renderer)

	return ret0
}

// Surface returns the surface of this `GtkNative`.
func (s native) Surface() gdk.Surface {
	var arg0 *C.GtkNative

	arg0 = (*C.GtkNative)(s.Native())

	ret := C.gtk_native_get_surface(arg0)

	var ret0 gdk.Surface

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(gdk.Surface)

	return ret0
}

// SurfaceTransform retrieves the surface transform of @self.
//
// This is the translation from @self's surface coordinates into @self's
// widget coordinates.
func (s native) SurfaceTransform() (x float64, y float64) {
	var arg0 *C.GtkNative
	var arg1 *C.double // out
	var arg2 *C.double // out

	arg0 = (*C.GtkNative)(s.Native())

	C.gtk_native_get_surface_transform(arg0, &arg1, &arg2)

	var ret0 float64
	var ret1 float64

	ret0 = float64(arg1)

	ret1 = float64(arg2)

	return ret0, ret1
}

// Realize realizes a `GtkNative`.
//
// This should only be used by subclasses.
func (s native) Realize() {
	var arg0 *C.GtkNative

	arg0 = (*C.GtkNative)(s.Native())

	C.gtk_native_realize(arg0)
}

// Unrealize unrealizes a `GtkNative`.
//
// This should only be used by subclasses.
func (s native) Unrealize() {
	var arg0 *C.GtkNative

	arg0 = (*C.GtkNative)(s.Native())

	C.gtk_native_unrealize(arg0)
}
