// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeViewport = coreglib.Type(C.gtk_viewport_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeViewport, F: marshalViewport},
	})
}

// ViewportOverrider contains methods that are overridable.
type ViewportOverrider interface {
}

// Viewport widget acts as an adaptor class, implementing scrollability for
// child widgets that lack their own scrolling capabilities. Use GtkViewport to
// scroll child widgets such as Grid, Box, and so on.
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
type Viewport struct {
	_ [0]func() // equal guard
	Bin

	*coreglib.Object
	Scrollable
}

var (
	_ Binner            = (*Viewport)(nil)
	_ coreglib.Objector = (*Viewport)(nil)
)

func initClassViewport(gclass unsafe.Pointer, goval any) {
}

func wrapViewport(obj *coreglib.Object) *Viewport {
	return &Viewport{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
		},
		Object: obj,
		Scrollable: Scrollable{
			Object: obj,
		},
	}
}

func marshalViewport(p uintptr) (interface{}, error) {
	return wrapViewport(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewViewport creates a new Viewport with the given adjustments, or with
// default adjustments if none are given.
//
// The function takes the following parameters:
//
//    - hadjustment (optional): horizontal adjustment.
//    - vadjustment (optional): vertical adjustment.
//
// The function returns the following values:
//
//    - viewport: new Viewport.
//
func NewViewport(hadjustment, vadjustment *Adjustment) *Viewport {
	var _arg1 *C.GtkAdjustment // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	if hadjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(hadjustment).Native()))
	}
	if vadjustment != nil {
		_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(vadjustment).Native()))
	}

	_cret = C.gtk_viewport_new(_arg1, _arg2)
	runtime.KeepAlive(hadjustment)
	runtime.KeepAlive(vadjustment)

	var _viewport *Viewport // out

	_viewport = wrapViewport(coreglib.Take(unsafe.Pointer(_cret)))

	return _viewport
}

// BinWindow gets the bin window of the Viewport.
//
// The function returns the following values:
//
//    - window: Window.
//
func (viewport *Viewport) BinWindow() gdk.Windower {
	var _arg0 *C.GtkViewport // out
	var _cret *C.GdkWindow   // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))

	_cret = C.gtk_viewport_get_bin_window(_arg0)
	runtime.KeepAlive(viewport)

	var _window gdk.Windower // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.Windower)
			return ok
		})
		rv, ok := casted.(gdk.Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_window = rv
	}

	return _window
}

// HAdjustment returns the horizontal adjustment of the viewport.
//
// Deprecated: Use gtk_scrollable_get_hadjustment().
//
// The function returns the following values:
//
//    - adjustment: horizontal adjustment of viewport.
//
func (viewport *Viewport) HAdjustment() *Adjustment {
	var _arg0 *C.GtkViewport   // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))

	_cret = C.gtk_viewport_get_hadjustment(_arg0)
	runtime.KeepAlive(viewport)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// ShadowType gets the shadow type of the Viewport. See
// gtk_viewport_set_shadow_type().
//
// The function returns the following values:
//
//    - shadowType: shadow type.
//
func (viewport *Viewport) ShadowType() ShadowType {
	var _arg0 *C.GtkViewport  // out
	var _cret C.GtkShadowType // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))

	_cret = C.gtk_viewport_get_shadow_type(_arg0)
	runtime.KeepAlive(viewport)

	var _shadowType ShadowType // out

	_shadowType = ShadowType(_cret)

	return _shadowType
}

// VAdjustment returns the vertical adjustment of the viewport.
//
// Deprecated: Use gtk_scrollable_get_vadjustment().
//
// The function returns the following values:
//
//    - adjustment: vertical adjustment of viewport.
//
func (viewport *Viewport) VAdjustment() *Adjustment {
	var _arg0 *C.GtkViewport   // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))

	_cret = C.gtk_viewport_get_vadjustment(_arg0)
	runtime.KeepAlive(viewport)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// ViewWindow gets the view window of the Viewport.
//
// The function returns the following values:
//
//    - window: Window.
//
func (viewport *Viewport) ViewWindow() gdk.Windower {
	var _arg0 *C.GtkViewport // out
	var _cret *C.GdkWindow   // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))

	_cret = C.gtk_viewport_get_view_window(_arg0)
	runtime.KeepAlive(viewport)

	var _window gdk.Windower // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.Windower)
			return ok
		})
		rv, ok := casted.(gdk.Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_window = rv
	}

	return _window
}

// SetHAdjustment sets the horizontal adjustment of the viewport.
//
// Deprecated: Use gtk_scrollable_set_hadjustment().
//
// The function takes the following parameters:
//
//    - adjustment (optional): Adjustment.
//
func (viewport *Viewport) SetHAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkViewport   // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))
	if adjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}

	C.gtk_viewport_set_hadjustment(_arg0, _arg1)
	runtime.KeepAlive(viewport)
	runtime.KeepAlive(adjustment)
}

// SetShadowType sets the shadow type of the viewport.
//
// The function takes the following parameters:
//
//    - typ: new shadow type.
//
func (viewport *Viewport) SetShadowType(typ ShadowType) {
	var _arg0 *C.GtkViewport  // out
	var _arg1 C.GtkShadowType // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))
	_arg1 = C.GtkShadowType(typ)

	C.gtk_viewport_set_shadow_type(_arg0, _arg1)
	runtime.KeepAlive(viewport)
	runtime.KeepAlive(typ)
}

// SetVAdjustment sets the vertical adjustment of the viewport.
//
// Deprecated: Use gtk_scrollable_set_vadjustment().
//
// The function takes the following parameters:
//
//    - adjustment (optional): Adjustment.
//
func (viewport *Viewport) SetVAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkViewport   // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))
	if adjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}

	C.gtk_viewport_set_vadjustment(_arg0, _arg1)
	runtime.KeepAlive(viewport)
	runtime.KeepAlive(adjustment)
}
