// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_cursor_type_get_type()), F: marshalCursorType},
		{T: externglib.Type(C.gdk_cursor_get_type()), F: marshalCursor},
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
	// XCursor: ! (X_cursor.png)
	CursorTypeXCursor CursorType = 0
	// arrow: ! (arrow.png)
	CursorTypeArrow CursorType = 2
	// BasedArrowDown: ! (based_arrow_down.png)
	CursorTypeBasedArrowDown CursorType = 4
	// BasedArrowUp: ! (based_arrow_up.png)
	CursorTypeBasedArrowUp CursorType = 6
	// boat: ! (boat.png)
	CursorTypeBoat CursorType = 8
	// bogosity: ! (bogosity.png)
	CursorTypeBogosity CursorType = 10
	// BottomLeftCorner: ! (bottom_left_corner.png)
	CursorTypeBottomLeftCorner CursorType = 12
	// BottomRightCorner: ! (bottom_right_corner.png)
	CursorTypeBottomRightCorner CursorType = 14
	// BottomSide: ! (bottom_side.png)
	CursorTypeBottomSide CursorType = 16
	// BottomTee: ! (bottom_tee.png)
	CursorTypeBottomTee CursorType = 18
	// BoxSpiral: ! (box_spiral.png)
	CursorTypeBoxSpiral CursorType = 20
	// CenterPtr: ! (center_ptr.png)
	CursorTypeCenterPtr CursorType = 22
	// circle: ! (circle.png)
	CursorTypeCircle CursorType = 24
	// clock: ! (clock.png)
	CursorTypeClock CursorType = 26
	// CoffeeMug: ! (coffee_mug.png)
	CursorTypeCoffeeMug CursorType = 28
	// cross: ! (cross.png)
	CursorTypeCross CursorType = 30
	// CrossReverse: ! (cross_reverse.png)
	CursorTypeCrossReverse CursorType = 32
	// crosshair: ! (crosshair.png)
	CursorTypeCrosshair CursorType = 34
	// DiamondCross: ! (diamond_cross.png)
	CursorTypeDiamondCross CursorType = 36
	// dot: ! (dot.png)
	CursorTypeDot CursorType = 38
	// dotbox: ! (dotbox.png)
	CursorTypeDotbox CursorType = 40
	// DoubleArrow: ! (double_arrow.png)
	CursorTypeDoubleArrow CursorType = 42
	// DraftLarge: ! (draft_large.png)
	CursorTypeDraftLarge CursorType = 44
	// DraftSmall: ! (draft_small.png)
	CursorTypeDraftSmall CursorType = 46
	// DrapedBox: ! (draped_box.png)
	CursorTypeDrapedBox CursorType = 48
	// exchange: ! (exchange.png)
	CursorTypeExchange CursorType = 50
	// fleur: ! (fleur.png)
	CursorTypeFleur CursorType = 52
	// gobbler: ! (gobbler.png)
	CursorTypeGobbler CursorType = 54
	// gumby: ! (gumby.png)
	CursorTypeGumby CursorType = 56
	// hand1: ! (hand1.png)
	CursorTypeHand1 CursorType = 58
	// hand2: ! (hand2.png)
	CursorTypeHand2 CursorType = 60
	// heart: ! (heart.png)
	CursorTypeHeart CursorType = 62
	// icon: ! (icon.png)
	CursorTypeIcon CursorType = 64
	// IronCross: ! (iron_cross.png)
	CursorTypeIronCross CursorType = 66
	// LeftPtr: ! (left_ptr.png)
	CursorTypeLeftPtr CursorType = 68
	// LeftSide: ! (left_side.png)
	CursorTypeLeftSide CursorType = 70
	// LeftTee: ! (left_tee.png)
	CursorTypeLeftTee CursorType = 72
	// leftbutton: ! (leftbutton.png)
	CursorTypeLeftbutton CursorType = 74
	// LlAngle: ! (ll_angle.png)
	CursorTypeLlAngle CursorType = 76
	// LrAngle: ! (lr_angle.png)
	CursorTypeLrAngle CursorType = 78
	// man: ! (man.png)
	CursorTypeMan CursorType = 80
	// middlebutton: ! (middlebutton.png)
	CursorTypeMiddlebutton CursorType = 82
	// mouse: ! (mouse.png)
	CursorTypeMouse CursorType = 84
	// pencil: ! (pencil.png)
	CursorTypePencil CursorType = 86
	// pirate: ! (pirate.png)
	CursorTypePirate CursorType = 88
	// plus: ! (plus.png)
	CursorTypePlus CursorType = 90
	// QuestionArrow: ! (question_arrow.png)
	CursorTypeQuestionArrow CursorType = 92
	// RightPtr: ! (right_ptr.png)
	CursorTypeRightPtr CursorType = 94
	// RightSide: ! (right_side.png)
	CursorTypeRightSide CursorType = 96
	// RightTee: ! (right_tee.png)
	CursorTypeRightTee CursorType = 98
	// rightbutton: ! (rightbutton.png)
	CursorTypeRightbutton CursorType = 100
	// RTLLogo: ! (rtl_logo.png)
	CursorTypeRTLLogo CursorType = 102
	// sailboat: ! (sailboat.png)
	CursorTypeSailboat CursorType = 104
	// SbDownArrow: ! (sb_down_arrow.png)
	CursorTypeSbDownArrow CursorType = 106
	// SbHDoubleArrow: ! (sb_h_double_arrow.png)
	CursorTypeSbHDoubleArrow CursorType = 108
	// SbLeftArrow: ! (sb_left_arrow.png)
	CursorTypeSbLeftArrow CursorType = 110
	// SbRightArrow: ! (sb_right_arrow.png)
	CursorTypeSbRightArrow CursorType = 112
	// SbUpArrow: ! (sb_up_arrow.png)
	CursorTypeSbUpArrow CursorType = 114
	// SbVDoubleArrow: ! (sb_v_double_arrow.png)
	CursorTypeSbVDoubleArrow CursorType = 116
	// shuttle: ! (shuttle.png)
	CursorTypeShuttle CursorType = 118
	// sizing: ! (sizing.png)
	CursorTypeSizing CursorType = 120
	// spider: ! (spider.png)
	CursorTypeSpider CursorType = 122
	// spraycan: ! (spraycan.png)
	CursorTypeSpraycan CursorType = 124
	// star: ! (star.png)
	CursorTypeStar CursorType = 126
	// target: ! (target.png)
	CursorTypeTarget CursorType = 128
	// tcross: ! (tcross.png)
	CursorTypeTcross CursorType = 130
	// TopLeftArrow: ! (top_left_arrow.png)
	CursorTypeTopLeftArrow CursorType = 132
	// TopLeftCorner: ! (top_left_corner.png)
	CursorTypeTopLeftCorner CursorType = 134
	// TopRightCorner: ! (top_right_corner.png)
	CursorTypeTopRightCorner CursorType = 136
	// TopSide: ! (top_side.png)
	CursorTypeTopSide CursorType = 138
	// TopTee: ! (top_tee.png)
	CursorTypeTopTee CursorType = 140
	// trek: ! (trek.png)
	CursorTypeTrek CursorType = 142
	// UlAngle: ! (ul_angle.png)
	CursorTypeUlAngle CursorType = 144
	// umbrella: ! (umbrella.png)
	CursorTypeUmbrella CursorType = 146
	// UrAngle: ! (ur_angle.png)
	CursorTypeUrAngle CursorType = 148
	// watch: ! (watch.png)
	CursorTypeWatch CursorType = 150
	// xterm: ! (xterm.png)
	CursorTypeXterm CursorType = 152
	// LastCursor: last cursor type
	CursorTypeLastCursor CursorType = 153
	// BlankCursor: blank cursor. Since 2.16
	CursorTypeBlankCursor CursorType = -2
	// CursorIsPixmap: type of cursors constructed with
	// gdk_cursor_new_from_pixbuf()
	CursorTypeCursorIsPixmap CursorType = -1
)

