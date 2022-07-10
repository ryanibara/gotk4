// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeAnchorHints returns the GType for the type AnchorHints.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeAnchorHints() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "AnchorHints").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalAnchorHints)
	return gtype
}

// GTypePopupLayout returns the GType for the type PopupLayout.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypePopupLayout() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "PopupLayout").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalPopupLayout)
	return gtype
}

// AnchorHints: positioning hints for aligning a surface relative to a
// rectangle.
//
// These hints determine how the surface should be positioned in the case that
// the surface would fall off-screen if placed in its ideal position.
//
// For example, GDK_ANCHOR_FLIP_X will replace GDK_GRAVITY_NORTH_WEST with
// GDK_GRAVITY_NORTH_EAST and vice versa if the surface extends beyond the left
// or right edges of the monitor.
//
// If GDK_ANCHOR_SLIDE_X is set, the surface can be shifted horizontally to fit
// on-screen. If GDK_ANCHOR_RESIZE_X is set, the surface can be shrunken
// horizontally to fit.
//
// In general, when multiple flags are set, flipping should take precedence over
// sliding, which should take precedence over resizing.
type AnchorHints C.guint

const (
	// AnchorFlipX: allow flipping anchors horizontally.
	AnchorFlipX AnchorHints = 0b1
	// AnchorFlipY: allow flipping anchors vertically.
	AnchorFlipY AnchorHints = 0b10
	// AnchorSlideX: allow sliding surface horizontally.
	AnchorSlideX AnchorHints = 0b100
	// AnchorSlideY: allow sliding surface vertically.
	AnchorSlideY AnchorHints = 0b1000
	// AnchorResizeX: allow resizing surface horizontally.
	AnchorResizeX AnchorHints = 0b10000
	// AnchorResizeY: allow resizing surface vertically.
	AnchorResizeY AnchorHints = 0b100000
	// AnchorFlip: allow flipping anchors on both axes.
	AnchorFlip AnchorHints = 0b11
	// AnchorSlide: allow sliding surface on both axes.
	AnchorSlide AnchorHints = 0b1100
	// AnchorResize: allow resizing surface on both axes.
	AnchorResize AnchorHints = 0b110000
)

func marshalAnchorHints(p uintptr) (interface{}, error) {
	return AnchorHints(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for AnchorHints.
func (a AnchorHints) String() string {
	if a == 0 {
		return "AnchorHints(0)"
	}

	var builder strings.Builder
	builder.Grow(113)

	for a != 0 {
		next := a & (a - 1)
		bit := a - next

		switch bit {
		case AnchorFlipX:
			builder.WriteString("FlipX|")
		case AnchorFlipY:
			builder.WriteString("FlipY|")
		case AnchorSlideX:
			builder.WriteString("SlideX|")
		case AnchorSlideY:
			builder.WriteString("SlideY|")
		case AnchorResizeX:
			builder.WriteString("ResizeX|")
		case AnchorResizeY:
			builder.WriteString("ResizeY|")
		case AnchorFlip:
			builder.WriteString("Flip|")
		case AnchorSlide:
			builder.WriteString("Slide|")
		case AnchorResize:
			builder.WriteString("Resize|")
		default:
			builder.WriteString(fmt.Sprintf("AnchorHints(0b%b)|", bit))
		}

		a = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if a contains other.
func (a AnchorHints) Has(other AnchorHints) bool {
	return (a & other) == other
}

// PopupLayout: GdkPopupLayout struct contains information that is necessary
// position a gdk.Popup relative to its parent.
//
// The positioning requires a negotiation with the windowing system, since it
// depends on external constraints, such as the position of the parent surface,
// and the screen dimensions.
//
// The basic ingredients are a rectangle on the parent surface, and the anchor
// on both that rectangle and the popup. The anchors specify a side or corner to
// place next to each other.
//
// !Popup anchors (popup-anchors.png)
//
// For cases where placing the anchors next to each other would make the popup
// extend offscreen, the layout includes some hints for how to resolve this
// problem. The hints may suggest to flip the anchor position to the other side,
// or to 'slide' the popup along a side, or to resize it.
//
// !Flipping popups (popup-flip.png)
//
// !Sliding popups (popup-slide.png)
//
// These hints may be combined.
//
// Ultimatively, it is up to the windowing system to determine the position and
// size of the popup. You can learn about the result by calling
// gdk.Popup.GetPositionX(), gdk.Popup.GetPositionY(), gdk.Popup.GetRectAnchor()
// and gdk.Popup.GetSurfaceAnchor() after the popup has been presented. This can
// be used to adjust the rendering. For example, gtk.Popover changes its arrow
// position accordingly. But you have to be careful avoid changing the size of
// the popover, or it has to be presented again.
//
// An instance of this type is always passed by reference.
type PopupLayout struct {
	*popupLayout
}

// popupLayout is the struct that's finalized.
type popupLayout struct {
	native unsafe.Pointer
}

func marshalPopupLayout(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &PopupLayout{&popupLayout{(unsafe.Pointer)(b)}}, nil
}

// Copy makes a copy of layout.
//
// The function returns the following values:
//
//    - popupLayout: copy of layout.
//
func (layout *PopupLayout) Copy() *PopupLayout {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))

	_info := girepository.MustFind("Gdk", "PopupLayout")
	_gret := _info.InvokeRecordMethod("copy", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(layout)

	var _popupLayout *PopupLayout // out

	_popupLayout = (*PopupLayout)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_popupLayout)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _popupLayout
}

// Equal: check whether layout and other has identical layout properties.
//
// The function takes the following parameters:
//
//    - other GdkPopupLayout.
//
// The function returns the following values:
//
//    - ok: TRUE if layout and other have identical layout properties, otherwise
//      FALSE.
//
func (layout *PopupLayout) Equal(other *PopupLayout) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(other)))

	_info := girepository.MustFind("Gdk", "PopupLayout")
	_gret := _info.InvokeRecordMethod("equal", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(layout)
	runtime.KeepAlive(other)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// AnchorRect: get the anchor rectangle.
//
// The function returns the following values:
//
//    - rectangle: anchor rectangle.
//
func (layout *PopupLayout) AnchorRect() *Rectangle {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))

	_info := girepository.MustFind("Gdk", "PopupLayout")
	_gret := _info.InvokeRecordMethod("get_anchor_rect", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(layout)

	var _rectangle *Rectangle // out

	_rectangle = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _rectangle
}

