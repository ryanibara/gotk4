// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_popup_get_type()), F: marshalPopuper},
	})
}

// Popuper describes Popup's methods.
type Popuper interface {
	// Autohide returns whether this popup is set to hide on outside clicks.
	Autohide() bool
	// Parent returns the parent surface of a popup.
	Parent() *Surface
	// PositionX obtains the position of the popup relative to its parent.
	PositionX() int
	// PositionY obtains the position of the popup relative to its parent.
	PositionY() int
	// RectAnchor gets the current popup rectangle anchor.
	RectAnchor() Gravity
	// SurfaceAnchor gets the current popup surface anchor.
	SurfaceAnchor() Gravity
	// Present popup after having processed the PopupLayout rules.
	Present(width int, height int, layout *PopupLayout) bool
}

// Popup: GdkPopup is a surface that is attached to another surface.
//
// The GdkPopup is positioned relative to its parent surface.
//
// GdkPopups are typically used to implement menus and similar popups. They can
// be modal, which is indicated by the gdkpopup:autohide property.
type Popup struct {
	Surface
}

var (
	_ Popuper         = (*Popup)(nil)
	_ gextras.Nativer = (*Popup)(nil)
)

func wrapPopup(obj *externglib.Object) *Popup {
	return &Popup{
		Surface: Surface{
			Object: obj,
		},
	}
}

func marshalPopuper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPopup(obj), nil
}

// Autohide returns whether this popup is set to hide on outside clicks.
func (popup *Popup) Autohide() bool {
	var _arg0 *C.GdkPopup // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(popup.Native()))

	_cret = C.gdk_popup_get_autohide(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Parent returns the parent surface of a popup.
func (popup *Popup) Parent() *Surface {
	var _arg0 *C.GdkPopup   // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(popup.Native()))

	_cret = C.gdk_popup_get_parent(_arg0)

	var _surface *Surface // out

	_surface = wrapSurface(externglib.Take(unsafe.Pointer(_cret)))

	return _surface
}

// PositionX obtains the position of the popup relative to its parent.
func (popup *Popup) PositionX() int {
	var _arg0 *C.GdkPopup // out
	var _cret C.int       // in

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(popup.Native()))

	_cret = C.gdk_popup_get_position_x(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PositionY obtains the position of the popup relative to its parent.
func (popup *Popup) PositionY() int {
	var _arg0 *C.GdkPopup // out
	var _cret C.int       // in

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(popup.Native()))

	_cret = C.gdk_popup_get_position_y(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// RectAnchor gets the current popup rectangle anchor.
//
// The value returned may change after calling gdk.Popup.Present(), or after the
// gdk.Surface::layout signal is emitted.
func (popup *Popup) RectAnchor() Gravity {
	var _arg0 *C.GdkPopup  // out
	var _cret C.GdkGravity // in

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(popup.Native()))

	_cret = C.gdk_popup_get_rect_anchor(_arg0)

	var _gravity Gravity // out

	_gravity = Gravity(_cret)

	return _gravity
}

// SurfaceAnchor gets the current popup surface anchor.
//
// The value returned may change after calling gdk.Popup.Present(), or after the
// gdk.Surface::layout signal is emitted.
func (popup *Popup) SurfaceAnchor() Gravity {
	var _arg0 *C.GdkPopup  // out
	var _cret C.GdkGravity // in

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(popup.Native()))

	_cret = C.gdk_popup_get_surface_anchor(_arg0)

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
func (popup *Popup) Present(width int, height int, layout *PopupLayout) bool {
	var _arg0 *C.GdkPopup       // out
	var _arg1 C.int             // out
	var _arg2 C.int             // out
	var _arg3 *C.GdkPopupLayout // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkPopup)(unsafe.Pointer(popup.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)
	_arg3 = (*C.GdkPopupLayout)(unsafe.Pointer(layout))

	_cret = C.gdk_popup_present(_arg0, _arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