func marshalCursorType(p uintptr) (interface{}, error) {
	return CursorType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Cursor represents a cursor. Its contents are private.
type Cursor interface {
	gextras.Objector

	// CursorType returns the cursor type for this cursor.
	CursorType() CursorType
	// Display returns the display on which the Cursor is defined.
	Display() Display
	// Image returns a Pixbuf with the image used to display the cursor.
	//
	// Note that depending on the capabilities of the windowing system and on
	// the cursor, GDK may not be able to obtain the image data. In this case,
	// nil is returned.
	Image() gdkpixbuf.Pixbuf
	// Surface returns a cairo image surface with the image used to display the
	// cursor.
	//
	// Note that depending on the capabilities of the windowing system and on
	// the cursor, GDK may not be able to obtain the image data. In this case,
	// nil is returned.
	Surface() (xHot float64, yHot float64, surface cairo.Surface)
	// RefCursor adds a reference to @cursor.
	//
	// Deprecated: since version 3.0.
	RefCursor() Cursor
	// UnrefCursor removes a reference from @cursor, deallocating the cursor if
	// no references remain.
	//
	// Deprecated: since version 3.0.
	UnrefCursor()
}

// cursor implements the Cursor class.
type cursor struct {
	gextras.Objector
}

// WrapCursor wraps a GObject to the right type. It is
// primarily used internally.
func WrapCursor(obj *externglib.Object) Cursor {
	return cursor{
		Objector: obj,
	}
}

func marshalCursor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCursor(obj), nil
}

