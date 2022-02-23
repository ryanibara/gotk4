// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gdkpopup.go.
var GTypePopup = externglib.Type(C.gdk_popup_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypePopup, F: marshalPopup},
	})
}

// PopupOverrider contains methods that are overridable.
type PopupOverrider interface {
}

// Popup: GdkPopup is a surface that is attached to another surface.
//
// The GdkPopup is positioned relative to its parent surface.
//
// GdkPopups are typically used to implement menus and similar popups. They can
// be modal, which is indicated by the gdkpopup:autohide property.
type Popup struct {
	_ [0]func() // equal guard
	Surface
}

var (
	_ Surfacer = (*Popup)(nil)
)

// Popupper describes Popup's interface methods.
type Popupper interface {
	externglib.Objector

	// Autohide returns whether this popup is set to hide on outside clicks.
	Autohide() bool
	// Parent returns the parent surface of a popup.
	Parent() Surfacer
	// PositionX obtains the position of the popup relative to its parent.
	PositionX() int
	// PositionY obtains the position of the popup relative to its parent.
	PositionY() int
	// RectAnchor gets the current popup rectangle anchor.
	RectAnchor() Gravity
	// SurfaceAnchor gets the current popup surface anchor.
	SurfaceAnchor() Gravity
	// Present popup after having processed the PopupLayout rules.
	Present(width, height int, layout *PopupLayout) bool
}

var _ Popupper = (*Popup)(nil)

func ifaceInitPopupper(gifacePtr, data C.gpointer) {
}

func wrapPopup(obj *externglib.Object) *Popup {
	return &Popup{
		Surface: Surface{
			Object: obj,
		},
	}
}

func marshalPopup(p uintptr) (interface{}, error) {
	return wrapPopup(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Autohide returns whether this popup is set to hide on outside clicks.
//
// The function returns the following values:
//
//    - ok: TRUE if popup will autohide.
//
func (popup *Popup) Autohide() bool {
	var _arg0 *C.GdkPopup // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(popup.Native()))

	_cret = C.gdk_popup_get_autohide(_arg0)
	runtime.KeepAlive(popup)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Parent returns the parent surface of a popup.
//
// The function returns the following values:
//
//    - surface: parent surface.
//
func (popup *Popup) Parent() Surfacer {
	var _arg0 *C.GdkPopup   // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(popup.Native()))

	_cret = C.gdk_popup_get_parent(_arg0)
	runtime.KeepAlive(popup)

	var _surface Surfacer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Surfacer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Surfacer)
			return ok
		})
		rv, ok := casted.(Surfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Surfacer")
		}
		_surface = rv
	}

	return _surface
}

// PositionX obtains the position of the popup relative to its parent.
//
// The function returns the following values:
//
//    - gint: x coordinate of popup position.
//
func (popup *Popup) PositionX() int {
	var _arg0 *C.GdkPopup // out
	var _cret C.int       // in

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(popup.Native()))

	_cret = C.gdk_popup_get_position_x(_arg0)
	runtime.KeepAlive(popup)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PositionY obtains the position of the popup relative to its parent.
//
// The function returns the following values:
//
//    - gint: y coordinate of popup position.
//
func (popup *Popup) PositionY() int {
	var _arg0 *C.GdkPopup // out
	var _cret C.int       // in

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(popup.Native()))

	_cret = C.gdk_popup_get_position_y(_arg0)
	runtime.KeepAlive(popup)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// RectAnchor gets the current popup rectangle anchor.
//
// The value returned may change after calling gdk.Popup.Present(), or after the
// gdk.Surface::layout signal is emitted.
//
// The function returns the following values:
//
//    - gravity: current rectangle anchor value of popup.
//
func (popup *Popup) RectAnchor() Gravity {
	var _arg0 *C.GdkPopup  // out
	var _cret C.GdkGravity // in

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(popup.Native()))

	_cret = C.gdk_popup_get_rect_anchor(_arg0)
	runtime.KeepAlive(popup)

	var _gravity Gravity // out

	_gravity = Gravity(_cret)

	return _gravity
}

// SurfaceAnchor gets the current popup surface anchor.
//
// The value returned may change after calling gdk.Popup.Present(), or after the
// gdk.Surface::layout signal is emitted.
//
// The function returns the following values:
//
//    - gravity: current surface anchor value of popup.
//
func (popup *Popup) SurfaceAnchor() Gravity {
	var _arg0 *C.GdkPopup  // out
	var _cret C.GdkGravity // in

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(popup.Native()))

	_cret = C.gdk_popup_get_surface_anchor(_arg0)
	runtime.KeepAlive(popup)

	var _gravity Gravity // out

	_gravity = Gravity(_cret)

	return _gravity
}

// Present popup after having processed the PopupLayout rules.
//
// If the popup was previously now showing, it will be showed, otherwise it will
// change position according to layout.
//
// After calling this function, the result should be handled in response to the
// gdksurface::layout signal being emitted. The resulting popup position can be
// queried using gdk.Popup.GetPositionX(), gdk.Popup.GetPositionY(), and the
// resulting size will be sent as parameters in the layout signal. Use
// gdk.Popup.GetRectAnchor() and gdk.Popup.GetSurfaceAnchor() to get the
// resulting anchors.
//
// Presenting may fail, for example if the popup is set to autohide and is
// immediately hidden upon being presented. If presenting failed, the
// gdk.Surface::layout signal will not me emitted.
//
// The function takes the following parameters:
//
//    - width: unconstrained popup width to layout.
//    - height: unconstrained popup height to layout.
//    - layout: GdkPopupLayout object used to layout.
//
// The function returns the following values:
//
//    - ok: FALSE if it failed to be presented, otherwise TRUE.
//
func (popup *Popup) Present(width, height int, layout *PopupLayout) bool {
	var _arg0 *C.GdkPopup       // out
	var _arg1 C.int             // out
	var _arg2 C.int             // out
	var _arg3 *C.GdkPopupLayout // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(popup.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)
	_arg3 = (*C.GdkPopupLayout)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = C.gdk_popup_present(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(popup)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(layout)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
