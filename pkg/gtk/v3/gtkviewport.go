// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_viewport_get_type()), F: marshalViewport},
	})
}

// Viewport: the Viewport widget acts as an adaptor class, implementing
// scrollability for child widgets that lack their own scrolling capabilities.
// Use GtkViewport to scroll child widgets such as Grid, Box, and so on.
//
// If a widget has native scrolling abilities, such as TextView, TreeView or
// IconView, it can be added to a ScrolledWindow with gtk_container_add(). If a
// widget does not, you must first add the widget to a Viewport, then add the
// viewport to the scrolled window. gtk_container_add() does this automatically
// if a child that does not implement Scrollable is added to a ScrolledWindow,
// so you can ignore the presence of the viewport.
//
// The GtkViewport will start scrolling content only if allocated less than the
// child widget’s minimum size in a given orientation.
//
//
// CSS nodes
//
// GtkViewport has a single CSS node with name viewport.
type Viewport interface {
	Bin
	Buildable
	Scrollable

	// BinWindow gets the bin window of the Viewport.
	BinWindow() gdk.Window
	// HAdjustment returns the horizontal adjustment of the viewport.
	HAdjustment() Adjustment
	// ShadowType gets the shadow type of the Viewport. See
	// gtk_viewport_set_shadow_type().
	ShadowType() ShadowType
	// VAdjustment returns the vertical adjustment of the viewport.
	VAdjustment() Adjustment
	// ViewWindow gets the view window of the Viewport.
	ViewWindow() gdk.Window
	// SetHAdjustment sets the horizontal adjustment of the viewport.
	SetHAdjustment(adjustment Adjustment)
	// SetShadowType sets the shadow type of the viewport.
	SetShadowType(typ ShadowType)
	// SetVAdjustment sets the vertical adjustment of the viewport.
	SetVAdjustment(adjustment Adjustment)
}

// viewport implements the Viewport interface.
type viewport struct {
	Bin
	Buildable
	Scrollable
}

var _ Viewport = (*viewport)(nil)

// WrapViewport wraps a GObject to the right type. It is
// primarily used internally.
func WrapViewport(obj *externglib.Object) Viewport {
	return Viewport{
		Bin:        WrapBin(obj),
		Buildable:  WrapBuildable(obj),
		Scrollable: WrapScrollable(obj),
	}
}

func marshalViewport(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapViewport(obj), nil
}

// NewViewport constructs a class Viewport.
func NewViewport(hadjustment Adjustment, vadjustment Adjustment) Viewport {
	var arg1 *C.GtkAdjustment
	var arg2 *C.GtkAdjustment

	arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))
	arg2 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	var cret C.GtkViewport
	var ret1 Viewport

	cret = C.gtk_viewport_new(hadjustment, vadjustment)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Viewport)

	return ret1
}

// BinWindow gets the bin window of the Viewport.
func (v viewport) BinWindow() gdk.Window {
	var arg0 *C.GtkViewport

	arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	var cret *C.GdkWindow
	var ret1 gdk.Window

	cret = C.gtk_viewport_get_bin_window(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Window)

	return ret1
}

// HAdjustment returns the horizontal adjustment of the viewport.
func (v viewport) HAdjustment() Adjustment {
	var arg0 *C.GtkViewport

	arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	var cret *C.GtkAdjustment
	var ret1 Adjustment

	cret = C.gtk_viewport_get_hadjustment(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Adjustment)

	return ret1
}

// ShadowType gets the shadow type of the Viewport. See
// gtk_viewport_set_shadow_type().
func (v viewport) ShadowType() ShadowType {
	var arg0 *C.GtkViewport

	arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	var cret C.GtkShadowType
	var ret1 ShadowType

	cret = C.gtk_viewport_get_shadow_type(arg0)

	ret1 = ShadowType(cret)

	return ret1
}

// VAdjustment returns the vertical adjustment of the viewport.
func (v viewport) VAdjustment() Adjustment {
	var arg0 *C.GtkViewport

	arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	var cret *C.GtkAdjustment
	var ret1 Adjustment

	cret = C.gtk_viewport_get_vadjustment(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Adjustment)

	return ret1
}

// ViewWindow gets the view window of the Viewport.
func (v viewport) ViewWindow() gdk.Window {
	var arg0 *C.GtkViewport

	arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	var cret *C.GdkWindow
	var ret1 gdk.Window

	cret = C.gtk_viewport_get_view_window(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Window)

	return ret1
}

// SetHAdjustment sets the horizontal adjustment of the viewport.
func (v viewport) SetHAdjustment(adjustment Adjustment) {
	var arg0 *C.GtkViewport
	var arg1 *C.GtkAdjustment

	arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))
	arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_viewport_set_hadjustment(arg0, adjustment)
}

// SetShadowType sets the shadow type of the viewport.
func (v viewport) SetShadowType(typ ShadowType) {
	var arg0 *C.GtkViewport
	var arg1 C.GtkShadowType

	arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))
	arg1 = (C.GtkShadowType)(typ)

	C.gtk_viewport_set_shadow_type(arg0, typ)
}

// SetVAdjustment sets the vertical adjustment of the viewport.
func (v viewport) SetVAdjustment(adjustment Adjustment) {
	var arg0 *C.GtkViewport
	var arg1 *C.GtkAdjustment

	arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))
	arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_viewport_set_vadjustment(arg0, adjustment)
}

type ViewportPrivate struct {
	native C.GtkViewportPrivate
}

// WrapViewportPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapViewportPrivate(ptr unsafe.Pointer) *ViewportPrivate {
	if ptr == nil {
		return nil
	}

	return (*ViewportPrivate)(ptr)
}

func marshalViewportPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapViewportPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (v *ViewportPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&v.native)
}
