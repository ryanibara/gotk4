// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_scroll_type_get_type()), F: marshalScrollType},
		{T: externglib.Type(C.atk_component_get_type()), F: marshalComponent},
		{T: externglib.Type(C.atk_rectangle_get_type()), F: marshalRectangle},
	})
}

// ScrollType specifies where an object should be placed on the screen when
// using scroll_to.
type ScrollType int

const (
	// TopLeft: scroll the object vertically and horizontally to bring its top
	// left corner to the top left corner of the window.
	ScrollTypeTopLeft ScrollType = 0
	// BottomRight: scroll the object vertically and horizontally to bring its
	// bottom right corner to the bottom right corner of the window.
	ScrollTypeBottomRight ScrollType = 1
	// TopEdge: scroll the object vertically to bring its top edge to the top
	// edge of the window.
	ScrollTypeTopEdge ScrollType = 2
	// BottomEdge: scroll the object vertically to bring its bottom edge to the
	// bottom edge of the window.
	ScrollTypeBottomEdge ScrollType = 3
	// LeftEdge: scroll the object vertically and horizontally to bring its left
	// edge to the left edge of the window.
	ScrollTypeLeftEdge ScrollType = 4
	// RightEdge: scroll the object vertically and horizontally to bring its
	// right edge to the right edge of the window.
	ScrollTypeRightEdge ScrollType = 5
	// anywhere: scroll the object vertically and horizontally so that as much
	// as possible of the object becomes visible. The exact placement is
	// determined by the application.
	ScrollTypeAnywhere ScrollType = 6
)

func marshalScrollType(p uintptr) (interface{}, error) {
	return ScrollType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Component: Component should be implemented by most if not all UI elements
// with an actual on-screen presence, i.e. components which can be said to have
// a screen-coordinate bounding box. Virtually all widgets will need to have
// Component implementations provided for their corresponding Object class. In
// short, only UI elements which are *not* GUI elements will omit this ATK
// interface.
//
// A possible exception might be textual information with a transparent
// background, in which case text glyph bounding box information is provided by
// Text.
type Component interface {
	gextras.Objector

	// Contains: set the size of the @component in terms of width and height.
	Contains(x int, y int, coordType CoordType) bool
	// Alpha: set the size of the @component in terms of width and height.
	Alpha() float64
	// Extents: set the size of the @component in terms of width and height.
	Extents(coordType CoordType) (x int, y int, width int, height int)
	// Layer: set the size of the @component in terms of width and height.
	Layer() Layer
	// MdiZorder: set the size of the @component in terms of width and height.
	MdiZorder() int
	// Position: set the size of the @component in terms of width and height.
	Position(coordType CoordType) (x int, y int)
	// Size: set the size of the @component in terms of width and height.
	Size() (width int, height int)
	// GrabFocus: set the size of the @component in terms of width and height.
	GrabFocus() bool
	// RefAccessibleAtPoint: set the size of the @component in terms of width
	// and height.
	RefAccessibleAtPoint(x int, y int, coordType CoordType) Object
	// RemoveFocusHandler: set the size of the @component in terms of width and
	// height.
	RemoveFocusHandler(handlerId uint)
	// ScrollTo: set the size of the @component in terms of width and height.
	ScrollTo(typ ScrollType) bool
	// ScrollToPoint: set the size of the @component in terms of width and
	// height.
	ScrollToPoint(coords CoordType, x int, y int) bool
	// SetExtents: set the size of the @component in terms of width and height.
	SetExtents(x int, y int, width int, height int, coordType CoordType) bool
	// SetPosition: set the size of the @component in terms of width and height.
	SetPosition(x int, y int, coordType CoordType) bool
	// SetSize: set the size of the @component in terms of width and height.
	SetSize(width int, height int) bool
}

// component implements the Component interface.
type component struct {
	gextras.Objector
}

var _ Component = (*component)(nil)

// WrapComponent wraps a GObject to a type that implements
// interface Component. It is primarily used internally.
func WrapComponent(obj *externglib.Object) Component {
	return component{
		Objector: obj,
	}
}

func marshalComponent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapComponent(obj), nil
}