// Offset retrieves the offset for the anchor rectangle.
//
// The function returns the following values:
//
//    - dx: return location for the delta X coordinate.
//    - dy: return location for the delta Y coordinate.
//
func (layout *PopupLayout) Offset() (dx int32, dy int32) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))

	_info := girepository.MustFind("Gdk", "PopupLayout")
	_info.InvokeRecordMethod("get_offset", _args[:], _outs[:])

	runtime.KeepAlive(layout)

	var _dx int32 // out
	var _dy int32 // out

	_dx = int32(*(*C.int)(unsafe.Pointer(&_outs[0])))
	_dy = int32(*(*C.int)(unsafe.Pointer(&_outs[1])))

	return _dx, _dy
}

// ShadowWidth obtains the shadow widths of this layout.
//
// The function returns the following values:
//
//    - left: return location for the left shadow width.
//    - right: return location for the right shadow width.
//    - top: return location for the top shadow width.
//    - bottom: return location for the bottom shadow width.
//
func (layout *PopupLayout) ShadowWidth() (left int32, right int32, top int32, bottom int32) {
	var _args [1]girepository.Argument
	var _outs [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))

	_info := girepository.MustFind("Gdk", "PopupLayout")
	_info.InvokeRecordMethod("get_shadow_width", _args[:], _outs[:])

	runtime.KeepAlive(layout)

	var _left int32   // out
	var _right int32  // out
	var _top int32    // out
	var _bottom int32 // out

	_left = int32(*(*C.int)(unsafe.Pointer(&_outs[0])))
	_right = int32(*(*C.int)(unsafe.Pointer(&_outs[1])))
	_top = int32(*(*C.int)(unsafe.Pointer(&_outs[2])))
	_bottom = int32(*(*C.int)(unsafe.Pointer(&_outs[3])))

	return _left, _right, _top, _bottom
}

// SetAnchorRect: set the anchor rectangle.
//
// The function takes the following parameters:
//
//    - anchorRect: new anchor rectangle.
//
func (layout *PopupLayout) SetAnchorRect(anchorRect *Rectangle) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(anchorRect)))

	_info := girepository.MustFind("Gdk", "PopupLayout")
	_info.InvokeRecordMethod("set_anchor_rect", _args[:], nil)

	runtime.KeepAlive(layout)
	runtime.KeepAlive(anchorRect)
}

// SetOffset: offset the position of the anchor rectangle with the given delta.
//
// The function takes the following parameters:
//
//    - dx: x delta to offset the anchor rectangle with.
//    - dy: y delta to offset the anchor rectangle with.
//
func (layout *PopupLayout) SetOffset(dx int32, dy int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(dx)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(dy)

	_info := girepository.MustFind("Gdk", "PopupLayout")
	_info.InvokeRecordMethod("set_offset", _args[:], nil)

	runtime.KeepAlive(layout)
	runtime.KeepAlive(dx)
	runtime.KeepAlive(dy)
}

// SetShadowWidth sets the shadow width of the popup.
//
// The shadow width corresponds to the part of the computed surface size that
// would consist of the shadow margin surrounding the window, would there be
// any.
//
// The function takes the following parameters:
//
//    - left: width of the left part of the shadow.
//    - right: width of the right part of the shadow.
//    - top: height of the top part of the shadow.
//    - bottom: height of the bottom part of the shadow.
//
func (layout *PopupLayout) SetShadowWidth(left int32, right int32, top int32, bottom int32) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(left)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(right)
	*(*C.int)(unsafe.Pointer(&_args[3])) = C.int(top)
	*(*C.int)(unsafe.Pointer(&_args[4])) = C.int(bottom)

	_info := girepository.MustFind("Gdk", "PopupLayout")
	_info.InvokeRecordMethod("set_shadow_width", _args[:], nil)

	runtime.KeepAlive(layout)
	runtime.KeepAlive(left)
	runtime.KeepAlive(right)
	runtime.KeepAlive(top)
	runtime.KeepAlive(bottom)
}
