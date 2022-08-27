// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"fmt"
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
// gboolean _gotk4_atk1_Component_virtual_scroll_to(void* fnptr, AtkComponent* arg0, AtkScrollType arg1) {
//   return ((gboolean (*)(AtkComponent*, AtkScrollType))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_atk1_Component_virtual_scroll_to_point(void* fnptr, AtkComponent* arg0, AtkCoordType arg1, gint arg2, gint arg3) {
//   return ((gboolean (*)(AtkComponent*, AtkCoordType, gint, gint))(fnptr))(arg0, arg1, arg2, arg3);
// };
import "C"

// GType values.
var (
	GTypeScrollType = coreglib.Type(C.atk_scroll_type_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeScrollType, F: marshalScrollType},
	})
}

// ScrollType specifies where an object should be placed on the screen when
// using scroll_to.
type ScrollType C.gint

const (
	// ScrollTopLeft: scroll the object vertically and horizontally to bring its
	// top left corner to the top left corner of the window.
	ScrollTopLeft ScrollType = iota
	// ScrollBottomRight: scroll the object vertically and horizontally to bring
	// its bottom right corner to the bottom right corner of the window.
	ScrollBottomRight
	// ScrollTopEdge: scroll the object vertically to bring its top edge to the
	// top edge of the window.
	ScrollTopEdge
	// ScrollBottomEdge: scroll the object vertically to bring its bottom edge
	// to the bottom edge of the window.
	ScrollBottomEdge
	// ScrollLeftEdge: scroll the object vertically and horizontally to bring
	// its left edge to the left edge of the window.
	ScrollLeftEdge
	// ScrollRightEdge: scroll the object vertically and horizontally to bring
	// its right edge to the right edge of the window.
	ScrollRightEdge
	// ScrollAnywhere: scroll the object vertically and horizontally so that as
	// much as possible of the object becomes visible. The exact placement is
	// determined by the application.
	ScrollAnywhere
)

func marshalScrollType(p uintptr) (interface{}, error) {
	return ScrollType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ScrollType.
func (s ScrollType) String() string {
	switch s {
	case ScrollTopLeft:
		return "TopLeft"
	case ScrollBottomRight:
		return "BottomRight"
	case ScrollTopEdge:
		return "TopEdge"
	case ScrollBottomEdge:
		return "BottomEdge"
	case ScrollLeftEdge:
		return "LeftEdge"
	case ScrollRightEdge:
		return "RightEdge"
	case ScrollAnywhere:
		return "Anywhere"
	default:
		return fmt.Sprintf("ScrollType(%d)", s)
	}
}

// ScrollTo makes component visible on the screen by scrolling all necessary
// parents.
//
// Contrary to atk_component_set_position, this does not actually move component
// in its parent, this only makes the parents scroll so that the object shows up
// on the screen, given its current position within the parents.
//
// The function takes the following parameters:
//
//    - typ: specify where the object should be made visible.
//
// The function returns the following values:
//
//    - ok: whether scrolling was successful.
//
func (component *Component) ScrollTo(typ ScrollType) bool {
	var _arg0 *C.AtkComponent // out
	var _arg1 C.AtkScrollType // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(coreglib.InternObject(component).Native()))
	_arg1 = C.AtkScrollType(typ)

	_cret = C.atk_component_scroll_to(_arg0, _arg1)
	runtime.KeepAlive(component)
	runtime.KeepAlive(typ)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ScrollToPoint: move the top-left of component to a given position of the
// screen by scrolling all necessary parents.
//
// The function takes the following parameters:
//
//    - coords: specify whether coordinates are relative to the screen or to the
//      parent object.
//    - x: x-position where to scroll to.
//    - y: y-position where to scroll to.
//
// The function returns the following values:
//
//    - ok: whether scrolling was successful.
//
func (component *Component) ScrollToPoint(coords CoordType, x, y int) bool {
	var _arg0 *C.AtkComponent // out
	var _arg1 C.AtkCoordType  // out
	var _arg2 C.gint          // out
	var _arg3 C.gint          // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(coreglib.InternObject(component).Native()))
	_arg1 = C.AtkCoordType(coords)
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	_cret = C.atk_component_scroll_to_point(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(component)
	runtime.KeepAlive(coords)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// scrollTo makes component visible on the screen by scrolling all necessary
// parents.
//
// Contrary to atk_component_set_position, this does not actually move component
// in its parent, this only makes the parents scroll so that the object shows up
// on the screen, given its current position within the parents.
//
// The function takes the following parameters:
//
//    - typ: specify where the object should be made visible.
//
// The function returns the following values:
//
//    - ok: whether scrolling was successful.
//
func (component *Component) scrollTo(typ ScrollType) bool {
	gclass := (*C.AtkComponentIface)(coreglib.PeekParentClass(component))
	fnarg := gclass.scroll_to

	var _arg0 *C.AtkComponent // out
	var _arg1 C.AtkScrollType // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(coreglib.InternObject(component).Native()))
	_arg1 = C.AtkScrollType(typ)

	_cret = C._gotk4_atk1_Component_virtual_scroll_to(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(component)
	runtime.KeepAlive(typ)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// scrollToPoint: move the top-left of component to a given position of the
// screen by scrolling all necessary parents.
//
// The function takes the following parameters:
//
//    - coords: specify whether coordinates are relative to the screen or to the
//      parent object.
//    - x: x-position where to scroll to.
//    - y: y-position where to scroll to.
//
// The function returns the following values:
//
//    - ok: whether scrolling was successful.
//
func (component *Component) scrollToPoint(coords CoordType, x, y int) bool {
	gclass := (*C.AtkComponentIface)(coreglib.PeekParentClass(component))
	fnarg := gclass.scroll_to_point

	var _arg0 *C.AtkComponent // out
	var _arg1 C.AtkCoordType  // out
	var _arg2 C.gint          // out
	var _arg3 C.gint          // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(coreglib.InternObject(component).Native()))
	_arg1 = C.AtkCoordType(coords)
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	_cret = C._gotk4_atk1_Component_virtual_scroll_to_point(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(component)
	runtime.KeepAlive(coords)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
