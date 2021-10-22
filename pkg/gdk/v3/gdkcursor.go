// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_cursor_type_get_type()), F: marshalCursorType},
		{T: externglib.Type(C.gdk_cursor_get_type()), F: marshalCursorrer},
	})
}

// CursorType: predefined cursors.
//
// Note that these IDs are directly taken from the X cursor font, and many of
// these cursors are either not useful, or are not available on other platforms.
//
// The recommended way to create cursors is to use gdk_cursor_new_from_name().
type CursorType int

const (
	// XCursor: ! (X_cursor.png).
	XCursor CursorType = 0
	// Arrow: ! (arrow.png).
	Arrow CursorType = 2
	// BasedArrowDown: ! (based_arrow_down.png).
	BasedArrowDown CursorType = 4
	// BasedArrowUp: ! (based_arrow_up.png).
	BasedArrowUp CursorType = 6
	// Boat: ! (boat.png).
	Boat CursorType = 8
	// Bogosity: ! (bogosity.png).
	Bogosity CursorType = 10
	// BottomLeftCorner: ! (bottom_left_corner.png).
	BottomLeftCorner CursorType = 12
	// BottomRightCorner: ! (bottom_right_corner.png).
	BottomRightCorner CursorType = 14
	// BottomSide: ! (bottom_side.png).
	BottomSide CursorType = 16
	// BottomTee: ! (bottom_tee.png).
	BottomTee CursorType = 18
	// BoxSpiral: ! (box_spiral.png).
	BoxSpiral CursorType = 20
	// CenterPtr: ! (center_ptr.png).
	CenterPtr CursorType = 22
	// Circle: ! (circle.png).
	Circle CursorType = 24
	// Clock: ! (clock.png).
	Clock CursorType = 26
	// CoffeeMug: ! (coffee_mug.png).
	CoffeeMug CursorType = 28
	// Cross: ! (cross.png).
	Cross CursorType = 30
	// CrossReverse: ! (cross_reverse.png).
	CrossReverse CursorType = 32
	// Crosshair: ! (crosshair.png).
	Crosshair CursorType = 34
	// DiamondCross: ! (diamond_cross.png).
	DiamondCross CursorType = 36
	// Dot: ! (dot.png).
	Dot CursorType = 38
	// Dotbox: ! (dotbox.png).
	Dotbox CursorType = 40
	// DoubleArrow: ! (double_arrow.png).
	DoubleArrow CursorType = 42
	// DraftLarge: ! (draft_large.png).
	DraftLarge CursorType = 44
	// DraftSmall: ! (draft_small.png).
	DraftSmall CursorType = 46
	// DrapedBox: ! (draped_box.png).
	DrapedBox CursorType = 48
	// Exchange: ! (exchange.png).
	Exchange CursorType = 50
	// Fleur: ! (fleur.png).
	Fleur CursorType = 52
	// Gobbler: ! (gobbler.png).
	Gobbler CursorType = 54
	// Gumby: ! (gumby.png).
	Gumby CursorType = 56
	// Hand1: ! (hand1.png).
	Hand1 CursorType = 58
	// Hand2: ! (hand2.png).
	Hand2 CursorType = 60
	// Heart: ! (heart.png).
	Heart CursorType = 62
	// Icon: ! (icon.png).
	Icon CursorType = 64
	// IronCross: ! (iron_cross.png).
	IronCross CursorType = 66
	// LeftPtr: ! (left_ptr.png).
	LeftPtr CursorType = 68
	// LeftSide: ! (left_side.png).
	LeftSide CursorType = 70
	// LeftTee: ! (left_tee.png).
	LeftTee CursorType = 72
	// Leftbutton: ! (leftbutton.png).
	Leftbutton CursorType = 74
	// LlAngle: ! (ll_angle.png).
	LlAngle CursorType = 76
	// LrAngle: ! (lr_angle.png).
	LrAngle CursorType = 78
	// Man: ! (man.png).
	Man CursorType = 80
	// Middlebutton: ! (middlebutton.png).
	Middlebutton CursorType = 82
	// Mouse: ! (mouse.png).
	Mouse CursorType = 84
	// Pencil: ! (pencil.png).
	Pencil CursorType = 86
	// Pirate: ! (pirate.png).
	Pirate CursorType = 88
	// Plus: ! (plus.png).
	Plus CursorType = 90
	// QuestionArrow: ! (question_arrow.png).
	QuestionArrow CursorType = 92
	// RightPtr: ! (right_ptr.png).
	RightPtr CursorType = 94
	// RightSide: ! (right_side.png).
	RightSide CursorType = 96
	// RightTee: ! (right_tee.png).
	RightTee CursorType = 98
	// Rightbutton: ! (rightbutton.png).
	Rightbutton CursorType = 100
	// RTLLogo: ! (rtl_logo.png).
	RTLLogo CursorType = 102
	// Sailboat: ! (sailboat.png).
	Sailboat CursorType = 104
	// SbDownArrow: ! (sb_down_arrow.png).
	SbDownArrow CursorType = 106
	// SbHDoubleArrow: ! (sb_h_double_arrow.png).
	SbHDoubleArrow CursorType = 108
	// SbLeftArrow: ! (sb_left_arrow.png).
	SbLeftArrow CursorType = 110
	// SbRightArrow: ! (sb_right_arrow.png).
	SbRightArrow CursorType = 112
	// SbUpArrow: ! (sb_up_arrow.png).
	SbUpArrow CursorType = 114
	// SbVDoubleArrow: ! (sb_v_double_arrow.png).
	SbVDoubleArrow CursorType = 116
	// Shuttle: ! (shuttle.png).
	Shuttle CursorType = 118
	// Sizing: ! (sizing.png).
	Sizing CursorType = 120
	// Spider: ! (spider.png).
	Spider CursorType = 122
	// Spraycan: ! (spraycan.png).
	Spraycan CursorType = 124
	// Star: ! (star.png).
	Star CursorType = 126
	// Target: ! (target.png).
	Target CursorType = 128
	// Tcross: ! (tcross.png).
	Tcross CursorType = 130
	// TopLeftArrow: ! (top_left_arrow.png).
	TopLeftArrow CursorType = 132
	// TopLeftCorner: ! (top_left_corner.png).
	TopLeftCorner CursorType = 134
	// TopRightCorner: ! (top_right_corner.png).
	TopRightCorner CursorType = 136
	// TopSide: ! (top_side.png).
	TopSide CursorType = 138
	// TopTee: ! (top_tee.png).
	TopTee CursorType = 140
	// Trek: ! (trek.png).
	Trek CursorType = 142
	// UlAngle: ! (ul_angle.png).
	UlAngle CursorType = 144
	// Umbrella: ! (umbrella.png).
	Umbrella CursorType = 146
	// UrAngle: ! (ur_angle.png).
	UrAngle CursorType = 148
	// Watch: ! (watch.png).
	Watch CursorType = 150
	// Xterm: ! (xterm.png).
	Xterm CursorType = 152
	// LastCursor: last cursor type.
	LastCursor CursorType = 153
	// BlankCursor: blank cursor. Since 2.16.
	BlankCursor CursorType = -2
	// CursorIsPixmap: type of cursors constructed with
	// gdk_cursor_new_from_pixbuf().
	CursorIsPixmap CursorType = -1
)