func (c component) Contains(x int, y int, coordType CoordType) bool {
	var _arg0 *C.AtkComponent // out
	var _arg1 C.gint          // out
	var _arg2 C.gint          // out
	var _arg3 C.AtkCoordType  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)
	_arg3 = C.AtkCoordType(coordType)

	_cret = C.atk_component_contains(_arg0, _arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c component) Alpha() float64 {
	var _arg0 *C.AtkComponent // out
	var _cret C.gdouble       // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(c.Native()))

	_cret = C.atk_component_get_alpha(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (c component) Extents(coordType CoordType) (x int, y int, width int, height int) {
	var _arg0 *C.AtkComponent // out
	var _arg1 C.gint          // in
	var _arg2 C.gint          // in
	var _arg3 C.gint          // in
	var _arg4 C.gint          // in
	var _arg5 C.AtkCoordType  // out

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(c.Native()))
	_arg5 = C.AtkCoordType(coordType)

	C.atk_component_get_extents(_arg0, &_arg1, &_arg2, &_arg3, &_arg4, _arg5)

	var _x int      // out
	var _y int      // out
	var _width int  // out
	var _height int // out

	_x = int(_arg1)
	_y = int(_arg2)
	_width = int(_arg3)
	_height = int(_arg4)

	return _x, _y, _width, _height
}

func (c component) Layer() Layer {
	var _arg0 *C.AtkComponent // out
	var _cret C.AtkLayer      // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(c.Native()))

	_cret = C.atk_component_get_layer(_arg0)

	var _layer Layer // out

	_layer = Layer(_cret)

	return _layer
}

func (c component) MdiZorder() int {
	var _arg0 *C.AtkComponent // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(c.Native()))

	_cret = C.atk_component_get_mdi_zorder(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (c component) Position(coordType CoordType) (x int, y int) {
	var _arg0 *C.AtkComponent // out
	var _arg1 C.gint          // in
	var _arg2 C.gint          // in
	var _arg3 C.AtkCoordType  // out

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(c.Native()))
	_arg3 = C.AtkCoordType(coordType)

	C.atk_component_get_position(_arg0, &_arg1, &_arg2, _arg3)

	var _x int // out
	var _y int // out

	_x = int(_arg1)
	_y = int(_arg2)

	return _x, _y
}

func (c component) Size() (width int, height int) {
	var _arg0 *C.AtkComponent // out
	var _arg1 C.gint          // in
	var _arg2 C.gint          // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(c.Native()))

	C.atk_component_get_size(_arg0, &_arg1, &_arg2)

	var _width int  // out
	var _height int // out

	_width = int(_arg1)
	_height = int(_arg2)

	return _width, _height
}

func (c component) GrabFocus() bool {
	var _arg0 *C.AtkComponent // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(c.Native()))

	_cret = C.atk_component_grab_focus(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c component) RefAccessibleAtPoint(x int, y int, coordType CoordType) Object {
	var _arg0 *C.AtkComponent // out
	var _arg1 C.gint          // out
	var _arg2 C.gint          // out
	var _arg3 C.AtkCoordType  // out
	var _cret *C.AtkObject    // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)
	_arg3 = C.AtkCoordType(coordType)

	_cret = C.atk_component_ref_accessible_at_point(_arg0, _arg1, _arg2, _arg3)

	var _object Object // out

	_object = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Object)

	return _object
}

func (c component) RemoveFocusHandler(handlerId uint) {
	var _arg0 *C.AtkComponent // out
	var _arg1 C.guint         // out

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(c.Native()))
	_arg1 = C.guint(handlerId)

	C.atk_component_remove_focus_handler(_arg0, _arg1)
}

func (c component) ScrollTo(typ ScrollType) bool {
	var _arg0 *C.AtkComponent // out
	var _arg1 C.AtkScrollType // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(c.Native()))
	_arg1 = C.AtkScrollType(typ)

	_cret = C.atk_component_scroll_to(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c component) ScrollToPoint(coords CoordType, x int, y int) bool {
	var _arg0 *C.AtkComponent // out
	var _arg1 C.AtkCoordType  // out
	var _arg2 C.gint          // out
	var _arg3 C.gint          // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(c.Native()))
	_arg1 = C.AtkCoordType(coords)
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	_cret = C.atk_component_scroll_to_point(_arg0, _arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c component) SetExtents(x int, y int, width int, height int, coordType CoordType) bool {
	var _arg0 *C.AtkComponent // out
	var _arg1 C.gint          // out
	var _arg2 C.gint          // out
	var _arg3 C.gint          // out
	var _arg4 C.gint          // out
	var _arg5 C.AtkCoordType  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)
	_arg3 = C.gint(width)
	_arg4 = C.gint(height)
	_arg5 = C.AtkCoordType(coordType)

	_cret = C.atk_component_set_extents(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c component) SetPosition(x int, y int, coordType CoordType) bool {
	var _arg0 *C.AtkComponent // out
	var _arg1 C.gint          // out
	var _arg2 C.gint          // out
	var _arg3 C.AtkCoordType  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)
	_arg3 = C.AtkCoordType(coordType)

	_cret = C.atk_component_set_position(_arg0, _arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c component) SetSize(width int, height int) bool {
	var _arg0 *C.AtkComponent // out
	var _arg1 C.gint          // out
	var _arg2 C.gint          // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkComponent)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(width)
	_arg2 = C.gint(height)

	_cret = C.atk_component_set_size(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Rectangle: a data structure for holding a rectangle. Those coordinates are
// relative to the component top-level parent.
type Rectangle C.AtkRectangle

// WrapRectangle wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRectangle(ptr unsafe.Pointer) *Rectangle {
	return (*Rectangle)(ptr)
}

func marshalRectangle(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Rectangle)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *Rectangle) Native() unsafe.Pointer {
	return unsafe.Pointer(r)
}