// NewCursor creates a new cursor from the set of builtin cursors for the
// default display. See gdk_cursor_new_for_display().
//
// To make the cursor invisible, use GDK_BLANK_CURSOR.
//
// Deprecated: since version 3.16.
func NewCursor(cursorType CursorType) Cursor {
	var _arg1 C.GdkCursorType // out
	var _cret *C.GdkCursor    // in

	_arg1 = C.GdkCursorType(cursorType)

	_cret = C.gdk_cursor_new(_arg1)

	var _cursor Cursor // out

	_cursor = WrapCursor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cursor
}

// NewCursorForDisplay creates a new cursor from the set of builtin cursors.
func NewCursorForDisplay(display Display, cursorType CursorType) Cursor {
	var _arg1 *C.GdkDisplay   // out
	var _arg2 C.GdkCursorType // out
	var _cret *C.GdkCursor    // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = C.GdkCursorType(cursorType)

	_cret = C.gdk_cursor_new_for_display(_arg1, _arg2)

	var _cursor Cursor // out

	_cursor = WrapCursor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cursor
}

// NewCursorFromName creates a new cursor by looking up @name in the current
// cursor theme.
//
// A recommended set of cursor names that will work across different platforms
// can be found in the CSS specification: - "none" - ! (default_cursor.png)
// "default" - ! (help_cursor.png) "help" - ! (pointer_cursor.png) "pointer" - !
// (context_menu_cursor.png) "context-menu" - ! (progress_cursor.png) "progress"
// - ! (wait_cursor.png) "wait" - ! (cell_cursor.png) "cell" - !
// (crosshair_cursor.png) "crosshair" - ! (text_cursor.png) "text" - !
// (vertical_text_cursor.png) "vertical-text" - ! (alias_cursor.png) "alias" - !
// (copy_cursor.png) "copy" - ! (no_drop_cursor.png) "no-drop" - !
// (move_cursor.png) "move" - ! (not_allowed_cursor.png) "not-allowed" - !
// (grab_cursor.png) "grab" - ! (grabbing_cursor.png) "grabbing" - !
// (all_scroll_cursor.png) "all-scroll" - ! (col_resize_cursor.png) "col-resize"
// - ! (row_resize_cursor.png) "row-resize" - ! (n_resize_cursor.png) "n-resize"
// - ! (e_resize_cursor.png) "e-resize" - ! (s_resize_cursor.png) "s-resize" - !
// (w_resize_cursor.png) "w-resize" - ! (ne_resize_cursor.png) "ne-resize" - !
// (nw_resize_cursor.png) "nw-resize" - ! (sw_resize_cursor.png) "sw-resize" - !
// (se_resize_cursor.png) "se-resize" - ! (ew_resize_cursor.png) "ew-resize" - !
// (ns_resize_cursor.png) "ns-resize" - ! (nesw_resize_cursor.png) "nesw-resize"
// - ! (nwse_resize_cursor.png) "nwse-resize" - ! (zoom_in_cursor.png) "zoom-in"
// - ! (zoom_out_cursor.png) "zoom-out"
func NewCursorFromName(display Display, name string) Cursor {
	var _arg1 *C.GdkDisplay // out
	var _arg2 *C.gchar      // out
	var _cret *C.GdkCursor  // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gdk_cursor_new_from_name(_arg1, _arg2)

	var _cursor Cursor // out

	_cursor = WrapCursor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

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
// If @x or @y are `-1`, the pixbuf must have options named “x_hot” and “y_hot”,
// resp., containing integer values between `0` and the width resp. height of
// the pixbuf. (Since: 3.0)
//
// On the X backend, support for RGBA cursors requires a sufficently new version
// of the X Render extension.
func NewCursorFromPixbuf(display Display, pixbuf gdkpixbuf.Pixbuf, x int, y int) Cursor {
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

	var _cursor Cursor // out

	_cursor = WrapCursor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

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
func NewCursorFromSurface(display Display, surface cairo.Surface, x float64, y float64) Cursor {
	var _arg1 *C.GdkDisplay      // out
	var _arg2 *C.cairo_surface_t // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _cret *C.GdkCursor       // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (*C.cairo_surface_t)(unsafe.Pointer(surface))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)

	_cret = C.gdk_cursor_new_from_surface(_arg1, _arg2, _arg3, _arg4)

	var _cursor Cursor // out

	_cursor = WrapCursor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cursor
}

func (c cursor) CursorType() CursorType {
	var _arg0 *C.GdkCursor    // out
	var _cret C.GdkCursorType // in

	_arg0 = (*C.GdkCursor)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_cursor_get_cursor_type(_arg0)

	var _cursorType CursorType // out

	_cursorType = CursorType(_cret)

	return _cursorType
}

func (c cursor) Display() Display {
	var _arg0 *C.GdkCursor  // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkCursor)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_cursor_get_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Display)

	return _display
}

func (c cursor) Image() gdkpixbuf.Pixbuf {
	var _arg0 *C.GdkCursor // out
	var _cret *C.GdkPixbuf // in

	_arg0 = (*C.GdkCursor)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_cursor_get_image(_arg0)

	var _pixbuf gdkpixbuf.Pixbuf // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(gdkpixbuf.Pixbuf)

	return _pixbuf
}

func (c cursor) Surface() (xHot float64, yHot float64, surface cairo.Surface) {
	var _arg0 *C.GdkCursor       // out
	var _arg1 *C.gdouble         // in
	var _arg2 *C.gdouble         // in
	var _cret *C.cairo_surface_t // in

	_arg0 = (*C.GdkCursor)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_cursor_get_surface(_arg0, &_arg1, &_arg2)

	var _xHot float64          // out
	var _yHot float64          // out
	var _surface cairo.Surface // out

	_xHot = float64(_arg1)
	_yHot = float64(_arg2)
	_surface = (cairo.Surface)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_surface, func(v cairo.Surface) {
		C.free(unsafe.Pointer(v))
	})

	return _xHot, _yHot, _surface
}

func (c cursor) RefCursor() Cursor {
	var _arg0 *C.GdkCursor // out
	var _cret *C.GdkCursor // in

	_arg0 = (*C.GdkCursor)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_cursor_ref(_arg0)

	var _ret Cursor // out

	_ret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Cursor)

	return _ret
}

func (c cursor) UnrefCursor() {
	var _arg0 *C.GdkCursor // out

	_arg0 = (*C.GdkCursor)(unsafe.Pointer(c.Native()))

	C.gdk_cursor_unref(_arg0)
}