func marshalCursorType(p uintptr) (interface{}, error) {
	return CursorType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CursorType.
func (c CursorType) String() string {
	switch c {
	case XCursor:
		return "XCursor"
	case Arrow:
		return "Arrow"
	case BasedArrowDown:
		return "BasedArrowDown"
	case BasedArrowUp:
		return "BasedArrowUp"
	case Boat:
		return "Boat"
	case Bogosity:
		return "Bogosity"
	case BottomLeftCorner:
		return "BottomLeftCorner"
	case BottomRightCorner:
		return "BottomRightCorner"
	case BottomSide:
		return "BottomSide"
	case BottomTee:
		return "BottomTee"
	case BoxSpiral:
		return "BoxSpiral"
	case CenterPtr:
		return "CenterPtr"
	case Circle:
		return "Circle"
	case Clock:
		return "Clock"
	case CoffeeMug:
		return "CoffeeMug"
	case Cross:
		return "Cross"
	case CrossReverse:
		return "CrossReverse"
	case Crosshair:
		return "Crosshair"
	case DiamondCross:
		return "DiamondCross"
	case Dot:
		return "Dot"
	case Dotbox:
		return "Dotbox"
	case DoubleArrow:
		return "DoubleArrow"
	case DraftLarge:
		return "DraftLarge"
	case DraftSmall:
		return "DraftSmall"
	case DrapedBox:
		return "DrapedBox"
	case Exchange:
		return "Exchange"
	case Fleur:
		return "Fleur"
	case Gobbler:
		return "Gobbler"
	case Gumby:
		return "Gumby"
	case Hand1:
		return "Hand1"
	case Hand2:
		return "Hand2"
	case Heart:
		return "Heart"
	case Icon:
		return "Icon"
	case IronCross:
		return "IronCross"
	case LeftPtr:
		return "LeftPtr"
	case LeftSide:
		return "LeftSide"
	case LeftTee:
		return "LeftTee"
	case Leftbutton:
		return "Leftbutton"
	case LlAngle:
		return "LlAngle"
	case LrAngle:
		return "LrAngle"
	case Man:
		return "Man"
	case Middlebutton:
		return "Middlebutton"
	case Mouse:
		return "Mouse"
	case Pencil:
		return "Pencil"
	case Pirate:
		return "Pirate"
	case Plus:
		return "Plus"
	case QuestionArrow:
		return "QuestionArrow"
	case RightPtr:
		return "RightPtr"
	case RightSide:
		return "RightSide"
	case RightTee:
		return "RightTee"
	case Rightbutton:
		return "Rightbutton"
	case RTLLogo:
		return "RTLLogo"
	case Sailboat:
		return "Sailboat"
	case SbDownArrow:
		return "SbDownArrow"
	case SbHDoubleArrow:
		return "SbHDoubleArrow"
	case SbLeftArrow:
		return "SbLeftArrow"
	case SbRightArrow:
		return "SbRightArrow"
	case SbUpArrow:
		return "SbUpArrow"
	case SbVDoubleArrow:
		return "SbVDoubleArrow"
	case Shuttle:
		return "Shuttle"
	case Sizing:
		return "Sizing"
	case Spider:
		return "Spider"
	case Spraycan:
		return "Spraycan"
	case Star:
		return "Star"
	case Target:
		return "Target"
	case Tcross:
		return "Tcross"
	case TopLeftArrow:
		return "TopLeftArrow"
	case TopLeftCorner:
		return "TopLeftCorner"
	case TopRightCorner:
		return "TopRightCorner"
	case TopSide:
		return "TopSide"
	case TopTee:
		return "TopTee"
	case Trek:
		return "Trek"
	case UlAngle:
		return "UlAngle"
	case Umbrella:
		return "Umbrella"
	case UrAngle:
		return "UrAngle"
	case Watch:
		return "Watch"
	case Xterm:
		return "Xterm"
	case LastCursor:
		return "LastCursor"
	case BlankCursor:
		return "BlankCursor"
	case CursorIsPixmap:
		return "CursorIsPixmap"
	default:
		return fmt.Sprintf("CursorType(%d)", c)
	}
}

// Cursor represents a cursor. Its contents are private.
type Cursor struct {
	*externglib.Object
}

// Cursorrer describes types inherited from class Cursor.
// To get the original type, the caller must assert this to an interface or
// another type.
type Cursorrer interface {
	externglib.Objector

	// BaseCursor returns the underlying base class.
	BaseCursor() *Cursor
}

var _ Cursorrer = (*Cursor)(nil)

func wrapCursor(obj *externglib.Object) *Cursor {
	return &Cursor{
		Object: obj,
	}
}

func marshalCursorrer(p uintptr) (interface{}, error) {
	return wrapCursor(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCursor creates a new cursor from the set of builtin cursors for the
// default display. See gdk_cursor_new_for_display().
//
// To make the cursor invisible, use GDK_BLANK_CURSOR.
//
// Deprecated: Use gdk_cursor_new_for_display() instead.
//
// The function takes the following parameters:
//
//    - cursorType: cursor to create.
//
func NewCursor(cursorType CursorType) *Cursor {
	var _arg1 C.GdkCursorType // out
	var _cret *C.GdkCursor    // in

	_arg1 = C.GdkCursorType(cursorType)

	_cret = C.gdk_cursor_new(_arg1)
	runtime.KeepAlive(cursorType)

	var _cursor *Cursor // out

	_cursor = wrapCursor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cursor
}

// NewCursorForDisplay creates a new cursor from the set of builtin cursors.
//
// The function takes the following parameters:
//
//    - display for which the cursor will be created.
//    - cursorType: cursor to create.
//
func NewCursorForDisplay(display *Display, cursorType CursorType) *Cursor {
	var _arg1 *C.GdkDisplay   // out
	var _arg2 C.GdkCursorType // out
	var _cret *C.GdkCursor    // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = C.GdkCursorType(cursorType)

	_cret = C.gdk_cursor_new_for_display(_arg1, _arg2)
	runtime.KeepAlive(display)
	runtime.KeepAlive(cursorType)

	var _cursor *Cursor // out

	if _cret != nil {
		_cursor = wrapCursor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _cursor
}

// NewCursorFromName creates a new cursor by looking up name in the current
// cursor theme.
//
// A recommended set of cursor names that will work across different platforms
// can be found in the CSS specification:
//
// - "none"
//
// - ! (default_cursor.png) "default"
//
// - ! (help_cursor.png) "help"
//
// - ! (pointer_cursor.png) "pointer"
//
// - ! (context_menu_cursor.png) "context-menu"
//
// - ! (progress_cursor.png) "progress"
//
// - ! (wait_cursor.png) "wait"
//
// - ! (cell_cursor.png) "cell"
//
// - ! (crosshair_cursor.png) "crosshair"
//
// - ! (text_cursor.png) "text"
//
// - ! (vertical_text_cursor.png) "vertical-text"
//
// - ! (alias_cursor.png) "alias"
//
// - ! (copy_cursor.png) "copy"
//
// - ! (no_drop_cursor.png) "no-drop"
//
// - ! (move_cursor.png) "move"
//
// - ! (not_allowed_cursor.png) "not-allowed"
//
// - ! (grab_cursor.png) "grab"
//
// - ! (grabbing_cursor.png) "grabbing"
//
// - ! (all_scroll_cursor.png) "all-scroll"
//
// - ! (col_resize_cursor.png) "col-resize"
//
// - ! (row_resize_cursor.png) "row-resize"
//
// - ! (n_resize_cursor.png) "n-resize"
//
// - ! (e_resize_cursor.png) "e-resize"
//
// - ! (s_resize_cursor.png) "s-resize"
//
// - ! (w_resize_cursor.png) "w-resize"
//
// - ! (ne_resize_cursor.png) "ne-resize"
//
// - ! (nw_resize_cursor.png) "nw-resize"
//
// - ! (sw_resize_cursor.png) "sw-resize"
//
// - ! (se_resize_cursor.png) "se-resize"
//
// - ! (ew_resize_cursor.png) "ew-resize"
//
// - ! (ns_resize_cursor.png) "ns-resize"
//
// - ! (nesw_resize_cursor.png) "nesw-resize"
//
// - ! (nwse_resize_cursor.png) "nwse-resize"
//
// - ! (zoom_in_cursor.png) "zoom-in"
//
// - ! (zoom_out_cursor.png) "zoom-out".
//
// The function takes the following parameters:
//
//    - display for which the cursor will be created.
//    - name of the cursor.
//
func NewCursorFromName(display *Display, name string) *Cursor {
	var _arg1 *C.GdkDisplay // out
	var _arg2 *C.gchar      // out
	var _cret *C.GdkCursor  // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gdk_cursor_new_from_name(_arg1, _arg2)
	runtime.KeepAlive(display)
	runtime.KeepAlive(name)

	var _cursor *Cursor // out

	if _cret != nil {
		_cursor = wrapCursor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _cursor
}

// NewCursorFromPixbuf creates a new cursor from a pixbuf.
//
// Not all GDK backends support RGBA cursors. If they are not supported, a
// monochrome approximation will be displayed. The functions
// gdk_display_supports_cursor_alpha() and gdk_display_supports_cursor_color()
// can be used to determine whether RGBA cursors are supported;
// gdk_display_get_default_cursor_size() and
// gdk_display_get_maximal_cursor_size() give information about cursor sizes.
//
// If x or y are -1, the pixbuf must have options named “x_hot” and “y_hot”,
// resp., containing integer values between 0 and the width resp. height of the
// pixbuf. (Since: 3.0)
//
// On the X backend, support for RGBA cursors requires a sufficently new version
// of the X Render extension.
//
// The function takes the following parameters:
//
//    - display for which the cursor will be created.
//    - pixbuf containing the cursor image.
//    - x: horizontal offset of the “hotspot” of the cursor.
//    - y: vertical offset of the “hotspot” of the cursor.
//
func NewCursorFromPixbuf(display *Display, pixbuf *gdkpixbuf.Pixbuf, x, y int) *Cursor {
	var _arg1 *C.GdkDisplay // out
	var _arg2 *C.GdkPixbuf  // out
	var _arg3 C.gint        // out
	var _arg4 C.gint        // out
	var _cret *C.GdkCursor  // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	_arg3 = C.gint(x)
	_arg4 = C.gint(y)

	_cret = C.gdk_cursor_new_from_pixbuf(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(display)
	runtime.KeepAlive(pixbuf)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _cursor *Cursor // out

	_cursor = wrapCursor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cursor
}

// NewCursorFromSurface creates a new cursor from a cairo image surface.
//
// Not all GDK backends support RGBA cursors. If they are not supported, a
// monochrome approximation will be displayed. The functions
// gdk_display_supports_cursor_alpha() and gdk_display_supports_cursor_color()
// can be used to determine whether RGBA cursors are supported;
// gdk_display_get_default_cursor_size() and
// gdk_display_get_maximal_cursor_size() give information about cursor sizes.
//
// On the X backend, support for RGBA cursors requires a sufficently new version
// of the X Render extension.
//
// The function takes the following parameters:
//
//    - display for which the cursor will be created.
//    - surface: cairo image surface containing the cursor pixel data.
//    - x: horizontal offset of the “hotspot” of the cursor.
//    - y: vertical offset of the “hotspot” of the cursor.
//
func NewCursorFromSurface(display *Display, surface *cairo.Surface, x, y float64) *Cursor {
	var _arg1 *C.GdkDisplay      // out
	var _arg2 *C.cairo_surface_t // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _cret *C.GdkCursor       // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)

	_cret = C.gdk_cursor_new_from_surface(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(display)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _cursor *Cursor // out

	_cursor = wrapCursor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cursor
}

// CursorType returns the cursor type for this cursor.
func (cursor *Cursor) CursorType() CursorType {
	var _arg0 *C.GdkCursor    // out
	var _cret C.GdkCursorType // in

	_arg0 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	_cret = C.gdk_cursor_get_cursor_type(_arg0)
	runtime.KeepAlive(cursor)

	var _cursorType CursorType // out

	_cursorType = CursorType(_cret)

	return _cursorType
}

// Display returns the display on which the Cursor is defined.
func (cursor *Cursor) Display() *Display {
	var _arg0 *C.GdkCursor  // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	_cret = C.gdk_cursor_get_display(_arg0)
	runtime.KeepAlive(cursor)

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// Image returns a Pixbuf with the image used to display the cursor.
//
// Note that depending on the capabilities of the windowing system and on the
// cursor, GDK may not be able to obtain the image data. In this case, NULL is
// returned.
func (cursor *Cursor) Image() *gdkpixbuf.Pixbuf {
	var _arg0 *C.GdkCursor // out
	var _cret *C.GdkPixbuf // in

	_arg0 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	_cret = C.gdk_cursor_get_image(_arg0)
	runtime.KeepAlive(cursor)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if _cret != nil {
		{
			obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
			_pixbuf = &gdkpixbuf.Pixbuf{
				Object: obj,
				LoadableIcon: gio.LoadableIcon{
					Icon: gio.Icon{
						Object: obj,
					},
				},
			}
		}
	}

	return _pixbuf
}

// Surface returns a cairo image surface with the image used to display the
// cursor.
//
// Note that depending on the capabilities of the windowing system and on the
// cursor, GDK may not be able to obtain the image data. In this case, NULL is
// returned.
func (cursor *Cursor) Surface() (xHot float64, yHot float64, surface *cairo.Surface) {
	var _arg0 *C.GdkCursor       // out
	var _arg1 C.gdouble          // in
	var _arg2 C.gdouble          // in
	var _cret *C.cairo_surface_t // in

	_arg0 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	_cret = C.gdk_cursor_get_surface(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(cursor)

	var _xHot float64           // out
	var _yHot float64           // out
	var _surface *cairo.Surface // out

	_xHot = float64(_arg1)
	_yHot = float64(_arg2)
	if _cret != nil {
		_surface = cairo.WrapSurface(uintptr(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(_surface, func(v *cairo.Surface) {
			C.cairo_surface_destroy((*C.cairo_surface_t)(unsafe.Pointer(v.Native())))
		})
	}

	return _xHot, _yHot, _surface
}

// BaseCursor returns cursor.
func (cursor *Cursor) BaseCursor() *Cursor {
	return cursor
}
